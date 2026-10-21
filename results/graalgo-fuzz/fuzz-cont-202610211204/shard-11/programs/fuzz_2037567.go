package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 uint, p1 bool) int {
    var v5 uint16 = 167
    if p1 {
        switch 42 {
        case -202:
            v5 = (v5 + (v5 >> ((p0) & 63)))
        case 32767:
            var v6 bool = true
            _ = v6
        case 42:
            v5 += v5
        }
    } else {
        if (p1 && false) {
            p0 = (^(p0))
            fmt.Println(v5)
            for fi7 := 0; fi7 < 4; fi7++ {
                v5 = (^(v5))
                _ = fi7
            }
        }
    }
    v8 := GenS1{F0: 2, F1: 32767}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 339
}

func main() {
    v1 := []string{}
    var v2 int16 = 0
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v9 bool = (v2 >= (v2 ^ (v2 * 32767)))
    var v10 string = "\n\t"
    if (v9 || (1 >= 4294967295)) {
        var v11 float32 = (3.14159 * (1e10 * (-2.25 + float32((float64(1e10) * (float64(((0.0 - 1.5) / 0.1)) * (3.14159 - (0.0 - 1e10))))))))
        switch 590 {
        case 128:
            var v12 int64 = 0
            _ = v12
        default:
            fmt.Println(v9, v11, v2)
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            v3 = append(v3, 65535)
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := GenS2{F0: 2147483647, F1: -997, F2: -299}
                _ = fi14
                _ = v15
            }
            if false {
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 int64 = 0
                    fmt.Println(fi13)
                    _ = fi16
                    _ = v17
                }
                fmt.Println(fi13, v11, v10)
                var v18 int8 = 1
                _ = v18
            } else {
                fmt.Println(v11)
            }
            _ = fi13
        }
        _ = v11
    } else {
        _ = len(v4)
    }
    v10 = (v10 + v10)
    if len(v1) > 0 {
        v1[0] = (("hello" + v10) + v10)
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int16 = v2
        var v21 float32 = (1.5 - 0.0)
        fmt.Println(v21, v10, v20)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float64 = float64(float32(float64(3.14159)))
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 bool = false
        v25 := GenS3{F0: (-(fi23)), F1: (!v9)}
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(fi23)
            v24 = ((12 != 42) || true)
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v27 int8 = -128
    if (0 <= 2147483648) {
        var v28 bool = v9
        _ = v28
    } else {
        var v29 int32 = 42
        var v30 int16 = v2
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        fmt.Println(v10, v2, v27)
        switch (fi31 >> ((803) & 63)) {
        case 32767:
            if v9 {
                v32 := GenS2{F0: 0, F1: 397, F2: (v2 | v2)}
                var v33 float32 = (0.0 * (float32((1e10 * ((v22 / 1.0) / v22))) + (3.14159 * float32((0.0 + -2.25)))))
                _ = v32
                _ = v33
            } else {
                fmt.Println(v27, fi31)
                var v34 string = "\n\t"
                _ = v34
            }
        case 32768:
            fmt.Println(v22)
        default:
            v22 = float64(float32(1e10))
        }
        _ = fi31
    }
    v3 = append(v3, 4294967295)
    delete(v4, 4294967295)
    v35 := GenS4{F0: 0}
    if false {
        v27 += (v27 % ((v27) | 1))
        v36 := GenS5{F0: 4294967295, F1: v10, F2: 65535}
        v10 = (("a" + v10) + v10)
        _ = v36
    } else {
        var v37 bool = false
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = v37
    }
    delete(v4, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v9
}
