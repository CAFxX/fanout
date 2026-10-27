package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0(p0 int64, p1 uint) int32 {
    for fi4 := 0; fi4 < 1; fi4++ {
        if false {
            var v5 uint8 = 42
            fmt.Println(p1, p1, v5)
            if false {
                fmt.Println(v5, fi4)
                var v6 int32 = -2147483648
                var v7 bool = false
                _ = v6
                _ = v7
            }
            _ = v5
        }
        _ = fi4
    }
    v8 := GenS1{F0: 4294967295, F1: 112, F2: 125}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int = fi9
        var v11 uint = (p1 << ((p1) & 63))
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{-184}
    switch "\n\t" {
    case "\n\t":
        var v13 uint8 = 0
        _ = v13
    case "foo bar":
        var v14 int64 = (p0 >> ((p1) & 63))
        _ = v14
    default:
        v15 := GenS1{F0: 0, F1: 1, F2: 1}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    return -713
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -315
    var v3 int64 = 66
    var v16 float32 = float32(float64((0.0 - (0.1 * (0.1 * (1e10 * 3.14159))))))
    v3 += 9223372036854775807
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := []int{fi17, 4294967295, fi17}
        _ = fi17
        _ = v18
    }
    switch "hello" {
    case "hello":
        if false {
            if (!false) {
                if true {
                    if ((true && true) || false) {
                        fmt.Println(v16, v3)
                        fmt.Println(v2)
                    }
                } else {
                    v19 := []int{382, 42, 127}
                    _ = v19
                }
                v20 := GenS2{F0: (^(v2)), F1: 2}
                v2 -= (^(v3))
                _ = v20
            }
            var v21 int16 = 42
            v22 := []int{-1}
            _ = v21
            _ = v22
        }
    case "foo bar":
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 int32 = 0
            fmt.Println(v24, v2)
            var v25 float32 = float32(3.14159)
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v3 -= (v3 - (v3 | v3))
    if (true || ((2147483647 >= 1) || (!(288 == 701)))) {
        if false {
            var v26 int = 490
            _ = v26
        } else {
            if (!false) {
                v27 := v1[("a" + "\n\t")]
                v2 = v2
                _ = v27
            }
            v3 -= (v2 / (((-(v2))) | 1))
            v16 -= v16
        }
    }
    v16 -= v16
    v16 -= -2.25
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := []int{int(v2), (fi28 << ((588) & 63))}
        v30 := GenS3{F0: fi28, F1: (v3 & 1)}
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := []int{725, -228, 0, 42}
    if (674 >= 42) {
        v3 -= v3
        v2 -= (v2 | -1)
        var v32 uint8 = 42
        _ = v32
    } else {
        v33 := GenS4{F0: 32767, F1: 1}
        if false {
            if (339 >= -956) {
                if true {
                    delete(v1, "hello")
                    var v34 uint = 2
                    _ = v34
                } else {
                    var v35 uint64 = 549
                    v16 = (3.14159 / v16)
                    _ = v35
                }
            }
            if true {
                var v36 int = 2147483647
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                _ = v36
            }
        } else {
            fmt.Println(v2)
            fmt.Println(v2, v16, v16)
        }
        v37 := GenS5{F0: -2147483648, F1: -1, F2: 42}
        _ = v33
        _ = v37
    }
    if false {
        v3 += (v2 * v2)
        v3 = (v2 % (((v3 ^ (^(v3)))) | 1))
        var v38 uint = 497
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
}
