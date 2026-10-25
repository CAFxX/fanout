package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int8) int {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 42, F1: 950, F2: (p1 >> ((0) & 63))}
        p1 = p1
        var v8 string = "\n\t"
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{(-(fi9))}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        switch fi11 {
        case 290:
            var v12 uint64 = 2
            _ = v12
        }
        fmt.Println(fi11, p0, p1)
        _ = fi11
    }
    p1 -= (p1 | (p1 % ((p1) | 1)))
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint32 = 0
    var v2 string = "foo bar"
    var v3 uint = 2
    var v4 float32 = (float32(-2.25) + 1e10)
    v5 := []uint32{(^(v1)), v1, (v1 - 2147483648)}
    v4 = v4
    if len(v5) > 0 {
        v5[0] = (v1 - (v1 >> (((^(v3))) & 63)))
    }
    v13 := GenS1{F0: v1, F1: -2147483648, F2: 7}
    v1 = (v1 - v1)
    if len(v5) > 0 {
        v5[0] = 1
    }
    _ = len(v5)
    var v14 float32 = v4
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: (v1 ^ v1), F1: -1, F2: 42}
        _ = fi15
        _ = v16
    }
    var v17 int = int(v1)
    _ = len(v5)
    switch v2 {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        var v18 uint8 = 0
        _ = v18
    case "\n\t":
        if false {
            v19 := GenS2{F0: 773, F1: 733}
            v17 += (v17 << (((v3 | 1)) & 63))
            _ = v19
        }
    default:
        v20 := GenS1{F0: v1, F1: 956, F2: 127}
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int64 = 42
        v23 := GenS3{F0: v2}
        _ = fi21
        _ = v22
        _ = v23
    }
    v17 = genFn0(true, 42)
    if len(v5) > 0 {
        v5[0] = (v1 + 914)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
