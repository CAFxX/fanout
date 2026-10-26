package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
func genFn0() uint32 {
    var v5 int32 = 2147483647
    switch "a" {
    case "":
        v5 = 910
    case "hello":
        switch 438 {
        case -2147483648:
            var v6 float32 = 0.1
            _ = v6
        case -9223372036854775808:
            fmt.Println(v5)
        }
    case "\n\t":
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v5 = -814
                    _ = fi8
                }
                v9 := []int{(fi7 + fi7)}
                var v10 bool = false
                _ = v10
                _ = v9
            }
            var v11 int = (fi7 - -9223372036854775808)
            switch "foo bar" {
            case "foo bar":
                var v12 string = (("" + "\n\t") + "\n\t")
                _ = v12
            case "hello":
                var v13 int = v11
                _ = v13
            }
            _ = fi7
            _ = v11
        }
    default:
        fmt.Println(v5)
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: v5, F1: false}
        var v16 float64 = (1e10 + 3.14159)
        _ = fi14
        _ = v15
        _ = v16
    }
    v5 = (-(v5))
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint8) uint8 {
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := []int{(fi17 ^ fi17), (fi17 - (fi17 << ((1) & 63))), fi17, fi17}
        p0 = p0
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: -2147483648, F1: true}
    v20 := []int{4294967295, -147}
    var v21 string = "foo bar"
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 0
    var v3 uint8 = (v2 ^ (v2 / ((v2) | 1)))
    v4 := make(map[string]uint32)
    var v22 int8 = 127
    var v23 int8 = (v22 & (v22 + (v22 << ((1) & 63))))
    v3 = v3
    v2 -= genFn1(1, genFn1(32767, v2))
    v3 -= v2
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            _ = len(v1)
            _ = fi24
        }
    }
    v1[9223372036854775807] = 65535
    var v25 uint = 255
    if (2 == 418) {
        var v26 string = (("\n\t" + "") + (("" + ("hello" + "")) + ("foo bar" + "\n\t")))
        v25 = v25
        _ = v26
    }
    v3 = (v2 << ((v25) & 63))
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
