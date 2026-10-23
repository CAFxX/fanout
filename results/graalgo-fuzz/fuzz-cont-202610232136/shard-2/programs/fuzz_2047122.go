package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 uint8; F2 int }
func genFn0(p0 uint32) int16 {
    var v4 int8 = -128
    v5 := []int{9223372036854775807, 263}
    v4 += (v4 | 1)
    v6 := GenS1{F0: (^(p0)), F1: -284}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float64 = -2.25
        p0 -= 65535
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 52
    v3 := make(map[int]int)
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: v2, F1: fi9}
        v1 = 128
        var v11 int64 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 2147483647, F1: -242}
    v13 := GenS1{F0: (v2 << ((18446744073709551615) & 63)), F1: -905}
    _ = len(v3)
    if true {
        v14 := GenS1{F0: v2, F1: int(v2)}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            v2 += (v2 << ((871) & 63))
            var v16 float64 = 0.1
            v16 = 1e10
            _ = fi15
            _ = v16
        }
        var v17 string = ("" + "")
        _ = v17
    }
    v18 := []int{967, 651, -536}
    for fi19 := 0; fi19 < 9; fi19++ {
        v1 += (^(v1))
        v20 := GenS2{F0: (v2 & v2), F1: 208, F2: (fi19 / ((fi19) | 1))}
        _ = fi19
        _ = v20
    }
    var v21 string = ("foo bar" + (("foo bar" + "hello") + "foo bar"))
    if false {
        v21 = "foo bar"
        var v22 uint16 = 0
        v2 += v2
        _ = v22
    }
    fmt.Println(v2, v1)
    var v23 uint8 = uint8(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
}
