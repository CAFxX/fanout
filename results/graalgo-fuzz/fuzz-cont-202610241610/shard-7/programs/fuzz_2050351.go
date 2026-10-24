package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 int8 = 127
    var v4 uint16 = 32767
    v4 = v4
    v5 := GenS1{F0: (v3 - v3), F1: 2, F2: 32767}
    v6 := []int{9223372036854775807, -432}
    var v7 int16 = -744
    if true {
        v6 = append(v6, -519)
        switch ("a" + "hello") {
        case "hello":
            var v8 int8 = v3
            _ = v8
        case "\n\t":
            switch -581 {
            case 32767:
                switch "a" {
                case "\n\t":
                    v7 += (v7 & v7)
                case "a":
                    if true {
                        v2 = 2
                    }
                case "foo bar":
                    var v9 uint64 = 56
                    _ = v9
                default:
                    v4 += (^(v4))
                }
            case -2147483648:
                v1[-150] = false
            default:
                var v10 int32 = 2147483647
                _ = v10
            }
        default:
            v3 -= 0
        }
        if (3.14159 < -2.25) {
            if true {
                var v11 int16 = (v7 + (v7 >> ((uint(v4)) & 63)))
                _ = v11
            }
        }
    }
    if len(v6) > 0 {
        v6[0] = -561
    }
    delete(v1, 2147483647)
    fmt.Println(v3)
    var v12 string = "a"
    v13 := GenS2{F0: 848, F1: 1, F2: 269}
    v1[-296] = true
    if (!(!true)) {
        var v14 int64 = 9223372036854775807
        _ = len(v1)
        _ = v14
    }
    v2 += v2
    v3 = v3
    v15 := GenS3{F0: 18446744073709551615, F1: false}
    switch "" {
    case "a":
        v3 = -69
    case "hello":
        if (243 >= -878) {
            v4 = v4
        } else {
            var v16 int = -995
            _ = v16
        }
    case "":
        v2 = 56
    default:
        fmt.Println(v12)
    }
    var v17 string = v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
