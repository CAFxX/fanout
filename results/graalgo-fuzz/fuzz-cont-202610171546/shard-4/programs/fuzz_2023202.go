package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -84
    var v3 uint8 = 127
    var v4 uint8 = (v3 % ((0) | 1))
    v4 -= (^(v4))
    v3 = v3
    var v5 uint16 = 680
    v6 := GenS1{F0: 635}
    var v7 string = ""
    for fi8 := 0; fi8 < 11; fi8++ {
        if false {
            switch v7 {
            case "":
                v9 := GenS1{F0: 275}
                _ = v9
            default:
                v10 := GenS1{F0: 4294967295}
                _ = v10
            }
            var v11 float32 = -2.25
            _ = v11
        }
        _ = fi8
    }
    v4 = v3
    var v12 uint64 = 933
    switch -45 {
    case 4294967295:
        var v13 int16 = 723
        _ = v13
    default:
        switch 128 {
        case 307:
            v14 := []int{}
            _ = v14
        }
    }
    fmt.Println(v5, v2)
    v7 = v7
    switch (v7 + v7) {
    case "":
        var v15 int = 4294967295
        _ = v15
    case "a":
        v16 := []int{-9223372036854775808}
        _ = v16
    }
    v12 -= 1
    fmt.Println(v5, v2, v7)
    var v17 int64 = 0
    if (true && false) {
        v7 = "hello"
        var v18 uint32 = 751
        _ = v18
    }
    if (false && ((((1e10 * 0.1) > (1.5 + float64(float32(float64(float32(1e10)))))) || (42 != 18446744073709551615)) || (("a" + ("hello" + (("hello" + (v7 + v7)) + ("\n\t" + v7)))) <= "hello"))) {
        fmt.Println(v4, v3)
        v17 -= v17
    } else {
        fmt.Println(v7, v17)
        var v19 float32 = 0.0
        v1[1] = 42
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v5 = (v5 * 273)
        if (true && true) {
            switch int(v17) {
            case 2147483647:
                v12 = v12
            default:
                fmt.Println(v3, v3, v3)
            }
            switch fi20 {
            case -9223372036854775808:
                if (!true) {
                    var v21 float64 = (3.14159 / 1.0)
                    _ = v21
                } else {
                    v2 = v2
                    v22 := []int{(^(fi20))}
                    v7 = ""
                    _ = v22
                }
            case 9223372036854775807:
                v12 += (v12 | v12)
            }
            v23 := GenS2{F0: false, F1: 0}
            _ = v23
        }
        v2 -= v2
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
