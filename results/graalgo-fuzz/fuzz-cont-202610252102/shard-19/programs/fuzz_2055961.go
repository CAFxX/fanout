package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 string, p1 int16) uint64 {
    switch p0 {
    case "":
        p1 = p1
    case "foo bar":
        p1 = p1
    case "\n\t":
        var v5 int = 32767
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6, fi6, fi6)
        p0 = p0
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (p0 + "hello") {
        case "hello":
            var v8 int32 = -2147483648
            _ = v8
        case "":
            var v9 uint64 = 254
            _ = v9
        case "foo bar":
            v10 := GenS1{F0: 565, F1: 1, F2: -1}
            _ = v10
        }
        if ((false && true) || true) {
            p0 = p0
            var v11 uint8 = 2
            _ = v11
        } else {
            v12 := []int{(^(fi7)), -1}
            _ = v12
        }
        _ = fi7
    }
    v13 := []int{-9223372036854775808, 127}
    _ = p0
    _ = p1
    _ = v13
    return 957
}
func genFn1(p0 int64, p1 int, p2 string) string {
    v14 := []int{4294967295}
    if len(v14) > 0 {
        v14[0] = p1
    }
    p1 = p1
    _ = len(v14)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p2 + "\n\t")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 127
    var v3 uint16 = uint16(v2)
    var v4 int8 = -15
    if ((599 >= 18446744073709551615) || (-9223372036854775808 >= 42)) {
        v4 += (^(v4))
        v4 -= v4
        switch "a" {
        case "":
            v15 := GenS1{F0: (v3 / ((v3) | 1)), F1: -432, F2: 0}
            _ = v15
        case "hello":
            _ = len(v1)
        case "a":
            v16 := GenS1{F0: v3, F1: 1, F2: 2147483647}
            _ = v16
        default:
            var v17 int = -1
            _ = v17
        }
    } else {
        fmt.Println(v2, v2, v2)
        for fi18 := 0; fi18 < 12; fi18++ {
            v4 = 1
            _ = fi18
        }
        v19 := []int{128, -9223372036854775808, 9223372036854775807, 42}
        _ = v19
    }
    v2 = 0
    v20 := []int{9223372036854775807, 2147483647, 42}
    v3 -= (v3 >> ((1) & 63))
    v21 := GenS1{F0: 0, F1: 897, F2: 2147483647}
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 int = fi22
        for fi24 := 0; fi24 < 2; fi24++ {
            if false {
                v25 := GenS1{F0: 0, F1: 1, F2: -845}
                _ = v25
            }
            v23 = fi22
            if false {
                v26 := GenS2{F0: -1}
                _ = v26
            } else {
                v3 -= 0
                v27 := GenS1{F0: 97, F1: -390, F2: 0}
                fmt.Println(fi22, v23)
                fmt.Println(fi22, v4)
                switch int(v2) {
                case 42:
                    var v28 int32 = -272
                    _ = v28
                default:
                    fmt.Println(v4)
                    var v29 string = "hello"
                    var v30 uint16 = (v3 % ((v3) | 1))
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    v31 := GenS3{F0: false}
    fmt.Println(v4, v4)
    v2 -= (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
