package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 0
    var v5 uint8 = v4
    v6 := GenS1{F0: -900, F1: -1}
    switch ("foo bar" + "foo bar") {
    case "":
        v1 = (v1 | v1)
    case "hello":
        v5 -= 162
    }
    if true {
        v7 := GenS1{F0: 354, F1: 822}
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint64 = 207
            fmt.Println(v9, v1)
            v9 = v9
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: (!true)}
        _ = v10
        _ = v7
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{}
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 uint16 = (v1 + (v1 ^ 32767))
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
        v4 = 2
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        switch ("hello" + "foo bar") {
        case "\n\t":
            if true {
                v16 := GenS2{F0: true}
                var v17 int16 = 42
                v18 := []int{(-(fi15)), fi15, fi15}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v19 := GenS3{F0: ""}
                var v20 int = fi15
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS4{F0: (v4 + v5)}
            _ = v21
        }
        var v22 uint8 = (^(v5))
        var v23 float64 = (0.0 / 3.14159)
        _ = fi15
        _ = v22
        _ = v23
    }
    v24 := []int{2147483647}
    _ = len(v24)
    v25 := GenS3{F0: "a"}
    switch 641 {
    case -243:
        v5 += (v5 ^ (v5 >> ((0) & 63)))
    case 9223372036854775807:
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 uint8 = (v4 + v4)
            v28 := GenS4{F0: 1}
            _ = fi26
            _ = v27
            _ = v28
        }
    case 2147483647:
        v4 = (v4 | 42)
    }
    v29 := GenS4{F0: 255}
    v30 := GenS1{F0: 0, F1: -695}
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(v4)
        _ = fi31
    }
    v32 := GenS5{F0: 0, F1: "foo bar"}
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
