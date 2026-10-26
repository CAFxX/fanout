package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 uint8) int8 {
    v7 := []int{4294967295}
    if (2147483647 < -2147483648) {
        v8 := GenS1{F0: 2}
        p0 = 47
        if len(v7) > 0 {
            v7[0] = 0
        }
        _ = v8
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return -128
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{2, 721, 65535, 65535}
    var v3 string = (("hello" + "a") + (("\n\t" + ((("" + (("\n\t" + "foo bar") + "foo bar")) + "foo bar") + "")) + ("" + "")))
    var v4 uint64 = 4294967295
    v5 := []uint32{65535, 666}
    v6 := make(map[uint32]int)
    var v9 int = 4294967295
    var v10 uint32 = uint32(v4)
    if false {
        var v11 uint64 = 18446744073709551615
        v12 := GenS2{F0: 2147483647, F1: 42}
        _ = v11
        _ = v12
    }
    v3 = "a"
    var v13 uint = 321
    var v14 int8 = 127
    fmt.Println(v1, v14, v10)
    var v15 uint = v13
    v16 := GenS2{F0: (v10 / ((v10) | 1)), F1: 64}
    var v17 float64 = 1.5
    if false {
        var v18 uint8 = 17
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS2{F0: v10, F1: 16}
            v13 = v15
            v21 := GenS3{F0: v4, F1: v4}
            _ = fi19
            _ = v20
            _ = v21
        }
        switch v3 {
        case "":
            fmt.Println(v17, v9)
        case "a":
            for fi22 := 0; fi22 < 9; fi22++ {
                v23 := GenS2{F0: 2147483648, F1: (v18 / ((v18) | 1))}
                _ = fi22
                _ = v23
            }
        }
        _ = v18
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS4{F0: -329, F1: false}
            v26 := GenS1{F0: 4294967295}
            v1 = v1
            _ = fi24
            _ = v25
            _ = v26
        }
        v27 := GenS2{F0: v10, F1: 109}
        _ = v27
    }
    switch v3 {
    case "a":
        v9 -= v9
    case "\n\t":
        v28 := GenS5{F0: false, F1: 128, F2: v10}
        _ = v28
    default:
        fmt.Println(v17, v14)
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        if false {
            var v30 uint = uint(v4)
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v17 = v17
            v10 = 230
            if (-1 > -32768) {
                fmt.Println(v3)
                var v32 float64 = float64(v1)
                var v33 uint32 = (v10 & (v10 << ((v13) & 63)))
                _ = v32
                _ = v33
            }
            _ = fi31
        }
        var v34 bool = (v13 != 18446744073709551615)
        _ = fi29
        _ = v34
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v13 = (v15 - v13)
        var v36 int32 = 2147483647
        _ = fi35
        _ = v36
    }
    var v37 int32 = -1
    for fi38 := 0; fi38 < 11; fi38++ {
        if len(v5) > 0 {
            v5[0] = v10
        }
        _ = fi38
    }
    var v39 float64 = (1e10 * (1e10 / (v17 / v17)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
