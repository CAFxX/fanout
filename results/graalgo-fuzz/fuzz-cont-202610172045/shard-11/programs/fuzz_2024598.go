package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
func genFn0() uint {
    var v6 float32 = (0.1 * (1e10 + 1e10))
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = 3.14159
        v8 := []int{}
        if true {
            v9 := GenS1{F0: 2, F1: 9223372036854775807, F2: 2}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    switch 42 {
    case 4294967295:
        var v10 int = -2147483648
        _ = v10
    }
    var v11 int16 = -32768
    v12 := GenS1{F0: 1, F1: 0, F2: 2}
    _ = v11
    _ = v12
    _ = v6
    return 947
}
func genFn1(p0 int, p1 uint, p2 int8) uint8 {
    var v13 uint32 = 2
    fmt.Println(p2)
    if false {
        var v14 uint32 = 829
        _ = v14
    } else {
        v15 := []int{p0}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = (3.14159 + -2.25)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 uint64 = 556
    switch "\n\t" {
    case "":
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS2{F0: 854, F1: "\n\t"}
            _ = fi16
            _ = v17
        }
    case "a":
        v18 := GenS3{F0: 0, F1: 65535, F2: 4294967295}
        _ = v18
    }
    v19 := GenS1{F0: 18446744073709551615, F1: 0, F2: 42}
    var v20 int64 = 1
    v21 := []int{127}
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint8 = uint8(v1)
        var v24 float64 = 3.14159
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint = 255
    var v26 uint64 = uint64(v25)
    var v27 uint32 = v1
    var v28 uint8 = genFn1(2147483647, v25, int8(v20))
    var v29 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
