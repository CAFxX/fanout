package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 bool = (true && (!false))
    var v2 int = 2147483647
    var v3 uint64 = 4294967295
    var v4 uint32 = 2147483648
    switch v2 {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 20; fi5++ {
            if ((0.1 / 1.0) == (-2.25 - (((float64(3.14159) / 1.0) / 0.1) * (1.5 * 1e10)))) {
                v6 := GenS1{F0: 1, F1: (v4 ^ v4), F2: 42}
                v7 := GenS1{F0: -806, F1: 4294967295, F2: 0}
                v4 -= (v4 / ((v4) | 1))
                _ = v6
                _ = v7
            }
            v2 = 4294967295
            _ = fi5
        }
    }
    if true {
        switch "\n\t" {
        case "":
            var v8 int32 = -340
            _ = v8
        default:
            var v9 uint = 42
            _ = v9
        }
    }
    switch (("\n\t" + "\n\t") + "\n\t") {
    case "a":
        v10 := []int{}
        _ = v10
    }
    v11 := GenS1{F0: 0, F1: (^(v4)), F2: 42}
    switch "foo bar" {
    case "hello":
        v12 := []int{}
        _ = v12
    }
    switch "\n\t" {
    case "a":
        v13 := []int{v2, (v2 * (v2 >> ((93) & 63))), v2, (-(v2))}
        _ = v13
    case "hello":
        fmt.Println(v2, v2)
    default:
        v14 := GenS2{F0: 127, F1: v3, F2: 196}
        _ = v14
    }
    var v15 uint8 = 1
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := []int{(fi16 << ((2) & 63)), (^(fi16))}
            fmt.Println(v15, v4, v1)
            _ = fi17
            _ = v18
        }
        var v19 uint = 201
        _ = fi16
        _ = v19
    }
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 float32 = 0.0
        for fi23 := 0; fi23 < 6; fi23++ {
            if ((v20 & (^(v20))) <= (v20 + int16(v15))) {
                v24 := GenS1{F0: int16(fi21), F1: v4, F2: -2147483648}
                switch 1 {
                case 794:
                    v25 := GenS3{F0: v15, F1: 42, F2: 65535}
                    _ = v25
                case 470:
                    var v26 uint8 = uint8(v4)
                    _ = v26
                }
                v27 := []int{70, (^(fi23)), (fi23 * (fi21 & v2)), 42}
                _ = v24
                _ = v27
            } else {
                fmt.Println(v1, v20)
            }
            v2 += v2
            var v28 uint8 = (^(v15))
            _ = fi23
            _ = v28
        }
        v3 += uint64(v15)
        _ = fi21
        _ = v22
    }
    v4 = v4
    v1 = (9223372036854775807 > 127)
    v29 := GenS4{F0: (-(v20))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
}
