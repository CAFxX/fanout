package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 int64) int8 {
    switch "foo bar" {
    case "":
        v4 := []int{32767, 0, 9223372036854775807}
        _ = v4
    case "a":
        var v5 int16 = -784
        _ = v5
    case "hello":
        var v6 int64 = p0
        _ = v6
    default:
        p1 += 9223372036854775807
    }
    p0 = p1
    for fi7 := 0; fi7 < 1; fi7++ {
        p1 += (p1 % (((p1 + (p0 / (((^(p0))) | 1)))) | 1))
        _ = fi7
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []int{128, 128, 32768}
    var v2 int = 32767
    var v3 int32 = -1
    if (42 > 41) {
        v3 = v3
        fmt.Println(v3)
        var v8 int8 = -128
        _ = v8
    }
    switch -799 {
    case 32767:
        v9 := GenS1{F0: 0, F1: 18446744073709551615, F2: 2}
        _ = v9
    default:
        var v10 bool = (("hello" + "a") != "\n\t")
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if len(v1) > 0 {
            v1[0] = v2
        }
        _ = fi11
    }
    var v12 uint32 = 241
    for fi13 := 0; fi13 < 11; fi13++ {
        v2 = (v2 % ((9223372036854775807) | 1))
        var v14 float32 = float32(0.1)
        v15 := GenS1{F0: -12, F1: 860, F2: 42}
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = v2
    var v16 float32 = 3.14159
    if (!false) {
        v17 := GenS2{F0: 0}
        _ = v17
    }
    fmt.Println(v2, v3, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
}
