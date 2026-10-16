package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int = -2147483648
    v2 := []uint32{123, 4294967295}
    var v3 float64 = ((float64(-2.25) / -2.25) + 1e10)
    v4 := GenS1{F0: false, F1: 0}
    _ = len(v2)
    if (false || true) {
        if false {
            v3 = v3
        } else {
            var v5 uint32 = 284
            v6 := GenS1{F0: false, F1: (^(v5))}
            var v7 bool = (!true)
            _ = v5
            _ = v6
            _ = v7
        }
        v3 = v3
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                switch "a" {
                case "\n\t":
                    v1 = (fi8 ^ (v1 + fi9))
                case "hello":
                    v2 = append(v2, 0)
                default:
                    var v10 float32 = 3.14159
                    _ = v10
                }
                v11 := GenS1{F0: (537 == 4294967295), F1: 2}
                _ = fi9
                _ = v11
            }
            if (1 != 42) {
                var v12 string = "a"
                var v13 float32 = -2.25
                if true {
                    var v14 float64 = v3
                    v3 = v14
                    _ = v14
                }
                _ = v12
                _ = v13
            }
            _ = fi8
        }
    }
    var v15 uint16 = 32767
    var v16 string = "hello"
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v1, fi17)
        _ = fi17
    }
    v15 -= (^(v15))
    var v18 int16 = 1000
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            switch (v16 + v16) {
            case "hello":
                var v22 uint16 = (^(v15))
                _ = v22
            case "\n\t":
                var v23 uint32 = 507
                _ = v23
            default:
                v19 += ((1e10 + float64((-2.25 / (-2.25 - (0.0 * float32(v19)))))) + v19)
            }
            v3 = v3
            _ = fi21
        }
        var v24 int8 = -63
        _ = fi20
        _ = v24
    }
    v25 := GenS1{F0: (v1 <= (v1 & v1)), F1: 386}
    fmt.Println(v18, v16)
    v26 := GenS2{F0: v18, F1: 1}
    v1 = v1
    for fi27 := 0; fi27 < 12; fi27++ {
        v18 += v18
        _ = fi27
    }
    v19 = float64(((1e10 + 0.0) * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
