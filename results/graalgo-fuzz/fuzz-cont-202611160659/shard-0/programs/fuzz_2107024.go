package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 string; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint16) bool {
    v4 := []int{-897, 4294967295}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = p0
        p0 = 570
        v6 := GenS1{F0: 0, F1: 1000, F2: -9223372036854775808}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return (127 > 2)
}
func genFn1(p0 int, p1 int64, p2 int32) uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 399
        v9 := []int{-1, (fi7 & fi7)}
        fmt.Println(p1, fi7, p1)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint8 = 255
        v3["hello"] = "foo bar"
        if true {
            v13 := GenS3{F0: (fi11 % (((fi11 ^ (-(fi11)))) | 1))}
            fmt.Println(v12, v12)
            _ = v13
        } else {
            v12 += v12
            v12 -= 227
            _ = len(v3)
        }
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS4{F0: 1, F1: 4294967295}
        var v16 uint32 = 1
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int16 = 0
    var v18 int = 2147483647
    var v19 uint8 = 255
    var v20 string = ("" + "a")
    v21 := GenS5{F0: 6, F1: 42}
    if (int(v19) != (v18 & (^(v18)))) {
        var v22 int32 = 0
        v22 = v22
        for fi23 := 0; fi23 < 4; fi23++ {
            v18 = v18
            var v24 float32 = 3.14159
            delete(v2, v20)
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v19 += (v19 ^ 255)
    switch -1 {
    case 0:
        var v25 int64 = 940
        _ = v25
    case 2147483647:
        if true {
            var v26 int32 = -850
            _ = v26
        }
    case -1:
        var v27 int64 = 200
        _ = v27
    default:
        v28 := GenS6{F0: "", F1: 4294967295, F2: 0}
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(v19, fi29)
        var v30 uint = 255
        _ = fi29
        _ = v30
    }
    v1 = append(v1, 2147483647)
    var v31 uint8 = (v19 | v19)
    v32 := GenS7{F0: 2}
    var v33 float32 = float32(3.14159)
    v18 = v18
    v34 := GenS8{F0: (-2.25 != 0.0), F1: 0}
    for fi35 := 0; fi35 < 17; fi35++ {
        if false {
            fmt.Println(v33, v31, v18)
            v31 = v19
            switch -920 {
            case 42:
                fmt.Println(v31, v33)
            case -2147483648:
                v18 -= 759
            case 32767:
                v33 -= (0.0 / 1.5)
            }
        } else {
            switch v20 {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    fmt.Println(fi36)
                    _ = fi36
                }
            }
        }
        var v37 string = "hello"
        _ = fi35
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
