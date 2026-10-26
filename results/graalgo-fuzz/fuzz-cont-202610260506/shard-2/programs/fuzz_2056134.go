package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int64, p1 int8) int8 {
    v4 := []int{201}
    v4 = append(v4, -2147483648)
    if true {
        var v5 uint = 42
        switch "hello" {
        case "hello":
            p0 -= -1
        default:
            var v6 uint32 = 2147483647
            _ = v6
        }
        var v7 uint64 = 188
        _ = v5
        _ = v7
    } else {
        var v8 uint32 = 42
        if true {
            p0 -= p0
            p0 -= (p0 / ((p0) | 1))
        }
        v9 := GenS1{F0: -34, F1: 0}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            p0 += 9223372036854775807
            switch "\n\t" {
            case "hello":
                v11 := GenS1{F0: (fi10 * fi10), F1: 0}
                _ = v11
            case "foo bar":
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 int = 4294967295
                    var v14 float64 = 1e10
                    var v15 int32 = -605
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            default:
                var v16 string = "a"
                _ = v16
            }
            if true {
                _ = len(v4)
                var v17 int64 = 1
                _ = v17
            } else {
                fmt.Println(fi10, p1, p1)
            }
        } else {
            var v18 uint16 = 255
            _ = v18
        }
        v19 := GenS1{F0: -1, F1: 18446744073709551615}
        if len(v4) > 0 {
            v4[0] = fi10
        }
        _ = fi10
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 uint = 2
    var v2 int32 = 238
    var v3 float64 = float64(1.5)
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{(fi20 << (((v1 - v1)) & 63)), fi20, fi20, (fi20 >> ((v1) & 63))}
        v21 = append(v21, fi20)
        var v22 float64 = ((((0.0 + 1e10) * 3.14159) * (float64(-2.25) / v3)) / (v3 * float64((1.5 - (((0.0 / 1.5) + (0.1 * (float32((v3 + ((v3 + 1.5) + (v3 * (((1.5 + 3.14159) - v3) - v3))))) * 1e10))) - 3.14159)))))
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 65535
    var v24 string = "\n\t"
    v25 := []int{-2147483648, -2147483648, -409}
    v1 = 564
    v1 -= v1
    v26 := GenS2{F0: "\n\t", F1: -2147483648}
    v27 := GenS3{F0: (v23 >> (((v1 * v1)) & 63)), F1: uint16(v1)}
    switch v24 {
    case "hello":
        var v28 uint = (^(v1))
        _ = v28
    }
    var v29 string = ""
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v2, v24, fi30)
        fmt.Println(v2, v24)
        if (!(false && true)) {
            switch 2147483647 {
            case -9223372036854775808:
                if (0.0 <= (v3 / 1.0)) {
                    var v31 uint = 18446744073709551615
                    _ = v31
                }
            case -55:
                if false {
                    fmt.Println(v29, v29, v24)
                    v1 = uint(v23)
                    _ = len(v25)
                }
            case 1:
                for fi32 := 0; fi32 < 18; fi32++ {
                    fmt.Println(v23, v24, v24)
                    _ = fi32
                }
            }
            v1 = (v1 >> ((v1) & 63))
        } else {
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 int64 = 0
                v3 = 0.1
                _ = fi33
                _ = v34
            }
            var v35 uint32 = v23
            _ = v35
        }
        _ = fi30
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        var v37 float32 = (1.5 - 3.14159)
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        v3 = 1.5
        v1 = (v1 % (((v1 / ((255) | 1))) | 1))
        v29 = v29
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
