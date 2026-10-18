package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int8; F2 int64 }
type GenS8 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0() string {
    var v7 uint = 42
    if (("a" + "a") == "hello") {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 3; fi8++ {
            v7 = v7
            var v9 uint32 = 338
            _ = fi8
            _ = v9
        }
        var v10 uint = v7
        _ = v10
    }
    _ = v7
    return "\n\t"
}
func genFn1(p0 string) int64 {
    v11 := GenS1{F0: 0, F1: -128, F2: 255}
    p0 = "foo bar"
    _ = p0
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := []int{1, 1}
    var v4 uint64 = 502
    var v5 int32 = 42
    var v6 uint16 = 255
    switch "foo bar" {
    case "foo bar":
        switch int(v6) {
        case 1:
            var v12 float32 = 0.1
            _ = v12
        default:
            if (!(716 == uint(v6))) {
                v3 = append(v3, 9223372036854775807)
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int32 = v5
                    v15 := GenS2{F0: v4, F1: (v6 * v6), F2: "\n\t"}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 int32 = -1
                _ = v16
            } else {
                var v17 float64 = 1.5
                if false {
                    var v18 float64 = 1.5
                    v5 += v5
                    _ = v18
                } else {
                    var v19 string = ("a" + (("a" + "") + ("" + "")))
                    if (!(65535 < v6)) {
                        if false {
                            v19 = v19
                            var v20 int64 = 0
                            var v21 string = v19
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v5, v4, v17)
                            fmt.Println(v4, v5)
                            var v22 uint = 205
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 1; fi23++ {
                            _ = len(v1)
                            _ = fi23
                        }
                        if (42 != 1) {
                            if false {
                                v24 := GenS1{F0: 668, F1: -92, F2: 255}
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    var v26 uint64 = (v4 / ((v4) | 1))
                                    _ = fi25
                                    _ = v26
                                }
                                v27 := GenS3{F0: -50, F1: v6, F2: 475}
                                v28 := GenS3{F0: -1, F1: (v6 << ((65535) & 63)), F2: 2147483648}
                                _ = v27
                                _ = v28
                            }
                        }
                    }
                    v29 := GenS3{F0: 100, F1: v6, F2: 872}
                    _ = v19
                    _ = v29
                }
                _ = v17
            }
        }
    default:
        v5 += v5
    }
    v4 -= (v4 * v4)
    switch 403 {
    case -9223372036854775808:
        v4 = (v4 / ((v4) | 1))
    case 32768:
        v30 := GenS3{F0: -63, F1: (^(v6)), F2: 1}
        _ = v30
    case 42:
        if (32767 != -1) {
            fmt.Println(v4, v4)
            v6 = 146
            if false {
                switch "" {
                case "a":
                    switch ((("" + ("" + "")) + "") + ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + "a"))) {
                    case "hello":
                        switch "hello" {
                        case "hello":
                            for fi31 := 0; fi31 < 17; fi31++ {
                                var v32 float64 = 0.0
                                fmt.Println(v6, v32)
                                fmt.Println(v5)
                                _ = fi31
                                _ = v32
                            }
                        }
                    default:
                        fmt.Println(v4, v4)
                    }
                case "foo bar":
                    if false {
                        var v33 uint8 = 128
                        _ = v33
                    } else {
                        if false {
                            switch ("foo bar" + "") {
                            case "a":
                                v5 += v5
                            case "foo bar":
                                v3 = append(v3, 42)
                            case "\n\t":
                                var v34 float64 = float64(3.14159)
                                _ = v34
                            }
                        }
                    }
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v35 := GenS4{F0: 32767, F1: v4}
                _ = v35
            } else {
                for fi36 := 0; fi36 < 5; fi36++ {
                    switch ("foo bar" + (("foo bar" + (("" + ("a" + ((("" + "\n\t") + "\n\t") + ("hello" + "")))) + ("a" + ""))) + ("" + "hello"))) {
                    case "foo bar":
                        fmt.Println(v4, v5, v6)
                    case "\n\t":
                        v37 := v1[(^(fi36))]
                        _ = v37
                    case "a":
                        v4 -= v4
                    default:
                        fmt.Println(fi36, v5, v5)
                    }
                    var v38 int32 = 821
                    v4 = (v4 % (((v4 & v4)) | 1))
                    _ = fi36
                    _ = v38
                }
                var v39 float32 = 1.5
                _ = v39
            }
        } else {
            if false {
                if true {
                    var v40 float64 = (1.5 + (3.14159 - 3.14159))
                    switch "hello" {
                    case "":
                        for fi41 := 0; fi41 < 19; fi41++ {
                            v4 = (v4 << ((512) & 63))
                            fmt.Println(v40, v4)
                            fmt.Println(v40, v6)
                            _ = fi41
                        }
                    case "a":
                        for fi42 := 0; fi42 < 20; fi42++ {
                            for fi43 := 0; fi43 < 14; fi43++ {
                                switch "foo bar" {
                                case "a":
                                    var v44 uint32 = 4294967295
                                    _ = v44
                                case "":
                                    var v45 uint8 = 1
                                    v46 := GenS2{F0: v4, F1: (v6 | v6), F2: "hello"}
                                    v4 = 1
                                    _ = v45
                                    _ = v46
                                case "foo bar":
                                    var v47 int8 = 1
                                    var v48 uint8 = 1
                                    var v49 float32 = (1e10 + 0.1)
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                default:
                                    var v50 float64 = (-2.25 * v40)
                                    _ = v50
                                }
                                _ = fi43
                            }
                            for fi51 := 0; fi51 < 7; fi51++ {
                                v52 := GenS3{F0: 42, F1: v6, F2: 42}
                                v40 = (v40 * float64(3.14159))
                                _ = fi51
                                _ = v52
                            }
                            v53 := GenS4{F0: 1, F1: v4}
                            _ = fi42
                            _ = v53
                        }
                    case "hello":
                        if true {
                            if false {
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    for fi55 := 0; fi55 < 9; fi55++ {
                                        switch ("hello" + "") {
                                        case "foo bar":
                                            v4 -= (v4 * (v4 - v4))
                                        }
                                        var v56 int8 = 100
                                        var v57 uint64 = v4
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                    v6 -= (v6 << ((4294967295) & 63))
                                    _ = fi54
                                }
                                v5 += v5
                            }
                        } else {
                            v58 := GenS4{F0: 1000, F1: 4294967295}
                            _ = v58
                        }
                    default:
                        for fi59 := 0; fi59 < 16; fi59++ {
                            var v60 int32 = v5
                            _ = fi59
                            _ = v60
                        }
                    }
                    v40 = 0.1
                    _ = v40
                } else {
                    var v61 bool = true
                    if (!true) {
                        var v62 uint16 = (v6 << ((65535) & 63))
                        v63 := GenS5{F0: v5, F1: 65535}
                        _ = v62
                        _ = v63
                    } else {
                        v64 := GenS2{F0: v4, F1: 42, F2: ("" + "")}
                        switch 0 {
                        case 364:
                            var v65 bool = true
                            _ = v65
                        case -1:
                            var v66 int = 127
                            _ = v66
                        default:
                            v4 += v4
                        }
                        _ = v64
                    }
                    _ = v61
                }
                var v67 float32 = 0.1
                _ = v67
            } else {
                if len(v3) > 0 {
                    v3[0] = -895
                }
                v68 := GenS6{F0: "a"}
                v5 = v5
                _ = v68
            }
            for fi69 := 0; fi69 < 1; fi69++ {
                fmt.Println(v4, v6)
                var v70 uint32 = 2147483648
                _ = len(v3)
                _ = fi69
                _ = v70
            }
            v6 -= v6
        }
    default:
        v71 := GenS4{F0: 32767, F1: v4}
        _ = v71
    }
    fmt.Println(v4, v4, v6)
    if false {
        if true {
            delete(v2, uint32(v6))
            var v72 uint32 = 147
            _ = v72
        } else {
            v6 = 2
        }
    }
    v4 += 96
    v73 := GenS2{F0: uint64(v6), F1: v6, F2: "hello"}
    for fi74 := 0; fi74 < 12; fi74++ {
        v5 += (v5 % ((84) | 1))
        v75 := v2[291]
        v6 += (v6 << ((43) & 63))
        _ = fi74
        _ = v75
    }
    v5 = v5
    if ((-105 < 127) || ("" < "a")) {
        if true {
            fmt.Println(v4, v6)
            var v76 string = "\n\t"
            _ = v76
        }
        v6 += 255
        v5 += v5
    } else {
        v77 := GenS7{F0: 1, F1: 42, F2: -9223372036854775808}
        var v78 int16 = -799
        switch 128 {
        case 32767:
            for fi79 := 0; fi79 < 5; fi79++ {
                v80 := v1[(fi79 * fi79)]
                v5 += v5
                _ = fi79
                _ = v80
            }
        case -9223372036854775808:
            v81 := GenS8{F0: -128, F1: -342, F2: v6}
            _ = v81
        default:
            var v82 uint8 = 2
            _ = v82
        }
        _ = v77
        _ = v78
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
