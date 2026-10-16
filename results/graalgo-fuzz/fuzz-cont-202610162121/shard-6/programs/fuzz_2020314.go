package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 string }
type GenS5 struct { F0 uint64; F1 int32; F2 int }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 int32; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint64; F1 string }

func main() {
    v1 := []uint32{0, 42, 2, 401}
    v2 := make(map[string]string)
    var v3 int = 2147483647
    var v4 int32 = 2147483647
    var v5 float32 = 0.0
    var v6 uint = 461
    var v7 uint8 = 158
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float32 = 0.1
        v7 += (v7 * v7)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 18446744073709551615, F1: 0}
    fmt.Println(v5, v7, v6)
    if false {
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                if true {
                    if false {
                        v7 += (v7 % (((^(v7))) | 1))
                        var v12 float32 = v5
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            switch ("hello" + ("" + "foo bar")) {
                            case "a":
                                v7 = (v7 * uint8(v6))
                            case "foo bar":
                                switch (("" + ((("\n\t" + "\n\t") + "\n\t") + ("a" + "\n\t"))) + "\n\t") {
                                case "hello":
                                    v7 = v7
                                }
                            }
                            _ = fi13
                        }
                        var v14 uint32 = 65535
                        _ = v14
                    }
                }
                _ = fi11
            }
            v5 = ((-2.25 - v5) * (v5 - v5))
            fmt.Println(v5, v5)
        }
        if true {
            switch -2147483648 {
            case 876:
                if true {
                    var v15 int32 = (v4 << ((v6) & 63))
                    fmt.Println(v15, v3, v7)
                    v6 += 4294967295
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS1{F0: 688, F1: -128}
                        var v18 string = "a"
                        switch 4294967295 {
                        case -809:
                            fmt.Println(fi16)
                        default:
                            delete(v2, v18)
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            default:
                var v19 uint16 = 32767
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS1{F0: 939, F1: 42}
                v22 := GenS1{F0: 715, F1: -128}
                _ = fi20
                _ = v21
                _ = v22
            }
            if (493 == 32767) {
                if (-1 < 42) {
                    v23 := GenS1{F0: 4294967295, F1: -1}
                    v2["a"] = "\n\t"
                    var v24 string = "hello"
                    _ = v23
                    _ = v24
                } else {
                    switch v3 {
                    case -65:
                        v25 := GenS2{F0: v6, F1: v7, F2: 127}
                        _ = v25
                    }
                    if (374 == 42) {
                        var v26 float64 = 3.14159
                        var v27 float32 = v5
                        _ = v26
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                fmt.Println(v6, v4, fi29)
                                _ = fi29
                            }
                            _ = fi28
                        }
                        var v30 bool = false
                        for fi31 := 0; fi31 < 19; fi31++ {
                            if v30 {
                                v32 := GenS3{F0: 42}
                                _ = v32
                            }
                            if (4294967295 < 2147483648) {
                                fmt.Println(v3, v4)
                                var v33 int64 = 262
                                _ = v33
                            }
                            _ = fi31
                        }
                        _ = v30
                    }
                }
                v34 := GenS4{F0: 696, F1: 103, F2: "hello"}
                var v35 float32 = 1.5
                _ = v34
                _ = v35
            } else {
                var v36 uint8 = 242
                var v37 bool = true
                for fi38 := 0; fi38 < 7; fi38++ {
                    if v37 {
                        switch ("" + "") {
                        case "":
                            for fi39 := 0; fi39 < 8; fi39++ {
                                v40 := GenS5{F0: 856, F1: v4, F2: (fi39 + fi39)}
                                _ = fi39
                                _ = v40
                            }
                        }
                        _ = len(v1)
                    }
                    for fi41 := 0; fi41 < 16; fi41++ {
                        var v42 int32 = 657
                        v4 += -909
                        _ = fi41
                        _ = v42
                    }
                    v43 := GenS5{F0: 42, F1: v4, F2: (-(fi38))}
                    _ = fi38
                    _ = v43
                }
                _ = v36
                _ = v37
            }
        }
        for fi44 := 0; fi44 < 1; fi44++ {
            v7 -= 45
            v4 = (v4 + v4)
            var v45 uint16 = 520
            _ = fi44
            _ = v45
        }
    }
    if false {
        switch v3 {
        case -1:
            v5 = (((v5 / v5) + v5) / 1.0)
        }
    }
    if (true && false) {
        switch (v3 & (v3 - v3)) {
        case 128:
            v7 = v7
        case 9223372036854775807:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if (v3 < (v3 - v3)) {
                delete(v2, "foo bar")
                _ = len(v1)
                fmt.Println(v5, v7, v3)
            }
        }
        var v46 bool = (!true)
        _ = v46
    } else {
        var v47 float64 = 3.14159
        v48 := GenS6{F0: (v6 ^ 42), F1: int(v6)}
        _ = v47
        _ = v48
    }
    v49 := GenS7{F0: v7, F1: 4294967295}
    v2["\n\t"] = ("\n\t" + ("" + "a"))
    switch "\n\t" {
    case "foo bar":
        v50 := v2[(((("foo bar" + "hello") + "\n\t") + "\n\t") + ("foo bar" + ("\n\t" + "a")))]
        _ = v50
    default:
        v5 -= v5
    }
    if (!true) {
        v51 := GenS5{F0: 18446744073709551615, F1: v4, F2: (-(v3))}
        switch "\n\t" {
        case "a":
            for fi52 := 0; fi52 < 1; fi52++ {
                var v53 uint32 = 967
                if false {
                    var v54 int8 = -75
                    var v55 uint32 = 0
                    _ = v54
                    _ = v55
                }
                _ = fi52
                _ = v53
            }
        }
        for fi56 := 0; fi56 < 5; fi56++ {
            for fi57 := 0; fi57 < 10; fi57++ {
                if (3.14159 < ((1.5 * (1.5 - (((1e10 * -2.25) + (float64(v5) + 3.14159)) / 1.5))) * -2.25)) {
                    if (false || (!(1 < 51))) {
                        v6 += 18446744073709551615
                        v5 = 0.0
                    } else {
                        if true {
                            var v58 int8 = 42
                            switch fi56 {
                            case 904:
                                v59 := GenS5{F0: 756, F1: (-(v4)), F2: 1}
                                _ = v59
                            case 32768:
                                var v60 uint = 65535
                                _ = v60
                            case 367:
                                switch "hello" {
                                case "foo bar":
                                    switch (-(fi56)) {
                                    case 0:
                                        delete(v2, "foo bar")
                                        var v61 float64 = -2.25
                                        _ = v61
                                    }
                                case "":
                                    if (!true) {
                                        v6 = (^(v6))
                                        switch "" {
                                        case "":
                                            var v62 int8 = v58
                                            _ = v62
                                        }
                                        fmt.Println(v58, v4, fi56)
                                    }
                                case "hello":
                                    v7 -= v7
                                    v63 := GenS8{F0: 215, F1: 255, F2: 255}
                                    v58 += 100
                                    _ = v63
                                default:
                                    var v64 string = ""
                                    _ = v64
                                }
                            }
                            _ = v58
                        }
                        v65 := GenS9{F0: 42, F1: ""}
                        _ = v65
                    }
                }
                if (!true) {
                    var v66 uint16 = 1
                    _ = len(v2)
                    v4 -= 888
                    _ = v66
                }
                var v67 float32 = (v5 + (v5 / 3.14159))
                _ = fi57
                _ = v67
            }
            _ = fi56
        }
        _ = v51
    } else {
        v68 := GenS6{F0: (v6 % ((v6) | 1)), F1: (v3 * v3)}
        if true {
            v4 -= (v4 * 837)
        } else {
            var v69 float32 = v5
            fmt.Println(v5, v6)
            var v70 bool = (!false)
            _ = v69
            _ = v70
        }
        var v71 int16 = -639
        _ = v68
        _ = v71
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v72 := GenS6{F0: 370, F1: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v72
}
