package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint32 }
func genFn0() bool {
    v4 := GenS1{F0: 0}
    var v5 uint8 = 107
    var v6 uint64 = 687
    switch -9223372036854775808 {
    case -995:
        var v7 float64 = 1e10
        _ = v7
    case -1:
        var v8 string = (("hello" + "") + "")
        _ = v8
    case 9223372036854775807:
        v5 = v5
    }
    _ = v4
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 int8, p1 int, p2 int32) int32 {
    if false {
        p1 += 32767
        v9 := []int{(p1 ^ 127), 127, 0, p1}
        fmt.Println(p0, p2)
        _ = v9
    }
    if false {
        p0 -= p0
        var v10 uint16 = 0
        _ = v10
    } else {
        p1 = 26
        var v11 float64 = 1.5
        _ = v11
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = (1e10 / (((3.14159 + 1e10) - (1.5 - 1.5)) / 1.5))
    v3 := make(map[uint32]int)
    v12 := GenS1{F0: 0}
    if (0 > -128) {
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                delete(v3, 292)
                fmt.Println(fi14)
                if false {
                    v15 := GenS2{F0: -128, F1: 859}
                    var v16 int8 = 0
                    switch "foo bar" {
                    case "foo bar":
                        var v17 uint32 = 2147483648
                        _ = v17
                    case "hello":
                        switch 361 {
                        case 667:
                            v18 := GenS2{F0: v16, F1: 2147483647}
                            _ = v18
                        case 2147483647:
                            if (-2147483648 < 2147483647) {
                                var v19 float32 = (1e10 - 1.5)
                                if false {
                                    var v20 uint32 = 2147483647
                                    _ = v20
                                }
                                fmt.Println(v1, fi14, v19)
                                var v21 uint32 = 876
                                _ = v19
                                _ = v21
                            }
                        default:
                            var v22 int32 = -934
                            _ = v22
                        }
                    }
                    _ = v15
                    _ = v16
                } else {
                    var v23 float64 = v2
                    v1 = 0
                    _ = v23
                }
                _ = fi14
            }
            _ = fi13
        }
        var v24 uint8 = 0
        fmt.Println(v24, v2)
        _ = v24
    } else {
        v25 := GenS2{F0: 0, F1: 65535}
        fmt.Println(v2, v1)
        _ = v25
    }
    fmt.Println(v2, v2)
    if true {
        switch "" {
        case "\n\t":
            v26 := GenS1{F0: 127}
            _ = v26
        case "":
            var v27 float64 = v2
            _ = v27
        case "foo bar":
            v28 := []int{-199}
            _ = v28
        }
        v29 := GenS2{F0: 42, F1: 2}
        _ = v29
    }
    v30 := GenS1{F0: 0}
    var v31 int8 = -1
    _ = len(v3)
    fmt.Println(v31, v2, v31)
    var v32 int8 = 74
    v31 += (v31 - -115)
    v31 += v32
    var v33 int32 = -2147483648
    v32 += (v31 % (((^(v31))) | 1))
    if true {
        v1 = 845
        v3[0] = -783
        var v34 bool = (false || true)
        _ = v34
    } else {
        v1 = 255
        v35 := GenS3{F0: 1}
        if false {
            if (32767 <= 369) {
                var v36 int64 = -1
                _ = v36
            }
            v31 -= (v32 | v31)
        }
        _ = v35
    }
    v37 := []int{-9223372036854775808, 201}
    if (100 <= -35) {
        switch 0 {
        case 255:
            v37 = append(v37, 297)
        default:
            if true {
                if ((v1 + v1) != v1) {
                    if false {
                        _ = len(v3)
                        for fi38 := 0; fi38 < 16; fi38++ {
                            var v39 int32 = v33
                            _ = fi38
                            _ = v39
                        }
                    }
                } else {
                    v40 := GenS4{F0: 4294967295, F1: 244}
                    delete(v3, 885)
                    _ = v40
                }
            } else {
                var v41 bool = (true || true)
                _ = v41
            }
        }
    } else {
        if len(v37) > 0 {
            v37[0] = 957
        }
    }
    var v42 float32 = 0.0
    if true {
        v3[2147483648] = 388
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v42
}
