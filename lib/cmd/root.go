package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"

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

  capacitor.Discharge()

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
