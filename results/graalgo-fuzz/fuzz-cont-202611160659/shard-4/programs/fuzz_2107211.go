package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 string; F1 string; F2 int }
type GenS6 struct { F0 int }

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 0
    var v3 uint32 = 2147483647
    var v4 int32 = 2147483647
    var v5 uint64 = 2
    v1 += v1
    v6 := GenS1{F0: (v5 * 2), F1: "foo bar"}
    v2 = (^(v2))
    if (("" == "hello") || true) {
        v5 = 2
        fmt.Println(v1)
        v7 := GenS1{F0: (v5 + 18446744073709551615), F1: ("foo bar" + "foo bar")}
        _ = v7
    }
    if ((!true) && false) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v4 = 42
            v9 := GenS1{F0: (^(v5)), F1: "foo bar"}
            v10 := GenS1{F0: v5, F1: "hello"}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: uint8(v3), F1: (v4 & (v4 << ((67) & 63)))}
        v1 -= v1
        _ = v11
    }
    v12 := []int{0}
    if len(v12) > 0 {
        v12[0] = 774
    }
    var v13 float32 = (3.14159 - ((((0.0 / ((0.0 / 1e10) + ((0.1 + float32(((v1 * v1) / (1.5 - v1)))) * (-2.25 + -2.25)))) - (3.14159 - 0.1)) / (((float32((float64(1.5) - 0.1)) + 1e10) / 1.0) + 0.1)) - 0.0))
    v1 = 1e10
    v14 := GenS1{F0: uint64(v3), F1: ("" + "")}
    fmt.Println(v4, v2, v4)
    fmt.Println(v4)
    var v15 bool = ("a" > "foo bar")
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 17; fi17++ {
        v4 -= (v4 >> ((183) & 63))
        if true {
            if false {
                if v15 {
                    var v18 int64 = 9223372036854775807
                    _ = v18
                } else {
                    if (!(4294967295 >= (fi17 | fi17))) {
                        switch (fi17 ^ -229) {
                        case 641:
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v16 = "foo bar"
                                fmt.Println(v2)
                                v2 -= v2
                                if (!(!((42 < 4294967295) || (2 <= 531)))) {
                                    v13 += v13
                                    var v20 uint64 = (^(v5))
                                    _ = len(v12)
                                    _ = v20
                                } else {
                                    v13 = ((3.14159 + float32((v1 * v1))) + v13)
                                }
                                v21 := GenS1{F0: (^(v5)), F1: v16}
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v22 := GenS2{F0: v2, F1: 0}
                                var v23 float64 = (((float64(1.5) * 1.5) - 1e10) * v1)
                                _ = fi19
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                        case -300:
                            fmt.Println(v1, v16, fi17)
                        default:
                            switch "hello" {
                            case "hello":
                                if v15 {
                                    v24 := GenS3{F0: 42, F1: v16}
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        var v26 uint16 = 65535
                                        switch v16 {
                                        case "a":
                                            v16 = ((v16 + v16) + v16)
                                        case "foo bar":
                                            var v27 uint16 = v26
                                            _ = v27
                                        default:
                                            v3 -= 2
                                        }
                                        switch (fi25 % ((4294967295) | 1)) {
                                        case 5:
                                            if v15 {
                                                v15 = false
                                                v28 := GenS4{F0: -694, F1: (v5 * uint64(v2)), F2: v26}
                                                _ = v28
                                            }
                                        case 128:
                                            var v29 string = "foo bar"
                                            _ = v29
                                        default:
                                            if (((v1 - v1) - v1) > 3.14159) {
                                                var v30 uint8 = v2
                                                _ = v30
                                            } else {
                                                fmt.Println(v5, v26, v5)
                                                var v31 uint8 = (v2 << ((42) & 63))
                                                _ = v31
                                            }
                                        }
                                        var v32 int8 = -103
                                        v12 = append(v12, (fi25 + (fi25 ^ (fi17 - 2147483647))))
                                        var v33 uint64 = (v5 ^ (v5 << ((1) & 63)))
                                        var v34 float32 = (3.14159 / v13)
                                        _ = fi25
                                        _ = v26
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = v24
                                } else {
                                    if ((v4 + -2147483648) != -810) {
                                        var v35 uint8 = v2
                                        _ = v35
                                    }
                                    v3 += 774
                                }
                            case "\n\t":
                                var v36 int8 = 0
                                _ = v36
                            default:
                                _ = len(v12)
                            }
                        }
                        var v37 uint8 = (v2 * v2)
                        _ = v37
                    }
                }
            } else {
                v13 += (v13 + 0.1)
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                v13 += v13
                if (42 >= v2) {
                    if false {
                        v39 := GenS5{F0: ("foo bar" + (v16 + (v16 + v16))), F1: v16, F2: (fi17 * fi38)}
                        _ = v39
                    } else {
                        switch "a" {
                        case "a":
                            v12 = append(v12, fi17)
                        }
                        var v40 int8 = 42
                        _ = v40
                    }
                    switch -387 {
                    case 127:
                        v15 = (false && true)
                        if v15 {
                            fmt.Println(v13, v3)
                        } else {
                            fmt.Println(v1, fi38, v15)
                        }
                    case 1:
                        _ = len(v12)
                        v41 := GenS5{F0: v16, F1: "hello", F2: (fi38 / ((-2147483648) | 1))}
                        v12 = append(v12, (^(fi38)))
                        if len(v12) > 0 {
                            v12[0] = fi17
                        }
                        _ = v41
                    }
                    if v15 {
                        var v42 int8 = 107
                        v2 = v2
                        v43 := GenS5{F0: "hello", F1: (v16 + ("" + v16)), F2: (fi38 << ((736) & 63))}
                        _ = v42
                        _ = v43
                    } else {
                        _ = len(v12)
                    }
                } else {
                    if v15 {
                        if v15 {
                            var v44 uint8 = v2
                            v45 := GenS6{F0: fi17}
                            _ = v44
                            _ = v45
                        }
                    }
                    if v15 {
                        if len(v12) > 0 {
                            v12[0] = (fi38 % ((fi38) | 1))
                        }
                    }
                }
                v46 := GenS3{F0: 127, F1: "\n\t"}
                _ = fi38
                _ = v46
            }
            v4 = (v4 / (((v4 & -660)) | 1))
        } else {
            if v15 {
                var v47 int = (fi17 * (fi17 * fi17))
                var v48 uint64 = v5
                _ = v47
                _ = v48
            } else {
                v13 = v13
                for fi49 := 0; fi49 < 15; fi49++ {
                    var v50 bool = true
                    var v51 int = fi49
                    v51 = -894
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            }
            v16 = v16
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
