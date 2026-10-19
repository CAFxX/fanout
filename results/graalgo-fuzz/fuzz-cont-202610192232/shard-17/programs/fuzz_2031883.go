package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int8; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 87
    v3 := []string{"", (("a" + "") + ("hello" + "\n\t")), "foo bar", ("a" + "a")}
    var v4 uint32 = 1
    var v5 float32 = (3.14159 / (v1 + 1e10))
    var v6 int16 = 32767
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 bool = false
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8, v5)
        _ = fi8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if (!(0 > v6)) {
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint64 = 531
                v11 -= 0
                v6 = v6
                _ = fi10
                _ = v11
            }
            var v12 string = "\n\t"
            for fi13 := 0; fi13 < 5; fi13++ {
                if len(v3) > 0 {
                    v3[0] = v12
                }
                fmt.Println(v6)
                var v14 int = fi9
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        if (int(v6) > -893) {
            if ((v2 / (((^(v2))) | 1)) >= v2) {
                v15 := GenS1{F0: "\n\t", F1: (v2 * v2)}
                _ = v15
            } else {
                var v16 bool = true
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint64 = 428
                _ = fi17
                _ = v18
            }
        }
        if v7 {
            var v19 int = (-(fi9))
            var v20 int = (^(fi9))
            if v7 {
                var v21 uint16 = 1
                var v22 string = "a"
                if (65535 != 65535) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v2 -= (v2 + (v2 + 18446744073709551615))
                        var v24 uint64 = (v2 ^ 0)
                        v25 := GenS2{F0: 42, F1: 35}
                        if (127 == 42) {
                            v26 := GenS3{F0: fi9, F1: (^(v21)), F2: v20}
                            if v7 {
                                v5 -= (float32(float64(1e10)) * float32(3.14159))
                                var v27 uint8 = 127
                                _ = v27
                            }
                            if len(v3) > 0 {
                                v3[0] = v22
                            }
                            _ = v26
                        } else {
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = v22
                                }
                                fmt.Println(v19, v4)
                            } else {
                                var v28 int8 = -120
                                _ = v28
                            }
                            _ = len(v3)
                            var v29 uint8 = 255
                            _ = v29
                        }
                        switch v22 {
                        case "a":
                            var v30 string = (("a" + ("a" + v22)) + "foo bar")
                            _ = v30
                        case "\n\t":
                            v31 := GenS4{F0: v22}
                            _ = v31
                        case "foo bar":
                            switch 0 {
                            case 127:
                                fmt.Println(v5)
                            case 128:
                                v32 := GenS4{F0: "hello"}
                                _ = v32
                            case 15:
                                v1 = v1
                            }
                        default:
                            fmt.Println(v20, v4)
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                } else {
                    var v33 float32 = v1
                    _ = v33
                }
                _ = v21
                _ = v22
            } else {
                v7 = true
                if (0 > 0) {
                    if v7 {
                        switch "foo bar" {
                        case "\n\t":
                            var v34 int16 = 0
                            _ = v34
                        case "":
                            v35 := GenS5{F0: 481, F1: 42, F2: 97}
                            _ = v35
                        }
                        var v36 int8 = 100
                        if v7 {
                            if false {
                                switch "" {
                                case "a":
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        v6 -= (v6 | 120)
                                        v38 := GenS6{F0: 1}
                                        _ = fi37
                                        _ = v38
                                    }
                                case "\n\t":
                                    for fi39 := 0; fi39 < 14; fi39++ {
                                        var v40 uint64 = v2
                                        switch (v20 / (((fi39 | fi9)) | 1)) {
                                        case 548:
                                            var v41 int16 = (v6 << ((uint(v4)) & 63))
                                            _ = v41
                                        case -1:
                                            if v7 {
                                                v19 += 753
                                                switch (v20 + 287) {
                                                case 9223372036854775807:
                                                    if v7 {
                                                        var v42 float32 = -2.25
                                                        switch "hello" {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "a":
                                                                fmt.Println(v5, fi39)
                                                            default:
                                                                switch (fi39 - fi39) {
                                                                case -166:
                                                                    if (v7 && (127 <= (^(v36)))) {
                                                                        var v43 string = ("a" + "hello")
                                                                        if v7 {
                                                                            v44 := GenS6{F0: 0}
                                                                            v45 := GenS6{F0: 42}
                                                                            if (!v7) {
                                                                                fmt.Println(v1)
                                                                            } else {
                                                                                var v46 uint64 = 42
                                                                                _ = v46
                                                                            }
                                                                            _ = v44
                                                                            _ = v45
                                                                        }
                                                                        _ = v43
                                                                    }
                                                                case 1:
                                                                    if true {
                                                                        fmt.Println(v7, v1, v42)
                                                                        fmt.Println(v6)
                                                                    } else {
                                                                        var v47 uint16 = 1
                                                                        v5 += (float32(3.14159) + 0.0)
                                                                        _ = v47
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            var v48 bool = v7
                                                            var v49 int32 = -13
                                                            var v50 bool = v48
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        _ = v42
                                                    } else {
                                                        switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                        case "\n\t":
                                                            fmt.Println(v2)
                                                        case "foo bar":
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                        case "":
                                                            switch (fi39 / ((fi9) | 1)) {
                                                            case 125:
                                                                fmt.Println(v20, v20)
                                                            case -363:
                                                                v4 -= v4
                                                            case 2147483647:
                                                                var v51 int64 = 42
                                                                switch v19 {
                                                                case 128:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v52 uint64 = 62
                                                                        _ = v52
                                                                    case "\n\t":
                                                                        v53 := GenS2{F0: 42, F1: (v36 | -121)}
                                                                        _ = v53
                                                                    }
                                                                default:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        var v54 float64 = (float64(-2.25) - 0.0)
                                                                        _ = v54
                                                                    }
                                                                }
                                                                v55 := GenS4{F0: "hello"}
                                                                _ = v51
                                                                _ = v55
                                                            }
                                                        default:
                                                            var v56 float64 = 0.1
                                                            _ = v56
                                                        }
                                                        fmt.Println(v4)
                                                    }
                                                case 1:
                                                    var v57 bool = ((v1 != v5) || (0.0 == 0.0))
                                                    _ = v57
                                                default:
                                                    fmt.Println(v1, fi39, v1)
                                                }
                                            } else {
                                                v7 = true
                                                v58 := GenS3{F0: v19, F1: 704, F2: (v20 | -2147483648)}
                                                _ = v58
                                            }
                                            fmt.Println(v19, v1, v20)
                                            var v59 uint8 = 42
                                            var v60 float64 = 1e10
                                            _ = v59
                                            _ = v60
                                        }
                                        switch ("" + "a") {
                                        case "":
                                            v61 := GenS6{F0: 677}
                                            switch (fi9 & v19) {
                                            case -9223372036854775808:
                                                v4 = 2147483647
                                            }
                                            v6 = 0
                                            if (!(!v7)) {
                                                v3 = append(v3, "a")
                                                v40 = v40
                                            }
                                            _ = v61
                                        case "foo bar":
                                            if v7 {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        case "hello":
                                            var v62 uint = 4294967295
                                            _ = v62
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                v5 = (v1 + 3.14159)
                                fmt.Println(v36, v7)
                            }
                            v63 := GenS7{F0: (v2 & 4294967295)}
                            var v64 int32 = 0
                            _ = v63
                            _ = v64
                        } else {
                            switch ("" + (("" + "") + "hello")) {
                            case "\n\t":
                                if (2147483647 >= 932) {
                                    v1 += v1
                                    if (!v7) {
                                        fmt.Println(v7)
                                        v36 = (v36 / (((v36 % (((v36 - (v36 & (v36 - (-(v36)))))) | 1))) | 1))
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = ("" + ("a" + ""))
                                        }
                                        switch int(v6) {
                                        case 32767:
                                            v65 := GenS5{F0: -406, F1: v36, F2: 1}
                                            _ = v65
                                        }
                                    }
                                }
                            default:
                                var v66 bool = (1e10 != (1e10 + 1e10))
                                _ = v66
                            }
                        }
                        _ = v36
                    } else {
                        if (!v7) {
                            v20 = 42
                            fmt.Println(v19, v7)
                        } else {
                            v7 = (179 < 0)
                        }
                    }
                    fmt.Println(v6)
                }
                v6 += v6
            }
            _ = v19
            _ = v20
        }
        _ = fi9
    }
    v3 = append(v3, ("\n\t" + "hello"))
    for fi67 := 0; fi67 < 7; fi67++ {
        v6 += (v6 ^ (v6 << ((0) & 63)))
        _ = fi67
    }
    var v68 string = ""
    for fi69 := 0; fi69 < 16; fi69++ {
        v68 = ("\n\t" + v68)
        if (100 <= 42) {
            if (!v7) {
                v70 := GenS4{F0: ((v68 + v68) + v68)}
                var v71 bool = v7
                _ = v70
                _ = v71
            } else {
                switch "hello" {
                case "foo bar":
                    v72 := GenS2{F0: 0, F1: 0}
                    _ = v72
                case "hello":
                    v4 -= v4
                case "a":
                    v73 := GenS8{F0: 303}
                    _ = v73
                }
            }
            var v74 string = "hello"
            for fi75 := 0; fi75 < 7; fi75++ {
                v76 := GenS7{F0: (v2 | (v2 >> ((42) & 63)))}
                switch (fi75 / ((127) | 1)) {
                case -651:
                    var v77 string = "foo bar"
                    _ = v77
                }
                _ = fi75
                _ = v76
            }
            _ = v74
        } else {
            v4 = v4
        }
        for fi78 := 0; fi78 < 20; fi78++ {
            var v79 uint8 = 2
            _ = fi78
            _ = v79
        }
        _ = fi69
    }
    if true {
        var v80 bool = v7
        _ = v80
    } else {
        var v81 int16 = (v6 - (v6 >> ((uint(v2)) & 63)))
        for fi82 := 0; fi82 < 14; fi82++ {
            v5 = (v1 + v1)
            var v83 int = 2147483647
            _ = fi82
            _ = v83
        }
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
}
