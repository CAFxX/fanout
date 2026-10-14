package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 uint8, p1 int8) int64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            p1 = 100
            v7 := []int{1}
            var v8 float32 = ((-2.25 / (((0.0 / (0.1 - 1e10)) * -2.25) - 1.5)) - ((float32(-2.25) - -2.25) + -2.25))
            _ = fi6
            _ = v7
            _ = v8
        }
        if ((float32(-2.25) / (3.14159 / 1e10)) != 1.5) {
            if false {
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := GenS1{F0: (fi9 - 2147483647), F1: 745, F2: 255}
                    var v11 int16 = -229
                    p0 = (^(p0))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            var v12 int = (-(fi5))
            switch "" {
            case "foo bar":
                var v13 float64 = float64(((3.14159 * -2.25) * 0.1))
                _ = v13
            case "hello":
                v12 -= (fi5 & (^(v12)))
            }
            _ = v12
        } else {
            var v14 uint = 936
            if (1.5 <= 0.0) {
                if false {
                    var v15 uint64 = 18446744073709551615
                    var v16 uint = (v14 ^ (v14 / ((v14) | 1)))
                    _ = v15
                    _ = v16
                }
                p0 += p0
            }
            _ = v14
        }
        if (!true) {
            switch "foo bar" {
            case "hello":
                v17 := GenS2{F0: 661}
                _ = v17
            case "a":
                v18 := []int{fi5, fi5, int(p1)}
                _ = v18
            case "foo bar":
                var v19 int8 = -110
                _ = v19
            default:
                v20 := GenS2{F0: 0}
                _ = v20
            }
        }
        _ = fi5
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint8 = (p0 << ((308) & 63))
        v23 := []int{-1, 936, (fi21 ^ (fi21 ^ (fi21 - (fi21 / (((fi21 % ((fi21) | 1))) | 1)))))}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: 1}
    if (true || true) {
        p0 += 1
        v25 := []int{1, -413, -115}
        _ = v25
    } else {
        fmt.Println(p1, p1)
        switch -2147483648 {
        case 32768:
            if true {
                v26 := []int{-873, 1}
                var v27 bool = true
                _ = v26
                _ = v27
            }
        }
        fmt.Println(p0)
    }
    v28 := GenS2{F0: 441}
    _ = p0
    _ = p1
    _ = v24
    _ = v28
    return 228
}
func genFn1(p0 int32, p1 int8) uint64 {
    fmt.Println(p0, p0)
    v29 := GenS4{F0: -286, F1: ""}
    _ = p0
    _ = p1
    _ = v29
    return 2
}

func main() {
    var v1 int16 = 278
    var v2 uint = 0
    var v3 int8 = 115
    var v4 uint8 = uint8(v2)
    switch 4294967295 {
    case 675:
        fmt.Println(v4, v3, v1)
    case 127:
        switch -1 {
        case -9223372036854775808:
            v1 = 1
        }
    case 870:
        v30 := []int{32767, 0, -230}
        _ = v30
    default:
        v2 -= (^(v2))
    }
    var v31 uint32 = 763
    v31 = v31
    if (false && true) {
        v2 = (^(v2))
    } else {
        fmt.Println(v4, v3, v2)
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 int8 = (v3 | v3)
        v2 -= v2
        _ = fi32
        _ = v33
    }
    v31 += v31
    var v34 uint = (v2 - (v2 & v2))
    if true {
        var v35 float64 = 1e10
        _ = v35
    }
    v3 += (v3 % ((v3) | 1))
    v31 -= uint32(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
