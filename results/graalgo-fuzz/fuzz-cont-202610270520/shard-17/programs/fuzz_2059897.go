package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 int32, p1 int16) int64 {
    if (255 != 1) {
        var v6 int = 32768
        if ((0.1 * (0.1 - 1.5)) != 1e10) {
            v7 := []int{(v6 << ((216) & 63))}
            var v8 float32 = (float32((0.0 - float64(1.5))) / (((1e10 * (1.5 * float32(0.1))) - (float32(((0.1 * (-2.25 / float64((-2.25 / 1.0)))) * (float64(1e10) * 0.0))) / 1.5)) + 1.5))
            _ = v7
            _ = v8
        }
        _ = v6
    }
    var v9 int8 = -1
    v10 := GenS1{F0: 802}
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint = 2
        var v13 uint16 = 814
        var v14 string = ""
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if false {
        var v15 uint = 1
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -10
}
func genFn1(p0 int) uint64 {
    for fi16 := 0; fi16 < 3; fi16++ {
        switch "a" {
        case "foo bar":
            p0 = (p0 & (fi16 * (^(p0))))
        default:
            switch (p0 % ((649) | 1)) {
            case 2147483647:
                v17 := GenS2{F0: "", F1: 42}
                _ = v17
            case 127:
                if true {
                    var v18 int8 = 118
                    fmt.Println(fi16)
                    _ = v18
                }
            case 0:
                fmt.Println(p0, p0)
            }
        }
        _ = fi16
    }
    var v19 uint32 = 4294967295
    p0 += -1
    _ = p0
    _ = v19
    return 2
}

func main() {
    var v1 int = 0
    v2 := []uint32{}
    var v3 uint = 212
    v4 := make(map[string]uint32)
    var v5 float64 = float64(1e10)
    var v20 float64 = v5
    if (true || false) {
        var v21 int8 = -72
        v3 += 255
        if ((v3 | (v3 ^ (v3 | v3))) <= v3) {
            for fi22 := 0; fi22 < 19; fi22++ {
                switch "" {
                case "foo bar":
                    var v23 uint = 2
                    _ = v23
                case "hello":
                    fmt.Println(v1, v20)
                default:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        if (!(128 == 2)) {
                            v25 := GenS1{F0: 1}
                            v26 := v4["a"]
                            var v27 string = ("" + "\n\t")
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            if ((false && (0 >= 4294967295)) && (-276 != v1)) {
                                var v28 int32 = -1
                                fmt.Println(v28, v20, v3)
                                v28 = (v28 - -2147483648)
                                _ = v28
                            } else {
                                switch fi22 {
                                case 974:
                                    fmt.Println(v5)
                                case -451:
                                    v20 = v5
                                case 1:
                                    v2 = append(v2, 2)
                                default:
                                    if false {
                                        if (125 >= 836) {
                                            var v29 int16 = 224
                                            _ = len(v2)
                                            _ = v29
                                        } else {
                                            v3 = v3
                                        }
                                        v30 := GenS3{F0: 128, F1: "\n\t", F2: (^(v21))}
                                        var v31 uint64 = 778
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if false {
                                            var v32 uint32 = 1
                                            if len(v2) > 0 {
                                                v2[0] = v32
                                            }
                                            _ = v32
                                        }
                                        if ("" != "") {
                                            var v33 uint16 = 42
                                            v2 = append(v2, 2147483648)
                                            _ = v33
                                        }
                                    }
                                }
                                fmt.Println(v21)
                            }
                            switch "foo bar" {
                            case "foo bar":
                                var v34 uint16 = 32767
                                _ = v34
                            case "a":
                                if (0 < 0) {
                                    v5 = float64((0.1 / float32(v5)))
                                    v3 -= (v3 | 18446744073709551615)
                                }
                            case "\n\t":
                                fmt.Println(v20, fi22)
                            default:
                                _ = len(v2)
                            }
                        }
                        v35 := v4["foo bar"]
                        var v36 int32 = 1
                        _ = fi24
                        _ = v35
                        _ = v36
                    }
                }
                v37 := GenS3{F0: 3, F1: "\n\t", F2: v21}
                _ = fi22
                _ = v37
            }
            fmt.Println(v3, v5, v20)
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(v21)
                _ = fi38
            }
        } else {
            var v39 uint8 = 127
            v40 := GenS4{F0: v3}
            _ = v39
            _ = v40
        }
        _ = v21
    } else {
        var v41 uint32 = 347
        var v42 bool = true
        _ = v41
        _ = v42
    }
    var v43 float64 = 1e10
    fmt.Println(v43, v20, v5)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5, v1, v20)
    case "":
        v44 := GenS5{F0: genFn1(32768), F1: -2147483648}
        _ = v44
    }
    var v45 uint8 = 42
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS1{F0: 1}
        v5 -= v43
        var v48 int8 = -128
        _ = fi46
        _ = v47
        _ = v48
    }
    v3 += (v3 ^ (v3 * v3))
    if len(v2) > 0 {
        v2[0] = 262
    }
    v20 = float64((((1.5 / 1.0) / ((0.0 - 0.0) - 1.5)) / (1e10 - -2.25)))
    for fi49 := 0; fi49 < 1; fi49++ {
        if len(v2) > 0 {
            v2[0] = uint32(v3)
        }
        _ = fi49
    }
    var v50 float32 = 1.5
    v5 = (v43 / ((v5 * v20) / v43))
    if false {
        if ("\n\t" != "foo bar") {
            v43 += v5
            for fi51 := 0; fi51 < 3; fi51++ {
                var v52 uint = v3
                fmt.Println(v50, v45, v45)
                v43 = v5
                _ = fi51
                _ = v52
            }
            var v53 int8 = -79
            _ = v53
        }
        var v54 bool = (((425 != uint64(v3)) && false) || (true || (42 >= int32(v3))))
        _ = v54
    }
    var v55 float64 = ((v5 * v43) - 1e10)
    v56 := GenS5{F0: genFn1((v1 & (^(v1)))), F1: 4294967295}
    fmt.Println(v43, v55, v45)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v50
    _ = v55
    _ = v56
}
