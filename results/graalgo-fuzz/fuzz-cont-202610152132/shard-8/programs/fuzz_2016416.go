package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 int64; F2 bool }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []int{32768, -526, 0}
    var v2 int64 = 26
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    var v5 uint64 = (v3 & v3)
    switch -9223372036854775808 {
    case 32768:
        v5 -= 1
    case -266:
        v6 := v4[2]
        _ = v6
    case 4294967295:
        v2 += v2
    default:
        _ = len(v4)
    }
    v3 = v3
    v7 := v4[923]
    v5 -= 242
    var v8 int16 = -1
    fmt.Println(v8, v5)
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(fi9, v3, v7)
        if ((2 != 18446744073709551615) || (!(!(660 != 4294967295)))) {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: "a"}
                var v12 float32 = -2.25
                if true {
                    fmt.Println(v12, fi9, v12)
                    var v13 int = fi10
                    v8 = int16(v3)
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            var v14 int32 = 2147483647
            var v15 uint16 = 1
            v16 := GenS1{F0: ("" + "")}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    var v17 string = ""
    switch (v7 / (((v7 ^ -1)) | 1)) {
    case 867:
        v17 = "hello"
    case -757:
        if false {
            var v18 int16 = (v8 * 0)
            if (32768 <= (v7 * v7)) {
                switch v17 {
                case "a":
                    fmt.Println(v8, v3, v5)
                case "":
                    v17 = v17
                default:
                    v19 := GenS1{F0: "hello"}
                    _ = v19
                }
                v20 := GenS2{F0: 2, F1: (-(v7)), F2: 305}
                v21 := GenS3{F0: (v7 & -387), F1: (v17 + v17), F2: 9223372036854775807}
                _ = v20
                _ = v21
            }
            v8 += v18
            _ = v18
        } else {
            v22 := v4[938]
            fmt.Println(v22, v17)
            switch v17 {
            case "":
                for fi23 := 0; fi23 < 15; fi23++ {
                    if (42 < 42) {
                        v24 := GenS4{F0: 1, F1: -9223372036854775808, F2: (!(0 >= 140))}
                        v25 := GenS3{F0: (fi23 / ((fi23) | 1)), F1: (v17 + v17), F2: int(v2)}
                        var v26 int = (v7 | (v22 ^ v22))
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 uint = 816
                    v28 := GenS5{F0: (v5 << ((v27) & 63))}
                    _ = fi23
                    _ = v27
                    _ = v28
                }
            default:
                v5 = 2
            }
            _ = v22
        }
    }
    var v29 uint64 = (v5 | 18446744073709551615)
    switch "hello" {
    case "\n\t":
        var v30 float64 = (float64(1e10) / 3.14159)
        _ = v30
    case "a":
        v31 := GenS5{F0: (v5 ^ 0)}
        _ = v31
    default:
        for fi32 := 0; fi32 < 5; fi32++ {
            fmt.Println(v2)
            _ = fi32
        }
    }
    v33 := GenS3{F0: -2147483648, F1: ("hello" + v17), F2: v7}
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
