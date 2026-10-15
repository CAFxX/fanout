package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 2
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 751
    fmt.Println(v8)
    _ = v8
    return 697
}
func genFn1(p0 int, p1 uint64, p2 int32) string {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = 3.14159
        v11 := GenS1{F0: 2}
        fmt.Println(p0)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := []int{(p0 + (fi12 >> ((745) & 63)))}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 int8 = 101
    var v2 int16 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    var v5 uint16 = 256
    v3[2] = -2147483648
    v14 := GenS2{F0: 42, F1: -61, F2: v1}
    switch "" {
    case "hello":
        v5 += 42
    case "":
        switch "hello" {
        case "a":
            v15 := GenS3{F0: 0, F1: (v1 - (v1 ^ v1)), F2: 42}
            _ = v15
        case "\n\t":
            v2 = -32768
        case "foo bar":
            v16 := []int{-1, -57, -2147483648}
            _ = v16
        }
    case "foo bar":
        var v17 string = "foo bar"
        _ = v17
    default:
        var v18 uint = 18446744073709551615
        _ = v18
    }
    v19 := GenS2{F0: 0, F1: int8(v5), F2: (v1 & (v1 << ((0) & 63)))}
    if false {
        for fi20 := 0; fi20 < 7; fi20++ {
            v5 = (v5 % ((v5) | 1))
            _ = fi20
        }
        var v21 int32 = -2147483648
        _ = v21
    } else {
        v1 = (v1 ^ -128)
        v4[2] = "a"
    }
    var v22 uint8 = 127
    if (float32(0.1) <= 0.0) {
        if false {
            fmt.Println(v5, v2)
            v23 := []int{}
            _ = v23
        } else {
            v24 := GenS4{F0: genFn0()}
            v5 = uint16(v22)
            var v25 uint = 4294967295
            _ = v24
            _ = v25
        }
        fmt.Println(v1, v1, v22)
    } else {
        fmt.Println(v22, v1)
        v5 -= (v5 / (((v5 % ((v5) | 1))) | 1))
        switch "a" {
        case "hello":
            v26 := GenS1{F0: v5}
            _ = v26
        case "a":
            var v27 bool = true
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS5{F0: genFn0()}
        v30 := GenS2{F0: (v22 | 42), F1: v1, F2: v1}
        fmt.Println(v2, fi28)
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v5)
    v31 := []int{1}
    var v32 int = 911
    v33 := v3[1]
    if (0 < 4294967295) {
        if (1 != 871) {
            switch (v33 >> ((255) & 63)) {
            case 32768:
                if len(v31) > 0 {
                    v31[0] = (-(v32))
                }
            default:
                for fi34 := 0; fi34 < 16; fi34++ {
                    switch (("" + (("foo bar" + ((("foo bar" + ((("a" + "a") + "hello") + "hello")) + "a") + "\n\t")) + "a")) + "foo bar") {
                    case "\n\t":
                        var v35 int8 = v1
                        _ = v35
                    case "hello":
                        fmt.Println(v22)
                    case "":
                        v36 := GenS3{F0: 83, F1: (v1 ^ v1), F2: 255}
                        _ = v36
                    }
                    if (909 < 343) {
                        fmt.Println(v33, fi34, v5)
                    } else {
                        var v37 int16 = 286
                        _ = v37
                    }
                    v5 += (v5 | (v5 * (^(v5))))
                    _ = fi34
                }
            }
        }
        fmt.Println(v33)
    }
    _ = len(v3)
    if (!true) {
        v5 = 255
        v5 += (v5 | 2)
        fmt.Println(v5, v22, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
