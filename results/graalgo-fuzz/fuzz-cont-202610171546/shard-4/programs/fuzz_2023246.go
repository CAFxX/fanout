package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 uint) bool {
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7, p1)
        var v8 uint32 = 2147483648
        v9 := []int{p0}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = -470
    p0 -= 127
    switch -1 {
    case -9223372036854775808:
        p1 = (p1 >> ((p1) & 63))
    case 9223372036854775807:
        switch p0 {
        case 436:
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := []int{127, p0, fi11}
                for fi13 := 0; fi13 < 14; fi13++ {
                    switch p0 {
                    case -2147483648:
                        var v14 uint = (p1 * (p1 ^ (p1 & 622)))
                        _ = v14
                    case 9223372036854775807:
                        var v15 int8 = -1
                        _ = v15
                    default:
                        var v16 uint32 = 2147483647
                        _ = v16
                    }
                    switch ("" + "hello") {
                    case "a":
                        var v17 int = fi13
                        _ = v17
                    case "":
                        var v18 int64 = 0
                        _ = v18
                    case "foo bar":
                        p1 -= p1
                    }
                    var v19 int = fi13
                    _ = fi13
                    _ = v19
                }
                _ = fi11
                _ = v12
            }
        case -935:
            p1 += (^(p1))
        default:
            if true {
                v20 := GenS1{F0: 803, F1: (p1 % (((p1 % (((p1 & (p1 + p1))) | 1))) | 1))}
                _ = v20
            }
        }
    default:
        v21 := []int{}
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 bool = (1 <= 65535)
        for fi24 := 0; fi24 < 17; fi24++ {
            v10 += v10
            var v25 int8 = 0
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v10
    return true
}
func genFn1(p0 int16, p1 int16) int32 {
    fmt.Println(p0)
    v26 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v26
    return 858
}

func main() {
    var v1 int16 = -1
    v2 := []string{"\n\t"}
    var v3 int16 = (v1 >> ((255) & 63))
    v4 := []int{32767, 1}
    v5 := make(map[int]int)
    var v6 uint8 = 127
    v27 := GenS2{F0: 255}
    var v28 uint = 1
    v1 = (^(v3))
    if true {
        v1 -= 0
        fmt.Println(v3, v3, v1)
    }
    var v29 int32 = -310
    v29 = genFn1((v3 ^ (v1 << (((^(v28))) & 63))), 21)
    fmt.Println(v1, v6, v6)
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    fmt.Println(v3, v6, v1)
    v28 -= (^(v28))
    if false {
        var v30 int8 = 100
        switch "\n\t" {
        case "a":
            fmt.Println(v1)
        }
        _ = len(v2)
        _ = v30
    }
    v1 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
