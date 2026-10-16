package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int8; F1 uint; F2 int64 }

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    if ((1e10 - 3.14159) != 1e10) {
        var v4 int8 = 73
        switch "" {
        case "foo bar":
            var v5 int32 = 0
            _ = v5
        }
        _ = v4
    } else {
        var v6 uint64 = 1
        _ = len(v2)
        v7 := GenS1{F0: int8(v6), F1: (false || false)}
        _ = v6
        _ = v7
    }
    switch -1 {
    case 890:
        var v8 uint16 = 322
        _ = v8
    case 32767:
        var v9 uint8 = 24
        _ = v9
    default:
        v10 := GenS1{F0: 127, F1: (0 > 0)}
        _ = v10
    }
    v11 := GenS1{F0: 34, F1: ((!(false || (!(2 <= 60)))) || false)}
    v12 := GenS2{F0: 42, F1: 4294967295, F2: 432}
    var v13 int32 = 0
    v13 = (^(v13))
    switch (("\n\t" + "a") + "a") {
    case "a":
        var v14 int8 = -1
        _ = v14
    case "hello":
        switch "\n\t" {
        case "foo bar":
            var v15 uint = 4294967295
            _ = v15
        case "hello":
            v16 := GenS1{F0: -109, F1: (2147483647 != -1)}
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: 42, F1: (!true)}
                fmt.Println(v13, v13)
                v3["hello"] = (fi17 % (((fi17 >> ((2) & 63))) | 1))
                _ = fi17
                _ = v18
            }
        }
    default:
        v13 = (v13 ^ v13)
    }
    var v19 int = 548
    var v20 int64 = -191
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
