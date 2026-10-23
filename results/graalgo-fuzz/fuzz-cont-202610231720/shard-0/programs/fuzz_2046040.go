package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int8; F1 int32 }
func genFn0(p0 int64) int8 {
    var v7 uint64 = 910
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint8 = 1
        var v10 float32 = 1.5
        v11 := []int{1, fi8, fi8, fi8}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float32 = 0.0
    fmt.Println(v12, v12, v12)
    var v13 float64 = 1e10
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 255
    var v3 int32 = 311
    var v4 uint16 = 32767
    v5 := []string{"", "a", "hello"}
    var v6 int64 = 42
    if (!((v3 - (v3 / (((v3 | (-(v3)))) | 1))) <= v3)) {
        v3 = v3
        for fi14 := 0; fi14 < 17; fi14++ {
            v1[2] = (("foo bar" + "a") + "foo bar")
            if (true || true) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v2 = 42
                    v16 := GenS1{F0: (^(fi15)), F1: 0}
                    if ((2 != 125) && (true && false)) {
                        v17 := GenS2{F0: 118}
                        if true {
                            if (false && ((!((1.5 - 1.5) > (1.5 + 1.5))) && true)) {
                                switch "\n\t" {
                                case "a":
                                    v2 = 42
                                case "hello":
                                    v6 += (^(v6))
                                case "":
                                    var v18 int8 = 2
                                    _ = v18
                                default:
                                    fmt.Println(v2, fi15, v3)
                                }
                                v3 = (^(v3))
                                var v19 int16 = -728
                                _ = v19
                            }
                            v20 := GenS2{F0: 543}
                            _ = v20
                        } else {
                            if (true || true) {
                                v21 := GenS2{F0: 2}
                                switch "hello" {
                                case "foo bar":
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        _ = len(v1)
                                    }
                                case "\n\t":
                                    v6 = 374
                                }
                                v22 := GenS3{F0: 127}
                                _ = v21
                                _ = v22
                            }
                            v2 -= (^(v4))
                        }
                        if true {
                            v3 = 0
                            fmt.Println(v3)
                            v4 += 661
                        } else {
                            var v23 float32 = 0.1
                            switch "a" {
                            case "foo bar":
                                v24 := GenS1{F0: (fi14 & (^(fi14))), F1: 4294967295}
                                _ = v24
                            case "hello":
                                if true {
                                    v2 = v4
                                }
                            }
                            _ = v23
                        }
                        v5 = append(v5, "")
                        _ = v17
                    }
                    v2 -= 0
                    _ = fi15
                    _ = v16
                }
                switch "a" {
                case "hello":
                    v25 := GenS3{F0: 42}
                    _ = v25
                case "a":
                    v6 = (-(v6))
                case "\n\t":
                    if false {
                        v5 = append(v5, "\n\t")
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 int8 = 1
                            var v28 int32 = (v3 | (^(v3)))
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                default:
                    _ = len(v5)
                }
            } else {
                v29 := v1[1]
                var v30 uint = 255
                v31 := GenS4{F0: (v6 + v6), F1: 1, F2: 2}
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 float32 = float32((3.14159 + 3.14159))
            _ = fi14
            _ = v32
        }
        delete(v1, 65535)
    } else {
        v4 = v4
        fmt.Println(v2)
        var v33 float64 = 1e10
        _ = v33
    }
    v4 += v4
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            _ = len(v5)
            _ = fi35
        }
        var v36 bool = (!(!false))
        var v37 uint32 = 4294967295
        _ = fi34
        _ = v36
        _ = v37
    }
    switch 32768 {
    case -454:
        v38 := GenS2{F0: 508}
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(fi39, v2, fi39)
        _ = fi39
    }
    v40 := GenS1{F0: -2147483648, F1: 210}
    v4 = (v4 * 2)
    var v41 int32 = -2147483648
    if (-868 != 9223372036854775807) {
        var v42 int64 = (v6 % ((v6) | 1))
        v43 := GenS5{F0: -120, F1: v41}
        var v44 string = ""
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
