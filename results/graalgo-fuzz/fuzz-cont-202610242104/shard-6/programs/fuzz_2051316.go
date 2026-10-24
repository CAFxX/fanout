package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint64; F2 int8 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint64 }
func genFn0() int8 {
    var v4 bool = ((((1e10 + -2.25) * 1e10) + ((0.1 * (0.1 / -2.25)) - 0.1)) == (-2.25 + ((-2.25 - 0.0) + 1.5)))
    v4 = (26 >= 65535)
    var v5 int64 = -639
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{-283}
        if len(v7) > 0 {
            v7[0] = fi6
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            switch (fi8 >> ((0) & 63)) {
            case -2147483648:
                v4 = v4
            case -119:
                switch fi8 {
                case 528:
                    var v9 int64 = (v5 >> ((18446744073709551615) & 63))
                    _ = v9
                case 1:
                    if ((1.5 - 1e10) == 0.1) {
                        fmt.Println(fi8)
                    } else {
                        var v10 uint16 = 389
                        var v11 int32 = 2147483647
                        _ = v10
                        _ = v11
                    }
                case 32767:
                    var v12 int32 = 0
                    _ = v12
                }
            }
            v5 -= 42
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 float32 = (1.5 + 0.1)
    v2 := []string{}
    var v3 bool = true
    if (0.0 == 1.5) {
        v13 := GenS1{F0: 2147483648, F1: 9223372036854775807, F2: -909}
        switch ((("" + ("a" + "a")) + "hello") + "hello") {
        case "a":
            fmt.Println(v3, v1)
        case "\n\t":
            var v14 bool = false
            _ = v14
        case "hello":
            if true {
                var v15 bool = v3
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS2{F0: 100}
                _ = fi16
                _ = v17
            }
        }
        _ = v13
    }
    if (100 >= 42) {
        v18 := GenS1{F0: 141, F1: 127, F2: 1}
        v19 := GenS3{F0: ("a" + "\n\t")}
        _ = v18
        _ = v19
    }
    switch -465 {
    case -923:
        v1 -= ((float32((0.0 + 3.14159)) * (v1 - ((v1 / v1) - 0.1))) - v1)
    case 1:
        var v20 int16 = -32768
        _ = v20
    }
    v21 := GenS3{F0: ""}
    var v22 uint32 = 1
    var v23 uint64 = 18446744073709551615
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 float64 = float64(v1)
        for fi26 := 0; fi26 < 4; fi26++ {
            v2 = append(v2, "")
            var v27 float32 = (v1 + v1)
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        if (float32(3.14159) != 1e10) {
            var v29 uint = 915
            if false {
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "foo bar")
                }
                v30 := GenS2{F0: 127}
                _ = v30
            } else {
                v29 += 614
                switch "\n\t" {
                case "a":
                    if v3 {
                        if false {
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        } else {
                            v23 += (v23 & (v23 % (((v23 ^ (^(v23)))) | 1)))
                            v29 += 65535
                            var v31 int8 = genFn0()
                            _ = v31
                        }
                        v32 := GenS4{F0: 42, F1: (^(v23)), F2: 127}
                        var v33 int8 = 1
                        _ = v32
                        _ = v33
                    } else {
                        v23 += (v23 ^ (v23 & (v23 | (v23 << (((v29 + v29)) & 63)))))
                        if (!v3) {
                            if len(v2) > 0 {
                                v2[0] = ("" + "foo bar")
                            }
                            if (2 > 891) {
                                if (!(850 != 1)) {
                                    fmt.Println(v29, v29)
                                    var v34 int16 = -32768
                                    _ = v34
                                }
                                if (127 != 113) {
                                    v35 := GenS5{F0: -863, F1: 4294967295}
                                    switch (^(fi28)) {
                                    case 4294967295:
                                        for fi36 := 0; fi36 < 9; fi36++ {
                                            var v37 int16 = 32767
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                            _ = fi36
                                            _ = v37
                                        }
                                    case -245:
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                    default:
                                        for fi38 := 0; fi38 < 15; fi38++ {
                                            var v39 uint32 = (v22 ^ v22)
                                            if v3 {
                                                var v40 uint = (^(v29))
                                                fmt.Println(v39)
                                                switch ("" + ("" + "\n\t")) {
                                                case "hello":
                                                    v41 := GenS6{F0: "hello", F1: 5}
                                                    _ = v41
                                                case "a":
                                                    if true {
                                                        v22 = v22
                                                    }
                                                    var v42 int16 = -57
                                                    var v43 uint16 = 42
                                                    _ = v42
                                                    _ = v43
                                                default:
                                                    if (-2.25 <= v1) {
                                                        v44 := GenS2{F0: int8(fi38)}
                                                        _ = v44
                                                    } else {
                                                        if v3 {
                                                            v45 := GenS7{F0: 0, F1: 4294967295}
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = ""
                                                            }
                                                            _ = v45
                                                        }
                                                        var v46 int16 = -32768
                                                        _ = v46
                                                    }
                                                }
                                                _ = v40
                                            }
                                            v3 = v3
                                            _ = fi38
                                            _ = v39
                                        }
                                    }
                                    v29 += 18446744073709551615
                                    _ = v35
                                } else {
                                    v23 = 137
                                    var v47 float32 = 1.5
                                    _ = v47
                                }
                                v48 := GenS5{F0: -614, F1: fi28}
                                _ = v48
                            } else {
                                if v3 {
                                    v49 := GenS8{F0: v23}
                                    _ = v49
                                } else {
                                    var v50 string = "hello"
                                    _ = v50
                                }
                            }
                            var v51 uint = 531
                            _ = v51
                        } else {
                            if true {
                                switch (fi28 << ((18446744073709551615) & 63)) {
                                case -81:
                                    v22 -= v22
                                default:
                                    var v52 int32 = -2147483648
                                    _ = v52
                                }
                                for fi53 := 0; fi53 < 8; fi53++ {
                                    fmt.Println(v1)
                                    var v54 int16 = 1000
                                    v3 = v3
                                    _ = fi53
                                    _ = v54
                                }
                            }
                            if v3 {
                                if ((v22 - (v22 << ((v29) & 63))) < (v22 / ((v22) | 1))) {
                                    var v55 uint32 = (v22 / (((v22 | v22)) | 1))
                                    v29 = 2
                                    _ = v55
                                } else {
                                    var v56 int32 = 2147483647
                                    v2 = append(v2, ("\n\t" + ("hello" + (("a" + "a") + ((("" + ("\n\t" + ("foo bar" + "\n\t"))) + "") + "")))))
                                    _ = v56
                                }
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v3, v23)
                                case "foo bar":
                                    var v57 uint = 0
                                    _ = v57
                                case "\n\t":
                                    var v58 uint = v29
                                    _ = v58
                                default:
                                    v29 -= (v29 | (v29 & (v29 / ((v29) | 1))))
                                }
                                switch 0 {
                                case -545:
                                    var v59 uint16 = 32767
                                    _ = v59
                                }
                            }
                        }
                    }
                case "foo bar":
                    v1 += v1
                case "hello":
                    v60 := GenS8{F0: v23}
                    _ = v60
                }
            }
            _ = v29
        } else {
            v61 := GenS5{F0: -9223372036854775808, F1: 0}
            v62 := GenS5{F0: 379, F1: 1}
            v1 += ((v1 + (1e10 + ((-2.25 / 1.5) - 0.1))) / v1)
            _ = v61
            _ = v62
        }
        _ = fi28
    }
    if len(v2) > 0 {
        v2[0] = ("a" + "hello")
    }
    for fi63 := 0; fi63 < 16; fi63++ {
        for fi64 := 0; fi64 < 12; fi64++ {
            var v65 bool = (!(2 > 2))
            _ = fi64
            _ = v65
        }
        fmt.Println(v22)
        v22 = 985
        _ = fi63
    }
    var v66 int8 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v66
}
