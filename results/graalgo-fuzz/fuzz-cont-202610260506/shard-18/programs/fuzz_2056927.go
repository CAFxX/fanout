package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int32 = -1
    var v2 uint32 = 42
    var v3 int64 = 428
    var v4 string = "hello"
    v5 := []int{881}
    if (!(1 <= 0)) {
        fmt.Println(v4)
        var v6 string = v4
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: -128, F1: false}
                _ = fi7
                _ = v8
            }
        } else {
            if (v1 == -1) {
                var v9 int64 = (-(v3))
                _ = v9
            }
            fmt.Println(v1, v1)
        }
        _ = v6
    } else {
        var v10 int32 = (v1 ^ (v1 | 1))
        v1 = v10
        _ = v10
    }
    v1 -= 1
    switch -128 {
    case 816:
        if (!(-306 == -2147483648)) {
            v2 = (v2 / (((v2 | v2)) | 1))
        } else {
            switch 766 {
            case -974:
                var v11 int8 = -128
                _ = v11
            case -859:
                v1 = 304
            case 127:
                v1 = v1
            default:
                v12 := GenS1{F0: 100, F1: (!((!false) || (true || (v3 <= (v3 + 9223372036854775807)))))}
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                if len(v5) > 0 {
                    v5[0] = fi13
                }
                if len(v5) > 0 {
                    v5[0] = (fi13 + fi13)
                }
                var v14 int64 = v3
                _ = fi13
                _ = v14
            }
            var v15 int = 0
            _ = v15
        }
    case 32768:
        v4 = "\n\t"
    }
    v3 = v3
    fmt.Println(v4)
    v16 := GenS2{F0: 1, F1: (0 != 42)}
    var v17 float64 = 1.5
    v17 -= 0.0
    if ((float64(-2.25) <= 3.14159) || (2147483647 <= 0)) {
        for fi18 := 0; fi18 < 2; fi18++ {
            v17 -= v17
            _ = fi18
        }
        v19 := GenS1{F0: 75, F1: (2 >= 255)}
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        switch fi20 {
        case 32768:
            v4 = v4
        case -9223372036854775808:
            var v21 string = v4
            _ = v21
        }
        var v22 int64 = -1
        var v23 uint32 = v2
        _ = fi20
        _ = v22
        _ = v23
    }
    var v24 bool = (1 == 251)
    fmt.Println(v24, v2, v24)
    v25 := GenS3{F0: (v1 - (v1 + (v1 << ((42) & 63))))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
