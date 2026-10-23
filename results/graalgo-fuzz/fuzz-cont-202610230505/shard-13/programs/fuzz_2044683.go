package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -435
    v3 := make(map[int]int)
    var v4 float64 = 1e10
    if (true || true) {
        delete(v3, 0)
    } else {
        v5 := []int{0, 4294967295}
        if (9223372036854775807 >= -9223372036854775808) {
            fmt.Println(v2, v4)
            v6 := GenS1{F0: ("\n\t" + "hello")}
            var v7 float32 = -2.25
            _ = v6
            _ = v7
        } else {
            switch "\n\t" {
            case "foo bar":
                v8 := GenS1{F0: ((("foo bar" + "hello") + ("hello" + "a")) + (("a" + "") + ""))}
                _ = v8
            case "\n\t":
                fmt.Println(v2, v2)
            default:
                v9 := GenS2{F0: 156}
                _ = v9
            }
            fmt.Println(v2, v4)
        }
        var v10 int = 4294967295
        _ = v10
        _ = v5
    }
    switch -9223372036854775808 {
    case -870:
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int = 4294967295
            if (!true) {
                var v13 float64 = 0.1
                _ = v13
            }
            _ = fi11
            _ = v12
        }
    }
    var v14 int = -2147483648
    if true {
        if ("hello" == (("foo bar" + "foo bar") + (("\n\t" + "") + ("" + "foo bar")))) {
            for fi15 := 0; fi15 < 3; fi15++ {
                if (!(105 >= 65535)) {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 uint = 42
                        v18 := GenS2{F0: 42}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint32 = 97
                    switch ("foo bar" + "") {
                    case "":
                        v14 = 32768
                    }
                    _ = v19
                }
                _ = fi15
            }
        } else {
            if true {
                var v20 int64 = 0
                _ = v20
            } else {
                if false {
                    if false {
                        var v21 int = v14
                        v22 := GenS2{F0: 128}
                        var v23 int16 = v2
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(v4)
                    }
                }
                var v24 int64 = -1
                var v25 uint = 255
                _ = v24
                _ = v25
            }
        }
    } else {
        var v26 uint64 = 0
        switch 128 {
        case -1:
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(fi27)
                v3[(^(fi27))] = 2147483647
                _ = fi27
            }
        case 128:
            for fi28 := 0; fi28 < 5; fi28++ {
                if true {
                    var v29 bool = true
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v31 := GenS3{F0: -128, F1: -71, F2: 42}
                        _ = fi30
                        _ = v31
                    }
                    v26 += v26
                    _ = v29
                }
                if false {
                    v32 := GenS4{F0: 317, F1: fi28}
                    switch 0 {
                    case -36:
                        fmt.Println(v2)
                    default:
                        v2 = v2
                    }
                    delete(v3, (v14 | 9223372036854775807))
                    _ = v32
                } else {
                    if (659 != v26) {
                        v33 := GenS4{F0: 0, F1: v14}
                        _ = v33
                    }
                    v4 += 0.0
                }
                v2 -= 1000
                _ = fi28
            }
        case 127:
            v34 := GenS3{F0: 1, F1: 1, F2: 266}
            _ = v34
        default:
            for fi35 := 0; fi35 < 17; fi35++ {
                v26 = (v26 % ((v26) | 1))
                var v36 float32 = 0.1
                _ = fi35
                _ = v36
            }
        }
        v4 = (v4 - float64(0.0))
        _ = v26
    }
    v4 = ((((((v4 + v4) - (v4 - 0.1)) + float64(0.0)) - float64(0.0)) - ((1e10 / (v4 + float64(3.14159))) + v4)) + v4)
    v4 += v4
    var v37 int64 = -9223372036854775808
    for fi38 := 0; fi38 < 16; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            switch ("a" + ("\n\t" + "foo bar")) {
            case "hello":
                var v40 int = v14
                _ = v40
            case "foo bar":
                fmt.Println(fi38, v37, fi39)
            case "\n\t":
                var v41 int64 = 686
                _ = v41
            }
            v42 := GenS2{F0: 0}
            _ = fi39
            _ = v42
        }
        _ = fi38
    }
    v43 := GenS3{F0: 127, F1: -60, F2: -842}
    switch "" {
    case "":
        for fi44 := 0; fi44 < 12; fi44++ {
            v2 += 42
            var v45 uint32 = 2147483647
            _ = fi44
            _ = v45
        }
    }
    switch ("a" + "a") {
    case "a":
        for fi46 := 0; fi46 < 12; fi46++ {
            fmt.Println(fi46)
            _ = fi46
        }
    case "hello":
        var v47 uint32 = 65535
        _ = v47
    case "":
        v37 = (v37 - v37)
    }
    v48 := []int{}
    var v49 uint64 = 0
    if false {
        var v50 uint64 = v49
        v14 = (-(v14))
        _ = v50
    } else {
        var v51 int32 = 179
        v52 := GenS2{F0: 97}
        _ = len(v1)
        _ = v51
        _ = v52
    }
    _ = len(v48)
    var v53 float32 = -2.25
    switch (-(v14)) {
    case 32767:
        if false {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v53, v53)
            default:
                var v54 string = ("hello" + "")
                _ = v54
            }
            if false {
                if true {
                    fmt.Println(v2, v2)
                }
                delete(v3, 128)
                if false {
                    if false {
                        v37 += -1
                        v55 := GenS5{F0: (("a" + "\n\t") + "\n\t"), F1: 120}
                        _ = v55
                    }
                    var v56 uint8 = 10
                    var v57 float64 = (0.1 - -2.25)
                    _ = v56
                    _ = v57
                } else {
                    if (v49 > (v49 + (^(v49)))) {
                        v58 := v1["a"]
                        _ = v58
                    } else {
                        v2 = 32767
                        var v59 uint32 = 2147483647
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 2; fi60++ {
                        v61 := GenS6{F0: 255, F1: 42, F2: (!(0 <= 0))}
                        v3[v14] = (fi60 / ((1) | 1))
                        _ = fi60
                        _ = v61
                    }
                    fmt.Println(v49, v53)
                }
            }
            switch v14 {
            case 32767:
                v4 -= v4
            default:
                if true {
                    fmt.Println(v53)
                    switch "\n\t" {
                    case "a":
                        v37 = v37
                    case "foo bar":
                        var v62 int = v14
                        _ = v62
                    case "\n\t":
                        for fi63 := 0; fi63 < 17; fi63++ {
                            switch (-(fi63)) {
                            case 4294967295:
                                v64 := GenS1{F0: (("foo bar" + "foo bar") + ("foo bar" + ""))}
                                _ = v64
                            case 2147483647:
                                for fi65 := 0; fi65 < 4; fi65++ {
                                    var v66 int32 = 2147483647
                                    _ = fi65
                                    _ = v66
                                }
                            case 127:
                                for fi67 := 0; fi67 < 1; fi67++ {
                                    v68 := v3[(v14 - int(v49))]
                                    v68 += (fi63 ^ v68)
                                    fmt.Println(v53, v53)
                                    v48 = append(v48, (fi63 >> ((0) & 63)))
                                    _ = fi67
                                    _ = v68
                                }
                            default:
                                if (false && ((float64((v53 - (v53 * float32((float64(v53) - (v4 + v4)))))) < 1.5) && ((1000 >= (v2 / ((1) | 1))) || (!true)))) {
                                    v14 = (v14 << ((2) & 63))
                                    if (true || true) {
                                        v49 = (v49 - v49)
                                    }
                                    for fi69 := 0; fi69 < 8; fi69++ {
                                        v49 += v49
                                        _ = fi69
                                    }
                                }
                            }
                            v70 := GenS2{F0: 2}
                            var v71 uint = 42
                            _ = fi63
                            _ = v70
                            _ = v71
                        }
                    }
                } else {
                    v48 = append(v48, -854)
                    v4 = v4
                }
            }
        }
    case 9223372036854775807:
        switch "" {
        case "\n\t":
            var v72 uint8 = 2
            _ = v72
        case "":
            v73 := GenS7{F0: 1}
            _ = v73
        case "hello":
            switch -102 {
            case 891:
                _ = len(v3)
            case 4294967295:
                var v74 bool = (127 < -128)
                _ = v74
            default:
                v37 = v37
            }
        default:
            var v75 uint16 = 446
            _ = v75
        }
    default:
        if false {
            v2 = (v2 << ((18446744073709551615) & 63))
            var v76 uint = 255
            var v77 uint8 = 42
            _ = v76
            _ = v77
        } else {
            v78 := GenS4{F0: 1, F1: (v14 ^ (v14 % (((v14 | v14)) | 1)))}
            var v79 float32 = float32((v4 * 1.5))
            v80 := GenS4{F0: 42, F1: (^(v14))}
            _ = v78
            _ = v79
            _ = v80
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v48
    _ = v49
    _ = v53
}
