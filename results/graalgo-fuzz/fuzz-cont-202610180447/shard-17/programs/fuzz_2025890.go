package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0(p0 int16, p1 uint32, p2 uint64) int16 {
    var v7 uint32 = (p1 | (p1 % ((2) | 1)))
    var v8 int32 = 1
    v9 := []int{4294967295, -1, -148}
    if len(v9) > 0 {
        v9[0] = 0
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 34
}
func genFn1(p0 uint16) string {
    p0 -= p0
    p0 -= 127
    if (4294967295 > -2147483648) {
        if true {
            var v10 bool = (true && (true && false))
            fmt.Println(v10, v10)
            var v11 uint32 = uint32(p0)
            _ = v10
            _ = v11
        } else {
            if (true && false) {
                var v12 uint = 2
                switch "foo bar" {
                case "foo bar":
                    v13 := GenS1{F0: 214, F1: false}
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 11; fi14++ {
                        switch ("" + ("\n\t" + (("foo bar" + "hello") + "hello"))) {
                        case "foo bar":
                            v15 := GenS1{F0: 32768, F1: (32767 == -32768)}
                            _ = v15
                        case "a":
                            v16 := []int{(^(fi14)), fi14}
                            _ = v16
                        case "hello":
                            var v17 int32 = 555
                            _ = v17
                        default:
                            var v18 float64 = 1.5
                            _ = v18
                        }
                        if (1.5 > -2.25) {
                            if (false && (false || ((v12 + v12) != 2))) {
                                var v19 uint8 = 2
                                if (!false) {
                                    v20 := GenS1{F0: 499, F1: (true || (1e10 > -2.25))}
                                    _ = v20
                                } else {
                                    fmt.Println(fi14, v12, fi14)
                                    p0 = (p0 & (^(p0)))
                                    var v21 bool = true
                                    _ = v21
                                }
                                fmt.Println(fi14, v12)
                                _ = v19
                            }
                            for fi22 := 0; fi22 < 15; fi22++ {
                                p0 -= p0
                                _ = fi22
                            }
                        } else {
                            v23 := GenS1{F0: fi14, F1: false}
                            v12 += (v12 / ((uint(p0)) | 1))
                            for fi24 := 0; fi24 < 12; fi24++ {
                                var v25 int16 = 32767
                                v12 += 173
                                var v26 uint32 = 677
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            _ = v23
                        }
                        _ = fi14
                    }
                case "hello":
                    if false {
                        for fi27 := 0; fi27 < 9; fi27++ {
                            v28 := GenS1{F0: (fi27 ^ (fi27 | fi27)), F1: true}
                            p0 += 2
                            _ = fi27
                            _ = v28
                        }
                        switch 2147483647 {
                        case -1:
                            var v29 uint32 = 2147483648
                            _ = v29
                        default:
                            for fi30 := 0; fi30 < 4; fi30++ {
                                for fi31 := 0; fi31 < 11; fi31++ {
                                    var v32 uint64 = 0
                                    v32 = v32
                                    _ = fi31
                                    _ = v32
                                }
                                if (42 == -1) {
                                    var v33 string = (("\n\t" + "\n\t") + "")
                                    var v34 uint16 = 255
                                    _ = v33
                                    _ = v34
                                } else {
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        for fi35 := 0; fi35 < 5; fi35++ {
                                            v12 -= (v12 >> ((v12) & 63))
                                            var v36 uint8 = 0
                                            if (!(0 != 0)) {
                                                fmt.Println(fi35, fi35, v36)
                                                if (!((!(!true)) && true)) {
                                                    v37 := []int{4294967295, fi30, fi30}
                                                    var v38 uint8 = 177
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    switch (^(fi30)) {
                                                    case 149:
                                                        v39 := []int{}
                                                        _ = v39
                                                    case 2147483647:
                                                        v40 := GenS1{F0: (fi35 >> ((42) & 63)), F1: (int64(fi30) <= 1)}
                                                        _ = v40
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v12, v12)
                                                v12 = uint(v36)
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                    }
                                    v41 := GenS2{F0: 267, F1: 4294967295, F2: 1}
                                    if true {
                                        fmt.Println(p0)
                                    }
                                    _ = v41
                                }
                                var v42 float64 = 1.5
                                _ = fi30
                                _ = v42
                            }
                        }
                    }
                default:
                    p0 -= (p0 + 220)
                }
                v43 := GenS1{F0: 4294967295, F1: ((-85 == 127) && (2147483648 > 42))}
                _ = v12
                _ = v43
            }
            var v44 float64 = ((3.14159 + (((1e10 + 3.14159) + 0.1) + 0.0)) - (float64((1.5 - (0.0 / 1.0))) / (1.5 / (1.5 - -2.25))))
            v45 := GenS2{F0: 128, F1: 2, F2: 42}
            _ = v44
            _ = v45
        }
        p0 -= 255
    } else {
        fmt.Println(p0)
    }
    v46 := GenS1{F0: 2147483647, F1: false}
    if false {
        p0 += (p0 + (p0 + p0))
        v47 := []int{}
        var v48 int8 = -128
        _ = v47
        _ = v48
    }
    _ = p0
    _ = v46
    return "\n\t"
}

func main() {
    var v1 uint8 = 1
    v2 := []int{int(v1), -699, -9223372036854775808, 128}
    var v3 uint32 = 2147483647
    var v4 int = 58
    v5 := []string{"foo bar", "a", ("foo bar" + "hello"), "hello"}
    v6 := []int{-2147483648, v4, (v4 - (v4 ^ v4)), (v4 | (v4 - (v4 - v4)))}
    v4 = v4
    var v49 float64 = 1e10
    var v50 uint32 = v3
    for fi51 := 0; fi51 < 1; fi51++ {
        fmt.Println(fi51, v3)
        if true {
            if true {
                _ = len(v2)
            } else {
                v3 -= (v3 / (((v50 - 42)) | 1))
            }
            switch "\n\t" {
            case "":
                for fi52 := 0; fi52 < 5; fi52++ {
                    var v53 bool = true
                    _ = fi52
                    _ = v53
                }
            case "hello":
                var v54 int16 = -112
                _ = v54
            case "\n\t":
                _ = len(v5)
            default:
                if (false && true) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
        for fi55 := 0; fi55 < 5; fi55++ {
            if false {
                if (!false) {
                    v1 = (v1 % ((v1) | 1))
                } else {
                    var v56 int32 = 0
                    _ = v56
                }
                switch -220 {
                case 2147483647:
                    var v57 int = fi51
                    _ = v57
                case -609:
                    fmt.Println(fi55, v3)
                case 976:
                    v3 += (v3 ^ (v50 + uint32(v1)))
                default:
                    v58 := GenS2{F0: -2147483648, F1: 0, F2: 42}
                    _ = v58
                }
            } else {
                _ = len(v6)
                var v59 int = (fi51 | v4)
                _ = v59
            }
            if false {
                if true {
                    if (-949 > 0) {
                        v60 := GenS2{F0: v4, F1: 377, F2: -1}
                        _ = v60
                    } else {
                        fmt.Println(v1, v4, v50)
                    }
                    if true {
                        var v61 float64 = v49
                        _ = v61
                    } else {
                        if (255 <= 0) {
                            var v62 int32 = 2147483647
                            v63 := GenS3{F0: uint32(v1), F1: 165}
                            _ = v62
                            _ = v63
                        }
                        var v64 int16 = -32768
                        v3 = 2
                        _ = v64
                    }
                    v65 := GenS4{F0: (4294967295 < 4294967295)}
                    v1 = (^(v1))
                    var v66 int16 = 681
                    var v67 uint8 = v1
                    _ = v65
                    _ = v66
                    _ = v67
                } else {
                    var v68 uint = 550
                    _ = v68
                }
                v69 := GenS5{F0: 2, F1: (v1 - 42), F2: 4294967295}
                _ = v69
            } else {
                v49 -= v49
            }
            if len(v6) > 0 {
                v6[0] = fi51
            }
            _ = fi55
        }
        _ = fi51
    }
    v3 -= 4294967295
    fmt.Println(v4, v49, v4)
    v70 := GenS3{F0: v50, F1: 18446744073709551615}
    switch "hello" {
    case "a":
        var v71 int64 = 1
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v70
}
