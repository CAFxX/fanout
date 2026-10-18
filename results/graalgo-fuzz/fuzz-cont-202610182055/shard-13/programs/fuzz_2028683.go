package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int; F1 uint32; F2 int16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int }
type GenS10 struct { F0 uint32; F1 bool; F2 int16 }

func main() {
    var v1 bool = false
    v2 := []uint32{}
    v3 := []string{}
    v4 := make(map[string]bool)
    var v5 int = 1
    var v6 uint8 = 1
    v1 = (-1 != 1)
    var v7 int = (v5 % ((v5) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "foo bar" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    case "\n\t":
        var v9 bool = v1
        _ = v9
    case "":
        _ = len(v2)
    default:
        v10 := GenS1{F0: 725}
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            if false {
                v12 := GenS2{F0: 42, F1: (!v1), F2: -310}
                var v13 int16 = 1000
                v13 -= 517
                _ = v12
                _ = v13
            } else {
                if v1 {
                    v14 := GenS1{F0: 953}
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v2 = append(v2, 4294967295)
                        var v16 bool = false
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v17 int = fi11
                        v18 := GenS3{F0: -1, F1: "\n\t"}
                        v4[("hello" + "\n\t")] = v16
                        switch -264 {
                        case 128:
                            v5 = (fi11 & (v5 - (v17 - 42)))
                        case 4294967295:
                            v7 -= (fi11 << ((255) & 63))
                        default:
                            if v1 {
                                v17 = fi11
                            } else {
                                switch ("" + "\n\t") {
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v19 int16 = 1
                                    if (v16 && v1) {
                                        if (739 == 337) {
                                            var v20 int64 = 1
                                            var v21 int8 = -1
                                            v4[(("" + (("" + "hello") + "")) + "\n\t")] = (42 > 4294967295)
                                            _ = v20
                                            _ = v21
                                        }
                                        v2 = append(v2, 2147483648)
                                        switch ("\n\t" + "a") {
                                        case "":
                                            v6 -= (v6 + v6)
                                        case "hello":
                                            switch ("" + ("\n\t" + (("hello" + "foo bar") + (("hello" + "hello") + "foo bar")))) {
                                            case "hello":
                                                v16 = v16
                                            case "foo bar":
                                                v16 = v1
                                            case "":
                                                v22 := GenS3{F0: 0, F1: ((("hello" + ("hello" + "")) + "foo bar") + "")}
                                                _ = v22
                                            default:
                                                switch (fi11 - v7) {
                                                case 888:
                                                    if (!v16) {
                                                        if v16 {
                                                            v19 -= v19
                                                            var v23 uint = uint(v6)
                                                            var v24 uint16 = uint16(v23)
                                                            var v25 string = ("\n\t" + (("hello" + (("\n\t" + "a") + "\n\t")) + ""))
                                                            _ = v23
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                        _ = len(v4)
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            if true {
                                                if (!true) {
                                                    v26 := GenS4{F0: v5, F1: 0, F2: -1}
                                                    _ = v26
                                                } else {
                                                    var v27 uint32 = 1
                                                    _ = v27
                                                }
                                            } else {
                                                v28 := GenS1{F0: 121}
                                                var v29 bool = v16
                                                var v30 int16 = (v19 + (v19 >> ((336) & 63)))
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            v31 := GenS5{F0: 104, F1: v19}
                                            _ = v31
                                        default:
                                            var v32 int16 = v19
                                            _ = v32
                                        }
                                    } else {
                                        v1 = ((127 > -1) || (false && false))
                                    }
                                    _ = v19
                                case "a":
                                    var v33 bool = (!(2147483647 <= 2))
                                    _ = v33
                                default:
                                    var v34 bool = v1
                                    var v35 string = "\n\t"
                                    _ = v34
                                    _ = v35
                                }
                            }
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                }
            }
            _ = fi11
        }
        v36 := GenS3{F0: -128, F1: "hello"}
        _ = v36
    } else {
        v5 = 930
        if ((v7 ^ -1) >= 0) {
            if ((v7 / ((v7) | 1)) < v5) {
                fmt.Println(v6, v5, v7)
                for fi37 := 0; fi37 < 18; fi37++ {
                    for fi38 := 0; fi38 < 6; fi38++ {
                        switch 4294967295 {
                        case 0:
                            v39 := GenS1{F0: 65535}
                            _ = v39
                        case 42:
                            _ = len(v3)
                        case -72:
                            var v40 int16 = -1
                            _ = v40
                        }
                        var v41 float32 = -2.25
                        v42 := GenS5{F0: 0, F1: 1}
                        _ = fi38
                        _ = v41
                        _ = v42
                    }
                    _ = fi37
                }
            } else {
                fmt.Println(v7)
                v1 = (255 <= 1)
            }
            switch ("\n\t" + (("a" + "foo bar") + (("a" + "") + "a"))) {
            case "hello":
                v1 = (!v1)
            case "a":
                for fi43 := 0; fi43 < 20; fi43++ {
                    for fi44 := 0; fi44 < 3; fi44++ {
                        if v1 {
                            v7 = (fi44 % (((v5 >> ((42) & 63))) | 1))
                            fmt.Println(fi43, fi43, v1)
                            v6 = v6
                        } else {
                            var v45 int32 = 1
                            var v46 uint32 = 1
                            var v47 float32 = 3.14159
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        var v48 int16 = -122
                        _ = fi44
                        _ = v48
                    }
                    _ = fi43
                }
            case "\n\t":
                v49 := GenS6{F0: 9223372036854775807}
                _ = v49
            default:
                switch ("" + "hello") {
                case "foo bar":
                    var v50 bool = true
                    _ = v50
                default:
                    if true {
                        v7 = v7
                    } else {
                        switch ("hello" + "foo bar") {
                        case "a":
                            if (!((824 <= 42) && v1)) {
                                v51 := GenS7{F0: 49, F1: (v5 >> ((0) & 63))}
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    v53 := GenS8{F0: 42}
                                    v6 = (v6 + 42)
                                    switch 2147483647 {
                                    case -623:
                                        v5 = (v5 & fi52)
                                    case 127:
                                        for fi54 := 0; fi54 < 4; fi54++ {
                                            v7 = -661
                                            _ = fi54
                                        }
                                    }
                                    _ = fi52
                                    _ = v53
                                }
                                _ = v51
                            }
                        default:
                            for fi55 := 0; fi55 < 16; fi55++ {
                                fmt.Println(v1)
                                _ = fi55
                            }
                        }
                        v6 += v6
                    }
                }
            }
            for fi56 := 0; fi56 < 13; fi56++ {
                if (v1 && (true && v1)) {
                    for fi57 := 0; fi57 < 12; fi57++ {
                        v7 = -424
                        fmt.Println(v7)
                        _ = fi57
                    }
                    var v58 float64 = 1.5
                    v58 -= 3.14159
                    _ = v58
                } else {
                    var v59 uint64 = 675
                    _ = len(v2)
                    if v1 {
                        for fi60 := 0; fi60 < 14; fi60++ {
                            var v61 int16 = -25
                            var v62 int32 = -2147483648
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        if true {
                            v63 := GenS4{F0: fi56, F1: 2, F2: 42}
                            _ = v63
                        }
                    } else {
                        fmt.Println(v1, v5)
                        v64 := GenS9{F0: 4294967295, F1: 42, F2: 764}
                        _ = v64
                    }
                    _ = v59
                }
                var v65 float32 = (((((1.5 - 3.14159) + 0.1) - 1.5) * 0.1) * 3.14159)
                _ = fi56
                _ = v65
            }
        } else {
            v66 := GenS7{F0: -606, F1: 128}
            if (65535 > 68) {
                switch ("\n\t" + "hello") {
                case "a":
                    if false {
                        v6 -= v6
                        var v67 int32 = 2147483647
                        _ = v67
                    } else {
                        v5 -= -1
                    }
                }
                if true {
                    if (552 < 18446744073709551615) {
                        var v68 int8 = 127
                        v69 := GenS3{F0: 0, F1: "hello"}
                        _ = v68
                        _ = v69
                    } else {
                        var v70 bool = (!(!(!(1.5 >= 0.0))))
                        _ = v70
                    }
                    fmt.Println(v7, v1, v5)
                    fmt.Println(v7, v7)
                } else {
                    fmt.Println(v7)
                    v5 += (v7 >> ((784) & 63))
                    if v1 {
                        v7 = -9223372036854775808
                        var v71 float64 = 1e10
                        _ = v71
                    }
                }
                v72 := GenS10{F0: 65535, F1: ((-370 >= 42) && (879 <= 32767)), F2: 1000}
                _ = v72
            }
            _ = v66
        }
        switch "foo bar" {
        case "hello":
            _ = len(v2)
        case "":
            v7 -= v7
        case "foo bar":
            v7 = v5
        default:
            v1 = v1
        }
    }
    switch (v7 & v7) {
    case -1:
        switch ("\n\t" + ("foo bar" + "\n\t")) {
        case "a":
            v6 = v6
        }
    case 127:
        v6 = v6
    }
    v1 = true
    switch (v5 & v7) {
    case 931:
        var v73 uint = 255
        _ = v73
    case -9223372036854775808:
        v4["\n\t"] = v1
    }
    v7 -= (v7 % (((v5 >> ((255) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
