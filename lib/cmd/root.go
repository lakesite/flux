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

var cmdAddOrganization = &cobra.Command{
  Use:   "addorganization []",
  Short: "Adds an organization record",
  Long: `Given a name and description, add an organization record.`,
  Run: func(cmd *cobra.Command, args []string) {
      name, _:= cmd.Flags().GetString("name")
      description, _ := cmd.Flags().GetString("description")

      organization := capacitor.Organization{name, description}
      resp, err := capacitor.SetRecord("organizations", organization)
      if err == nil {
        fmt.Println("Result: " + resp)
      }
  },
}

var cmdGetOrganization = &cobra.Command{
  Use:   "getorganization [id]",
  Short: "Gets an organization record by id",
  Long: `Given an id, return the organization record.`,
  Run: func(cmd *cobra.Command, args []string) {
      var organization capacitor.Organization

      id, _:= cmd.Flags().GetInt("id")
      resp, entity, err := capacitor.GetRecord("organizations", id)
      if err == nil {
        // now bind the struct
        bindErr := mapstructure.Decode(entity, &organization)
        if (bindErr != nil) {
          panic(bindErr)
        }

        fmt.Printf("JSON response: %v\n", resp)
        fmt.Printf("Organization entity: %#v\n", organization)
      }
  },
}

var cmdAddSystem = &cobra.Command{
  Use:   "addsystem []",
  Short: "Adds a system record",
  Long: `Given a name, organization id, ip address, and description, add a
         system record.`,
  Run: func(cmd *cobra.Command, args []string) {
      name, _:= cmd.Flags().GetString("name")
      organizationId, _:= cmd.Flags().GetString("id")
      ipAddress, _ := cmd.Flags().GetString("ip")
      description, _ := cmd.Flags().GetString("description")

      system := capacitor.System{name, organizationId, ipAddress, description}
      resp, err := capacitor.SetRecord("systems", system)
      if err == nil {
        fmt.Println("Result: " + resp)
      }
  },
}

var cmdGetSystem = &cobra.Command{
  Use:   "getsystem [id]",
  Short: "Gets a system event record by id",
  Long: `Given an id, return the system record.`,
  Run: func(cmd *cobra.Command, args []string) {
      var system capacitor.System

      id, _:= cmd.Flags().GetInt("id")
      resp, entity, err := capacitor.GetRecord("systems", id)
      if err == nil {
        // now bind the struct
        bindErr := mapstructure.Decode(entity, &system)
        if (bindErr != nil) {
          panic(bindErr)
        }

        fmt.Printf("JSON response: %v\n", resp)
        fmt.Printf("System entity: %#v\n", system)
      }
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

  cmdAddSystem.Flags().StringP("name", "", "", "The name of the system.")
  cmdAddSystem.Flags().StringP("id", "", "", "The corresponding organization URL of the system record.")
  cmdAddSystem.Flags().StringP("ip", "", "", "The ip address of the system record.")
  cmdAddSystem.Flags().StringP("description", "", "", "The corresponding description of the system record.")
  cmdAddSystem.MarkFlagRequired("name")
  cmdAddSystem.MarkFlagRequired("id")
  cmdAddSystem.MarkFlagRequired("ip")
  cmdAddSystem.MarkFlagRequired("description")
  rootCmd.AddCommand(cmdAddSystem)

  cmdGetSystem.Flags().IntP("id", "", 0, "The id of the system record.")
  cmdGetSystem.MarkFlagRequired("id")
  rootCmd.AddCommand(cmdGetSystem)

  cmdAddOrganization.Flags().StringP("name", "", "", "The name of the organization.")
  cmdAddOrganization.Flags().StringP("description", "", "", "The description of the organization.")
  cmdAddOrganization.MarkFlagRequired("name")
  cmdAddOrganization.MarkFlagRequired("description")
  rootCmd.AddCommand(cmdAddOrganization)

  cmdGetOrganization.Flags().IntP("id", "", 0, "The id of the organization record.")
  cmdGetOrganization.MarkFlagRequired("id")
  rootCmd.AddCommand(cmdGetOrganization)

  capacitor.Discharge()

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
