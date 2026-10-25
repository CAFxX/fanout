package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 int8, p1 uint64) string {
    switch ("hello" + (("\n\t" + "a") + "")) {
    case "\n\t":
        v7 := []int{-2147483648}
        _ = v7
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            if false {
                fmt.Println(p0, p1, p0)
            }
            for fi9 := 0; fi9 < 20; fi9++ {
                if (1 > -2147483648) {
                    p1 -= (p1 | 403)
                    v10 := GenS1{F0: 2147483647, F1: (p1 * p1)}
                    var v11 int = fi8
                    _ = v10
                    _ = v11
                } else {
                    if true {
                        p0 += p0
                        switch "\n\t" {
                        case "foo bar":
                            var v12 float64 = 3.14159
                            _ = v12
                        case "\n\t":
                            v13 := []int{fi9, (fi9 * fi8), -1, int(p1)}
                            _ = v13
                        case "hello":
                            var v14 int32 = 1
                            _ = v14
                        default:
                            p0 = p0
                        }
                    }
                    v15 := []int{(fi8 % (((fi8 >> ((42) & 63))) | 1)), (fi9 >> ((163) & 63)), fi8, 32767}
                    var v16 string = ""
                    p0 += 100
                    _ = v15
                    _ = v16
                }
                var v17 float64 = 3.14159
                _ = fi9
                _ = v17
            }
            _ = fi8
        }
    }
    var v18 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v18
    return "hello"
}

func main() {
    v1 := []uint32{2147483648, 583}
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    v4 := make(map[uint32]uint32)
    var v5 float64 = float64((float32((1.5 / (1e10 * float64(float32(-2.25))))) - 0.1))
    var v6 int = 929
    v19 := GenS1{F0: 2147483647, F1: v3}
    switch v6 {
    case 42:
        var v20 uint = 0
        _ = v20
    }
    if (!(656 > -1)) {
        if false {
            v21 := GenS1{F0: -760, F1: v3}
            _ = v21
        } else {
            switch "a" {
            case "a":
                v6 -= 32768
            case "\n\t":
                if ((^(v3)) > 0) {
                    if true {
                        var v22 float64 = 3.14159
                        _ = v22
                    } else {
                        if (0.1 <= 1.5) {
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            v3 -= v3
                            v3 -= v3
                        } else {
                            v5 = (v5 - 1.5)
                        }
                        var v23 uint = 4294967295
                        _ = v23
                    }
                }
            case "":
                switch 32768 {
                case 42:
                    var v24 int16 = -977
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 5; fi25++ {
                v5 -= (v5 * (float64((1e10 + (-2.25 - 3.14159))) * v5))
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(fi26, fi25, v3)
                    var v27 float32 = (3.14159 + 0.0)
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            var v28 uint64 = (v3 - 346)
            _ = v28
        }
        v5 = float64(1e10)
    } else {
        var v29 int32 = 0
        fmt.Println(v29, v5, v29)
        var v30 int32 = 2147483647
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: "hello", F1: 178}
    v32 := GenS1{F0: -2147483648, F1: v3}
    fmt.Println(v3, v5, v6)
    v6 = 4294967295
    var v33 int64 = -9223372036854775808
    v1 = append(v1, 363)
    var v34 uint = 579
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
