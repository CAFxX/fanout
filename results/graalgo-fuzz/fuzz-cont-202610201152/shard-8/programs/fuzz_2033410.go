package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 uint8; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 string = "a"
    v3 := make(map[string]uint32)
    var v4 float64 = -2.25
    v2 = v2
    if (-2.25 < (1e10 - 0.0)) {
        v5 := GenS1{F0: 0, F1: 2147483647}
        var v6 int64 = -1
        _ = v5
        _ = v6
    }
    var v7 uint = 213
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 1000, F1: (^(fi8))}
        _ = fi8
        _ = v9
    }
    var v10 float64 = 0.1
    delete(v1, "a")
    v2 = v2
    if (2 <= 2) {
        v11 := GenS1{F0: -32768, F1: 0}
        var v12 bool = false
        _ = v11
        _ = v12
    } else {
        var v13 uint8 = 203
        var v14 int64 = -241
        _ = v13
        _ = v14
    }
    var v15 int8 = -128
    if true {
        switch -1 {
        case 32767:
            if true {
                _ = len(v3)
                if true {
                    v16 := GenS2{F0: (v15 | -1), F1: 128, F2: (v2 + v2)}
                    _ = v16
                } else {
                    v17 := GenS2{F0: v15, F1: 128, F2: (v2 + "")}
                    if (476 <= 32767) {
                        var v18 int32 = -972
                        v19 := []int{-268, 2147483647, -9223372036854775808, 778}
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        if (262 <= 492) {
                            var v21 uint64 = 0
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 15; fi22++ {
                                v15 = int8(v7)
                                v2 = v2
                                _ = fi22
                            }
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 uint32 = 304
                                v7 = (^(v7))
                                v25 := GenS1{F0: 236, F1: fi20}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                        var v26 int = (fi20 + fi20)
                        v2 = v2
                        _ = fi20
                        _ = v26
                    }
                    _ = v17
                }
                var v27 float64 = v4
                _ = v27
            } else {
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch (fi28 + -1) {
                    case 163:
                        var v29 float32 = float32(((-2.25 * v10) * (v4 - (v4 * v4))))
                        _ = v29
                    case -2147483648:
                        for fi30 := 0; fi30 < 17; fi30++ {
                            v2 = v2
                            _ = fi30
                        }
                    case 127:
                        _ = len(v3)
                    default:
                        v2 = (v2 + (("foo bar" + "foo bar") + ("" + ("\n\t" + v2))))
                    }
                    var v31 uint = 255
                    v4 = 1.5
                    _ = fi28
                    _ = v31
                }
            }
        case 0:
            v32 := []int{42, 128, 127}
            _ = v32
        case -1:
            var v33 uint8 = 128
            _ = v33
        }
    } else {
        fmt.Println(v10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
