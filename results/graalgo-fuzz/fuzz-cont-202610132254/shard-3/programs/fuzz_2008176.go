package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint16; F1 string; F2 uint16 }
func genFn0(p0 int16) uint32 {
    p0 = (p0 + (-(p0)))
    p0 = (p0 * p0)
    _ = p0
    return 0
}
func genFn1(p0 uint8, p1 uint64) int32 {
    fmt.Println(p1, p1)
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = 9
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{1, 60}
    v2 := make(map[string]int)
    var v3 float32 = 3.14159
    var v5 string = "a"
    var v6 int = 416
    fmt.Println(v5, v3, v6)
    switch (^(v6)) {
    case 1:
        if true {
            fmt.Println(v6, v6)
            v7 := GenS1{F0: 574, F1: 876, F2: -1}
            _ = v7
        }
    }
    var v8 uint = 601
    switch v5 {
    case "a":
        v9 := v2[("foo bar" + v5)]
        _ = v9
    case "":
        v10 := GenS2{F0: 2147483647}
        _ = v10
    default:
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v8, v3, v6)
            v1 = append(v1, 4294967295)
            _ = fi11
        }
    }
    if false {
        v8 = (v8 / ((0) | 1))
        v3 = v3
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS3{F0: ((!(437 == v6)) || (true && false)), F1: v8}
        v6 -= (fi12 * (fi12 << (((^(v8))) & 63)))
        if (!(float64(v3) <= (3.14159 - 3.14159))) {
            v3 = 0.0
            for fi14 := 0; fi14 < 9; fi14++ {
                if (v8 <= (v8 / (((v8 << ((v8) & 63))) | 1))) {
                    if false {
                        v15 := GenS3{F0: false, F1: (v8 - 42)}
                        switch v6 {
                        case 127:
                            fmt.Println(fi14)
                        case 4294967295:
                            if ((v6 == (v6 & fi14)) || false) {
                                switch ("\n\t" + v5) {
                                case "":
                                    v16 := GenS4{F0: -1, F1: 4294967295}
                                    _ = v16
                                default:
                                    switch "a" {
                                    case "a":
                                        switch fi12 {
                                        case 383:
                                            if true {
                                                v1 = append(v1, 0)
                                            }
                                        case 579:
                                            var v17 uint8 = 99
                                            var v18 uint32 = 2
                                            _ = v17
                                            _ = v18
                                        case 128:
                                            if (!false) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v19 := GenS5{F0: (-720 != -1), F1: 0}
                                                _ = v19
                                            } else {
                                                v8 += 915
                                                v20 := GenS6{F0: 53, F1: (true || (127 >= 62)), F2: 255}
                                                v5 = "a"
                                                _ = v20
                                            }
                                        default:
                                            fmt.Println(v5, fi14)
                                        }
                                    default:
                                        if (0 == 596) {
                                            v6 -= -745
                                        }
                                    }
                                }
                                if true {
                                    var v21 uint64 = 839
                                    _ = v21
                                } else {
                                    var v22 uint8 = 1
                                    if (false && true) {
                                        v2[v5] = (fi14 / ((v6) | 1))
                                        switch (-(fi14)) {
                                        case 4294967295:
                                            if (!(!(95 >= (^(v22))))) {
                                                if true {
                                                    var v23 uint64 = 998
                                                    _ = v23
                                                }
                                                if true {
                                                    v8 += (v8 ^ 0)
                                                    v24 := GenS7{F0: 255, F1: "", F2: 481}
                                                    v6 = 793
                                                    _ = v24
                                                }
                                                fmt.Println(fi14)
                                            } else {
                                                v3 += (1e10 / 1.0)
                                            }
                                        }
                                        var v25 uint32 = 65535
                                        _ = v25
                                    }
                                    var v26 int = (fi14 | 42)
                                    _ = v22
                                    _ = v26
                                }
                                if false {
                                    var v27 float32 = v3
                                    var v28 float32 = 0.1
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                v29 := GenS3{F0: false, F1: 253}
                                switch fi12 {
                                case 902:
                                    v6 = (v6 % ((fi14) | 1))
                                case -975:
                                    var v30 uint8 = 128
                                    _ = v30
                                case 42:
                                    var v31 float64 = 0.0
                                    v5 = ""
                                    _ = v31
                                default:
                                    v32 := v2[v5]
                                    _ = v32
                                }
                                v33 := GenS3{F0: ((v8 | (v8 * v8)) != v8), F1: 456}
                                _ = v29
                                _ = v33
                            }
                        }
                        if ((v8 != (v8 + (v8 - 18446744073709551615))) || true) {
                            var v34 bool = true
                            var v35 int = v6
                            switch "a" {
                            case "foo bar":
                                v36 := v2[v5]
                                _ = v36
                            }
                            _ = v34
                            _ = v35
                        } else {
                            var v37 uint16 = 2
                            fmt.Println(v8, v3, v3)
                            _ = len(v1)
                            _ = v37
                        }
                        _ = v15
                    }
                    fmt.Println(v8, v5, v8)
                    fmt.Println(v8, v5)
                }
                _ = fi14
            }
        }
        _ = fi12
        _ = v13
    }
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 float32 = (0.1 / v3)
        v5 = ((v5 + ((v5 + v5) + ("hello" + ((v5 + v5) + v5)))) + "a")
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        v6 += -593
        var v41 float64 = (3.14159 * 1e10)
        _ = fi40
        _ = v41
    }
    v5 = (v5 + (((v5 + v5) + v5) + v5))
    fmt.Println(v6, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
