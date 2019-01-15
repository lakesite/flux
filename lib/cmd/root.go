package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/mitchellh/mapstructure"

  "github.com/lakesite/flux/lib/capacitor"

)

var rootCmd = &cobra.Command{
  Use:   "flux",
  Short: "Flux consumes aeon api",
  Long: `Flux is a command line interface to the aeon api`,
  Run: func(cmd *cobra.Command, args []string) {
    capacitor.Discharge()
  },
}

var cmdAddStatus = &cobra.Command{
  Use:   "addstatus []",
  Short: "Adds a status event record",
  Long: `Given a system id, state and description (with other options),
  add a status event record.`,
  Run: func(cmd *cobra.Command, args []string) {
      system, _:= cmd.Flags().GetString("system")
      application, _:= cmd.Flags().GetString("application")
      service, _ := cmd.Flags().GetString("service")
      state, _ := cmd.Flags().GetString("state")
      description, _ := cmd.Flags().GetString("description")
      status := capacitor.Status{system, application, service, state, description}
      resp, err := capacitor.SetRecord("statuses", status)
      if err == nil {
        fmt.Println("Result: " + resp)
      }
  },
}

var cmdGetStatus = &cobra.Command{
  Use:   "getstatus [id]",
  Short: "Gets a status event record by id",
  Long: `Given an id, return the status event record.`,
  Run: func(cmd *cobra.Command, args []string) {
      var status capacitor.Status

      id, _:= cmd.Flags().GetInt("id")
      resp, entity, err := capacitor.GetRecord("statuses", id)
      if err == nil {
        // now bind the struct
        bindErr := mapstructure.Decode(entity, &status)
        if (bindErr != nil) {
          panic(bindErr)
        }

        fmt.Printf("JSON response: %v\n", resp)
        fmt.Printf("Status entity: %#v\n", status)
        fmt.Printf("")
      }
  },
}

func Execute() {
  // setup status record interface
  cmdAddStatus.Flags().StringP("system", "", "", "The corresponding system URL of the event record.")
  cmdAddStatus.Flags().StringP("application", "", "", "The corresponding application URL of the event record.")
  cmdAddStatus.Flags().StringP("service", "", "", "The corresponding service URL of the event record.")
  cmdAddStatus.Flags().StringP("state", "", "OKAY", "The corresponding state of the event record.")
  cmdAddStatus.Flags().StringP("description", "", "", "The corresponding description of the event record.")
  cmdAddStatus.MarkFlagRequired("system")
  cmdAddStatus.MarkFlagRequired("state")
  cmdAddStatus.MarkFlagRequired("description")
  rootCmd.AddCommand(cmdAddStatus)

  cmdGetStatus.Flags().IntP("id", "", 0, "The id of the event record.")
  cmdGetStatus.MarkFlagRequired("id")
  rootCmd.AddCommand(cmdGetStatus)

  capacitor.Discharge()

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
