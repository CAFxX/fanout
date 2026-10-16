package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 int, p1 uint32) uint64 {
    if (0 <= -1) {
        for fi4 := 0; fi4 < 18; fi4++ {
            p0 += p0
            _ = fi4
        }
    } else {
        p1 += (^(p1))
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{int(p1), 32768, (fi5 | fi5), (fi5 % ((fi5) | 1))}
        var v7 int16 = 1
        fmt.Println(p1, fi5, p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 bool = false
    var v2 uint16 = 65535
    var v3 string = ""
    v8 := []int{-9223372036854775808}
    v8 = append(v8, -823)
    for fi9 := 0; fi9 < 7; fi9++ {
        if v1 {
            if (92 <= -1) {
                v2 -= 808
                if (2147483647 > 42) {
                    var v10 uint64 = 1
                    _ = v10
                } else {
                    var v11 float64 = -2.25
                    _ = v11
                }
            }
        } else {
            v12 := GenS1{F0: 32767, F1: 587, F2: v2}
            v2 -= 953
            _ = v12
        }
        switch v3 {
        case "\n\t":
            var v13 bool = ((2 != uint64(v2)) || ((v1 && (true || v1)) && false))
            _ = v13
        default:
            var v14 float64 = float64(float32((float64(-2.25) / 0.1)))
            _ = v14
        }
        fmt.Println(v3, v3, v2)
        _ = fi9
    }
    v1 = ((1e10 + 1.5) > 1.5)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if v1 {
        v15 := GenS2{F0: 2, F1: 60, F2: v3}
        v16 := GenS3{F0: -9223372036854775808, F1: 32768, F2: -812}
        _ = v15
        _ = v16
    }
    fmt.Println(v3, v2, v1)
    if ((((v2 % (((v2 - (v2 / ((v2) | 1)))) | 1)) == 486) || v1) || (23 < 127)) {
        fmt.Println(v1, v3, v2)
        if len(v8) > 0 {
            v8[0] = 32768
        }
        if v1 {
            v8 = append(v8, 4294967295)
        }
    }
    switch -2147483648 {
    case 32768:
        var v17 uint = 1
        _ = v17
    case 219:
        if (!v1) {
            var v18 int8 = 1
            v3 = (v3 + (((v3 + v3) + ("a" + v3)) + ("a" + v3)))
            var v19 uint16 = (v2 / (((v2 + v2)) | 1))
            _ = v18
            _ = v19
        }
    case 449:
        fmt.Println(v3)
    default:
        v20 := GenS4{F0: 4294967295, F1: 42, F2: 127}
        _ = v20
    }
    v2 -= v2
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint16 = v2
        if false {
            v3 = (v3 + v3)
            if (100 >= -1) {
                v3 = v3
                v22 -= 2
                for fi23 := 0; fi23 < 16; fi23++ {
                    v1 = (!v1)
                    var v24 float64 = 1e10
                    v3 = v3
                    _ = fi23
                    _ = v24
                }
            } else {
                var v25 int = -566
                var v26 int64 = 1
                _ = v25
                _ = v26
            }
        }
        _ = len(v8)
        _ = fi21
        _ = v22
    }
    v1 = ("foo bar" > v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
