package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }

func main() {
    var v1 bool = (710 <= 42)
    v2 := []string{"\n\t", (("hello" + ((("" + "hello") + "hello") + "\n\t")) + ("" + ((("\n\t" + "foo bar") + "") + "foo bar"))), "a"}
    var v3 uint64 = 698
    switch "hello" {
    case "a":
        v3 = (v3 & v3)
    }
    v4 := GenS1{F0: 127, F1: 245}
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v5 uint64 = (v3 - (v3 % (((v3 ^ (^(v3)))) | 1)))
    v3 = v5
    var v6 uint8 = 2
    var v7 uint8 = (v6 * 2)
    v8 := GenS1{F0: 2147483647, F1: -32768}
    v9 := GenS2{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
