package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 0
    v2 := []uint32{2147483647, 2147483647}
    v3 := []uint32{361, 0, 885, 2147483648}
    var v4 float64 = 1e10
    var v5 float64 = (v4 * v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 int = 690
    v7 := GenS1{F0: 546, F1: 42}
    if (int32(v1) > -1) {
        var v8 float64 = v5
        var v9 int32 = 0
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS1{F0: 376, F1: 42}
            if ((447 != 848) && (!true)) {
                if (9223372036854775807 != -9223372036854775808) {
                    v4 -= (float64((1.5 / (float32(-2.25) - 0.1))) + -2.25)
                    switch "foo bar" {
                    case "":
                        var v12 bool = false
                        _ = v12
                    case "a":
                        v9 = (v9 | v9)
                    case "foo bar":
                        v2 = append(v2, 65535)
                    }
                    fmt.Println(v4, v1)
                }
            }
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        v1 += 0
    }
    v13 := GenS2{F0: 4294967295, F1: 1}
    switch ("hello" + "foo bar") {
    case "hello":
        v6 = -1
    case "":
        v14 := GenS3{F0: 0, F1: 1}
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 float32 = float32(v4)
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS1{F0: 326, F1: 1}
    switch 128 {
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = 65535
        }
    case 4294967295:
        v5 = (0.0 * 1e10)
    case 32767:
        var v18 uint16 = (v1 ^ (v1 * v1))
        _ = v18
    default:
        fmt.Println(v1, v6, v6)
    }
    v1 += (v1 % (((v1 & v1)) | 1))
    v5 = v4
    if (!(v6 >= -229)) {
        switch "a" {
        case "":
            var v19 int = v6
            _ = v19
        default:
            var v20 uint64 = 109
            _ = v20
        }
    } else {
        var v21 float64 = float64(-2.25)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
