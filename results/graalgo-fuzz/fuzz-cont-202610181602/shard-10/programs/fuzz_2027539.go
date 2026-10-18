package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint) int16 {
    v6 := GenS1{F0: "foo bar", F1: 1}
    var v7 string = "\n\t"
    switch -2147483648 {
    case 9223372036854775807:
        var v8 int = -327
        _ = v8
    case 128:
        if (true && (!false)) {
            v9 := GenS2{F0: 101}
            _ = v9
        } else {
            switch v7 {
            case "a":
                fmt.Println(v7, p0)
            case "\n\t":
                v7 = ""
            default:
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS3{F0: fi10}
                    fmt.Println(p0)
                    _ = fi10
                    _ = v11
                }
            }
        }
    default:
        var v12 uint32 = 2147483648
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 int32) int8 {
    v13 := GenS4{F0: "foo bar", F1: false, F2: 2}
    p0 = p0
    var v14 int32 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 78
}

func main() {
    var v1 uint16 = 32767
    var v2 bool = (!(v1 >= (v1 * v1)))
    v3 := make(map[string]uint32)
    var v4 uint16 = (^(v1))
    var v5 float64 = float64(((0.0 * (0.1 * (1.5 + float32(-2.25)))) + 1.5))
    v15 := GenS5{F0: 255}
    v16 := GenS3{F0: 127}
    v4 = (v4 & v4)
    v1 = (^(v1))
    v17 := GenS5{F0: 65535}
    var v18 uint32 = 65535
    if ((-842 > -9223372036854775808) || (v2 && v2)) {
        fmt.Println(v5)
    } else {
        _ = len(v3)
        v18 = (^(v18))
    }
    if v2 {
        for fi19 := 0; fi19 < 1; fi19++ {
            if v2 {
                var v20 uint = 4294967295
                _ = v20
            } else {
                v3[("" + "a")] = v18
            }
            _ = fi19
        }
    }
    v21 := v3["foo bar"]
    var v22 uint64 = 0
    if (true && (v22 >= 18446744073709551615)) {
        v4 = v4
    }
    if (-1 <= 1) {
        v22 += (v22 + (v22 | (v22 >> ((768) & 63))))
        v23 := GenS5{F0: 63}
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint32 = 65535
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if v2 {
            var v27 uint = 700
            _ = v27
        }
        _ = fi26
    }
    fmt.Println(v22, v4, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
