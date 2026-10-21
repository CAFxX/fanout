package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 1
    v3 := []uint32{4294967295}
    var v4 int64 = -728
    fmt.Println(v4, v4)
    v5 := GenS1{F0: 32768, F1: 32767, F2: 205}
    v2 = v2
    for fi6 := 0; fi6 < 15; fi6++ {
        _ = len(v1)
        _ = fi6
    }
    switch "\n\t" {
    case "\n\t":
        v7 := GenS1{F0: 32767, F1: 1, F2: 32767}
        _ = v7
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                v9 := GenS2{F0: v2}
                v10 := GenS2{F0: v2}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    v4 = v4
    switch 4294967295 {
    case 32767:
        switch 9223372036854775807 {
        case 128:
            var v11 string = ""
            _ = v11
        case -2147483648:
            v4 = (v4 + v4)
        case 2147483647:
            if false {
                var v12 int8 = 42
                v13 := GenS2{F0: (^(v2))}
                var v14 uint32 = 243
                _ = v12
                _ = v13
                _ = v14
            }
        default:
            _ = len(v3)
        }
    default:
        for fi15 := 0; fi15 < 13; fi15++ {
            v4 = (^(v4))
            _ = fi15
        }
    }
    v4 = v4
    var v16 uint = 2
    if ((v2 > (v2 >> (((v16 & (^(v16)))) & 63))) && false) {
        var v17 int8 = 42
        var v18 uint8 = 128
        _ = v17
        _ = v18
    } else {
        switch "\n\t" {
        case "a":
            switch 42 {
            case 1:
                if true {
                    var v19 int32 = 42
                    var v20 uint = v16
                    v21 := GenS3{F0: -101, F1: (v4 - (v4 >> ((v16) & 63))), F2: (v20 % ((v16) | 1))}
                    _ = v19
                    _ = v20
                    _ = v21
                }
            case -1:
                fmt.Println(v2, v16)
            case 235:
                switch 1 {
                case 127:
                    fmt.Println(v16, v4, v16)
                }
            default:
                v4 += -1
            }
        case "hello":
            if (!false) {
                v22 := GenS1{F0: 4294967295, F1: 301, F2: -2147483648}
                var v23 bool = (127 >= 228)
                _ = v22
                _ = v23
            }
        case "\n\t":
            v4 = v4
        default:
            v3 = append(v3, 675)
        }
        v4 = 85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
