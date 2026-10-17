package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int64; F1 string }
func genFn0() int8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := []int{-973, 2147483647, -899, 1}
    _ = v7
    return 127
}
func genFn1() uint {
    v8 := []int{127, 828, -641, 128}
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            if len(v8) > 0 {
                v8[0] = fi10
            }
            var v11 int64 = 9223372036854775807
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    if len(v8) > 0 {
        v8[0] = 9223372036854775807
    }
    var v12 uint32 = 4294967295
    v13 := GenS1{F0: -348, F1: 0}
    _ = v12
    _ = v13
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 2
    var v3 string = ("" + "foo bar")
    var v4 uint8 = 42
    v5 := make(map[uint32]bool)
    var v14 uint32 = v1
    v2 = v2
    var v15 uint16 = 454
    v16 := []int{}
    v3 = (((v3 + ((v3 + v3) + ("hello" + ((("" + (v3 + (("hello" + "hello") + (v3 + (("foo bar" + v3) + v3))))) + v3) + v3)))) + (v3 + ("hello" + ("a" + v3)))) + v3)
    for fi17 := 0; fi17 < 8; fi17++ {
        if len(v16) > 0 {
            v16[0] = (^(fi17))
        }
        if ("" > (v3 + v3)) {
            fmt.Println(v2, v3)
            switch fi17 {
            case 127:
                _ = len(v5)
            case -147:
                v14 += (v1 / ((v1) | 1))
            case 977:
                v5[v14] = true
            default:
                var v18 uint32 = (v14 - (^(v14)))
                _ = v18
            }
        }
        v16 = append(v16, (fi17 ^ (fi17 / ((fi17) | 1))))
        _ = fi17
    }
    v19 := GenS1{F0: -298, F1: 42}
    var v20 int8 = 101
    var v21 uint32 = v14
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if (v4 > (^(v4))) {
        switch v3 {
        case "a":
            if len(v16) > 0 {
                v16[0] = -1
            }
        case "":
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 bool = true
                _ = fi22
                _ = v23
            }
        case "\n\t":
            if false {
                v2 += v2
                var v24 int8 = genFn0()
                _ = v24
            } else {
                v25 := v5[v1]
                _ = v25
            }
        default:
            v26 := GenS2{F0: 42, F1: v3}
            _ = v26
        }
        v27 := GenS2{F0: int64(v20), F1: (v3 + "")}
        v14 -= v21
        _ = v27
    }
    v3 = v3
    switch -2147483648 {
    case -425:
        v15 += v15
    case -2147483648:
        fmt.Println(v4, v21)
    case -453:
        switch 32768 {
        case 66:
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 int32 = -1
                _ = fi28
                _ = v29
            }
        case 42:
            v30 := GenS2{F0: -417, F1: ""}
            _ = v30
        }
    }
    switch v3 {
    case "\n\t":
        delete(v5, (v21 >> ((920) & 63)))
    case "":
        v1 = 42
    default:
        var v31 uint16 = (v15 ^ (v15 * v15))
        _ = v31
    }
    v15 -= uint16(v21)
    var v32 int8 = v20
    fmt.Println(v1, v20, v32)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
