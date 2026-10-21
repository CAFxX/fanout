package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0(p0 int16, p1 uint8, p2 uint32) int8 {
    v5 := []int{4294967295, 9223372036854775807}
    if (true || true) {
        fmt.Println(p2, p1)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        p1 -= 253
        v6 := GenS1{F0: "foo bar", F1: 80}
        v7 := GenS2{F0: -2147483648, F1: 1, F2: 1}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1(p0 uint8) int32 {
    v8 := GenS2{F0: -401, F1: -32768, F2: -55}
    var v9 int = -534
    v10 := []int{128, (v9 << ((119) & 63)), 32768}
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v9, fi11, fi11)
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -223
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{166, 65535}
    var v3 uint32 = uint32(v1)
    v4 := []int{128}
    var v12 string = "a"
    var v13 uint16 = 2
    if (true && (42 == 42)) {
        v13 += v13
        v3 -= (v3 % ((v3) | 1))
    }
    v1 += (v1 ^ v1)
    var v14 string = v12
    v1 = (v1 << ((2) & 63))
    for fi15 := 0; fi15 < 8; fi15++ {
        if (int16(v13) < 1000) {
            switch (v14 + (v12 + v12)) {
            case "a":
                v12 = ("hello" + ("" + v12))
            }
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                v14 = ""
                if ((false || (!(!true))) && ((18446744073709551615 < 0) && (70 <= genFn0(42, v1, (^(v3)))))) {
                    if false {
                        switch fi15 {
                        case 9223372036854775807:
                            _ = len(v4)
                        case 0:
                            if len(v2) > 0 {
                                v2[0] = v3
                            }
                        case -1:
                            v17 := GenS2{F0: (-(fi16)), F1: -143, F2: -49}
                            _ = v17
                        default:
                            v14 = "a"
                        }
                        if true {
                            v18 := GenS2{F0: fi15, F1: 42, F2: -61}
                            _ = v18
                        }
                        switch fi15 {
                        case 9223372036854775807:
                            fmt.Println(v13, v13, v13)
                        case 2147483647:
                            if false {
                                v13 = v13
                            } else {
                                var v19 bool = (!false)
                                _ = v19
                            }
                        case 733:
                            switch "" {
                            case "foo bar":
                                if (!true) {
                                    if false {
                                        var v20 int16 = 1000
                                        var v21 float64 = 1e10
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 uint8 = 108
                                        var v23 uint16 = (v13 % ((v13) | 1))
                                        switch fi15 {
                                        case 32768:
                                            switch ((v14 + ("\n\t" + v14)) + (v12 + v12)) {
                                            case "foo bar":
                                                v14 = v12
                                            case "hello":
                                                var v24 uint16 = (v23 ^ (v23 >> ((729) & 63)))
                                                if true {
                                                    switch (v14 + "a") {
                                                    case "hello":
                                                        v4 = append(v4, fi15)
                                                    case "":
                                                        var v25 int16 = -1
                                                        _ = v25
                                                    }
                                                    v26 := GenS1{F0: (v12 + v14), F1: -111}
                                                    fmt.Println(v13, v14, fi16)
                                                    _ = v26
                                                } else {
                                                    var v27 bool = true
                                                    _ = v27
                                                }
                                                _ = v24
                                            case "":
                                                var v28 uint64 = 2
                                                _ = v28
                                            default:
                                                v4 = append(v4, (-(fi15)))
                                            }
                                        case 129:
                                            v22 = uint8(v23)
                                        default:
                                            switch (fi16 >> ((47) & 63)) {
                                            case -332:
                                                v13 += (v13 * 65535)
                                            }
                                        }
                                        _ = len(v4)
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            case "":
                                switch 1 {
                                case 0:
                                    var v29 int16 = -141
                                    _ = v29
                                case 224:
                                    if (0 != -1) {
                                        if false {
                                            var v30 int8 = 0
                                            _ = v30
                                        } else {
                                            switch ((v12 + "foo bar") + "") {
                                            case "a":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            default:
                                                fmt.Println(v1, v14, v3)
                                            }
                                            if false {
                                                v31 := GenS3{F0: 1, F1: (^(v1))}
                                                if false {
                                                    v1 = (v1 | 1)
                                                } else {
                                                    v32 := GenS1{F0: v12, F1: 0}
                                                    _ = v32
                                                }
                                                _ = v31
                                            } else {
                                                fmt.Println(fi15, fi15, v3)
                                                var v33 int64 = -9223372036854775808
                                                var v34 uint8 = 110
                                                _ = v33
                                                _ = v34
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v35 uint8 = (v1 + (^(v1)))
                                            if (false || (int16(v35) >= -32768)) {
                                                if false {
                                                    v3 += (v3 >> ((0) & 63))
                                                    v12 = (v12 + ("\n\t" + (v12 + v12)))
                                                }
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v36 uint = 2
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                        var v37 int32 = -2147483648
                                        _ = v37
                                    } else {
                                        if false {
                                            v38 := GenS1{F0: v12, F1: 9223372036854775807}
                                            fmt.Println(fi16, v13)
                                            if false {
                                                switch (v12 + ("" + ((((v12 + "a") + (((v12 + (((v12 + (("hello" + v12) + v12)) + v14) + (v14 + "a"))) + "") + ((v12 + v12) + ((("hello" + "a") + (v14 + v12)) + v14)))) + "foo bar") + v14))) {
                                                case "hello":
                                                    v3 -= v3
                                                }
                                                v12 = ""
                                                if (true || (305 >= 65535)) {
                                                    var v39 uint16 = 32767
                                                    v40 := GenS2{F0: fi16, F1: 483, F2: 31}
                                                    _ = len(v2)
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v14 = ((v14 + ((v14 + "hello") + ((v12 + v14) + ""))) + "hello")
                                                    v3 -= 523
                                                    fmt.Println(v12, v1, v1)
                                                }
                                            } else {
                                                v13 += v13
                                            }
                                            _ = v38
                                        } else {
                                            fmt.Println(v14)
                                        }
                                        var v41 int64 = 42
                                        var v42 bool = false
                                        _ = v41
                                        _ = v42
                                    }
                                default:
                                    v12 = (v14 + ((("foo bar" + v14) + (v14 + v14)) + v14))
                                }
                            default:
                                var v43 int = fi16
                                _ = v43
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    } else {
                        var v44 float32 = -2.25
                        var v45 uint16 = v13
                        _ = v44
                        _ = v45
                    }
                    var v46 uint = 2
                    _ = v46
                }
                v47 := GenS4{F0: 727, F1: v3}
                v12 = (("" + v14) + v12)
                _ = v47
            }
            v3 = 667
            if len(v2) > 0 {
                v2[0] = 42
            }
            var v48 int32 = 2147483647
            var v49 int64 = 9223372036854775807
            _ = fi16
            _ = v48
            _ = v49
        }
        var v50 uint64 = 446
        _ = fi15
        _ = v50
    }
    var v51 uint64 = 0
    v51 = 0
    v52 := GenS1{F0: "foo bar", F1: 1}
    if (18446744073709551615 <= v51) {
        for fi53 := 0; fi53 < 7; fi53++ {
            var v54 int32 = -2147483648
            v1 = (^(v1))
            v14 = v12
            _ = fi53
            _ = v54
        }
        for fi55 := 0; fi55 < 16; fi55++ {
            var v56 uint = 619
            v1 += uint8(v51)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi55
            _ = v56
        }
    } else {
        v4 = append(v4, 9223372036854775807)
    }
    v14 = v12
    switch v12 {
    case "hello":
        _ = len(v2)
    default:
        if ((v3 | v3) <= 2147483647) {
            var v57 uint32 = (v3 << ((65535) & 63))
            _ = v57
        } else {
            v4 = append(v4, 78)
            v58 := GenS4{F0: 486, F1: v3}
            _ = v58
        }
    }
    if false {
        var v59 uint16 = (v13 / (((v13 ^ v13)) | 1))
        if len(v2) > 0 {
            v2[0] = v3
        }
        switch "\n\t" {
        case "hello":
            v51 += v51
        }
        _ = v59
    } else {
        v60 := GenS5{F0: 42, F1: v3}
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
}
