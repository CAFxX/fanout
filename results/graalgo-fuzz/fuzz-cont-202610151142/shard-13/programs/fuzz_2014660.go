package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int8; F1 int64; F2 uint }
func genFn0(p0 int) uint {
    fmt.Println(p0)
    p0 = 1
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: ("a" + "a"), F1: (((3.14159 * (0.1 * float64(float32((3.14159 / 1e10))))) != -2.25) && true), F2: 781}
        var v7 uint16 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        if (0 < 1) {
            if (!true) {
                v8 := []int{(p0 / ((p0) | 1)), (p0 >> ((0) & 63)), (p0 * (-(p0))), (p0 << ((1) & 63))}
                p0 = (p0 << ((42) & 63))
                _ = v8
            } else {
                v9 := GenS2{F0: 127, F1: (334 != 32767)}
                var v10 int32 = -1
                fmt.Println(v10)
                _ = v10
                _ = v9
            }
            if true {
                if (28 >= 0) {
                    p0 = p0
                    var v11 int8 = -86
                    _ = v11
                }
                v12 := []int{4294967295}
                _ = v12
            }
            fmt.Println(p0)
        }
        p0 -= 128
        var v13 int = (p0 - 128)
        _ = v13
    }
    _ = p0
    return 42
}
func genFn1() int64 {
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 float32 = 3.14159
        var v16 int8 = -126
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{1}
    var v18 uint32 = 42
    v18 = (v18 % (((v18 - (v18 << ((65535) & 63)))) | 1))
    var v19 int64 = 610
    _ = v17
    _ = v18
    _ = v19
    return v19
}

func main() {
    v1 := []int{2147483647, 4294967295, 128, 32768}
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 int = 32767
    if (!(true && ((!(false || (!(-2.25 != -2.25)))) || (((0.0 - 0.0) * 1.5) < (float64(1e10) * (1e10 * (float64((((3.14159 * 1.5) - 3.14159) / 0.1)) / (1.5 * ((0.0 - 0.1) - 0.0))))))))) {
        for fi20 := 0; fi20 < 19; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = -445
            _ = fi20
        }
    } else {
        v4 -= 2147483647
        v4 -= 0
    }
    v21 := GenS1{F0: "foo bar", F1: false, F2: -1}
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS3{F0: 1, F1: genFn1()}
            var v25 uint16 = 255
            fmt.Println(fi23, v4, fi22)
            _ = fi23
            _ = v24
            _ = v25
        }
        v4 = 983
        v26 := GenS4{F0: 100, F1: 1, F2: 890}
        _ = fi22
        _ = v26
    }
    var v27 int8 = 40
    v27 += (v27 + v27)
    v27 += v27
    for fi28 := 0; fi28 < 9; fi28++ {
        v27 -= (^(v27))
        var v29 int = (fi28 ^ v4)
        v4 = 1
        _ = fi28
        _ = v29
    }
    var v30 int32 = 0
    if false {
        v31 := v2["a"]
        _ = v31
    } else {
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 int32 = -1
            v34 := GenS4{F0: v27, F1: 42, F2: 4294967295}
            _ = fi32
            _ = v33
            _ = v34
        }
        v27 -= (-(v27))
        fmt.Println(v4, v27)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
