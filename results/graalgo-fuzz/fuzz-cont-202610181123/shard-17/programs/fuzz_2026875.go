package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int32; F1 int8 }

func main() {
    var v1 string = "a"
    var v2 float32 = (float32(float64((0.1 + 0.1))) / (1e10 * ((0.0 - 1.5) + (1e10 / -2.25))))
    v3 := []uint32{808, 409, 306, 42}
    var v4 bool = (42 <= 9223372036854775807)
    v5 := []uint32{}
    if false {
        if v4 {
            fmt.Println(v1)
            v2 += v2
            v1 = ""
        }
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int64 = 9223372036854775807
            switch "foo bar" {
            case "hello":
                if ((32767 != 155) && (!(!(1 > 42)))) {
                    if false {
                        var v8 uint8 = 255
                        var v9 int = fi6
                        v8 -= v8
                        _ = v8
                        _ = v9
                    }
                    v4 = v4
                    var v10 int = fi6
                    _ = v10
                }
            }
            switch 4294967295 {
            case 0:
                v11 := GenS1{F0: true, F1: 32767}
                _ = v11
            default:
                switch (v1 + v1) {
                case "\n\t":
                    var v12 uint16 = 550
                    _ = v12
                }
            }
            _ = fi6
            _ = v7
        }
        v2 = 3.14159
    } else {
        v13 := GenS2{F0: -691, F1: 0}
        var v14 bool = (1 != -1)
        var v15 int16 = 858
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v1)
    var v16 uint8 = 1
    switch ("\n\t" + v1) {
    case "foo bar":
        var v17 string = "\n\t"
        _ = v17
    case "hello":
        var v18 int16 = 308
        _ = v18
    case "":
        var v19 int16 = 42
        _ = v19
    }
    v20 := GenS3{F0: 42, F1: v16, F2: 1}
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = 1
        }
    }
    if ((3.14159 == (0.0 * 0.0)) && (135 > 1000)) {
        v21 := GenS3{F0: 42, F1: (v16 % ((v16) | 1)), F2: uint16(v16)}
        var v22 int64 = 538
        var v23 int16 = 1
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: 2147483647, F1: 1}
    fmt.Println(v1, v4, v4)
    var v25 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
