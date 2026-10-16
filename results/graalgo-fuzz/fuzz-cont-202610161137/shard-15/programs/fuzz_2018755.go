package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 int64; F1 int32; F2 int32 }
func genFn0(p0 int, p1 int16) int16 {
    switch ("hello" + (("\n\t" + ("\n\t" + "hello")) + "")) {
    case "a":
        v5 := GenS1{F0: int32(p1)}
        _ = v5
    case "hello":
        v6 := GenS1{F0: -2147483648}
        _ = v6
    }
    if ((-(p1)) == (p1 * 207)) {
        if ((32767 > 387) && true) {
            v7 := GenS2{F0: 775, F1: 1, F2: 2}
            if true {
                var v8 int = p0
                fmt.Println(p1, v8)
                _ = v8
            }
            if ((766 != 32767) && false) {
                for fi9 := 0; fi9 < 14; fi9++ {
                    switch "\n\t" {
                    case "hello":
                        switch ("" + ("" + ("hello" + (("a" + "a") + "\n\t")))) {
                        case "hello":
                            p1 -= p1
                        default:
                            v10 := []int{}
                            _ = v10
                        }
                    case "foo bar":
                        var v11 float64 = ((-2.25 + 0.1) + (3.14159 * 1e10))
                        _ = v11
                    case "\n\t":
                        for fi12 := 0; fi12 < 2; fi12++ {
                            var v13 bool = (!false)
                            var v14 int16 = p1
                            var v15 float64 = 0.1
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    default:
                        v16 := GenS1{F0: 718}
                        _ = v16
                    }
                    _ = fi9
                }
                if true {
                    v17 := GenS2{F0: 4294967295, F1: 69, F2: 255}
                    v18 := GenS2{F0: 791, F1: -1, F2: 65535}
                    p1 = 972
                    _ = v17
                    _ = v18
                }
            }
            _ = v7
        }
        fmt.Println(p0, p0)
    } else {
        var v19 uint32 = 1
        var v20 uint = 1
        if true {
            fmt.Println(p1, v19)
        } else {
            p1 += (-(p1))
        }
        _ = v19
        _ = v20
    }
    var v21 uint64 = 177
    for fi22 := 0; fi22 < 14; fi22++ {
        if false {
            fmt.Println(v21)
            var v23 uint32 = 0
            var v24 uint64 = v21
            _ = v23
            _ = v24
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = v21
    return (p1 << ((42) & 63))
}
func genFn1() uint {
    v25 := []int{1}
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int8 = 127
        _ = fi26
        _ = v27
    }
    var v28 int16 = -509
    _ = v25
    _ = v28
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 224
    v3 := make(map[uint32]bool)
    var v4 uint32 = 4294967295
    for fi29 := 0; fi29 < 1; fi29++ {
        v4 += v4
        _ = fi29
    }
    v4 -= (v4 ^ v4)
    var v30 bool = (18446744073709551615 < 928)
    var v31 uint64 = 18446744073709551615
    if ((true && (v4 > v4)) && ("" < ("a" + ("foo bar" + "foo bar")))) {
        for fi32 := 0; fi32 < 12; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 bool = v30
                v31 -= v31
                _ = fi33
                _ = v34
            }
            _ = fi32
        }
        v3[2147483648] = (406 <= -957)
    }
    v35 := v3[(v4 * (^(v4)))]
    fmt.Println(v35, v30, v30)
    v36 := GenS3{F0: int64(v4), F1: 1, F2: -435}
    v30 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
