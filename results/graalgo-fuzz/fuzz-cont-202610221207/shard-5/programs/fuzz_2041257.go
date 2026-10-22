package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 int8 }
func genFn0(p0 string) string {
    v7 := []int{32767, 2147483647, -1}
    if true {
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return ""
}

func main() {
    v1 := []uint32{625, 771, 207}
    var v2 int8 = 0
    var v3 int64 = -1
    v4 := make(map[int]bool)
    var v5 int8 = (v2 ^ v2)
    var v6 int8 = (v2 + 127)
    v3 = (v3 & v3)
    fmt.Println(v3, v6)
    var v8 int = 0
    v9 := GenS1{F0: 2147483647}
    for fi10 := 0; fi10 < 15; fi10++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = fi10
    }
    v11 := GenS2{F0: 1000}
    var v12 int8 = (v5 << ((239) & 63))
    switch "hello" {
    case "a":
        v3 -= (v3 | (v3 >> ((1) & 63)))
    default:
        v13 := GenS1{F0: 48}
        _ = v13
    }
    switch "a" {
    case "a":
        var v14 int32 = 599
        _ = v14
    case "hello":
        var v15 bool = false
        _ = v15
    case "foo bar":
        switch (v8 | -760) {
        case -2147483648:
            if false {
                var v16 int8 = v12
                v17 := GenS1{F0: 2147483647}
                v4[v8] = false
                _ = v16
                _ = v17
            }
        default:
            var v18 int = v8
            _ = v18
        }
    }
    var v19 int16 = 1000
    if true {
        switch -225 {
        case -9223372036854775808:
            var v20 uint16 = 32767
            _ = v20
        case -1:
            var v21 int = (v8 / ((v8) | 1))
            _ = v21
        case 128:
            delete(v4, (v8 << ((18446744073709551615) & 63)))
        }
        v3 = v3
    }
    if ((false && false) || ((false && false) && (!true))) {
        for fi22 := 0; fi22 < 12; fi22++ {
            v8 += fi22
            v23 := GenS3{F0: v3, F1: v5, F2: 255}
            _ = fi22
            _ = v23
        }
    }
    var v24 uint32 = 1
    v25 := GenS3{F0: (v3 + v3), F1: v12, F2: 2}
    v26 := GenS4{F0: (v5 | int8(v24))}
    var v27 uint = 65535
    v8 = (v8 & v8)
    var v28 uint64 = 368
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
