package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int; F1 uint }

func main() {
    var v1 int64 = 574
    v2 := []uint32{}
    var v3 bool = (!(0 < 831))
    v4 := GenS1{F0: -1, F1: "a", F2: 943}
    if false {
        var v5 int32 = -2147483648
        if (!true) {
            for fi6 := 0; fi6 < 11; fi6++ {
                v2 = append(v2, 2)
                var v7 int = fi6
                var v8 int = 9223372036854775807
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            var v9 float32 = 0.0
            v3 = v3
            _ = v9
        }
        switch -1 {
        case 127:
            v1 += 9223372036854775807
        case 32768:
            var v10 int = 1
            _ = v10
        case 0:
            v11 := GenS1{F0: 100, F1: "\n\t", F2: (v1 / (((v1 - v1)) | 1))}
            _ = v11
        }
        _ = v5
    }
    var v12 int8 = 127
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = (32768 == 447)
    if v3 {
        if (v3 || true) {
            if true {
                var v13 float64 = 0.1
                if false {
                    var v14 float64 = v13
                    fmt.Println(v3, v13)
                    _ = v14
                }
                var v15 float32 = 0.0
                _ = v13
                _ = v15
            } else {
                var v16 bool = false
                for fi17 := 0; fi17 < 19; fi17++ {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        _ = len(v2)
                        switch "a" {
                        case "a":
                            v19 := GenS2{F0: v1, F1: 1}
                            _ = v19
                        }
                        _ = fi18
                    }
                    switch -1 {
                    case 128:
                        v1 = 1
                    case -237:
                        if (3.14159 == (float32((0.0 / 1.0)) - (0.1 + 1e10))) {
                            var v20 int32 = 397
                            var v21 int32 = (-(v20))
                            var v22 uint32 = 212
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    default:
                        v23 := GenS3{F0: v12, F1: v12}
                        _ = v23
                    }
                    _ = fi17
                }
                _ = v16
            }
            v1 = v1
            var v24 int32 = 42
            _ = v24
        }
        var v25 string = (("foo bar" + ("\n\t" + "a")) + "a")
        for fi26 := 0; fi26 < 6; fi26++ {
            if v3 {
                switch v25 {
                case "hello":
                    v3 = v3
                case "foo bar":
                    switch 42 {
                    case 32767:
                        var v27 uint8 = 162
                        _ = v27
                    case 678:
                        for fi28 := 0; fi28 < 9; fi28++ {
                            v3 = v3
                            _ = fi28
                        }
                    case 4294967295:
                        var v29 uint8 = 204
                        _ = v29
                    }
                case "":
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                default:
                    v3 = false
                }
            } else {
                v12 -= v12
                if v3 {
                    var v30 int64 = v1
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int16 = -1
                        var v33 float32 = -2.25
                        switch 4294967295 {
                        case 942:
                            if (((v12 * (v12 << ((575) & 63))) > 100) || (!((v30 & (v30 + v30)) == (^(v30))))) {
                                var v34 uint64 = 992
                                fmt.Println(v12)
                                v35 := GenS2{F0: (v30 << ((1) & 63)), F1: 65535}
                                _ = v34
                                _ = v35
                            }
                        case 2147483647:
                            var v36 uint16 = 255
                            _ = v36
                        case 4294967295:
                            switch (fi31 << ((4294967295) & 63)) {
                            case 4294967295:
                                v32 = 0
                            case 32767:
                                fmt.Println(v1, v12, fi26)
                            case 2147483647:
                                switch (v25 + (((v25 + "a") + "hello") + "\n\t")) {
                                case "a":
                                    v37 := GenS2{F0: (v30 ^ (v30 | (-(v30)))), F1: 868}
                                    v2 = append(v2, 2)
                                    _ = v37
                                case "\n\t":
                                    if (255 >= 1) {
                                        var v38 uint8 = 128
                                        v2 = append(v2, 2147483647)
                                        _ = v38
                                    }
                                case "foo bar":
                                    fmt.Println(fi31, v33)
                                default:
                                    _ = len(v2)
                                }
                            default:
                                var v39 int16 = v32
                                _ = v39
                            }
                        default:
                            var v40 uint = 0
                            _ = v40
                        }
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v41 := GenS4{F0: v3, F1: -32768}
                    _ = v30
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        fmt.Println(fi26, v1, v3)
                        _ = fi42
                    }
                    v3 = (42 <= (fi26 + fi26))
                }
                if ((v3 || v3) || false) {
                    _ = len(v2)
                    var v43 int32 = 2147483647
                    _ = v43
                }
            }
            _ = fi26
        }
        _ = v25
    } else {
        v44 := GenS3{F0: v12, F1: (v12 << ((42) & 63))}
        if v3 {
            if v3 {
                switch 270 {
                case 495:
                    var v45 uint32 = 617
                    _ = v45
                case 32768:
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            var v46 int32 = -741
                            _ = v46
                        case "":
                            v3 = (32767 > 32767)
                        }
                        var v47 uint64 = 907
                        if true {
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                v47 -= v47
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        _ = v47
                    }
                case 128:
                    var v48 int16 = 139
                    _ = v48
                default:
                    for fi49 := 0; fi49 < 8; fi49++ {
                        v50 := GenS1{F0: (v12 / (((v12 ^ v12)) | 1)), F1: "\n\t", F2: 42}
                        _ = fi49
                        _ = v50
                    }
                }
                var v51 int32 = -508
                v51 = (v51 & (^(v51)))
                _ = v51
            } else {
                _ = len(v2)
                v52 := GenS1{F0: v12, F1: "foo bar", F2: v1}
                _ = v52
            }
            if (1 >= 4294967295) {
                v1 = v1
            }
            var v53 string = ((("a" + ("" + "")) + ("hello" + "a")) + "\n\t")
            _ = v53
        }
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 float64 = 1e10
            v1 = (v1 | -931)
            v55 = v55
            _ = fi54
            _ = v55
        }
        _ = v44
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        v57 := GenS5{F0: (fi56 | fi56), F1: 65535}
        _ = fi56
        _ = v57
    }
    var v58 int8 = v12
    v12 = (-(v58))
    for fi59 := 0; fi59 < 4; fi59++ {
        v58 += (v12 - (v12 >> ((997) & 63)))
        _ = fi59
    }
    v3 = true
    fmt.Println(v58, v12)
    fmt.Println(v12, v1)
    if true {
        v12 = 42
        if v3 {
            v1 = (v1 - v1)
            for fi60 := 0; fi60 < 4; fi60++ {
                var v61 string = ("\n\t" + "hello")
                _ = fi60
                _ = v61
            }
            if false {
                var v62 float32 = (3.14159 - 0.0)
                switch ("a" + (("hello" + ("" + "foo bar")) + ("\n\t" + "foo bar"))) {
                case "foo bar":
                    fmt.Println(v3, v1)
                default:
                    v58 = v58
                }
                _ = v62
            } else {
                fmt.Println(v12, v3, v12)
            }
        }
        var v63 int8 = (v58 ^ (v12 << ((130) & 63)))
        _ = v63
    } else {
        v1 = v1
        var v64 int16 = 0
        v65 := GenS5{F0: 128, F1: 227}
        _ = v64
        _ = v65
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v58
}
