package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint; F2 bool }
type GenS6 struct { F0 int; F1 int16; F2 uint64 }
type GenS7 struct { F0 int16 }
func genFn0() int8 {
    var v7 uint32 = 2
    var v8 int8 = 100
    var v9 int16 = -874
    var v10 int8 = v8
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (758 < 65535)
    v3 := []int{-878, 127, -611}
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    var v6 int16 = -32768
    v2 = false
    v1[4294967295] = 4294967295
    var v11 uint8 = 42
    var v12 float64 = 1.5
    v13 := GenS1{F0: -1, F1: 42, F2: 0}
    v2 = true
    var v14 uint32 = 0
    delete(v1, (^(v14)))
    v1[2147483648] = v14
    var v15 uint32 = (v14 + v14)
    fmt.Println(v14, v14)
    var v16 int64 = -608
    v12 -= (-2.25 / v12)
    v11 = (v11 % (((^(v11))) | 1))
    if v2 {
        var v17 int64 = -902
        if (!v2) {
            var v18 uint = 65535
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v6, v17, v12)
                v20 := GenS2{F0: v2, F1: genFn0()}
                _ = fi19
                _ = v20
            }
            switch -161 {
            case 9223372036854775807:
                v3 = append(v3, -1)
            case 1:
                for fi21 := 0; fi21 < 17; fi21++ {
                    if true {
                        for fi22 := 0; fi22 < 13; fi22++ {
                            _ = len(v5)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi22
                        }
                        v23 := GenS3{F0: (v16 ^ (v16 | (v17 * (-(v16))))), F1: (v16 | v17)}
                        v18 -= uint(v11)
                        _ = v23
                    }
                    if false {
                        v18 = 42
                    }
                    _ = fi21
                }
            case 389:
                v24 := GenS1{F0: 42, F1: 912, F2: 2147483647}
                _ = v24
            default:
                if (v2 && (v2 || v2)) {
                    var v25 uint64 = 895
                    _ = v25
                }
            }
            _ = v18
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 bool = (v2 || (!(v2 || false)))
            _ = fi26
            _ = v27
        }
        _ = v17
    } else {
        var v28 bool = (v14 < 2147483647)
        v15 = v14
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        switch 128 {
        case -1:
            if v2 {
                var v30 uint64 = 42
                _ = v30
            }
        case 4294967295:
            v31 := GenS4{F0: fi29}
            _ = v31
        default:
            switch ("hello" + ("\n\t" + "\n\t")) {
            case "":
                if true {
                    v32 := GenS5{F0: 1, F1: 4294967295, F2: false}
                    _ = v32
                } else {
                    if v2 {
                        v6 = v6
                    }
                    var v33 bool = ((345 != 240) && ("\n\t" != "hello"))
                    _ = v33
                }
            case "foo bar":
                for fi34 := 0; fi34 < 6; fi34++ {
                    var v35 int16 = -291
                    _ = fi34
                    _ = v35
                }
            }
        }
        if v2 {
            if v2 {
                v36 := GenS2{F0: v2, F1: -128}
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 uint16 = 0
                    v39 := GenS6{F0: fi37, F1: v6, F2: uint64(v38)}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = v36
            } else {
                v11 = 255
                v6 -= v6
            }
            v2 = ((v15 >= (v14 | (^(v15)))) || (18446744073709551615 == 1))
        }
        for fi40 := 0; fi40 < 8; fi40++ {
            v11 -= v11
            v41 := GenS5{F0: 1, F1: 757, F2: (100 < -85)}
            _ = fi40
            _ = v41
        }
        _ = fi29
    }
    switch 127 {
    case -845:
        v16 -= v16
    case 716:
        var v42 int16 = v6
        _ = v42
    case 0:
        v43 := GenS7{F0: 0}
        _ = v43
    default:
        for fi44 := 0; fi44 < 7; fi44++ {
            v12 = (1e10 / (v12 * float64((0.1 / ((3.14159 * 0.0) * ((1.5 * (0.1 * (1.5 * ((-2.25 + ((float32(float64((-2.25 + (-2.25 * float32(v12))))) * (1e10 + 0.1)) * -2.25)) - (0.1 * (-2.25 + 0.0)))))) - (0.0 - 0.1)))))))
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 float32 = 1e10
                _ = fi45
                _ = v46
            }
            var v47 int16 = v6
            _ = fi44
            _ = v47
        }
    }
    var v48 int = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
