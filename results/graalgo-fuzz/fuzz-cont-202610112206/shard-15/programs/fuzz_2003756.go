package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0(p0 uint64, p1 uint, p2 uint32) int16 {
    var v4 uint16 = 1
    switch "" {
    case "\n\t":
        fmt.Println(v4)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -542
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 int = 541
    var v5 int32 = 42
    for fi6 := 0; fi6 < 17; fi6++ {
        switch "\n\t" {
        case "\n\t":
            var v7 int8 = 127
            _ = v7
        case "hello":
            var v8 uint16 = 721
            _ = v8
        case "foo bar":
            for fi9 := 0; fi9 < 9; fi9++ {
                if (2147483647 != 65535) {
                    fmt.Println(fi9, v3, fi6)
                    var v10 float64 = 0.0
                    v5 += 410
                    var v11 float64 = 3.14159
                    _ = v10
                    _ = v11
                }
                switch "hello" {
                case "":
                    if v1 {
                        v5 += (v5 + (v5 >> ((4294967295) & 63)))
                        v3 -= 128
                        v12 := GenS1{F0: 241}
                        _ = v12
                    }
                case "a":
                    var v13 uint8 = 2
                    var v14 string = ""
                    _ = v13
                    _ = v14
                default:
                    if (v1 && ((0.0 / (3.14159 - (0.0 + (0.0 + 1e10)))) <= -2.25)) {
                        fmt.Println(v3, fi9, v5)
                        v15 := GenS1{F0: 127}
                        if (1000 >= -1) {
                            var v16 int = (^(fi9))
                            var v17 int32 = (v5 * v5)
                            v1 = v1
                            _ = v16
                            _ = v17
                        } else {
                            switch 127 {
                            case 1:
                                if (1 <= 945) {
                                    switch v3 {
                                    case 42:
                                        if (false || v1) {
                                            v3 -= (fi9 >> ((180) & 63))
                                            var v18 uint = 65535
                                            _ = v18
                                        }
                                    case -1:
                                        if (v1 && (42 == 2147483647)) {
                                            v5 = 42
                                        }
                                    default:
                                        v1 = v1
                                    }
                                    v1 = v1
                                } else {
                                    if (42 <= 2) {
                                        v19 := []int{v3}
                                        v20 := GenS1{F0: 2}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := GenS1{F0: 127}
                                        var v22 uint16 = 2
                                        v23 := []int{(fi9 + 4294967295), -9223372036854775808, 127, 48}
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                v3 = v3
                                v24 := []int{}
                                _ = v24
                            }
                            v1 = true
                            var v25 uint64 = 923
                            _ = v25
                        }
                        if false {
                            v26 := GenS1{F0: 2}
                            if (v1 && v1) {
                                var v27 uint32 = 2147483648
                                v28 := GenS1{F0: 2}
                                v29 := GenS2{F0: 63, F1: -1}
                                fmt.Println(v27, v27)
                                v5 = (v5 & (v5 << ((1) & 63)))
                                v27 -= (v27 << ((4294967295) & 63))
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 bool = v1
                            v3 = v3
                            _ = v26
                            _ = v30
                        }
                        var v31 int = (-(v3))
                        switch (("" + "") + "") {
                        case "a":
                            var v32 uint8 = 10
                            v33 := GenS1{F0: (v32 / (((v32 + (v32 * (^(v32))))) | 1))}
                            _ = v32
                            _ = v33
                        case "\n\t":
                            var v34 int32 = v5
                            _ = v34
                        }
                        fmt.Println(v1, fi9, v3)
                        _ = v15
                        _ = v31
                    }
                }
                var v35 uint16 = 214
                v3 = (v3 | (-(fi6)))
                v36 := GenS3{F0: ""}
                v3 -= 316
                v1 = (v1 || false)
                _ = fi9
                _ = v35
                _ = v36
            }
        default:
            var v37 int = 32767
            _ = v37
        }
        v3 = 0
        _ = fi6
    }
    v38 := []int{}
    var v39 float32 = float32(-2.25)
    switch "" {
    case "\n\t":
        var v40 uint16 = 769
        _ = v40
    case "foo bar":
        var v41 uint = 42
        _ = v41
    default:
        delete(v2, "foo bar")
    }
    var v42 int8 = 55
    v39 += (1e10 - v39)
    fmt.Println(v39, v1)
    var v43 int8 = (-(v42))
    var v44 int64 = int64(v5)
    v45 := GenS4{F0: 724, F1: 65535, F2: 1000}
    v3 = 467
    _ = len(v38)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
