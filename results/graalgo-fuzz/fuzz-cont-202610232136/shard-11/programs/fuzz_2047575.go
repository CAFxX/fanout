package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 uint32, p1 uint64) string {
    v7 := []int{-884, 127, -2147483648}
    var v8 int32 = -1
    var v9 bool = true
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v8, p1)
    default:
        var v10 uint64 = (^(p1))
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 string = ("\n\t" + ("foo bar" + ("hello" + ("" + (("" + ("a" + ("a" + "a"))) + ("" + "foo bar"))))))
    var v2 float64 = ((float64(3.14159) + float64((1.5 - 3.14159))) - 0.0)
    var v3 int64 = -908
    var v4 int8 = 1
    var v5 int16 = 52
    var v6 bool = (((false && true) && (v4 == 127)) && true)
    var v11 float64 = v2
    v12 := GenS1{F0: 2147483647, F1: 42}
    v6 = (4294967295 > 408)
    switch 919 {
    case 127:
        fmt.Println(v6, v1)
    }
    var v13 float32 = (1e10 - 3.14159)
    var v14 string = v1
    if (-9223372036854775808 <= 32768) {
        if v6 {
            var v15 uint8 = 42
            switch 32768 {
            case 4294967295:
                v16 := []int{32768, 9223372036854775807, 127}
                _ = v16
            case 32767:
                v4 -= (v4 + v4)
            }
            _ = v15
        }
        fmt.Println(v14, v14, v11)
    }
    v17 := []int{}
    for fi18 := 0; fi18 < 13; fi18++ {
        switch "\n\t" {
        case "hello":
            if ((1 > 42) && (false && v6)) {
                v13 = ((1.5 + 1e10) - 1e10)
            }
        case "foo bar":
            fmt.Println(v2)
        }
        switch fi18 {
        case -2147483648:
            var v19 int64 = 9223372036854775807
            _ = v19
        case -9223372036854775808:
            var v20 float64 = 0.1
            _ = v20
        case -1:
            var v21 int = 2147483647
            _ = v21
        }
        if (v3 <= (v3 + (v3 << ((18446744073709551615) & 63)))) {
            fmt.Println(v11)
            var v22 float32 = -2.25
            _ = v22
        }
        _ = fi18
    }
    _ = len(v17)
    var v23 uint16 = 65535
    v24 := GenS2{F0: v5, F1: 255}
    fmt.Println(v11)
    for fi25 := 0; fi25 < 9; fi25++ {
        switch fi25 {
        case 127:
            v14 = "hello"
        case 0:
            fmt.Println(v11, v11)
        default:
            switch fi25 {
            case -452:
                for fi26 := 0; fi26 < 10; fi26++ {
                    v17 = append(v17, (fi25 + -288))
                    var v27 uint = 265
                    _ = fi26
                    _ = v27
                }
            }
        }
        _ = fi25
    }
    var v28 int16 = v5
    if v6 {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 float32 = v13
            var v31 uint = 2
            switch v1 {
            case "hello":
                v32 := GenS3{F0: v6, F1: -2147483648, F2: ((42 == uint8(v23)) || false)}
                _ = v32
            default:
                v14 = ("a" + v1)
            }
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    v33 := GenS1{F0: 34, F1: 747}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
