package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16) uint32 {
    var v5 bool = (-242 < 9223372036854775807)
    if (v5 || v5) {
        var v6 string = ("a" + "a")
        v7 := []int{-1, 42, 646, 42}
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint32 = 2147483647
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        v5 = false
    }
    if (v5 && (true || v5)) {
        v10 := []int{-489, 402}
        _ = v10
    }
    var v11 uint16 = p0
    _ = p0
    _ = v11
    _ = v5
    return 0
}
func genFn1() int16 {
    switch "foo bar" {
    case "foo bar":
        v12 := GenS1{F0: 42, F1: true}
        _ = v12
    case "\n\t":
        var v13 uint8 = 42
        _ = v13
    case "":
        var v14 uint64 = 42
        _ = v14
    default:
        var v15 float32 = float32(float64(0.0))
        _ = v15
    }
    var v16 bool = (true && (false || true))
    for fi17 := 0; fi17 < 15; fi17++ {
        v16 = (1 > 0)
        v18 := GenS1{F0: 1, F1: v16}
        _ = fi17
        _ = v18
    }
    _ = v16
    return -1
}

func main() {
    var v1 bool = (-1 == -1)
    var v2 int16 = 0
    var v3 int8 = 111
    v4 := make(map[int]string)
    v19 := GenS2{F0: 2147483647}
    fmt.Println(v2, v3, v1)
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            v3 = (v3 / ((v3) | 1))
            v4[fi20] = "a"
            _ = len(v4)
            _ = fi21
        }
        _ = fi20
    }
    v3 = v3
    v2 = (v2 * (v2 << ((18446744073709551615) & 63)))
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS2{F0: 2147483647}
        var v24 uint8 = 127
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint = 4294967295
    v26 := v4[128]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
