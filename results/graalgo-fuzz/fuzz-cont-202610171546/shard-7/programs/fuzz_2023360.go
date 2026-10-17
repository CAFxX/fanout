package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int32) int64 {
    p1 = 1
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = p1
        var v7 int16 = -32768
        v8 := GenS1{F0: 220, F1: 1, F2: (v7 % (((v7 * v7)) | 1))}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int8 = 107
    v10 := GenS2{F0: 43}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -1
}
func genFn1() uint32 {
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint = 843
        fmt.Println(v12)
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 65535
    var v14 uint = 847
    var v15 int32 = -2147483648
    v16 := GenS1{F0: 126, F1: 127, F2: -530}
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return (v13 * 2147483648)
}

func main() {
    var v1 uint8 = 248
    v2 := []int{4294967295}
    v3 := []string{"hello", "foo bar", "", ("foo bar" + "foo bar")}
    v4 := make(map[string]int)
    fmt.Println(v1)
    v17 := GenS1{F0: 355, F1: -107, F2: -850}
    var v18 uint64 = 964
    v18 += v18
    var v19 int64 = genFn0((v1 % ((v1) | 1)), 0)
    fmt.Println(v19)
    v18 += v18
    var v20 uint = 997
    v19 = (v19 | v19)
    v21 := GenS1{F0: (v18 >> (((v20 | v20)) & 63)), F1: 42, F2: 1}
    v22 := GenS1{F0: v18, F1: 56, F2: 1000}
    v18 -= (v18 + (v18 << (((v20 % ((18446744073709551615) | 1))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
