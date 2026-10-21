package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    var v1 float64 = (1e10 * (3.14159 * 0.1))
    var v2 int8 = 42
    var v3 bool = true
    v4 := []int{42, -346}
    if ("foo bar" > "foo bar") {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 string = "foo bar"
            _ = fi5
            _ = v6
        }
    } else {
        switch "\n\t" {
        case "a":
            var v7 string = (((("hello" + ("" + "a")) + "hello") + "\n\t") + ("\n\t" + ""))
            _ = v7
        case "foo bar":
            v1 += v1
        }
        var v8 int64 = -1
        _ = v8
    }
    var v9 uint32 = 0
    for fi10 := 0; fi10 < 1; fi10++ {
        if false {
            var v11 int64 = 930
            v1 = float64(3.14159)
            fmt.Println(fi10, v2)
            _ = v11
        } else {
            var v12 int8 = v2
            v2 += v12
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(v1, v1, v1)
    v9 = 2147483647
    fmt.Println(v9, v1)
    v13 := GenS1{F0: (v3 && v3), F1: v2}
    switch 2147483647 {
    case 42:
        v14 := GenS1{F0: (9223372036854775807 != 0), F1: v2}
        _ = v14
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -2147483648:
        if v3 {
            _ = len(v4)
            var v15 uint16 = 42
            var v16 uint8 = 243
            _ = v15
            _ = v16
        }
    default:
        _ = len(v4)
    }
    if (((v2 / ((v2) | 1)) >= (^(v2))) || false) {
        var v17 float64 = v1
        _ = v17
    }
    v4 = append(v4, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
