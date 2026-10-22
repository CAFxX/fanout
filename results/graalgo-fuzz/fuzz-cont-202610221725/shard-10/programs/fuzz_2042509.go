package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{-925, 2147483647, 2, -303}
    var v2 int16 = 32767
    var v3 uint16 = 0
    var v4 float64 = 3.14159
    if (248 >= 18446744073709551615) {
        var v5 uint64 = 4294967295
        _ = v5
    } else {
        if (!("\n\t" < (("a" + "") + "a"))) {
            switch 4294967295 {
            case 32768:
                v3 += (v3 >> ((0) & 63))
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            var v6 int32 = -1
            _ = len(v1)
            v7 := GenS1{F0: -2147483648, F1: true}
            _ = v6
            _ = v7
        }
        var v8 int = 32767
        _ = v8
    }
    switch "\n\t" {
    case "a":
        if true {
            switch -9223372036854775808 {
            case 128:
                v1 = append(v1, 32767)
            default:
                if (v3 > (v3 * v3)) {
                    v9 := GenS1{F0: int32(v3), F1: false}
                    _ = v9
                }
            }
            v3 = 771
        }
    case "\n\t":
        var v10 int16 = 825
        _ = v10
    }
    if (!false) {
        v1 = append(v1, -924)
        fmt.Println(v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        switch ("hello" + "hello") {
        case "a":
            v11 := GenS1{F0: 140, F1: (4294967295 == 18446744073709551615)}
            _ = v11
        default:
            switch -541 {
            case 680:
                var v12 int16 = (v2 ^ v2)
                _ = v12
            default:
                var v13 uint16 = (v3 % (((v3 + v3)) | 1))
                _ = v13
            }
        }
        var v14 uint64 = 309
        _ = v14
    }
    var v15 uint = 65535
    var v16 int32 = 0
    if false {
        v17 := GenS2{F0: -2147483648}
        _ = v17
    }
    v4 -= -2.25
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
