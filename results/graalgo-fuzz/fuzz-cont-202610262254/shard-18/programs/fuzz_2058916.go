package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 uint, p1 uint64, p2 int8) uint64 {
    var v4 float64 = (0.1 - (0.1 - 1e10))
    var v5 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint) uint8 {
    if true {
        p0 -= 4294967295
    } else {
        p0 = p0
        p0 += p0
        if (true || true) {
            var v6 int32 = 884
            if true {
                v6 -= (v6 * v6)
                v7 := GenS1{F0: false, F1: 42}
                _ = v7
            } else {
                p0 = (^(p0))
            }
            _ = v6
        }
    }
    p0 = 2
    v8 := GenS1{F0: (1e10 >= 3.14159), F1: -32768}
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 uint = 255
    var v2 int16 = 1
    v3 := make(map[string]bool)
    var v9 uint = v1
    switch 836 {
    case 32767:
        v10 := []int{-313, -496}
        _ = v10
    case -560:
        var v11 int = 9223372036854775807
        _ = v11
    case 32768:
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v9, v1)
            if (v2 >= int16(v1)) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint = v9
                    fmt.Println(v14, v2)
                    var v15 uint8 = 2
                    v16 := GenS1{F0: true, F1: v2}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                var v17 int8 = 0
                v18 := GenS2{F0: ("" + "a"), F1: 684, F2: (v1 + 227)}
                _ = v17
                _ = v18
            } else {
                if (!false) {
                    fmt.Println(v2, fi12)
                    switch "" {
                    case "\n\t":
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    }
                    var v20 string = "a"
                    _ = v20
                } else {
                    v21 := GenS2{F0: "a", F1: 0, F2: v1}
                    _ = v21
                }
            }
            delete(v3, ("foo bar" + "a"))
            _ = fi12
        }
    default:
        switch ("foo bar" + "") {
        case "":
            var v22 uint8 = 127
            _ = v22
        case "a":
            v2 = v2
        case "\n\t":
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v1 = v1
                    if false {
                        v2 = v2
                    }
                    v1 -= (v9 ^ v1)
                    _ = fi24
                }
                v2 = (-(v2))
                _ = fi23
            }
        default:
            var v25 uint8 = uint8(v1)
            _ = v25
        }
    }
    v2 += v2
    var v26 int32 = -208
    v2 = int16(v26)
    if false {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 uint8 = 128
            _ = fi27
            _ = v28
        }
        v9 -= (^(v1))
        for fi29 := 0; fi29 < 12; fi29++ {
            if false {
                v2 -= (-(v2))
                v3["foo bar"] = (-1 <= -501)
            }
            var v30 uint32 = 2
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v2, v9)
    for fi31 := 0; fi31 < 7; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            v26 = (v26 % ((int32(v2)) | 1))
            _ = fi32
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v9
}
