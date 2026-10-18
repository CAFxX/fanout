package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 string }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int = 616
    var v2 uint8 = 2
    v3 := make(map[uint32]uint32)
    v2 = (v2 / ((v2) | 1))
    for fi4 := 0; fi4 < 10; fi4++ {
        if (("\n\t" + "") <= "") {
            v5 := GenS1{F0: 42}
            for fi6 := 0; fi6 < 14; fi6++ {
                fmt.Println(v1, v1)
                v7 := GenS2{F0: "foo bar", F1: true}
                _ = fi6
                _ = v7
            }
            v8 := v3[2147483647]
            _ = v5
            _ = v8
        } else {
            v2 = (v2 * v2)
            var v9 int32 = -710
            var v10 int = v1
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    var v11 uint64 = 2
    fmt.Println(v11)
    var v12 int32 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS3{F0: (v11 * v11), F1: 155}
        if (!(int64(v11) != 0)) {
            switch v1 {
            case 1:
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := GenS4{F0: "\n\t", F1: (^(v12)), F2: ""}
                    _ = fi15
                    _ = v16
                }
            case 0:
                if (1e10 == (0.1 + 3.14159)) {
                    v17 := GenS1{F0: 4294967295}
                    if true {
                        var v18 float64 = 1e10
                        v19 := GenS2{F0: "hello", F1: false}
                        var v20 uint64 = v11
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v1, v2, v2)
                    }
                    _ = v17
                } else {
                    v3[972] = 671
                    fmt.Println(v12, v2, v12)
                }
            }
        }
        var v21 int = (^(v1))
        _ = fi13
        _ = v14
        _ = v21
    }
    v22 := []int{128, v1, -2147483648, (v1 << ((255) & 63))}
    v23 := GenS5{F0: v2}
    v2 += v2
    v12 -= v12
    switch (v1 - v1) {
    case 128:
        var v24 float64 = (0.1 + 0.0)
        _ = v24
    case -753:
        v22 = append(v22, v1)
    case 144:
        var v25 int8 = 100
        _ = v25
    default:
        switch v1 {
        case 127:
            switch (v1 % ((v1) | 1)) {
            case 128:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
        case -9223372036854775808:
            var v26 uint32 = 94
            _ = v26
        case -394:
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                if false {
                    v12 = (^(v12))
                } else {
                    v28 := GenS5{F0: uint8(v11)}
                    v29 := GenS4{F0: ("a" + "\n\t"), F1: v12, F2: ("foo bar" + "a")}
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
        }
    }
    v11 += (^(v11))
    if len(v22) > 0 {
        v22[0] = (v1 >> ((2) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
