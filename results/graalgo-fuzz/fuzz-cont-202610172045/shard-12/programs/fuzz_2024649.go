package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int32 }
func genFn0(p0 uint32, p1 int8, p2 int) uint64 {
    var v7 float32 = 0.1
    v7 = v7
    if true {
        if (-1 <= -197) {
            p1 = p1
        } else {
            if false {
                v7 += v7
                p2 = p2
            }
            var v8 int8 = (p1 * p1)
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 bool) int8 {
    v9 := GenS1{F0: 32768, F1: 2}
    if (!p0) {
        switch 42 {
        case 42:
            v10 := []int{}
            _ = v10
        default:
            var v11 int = 9223372036854775807
            _ = v11
        }
        switch "foo bar" {
        case "\n\t":
            switch "hello" {
            case "hello":
                switch 0 {
                case 2147483647:
                    if p0 {
                        for fi12 := 0; fi12 < 8; fi12++ {
                            p0 = p0
                            _ = fi12
                        }
                    }
                case 9223372036854775807:
                    p0 = p0
                default:
                    var v13 int32 = -2147483648
                    _ = v13
                }
            case "foo bar":
                p0 = (float32((1e10 / ((1e10 - (1.5 + 3.14159)) * 1.5))) >= 0.0)
            case "\n\t":
                fmt.Println(p0)
            }
        }
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint32 = 2147483647
            var v16 int32 = 0
            if true {
                var v17 float64 = ((1.5 * 0.1) * -2.25)
                fmt.Println(v16, fi14)
                _ = v17
            } else {
                var v18 bool = (!false)
                p0 = (114 != 92)
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v19 int16 = 1000
        _ = v19
    }
    fmt.Println(p0)
    v20 := GenS1{F0: -34, F1: 806}
    _ = p0
    _ = v20
    _ = v9
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 32767
    var v3 bool = true
    v4 := make(map[string]string)
    v5 := []int{-9223372036854775808}
    v6 := []uint32{0, 4294967295, 4294967295}
    fmt.Println(v2, v1, v2)
    v3 = v3
    v21 := GenS2{F0: 2, F1: 112}
    switch ("a" + "") {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v1, v1)
    }
    v3 = (!(v2 <= v2))
    if true {
        var v22 uint16 = (v2 % ((v2) | 1))
        v2 += (^(v22))
        _ = v22
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        fmt.Println(v2, v1)
    }
    if v3 {
        if v3 {
            if (!(((1.5 * (0.0 + (0.1 * 0.1))) >= 1.5) || v3)) {
                if v3 {
                    v23 := GenS3{F0: v3}
                    _ = v23
                } else {
                    var v24 float64 = (1e10 + 1.5)
                    switch 831 {
                    case 1:
                        for fi25 := 0; fi25 < 19; fi25++ {
                            var v26 string = ""
                            fmt.Println(v26, fi25, v3)
                            _ = fi25
                            _ = v26
                        }
                    case -2147483648:
                        var v27 uint16 = v2
                        _ = v27
                    default:
                        v28 := GenS1{F0: 9223372036854775807, F1: 876}
                        _ = v28
                    }
                    v1 = (v1 | v1)
                    _ = v24
                }
                fmt.Println(v2, v2)
                for fi29 := 0; fi29 < 14; fi29++ {
                    if false {
                        v2 = v2
                        switch (fi29 + (fi29 & (fi29 << ((18446744073709551615) & 63)))) {
                        case 32767:
                            v30 := v4[("foo bar" + "foo bar")]
                            _ = v30
                        default:
                            switch 127 {
                            case -327:
                                for fi31 := 0; fi31 < 14; fi31++ {
                                    if v3 {
                                        v2 = v2
                                        v1 -= 276
                                    } else {
                                        var v32 uint = 0
                                        _ = len(v6)
                                        _ = v32
                                    }
                                    _ = fi31
                                }
                            case -2147483648:
                                if v3 {
                                    v2 = v2
                                    _ = len(v6)
                                    if v3 {
                                        if v3 {
                                            v33 := GenS3{F0: (!(1 <= v2))}
                                            v1 += (^(v1))
                                            v34 := v4["hello"]
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 int = (fi29 + (^(fi29)))
                                        _ = v35
                                    }
                                }
                            default:
                                v36 := GenS3{F0: v3}
                                _ = v36
                            }
                        }
                        v5 = append(v5, fi29)
                    }
                    v37 := GenS2{F0: 4294967295, F1: 4294967295}
                    for fi38 := 0; fi38 < 4; fi38++ {
                        v39 := GenS4{F0: (fi29 ^ fi38), F1: 2147483647}
                        _ = fi38
                        _ = v39
                    }
                    _ = fi29
                    _ = v37
                }
            }
            fmt.Println(v3, v2, v1)
        }
        var v40 float64 = float64(-2.25)
        if true {
            var v41 int32 = (v1 | v1)
            var v42 int32 = -1
            if ((32767 >= 9223372036854775807) || (((!(1e10 <= float64(1e10))) || (-32768 != 440)) || (42 == 42))) {
                v5 = append(v5, 4294967295)
            } else {
                var v43 uint64 = 4294967295
                var v44 float32 = (0.0 * (0.1 - ((3.14159 + 0.1) / 1.0)))
                _ = v43
                _ = v44
            }
            _ = v41
            _ = v42
        } else {
            for fi45 := 0; fi45 < 4; fi45++ {
                v1 -= v1
                v46 := GenS3{F0: false}
                _ = fi45
                _ = v46
            }
            var v47 uint32 = 65535
            _ = v47
        }
        _ = v40
    } else {
        v48 := GenS2{F0: 0, F1: 0}
        var v49 uint16 = (^(v2))
        if len(v6) > 0 {
            v6[0] = 858
        }
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
