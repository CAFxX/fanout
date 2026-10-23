package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }

func main() {
    var v1 bool = (false && false)
    var v2 uint64 = 4294967295
    var v3 bool = ((float64(1e10) - 0.0) >= -2.25)
    var v4 float64 = 3.14159
    v5 := GenS1{F0: "foo bar", F1: 1, F2: 816}
    if false {
        if (v3 && v3) {
            var v6 int64 = -1
            v4 -= v4
            _ = v6
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint16 = 1
            switch 4294967295 {
            case 0:
                for fi9 := 0; fi9 < 15; fi9++ {
                    v3 = false
                    fmt.Println(fi9)
                    _ = fi9
                }
            case 32767:
                v10 := GenS1{F0: ((("\n\t" + "\n\t") + "") + ""), F1: v8, F2: 2}
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    case 128:
        v3 = v3
    case -9223372036854775808:
        if v3 {
            if (65535 == 42) {
                if v3 {
                    v3 = true
                    v11 := []int{2147483647, 0}
                    fmt.Println(v1, v4, v2)
                    _ = v11
                } else {
                    v12 := GenS2{F0: 42}
                    _ = v12
                }
            }
            var v13 float32 = 1.5
            _ = v13
        }
    default:
        fmt.Println(v3, v2, v4)
    }
    switch "foo bar" {
    case "hello":
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
        }
    case "":
        fmt.Println(v1)
    default:
        fmt.Println(v2)
    }
    if ((-94 != -459) || v1) {
        var v16 uint16 = 2
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := []int{fi17, fi17}
            fmt.Println(v3, v2)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v4 += (float64((-2.25 + 0.1)) / v4)
    v3 = true
    v19 := GenS3{F0: 195, F1: 42, F2: (^(v2))}
    if (v1 || v3) {
        v4 = 0.1
        for fi20 := 0; fi20 < 18; fi20++ {
            if (v3 && v1) {
                var v21 int64 = 0
                switch fi20 {
                case 32767:
                    v22 := []int{fi20, 1, (fi20 << ((157) & 63))}
                    _ = v22
                case 127:
                    var v23 string = "foo bar"
                    _ = v23
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := []int{-318, (fi20 ^ (fi24 | (fi24 % (((fi24 ^ fi20)) | 1))))}
                    _ = fi24
                    _ = v25
                }
                _ = v21
            }
            fmt.Println(v1)
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := []int{0}
                v3 = (0.1 == ((-2.25 - -2.25) * (1.5 / -2.25)))
                var v28 float64 = ((v4 / 1e10) / 0.1)
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi20
        }
    } else {
        v1 = false
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int32 = -325
            v31 := []int{9223372036854775807, fi29, (fi29 + (fi29 | (fi29 * fi29))), fi29}
            _ = fi29
            _ = v30
            _ = v31
        }
        fmt.Println(v4, v3)
    }
    var v32 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
