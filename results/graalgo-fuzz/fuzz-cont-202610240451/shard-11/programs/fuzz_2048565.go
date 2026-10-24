package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0() uint16 {
    if true {
        if true {
            var v7 string = "hello"
            switch "a" {
            case "hello":
                v8 := []int{}
                _ = v8
            case "foo bar":
                v9 := []int{-946, -9223372036854775808, 32767}
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := []int{(fi10 - fi10), (fi10 - (fi10 | (-(fi10))))}
                    var v12 int16 = -1
                    v7 = (v7 + ((v7 + v7) + "foo bar"))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            _ = v7
        } else {
            var v13 uint64 = 42
            v14 := GenS1{F0: 1, F1: 42, F2: -1}
            v13 = v13
            _ = v13
            _ = v14
        }
    }
    v15 := []int{285, -21, -2147483648}
    if false {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: 32767, F1: 455, F2: 0}
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: 0, F1: 2, F2: 154}
        _ = v18
    } else {
        var v19 int8 = -128
        _ = v19
    }
    v20 := GenS2{F0: 0, F1: 417, F2: 784}
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 int = -9223372036854775808
            _ = fi22
            _ = v23
        }
        switch (fi21 - (fi21 >> ((1) & 63))) {
        case 4294967295:
            v24 := GenS3{F0: 1, F1: 42}
            _ = v24
        case -2147483648:
            var v25 uint = 42
            _ = v25
        default:
            var v26 uint8 = 86
            _ = v26
        }
        _ = fi21
    }
    _ = v15
    _ = v20
    return 42
}
func genFn1(p0 int, p1 uint8, p2 int64) uint {
    v27 := GenS4{F0: -1, F1: p1}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 4294967295
}

func main() {
    var v1 int32 = 42
    v2 := []uint32{566, 65535, 2147483648, 2147483647}
    v3 := []string{}
    var v4 uint8 = 255
    v5 := []uint32{0, 394, 4294967295, 2}
    var v6 int64 = 0
    if false {
        v4 = v4
    }
    v28 := GenS5{F0: -267}
    v4 += v4
    var v29 uint64 = 42
    for fi30 := 0; fi30 < 1; fi30++ {
        v29 = 0
        _ = fi30
    }
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v31 int16 = 1000
    var v32 int = int(v31)
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS6{F0: 65535, F1: v4}
        switch 42 {
        case -12:
            v31 = v31
        case 1:
            v4 = v4
        case 128:
            v35 := GenS6{F0: 2, F1: 125}
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    var v36 uint8 = (v4 & (v4 | v4))
    fmt.Println(v1, v31)
    v29 -= 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
