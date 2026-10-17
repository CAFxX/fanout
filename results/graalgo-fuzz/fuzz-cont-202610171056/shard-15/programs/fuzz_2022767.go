package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int32; F2 bool }
func genFn0() uint16 {
    var v6 bool = (true || true)
    v6 = (-2147483648 > 2147483647)
    switch 9223372036854775807 {
    case 1:
        if v6 {
            var v7 uint16 = 32767
            _ = v7
        } else {
            v8 := []int{-1, 4294967295}
            v9 := GenS1{F0: 417, F1: 9223372036854775807}
            var v10 string = "hello"
            _ = v10
            _ = v8
            _ = v9
        }
    case 32767:
        v6 = (0.0 >= 0.0)
    }
    v11 := GenS1{F0: 0, F1: -9223372036854775808}
    _ = v11
    _ = v6
    return 42
}

func main() {
    var v1 uint = 255
    var v2 int32 = 42
    v3 := []uint32{42, 661, 843, 4294967295}
    var v4 float64 = (((3.14159 * (0.1 * 3.14159)) * (-2.25 + 0.1)) / 0.1)
    var v5 uint32 = 2147483648
    _ = len(v3)
    fmt.Println(v2, v4)
    v2 -= (v2 - -2147483648)
    v5 += (v5 | v5)
    v2 = v2
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS2{F0: false, F1: 255}
        _ = fi12
        _ = v13
    }
    var v14 float64 = v4
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS3{F0: 153}
        var v17 int = -273
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 -= (v1 - v1)
    var v18 int32 = (v2 + v2)
    var v19 uint = 646
    var v20 uint64 = 18446744073709551615
    if false {
        var v21 int64 = 1
        v22 := GenS3{F0: v20}
        if len(v3) > 0 {
            v3[0] = v5
        }
        _ = v21
        _ = v22
    } else {
        if false {
            fmt.Println(v19, v5)
            for fi23 := 0; fi23 < 4; fi23++ {
                v14 += (v4 / 0.1)
                v24 := GenS3{F0: uint64(v5)}
                _ = fi23
                _ = v24
            }
            switch "hello" {
            case "\n\t":
                if (0 < 4294967295) {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    v19 += v1
                }
            case "foo bar":
                for fi25 := 0; fi25 < 11; fi25++ {
                    v4 = 0.1
                    _ = len(v3)
                    v26 := GenS4{F0: false, F1: v18, F2: (42 < -419)}
                    _ = fi25
                    _ = v26
                }
            case "":
                var v27 int16 = 32767
                _ = v27
            }
        } else {
            v18 = v18
        }
        var v28 float32 = float32((v4 - 0.1))
        _ = v28
    }
    var v29 int8 = 34
    v4 = ((v14 * v14) / (0.0 / (((0.1 + float64(1e10)) * v14) - v4)))
    switch 4294967295 {
    case -359:
        v20 = (v20 >> ((uint(v20)) & 63))
    case 32767:
        v1 -= v1
    default:
        for fi30 := 0; fi30 < 11; fi30++ {
            v4 += (float64(0.1) / -2.25)
            var v31 uint8 = 1
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
