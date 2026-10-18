package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int32, p1 uint16) uint {
    if true {
        fmt.Println(p0)
        v6 := []int{971, 4294967295}
        _ = v6
    } else {
        fmt.Println(p0, p0)
        switch ("a" + "foo bar") {
        case "hello":
            p0 -= -2147483648
        case "\n\t":
            var v7 uint16 = (p1 + p1)
            _ = v7
        case "":
            switch -259 {
            case -682:
                v8 := []int{0, -139, 32767}
                _ = v8
            case -2147483648:
                var v9 int64 = 827
                _ = v9
            default:
                p0 -= p0
            }
        default:
            var v10 bool = (18446744073709551615 <= 255)
            _ = v10
        }
    }
    v11 := GenS1{F0: 1000, F1: 2}
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2, 174}
    var v2 bool = (!(!(154 < 35)))
    var v3 float32 = (-2.25 + -2.25)
    var v4 int32 = 685
    v5 := make(map[string]int)
    v5[""] = 32767
    var v12 int16 = 1000
    v13 := GenS1{F0: 32767, F1: 0}
    switch -329 {
    case -9223372036854775808:
        v14 := GenS2{F0: 2147483647, F1: 24}
        _ = v14
    default:
        var v15 int64 = 9223372036854775807
        _ = v15
    }
    v16 := GenS2{F0: 4294967295, F1: 840}
    v1 = append(v1, 2)
    v4 += (v4 & v4)
    v2 = true
    v3 += 0.1
    if ("foo bar" != "hello") {
        switch -773 {
        case 42:
            var v17 int32 = v4
            _ = v17
        case 128:
            v4 = (v4 * v4)
        case 1:
            v4 += 0
        default:
            v1 = append(v1, 2147483648)
        }
        var v18 uint64 = 2
        fmt.Println(v12, v12, v4)
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
