package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0() uint16 {
    if false {
        var v6 bool = false
        _ = v6
    }
    if false {
        var v7 uint32 = 65535
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := []int{fi9}
            var v11 int = (fi8 >> ((18446744073709551615) & 63))
            if true {
                v12 := GenS1{F0: 42, F1: -9223372036854775808}
                v10 = append(v10, (v11 | (v11 % (((fi9 ^ (^(fi8)))) | 1))))
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        switch fi8 {
        case 42:
            var v13 int64 = 42
            _ = v13
        default:
            fmt.Println(fi8)
        }
        v14 := GenS1{F0: 65535, F1: 0}
        _ = fi8
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := []int{(fi15 | (fi15 >> ((65535) & 63))), (fi15 / ((9223372036854775807) | 1)), 127}
        var v17 float32 = 3.14159
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint64 = 986
    _ = v18
    return 432
}
func genFn1(p0 uint64, p1 uint, p2 uint32) int {
    v19 := GenS1{F0: p1, F1: -9223372036854775808}
    for fi20 := 0; fi20 < 15; fi20++ {
        if false {
            var v21 uint = p1
            switch 1 {
            case -1:
                p0 -= p0
            case 32768:
                var v22 uint64 = 119
                _ = v22
            }
            _ = v21
        } else {
            var v23 int = fi20
            _ = v23
        }
        if true {
            var v24 bool = (3.14159 >= 0.1)
            var v25 int8 = 0
            for fi26 := 0; fi26 < 13; fi26++ {
                v24 = true
                _ = fi26
            }
            _ = v24
            _ = v25
        }
        v27 := []int{(fi20 & -573), -1, fi20, fi20}
        _ = fi20
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return int(p1)
}

func main() {
    var v1 uint16 = 384
    var v2 uint16 = (v1 & 42)
    v3 := []int{-1, -380, 2147483647}
    v4 := make(map[uint32]string)
    v5 := []uint32{0, 948, 0}
    var v28 int16 = -777
    if true {
        v29 := GenS1{F0: 887, F1: 0}
        if len(v5) > 0 {
            v5[0] = 65535
        }
        var v30 int = -45
        _ = v29
        _ = v30
    }
    v5 = append(v5, 595)
    if (!true) {
        var v31 uint16 = v1
        _ = v31
    }
    v32 := GenS1{F0: 42, F1: 205}
    v33 := v4[340]
    var v34 int32 = -1
    _ = len(v4)
    v33 = "\n\t"
    v1 += (v1 - v1)
    fmt.Println(v33, v33)
    switch ((v33 + v33) + v33) {
    case "foo bar":
        v4[641] = "\n\t"
    case "":
        var v35 int16 = (v28 + v28)
        _ = v35
    default:
        if (false || (false && (false || true))) {
            v28 = (v28 / (((^(v28))) | 1))
            fmt.Println(v28, v1)
            var v36 bool = (1e10 < 3.14159)
            _ = v36
        } else {
            v34 -= (v34 + v34)
        }
    }
    v34 = v34
    v33 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
