package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }

func main() {
    v1 := []string{"foo bar", "a", ("foo bar" + (("foo bar" + "\n\t") + "a"))}
    var v2 uint8 = 0
    var v3 uint32 = 4294967295
    v4 := GenS1{F0: 42, F1: "a", F2: 65535}
    var v5 int16 = -32768
    switch "hello" {
    case "a":
        var v6 string = "\n\t"
        _ = v6
    case "hello":
        v7 := GenS2{F0: "", F1: 1, F2: 65535}
        _ = v7
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    _ = len(v1)
    v8 := GenS1{F0: 0, F1: ("a" + "foo bar"), F2: 4294967295}
    var v9 int16 = v5
    var v10 uint16 = 32767
    v3 += 60
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
