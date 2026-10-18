package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]bool)
    var v3 float64 = (float64(0.1) + -2.25)
    var v4 bool = (!(324 != 1))
    v5 := make(map[uint32]uint32)
    var v6 bool = false
    if v4 {
        v4 = v4
    }
    v4 = (!v4)
    if (v3 < 0.0) {
        v1 -= (v1 + -142)
        if v6 {
            switch -134 {
            case -2147483648:
                delete(v5, 65535)
            default:
                var v7 float64 = v3
                _ = v7
            }
            switch 2147483647 {
            case 2147483647:
                for fi8 := 0; fi8 < 18; fi8++ {
                    if (0 == 1) {
                        if v4 {
                            v4 = true
                            var v9 uint = 0
                            _ = v9
                        } else {
                            switch 413 {
                            case -461:
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    v10 := []int{-2147483648, -406, 9223372036854775807}
                                    _ = v10
                                case "hello":
                                    if v6 {
                                        v11 := GenS1{F0: 1, F1: -889, F2: int16(v1)}
                                        v12 := GenS1{F0: 32767, F1: v1, F2: 42}
                                        _ = v11
                                        _ = v12
                                    }
                                case "a":
                                    v1 = (v1 * 871)
                                }
                            case 700:
                                delete(v5, 42)
                            }
                        }
                        if v6 {
                            var v13 uint16 = 0
                            _ = v13
                        }
                    }
                    if v4 {
                        var v14 uint32 = 2
                        _ = v14
                    }
                    _ = fi8
                }
            case 0:
                for fi15 := 0; fi15 < 7; fi15++ {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch (^(fi16)) {
                        case 0:
                            v17 := []int{(-(fi16)), fi16, (-(fi16)), 61}
                            _ = v17
                        default:
                            var v18 uint = 4294967295
                            _ = v18
                        }
                        v4 = v4
                        _ = len(v2)
                        _ = fi16
                    }
                    _ = fi15
                }
            default:
                if (!(!(!(!(65535 <= 534))))) {
                    _ = len(v2)
                    v4 = v4
                } else {
                    var v19 int64 = -1
                    if ((v4 && (42 > -2147483648)) || v4) {
                        v5[869] = 0
                        v4 = (!(!v6))
                    } else {
                        v3 = v3
                        v19 += v19
                        for fi20 := 0; fi20 < 14; fi20++ {
                            for fi21 := 0; fi21 < 11; fi21++ {
                                if (-1 > 886) {
                                    v2["\n\t"] = (65535 == 32767)
                                    var v22 uint16 = 1
                                    var v23 uint32 = 0
                                    fmt.Println(v22, v1, v4)
                                    _ = v22
                                    _ = v23
                                }
                                v3 += v3
                                _ = fi21
                            }
                            _ = fi20
                        }
                    }
                    v19 = (v19 | v1)
                    _ = v19
                }
            }
            var v24 int32 = -636
            _ = v24
        }
        if (!(false || (v1 > v1))) {
            fmt.Println(v4)
            v25 := GenS1{F0: 1, F1: v1, F2: 32767}
            var v26 uint8 = 99
            _ = v25
            _ = v26
        } else {
            var v27 uint64 = 2
            _ = v27
        }
    } else {
        fmt.Println(v4, v3, v3)
    }
    switch "foo bar" {
    case "\n\t":
        v6 = (v6 || true)
    default:
        v3 = ((v3 * ((((0.1 / -2.25) / v3) + ((0.1 * 0.1) * float64(float32(v3)))) + (3.14159 / v3))) / 1.0)
    }
    v28 := v2[("hello" + "foo bar")]
    var v29 uint64 = 0
    v30 := []int{}
    if v4 {
        switch 4294967295 {
        case 128:
            switch 32767 {
            case 32767:
                if (1 == 35) {
                    v31 := GenS2{F0: v1, F1: 42}
                    _ = len(v5)
                    _ = v31
                }
            case 1:
                v28 = v4
            default:
                v32 := GenS2{F0: 0, F1: 255}
                _ = v32
            }
        case 34:
            switch 32767 {
            case -9223372036854775808:
                if false {
                    var v33 float32 = 0.0
                    v4 = true
                    _ = v33
                }
            case 945:
                v29 = (v29 * v29)
            case 127:
                var v34 float32 = (0.1 / 1.5)
                _ = v34
            default:
                _ = len(v5)
            }
        }
        var v35 int = 32767
        _ = v35
    }
    v30 = append(v30, 4294967295)
    delete(v5, 72)
    var v36 int64 = v1
    var v37 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
