package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
func genFn0(p0 uint8) int16 {
    v5 := []int{404, -990}
    v6 := GenS1{F0: 819, F1: (^(p0)), F2: (^(p0))}
    if false {
        var v7 uint32 = 2
        var v8 int8 = -1
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 int8) int16 {
    p1 -= (p1 * (p1 >> ((65535) & 63)))
    switch 2147483647 {
    case 128:
        var v9 int32 = 42
        _ = v9
    case -2147483648:
        fmt.Println(p2, p2, p1)
    case 4294967295:
        var v10 int = 127
        _ = v10
    default:
        var v11 uint = 65535
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 9223372036854775807
    var v3 uint8 = 128
    var v4 int = 128
    v2 -= -9223372036854775808
    var v12 uint8 = (^(v3))
    if (false && true) {
        var v13 int32 = 42
        _ = v13
    }
    var v14 uint8 = v3
    switch 442 {
    case 527:
        v15 := GenS1{F0: 143, F1: v12, F2: v14}
        _ = v15
    case 128:
        switch "\n\t" {
        case "":
            v16 := []int{}
            _ = v16
        }
    default:
        fmt.Println(v2, v12)
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint32 = 265
        var v19 int32 = int32(v2)
        _ = fi17
        _ = v18
        _ = v19
    }
    if (false && false) {
        fmt.Println(v12, v14)
        for fi20 := 0; fi20 < 6; fi20++ {
            v3 = (^(v12))
            _ = fi20
        }
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: 1, F1: v3, F2: (v12 >> ((4294967295) & 63))}
        switch (("hello" + (("foo bar" + "a") + (("" + "") + "\n\t"))) + "a") {
        case "\n\t":
            var v23 uint32 = 42
            _ = v23
        default:
            var v24 uint32 = 2147483648
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v14)
    for fi25 := 0; fi25 < 9; fi25++ {
        v14 += v12
        var v26 uint = 255
        _ = fi25
        _ = v26
    }
    if true {
        v27 := GenS2{F0: 0, F1: v2, F2: 2}
        _ = v27
    }
    v3 -= (v12 << ((980) & 63))
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint64 = 219
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: 513, F1: v3, F2: 1}
    v31 := []int{}
    var v32 int64 = 42
    var v33 uint = 65535
    var v34 uint = v33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
