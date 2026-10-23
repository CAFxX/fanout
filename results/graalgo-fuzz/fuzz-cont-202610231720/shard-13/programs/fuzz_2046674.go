package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int; F2 bool }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int = -1
    var v2 int32 = 0
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    var v5 float32 = ((3.14159 * ((3.14159 * 1e10) / 1.0)) + (3.14159 - ((3.14159 - 1e10) / -2.25)))
    var v6 float32 = float32(-2.25)
    v1 += v1
    switch ("" + "hello") {
    case "foo bar":
        v6 = (0.1 + v5)
    case "\n\t":
        if (false && ((!(!false)) || true)) {
            for fi7 := 0; fi7 < 4; fi7++ {
                switch v1 {
                case 32767:
                    if (v1 <= (-(v1))) {
                        fmt.Println(fi7, v5, v1)
                        v8 := GenS1{F0: 11, F1: 0}
                        v1 = (-(fi7))
                        _ = v8
                    }
                case 46:
                    v9 := []int{fi7, (^(v1))}
                    _ = v9
                }
                _ = fi7
            }
        } else {
            if false {
                if (127 >= 0) {
                    v5 += (v5 + ((float32((3.14159 * float64(v5))) - -2.25) - 3.14159))
                    v10 := GenS1{F0: 97, F1: 111}
                    _ = v10
                } else {
                    var v11 uint16 = 135
                    _ = v11
                }
                _ = len(v3)
                var v12 int = 4294967295
                _ = v12
            } else {
                var v13 int8 = 0
                var v14 int = -203
                for fi15 := 0; fi15 < 5; fi15++ {
                    if false {
                        v16 := GenS1{F0: -88, F1: 0}
                        _ = v16
                    } else {
                        v13 -= (v13 + (v13 ^ (v13 ^ (v13 >> ((4294967295) & 63)))))
                    }
                    v17 := []int{(v14 % ((42) | 1))}
                    _ = fi15
                    _ = v17
                }
                _ = v13
                _ = v14
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                v5 = (v6 / ((v6 * (0.0 - 0.0)) + v5))
                if (false && ((true || ((1 == 494) || true)) || (-433 < -1))) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint8 = 1
                        v2 -= int32(v20)
                        var v21 float64 = (-2.25 + 0.0)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := []int{(v1 / (((-(fi18))) | 1))}
                    _ = v22
                }
                _ = fi18
            }
        }
    }
    fmt.Println(v5, v5)
    v3[v1] = 860
    v6 -= (v5 + v6)
    for fi23 := 0; fi23 < 5; fi23++ {
        if true {
            v5 = (v6 * v6)
            var v24 bool = false
            _ = v24
        } else {
            var v25 float32 = v6
            var v26 bool = ("foo bar" == ("a" + ((("\n\t" + "foo bar") + "\n\t") + ("\n\t" + ("a" + ("\n\t" + "\n\t"))))))
            _ = v25
            _ = v26
        }
        var v27 int32 = (v2 >> ((42) & 63))
        _ = fi23
        _ = v27
    }
    v28 := []int{(v1 / (((v1 >> ((4294967295) & 63))) | 1)), (v1 ^ (^(v1))), v1}
    v2 = (v2 - (v2 / ((v2) | 1)))
    for fi29 := 0; fi29 < 16; fi29++ {
        if (2 == 65535) {
            v30 := v4["foo bar"]
            if false {
                var v31 int8 = 42
                v32 := GenS1{F0: 0, F1: 4294967295}
                switch "" {
                case "a":
                    fmt.Println(v31, v30)
                case "foo bar":
                    if false {
                        switch ("\n\t" + v30) {
                        case "\n\t":
                            var v33 int16 = -1
                            _ = v33
                        case "hello":
                            if true {
                                if false {
                                    if (!(!("a" <= v30))) {
                                        if len(v28) > 0 {
                                            v28[0] = (v1 - (v1 - (^(fi29))))
                                        }
                                        v6 = v5
                                        switch v30 {
                                        case "foo bar":
                                            if ((0 <= 311) || false) {
                                                var v34 int16 = 32767
                                                fmt.Println(v31)
                                                _ = v34
                                            } else {
                                                fmt.Println(v2, v1, v1)
                                                v35 := GenS1{F0: v31, F1: 4294967295}
                                                for fi36 := 0; fi36 < 17; fi36++ {
                                                    if (!true) {
                                                        _ = len(v28)
                                                        v28 = append(v28, (fi36 << ((1) & 63)))
                                                        if len(v28) > 0 {
                                                            v28[0] = v1
                                                        }
                                                    } else {
                                                        v2 = (v2 & (^(v2)))
                                                    }
                                                    switch ((v30 + v30) + ("" + v30)) {
                                                    case "a":
                                                        var v37 float32 = (3.14159 + v6)
                                                        _ = v37
                                                    case "\n\t":
                                                        v28 = append(v28, -9223372036854775808)
                                                    }
                                                    _ = fi36
                                                }
                                                _ = v35
                                            }
                                        }
                                    } else {
                                        v31 -= v31
                                    }
                                    fmt.Println(v31)
                                    for fi38 := 0; fi38 < 7; fi38++ {
                                        delete(v3, (fi29 | (v1 / ((fi29) | 1))))
                                        _ = fi38
                                    }
                                }
                            }
                        default:
                            for fi39 := 0; fi39 < 17; fi39++ {
                                var v40 int16 = -902
                                _ = fi39
                                _ = v40
                            }
                        }
                        _ = len(v28)
                        v41 := GenS2{F0: v31, F1: v31, F2: 0}
                        _ = v41
                    }
                }
                _ = v31
                _ = v32
            }
            switch v30 {
            case "hello":
                v2 = 1
            case "foo bar":
                v42 := GenS1{F0: -1, F1: 4294967295}
                _ = v42
            }
            _ = v30
        } else {
            for fi43 := 0; fi43 < 8; fi43++ {
                var v44 float64 = 0.1
                if ("\n\t" < ((("foo bar" + "a") + ("a" + "a")) + "")) {
                    if true {
                        fmt.Println(v1, v44)
                        v1 -= v1
                        v1 += (fi29 - fi29)
                    } else {
                        v45 := v4["foo bar"]
                        _ = v45
                    }
                } else {
                    var v46 int64 = -9223372036854775808
                    _ = v46
                }
                switch fi43 {
                case -1:
                    switch fi29 {
                    case 1:
                        var v47 string = "foo bar"
                        _ = v47
                    case -36:
                        v48 := GenS1{F0: 1, F1: 1}
                        v49 := GenS3{F0: "foo bar"}
                        var v50 uint = 65535
                        if (2 >= 0) {
                            if false {
                                v51 := GenS3{F0: ((("foo bar" + "\n\t") + "\n\t") + "\n\t")}
                                v52 := GenS4{F0: 32767, F1: (v1 >> (((v50 & 4294967295)) & 63)), F2: true}
                                if ((true && true) || true) {
                                    var v53 int = 836
                                    v2 = (v2 * 1)
                                    _ = v53
                                }
                                _ = v51
                                _ = v52
                            } else {
                                v50 += (v50 + (v50 & (v50 >> ((0) & 63))))
                            }
                            v54 := GenS2{F0: 1, F1: -128, F2: 9223372036854775807}
                            fmt.Println(v44)
                            _ = v54
                        } else {
                            v1 = fi43
                        }
                        v5 = ((0.1 - float32(((3.14159 + 0.0) / float64((v5 * (v6 / -2.25)))))) / (v6 + 3.14159))
                        v55 := GenS5{F0: 1, F1: (v2 << ((v50) & 63))}
                        _ = v48
                        _ = v49
                        _ = v50
                        _ = v55
                    case -9223372036854775808:
                        delete(v3, int(v2))
                    }
                }
                var v56 float32 = v5
                var v57 uint = 517
                _ = fi43
                _ = v44
                _ = v56
                _ = v57
            }
        }
        for fi58 := 0; fi58 < 20; fi58++ {
            v1 -= (fi29 | -1000)
            var v59 uint64 = 0
            v6 -= ((v6 / (((v6 / v6) * (-2.25 / v6)) - 3.14159)) / ((v5 / v6) * (v5 + 0.1)))
            _ = fi58
            _ = v59
        }
        if (0 != -128) {
            for fi60 := 0; fi60 < 15; fi60++ {
                var v61 float64 = float64(v6)
                _ = fi60
                _ = v61
            }
            switch "\n\t" {
            case "foo bar":
                v5 = v5
            case "a":
                fmt.Println(v2, v2, v1)
            case "":
                if true {
                    v1 = (fi29 | -459)
                    switch v1 {
                    case 711:
                        v62 := GenS2{F0: 101, F1: 42, F2: -1}
                        _ = v62
                    case -2147483648:
                        var v63 uint8 = 0
                        _ = v63
                    case 2147483647:
                        v6 += ((3.14159 + v5) - float32(0.1))
                    }
                } else {
                    fmt.Println(fi29, v5)
                    for fi64 := 0; fi64 < 9; fi64++ {
                        v1 -= (fi29 - v1)
                        var v65 string = (("\n\t" + (("a" + "a") + "")) + "foo bar")
                        _ = fi64
                        _ = v65
                    }
                    if (219 == 65535) {
                        for fi66 := 0; fi66 < 8; fi66++ {
                            var v67 uint = 18446744073709551615
                            _ = fi66
                            _ = v67
                        }
                    } else {
                        var v68 uint64 = 18446744073709551615
                        fmt.Println(fi29, v5)
                        v69 := v3[-1]
                        _ = v68
                        _ = v69
                    }
                }
            default:
                switch "foo bar" {
                case "a":
                    var v70 float64 = 0.0
                    _ = v70
                }
            }
        }
        _ = fi29
    }
    fmt.Println(v5, v2, v6)
    if false {
        if true {
            fmt.Println(v5, v1)
        }
        v71 := GenS1{F0: -16, F1: 928}
        _ = v71
    }
    var v72 uint = 0
    v73 := GenS6{F0: 890}
    v74 := GenS4{F0: 2, F1: v1, F2: false}
    switch v1 {
    case 744:
        fmt.Println(v6, v5, v1)
    case 1:
        for fi75 := 0; fi75 < 20; fi75++ {
            var v76 uint32 = 65535
            var v77 int = 32767
            _ = fi75
            _ = v76
            _ = v77
        }
    case 32768:
        v2 -= 982
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v74
}
