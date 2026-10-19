package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
func genFn0() int {
    var v6 int16 = -32768
    v6 = (v6 / (((v6 * (^(v6)))) | 1))
    v7 := []int{-1, 838, -1}
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() uint64 {
    switch 433 {
    case 4294967295:
        v8 := GenS1{F0: 0, F1: -199}
        _ = v8
    case 9223372036854775807:
        var v9 int = 32767
        _ = v9
    default:
        v10 := GenS1{F0: 42, F1: 32767}
        _ = v10
    }
    var v11 int16 = 1000
    _ = v11
    return 2
}

func main() {
    var v1 int8 = 100
    var v2 int8 = (v1 << ((18446744073709551615) & 63))
    v3 := make(map[int]string)
    var v4 uint64 = 0
    var v5 uint = 18446744073709551615
    fmt.Println(v2)
    _ = len(v3)
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint = (v5 << ((2) & 63))
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 681, F1: -32768}
            if false {
                var v16 uint32 = 170
                v17 := GenS1{F0: 365, F1: -1}
                _ = v16
                _ = v17
            } else {
                var v18 int64 = 1
                v13 -= (v5 | 382)
                fmt.Println(v13)
                _ = v18
            }
            var v19 uint64 = (v4 ^ 13)
            v2 -= v2
            v4 += uint64(v5)
            _ = fi14
            _ = v15
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v2, v13)
            _ = fi20
        }
        _ = fi12
        _ = v13
    }
    if true {
        delete(v3, 32767)
    } else {
        switch "a" {
        case "a":
            v2 = v2
        case "hello":
            v21 := GenS1{F0: 68, F1: -303}
            _ = v21
        }
        v2 += (v2 & (v1 - -128))
    }
    if ((v1 < 100) && (("hello" + "hello") != "a")) {
        var v22 int = genFn0()
        fmt.Println(v1, v2)
        _ = v22
    } else {
        if (46 < 0) {
            var v23 float64 = (1e10 - 0.1)
            for fi24 := 0; fi24 < 3; fi24++ {
                v3[genFn0()] = ""
                var v25 uint16 = 1
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    }
    v4 = (v4 - v4)
    for fi26 := 0; fi26 < 4; fi26++ {
        if (false || ((true && (((100 >= v1) || true) || false)) || true)) {
            v27 := []int{fi26, fi26, genFn0(), genFn0()}
            switch fi26 {
            case 2147483647:
                fmt.Println(v5, v5)
            }
            v28 := GenS2{F0: -1, F1: 889, F2: (^(fi26))}
            _ = v27
            _ = v28
        }
        var v29 uint32 = 2
        _ = fi26
        _ = v29
    }
    v30 := []int{0, 128, 1}
    v5 = (^(v5))
    v4 -= (v4 * v4)
    var v31 float64 = (float64(1.5) / 1.5)
    fmt.Println(v1)
    v32 := GenS1{F0: -1, F1: -47}
    var v33 int16 = 1000
    fmt.Println(v5)
    var v34 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
