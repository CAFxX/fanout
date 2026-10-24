package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int32 = 484
    var v2 float64 = 1.5
    v3 := make(map[string]string)
    if false {
        var v4 bool = false
        switch 42 {
        case 32768:
            v5 := []int{-9223372036854775808, 0}
            _ = v5
        case 4294967295:
            v3[("a" + "\n\t")] = ("hello" + "a")
        }
        _ = v4
    }
    var v6 uint64 = 1
    switch -1 {
    case 127:
        var v7 int = -253
        _ = v7
    case 128:
        v2 = 1.5
    case 685:
        if true {
            v8 := GenS1{F0: 1, F1: -9223372036854775808, F2: (4294967295 > 18446744073709551615)}
            _ = v8
        } else {
            _ = len(v3)
        }
    default:
        var v9 int16 = 1
        _ = v9
    }
    var v10 string = ""
    if true {
        fmt.Println(v6, v6)
        var v11 uint = 0
        switch v10 {
        case "foo bar":
            var v12 float32 = 1.5
            _ = v12
        case "\n\t":
            v11 = v11
        case "hello":
            v13 := []int{-343, 9223372036854775807, 334, -2147483648}
            _ = v13
        default:
            if false {
                v10 = "\n\t"
                switch (("\n\t" + "a") + "foo bar") {
                case "\n\t":
                    var v14 uint32 = 42
                    _ = v14
                default:
                    v6 = 614
                }
                fmt.Println(v11)
            }
        }
        _ = v11
    } else {
        _ = len(v3)
        v3[(v10 + "a")] = ("\n\t" + v10)
    }
    v10 = v10
    v15 := GenS1{F0: uint8(v6), F1: 128, F2: ((!false) && true)}
    v10 = (v10 + (((v10 + "") + v10) + v10))
    v1 -= (v1 | v1)
    var v16 int16 = -32768
    v17 := GenS2{F0: 2147483648}
    v18 := []int{81, 4294967295}
    v1 -= -2147483648
    var v19 int64 = -9223372036854775808
    switch (v10 + (v10 + v10)) {
    case "hello":
        fmt.Println(v2, v6, v6)
    case "\n\t":
        var v20 uint16 = 42
        _ = v20
    case "":
        v10 = ("\n\t" + v10)
    default:
        var v21 uint32 = 0
        _ = v21
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
