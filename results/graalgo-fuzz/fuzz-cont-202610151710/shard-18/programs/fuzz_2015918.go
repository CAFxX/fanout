package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }

func main() {
    var v1 int = 0
    v2 := make(map[int]uint32)
    var v3 int16 = -186
    var v4 int16 = (v3 ^ (^(v3)))
    v5 := []int{(v1 - (-(v1))), v1, (-(v1)), (v1 << ((65535) & 63))}
    v4 = v4
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 18446744073709551615, F1: -128, F2: (1e10 < 0.0)}
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 18446744073709551615, F1: -128, F2: (1 > (fi8 & 4294967295))}
            var v10 float64 = float64(1.5)
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 float32 = 1e10
        _ = v11
    }
    v4 = v3
    var v12 uint8 = 243
    v13 := GenS1{F0: 2, F1: 1, F2: ((42 > 42) && (v4 < (v4 ^ v4)))}
    fmt.Println(v12)
    var v14 string = ""
    switch 200 {
    case 793:
        if (42 < 1) {
            if ((false && (180 == v12)) || ((v3 * (v3 / (((v4 | v3)) | 1))) < 1)) {
                v12 = v12
            } else {
                var v15 int = 594
                _ = v15
            }
            v14 = v14
        } else {
            if false {
                switch v14 {
                case "":
                    var v16 uint = 892
                    _ = v16
                case "a":
                    var v17 int32 = 0
                    _ = v17
                case "hello":
                    var v18 int64 = 829
                    _ = v18
                default:
                    v12 += 0
                }
            }
        }
    case 1:
        var v19 uint8 = v12
        _ = v19
    case 128:
        switch -955 {
        case 374:
            switch 0 {
            case -137:
                var v20 string = ("\n\t" + (v14 + "hello"))
                _ = v20
            case 32767:
                var v21 string = v14
                _ = v21
            case 0:
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 int16 = 32767
                    v4 += -1
                    _ = fi22
                    _ = v23
                }
            default:
                v1 += (v1 ^ v1)
            }
        }
    }
    var v24 int64 = 42
    v12 = 165
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
