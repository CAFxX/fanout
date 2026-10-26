package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
func genFn0() int32 {
    if ("foo bar" > "a") {
        var v6 int = -2147483648
        var v7 uint8 = 255
        _ = v6
        _ = v7
    }
    var v8 int16 = -32768
    _ = v8
    return -1
}
func genFn1(p0 int, p1 int8, p2 bool) uint64 {
    if true {
        p1 += p1
        v9 := GenS1{F0: p2, F1: 976, F2: genFn0()}
        p1 -= (-(p1))
        _ = v9
    }
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 814
    var v3 int = -1
    var v4 int16 = 1
    v5 := make(map[int]bool)
    v2 = (v4 / ((-1) | 1))
    var v10 uint = 1
    v3 -= v3
    switch int(v1) {
    case 127:
        fmt.Println(v1, v1)
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v2 = 1
                var v13 uint = 4294967295
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    v1 += (v1 << (((v10 | v10)) & 63))
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int16 = v2
        v4 = v2
        _ = fi14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 7; fi16++ {
            v10 += (^(v10))
            v17 := GenS2{F0: (-(fi16))}
            _ = fi16
            _ = v17
        }
    }
    var v18 int64 = 42
    v19 := []int{}
    for fi20 := 0; fi20 < 12; fi20++ {
        v19 = append(v19, (-(fi20)))
        v21 := GenS3{F0: int64(v10)}
        _ = fi20
        _ = v21
    }
    if true {
        fmt.Println(v10)
        var v22 float32 = 3.14159
        _ = v22
    } else {
        if len(v19) > 0 {
            v19[0] = 255
        }
        v2 = (^(v2))
    }
    if (true || false) {
        v18 = v18
        var v23 int64 = -1
        if ((654 < 65535) || ((!false) || (-55 < 52))) {
            v18 = (v18 >> (((v10 * 4294967295)) & 63))
        }
        _ = v23
    }
    if ((^(v4)) >= v4) {
        _ = len(v5)
    } else {
        switch v3 {
        case 128:
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v10)
                _ = fi24
            }
        default:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
        v3 += (v3 - v3)
        fmt.Println(v1, v3, v3)
    }
    v18 -= v18
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
