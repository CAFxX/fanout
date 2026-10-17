package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 int8, p1 int) int32 {
    v4 := []int{(^(p1)), (p1 & (p1 * p1))}
    switch "foo bar" {
    case "hello":
        v5 := GenS1{F0: 2147483647, F1: true, F2: -1}
        _ = v5
    case "":
        var v6 uint8 = 128
        _ = v6
    case "\n\t":
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := []string{("" + ("a" + ""))}
    var v2 int64 = 433
    var v3 int32 = -630
    var v7 uint8 = 42
    v8 := GenS1{F0: 745, F1: false, F2: 9223372036854775807}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v7, v7, v3)
    v2 -= v2
    for fi9 := 0; fi9 < 11; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = (v3 * genFn0(-1, (fi9 << ((18446744073709551615) & 63))))
        _ = fi9
    }
    if true {
        var v10 uint32 = 0
        v2 = 9223372036854775807
        v3 = (v3 >> ((2) & 63))
        _ = v10
    } else {
        v1 = append(v1, ("hello" + (((("a" + "") + "hello") + "") + ("hello" + "\n\t"))))
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint = 0
            _ = fi11
            _ = v12
        }
    }
    _ = len(v1)
    v13 := GenS1{F0: (v3 | (v3 / (((v3 % ((224) | 1))) | 1))), F1: (42 >= -9223372036854775808), F2: (v2 | v2)}
    if ((false || false) && (0 < -1)) {
        fmt.Println(v7)
    } else {
        v14 := GenS2{F0: 974, F1: 42}
        _ = v14
    }
    var v15 uint = 263
    v3 = 2147483647
    var v16 int8 = -5
    if true {
        switch "foo bar" {
        case "hello":
            _ = len(v1)
        default:
            _ = len(v1)
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint16 = 341
            for fi19 := 0; fi19 < 17; fi19++ {
                fmt.Println(v18)
                v20 := GenS1{F0: genFn0((v16 * (^(v16))), fi19), F1: true, F2: 9223372036854775807}
                _ = fi19
                _ = v20
            }
            fmt.Println(v18, v3)
            _ = fi17
            _ = v18
        }
    } else {
        var v21 uint = (v15 * v15)
        _ = v21
    }
    v2 = v2
    v16 += -128
    var v22 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v7
    _ = v8
}
