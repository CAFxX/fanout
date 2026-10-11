package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint }
func genFn0(p0 int8, p1 int) int8 {
    p1 = p1
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 1}
        v9 := []int{(p1 | fi7), p1, 127, 9223372036854775807}
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 += p1
    _ = p0
    _ = p1
    return (p0 << ((2) & 63))
}

func main() {
    var v1 int16 = -1
    v2 := []int{32768, 467, -439}
    var v3 uint = 627
    v4 := make(map[string]bool)
    var v5 int16 = (v1 / ((-32768) | 1))
    v6 := make(map[uint32]bool)
    var v10 int8 = 127
    if (false || true) {
        if true {
            v11 := GenS2{F0: (v3 >> ((v3) & 63)), F1: v1, F2: 42}
            for fi12 := 0; fi12 < 8; fi12++ {
                v4["a"] = (255 < 189)
                _ = fi12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v5 = (v1 | (v1 >> ((v3) & 63)))
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 string = (("\n\t" + (("a" + "hello") + ((("foo bar" + (("foo bar" + "a") + "\n\t")) + "\n\t") + "foo bar"))) + "a")
                    v4[v15] = false
                    if (1.5 == ((((0.0 - (0.0 * (0.1 + 1e10))) - (3.14159 - 1e10)) * (0.1 * (((((((1.5 / float32(float64(3.14159))) - 0.1) * (((1e10 + 1e10) + 0.0) + ((0.1 / -2.25) * 0.0))) / float32(float64((0.0 + 0.1)))) - -2.25) + ((0.1 / -2.25) + 1.5)) - -2.25))) * (float32(-2.25) + 0.1))) {
                        var v16 float64 = (-2.25 - 0.1)
                        fmt.Println(v5, v10)
                        _ = v16
                    }
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            _ = v11
        }
    }
    var v17 int8 = (v10 * 100)
    var v18 int64 = -1
    v19 := GenS1{F0: 1}
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v1, v1, v18)
        v17 -= v10
        var v21 float64 = -2.25
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: 18446744073709551615, F1: v18}
    v23 := GenS4{F0: 106}
    if len(v2) > 0 {
        v2[0] = -403
    }
    if true {
        var v24 int8 = v17
        v10 += v24
        _ = v24
    }
    v1 -= (v1 - v1)
    if (!((v3 + (^(v3))) == v3)) {
        var v25 int64 = (v18 + (v18 ^ (v18 % (((v18 << ((4294967295) & 63))) | 1))))
        if false {
            switch "" {
            case "foo bar":
                v5 += (v5 - v5)
            }
            v26 := GenS5{F0: -111, F1: 0}
            _ = v26
        } else {
            fmt.Println(v18)
        }
        _ = v25
    } else {
        switch 1 {
        case 754:
            switch "hello" {
            case "hello":
                if false {
                    fmt.Println(v5, v17)
                }
            default:
                var v27 string = "foo bar"
                _ = v27
            }
        case 4294967295:
            var v28 bool = true
            _ = v28
        case 0:
            fmt.Println(v3, v17)
        default:
            v1 = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
