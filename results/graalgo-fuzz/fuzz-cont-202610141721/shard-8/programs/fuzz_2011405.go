package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0() bool {
    switch ("hello" + ((((("a" + (("a" + ("hello" + (("a" + ("\n\t" + ("\n\t" + ""))) + ""))) + "\n\t")) + "hello") + "foo bar") + ("\n\t" + "hello")) + "a")) {
    case "foo bar":
        v5 := []int{-488}
        _ = v5
    default:
        var v6 int32 = 2147483647
        _ = v6
    }
    switch 32767 {
    case 9223372036854775807:
        var v7 uint = 18446744073709551615
        _ = v7
    case 0:
        var v8 uint = 2
        _ = v8
    case 4294967295:
        var v9 string = ("" + "foo bar")
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    _ = v10
    return (-215 == -32768)
}

func main() {
    var v1 float32 = ((float32(3.14159) / 1e10) + 0.1)
    v2 := make(map[string]int)
    v3 := []uint32{2, 179, 2147483647, 221}
    var v4 int = -2147483648
    v11 := GenS1{F0: 454, F1: 572}
    v12 := GenS1{F0: -1, F1: 4294967295}
    _ = len(v3)
    fmt.Println(v1, v4)
    for fi13 := 0; fi13 < 16; fi13++ {
        switch "a" {
        case "hello":
            v4 -= v4
        }
        delete(v2, "hello")
        _ = fi13
    }
    var v14 int64 = 899
    v1 += float32(1e10)
    v15 := v2[""]
    var v16 int32 = 42
    var v17 bool = true
    switch "a" {
    case "a":
        for fi18 := 0; fi18 < 8; fi18++ {
            fmt.Println(v16)
            _ = fi18
        }
    case "":
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int16 = 189
            v15 = v4
            _ = fi19
            _ = v20
        }
    case "\n\t":
        var v21 float64 = (3.14159 / 3.14159)
        _ = v21
    }
    fmt.Println(v1, v15)
    var v22 int8 = -82
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS2{F0: 58}
        var v25 float32 = 3.14159
        _ = fi23
        _ = v24
        _ = v25
    }
    switch "hello" {
    case "hello":
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 float32 = (0.1 - (v1 - (v1 / 3.14159)))
            _ = len(v2)
            _ = fi26
            _ = v27
        }
    case "a":
        v28 := GenS3{F0: v14, F1: v22, F2: 4294967295}
        _ = v28
    case "\n\t":
        var v29 int16 = 32767
        _ = v29
    default:
        v22 = (v22 << ((18446744073709551615) & 63))
    }
    if (false && ((v22 - (-(v22))) < 0)) {
        v15 -= 2147483647
    } else {
        v17 = v17
        v30 := GenS4{F0: 255, F1: 42}
        if (!v17) {
            var v31 uint16 = 2
            fmt.Println(v14)
            _ = v31
        } else {
            if (!v17) {
                switch 32767 {
                case -161:
                    v1 += ((((1.5 - 1.5) - v1) / float32(0.1)) + 3.14159)
                case -2147483648:
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v33 := GenS1{F0: 1, F1: 255}
                        _ = fi32
                        _ = v33
                    }
                case 1:
                    v14 += (-(v14))
                }
            } else {
                _ = len(v2)
            }
        }
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
