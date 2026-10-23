package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 int16) int {
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 = p0
        _ = fi6
    }
    v7 := GenS1{F0: p0, F1: -1}
    if false {
        v8 := []int{0, -289, 42, -134}
        _ = v8
    } else {
        fmt.Println(p0)
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{fi9, int(p0), (-(fi9))}
            _ = fi9
            _ = v10
        }
        var v11 uint = 2
        _ = v11
    }
    v12 := GenS2{F0: "foo bar", F1: -1, F2: 42}
    _ = p0
    _ = v12
    _ = v7
    return -672
}

func main() {
    var v1 uint64 = 306
    var v2 float64 = (0.0 * 0.0)
    var v3 uint32 = 548
    v4 := make(map[uint32]string)
    var v5 int8 = -128
    if true {
        var v13 uint64 = v1
        v14 := GenS3{F0: (-23 == 1)}
        v15 := []int{-9223372036854775808, 2147483647, 334}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v16 := GenS4{F0: (v1 ^ 18446744073709551615), F1: 255}
        _ = v16
    }
    v2 -= 1e10
    fmt.Println(v1, v3, v3)
    v17 := GenS3{F0: (false || ((v5 & v5) >= v5))}
    var v18 int16 = -963
    switch -12 {
    case 1:
        v1 -= v1
    case -221:
        v19 := GenS4{F0: v1, F1: 4294967295}
        _ = v19
    case -619:
        var v20 float32 = 0.0
        _ = v20
    default:
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 float32 = (float32((0.0 * v2)) - 3.14159)
            _ = fi21
            _ = v22
        }
    }
    var v23 int = -1
    fmt.Println(v3)
    if true {
        v1 += v1
        var v24 uint16 = 656
        for fi25 := 0; fi25 < 11; fi25++ {
            delete(v4, (^(v3)))
            if true {
                v26 := GenS5{F0: (v23 & genFn0(32767)), F1: -768}
                _ = v26
            }
            if false {
                var v27 int16 = (v18 ^ (^(v18)))
                _ = v27
            } else {
                for fi28 := 0; fi28 < 3; fi28++ {
                    fmt.Println(v3, fi25, v24)
                    _ = fi28
                }
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int8 = (v5 << ((2) & 63))
                    _ = fi29
                    _ = v30
                }
                var v31 uint16 = (^(v24))
                _ = v31
            }
            _ = fi25
        }
        _ = v24
    } else {
        v32 := GenS4{F0: (^(v1)), F1: 620}
        _ = v32
    }
    v33 := []int{genFn0(v18), v23}
    var v34 int = genFn0(v18)
    v35 := GenS3{F0: true}
    switch "a" {
    case "hello":
        fmt.Println(v23, v1, v3)
    default:
        var v36 uint32 = 2
        _ = v36
    }
    var v37 int64 = 0
    switch "hello" {
    case "a":
        if true {
            for fi38 := 0; fi38 < 8; fi38++ {
                _ = len(v4)
                var v39 int32 = -770
                _ = fi38
                _ = v39
            }
            var v40 uint32 = (^(v3))
            var v41 bool = (4294967295 < 0)
            _ = v40
            _ = v41
        }
    case "":
        if ((^(v3)) <= (v3 + (v3 | (v3 - (^(v3)))))) {
            v23 += genFn0(v18)
        } else {
            v42 := GenS5{F0: -632, F1: int16(v3)}
            v23 += v34
            _ = v42
        }
    default:
        v43 := v4[574]
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
