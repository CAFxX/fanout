package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
func genFn0(p0 int64, p1 uint32) int32 {
    switch -2147483648 {
    case -657:
        for fi6 := 0; fi6 < 7; fi6++ {
            if (fi6 < fi6) {
                switch "foo bar" {
                case "":
                    fmt.Println(p1, p1, fi6)
                case "hello":
                    var v7 float64 = (0.0 - 3.14159)
                    _ = v7
                }
            } else {
                if (false || false) {
                    if false {
                        if true {
                            v8 := GenS1{F0: 4294967295}
                            fmt.Println(fi6, p0)
                            _ = v8
                        } else {
                            for fi9 := 0; fi9 < 10; fi9++ {
                                v10 := GenS1{F0: 65535}
                                p0 = (p0 >> ((2) & 63))
                                fmt.Println(p1, p1, fi6)
                                var v11 uint = 42
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            for fi12 := 0; fi12 < 19; fi12++ {
                                v13 := GenS1{F0: 2}
                                var v14 float64 = ((-2.25 + 1.5) * 1.5)
                                if (0 != 35) {
                                    switch fi12 {
                                    case -149:
                                        v15 := []int{}
                                        _ = v15
                                    default:
                                        fmt.Println(p1)
                                    }
                                    if false {
                                        if false {
                                            v16 := []int{fi6, (fi6 ^ (fi6 + (^(fi6)))), (fi12 << ((171) & 63)), fi12}
                                            var v17 uint8 = 0
                                            v16 = append(v16, (^(fi12)))
                                            _ = v16
                                            _ = v17
                                        }
                                        fmt.Println(v14, fi6, fi12)
                                        p0 = -9223372036854775808
                                    }
                                    var v18 float64 = v14
                                    fmt.Println(p1, fi6)
                                    v19 := GenS2{F0: -383, F1: 42}
                                    if (!true) {
                                        v20 := GenS3{F0: fi12, F1: -1, F2: 2147483647}
                                        _ = v20
                                    }
                                    v21 := []int{(-(fi6)), (fi6 * -1)}
                                    _ = v18
                                    _ = v19
                                    _ = v21
                                } else {
                                    p1 = p1
                                }
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            p0 = p0
                        }
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v23 := GenS3{F0: 42, F1: 0, F2: 1}
                            p0 += 1
                            fmt.Println(fi6, fi22)
                            _ = fi22
                            _ = v23
                        }
                        var v24 int16 = 32767
                        _ = v24
                    }
                    p1 = (p1 << ((255) & 63))
                }
            }
            v25 := []int{1, (fi6 * -1)}
            _ = fi6
            _ = v25
        }
    case -905:
        p1 = p1
    case 32768:
        v26 := GenS3{F0: 32767, F1: 0, F2: -2147483648}
        _ = v26
    }
    switch 128 {
    case 1:
        var v27 uint = 0
        _ = v27
    }
    v28 := []int{32767, 0, 127, 127}
    if (!(true || true)) {
        var v29 bool = true
        _ = v29
    } else {
        p1 += (p1 % ((403) | 1))
    }
    _ = p0
    _ = p1
    _ = v28
    return -595
}
func genFn1(p0 uint, p1 bool, p2 uint8) uint {
    var v30 int8 = 100
    v30 -= v30
    switch "" {
    case "foo bar":
        fmt.Println(p2)
    case "hello":
        var v31 int64 = 1
        _ = v31
    case "\n\t":
        v32 := []int{127}
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return uint(p2)
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{9223372036854775807}
    var v3 uint64 = 1
    var v4 string = "\n\t"
    var v5 float64 = float64(0.0)
    for fi33 := 0; fi33 < 3; fi33++ {
        v34 := GenS1{F0: 913}
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: 2147483647, F1: 18446744073709551615}
    var v36 bool = (-2.25 > 1.5)
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v1)
    for fi37 := 0; fi37 < 4; fi37++ {
        v2 = append(v2, 128)
        var v38 string = ("" + "foo bar")
        _ = fi37
        _ = v38
    }
    fmt.Println(v5)
    var v39 int = -413
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
