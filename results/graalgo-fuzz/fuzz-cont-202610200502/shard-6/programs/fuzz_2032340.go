package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint; F2 int8 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 float64 = ((0.0 / ((-2.25 + 3.14159) + 3.14159)) + 1e10)
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v1 -= ((((v1 * v1) + 0.0) / (v1 * v1)) / (float64(float32(v1)) / 1.0))
    if (1 == -893) {
        var v5 int32 = 1
        var v6 float64 = v1
        _ = v5
        _ = v6
    }
    v1 = v1
    switch 2147483647 {
    case 32767:
        if (-32768 != 32767) {
            var v7 int8 = -1
            _ = v7
        }
    case -2147483648:
        fmt.Println(v1)
    case 32768:
        v8 := GenS1{F0: -81, F1: 255, F2: -977}
        _ = v8
    default:
        v9 := GenS2{F0: -1, F1: -81, F2: 657}
        _ = v9
    }
    if (false && true) {
        v1 -= (v1 + ((v1 + float64((1.5 / 0.1))) + (v1 - v1)))
        for fi10 := 0; fi10 < 18; fi10++ {
            switch (fi10 / (((fi10 ^ fi10)) | 1)) {
            case 568:
                if false {
                    var v11 bool = true
                    var v12 uint16 = 2
                    v11 = ((((0 == -1) || v11) || (!false)) && v11)
                    _ = v11
                    _ = v12
                }
            case -9223372036854775808:
                delete(v3, 0)
            case 32767:
                fmt.Println(v1, v1)
            }
            _ = fi10
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 float32 = (0.0 * 3.14159)
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    switch 128 {
    case -197:
        v1 += (((v1 * (v1 - (float64(-2.25) - v1))) + 1.5) + ((-2.25 * (v1 / (0.1 * (((float64((0.0 + 1e10)) + (3.14159 / v1)) * (v1 + 0.1)) + float64((0.1 * 0.1)))))) + (float64((-2.25 / -2.25)) + (v1 * ((v1 / (v1 / (0.0 - ((-2.25 / (v1 * v1)) - v1)))) - (v1 + v1))))))
    case 2147483647:
        v15 := []int{128}
        _ = v15
    case -1:
        v16 := GenS3{F0: 608, F1: 1, F2: 100}
        _ = v16
    default:
        v17 := GenS4{F0: 65535, F1: -598}
        _ = v17
    }
    var v18 int64 = 1
    v19 := GenS2{F0: (^(v18)), F1: 100, F2: (v18 & v18)}
    var v20 float64 = (0.0 + 0.1)
    if ("hello" > "foo bar") {
        for fi21 := 0; fi21 < 14; fi21++ {
            v20 -= 0.0
            v22 := GenS2{F0: v18, F1: 65, F2: (v18 + v18)}
            _ = fi21
            _ = v22
        }
    }
    var v23 uint16 = 0
    var v24 uint16 = (v23 % ((24) | 1))
    for fi25 := 0; fi25 < 9; fi25++ {
        v20 = v1
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 uint32 = 1
            if (!(!false)) {
                _ = len(v4)
                if (!true) {
                    v24 -= 2
                }
                v28 := GenS4{F0: 18446744073709551615, F1: v18}
                _ = v28
            } else {
                v20 = (0.0 - 0.1)
            }
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    v29 := GenS3{F0: 42, F1: 42, F2: -84}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
}
