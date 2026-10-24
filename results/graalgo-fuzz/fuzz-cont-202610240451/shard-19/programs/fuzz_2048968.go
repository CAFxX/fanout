package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0() uint16 {
    var v5 uint = 352
    v6 := []int{-611, 0}
    var v7 int32 = -1
    switch "" {
    case "foo bar":
        v8 := GenS1{F0: 1, F1: 56}
        _ = v8
    case "":
        v6 = append(v6, -9223372036854775808)
    default:
        var v9 uint64 = 2
        _ = v9
    }
    var v10 uint8 = 255
    _ = v10
    _ = v5
    _ = v6
    _ = v7
    return 32767
}
func genFn1() bool {
    var v11 float64 = -2.25
    var v12 uint32 = 2147483647
    fmt.Println(v11, v12)
    var v13 float64 = float64(1.5)
    v14 := GenS2{F0: 9223372036854775807}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return false
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]uint32)
    var v3 int16 = 32767
    var v4 int = 128
    switch "a" {
    case "":
        var v15 bool = false
        _ = v15
    case "hello":
        v3 += 1000
    case "\n\t":
        v4 = v4
    default:
        if (!false) {
            switch "" {
            case "":
                v16 := GenS2{F0: (v1 ^ (v1 & -1))}
                _ = v16
            case "hello":
                var v17 int32 = 316
                _ = v17
            default:
                v3 = v3
            }
            v3 -= (v3 << ((527) & 63))
        }
    }
    v1 -= (v1 ^ (v1 >> ((255) & 63)))
    var v18 int = v4
    var v19 int = (v4 - v18)
    switch (v4 + v19) {
    case 1:
        var v20 string = ("a" + "")
        _ = v20
    case 9223372036854775807:
        var v21 int = (v18 << ((4294967295) & 63))
        _ = v21
    case 32768:
        var v22 uint = 42
        _ = v22
    default:
        var v23 float64 = ((((3.14159 - 1e10) + -2.25) + 0.0) + 0.1)
        _ = v23
    }
    v24 := []int{(v4 * (v19 | v19)), v18, (-(v19)), v18}
    v25 := GenS2{F0: (v1 + (-(v1)))}
    v26 := GenS2{F0: -1}
    var v27 float32 = (-2.25 - ((-2.25 / 1.0) + 3.14159))
    v1 -= (v1 + (v1 & (-(v1))))
    var v28 int8 = 42
    var v29 float32 = -2.25
    var v30 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
