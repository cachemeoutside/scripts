package main

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

type T struct {
  VPN []string `yaml:"ucla_vpn_networks",omitempty"`
  LIB []string `yaml:"ucla_library_staff_networks",omitempty"`
}

func main() {
  t := T{}

  yamlFile, err := ioutil.ReadFile("/Users/avuong/ucla-workspace/github/ansible_uclalib_configs/group_vars/all/main.yml")
  err = yaml.Unmarshal(yamlFile, &t)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(t.VPN)
  fmt.Println(t.LIB)
}
