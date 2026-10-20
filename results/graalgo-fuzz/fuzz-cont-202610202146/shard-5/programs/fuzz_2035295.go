package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0(p0 bool) string {
    v4 := []int{-9223372036854775808, -1, 966, 551}
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(fi5, fi5)
        if (136 > 540) {
            v6 := GenS1{F0: 1, F1: 730}
            fmt.Println(p0, p0)
            if p0 {
                var v7 int = (fi5 * (-(fi5)))
                _ = v7
            }
            _ = v6
        } else {
            if p0 {
                p0 = p0
                for fi8 := 0; fi8 < 15; fi8++ {
                    var v9 float32 = 0.0
                    fmt.Println(fi8)
                    v4 = append(v4, fi5)
                    v10 := GenS1{F0: 1, F1: 804}
                    switch ("\n\t" + "") {
                    case "":
                        _ = len(v4)
                    case "a":
                        var v11 uint32 = 42
                        _ = v11
                    default:
                        fmt.Println(p0, fi5, fi5)
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                fmt.Println(fi5, fi5)
            } else {
                v12 := GenS1{F0: 42, F1: 1000}
                if p0 {
                    var v13 uint32 = 4294967295
                    switch fi5 {
                    case -1:
                        v13 = v13
                    case 1:
                        _ = len(v4)
                    case 128:
                        var v14 uint16 = 1
                        _ = v14
                    }
                    _ = v13
                }
                var v15 uint16 = 171
                _ = v12
                _ = v15
            }
        }
        var v16 int8 = 100
        _ = fi5
        _ = v16
    }
    v17 := GenS2{F0: 42, F1: 547}
    var v18 float32 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v4
    return ((("hello" + "") + "foo bar") + ("a" + "hello"))
}
func genFn1(p0 uint64, p1 bool, p2 bool) int {
    v19 := GenS2{F0: 4294967295, F1: (p0 ^ p0)}
    var v20 int32 = -112
    v21 := GenS1{F0: 65535, F1: -19}
    switch -17 {
    case 0:
        v22 := GenS2{F0: p0, F1: (p0 % (((^(p0))) | 1))}
        _ = v22
    default:
        fmt.Println(p2, v20, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return 380
}

func main() {
    v1 := []uint32{42, 472}
    v2 := make(map[string]int)
    var v3 int8 = 42
    var v23 float64 = 3.14159
    v3 = -1
    v24 := GenS1{F0: 2, F1: -1}
    v23 = v23
    v25 := GenS3{F0: 42}
    _ = len(v1)
    if (442 >= 42) {
        v23 = 0.0
        v26 := GenS2{F0: 1, F1: 382}
        for fi27 := 0; fi27 < 18; fi27++ {
            v2["\n\t"] = fi27
            _ = fi27
        }
        _ = v26
    } else {
        if false {
            switch "hello" {
            case "a":
                v3 += 42
            case "":
                v1 = append(v1, 65535)
            }
        } else {
            v28 := GenS4{F0: 0, F1: 9223372036854775807, F2: 416}
            _ = v28
        }
        v29 := GenS5{F0: 3, F1: 42, F2: 42}
        _ = v29
    }
    v3 -= v3
    v23 = (v23 + 1e10)
    v23 -= 1e10
    v23 = v23
    var v30 uint32 = 1
    v31 := GenS4{F0: 1, F1: 127, F2: 709}
    if (!true) {
        var v32 uint16 = 2
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 int8 = -91
            v35 := GenS6{F0: v32, F1: uint32(v32), F2: v34}
            switch (fi33 >> ((1) & 63)) {
            case -2147483648:
                var v36 int8 = 1
                _ = v36
            case 128:
                _ = len(v2)
            default:
                v32 -= v32
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v32
    } else {
        v37 := GenS7{F0: 0, F1: 244}
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
}
