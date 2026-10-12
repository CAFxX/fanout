package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
func genFn0(p0 bool) uint16 {
    p0 = (-1 > 0)
    v5 := GenS1{F0: 210, F1: 1}
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 2, F1: 962}
        switch (fi6 % (((fi6 ^ (fi6 * fi6))) | 1)) {
        case -936:
            switch "hello" {
            case "a":
                for fi8 := 0; fi8 < 7; fi8++ {
                    v9 := []int{(fi8 * (fi6 | fi8)), fi6}
                    _ = fi8
                    _ = v9
                }
            case "hello":
                if p0 {
                    v10 := []int{(^(fi6)), (fi6 >> ((18446744073709551615) & 63)), (fi6 % (((fi6 * (fi6 / ((1) | 1)))) | 1)), (fi6 ^ fi6)}
                    _ = v10
                }
            case "":
                v11 := GenS1{F0: 0, F1: 449}
                _ = v11
            default:
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := []int{fi6, 283, fi12, fi6}
                    _ = fi12
                    _ = v13
                }
            }
        }
        var v14 int64 = 42
        _ = fi6
        _ = v14
        _ = v7
    }
    v15 := GenS1{F0: 42, F1: 2147483647}
    p0 = true
    _ = p0
    _ = v15
    _ = v5
    return 345
}
func genFn1(p0 int64, p1 string) string {
    var v16 uint = 2
    v17 := GenS1{F0: 2147483647, F1: 2147483647}
    v18 := GenS1{F0: 42, F1: 0}
    var v19 string = (("\n\t" + p1) + p1)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return p1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    v3 := []string{}
    var v4 int64 = 57
    var v20 int32 = 2147483647
    v4 -= -751
    v21 := v1[(("" + "") + ("hello" + "a"))]
    v4 = (v4 + -1)
    var v22 uint8 = 127
    var v23 int64 = (-(v4))
    if v21 {
        var v24 uint64 = 1
        switch 0 {
        case 32768:
            var v25 uint64 = v24
            _ = v25
        case 896:
            v26 := GenS2{F0: 785, F1: (v24 ^ v24), F2: v20}
            _ = v26
        }
        v3 = append(v3, "hello")
        _ = v24
    } else {
        var v27 uint32 = 805
        v21 = false
        _ = v27
    }
    var v28 uint64 = 4294967295
    for fi29 := 0; fi29 < 11; fi29++ {
        v23 = 0
        var v30 int16 = v2
        _ = fi29
        _ = v30
    }
    var v31 uint8 = (^(v22))
    if v21 {
        v22 -= 42
    }
    v20 -= (v20 & v20)
    v32 := GenS1{F0: 1, F1: 42}
    var v33 uint64 = (^(v28))
    for fi34 := 0; fi34 < 18; fi34++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi34
    }
    if (!(0.1 == 0.0)) {
        v35 := GenS1{F0: 1, F1: 1}
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
