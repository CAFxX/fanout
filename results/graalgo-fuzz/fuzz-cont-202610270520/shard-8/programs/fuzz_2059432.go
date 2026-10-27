package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 int64 = 9223372036854775807
    if ((!false) || true) {
        fmt.Println(v4)
    }
    v1 = v1
    var v5 bool = true
    v1 = (v1 - v1)
    fmt.Println(v1, v4, v4)
    v5 = (((!v5) && ((3.14159 / 1.5) == (-2.25 - (0.0 / 1e10)))) && v5)
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            v5 = (!false)
            v7 := GenS1{F0: 647}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{482, fi6, 1, (fi6 | 32767)}
                var v10 int64 = (v4 / ((v4) | 1))
                if v5 {
                    if (-128 != 1) {
                        delete(v2, 65535)
                        switch "hello" {
                        case "":
                            v5 = (370 == 0)
                        }
                        fmt.Println(v5, v1)
                    }
                    var v11 uint = v1
                    var v12 uint16 = 65535
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(fi8, v10)
                    var v13 int64 = -194
                    _ = v13
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            if ((!true) || (-1 <= 2147483647)) {
                if (v5 || v5) {
                    var v14 bool = false
                    var v15 float32 = 3.14159
                    v16 := GenS2{F0: 2, F1: (v1 % ((v1) | 1))}
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint8 = 81
                    fmt.Println(v17, v4, v1)
                    v4 = v4
                    _ = v17
                }
                switch fi6 {
                case 127:
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if (true || true) {
                            v19 := GenS1{F0: fi6}
                            if (!true) {
                                v5 = v5
                                var v20 uint64 = 1
                                _ = v20
                            }
                            switch "a" {
                            case "a":
                                v1 = v1
                            case "foo bar":
                                v21 := GenS2{F0: 128, F1: v1}
                                _ = v21
                            default:
                                switch fi6 {
                                case 0:
                                    var v22 bool = (!v5)
                                    _ = v22
                                case 32767:
                                    v23 := []int{fi6, (fi18 * (^(fi18)))}
                                    _ = v23
                                default:
                                    v4 -= (v4 & v4)
                                }
                            }
                            v5 = v5
                            fmt.Println(fi6, fi18)
                            _ = v19
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                var v24 float64 = 1.5
                                _ = v24
                            case "":
                                v25 := []int{}
                                _ = v25
                            }
                            var v26 int8 = -1
                            _ = v26
                        }
                        _ = fi18
                    }
                default:
                    v1 += 18446744073709551615
                }
            } else {
                _ = len(v2)
            }
        }
        _ = fi6
    }
    var v27 uint64 = 974
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            switch "" {
            case "foo bar":
                var v30 uint8 = 255
                _ = v30
            default:
                fmt.Println(v27, v27)
            }
            _ = fi29
        }
        if v5 {
            var v31 uint32 = 2147483648
            _ = v31
        }
        _ = fi28
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        if (((1.5 * 1.5) + 0.1) != ((0.1 / 1.0) * (((0.1 * 1e10) + 3.14159) * (0.1 * (0.1 * 0.0))))) {
            var v33 bool = (!(128 > 2))
            for fi34 := 0; fi34 < 13; fi34++ {
                fmt.Println(fi34, v27)
                v35 := []int{4294967295, fi34}
                var v36 uint8 = 1
                _ = fi34
                _ = v35
                _ = v36
            }
            var v37 uint = 4294967295
            _ = v33
            _ = v37
        }
        var v38 float32 = 0.0
        fmt.Println(v1)
        _ = fi32
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
