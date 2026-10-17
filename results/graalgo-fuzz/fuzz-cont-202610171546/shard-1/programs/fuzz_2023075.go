package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 int, p1 int32) int32 {
    v5 := []int{p0, (p0 + 867)}
    if (false || true) {
        if (863 > 255) {
            fmt.Println(p0, p1)
            v6 := GenS1{F0: "\n\t", F1: 1}
            v5 = append(v5, (p0 & 128))
            _ = v6
        } else {
            if true {
                v7 := GenS1{F0: "hello", F1: 255}
                v5 = append(v5, (p0 * p0))
                _ = v7
            }
        }
        var v8 string = ""
        v8 = v8
        _ = v8
    }
    p1 = int32(p0)
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9, fi9, p0)
        _ = fi9
    }
    p0 += 1
    _ = p0
    _ = p1
    _ = v5
    return (p1 ^ p1)
}

func main() {
    var v1 uint8 = 173
    var v2 uint64 = 4294967295
    var v3 string = "a"
    v4 := make(map[uint32]string)
    for fi10 := 0; fi10 < 20; fi10++ {
        if false {
            var v11 uint16 = 472
            v11 -= (v11 & 756)
            _ = v11
        }
        _ = fi10
    }
    v12 := []int{32768, -9223372036854775808, 292}
    v2 -= v2
    var v13 int = 1
    v13 -= 9223372036854775807
    switch 32767 {
    case 4294967295:
        _ = len(v12)
    case 127:
        v1 = (v1 % ((v1) | 1))
    default:
        fmt.Println(v2)
    }
    var v14 int = 29
    if false {
        v1 -= 2
        v13 += v14
    } else {
        v2 += v2
        var v15 int16 = 1
        _ = v15
    }
    v3 = v3
    v2 -= (v2 % ((v2) | 1))
    _ = len(v4)
    for fi16 := 0; fi16 < 8; fi16++ {
        v13 -= (fi16 + v14)
        var v17 uint64 = (v2 << ((761) & 63))
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
