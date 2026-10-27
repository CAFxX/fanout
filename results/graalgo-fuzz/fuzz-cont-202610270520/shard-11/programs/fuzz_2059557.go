package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 string; F2 string }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0() uint64 {
    var v7 uint8 = 2
    switch "" {
    case "\n\t":
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                switch ("" + "hello") {
                case "foo bar":
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 float64 = 0.0
                        var v11 int = fi8
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case "a":
                    var v12 int16 = 1000
                    _ = v12
                case "\n\t":
                    v13 := []int{fi8, (fi8 << ((818) & 63)), 32767, (fi8 & fi8)}
                    _ = v13
                }
                v14 := GenS1{F0: 42, F1: 42, F2: fi8}
                var v15 uint32 = 65535
                _ = fi8
                _ = v14
                _ = v15
            }
        } else {
            var v16 int32 = 1
            var v17 float32 = (-2.25 - 3.14159)
            if (!(("" + "hello") != "hello")) {
                v17 = v17
            }
            _ = v16
            _ = v17
        }
    case "a":
        v7 = v7
    case "":
        v18 := GenS1{F0: (v7 * (v7 * 127)), F1: 4294967295, F2: int(v7)}
        _ = v18
    default:
        var v19 int16 = 32767
        _ = v19
    }
    _ = v7
    return 110
}
func genFn1(p0 bool) int64 {
    v20 := GenS2{F0: 1, F1: ((("hello" + "a") + "foo bar") + "\n\t"), F2: "a"}
    if true {
        p0 = (p0 || (-1 >= 1))
        v21 := []int{}
        _ = v21
    }
    _ = p0
    _ = v20
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    var v3 string = "\n\t"
    v4 := make(map[int]uint32)
    var v5 float64 = (3.14159 - 1e10)
    v6 := make(map[uint32]bool)
    fmt.Println(v2, v2)
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            v2 += v2
            delete(v6, 4294967295)
            v5 -= v5
            if (32767 > 610) {
                var v24 uint64 = 356
                v3 = v3
                _ = v24
            }
            v2 = v2
            _ = fi23
        }
        v2 = v2
        _ = fi22
    }
    var v25 uint64 = 4294967295
    v2 += v2
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 bool = true
        _ = fi26
        _ = v27
    }
    switch 2147483647 {
    case 128:
        var v28 int64 = -9223372036854775808
        _ = v28
    case 2147483647:
        v25 = (v25 - (v25 + (v25 % (((v25 % ((v25) | 1))) | 1))))
    case 32767:
        fmt.Println(v3)
    }
    v29 := GenS1{F0: 154, F1: 2, F2: 375}
    v30 := []int{42, 127, -558, 985}
    var v31 float32 = (0.1 * (-2.25 + -2.25))
    fmt.Println(v31, v5)
    switch 870 {
    case 0:
        for fi32 := 0; fi32 < 13; fi32++ {
            v3 = (v3 + v3)
            v33 := GenS2{F0: -2147483648, F1: v3, F2: "a"}
            if (0 != 1) {
                v34 := GenS1{F0: 255, F1: 396, F2: (fi32 << ((514) & 63))}
                _ = v34
            } else {
                var v35 int16 = -32768
                var v36 bool = false
                v25 -= v25
                _ = v35
                _ = v36
            }
            _ = fi32
            _ = v33
        }
    case -850:
        v37 := GenS1{F0: 24, F1: 436, F2: 825}
        _ = v37
    default:
        fmt.Println(v2, v2)
    }
    switch 2147483647 {
    case 128:
        fmt.Println(v2, v25)
    case 42:
        var v38 int32 = 42
        _ = v38
    default:
        switch -2147483648 {
        case -609:
            _ = len(v4)
        case 32768:
            v39 := GenS1{F0: 1, F1: 4294967295, F2: -9223372036854775808}
            _ = v39
        default:
            v25 -= genFn0()
        }
    }
    if false {
        var v40 float64 = 1.5
        switch v3 {
        case "a":
            v3 = (v3 + "")
        case "\n\t":
            var v41 bool = false
            _ = v41
        case "":
            v40 -= v5
        default:
            v4[-30] = 2147483647
        }
        _ = v40
    } else {
        v42 := GenS3{F0: 255, F1: (^(v25))}
        _ = v42
    }
    if (127 > 0) {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
