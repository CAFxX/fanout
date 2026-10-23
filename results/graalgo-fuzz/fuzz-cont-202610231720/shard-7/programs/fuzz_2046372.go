package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 int16) int16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: 2147483647, F1: 320}
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 1
    if false {
        if ((1e10 / ((-2.25 - 0.1) - 0.1)) <= 3.14159) {
            fmt.Println(v6, p0)
        } else {
            var v7 float32 = (((1.5 - 1.5) * (-2.25 - ((0.0 + 0.0) + 3.14159))) + ((-2.25 / 3.14159) * -2.25))
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int16 = p0
                var v10 uint16 = 890
                switch (fi8 >> ((226) & 63)) {
                case 32767:
                    v11 := []int{128}
                    _ = v11
                case 9223372036854775807:
                    v9 = p0
                case -2147483648:
                    var v12 int64 = 42
                    _ = v12
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            var v13 int16 = (p0 & (-(p0)))
            _ = v13
            _ = v7
        }
    } else {
        var v14 float64 = 1e10
        v15 := []int{298, 42}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: 255, F1: (42 < 261)}
    _ = p0
    _ = v16
    _ = v6
    return p0
}
func genFn1(p0 int16, p1 int8) uint16 {
    if (2 <= 8) {
        if (false && ((0.0 / ((-2.25 / -2.25) - 1.5)) <= (0.1 / (float32(0.1) / 1.0)))) {
            if true {
                var v17 string = "hello"
                _ = v17
            }
        }
    }
    if (0 != 4294967295) {
        var v18 uint32 = 2147483647
        _ = v18
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 float32 = float32((-2.25 / 1e10))
    var v2 uint = 2
    var v3 int = 32767
    v1 = (float32(1.5) - (1.5 + (v1 * v1)))
    var v19 int64 = -841
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS2{F0: 245, F1: false}
        v3 = fi20
        var v22 uint8 = 80
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := []int{1, 32768}
    for fi24 := 0; fi24 < 3; fi24++ {
        v2 -= (^(v2))
        switch "a" {
        case "hello":
            var v25 uint32 = 2147483648
            _ = v25
        }
        _ = fi24
    }
    _ = len(v23)
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 uint32 = 4294967295
        _ = fi26
        _ = v27
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
