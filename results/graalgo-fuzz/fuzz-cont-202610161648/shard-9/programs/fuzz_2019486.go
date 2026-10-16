package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 uint32) bool {
    v6 := []int{}
    p0 = (-(p0))
    fmt.Println(p1, p1)
    switch 32767 {
    case 0:
        if (1 != p1) {
            fmt.Println(p0, p1)
            fmt.Println(p1, p0)
        }
    default:
        p1 = (p1 >> ((313) & 63))
    }
    _ = p0
    _ = p1
    _ = v6
    return false
}
func genFn1(p0 uint8, p1 int64, p2 uint16) int8 {
    p1 += -1
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "foo bar" {
        case "a":
            fmt.Println(fi7)
        default:
            v8 := GenS1{F0: 0, F1: p1}
            _ = v8
        }
        _ = fi7
    }
    var v9 uint8 = (p0 + 19)
    var v10 uint32 = 575
    v11 := GenS2{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{127, 1, 35}
    var v3 int16 = 0
    var v4 int16 = (v3 & -840)
    var v5 uint16 = 2
    switch "hello" {
    case "":
        var v12 string = ""
        _ = v12
    case "foo bar":
        if true {
            var v13 uint64 = 18446744073709551615
            v3 += v4
            _ = v13
        } else {
            switch -307 {
            case 0:
                v14 := GenS1{F0: 2147483648, F1: -1}
                _ = v14
            case 1:
                var v15 int16 = v3
                _ = v15
            default:
                var v16 int8 = 42
                _ = v16
            }
            var v17 uint = 253
            _ = v17
        }
    default:
        var v18 int8 = 127
        _ = v18
    }
    fmt.Println(v3, v3)
    v4 = v4
    if false {
        v19 := GenS1{F0: 903, F1: -951}
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        _ = len(v1)
        if true {
            v4 += v3
            v21 := GenS3{F0: "hello"}
            fmt.Println(v3, v5)
            _ = v21
        } else {
            if false {
                var v22 float64 = ((-2.25 / float64(((((3.14159 * 0.1) / 0.1) + 0.1) / -2.25))) + float64((float32((0.0 + 0.1)) - 0.0)))
                fmt.Println(fi20, v4)
                _ = v22
            }
            v23 := GenS2{F0: 42}
            _ = v23
        }
        _ = fi20
    }
    fmt.Println(v5, v3)
    if true {
        var v24 int16 = 0
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v25 int = 4294967295
            _ = v25
        case -2147483648:
            var v26 int32 = 42
            _ = v26
        }
        _ = v24
    }
    var v27 float64 = 3.14159
    v28 := GenS2{F0: 1}
    fmt.Println(v27, v3, v5)
    var v29 bool = (!false)
    var v30 int16 = (v4 ^ 32767)
    switch ((("foo bar" + "a") + "hello") + "") {
    case "hello":
        v3 = 32767
    default:
        if false {
            v29 = v29
        } else {
            var v31 uint8 = 201
            _ = v31
        }
    }
    v32 := v1[("hello" + "\n\t")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
