package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 uint) uint8 {
    v7 := GenS1{F0: 128, F1: 731}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int = fi8
        if ((true && true) || false) {
            p0 = p0
            var v10 uint16 = 255
            _ = v10
        } else {
            p0 += p0
            v11 := GenS2{F0: "hello", F1: 804}
            fmt.Println(p0)
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            if (!(!(true && false))) {
                var v13 int16 = -32768
                v14 := []int{fi8, 46, 9223372036854775807}
                _ = v13
                _ = v14
            } else {
                fmt.Println(fi12)
                var v15 int8 = -128
                _ = v15
            }
            if true {
                v16 := []int{v9}
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = v16
            }
            p0 += 4294967295
            v17 := []int{}
            _ = fi12
            _ = v17
        }
        _ = fi8
        _ = v9
    }
    p0 += 4294967295
    p0 -= p0
    switch 586 {
    case 128:
        v18 := []int{2147483647, 570}
        _ = v18
    case 42:
        if (p0 < 65535) {
            p0 = p0
        } else {
            if ((32767 < -14) && true) {
                v19 := []int{-9223372036854775808, 127}
                v20 := GenS1{F0: 66, F1: -743}
                fmt.Println(p0)
                _ = v19
                _ = v20
            }
        }
    }
    _ = p0
    _ = v7
    return 127
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    var v3 int64 = 0
    v4 := []uint32{}
    v5 := make(map[int]string)
    var v6 float32 = 0.0
    var v21 bool = (42 != 1)
    var v22 float32 = v6
    if (32767 != 2147483647) {
        var v23 uint32 = 652
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v21 = (-128 == 127)
        if (true || v21) {
            var v25 float64 = 0.0
            v26 := GenS2{F0: "a", F1: 0}
            for fi27 := 0; fi27 < 10; fi27++ {
                v4 = append(v4, 31)
                v28 := GenS2{F0: "hello", F1: -2147483648}
                fmt.Println(v21, fi24)
                _ = fi27
                _ = v28
            }
            _ = v25
            _ = v26
        }
        fmt.Println(v22, v3, v3)
        _ = fi24
    }
    switch -650 {
    case -633:
        if ("a" <= "foo bar") {
            v29 := GenS2{F0: "", F1: 55}
            _ = v29
        }
    case -9223372036854775808:
        var v30 uint32 = 2147483648
        _ = v30
    }
    v3 += (v3 + v3)
    var v31 int = -9223372036854775808
    fmt.Println(v31, v6, v6)
    for fi32 := 0; fi32 < 6; fi32++ {
        if v21 {
            fmt.Println(v3, v3, v31)
        }
        _ = fi32
    }
    v6 -= (((v22 + (v22 + v6)) - v22) * float32(1e10))
    var v33 int32 = 798
    switch "hello" {
    case "hello":
        v3 = -9223372036854775808
    default:
        if true {
            v6 += ((((v22 / (v6 / ((v22 * 3.14159) / (v6 - float32(1e10))))) - v22) / v6) / v22)
        }
    }
    v21 = ((float64(3.14159) - (1e10 + float64(-2.25))) >= -2.25)
    v3 -= (v3 * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
