package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 int; F2 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int, p1 uint8) int8 {
    v6 := GenS1{F0: 2, F1: 42}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int64) uint32 {
    v8 := []int{-896}
    switch ("hello" + "\n\t") {
    case "\n\t":
        p0 = (p0 && p0)
    case "foo bar":
        var v9 float32 = 0.0
        _ = v9
    default:
        v10 := GenS1{F0: 2, F1: 4294967295}
        _ = v10
    }
    var v11 uint = 1
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 int = 962
    v5 := []string{"\n\t", ("foo bar" + "foo bar")}
    if true {
        var v12 int16 = 42
        v13 := GenS1{F0: 809, F1: 42}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            if ((fi14 + (v4 ^ 0)) <= fi14) {
                v4 -= 108
                v15 := GenS2{F0: (fi14 >> ((0) & 63)), F1: (v4 * v4), F2: true}
                _ = v15
            } else {
                v16 := GenS2{F0: 2147483647, F1: (fi14 << ((1) & 63)), F2: false}
                switch fi14 {
                case -86:
                    var v17 uint16 = 2
                    _ = v17
                case 4294967295:
                    if (!(4294967295 >= 2147483647)) {
                        var v18 int64 = 9223372036854775807
                        v18 += v18
                        _ = v18
                    } else {
                        v19 := GenS2{F0: -1, F1: (fi14 - -9223372036854775808), F2: ((!(!(-128 <= 100))) && false)}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v20 := GenS1{F0: 18446744073709551615, F1: 4294967295}
                        _ = v19
                        _ = v20
                    }
                default:
                    fmt.Println(v4)
                }
                v3[-1] = 168
                _ = v16
            }
            var v21 int8 = 29
            v21 = v21
            _ = fi14
            _ = v21
        }
        var v22 int16 = -1
        _ = v22
    }
    v4 = (v4 + v4)
    switch "\n\t" {
    case "":
        v5 = append(v5, "foo bar")
    case "a":
        switch "a" {
        case "foo bar":
            fmt.Println(v4)
        case "\n\t":
            var v23 int8 = -74
            _ = v23
        }
    case "hello":
        for fi24 := 0; fi24 < 6; fi24++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi24
        }
    }
    v2[(v4 >> ((42) & 63))] = false
    v4 += v4
    var v25 float32 = (float32((float64(-2.25) + (0.0 + 3.14159))) + 3.14159)
    v26 := GenS2{F0: v4, F1: v4, F2: false}
    if (42 < 0) {
        var v27 int = (^(v4))
        v28 := GenS3{F0: 2147483647, F1: 2147483648}
        if (true || (2 <= 2)) {
            v25 -= ((v25 - -2.25) + 0.1)
            switch 9223372036854775807 {
            case 127:
                v29 := GenS4{F0: "", F1: -128}
                _ = v29
            default:
                v25 = v25
            }
        } else {
            fmt.Println(v27)
            v5 = append(v5, "hello")
        }
        _ = v27
        _ = v28
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v30 := GenS5{F0: int8(v4)}
        v4 = 928
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
