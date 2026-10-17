package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]uint32)
    var v3 int16 = 0
    var v4 uint64 = 133
    var v5 int16 = v3
    v6 := make(map[uint32]bool)
    v3 = (v3 % (((-(v3))) | 1))
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v3, v1)
        _ = fi7
    }
    v8 := []int{9223372036854775807, 1, 1, -9223372036854775808}
    switch "foo bar" {
    case "foo bar":
        v9 := GenS1{F0: 2147483648, F1: 97, F2: ((127 == 5) && (697 <= 42))}
        _ = v9
    case "\n\t":
        switch 9223372036854775807 {
        case 4294967295:
            v1 = v1
        case 355:
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: 42, F1: 1, F2: (42 == 100)}
                _ = fi10
                _ = v11
            }
        case -371:
            var v12 uint16 = 65535
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint16 = 2
            v14 = 65535
            v15 := GenS2{F0: v14}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    switch "hello" {
    case "foo bar":
        for fi16 := 0; fi16 < 16; fi16++ {
            v4 -= 0
            v8 = append(v8, (fi16 | (-(fi16))))
            _ = fi16
        }
    default:
        if true {
            v17 := v2[128]
            _ = v17
        } else {
            v8 = append(v8, 916)
            var v18 uint16 = 65535
            _ = v18
        }
    }
    v4 += 347
    fmt.Println(v5, v1, v4)
    var v19 uint = 255
    v1 = v1
    if true {
        v3 += v5
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int32 = -345
            _ = fi20
            _ = v21
        }
        delete(v6, 617)
    }
    switch "\n\t" {
    case "":
        v19 -= v19
    case "a":
        v1 = 1
    case "\n\t":
        v5 = v3
    default:
        if len(v8) > 0 {
            v8[0] = 128
        }
    }
    if true {
        delete(v2, -357)
        fmt.Println(v19, v5)
        var v22 uint8 = v1
        _ = v22
    } else {
        v4 += (v4 + v4)
        v23 := GenS3{F0: uint64(v19), F1: 843}
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v4)
            var v25 uint32 = uint32(v19)
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
