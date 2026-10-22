package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64; F2 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]bool)
    var v3 string = "\n\t"
    var v4 int64 = -9223372036854775808
    var v5 bool = false
    var v6 string = (v3 + "")
    v4 += v4
    v4 -= v4
    for fi7 := 0; fi7 < 14; fi7++ {
        if (18446744073709551615 == 0) {
            v8 := GenS1{F0: 42, F1: v1, F2: 4294967295}
            _ = v8
        } else {
            fmt.Println(v1)
        }
        _ = fi7
    }
    var v9 int8 = 0
    v10 := GenS2{F0: 9223372036854775807}
    if (!v5) {
        switch 78 {
        case -2147483648:
            v11 := GenS2{F0: 127}
            _ = v11
        case 4294967295:
            for fi12 := 0; fi12 < 20; fi12++ {
                switch (fi12 | 9223372036854775807) {
                case 9223372036854775807:
                    var v13 int16 = 0
                    _ = v13
                case 32767:
                    v5 = (1.5 != (0.0 + 3.14159))
                case -1:
                    v14 := []int{(fi12 ^ fi12), -2147483648, (fi12 >> ((2) & 63)), (fi12 >> ((286) & 63))}
                    _ = v14
                }
                _ = fi12
            }
        case -12:
            v15 := GenS3{F0: 42, F1: (v4 | (-(v4))), F2: 582}
            _ = v15
        }
        switch -225 {
        case 9223372036854775807:
            v16 := []int{32768, int(v1), 42}
            _ = v16
        case -864:
            v17 := GenS1{F0: 585, F1: (v1 | (-(v9))), F2: 4294967295}
            _ = v17
        case 940:
            v3 = "\n\t"
        }
        fmt.Println(v4, v5)
    }
    v18 := GenS4{F0: 4294967295}
    fmt.Println(v3, v3, v1)
    v19 := GenS3{F0: 0, F1: -9223372036854775808, F2: 255}
    var v20 uint32 = 0
    switch 32768 {
    case -970:
        switch 1 {
        case 2147483647:
            var v21 uint32 = v20
            _ = v21
        default:
            v22 := GenS3{F0: 1, F1: (v4 / ((int64(v20)) | 1)), F2: 0}
            _ = v22
        }
    }
    if v5 {
        v23 := GenS1{F0: 42, F1: -128, F2: v20}
        _ = v23
    }
    var v24 uint64 = 0
    var v25 uint8 = 42
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
