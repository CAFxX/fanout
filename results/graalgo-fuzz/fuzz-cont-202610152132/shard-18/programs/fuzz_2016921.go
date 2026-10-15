package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
func genFn0(p0 uint8, p1 int8, p2 uint8) uint16 {
    switch ("" + "a") {
    case "hello":
        fmt.Println(p2, p0)
    default:
        var v7 uint64 = 557
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 677
}
func genFn1(p0 int8, p1 int32, p2 int) int32 {
    switch (-(p2)) {
    case 405:
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{fi9, (fi9 % ((fi9) | 1))}
            _ = fi9
            _ = v10
        }
    }
    p0 = -83
    if (p0 > p0) {
        v11 := GenS1{F0: p0, F1: (("foo bar" + (("a" + "foo bar") + (("" + ((("hello" + "") + "\n\t") + "")) + (("\n\t" + "") + "foo bar")))) + "hello"), F2: 128}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := []int{32767, 127, 32767}
    var v2 bool = true
    var v3 int = 128
    var v4 int16 = -32768
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    if false {
        fmt.Println(v4, v4, v4)
    } else {
        if len(v1) > 0 {
            v1[0] = v3
        }
    }
    var v12 int16 = 1
    var v13 int16 = v12
    v12 -= (v4 + (v4 + v12))
    v2 = false
    for fi14 := 0; fi14 < 6; fi14++ {
        v2 = v2
        _ = fi14
    }
    var v15 bool = v2
    switch v3 {
    case 1:
        switch "hello" {
        case "\n\t":
            delete(v6, ("" + ("\n\t" + "a")))
        }
    default:
        v6["a"] = 2
    }
    v2 = v15
    for fi16 := 0; fi16 < 8; fi16++ {
        switch v3 {
        case 844:
            switch (fi16 / ((-375) | 1)) {
            case -1:
                v15 = false
            default:
                v2 = true
            }
        default:
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 int32 = 816
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    var v19 uint = 1
    if (v13 < v13) {
        v19 = 14
        if false {
            v19 = (^(v19))
        } else {
            fmt.Println(v4)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
