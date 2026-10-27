package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
func genFn0() uint8 {
    var v6 int = 0
    var v7 float64 = 0.1
    v8 := GenS1{F0: "\n\t"}
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    var v3 uint8 = 145
    var v4 float64 = float64(-2.25)
    v5 := []string{"a", ("a" + ""), ""}
    switch "hello" {
    case "foo bar":
        v9 := GenS1{F0: (("a" + "") + "hello")}
        _ = v9
    }
    switch "a" {
    case "foo bar":
        var v10 uint64 = 1
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: "hello"}
        _ = v11
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if true {
        switch -1 {
        case 2147483647:
            var v12 uint = uint(v1)
            _ = v12
        case 0:
            fmt.Println(v3, v1)
        default:
            _ = len(v2)
        }
        var v13 uint64 = 1
        fmt.Println(v13, v1)
        _ = v13
    } else {
        v14 := GenS1{F0: "foo bar"}
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int = (fi15 - fi15)
            var v17 float64 = float64(3.14159)
            v16 -= (fi15 - fi15)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    var v18 uint8 = v3
    v19 := GenS2{F0: 4294967295, F1: 2}
    switch "\n\t" {
    case "":
        _ = len(v2)
    case "hello":
        var v20 uint32 = 132
        _ = v20
    case "foo bar":
        v4 = v4
    }
    var v21 uint16 = (v1 - (^(v1)))
    var v22 uint = 2
    v22 -= v22
    _ = len(v2)
    v23 := GenS2{F0: -9223372036854775808, F1: uint64(v22)}
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS1{F0: ""}
        v21 = v21
        v26 := GenS3{F0: 42, F1: (v3 - (^(v3))), F2: (v22 << ((v22) & 63))}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
