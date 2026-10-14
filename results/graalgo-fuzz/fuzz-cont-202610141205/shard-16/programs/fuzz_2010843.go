package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint32; F1 uint8; F2 string }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint64) uint8 {
    switch "foo bar" {
    case "foo bar":
        var v7 int8 = 1
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = v8
    return 165
}
func genFn1(p0 string, p1 uint32, p2 uint16) int32 {
    p0 = p0
    for fi9 := 0; fi9 < 16; fi9++ {
        p0 = (p0 + p0)
        var v10 uint8 = uint8(p2)
        p1 = 1
        _ = fi9
        _ = v10
    }
    var v11 string = p0
    v12 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := []uint32{4294967295, 0, 2147483647}
    var v2 int = 9223372036854775807
    var v3 float32 = (((1.5 - 3.14159) / (1e10 + (-2.25 - -2.25))) - 3.14159)
    v4 := make(map[string]uint32)
    var v5 int = (v2 >> ((65535) & 63))
    v6 := make(map[int]int)
    var v13 int = (v2 + v2)
    _ = len(v1)
    switch (v5 << ((18446744073709551615) & 63)) {
    case 0:
        fmt.Println(v5)
    case 659:
        fmt.Println(v2, v13, v2)
    case 127:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v2 = (v5 & v13)
            if (-32768 < -56) {
                v3 = v3
                if ((false && (!false)) || false) {
                    switch 555 {
                    case 128:
                        var v14 string = "hello"
                        _ = v14
                    default:
                        if (2 <= 97) {
                            var v15 int32 = 2147483647
                            _ = v15
                        }
                    }
                } else {
                    fmt.Println(v5, v13)
                }
            }
        }
    default:
        fmt.Println(v5, v5, v5)
    }
    switch (v5 / ((v2) | 1)) {
    case -752:
        var v16 uint64 = 42
        _ = v16
    }
    var v17 uint64 = 869
    if true {
        v18 := GenS1{F0: v2}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v5, v5, v2)
            switch "foo bar" {
            case "foo bar":
                v20 := GenS1{F0: (fi19 % ((v13) | 1))}
                _ = v20
            case "":
                v6[(^(v5))] = v5
            }
            _ = fi19
        }
        v5 = (v5 | v2)
    }
    var v21 int64 = 42
    if false {
        var v22 int32 = genFn1(("a" + "hello"), 1, 0)
        fmt.Println(v3)
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v21 = 881
        v1 = append(v1, 2)
        v17 -= v17
        _ = fi23
    }
    v3 = (v3 + ((3.14159 - v3) - (v3 * v3)))
    v5 = (v2 * (v2 / ((v13) | 1)))
    if (32767 <= 2) {
        var v24 int16 = 296
        var v25 uint16 = 289
        _ = v24
        _ = v25
    }
    switch (("" + "\n\t") + (("a" + "\n\t") + "foo bar")) {
    case "foo bar":
        v3 -= 0.1
    }
    if (genFn1(("hello" + "hello"), 42, 2) < -905) {
        if (1 < 2) {
            var v26 float64 = 1e10
            switch "" {
            case "hello":
                var v27 int64 = (-(v21))
                _ = v27
            }
            if ((32768 <= (v13 - 127)) && (true && (167 == 126))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v4, "a")
            } else {
                var v28 bool = false
                _ = v28
            }
            _ = v26
        } else {
            var v29 uint = 42
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        if false {
            delete(v4, "\n\t")
            var v31 uint32 = 231
            switch ("foo bar" + "a") {
            case "hello":
                if (("foo bar" + "a") == ("foo bar" + ("a" + "a"))) {
                    v32 := GenS2{F0: 1, F1: (v21 + v21)}
                    v33 := GenS3{F0: 123, F1: 9223372036854775807}
                    _ = v32
                    _ = v33
                } else {
                    var v34 string = ""
                    v2 = 42
                    _ = v34
                }
            case "foo bar":
                fmt.Println(v2, v5, v2)
            case "\n\t":
                var v35 string = "\n\t"
                _ = v35
            default:
                if false {
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v5 = fi36
                        var v37 uint64 = uint64(v31)
                        _ = fi36
                        _ = v37
                    }
                }
            }
            _ = v31
        } else {
            switch "foo bar" {
            case "":
                _ = len(v1)
            case "a":
                if false {
                    if true {
                        if (!false) {
                            var v38 string = (("a" + (("a" + "a") + ("a" + ""))) + "")
                            _ = v38
                        } else {
                            if (v21 == (v21 | (v21 * v21))) {
                                fmt.Println(v5, v5, fi30)
                            } else {
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    fmt.Println(v5, fi39, fi39)
                                    var v40 uint16 = 237
                                    var v41 uint16 = (v40 % ((2) | 1))
                                    switch int(v40) {
                                    case 42:
                                        var v42 uint16 = v41
                                        _ = v42
                                    default:
                                        v17 = 4294967295
                                    }
                                    fmt.Println(v21)
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                            }
                            fmt.Println(v17)
                            v5 = v2
                        }
                        var v43 bool = false
                        _ = v43
                    } else {
                        switch (v5 % ((42) | 1)) {
                        case 127:
                            var v44 int16 = 0
                            _ = v44
                        }
                        if (!true) {
                            v5 = -86
                        } else {
                            if false {
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    var v46 uint8 = 0
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v47 int16 = 0
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                                if true {
                                    for fi48 := 0; fi48 < 10; fi48++ {
                                        fmt.Println(v3, v21, v3)
                                        v13 -= 0
                                        v3 = v3
                                        _ = fi48
                                    }
                                    v17 -= 2
                                }
                                var v49 int = v2
                                _ = v49
                            } else {
                                v50 := GenS1{F0: fi30}
                                if false {
                                    v2 = (fi30 | v2)
                                    if true {
                                        for fi51 := 0; fi51 < 16; fi51++ {
                                            v52 := GenS1{F0: v13}
                                            delete(v4, "a")
                                            _ = fi51
                                            _ = v52
                                        }
                                        switch ("foo bar" + "hello") {
                                        case "hello":
                                            v21 = v21
                                        case "":
                                            fmt.Println(v13, v13)
                                        case "\n\t":
                                            var v53 uint16 = 42
                                            _ = v53
                                        }
                                    }
                                }
                                var v54 uint64 = 1
                                _ = v50
                                _ = v54
                            }
                        }
                    }
                    delete(v6, (fi30 * 127))
                    fmt.Println(v21)
                } else {
                    var v55 int = v2
                    for fi56 := 0; fi56 < 19; fi56++ {
                        var v57 uint16 = 53
                        v58 := GenS4{F0: 42}
                        v57 -= (^(v57))
                        var v59 int = 32768
                        if true {
                            v60 := GenS5{F0: (^(v57)), F1: 0}
                            _ = len(v1)
                            fmt.Println(v5)
                            if false {
                                if true {
                                    v57 = v57
                                    var v61 int64 = 944
                                    _ = v61
                                } else {
                                    if true {
                                        v62 := GenS6{F0: 2147483647, F1: genFn0(191), F2: "foo bar"}
                                        var v63 int64 = (v21 >> ((255) & 63))
                                        switch (fi30 & int(v17)) {
                                        case 1:
                                            v64 := GenS2{F0: 4294967295, F1: (v21 >> ((871) & 63))}
                                            _ = v64
                                        case -9223372036854775808:
                                            var v65 int = 2147483647
                                            _ = v65
                                        default:
                                            var v66 uint16 = (v57 + v57)
                                            _ = v66
                                        }
                                        _ = v62
                                        _ = v63
                                    }
                                    v67 := GenS6{F0: 4294967295, F1: 123, F2: (("hello" + ("hello" + ("foo bar" + ("hello" + ("hello" + "hello"))))) + "a")}
                                    v5 = -2147483648
                                    _ = v67
                                }
                                v57 = v57
                                fmt.Println(v5, v55)
                            } else {
                                if true {
                                    fmt.Println(v57, v5)
                                    var v68 int64 = (v21 + v21)
                                    v55 += (v59 | v55)
                                    _ = v68
                                } else {
                                    fmt.Println(v57, v17)
                                }
                                if (false || false) {
                                    var v69 int32 = -849
                                    v70 := GenS5{F0: v57, F1: 1}
                                    _ = v69
                                    _ = v70
                                }
                                var v71 uint16 = (^(v57))
                                v72 := GenS7{F0: genFn0(v17), F1: 211}
                                _ = v71
                                _ = v72
                            }
                            v73 := GenS5{F0: v57, F1: 0}
                            _ = v60
                            _ = v73
                        }
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    _ = v55
                }
            default:
                v3 = 1e10
            }
            for fi74 := 0; fi74 < 13; fi74++ {
                v75 := GenS8{F0: 127}
                _ = fi74
                _ = v75
            }
        }
        _ = fi30
    }
    var v76 int8 = -63
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
