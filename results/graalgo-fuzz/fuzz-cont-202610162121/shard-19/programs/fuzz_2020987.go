package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint32, p1 uint64, p2 uint16) uint16 {
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 -= 65535
        _ = fi5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = p0
        p0 = (p0 - (p0 / (((p0 % ((p0) | 1))) | 1)))
        switch "hello" {
        case "a":
            fmt.Println(p0, fi6)
        case "hello":
            v7 := []int{(fi6 + (^(fi6))), 42}
            _ = v7
        case "\n\t":
            switch "hello" {
            case "hello":
                fmt.Println(p1, fi6)
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 163
    var v3 int64 = 598
    var v4 int32 = 2147483647
    v8 := GenS1{F0: (v2 % ((v2) | 1))}
    switch "a" {
    case "\n\t":
        v4 = (^(v4))
    case "a":
        v4 = (-(v4))
    case "foo bar":
        v9 := GenS2{F0: 188, F1: -655, F2: false}
        _ = v9
    }
    v2 = genFn0(65535, (v1 | 4294967295), (v2 << ((565) & 63)))
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v3, v1, v3)
        _ = fi10
    }
    v11 := GenS2{F0: 966, F1: 2147483647, F2: true}
    switch 4294967295 {
    case 936:
        fmt.Println(v2, v3)
    case 128:
        v12 := GenS1{F0: 255}
        _ = v12
    case 32768:
        v2 = (v2 << ((42) & 63))
    default:
        v13 := []int{}
        _ = v13
    }
    v14 := []int{42, 42}
    v4 += (v4 << ((132) & 63))
    v4 += v4
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int16 = 1
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v4, v2)
    var v17 uint64 = uint64(v2)
    v3 += (v3 - v3)
    v18 := GenS3{F0: int16(v2), F1: 119}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
