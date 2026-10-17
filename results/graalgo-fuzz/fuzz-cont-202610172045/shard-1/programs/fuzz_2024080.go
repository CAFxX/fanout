package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint8 = 64
    var v2 float64 = ((float64(float32(0.0)) + 1e10) / (0.1 * (3.14159 - (1e10 / float64((0.1 + (0.0 * float32(0.1))))))))
    v3 := make(map[string]uint32)
    var v4 int64 = 1
    v5 := []int{56}
    v4 = 645
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    if false {
        if (v2 > v2) {
            if len(v5) > 0 {
                v5[0] = int(v1)
            }
            var v6 uint16 = 951
            switch "\n\t" {
            case "\n\t":
                v7 := v3["a"]
                _ = v7
            case "hello":
                v1 = (v1 & v1)
            }
            _ = v6
        } else {
            v8 := GenS1{F0: "hello", F1: 0}
            var v9 uint32 = 2147483648
            _ = v8
            _ = v9
        }
        var v10 uint16 = 32767
        var v11 float64 = v2
        _ = v10
        _ = v11
    } else {
        if true {
            switch 659 {
            case 4294967295:
                fmt.Println(v1, v4)
            case 0:
                v12 := v3["\n\t"]
                _ = v12
            default:
                v13 := GenS2{F0: 2147483647, F1: v1, F2: 255}
                _ = v13
            }
            var v14 int64 = v4
            _ = v14
        } else {
            v15 := GenS3{F0: (0 > 1), F1: 0, F2: ("foo bar" > ("a" + "a"))}
            _ = v15
        }
        var v16 uint8 = v1
        if true {
            if false {
                fmt.Println(v1)
            }
            if true {
                switch "a" {
                case "a":
                    fmt.Println(v2)
                case "":
                    v4 = v4
                }
                if (0 == 0) {
                    var v17 float64 = float64(-2.25)
                    var v18 bool = true
                    var v19 float64 = float64(1.5)
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    var v20 bool = (false && false)
                    v20 = v20
                    _ = v20
                }
                var v21 int8 = 1
                _ = v21
            } else {
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v2)
                    switch "hello" {
                    case "hello":
                        if (uint(v16) == 42) {
                            if false {
                                v16 = v16
                                fmt.Println(fi22)
                            } else {
                                var v23 int16 = 32767
                                v16 += (^(v1))
                                if (42 != 4294967295) {
                                    var v24 int8 = 100
                                    fmt.Println(v4, v1, v4)
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        v5 = append(v5, (-(fi22)))
                                        fmt.Println(v16, v4, v1)
                                        _ = fi25
                                    }
                                    v26 := GenS2{F0: 419, F1: (v1 / (((v1 << ((788) & 63))) | 1)), F2: 65535}
                                    _ = v26
                                }
                                _ = v23
                            }
                        } else {
                            for fi27 := 0; fi27 < 11; fi27++ {
                                if false {
                                    v2 += v2
                                    _ = len(v3)
                                }
                                if (!true) {
                                    v28 := GenS4{F0: 0}
                                    v5 = append(v5, (fi27 & (fi27 ^ fi22)))
                                    if true {
                                        v16 = (v1 | (v16 >> ((801) & 63)))
                                    }
                                    _ = v28
                                }
                                v1 -= v16
                                _ = fi27
                            }
                        }
                    case "\n\t":
                        v4 = (-(v4))
                    default:
                        v2 -= v2
                    }
                    v1 += (^(v1))
                    _ = fi22
                }
            }
        } else {
            v3[(("foo bar" + ("foo bar" + "a")) + "a")] = 800
            v1 -= (v16 / ((v16) | 1))
        }
        _ = v16
    }
    var v29 uint64 = 855
    v4 += (v4 / (((v4 << ((2) & 63))) | 1))
    var v30 int = -855
    v31 := GenS1{F0: "foo bar", F1: 1}
    var v32 int32 = -1
    _ = len(v5)
    v29 = v29
    v33 := GenS5{F0: 42, F1: -128}
    var v34 float64 = (3.14159 + (v2 + v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
