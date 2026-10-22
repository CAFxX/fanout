package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 uint64) uint32 {
    v6 := GenS1{F0: 0, F1: ((669 <= uint32(p0)) && false), F2: 127}
    p0 += p0
    _ = p0
    _ = v6
    return 2147483648
}
func genFn1() int {
    if false {
        if true {
            for fi7 := 0; fi7 < 14; fi7++ {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := []int{}
                    _ = fi8
                    _ = v9
                }
                v10 := []int{fi7, (fi7 ^ (fi7 - 236)), (fi7 << ((42) & 63)), 127}
                _ = fi7
                _ = v10
            }
            if false {
                v11 := []int{749, -2147483648}
                var v12 string = "foo bar"
                var v13 int32 = 921
                _ = v11
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: 18446744073709551615, F1: ((((0.0 / ((-2.25 / 1.0) / float64(1.5))) * (0.1 / 1.0)) - ((3.14159 - -2.25) - 0.0)) <= 1e10), F2: -6}
                v15 := GenS1{F0: 816, F1: false, F2: 10}
                var v16 float32 = (1.5 / (3.14159 - (1e10 - (1.5 + 0.0))))
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 float32 = ((1e10 / (-2.25 / 1.0)) * 0.1)
            _ = v17
        }
    } else {
        var v18 uint16 = 317
        var v19 int8 = -128
        v19 = (v19 / (((v19 >> ((42) & 63))) | 1))
        _ = v18
        _ = v19
    }
    var v20 int16 = 1000
    _ = v20
    return 128
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 float64 = (1.5 / 1.0)
    var v5 float64 = 3.14159
    v21 := []int{}
    var v22 int = 451
    fmt.Println(v4, v1, v1)
    v22 += 457
    v23 := GenS2{F0: 42, F1: (v1 & (v1 + (v1 / ((123) | 1)))), F2: 151}
    fmt.Println(v22, v1, v22)
    switch "foo bar" {
    case "a":
        for fi24 := 0; fi24 < 4; fi24++ {
            v21 = append(v21, fi24)
            delete(v2, genFn1())
            _ = fi24
        }
    case "foo bar":
        fmt.Println(v22)
    default:
        var v25 int64 = (^(v3))
        _ = v25
    }
    v4 = v5
    var v26 int32 = -525
    _ = len(v2)
    var v27 float32 = (3.14159 / ((-2.25 + (-2.25 + 3.14159)) * float32(0.1)))
    v28 := GenS3{F0: (("a" + "foo bar") + ("foo bar" + "")), F1: 587}
    v29 := v2[-9223372036854775808]
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
