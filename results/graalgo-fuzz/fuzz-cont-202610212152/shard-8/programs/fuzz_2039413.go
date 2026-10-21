package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 uint) uint32 {
    fmt.Println(p0, p0)
    var v6 float64 = (1e10 - (1.5 * (-2.25 / (1.5 / (0.1 * -2.25)))))
    _ = p0
    _ = p1
    _ = v6
    return 2147483648
}
func genFn1(p0 bool, p1 int64) string {
    fmt.Println(p1)
    var v7 uint16 = 858
    _ = p0
    _ = p1
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 2147483647
    var v3 uint32 = (v2 * v2)
    var v4 float32 = 0.0
    v5 := []int{}
    switch 42 {
    case -15:
        var v8 int64 = 0
        _ = v8
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi9
        }
    }
    if false {
        v2 -= genFn0(((-79 == 0) || true), 738)
        v2 = genFn0((979 > 65535), 255)
        if (1 != 2) {
            v4 += v4
        } else {
            fmt.Println(v4)
            var v10 int8 = 102
            v11 := GenS1{F0: (v3 | v2), F1: 2}
            _ = v10
            _ = v11
        }
    } else {
        var v12 int16 = 433
        v12 -= v12
        for fi13 := 0; fi13 < 3; fi13++ {
            delete(v1, "a")
            v14 := GenS1{F0: (v2 | v2), F1: 255}
            v15 := GenS2{F0: 0, F1: 468}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    if (false && false) {
        v4 = (v4 - v4)
        v3 = (v2 ^ v3)
        v4 -= v4
    } else {
        var v16 string = "hello"
        _ = v16
    }
    fmt.Println(v2, v2, v3)
    var v17 int8 = 100
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if ((v2 ^ v2) == genFn0((915 >= 1), 2)) {
            v18 := GenS3{F0: "a"}
            _ = len(v5)
            _ = v18
        }
    }
    if true {
        var v19 int8 = (-(v17))
        switch 32768 {
        case 4294967295:
            v20 := GenS1{F0: 2147483647, F1: 435}
            _ = v20
        case 42:
            var v21 float32 = 1e10
            _ = v21
        case -634:
            v17 = 1
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                if false {
                    var v23 string = ""
                    _ = v23
                } else {
                    v4 -= ((-2.25 - 1e10) - v4)
                }
                _ = fi22
            }
        }
        _ = v19
    }
    v2 = 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -633 {
    case -398:
        v2 = v3
    case 2147483647:
        switch 128 {
        case 0:
            fmt.Println(v17, v2)
        case 2147483647:
            var v24 uint32 = v2
            _ = v24
        default:
            var v25 string = "foo bar"
            _ = v25
        }
    case 0:
        var v26 int64 = 9223372036854775807
        _ = v26
    default:
        v27 := GenS4{F0: 255, F1: (false && false)}
        _ = v27
    }
    v17 -= 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
