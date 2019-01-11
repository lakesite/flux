package capacitor


// data models

type Status struct {
  System string
  Application string
  Service string
  State string
  Description string
}

type System struct {
  Name string
  Organization string
  Ip_address string
  Description string
}

type Organization struct {
  Name string
  Description string
}
