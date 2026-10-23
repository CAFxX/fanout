package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    var v1 uint32 = 2
    var v2 uint = 42
    var v3 uint8 = 255
    var v4 float64 = float64(0.0)
    v5 := []string{("hello" + "a"), (((("foo bar" + ("hello" + "a")) + (("foo bar" + "\n\t") + "")) + ("hello" + "hello")) + "foo bar"), "foo bar", "\n\t"}
    switch "foo bar" {
    case "a":
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 uint64 = 571
                v9 := GenS1{F0: v1, F1: 1000, F2: "\n\t"}
                var v10 int16 = 696
                v5 = append(v5, "hello")
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 13; fi11++ {
                v2 -= v2
                _ = fi11
            }
            _ = fi6
        }
    default:
        fmt.Println(v4, v1)
    }
    switch -523 {
    case 127:
        switch "foo bar" {
        case "\n\t":
            var v12 string = "a"
            _ = v12
        case "hello":
            switch -1 {
            case 184:
                var v13 uint8 = v3
                _ = v13
            default:
                v4 += v4
            }
        default:
            var v14 uint8 = (^(v3))
            _ = v14
        }
    case -2147483648:
        if true {
            if len(v5) > 0 {
                v5[0] = (("foo bar" + "") + "\n\t")
            }
            v15 := GenS2{F0: 42, F1: (^(v1)), F2: ""}
            v1 += 2
            _ = v15
        }
    }
    fmt.Println(v2, v1)
    v2 -= (v2 ^ v2)
    v1 -= 2147483647
    switch 32768 {
    case 4294967295:
        v16 := GenS1{F0: v1, F1: -32768, F2: ""}
        _ = v16
    case -2147483648:
        var v17 int32 = -823
        _ = v17
    case 461:
        var v18 float32 = -2.25
        _ = v18
    }
    var v19 float32 = (0.0 / 1.0)
    switch -9223372036854775808 {
    case -1:
        v4 += v4
    case -2147483648:
        for fi20 := 0; fi20 < 20; fi20++ {
            v1 -= (v1 % (((v1 | (v1 % (((v1 / (((v1 * v1)) | 1))) | 1)))) | 1))
            _ = fi20
        }
    }
    var v21 uint8 = v3
    v19 = (((1e10 - v19) - float32(v4)) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
