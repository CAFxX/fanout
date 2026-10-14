package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 uint64 }
type GenS6 struct { F0 int32 }
func genFn0() bool {
    var v7 uint8 = 55
    var v8 uint32 = 2147483647
    v9 := []int{-9223372036854775808, 9223372036854775807, 127}
    v8 = v8
    var v10 uint8 = 1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 int32) uint32 {
    p0 = 705
    p0 -= p0
    _ = p0
    return 504
}

func main() {
    v1 := []int{-514, 32767}
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 string = ""
    var v11 float64 = -2.25
    v11 -= v11
    var v12 int8 = 0
    switch 128 {
    case -9223372036854775808:
        v12 += (v12 ^ 97)
    case 127:
        if (float64((0.0 / 1.0)) != (3.14159 + v11)) {
            var v13 int16 = 0
            if ((!false) || (true && (1e10 == (1.5 + 0.0)))) {
                v11 += v11
            }
            v14 := v4[v6]
            _ = v13
            _ = v14
        } else {
            if (-2.25 == 1.5) {
                for fi15 := 0; fi15 < 9; fi15++ {
                    delete(v4, "\n\t")
                    _ = fi15
                }
            }
            var v16 uint16 = 0
            _ = v16
        }
    default:
        v17 := GenS1{F0: 89, F1: 351}
        _ = v17
    }
    var v18 bool = false
    v6 = v6
    if v18 {
        if v18 {
            v19 := GenS1{F0: 1, F1: 0}
            _ = v19
        }
    }
    if (18446744073709551615 <= 38) {
        fmt.Println(v11, v18, v11)
        if true {
            var v20 string = "foo bar"
            if (373 == 0) {
                v21 := GenS1{F0: 255, F1: 554}
                for fi22 := 0; fi22 < 15; fi22++ {
                    v12 -= (v12 << ((42) & 63))
                    v11 = float64(3.14159)
                    var v23 uint8 = 255
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                if (((true || false) && ((9223372036854775807 >= -895) && v18)) || ((671 > -460) && true)) {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        fmt.Println(v20, fi24, v12)
                        switch (fi24 & 1) {
                        case 127:
                            v6 = v6
                        case -486:
                            v25 := GenS2{F0: true, F1: 477, F2: 255}
                            _ = v25
                        case 42:
                            v18 = (-797 <= 42)
                        }
                        v20 = "foo bar"
                        _ = fi24
                    }
                    v26 := GenS3{F0: 2147483647}
                    _ = v26
                }
                fmt.Println(v12)
                if ((-(v12)) < (v12 * -128)) {
                    v27 := GenS2{F0: (-500 <= -1), F1: 128, F2: 68}
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS4{F0: 0}
                        v12 += v12
                        fmt.Println(v11, fi28)
                        _ = fi28
                        _ = v29
                    }
                    v12 += v12
                    _ = v27
                }
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int8 = v12
                v31 -= v12
                switch 1 {
                case 138:
                    for fi32 := 0; fi32 < 12; fi32++ {
                        switch v6 {
                        case "foo bar":
                            var v33 int64 = -9223372036854775808
                            _ = v33
                        case "":
                            v18 = v18
                        case "\n\t":
                            var v34 int64 = 1
                            _ = v34
                        default:
                            v1 = append(v1, fi32)
                        }
                        if len(v1) > 0 {
                            v1[0] = (fi30 << ((42) & 63))
                        }
                        if v18 {
                            v35 := GenS5{F0: 942, F1: (("foo bar" + "") + (v6 + v6)), F2: 4294967295}
                            switch 0 {
                            case 2147483647:
                                var v36 uint = 255
                                _ = v36
                            default:
                                v37 := GenS1{F0: 42, F1: 2}
                                _ = v37
                            }
                            fmt.Println(v20)
                            if ((((v20 + (("hello" + "") + "\n\t")) + ((v6 + (((((v20 + v6) + (v6 + "hello")) + "\n\t") + "foo bar") + (((("" + v20) + (v6 + v20)) + v6) + ((v6 + "hello") + (v6 + ""))))) + v20)) + (("foo bar" + "a") + v6)) != v6) {
                                v31 += v12
                            }
                            if v18 {
                                v20 = (((v6 + "") + "a") + v6)
                            } else {
                                v38 := GenS6{F0: -2147483648}
                                v31 += (-(v31))
                                if len(v1) > 0 {
                                    v1[0] = fi30
                                }
                                _ = v38
                            }
                            v39 := v2[2]
                            _ = v35
                            _ = v39
                        }
                        _ = fi32
                    }
                case 128:
                    v11 = (v11 + (v11 * 0.0))
                case -9223372036854775808:
                    fmt.Println(fi30, v20, v12)
                default:
                    fmt.Println(v31)
                }
                _ = fi30
                _ = v31
            }
            _ = v20
        }
    } else {
        var v40 int32 = 2147483647
        v6 = v6
        v6 = (v6 + "")
        _ = v40
    }
    var v41 uint = 534
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
