package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 bool; F1 int32 }
func genFn0(p0 uint64, p1 uint64, p2 uint32) bool {
    if true {
        fmt.Println(p0, p0)
    }
    v4 := GenS1{F0: 42, F1: 1, F2: 65535}
    v5 := []int{}
    v6 := GenS2{F0: 0, F1: "hello", F2: 2}
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(p0, fi7, p1)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (float64(((3.14159 / 3.14159) - (((float32(((1e10 / 3.14159) + ((1e10 * -2.25) * 3.14159))) / 0.1) * ((0.1 * 3.14159) * float32((((1.5 * 1.5) / 1.0) + 3.14159)))) * 0.0))) > (1e10 + -2.25))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{4294967295, 9223372036854775807, 32768}
    v3 := make(map[string]bool)
    var v8 uint = 65535
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS3{F0: 1000, F1: false}
        v9 -= v9
        v8 += (v8 | v8)
        _ = fi10
        _ = v11
    }
    if false {
        v8 -= v8
        var v12 string = ("foo bar" + "")
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v3["\n\t"] = (((0.0 / 3.14159) / 1e10) >= ((3.14159 + 0.0) / 1.5))
        v14 := GenS2{F0: 42, F1: (("\n\t" + ("" + "foo bar")) + ""), F2: 0}
        v9 += v9
        _ = fi13
        _ = v14
    }
    v9 = 0
    for fi15 := 0; fi15 < 7; fi15++ {
        v9 += v9
        v8 = (v8 << (((v8 >> (((v8 >> ((516) & 63))) & 63))) & 63))
        _ = fi15
    }
    v16 := GenS4{F0: (!true), F1: -1}
    v8 += v8
    if false {
        v17 := GenS4{F0: (-2147483648 > 42), F1: 42}
        _ = v17
    } else {
        v9 = v9
        var v18 int64 = -968
        _ = v18
    }
    fmt.Println(v8, v9)
    var v19 uint32 = 683
    v20 := GenS2{F0: 0, F1: "\n\t", F2: 127}
    delete(v3, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
