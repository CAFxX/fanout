package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 42
        if false {
            var v8 uint8 = 1
            v8 += v7
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    var v9 int8 = -128
    fmt.Println(v9)
    switch ("foo bar" + "foo bar") {
    case "a":
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    default:
        var v12 int16 = -44
        _ = v12
    }
    var v13 float64 = (0.0 - -2.25)
    _ = v13
    _ = v9
    return 32767
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = (0.0 - 0.1)
    var v3 int16 = -7
    var v4 uint64 = 654
    v5 := make(map[string]bool)
    v4 -= (v4 << ((65535) & 63))
    switch "a" {
    case "a":
        v14 := []int{32767, -118}
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint32 = 0
            _ = fi15
            _ = v16
        }
    default:
        var v17 uint8 = 242
        _ = v17
    }
    v2 -= v2
    v18 := []int{44}
    var v19 uint64 = (v4 << ((4294967295) & 63))
    var v20 uint = 255
    if false {
        v4 = v19
        v21 := GenS1{F0: 2147483647, F1: 64}
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v20 = uint(v19)
        v1 = (v1 - (v1 % ((2) | 1)))
        _ = fi22
    }
    fmt.Println(v20)
    v2 += v2
    fmt.Println(v19)
    switch ("" + "hello") {
    case "foo bar":
        fmt.Println(v3, v2)
    case "hello":
        fmt.Println(v1, v3, v4)
    }
    switch 965 {
    case 224:
        if (true || (false || true)) {
            var v23 int32 = 65
            switch -847 {
            case 1:
                v4 = v4
            }
            var v24 int16 = (v3 << ((314) & 63))
            _ = v23
            _ = v24
        } else {
            if true {
                v25 := GenS1{F0: 4294967295, F1: 127}
                _ = len(v18)
                _ = v25
            }
        }
    case -562:
        switch -9223372036854775808 {
        case 2147483647:
            v2 = (v2 - 1.5)
        default:
            _ = len(v18)
        }
    case -208:
        v26 := v5[("a" + ("\n\t" + ""))]
        _ = v26
    default:
        var v27 float32 = 3.14159
        _ = v27
    }
    switch 9223372036854775807 {
    case 268:
        if (42 != 0) {
            var v28 int = -9223372036854775808
            _ = v28
        }
    case 9223372036854775807:
        var v29 float64 = -2.25
        _ = v29
    case 594:
        _ = len(v18)
    }
    v30 := v5[("\n\t" + "foo bar")]
    v5[""] = false
    v31 := GenS1{F0: 42, F1: -128}
    var v32 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
