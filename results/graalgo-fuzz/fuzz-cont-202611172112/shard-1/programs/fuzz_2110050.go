package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 int64, p2 int8) uint8 {
    var v6 int = 2147483647
    v7 := []int{}
    var v8 int = (v6 % ((9223372036854775807) | 1))
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 uint8 = 2
            if (-801 >= 42) {
                v8 = 9223372036854775807
                v7 = append(v7, fi9)
            } else {
                p1 -= p1
                v7 = append(v7, v8)
                v6 = v6
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        _ = len(v7)
        if false {
            p2 += int8(p1)
            if (fi12 < v8) {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 int8 = -95
                    p2 += (v14 << ((758) & 63))
                    v14 = (p2 % ((v14) | 1))
                    _ = fi13
                    _ = v14
                }
            }
            v15 := GenS1{F0: 0, F1: v8, F2: p2}
            _ = v15
        } else {
            switch v8 {
            case 128:
                v16 := GenS2{F0: (p0 & (p0 + 128)), F1: 168, F2: -2147483648}
                _ = v16
            case 32767:
                var v17 int = (fi12 * 4294967295)
                _ = v17
            default:
                var v18 int64 = p1
                _ = v18
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 float32 = float32((1.5 - 1e10))
    var v2 float32 = 1.5
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 int32 = -1
    switch -2147483648 {
    case 192:
        var v19 uint16 = 65535
        _ = v19
    }
    fmt.Println(v2, v1)
    v3 = (278 <= 18446744073709551615)
    v3 = (v3 && true)
    v20 := []int{4294967295}
    delete(v4, "foo bar")
    v1 = -2.25
    switch "hello" {
    case "hello":
        var v21 int32 = v5
        _ = v21
    case "foo bar":
        var v22 float32 = float32((0.1 - 1.5))
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        _ = len(v4)
        _ = fi23
    }
    var v24 int = -337
    if (42 != 0) {
        v24 = v24
        if true {
            if true {
                var v25 int32 = v5
                v24 = v24
                _ = v25
            }
            var v26 uint16 = 65535
            _ = v26
        }
    }
    var v27 uint64 = 120
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
