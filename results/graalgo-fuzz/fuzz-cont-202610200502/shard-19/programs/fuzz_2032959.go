package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0(p0 string) int64 {
    var v4 uint8 = 128
    switch -141 {
    case 108:
        var v5 uint = 38
        _ = v5
    default:
        v6 := GenS1{F0: 1, F1: -669, F2: -2147483648}
        _ = v6
    }
    var v7 int = 0
    switch "" {
    case "":
        var v8 int = (v7 & int(v4))
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 uint64) int8 {
    v9 := GenS1{F0: genFn0("\n\t"), F1: 42, F2: 0}
    p0 = (p0 / (((p0 << ((255) & 63))) | 1))
    p0 = (p0 & 2)
    _ = p0
    _ = v9
    return 127
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]bool)
    v3 := []int{32768, -470, 127, 0}
    switch 830 {
    case 0:
        if true {
            if false {
                v10 := v2[42]
                var v11 int32 = 1
                v12 := GenS2{F0: v10}
                _ = v10
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch fi13 {
                    case 9223372036854775807:
                        v1 += (-(v1))
                    case 4294967295:
                        v2[(fi13 + (fi13 - 127))] = false
                    }
                    if false {
                        _ = len(v3)
                        if len(v3) > 0 {
                            v3[0] = fi13
                        }
                        v1 -= -1
                    }
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS2{F0: true}
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
                var v16 uint32 = 2147483647
                for fi17 := 0; fi17 < 2; fi17++ {
                    v16 -= (v16 % (((v16 | v16)) | 1))
                    var v18 int = (^(fi17))
                    switch -2147483648 {
                    case -1:
                        var v19 uint32 = v16
                        _ = v19
                    case 2147483647:
                        if len(v3) > 0 {
                            v3[0] = 578
                        }
                    case 759:
                        var v20 int8 = (v1 >> ((0) & 63))
                        _ = v20
                    default:
                        switch fi17 {
                        case -9223372036854775808:
                            _ = len(v3)
                        case 0:
                            var v21 uint = 271
                            _ = v21
                        case 9223372036854775807:
                            for fi22 := 0; fi22 < 3; fi22++ {
                                if len(v3) > 0 {
                                    v3[0] = v18
                                }
                                v18 += (v18 + -1)
                                if (true || (0 < 42)) {
                                    _ = len(v3)
                                    v23 := v2[fi17]
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    _ = v23
                                } else {
                                    if (255 != 255) {
                                        var v24 float64 = -2.25
                                        _ = v24
                                    }
                                }
                                _ = fi22
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            if (128 > 1) {
                var v25 uint16 = 65535
                fmt.Println(v25, v1)
                _ = v25
            } else {
                v1 = (-(v1))
                for fi26 := 0; fi26 < 18; fi26++ {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v28 := GenS3{F0: (696 == 209)}
                        var v29 float64 = (3.14159 + (0.1 + float64(0.0)))
                        fmt.Println(fi26, v29, v1)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    var v30 int16 = 0
                    _ = fi26
                    _ = v30
                }
            }
            v31 := GenS4{F0: "foo bar"}
            _ = v31
        } else {
            var v32 uint32 = 89
            var v33 string = ("" + "a")
            _ = v32
            _ = v33
        }
    case 127:
        for fi34 := 0; fi34 < 13; fi34++ {
            fmt.Println(fi34)
            var v35 uint = 42
            _ = fi34
            _ = v35
        }
    default:
        fmt.Println(v1)
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 127
        }
    }
    v36 := GenS5{F0: "hello"}
    v37 := v2[223]
    var v38 string = ("hello" + "hello")
    _ = len(v3)
    var v39 uint8 = 128
    if (!v37) {
        if true {
            var v40 float32 = 0.1
            fmt.Println(v1, v39, v37)
            _ = v40
        }
    } else {
        v41 := GenS1{F0: -9223372036854775808, F1: -327, F2: 0}
        v42 := GenS6{F0: (v1 ^ v1), F1: 0}
        var v43 bool = false
        _ = v41
        _ = v42
        _ = v43
    }
    var v44 float32 = 1.5
    var v45 int16 = -717
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v44
    _ = v45
}
