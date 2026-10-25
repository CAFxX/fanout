package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
func genFn0(p0 bool, p1 uint) int32 {
    v4 := []int{34, 0, 2147483647, 127}
    p1 += (^(p1))
    p0 = p0
    v5 := GenS1{F0: 1, F1: (0 != 42)}
    p0 = ((((!((!p0) && p0)) && true) && true) || (!p0))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 3.14159
    var v3 bool = true
    var v6 bool = false
    fmt.Println(v3, v2)
    if v6 {
        v7 := GenS2{F0: 127, F1: 42}
        v8 := []int{32767, 128, 32768, -2147483648}
        _ = v7
        _ = v8
    }
    var v9 int64 = 217
    var v10 int64 = (v9 * v9)
    v11 := []int{-2147483648, -2147483648}
    v10 = (-(v9))
    var v12 uint16 = 124
    switch 127 {
    case 497:
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                if (127 != 2) {
                    var v15 float64 = v2
                    _ = v15
                } else {
                    v12 -= 65535
                    v6 = false
                    switch (-(fi13)) {
                    case 9223372036854775807:
                        var v16 string = "a"
                        _ = v16
                    case 943:
                        if ((1.5 - v2) != float64(0.0)) {
                            var v17 float32 = 0.1
                            var v18 bool = (0 >= 42)
                            _ = v17
                            _ = v18
                        } else {
                            var v19 bool = v3
                            v1[2] = true
                            if true {
                                if false {
                                    v6 = v19
                                    fmt.Println(v19, v10)
                                }
                            } else {
                                var v20 int16 = 746
                                _ = v20
                            }
                            if true {
                                if len(v11) > 0 {
                                    v11[0] = 32768
                                }
                                v3 = true
                                v12 = v12
                            } else {
                                switch "foo bar" {
                                case "":
                                    v10 = (-(v10))
                                case "hello":
                                    fmt.Println(v2, v19, fi13)
                                case "\n\t":
                                    v10 -= (v9 + -9223372036854775808)
                                }
                                if true {
                                    fmt.Println(v10)
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                                v3 = v19
                                var v21 int32 = 323
                                v2 = 3.14159
                                _ = v21
                            }
                            _ = v19
                        }
                    default:
                        v22 := GenS3{F0: 127, F1: 0}
                        _ = v22
                    }
                }
                switch fi14 {
                case 9223372036854775807:
                    var v23 int16 = 1
                    _ = v23
                default:
                    v24 := GenS2{F0: 0, F1: 35}
                    _ = v24
                }
                if (!v6) {
                    var v25 int16 = 32767
                    v1[2147483648] = v3
                    switch -645 {
                    case 128:
                        switch ("" + "") {
                        case "a":
                            fmt.Println(v10, v10)
                        case "":
                            var v26 uint8 = 238
                            _ = v26
                        default:
                            var v27 float32 = 0.0
                            var v28 float64 = 0.0
                            _ = v27
                            _ = v28
                        }
                    case 2147483647:
                        switch (fi13 + (fi13 / ((fi13) | 1))) {
                        case 424:
                            v29 := GenS4{F0: (-128 <= 75), F1: "", F2: v6}
                            _ = v29
                        default:
                            fmt.Println(fi13, v6, fi14)
                        }
                    default:
                        var v30 uint8 = 1
                        _ = v30
                    }
                    _ = v25
                }
                _ = fi14
            }
            _ = fi13
        }
    case -9223372036854775808:
        fmt.Println(v9, v6, v10)
    case 127:
        var v31 float64 = v2
        _ = v31
    default:
        v32 := GenS5{F0: 133}
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v10 += v9
        if ((true && v6) || false) {
            if ((v12 ^ v12) > (^(v12))) {
                for fi34 := 0; fi34 < 1; fi34++ {
                    v10 = (v9 % (((-(v10))) | 1))
                    _ = fi34
                }
                v1[16] = v3
                var v35 float64 = v2
                _ = v35
            }
            v3 = (!(373 >= 2))
        }
        _ = fi33
    }
    if ((127 == 0) && (v2 < ((v2 * 3.14159) - 1e10))) {
        v36 := GenS6{F0: 0, F1: 2147483647}
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
