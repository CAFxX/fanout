package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int; F2 bool }
type GenS6 struct { F0 int16; F1 int; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 uint8) string {
    if (false && (false || false)) {
        var v6 bool = (69 != 898)
        v7 := []int{128, 31}
        v8 := GenS1{F0: 752, F1: 0, F2: "\n\t"}
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 uint = 205
        v11 := GenS2{F0: 255, F1: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return (("\n\t" + "foo bar") + "foo bar")
}
func genFn1(p0 uint32) bool {
    v12 := GenS3{F0: -968}
    for fi13 := 0; fi13 < 8; fi13++ {
        switch (fi13 / (((fi13 << ((42) & 63))) | 1)) {
        case 2147483647:
            v14 := []int{128, fi13, (fi13 % ((0) | 1))}
            _ = v14
        default:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float64 = 0.1
                v17 := GenS1{F0: 2, F1: 4294967295, F2: ""}
                var v18 int = -492
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        p0 = 1
        v19 := GenS1{F0: 633, F1: 4294967295, F2: "hello"}
        _ = fi13
        _ = v19
    }
    fmt.Println(p0)
    switch -990 {
    case -890:
        if (858 >= 255) {
            p0 = 642
            fmt.Println(p0)
        } else {
            v20 := GenS1{F0: 218, F1: 986, F2: "a"}
            switch 0 {
            case -301:
                v21 := []int{546, 42}
                _ = v21
            case 2147483647:
                for fi22 := 0; fi22 < 16; fi22++ {
                    fmt.Println(fi22, p0)
                    v23 := GenS3{F0: 1}
                    if ((false || false) || (!false)) {
                        fmt.Println(fi22, fi22)
                        v24 := []int{(fi22 * -9223372036854775808)}
                        p0 = (p0 / ((p0) | 1))
                        _ = v24
                    }
                    _ = fi22
                    _ = v23
                }
            default:
                v25 := []int{4294967295, -9223372036854775808, 127}
                _ = v25
            }
            fmt.Println(p0)
            _ = v20
        }
    default:
        v26 := GenS4{F0: "", F1: 255, F2: 4294967295}
        _ = v26
    }
    _ = p0
    _ = v12
    return true
}

func main() {
    var v1 int = 2147483647
    var v2 int8 = -1
    var v3 uint8 = 82
    var v4 int16 = 314
    var v5 bool = true
    var v27 bool = true
    v28 := []int{}
    for fi29 := 0; fi29 < 7; fi29++ {
        if v27 {
            if (-1 > 2147483647) {
                switch fi29 {
                case 127:
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
            } else {
                v27 = false
                fmt.Println(v4)
            }
            v1 = (fi29 & (fi29 ^ (fi29 ^ 128)))
        }
        var v30 int8 = (v2 >> ((4294967295) & 63))
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 4294967295
    if v27 {
        if (475 >= 819) {
            switch 591 {
            case 32768:
                var v32 float32 = -2.25
                _ = v32
            case 2147483647:
                switch "" {
                case "foo bar":
                    if v27 {
                        for fi33 := 0; fi33 < 8; fi33++ {
                            if v27 {
                                var v34 uint16 = 0
                                _ = len(v28)
                                if v5 {
                                    if false {
                                        fmt.Println(fi33)
                                        var v35 uint8 = v3
                                        if true {
                                            if (v31 == v31) {
                                                v36 := GenS4{F0: "", F1: v35, F2: 42}
                                                var v37 int8 = -1
                                                v1 -= -2147483648
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v38 int16 = -162
                                                var v39 int32 = 2147483647
                                                var v40 int64 = -721
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                            v41 := GenS1{F0: v34, F1: 1, F2: "foo bar"}
                                            if len(v28) > 0 {
                                                v28[0] = fi33
                                            }
                                            _ = v41
                                        }
                                        _ = v35
                                    } else {
                                        var v42 string = "a"
                                        switch v42 {
                                        case "":
                                            if len(v28) > 0 {
                                                v28[0] = fi33
                                            }
                                        case "hello":
                                            if (!(-2.25 < 0.1)) {
                                                _ = len(v28)
                                                for fi43 := 0; fi43 < 3; fi43++ {
                                                    var v44 float32 = (1e10 - float32(-2.25))
                                                    v45 := GenS5{F0: (^(v34)), F1: -947, F2: (((!(42 > -1)) && true) || (0 >= 18446744073709551615))}
                                                    var v46 int16 = v4
                                                    _ = fi43
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 uint64 = 4294967295
                                                _ = v47
                                            }
                                        }
                                        _ = v42
                                    }
                                    v5 = v5
                                    fmt.Println(v31)
                                }
                                _ = v34
                            }
                            v48 := GenS6{F0: (v4 | (v4 + v4)), F1: (v1 - v1), F2: v3}
                            _ = fi33
                            _ = v48
                        }
                    } else {
                        var v49 int16 = (-(v4))
                        for fi50 := 0; fi50 < 2; fi50++ {
                            var v51 uint64 = 18446744073709551615
                            _ = fi50
                            _ = v51
                        }
                        _ = v49
                    }
                case "hello":
                    var v52 string = ""
                    _ = v52
                default:
                    var v53 int16 = v4
                    _ = v53
                }
            default:
                for fi54 := 0; fi54 < 14; fi54++ {
                    if len(v28) > 0 {
                        v28[0] = v1
                    }
                    _ = fi54
                }
            }
            if (0 == v2) {
                var v55 int32 = 0
                _ = len(v28)
                v28 = append(v28, v1)
                _ = v55
            } else {
                var v56 uint8 = v3
                if len(v28) > 0 {
                    v28[0] = (v1 / (((v1 | v1)) | 1))
                }
                _ = v56
            }
            var v57 int16 = (v4 | (v4 | (v4 + (-(v4)))))
            _ = v57
        }
    } else {
        var v58 uint8 = 241
        _ = v58
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        var v60 int = -1
        _ = fi59
        _ = v60
    }
    if len(v28) > 0 {
        v28[0] = v1
    }
    v61 := GenS7{F0: v31, F1: -2147483648, F2: -481}
    if true {
        var v62 uint8 = 42
        var v63 uint16 = 255
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v61
}
