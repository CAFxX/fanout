package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 int32 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8 }
func genFn0() uint64 {
    v5 := GenS1{F0: 0}
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 54
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint8, p1 uint, p2 int32) uint16 {
    var v9 bool = false
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int8 = 100
    v4 := []int{1, -2147483648, 766}
    if (42 <= 118) {
        v10 := GenS1{F0: -782}
        for fi11 := 0; fi11 < 5; fi11++ {
            if true {
                for fi12 := 0; fi12 < 17; fi12++ {
                    if true {
                        fmt.Println(fi11, fi12, v3)
                        switch "\n\t" {
                        case "hello":
                            v13 := GenS2{F0: (v3 / (((-(v3))) | 1))}
                            _ = v13
                        case "":
                            fmt.Println(fi11, fi12, v3)
                        }
                    } else {
                        var v14 uint16 = 1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if false {
                            v3 += 0
                        } else {
                            v3 += -75
                            v14 += v14
                            v14 -= (v14 & (v14 % ((genFn1(42, 1, 42)) | 1)))
                            if false {
                                v3 += (v3 % ((v3) | 1))
                            }
                        }
                        _ = v14
                    }
                    _ = fi12
                }
                fmt.Println(v3)
            } else {
                var v15 string = "foo bar"
                var v16 string = ("foo bar" + "hello")
                _ = v15
                _ = v16
            }
            v3 = v3
            v17 := GenS3{F0: 1, F1: "", F2: 2147483647}
            _ = fi11
            _ = v17
        }
        v18 := GenS4{F0: 2147483648, F1: 480, F2: 430}
        _ = v10
        _ = v18
    } else {
        switch "foo bar" {
        case "a":
            fmt.Println(v3)
        case "foo bar":
            v19 := GenS4{F0: 4294967295, F1: 190, F2: 4294967295}
            _ = v19
        default:
            fmt.Println(v3)
        }
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        switch "hello" {
        case "\n\t":
            v3 -= -50
        case "a":
            var v21 int32 = -20
            _ = v21
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            v3 = -128
            _ = fi22
        }
        if (true && (true || false)) {
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 int64 = 264
                fmt.Println(v24)
                _ = fi23
                _ = v24
            }
            v25 := GenS5{F0: v3, F1: ((!false) || ((true && (1e10 != float64(float32(3.14159)))) || true)), F2: 42}
            if true {
                var v26 string = "\n\t"
                if false {
                    if (-2.25 <= float64(1e10)) {
                        if (genFn0() >= 2) {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                var v28 int8 = v3
                                _ = fi27
                                _ = v28
                            }
                            switch "hello" {
                            case "foo bar":
                                var v29 uint32 = 2147483648
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 8; fi30++ {
                                v31 := GenS6{F0: false, F1: v3, F2: -1}
                                _ = fi30
                                _ = v31
                            }
                        } else {
                            v26 = "\n\t"
                            if false {
                                if (false || true) {
                                    var v32 uint = 0
                                    for fi33 := 0; fi33 < 16; fi33++ {
                                        if (-2.25 >= -2.25) {
                                            v4 = append(v4, 9223372036854775807)
                                            fmt.Println(v3, v32)
                                        }
                                        v32 = 18446744073709551615
                                        var v34 uint16 = 255
                                        _ = fi33
                                        _ = v34
                                    }
                                    for fi35 := 0; fi35 < 20; fi35++ {
                                        if (!(!true)) {
                                            if (0.1 == float32(0.1)) {
                                                fmt.Println(fi20)
                                            } else {
                                                _ = len(v1)
                                            }
                                            v3 += v3
                                            switch (v26 + "foo bar") {
                                            case "hello":
                                                var v36 uint16 = 255
                                                _ = v36
                                            }
                                            v26 = v26
                                            switch "foo bar" {
                                            case "\n\t":
                                                v26 = (("a" + v26) + v26)
                                            }
                                        }
                                        fmt.Println(v26)
                                        _ = fi35
                                    }
                                    _ = v32
                                }
                                v26 = ("\n\t" + v26)
                            }
                        }
                        v37 := GenS5{F0: int8(fi20), F1: (!(2147483648 < 515)), F2: 0}
                        for fi38 := 0; fi38 < 17; fi38++ {
                            if (!true) {
                                fmt.Println(fi38, fi38)
                            }
                            var v39 int64 = -9223372036854775808
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                    if true {
                        if ((v3 > v3) && (!true)) {
                            if (164 <= 0) {
                                for fi40 := 0; fi40 < 9; fi40++ {
                                    switch (^(fi20)) {
                                    case -9223372036854775808:
                                        v41 := GenS3{F0: -1, F1: v26, F2: 1}
                                        _ = v41
                                    case -362:
                                        v2[v26] = v26
                                    case 9223372036854775807:
                                        v3 += v3
                                    }
                                    _ = fi40
                                }
                            } else {
                                if (0 <= 2) {
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        var v43 int8 = (v3 | (-(v3)))
                                        v43 -= 42
                                        _ = fi42
                                        _ = v43
                                    }
                                    v44 := GenS4{F0: 599, F1: 65535, F2: 2}
                                    _ = v44
                                }
                                v45 := GenS6{F0: false, F1: v3, F2: -890}
                                _ = v45
                            }
                            var v46 float64 = 1e10
                            _ = v46
                        } else {
                            for fi47 := 0; fi47 < 2; fi47++ {
                                v26 = "a"
                                v3 = (v3 + v3)
                                _ = fi47
                            }
                            var v48 string = v26
                            _ = v48
                        }
                        var v49 uint64 = 42
                        _ = v49
                    }
                    delete(v2, "hello")
                }
                v50 := GenS7{F0: fi20, F1: -1}
                _ = v26
                _ = v50
            } else {
                _ = len(v1)
            }
            _ = v25
        } else {
            fmt.Println(fi20)
            var v51 bool = true
            var v52 float64 = (1e10 * (((1e10 * 1.5) / ((float64(3.14159) / 0.1) * ((1.5 + 3.14159) - 1.5))) / (0.1 - (float64(-2.25) / float64(((-2.25 + (1e10 - (0.0 / 1e10))) - 0.1))))))
            _ = v51
            _ = v52
        }
        _ = fi20
    }
    var v53 float64 = float64(((0.0 * 1e10) + (-2.25 + (0.0 - 0.1))))
    switch 2147483647 {
    case 0:
        fmt.Println(v3, v53)
    case 9223372036854775807:
        fmt.Println(v3)
    }
    var v54 bool = false
    for fi55 := 0; fi55 < 13; fi55++ {
        switch (fi55 >> ((4294967295) & 63)) {
        case 128:
            switch "\n\t" {
            case "":
                var v56 int32 = -879
                _ = v56
            }
        }
        v57 := GenS8{F0: 893}
        for fi58 := 0; fi58 < 16; fi58++ {
            delete(v1, ("a" + "hello"))
            v59 := GenS5{F0: (v3 + v3), F1: v54, F2: genFn0()}
            _ = fi58
            _ = v59
        }
        _ = fi55
        _ = v57
    }
    v53 += 3.14159
    _ = len(v1)
    var v60 uint = 65535
    v3 -= v3
    v61 := GenS9{F0: (^(v3))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v60
    _ = v61
}
