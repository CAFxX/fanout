package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
func genFn0(p0 uint64, p1 uint32) int16 {
    if true {
        var v7 int64 = -564
        for fi8 := 0; fi8 < 5; fi8++ {
            switch "foo bar" {
            case "\n\t":
                if true {
                    v9 := []int{572, -2147483648, int(p1)}
                    fmt.Println(p1, fi8)
                    var v10 uint64 = p0
                    _ = v10
                    _ = v9
                } else {
                    v7 -= v7
                    switch (^(fi8)) {
                    case 0:
                        v7 = (v7 - v7)
                    default:
                        v11 := GenS1{F0: 1, F1: 572, F2: 117}
                        _ = v11
                    }
                }
            case "hello":
                var v12 uint64 = p0
                _ = v12
            }
            _ = fi8
        }
        p1 = (p1 ^ 65535)
        _ = v7
    }
    switch 489 {
    case -95:
        v13 := GenS2{F0: 154, F1: false, F2: "\n\t"}
        _ = v13
    }
    p0 += 0
    p1 = (p1 >> ((390) & 63))
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint16, p1 int8) uint32 {
    v14 := []int{-2147483648, 2147483647, -9223372036854775808, 42}
    p1 = (p1 ^ (p1 / ((p1) | 1)))
    v15 := GenS3{F0: 383}
    if true {
        p1 = (p1 / ((65) | 1))
        var v16 float32 = -2.25
        var v17 uint8 = 255
        _ = v16
        _ = v17
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint64 = 42
    v3 := []uint32{2}
    v4 := make(map[int]bool)
    v5 := []string{"\n\t", ("\n\t" + "\n\t"), ("a" + "\n\t")}
    var v6 uint8 = 2
    v1 = v1
    fmt.Println(v6, v2, v1)
    v2 = v2
    var v18 int8 = 96
    if (32767 >= 255) {
        v19 := GenS2{F0: -157, F1: ("\n\t" == ("hello" + "foo bar")), F2: "foo bar"}
        fmt.Println(v18, v6, v6)
        _ = v19
    } else {
        v20 := GenS4{F0: 42, F1: 42, F2: -1}
        v21 := GenS3{F0: 42}
        v6 = (^(v6))
        _ = v20
        _ = v21
    }
    v5 = append(v5, ("hello" + ("\n\t" + "hello")))
    v1 = v1
    var v22 uint = 493
    v22 = v22
    var v23 float32 = 0.1
    switch "a" {
    case "hello":
        v6 = 0
    case "foo bar":
        v24 := GenS3{F0: uint16(v22)}
        _ = v24
    default:
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 string = "hello"
            v27 := GenS1{F0: -1, F1: v22, F2: uint32(v22)}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        switch fi28 {
        case 0:
            var v29 uint = (^(v22))
            _ = v29
        default:
            v30 := GenS2{F0: 1, F1: v1, F2: ""}
            _ = v30
        }
        v31 := GenS4{F0: -200, F1: -9223372036854775808, F2: -1}
        _ = fi28
        _ = v31
    }
    fmt.Println(v23, v22)
    if (v18 > v18) {
        fmt.Println(v23)
        fmt.Println(v22)
    } else {
        fmt.Println(v2)
    }
    v18 += v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
