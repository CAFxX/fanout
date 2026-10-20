package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint16 }
func genFn0(p0 int8, p1 uint) uint32 {
    var v6 uint8 = 26
    v7 := GenS1{F0: 0, F1: 315}
    fmt.Println(p1, p0, p1)
    switch (("" + "hello") + "\n\t") {
    case "":
        v8 := GenS2{F0: p1, F1: 42}
        _ = v8
    case "\n\t":
        v9 := []int{32768, -9223372036854775808, int(v6), 4294967295}
        _ = v9
    case "foo bar":
        var v10 int16 = 32767
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            v6 += 128
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 uint32 = 1
                p0 = (p0 - 1)
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint64 = 18446744073709551615
                v16 := []int{fi14, fi11, -2147483648}
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 475
}
func genFn1() uint64 {
    var v17 uint = 255
    var v18 int64 = -1
    _ = v17
    _ = v18
    return 845
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 int64 = 0
    v5 := make(map[string]bool)
    var v19 string = ("hello" + (("a" + "\n\t") + "a"))
    v1 = 0.0
    if (!true) {
        var v20 int8 = 25
        if false {
            v21 := GenS3{F0: 113, F1: (465 == 0), F2: v20}
            _ = v21
        }
        fmt.Println(v19, v4)
        _ = v20
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v19 = v19
        _ = fi22
    }
    var v23 bool = true
    if (!v23) {
        v24 := GenS3{F0: 32767, F1: v23, F2: 0}
        v25 := []int{-2147483648, 484, -1}
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS4{F0: -46, F1: 238, F2: 255}
            _ = fi26
            _ = v27
        }
        _ = v24
        _ = v25
    } else {
        var v28 string = ("" + v19)
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int32 = 691
            var v31 uint8 = 255
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    var v32 int = 201
    v23 = (247 > 845)
    switch v19 {
    case "":
        fmt.Println(v23, v32, v19)
    }
    var v33 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
