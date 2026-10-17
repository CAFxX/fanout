package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
func genFn0(p0 uint16) string {
    p0 = 360
    var v7 uint = 149
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8)
        _ = fi8
    }
    v7 += 42
    if true {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint32 = 0
            v7 -= (^(v7))
            _ = fi11
            _ = v12
        }
        v13 := []int{32768, -989, 9223372036854775807, -2147483648}
        v7 = (v7 - 65535)
        _ = v13
    }
    _ = p0
    _ = v7
    return ("" + "hello")
}
func genFn1(p0 int8, p1 uint64, p2 string) uint8 {
    p0 = p0
    p1 = (p1 / (((p1 / ((1) | 1))) | 1))
    switch 827 {
    case 32767:
        p0 = -128
    default:
        v14 := []int{}
        _ = v14
    }
    p2 = "a"
    var v15 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 128
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = 32767
    var v3 float32 = (1.5 + (float32(-2.25) * (((v1 - v1) - v1) * float32(-2.25))))
    var v4 uint = 0
    v5 := []string{"foo bar", ("a" + "a"), "a"}
    v6 := make(map[string]uint32)
    v1 = v1
    var v16 string = ("\n\t" + "")
    v4 += 969
    var v17 float32 = (1e10 / (v1 - v1))
    fmt.Println(v16)
    v18 := GenS1{F0: 32767, F1: -1}
    v5 = append(v5, v16)
    fmt.Println(v16, v1)
    v19 := GenS2{F0: (!((0.0 / 1.0) < 1.5)), F1: (v4 & 2), F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
