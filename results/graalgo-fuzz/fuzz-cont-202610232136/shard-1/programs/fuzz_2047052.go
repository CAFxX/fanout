package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 string; F1 uint; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 bool }
func genFn0(p0 int16) int16 {
    switch 9223372036854775807 {
    case 1:
        var v5 int8 = 25
        _ = v5
    case 884:
        v6 := GenS1{F0: 65535, F1: 9223372036854775807, F2: "hello"}
        _ = v6
    default:
        p0 -= p0
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        p0 -= (p0 / ((p0) | 1))
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float64 = -2.25
            v10 := []int{(fi8 | fi7), (fi8 << ((65535) & 63)), 0}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(p0, p0)
        fmt.Println(p0)
        var v12 int = fi11
        _ = fi11
        _ = v12
    }
    v13 := []int{-2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v13
    return (^(p0))
}
func genFn1(p0 uint8, p1 uint32, p2 bool) int {
    for fi14 := 0; fi14 < 14; fi14++ {
        if (false || (!false)) {
            v15 := GenS2{F0: 693}
            _ = v15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := []int{}
            p2 = p2
            _ = fi17
            _ = v18
        }
        fmt.Println(fi16, p0, p2)
        _ = fi16
    }
    v19 := []int{}
    v19 = append(v19, 128)
    var v20 uint64 = 362
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 uint = 566
    v2 := make(map[uint32]string)
    var v3 uint = (v1 * v1)
    v4 := make(map[string]bool)
    v3 = v3
    var v21 uint32 = 1
    var v22 uint = (v3 >> ((374) & 63))
    v23 := GenS2{F0: -851}
    var v24 string = ("" + "\n\t")
    var v25 int64 = int64(v22)
    if true {
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 int64 = v25
            var v28 bool = false
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := []int{1, -319, -407, -20}
        switch 4294967295 {
        case -4:
            if false {
                var v30 uint32 = 145
                var v31 float32 = 1e10
                _ = v30
                _ = v31
            } else {
                var v32 bool = (!(9223372036854775807 == -646))
                fmt.Println(v24, v3)
                _ = v32
            }
        default:
            if true {
                switch "a" {
                case "foo bar":
                    v33 := GenS2{F0: -1}
                    _ = v33
                case "hello":
                    v34 := GenS3{F0: 2, F1: 42, F2: 2}
                    _ = v34
                }
                v1 += v1
            } else {
                v1 -= (v1 * (v1 ^ v3))
                for fi35 := 0; fi35 < 20; fi35++ {
                    v22 += (^(v3))
                    v29 = append(v29, 944)
                    var v36 int64 = 0
                    _ = fi35
                    _ = v36
                }
            }
        }
        _ = v29
    } else {
        var v37 uint = (v3 + (v22 / (((v22 / ((v22) | 1))) | 1)))
        _ = v37
    }
    v38 := GenS4{F0: v24, F1: v22, F2: v22}
    v39 := []int{1, 518}
    v39 = append(v39, -9223372036854775808)
    var v40 int8 = -1
    for fi41 := 0; fi41 < 16; fi41++ {
        v24 = ""
        v42 := GenS4{F0: ("hello" + v24), F1: 4294967295, F2: (v22 << ((v1) & 63))}
        var v43 uint64 = 4294967295
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println(v22)
    if true {
        var v44 bool = (false && false)
        _ = v44
    } else {
        var v45 uint64 = 128
        v4[(v24 + v24)] = (true && (1 == 2))
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := GenS5{F0: 127}
            if (1e10 > (1.5 * 0.0)) {
                v1 = 65535
                for fi48 := 0; fi48 < 12; fi48++ {
                    if (0.0 != 0.0) {
                        v40 = v40
                        fmt.Println(fi48)
                        v49 := GenS6{F0: 255, F1: v45}
                        _ = v49
                    } else {
                        fmt.Println(v40, v22, v25)
                        if len(v39) > 0 {
                            v39[0] = fi48
                        }
                        var v50 int16 = 32767
                        _ = v50
                    }
                    var v51 uint64 = v45
                    v52 := GenS7{F0: false}
                    _ = fi48
                    _ = v51
                    _ = v52
                }
            } else {
                var v53 int = (fi46 << (((^(v22))) & 63))
                _ = v53
            }
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
