package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 uint, p1 uint8, p2 uint64) bool {
    fmt.Println(p1)
    if false {
        fmt.Println(p1)
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{(fi6 / (((^(fi6))) | 1))}
            p2 = p2
            fmt.Println(p2, p1)
            _ = fi6
            _ = v7
        }
    } else {
        var v8 int16 = -32768
        if (0 > 1) {
            v9 := GenS1{F0: (p1 | (p1 / ((35) | 1))), F1: uint64(p0), F2: v8}
            var v10 int16 = v8
            v11 := []int{-888, int(p0), -906}
            _ = v10
            _ = v11
            _ = v9
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                p0 += p0
                var v13 uint32 = 4294967295
                if true {
                    v13 += v13
                }
                _ = fi12
                _ = v13
            }
        }
        _ = v8
    }
    v14 := GenS1{F0: p1, F1: uint64(p0), F2: -114}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (true && false)
}
func genFn1() uint32 {
    var v15 int8 = -92
    v16 := []int{9223372036854775807, 650, -2147483648}
    fmt.Println(v15)
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 113
    v3 := make(map[uint32]uint32)
    var v4 float32 = (0.0 * ((0.1 / float32((0.1 * 3.14159))) + 1.5))
    v5 := make(map[int]int)
    v3[2147483648] = 0
    v1 -= v1
    var v17 float32 = (3.14159 + 0.1)
    var v18 string = "hello"
    switch -2147483648 {
    case -641:
        fmt.Println(v1, v1)
    }
    if (767 > 620) {
        v19 := GenS2{F0: 822}
        fmt.Println(v1, v17, v17)
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS3{F0: -32768, F1: -190, F2: (v2 | (v2 & (v2 + (v2 / ((42) | 1)))))}
            var v22 int8 = 127
            var v23 int32 = 42
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = v19
    } else {
        switch (v18 + ("hello" + (((v18 + (v18 + v18)) + (("hello" + v18) + v18)) + v18))) {
        case "":
            v24 := []int{-1, -50}
            _ = v24
        case "hello":
            var v25 int = 42
            _ = v25
        }
    }
    v2 = v2
    fmt.Println(v18)
    v26 := GenS1{F0: 127, F1: (v1 + (v1 * (v1 + (v1 >> ((331) & 63))))), F2: 1}
    v17 -= 0.0
    var v27 uint64 = 2
    v17 += 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
