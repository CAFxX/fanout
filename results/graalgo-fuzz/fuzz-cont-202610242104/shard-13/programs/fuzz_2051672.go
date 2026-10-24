package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint64) int {
    var v4 uint = 65535
    fmt.Println(p1, v4, p1)
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            p1 += p1
            _ = fi6
        }
        _ = fi5
    }
    var v7 int = 127
    v8 := GenS1{F0: (v7 / ((v7) | 1)), F1: 32768}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    _ = v8
    return -889
}
func genFn1(p0 string, p1 int8, p2 int16) string {
    v9 := []int{4294967295, -690, 323, -1}
    p1 -= -128
    var v10 float64 = (0.0 / 1.5)
    p0 = p0
    p2 -= (p2 % ((-1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (((p0 + p0) + ("a" + (p0 + p0))) + ("foo bar" + (((("hello" + p0) + p0) + (p0 + "\n\t")) + "")))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 1
    var v3 float32 = 0.1
    v11 := []int{127}
    switch 42 {
    case 4294967295:
        if (!false) {
            v3 = ((v3 - 1e10) / 1.0)
            v1[4294967295] = 251
        }
    default:
        v12 := GenS2{F0: 687, F1: -861}
        _ = v12
    }
    var v13 float32 = v3
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            v1[4294967295] = 65535
            _ = fi15
        }
        _ = fi14
    }
    switch "" {
    case "a":
        switch 2147483647 {
        case -850:
            v3 = (v13 + 0.1)
        case -1:
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS3{F0: 4294967295, F1: "\n\t"}
                switch (fi16 - 32767) {
                case 32768:
                    var v18 bool = true
                    _ = v18
                default:
                    v3 -= 0.1
                }
                var v19 int16 = 1
                _ = fi16
                _ = v17
                _ = v19
            }
        case -218:
            v1[2147483647] = 2147483647
        default:
            var v20 string = (("" + "\n\t") + (("foo bar" + "foo bar") + "hello"))
            _ = v20
        }
    case "foo bar":
        v21 := GenS4{F0: 32767}
        _ = v21
    default:
        fmt.Println(v2, v2)
    }
    var v22 uint16 = 2
    v22 -= (v22 ^ v22)
    v13 -= (3.14159 - v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v3
}
