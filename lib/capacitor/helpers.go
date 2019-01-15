package capacitor


type AbstractContainer struct {
  Package map[string]string
}

/*
    users := gjson.Get(getJSONCollection("users"), "results")
    users.ForEach(func(key, value gjson.Result) bool {
      record := value.String()
      fmt.Println(gjson.Parse(record).Get("email"))
      return true
    })
*/
