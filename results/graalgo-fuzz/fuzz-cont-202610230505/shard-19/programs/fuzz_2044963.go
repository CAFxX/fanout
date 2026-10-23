package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0() int64 {
    v7 := []int{2147483647, 199}
    if (true && true) {
        var v8 uint = 42
        _ = v8
    } else {
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        if true {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v7 = append(v7, 127)
            switch "\n\t" {
            case "":
                switch -493 {
                case 127:
                    if len(v7) > 0 {
                        v7[0] = 582
                    }
                }
            case "hello":
                var v10 int8 = 127
                _ = v10
            case "foo bar":
                var v11 uint = 743
                _ = v11
            }
        }
        v12 := GenS1{F0: 2147483647, F1: 392, F2: -1}
        _ = v12
    }
    _ = len(v7)
    var v13 uint16 = 255
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: 128, F1: 1, F2: 690}
        _ = len(v7)
        _ = fi14
        _ = v15
    }
    _ = v13
    _ = v7
    return -809
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{384, 32768, -589, -9223372036854775808}
    var v3 int8 = -128
    var v4 int32 = 887
    v5 := make(map[string]int)
    var v6 uint = 578
    var v16 float64 = (3.14159 - 1e10)
    v16 = v16
    if false {
        var v17 string = "hello"
        _ = len(v1)
        _ = v17
    }
    var v18 bool = ((0.1 + (1e10 - 3.14159)) < 1e10)
    v2 = append(v2, -1)
    v19 := GenS1{F0: (v4 | 557), F1: 4294967295, F2: (v3 / ((42) | 1))}
    switch -9223372036854775808 {
    case 1:
        v20 := GenS3{F0: -9223372036854775808, F1: 1}
        _ = v20
    case -1:
        var v21 uint = 0
        _ = v21
    case 9223372036854775807:
        v4 = (^(v4))
    default:
        if (-323 != -1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v16 += v16
        }
    }
    fmt.Println(v6)
    var v22 string = ""
    fmt.Println(v3)
    v2 = append(v2, 68)
    var v23 uint32 = 2
    if (false && ((v6 >> (((v6 / (((v6 << ((uint(v23)) & 63))) | 1))) & 63)) > 2)) {
        v3 -= v3
        if v18 {
            var v24 float32 = (float32(1.5) - (0.0 * ((-2.25 * -2.25) + 1.5)))
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 bool = (!true)
            _ = len(v5)
            var v27 uint16 = 713
            _ = fi25
            _ = v26
            _ = v27
        }
    } else {
        v3 = (-(v3))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
