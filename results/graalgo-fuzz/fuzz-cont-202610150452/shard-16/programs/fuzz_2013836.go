package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 int) uint8 {
    var v4 float32 = float32(3.14159)
    v5 := GenS1{F0: "", F1: 65535}
    v4 = v4
    if false {
        p0 = p0
    } else {
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    v7 := []int{p0, (p0 / (((p0 | p0)) | 1)), p0, p0}
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    return 253
}
func genFn1(p0 uint32, p1 int64) uint8 {
    var v8 float64 = float64(-2.25)
    fmt.Println(p0, v8, v8)
    v9 := []int{-286, 1}
    v10 := GenS1{F0: "\n\t", F1: (p0 + p0)}
    v11 := GenS1{F0: "hello", F1: (p0 | (^(p0)))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -18
    var v3 bool = true
    var v12 float64 = 0.1
    v13 := GenS1{F0: ("hello" + "hello"), F1: 928}
    if (v2 <= v2) {
        v2 -= (v2 ^ (v2 >> ((1) & 63)))
        for fi14 := 0; fi14 < 1; fi14++ {
            v12 += 1.5
            v15 := GenS1{F0: "hello", F1: 2}
            if true {
                v12 = float64(((0.1 * 1.5) + 1.5))
            }
            _ = fi14
            _ = v15
        }
        v16 := v1[2]
        _ = v16
    }
    switch ("a" + "") {
    case "foo bar":
        v17 := GenS1{F0: "a", F1: 2147483647}
        _ = v17
    default:
        var v18 int64 = 1
        _ = v18
    }
    var v19 float64 = (float64(1.5) - 0.1)
    var v20 int32 = -1
    for fi21 := 0; fi21 < 17; fi21++ {
        v20 = v20
        switch "\n\t" {
        case "a":
            if (42 == 1) {
                fmt.Println(fi21, v19)
            }
        case "":
            var v22 uint64 = 42
            _ = v22
        case "hello":
            v1[42] = ("\n\t" + "hello")
        default:
            v23 := []int{32767, (fi21 / ((-2147483648) | 1))}
            _ = v23
        }
        _ = fi21
    }
    v12 = v12
    switch 0 {
    case 32768:
        fmt.Println(v12)
    }
    v24 := GenS1{F0: (("foo bar" + "hello") + "foo bar"), F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
