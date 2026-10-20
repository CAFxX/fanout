package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 bool, p1 int8) uint16 {
    var v6 uint32 = 0
    switch 42 {
    case 128:
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                if (!p0) {
                    p1 -= p1
                    if (1.5 <= (-2.25 - 0.0)) {
                        var v9 int = (fi7 ^ 32768)
                        _ = v9
                    } else {
                        fmt.Println(fi7, v6)
                        if p0 {
                            v10 := GenS1{F0: -913, F1: 255, F2: -1}
                            _ = v10
                        }
                    }
                    var v11 float32 = ((0.0 - ((3.14159 + (float32((-2.25 / 1e10)) / float32((3.14159 + -2.25)))) / 1.5)) * 1.5)
                    _ = v11
                } else {
                    v12 := GenS1{F0: 916, F1: 42, F2: -1}
                    _ = v12
                }
                v13 := []int{}
                _ = fi8
                _ = v13
            }
            _ = fi7
        }
    default:
        if false {
            fmt.Println(v6)
            var v14 uint = 899
            var v15 int = -901
            _ = v14
            _ = v15
        } else {
            var v16 uint64 = 4294967295
            for fi17 := 0; fi17 < 19; fi17++ {
                if true {
                    fmt.Println(fi17, v6)
                    p0 = p0
                } else {
                    var v18 float64 = (3.14159 * float64(0.1))
                    _ = v18
                }
                _ = fi17
            }
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 553
}
func genFn1(p0 string, p1 uint32) int16 {
    v19 := []int{32767, -983}
    p0 = "\n\t"
    var v20 uint64 = 0
    p0 = p0
    for fi21 := 0; fi21 < 10; fi21++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        var v22 int32 = -386
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return int16(p1)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -853
    v3 := make(map[string]uint32)
    var v4 int16 = 0
    v5 := make(map[uint32]uint32)
    if (4294967295 < 42) {
        v4 += v4
        fmt.Println(v4, v4)
    }
    var v23 int8 = 96
    var v24 int64 = -1
    v4 -= (v4 >> ((255) & 63))
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 float64 = 1.5
        if true {
            switch "hello" {
            case "a":
                fmt.Println(v4)
            }
        } else {
            fmt.Println(v23, v24)
        }
        _ = fi25
        _ = v26
    }
    var v27 int64 = v24
    v28 := GenS2{F0: (!((^(v24)) >= v24)), F1: 65535, F2: v4}
    var v29 float32 = -2.25
    switch "hello" {
    case "a":
        if false {
            if true {
                v30 := GenS2{F0: true, F1: 65535, F2: v4}
                var v31 int = 66
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 uint16 = 0
                    _ = fi32
                    _ = v33
                }
                _ = v30
                _ = v31
            }
            v4 += (v4 | 412)
            v34 := []int{-928, 1, 32768}
            _ = v34
        }
    case "":
        var v35 int32 = 1
        _ = v35
    default:
        var v36 int32 = 0
        _ = v36
    }
    switch -687 {
    case 42:
        fmt.Println(v29)
    }
    if ((88 >= (v23 % ((v23) | 1))) || (genFn1(("" + "foo bar"), 147) <= (v4 >> ((0) & 63)))) {
        var v37 uint64 = 1
        v29 = ((float32(((float64(((1e10 / 0.1) / v29)) - (0.0 / (0.1 * -2.25))) - 3.14159)) - float32(0.0)) - 0.0)
        if (0 >= 63) {
            var v38 float64 = -2.25
            if true {
                v4 -= genFn1("\n\t", 2147483648)
                var v39 uint16 = 965
                v39 = (^(v39))
                _ = v39
            } else {
                for fi40 := 0; fi40 < 15; fi40++ {
                    v41 := []int{int(v27)}
                    _ = fi40
                    _ = v41
                }
                v4 -= v4
                v42 := GenS1{F0: 1, F1: 1, F2: 2147483647}
                _ = v42
            }
            var v43 float32 = (float32(((v38 / (-2.25 - float64(0.1))) - v38)) + (-2.25 - -2.25))
            _ = v38
            _ = v43
        }
        _ = v37
    }
    _ = len(v5)
    fmt.Println(v4, v29)
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS3{F0: 948, F1: (!true), F2: (fi44 | (^(fi44)))}
        v24 = 42
        _ = fi44
        _ = v45
    }
    var v46 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
