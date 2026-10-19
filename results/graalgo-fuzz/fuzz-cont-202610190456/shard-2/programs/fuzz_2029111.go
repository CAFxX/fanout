package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 bool = false
    v2 := []string{("" + (("hello" + "hello") + "foo bar")), "foo bar", "hello", ("a" + "")}
    var v3 float32 = (1e10 + ((-2.25 - 1e10) / 1.5))
    switch (("" + ("\n\t" + "hello")) + "foo bar") {
    case "hello":
        v4 := GenS1{F0: 42, F1: 0}
        _ = v4
    case "a":
        v5 := GenS1{F0: 54, F1: 2}
        _ = v5
    case "foo bar":
        var v6 int8 = 67
        _ = v6
    }
    v7 := GenS1{F0: 127, F1: 42}
    if (v3 <= v3) {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int8 = 127
            v2 = append(v2, "hello")
            var v10 float64 = 0.0
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        _ = len(v2)
        var v11 string = "hello"
        _ = v11
    }
    v1 = (42 <= 100)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v12 uint16 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v12 = v12
        _ = fi13
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
