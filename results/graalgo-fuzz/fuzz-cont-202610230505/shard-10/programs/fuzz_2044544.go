package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 int16, p2 uint) int {
    if false {
        fmt.Println(p1, p0)
        v4 := []int{-2147483648, 1, 128}
        _ = v4
    }
    v5 := GenS1{F0: 100}
    v6 := []int{9223372036854775807, 128, -9223372036854775808}
    if (-403 < -9223372036854775808) {
        switch ("" + "\n\t") {
        case "hello":
            var v7 int = -1
            _ = v7
        case "foo bar":
            v8 := GenS1{F0: -1}
            _ = v8
        default:
            fmt.Println(p1, p2, p2)
        }
    }
    v9 := GenS2{F0: 2, F1: 363, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v9
    return 146
}

func main() {
    var v1 float32 = float32(3.14159)
    v2 := []string{"a", "foo bar"}
    var v3 int32 = 42
    v10 := GenS3{F0: 255}
    var v11 int16 = -1
    switch "hello" {
    case "\n\t":
        v3 -= v3
    case "":
        var v12 float64 = float64(float32((-2.25 * 1e10)))
        _ = v12
    default:
        fmt.Println(v11, v1)
    }
    v3 -= (v3 / ((v3) | 1))
    v1 -= v1
    v2 = append(v2, "\n\t")
    v13 := GenS2{F0: 4294967295, F1: 1, F2: 86}
    if len(v2) > 0 {
        v2[0] = "a"
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint8 = 255
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi14
        _ = v15
    }
    var v16 int = -1
    v17 := GenS4{F0: v3}
    var v18 int32 = (v3 >> ((2) & 63))
    var v19 int8 = -1
    v20 := GenS1{F0: 42}
    for fi21 := 0; fi21 < 10; fi21++ {
        v2 = append(v2, "a")
        if true {
            v22 := GenS4{F0: v3}
            switch fi21 {
            case -9223372036854775808:
                var v23 uint16 = 35
                _ = v23
            case 42:
                v18 += 0
            }
            _ = v22
        } else {
            fmt.Println(v11)
        }
        _ = fi21
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
