package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 int16, p1 int32, p2 uint64) int {
    var v5 uint = 2
    v6 := []int{int(p2), 4294967295, int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 481
}

func main() {
    var v1 int16 = 773
    var v2 uint = 2
    v3 := make(map[int]uint32)
    var v4 string = ("\n\t" + "hello")
    fmt.Println(v4)
    delete(v3, 127)
    var v7 float32 = 0.1
    v8 := GenS1{F0: 9223372036854775807, F1: 32768, F2: "foo bar"}
    if ((uint8(v2) <= 255) && true) {
        v4 = v4
        fmt.Println(v4, v7)
    } else {
        var v9 uint64 = 1
        _ = v9
    }
    switch v4 {
    case "a":
        var v10 uint32 = 354
        _ = v10
    }
    v11 := v3[-2147483648]
    var v12 int64 = -598
    v13 := []int{32768, -9223372036854775808, -9223372036854775808, 0}
    v14 := GenS2{F0: 934, F1: (^(v11))}
    var v15 float64 = 3.14159
    v16 := GenS2{F0: 0, F1: uint32(v2)}
    switch v4 {
    case "hello":
        v17 := GenS1{F0: (v12 + 553), F1: genFn0(v1, 1, 310), F2: v4}
        _ = v17
    case "":
        _ = len(v3)
    case "\n\t":
        fmt.Println(v4, v2, v11)
    }
    var v18 float32 = v7
    v19 := GenS2{F0: 2, F1: 4294967295}
    fmt.Println(v1, v2)
    if false {
        _ = len(v13)
        if false {
            switch 1 {
            case 595:
                v20 := GenS1{F0: (-(v12)), F1: genFn0(v1, 622, 4294967295), F2: "foo bar"}
                _ = v20
            case 2147483647:
                v21 := v3[-61]
                _ = v21
            case -454:
                v22 := GenS3{F0: v1, F1: -1, F2: -1}
                _ = v22
            default:
                var v23 uint32 = (v11 << (((v2 - 42)) & 63))
                _ = v23
            }
            var v24 float64 = 1e10
            _ = v24
        } else {
            if true {
                v4 = ("a" + (v4 + "a"))
            } else {
                var v25 float64 = -2.25
                var v26 int32 = -658
                _ = v25
                _ = v26
            }
        }
    }
    v27 := GenS1{F0: (v12 >> (((v2 - v2)) & 63)), F1: 470, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
