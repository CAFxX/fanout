package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int16; F1 uint64 }
func genFn0() uint32 {
    var v4 uint = 2
    v5 := []int{}
    v5 = append(v5, 405)
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]string)
    var v3 uint8 = 128
    if false {
        var v6 uint64 = 0
        var v7 uint = 869
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 1000, F1: (v3 ^ v3)}
        _ = len(v2)
        _ = v8
    }
    var v9 string = "hello"
    v10 := GenS2{F0: 963, F1: 210}
    v9 = v9
    var v11 float64 = float64(1e10)
    v11 = ((((v11 * v11) + v11) * (0.0 + (float64((-2.25 * 0.0)) / v11))) / 1.5)
    v12 := []int{-316, 454, 32767, 4294967295}
    v13 := GenS3{F0: 42, F1: v3}
    v11 = 0.0
    var v14 float32 = 0.1
    v15 := GenS4{F0: 65535, F1: v3}
    fmt.Println(v1, v9, v9)
    if false {
        v2["a"] = (v9 + v9)
        var v16 uint8 = v3
        if false {
            var v17 uint8 = (v3 >> ((611) & 63))
            _ = v17
        } else {
            switch 912 {
            case 42:
                fmt.Println(v1, v14, v1)
            case -2147483648:
                _ = len(v2)
            }
            v3 = v3
        }
        _ = v16
    }
    v18 := GenS3{F0: -103, F1: (v3 ^ v3)}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if (true || (v3 != 255)) {
        v19 := GenS5{F0: 4294967295}
        v20 := GenS6{F0: 582, F1: 255}
        _ = v19
        _ = v20
    } else {
        v14 -= v14
        v21 := GenS3{F0: v1, F1: v3}
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 3; fi23++ {
                if true {
                    v24 := GenS7{F0: 0, F1: 83}
                    delete(v2, "foo bar")
                    _ = v24
                } else {
                    var v25 float32 = 0.1
                    _ = v25
                }
                v9 = v9
                _ = fi23
            }
            var v26 float32 = (v14 / 0.1)
            _ = fi22
            _ = v26
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
