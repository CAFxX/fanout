package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]uint32)
    v3 := []int{179, -1}
    v4 := []int{}
    v5 := make(map[uint32]bool)
    v6 := v5[42]
    switch ("" + "foo bar") {
    case "\n\t":
        v7 := GenS1{F0: 506, F1: 0}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 19; fi8++ {
            v6 = v6
            switch fi8 {
            case 198:
                var v9 int64 = 0
                _ = v9
            case -1:
                v10 := v5[42]
                _ = v10
            default:
                switch "foo bar" {
                case "":
                    v1 -= 1
                case "hello":
                    if len(v4) > 0 {
                        v4[0] = fi8
                    }
                case "\n\t":
                    v11 := GenS1{F0: 42, F1: 42}
                    _ = v11
                default:
                    if v6 {
                        var v12 uint32 = 0
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v12 = 877
                            _ = fi13
                        }
                        var v14 string = "\n\t"
                        _ = v12
                        _ = v14
                    }
                }
            }
            _ = fi8
        }
    }
    v15 := GenS1{F0: 87, F1: 4294967295}
    if (0 <= 42) {
        v2[1] = 1
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS2{F0: v1, F1: 1}
            var v18 int32 = -1
            _ = fi16
            _ = v17
            _ = v18
        }
        switch -2147483648 {
        case 2147483647:
            var v19 uint8 = 42
            _ = v19
        case 447:
            v6 = v6
        case 4294967295:
            var v20 int8 = 86
            _ = v20
        }
    }
    v21 := v5[0]
    v1 -= 65535
    var v22 int8 = 42
    v23 := GenS3{F0: 1, F1: 2147483647, F2: 127}
    var v24 uint8 = 42
    v6 = (!true)
    var v25 float64 = 0.0
    fmt.Println(v1, v24, v1)
    switch 203 {
    case -9223372036854775808:
        v26 := GenS4{F0: 0, F1: -43}
        _ = v26
    case -29:
        v4 = append(v4, 460)
    case 128:
        if ((v24 >> ((v1) & 63)) < (v24 * (v24 & (v24 ^ 66)))) {
            v1 = v1
        }
    }
    v27 := v5[42]
    var v28 uint64 = 695
    v29 := GenS2{F0: (v1 % ((1) | 1)), F1: uint8(v28)}
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint = (v1 / ((v1) | 1))
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
