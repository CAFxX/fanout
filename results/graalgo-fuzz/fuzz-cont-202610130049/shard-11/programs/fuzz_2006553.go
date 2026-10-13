package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0() int32 {
    switch "a" {
    case "foo bar":
        v6 := []int{}
        _ = v6
    case "":
        v7 := []int{-9223372036854775808}
        _ = v7
    case "hello":
        var v8 bool = false
        _ = v8
    default:
        switch "foo bar" {
        case "a":
            var v9 int8 = -128
            _ = v9
        default:
            var v10 int = 242
            _ = v10
        }
    }
    v11 := GenS1{F0: (42 != -84), F1: 42, F2: 537}
    var v12 int8 = -128
    var v13 float32 = 0.1
    _ = v11
    _ = v12
    _ = v13
    return 2147483647
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{42, 227, 4294967295}
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    var v5 uint64 = 103
    var v14 uint64 = v5
    fmt.Println(v5, v14)
    var v15 float32 = 0.0
    var v16 string = (("\n\t" + "\n\t") + ("\n\t" + (("a" + "foo bar") + "a")))
    v14 -= (v14 << ((217) & 63))
    var v17 float32 = float32((1.5 - 1e10))
    v14 += (v14 - (v14 << ((1) & 63)))
    fmt.Println(v17, v15)
    var v18 uint32 = 65535
    var v19 uint64 = 2
    var v20 bool = (v17 > v17)
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v14, v15)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
    }
    if ((!v20) && ((v18 >> ((18446744073709551615) & 63)) == v18)) {
        v22 := v3[4294967295]
        _ = v22
    } else {
        v5 = v19
        v23 := GenS2{F0: (^(v18))}
        _ = v23
    }
    v20 = (((0 < 1) && v20) && v20)
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            switch v16 {
            case "foo bar":
                var v26 uint16 = 42
                _ = v26
            case "hello":
                fmt.Println(v16, v15)
            case "\n\t":
                var v27 uint16 = 0
                _ = v27
            default:
                var v28 uint8 = 246
                _ = v28
            }
            v16 = "hello"
            _ = fi25
        }
        _ = fi24
    }
    v5 -= v19
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
