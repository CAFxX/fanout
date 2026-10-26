package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string, p1 int64) int16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 983, F1: 42}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 2}
    p0 = (p0 + (("\n\t" + (p0 + "a")) + ("" + "")))
    switch 128 {
    case 899:
        fmt.Println(p1, p0)
    }
    v10 := GenS1{F0: 42, F1: 65535}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -1
}
func genFn1(p0 uint32, p1 uint8, p2 int16) int64 {
    p2 += int16(p1)
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS2{F0: 0}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 895
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    var v4 uint16 = 32767
    var v5 int8 = 100
    var v6 int = -9223372036854775808
    v13 := GenS3{F0: 148, F1: (v6 <= v6)}
    v1[956] = 2147483648
    fmt.Println(v4, v4, v4)
    v4 = (v4 % ((v4) | 1))
    if false {
        v14 := v3[(^(v6))]
        v15 := GenS4{F0: 65535, F1: 501}
        _ = v14
        _ = v15
    }
    _ = len(v3)
    v5 += v5
    v16 := GenS5{F0: 201}
    var v17 uint = 65535
    switch -178 {
    case 489:
        for fi18 := 0; fi18 < 5; fi18++ {
            v6 = fi18
            var v19 bool = false
            _ = fi18
            _ = v19
        }
    case 0:
        var v20 int16 = 1
        _ = v20
    case 9223372036854775807:
        v21 := GenS5{F0: 197}
        _ = v21
    }
    var v22 int32 = -2147483648
    var v23 uint16 = (^(v4))
    v24 := []int{(v6 + -1), v6}
    var v25 uint16 = (^(v4))
    var v26 int64 = 1
    if (false && true) {
        v6 = 42
        var v27 uint = v17
        v28 := GenS2{F0: 4294967295}
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
