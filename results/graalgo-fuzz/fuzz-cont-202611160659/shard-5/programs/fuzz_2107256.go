package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 string) uint {
    p1 = "foo bar"
    v4 := GenS1{F0: 65535, F1: p1}
    fmt.Println(p1, p0)
    v5 := []int{32768, 32768, 696}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 uint8) bool {
    v6 := []int{42, 521, 0, 9223372036854775807}
    switch 0 {
    case 42:
        if len(v6) > 0 {
            v6[0] = -2147483648
        }
    case 9223372036854775807:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case -1:
        var v7 int32 = -2147483648
        _ = v7
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v8 int32 = 592
    p1 += 128
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return (3.14159 <= (3.14159 * 1e10))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    var v9 uint16 = 609
    v10 := GenS2{F0: 42, F1: 1}
    v11 := GenS2{F0: -2147483648, F1: 224}
    v12 := GenS3{F0: 1}
    v13 := []int{-22, -201}
    fmt.Println(v9)
    switch "" {
    case "\n\t":
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    _ = len(v2)
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 16; fi15++ {
        v9 = (^(v9))
        v9 += 829
        if len(v13) > 0 {
            v13[0] = (^(fi15))
        }
        _ = fi15
    }
    var v16 string = "a"
    var v17 int32 = 1
    delete(v1, 127)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
