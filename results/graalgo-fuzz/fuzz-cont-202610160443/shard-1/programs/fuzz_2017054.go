package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int64, p2 uint64) int16 {
    if ((!true) || (3.14159 < 1.5)) {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 137
            _ = fi7
            _ = v8
        }
        p0 -= uint16(p2)
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = "foo bar"
            p2 -= (p2 & p2)
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(p0)
                var v12 uint8 = 127
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    } else {
        var v13 int16 = -300
        p1 -= -9223372036854775808
        v14 := []int{int(v13), 32768, -9223372036854775808, 970}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := []int{fi15}
        p2 = p2
        if false {
            var v17 int32 = 439
            _ = v17
        } else {
            v18 := GenS1{F0: 2147483647}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := []int{42, 1, -532}
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 1000
}
func genFn1(p0 uint16, p1 uint8, p2 int64) uint8 {
    for fi20 := 0; fi20 < 2; fi20++ {
        p1 = p1
        _ = fi20
    }
    var v21 uint32 = 2
    p2 -= 159
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return p1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    v5 := []int{32768, 0}
    var v6 int64 = (v4 ^ v4)
    var v22 uint8 = 42
    var v23 int = 285
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 int = 343
        _ = fi24
        _ = v25
    }
    switch "" {
    case "foo bar":
        v26 := GenS1{F0: 437}
        _ = v26
    default:
        v23 = v23
    }
    switch "" {
    case "a":
        v23 -= (v23 >> ((uint(v22)) & 63))
    case "\n\t":
        v22 -= (v22 - genFn1(255, v22, -9223372036854775808))
    }
    var v27 uint16 = 1
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(v6)
        _ = fi28
    }
    var v29 int8 = 100
    var v30 string = "hello"
    if true {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 int16 = 1
            _ = fi31
            _ = v32
        }
        var v33 float32 = 0.0
        fmt.Println(v29, v33)
        _ = v33
    } else {
        v23 -= (v23 & (-(v23)))
        var v34 int64 = -421
        v27 += 42
        _ = v34
    }
    v35 := GenS1{F0: 911}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
