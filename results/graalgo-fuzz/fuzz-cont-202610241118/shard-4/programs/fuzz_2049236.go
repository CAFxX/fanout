package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint64; F2 string }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = (3.14159 / 1.5)
    v3 := make(map[string]string)
    v4 := []uint32{}
    var v5 uint16 = 0
    v6 := []int{2147483647, -1}
    var v7 int8 = 30
    v7 = (-(v7))
    var v8 float32 = v2
    for fi9 := 0; fi9 < 20; fi9++ {
        if (!(v5 >= 255)) {
            var v10 string = v1
            if len(v4) > 0 {
                v4[0] = 2
            }
            _ = len(v3)
            _ = v10
        } else {
            delete(v3, ((("" + v1) + v1) + ("\n\t" + ("" + (v1 + v1)))))
            v8 = ((v8 - ((float32(1.5) + (1e10 - v2)) * (v8 - v8))) + (v8 * 1.5))
            if false {
                var v11 uint16 = (^(v5))
                switch (v1 + "hello") {
                case "foo bar":
                    fmt.Println(v1)
                default:
                    if true {
                        v12 := GenS1{F0: 128, F1: (v5 * 301)}
                        for fi13 := 0; fi13 < 13; fi13++ {
                            var v14 uint = 1
                            _ = fi13
                            _ = v14
                        }
                        _ = v12
                    }
                }
                if (true || (779 < 2147483648)) {
                    fmt.Println(v2, v2, v1)
                } else {
                    v11 += 255
                }
                _ = v11
            }
        }
        switch fi9 {
        case 32768:
            if false {
                for fi15 := 0; fi15 < 15; fi15++ {
                    if false {
                        var v16 bool = (fi9 > 128)
                        _ = len(v3)
                        switch "" {
                        case "a":
                            switch v1 {
                            case "foo bar":
                                switch (("\n\t" + ("a" + v1)) + (("" + v1) + v1)) {
                                case "hello":
                                    if v16 {
                                        var v17 int16 = 301
                                        v2 += v2
                                        _ = v17
                                    } else {
                                        fmt.Println(v1)
                                    }
                                default:
                                    if v16 {
                                        v18 := GenS1{F0: 128, F1: v5}
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: 1, F1: (v5 * (v5 & 65535))}
                                        switch v1 {
                                        case "":
                                            v20 := GenS2{F0: (v7 / ((-128) | 1))}
                                            _ = v20
                                        case "\n\t":
                                            if (69 <= 25) {
                                                v2 += -2.25
                                            }
                                        case "a":
                                            var v21 uint = 255
                                            _ = v21
                                        default:
                                            var v22 uint = 255
                                            _ = v22
                                        }
                                        _ = v19
                                    }
                                }
                            case "":
                                v5 = (^(v5))
                            }
                        }
                        _ = v16
                    } else {
                        switch fi9 {
                        case 32767:
                            var v23 int16 = int16(fi15)
                            _ = v23
                        default:
                            var v24 int32 = 1
                            var v25 uint8 = 42
                            v7 = v7
                            _ = v24
                            _ = v25
                        }
                        if ((true && ((!true) || (!false))) && true) {
                            if len(v6) > 0 {
                                v6[0] = fi15
                            }
                            if false {
                                v7 -= (v7 * (v7 << ((69) & 63)))
                                _ = len(v4)
                            }
                            _ = len(v3)
                        } else {
                            var v26 uint16 = 1
                            v26 += 2
                            _ = v26
                        }
                        v3[""] = ("" + (v1 + (("" + v1) + v1)))
                        if true {
                            v8 = v8
                            if true {
                                if false {
                                    v5 = (v5 >> ((843) & 63))
                                    if (531 != 0) {
                                        var v27 int8 = v7
                                        _ = len(v6)
                                        v28 := GenS2{F0: v27}
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v1 = (v1 + "hello")
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v2 = 1.5
                                } else {
                                    _ = len(v6)
                                }
                                if (265 > 1) {
                                    v29 := GenS2{F0: 127}
                                    _ = v29
                                }
                            } else {
                                switch fi9 {
                                case 32768:
                                    v2 = v2
                                default:
                                    fmt.Println(v8)
                                    if (!false) {
                                        fmt.Println(v7, fi9)
                                        v30 := v3[v1]
                                        v31 := GenS3{F0: (fi15 - (fi9 * (fi15 << ((65535) & 63)))), F1: (^(v5)), F2: 127}
                                        fmt.Println(fi9, v30, v1)
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if (!(false && true)) {
                                            var v32 uint32 = 2147483648
                                            _ = v32
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = (fi9 << ((42) & 63))
                                        }
                                    }
                                }
                                var v33 string = ("" + v1)
                                _ = v33
                            }
                        }
                    }
                    if true {
                        var v34 uint8 = 0
                        _ = v34
                    }
                    _ = fi15
                }
                var v35 bool = true
                _ = v35
            } else {
                v8 -= ((v8 - 3.14159) * 3.14159)
                var v36 float32 = float32((3.14159 * 1.5))
                if true {
                    fmt.Println(v5, v5)
                    _ = len(v4)
                    _ = len(v4)
                }
                _ = v36
            }
        }
        _ = fi9
    }
    delete(v3, v1)
    v1 = v1
    v37 := GenS4{F0: v7, F1: ("\n\t" + (v1 + v1))}
    var v38 uint32 = 892
    v2 += float32(0.1)
    if (-2147483648 == 1) {
        for fi39 := 0; fi39 < 17; fi39++ {
            delete(v3, (v1 + v1))
            _ = fi39
        }
        fmt.Println(v5, v5)
        switch 578 {
        case -822:
            switch -1 {
            case -111:
                if (((false || (!false)) || false) || (!false)) {
                    v38 = uint32(v5)
                } else {
                    v40 := v3[v1]
                    for fi41 := 0; fi41 < 15; fi41++ {
                        var v42 string = ("a" + v1)
                        for fi43 := 0; fi43 < 9; fi43++ {
                            v44 := GenS1{F0: 255, F1: v5}
                            _ = fi43
                            _ = v44
                        }
                        if (!false) {
                            for fi45 := 0; fi45 < 17; fi45++ {
                                switch fi45 {
                                case -906:
                                    var v46 int16 = 0
                                    _ = v46
                                }
                                v1 = (v42 + "")
                                if (1 >= 1) {
                                    var v47 uint32 = (v38 + 4294967295)
                                    _ = v47
                                } else {
                                    var v48 uint64 = 1
                                    _ = v48
                                }
                                _ = fi45
                            }
                        } else {
                            for fi49 := 0; fi49 < 16; fi49++ {
                                v8 = (0.1 * (v8 + ((v8 * (v2 + (v2 * -2.25))) - float32(1e10))))
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                switch 933 {
                                case 0:
                                    v40 = v40
                                }
                                _ = fi49
                            }
                            var v50 float64 = (float64(0.0) - 0.1)
                            v40 = ((v1 + v42) + "")
                            _ = v50
                        }
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
            default:
                fmt.Println(v2, v1, v8)
            }
        case -562:
            for fi51 := 0; fi51 < 9; fi51++ {
                v6 = append(v6, (^(fi51)))
                var v52 uint32 = v38
                _ = fi51
                _ = v52
            }
        case -1:
            fmt.Println(v7, v7, v8)
        }
    }
    if false {
        delete(v3, v1)
        if false {
            var v53 uint64 = 944
            _ = v53
        }
    } else {
        v54 := GenS4{F0: v7, F1: v1}
        if true {
            v38 -= (v38 + (^(v38)))
            if ((9223372036854775807 > -1000) && true) {
                if true {
                    v8 -= ((v8 + 0.1) * (1.5 - float32(0.0)))
                } else {
                    switch v1 {
                    case "hello":
                        for fi55 := 0; fi55 < 6; fi55++ {
                            for fi56 := 0; fi56 < 3; fi56++ {
                                switch (fi55 >> ((42) & 63)) {
                                case -40:
                                    switch (fi55 >> ((65535) & 63)) {
                                    case -314:
                                        v57 := GenS2{F0: (v7 & (-(v7)))}
                                        _ = v57
                                    }
                                case 9223372036854775807:
                                    var v58 int8 = (^(v7))
                                    _ = v58
                                case 127:
                                    var v59 string = v1
                                    _ = v59
                                }
                                var v60 uint32 = v38
                                _ = fi56
                                _ = v60
                            }
                            v38 = (v38 - (v38 % ((42) | 1)))
                            v61 := GenS5{F0: (fi55 & fi55)}
                            _ = fi55
                            _ = v61
                        }
                    case "":
                        v62 := GenS6{F0: 1, F1: 42, F2: v1}
                        _ = v62
                    case "a":
                        for fi63 := 0; fi63 < 2; fi63++ {
                            if false {
                                v38 = 584
                                if len(v6) > 0 {
                                    v6[0] = fi63
                                }
                                var v64 uint64 = 862
                                _ = v64
                            } else {
                                v65 := GenS4{F0: (v7 >> ((18446744073709551615) & 63)), F1: ((v1 + v1) + ("foo bar" + "a"))}
                                _ = v65
                            }
                            for fi66 := 0; fi66 < 5; fi66++ {
                                v38 += (v38 / ((v38) | 1))
                                v8 += 1e10
                                if ((v1 <= (v1 + "\n\t")) && (!true)) {
                                    if ((1.5 * (0.1 - -2.25)) == 1e10) {
                                        v2 += 1.5
                                        v8 += 0.1
                                        if true {
                                            v67 := GenS5{F0: (fi66 % ((32767) | 1))}
                                            _ = v67
                                        }
                                    }
                                    var v68 int64 = -9223372036854775808
                                    _ = v68
                                }
                                var v69 bool = (((false || (42 >= 202)) && false) && (!true))
                                switch 998 {
                                case 9223372036854775807:
                                    v1 = v1
                                case -9223372036854775808:
                                    v70 := GenS2{F0: v7}
                                    _ = v70
                                case -2147483648:
                                    var v71 int32 = 577
                                    _ = v71
                                }
                                var v72 uint8 = 255
                                _ = fi66
                                _ = v69
                                _ = v72
                            }
                            v3[v1] = ("" + v1)
                            _ = fi63
                        }
                    default:
                        var v73 uint16 = 65535
                        _ = v73
                    }
                    v7 = v7
                }
                v2 = ((((0.0 * 3.14159) / v8) * (1.5 + v8)) - -2.25)
            }
        } else {
            for fi74 := 0; fi74 < 19; fi74++ {
                for fi75 := 0; fi75 < 15; fi75++ {
                    v8 -= 1.5
                    _ = fi75
                }
                _ = fi74
            }
            if (v5 <= (v5 | v5)) {
                fmt.Println(v2, v38)
                v76 := GenS2{F0: (v7 | (^(v7)))}
                _ = v76
            } else {
                v1 = ("" + "hello")
                fmt.Println(v1, v5)
            }
        }
        v77 := v3[("hello" + "a")]
        _ = v54
        _ = v77
    }
    v38 = v38
    var v78 uint8 = 111
    fmt.Println(v78, v8, v5)
    v8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
    _ = v8
}
