package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
func genFn0(p0 int32, p1 bool, p2 uint64) int32 {
    var v4 uint32 = 2147483647
    if p1 {
        switch -139 {
        case 9223372036854775807:
            var v5 int16 = 1
            _ = v5
        case 128:
            fmt.Println(v4, p0)
        case 127:
            var v6 int = 2147483647
            _ = v6
        default:
            var v7 int16 = -969
            _ = v7
        }
    }
    switch 127 {
    case 4294967295:
        switch 42 {
        case 1:
            v8 := []int{9223372036854775807, 1, 4294967295, 0}
            _ = v8
        }
    default:
        switch -781 {
        case 32767:
            if true {
                var v9 string = "hello"
                _ = v9
            } else {
                p1 = ("foo bar" == "a")
                v10 := GenS1{F0: 42, F1: 367, F2: 255}
                var v11 bool = false
                _ = v10
                _ = v11
            }
        case 128:
            p2 -= (p2 % ((p2) | 1))
        }
    }
    var v12 uint32 = (v4 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return p0
}
func genFn1(p0 uint32, p1 uint8, p2 int) int {
    var v13 uint = 255
    fmt.Println(p0)
    switch "hello" {
    case "":
        v14 := []int{}
        _ = v14
    case "a":
        var v15 float32 = 1e10
        _ = v15
    case "foo bar":
        if true {
            var v16 uint = 2
            p1 += 11
            _ = v16
        } else {
            var v17 int16 = 1000
            var v18 int16 = int16(p1)
            switch ("a" + "foo bar") {
            case "a":
                v13 = v13
            default:
                p1 = (p1 >> ((0) & 63))
            }
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 0
    v3 := []string{}
    v19 := GenS1{F0: 286, F1: 42, F2: 0}
    v20 := GenS2{F0: 1, F1: 32767, F2: 42}
    v21 := v1[(^(v2))]
    var v22 uint32 = v2
    v21 = 4294967295
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS1{F0: 2, F1: -1, F2: 65535}
        v25 := v1[v21]
        _ = fi23
        _ = v24
        _ = v25
    }
    v3 = append(v3, "\n\t")
    var v26 string = ("foo bar" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
}
