package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int }
type GenS5 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0() uint16 {
    v7 := GenS1{F0: 127, F1: -9223372036854775808, F2: 333}
    var v8 uint16 = 2
    fmt.Println(v8)
    v9 := []int{-1, -207}
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 string = (("\n\t" + "hello") + "\n\t")
    var v2 uint8 = 24
    var v3 uint64 = 282
    var v4 float64 = 0.1
    v5 := []string{("foo bar" + "hello"), "foo bar"}
    v6 := make(map[string]int)
    var v10 uint8 = v2
    if true {
        if false {
            v2 = v2
            if true {
                fmt.Println(v4, v10)
            }
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v2)
                var v12 int32 = 42
                _ = fi11
                _ = v12
            }
        }
        var v13 float64 = ((v4 / v4) * ((v4 + (1.5 - float64(1e10))) * ((((v4 / v4) - 1.5) / v4) * v4)))
        var v14 uint32 = 42
        _ = v13
        _ = v14
    } else {
        fmt.Println(v10, v2, v1)
    }
    v15 := v6[(((v1 + "foo bar") + v1) + v1)]
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 string = (v1 + v1)
        v17 = v17
        v18 := GenS1{F0: fi16, F1: -468, F2: 552}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 42, F1: 1, F2: 65535}
    switch "hello" {
    case "a":
        v10 = v2
    case "foo bar":
        v4 = ((v4 - v4) / v4)
    default:
        v5 = append(v5, v1)
    }
    var v20 int16 = 479
    v21 := GenS3{F0: -613, F1: 1}
    _ = len(v5)
    for fi22 := 0; fi22 < 6; fi22++ {
        if (v1 < v1) {
            fmt.Println(v15)
            switch v15 {
            case -2147483648:
                _ = len(v5)
            case -781:
                fmt.Println(v2)
            case 32768:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        } else {
            switch 127 {
            case 683:
                switch fi22 {
                case -1:
                    if (v2 < 255) {
                        fmt.Println(v15, v3)
                        v23 := GenS1{F0: v15, F1: (-(v15)), F2: 612}
                        _ = v23
                    }
                case 1:
                    var v24 int64 = 9223372036854775807
                    _ = v24
                }
            case 680:
                fmt.Println(v2)
            case -2147483648:
                v4 += v4
            default:
                v20 -= 450
            }
            var v25 float64 = v4
            _ = v25
        }
        _ = fi22
    }
    v3 = 4294967295
    fmt.Println(v4, v20, v2)
    if (((genFn0() == 255) || ((uint(v10) > 65535) || false)) && ((v2 & v2) > v10)) {
        _ = len(v6)
        var v26 int = 712
        switch (v15 & 32767) {
        case 0:
            var v27 uint64 = (^(v3))
            _ = v27
        default:
            switch v26 {
            case 316:
                delete(v6, (("a" + v1) + v1))
            case -1:
                if (v15 < v15) {
                    fmt.Println(v2)
                } else {
                    var v28 uint32 = 65535
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v20 -= (v20 >> ((2) & 63))
                        _ = fi29
                    }
                    _ = v28
                }
            default:
                fmt.Println(v2, v3, v10)
            }
        }
        _ = v26
    }
    v2 -= (v10 + (v2 / ((128) | 1)))
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 uint = 1
        for fi32 := 0; fi32 < 4; fi32++ {
            if (((138 >= uint8(v3)) || false) || false) {
                if (0 == 2) {
                    v1 = ("\n\t" + v1)
                    var v33 float64 = v4
                    _ = v33
                } else {
                    var v34 uint16 = 65535
                    v5 = append(v5, "\n\t")
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v34
                }
                v35 := v6["\n\t"]
                v36 := GenS4{F0: (v31 ^ 4294967295), F1: 49, F2: 170}
                v37 := GenS5{F0: v20, F1: false, F2: uint16(v31)}
                _ = v35
                _ = v36
                _ = v37
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v38 int = fi32
            _ = fi32
            _ = v38
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            var v40 uint16 = 65535
            v41 := GenS5{F0: v20, F1: false, F2: genFn0()}
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
