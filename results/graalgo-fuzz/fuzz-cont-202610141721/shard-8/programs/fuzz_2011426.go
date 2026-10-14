package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0() int {
    var v5 int32 = 0
    v6 := GenS1{F0: -969}
    v5 = v5
    var v7 int = 249
    v8 := []int{}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1() int8 {
    var v9 float32 = 1e10
    v9 = (v9 * v9)
    v10 := GenS2{F0: 912, F1: -2147483648}
    var v11 int16 = 0
    _ = v10
    _ = v11
    _ = v9
    return 107
}

func main() {
    v1 := []string{"foo bar", "foo bar", (((("" + ("foo bar" + ("" + ""))) + ("" + ("" + "foo bar"))) + ("\n\t" + ("hello" + ("\n\t" + ("a" + ("foo bar" + ("hello" + "a"))))))) + "a")}
    v2 := make(map[int]uint32)
    var v3 uint16 = 2
    v4 := []uint32{816, 42}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 float32 = float32(0.1)
        _ = fi12
        _ = v13
    }
    if false {
        v14 := GenS2{F0: -1, F1: 42}
        fmt.Println(v3)
        _ = v14
    }
    var v15 float64 = 0.0
    for fi16 := 0; fi16 < 7; fi16++ {
        switch (fi16 / ((fi16) | 1)) {
        case 32768:
            v15 -= v15
        default:
            if (float32((0.1 - float64((0.0 - -2.25)))) != (float32(-2.25) / 1.0)) {
                fmt.Println(fi16, v15)
            } else {
                v3 = v3
                var v17 float32 = 3.14159
                _ = v17
            }
        }
        v15 -= v15
        _ = fi16
    }
    switch ("\n\t" + "hello") {
    case "hello":
        v15 = (v15 + ((v15 - 0.0) / 1.5))
    default:
        fmt.Println(v15)
    }
    var v18 int8 = -99
    v19 := v2[-1]
    var v20 int = -2147483648
    fmt.Println(v18, v3, v20)
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS1{F0: 32767}
        v3 += v3
        _ = fi21
        _ = v22
    }
    if (-1 >= 1) {
        var v23 int = genFn0()
        fmt.Println(v19)
        _ = v23
    } else {
        v24 := GenS3{F0: -391, F1: 428}
        switch "hello" {
        case "\n\t":
            for fi25 := 0; fi25 < 4; fi25++ {
                v19 = (v19 % ((288) | 1))
                _ = fi25
            }
        default:
            v18 = v18
        }
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
