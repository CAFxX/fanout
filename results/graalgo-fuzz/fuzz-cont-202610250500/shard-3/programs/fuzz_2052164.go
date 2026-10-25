package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 uint8 = 42
    v5 := GenS1{F0: 1, F1: 2147483648, F2: "hello"}
    v4 = v4
    v6 := v1[1]
    v6 = (v6 >> ((749) & 63))
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int16 = -293
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := v1[983]
            switch (v10 & (-(fi7))) {
            case 4294967295:
                var v11 uint64 = 0
                _ = v11
            case 842:
                v6 = (fi7 | (^(fi7)))
            case -969:
                if false {
                    v12 := GenS1{F0: -9223372036854775808, F1: 373, F2: "\n\t"}
                    var v13 int32 = -2147483648
                    var v14 bool = true
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    var v15 uint8 = 0
                    _ = v15
                }
            default:
                var v16 int32 = 969
                _ = v16
            }
            v6 = v10
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v17 := []int{-9223372036854775808, (v6 * 663)}
    if true {
        switch "hello" {
        case "hello":
            v4 += 11
        case "":
            v18 := GenS2{F0: v4, F1: 0}
            _ = v18
        default:
            v19 := GenS3{F0: 1}
            _ = v19
        }
        v20 := GenS4{F0: 42}
        var v21 bool = (v4 < (v4 * 1))
        _ = v20
        _ = v21
    } else {
        var v22 int8 = 94
        _ = len(v2)
        for fi23 := 0; fi23 < 20; fi23++ {
            v3 -= v3
            var v24 int8 = -103
            v4 -= 30
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    if (v3 != v3) {
        var v25 string = "\n\t"
        fmt.Println(v6)
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 uint16 = 2
            if false {
                if false {
                    var v28 int64 = 1
                    _ = v28
                }
                v4 -= (v4 * (v4 >> ((4294967295) & 63)))
                if false {
                    var v29 uint16 = v27
                    _ = v29
                }
            }
            v30 := GenS3{F0: (v27 % ((v27) | 1))}
            _ = fi26
            _ = v27
            _ = v30
        }
        _ = v25
    }
    v31 := v2[v6]
    for fi32 := 0; fi32 < 6; fi32++ {
        var v33 int16 = v3
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
