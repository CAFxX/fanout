package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 int16, p2 int16) int32 {
    p0 = 42
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint32 = 168
        for fi6 := 0; fi6 < 1; fi6++ {
            if false {
                var v7 string = ""
                fmt.Println(v5, v7)
                v8 := GenS1{F0: 498, F1: 90}
                _ = v7
                _ = v8
            } else {
                switch (p0 | -9223372036854775808) {
                case 252:
                    v9 := GenS1{F0: 579, F1: 1}
                    _ = v9
                default:
                    v10 := []int{634}
                    _ = v10
                }
                switch "a" {
                case "foo bar":
                    v5 = v5
                case "a":
                    var v11 uint8 = 1
                    _ = v11
                default:
                    p2 = p2
                }
            }
            v12 := []int{}
            _ = fi6
            _ = v12
        }
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 769
}
func genFn1(p0 int8, p1 bool, p2 uint8) uint16 {
    v13 := []int{-157}
    v14 := GenS1{F0: -9223372036854775808, F1: p0}
    var v15 int = -61
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 32767
}

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 0
    var v3 int16 = -511
    switch "foo bar" {
    case "hello":
        v16 := []int{42, -185, -9223372036854775808, -1}
        _ = v16
    case "\n\t":
        if true {
            var v17 float32 = (v1 + v1)
            v17 = v1
            _ = v17
        }
    default:
        var v18 string = ("foo bar" + "\n\t")
        _ = v18
    }
    fmt.Println(v1, v3)
    v2 = (v2 ^ v2)
    switch "hello" {
    case "hello":
        v19 := []int{32767, 9223372036854775807, -77}
        _ = v19
    case "\n\t":
        if (!(!(true && (65535 > 42)))) {
            var v20 int64 = 1
            v1 = -2.25
            v21 := []int{957, 4294967295, 710, 127}
            _ = v20
            _ = v21
        } else {
            if false {
                for fi22 := 0; fi22 < 10; fi22++ {
                    if false {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v1 = 0.1
                            v3 -= (^(v3))
                            v1 -= 0.0
                            if true {
                                switch ("\n\t" + (("a" + "a") + "a")) {
                                case "a":
                                    v1 = v1
                                    v24 := []int{}
                                    _ = v24
                                }
                            }
                            _ = fi23
                        }
                        v25 := []int{}
                        var v26 uint64 = 42
                        _ = v25
                        _ = v26
                    }
                    v27 := []int{(fi22 / ((1) | 1))}
                    _ = fi22
                    _ = v27
                }
                if false {
                    v3 = (v3 * int16(v2))
                    v3 -= 0
                }
                switch -140 {
                case -2147483648:
                    v28 := GenS2{F0: 4294967295}
                    _ = v28
                case -9223372036854775808:
                    v2 = -641
                }
            } else {
                switch "" {
                case "foo bar":
                    for fi29 := 0; fi29 < 12; fi29++ {
                        v30 := []int{0}
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v32 := GenS1{F0: -9223372036854775808, F1: -1}
                            var v33 float32 = ((1.5 - ((v1 / v1) - 3.14159)) / 3.14159)
                            v30 = append(v30, int(v3))
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = fi29
                        _ = v30
                    }
                default:
                    var v34 int8 = 19
                    _ = v34
                }
            }
        }
    }
    v35 := GenS2{F0: 28}
    v36 := []int{2147483647, -2147483648, 2147483647, 32767}
    if true {
        v37 := GenS1{F0: -9223372036854775808, F1: -101}
        fmt.Println(v1)
        var v38 int64 = 0
        _ = v37
        _ = v38
    }
    v3 += (v3 | (-(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
