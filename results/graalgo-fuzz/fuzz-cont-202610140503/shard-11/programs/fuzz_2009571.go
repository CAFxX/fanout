package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int64, p1 uint, p2 int32) bool {
    p0 -= (p0 % ((p0) | 1))
    if (1e10 <= (((((1e10 + -2.25) - float32(((1e10 * 1.5) * -2.25))) + float32((1.5 * 3.14159))) / 1.5) + 0.1)) {
        switch ("" + (("foo bar" + "\n\t") + "foo bar")) {
        case "":
            v6 := []int{-398, -63}
            _ = v6
        case "\n\t":
            if false {
                fmt.Println(p1)
                v7 := GenS1{F0: ("hello" + "hello"), F1: "a", F2: 1}
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(p2, fi8)
                    var v9 float32 = -2.25
                    fmt.Println(p2, p0)
                    _ = fi8
                    _ = v9
                }
                _ = v7
            }
        case "a":
            v10 := []int{0}
            _ = v10
        default:
            p1 += (p1 ^ (p1 << (((^(p1))) & 63)))
        }
        var v11 uint64 = 4294967295
        if (32767 < -2147483648) {
            v11 -= v11
            var v12 int = 127
            _ = v12
        }
        _ = v11
    }
    p0 = (p0 - -1)
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    var v5 bool = (42 < 42)
    v13 := GenS2{F0: 1}
    var v14 int8 = 0
    var v15 int16 = 1
    var v16 int = 9223372036854775807
    fmt.Println(v5, v5)
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 float64 = (v1 + v1)
        var v19 float64 = float64(1e10)
        if (0 < 529) {
            var v20 int32 = 0
            v20 = -1
            _ = v20
        } else {
            v14 -= -18
            if (664 == 65535) {
                fmt.Println(v5, v5, fi17)
                var v21 float32 = float32(v18)
                var v22 uint = 1
                _ = v21
                _ = v22
            } else {
                v5 = v5
            }
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    v14 += (v14 * v14)
    fmt.Println(v14, v5)
    if (!v5) {
        v23 := v3[("hello" + "")]
        _ = v23
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                v5 = v5
                var v26 int16 = v15
                v27 := GenS3{F0: 2147483647, F1: 0}
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := GenS2{F0: 0}
            _ = fi24
            _ = v28
        }
        v29 := GenS2{F0: 2147483647}
        _ = v29
    }
    switch "" {
    case "a":
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
