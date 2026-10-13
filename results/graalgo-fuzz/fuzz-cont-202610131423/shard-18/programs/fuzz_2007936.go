package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 uint64) int16 {
    if (false || (18446744073709551615 == 731)) {
        v7 := []int{(p0 << ((4294967295) & 63)), p0, (p0 + p0)}
        _ = v7
    } else {
        var v8 uint64 = (p1 >> ((18446744073709551615) & 63))
        switch "\n\t" {
        case "foo bar":
            p0 -= (p0 << ((255) & 63))
        }
        var v9 uint64 = (^(p1))
        _ = v8
        _ = v9
    }
    var v10 float64 = (0.1 / float64(0.1))
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -763
}
func genFn1(p0 int16, p1 int64) bool {
    if (0 != p0) {
        p1 += 506
        for fi12 := 0; fi12 < 10; fi12++ {
            p1 = 0
            _ = fi12
        }
    }
    if ((false && (938 >= 42)) || true) {
        var v13 float32 = float32((0.1 - (3.14159 - ((float64(-2.25) + 3.14159) * (((0.0 / 0.1) + (1e10 / -2.25)) - (1.5 - float64(3.14159)))))))
        var v14 int8 = -1
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{-758}
    var v3 string = ("a" + "foo bar")
    v4 := make(map[int]uint32)
    var v5 uint = 2
    var v6 int8 = 64
    var v15 float32 = (v1 * -2.25)
    switch -1 {
    case -2147483648:
        var v16 int32 = 2147483647
        _ = v16
    case 32768:
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                v6 = 127
                if true {
                    var v19 int = (fi17 | fi17)
                    v20 := GenS1{F0: -659, F1: 0, F2: 65535}
                    _ = v19
                    _ = v20
                }
                if false {
                    var v21 uint64 = 642
                    _ = v21
                }
                _ = fi18
            }
            _ = fi17
        }
    case -9223372036854775808:
        v5 = (v5 | v5)
    default:
        var v22 float32 = 0.0
        _ = v22
    }
    v3 = "foo bar"
    v23 := GenS1{F0: 1, F1: 32767, F2: 4294967295}
    v2 = append(v2, -475)
    fmt.Println(v6)
    v6 = (-(v6))
    if (-1 == 951) {
        v24 := GenS2{F0: 0, F1: 42}
        v25 := v4[-9223372036854775808]
        v25 += (v25 / ((v25) | 1))
        _ = v24
        _ = v25
    }
    v1 -= v1
    var v26 uint8 = 1
    var v27 uint8 = v26
    delete(v4, 32768)
    switch v3 {
    case "foo bar":
        v4[-2147483648] = 2147483647
    case "hello":
        fmt.Println(v15)
    }
    fmt.Println(v3, v1)
    v28 := GenS2{F0: -32768, F1: uint16(v27)}
    v26 -= v27
    for fi29 := 0; fi29 < 10; fi29++ {
        v26 = 2
        switch ("\n\t" + v3) {
        case "":
            var v30 uint8 = 127
            _ = v30
        case "\n\t":
            var v31 bool = (255 < v26)
            _ = v31
        case "foo bar":
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 bool = true
                v6 = (-(v6))
                _ = fi32
                _ = v33
            }
        }
        var v34 uint16 = 852
        _ = fi29
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS1{F0: 9223372036854775807, F1: 460, F2: 2147483648}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
