package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint }
func genFn0() int {
    var v5 string = ""
    for fi6 := 0; fi6 < 8; fi6++ {
        switch ("" + "") {
        case "a":
            for fi7 := 0; fi7 < 9; fi7++ {
                v5 = "hello"
                _ = fi7
            }
        case "foo bar":
            var v8 uint16 = 497
            _ = v8
        case "\n\t":
            var v9 uint8 = 1
            _ = v9
        }
        v10 := GenS1{F0: 430, F1: 2, F2: -2147483648}
        if ((!false) && (32767 == -975)) {
            v11 := []int{fi6, (fi6 + (^(fi6))), (fi6 * (fi6 % (((fi6 & (fi6 / (((-(fi6))) | 1)))) | 1))), (^(fi6))}
            fmt.Println(fi6, fi6)
            _ = v11
        } else {
            switch ("foo bar" + ((v5 + v5) + v5)) {
            case "a":
                v12 := []int{(fi6 | fi6), -617, 0, 128}
                _ = v12
            default:
                var v13 int64 = -872
                _ = v13
            }
            switch fi6 {
            case 42:
                fmt.Println(fi6, v5)
            case 1:
                v5 = "hello"
            case 840:
                v5 = v5
            }
        }
        _ = fi6
        _ = v10
    }
    v14 := []int{1, 321}
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v5, fi15)
        _ = fi15
    }
    _ = v14
    _ = v5
    return 122
}

func main() {
    v1 := []int{}
    var v2 bool = (2147483647 <= -521)
    var v3 uint8 = 2
    var v4 uint32 = 2147483647
    v4 = 599
    v4 -= (v4 & 622)
    v2 = v2
    v2 = (v2 || v2)
    var v16 uint8 = (v3 & (v3 % ((51) | 1)))
    var v17 uint16 = 568
    if (898 < -9223372036854775808) {
        var v18 int64 = 1
        v19 := GenS1{F0: 883, F1: 32767, F2: 0}
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint64 = 4294967295
            fmt.Println(v21)
            _ = fi20
            _ = v21
        }
        _ = v18
        _ = v19
    }
    var v22 float64 = (0.0 * (1.5 + 0.0))
    var v23 uint64 = 4294967295
    fmt.Println(v16)
    v3 = 42
    if ("hello" >= "\n\t") {
        switch "a" {
        case "hello":
            var v24 string = "hello"
            _ = v24
        }
        v22 = -2.25
        var v25 uint16 = (v17 ^ (v17 << ((0) & 63)))
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS2{F0: 1, F1: -128, F2: uint(v4)}
        if false {
            v23 = v23
            var v28 int8 = -104
            v2 = (4294967295 <= uint64(v17))
            _ = v28
        }
        v22 = -2.25
        _ = fi26
        _ = v27
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 string = ((("foo bar" + ("hello" + ("a" + (("a" + (("a" + "\n\t") + "hello")) + "hello")))) + "a") + "\n\t")
        _ = fi29
        _ = v30
    }
    fmt.Println(v4, v22, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
