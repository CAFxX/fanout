package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0(p0 bool, p1 uint, p2 int8) uint64 {
    var v6 int16 = -267
    v6 = (v6 >> (((p1 << ((p1) & 63))) & 63))
    if p0 {
        var v7 uint64 = 4294967295
        if (0 != (^(v7))) {
            switch -156 {
            case 42:
                switch -1 {
                case 4294967295:
                    p1 = p1
                case 127:
                    p2 += p2
                default:
                    var v8 uint8 = 127
                    _ = v8
                }
            case -258:
                var v9 uint64 = v7
                _ = v9
            case -412:
                fmt.Println(p1, v6, p1)
            }
            var v10 string = ("a" + ("\n\t" + "a"))
            _ = v10
        } else {
            v11 := GenS1{F0: (p0 || (!((^(v7)) != 524))), F1: -1}
            _ = v11
        }
        p1 = (^(p1))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 952
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int8 = 114
    v5 := make(map[uint32]int)
    if false {
        v2[0] = 873
        v4 -= (v4 / ((-74) | 1))
        v1 = v1
    }
    v1 += 0.1
    v12 := GenS2{F0: 778}
    if ((("foo bar" + "\n\t") < ("a" + "hello")) && (!true)) {
        v4 += -1
        fmt.Println(v4, v1)
    } else {
        fmt.Println(v4, v4)
        var v13 float32 = 0.0
        var v14 uint16 = 510
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v1)
    v15 := GenS1{F0: (v4 < (v4 / (((v4 | 100)) | 1))), F1: 2147483647}
    if true {
        v16 := GenS1{F0: ((!(-698 >= 1)) && false), F1: -1}
        _ = v16
    }
    var v17 int8 = v4
    if (4294967295 >= 4294967295) {
        v18 := v3[42]
        if (1000 < 1000) {
            var v19 int32 = -2147483648
            _ = v19
        }
        _ = v18
    } else {
        fmt.Println(v1, v4)
        v20 := GenS3{F0: 331, F1: -16}
        _ = len(v3)
        _ = v20
    }
    var v21 int8 = v17
    switch 32768 {
    case 42:
        var v22 uint32 = 4294967295
        _ = v22
    case -2147483648:
        v23 := GenS2{F0: 918}
        _ = v23
    case -9223372036854775808:
        var v24 int64 = -412
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
