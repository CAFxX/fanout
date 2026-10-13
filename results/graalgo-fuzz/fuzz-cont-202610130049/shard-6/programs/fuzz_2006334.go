package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint64) bool {
    v6 := []int{32767}
    v6 = append(v6, 127)
    fmt.Println(p1, p0)
    switch -274 {
    case 457:
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi7, p0)
            _ = fi7
        }
    case 4294967295:
        var v8 int64 = -9223372036854775808
        _ = v8
    case 0:
        if (1 >= -1) {
            p1 += 4294967295
            var v9 bool = false
            var v10 int8 = int8(p0)
            _ = v10
            _ = v9
        } else {
            p0 = p0
            p1 = (^(p1))
        }
    default:
        var v11 int32 = 932
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return (0 > p1)
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{42, 351, 2147483647, -9223372036854775808}
    v3 := make(map[string]uint32)
    var v4 string = "a"
    var v5 float32 = float32((3.14159 / 1e10))
    var v12 float64 = -2.25
    delete(v3, v4)
    if false {
        v3[(v4 + "foo bar")] = 2147483648
        if ((2 == 433) || (false && false)) {
            v13 := GenS1{F0: (false || false), F1: 420, F2: 1}
            _ = v13
        }
    } else {
        v12 = 1.5
    }
    fmt.Println(v12)
    v5 = v5
    if ((1e10 + v1) >= float32(-2.25)) {
        v14 := GenS1{F0: (972 <= 359), F1: 0, F2: 1}
        v15 := GenS1{F0: (!(false || false)), F1: 42, F2: 0}
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v4)
            v4 = v4
            _ = fi16
        }
        switch -85 {
        case 127:
            if (false || (true && ((700 != 81) || false))) {
                switch v4 {
                case "hello":
                    if false {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            v1 = (((((v1 / (v5 + 0.0)) + (v5 - v5)) / 0.1) - 0.0) * 3.14159)
                            v18 := GenS1{F0: false, F1: -2147483648, F2: -1}
                            v12 += v12
                            _ = fi17
                            _ = v18
                        }
                        _ = len(v3)
                        var v19 uint8 = 127
                        _ = v19
                    } else {
                        fmt.Println(v5)
                        var v20 string = v4
                        _ = v20
                    }
                case "":
                    var v21 uint64 = 891
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 17; fi22++ {
                        switch "foo bar" {
                        case "foo bar":
                            v12 = (0.1 - ((float64(3.14159) + v12) - v12))
                        case "a":
                            fmt.Println(v5, fi22)
                        case "hello":
                            v4 = (v4 + v4)
                        }
                        var v23 bool = (true && (524 < 42))
                        var v24 int = fi22
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            } else {
                fmt.Println(v4)
            }
        default:
            if true {
                fmt.Println(v4)
                var v25 int64 = -567
                _ = v25
            }
        }
        v1 += (v5 * float32(float64(v5)))
    }
    var v26 uint = 18
    v27 := GenS1{F0: false, F1: -829, F2: 32767}
    _ = len(v3)
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 14; fi29++ {
            if true {
                if (v4 <= ((v4 + "a") + "")) {
                    v26 = v26
                    switch "a" {
                    case "\n\t":
                        v30 := GenS2{F0: 2}
                        _ = v30
                    case "":
                        v31 := GenS3{F0: -1}
                        _ = v31
                    default:
                        if false {
                            if (!false) {
                                v26 -= (v26 % (((v26 / (((v26 ^ 0)) | 1))) | 1))
                                v1 -= v1
                            }
                            if (v4 == "") {
                                v32 := GenS1{F0: false, F1: 0, F2: 32767}
                                var v33 uint8 = 128
                                _ = v32
                                _ = v33
                            } else {
                                switch 128 {
                                case -320:
                                    v34 := GenS4{F0: 65535, F1: ("a" + v4), F2: 2}
                                    _ = v34
                                case 127:
                                    switch ("" + v4) {
                                    case "hello":
                                        v35 := GenS5{F0: (v26 - 857), F1: 42}
                                        _ = v35
                                    case "\n\t":
                                        var v36 uint32 = 2
                                        _ = v36
                                    }
                                case 128:
                                    v37 := GenS3{F0: 0}
                                    _ = v37
                                }
                            }
                            v4 = ("\n\t" + (v4 + v4))
                        } else {
                            switch v4 {
                            case "":
                                fmt.Println(fi28, v12, v26)
                            case "\n\t":
                                v38 := GenS6{F0: 877}
                                _ = v38
                            case "a":
                                v1 -= v1
                            default:
                                if (v12 != (0.0 - (v12 / 0.1))) {
                                    v39 := GenS5{F0: 2, F1: -9223372036854775808}
                                    _ = v39
                                }
                            }
                        }
                    }
                }
                fmt.Println(v12)
                v40 := v3[("a" + "hello")]
                _ = v40
            }
            v41 := v3[("a" + v4)]
            _ = fi29
            _ = v41
        }
        v42 := GenS6{F0: 4294967295}
        _ = fi28
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
