package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 uint, p2 int64) uint16 {
    p1 = p1
    var v5 int32 = -106
    v6 := GenS1{F0: 38, F1: (p0 % ((p0) | 1)), F2: 255}
    var v7 int64 = (p2 - -875)
    var v8 int64 = 73
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int8, p1 uint64, p2 int) int64 {
    if (1 < uint8(p1)) {
        p2 -= p2
        var v9 bool = false
        _ = v9
    }
    p1 -= (p1 / (((p1 ^ p1)) | 1))
    if false {
        fmt.Println(p1)
    }
    fmt.Println(p1)
    if true {
        switch p2 {
        case 128:
            var v10 int32 = 780
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    var v3 uint = 702
    var v4 uint64 = 99
    switch "a" {
    case "\n\t":
        var v11 uint64 = v4
        _ = v11
    }
    v12 := []int{32768, 9223372036854775807, -9223372036854775808}
    fmt.Println(v4)
    switch (("foo bar" + "foo bar") + ("" + "hello")) {
    case "\n\t":
        v3 = (v3 & 1)
    default:
        v13 := GenS2{F0: -977, F1: 9223372036854775807, F2: 42}
        _ = v13
    }
    if (-2147483648 == 0) {
        v4 = (v4 | (^(v4)))
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 = v4
            _ = fi14
        }
    } else {
        v15 := GenS3{F0: 304, F1: "a"}
        v16 := GenS3{F0: 42, F1: "a"}
        _ = v15
        _ = v16
    }
    if false {
        switch 127 {
        case -96:
            v2[-2147483648] = 1
        default:
            v17 := GenS3{F0: 127, F1: ("a" + "")}
            _ = v17
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 bool = (-463 > 1)
            if len(v12) > 0 {
                v12[0] = 1
            }
            if true {
                v20 := GenS4{F0: 4294967295, F1: 637}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 uint8 = 128
                    fmt.Println(fi18, fi21)
                    v1 = 42
                    switch -414 {
                    case -464:
                        var v23 int8 = 1
                        _ = v23
                    case 312:
                        switch (^(fi18)) {
                        case 1:
                            v1 -= (v1 | v22)
                        }
                    }
                    v24 := GenS5{F0: 111, F1: 732}
                    _ = fi21
                    _ = v22
                    _ = v24
                }
            }
            _ = fi18
            _ = v19
        }
    } else {
        if false {
            var v25 int64 = 0
            v4 = (v4 << ((v3) & 63))
            _ = v25
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 float32 = 1.5
                _ = fi26
                _ = v27
            }
        }
        fmt.Println(v3, v4, v3)
        v28 := v2[32767]
        _ = v28
    }
    v3 += (^(v3))
    var v29 int = -58
    _ = len(v12)
    v4 += uint64(v1)
    var v30 uint16 = 0
    fmt.Println(v30, v29, v30)
    fmt.Println(v4)
    for fi31 := 0; fi31 < 12; fi31++ {
        fmt.Println(v3)
        _ = fi31
    }
    var v32 int64 = 188
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
