package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 int32) int16 {
    var v6 uint = 389
    var v7 int32 = (p1 ^ 0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int16, p1 string, p2 int8) int64 {
    switch 127 {
    case -858:
        v8 := []int{32768, -1}
        _ = v8
    case 9223372036854775807:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint16 = 42
            switch fi9 {
            case 32767:
                p0 += (^(p0))
            case 561:
                fmt.Println(fi9, p2)
            case 1:
                var v11 uint16 = v10
                _ = v11
            }
            var v12 uint8 = 2
            _ = fi9
            _ = v10
            _ = v12
        }
    default:
        fmt.Println(p2, p1)
    }
    var v13 int16 = (p0 & (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -136
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    if true {
        var v14 uint16 = 2
        _ = v14
    }
    v15 := GenS1{F0: false, F1: 42}
    if true {
        var v16 uint64 = 818
        v17 := GenS2{F0: 42}
        switch -2147483648 {
        case -780:
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 int64 = -893
                _ = fi18
                _ = v19
            }
        default:
            v16 += 18446744073709551615
        }
        _ = v16
        _ = v17
    }
    var v20 uint64 = 42
    fmt.Println(v20)
    if false {
        if false {
            var v21 uint = 1
            v22 := []int{0, -571, -2147483648, 9223372036854775807}
            var v23 uint8 = 255
            _ = v21
            _ = v22
            _ = v23
        }
    } else {
        fmt.Println(v20)
        var v24 uint64 = v20
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 int16 = 0
            var v27 int64 = -948
            if (1 != v26) {
                fmt.Println(fi25, v27, v27)
            } else {
                v28 := GenS2{F0: 69}
                _ = v28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    }
    v20 = (v20 * v20)
    if false {
        v20 += v20
        fmt.Println(v20)
        var v29 int16 = -1
        _ = v29
    }
    v30 := GenS2{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
