package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint = 4294967295
    var v2 uint8 = 1
    v3 := make(map[int]string)
    var v4 int = -9223372036854775808
    var v5 uint = 255
    var v6 uint8 = 93
    fmt.Println(v4, v6, v6)
    var v7 uint32 = 588
    v8 := GenS1{F0: 0, F1: 865, F2: 943}
    switch "" {
    case "":
        fmt.Println(v5, v7, v6)
    case "\n\t":
        var v9 float64 = 1e10
        _ = v9
    case "hello":
        var v10 uint16 = 65535
        _ = v10
    }
    if ((v7 & (v7 / (((v7 % (((^(v7))) | 1))) | 1))) != 575) {
        if (128 <= 255) {
            v11 := GenS2{F0: 0, F1: v4}
            v5 -= (v1 >> ((v1) & 63))
            v12 := GenS3{F0: -1, F1: 42}
            _ = v11
            _ = v12
        } else {
            var v13 uint = 65535
            for fi14 := 0; fi14 < 2; fi14++ {
                v13 = 2
                _ = fi14
            }
            var v15 float32 = (3.14159 * (1.5 * 1e10))
            _ = v13
            _ = v15
        }
        if (910 >= (^(v4))) {
            switch "" {
            case "\n\t":
                v16 := GenS2{F0: 1, F1: 42}
                _ = v16
            case "a":
                v2 = v2
            default:
                v4 = v4
            }
            var v17 uint16 = 822
            v18 := GenS1{F0: -288, F1: 32767, F2: int64(v4)}
            _ = v17
            _ = v18
        } else {
            v7 += (v7 % ((320) | 1))
        }
    }
    v7 = 496
    switch v4 {
    case 32768:
        switch (^(v4)) {
        case 32767:
            v19 := []int{923, (-(v4)), v4, (v4 >> (((v1 / ((844) | 1))) & 63))}
            _ = v19
        case -9223372036854775808:
            v7 += uint32(v5)
        case -2147483648:
            if (2 == (v7 | (v7 << ((0) & 63)))) {
                v20 := []int{(-(v4))}
                switch "hello" {
                case "hello":
                    var v21 int = v4
                    _ = v21
                case "foo bar":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint8 = (v6 >> ((18446744073709551615) & 63))
                        var v24 int64 = 0
                        v25 := v3[(fi22 | fi22)]
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                case "\n\t":
                    v26 := GenS1{F0: -32768, F1: 1000, F2: 1}
                    _ = v26
                default:
                    v6 = (v6 + v6)
                }
                _ = v20
            }
        }
    case 1:
        var v27 int8 = 32
        _ = v27
    }
    if false {
        fmt.Println(v2, v5, v4)
        fmt.Println(v4, v1, v2)
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 int = fi28
            for fi30 := 0; fi30 < 8; fi30++ {
                if (v7 < (v7 - v7)) {
                    v31 := GenS3{F0: 2147483647, F1: -9223372036854775808}
                    _ = v31
                } else {
                    v5 = (^(v1))
                    var v32 uint = v5
                    var v33 uint32 = v7
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
            var v34 int = 32767
            _ = fi28
            _ = v29
            _ = v34
        }
    }
    var v35 float64 = (1e10 / 0.1)
    var v36 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
