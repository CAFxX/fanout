package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32) int64 {
    v6 := []int{258, 2147483647, -2147483648, 282}
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    var v7 uint16 = 32767
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int8 = 0
        _ = fi8
        _ = v9
    }
    switch "" {
    case "":
        for fi10 := 0; fi10 < 9; fi10++ {
            if (!((v7 >> ((1) & 63)) == v7)) {
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if true {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            var v12 int32 = p0
                            _ = fi11
                            _ = v12
                        }
                    }
                    var v13 string = ""
                    _ = v13
                } else {
                    if len(v6) > 0 {
                        v6[0] = (^(fi10))
                    }
                    p0 = (p0 >> ((1) & 63))
                }
                switch (fi10 & fi10) {
                case 9223372036854775807:
                    v14 := GenS1{F0: p0, F1: "\n\t", F2: 1}
                    _ = v14
                default:
                    var v15 int16 = int16(p0)
                    _ = v15
                }
            }
            _ = fi10
        }
    case "foo bar":
        var v16 uint8 = 186
        _ = v16
    case "a":
        var v17 uint8 = 2
        _ = v17
    default:
        if len(v6) > 0 {
            v6[0] = 0
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint32 = 378
    v2 := []int{1, 962, -2147483648}
    var v3 float64 = -2.25
    v4 := make(map[int]string)
    var v5 int32 = -2147483648
    switch 0 {
    case -9223372036854775808:
        v18 := GenS2{F0: ("" + "\n\t")}
        _ = v18
    default:
        v5 = v5
    }
    v3 += ((1e10 - v3) * (1.5 - v3))
    _ = len(v2)
    var v19 uint = 2
    var v20 int16 = -155
    delete(v4, 740)
    switch 32768 {
    case 1:
        fmt.Println(v1, v5, v3)
    default:
        v19 = (v19 << (((^(v19))) & 63))
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(fi21)
        if (false && (false || false)) {
            for fi22 := 0; fi22 < 3; fi22++ {
                delete(v4, (-(fi22)))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch fi22 {
                case 32768:
                    var v23 uint32 = (^(v1))
                    _ = v23
                case -713:
                    var v24 uint = v19
                    _ = v24
                }
                if true {
                    if false {
                        v25 := GenS3{F0: (v19 % (((v19 - v19)) | 1)), F1: (fi21 ^ 128)}
                        fmt.Println(v19, v19)
                        _ = v25
                    }
                    switch fi21 {
                    case 406:
                        if (("a" + "foo bar") > "") {
                            v19 = (v19 - (^(v19)))
                            v26 := GenS4{F0: 42}
                            _ = v26
                        } else {
                            v20 = 431
                            fmt.Println(v3)
                            var v27 uint32 = (^(v1))
                            _ = v27
                        }
                        v28 := GenS2{F0: "\n\t"}
                        v5 = v5
                        _ = v28
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v5 = (v5 / ((v5) | 1))
                }
                if true {
                    fmt.Println(v3, v1)
                    var v29 bool = false
                    var v30 float64 = float64(1.5)
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(fi21, v19)
                }
                fmt.Println(v1, v3, v3)
                _ = fi22
            }
            for fi31 := 0; fi31 < 7; fi31++ {
                if true {
                    v1 += v1
                }
                v32 := GenS1{F0: 2147483647, F1: ("hello" + ("\n\t" + "\n\t")), F2: 648}
                _ = fi31
                _ = v32
            }
        } else {
            fmt.Println(v20)
        }
        _ = fi21
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v1 = (v1 << (((v19 | (v19 % (((v19 * (v19 / ((18446744073709551615) | 1)))) | 1)))) & 63))
        if (((false && true) || false) || true) {
            v3 = v3
        } else {
            v5 = (v5 + -2147483648)
            v34 := GenS5{F0: 0}
            for fi35 := 0; fi35 < 15; fi35++ {
                if false {
                    var v36 bool = (!(-9223372036854775808 != 42))
                    v2 = append(v2, fi35)
                    var v37 int32 = (^(v5))
                    _ = v36
                    _ = v37
                } else {
                    v38 := GenS6{F0: (v5 * v5), F1: (v5 | (v5 / (((v5 % (((v5 + (v5 - -252))) | 1))) | 1))), F2: 0}
                    delete(v4, fi33)
                    v39 := GenS7{F0: 731, F1: 0}
                    _ = v38
                    _ = v39
                }
                _ = fi35
            }
            _ = v34
        }
        _ = len(v4)
        _ = fi33
    }
    var v40 int64 = 954
    switch "\n\t" {
    case "\n\t":
        var v41 uint8 = 128
        _ = v41
    case "":
        for fi42 := 0; fi42 < 13; fi42++ {
            for fi43 := 0; fi43 < 17; fi43++ {
                fmt.Println(v1, v19, v3)
                _ = fi43
            }
            v19 = (v19 >> ((v19) & 63))
            for fi44 := 0; fi44 < 5; fi44++ {
                var v45 int16 = int16(v1)
                v46 := GenS4{F0: 1}
                _ = fi44
                _ = v45
                _ = v46
            }
            _ = fi42
        }
    default:
        for fi47 := 0; fi47 < 2; fi47++ {
            for fi48 := 0; fi48 < 20; fi48++ {
                v49 := GenS1{F0: v5, F1: "foo bar", F2: 1}
                var v50 uint32 = v1
                fmt.Println(v5, v19)
                _ = fi48
                _ = v49
                _ = v50
            }
            v51 := GenS7{F0: 0, F1: 11}
            _ = fi47
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
