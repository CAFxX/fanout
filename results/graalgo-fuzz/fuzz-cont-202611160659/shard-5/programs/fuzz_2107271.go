package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint32, p1 int8) int32 {
    v7 := GenS1{F0: -414, F1: true}
    var v8 uint8 = 128
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 uint16) int {
    fmt.Println(p1, p0)
    var v9 uint8 = 167
    switch 32768 {
    case -721:
        var v10 uint = 560
        _ = v10
    case 127:
        v11 := []int{128, 127, 42, -2147483648}
        _ = v11
    default:
        if false {
            switch (("a" + "a") + "\n\t") {
            case "\n\t":
                fmt.Println(p0, p0, v9)
            case "hello":
                fmt.Println(p0, v9)
            case "foo bar":
                v9 += (v9 | 255)
            default:
                fmt.Println(p0)
            }
            v9 = 2
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    v3 := []uint32{2}
    var v4 uint8 = 201
    var v5 bool = (-1 != 9223372036854775807)
    var v6 int16 = -32768
    for fi12 := 0; fi12 < 16; fi12++ {
        _ = len(v1)
        _ = fi12
    }
    v13 := GenS1{F0: 42, F1: (v5 && (0 > 65535))}
    var v14 int8 = 0
    fmt.Println(v6, v14)
    v1["\n\t"] = 42
    delete(v1, "a")
    v6 = v6
    var v15 float64 = -2.25
    v2[("" + "\n\t")] = 42
    var v16 int8 = v14
    var v17 int16 = v6
    v18 := GenS2{F0: 9223372036854775807, F1: 65535}
    fmt.Println(v4, v17, v6)
    switch "a" {
    case "\n\t":
        v15 -= (((v15 / v15) * v15) - v15)
    case "a":
        switch genFn1(255, 107) {
        case 42:
            var v19 uint = 2
            _ = v19
        }
    }
    fmt.Println(v17)
    v16 = int8(v4)
    for fi20 := 0; fi20 < 6; fi20++ {
        switch "foo bar" {
        case "\n\t":
            if (711 <= 745) {
                v5 = v5
            } else {
                for fi21 := 0; fi21 < 15; fi21++ {
                    switch (fi20 << ((942) & 63)) {
                    case 42:
                        v22 := v1[("a" + ("\n\t" + "hello"))]
                        _ = v22
                    default:
                        var v23 float32 = 0.0
                        _ = v23
                    }
                    _ = fi21
                }
                fmt.Println(v6, v5, v16)
            }
        case "hello":
            if v5 {
                var v24 uint8 = 134
                v2[("\n\t" + ("a" + ""))] = 0
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int32 = 618
                    v27 := v2["foo bar"]
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = v24
            }
        case "":
            fmt.Println(v14)
        default:
            _ = len(v3)
        }
        _ = fi20
    }
    v17 += 467
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
