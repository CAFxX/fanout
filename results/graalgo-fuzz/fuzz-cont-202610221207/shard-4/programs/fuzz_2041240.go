package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint32, p2 uint8) uint {
    fmt.Println(p1)
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: fi6, F1: (p2 % (((p2 % (((p2 ^ (p2 / (((p2 - p2)) | 1)))) | 1))) | 1))}
        _ = fi6
        _ = v7
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint64, p1 int8, p2 uint) string {
    var v8 uint16 = 255
    v9 := []int{2147483647, 728}
    v10 := GenS2{F0: (p2 / ((uint(v8)) | 1))}
    v8 = 186
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int64 = -453
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return ((("" + "a") + "a") + "")
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 702
    var v3 uint64 = 42
    v4 := make(map[string]string)
    v5 := []int{1}
    if false {
        switch "" {
        case "foo bar":
            if ((42 != 1) && (!true)) {
                var v13 uint = genFn0(212, 0, 185)
                if false {
                    var v14 int32 = -1
                    _ = v14
                }
                _ = v13
            }
        case "hello":
            if false {
                fmt.Println(v3)
                v2 += 2
                if false {
                    v2 = (v2 - v2)
                    v2 = 42
                    v4[("a" + "\n\t")] = "a"
                } else {
                    var v15 int = 164
                    var v16 int16 = 1000
                    fmt.Println(v3)
                    _ = v15
                    _ = v16
                }
            }
        default:
            var v17 float64 = 0.1
            _ = v17
        }
    } else {
        delete(v4, "foo bar")
        v2 = 742
        if len(v5) > 0 {
            v5[0] = -931
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 float64 = ((1e10 - (((float64((1.5 + 1e10)) - 0.0) * float64(1.5)) / (-2.25 / 1.0))) + ((3.14159 * (1e10 - 1.5)) - (float64(0.0) + 0.1)))
    if (127 != 127) {
        var v19 uint64 = 0
        v19 = 0
        _ = v19
    } else {
        fmt.Println(v2)
        delete(v4, "foo bar")
    }
    if (true && false) {
        v20 := GenS3{F0: "a"}
        v5 = append(v5, 4294967295)
        fmt.Println(v2, v2)
        _ = v20
    } else {
        var v21 uint = v2
        _ = v21
    }
    switch 4294967295 {
    case 32768:
        if ((0 >= 32767) && false) {
            _ = len(v1)
        }
    case 559:
        for fi22 := 0; fi22 < 13; fi22++ {
            if (!(true || (int64(v2) < -9223372036854775808))) {
                switch 1 {
                case 127:
                    if false {
                        var v23 uint = v2
                        _ = v23
                    }
                case 9:
                    if (65535 >= 2147483648) {
                        v1[int(v3)] = false
                        var v24 uint16 = 42
                        _ = v24
                    }
                default:
                    var v25 int32 = -378
                    _ = v25
                }
                v26 := v4[("\n\t" + "")]
                _ = v26
            }
            _ = fi22
        }
    case 4294967295:
        fmt.Println(v18, v2)
    default:
        if (false && false) {
            if true {
                for fi27 := 0; fi27 < 16; fi27++ {
                    var v28 uint = v2
                    if ((fi27 + fi27) >= (fi27 * (fi27 % ((fi27) | 1)))) {
                        var v29 uint = (v2 & (v28 & 255))
                        fmt.Println(v2)
                        v30 := GenS4{F0: 1, F1: v2, F2: -2147483648}
                        _ = v29
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 4; fi31++ {
                            var v32 uint16 = 0
                            if len(v5) > 0 {
                                v5[0] = (fi27 | fi31)
                            }
                            _ = fi31
                            _ = v32
                        }
                        switch fi27 {
                        case -351:
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                if true {
                                    if (v3 < (v3 * 42)) {
                                        v33 := GenS3{F0: "a"}
                                        v34 := GenS5{F0: true, F1: (v3 << (((v28 | v28)) & 63))}
                                        var v35 int32 = 1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            case "":
                                for fi36 := 0; fi36 < 6; fi36++ {
                                    var v37 uint32 = 96
                                    var v38 int = 42
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            }
                        case 0:
                            for fi39 := 0; fi39 < 13; fi39++ {
                                if true {
                                    fmt.Println(v3, fi27, fi39)
                                    var v40 int8 = 59
                                    _ = v40
                                }
                                _ = fi39
                            }
                        }
                        v41 := GenS4{F0: 42, F1: v2, F2: -341}
                        _ = v41
                    }
                    _ = fi27
                    _ = v28
                }
            }
            v42 := v4[("foo bar" + "")]
            v18 = v18
            _ = v42
        }
    }
    var v43 uint64 = 0
    v43 += (^(v3))
    var v44 uint = genFn0(65535, uint32(v2), 96)
    if (float32(v18) >= (3.14159 / 1.5)) {
        if true {
            var v45 float64 = (v18 + -2.25)
            switch (("\n\t" + "hello") + ("hello" + "foo bar")) {
            case "a":
                v46 := GenS1{F0: -9223372036854775808, F1: 105}
                _ = v46
            }
            if (100 == 42) {
                v47 := GenS6{F0: 42}
                for fi48 := 0; fi48 < 7; fi48++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v3 -= v3
                    _ = fi48
                }
                if (!true) {
                    v44 += (v2 % ((genFn0(42, 42, 128)) | 1))
                }
                _ = v47
            } else {
                var v49 float32 = 3.14159
                _ = v49
            }
            _ = v45
        }
    } else {
        var v50 int64 = 596
        v43 = v43
        _ = v50
    }
    var v51 uint8 = 245
    v3 = (v43 * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v51
}
