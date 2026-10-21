package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 bool) uint64 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int64 = -9223372036854775808
        var v7 uint16 = 2
        if (45 != 42) {
            v8 := GenS1{F0: fi5, F1: int16(v7)}
            _ = v8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = p0
    p0 = (false && p0)
    if (!p0) {
        var v9 int = 1
        v10 := GenS1{F0: (^(v9)), F1: 1000}
        _ = v10
        _ = v9
    } else {
        p0 = (true || false)
    }
    _ = p0
    return 4294967295
}
func genFn1(p0 uint8, p1 int64, p2 string) string {
    var v11 int = int(p0)
    v11 += v11
    p0 = (p0 + 95)
    var v12 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p2
}

func main() {
    var v1 int = 789
    var v2 bool = false
    var v3 int64 = 0
    var v4 bool = (32767 == 2)
    for fi13 := 0; fi13 < 5; fi13++ {
        v4 = (("\n\t" >= "foo bar") || true)
        var v14 uint16 = 255
        _ = fi13
        _ = v14
    }
    if false {
        v4 = (v4 || v2)
        v15 := []int{(-(v1))}
        _ = v15
    }
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 float64 = (0.0 + 3.14159)
        v18 := GenS2{F0: 18446744073709551615, F1: 1}
        switch "\n\t" {
        case "a":
            var v19 uint64 = 1
            _ = v19
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                switch "a" {
                case "foo bar":
                    var v21 uint16 = 255
                    _ = v21
                }
                v17 = v17
                v22 := []int{v1}
                _ = fi20
                _ = v22
            }
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v23 bool = (v3 == 0)
    if v23 {
        switch 9223372036854775807 {
        case -635:
            v24 := GenS3{F0: v3, F1: 1}
            _ = v24
        case -2147483648:
            v3 = v3
        case -9223372036854775808:
            for fi25 := 0; fi25 < 20; fi25++ {
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := []int{-714, (v1 | 395), 559, (fi26 << ((255) & 63))}
                    _ = fi26
                    _ = v27
                }
                switch v1 {
                case 1:
                    v3 = v3
                }
                var v28 uint32 = 741
                _ = fi25
                _ = v28
            }
        }
    } else {
        v23 = v23
    }
    fmt.Println(v2)
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS3{F0: 9223372036854775807, F1: 32767}
        var v31 int16 = 32767
        var v32 int = fi29
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int32 = 2147483647
    v3 = v3
    v1 += (v1 % (((v1 & v1)) | 1))
    var v34 int = v1
    v4 = ((v2 && ((-2.25 + 0.1) < -2.25)) && true)
    v2 = (v3 <= v3)
    switch "\n\t" {
    case "\n\t":
        v3 -= v3
    case "":
        switch v34 {
        case 45:
            v4 = false
        case -23:
            var v35 float64 = (0.1 - 0.1)
            _ = v35
        default:
            v1 -= (v1 << ((65535) & 63))
        }
    default:
        if (1e10 > ((-2.25 + (1e10 + ((-2.25 + 0.1) + 0.0))) + float64((((0.0 * 0.1) - (((-2.25 + 1.5) + 1.5) + (((0.1 * -2.25) - (1.5 + -2.25)) - float32((-2.25 - 1e10))))) * (1e10 / 1e10))))) {
            v2 = (((((1.5 / (-2.25 * 1e10)) * 3.14159) + (0.1 / (float64(3.14159) - (1e10 * 0.0)))) / 3.14159) <= float64((float32(1.5) / 1.0)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
