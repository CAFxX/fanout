package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool, p1 uint32, p2 int32) int64 {
    if false {
        var v7 uint64 = 42
        _ = v7
    }
    var v8 uint = 124
    v8 = (v8 & (v8 & 2))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}
func genFn1(p0 string) uint8 {
    v9 := []int{128}
    if (65535 > 985) {
        if false {
            var v10 uint64 = 4294967295
            switch "a" {
            case "\n\t":
                _ = len(v9)
            case "foo bar":
                v10 -= 520
            }
            fmt.Println(v10)
            _ = v10
        }
        p0 = ("foo bar" + "foo bar")
    }
    var v11 bool = (0 >= 127)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = "\n\t"
        var v14 uint16 = 180
        switch fi12 {
        case -597:
            var v15 int64 = genFn0(v11, 2147483648, 1)
            _ = v15
        case -1:
            var v16 uint32 = 611
            _ = v16
        case 766:
            var v17 float64 = 1e10
            _ = v17
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v11
    _ = v9
    return 128
}

func main() {
    var v1 uint = 255
    var v2 int16 = -32768
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    var v5 uint8 = 1
    v6 := []string{("hello" + ""), "hello", "", ("\n\t" + "foo bar")}
    var v18 uint64 = 18446744073709551615
    v19 := GenS1{F0: (v2 & (v2 / (((v2 + (v2 >> (((v1 % ((v1) | 1))) & 63)))) | 1))), F1: 128}
    v5 += 2
    v20 := GenS2{F0: 602}
    if (2147483647 == 883) {
        v18 -= v18
        var v21 uint64 = (^(v18))
        var v22 float32 = v3
        _ = v21
        _ = v22
    } else {
        fmt.Println(v2, v18, v3)
        var v23 string = ("hello" + "a")
        _ = v23
    }
    switch -124 {
    case 127:
        if (("" > "a") && (true || false)) {
            v1 = v1
            var v24 bool = (2147483647 > -9223372036854775808)
            v5 += 10
            _ = v24
        }
    default:
        v2 += (v2 | (-(v2)))
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v25 uint = v1
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
