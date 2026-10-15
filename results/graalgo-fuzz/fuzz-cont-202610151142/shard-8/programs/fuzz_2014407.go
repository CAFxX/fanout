package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0() int16 {
    var v5 uint = 913
    v5 = (v5 << ((v5) & 63))
    _ = v5
    return 0
}
func genFn1() uint {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = (3.14159 / 1.0)
        _ = fi6
        _ = v7
    }
    var v8 int8 = 115
    if true {
        var v9 int64 = 42
        _ = v9
    } else {
        var v10 float64 = (float64(((-2.25 - float32(-2.25)) + 3.14159)) - 1.5)
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint32 = 0
        v8 += -128
        var v13 uint64 = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 string = ("a" + "hello")
        v8 += (-(v8))
        for fi16 := 0; fi16 < 20; fi16++ {
            if (2147483648 <= 2147483648) {
                if true {
                    v17 := []int{(fi16 | 127), fi16, (fi16 & fi14), fi14}
                    fmt.Println(v15, v8, v8)
                    _ = v17
                }
                var v18 bool = (599 == int16(fi16))
                _ = v18
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    _ = v8
    return 788
}

func main() {
    var v1 uint32 = 394
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v4 := make(map[string]bool)
    v1 = v1
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            if false {
                if ((1 > 44) || (211 >= 1)) {
                    v3 += 42
                    var v21 bool = false
                    _ = v21
                }
                var v22 int8 = 106
                _ = v22
            }
            v23 := GenS1{F0: 797, F1: -32768, F2: (-(fi20))}
            v3 += 0
            _ = fi20
            _ = v23
        }
        _ = fi19
    }
    fmt.Println(v1, v3)
    switch "" {
    case "":
        var v24 uint16 = 167
        _ = v24
    case "hello":
        if (1 < 0) {
            if false {
                _ = len(v2)
                v1 += v1
            } else {
                fmt.Println(v1, v3)
                v3 += v3
            }
            v3 += v3
            if (true || (-1 > 37)) {
                if (("foo bar" + "a") >= "") {
                    switch 664 {
                    case 127:
                        v1 -= 4294967295
                    default:
                        var v25 int8 = 100
                        _ = v25
                    }
                    if (float64(float32(0.0)) >= -2.25) {
                        v26 := GenS2{F0: true, F1: 0, F2: 4294967295}
                        v27 := GenS1{F0: -75, F1: -1, F2: 128}
                        v1 = v1
                        _ = v26
                        _ = v27
                    } else {
                        var v28 int32 = (v3 << ((4294967295) & 63))
                        _ = v28
                    }
                }
                v3 = int32(v1)
                if (990 == 573) {
                    var v29 int = 128
                    for fi30 := 0; fi30 < 7; fi30++ {
                        for fi31 := 0; fi31 < 6; fi31++ {
                            if true {
                                var v32 string = ""
                                _ = v32
                            }
                            _ = fi31
                        }
                        _ = fi30
                    }
                    _ = v29
                } else {
                    var v33 int64 = -977
                    _ = v33
                }
            } else {
                if (2 != 32767) {
                    var v34 int64 = -9223372036854775808
                    if false {
                        var v35 uint16 = 718
                        for fi36 := 0; fi36 < 16; fi36++ {
                            var v37 uint = 686
                            _ = fi36
                            _ = v37
                        }
                        _ = v35
                    } else {
                        v1 = v1
                        var v38 int8 = 127
                        for fi39 := 0; fi39 < 10; fi39++ {
                            if ((!(!true)) && ((3.14159 / (float32(0.0) + 1e10)) >= -2.25)) {
                                var v40 uint32 = v1
                                var v41 int = (fi39 | fi39)
                                _ = v40
                                _ = v41
                            } else {
                                v42 := []int{(fi39 % ((32768) | 1))}
                                if len(v42) >= 2 {
                                    v42 = v42[1:len(v42)]
                                }
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 16; fi43++ {
                                var v44 uint8 = 128
                                v45 := GenS3{F0: (fi43 ^ fi43)}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            _ = fi39
                        }
                        _ = v38
                    }
                    _ = v34
                }
                v46 := GenS4{F0: -9223372036854775808, F1: uint(v1)}
                _ = v46
            }
        }
    }
    var v47 int8 = 42
    for fi48 := 0; fi48 < 1; fi48++ {
        var v49 int64 = int64(v1)
        var v50 string = ("\n\t" + ("a" + "\n\t"))
        var v51 float32 = float32(1.5)
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        _ = len(v4)
        for fi53 := 0; fi53 < 9; fi53++ {
            v47 = 1
            var v54 uint32 = v1
            v55 := []int{(fi52 ^ (-(fi52)))}
            switch (fi53 % ((32767) | 1)) {
            case -705:
                v47 -= (v47 | (^(v47)))
            case -749:
                fmt.Println(v54, fi52, fi52)
            default:
                v3 -= v3
            }
            var v56 int64 = -9223372036854775808
            _ = fi53
            _ = v54
            _ = v55
            _ = v56
        }
        _ = fi52
    }
    var v57 uint32 = 2147483648
    switch "a" {
    case "\n\t":
        v4["foo bar"] = (genFn1() == 1)
    }
    fmt.Println(v3)
    var v58 uint32 = v1
    if true {
        fmt.Println(v47, v1, v47)
    } else {
        var v59 uint = 2
        v60 := GenS5{F0: (v3 >> ((v59) & 63))}
        _ = v59
        _ = v60
    }
    v61 := GenS6{F0: 127, F1: -667, F2: 65535}
    v58 = v1
    if true {
        fmt.Println(v3, v47, v57)
        if true {
            delete(v2, v57)
            switch -2147483648 {
            case -1:
                fmt.Println(v47, v47)
            }
        } else {
            v57 -= (^(v57))
        }
    }
    var v62 uint16 = 65535
    v63 := v4["a"]
    for fi64 := 0; fi64 < 20; fi64++ {
        var v65 uint64 = 2
        delete(v2, v57)
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v57
    _ = v58
    _ = v61
    _ = v62
    _ = v63
}
