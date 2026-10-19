package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    v1 := []string{"\n\t"}
    v2 := []string{"", "\n\t", "foo bar", ("" + ("hello" + "hello"))}
    var v3 bool = false
    var v4 int64 = 9223372036854775807
    var v5 uint = 42
    v3 = true
    for fi6 := 0; fi6 < 3; fi6++ {
        v2 = append(v2, "foo bar")
        if false {
            v7 := GenS1{F0: "a", F1: 0}
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v3, v3, v5)
    v8 := GenS1{F0: ("hello" + "foo bar"), F1: 37}
    v9 := GenS2{F0: -96, F1: 119}
    var v10 uint8 = 143
    v4 = v4
    var v11 uint16 = 42
    var v12 float32 = (1e10 - 0.0)
    v2 = append(v2, "\n\t")
    v1 = append(v1, ("foo bar" + (("" + "a") + "")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
