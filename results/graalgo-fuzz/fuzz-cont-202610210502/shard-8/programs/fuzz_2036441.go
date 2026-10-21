package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0() uint32 {
    v7 := []int{9223372036854775807, -1, 32767, 32767}
    v7 = append(v7, 337)
    var v8 int = 128
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 42, F1: 2}
        var v11 int = v8
        fmt.Println(v11, v8, v11)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return 303
}
func genFn1() uint {
    var v12 float64 = 1e10
    if false {
        switch 787 {
        case -120:
            v13 := GenS2{F0: 0, F1: 255}
            _ = v13
        case 715:
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 uint32 = 355
                _ = fi14
                _ = v15
            }
        default:
            v12 = 0.0
        }
        var v16 uint64 = 42
        _ = v16
    }
    _ = v12
    return 2
}

func main() {
    v1 := []int{127}
    v2 := make(map[string]bool)
    var v3 float32 = float32(3.14159)
    var v4 string = "\n\t"
    var v5 float64 = (float64(v3) + 1e10)
    v6 := make(map[uint32]uint32)
    var v17 string = v4
    v18 := v2[v4]
    var v19 int8 = 100
    var v20 int64 = -9223372036854775808
    v21 := GenS3{F0: (v17 + "foo bar"), F1: 65535}
    v18 = v18
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS3{F0: v4, F1: 1}
        var v24 int16 = -1
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint = 65535
    v2["\n\t"] = v18
    v26 := GenS2{F0: 2, F1: 144}
    v20 += v20
    v27 := v2["a"]
    if true {
        for fi28 := 0; fi28 < 19; fi28++ {
            switch v4 {
            case "foo bar":
                if (4294967295 < 18446744073709551615) {
                    v18 = false
                    v1 = append(v1, 419)
                } else {
                    switch 127 {
                    case -1:
                        v20 -= (v20 | 1)
                    case 737:
                        if true {
                            switch v17 {
                            case "":
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    if false {
                                        if false {
                                            var v30 float64 = 1.5
                                            var v31 string = ((v17 + v4) + ("foo bar" + v4))
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    v20 = (v20 | -9223372036854775808)
                                    if (65535 >= 264) {
                                        fmt.Println(fi29, fi28, v20)
                                        var v32 string = ("\n\t" + ("a" + v17))
                                        var v33 float32 = (((v3 * 0.1) / v3) / v3)
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v34 := GenS1{F0: fi28, F1: 1}
                                        _ = v34
                                    }
                                    _ = fi29
                                }
                            case "\n\t":
                                switch (fi28 << ((genFn1()) & 63)) {
                                case 0:
                                    if (2147483647 <= 42) {
                                        v3 = v3
                                    } else {
                                        v17 = "a"
                                        v35 := GenS3{F0: v17, F1: 620}
                                        var v36 uint = 710
                                        _ = v35
                                        _ = v36
                                    }
                                case 9223372036854775807:
                                    var v37 uint64 = 1
                                    _ = v37
                                }
                            }
                        } else {
                            v4 = v4
                            v27 = true
                            v18 = (2 <= 18446744073709551615)
                        }
                    case 127:
                        if (("\n\t" == "\n\t") && false) {
                            _ = len(v1)
                            v27 = ((!false) && (float32(3.14159) < v3))
                        }
                    }
                }
            case "a":
                fmt.Println(v19, v18)
            case "\n\t":
                if false {
                    _ = len(v6)
                    v18 = false
                    v20 = v20
                } else {
                    delete(v2, v4)
                }
            }
            _ = fi28
        }
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        v19 = v19
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := GenS4{F0: (v27 || (507 < 42))}
            v18 = false
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v41 := GenS5{F0: 42, F1: 240}
    var v42 uint64 = 927
    if v27 {
        v43 := GenS1{F0: 644, F1: 65535}
        if v27 {
            v3 = float32(3.14159)
            fmt.Println(v4, v18)
        }
        v44 := GenS5{F0: 1, F1: genFn0()}
        _ = v43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
