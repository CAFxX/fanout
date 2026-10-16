package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 string = ("" + "a")
    var v3 string = "a"
    var v4 uint8 = 0
    var v5 int8 = 62
    var v6 uint64 = 0
    var v7 float32 = float32(float64(1.5))
    for fi8 := 0; fi8 < 8; fi8++ {
        switch v3 {
        case "hello":
            var v9 float32 = 0.1
            _ = v9
        }
        _ = fi8
    }
    if true {
        switch -9223372036854775808 {
        case 2147483647:
            _ = len(v1)
        case -2147483648:
            v10 := GenS1{F0: 42, F1: v4}
            _ = v10
        }
        v3 = ("hello" + (v3 + v2))
    }
    v6 = v6
    var v11 int8 = (v5 / (((v5 >> ((4294967295) & 63))) | 1))
    switch 4294967295 {
    case 127:
        var v12 string = v3
        _ = v12
    default:
        if ((-2.25 == 1.5) && (true && ((264 != 2) || true))) {
            if false {
                v11 += -1
                v7 = ((0.0 - -2.25) / v7)
                v2 = ((v2 + v2) + v2)
            } else {
                if (!false) {
                    if (v4 >= v4) {
                        var v13 int16 = -1
                        var v14 int = -1
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint8 = 2
                        _ = v15
                    }
                    v16 := GenS2{F0: 32768, F1: (v2 + ""), F2: v3}
                    _ = v16
                }
                v17 := GenS3{F0: 1, F1: 1, F2: 2147483647}
                v18 := GenS3{F0: 9223372036854775807, F1: 32767, F2: 4294967295}
                _ = v17
                _ = v18
            }
            _ = len(v1)
        }
    }
    var v19 float64 = (1e10 - 0.0)
    v20 := GenS4{F0: true}
    v3 = v3
    var v21 string = ("" + v3)
    _ = len(v1)
    if false {
        var v22 uint16 = 1
        _ = v22
    }
    delete(v1, 139)
    var v23 string = (("a" + v2) + ("a" + v21))
    v7 -= v7
    v4 = 206
    var v24 int32 = -84
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
