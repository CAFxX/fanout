package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int64) int {
    v6 := []int{42, -1}
    p0 = (p0 - (p0 % ((0) | 1)))
    var v7 int64 = 9223372036854775807
    _ = p0
    _ = v6
    _ = v7
    return 312
}
func genFn1(p0 uint) string {
    v8 := []int{}
    for fi9 := 0; fi9 < 13; fi9++ {
        if len(v8) > 0 {
            v8[0] = (fi9 << ((p0) & 63))
        }
        _ = fi9
    }
    var v10 float32 = 0.0
    _ = p0
    _ = v10
    _ = v8
    return ("hello" + "\n\t")
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    var v4 uint32 = 2147483647
    var v5 bool = false
    _ = len(v1)
    switch "foo bar" {
    case "":
        fmt.Println(v4, v5)
    case "a":
        _ = len(v2)
    default:
        switch -2147483648 {
        case 128:
            switch 4294967295 {
            case 2147483647:
                if v5 {
                    switch -590 {
                    case 5:
                        v3 += 1e10
                    case -1:
                        v11 := GenS1{F0: (^(v4)), F1: -67}
                        _ = v11
                    }
                } else {
                    v5 = (!(!v5))
                }
            case 32768:
                v12 := GenS1{F0: (v4 | v4), F1: genFn0(1)}
                _ = v12
            default:
                for fi13 := 0; fi13 < 10; fi13++ {
                    v14 := GenS1{F0: (v4 << ((469) & 63)), F1: (^(fi13))}
                    v3 += 0.0
                    v3 -= 0.1
                    _ = fi13
                    _ = v14
                }
            }
        default:
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int16 = -1
                fmt.Println(fi15, fi15)
                _ = fi15
                _ = v16
            }
        }
    }
    fmt.Println(v4)
    v5 = v5
    if (0 == 0) {
        var v17 bool = v5
        for fi18 := 0; fi18 < 11; fi18++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = float32(1.5)
            var v19 bool = v5
            _ = fi18
            _ = v19
        }
        fmt.Println(v17)
        _ = v17
    }
    v20 := GenS2{F0: 2147483647, F1: -79}
    switch "hello" {
    case "":
        for fi21 := 0; fi21 < 18; fi21++ {
            _ = len(v1)
            var v22 float32 = v3
            _ = fi21
            _ = v22
        }
    case "a":
        v4 += (v4 / ((v4) | 1))
    }
    v23 := GenS3{F0: 155, F1: -2147483648}
    fmt.Println(v4, v3, v5)
    v24 := GenS3{F0: -840, F1: 1}
    switch "a" {
    case "a":
        var v25 uint8 = 221
        _ = v25
    default:
        v3 -= v3
    }
    v3 = ((1e10 - v3) * float32(0.1))
    v5 = (((false || true) || (v5 && ((float64(-2.25) <= 0.0) && true))) && (727 <= 9223372036854775807))
    var v26 uint32 = (v4 >> ((255) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
