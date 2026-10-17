package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int64; F2 int8 }
func genFn0() int8 {
    if (116 != 42) {
        var v7 int8 = 0
        v7 = v7
        _ = v7
    } else {
        var v8 int64 = -9223372036854775808
        v9 := []int{2147483647, 250, 127}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: true, F2: 0}
    switch (("foo bar" + "hello") + (("\n\t" + "foo bar") + "hello")) {
    case "foo bar":
        if (!(!(!(true && true)))) {
            v11 := GenS2{F0: 127, F1: 2}
            _ = v11
        } else {
            v12 := GenS3{F0: 32767}
            if ((!(true || ("\n\t" != ""))) || (!(2147483647 != 4294967295))) {
                if false {
                    if (153 == 128) {
                        v13 := GenS2{F0: 2, F1: 4294967295}
                        var v14 float32 = (((-2.25 + 0.0) - 3.14159) - (1e10 * -2.25))
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint64 = 0
                        _ = v15
                    }
                }
            } else {
                var v16 int16 = -32768
                _ = v16
            }
            var v17 int16 = 1
            _ = v12
            _ = v17
        }
    case "hello":
        var v18 uint = 0
        _ = v18
    case "":
        var v19 int8 = 0
        _ = v19
    default:
        switch -9223372036854775808 {
        case -2147483648:
            var v20 string = "a"
            _ = v20
        case -1:
            v21 := []int{-16, 0, 127}
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int16 = -32768
        v23 = v23
        _ = fi22
        _ = v23
    }
    v24 := GenS4{F0: 699, F1: 42, F2: 100}
    _ = v10
    _ = v24
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{0, -1}
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    var v6 int32 = 2147483647
    if ((128 != 124) && false) {
        v6 += v6
    }
    v25 := v5[4294967295]
    var v26 int8 = 0
    if v25 {
        var v27 bool = v25
        v25 = false
        fmt.Println(v6, v6, v6)
        _ = v27
    } else {
        v6 += (-(v6))
        var v28 string = ("a" + "foo bar")
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 float32 = 1e10
        _ = fi29
        _ = v30
    }
    var v31 uint8 = 2
    v26 = (v26 / ((v26) | 1))
    var v32 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
