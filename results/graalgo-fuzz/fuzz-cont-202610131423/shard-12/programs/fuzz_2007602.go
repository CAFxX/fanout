package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 uint16 = 42
    v4 := []uint32{65535}
    v5 := make(map[int]int)
    fmt.Println(v3)
    var v6 float64 = 3.14159
    v7 := GenS1{F0: 127, F1: (v3 / ((v3) | 1))}
    v8 := GenS2{F0: (2147483647 == 1), F1: 903, F2: uint(v3)}
    v2 = v2
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS3{F0: v2, F1: (v2 % ((int8(v3)) | 1)), F2: 0}
        v11 := GenS3{F0: (v2 + -128), F1: (v2 - (v2 >> ((4294967295) & 63))), F2: 2147483647}
        _ = fi9
        _ = v10
        _ = v11
    }
    v6 = v6
    fmt.Println(v6)
    var v12 int64 = 0
    fmt.Println(v2, v2)
    if false {
        var v13 string = "foo bar"
        v13 = (v13 + ("" + v13))
        _ = v13
    } else {
        fmt.Println(v3, v12)
        fmt.Println(v12)
    }
    v12 = v12
    delete(v5, 2147483647)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float32 = 0.0
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v3)
    var v16 uint32 = 360
    for fi17 := 0; fi17 < 12; fi17++ {
        if (v2 > 127) {
            v6 = v6
            v12 = (v12 * v12)
            if true {
                v18 := v5[(fi17 ^ (fi17 | -718))]
                var v19 int16 = -32768
                _ = v18
                _ = v19
            }
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
