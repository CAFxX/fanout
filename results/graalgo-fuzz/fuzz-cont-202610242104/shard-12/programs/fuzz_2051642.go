package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int8; F1 int64; F2 int16 }
type GenS6 struct { F0 int }
func genFn0(p0 int, p1 uint64, p2 int64) uint64 {
    v5 := GenS1{F0: true}
    v6 := GenS1{F0: ((!(!(("a" <= "") || true))) && ((uint(p1) <= 837) && ((true || true) && true)))}
    v7 := []int{}
    var v8 uint = 1
    v9 := GenS2{F0: -1, F1: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]bool)
    var v3 string = (v1 + ("a" + "\n\t"))
    v4 := make(map[string]bool)
    var v10 float64 = 0.0
    _ = len(v2)
    if (255 > 0) {
        var v11 int = -1
        v12 := GenS2{F0: -128, F1: -9223372036854775808}
        _ = v11
        _ = v12
    } else {
        fmt.Println(v10, v1)
        if true {
            var v13 int64 = 1
            _ = v13
        } else {
            switch (("hello" + v3) + v1) {
            case "\n\t":
                for fi14 := 0; fi14 < 3; fi14++ {
                    if true {
                        if ((v3 + "a") <= (v3 + v1)) {
                            v15 := GenS3{F0: 42}
                            _ = v15
                        } else {
                            var v16 string = ("foo bar" + v1)
                            _ = v16
                        }
                        v17 := []int{(-(fi14))}
                        v3 = v3
                        _ = v17
                    } else {
                        switch 377 {
                        case 128:
                            if false {
                                var v18 uint16 = 974
                                v19 := GenS4{F0: (fi14 % ((fi14) | 1)), F1: 4294967295}
                                v20 := GenS5{F0: -124, F1: 9223372036854775807, F2: 1}
                                _ = v18
                                _ = v19
                                _ = v20
                            } else {
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    v22 := GenS3{F0: 1}
                                    v3 = v1
                                    _ = fi21
                                    _ = v22
                                }
                                var v23 int64 = 9223372036854775807
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    if (!true) {
                                        if true {
                                            switch "a" {
                                            case "hello":
                                                v23 -= v23
                                                switch ((v3 + "") + (v3 + "")) {
                                                case "\n\t":
                                                    v25 := []int{}
                                                    _ = v25
                                                default:
                                                    var v26 float64 = v10
                                                    _ = v26
                                                }
                                                var v27 uint64 = 18446744073709551615
                                                _ = v27
                                            case "a":
                                                var v28 uint16 = 65535
                                                _ = v28
                                            case "foo bar":
                                                var v29 float32 = (0.1 / float32((((1.5 - (0.0 + v10)) / (0.1 / (v10 - -2.25))) / ((0.1 + v10) + 0.1))))
                                                _ = v29
                                            }
                                            v30 := GenS5{F0: 127, F1: -9223372036854775808, F2: 677}
                                            v1 = v1
                                            _ = v30
                                        }
                                        v31 := []int{fi24}
                                        _ = v31
                                    }
                                    v32 := GenS2{F0: 31, F1: v23}
                                    v10 += ((v10 + v10) / v10)
                                    _ = fi24
                                    _ = v32
                                }
                                _ = v23
                            }
                        case -9223372036854775808:
                            v33 := GenS2{F0: 0, F1: 1}
                            _ = v33
                        case 1:
                            fmt.Println(v1, v10)
                        }
                    }
                    _ = fi14
                }
            case "foo bar":
                var v34 bool = (!true)
                _ = v34
            }
        }
    }
    v35 := []int{-2147483648, 9223372036854775807, 575}
    fmt.Println(v10, v10)
    v36 := GenS6{F0: 525}
    var v37 uint32 = 114
    var v38 uint64 = uint64(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
