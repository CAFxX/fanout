package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0() int64 {
    v7 := GenS1{F0: "hello", F1: -1, F2: 32767}
    switch 140 {
    case -286:
        v8 := []int{-1, 4294967295, 14, 2147483647}
        _ = v8
    case -697:
        var v9 string = ("\n\t" + ("" + "\n\t"))
        _ = v9
    case 2147483647:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: "foo bar", F1: 64, F2: 1}
            fmt.Println(fi10)
            v12 := GenS1{F0: ((("\n\t" + "\n\t") + ("\n\t" + "a")) + ""), F1: -28, F2: 173}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if (2147483647 >= 226) {
        var v13 uint8 = 172
        _ = v13
    }
    var v14 uint16 = 600
    v15 := GenS2{F0: 1}
    _ = v14
    _ = v15
    _ = v7
    return -9223372036854775808
}
func genFn1() uint64 {
    v16 := []int{32767}
    v17 := GenS2{F0: 0}
    v18 := GenS1{F0: "", F1: -1, F2: 65535}
    _ = v16
    _ = v17
    _ = v18
    return 724
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]uint32)
    v3 := []int{32768, 719, 2147483647, 9223372036854775807}
    v4 := make(map[uint32]bool)
    var v5 int = 1
    var v6 int64 = -1
    var v19 uint64 = 4294967295
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 int = (-(v5))
    v21 := GenS1{F0: (("" + "foo bar") + "\n\t"), F1: 18, F2: 65535}
    v22 := GenS2{F0: 2147483647}
    var v23 uint8 = 42
    var v24 uint16 = 255
    v23 = (v23 - (v23 | 255))
    for fi25 := 0; fi25 < 17; fi25++ {
        v26 := GenS1{F0: ("" + "a"), F1: 78, F2: v24}
        v27 := GenS1{F0: "foo bar", F1: 78, F2: (v24 + v24)}
        var v28 uint = 18446744073709551615
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 uint8 = v23
    fmt.Println(v29, v23)
    _ = len(v2)
    v1 += 0.0
    fmt.Println(v5)
    switch "foo bar" {
    case "foo bar":
        v1 = float64((0.1 - ((1e10 / float32(v1)) * (0.0 * 3.14159))))
    case "a":
        v24 = 629
    case "hello":
        fmt.Println(v6, v1)
    default:
        v23 = v23
    }
    var v30 int64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
