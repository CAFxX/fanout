package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 int8) int {
    var v6 uint8 = 255
    var v7 float32 = float32(0.1)
    _ = p0
    _ = v6
    _ = v7
    return -424
}

func main() {
    var v1 uint32 = 14
    var v2 uint64 = 18446744073709551615
    v3 := []string{((("foo bar" + "foo bar") + "a") + "hello"), ("foo bar" + ("hello" + "foo bar"))}
    v4 := make(map[uint32]uint32)
    var v5 float64 = 3.14159
    if len(v3) > 0 {
        v3[0] = (("hello" + ("hello" + ("\n\t" + ""))) + ("a" + ""))
    }
    v2 = 82
    var v8 float32 = ((((-2.25 / 1e10) + (0.1 - 0.1)) + 0.1) + (0.1 - 3.14159))
    v9 := GenS1{F0: -1, F1: -32768}
    v10 := GenS2{F0: genFn0(-1), F1: 65535}
    var v11 float64 = (0.1 * 1e10)
    if (((float64(3.14159) / float64(v8)) <= v11) || true) {
        var v12 uint = 42
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            if false {
                fmt.Println(v8, fi13)
                if (v8 > 0.0) {
                    v11 -= (1.5 * 1e10)
                    var v14 uint16 = 255
                    v2 = v2
                    _ = v14
                }
            }
            _ = fi13
        }
        var v15 float64 = v5
        _ = v15
    }
    var v16 int16 = -32768
    fmt.Println(v16)
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS2{F0: (fi17 >> ((4294967295) & 63)), F1: 65}
        _ = fi17
        _ = v18
    }
    var v19 int64 = 629
    var v20 bool = ((v19 * -1) >= -9223372036854775808)
    for fi21 := 0; fi21 < 9; fi21++ {
        _ = len(v3)
        fmt.Println(v16, v19, v1)
        fmt.Println(fi21)
        _ = fi21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS3{F0: 4294967295, F1: v19}
        _ = fi22
        _ = v23
    }
    var v24 uint32 = v1
    var v25 int64 = -1
    switch "hello" {
    case "foo bar":
        var v26 string = "\n\t"
        _ = v26
    case "a":
        if true {
            v27 := GenS3{F0: 719, F1: -379}
            _ = v27
        } else {
            v3 = append(v3, "\n\t")
        }
    default:
        if (234 >= 2) {
            var v28 int32 = -1
            _ = v28
        }
    }
    v3 = append(v3, ("" + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
