package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
func genFn0(p0 int8, p1 int32, p2 int) uint64 {
    p0 += (p0 - (p0 + 100))
    v7 := []int{p2}
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 839
}
func genFn1(p0 uint64) uint {
    var v9 float32 = 1e10
    var v10 uint8 = uint8(p0)
    _ = p0
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    var v4 int = 128
    v5 := make(map[string]bool)
    v6 := []uint32{2, 1, 65535, 851}
    v3 += v3
    v11 := GenS1{F0: (((((((((((("hello" + "") + "foo bar") + "\n\t") + ("\n\t" + (("" + "a") + "foo bar"))) + "\n\t") + "") + ("a" + ("\n\t" + "hello"))) + "foo bar") + "") + "a") + "hello") + ""), F1: false, F2: 447}
    v1 -= 1
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    switch 32768 {
    case 0:
        v12 := GenS2{F0: (v1 ^ 46), F1: -1, F2: -1}
        _ = v12
    case 1:
        if (-455 < 221) {
            var v13 bool = false
            v1 -= (v1 << ((255) & 63))
            _ = v13
        }
    case 127:
        fmt.Println(v1)
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v1 = genFn0(127, 0, (v4 | (v4 - 128)))
        _ = fi14
    }
    v15 := GenS2{F0: v1, F1: 1000, F2: -1}
    v3 = v3
    var v16 int8 = 1
    if false {
        if false {
            v3 = ((float32(((0.1 + 1e10) * float64(((v3 * v3) - v3)))) + 0.0) - -2.25)
        }
    } else {
        var v17 uint8 = 2
        if (!(v1 > (^(v1)))) {
            var v18 uint64 = genFn0(v16, 42, 32767)
            var v19 float64 = ((1.5 * 0.1) * ((3.14159 + (-2.25 - 1.5)) + (3.14159 / -2.25)))
            _ = v18
            _ = v19
        }
        _ = v17
    }
    if true {
        if false {
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 float32 = float32((float64(v3) - float64(v3)))
                _ = len(v6)
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 float32 = (v3 * v3)
                _ = len(v6)
                _ = fi22
                _ = v23
            }
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(v1, v3)
        var v25 uint16 = 2
        _ = fi24
        _ = v25
    }
    v1 = (v1 - 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
