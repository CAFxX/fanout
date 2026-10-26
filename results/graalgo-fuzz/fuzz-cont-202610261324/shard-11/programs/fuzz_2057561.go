package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint8; F1 int16; F2 int16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int32, p1 int8, p2 int16) string {
    var v5 bool = (!true)
    var v6 float32 = -2.25
    if v5 {
        p1 = p1
        p2 -= (p2 - p2)
        v7 := GenS1{F0: 42, F1: 1, F2: -96}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 int = 127
    var v2 int = (v1 | (v1 >> ((70) & 63)))
    var v3 uint8 = 1
    v4 := make(map[uint32]string)
    delete(v4, 1)
    switch "hello" {
    case "hello":
        var v8 uint32 = 2
        _ = v8
    case "foo bar":
        v9 := GenS2{F0: 1, F1: (v3 / ((v3) | 1))}
        _ = v9
    }
    v1 = v1
    v2 += -9223372036854775808
    v10 := []int{(v2 & v1), (v2 * v2), (v1 % ((v2) | 1))}
    if true {
        delete(v4, 2)
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            if (false && false) {
                v12 := GenS3{F0: 42}
                var v13 uint = 255
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS1{F0: 32767, F1: 208, F2: -1}
                    _ = len(v10)
                    _ = fi14
                    _ = v15
                }
                _ = v12
                _ = v13
            }
            if len(v10) > 0 {
                v10[0] = -386
            }
            v16 := v4[2147483647]
            _ = fi11
            _ = v16
        }
        var v17 uint8 = 0
        _ = v17
    }
    if true {
        fmt.Println(v1)
    } else {
        if true {
            if false {
                switch 128 {
                case -2147483648:
                    switch int(v3) {
                    case 42:
                        var v18 uint = 18446744073709551615
                        _ = v18
                    case 128:
                        for fi19 := 0; fi19 < 14; fi19++ {
                            if (-388 < -1) {
                                var v20 int16 = 32767
                                _ = v20
                            } else {
                                var v21 int8 = 4
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    var v23 float64 = 0.0
                                    v24 := GenS4{F0: 18446744073709551615, F1: 2147483647}
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                var v25 int16 = 8
                                _ = v21
                                _ = v25
                            }
                            _ = len(v4)
                            var v26 int64 = 9223372036854775807
                            _ = fi19
                            _ = v26
                        }
                    case 976:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            var v28 int64 = -1
                            switch "" {
                            case "foo bar":
                                if false {
                                    for fi29 := 0; fi29 < 4; fi29++ {
                                        var v30 int8 = 102
                                        v1 = 284
                                        fmt.Println(v30, v2, v3)
                                        _ = fi29
                                        _ = v30
                                    }
                                    v31 := GenS5{F0: (v1 - v2), F1: v3, F2: (v28 - v28)}
                                    v32 := GenS6{F0: 0, F1: 0, F2: 923}
                                    _ = v31
                                    _ = v32
                                }
                            default:
                                var v33 float32 = 0.1
                                _ = v33
                            }
                            _ = fi27
                            _ = v28
                        }
                    default:
                        var v34 int = (v1 * 574)
                        _ = v34
                    }
                default:
                    var v35 float64 = -2.25
                    _ = v35
                }
            }
        } else {
            var v36 int64 = 42
            _ = v36
        }
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 float64 = 0.0
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS6{F0: v3, F1: -32768, F2: 1}
                v41 := GenS3{F0: 1000}
                v38 += (v38 * v38)
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
        v42 := GenS7{F0: 327, F1: -9223372036854775808}
        _ = v42
    }
    delete(v4, 2147483648)
    fmt.Println(v1)
    fmt.Println(v1)
    if (818 <= 662) {
        for fi43 := 0; fi43 < 11; fi43++ {
            for fi44 := 0; fi44 < 13; fi44++ {
                v1 = fi44
                v1 -= 128
                delete(v4, 2147483647)
                _ = fi44
            }
            v3 = v3
            _ = fi43
        }
    }
    if len(v10) > 0 {
        v10[0] = v1
    }
    fmt.Println(v1, v3, v3)
    v45 := GenS8{F0: (v3 * v3), F1: 127}
    var v46 int16 = -83
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
