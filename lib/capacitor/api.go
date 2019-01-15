package capacitor

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "reflect"
    "strconv"
    "strings"

    "github.com/tidwall/gjson"
)

// Given a username and password, post to the api-token-auth
// endpoint to retrieve (or create) a token using DRF.
func getToken(username, password string) string {
  formdata := url.Values{}
  formdata.Set("username", username)
  formdata.Set("password", password)

  req, err := http.NewRequest("POST", BASE_API + "api-token-auth/", strings.NewReader(formdata.Encode()))
  req.Close = true
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  response, err := httpClient.Do(req)
  if err != nil {
      log.Fatalf("The HTTP request failed with error %s\n", err)
  } else {
      defer response.Body.Close()
      data, _ := ioutil.ReadAll(response.Body)
      token = gjson.Get(string(data), "token").String()
  }
  return token
}

// Given an api endpoint and id for an entity, get the result
// and bind to a struct which represents the entity.
// Return the response body, struct and error if any
func GetRecord(endpoint string, id int) (string, interface{}, error) {
  result := ""
  entity := AbstractContainer{make(map[string]string)}

  req, err := http.NewRequest("GET", API + endpoint + "/" + strconv.Itoa(id), nil)
  req.Close = true
  req.Header.Add("Authorization", token)
  response, err := httpClient.Do(req)

  if err != nil {
    log.Fatalf("The HTTP request failed with error %s\n", err)
  } else {
    defer response.Body.Close()
    data, _ := ioutil.ReadAll(response.Body)
    result = string(data)
    gjson.Parse(result).ForEach(func(key, value gjson.Result) bool {
      if (value.String() != "") {
        entity.Package[key.String()] = value.String()
        fmt.Printf("Got key: %v, value: %v\n", key, value)
      }
      return true
    })

  }
  return result, entity.Package, err
}

// Given a model which maps directly into the api,
// meaning the name of the structure attributes map to the api
// post them and return the response body and error if any
func SetRecord(entity string, m interface{}) (string, error) {
  result := ""
  formdata := url.Values{}

  s := reflect.ValueOf(m)

  if s.Kind() == reflect.Ptr {
    s = s.Elem()
  }

  if s.Kind() != reflect.Struct {
    log.Fatalf("Unexpected type: %s", s.Kind())
  }

  for i := 0; i < s.NumField(); i++ {
    valueField := s.Field(i)
    typeField := s.Type().Field(i)
    fieldName := strings.ToLower(typeField.Name)
    fieldValue := reflect.ValueOf(valueField.Interface()).String()

    if len(fieldValue) > 0 {
      formdata.Set(fieldName, fieldValue)
    }
  }

  req, err := http.NewRequest("POST", API + entity + "/", strings.NewReader(formdata.Encode()))
  req.Close = true
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Authorization", token)
  response, err := httpClient.Do(req)

  if err != nil {
    log.Fatalf("The HTTP request failed with error %s\n", err)
  } else {
    defer response.Body.Close()
    data, _ := ioutil.ReadAll(response.Body)
    result = string(data)
  }
  return result, err
}

// Given an api endpoint (name) in the format DRF uses,
// return the json result subset
func getJSONCollection(name string) string {
  jsonResult := ""
  req, err := http.NewRequest("GET", API + name + "/", nil)
  req.Header.Add("Authorization", token)
  response, err := httpClient.Do(req)
  if err != nil {
      fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      jsonResult = string(data)
  }
  return jsonResult
}
