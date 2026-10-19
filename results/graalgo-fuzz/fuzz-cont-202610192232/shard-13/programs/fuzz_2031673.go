package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 int64 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 int16, p1 uint, p2 uint8) string {
    for fi6 := 0; fi6 < 11; fi6++ {
        p2 = 255
        p0 -= p0
        v7 := GenS1{F0: 42, F1: int(p2)}
        _ = fi6
        _ = v7
    }
    fmt.Println(p2, p1)
    var v8 uint = (p1 ^ (^(p1)))
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(v8)
        v10 := []int{(fi9 ^ (-(fi9))), (fi9 * 127)}
        _ = fi9
        _ = v10
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "a"
}
func genFn1(p0 int16) bool {
    v11 := GenS2{F0: 1}
    if false {
        p0 -= (p0 | p0)
        var v12 uint = 65535
        var v13 int64 = 42
        _ = v12
        _ = v13
    } else {
        switch ((("\n\t" + ("a" + ("foo bar" + "a"))) + ("\n\t" + ("foo bar" + ("foo bar" + "")))) + "") {
        case "\n\t":
            var v14 int = -2147483648
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 9; fi15++ {
                if false {
                    v16 := []int{127, fi15}
                    if false {
                        if len(v16) > 0 {
                            v16[0] = fi15
                        }
                        fmt.Println(fi15)
                    } else {
                        var v17 int8 = 42
                        p0 -= p0
                        _ = v17
                    }
                    _ = v16
                }
                v18 := []int{}
                _ = fi15
                _ = v18
            }
        case "hello":
            if ((!(false && false)) || false) {
                var v19 int8 = 1
                v20 := GenS3{F0: ""}
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS1{F0: 42, F1: 23}
            _ = v21
        }
        p0 -= p0
        v22 := GenS4{F0: (2 <= 4294967295), F1: 42, F2: -9223372036854775808}
        _ = v22
    }
    v23 := GenS2{F0: 0}
    fmt.Println(p0)
    _ = p0
    _ = v11
    _ = v23
    return (((true && true) || false) || false)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-2147483648, 2147483647, 2147483647}
    var v3 uint64 = 0
    var v4 uint8 = 255
    var v5 int8 = 70
    for fi24 := 0; fi24 < 14; fi24++ {
        if (32767 < 2) {
            _ = len(v2)
        }
        _ = fi24
    }
    var v25 uint32 = 297
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS5{F0: 1000, F1: (v4 << ((uint(v25)) & 63)), F2: v3}
        v28 := GenS3{F0: "foo bar"}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int32 = 0
    if false {
        var v30 int = int(v4)
        _ = len(v1)
        _ = v30
    }
    fmt.Println(v29)
    var v31 float32 = 0.0
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 8; fi33++ {
            _ = len(v1)
            v3 = 1
            _ = fi33
        }
        _ = fi32
    }
    v34 := GenS1{F0: -11, F1: -918}
    if (1 != v25) {
        var v35 float32 = (3.14159 - v31)
        v36 := GenS3{F0: "hello"}
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
