package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint32, p1 int64) int64 {
    p0 = 423
    if false {
        v4 := []int{-2147483648, -646, 329, 42}
        switch "" {
        case "hello":
            fmt.Println(p1, p0)
        }
        _ = v4
    } else {
        var v5 uint = 18446744073709551615
        _ = v5
    }
    if false {
        switch -2147483648 {
        case 42:
            p1 = -9223372036854775808
        case 1:
            p0 = (p0 % ((p0) | 1))
        default:
            v6 := []int{}
            _ = v6
        }
        v7 := []int{127}
        _ = v7
    }
    v8 := GenS1{F0: -1, F1: (!false)}
    v9 := GenS2{F0: -2147483648, F1: p1}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p1
}
func genFn1(p0 int32, p1 uint16, p2 uint) uint8 {
    v10 := []int{9223372036854775807, -982, 770}
    v11 := GenS1{F0: -9223372036854775808, F1: (63 >= -9223372036854775808)}
    fmt.Println(p1)
    var v12 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 32767
    var v3 int64 = -9223372036854775808
    v2 += -1
    if (false && (43 != 2)) {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := v1[2]
            v2 = (v2 % ((v2) | 1))
            _ = fi13
            _ = v14
        }
        v15 := GenS3{F0: v3, F1: v2}
        _ = v15
    } else {
        v3 = (v3 << ((1) & 63))
        if (true && (42 > 202)) {
            var v16 float64 = 1.5
            switch 72 {
            case -949:
                v17 := []int{42, 32768, 2147483647}
                _ = v17
            case -523:
                var v18 uint16 = 0
                _ = v18
            }
            _ = v16
        } else {
            v3 -= (-(v3))
        }
    }
    if (("\n\t" + (("\n\t" + "\n\t") + "hello")) == "foo bar") {
        v3 = v3
        v19 := GenS4{F0: v3}
        _ = v19
    }
    if ((-2147483648 <= -1) || false) {
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS5{F0: ("\n\t" + "a"), F1: 548, F2: 2147483648}
            var v22 int16 = (v2 | v2)
            v23 := GenS2{F0: 42, F1: genFn0(2147483647, (^(v3)))}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    } else {
        var v24 float64 = -2.25
        switch ("a" + ((("a" + "a") + (("hello" + "foo bar") + "\n\t")) + "hello")) {
        case "hello":
            var v25 float64 = ((0.1 - 1e10) / (v24 * (v24 - -2.25)))
            _ = v25
        default:
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := GenS4{F0: (v3 >> ((42) & 63))}
                if ((v2 <= 32767) && true) {
                    fmt.Println(v24)
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        var v28 int16 = v2
                        _ = v28
                    case "a":
                        for fi29 := 0; fi29 < 16; fi29++ {
                            switch "foo bar" {
                            case "":
                                v30 := v1[2]
                                _ = v30
                            }
                            v3 += (-(v3))
                            var v31 int = (fi29 / ((fi26) | 1))
                            _ = fi29
                            _ = v31
                        }
                    default:
                        _ = len(v1)
                    }
                }
                v24 -= v24
                _ = fi26
                _ = v27
            }
        }
        _ = v24
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        if true {
            v33 := GenS1{F0: (v3 - (v3 >> ((763) & 63))), F1: false}
            _ = v33
        }
        var v34 uint64 = 42
        _ = fi32
        _ = v34
    }
    v3 -= v3
    var v35 uint64 = 42
    if false {
        v3 = v3
        switch 127 {
        case -9223372036854775808:
            var v36 uint32 = 2147483647
            _ = v36
        case 1:
            v37 := []int{-2147483648, 32768, 500}
            _ = v37
        case 2147483647:
            v38 := GenS6{F0: -9223372036854775808}
            _ = v38
        default:
            v2 -= -1
        }
        fmt.Println(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
