package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 uint32) int16 {
    var v6 float64 = float64(1e10)
    if (1 != -112) {
        fmt.Println(p0)
    }
    v7 := []int{32768, 42, -9223372036854775808, 9223372036854775807}
    switch "" {
    case "":
        var v8 int16 = 161
        _ = v8
    case "\n\t":
        fmt.Println(p0, v6)
    }
    _ = p0
    _ = v6
    _ = v7
    return -143
}
func genFn1(p0 uint64, p1 int8) uint8 {
    v9 := []int{-344, -834, 9223372036854775807, 42}
    p1 = 42
    var v10 int16 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return uint8(p0)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    if false {
        v11 := v4[478]
        switch v11 {
        case 403:
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: -488, F1: 32768}
                var v14 uint16 = 42
                _ = fi12
                _ = v13
                _ = v14
            }
        case 32768:
            switch 9223372036854775807 {
            case 47:
                v3 = v3
            case 42:
                fmt.Println(v3, v11)
            case -1:
                fmt.Println(v11, v11)
            default:
                v15 := GenS2{F0: 255}
                _ = v15
            }
        }
        v3 = -2.25
        _ = v11
    } else {
        v3 -= v3
        var v16 bool = true
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float32 = (0.1 * -2.25)
        v19 := GenS3{F0: 237}
        switch (fi17 ^ 42) {
        case -600:
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS2{F0: 42}
                v18 = (1e10 - v18)
                switch "\n\t" {
                case "\n\t":
                    v18 += float32((0.0 * v3))
                case "a":
                    v18 = (v18 * (v18 / (1.5 / v18)))
                case "foo bar":
                    switch ("" + "hello") {
                    case "\n\t":
                        v22 := GenS3{F0: 821}
                        _ = v22
                    }
                }
                _ = fi20
                _ = v21
            }
        case 1:
            v18 = v18
        case 4294967295:
            v2["hello"] = false
        default:
            v18 = (v18 - v18)
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v23 int16 = -1
    for fi24 := 0; fi24 < 6; fi24++ {
        v3 = -2.25
        var v25 uint = 2
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: 33, F1: 2147483647, F2: 100}
    var v27 int8 = -128
    v3 += v3
    var v28 float32 = (((((float32(0.1) * 1.5) * 1.5) / float32(((v3 / (v3 - float64(3.14159))) - v3))) + (-2.25 / 0.1)) - -2.25)
    v29 := GenS1{F0: 0, F1: 702}
    v30 := GenS3{F0: 513}
    v2["foo bar"] = (18446744073709551615 >= 652)
    delete(v1, ("foo bar" + ("a" + ("" + (("" + (((((("hello" + (("a" + "foo bar") + "a")) + ("" + ("a" + "a"))) + "a") + "foo bar") + "\n\t") + "\n\t")) + ("" + "a"))))))
    v31 := GenS1{F0: -1, F1: 1}
    v23 = (-(v23))
    var v32 float64 = v3
    for fi33 := 0; fi33 < 4; fi33++ {
        v3 -= 1.5
        var v34 uint64 = 18446744073709551615
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
