package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 string; F1 int8; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 uint) uint16 {
    fmt.Println(p0)
    var v4 int16 = 42
    _ = p0
    _ = v4
    return 990
}
func genFn1(p0 uint, p1 string, p2 int16) int64 {
    var v5 int8 = -115
    if true {
        switch (p1 + p1) {
        case "a":
            v6 := []int{-9223372036854775808}
            _ = v6
        case "hello":
            v7 := []int{4294967295}
            _ = v7
        default:
            p2 = 42
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 741
}

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 1
    v3 := make(map[uint32]bool)
    v8 := GenS1{F0: 127}
    v9 := []int{-604, 9223372036854775807, 4294967295}
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int8 = int8(v1)
        _ = fi10
        _ = v11
    }
    _ = len(v3)
    v12 := GenS1{F0: 52}
    var v13 string = ("hello" + (("foo bar" + ("a" + "a")) + ""))
    v14 := GenS1{F0: 127}
    _ = len(v9)
    var v15 int64 = int64(v1)
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v13 = "a"
            var v18 int16 = 1000
            if false {
                switch (fi16 + 2147483647) {
                case -1:
                    if (fi16 > int(v2)) {
                        var v19 uint16 = 852
                        var v20 uint8 = (v1 + (^(v1)))
                        v19 = genFn0(255)
                        _ = v19
                        _ = v20
                    } else {
                        v1 -= v1
                        v1 += v1
                    }
                case 127:
                    switch "foo bar" {
                    case "hello":
                        if (false || (1.5 <= -2.25)) {
                            if false {
                                v21 := GenS2{F0: true, F1: 42, F2: (^(v1))}
                                if true {
                                    fmt.Println(v18)
                                }
                                _ = v21
                            }
                        } else {
                            fmt.Println(v1, v1)
                            var v22 int16 = (v18 + (^(v18)))
                            v1 = v1
                            _ = v22
                        }
                        var v23 int32 = -17
                        _ = v23
                    case "a":
                        var v24 uint8 = 2
                        _ = v24
                    case "foo bar":
                        var v25 float64 = 1.5
                        _ = v25
                    default:
                        var v26 float64 = (1e10 + 3.14159)
                        _ = v26
                    }
                }
                if false {
                    fmt.Println(v15)
                    switch (fi17 | fi17) {
                    case 1:
                        v3[v2] = true
                    case 32768:
                        fmt.Println(v18)
                    default:
                        v27 := v3[(^(v2))]
                        _ = v27
                    }
                } else {
                    var v28 string = ("\n\t" + v13)
                    v29 := GenS3{F0: 82, F1: -1}
                    _ = v28
                    _ = v29
                }
                v30 := GenS4{F0: (v13 + (v13 + (("a" + "a") + v13))), F1: -13, F2: 2}
                _ = v30
            } else {
                switch (fi17 & fi17) {
                case -1:
                    v13 = v13
                default:
                    switch (fi17 << ((778) & 63)) {
                    case 290:
                        v31 := GenS5{F0: v15}
                        _ = v31
                    case -9223372036854775808:
                        switch (("\n\t" + "foo bar") + (v13 + "a")) {
                        case "":
                            v32 := GenS6{F0: 42, F1: v2, F2: 42}
                            _ = v32
                        case "a":
                            v1 = v1
                        default:
                            v3[(v2 | (v2 / ((v2) | 1)))] = ((v1 >= v1) || (42 > 4294967295))
                            fmt.Println(v2)
                            var v33 int = fi17
                            v9 = append(v9, (-(v33)))
                            _ = v33
                        }
                    }
                }
                fmt.Println(fi17)
                v34 := GenS5{F0: (v15 ^ (v15 * (v15 | (v15 + v15))))}
                _ = v34
            }
            v18 += (v18 * 0)
            fmt.Println(v1, v15, fi17)
            _ = fi17
            _ = v18
        }
        var v35 bool = true
        _ = fi16
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS6{F0: 0, F1: 2147483648, F2: 4294967295}
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 11; fi38++ {
        var v39 int64 = genFn1(540, (v13 + v13), -182)
        var v40 int = (fi38 & (-(fi38)))
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
