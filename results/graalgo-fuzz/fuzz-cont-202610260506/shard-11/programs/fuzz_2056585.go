package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint8 }
func genFn0(p0 uint8) uint32 {
    var v7 uint32 = 2
    v8 := []int{}
    var v9 int8 = 127
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v7 + 0)
}
func genFn1(p0 bool) int32 {
    var v10 float64 = ((-2.25 - 0.1) / (-2.25 - 0.1))
    v11 := []int{45, 0, 0}
    if true {
        if p0 {
            if ((0.1 >= 0.0) && true) {
                if (0.1 != (1e10 + 3.14159)) {
                    switch "foo bar" {
                    case "hello":
                        var v12 bool = (4294967295 <= 203)
                        _ = v12
                    case "":
                        if (63 <= 1) {
                            if p0 {
                                v11 = append(v11, 127)
                                if p0 {
                                    p0 = p0
                                    if len(v11) > 0 {
                                        v11[0] = 4294967295
                                    }
                                    p0 = (!(p0 && false))
                                }
                            } else {
                                v13 := GenS1{F0: 42, F1: 127}
                                v10 = (v10 / v10)
                                _ = v13
                            }
                        } else {
                            var v14 string = "a"
                            _ = v14
                        }
                    }
                } else {
                    var v15 float64 = ((1e10 * v10) / float64(1e10))
                    _ = v15
                }
                if len(v11) > 0 {
                    v11[0] = 32767
                }
                var v16 int8 = 0
                _ = v16
            } else {
                p0 = false
            }
            fmt.Println(v10, v10)
        }
    } else {
        v17 := GenS2{F0: 42, F1: 32767}
        fmt.Println(p0, v10)
        v18 := GenS1{F0: 100, F1: 127}
        _ = v17
        _ = v18
    }
    _ = p0
    _ = v10
    _ = v11
    return -257
}

func main() {
    var v1 uint8 = 168
    v2 := make(map[int]bool)
    var v3 int8 = 59
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    v6 := make(map[int]string)
    if (42 >= 2) {
        _ = len(v5)
        _ = len(v2)
        var v19 int = 812
        _ = v19
    }
    var v20 bool = (0 < 969)
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 -= (v1 / ((v1) | 1))
        var v22 int = (fi21 + -1)
        for fi23 := 0; fi23 < 7; fi23++ {
            v1 -= v1
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v1 = 0
    fmt.Println(v1)
    var v24 int64 = 971
    v6[-474] = "foo bar"
    switch ((("foo bar" + "\n\t") + "foo bar") + "foo bar") {
    case "hello":
        var v25 uint64 = 1
        _ = v25
    case "":
        v26 := GenS2{F0: 4294967295, F1: 745}
        _ = v26
    default:
        fmt.Println(v1)
    }
    v1 = (v1 + v1)
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS2{F0: 954, F1: 2}
        v1 = v1
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS1{F0: 42, F1: v3}
            v3 = 127
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    v31 := GenS3{F0: 308, F1: "hello"}
    fmt.Println(v20, v20)
    if (!true) {
        v6[32767] = "hello"
    } else {
        switch "" {
        case "a":
            if ((v24 - v24) < v24) {
                v5 = append(v5, 2147483648)
                fmt.Println(v24, v1, v3)
                var v32 float64 = -2.25
                _ = v32
            }
        case "foo bar":
            fmt.Println(v20, v24, v20)
        }
        v33 := v2[829]
        _ = v33
    }
    v34 := GenS4{F0: 65535, F1: v1, F2: v1}
    switch 127 {
    case -1:
        for fi35 := 0; fi35 < 20; fi35++ {
            v20 = (!((v24 | (v24 % ((v24) | 1))) != v24))
            v1 = (^(v1))
            _ = fi35
        }
    case 32767:
        var v36 uint32 = 2147483647
        _ = v36
    default:
        var v37 uint16 = 32767
        _ = v37
    }
    var v38 float64 = 0.1
    switch 32768 {
    case 32768:
        if v20 {
            if len(v5) > 0 {
                v5[0] = 65535
            }
            v38 = 0.0
        }
    default:
        for fi39 := 0; fi39 < 2; fi39++ {
            v38 = 1.5
            v40 := GenS3{F0: 0, F1: "\n\t"}
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v31
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
