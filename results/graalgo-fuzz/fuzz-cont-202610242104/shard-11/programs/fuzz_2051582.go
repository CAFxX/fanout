package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{65535}
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    var v6 int16 = -1
    for fi7 := 0; fi7 < 10; fi7++ {
        v3 -= float64(((3.14159 / 3.14159) + (0.0 / 1.5)))
        var v8 string = ""
        switch (fi7 | (^(fi7))) {
        case 0:
            for fi9 := 0; fi9 < 15; fi9++ {
                v8 = ((v8 + v8) + v8)
                var v10 int = (fi9 << ((255) & 63))
                var v11 int32 = -997
                _ = fi9
                _ = v10
                _ = v11
            }
        case -9223372036854775808:
            var v12 int16 = (v6 >> ((18446744073709551615) & 63))
            _ = v12
        default:
            v13 := GenS1{F0: (2 < 149), F1: 4294967295}
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v3)
    _ = len(v1)
    v6 = 895
    v14 := v4[("foo bar" + "hello")]
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = float64((0.1 * (1e10 / ((0.0 / -2.25) - 1.5))))
        _ = fi15
    }
    var v16 uint8 = 2
    if (2 < 702) {
        switch 128 {
        case -1:
            v17 := GenS2{F0: 32767}
            _ = v17
        case 42:
            fmt.Println(v3)
        }
        if (32767 > 2147483647) {
            var v18 float64 = float64((-2.25 * 3.14159))
            _ = v18
        } else {
            v16 += (v16 | v16)
        }
    }
    var v19 int16 = 311
    var v20 string = ("foo bar" + (("a" + "") + "foo bar"))
    var v21 uint16 = 2
    var v22 uint = 255
    switch "foo bar" {
    case "a":
        switch "foo bar" {
        case "":
            v6 += (-(v19))
        case "\n\t":
            v22 -= (v22 >> (((v22 / (((v22 % (((^(v22))) | 1))) | 1))) & 63))
        default:
            var v23 int32 = 42
            _ = v23
        }
    case "hello":
        switch 0 {
        case 308:
            v24 := GenS2{F0: (v21 >> ((v22) & 63))}
            _ = v24
        case 42:
            for fi25 := 0; fi25 < 5; fi25++ {
                v14 = v14
                fmt.Println(v14, v14)
                fmt.Println(v19)
                _ = fi25
            }
        case -29:
            var v26 float32 = ((0.0 + (((float32(v3) + 3.14159) * 0.1) - 1.5)) + 1.5)
            _ = v26
        default:
            v6 += (v19 >> ((v22) & 63))
        }
    case "":
        fmt.Println(v20, v16)
    }
    var v27 int64 = -821
    v21 -= v21
    v19 = v19
    v28 := v1[-2147483648]
    var v29 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
