package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    v4 := []int{}
    delete(v2, 21)
    var v5 int8 = 100
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = 0.0
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 42
    switch (("\n\t" + (("foo bar" + "") + (("foo bar" + ("a" + "foo bar")) + ("foo bar" + "")))) + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v5, v8)
    }
    switch 32768 {
    case 1:
        v9 := GenS1{F0: 128}
        _ = v9
    default:
        var v10 int64 = 0
        _ = v10
    }
    v11 := GenS2{F0: 211, F1: "hello", F2: 255}
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            if true {
                v14 := GenS2{F0: 593, F1: ("\n\t" + "foo bar"), F2: 255}
                _ = v14
            } else {
                v5 += v5
            }
            switch (-(fi13)) {
            case 127:
                if ((fi13 & int(v1)) <= fi13) {
                    var v15 int16 = -786
                    v15 += v15
                    var v16 float32 = float32(-2.25)
                    _ = v15
                    _ = v16
                } else {
                    v8 -= uint8(v1)
                    var v17 uint = 1
                    v18 := GenS2{F0: 46, F1: "\n\t", F2: 32767}
                    _ = v17
                    _ = v18
                }
            }
            var v19 uint64 = 2
            var v20 int = fi12
            if true {
                if (65535 == 732) {
                    v1 = v1
                    if (true && false) {
                        _ = len(v4)
                    } else {
                        var v21 uint16 = (v1 & (v1 * v1))
                        v22 := GenS2{F0: v19, F1: "\n\t", F2: (^(v21))}
                        var v23 string = "foo bar"
                        var v24 uint32 = 42
                        switch (v23 + v23) {
                        case "hello":
                            if (false || (((v23 + "") == v23) && false)) {
                                v1 = v1
                                switch v23 {
                                case "\n\t":
                                    v4 = append(v4, int(v19))
                                case "":
                                    var v25 uint = 527
                                    _ = v25
                                case "hello":
                                    v3 = v3
                                }
                                if true {
                                    if (v19 >= 4294967295) {
                                        v23 = ("hello" + v23)
                                        v26 := GenS1{F0: 897}
                                        var v27 uint8 = v8
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v28 := GenS1{F0: -798}
                                        v23 = v23
                                        _ = v28
                                    }
                                    fmt.Println(v19, fi13, v5)
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        _ = len(v4)
                                    case "a":
                                        fmt.Println(v20, v21)
                                    case "hello":
                                        v29 := GenS3{F0: v19, F1: (!false)}
                                        _ = v29
                                    }
                                    if false {
                                        var v30 int64 = int64(v5)
                                        if ((!false) && ((v8 >> ((18446744073709551615) & 63)) != (v8 - v8))) {
                                            v24 = v24
                                        }
                                        v31 := GenS1{F0: (fi13 & (v20 % ((int(v24)) | 1)))}
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v5 = (v5 ^ (v5 * 100))
                                        fmt.Println(v24)
                                    }
                                    if ((fi12 ^ fi13) < fi12) {
                                        _ = len(v4)
                                        fmt.Println(v21, v3)
                                        var v32 uint8 = (^(v8))
                                        _ = v32
                                    } else {
                                        var v33 int8 = -128
                                        _ = v33
                                    }
                                }
                            }
                        case "foo bar":
                            if (true && ((!(((!true) && true) || (1.5 > 0.1))) || (-425 == -2147483648))) {
                                v34 := GenS4{F0: 255, F1: v3, F2: 1000}
                                if false {
                                    var v35 int16 = -874
                                    _ = v35
                                } else {
                                    v1 += (v1 / (((v1 << ((42) & 63))) | 1))
                                }
                                v36 := GenS1{F0: (fi13 >> ((517) & 63))}
                                v3 = -9223372036854775808
                                switch v20 {
                                case 9223372036854775807:
                                    if (1.5 >= 1.5) {
                                        var v37 string = v23
                                        if (false && false) {
                                            var v38 float64 = 1.5
                                            v39 := GenS3{F0: 4294967295, F1: false}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v40 := GenS5{F0: v19}
                                            fmt.Println(v19, v5, v24)
                                            _ = v40
                                        }
                                        _ = v37
                                    }
                                case 32768:
                                    if ("hello" >= ("a" + v23)) {
                                        switch v23 {
                                        case "":
                                            v20 = fi12
                                        case "\n\t":
                                            var v41 uint8 = (v8 / ((61) | 1))
                                            _ = v41
                                        case "hello":
                                            v23 = ((v23 + ("" + v23)) + ((((v23 + "") + "") + "") + v23))
                                        }
                                        v24 = v24
                                    }
                                default:
                                    v8 = (v8 & v8)
                                }
                                v42 := GenS6{F0: false, F1: uint(v21)}
                                fmt.Println(v24, v5)
                                _ = v34
                                _ = v36
                                _ = v42
                            } else {
                                if ((fi13 * fi12) > 9223372036854775807) {
                                    v24 = (v24 * (v24 - (^(v24))))
                                    v4 = append(v4, 32768)
                                }
                                _ = len(v4)
                                v8 = (v8 % ((v8) | 1))
                            }
                        default:
                            if (!true) {
                                switch (("\n\t" + (("foo bar" + ((v23 + "\n\t") + ((v23 + "") + (v23 + ("a" + "a"))))) + "foo bar")) + (v23 + v23)) {
                                case "a":
                                    v43 := GenS7{F0: v24, F1: (1.5 == ((3.14159 / 1.5) / 1.0))}
                                    _ = v43
                                }
                                switch fi13 {
                                case -412:
                                    var v44 int32 = 42
                                    _ = v44
                                case 0:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v24 = v24
                                    case "a":
                                        v45 := GenS3{F0: 2, F1: ((v5 - (v5 | (v5 << ((964) & 63)))) > (v5 | 127))}
                                        _ = v45
                                    }
                                    if true {
                                        var v46 uint = 4294967295
                                        var v47 uint16 = uint16(v46)
                                        _ = v46
                                        _ = v47
                                    } else {
                                        v48 := GenS5{F0: 4294967295}
                                        if ((1 == 18446744073709551615) || true) {
                                            switch v23 {
                                            case "foo bar":
                                                var v49 float64 = float64((3.14159 / (1e10 / 3.14159)))
                                                fmt.Println(fi12, fi12, v49)
                                                _ = v49
                                            case "\n\t":
                                                v50 := GenS3{F0: v19, F1: false}
                                                v5 += (-(v5))
                                                _ = v50
                                            case "":
                                                var v51 uint8 = 128
                                                v3 = v3
                                                v8 = v8
                                                _ = v51
                                            default:
                                                if false {
                                                    if false {
                                                        var v52 uint64 = 1
                                                        var v53 string = v23
                                                        var v54 uint16 = (v1 + 32767)
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v2[(v20 & fi12)] = 2147483647
                                                        var v55 int = 9223372036854775807
                                                        _ = v55
                                                    }
                                                    if true {
                                                        var v56 int16 = 0
                                                        var v57 bool = false
                                                        v58 := GenS6{F0: v57, F1: 65535}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = len(v2)
                                                }
                                            }
                                            v8 = 60
                                        }
                                        v3 += v3
                                        _ = v48
                                    }
                                }
                                v59 := GenS3{F0: 977, F1: (v8 != v8)}
                                _ = v59
                            } else {
                                var v60 uint32 = (v24 * (^(v24)))
                                v61 := GenS1{F0: (fi13 & (v20 << ((239) & 63)))}
                                v62 := GenS5{F0: 1}
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        }
                        if (159 != uint(v8)) {
                            v23 = v23
                            _ = len(v2)
                            var v63 float64 = 1e10
                            _ = v63
                        }
                        v64 := v2[fi12]
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                        _ = v64
                    }
                }
                v65 := GenS7{F0: 954, F1: false}
                v66 := GenS3{F0: uint64(v8), F1: (!false)}
                v8 += 1
                fmt.Println(fi13, v3)
                _ = v65
                _ = v66
            } else {
                v67 := GenS5{F0: v19}
                _ = v67
            }
            _ = fi13
            _ = v19
            _ = v20
        }
        var v68 uint16 = 65535
        _ = fi12
        _ = v68
    }
    var v69 int = 32767
    var v70 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v8
}
