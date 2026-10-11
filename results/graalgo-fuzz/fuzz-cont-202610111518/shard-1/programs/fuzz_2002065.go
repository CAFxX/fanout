package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 string; F2 bool }
type GenS4 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 int8) bool {
    p0 = p0
    v6 := []int{}
    var v7 uint16 = 42
    v7 = 652
    _ = p0
    _ = v6
    _ = v7
    return ((-2.25 * float64(((((3.14159 * 0.0) / ((3.14159 + (-2.25 + (((1e10 / ((3.14159 / 0.1) * (-2.25 - (1.5 + 0.1)))) / float32(float64(0.1))) / 1e10))) / -2.25)) * 0.1) + 0.1))) < 1e10)
}
func genFn1() uint8 {
    v8 := GenS1{F0: 42, F1: -32768}
    var v9 int32 = 0
    var v10 string = "foo bar"
    var v11 int32 = (-(v9))
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -1
    var v3 string = "hello"
    v4 := []uint32{42, 42, 65535, 962}
    var v5 uint8 = 2
    v12 := GenS2{F0: 2147483647}
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 = "a"
        _ = fi13
    }
    var v14 string = "\n\t"
    v15 := GenS1{F0: 32767, F1: -32768}
    v2 -= v2
    var v16 int16 = -923
    if (!false) {
        if (255 >= 802) {
            var v17 uint = 65535
            v18 := GenS3{F0: 237, F1: "foo bar", F2: true}
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := v1[(v14 + v14)]
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            if len(v4) > 0 {
                v4[0] = 65535
            }
            fmt.Println(v14, v5)
            _ = fi21
        }
        var v22 uint8 = v5
        _ = v22
    }
    var v23 int8 = v2
    var v24 float32 = ((((-2.25 + -2.25) * 0.1) - 1.5) + 0.1)
    v3 = "a"
    var v25 int8 = 100
    fmt.Println(v5, v16)
    v26 := GenS4{F0: "\n\t", F1: 128, F2: (^(v23))}
    var v27 uint64 = 4294967295
    v27 = (v27 | v27)
    var v28 int8 = (^(v23))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
