package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
func genFn0(p0 string) uint16 {
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 255
        var v8 uint32 = 42
        if false {
            v9 := GenS1{F0: 9223372036854775807}
            v10 := GenS1{F0: 9223372036854775807}
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 1}
            var v12 uint32 = v8
            var v13 uint64 = 4294967295
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 98
    var v3 string = ("" + (("\n\t" + "\n\t") + "foo bar"))
    v4 := make(map[string]uint32)
    var v14 int = -1
    switch v3 {
    case "\n\t":
        v2 = (v2 % (((v2 - (v2 ^ 18446744073709551615))) | 1))
    default:
        if (-32768 > 1) {
            var v15 string = ("a" + ("foo bar" + v3))
            v16 := []int{(v14 % ((2147483647) | 1))}
            _ = v15
            _ = v16
        }
    }
    var v17 float32 = (((float32(1e10) * -2.25) / -2.25) - 1e10)
    var v18 int8 = 127
    v19 := GenS2{F0: v14, F1: (-(v18)), F2: -1}
    var v20 int8 = -128
    if (!true) {
        v21 := []int{v14, v14}
        var v22 uint = 1
        _ = v21
        _ = v22
    } else {
        var v23 int = (v14 / ((1) | 1))
        var v24 string = v3
        _ = v23
        _ = v24
    }
    var v25 string = (v3 + v3)
    v26 := GenS1{F0: 440}
    v20 -= v18
    for fi27 := 0; fi27 < 6; fi27++ {
        v2 = v2
        _ = fi27
    }
    fmt.Println(v2, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
