package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16; F1 int64; F2 string }
func genFn0(p0 int8, p1 int32) uint32 {
    switch "hello" {
    case "foo bar":
        p1 = (p1 << ((1) & 63))
    case "":
        v7 := GenS1{F0: 42, F1: "foo bar"}
        _ = v7
    case "a":
        var v8 float64 = 1.5
        _ = v8
    }
    var v9 float32 = (3.14159 * float32(-2.25))
    switch (((("foo bar" + "hello") + "\n\t") + ("" + ((("foo bar" + ("" + ("hello" + "foo bar"))) + "hello") + ("hello" + "hello")))) + ("\n\t" + ("hello" + "\n\t"))) {
    case "foo bar":
        v10 := GenS2{F0: -1, F1: -1, F2: ""}
        _ = v10
    case "a":
        switch -9223372036854775808 {
        case 32768:
            switch "foo bar" {
            case "foo bar":
                v11 := []int{-464, 42, -502}
                _ = v11
            case "\n\t":
                v9 -= 0.1
            case "":
                fmt.Println(v9)
            default:
                v12 := []int{4294967295, -2147483648, -9223372036854775808}
                _ = v12
            }
        case 4294967295:
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := []int{}
                var v15 uint8 = 119
                v16 := GenS1{F0: 158, F1: ""}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        case 32767:
            var v17 float32 = (float32(3.14159) * (v9 * v9))
            _ = v17
        default:
            fmt.Println(p0, p1, p1)
        }
    case "":
        var v18 string = ("foo bar" + "")
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 1
    var v2 int16 = -257
    var v3 uint8 = (v1 % ((227) | 1))
    var v4 uint64 = 18446744073709551615
    var v5 uint = 255
    v6 := []uint32{4294967295}
    fmt.Println(v3, v4)
    fmt.Println(v5, v4, v4)
    fmt.Println(v3, v5, v4)
    v4 += v4
    v2 += v2
    var v19 float32 = (0.0 / 1e10)
    var v20 int32 = 1
    var v21 float32 = (v19 / v19)
    v6 = append(v6, 95)
    if false {
        fmt.Println(v1, v3, v21)
    }
    if len(v6) > 0 {
        v6[0] = 2
    }
    v6 = append(v6, 65535)
    v5 += (^(v5))
    v3 += (v1 ^ 128)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
