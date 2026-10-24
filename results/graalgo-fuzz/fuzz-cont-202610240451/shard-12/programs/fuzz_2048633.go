package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 int32, p1 int, p2 int) uint32 {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{(fi7 * (fi7 - (-(p2)))), -9223372036854775808, (fi7 - p1)}
        _ = fi7
        _ = v8
    }
    var v9 bool = (!(false || (-9223372036854775808 > -9223372036854775808)))
    v10 := GenS1{F0: 0, F1: 924}
    for fi11 := 0; fi11 < 7; fi11++ {
        if (!(v9 && (false && (!v9)))) {
            if false {
                if (-1 == (^(p2))) {
                    v12 := []int{(p1 * p1)}
                    if len(v12) > 0 {
                        v12[0] = fi11
                    }
                    _ = v12
                }
            } else {
                v13 := []int{(-(fi11)), p1}
                v14 := GenS1{F0: 57, F1: 65535}
                _ = v13
                _ = v14
            }
            p2 = 32767
            var v15 int32 = (p0 & p0)
            _ = v15
        }
        var v16 uint8 = 192
        _ = fi11
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2147483648
}

func main() {
    var v1 int = 127
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    var v6 float64 = ((float64((3.14159 - 0.1)) + v3) / (v3 / (-2.25 - v3)))
    if (42 == 0) {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint16 = 637
                _ = fi18
                _ = v19
            }
            if v5 {
                v2[2147483648] = (fi17 | fi17)
            } else {
                v20 := GenS1{F0: 4294967295, F1: 731}
                _ = v20
            }
            v1 += (fi17 % ((v1) | 1))
            _ = fi17
        }
    } else {
        v21 := GenS1{F0: 1, F1: 296}
        _ = v21
    }
    v6 = ((3.14159 / v6) * ((-2.25 / ((v6 - ((v6 / ((v6 + (1e10 / ((((1.5 - v6) * ((v6 - -2.25) - (float64((-2.25 + 0.1)) / v6))) * (1.5 * (float64(1e10) - (3.14159 + 1.5)))) * v3))) * ((1.5 * 3.14159) - v3))) - (((v3 - v6) + float64(((float32(1e10) + (0.0 - (0.1 / 3.14159))) * 1.5))) - (v6 - v6)))) * v6)) - v6))
    switch "foo bar" {
    case "hello":
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(fi22, v6)
                var v24 int64 = -1
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    case "foo bar":
        var v25 uint64 = 18446744073709551615
        _ = v25
    default:
        v26 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v26
    }
    var v27 uint = 18446744073709551615
    var v28 uint16 = 565
    v29 := []int{-9223372036854775808, 32768, (v1 - v1)}
    var v30 int8 = 100
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    v30 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
