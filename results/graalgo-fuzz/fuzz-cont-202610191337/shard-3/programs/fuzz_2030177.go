package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 int8 = -1
    var v4 int16 = 42
    var v5 int32 = 2147483647
    v6 := make(map[int]string)
    var v7 uint16 = 399
    if true {
        switch -34 {
        case 32767:
            var v8 uint = 18446744073709551615
            _ = v8
        default:
            v3 += v3
        }
    } else {
        v4 += 0
    }
    switch 128 {
    case 128:
        var v9 uint32 = 2147483648
        _ = v9
    case 72:
        var v10 uint16 = (v7 % (((v7 * v7)) | 1))
        _ = v10
    case 761:
        v11 := GenS1{F0: (v3 * (^(v3))), F1: 1, F2: v3}
        _ = v11
    default:
        switch "foo bar" {
        case "\n\t":
            v4 = (v4 / ((v4) | 1))
        case "hello":
            for fi12 := 0; fi12 < 13; fi12++ {
                if true {
                    v13 := GenS1{F0: (v3 ^ v3), F1: 1, F2: (v3 >> ((18446744073709551615) & 63))}
                    var v14 int64 = 449
                    var v15 bool = false
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        case "a":
            var v16 bool = (391 == 121)
            _ = v16
        }
    }
    var v17 int16 = -705
    v5 += (v5 >> ((0) & 63))
    v18 := []int{128, -2147483648}
    var v19 int = 32768
    var v20 string = "a"
    var v21 uint32 = 402
    if len(v18) > 0 {
        v18[0] = v19
    }
    fmt.Println(v3, v20, v7)
    var v22 uint = 18446744073709551615
    if ((!(!true)) && false) {
        v23 := GenS2{F0: v21}
        _ = v23
    } else {
        var v24 bool = false
        var v25 uint64 = 1
        var v26 int16 = 1
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
