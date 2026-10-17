package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 bool; F1 int16; F2 int }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    var v5 uint = 65535
    if (0 <= v4) {
        if (!(2 == 2)) {
            var v6 float64 = -2.25
            _ = v6
        }
        fmt.Println(v5)
    } else {
        var v7 int64 = 144
        var v8 string = ""
        var v9 int16 = 641
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 32768}
    v11 := GenS1{F0: 1}
    v1 = (v1 - v1)
    v12 := []int{9223372036854775807, 732, -1, 50}
    v13 := GenS2{F0: v1, F1: -663}
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            v5 -= v5
            var v15 uint8 = (v1 - (v1 / ((v1) | 1)))
            _ = len(v12)
            _ = fi14
            _ = v15
        }
        fmt.Println(v4, v4)
    }
    _ = len(v12)
    delete(v3, 65535)
    v16 := v3[uint32(v1)]
    v12 = append(v12, v16)
    v5 -= (v5 >> (((v5 % (((v5 / (((v5 >> (((v5 & 4294967295)) & 63))) | 1))) | 1))) & 63))
    for fi17 := 0; fi17 < 17; fi17++ {
        switch "hello" {
        case "a":
            v18 := v3[65535]
            _ = v18
        default:
            fmt.Println(v1)
        }
        _ = fi17
    }
    fmt.Println(v1)
    var v19 uint32 = 588
    var v20 uint16 = 186
    if ((1 != -1) || true) {
        fmt.Println(v5, v19)
        fmt.Println(v4)
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                if true {
                    v16 -= fi22
                } else {
                    v19 = (v19 & v19)
                    v12 = append(v12, (fi21 | fi21))
                    fmt.Println(v16)
                }
                if false {
                    var v23 int8 = -109
                    switch fi21 {
                    case 32767:
                        var v24 uint16 = v20
                        _ = v24
                    default:
                        v25 := GenS2{F0: 42, F1: (-(fi21))}
                        v26 := v2[(fi22 / ((-2147483648) | 1))]
                        var v27 bool = false
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = v23
                }
                _ = fi22
            }
            v28 := GenS3{F0: (885 != 753), F1: 0, F2: (v16 * int(v20))}
            _ = fi21
            _ = v28
        }
    } else {
        v1 -= 42
        var v29 bool = (1 <= -1)
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
