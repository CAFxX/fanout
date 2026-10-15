package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 bool, p1 uint8) int64 {
    fmt.Println(p1, p1)
    var v4 uint = 255
    v5 := GenS1{F0: -2147483648}
    if (p0 || true) {
        v6 := GenS2{F0: 32767, F1: 2, F2: -32768}
        _ = v6
    }
    v7 := []int{540, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    return -717
}
func genFn1(p0 int32, p1 int) uint8 {
    switch "foo bar" {
    case "hello":
        p1 = p1
    case "\n\t":
        p1 += 32767
    }
    if false {
        switch p1 {
        case 42:
            if false {
                if true {
                    p1 -= 127
                    v8 := GenS3{F0: genFn0(true, 151), F1: 2}
                    _ = v8
                } else {
                    v9 := []int{p1}
                    p1 += p1
                    _ = v9
                }
            } else {
                v10 := GenS3{F0: -9223372036854775808, F1: 787}
                _ = v10
            }
        }
        var v11 uint = 446
        v11 = 255
        _ = v11
    } else {
        v12 := GenS4{F0: 2147483647, F1: -1}
        p0 -= -1
        _ = v12
    }
    p1 += p1
    v13 := []int{2147483647, p1}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = -9223372036854775808
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 int = 0
    var v16 int64 = 42
    v17 := []int{32768, v3, (v3 - v3)}
    switch "\n\t" {
    case "foo bar":
        var v18 int32 = -122
        _ = v18
    case "":
        if (42 != 127) {
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 int8 = -128
                var v21 uint8 = 193
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 int32 = -2147483648
            _ = v22
        } else {
            var v23 float32 = 1e10
            _ = v23
        }
    case "\n\t":
        switch -446 {
        case 127:
            v16 = (v16 >> ((1) & 63))
        case 42:
            v16 -= genFn0(true, 169)
        case -663:
            switch "a" {
            case "a":
                v24 := v1[370]
                _ = v24
            case "hello":
                v16 = 489
            }
        }
    default:
        switch (v3 | (v3 << ((330) & 63))) {
        case 4294967295:
            v25 := GenS5{F0: 42, F1: 0, F2: 535}
            _ = v25
        case 32767:
            if (!((float64(0.1) + float64(1.5)) <= 1e10)) {
                var v26 float32 = 1e10
                var v27 int = v3
                _ = v26
                _ = v27
            }
        default:
            var v28 float32 = 0.0
            _ = v28
        }
    }
    v16 = 9223372036854775807
    v16 = genFn0(((3.14159 * (((((0.0 / 1.0) / -2.25) + 0.0) + -2.25) + 0.1)) <= (-2.25 / -2.25)), 0)
    var v29 uint = 42
    var v30 uint64 = 979
    for fi31 := 0; fi31 < 1; fi31++ {
        v29 += 663
        v3 = (fi31 << (((v29 << ((v29) & 63))) & 63))
        _ = fi31
    }
    var v32 uint16 = 337
    v33 := GenS5{F0: 100, F1: (v29 % ((1) | 1)), F2: v30}
    var v34 bool = (false && ((4294967295 < 1) || (-1 == -32768)))
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
}
