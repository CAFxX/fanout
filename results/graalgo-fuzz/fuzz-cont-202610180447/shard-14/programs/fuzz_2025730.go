package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 uint8, p2 int16) bool {
    fmt.Println(p2)
    var v6 int8 = -37
    var v7 int64 = (-(p0))
    v8 := []int{32767, 4294967295, -626}
    switch 32768 {
    case 42:
        p1 += p1
    case -1:
        var v9 uint16 = 255
        _ = v9
    default:
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int, p1 int16, p2 uint16) string {
    if (-2.25 > 3.14159) {
        p0 -= (p0 & p0)
    }
    p1 += p1
    if (42 != p2) {
        switch "" {
        case "foo bar":
            var v10 uint32 = 728
            _ = v10
        default:
            p2 = (p2 - 32767)
        }
        v11 := []int{}
        fmt.Println(p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 42
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 int = 170
    if false {
        if ((2147483648 == 1) || (true || (116 <= 2))) {
            _ = len(v4)
        } else {
            var v12 uint32 = 65535
            switch v5 {
            case 128:
                v2 += (v2 * 65535)
            case -2147483648:
                switch (v5 | (-(v5))) {
                case -9223372036854775808:
                    v1 = v1
                }
            case 910:
                fmt.Println(v2, v12, v12)
            default:
                switch 42 {
                case 128:
                    switch v5 {
                    case 128:
                        v13 := GenS1{F0: false, F1: (^(v1)), F2: uint16(v2)}
                        _ = v13
                    }
                case -1:
                    v5 = 9223372036854775807
                }
            }
            _ = v12
        }
        var v14 int8 = 32
        v15 := GenS1{F0: false, F1: (v1 | (v1 % ((143) | 1))), F2: (v1 - 744)}
        _ = v14
        _ = v15
    } else {
        var v16 uint32 = 2147483648
        v17 := GenS2{F0: -470}
        _ = v16
        _ = v17
    }
    v2 = (v2 | (v2 % ((893) | 1)))
    var v18 uint8 = 209
    var v19 uint = v2
    var v20 int8 = 42
    var v21 uint8 = v18
    for fi22 := 0; fi22 < 1; fi22++ {
        v18 = v21
        for fi23 := 0; fi23 < 10; fi23++ {
            switch fi22 {
            case -1:
                v24 := []int{(fi22 % ((127) | 1)), (fi23 | (fi22 % ((127) | 1))), v5}
                _ = v24
            case 463:
                switch ("" + "a") {
                case "foo bar":
                    v20 += (^(v20))
                case "":
                    fmt.Println(v21, v18, v20)
                    fmt.Println(v18, v1, v20)
                default:
                    v25 := GenS2{F0: 628}
                    _ = v25
                }
            default:
                var v26 float64 = 1.5
                _ = v26
            }
            _ = fi23
        }
        fmt.Println(fi22, v2)
        _ = fi22
    }
    v27 := []int{int(v2)}
    v18 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
