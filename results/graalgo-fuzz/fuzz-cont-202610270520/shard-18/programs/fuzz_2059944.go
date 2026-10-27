package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 int8; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int64; F2 string }
func genFn0() int16 {
    var v6 uint64 = 0
    fmt.Println(v6)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: 42, F1: 1}
        _ = fi7
        _ = v8
    }
    _ = v6
    return -814
}

func main() {
    v1 := []int{32768}
    var v2 float32 = 1.5
    var v3 uint = 4294967295
    v4 := make(map[uint32]int)
    var v5 string = ("a" + "\n\t")
    var v9 int16 = -432
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            if (fi11 != fi11) {
                v2 -= (float32(float64((v2 * v2))) + v2)
            } else {
                if (!((fi10 % ((2147483647) | 1)) < (fi11 - fi11))) {
                    v2 += v2
                    var v12 int32 = -2147483648
                    _ = v12
                }
            }
            _ = fi11
        }
        fmt.Println(v5, v5)
        _ = fi10
    }
    fmt.Println(v3, v9, v5)
    var v13 uint = v3
    var v14 int = 0
    for fi15 := 0; fi15 < 2; fi15++ {
        switch v14 {
        case 1:
            for fi16 := 0; fi16 < 7; fi16++ {
                if false {
                    v14 = v14
                    fmt.Println(v2, v5, v2)
                    var v17 int64 = -9223372036854775808
                    _ = v17
                } else {
                    v2 = v2
                    v1 = append(v1, fi15)
                    v18 := GenS1{F0: 218, F1: 0}
                    _ = v18
                }
                _ = fi16
            }
        case -832:
            switch (-(v14)) {
            case -1:
                v19 := v4[610]
                _ = v19
            case 2147483647:
                v9 -= v9
            case -2147483648:
                var v20 float32 = float32(((1e10 + (float64(-2.25) * 1.5)) / 1e10))
                _ = v20
            default:
                v21 := v4[4294967295]
                _ = v21
            }
        case -2147483648:
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 float64 = float64(0.0)
                v24 := GenS1{F0: v9, F1: (v13 | v3)}
                v25 := GenS2{F0: 1, F1: -1, F2: v3}
                delete(v4, 2)
                var v26 uint32 = 1
                if false {
                    v27 := GenS3{F0: v5, F1: -128, F2: 255}
                    v28 := GenS4{F0: false}
                    v29 := GenS5{F0: (0 < 42), F1: 0, F2: v5}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 int8 = 0
                    v31 := GenS5{F0: false, F1: -9223372036854775808, F2: ((v5 + v5) + (v5 + (v5 + v5)))}
                    v13 -= 18446744073709551615
                    v14 = fi15
                    _ = v30
                    _ = v31
                }
                var v32 uint32 = uint32(v13)
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
                _ = v26
                _ = v32
            }
        }
        if (false && true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch "" {
        case "foo bar":
            fmt.Println(v14, fi15)
        case "":
            v14 += v14
        default:
            fmt.Println(v14)
        }
        _ = fi15
    }
    switch 400 {
    case -1:
        v2 -= float32(1e10)
    case 861:
        fmt.Println(v3)
    default:
        for fi33 := 0; fi33 < 2; fi33++ {
            if false {
                var v34 int32 = 867
                if (!(!(uint8(v3) >= 127))) {
                    var v35 string = v5
                    _ = len(v1)
                    _ = v35
                }
                _ = v34
            }
            _ = fi33
        }
    }
    fmt.Println(v3, v9, v13)
    v1 = append(v1, (-(v14)))
    var v36 uint16 = 65535
    v37 := GenS3{F0: v5, F1: 1, F2: 0}
    _ = len(v4)
    v36 -= (v36 << ((v13) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
