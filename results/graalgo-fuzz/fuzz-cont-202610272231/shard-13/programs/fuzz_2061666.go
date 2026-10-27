package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
func genFn0(p0 int64) int32 {
    var v4 uint8 = 205
    var v5 int8 = -128
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 4294967295
        _ = fi6
        _ = v7
    }
    if true {
        var v8 uint8 = 42
        var v9 float64 = (1.5 - 1.5)
        var v10 uint64 = 42
        _ = v10
        _ = v8
        _ = v9
    }
    if false {
        v11 := []int{}
        _ = len(v11)
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int32 {
    var v12 string = ""
    fmt.Println(v12)
    v13 := GenS1{F0: 158, F1: 9223372036854775807, F2: 32767}
    v14 := GenS1{F0: 255, F1: 0, F2: 1}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 float32 = 1e10
        var v17 float64 = ((0.0 + 0.0) + 1.5)
        var v18 float32 = (1.5 - v16)
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    _ = v12
    _ = v13
    _ = v14
    return -2147483648
}

func main() {
    var v1 int = 32768
    var v2 float64 = 1.5
    var v3 bool = true
    v19 := []int{}
    v20 := GenS1{F0: 128, F1: -9223372036854775808, F2: -932}
    v21 := GenS1{F0: 2, F1: 9223372036854775807, F2: 1000}
    var v22 int64 = 0
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            _ = len(v19)
            _ = fi24
        }
        _ = fi23
    }
    _ = len(v19)
    switch ("hello" + "\n\t") {
    case "":
        v22 = (v22 + v22)
    case "hello":
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    }
    if ((v22 << ((65535) & 63)) <= 1) {
        v2 = 3.14159
        if (((float32((3.14159 * (v2 * 1.5))) / 1.0) + (float32(1e10) + 0.0)) <= 3.14159) {
            if len(v19) > 0 {
                v19[0] = v1
            }
        }
    }
    v25 := GenS2{F0: 0, F1: ((float64(3.14159) + (v2 / (v2 / v2))) >= v2), F2: 97}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
}
