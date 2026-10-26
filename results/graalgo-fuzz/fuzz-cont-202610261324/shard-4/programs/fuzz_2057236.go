package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() uint32 {
    v5 := GenS1{F0: (76 > 42), F1: 42}
    v6 := GenS1{F0: false, F1: 2}
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8) bool {
    p0 = (p0 * p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (^(p0))
        _ = fi7
    }
    _ = p0
    return (false || (true || (!false)))
}

func main() {
    var v1 uint32 = 2
    var v2 int64 = 1
    var v3 int = 128
    v4 := []string{}
    var v8 uint = 922
    switch "a" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v9 int8 = -1
    var v10 int = int(v2)
    var v11 int64 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v9 = (v9 & (v9 ^ int8(v8)))
    if (false && false) {
        v8 += (v8 - 4294967295)
        var v12 uint16 = 334
        _ = v12
    }
    switch "" {
    case "\n\t":
        for fi13 := 0; fi13 < 6; fi13++ {
            switch "foo bar" {
            case "":
                v14 := GenS1{F0: false, F1: 0}
                _ = v14
            }
            _ = fi13
        }
    case "a":
        var v15 int16 = 1
        _ = v15
    }
    switch (v3 - (^(v3))) {
    case 9223372036854775807:
        var v16 uint = 0
        _ = v16
    case 42:
        for fi17 := 0; fi17 < 7; fi17++ {
            if true {
                fmt.Println(v8, v9)
                var v18 int8 = (v9 << (((v8 * 0)) & 63))
                if (!((0.1 / 1.0) <= ((1e10 + (1e10 - float64((1.5 * (((0.0 * 1.5) / (-2.25 / (1e10 / 1e10))) * -2.25))))) + float64(((0.1 / (float32(1e10) + 1e10)) / (((1e10 / 1e10) * (-2.25 * 0.1)) / 0.1)))))) {
                    v2 = (-(v11))
                }
                _ = v18
            }
            var v19 string = ("" + "\n\t")
            _ = fi17
            _ = v19
        }
    default:
        v1 = v1
    }
    v11 = v11
    for fi20 := 0; fi20 < 16; fi20++ {
        v11 += (-(v2))
        v11 -= (^(v2))
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 uint64 = 4294967295
            v23 := GenS2{F0: (!(0.0 != ((-2.25 - 0.0) + 0.0))), F1: uint(v1)}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v8 -= v8
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
