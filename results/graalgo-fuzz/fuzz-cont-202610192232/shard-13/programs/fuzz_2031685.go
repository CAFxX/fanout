package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
func genFn0(p0 uint8) uint8 {
    p0 = 255
    p0 = (p0 - (p0 + 0))
    v5 := []int{128, 127, -587, 32768}
    var v6 int8 = 100
    var v7 bool = (0 != 2147483647)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (p0 << ((223) & 63))
}
func genFn1(p0 uint16, p1 int) uint {
    switch p1 {
    case 4294967295:
        var v8 float32 = ((0.0 * (float32(1.5) + 0.0)) * 1.5)
        _ = v8
    case 0:
        fmt.Println(p0, p0)
    }
    v9 := GenS1{F0: 0, F1: 42, F2: "a"}
    _ = p0
    _ = p1
    _ = v9
    return 562
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[string]bool)
    var v3 bool = ("\n\t" >= "\n\t")
    var v4 int32 = -1
    var v10 int32 = (v4 ^ (v4 * (v4 / (((v4 * v4)) | 1))))
    fmt.Println(v3)
    v4 += v4
    v11 := v2[("\n\t" + (("\n\t" + "hello") + "a"))]
    fmt.Println(v10)
    switch -855 {
    case -971:
        switch ("" + "") {
        case "":
            fmt.Println(v11)
        case "hello":
            v4 += v4
        case "\n\t":
            v3 = v11
        default:
            if v11 {
                v3 = (0.1 <= 1e10)
            }
        }
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        switch (fi12 * fi12) {
        case -2147483648:
            var v13 uint16 = 255
            _ = v13
        case -472:
            v14 := GenS1{F0: 32767, F1: 127, F2: "hello"}
            _ = v14
        default:
            v10 = v4
        }
        v3 = (65535 > 124)
        v4 = (v4 >> ((185) & 63))
        _ = fi12
    }
    v15 := GenS1{F0: 32767, F1: 1, F2: (("foo bar" + "\n\t") + (("a" + "foo bar") + "hello"))}
    v16 := GenS1{F0: -1, F1: 1, F2: ("a" + ("" + "hello"))}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
