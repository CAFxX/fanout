package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint32 = 794
    var v2 string = "a"
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 float32 = float32((0.0 - float64((((((((0.1 + 0.1) + ((0.0 + 3.14159) * 3.14159)) - 0.1) * 0.1) / (-2.25 - (((0.0 - (-2.25 - 3.14159)) + -2.25) / -2.25))) / 3.14159) + 1.5))))
    v1 = (v1 & v1)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = 32767
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: -32768, F2: -128}
        for fi10 := 0; fi10 < 14; fi10++ {
            if true {
                var v11 int64 = -1
                var v12 uint64 = 482
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 30, F1: 654, F2: 100}
            var v14 float64 = 3.14159
            _ = fi10
            _ = v13
            _ = v14
        }
        v2 = v2
        _ = fi8
        _ = v9
    }
    v15 := GenS2{F0: 214, F1: 10}
    v16 := []int{2147483647}
    var v17 bool = false
    var v18 uint64 = 133
    if (((float32(float64(((0.1 + v5) + float32(0.1)))) - 0.0) / v5) >= v5) {
        if (13 >= -106) {
            v19 := GenS1{F0: 0, F1: 1, F2: 42}
            v5 -= (1e10 + (1.5 + v5))
            _ = v19
        } else {
            switch (v2 + v2) {
            case "foo bar":
                v17 = false
            case "":
                var v20 float64 = 3.14159
                _ = v20
            default:
                delete(v4, -1)
            }
        }
    }
    switch 4294967295 {
    case 920:
        var v21 float32 = 1e10
        _ = v21
    case 4294967295:
        var v22 string = (v2 + v2)
        _ = v22
    case -1:
        var v23 uint8 = 128
        _ = v23
    default:
        if (9223372036854775807 > -9223372036854775808) {
            var v24 uint64 = (^(v18))
            var v25 int = 2147483647
            _ = v24
            _ = v25
        } else {
            v5 = 1.5
            var v26 uint8 = 42
            switch -858 {
            case 332:
                for fi27 := 0; fi27 < 7; fi27++ {
                    fmt.Println(v1, v1, fi27)
                    fmt.Println(v17, v18)
                    switch fi27 {
                    case 32767:
                        var v28 uint16 = 32767
                        _ = v28
                    case 869:
                        var v29 string = v2
                        _ = v29
                    }
                    _ = fi27
                }
            }
            _ = v26
        }
    }
    fmt.Println(v1, v5)
    for fi30 := 0; fi30 < 10; fi30++ {
        v18 -= (v18 - v18)
        var v31 float32 = v5
        var v32 uint16 = 257
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 bool = v17
    if v17 {
        v34 := GenS3{F0: -2147483648}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
