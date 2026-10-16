package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint16; F2 bool }
func genFn0(p0 int64, p1 string) int8 {
    v6 := GenS1{F0: 955, F1: 32767, F2: p1}
    p0 = -1
    v7 := GenS1{F0: 65535, F1: -1, F2: "a"}
    p1 = "hello"
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: (p1 + p1)}
        if false {
            var v10 float32 = (3.14159 + 3.14159)
            fmt.Println(p1, v10, v10)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 uint64) string {
    var v11 string = (("foo bar" + (("a" + "a") + "")) + ((("" + "a") + ((("foo bar" + ("\n\t" + "")) + "a") + "a")) + "\n\t"))
    var v12 float64 = 1.5
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return v11
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = (0.0 / float32(3.14159))
    v3 := []int{9223372036854775807, -130, 320}
    var v4 bool = false
    var v5 int16 = -32768
    fmt.Println(v5, v4)
    if v4 {
        _ = len(v3)
        for fi13 := 0; fi13 < 17; fi13++ {
            v1 = v1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi13
        }
    } else {
        var v14 uint32 = 1
        var v15 int16 = (v5 | 453)
        _ = v14
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    var v17 int8 = 58
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v18 := GenS3{F0: 173, F1: 1, F2: (((("foo bar" >= "hello") && (!(215 == 127))) || v4) && v4)}
    var v19 float64 = -2.25
    v16 += -1
    switch "foo bar" {
    case "\n\t":
        var v20 int64 = (^(v16))
        _ = v20
    default:
        switch "hello" {
        case "":
            v5 -= (v5 + -1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
