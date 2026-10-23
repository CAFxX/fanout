package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint; F1 int; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int8 }
func genFn0(p0 string, p1 int64) uint64 {
    v6 := GenS1{F0: p0, F1: 844}
    v7 := GenS2{F0: 2147483648, F1: (32767 != 0)}
    if (3.14159 < 1e10) {
        for fi8 := 0; fi8 < 11; fi8++ {
            if false {
                fmt.Println(p0)
            } else {
                var v9 int32 = 751
                _ = v9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS3{F0: 42, F1: 201, F2: -2147483648}
                var v12 uint64 = 42
                switch fi8 {
                case -974:
                    p1 = p1
                case 127:
                    v13 := []int{(fi10 ^ 128), (fi8 >> ((65535) & 63)), fi8, fi8}
                    _ = v13
                default:
                    v14 := GenS2{F0: 2, F1: (65535 != 0)}
                    _ = v14
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    } else {
        p0 = ((p0 + p0) + p0)
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint) int8 {
    var v15 float64 = (-2.25 - (1.5 / (1e10 - float64(float32(3.14159)))))
    for fi16 := 0; fi16 < 10; fi16++ {
        v15 -= (v15 + 0.0)
        var v17 uint64 = 4294967295
        _ = fi16
        _ = v17
    }
    v18 := []int{-561, -1}
    if true {
        var v19 int16 = 32767
        var v20 uint64 = 18446744073709551615
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v15
    _ = v18
    return -128
}

func main() {
    var v1 int32 = -1
    var v2 float64 = (1.5 / 1e10)
    v3 := []uint32{}
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    if (!false) {
        for fi21 := 0; fi21 < 3; fi21++ {
            switch "foo bar" {
            case "a":
                var v22 float64 = v2
                _ = v22
            case "":
                v23 := GenS3{F0: 42, F1: -896, F2: -1}
                _ = v23
            default:
                v1 = 2147483647
            }
            _ = fi21
        }
    }
    v24 := GenS3{F0: 2, F1: 127, F2: -1}
    v25 := GenS1{F0: ("\n\t" + ""), F1: 42}
    if true {
        var v26 int64 = 9223372036854775807
        v27 := v5[127]
        _ = v26
        _ = v27
    } else {
        v28 := GenS4{F0: 251, F1: 1, F2: 1}
        var v29 int32 = v1
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 += -2147483648
        fmt.Println(v2)
        v1 -= v1
        _ = fi30
    }
    if (!true) {
        var v31 string = ("foo bar" + "a")
        switch 4294967295 {
        case -1:
            delete(v4, v31)
        case 32767:
            if true {
                switch v31 {
                case "foo bar":
                    if (float64(-2.25) >= (float64(0.0) * v2)) {
                        fmt.Println(v1, v31)
                    }
                case "a":
                    _ = len(v5)
                default:
                    for fi32 := 0; fi32 < 10; fi32++ {
                        var v33 int = fi32
                        for fi34 := 0; fi34 < 20; fi34++ {
                            v35 := GenS5{F0: v1, F1: 42, F2: 42}
                            v36 := GenS6{F0: v33, F1: 18446744073709551615}
                            var v37 float32 = 0.1
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                for fi38 := 0; fi38 < 11; fi38++ {
                    v4[("" + ("foo bar" + "hello"))] = ("\n\t" + (v31 + v31))
                    v39 := GenS7{F0: 78}
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 uint32 = 42
                        var v42 uint16 = 1
                        v2 = 1e10
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = fi38
                    _ = v39
                }
            }
        case 235:
            if (!false) {
                fmt.Println(v1, v1, v31)
                var v43 string = ""
                _ = v43
            } else {
                var v44 int64 = -1
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 int64 = (-(v44))
                    fmt.Println(v1, fi45, v44)
                    var v47 int8 = -128
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = v44
            }
        }
        _ = v31
    } else {
        v1 -= v1
    }
    fmt.Println(v1, v1)
    for fi48 := 0; fi48 < 2; fi48++ {
        var v49 int64 = -9223372036854775808
        var v50 int = fi48
        _ = fi48
        _ = v49
        _ = v50
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
