package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 bool; F2 int8 }
type GenS7 struct { F0 int16; F1 uint8 }
func genFn0() string {
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: "", F1: (fi8 | fi8), F2: 42}
        v7 = 1
        v7 += 4294967295
        _ = fi8
        _ = v9
    }
    _ = v7
    return "a"
}
func genFn1() int8 {
    if false {
        switch 1 {
        case 127:
            switch "foo bar" {
            case "\n\t":
                var v10 int32 = 1
                _ = v10
            case "a":
                v11 := []int{128, 4294967295}
                _ = v11
            default:
                var v12 int8 = 64
                _ = v12
            }
        default:
            var v13 float64 = -2.25
            _ = v13
        }
    } else {
        v14 := []int{32767, -457}
        _ = v14
    }
    var v15 uint = 940
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            if true {
                v18 := []int{(fi16 - fi17), 4294967295}
                var v19 float32 = 1e10
                v15 = (^(v15))
                _ = len(v18)
                _ = v18
                _ = v19
            }
            v15 = (^(v15))
            _ = fi17
        }
        _ = fi16
    }
    _ = v15
    return 1
}

func main() {
    var v1 float64 = (3.14159 - (3.14159 - -2.25))
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 uint16 = 42
    var v6 int16 = 946
    var v20 float32 = (((-2.25 - 0.0) * 1.5) - (1.5 + (((float32(v1) * 3.14159) + 0.0) * 3.14159)))
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int = 127
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(v1)
            switch "" {
            case "foo bar":
                v2 -= (v2 * ((v1 * v1) / 1.0))
            case "hello":
                fmt.Println(v2)
                if true {
                    _ = len(v4)
                }
                fmt.Println(v6, v5, fi23)
            default:
                var v25 uint8 = 2
                _ = v25
            }
            v26 := []int{fi23, -859, (fi24 & (fi23 + (^(fi23)))), (fi23 | fi24)}
            _ = fi24
            _ = v26
        }
        switch fi23 {
        case 4294967295:
            switch 4294967295 {
            case 42:
                v5 = v5
            case 1:
                var v27 int = (-(fi23))
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            fmt.Println(fi23, v6, v1)
            v29 := GenS1{F0: "hello", F1: fi23, F2: 2}
            var v30 int64 = 9223372036854775807
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi23
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        if true {
            if ((true && false) && (!(uint32(v5) <= 2147483648))) {
                var v32 uint8 = 128
                switch (fi31 | (-(fi31))) {
                case -433:
                    v33 := GenS2{F0: ((((v1 * 0.0) * v2) * (1e10 / (((((-2.25 / ((v2 + v2) + float64(v20))) / 3.14159) + 1e10) + v2) / float64((v20 + 0.1))))) <= ((1.5 - 0.1) - (v2 - 0.1)))}
                    _ = v33
                }
                v34 := GenS3{F0: 0, F1: -632}
                _ = v32
                _ = v34
            }
            var v35 uint64 = 435
            _ = v35
        }
        var v36 int = fi31
        _ = fi31
        _ = v36
    }
    v37 := GenS4{F0: 2147483647, F1: 55}
    v38 := []int{487, -934, 9223372036854775807}
    v39 := GenS5{F0: -634}
    if len(v38) > 0 {
        v38[0] = -297
    }
    v1 += v2
    for fi40 := 0; fi40 < 1; fi40++ {
        fmt.Println(v2, v20)
        var v41 bool = (false || true)
        v1 += v2
        _ = fi40
        _ = v41
    }
    v2 = v1
    switch 32768 {
    case 167:
        switch "\n\t" {
        case "hello":
            v2 += (v1 / float64((((v20 * (v20 * -2.25)) - (v20 / v20)) + v20)))
        case "":
            var v42 int32 = -321
            _ = v42
        default:
            v2 -= (v2 - 3.14159)
        }
    }
    if ((v5 + v5) == v5) {
        if (true || (int16(v5) <= (-(v6)))) {
            v6 += 0
            var v43 uint = 255
            v2 = v2
            _ = v43
        }
    }
    var v44 uint16 = v5
    if ((386 < 1) && true) {
        fmt.Println(v1, v5, v44)
    }
    if ((^(v5)) > (v5 >> ((42) & 63))) {
        fmt.Println(v2, v44, v20)
        if (42 > 0) {
            if true {
                if false {
                    var v45 float64 = v1
                    _ = v45
                } else {
                    v44 -= (v5 * 2)
                }
            } else {
                for fi46 := 0; fi46 < 17; fi46++ {
                    v6 -= 0
                    fmt.Println(v5, v6, v1)
                    var v47 int32 = 42
                    _ = fi46
                    _ = v47
                }
            }
            if (!false) {
                v38 = append(v38, 127)
                v5 += (v44 - 42)
            }
        }
    }
    if (true || (2147483647 < 2)) {
        fmt.Println(v6, v20, v6)
        v6 = (v6 % ((-32768) | 1))
    } else {
        var v48 uint = 175
        for fi49 := 0; fi49 < 10; fi49++ {
            v20 = v20
            if false {
                var v50 int16 = (v6 << (((^(v48))) & 63))
                var v51 int = (fi49 | (fi49 + fi49))
                if false {
                    if (!(-32768 != v6)) {
                        _ = len(v38)
                        v44 += 32767
                    } else {
                        switch "a" {
                        case "foo bar":
                            for fi52 := 0; fi52 < 1; fi52++ {
                                if (-9223372036854775808 <= 765) {
                                    if (2147483647 > 1) {
                                        v44 -= (v44 - (^(v5)))
                                    } else {
                                        var v53 uint16 = v44
                                        v38 = append(v38, v51)
                                        var v54 uint8 = 0
                                        _ = v53
                                        _ = v54
                                    }
                                    if ((((("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + "") + "a") + ("a" + "")) == ("\n\t" + "foo bar")) {
                                        if (false && (1 <= 905)) {
                                            var v55 bool = (365 != 42)
                                            var v56 bool = (!v55)
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v48 = 214
                                            if false {
                                                v5 += v44
                                                v44 = 1
                                            }
                                            var v57 int32 = -967
                                            _ = v57
                                        }
                                        v58 := GenS6{F0: fi52, F1: true, F2: 1}
                                        _ = v58
                                    }
                                }
                                switch (("foo bar" + "") + "") {
                                case "hello":
                                    fmt.Println(v50)
                                case "\n\t":
                                    if (0 <= 100) {
                                        var v59 uint64 = 56
                                        v44 -= (v44 ^ (v5 ^ v44))
                                        if false {
                                            v5 = (v5 / (((v44 + 65535)) | 1))
                                            var v60 uint32 = 371
                                            _ = v60
                                        }
                                        _ = v59
                                    } else {
                                        v1 = 3.14159
                                        v1 -= ((v2 / v1) * ((v1 + v2) + -2.25))
                                    }
                                case "a":
                                    var v61 uint64 = 987
                                    _ = v61
                                }
                                _ = fi52
                            }
                        case "":
                            switch "" {
                            case "foo bar":
                                var v62 int8 = 42
                                _ = v62
                            case "\n\t":
                                fmt.Println(v48, v6)
                            }
                        case "a":
                            var v63 int = (v51 - fi49)
                            _ = v63
                        }
                    }
                    v4[0] = (("\n\t" + ("\n\t" + "hello")) + ("\n\t" + (("a" + "foo bar") + ("\n\t" + "foo bar"))))
                    switch -502 {
                    case 762:
                        v64 := GenS7{F0: 1, F1: 30}
                        _ = v64
                    case 0:
                        switch fi49 {
                        case 903:
                            var v65 uint = v48
                            _ = v65
                        }
                    }
                }
                _ = v50
                _ = v51
            }
            var v66 int8 = -128
            _ = fi49
            _ = v66
        }
        var v67 uint8 = 0
        _ = v48
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
