package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint16; F2 uint }
func genFn0(p0 uint8) int64 {
    if true {
        p0 = p0
    } else {
        var v4 bool = (-32768 != 686)
        p0 = (^(p0))
        var v5 float64 = 1.5
        _ = v4
        _ = v5
    }
    if true {
        p0 = (p0 >> ((1) & 63))
        v6 := []int{-2147483648, 42, 1}
        if (!true) {
            var v7 uint32 = 2
            if (!(!((32767 >= 1000) && (!true)))) {
                v7 = (^(v7))
                v6 = append(v6, 128)
            }
            _ = v7
        }
        _ = v6
    }
    v8 := []int{127, 128, -493}
    v9 := GenS1{F0: 42, F1: 100, F2: (1 <= -32768)}
    v10 := GenS2{F0: 307}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (99 != 42)
    v3 := make(map[string]uint32)
    v11 := GenS2{F0: 2}
    if (float32((float64(3.14159) / 1e10)) < 0.1) {
        v2 = v2
        v2 = (-375 != 1000)
    } else {
        v12 := GenS3{F0: 1, F1: ""}
        v2 = v2
        _ = v12
    }
    var v13 string = "foo bar"
    fmt.Println(v13, v13)
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float64 = 3.14159
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(fi14, fi14, v2)
            if ((v2 || v2) && v2) {
                switch fi14 {
                case 32768:
                    if ((fi16 > (^(fi14))) || (v15 != v15)) {
                        fmt.Println(fi14, v2)
                        if (!(v15 < 1.5)) {
                            switch "" {
                            case "":
                                if ((0 == -1) && (366 == 1000)) {
                                    switch "" {
                                    case "":
                                        v17 := []int{}
                                        _ = v17
                                    case "a":
                                        switch (fi14 & (fi14 + fi14)) {
                                        case 9223372036854775807:
                                            var v18 string = (v13 + "\n\t")
                                            _ = v18
                                        }
                                    default:
                                        v19 := []int{}
                                        _ = v19
                                    }
                                }
                                v2 = (39 == -60)
                            case "hello":
                                fmt.Println(v15, v15, v2)
                            default:
                                if v2 {
                                    var v20 int32 = 2147483647
                                    if (!false) {
                                        v13 = ("a" + v13)
                                        var v21 int16 = 42
                                        _ = v21
                                    } else {
                                        v22 := GenS4{F0: 64}
                                        if (1 < -1) {
                                            var v23 uint = 0
                                            _ = v23
                                        }
                                        _ = v22
                                    }
                                    _ = v20
                                }
                            }
                        } else {
                            if v2 {
                                var v24 uint16 = 32767
                                _ = v24
                            } else {
                                switch fi16 {
                                case -229:
                                    var v25 int64 = 42
                                    _ = v25
                                default:
                                    v26 := GenS4{F0: 0}
                                    _ = v26
                                }
                            }
                            if true {
                                switch fi14 {
                                case 709:
                                    v13 = v13
                                }
                                v15 += (1e10 - 0.0)
                                v15 = (v15 + v15)
                            }
                        }
                        var v27 int16 = -32768
                        _ = v27
                    } else {
                        _ = len(v3)
                        if (527 < 286) {
                            var v28 bool = true
                            delete(v1, (fi14 >> ((495) & 63)))
                            if v28 {
                                v29 := []int{-9223372036854775808}
                                switch (fi14 / ((fi16) | 1)) {
                                case 2147483647:
                                    v30 := GenS3{F0: 1, F1: (v13 + ((("" + ("" + v13)) + (v13 + v13)) + v13))}
                                    _ = v30
                                case 32768:
                                    v2 = v2
                                }
                                _ = v29
                            }
                            _ = v28
                        } else {
                            v31 := GenS5{F0: 18446744073709551615, F1: 42}
                            v32 := GenS3{F0: 42, F1: v13}
                            _ = v31
                            _ = v32
                        }
                        v33 := GenS6{F0: 1}
                        _ = v33
                    }
                }
            } else {
                var v34 uint32 = 487
                v35 := GenS3{F0: 1, F1: v13}
                switch fi14 {
                case -2147483648:
                    v13 = v13
                }
                _ = v34
                _ = v35
            }
            if (genFn0(207) >= -420) {
                if v2 {
                    v36 := GenS3{F0: -1, F1: v13}
                    var v37 uint8 = 42
                    _ = v36
                    _ = v37
                }
                v38 := v1[(-(fi16))]
                if ((!(1 != 1000)) || (!("a" <= v13))) {
                    v13 = (v13 + (v13 + v13))
                    v15 = v15
                    v38 -= (v38 - v38)
                    v15 = ((v15 * (v15 * float64(3.14159))) - v15)
                    v2 = (982 >= 18446744073709551615)
                    switch 42 {
                    case 300:
                        if true {
                            var v39 int32 = 0
                            _ = v39
                        }
                    case 1:
                        if ((!(!v2)) || v2) {
                            v15 = v15
                            v2 = (false || false)
                            v15 -= v15
                            var v40 uint32 = v38
                            v41 := GenS7{F0: -1, F1: 586}
                            _ = v40
                            _ = v41
                        }
                    }
                } else {
                    fmt.Println(fi16, v38, fi16)
                    v38 = v38
                }
                v42 := GenS8{F0: -26, F1: v13}
                _ = v38
                _ = v42
            } else {
                v2 = (v2 || (true && (4294967295 >= 593)))
                switch fi14 {
                case -677:
                    v15 += ((v15 - (v15 + ((v15 - (((v15 + (v15 * v15)) + (v15 * (1.5 * 1.5))) + (v15 - 3.14159))) * 3.14159))) / v15)
                case -924:
                    fmt.Println(v15, v2)
                default:
                    v15 += (float64(-2.25) + v15)
                }
                switch ((((v13 + "foo bar") + v13) + "foo bar") + v13) {
                case "":
                    var v43 uint16 = 42
                    _ = v43
                case "a":
                    var v44 uint32 = 4294967295
                    _ = v44
                }
            }
            _ = fi16
        }
        if (!(-588 == 0)) {
            for fi45 := 0; fi45 < 15; fi45++ {
                var v46 uint64 = 702
                _ = fi45
                _ = v46
            }
            v47 := GenS9{F0: 18446744073709551615}
            var v48 float64 = 0.1
            _ = v47
            _ = v48
        } else {
            if (69 <= 2) {
                v49 := GenS3{F0: 380, F1: v13}
                var v50 int16 = 1000
                _ = v49
                _ = v50
            }
            var v51 uint = 1
            _ = v51
        }
        _ = fi14
        _ = v15
    }
    var v52 uint64 = 2
    v53 := []int{}
    if len(v53) >= 2 {
        v53 = v53[1:len(v53)]
    }
    if ((128 <= 2) || v2) {
        var v54 string = "hello"
        v55 := v3["\n\t"]
        v56 := GenS3{F0: 951, F1: "a"}
        _ = v54
        _ = v55
        _ = v56
    }
    fmt.Println(v13)
    for fi57 := 0; fi57 < 14; fi57++ {
        var v58 int8 = 127
        _ = fi57
        _ = v58
    }
    var v59 uint = 0
    var v60 int8 = int8(v59)
    var v61 float32 = float32((0.0 * 0.1))
    v62 := GenS10{F0: 0, F1: 42, F2: (v59 / ((v59) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
