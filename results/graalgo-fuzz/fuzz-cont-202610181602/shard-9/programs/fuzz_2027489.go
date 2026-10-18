package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 uint64) int64 {
    fmt.Println(p0)
    if false {
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 uint = p0
            _ = fi5
            _ = v6
        }
        v7 := []int{-1, 203, 32768}
        _ = v7
    }
    if ((42 <= -802) && true) {
        v8 := []int{}
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -501
    var v3 uint64 = 463
    var v4 int = (v2 % ((v2) | 1))
    v3 += v3
    if true {
        fmt.Println(v3, v2, v2)
    } else {
        if (("foo bar" + "hello") <= (("" + "\n\t") + "\n\t")) {
            fmt.Println(v2)
            switch (v4 - v4) {
            case 0:
                v9 := v1[842]
                _ = v9
            }
            if (!(-119 >= 42)) {
                var v10 int = v4
                var v11 string = "\n\t"
                v12 := GenS1{F0: 731, F1: genFn0(126, (v3 / (((v3 ^ v3)) | 1))), F2: 42}
                _ = v10
                _ = v11
                _ = v12
            } else {
                v2 = -617
            }
        } else {
            var v13 int32 = 42
            _ = v13
        }
    }
    v4 = (-(v4))
    v14 := GenS2{F0: 42, F1: -702, F2: 285}
    var v15 string = "hello"
    v2 = v2
    switch (-(v2)) {
    case -678:
        if true {
            var v16 string = ("a" + "")
            if (2147483647 != 2147483647) {
                var v17 uint64 = 817
                fmt.Println(v2)
                _ = v17
            }
            _ = v16
        } else {
            switch 2147483647 {
            case -688:
                switch 0 {
                case 9223372036854775807:
                    v18 := []int{(v4 << ((4294967295) & 63)), 163, (v2 | v2), (v2 & 42)}
                    _ = v18
                }
            case 2147483647:
                _ = len(v1)
            case -853:
                var v19 float32 = ((((float32(((1e10 - 3.14159) + 0.1)) * -2.25) * ((0.0 / 3.14159) * 0.0)) * -2.25) * (-2.25 + 0.0))
                _ = v19
            default:
                v20 := GenS1{F0: 1, F1: 364, F2: 89}
                _ = v20
            }
            var v21 uint64 = v3
            v22 := v1[65535]
            _ = v21
            _ = v22
        }
    case 1:
        var v23 int8 = 107
        _ = v23
    case -2147483648:
        var v24 float32 = float32((1e10 * (0.0 - 0.1)))
        _ = v24
    default:
        v25 := v1[2147483647]
        _ = v25
    }
    fmt.Println(v3)
    v26 := []int{(^(v2)), 0, (v4 & 42), 1}
    fmt.Println(v15, v4)
    var v27 int = v4
    if true {
        var v28 int = (v27 % ((v4) | 1))
        _ = v28
    }
    _ = len(v1)
    v1[131] = 65535
    v29 := GenS3{F0: (true && ((3.14159 * 0.0) >= 0.0))}
    var v30 uint8 = 241
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
