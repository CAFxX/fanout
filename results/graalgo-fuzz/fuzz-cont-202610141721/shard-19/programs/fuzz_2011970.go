package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
func genFn0(p0 uint, p1 uint) bool {
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(fi6, fi6, p1)
        _ = fi6
    }
    if true {
        fmt.Println(p0, p1)
        fmt.Println(p1, p1)
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 65535}
            _ = fi7
            _ = v8
        }
    }
    if (399 >= 733) {
        v9 := GenS1{F0: 363}
        _ = v9
    }
    var v10 uint = 661
    var v11 uint64 = 510
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return false
}
func genFn1(p0 uint64, p1 string) uint16 {
    v12 := GenS2{F0: -229}
    fmt.Println(p1, p1)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 uint16 = 1
    v5 := []uint32{0, 42}
    if true {
        v4 = (v4 + (v4 | v4))
    }
    _ = len(v3)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint64 = 155
        v15 := GenS2{F0: -1}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint8 = 42
    v4 -= v4
    var v17 int32 = 1
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS3{F0: 255, F1: 2}
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: (^(v17))}
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(fi21, fi21)
        v16 = v16
        _ = fi21
    }
    fmt.Println(v17, v17, v17)
    v22 := GenS4{F0: genFn1(0, "hello"), F1: 2147483647, F2: "foo bar"}
    if true {
        v16 = v16
        v17 = v17
    } else {
        var v23 uint16 = (v4 + v4)
        var v24 uint32 = 0
        _ = v23
        _ = v24
    }
    if false {
        v25 := GenS3{F0: (v4 >> ((1) & 63)), F1: 4294967295}
        var v26 int16 = -32768
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v17 += v17
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float64 = ((((1.5 / 1.0) / 3.14159) - 1.5) - (1.5 - 1.5))
            fmt.Println(v16, fi28)
            _ = fi28
            _ = v29
        }
        var v30 int16 = 32767
        _ = fi27
        _ = v30
    }
    v16 -= 2
    v16 = v16
    v16 -= v16
    v17 += 176
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
