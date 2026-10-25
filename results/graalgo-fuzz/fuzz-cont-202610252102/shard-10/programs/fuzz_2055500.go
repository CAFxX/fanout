package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int64; F1 int64; F2 int }
type GenS4 struct { F0 uint; F1 string }
func genFn0(p0 bool, p1 string, p2 int64) int8 {
    v6 := []int{828}
    var v7 uint16 = 568
    p1 = p1
    v8 := GenS1{F0: 515, F1: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 bool = false
    var v2 int = 127
    v3 := []int{(^(v2))}
    v4 := make(map[uint32]int)
    var v5 string = "a"
    var v9 float32 = (1.5 - -2.25)
    var v10 uint64 = 571
    var v11 int8 = genFn0(false, "foo bar", 1)
    v4[758] = (v2 >> ((18446744073709551615) & 63))
    v12 := v4[1]
    _ = len(v3)
    v13 := GenS2{F0: 42, F1: v1}
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v1, v2)
            v10 = v10
            switch v2 {
            case -1:
                v16 := GenS3{F0: -1, F1: 605, F2: fi15}
                _ = v16
            case -239:
                switch v12 {
                case 9223372036854775807:
                    if v1 {
                        v17 := v4[793]
                        switch fi15 {
                        case 1:
                            v18 := v4[61]
                            _ = v18
                        case -1:
                            if true {
                                if (((v11 & genFn0(((v10 ^ 2) <= v10), v5, -9223372036854775808)) >= -128) || (2 == 42)) {
                                    v19 := v4[1]
                                    _ = v19
                                }
                                fmt.Println(v1, v1, v9)
                                v10 = v10
                            }
                        }
                        _ = v17
                    } else {
                        var v20 bool = (v1 && (953 >= (v10 / ((2) | 1))))
                        delete(v4, 65535)
                        var v21 uint32 = 560
                        _ = v20
                        _ = v21
                    }
                    v4[1] = fi15
                    v12 += (^(fi15))
                case -768:
                    v1 = (v9 > v9)
                case 0:
                    delete(v4, 0)
                default:
                    v2 -= fi15
                    var v22 float64 = 0.0
                    v5 = ((v5 + v5) + "")
                    _ = v22
                }
            case 32768:
                switch ("\n\t" + "a") {
                case "":
                    fmt.Println(v12, v2)
                case "hello":
                    var v23 int32 = 195
                    _ = v23
                }
            default:
                v24 := GenS3{F0: 0, F1: -9223372036854775808, F2: v12}
                _ = v24
            }
            var v25 int64 = 0
            _ = fi15
            _ = v25
        }
        _ = fi14
    }
    fmt.Println(v10, v5, v1)
    var v26 uint64 = (v10 & (v10 * v10))
    var v27 float64 = 1e10
    v1 = v1
    var v28 uint32 = 262
    v9 = v9
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS3{F0: 921, F1: -9223372036854775808, F2: (v12 + -739)}
            v32 := GenS2{F0: 243, F1: v1}
            if (2 < uint8(v28)) {
                if true {
                    fmt.Println(fi30, v27)
                }
                v28 += (^(v28))
                fmt.Println(v27, fi30)
                var v33 uint32 = (^(v28))
                if (!(854 > 1)) {
                    if (!(42 < 267)) {
                        if v1 {
                            var v34 uint = 1
                            switch fi30 {
                            case 32768:
                                v28 -= v33
                            default:
                                if v1 {
                                    var v35 uint = v34
                                    _ = v35
                                } else {
                                    v11 = (v11 & v11)
                                }
                            }
                            v9 = v9
                            _ = v34
                        } else {
                            fmt.Println(v1, v1, v10)
                            var v36 int32 = -2147483648
                            var v37 float64 = (((0.0 * 1e10) * v27) - float64(((v9 * 0.1) / v9)))
                            _ = v36
                            _ = v37
                        }
                        v28 -= v28
                        delete(v4, uint32(v10))
                    }
                    switch (v5 + "a") {
                    case "foo bar":
                        var v38 uint16 = 65535
                        _ = v38
                    default:
                        switch (v5 + (v5 + (v5 + v5))) {
                        case "":
                            var v39 int32 = 42
                            _ = v39
                        case "a":
                            fmt.Println(fi29, v33)
                        case "hello":
                            var v40 uint8 = 1
                            _ = v40
                        default:
                            if (1000 < 104) {
                                v9 -= (v9 - v9)
                            } else {
                                v33 -= v33
                            }
                        }
                    }
                }
                v28 = (v28 >> ((1) & 63))
                _ = v33
            } else {
                if v1 {
                    switch v5 {
                    case "foo bar":
                        v28 = (v28 - (^(v28)))
                    }
                    var v41 int8 = v11
                    fmt.Println(v2, v26, v41)
                    if ((!v1) || v1) {
                        v27 = (1.5 - 3.14159)
                    }
                    if ((false && (255 > 437)) && v1) {
                        v10 = v26
                    }
                    _ = v41
                } else {
                    v10 += v10
                    v42 := GenS4{F0: 1, F1: (v5 + v5)}
                    _ = v42
                }
                v12 = (v2 % ((v12) | 1))
                var v43 uint8 = 42
                _ = v43
            }
            v28 -= v28
            var v44 int8 = (v11 % ((v11) | 1))
            v45 := GenS3{F0: -9223372036854775808, F1: -44, F2: 269}
            _ = fi30
            _ = v31
            _ = v32
            _ = v44
            _ = v45
        }
        v46 := GenS4{F0: 93, F1: v5}
        if ((!v1) && v1) {
            var v47 bool = (1 > (^(v11)))
            fmt.Println(v9, v1, v47)
            _ = v47
        } else {
            for fi48 := 0; fi48 < 3; fi48++ {
                v49 := GenS4{F0: 458, F1: (v5 + v5)}
                _ = fi48
                _ = v49
            }
            for fi50 := 0; fi50 < 13; fi50++ {
                var v51 uint64 = v10
                fmt.Println(v26)
                if false {
                    if len(v3) > 0 {
                        v3[0] = (v2 | 32768)
                    }
                }
                fmt.Println(v11, v10, fi29)
                v9 = v9
                _ = fi50
                _ = v51
            }
            v1 = (0 > 0)
        }
        _ = fi29
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
