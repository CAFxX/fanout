package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 int8 }
func genFn0() bool {
    if false {
        switch "hello" {
        case "hello":
            var v7 uint64 = 42
            _ = v7
        case "foo bar":
            var v8 uint8 = 194
            _ = v8
        }
        var v9 uint32 = 65535
        _ = v9
    } else {
        var v10 uint8 = 128
        _ = v10
    }
    v11 := []int{32767, 2147483647, 9223372036854775807, 0}
    _ = v11
    return true
}
func genFn1(p0 uint16, p1 uint, p2 bool) bool {
    p2 = (-222 >= 0)
    if false {
        if p2 {
            v12 := []int{1, -27, 4294967295}
            p2 = p2
            _ = v12
        }
    }
    v13 := []int{417, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}

func main() {
    var v1 int = -695
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    var v4 uint16 = 1
    var v5 uint16 = (^(v4))
    var v6 float32 = 0.1
    var v14 string = ("a" + "")
    v15 := GenS1{F0: 9223372036854775807, F1: v2}
    var v16 int64 = -1
    v4 = v5
    fmt.Println(v16)
    var v17 uint16 = 255
    v18 := GenS1{F0: (v16 / ((v16) | 1)), F1: 0}
    var v19 uint32 = 2147483647
    var v20 uint16 = (v5 - v4)
    var v21 uint32 = v19
    var v22 float32 = v6
    fmt.Println(v1, v4, v16)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v1 += 1
            if false {
                v4 = 239
                var v25 int16 = -967
                v6 = v22
                _ = v25
            }
            v26 := GenS2{F0: false}
            v27 := GenS3{F0: 2, F1: v21, F2: (v17 % ((v20) | 1))}
            _ = fi24
            _ = v26
            _ = v27
        }
        v28 := GenS4{F0: 65535, F1: fi23, F2: -128}
        _ = fi23
        _ = v28
    }
    var v29 uint32 = (v19 >> ((774) & 63))
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
