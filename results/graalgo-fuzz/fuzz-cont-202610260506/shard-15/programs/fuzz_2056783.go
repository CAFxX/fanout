package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 bool }
func genFn0() uint32 {
    var v7 float64 = -2.25
    v8 := []int{9223372036854775807, 9223372036854775807, -2147483648}
    v7 -= v7
    v9 := GenS1{F0: ("foo bar" + ""), F1: 4294967295}
    if ((726 > 32767) && (false || (false || (255 != 1)))) {
        if (-69 > 100) {
            var v10 int8 = -1
            v7 = 3.14159
            _ = v10
        } else {
            fmt.Println(v7)
        }
        if len(v8) > 0 {
            v8[0] = -1
        }
        v11 := GenS2{F0: 18446744073709551615}
        _ = v11
    } else {
        v7 = (0.0 * (1e10 - (v7 + v7)))
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS3{F0: 42, F1: 109}
            _ = fi12
            _ = v13
        }
        v14 := GenS3{F0: 127, F1: 1}
        _ = v14
    }
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 int32, p1 int) uint32 {
    v15 := GenS2{F0: 137}
    p0 = p0
    var v16 uint = 42
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := []int{(p1 ^ fi18)}
            p1 = -9223372036854775808
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 int8 = 100
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v20
    return 0
}

func main() {
    var v1 int32 = -1
    v2 := []int{4294967295, 32768, 9223372036854775807, -2147483648}
    v3 := []int{-380, -9223372036854775808, 2147483647, 127}
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    v21 := GenS3{F0: 147, F1: 116}
    v4 = (v4 + v4)
    for fi22 := 0; fi22 < 9; fi22++ {
        if true {
            switch fi22 {
            case -9223372036854775808:
                switch (fi22 + (fi22 / ((fi22) | 1))) {
                case 32767:
                    v2 = append(v2, fi22)
                }
            }
        } else {
            v1 -= (v1 - v1)
        }
        v23 := v6[32767]
        var v24 uint = 4294967295
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v1 = -2147483648
        v4 = (v4 << ((4294967295) & 63))
        v26 := v6[887]
        _ = fi25
        _ = v26
    }
    switch 42 {
    case 128:
        delete(v6, -1)
    case -1:
        v1 += -2147483648
    case 4294967295:
        switch 32767 {
        case -346:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 42:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 805:
            fmt.Println(v1, v1)
        default:
            fmt.Println(v1)
        }
    default:
        v1 -= v1
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS4{F0: false}
        v6[(fi27 << ((0) & 63))] = fi27
        v29 := GenS2{F0: 776}
        _ = fi27
        _ = v28
        _ = v29
    }
    switch "" {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        switch -305 {
        case 42:
            var v30 string = ("" + "foo bar")
            _ = v30
        default:
            var v31 int32 = (v1 << ((0) & 63))
            _ = v31
        }
    case "foo bar":
        _ = len(v2)
    default:
        for fi32 := 0; fi32 < 20; fi32++ {
            _ = len(v3)
            v33 := GenS1{F0: "a", F1: 42}
            var v34 uint8 = 2
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v1 = -590
        var v36 uint64 = v4
        v6[fi35] = (fi35 % (((^(fi35))) | 1))
        _ = fi35
        _ = v36
    }
    v6[128] = -2147483648
    v1 -= (v1 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
