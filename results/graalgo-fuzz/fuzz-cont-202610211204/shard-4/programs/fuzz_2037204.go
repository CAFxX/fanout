package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0(p0 uint64) int8 {
    p0 = p0
    v4 := []int{781, 1}
    if (true && (false && false)) {
        p0 -= 1
    } else {
        v5 := GenS1{F0: 2, F1: 131}
        v6 := GenS2{F0: 255, F1: 9223372036854775807, F2: 0}
        _ = v5
        _ = v6
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 int8, p2 uint) uint32 {
    v7 := []int{32767, 872, 9223372036854775807}
    fmt.Println(p2, p2, p1)
    switch -975 {
    case 724:
        var v8 float64 = 1.5
        _ = v8
    }
    if len(v7) > 0 {
        v7[0] = 685
    }
    p1 += (p1 - (p1 << (((p0 & p0)) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    var v3 uint16 = 1
    fmt.Println(v3)
    var v9 uint64 = 1
    v2 = v2
    var v10 uint64 = v9
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int16 = 0
        if true {
            v9 += 573
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v9)
    fmt.Println(v10)
    v13 := []int{128}
    v10 -= v9
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int16 = 115
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            var v17 int64 = 0
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            fmt.Println(fi18, fi18, v3)
            if true {
                v19 := GenS2{F0: 127, F1: 764, F2: 375}
                var v20 int32 = 0
                _ = v19
                _ = v20
            } else {
                v2 = v2
            }
            v2 = ("hello" + (("hello" + "foo bar") + (v2 + v2)))
            _ = fi18
        }
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println(v10, v2, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
