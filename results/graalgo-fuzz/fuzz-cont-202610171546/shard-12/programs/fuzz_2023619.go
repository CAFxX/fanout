package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int16; F2 int32 }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
func genFn0(p0 uint) int32 {
    var v6 uint32 = 4294967295
    v7 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 1}
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint8, p1 uint) uint16 {
    v8 := GenS2{F0: 2147483648}
    p1 = 4294967295
    if (false && (!true)) {
        switch 0 {
        case 127:
            v9 := GenS3{F0: 475}
            _ = v9
        case 240:
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi10}
                var v12 uint64 = 1
                v11 = append(v11, fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
        case -2147483648:
            v13 := GenS3{F0: 2147483647}
            _ = v13
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return 466
}

func main() {
    v1 := []uint32{}
    v2 := []string{(((("foo bar" + "foo bar") + "a") + ("a" + ((("a" + (((("\n\t" + "foo bar") + ("hello" + "")) + "") + "hello")) + "\n\t") + "\n\t"))) + "a"), (("foo bar" + "foo bar") + "\n\t"), (("hello" + "hello") + "\n\t")}
    v3 := make(map[uint32]uint32)
    var v4 bool = (0.0 == (1.5 * 1.5))
    v5 := make(map[int]int)
    switch "foo bar" {
    case "a":
        fmt.Println(v4)
    case "hello":
        v14 := GenS1{F0: -1, F1: 42, F2: -1}
        _ = v14
    }
    v4 = (148 < 42)
    v4 = (0 < 0)
    switch (("\n\t" + "hello") + "\n\t") {
    case "":
        var v15 uint64 = 888
        _ = v15
    default:
        _ = len(v5)
    }
    if (1e10 >= ((0.0 + 3.14159) / ((float64(3.14159) + 1e10) + 0.1))) {
        var v16 float32 = float32(1.5)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "hello" {
        case "foo bar":
            if (2 < 255) {
                v16 -= 0.0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v4, v16)
            } else {
                v16 -= 0.1
                var v17 uint = 1
                if true {
                    var v18 uint = (^(v17))
                    for fi19 := 0; fi19 < 7; fi19++ {
                        if (0 >= 2) {
                            var v20 int8 = -1
                            _ = v20
                        } else {
                            v21 := GenS4{F0: (^(fi19)), F1: -306, F2: 42}
                            _ = v21
                        }
                        switch (fi19 << ((v17) & 63)) {
                        case 9223372036854775807:
                            v22 := GenS3{F0: genFn0(v17)}
                            _ = v22
                        case 128:
                            var v23 uint8 = 127
                            _ = v23
                        case 2147483647:
                            fmt.Println(v16)
                        default:
                            v18 -= (v17 + 4294967295)
                        }
                        var v24 uint32 = 0
                        _ = fi19
                        _ = v24
                    }
                    v25 := v5[42]
                    _ = v18
                    _ = v25
                } else {
                    v17 = v17
                    if ((-2.25 / (float64((1e10 + (v16 - 1.5))) / 1.5)) > 1e10) {
                        v17 = (v17 | v17)
                    } else {
                        if (-713 == -1) {
                            switch 355 {
                            case 128:
                                if len(v2) > 0 {
                                    v2[0] = (("foo bar" + "\n\t") + ("" + "foo bar"))
                                }
                            case 127:
                                if v4 {
                                    for fi26 := 0; fi26 < 15; fi26++ {
                                        var v27 int8 = 100
                                        v4 = (!v4)
                                        switch "\n\t" {
                                        case "foo bar":
                                            v28 := GenS4{F0: fi26, F1: 32767, F2: -109}
                                            _ = v28
                                        case "a":
                                            switch -107 {
                                            case -135:
                                                fmt.Println(v17)
                                            default:
                                                if true {
                                                    v5[fi26] = fi26
                                                }
                                            }
                                        case "hello":
                                            var v29 int = -813
                                            _ = v29
                                        default:
                                            if v4 {
                                                for fi30 := 0; fi30 < 2; fi30++ {
                                                    v17 += 0
                                                    v27 = v27
                                                    v27 -= v27
                                                    _ = fi30
                                                }
                                                v31 := GenS3{F0: 0}
                                                if v4 {
                                                    v4 = true
                                                    var v32 uint16 = 669
                                                    _ = v32
                                                } else {
                                                    fmt.Println(v16, fi26)
                                                }
                                                _ = v31
                                            }
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                    if (-1 >= 127) {
                                        var v33 uint16 = 2
                                        _ = v33
                                    } else {
                                        switch 1 {
                                        case -209:
                                            v17 -= (v17 % ((846) | 1))
                                        case 4294967295:
                                            var v34 uint = v17
                                            _ = v34
                                        default:
                                            v17 = v17
                                        }
                                        var v35 uint = (v17 / (((v17 & v17)) | 1))
                                        v36 := GenS4{F0: 2147483647, F1: 0, F2: 1}
                                        _ = v35
                                        _ = v36
                                    }
                                    for fi37 := 0; fi37 < 4; fi37++ {
                                        v17 -= 982
                                        for fi38 := 0; fi38 < 13; fi38++ {
                                            v39 := GenS4{F0: (fi37 & fi37), F1: 1, F2: 0}
                                            v17 -= v17
                                            _ = fi38
                                            _ = v39
                                        }
                                        _ = fi37
                                    }
                                }
                            case 661:
                                var v40 float32 = v16
                                _ = v40
                            default:
                                fmt.Println(v16, v16)
                            }
                            for fi41 := 0; fi41 < 3; fi41++ {
                                for fi42 := 0; fi42 < 5; fi42++ {
                                    var v43 int8 = 0
                                    fmt.Println(fi41)
                                    _ = fi42
                                    _ = v43
                                }
                                _ = fi41
                            }
                        }
                        for fi44 := 0; fi44 < 14; fi44++ {
                            fmt.Println(v16, v4, v4)
                            _ = fi44
                        }
                        v17 -= (^(v17))
                    }
                }
                _ = v17
            }
        default:
            var v45 float64 = float64(v16)
            _ = v45
        }
        _ = v16
    }
    fmt.Println(v4)
    for fi46 := 0; fi46 < 10; fi46++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        v47 := GenS5{F0: 1000, F1: "foo bar", F2: -2147483648}
        _ = fi46
        _ = v47
    }
    switch "hello" {
    case "":
        v4 = false
    default:
        var v48 int32 = 1
        _ = v48
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
