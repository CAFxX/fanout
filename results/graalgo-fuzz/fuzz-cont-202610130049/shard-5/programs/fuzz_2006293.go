package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 bool }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint16; F1 bool; F2 int32 }
func genFn0() uint32 {
    var v6 int32 = -993
    if false {
        v7 := GenS1{F0: "a"}
        _ = v7
    } else {
        v8 := []int{128}
        _ = v8
    }
    v6 = (v6 * v6)
    _ = v6
    return 879
}

func main() {
    v1 := []uint32{928}
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    v4 := []int{}
    var v5 int32 = 946
    v9 := GenS1{F0: ""}
    var v10 uint8 = 39
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            switch ("a" + ("foo bar" + "a")) {
            case "foo bar":
                v3 = -1
            case "":
                v5 = 1
            }
        } else {
            fmt.Println(v3, v5, v10)
            v3 = -9223372036854775808
            if true {
                var v12 int16 = 1
                var v13 bool = true
                if (v12 <= -33) {
                    v14 := GenS1{F0: ("foo bar" + "foo bar")}
                    v3 = (^(v3))
                    _ = v14
                } else {
                    if v13 {
                        var v15 int64 = (-(v3))
                        v15 = (v3 ^ 283)
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 int = fi16
                            fmt.Println(fi16, fi11, v5)
                            switch "\n\t" {
                            case "foo bar":
                                v2[0] = fi16
                            }
                            v4 = append(v4, 996)
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                _ = v12
                _ = v13
            } else {
                v18 := GenS2{F0: 1, F1: 128, F2: -778}
                var v19 float32 = (((0.1 / 1.0) - (0.1 * (0.1 / (-2.25 + ((0.0 * 0.0) + 3.14159))))) / (3.14159 - float32(1.5)))
                v10 = v10
                _ = v18
                _ = v19
            }
        }
        _ = fi11
    }
    v20 := GenS3{F0: v5}
    v21 := v2[876]
    if false {
        var v22 int16 = -1
        _ = v22
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch "a" {
            case "\n\t":
                v24 := GenS4{F0: v10, F1: 4294967295}
                _ = v24
            case "a":
                if false {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v2[2147483648] = fi25
                        _ = fi25
                    }
                    v26 := GenS5{F0: 100}
                    var v27 int64 = (^(v3))
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(fi23, v5)
                    var v28 uint64 = 880
                    _ = v28
                }
            }
            v10 -= (v10 - (v10 >> ((42) & 63)))
            var v29 uint64 = 18446744073709551615
            _ = fi23
            _ = v29
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            if false {
                v10 = (v10 | v10)
                var v31 string = ""
                for fi32 := 0; fi32 < 15; fi32++ {
                    v3 += v3
                    var v33 int = fi32
                    if len(v1) > 0 {
                        v1[0] = 537
                    }
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
            _ = fi30
        }
        if ((389 >= 844) || false) {
            v34 := GenS3{F0: v5}
            _ = v34
        } else {
            v35 := GenS6{F0: -832, F1: (v5 ^ -138), F2: false}
            if (1 != 4294967295) {
                v21 = v21
                v3 = (v3 % ((-1) | 1))
                _ = len(v2)
            } else {
                var v36 uint = 18446744073709551615
                v21 += (-(v21))
                for fi37 := 0; fi37 < 12; fi37++ {
                    if false {
                        v36 -= v36
                        switch -784 {
                        case 32767:
                            v3 = v3
                        default:
                            var v38 uint16 = 32767
                            _ = v38
                        }
                    } else {
                        for fi39 := 0; fi39 < 17; fi39++ {
                            if (1 < 32767) {
                                fmt.Println(fi39, fi39, v3)
                            } else {
                                var v40 uint64 = 476
                                v41 := GenS7{F0: (v40 * uint64(v36)), F1: 1}
                                var v42 int16 = 0
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            if ((v36 % (((v36 % ((4294967295) | 1))) | 1)) <= v36) {
                                fmt.Println(v10)
                                v3 = v3
                                var v43 int16 = 0
                                _ = v43
                            } else {
                                v5 += (v5 << (((v36 << ((v36) & 63))) & 63))
                                var v44 int64 = (v3 | -565)
                                v10 += v10
                                _ = v44
                            }
                            _ = fi39
                        }
                        var v45 uint8 = v10
                        if (680 == -369) {
                            v2[2147483648] = fi37
                            var v46 uint16 = 32767
                            var v47 uint64 = 2
                            _ = v46
                            _ = v47
                        } else {
                            var v48 uint64 = 42
                            v1 = append(v1, uint32(v36))
                            _ = v48
                        }
                        _ = v45
                    }
                    _ = fi37
                }
                _ = v36
            }
            _ = v35
        }
    }
    delete(v2, genFn0())
    v10 = 2
    v10 += v10
    _ = len(v4)
    fmt.Println(v3, v5, v21)
    v21 = 128
    fmt.Println(v3, v3, v21)
    v10 = (v10 ^ v10)
    switch "a" {
    case "hello":
        var v49 uint8 = (^(v10))
        _ = v49
    }
    v3 -= v3
    if (genFn0() < 42) {
        if ((!(((3.14159 + 0.1) - (1.5 + ((float32((3.14159 + 0.0)) + (float32(1.5) - (1.5 - float32(float64(3.14159))))) - 1.5))) < -2.25)) || (255 >= 65535)) {
            var v50 uint = 255
            var v51 int64 = v3
            var v52 float32 = (3.14159 * float32(0.0))
            _ = v50
            _ = v51
            _ = v52
        } else {
            _ = len(v2)
        }
        v21 = (-(v21))
    } else {
        if (v21 > (v21 ^ v21)) {
            if (!(true && false)) {
                for fi53 := 0; fi53 < 9; fi53++ {
                    v54 := GenS1{F0: "foo bar"}
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v10)
                    default:
                        for fi55 := 0; fi55 < 4; fi55++ {
                            fmt.Println(v3, v10, v3)
                            if len(v1) > 0 {
                                v1[0] = genFn0()
                            }
                            if true {
                                if ((-1 <= 0) && (true && false)) {
                                    switch fi53 {
                                    case 2147483647:
                                        v21 += fi55
                                    case 179:
                                        v56 := GenS5{F0: -21}
                                        _ = v56
                                    default:
                                        v3 = 202
                                    }
                                    v57 := GenS8{F0: 255, F1: false, F2: (v5 << ((18446744073709551615) & 63))}
                                    _ = v57
                                }
                                fmt.Println(fi53, v21)
                            } else {
                                if false {
                                    var v58 int32 = -2147483648
                                    var v59 int8 = -128
                                    v60 := GenS1{F0: ("a" + "a")}
                                    _ = len(v1)
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                v10 -= (v10 % (((v10 ^ v10)) | 1))
                            }
                            _ = fi55
                        }
                    }
                    var v61 uint32 = 1
                    _ = fi53
                    _ = v54
                    _ = v61
                }
            }
        }
        v5 = (v5 / ((v5) | 1))
    }
    switch 391 {
    case 42:
        var v62 bool = false
        _ = v62
    case 1:
        fmt.Println(v3, v21)
    case 277:
        v10 = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
