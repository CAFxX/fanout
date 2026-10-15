package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }
type GenS2 struct { F0 int16; F1 uint; F2 string }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 uint, p1 uint64, p2 int16) uint {
    var v7 int = 1
    v8 := GenS1{F0: (p1 + 400), F1: (v7 ^ (v7 + v7)), F2: 0}
    v9 := GenS1{F0: p1, F1: 0, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 uint, p1 uint32, p2 int8) int8 {
    switch 700 {
    case 42:
        p1 = uint32(p0)
    case 32767:
        var v10 uint64 = 0
        _ = v10
    case 2147483647:
        if (32767 > 1) {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 int64 = -1
                v13 := []int{fi11}
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint64 = 568
            v14 = v14
            _ = v14
        }
    default:
        if true {
            v15 := []int{}
            v15 = append(v15, 9223372036854775807)
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint8 = 141
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            p2 = (p2 * (p2 * (p2 * p2)))
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        if (654 < 42) {
            var v19 uint = genFn0((p0 - p0), 0, 32767)
            switch "hello" {
            case "\n\t":
                var v20 float32 = (0.0 + 1.5)
                _ = v20
            case "a":
                switch "\n\t" {
                case "":
                    var v21 bool = false
                    _ = v21
                case "\n\t":
                    p1 = 65535
                default:
                    p0 -= p0
                }
            default:
                switch 4294967295 {
                case -1:
                    p1 = p1
                }
            }
            if true {
                v22 := []int{fi18, (fi18 | 127), 4294967295}
                v23 := GenS2{F0: 938, F1: (^(v19)), F2: ("" + ("hello" + ""))}
                var v24 int32 = -261
                _ = v22
                _ = v23
                _ = v24
            }
            _ = v19
        } else {
            var v25 bool = false
            _ = v25
        }
        var v26 int32 = 42
        _ = fi18
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (float32(((0.1 * 0.0) - float64(0.1))) + 0.1)
    v3 := []string{}
    var v4 string = "foo bar"
    var v5 int = 112
    v6 := make(map[uint32]uint32)
    if (1.5 != ((float64((0.1 - v2)) / (-2.25 / 1.0)) - 1e10)) {
        v2 += v2
        var v27 bool = (1 < -1)
        _ = v27
    }
    if false {
        if (1 < 255) {
            if true {
                var v28 bool = (2 < 2)
                v29 := GenS1{F0: 2, F1: (v5 >> ((genFn0(genFn0(4294967295, 179, 935), 42, 32767)) & 63)), F2: 42}
                _ = v28
                _ = v29
            } else {
                v5 = -141
                var v30 int16 = 0
                v2 += v2
                _ = v30
            }
            fmt.Println(v2, v2)
            var v31 int8 = 100
            _ = v31
        } else {
            fmt.Println(v4)
            v5 = 9223372036854775807
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            v5 = fi32
            _ = fi32
        }
        if true {
            var v33 uint16 = 42
            var v34 float32 = v2
            _ = v33
            _ = v34
        }
    } else {
        v35 := GenS2{F0: 449, F1: 65535, F2: "foo bar"}
        var v36 uint64 = 4294967295
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        _ = len(v1)
        var v38 uint64 = 94
        _ = fi37
        _ = v38
    }
    var v39 int64 = -9223372036854775808
    var v40 uint32 = 65535
    fmt.Println(v2, v39, v5)
    if false {
        v41 := GenS3{F0: 255, F1: 18446744073709551615}
        var v42 int32 = 0
        _ = v41
        _ = v42
    } else {
        var v43 uint = 298
        fmt.Println(v43, v39)
        v44 := GenS2{F0: -471, F1: v43, F2: (("foo bar" + ("\n\t" + ((v4 + (v4 + ((v4 + "") + v4))) + "foo bar"))) + (v4 + (v4 + (v4 + ((((v4 + ((((("foo bar" + "") + ((v4 + v4) + ("\n\t" + v4))) + v4) + "foo bar") + v4)) + v4) + "a") + v4)))))}
        _ = v43
        _ = v44
    }
    v40 = v40
    if (3.14159 > 3.14159) {
        v40 -= v40
        fmt.Println(v5, v40, v40)
    }
    switch v5 {
    case -1:
        var v45 int = v5
        _ = v45
    case 42:
        delete(v1, v40)
    case 932:
        v4 = "a"
    default:
        v46 := GenS3{F0: 65535, F1: 18446744073709551615}
        _ = v46
    }
    fmt.Println(v39)
    switch 127 {
    case -124:
        if false {
            var v47 int16 = 1000
            var v48 uint = 0
            _ = v47
            _ = v48
        }
    }
    v39 += v39
    fmt.Println(v2, v4, v2)
    v5 -= (v5 | (v5 & v5))
    v4 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
