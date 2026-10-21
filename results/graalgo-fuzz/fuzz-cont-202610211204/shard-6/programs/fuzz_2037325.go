package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    var v4 int8 = -128
    var v5 uint8 = 255
    _ = p0
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 int8 = 0
    var v2 uint = 2
    var v3 int16 = 42
    v6 := GenS1{F0: 42, F1: 1, F2: 2147483647}
    v7 := []int{9223372036854775807, -1, 42, 127}
    fmt.Println(v2, v3)
    for fi8 := 0; fi8 < 5; fi8++ {
        v2 -= v2
        var v9 uint8 = 2
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v1)
        v2 += v2
        v1 = 124
        _ = fi10
    }
    _ = len(v7)
    v3 = v3
    fmt.Println(v1, v2)
    v1 = v1
    fmt.Println(v2, v3, v2)
    fmt.Println(v2, v2)
    if (false && (1 >= -2147483648)) {
        v11 := GenS2{F0: v1}
        _ = v11
    }
    switch -1 {
    case 642:
        if (0 != 1) {
            var v12 int8 = v1
            if true {
                fmt.Println(v12, v12)
                v1 -= (v12 + v12)
            } else {
                if false {
                    var v13 int = int(v2)
                    v14 := GenS3{F0: (v3 % ((32767) | 1)), F1: v2}
                    v2 -= (v2 - 2)
                    _ = v13
                    _ = v14
                } else {
                    var v15 bool = ((v12 >> ((1) & 63)) != (v12 / ((v1) | 1)))
                    var v16 bool = v15
                    _ = v15
                    _ = v16
                }
            }
            _ = v12
        } else {
            if (-9223372036854775808 == -1) {
                var v17 int = 376
                v1 += 0
                var v18 int = v17
                _ = v17
                _ = v18
            } else {
                v19 := GenS1{F0: 1, F1: 88, F2: -9223372036854775808}
                var v20 uint64 = 18446744073709551615
                _ = v19
                _ = v20
            }
        }
    case 32768:
        fmt.Println(v3)
    case 2147483647:
        switch -9223372036854775808 {
        case -2147483648:
            v1 = (v1 - (v1 ^ 42))
        case 9223372036854775807:
            for fi21 := 0; fi21 < 19; fi21++ {
                if len(v7) > 0 {
                    v7[0] = fi21
                }
                switch "" {
                case "\n\t":
                    v1 += -128
                case "hello":
                    v7 = append(v7, fi21)
                }
                v2 = v2
                _ = fi21
            }
        default:
            switch 1 {
            case 953:
                switch 42 {
                case 103:
                    v22 := GenS1{F0: 42, F1: 127, F2: 4294967295}
                    _ = v22
                }
            default:
                var v23 int16 = (v3 / ((v3) | 1))
                _ = v23
            }
        }
    default:
        var v24 uint8 = 188
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 uint16 = 285
        _ = len(v7)
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
