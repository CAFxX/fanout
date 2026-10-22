package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int }
func genFn0(p0 uint8) uint {
    v4 := []int{2147483647, 9223372036854775807, 32768, 128}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 13; fi5++ {
        _ = len(v4)
        switch fi5 {
        case 54:
            var v6 bool = true
            _ = v6
        default:
            var v7 int = (fi5 | 128)
            _ = v7
        }
        _ = fi5
    }
    if false {
        v8 := GenS1{F0: 1, F1: 0}
        v9 := GenS1{F0: 0, F1: 0}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: fi10, F1: 918}
            _ = fi10
            _ = v11
        }
        var v12 int16 = 32767
        p0 = p0
        _ = v12
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 uint = 941
    var v2 uint32 = 65535
    var v3 int64 = 1
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: 92, F1: 42}
        var v15 float32 = ((0.1 + 3.14159) * 1.5)
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 = v3
    v16 := []int{0, 981, -2147483648}
    switch -842 {
    case 4294967295:
        v3 += (v3 * (v3 / ((v3) | 1)))
    default:
        v17 := GenS2{F0: 0, F1: "\n\t", F2: 0}
        _ = v17
    }
    var v18 float64 = (0.0 / 1.0)
    v19 := GenS2{F0: 120, F1: (("a" + ("foo bar" + "\n\t")) + "a"), F2: 32767}
    v20 := GenS2{F0: 127, F1: "foo bar", F2: 29}
    var v21 int16 = -243
    var v22 int8 = 100
    for fi23 := 0; fi23 < 18; fi23++ {
        if true {
            if (-2147483648 <= 315) {
                var v24 float64 = v18
                _ = v24
            }
            fmt.Println(v21, v21)
        } else {
            v25 := GenS1{F0: 1, F1: 2}
            var v26 bool = false
            _ = v25
            _ = v26
        }
        var v27 float32 = 0.1
        var v28 bool = (1 >= -1)
        _ = fi23
        _ = v27
        _ = v28
    }
    v22 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
