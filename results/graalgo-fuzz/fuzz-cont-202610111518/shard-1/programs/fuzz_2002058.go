package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0() int64 {
    var v6 int16 = -453
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{(fi7 << ((0) & 63))}
            if (((v6 << ((255) & 63)) <= v6) || ((!false) || ((!false) || (42 <= 2)))) {
                if (v6 > v6) {
                    v6 -= (v6 + 163)
                    switch "a" {
                    case "foo bar":
                        switch fi8 {
                        case 9223372036854775807:
                            v6 += (v6 - (v6 & (v6 ^ 1000)))
                        case 682:
                            if (fi7 > fi7) {
                                fmt.Println(v6, fi8)
                            } else {
                                v6 = (v6 & 32767)
                                v6 = (^(v6))
                            }
                        default:
                            var v10 int64 = 1
                            _ = v10
                        }
                    case "\n\t":
                        var v11 int = fi8
                        _ = v11
                    case "":
                        var v12 int8 = 42
                        _ = v12
                    default:
                        if false {
                            v6 = (v6 >> ((370) & 63))
                            var v13 float64 = -2.25
                            _ = v13
                        }
                    }
                } else {
                    var v14 int8 = -116
                    v15 := GenS1{F0: 856, F1: 128, F2: 255}
                    if false {
                        fmt.Println(fi7)
                        v6 = -32768
                        _ = len(v9)
                        v16 := GenS1{F0: 1, F1: -2147483648, F2: 568}
                        fmt.Println(v6)
                        _ = v16
                    }
                    fmt.Println(v14, v14, v14)
                    _ = v14
                    _ = v15
                }
                if false {
                    fmt.Println(v6, fi7, v6)
                }
                var v17 float32 = 1.5
                _ = v17
            }
            _ = fi8
            _ = v9
        }
        v18 := []int{}
        _ = fi7
        _ = v18
    }
    _ = v6
    return 42
}

func main() {
    var v1 bool = (4294967295 <= 71)
    var v2 int = -2147483648
    var v3 uint64 = 42
    v4 := make(map[string]int)
    var v5 uint8 = 42
    if v1 {
        var v19 uint64 = 330
        v2 += (v2 | -450)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := []int{fi21}
                v2 += fi21
                if (1 >= 0) {
                    v23 := GenS2{F0: v3}
                    _ = v23
                }
                var v24 float32 = (1e10 / float32((1.5 / 1.0)))
                v4[""] = (v2 + (-(fi20)))
                _ = fi21
                _ = v22
                _ = v24
            }
            v1 = true
            _ = fi20
        }
        fmt.Println(v1, v2)
    }
    delete(v4, "hello")
    v1 = (v1 && (!false))
    for fi25 := 0; fi25 < 14; fi25++ {
        v1 = (2 > 65535)
        _ = fi25
    }
    v26 := GenS3{F0: (v1 && (("" + "\n\t") > "foo bar")), F1: -9223372036854775808}
    var v27 float64 = 1e10
    if (false || ((false && ("" >= "\n\t")) && v1)) {
        v28 := GenS4{F0: -1, F1: 2}
        _ = v28
    } else {
        v29 := GenS3{F0: (-9223372036854775808 != v2), F1: 42}
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        if (1 < 42) {
            var v31 uint = 715
            var v32 int32 = 732
            v33 := v4[("" + "a")]
            _ = v31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := []int{fi30, (fi34 & (v2 | 42)), (fi30 >> ((65535) & 63)), (fi34 & (fi34 ^ fi34))}
            v1 = v1
            fmt.Println(fi30)
            _ = fi34
            _ = v35
        }
        _ = fi30
    }
    var v36 uint8 = (v5 * (v5 % ((0) | 1)))
    if (!false) {
        for fi37 := 0; fi37 < 3; fi37++ {
            v5 = v5
            _ = fi37
        }
        v1 = v1
        v2 -= -9223372036854775808
    } else {
        var v38 uint = 255
        v27 -= v27
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
