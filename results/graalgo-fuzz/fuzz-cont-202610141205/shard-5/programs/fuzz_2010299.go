package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 uint32 = 1
    if (4294967295 >= 451) {
        var v6 uint = 255
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: (v5 % ((2147483648) | 1)), F1: 42, F2: v5}
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        v9 := GenS1{F0: (v5 / ((v5) | 1)), F1: v5, F2: 1}
        _ = v9
    }
    fmt.Println(v5)
    v10 := GenS2{F0: 386, F1: -9223372036854775808, F2: -1}
    v5 -= v5
    fmt.Println(v5)
    if (!(((1e10 + 0.0) - (1e10 / ((1.5 + float64((float32((0.1 + 0.0)) - 1.5))) - (1.5 * 3.14159)))) == 0.0)) {
        if (!true) {
            var v11 bool = false
            var v12 uint32 = v5
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(v5, v11)
                switch "foo bar" {
                case "a":
                    if (!((fi13 < (fi13 | fi13)) && false)) {
                        var v14 uint16 = 2
                        _ = v14
                    }
                case "":
                    v5 -= (v5 * v12)
                case "foo bar":
                    var v15 bool = v11
                    _ = v15
                }
                v16 := GenS3{F0: v11, F1: 42, F2: 127}
                _ = fi13
                _ = v16
            }
            _ = v11
            _ = v12
        }
    }
    v17 := []int{-2147483648}
    v18 := GenS3{F0: (true || (-676 <= 699)), F1: 1, F2: 1}
    v5 += v5
    for fi19 := 0; fi19 < 16; fi19++ {
        v5 += v5
        v2[32768] = -774
        v17 = append(v17, (fi19 / (((-(fi19))) | 1)))
        _ = fi19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v5 = v5
        v5 = 764
        _ = fi20
    }
    var v21 int32 = -505
    if (1.5 >= ((((1.5 - 1.5) * 0.0) * 0.0) + (-2.25 - 1e10))) {
        fmt.Println(v5, v5)
        v3["\n\t"] = false
        for fi22 := 0; fi22 < 2; fi22++ {
            if (true && true) {
                var v23 uint64 = 590
                if (v5 >= 65535) {
                    v5 = (v5 | v5)
                    var v24 bool = true
                    var v25 bool = v24
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint64 = (v23 | (v23 << ((255) & 63)))
                    v23 -= v26
                    if true {
                        var v27 float64 = float64(0.0)
                        v23 = (v26 % ((v26) | 1))
                        _ = v27
                    } else {
                        v28 := GenS3{F0: false, F1: 1, F2: 127}
                        if len(v17) > 0 {
                            v17[0] = fi22
                        }
                        fmt.Println(v26)
                        _ = v28
                    }
                    _ = v26
                }
                _ = v23
            } else {
                fmt.Println(fi22, fi22)
            }
            v29 := GenS2{F0: 588, F1: 0, F2: (v21 & v21)}
            var v30 float32 = 1e10
            _ = fi22
            _ = v29
            _ = v30
        }
    }
    var v31 int = 9223372036854775807
    var v32 float32 = 1e10
    if len(v17) > 0 {
        v17[0] = 9223372036854775807
    }
    if len(v17) > 0 {
        v17[0] = (v31 + 42)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
