package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 int32) uint8 {
    v5 := GenS1{F0: 42, F1: 9, F2: 65535}
    v6 := []int{42, 32768, -929}
    v7 := GenS1{F0: 4294967295, F1: -1, F2: 65535}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 39
}
func genFn1(p0 int8) int8 {
    fmt.Println(p0)
    v8 := []int{-228, -942, -9223372036854775808, 0}
    _ = p0
    _ = v8
    return p0
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 int8 = 0
    v1 += float32((((float64(3.14159) / 1.0) + (3.14159 / 3.14159)) + -2.25))
    if (-2.25 >= v1) {
        v4 += (^(v4))
        if (v1 >= v1) {
            v1 += v1
        }
    }
    var v9 int64 = -1
    var v10 uint16 = 104
    v1 += (float32(0.0) * 0.1)
    fmt.Println(v10)
    v1 = 1e10
    var v11 int = 32767
    switch "hello" {
    case "a":
        v12 := []int{(v11 >> ((255) & 63)), 0}
        _ = v12
    case "foo bar":
        switch (v11 + (v11 ^ 931)) {
        case -523:
            switch v11 {
            case -471:
                v1 += (v1 - 1.5)
            case -656:
                if false {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v15 := []int{(-(v11))}
                            _ = len(v15)
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                    v16 := []int{v11, 356, (v11 - (-(v11)))}
                    _ = v16
                } else {
                    if (false && ("\n\t" == "a")) {
                        v17 := []int{v11}
                        fmt.Println(v4, v1, v9)
                        _ = v17
                    } else {
                        v18 := []int{9223372036854775807}
                        v9 -= (v9 & -9223372036854775808)
                        _ = v18
                    }
                    v19 := GenS1{F0: 42, F1: 100, F2: 477}
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 int64 = (v9 >> ((uint(v10)) & 63))
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            case 1:
                v9 = v9
            }
        }
    case "\n\t":
        if (!((-110 >= v11) && true)) {
            var v22 string = "\n\t"
            _ = v22
        } else {
            for fi23 := 0; fi23 < 2; fi23++ {
                if (false && (false || true)) {
                    switch fi23 {
                    case -572:
                        var v24 uint = 361
                        _ = v24
                    case -293:
                        fmt.Println(v1)
                    default:
                        var v25 uint8 = 2
                        _ = v25
                    }
                    if ("foo bar" < (("a" + "hello") + "hello")) {
                        var v26 int = v11
                        v4 -= v4
                        for fi27 := 0; fi27 < 13; fi27++ {
                            v28 := GenS2{F0: v10, F1: 1000}
                            var v29 int8 = v4
                            _ = len(v3)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = v26
                    } else {
                        var v30 int64 = (v9 / (((^(v9))) | 1))
                        v30 = 42
                        v31 := GenS2{F0: 255, F1: -32768}
                        _ = v30
                        _ = v31
                    }
                    v10 = (v10 | v10)
                }
                var v32 int = (fi23 + fi23)
                v33 := []int{v32, fi23}
                _ = fi23
                _ = v32
                _ = v33
            }
            v4 = 127
        }
    }
    v34 := GenS2{F0: v10, F1: 32767}
    v35 := []int{(v11 + v11), (v11 & v11)}
    fmt.Println(v9, v1, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v9
}
