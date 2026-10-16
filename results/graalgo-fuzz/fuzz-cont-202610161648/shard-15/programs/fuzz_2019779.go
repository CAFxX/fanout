package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 int64) string {
    if (false && (1 < 255)) {
        p0 = p0
        var v4 bool = ((!true) || (0.1 > 0.1))
        _ = v4
    }
    if (((42 == 184) || ((-62 < 1) && true)) || (1 > 42)) {
        switch "hello" {
        case "hello":
            if (1 != 451) {
                switch "\n\t" {
                case "a":
                    for fi5 := 0; fi5 < 2; fi5++ {
                        var v6 uint8 = 68
                        fmt.Println(v6, p0)
                        _ = fi5
                        _ = v6
                    }
                case "":
                    switch "\n\t" {
                    case "":
                        if true {
                            var v7 float64 = 0.0
                            var v8 int32 = -214
                            for fi9 := 0; fi9 < 2; fi9++ {
                                p0 += p0
                                _ = fi9
                            }
                            _ = v7
                            _ = v8
                        } else {
                            var v10 uint64 = 419
                            v11 := []int{}
                            v12 := GenS1{F0: true, F1: p0, F2: 42}
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    default:
                        v13 := GenS1{F0: true, F1: p0, F2: 42}
                        _ = v13
                    }
                default:
                    v14 := []int{-2147483648, 887}
                    _ = v14
                }
                var v15 int32 = 0
                _ = v15
            }
        }
        p0 += (p0 % ((p0) | 1))
        v16 := GenS2{F0: 682, F1: "foo bar"}
        _ = v16
    } else {
        v17 := []int{-190, 32768}
        _ = v17
    }
    _ = p0
    return "foo bar"
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -1
    var v3 float64 = -2.25
    v2 += (v2 & -1)
    if true {
        if (1 <= 127) {
            v1 -= v1
            if true {
                var v18 uint8 = 2
                v1 = (v1 - (v1 ^ (v1 * v1)))
                _ = v18
            }
        }
    }
    v3 = ((3.14159 - v3) + (1e10 * (((v3 + v3) * (v3 - v3)) + ((v3 - float64((float32((0.0 * (0.1 * (0.0 + ((v3 - 1e10) / v3))))) * (float32(v3) * (1.5 - (0.0 / (float32(-2.25) + 0.1))))))) - -2.25))))
    var v19 int8 = 102
    var v20 uint32 = 529
    var v21 bool = true
    v1 += v1
    switch 32768 {
    case 1:
        var v22 int32 = (v2 - v2)
        _ = v22
    case -109:
        v21 = true
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int = -968
        v25 := []int{(fi23 | (v24 << ((18446744073709551615) & 63))), (fi23 ^ v24), fi23, (fi23 & 696)}
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v20, fi26)
        for fi27 := 0; fi27 < 17; fi27++ {
            if v21 {
                var v28 uint8 = 255
                var v29 int16 = 1
                v21 = v21
                _ = v28
                _ = v29
            }
            _ = fi27
        }
        _ = fi26
    }
    v30 := []int{2147483647}
    fmt.Println(v3)
    v1 -= (v1 % (((^(v1))) | 1))
    _ = len(v30)
    if (v19 <= (v19 >> ((2) & 63))) {
        var v31 float32 = -2.25
        var v32 uint64 = v1
        _ = v31
        _ = v32
    } else {
        var v33 int16 = -32768
        var v34 int16 = v33
        var v35 int8 = v19
        _ = v33
        _ = v34
        _ = v35
    }
    switch -229 {
    case -529:
        switch "foo bar" {
        case "hello":
            for fi36 := 0; fi36 < 7; fi36++ {
                v21 = (-1 > -356)
                v20 = (v20 ^ 1)
                v1 = v1
                _ = fi36
            }
        }
    default:
        var v37 uint64 = v1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
}
