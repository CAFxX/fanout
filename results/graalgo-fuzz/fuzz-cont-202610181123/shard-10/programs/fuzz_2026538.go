package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 string }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 string = ("a" + "hello")
    var v5 uint = 133
    var v6 int8 = -128
    _ = len(v3)
    v7 := GenS1{F0: 815, F1: v4}
    if (false && true) {
        v8 := GenS2{F0: 42, F1: 100, F2: ((v4 + v4) + ((v4 + ("foo bar" + "\n\t")) + v4))}
        v5 = (^(v5))
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi9, v5, fi10)
                if true {
                    v4 = "\n\t"
                    fmt.Println(fi9, v1)
                }
                if true {
                    var v11 uint = (v5 & v5)
                    var v12 uint32 = 825
                    v13 := GenS1{F0: 42, F1: ("foo bar" + v4)}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            var v14 uint = v5
            var v15 uint = (v14 & 201)
            _ = fi9
            _ = v14
            _ = v15
        }
        _ = v8
    } else {
        for fi16 := 0; fi16 < 11; fi16++ {
            if (true || true) {
                if (((("foo bar" + "\n\t") + (v4 + ("a" + ""))) + v4) != "") {
                    v2["foo bar"] = (v4 + v4)
                    var v17 float64 = (float64(((v1 * 0.1) - v1)) * -2.25)
                    var v18 uint64 = 907
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint64 = 18446744073709551615
                    _ = v19
                }
            } else {
                v4 = "a"
            }
            var v20 float64 = float64((1e10 * float32(0.0)))
            v20 -= v20
            _ = fi16
            _ = v20
        }
        var v21 uint8 = 132
        var v22 float64 = float64(1.5)
        _ = v21
        _ = v22
    }
    var v23 uint64 = 2
    var v24 int32 = -220
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS3{F0: 1, F1: v24}
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
        _ = fi25
        _ = v26
    }
    if (!(65535 >= 2147483647)) {
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(v4, v1, fi27)
            _ = fi27
        }
        v6 = (v6 & (v6 << ((v5) & 63)))
        for fi28 := 0; fi28 < 4; fi28++ {
            switch fi28 {
            case 392:
                v1 -= v1
            case 465:
                v5 += (v5 % (((v5 * (^(v5)))) | 1))
            case -9223372036854775808:
                var v29 bool = (uint32(v5) <= 2)
                _ = v29
            }
            switch v4 {
            case "\n\t":
                switch v4 {
                case "a":
                    v23 += (^(v23))
                }
            }
            _ = fi28
        }
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        if true {
            for fi31 := 0; fi31 < 19; fi31++ {
                if true {
                    if (-720 != (fi30 - (fi31 | fi31))) {
                        var v32 bool = true
                        _ = v32
                    }
                    v5 -= (v5 | v5)
                } else {
                    switch (fi30 * (fi30 | fi30)) {
                    case 4294967295:
                        _ = len(v3)
                    case 32768:
                        if ((!true) && false) {
                            delete(v2, "hello")
                        }
                    }
                    var v33 uint16 = 255
                    _ = v33
                }
                _ = fi31
            }
            var v34 int32 = v24
            var v35 int8 = v6
            _ = v34
            _ = v35
        }
        if ((^(v23)) < 42) {
            v36 := v2[v4]
            _ = v36
        }
        v3 = append(v3, -899)
        _ = fi30
    }
    v1 = (float32((1.5 + 3.14159)) - ((v1 - v1) + float32(-2.25)))
    v5 = (^(v5))
    if true {
        if (!((v23 - v23) > v23)) {
            var v37 float64 = 1.5
            _ = v37
        }
        var v38 uint8 = 128
        _ = v38
    } else {
        v39 := GenS4{F0: 160, F1: (v4 + v4)}
        var v40 int16 = 1
        for fi41 := 0; fi41 < 2; fi41++ {
            v4 = (v4 + "")
            if (!(false || ((!false) || false))) {
                var v42 uint8 = 128
                _ = v42
            }
            _ = fi41
        }
        _ = v39
        _ = v40
    }
    v43 := GenS5{F0: -9223372036854775808}
    if false {
        v24 = (v24 % (((v24 << (((v5 * (^(v5)))) & 63))) | 1))
        for fi44 := 0; fi44 < 4; fi44++ {
            if true {
                var v45 uint64 = (v23 | v23)
                _ = v45
            } else {
                if false {
                    var v46 uint = v5
                    var v47 uint64 = (v23 + 0)
                    _ = v46
                    _ = v47
                }
                if (false && true) {
                    v48 := GenS1{F0: 0, F1: ""}
                    fmt.Println(v5, v6)
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v5 += (v5 * v5)
                        var v50 uint8 = uint8(v23)
                        delete(v2, v4)
                        if (fi44 >= 128) {
                            if false {
                                v1 = v1
                            } else {
                                var v51 uint8 = 250
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v52 int32 = v24
                                    _ = v52
                                } else {
                                    if false {
                                        var v53 uint32 = 816
                                        _ = len(v3)
                                        _ = v53
                                    }
                                }
                                _ = v51
                            }
                        } else {
                            var v54 uint8 = v50
                            _ = v54
                        }
                        _ = fi49
                        _ = v50
                    }
                    v23 = v23
                    v55 := GenS1{F0: 323, F1: ("a" + (("a" + v4) + v4))}
                    _ = v55
                }
                for fi56 := 0; fi56 < 6; fi56++ {
                    v6 += v6
                    v2[v4] = v4
                    var v57 float32 = v1
                    _ = fi56
                    _ = v57
                }
            }
            for fi58 := 0; fi58 < 11; fi58++ {
                fmt.Println(v5, v1)
                v59 := GenS6{F0: uint32(v23), F1: ((v4 + ("" + v4)) + "foo bar")}
                _ = fi58
                _ = v59
            }
            v4 = ""
            _ = fi44
        }
        for fi60 := 0; fi60 < 2; fi60++ {
            if true {
                v6 = (v6 % ((int8(v24)) | 1))
            } else {
                if (0.1 != 0.0) {
                    if true {
                        delete(v2, "\n\t")
                    }
                    v61 := GenS4{F0: v5, F1: "\n\t"}
                    v1 -= (v1 + (-2.25 - v1))
                    _ = v61
                } else {
                    v62 := GenS4{F0: (v5 | v5), F1: v4}
                    var v63 uint8 = 7
                    v23 -= 766
                    _ = v62
                    _ = v63
                }
                var v64 float64 = 1e10
                _ = v64
            }
            v65 := GenS5{F0: (fi60 << ((v5) & 63))}
            for fi66 := 0; fi66 < 8; fi66++ {
                if true {
                    v67 := GenS2{F0: 0, F1: 2, F2: ((v4 + (v4 + v4)) + "\n\t")}
                    if true {
                        v4 = ("" + ("" + v4))
                        var v68 string = v4
                        v68 = v68
                        _ = v68
                    }
                    v3 = append(v3, (fi66 ^ (fi60 & (fi60 & (^(fi60))))))
                    _ = v67
                }
                fmt.Println(v4, fi66)
                _ = len(v3)
                _ = fi66
            }
            _ = fi60
            _ = v65
        }
    } else {
        for fi69 := 0; fi69 < 15; fi69++ {
            if (!(496 < 0)) {
                v1 += 1e10
                var v70 uint = v5
                v5 = (v5 / (((v70 ^ v70)) | 1))
                _ = v70
            } else {
                v24 = (v24 >> ((2) & 63))
                v71 := GenS5{F0: (fi69 >> ((v5) & 63))}
                v5 -= uint(v23)
                _ = v71
            }
            _ = fi69
        }
        var v72 int = 4294967295
        for fi73 := 0; fi73 < 7; fi73++ {
            var v74 int32 = v24
            switch (v4 + (v4 + v4)) {
            case "foo bar":
                switch 4294967295 {
                case 1:
                    fmt.Println(v72)
                case 2147483647:
                    for fi75 := 0; fi75 < 16; fi75++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi75
                    }
                case 888:
                    v3 = append(v3, (^(fi73)))
                }
            }
            fmt.Println(v6, v23)
            _ = fi73
            _ = v74
        }
        _ = v72
    }
    if (false && false) {
        var v76 int = 42
        _ = v76
    } else {
        var v77 uint32 = 460
        fmt.Println(v1)
        v1 -= v1
        _ = v77
    }
    var v78 string = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v78
}
