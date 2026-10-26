package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint32) uint64 {
    fmt.Println(p2)
    if false {
        var v4 float32 = 0.1
        p2 = (p0 + p2)
        _ = v4
    } else {
        p1 += (p1 * p1)
        v5 := []int{1, 0, -639, 32768}
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int = 487
            p2 = p2
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 = 1
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    v3 := []int{42, -996, 53, -687}
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v2, fi9)
        if (42 != 1) {
            switch (fi9 - (fi9 / (((fi9 % ((9223372036854775807) | 1))) | 1))) {
            case -2147483648:
                v10 := GenS1{F0: 592, F1: 40}
                _ = v10
            }
            v2 = (v2 << ((65535) & 63))
            v3 = append(v3, -1)
        }
        v11 := GenS2{F0: 32767}
        _ = fi9
        _ = v11
    }
    v2 = genFn0(1, 42, 65535)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 float64 = (((0.1 / 1.5) / (0.1 * 1e10)) * 1e10)
        _ = fi12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    var v14 bool = true
    var v15 uint8 = 2
    fmt.Println(v2, v2, v2)
    var v16 bool = (1000 > -1)
    switch "foo bar" {
    case "hello":
        _ = len(v3)
    case "":
        v2 -= (v2 % (((v2 / ((v2) | 1))) | 1))
    case "foo bar":
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 int32 = 1
            fmt.Println(v14, v14, fi17)
            switch (fi17 % ((fi17) | 1)) {
            case 488:
                v14 = v14
            case -9223372036854775808:
                switch (fi17 - (fi17 - 4294967295)) {
                case 1:
                    var v19 uint = uint(v2)
                    _ = v19
                default:
                    v20 := GenS3{F0: 100, F1: false}
                    _ = v20
                }
            case 32768:
                for fi21 := 0; fi21 < 12; fi21++ {
                    fmt.Println(v18, v18, v15)
                    _ = fi21
                }
            }
            _ = fi17
            _ = v18
        }
    }
    v14 = true
    v14 = v16
    var v22 int64 = -1
    if false {
        var v23 uint32 = 2
        fmt.Println(v16)
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
