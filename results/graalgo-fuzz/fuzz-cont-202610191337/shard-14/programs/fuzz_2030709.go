package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{}
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 12; fi4++ {
        v1 -= 2
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v5 int32 = -895
        _ = fi4
        _ = v5
    }
    var v6 int16 = 1
    v2 = append(v2, "")
    v7 := GenS1{F0: (!true), F1: (((^(v1)) > v1) || (1 != 18446744073709551615))}
    v3[""] = -1
    var v8 uint8 = 255
    if true {
        switch "\n\t" {
        case "\n\t":
            switch ("hello" + ("\n\t" + ("\n\t" + "foo bar"))) {
            case "\n\t":
                var v9 int32 = 42
                _ = v9
            }
        case "a":
            v10 := GenS1{F0: (32767 > 127), F1: (!false)}
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: true, F1: false}
                _ = fi11
                _ = v12
            }
        default:
            var v13 float64 = 1.5
            _ = v13
        }
        var v14 float32 = (0.1 * 0.0)
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint64 = (v1 + (v1 % ((v1) | 1)))
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        switch ("\n\t" + ("foo bar" + "a")) {
        case "hello":
            var v17 int8 = -78
            _ = v17
        case "":
            var v18 bool = false
            _ = v18
        case "a":
            v1 = 4294967295
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v8, v6, v8)
        case "hello":
            var v19 uint16 = 638
            _ = v19
        default:
            v6 = 37
        }
    }
    _ = len(v3)
    for fi20 := 0; fi20 < 17; fi20++ {
        v8 = v8
        _ = fi20
    }
    v6 += (v6 >> ((0) & 63))
    v1 = v1
    var v21 uint8 = 0
    if true {
        v8 += v21
        if len(v2) > 0 {
            v2[0] = ""
        }
        var v22 int = 4294967295
        _ = v22
    } else {
        fmt.Println(v1, v6)
    }
    fmt.Println(v21, v6, v8)
    var v23 uint16 = 493
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
