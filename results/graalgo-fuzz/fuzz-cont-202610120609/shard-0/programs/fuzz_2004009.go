package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []string{}
    var v2 float32 = -2.25
    var v3 float32 = 0.0
    var v4 int32 = 0
    v5 := []int{32767, 42}
    v6 := []int{-1, -2147483648, 128}
    var v7 uint8 = 69
    for fi8 := 0; fi8 < 17; fi8++ {
        v4 = (v4 >> ((42) & 63))
        fmt.Println(v3, v4)
        v4 += 2147483647
        _ = fi8
    }
    if (true && (-32768 == -32768)) {
        v9 := GenS1{F0: 32768, F1: v7, F2: 65535}
        for fi10 := 0; fi10 < 5; fi10++ {
            switch "\n\t" {
            case "":
                var v11 uint = 2
                _ = v11
            }
            var v12 string = (("\n\t" + ("" + ("foo bar" + ("hello" + "hello")))) + "foo bar")
            _ = fi10
            _ = v12
        }
        _ = v9
    }
    v7 += 127
    var v13 uint16 = 65535
    if (!true) {
        v3 -= 1.5
        v1 = append(v1, ("a" + "\n\t"))
    }
    if (-2147483648 == -2147483648) {
        v14 := GenS2{F0: 4, F1: -373}
        fmt.Println(v3, v2, v4)
        _ = v14
    } else {
        if (true || false) {
            if len(v6) > 0 {
                v6[0] = -1
            }
            v3 = (((v3 - -2.25) - ((1e10 + float32(float64(0.0))) - -2.25)) / ((v2 - (0.0 - (0.0 / v3))) - v2))
        }
        var v15 string = "a"
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    switch 42 {
    case 2147483647:
        fmt.Println(v3, v2)
    case 42:
        var v17 uint = 653
        _ = v17
    default:
        var v18 float32 = v3
        _ = v18
    }
    v4 = (v4 << ((167) & 63))
    var v19 uint8 = 160
    switch "foo bar" {
    case "a":
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v3)
            if (425 > 500) {
                var v21 float32 = (-2.25 * 1e10)
                var v22 uint64 = 2
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v4 = (-(v4))
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = fi23
                }
                switch "\n\t" {
                case "\n\t":
                    var v24 int = (fi20 / (((^(fi20))) | 1))
                    _ = v24
                }
                v4 -= v4
            }
            v4 = v4
            _ = fi20
        }
    case "foo bar":
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint = 241
            _ = fi25
            _ = v26
        }
    case "hello":
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS2{F0: 2, F1: 42}
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
