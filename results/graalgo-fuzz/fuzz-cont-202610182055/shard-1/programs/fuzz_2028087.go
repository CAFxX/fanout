package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint, p1 int64) int32 {
    v7 := []int{}
    if false {
        v8 := GenS1{F0: 1, F1: false}
        _ = v8
    }
    var v9 uint64 = 299
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint32 = 42
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    v5 := make(map[int]bool)
    var v6 int64 = 9223372036854775807
    var v10 uint = 4294967295
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint = v10
        if (false && ((!(true || false)) || true)) {
            var v13 uint16 = 0
            for fi14 := 0; fi14 < 17; fi14++ {
                v10 = (v10 % ((uint(v2)) | 1))
                var v15 string = "\n\t"
                var v16 int32 = 1
                if false {
                    switch v15 {
                    case "hello":
                        switch 128 {
                        case 0:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            v13 = (v13 + 2)
                        }
                    case "a":
                        var v17 int = fi14
                        _ = v17
                    case "\n\t":
                        var v18 uint = (v12 / ((v12) | 1))
                        _ = v18
                    }
                    v12 -= 575
                    var v19 uint64 = 42
                    _ = v19
                } else {
                    var v20 int32 = (v16 % (((v16 / (((v16 >> ((v10) & 63))) | 1))) | 1))
                    v2 += (v2 * (v2 ^ v2))
                    if (((^(v2)) < 2) && (1e10 == -2.25)) {
                        switch v15 {
                        case "a":
                            var v21 string = (v15 + "hello")
                            switch v21 {
                            case "foo bar":
                                v22 := GenS2{F0: 0}
                                _ = v22
                            default:
                                if false {
                                    var v23 uint = 309
                                    _ = v23
                                }
                                var v24 int16 = 225
                                var v25 uint = v10
                                _ = v24
                                _ = v25
                            }
                            _ = v21
                        case "\n\t":
                            v10 -= v10
                        case "":
                            if ((true || (((v16 - v20) >= v16) && (!(((v13 - 59) < (v13 ^ uint16(v2))) && (!(fi11 <= 936)))))) || (!(!(0 == 1)))) {
                                var v26 int16 = 280
                                v16 -= genFn0(v10, (v6 - 0))
                                var v27 uint16 = 65535
                                _ = v26
                                _ = v27
                            }
                        default:
                            v28 := GenS2{F0: 2}
                            if false {
                                var v29 uint16 = (v13 | (^(v13)))
                                _ = v29
                            } else {
                                if true {
                                    v16 += (v20 * v16)
                                } else {
                                    var v30 uint = (v10 - v10)
                                    _ = v30
                                }
                            }
                            _ = v28
                        }
                        v15 = "a"
                        var v31 int32 = v16
                        _ = v31
                    } else {
                        v15 = v15
                        v13 -= 2
                        v32 := GenS3{F0: 0, F1: v6}
                        _ = v32
                    }
                    v33 := GenS2{F0: 2}
                    _ = v20
                    _ = v33
                }
                if true {
                    fmt.Println(v15, v4, v6)
                    if false {
                        var v34 float64 = (-2.25 / -2.25)
                        fmt.Println(v13, v15)
                        fmt.Println(v34, fi14)
                        _ = v34
                    } else {
                        v35 := GenS4{F0: int32(v6), F1: (^(v16)), F2: 14}
                        v6 += v6
                        _ = v35
                    }
                    v5[fi14] = (-2 == 120)
                    v1 = append(v1, v2)
                    var v36 bool = (uint(v13) < v12)
                    _ = v36
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        }
        var v37 uint32 = (v2 & v2)
        _ = fi11
        _ = v12
        _ = v37
    }
    if ((!false) && (!true)) {
        var v38 float32 = 1.5
        v10 = v10
        v4 = 0.1
        _ = v38
    }
    var v39 int32 = 42
    var v40 float64 = (-2.25 + float64(-2.25))
    fmt.Println(v6)
    v6 -= (v6 - v6)
    for fi41 := 0; fi41 < 7; fi41++ {
        v3["a"] = ""
        _ = fi41
    }
    _ = len(v5)
    var v42 float32 = 0.0
    switch (("\n\t" + ("" + "hello")) + "\n\t") {
    case "foo bar":
        if (v10 <= v10) {
            if false {
                v43 := GenS5{F0: -1}
                var v44 bool = false
                _ = v43
                _ = v44
            } else {
                var v45 bool = false
                v2 = v2
                if (86 <= 100) {
                    switch -249 {
                    case 1:
                        if true {
                            switch 884 {
                            case -230:
                                var v46 int16 = 42
                                _ = v46
                            case -167:
                                var v47 string = "a"
                                _ = v47
                            case 128:
                                var v48 uint64 = 0
                                _ = v48
                            }
                            fmt.Println(v39, v45)
                        }
                    case -9223372036854775808:
                        v6 = int64(v10)
                    case -65:
                        fmt.Println(v45)
                    default:
                        v49 := GenS5{F0: 9}
                        _ = v49
                    }
                    v2 -= v2
                    if v45 {
                        for fi50 := 0; fi50 < 8; fi50++ {
                            var v51 int8 = 100
                            v4 = ((v4 / 1.5) * 0.1)
                            _ = fi50
                            _ = v51
                        }
                        v52 := GenS2{F0: 25}
                        _ = v52
                    }
                } else {
                    var v53 int32 = genFn0(v10, v6)
                    _ = v53
                }
                _ = v45
            }
        } else {
            for fi54 := 0; fi54 < 15; fi54++ {
                v5[1] = (v40 == v40)
                for fi55 := 0; fi55 < 4; fi55++ {
                    fmt.Println(v2)
                    var v56 float32 = v42
                    var v57 int32 = 1
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                _ = fi54
            }
            var v58 float64 = float64(-2.25)
            var v59 int32 = (v39 ^ genFn0((v10 | 2), (v6 ^ -40)))
            _ = v58
            _ = v59
        }
    case "":
        fmt.Println(v40, v10, v39)
    case "hello":
        v60 := GenS3{F0: 32767, F1: v6}
        _ = v60
    }
    var v61 uint16 = 2
    if ("\n\t" == "\n\t") {
        if ((-579 >= -2147483648) && (!((((0.0 + (v42 - (float32((v40 / 1.0)) / (v4 * 1e10)))) - v4) - (float32(float64(v4)) * ((v42 * 1.5) / 0.1))) >= ((v42 + float32(v40)) - (1.5 / (float32(v40) / (float32(v40) * v42))))))) {
            if (uint16(v2) != v61) {
                v10 += (^(v10))
                v62 := GenS1{F0: uint32(v10), F1: (v40 <= v40)}
                fmt.Println(v61)
                _ = v62
            }
            if true {
                if (!false) {
                    var v63 int8 = -128
                    if true {
                        v64 := GenS6{F0: 2}
                        _ = v64
                    }
                    fmt.Println(v42, v10)
                    _ = v63
                }
                v5[4294967295] = false
            } else {
                v2 += v2
                _ = len(v5)
                var v65 uint64 = 2
                _ = v65
            }
            v10 = v10
        }
        fmt.Println(v39, v61)
    }
    for fi66 := 0; fi66 < 3; fi66++ {
        if true {
            _ = len(v3)
            v67 := v5[int(v39)]
            _ = v67
        } else {
            fmt.Println(v39, fi66, v39)
            v40 = v40
            v68 := v3["\n\t"]
            _ = v68
        }
        _ = fi66
    }
    var v69 uint = v10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v42 += 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
    _ = v61
    _ = v69
}
