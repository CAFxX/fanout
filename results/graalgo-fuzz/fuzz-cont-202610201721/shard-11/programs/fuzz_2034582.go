package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint64 = 1
    var v2 uint32 = 2147483648
    var v3 float32 = 3.14159
    var v4 string = "\n\t"
    var v5 bool = (355 != 1)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float64 = 3.14159
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: (fi8 | (fi8 * fi6)), F1: (-9223372036854775808 != -9223372036854775808), F2: 2}
            if true {
                fmt.Println(v2)
                var v10 uint8 = 84
                if true {
                    var v11 int = fi6
                    _ = v11
                } else {
                    v12 := []int{-9223372036854775808, (fi8 / ((fi6) | 1)), -303, (fi6 ^ 128)}
                    v2 = (v2 & 324)
                    _ = v12
                }
                _ = v10
            }
            if v5 {
                v5 = ((v7 * float64(v3)) >= (v7 * (v7 / v7)))
                var v13 int = 2147483647
                if (!(v3 < (0.0 + v3))) {
                    var v14 int16 = 42
                    var v15 string = "hello"
                    var v16 float32 = v3
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    if (false && (82 != 42)) {
                        v17 := []int{(fi8 ^ (fi8 & fi6))}
                        _ = v17
                    } else {
                        if v5 {
                            var v18 bool = (true && false)
                            fmt.Println(v18, v13)
                            v19 := GenS1{F0: (fi6 << ((0) & 63)), F1: (v5 && v18), F2: 0}
                            _ = v18
                            _ = v19
                        } else {
                            v4 = v4
                            var v20 int = (fi6 % ((fi8) | 1))
                            v21 := []int{(fi6 | fi8), v20}
                            v22 := GenS1{F0: v20, F1: false, F2: 42}
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{fi8, fi6, fi6, (v13 ^ fi6)}
                        v1 = (v1 >> ((185) & 63))
                        _ = v23
                    }
                    var v24 float32 = 3.14159
                    v25 := []int{fi6}
                    _ = v24
                    _ = v25
                }
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v26 int8 = -19
    var v27 float32 = v3
    var v28 int64 = -196
    v26 -= v26
    v29 := []int{}
    v30 := GenS1{F0: 9223372036854775807, F1: (v28 <= 42), F2: 152}
    for fi31 := 0; fi31 < 14; fi31++ {
        v29 = append(v29, 652)
        _ = len(v29)
        _ = fi31
    }
    var v32 int16 = 301
    var v33 uint = 2
    v34 := GenS1{F0: 128, F1: (!true), F2: 173}
    v29 = append(v29, 2147483647)
    v35 := GenS2{F0: 128}
    var v36 int64 = v28
    switch -9223372036854775808 {
    case 127:
        if len(v29) > 0 {
            v29[0] = 42
        }
    case -9223372036854775808:
        v29 = append(v29, -298)
    default:
        v37 := GenS1{F0: 646, F1: (v5 || (("hello" < v4) || (529 >= 1))), F2: 1}
        _ = v37
    }
    v5 = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
