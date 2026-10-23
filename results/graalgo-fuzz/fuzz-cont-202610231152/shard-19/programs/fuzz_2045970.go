package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int16 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint }
func genFn0(p0 int8, p1 int64) int64 {
    if false {
        p0 = 77
        p1 = 286
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                p0 -= (p0 * (p0 - p0))
                p0 += p0
                p1 += (p1 / ((p1) | 1))
                _ = fi7
            }
            fmt.Println(p1, p1)
        } else {
            v8 := GenS1{F0: (p0 << ((430) & 63)), F1: -2147483648}
            if true {
                v9 := GenS1{F0: -1, F1: 42}
                p0 -= 90
                _ = v9
            } else {
                v10 := []int{4294967295, -597}
                var v11 uint8 = 1
                v12 := GenS1{F0: p0, F1: -141}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
        }
    } else {
        v13 := []int{}
        _ = v13
    }
    var v14 string = "hello"
    if false {
        var v15 string = ((("foo bar" + "foo bar") + (v14 + v14)) + "")
        _ = v15
    }
    switch 42 {
    case 128:
        v16 := GenS2{F0: "a", F1: v14, F2: 255}
        _ = v16
    case 32768:
        p1 = (p1 + p1)
    }
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 int8 = (^(v1))
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 string = "a"
    var v6 uint32 = 0
    fmt.Println(v5, v2)
    delete(v3, 128)
    var v17 uint = 635
    var v18 float64 = 3.14159
    var v19 string = "\n\t"
    v20 := GenS2{F0: "\n\t", F1: "\n\t", F2: 42}
    var v21 string = "a"
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            if false {
                v24 := GenS1{F0: (v2 + v1), F1: 0}
                _ = v24
            }
            v1 += (-(v1))
            if false {
                switch fi22 {
                case 32768:
                    v25 := []int{fi23}
                    _ = v25
                }
                switch fi23 {
                case 42:
                    var v26 bool = false
                    _ = v26
                case -654:
                    v1 += v2
                    fmt.Println(v2)
                    v21 = (v21 + "hello")
                    if true {
                        v2 += (v1 * (v2 - (v1 - v1)))
                        v27 := GenS1{F0: v1, F1: 42}
                        var v28 uint16 = 1
                        var v29 uint16 = (v28 & (^(v28)))
                        var v30 int32 = -1
                        if (9223372036854775807 != int64(v30)) {
                            v2 = 1
                            v30 = -636
                        } else {
                            _ = len(v3)
                            v18 = v18
                            fmt.Println(v21)
                        }
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    v17 += (v17 - v17)
                    fmt.Println(v17, v18, v18)
                    v31 := v4[v19]
                    _ = v31
                case 1:
                    if (-846 < -1) {
                        v19 = (((((v19 + "hello") + "hello") + "") + v19) + ((v19 + v19) + (v5 + v5)))
                    }
                default:
                    v21 = v19
                }
            }
            _ = fi23
        }
        v32 := GenS3{F0: 42, F1: 255, F2: 383}
        v6 -= v6
        _ = fi22
        _ = v32
    }
    if true {
        v19 = (((v5 + v21) + "foo bar") + "a")
        fmt.Println(v2, v1, v5)
    }
    v33 := GenS4{F0: 2}
    if (v1 != (v2 + 0)) {
        var v34 uint64 = 1
        delete(v3, 127)
        var v35 bool = false
        _ = v34
        _ = v35
    } else {
        v1 = (v2 + (v1 / ((v1) | 1)))
    }
    v36 := GenS5{F0: -556, F1: 2147483647, F2: 32767}
    delete(v3, 42)
    v37 := GenS6{F0: (v1 - 127), F1: (^(v6)), F2: v17}
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := v3[fi38]
        fmt.Println(v6)
        v40 := GenS2{F0: v21, F1: v5, F2: 128}
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
