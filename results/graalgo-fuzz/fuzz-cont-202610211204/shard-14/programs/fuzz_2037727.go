package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{fi4}
        _ = fi4
        _ = v5
    }
    switch 42 {
    case 4294967295:
        v6 := []int{-394, 0, -2147483648}
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: 32767}
    v8 := []int{158, 4294967295}
    _ = p0
    _ = v7
    _ = v8
    return 56
}
func genFn1(p0 string, p1 uint, p2 uint16) uint {
    var v9 int = 668
    fmt.Println(p1, p0)
    var v10 uint32 = 2147483648
    var v11 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (p1 * 2)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    var v3 uint16 = 702
    delete(v1, 32768)
    fmt.Println(v3, v3)
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS2{F0: -128, F1: 2147483648}
        _ = fi12
        _ = v13
    }
    var v14 int64 = -9223372036854775808
    v14 -= v14
    v15 := GenS3{F0: 51}
    if true {
        var v16 uint64 = 1
        if true {
            var v17 int16 = -32768
            var v18 uint32 = 94
            v3 = 65535
            _ = v17
            _ = v18
        } else {
            var v19 uint64 = v16
            _ = len(v1)
            _ = v19
        }
        _ = v16
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v14, v2, v3)
        v14 = v14
        _ = fi20
    }
    var v21 uint = 167
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int64 = 1
        v23 += 9223372036854775807
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
