package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 int) uint64 {
    v6 := GenS1{F0: 67, F1: 127, F2: (p0 | 469)}
    v7 := []int{(-(p1))}
    var v8 int32 = 1
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 255, F1: 100, F2: p0}
        switch "" {
        case "\n\t":
            var v11 float32 = (((1e10 / -2.25) * 1.5) + 3.14159)
            _ = v11
        case "hello":
            if (0 <= 1000) {
                for fi12 := 0; fi12 < 7; fi12++ {
                    p1 = (fi12 | 4294967295)
                    _ = fi12
                }
                if false {
                    p0 = p0
                }
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(v8, fi13)
                    var v14 uint16 = 42
                    var v15 int16 = 0
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        }
        if true {
            if (v8 == v8) {
                var v16 float32 = ((1.5 / 1e10) + 3.14159)
                _ = v16
            } else {
                if (int16(p0) < 1) {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    p0 += 2147483647
                } else {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        switch "" {
                        case "a":
                            var v18 int = 127
                            _ = v18
                        case "\n\t":
                            v19 := GenS1{F0: 127, F1: 42, F2: 65535}
                            _ = v19
                        case "hello":
                            v7 = append(v7, (fi9 % (((fi17 >> ((42) & 63))) | 1)))
                        default:
                            fmt.Println(fi9)
                        }
                        _ = fi17
                    }
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                p0 -= (p0 % (((p0 / (((p0 ^ (^(p0)))) | 1))) | 1))
                p1 -= (p1 / ((p1) | 1))
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi20
            }
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 591
}
func genFn1(p0 uint) uint {
    p0 -= p0
    p0 -= (p0 / ((255) | 1))
    v21 := GenS2{F0: 42, F1: 42, F2: (p0 & p0)}
    if (!(-52 == 48)) {
        v22 := GenS3{F0: 1, F1: 1}
        _ = v22
    }
    v23 := GenS4{F0: true}
    _ = p0
    _ = v21
    _ = v23
    return p0
}

func main() {
    v1 := []string{("a" + (("\n\t" + (("\n\t" + ("foo bar" + ((("foo bar" + "a") + "") + "foo bar"))) + "hello")) + ((("\n\t" + "hello") + ("a" + "a")) + ("foo bar" + "a"))))}
    v2 := []int{4294967295, -466, 946}
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    delete(v5, 1)
    v3 = v3
    v24 := GenS2{F0: -1, F1: 2, F2: 65535}
    fmt.Println(v3)
    fmt.Println(v3)
    v25 := GenS4{F0: true}
    _ = len(v2)
    var v26 int64 = 302
    for fi27 := 0; fi27 < 3; fi27++ {
        if (66 >= 254) {
            var v28 string = v3
            fmt.Println(fi27)
            _ = v28
        }
        var v29 uint8 = 127
        _ = fi27
        _ = v29
    }
    var v30 uint8 = 255
    for fi31 := 0; fi31 < 9; fi31++ {
        v30 += v30
        v26 = (v26 % ((-655) | 1))
        var v32 int = (fi31 % ((-1) | 1))
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
