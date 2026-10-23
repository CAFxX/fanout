package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 int16, p1 int) uint8 {
    p1 += (p1 ^ p1)
    var v6 float64 = (float64((1.5 / 3.14159)) * ((1.5 / 1.5) - 0.1))
    p1 += (-(p1))
    v7 := []int{-141, p1, p1, 27}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int8 = 0
        for fi10 := 0; fi10 < 2; fi10++ {
            _ = len(v7)
            var v11 int64 = 1
            var v12 string = "a"
            _ = fi10
            _ = v11
            _ = v12
        }
        switch p1 {
        case 1:
            if false {
                if ("foo bar" != "a") {
                    v13 := GenS1{F0: ("" + "a")}
                    for fi14 := 0; fi14 < 13; fi14++ {
                        if (true || (!false)) {
                            p1 = p1
                            v9 = (v9 % (((v9 | 42)) | 1))
                        } else {
                            var v15 string = "foo bar"
                            v16 := GenS1{F0: "a"}
                            _ = v15
                            _ = v16
                        }
                        v9 += (v9 | (^(v9)))
                        _ = fi14
                    }
                    fmt.Println(fi8, fi8, p0)
                    _ = v13
                }
                for fi17 := 0; fi17 < 14; fi17++ {
                    p0 += (p0 | -32768)
                    fmt.Println(p0, fi17)
                    _ = fi17
                }
                v18 := GenS2{F0: ""}
                _ = v18
            }
        case 117:
            var v19 uint64 = 2
            _ = v19
        case 0:
            if true {
                v20 := GenS2{F0: ""}
                var v21 int8 = (v9 + (v9 % ((v9) | 1)))
                for fi22 := 0; fi22 < 4; fi22++ {
                    p1 += fi22
                    _ = fi22
                }
                _ = v20
                _ = v21
            } else {
                v23 := GenS1{F0: "hello"}
                _ = len(v7)
                _ = v23
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18
}
func genFn1(p0 uint32) int8 {
    switch -2147483648 {
    case 32768:
        if false {
            var v24 string = ""
            _ = v24
        } else {
            if (2 <= 1) {
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            fmt.Println(p0)
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if (true && false) {
            if (!(true || false)) {
                var v27 int32 = -1
                var v28 bool = (true || (false && false))
                _ = v27
                _ = v28
            } else {
                var v29 int16 = -379
                v29 = (^(v29))
                _ = v29
            }
            v30 := []int{fi26, fi26, fi26}
            if true {
                v31 := GenS1{F0: "hello"}
                _ = v31
            }
            _ = v30
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            var v33 int8 = -128
            switch (fi26 ^ fi26) {
            case 32767:
                var v34 uint = 435
                _ = v34
            case 741:
                if (252 < 646) {
                    var v35 string = "a"
                    switch int(p0) {
                    case 42:
                        p0 -= (^(p0))
                    case 1:
                        v36 := GenS1{F0: ("a" + "foo bar")}
                        switch v35 {
                        case "a":
                            var v37 uint8 = 96
                            if (!true) {
                                v35 = (v35 + v35)
                                var v38 uint8 = 42
                                _ = v38
                            }
                            _ = v37
                        case "":
                            v39 := []int{(fi32 & (fi26 - fi26)), 32767, (fi32 * fi32)}
                            _ = v39
                        case "foo bar":
                            v40 := []int{fi32}
                            _ = v40
                        }
                        _ = v36
                    default:
                        if (fi32 == fi32) {
                            var v41 int32 = int32(v33)
                            p0 = p0
                            var v42 uint32 = (p0 >> ((42) & 63))
                            _ = v41
                            _ = v42
                        }
                    }
                    var v43 float32 = float32(0.0)
                    _ = v35
                    _ = v43
                }
            }
            fmt.Println(fi26, fi32, fi32)
            _ = fi32
            _ = v33
        }
        p0 = (p0 & (p0 << ((774) & 63)))
        _ = fi26
    }
    var v44 int32 = 0
    var v45 string = ("a" + "a")
    _ = p0
    _ = v44
    _ = v45
    return -42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2147483647, 234}
    var v3 int64 = -1
    var v4 int16 = 42
    var v5 uint = 4294967295
    if ((1.5 > (3.14159 / 1.5)) || (3.14159 > 3.14159)) {
        switch "\n\t" {
        case "foo bar":
            if false {
                var v46 float32 = 1e10
                _ = v46
            }
        case "a":
            if (9 >= 65535) {
                if true {
                    if (1.5 != 0.1) {
                        v47 := GenS2{F0: ("foo bar" + "hello")}
                        var v48 int8 = genFn1(978)
                        if (int8(v5) > 35) {
                            v3 = v3
                        } else {
                            var v49 uint32 = 127
                            if true {
                                if false {
                                    var v50 int32 = 425
                                    _ = v50
                                } else {
                                    switch ("" + ("a" + ("hello" + (("hello" + ("hello" + "\n\t")) + (("" + "") + ((("hello" + "hello") + "\n\t") + "\n\t")))))) {
                                    case "hello":
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            v52 := GenS2{F0: ""}
                                            v49 = (v49 & v49)
                                            v48 = 127
                                            _ = fi51
                                            _ = v52
                                        }
                                    case "a":
                                        var v53 uint32 = (v49 >> ((2) & 63))
                                        _ = v53
                                    case "\n\t":
                                        v54 := v1[-9223372036854775808]
                                        _ = v54
                                    default:
                                        v55 := GenS2{F0: "a"}
                                        _ = v55
                                    }
                                    for fi56 := 0; fi56 < 18; fi56++ {
                                        var v57 float32 = (0.1 + (0.0 + -2.25))
                                        for fi58 := 0; fi58 < 20; fi58++ {
                                            var v59 int32 = -1
                                            _ = fi58
                                            _ = v59
                                        }
                                        _ = fi56
                                        _ = v57
                                    }
                                }
                                var v60 uint16 = 0
                                switch 127 {
                                case 128:
                                    for fi61 := 0; fi61 < 13; fi61++ {
                                        v4 += 42
                                        _ = fi61
                                    }
                                case -289:
                                    var v62 uint16 = v60
                                    _ = v62
                                case 32767:
                                    var v63 string = ""
                                    _ = v63
                                }
                                _ = v60
                            } else {
                                var v64 int = 4294967295
                                var v65 bool = (!false)
                                _ = v64
                                _ = v65
                            }
                            _ = v49
                        }
                        _ = v47
                        _ = v48
                    } else {
                        var v66 int8 = -84
                        fmt.Println(v66, v66)
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 9; fi67++ {
                        fmt.Println(fi67)
                        _ = fi67
                    }
                    var v68 uint32 = 466
                    _ = v68
                } else {
                    var v69 uint64 = 4294967295
                    _ = v69
                }
                v5 = (v5 >> (((v5 + v5)) & 63))
                for fi70 := 0; fi70 < 16; fi70++ {
                    if (((0.1 / 3.14159) + (((-2.25 / float32(3.14159)) * -2.25) + (((0.1 + 0.1) + float32(0.1)) - 1.5))) >= 1e10) {
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    } else {
                        fmt.Println(v3, v5)
                    }
                    _ = fi70
                }
            }
        case "":
            if false {
                v71 := GenS3{F0: 986, F1: "", F2: (^(v4))}
                _ = v71
            }
        default:
            var v72 uint64 = 4294967295
            _ = v72
        }
        v5 += v5
    } else {
        var v73 int16 = (v4 / ((v4) | 1))
        _ = v73
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v74 float64 = 0.0
    v4 -= int16(v3)
    if ((v74 / ((((v74 + -2.25) * (0.1 * (-2.25 / (v74 * (v74 - v74))))) * (v74 / v74)) + 1e10)) <= (v74 / 0.1)) {
        if (!false) {
            var v75 uint64 = 4294967295
            for fi76 := 0; fi76 < 5; fi76++ {
                v77 := GenS2{F0: ((("foo bar" + "\n\t") + "hello") + ("foo bar" + ""))}
                _ = fi76
                _ = v77
            }
            v5 = v5
            _ = v75
        } else {
            v74 = (((((v74 + v74) / v74) / v74) - v74) - (float64(0.1) / 1.5))
        }
    } else {
        fmt.Println(v4, v74)
        v78 := GenS2{F0: ""}
        _ = v78
    }
    v4 = (^(v4))
    if true {
        v3 = (v3 % (((^(v3))) | 1))
        for fi79 := 0; fi79 < 11; fi79++ {
            v5 = (v5 >> (((v5 << ((679) & 63))) & 63))
            fmt.Println(v74)
            _ = fi79
        }
        v80 := GenS3{F0: 2147483647, F1: ("hello" + ("foo bar" + "")), F2: v4}
        _ = v80
    }
    var v81 int64 = v3
    var v82 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v81
    _ = v82
}
