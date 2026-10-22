package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 int16) int32 {
    if true {
        if true {
            var v4 int16 = p1
            _ = v4
        } else {
            v5 := []int{762, -180, 1}
            fmt.Println(p0, p0)
            _ = v5
        }
        p0 = -9223372036854775808
    }
    p1 = p1
    var v6 int16 = 1000
    if (-9223372036854775808 <= 185) {
        var v7 uint64 = 18446744073709551615
        var v8 int32 = -2147483648
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]bool)
    v3 := []int{}
    delete(v2, 1)
    v9 := GenS1{F0: 376}
    v3 = append(v3, 4294967295)
    _ = len(v2)
    if ((0.1 + 0.0) <= (((0.1 * 3.14159) + 3.14159) + (0.1 / 1.0))) {
        v1 = 0
    } else {
        v1 = v1
        v1 = 58
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "a":
        if false {
            v3 = append(v3, -917)
            v10 := GenS1{F0: 722}
            _ = v10
        }
    case "\n\t":
        v11 := GenS1{F0: 65535}
        _ = v11
    case "hello":
        var v12 int32 = 553
        _ = v12
    default:
        if false {
            var v13 int = -44
            switch v13 {
            case 9223372036854775807:
                fmt.Println(v1, v1)
            case 2147483647:
                var v14 int16 = -1
                _ = v14
            case 32767:
                v15 := GenS1{F0: 2}
                _ = v15
            default:
                _ = len(v3)
            }
            if len(v3) > 0 {
                v3[0] = (^(v13))
            }
            _ = v13
        }
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v1 = 100
    var v16 bool = true
    switch "foo bar" {
    case "hello":
        v17 := GenS1{F0: 42}
        _ = v17
    case "\n\t":
        v1 -= v1
    case "":
        _ = len(v2)
    }
    var v18 int8 = (-(v1))
    if (-664 > 128) {
        fmt.Println(v16, v18, v1)
        if false {
            var v19 int64 = 466
            v16 = (v16 || v16)
            _ = v19
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        fmt.Println(v18)
    }
    var v20 int64 = 902
    var v21 string = ("hello" + "foo bar")
    v22 := GenS2{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
