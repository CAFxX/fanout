package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }

func main() {
    var v1 int8 = 108
    v2 := []string{"a", "foo bar"}
    var v3 string = "foo bar"
    var v4 int32 = 318
    var v5 uint64 = 4294967295
    v5 = v5
    switch "" {
    case "foo bar":
        switch v3 {
        case "foo bar":
            v3 = v3
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "a":
            var v6 float32 = -2.25
            _ = v6
        }
    case "":
        var v7 int8 = 45
        _ = v7
    }
    if false {
        fmt.Println(v4, v4)
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int64 = 42
            var v10 float64 = (0.1 / 3.14159)
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v2 = append(v2, "foo bar")
        if true {
            if len(v2) > 0 {
                v2[0] = v3
            }
        } else {
            var v11 float64 = (1.5 + float64((1e10 + float32(0.1))))
            fmt.Println(v11, v4)
            var v12 string = ("a" + "hello")
            _ = v11
            _ = v12
        }
    }
    v5 = v5
    fmt.Println(v4, v5)
    var v13 int = -709
    switch ((v3 + (v3 + "hello")) + (v3 + "foo bar")) {
    case "\n\t":
        v3 = v3
    case "foo bar":
        _ = len(v2)
    case "hello":
        fmt.Println(v4)
    default:
        if true {
            _ = len(v2)
        } else {
            v13 = (^(v13))
        }
    }
    if (0 != 849) {
        var v14 bool = (true && (!false))
        _ = v14
    } else {
        v5 -= 42
    }
    fmt.Println(v4)
    var v15 string = "foo bar"
    v16 := GenS1{F0: 2, F1: 42, F2: (v5 - 42)}
    v13 -= (v13 | v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
