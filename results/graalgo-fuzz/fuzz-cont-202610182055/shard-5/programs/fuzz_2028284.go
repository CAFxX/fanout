package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 bool) uint64 {
    p0 = (4294967295 > 533)
    var v6 int64 = -841
    if true {
        var v7 int64 = -424
        var v8 int8 = 0
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return 219
}
func genFn1() bool {
    v9 := GenS1{F0: 458}
    if (!(-1 >= 1)) {
        var v10 float32 = -2.25
        v10 += v10
        var v11 float32 = 1.5
        _ = v10
        _ = v11
    }
    var v12 uint32 = 65535
    v13 := []int{4294967295, 2147483647, 9223372036854775807}
    var v14 float32 = (((1e10 + 3.14159) * 0.0) - (3.14159 * 1.5))
    _ = v12
    _ = v13
    _ = v14
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = 57
    var v3 float64 = ((0.0 / 3.14159) + ((float64((1e10 / float32(((0.1 * float64(1.5)) + 3.14159)))) - (3.14159 * -2.25)) / -2.25))
    var v4 uint16 = v1
    v5 := []int{}
    var v15 int16 = 0
    var v16 int = 659
    switch 127 {
    case -363:
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS1{F0: v15}
            _ = fi17
            _ = v18
        }
    case 127:
        switch "hello" {
        case "foo bar":
            switch 4294967295 {
            case 9223372036854775807:
                fmt.Println(v3, v1)
            default:
                v1 = 1
            }
        case "":
            _ = len(v5)
        default:
            v19 := GenS2{F0: 1, F1: 128, F2: v4}
            _ = v19
        }
    case 128:
        for fi20 := 0; fi20 < 17; fi20++ {
            switch v16 {
            case -9223372036854775808:
                var v21 float64 = float64(1e10)
                _ = v21
            case 32767:
                for fi22 := 0; fi22 < 18; fi22++ {
                    fmt.Println(v2, fi22, v16)
                    _ = fi22
                }
            default:
                var v23 int8 = 66
                _ = v23
            }
            fmt.Println(v3)
            _ = fi20
        }
    default:
        v24 := GenS2{F0: 26, F1: -2147483648, F2: (v1 + v1)}
        _ = v24
    }
    if (v2 <= v2) {
        v3 -= v3
        var v25 uint64 = 0
        _ = v25
    }
    if true {
        v3 += float64(-2.25)
        v2 -= 0
    }
    v26 := GenS1{F0: v15}
    v27 := GenS1{F0: v15}
    var v28 bool = true
    v4 -= v4
    _ = len(v5)
    for fi29 := 0; fi29 < 10; fi29++ {
        v5 = append(v5, 32767)
        switch "hello" {
        case "a":
            switch "" {
            case "a":
                if len(v5) > 0 {
                    v5[0] = fi29
                }
            default:
                var v30 float64 = float64(((1.5 + (0.0 / (1e10 * (0.0 - 1.5)))) * ((float32(0.0) + 0.0) - (-2.25 + 0.1))))
                _ = v30
            }
        case "foo bar":
            v1 = v4
        default:
            if (65535 > 65535) {
                switch "a" {
                case "\n\t":
                    v31 := GenS3{F0: 2, F1: 65535}
                    _ = v31
                case "":
                    var v32 int8 = 127
                    _ = v32
                }
            }
        }
        v28 = ("\n\t" > "a")
        _ = fi29
    }
    v5 = append(v5, (v16 * (v16 % (((v16 | v16)) | 1))))
    v33 := GenS1{F0: -408}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
