package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 uint64, p1 int, p2 uint) int64 {
    var v4 uint64 = (p0 / (((^(p0))) | 1))
    var v5 int32 = 441
    switch "" {
    case "foo bar":
        switch "a" {
        case "":
            for fi6 := 0; fi6 < 6; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    fmt.Println(p1, fi6)
                    v8 := []int{fi7}
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
        case "foo bar":
            switch p1 {
            case 0:
                if true {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 float64 = -2.25
                        v11 := []int{p1}
                        fmt.Println(p0, v10, p0)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    var v12 int = int(v4)
                    for fi13 := 0; fi13 < 4; fi13++ {
                        fmt.Println(fi13, p0)
                        _ = fi13
                    }
                    _ = v12
                }
            case 473:
                switch "foo bar" {
                case "a":
                    v14 := GenS1{F0: (-(v5)), F1: (p2 | p2), F2: 8}
                    _ = v14
                case "foo bar":
                    var v15 float64 = 3.14159
                    _ = v15
                case "hello":
                    v4 += v4
                default:
                    var v16 uint16 = 255
                    _ = v16
                }
            }
        case "hello":
            var v17 int16 = 0
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        p0 -= (p0 % ((v4) | 1))
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]uint32)
    v3 := []string{"hello", "hello", "", "\n\t"}
    var v19 int16 = -992
    var v20 int16 = 301
    if (float32(0.0) < (0.1 + -2.25)) {
        v21 := GenS1{F0: -2147483648, F1: 251, F2: 115}
        _ = v21
    } else {
        v2[v1] = 2
        for fi22 := 0; fi22 < 18; fi22++ {
            v19 = (v20 & v20)
            v1 += (v1 / (((v1 & (fi22 << ((4294967295) & 63)))) | 1))
            _ = fi22
        }
    }
    v23 := GenS1{F0: -2147483648, F1: 42, F2: 128}
    var v24 bool = false
    v19 = 862
    if v24 {
        var v25 bool = (((float64(3.14159) * float64(((1e10 * 0.0) - 3.14159))) >= 1e10) || v24)
        _ = v25
    } else {
        var v26 uint = 350
        v26 = (v26 & 42)
        _ = v26
    }
    if false {
        if v24 {
            for fi27 := 0; fi27 < 11; fi27++ {
                v2[-1] = 0
                fmt.Println(v20, v19)
                v19 += 0
                _ = fi27
            }
            v19 += v20
        }
        var v28 int = v1
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
}
