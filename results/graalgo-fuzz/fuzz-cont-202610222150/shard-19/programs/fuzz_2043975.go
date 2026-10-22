package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
func genFn0() int16 {
    var v5 float32 = 1.5
    v6 := GenS1{F0: -9223372036854775808, F1: true, F2: "\n\t"}
    fmt.Println(v5)
    var v7 uint64 = 0
    _ = v5
    _ = v6
    _ = v7
    return 1000
}
func genFn1(p0 uint16) bool {
    var v8 int16 = -1
    if (42 < 42) {
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                if (1 < 3) {
                    v10 := GenS2{F0: -85}
                    v8 = (v8 % (((v8 * (-(v8)))) | 1))
                    _ = v10
                } else {
                    var v11 bool = ((!(true && true)) || true)
                    v11 = v11
                    var v12 bool = v11
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: -869}
                p0 += (p0 >> ((4294967295) & 63))
                _ = v13
            }
            v8 -= 1
            p0 = p0
            _ = fi9
        }
    } else {
        v8 = (v8 + 1000)
        switch 1 {
        case 745:
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: 42, F1: (((p0 & 2) >= p0) && false), F2: "hello"}
                for fi16 := 0; fi16 < 17; fi16++ {
                    p0 += p0
                    v8 += genFn0()
                    v17 := []int{fi14, -254, fi16}
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        case 128:
            v8 = (v8 << ((976) & 63))
        default:
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 1, F1: false, F2: "foo bar"}
                _ = fi18
                _ = v19
            }
        }
    }
    var v20 uint64 = 1
    _ = p0
    _ = v20
    _ = v8
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 int = -1
    var v3 uint8 = 62
    var v4 float32 = (3.14159 - (0.0 - ((1.5 + 3.14159) - (float32((3.14159 * (0.0 + (1e10 + -2.25)))) / (float32(3.14159) - 0.1)))))
    fmt.Println(v3)
    v21 := GenS1{F0: -975, F1: true, F2: "a"}
    for fi22 := 0; fi22 < 19; fi22++ {
        v2 -= fi22
        var v23 float64 = (1.5 + (-2.25 - 1e10))
        _ = fi22
        _ = v23
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        v3 += v3
    case "hello":
        var v24 uint16 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        fmt.Println(v4)
        var v26 uint = 434
        v2 -= v2
        _ = fi25
        _ = v26
    }
    delete(v1, (-(v2)))
    v27 := GenS3{F0: -1}
    v2 = 9223372036854775807
    var v28 int32 = 743
    v4 = 3.14159
    v29 := []int{9223372036854775807}
    v3 = (v3 % (((v3 % (((v3 << ((18446744073709551615) & 63))) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
