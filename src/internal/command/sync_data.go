package command

import (
	"fmt"

	"feed/internal/broker"

	"github.com/spf13/cobra"
)

// TODO подумать над организацией кода
var syncCmd = &cobra.Command{
	Use:   "sync-data",
	Short: "data synchronization",
	Long:  "data synchronization for feed generation",
	Run: func(cmd *cobra.Command, args []string) {
		if err := syncBase(); err != nil {
			panic(fmt.Sprintf("base synchronization error: %v", err))
		}

		if err := syncCategory(); err != nil {
			panic(fmt.Sprintf("category synchronization error: %v", err))
		}
	},
}

func syncBase() error {
	// TODO убрать повторяющийся код
	ch, q, err := broker.QueueSyncData("sync_base")
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var process chan struct{}
	go func() {
		for d := range msgs {
			fmt.Println(d.Body)
		}
	}()
	<-process

	return nil
}

func syncCategory() error {
	ch, q, err := broker.QueueSyncData("sync_category")
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var process chan struct{}
	go func() {
		for d := range msgs {
			fmt.Println(d.Body)
		}
	}()
	<-process

	return nil
}
