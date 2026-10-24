package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 int64 }
func genFn0() int {
    var v5 uint32 = 2147483648
    if true {
        fmt.Println(v5)
    }
    var v6 int64 = 9223372036854775807
    v6 = v6
    _ = v5
    _ = v6
    return 590
}
func genFn1() int32 {
    var v7 float32 = 1.5
    v8 := []int{32768, 9223372036854775807, 42}
    v9 := GenS1{F0: 255, F1: 32767, F2: -1}
    var v10 float32 = (float32(((float64(v7) / 1.0) + float64(float32((-2.25 * (3.14159 - 3.14159)))))) / ((((0.0 + v7) * v7) / v7) - v7))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint = 18446744073709551615
    v4 := []int{42, 124, 399, -2147483648}
    v11 := GenS1{F0: 127, F1: 2, F2: 0}
    v12 := GenS2{F0: true, F1: 0, F2: 932}
    var v13 uint64 = 42
    var v14 uint8 = 0
    v15 := GenS1{F0: (^(v14)), F1: 42, F2: 32767}
    if (!v2) {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                if len(v4) > 0 {
                    v4[0] = 970
                }
                _ = fi17
            }
            _ = fi16
        }
        v18 := GenS2{F0: (!(!(((3.14159 / 1.0) != 3.14159) && v2))), F1: -683, F2: 42}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v2, v3)
            var v20 bool = (v14 > v14)
            var v21 int16 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
        v3 = v3
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int8 = 1
        _ = fi22
        _ = v23
    }
    v24 := v1[-681]
    v14 -= 2
    if false {
        var v25 string = (((((("\n\t" + "") + (v24 + v24)) + v24) + (v24 + v24)) + v24) + (v24 + "a"))
        _ = v25
    }
    v26 := GenS3{F0: 1000, F1: uint16(v14), F2: 9223372036854775807}
    if len(v4) > 0 {
        v4[0] = 128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
