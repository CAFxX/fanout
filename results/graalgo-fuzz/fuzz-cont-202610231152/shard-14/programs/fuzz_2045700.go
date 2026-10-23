package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 string; F1 string; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint64; F2 uint }
type GenS9 struct { F0 int8; F1 uint16; F2 int64 }

func main() {
    var v1 int64 = 40
    var v2 int16 = -32768
    var v3 uint16 = 943
    v4 := []string{"hello"}
    var v5 float32 = (-2.25 + 0.1)
    var v6 bool = false
    if (1 < 127) {
        switch "hello" {
        case "foo bar":
            fmt.Println(v6, v6, v1)
        case "hello":
            v4 = append(v4, "")
        default:
            var v7 int32 = 2147483647
            _ = v7
        }
        v8 := GenS1{F0: -59, F1: "a", F2: 255}
        if (false || v6) {
            v9 := GenS2{F0: 1}
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 int = (fi10 % ((fi10) | 1))
                if v6 {
                    switch -475 {
                    case 128:
                        if (0 < 127) {
                            var v12 uint32 = 0
                            switch -2147483648 {
                            case 127:
                                if (4294967295 >= 4294967295) {
                                    for fi13 := 0; fi13 < 5; fi13++ {
                                        var v14 int = (v11 & 2147483647)
                                        var v15 int16 = (v2 & v2)
                                        if (255 > uint8(v3)) {
                                            v6 = v6
                                            var v16 bool = (127 == 21)
                                            var v17 float64 = -2.25
                                            v4 = append(v4, ("" + "hello"))
                                            _ = v16
                                            _ = v17
                                        } else {
                                            v18 := GenS2{F0: 128}
                                            var v19 int8 = 1
                                            _ = v18
                                            _ = v19
                                        }
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                    }
                                    v4 = append(v4, "")
                                } else {
                                    switch 128 {
                                    case 262:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v20 := GenS3{F0: -128, F1: 18446744073709551615}
                                    _ = v20
                                }
                            default:
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                            }
                            _ = v12
                        } else {
                            v21 := GenS4{F0: (v3 | 42), F1: -1}
                            for fi22 := 0; fi22 < 1; fi22++ {
                                var v23 uint8 = 1
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        }
                    }
                } else {
                    var v24 uint16 = (v3 - 563)
                    if (!(!(42 == 4294967295))) {
                        if v6 {
                            var v25 bool = true
                            v11 += (-(fi10))
                            _ = v25
                        } else {
                            var v26 uint64 = 223
                            _ = v26
                        }
                    }
                    v27 := GenS5{F0: v6}
                    _ = v24
                    _ = v27
                }
                v6 = v6
                _ = fi10
                _ = v11
            }
            v28 := GenS6{F0: 42, F1: 1}
            _ = v28
            _ = v9
        }
        _ = v8
    }
    v6 = (-1 <= 42)
    var v29 bool = v6
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS6{F0: 18446744073709551615, F1: 2}
        for fi32 := 0; fi32 < 9; fi32++ {
            if ((-32768 <= v2) || false) {
                v33 := GenS5{F0: (3.14159 <= 1.5)}
                var v34 string = "foo bar"
                var v35 uint = 65535
                _ = v33
                _ = v34
                _ = v35
            }
            var v36 int = fi32
            _ = fi32
            _ = v36
        }
        _ = fi30
        _ = v31
    }
    var v37 uint32 = 2
    if (v6 && false) {
        if v6 {
            var v38 int16 = v2
            var v39 bool = v6
            for fi40 := 0; fi40 < 7; fi40++ {
                var v41 int32 = 400
                _ = fi40
                _ = v41
            }
            _ = v38
            _ = v39
        } else {
            var v42 uint32 = v37
            _ = v42
        }
    } else {
        v43 := GenS7{F0: ("foo bar" + "hello"), F1: ("" + "foo bar"), F2: (v37 % ((4294967295) | 1))}
        v6 = (-2.25 <= 3.14159)
        if (!(-26 == 127)) {
            if ((^(v37)) >= (v37 << ((582) & 63))) {
                switch "a" {
                case "\n\t":
                    v44 := GenS8{F0: -115, F1: 744, F2: 255}
                    _ = v44
                case "foo bar":
                    v37 -= v37
                case "":
                    fmt.Println(v6)
                default:
                    for fi45 := 0; fi45 < 12; fi45++ {
                        var v46 uint8 = 159
                        _ = fi45
                        _ = v46
                    }
                }
                v5 = 1e10
                var v47 int = 32767
                _ = v47
            } else {
                if true {
                    var v48 uint = 65535
                    var v49 int16 = (v2 * v2)
                    _ = v48
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 14; fi50++ {
                        for fi51 := 0; fi51 < 20; fi51++ {
                            var v52 float64 = 1e10
                            _ = fi51
                            _ = v52
                        }
                        switch fi50 {
                        case 9223372036854775807:
                            v53 := GenS9{F0: 95, F1: (v3 | v3), F2: 9223372036854775807}
                            _ = v53
                        case -494:
                            switch int(v1) {
                            case 1:
                                var v54 int = 4294967295
                                _ = v54
                            }
                        }
                        _ = fi50
                    }
                    if (v6 && (!true)) {
                        switch 985 {
                        case 1:
                            v29 = (-2147483648 >= 0)
                        default:
                            fmt.Println(v1, v5)
                        }
                        v37 += v37
                    }
                    v3 = (^(v3))
                }
                if true {
                    if false {
                        var v55 uint8 = 42
                        fmt.Println(v1, v37, v29)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v55
                    } else {
                        switch 42 {
                        case -153:
                            for fi56 := 0; fi56 < 16; fi56++ {
                                fmt.Println(v3, v6)
                                var v57 int = 128
                                v3 -= 0
                                _ = fi56
                                _ = v57
                            }
                        default:
                            var v58 int32 = 0
                            _ = v58
                        }
                    }
                    switch 687 {
                    case 2147483647:
                        v5 = v5
                    case -2147483648:
                        switch ("\n\t" + "\n\t") {
                        case "":
                            var v59 int16 = 0
                            _ = v59
                        case "a":
                            v37 += (v37 / ((v37) | 1))
                        case "\n\t":
                            v3 -= (v3 ^ (v3 / (((v3 ^ v3)) | 1)))
                        default:
                            switch "a" {
                            case "":
                                v37 -= v37
                            case "a":
                                v2 = 32767
                            }
                        }
                    default:
                        v4 = append(v4, "\n\t")
                    }
                    for fi60 := 0; fi60 < 19; fi60++ {
                        _ = len(v4)
                        _ = fi60
                    }
                }
            }
            var v61 int64 = v1
            if v6 {
                if false {
                    switch 4294967295 {
                    case 603:
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            v2 += v2
                        case -2147483648:
                            var v62 int8 = 100
                            _ = v62
                        case 4294967295:
                            v6 = v29
                        }
                    case 574:
                        var v63 int8 = -1
                        _ = v63
                    case 153:
                        var v64 uint = 2
                        _ = v64
                    }
                } else {
                    for fi65 := 0; fi65 < 8; fi65++ {
                        v66 := GenS6{F0: 185, F1: 474}
                        _ = fi65
                        _ = v66
                    }
                }
            } else {
                fmt.Println(v5, v1, v29)
                fmt.Println(v6)
                var v67 uint64 = 28
                _ = v67
            }
            _ = v61
        } else {
            var v68 int64 = v1
            var v69 uint8 = 2
            _ = v68
            _ = v69
        }
        _ = v43
    }
    v4 = append(v4, "foo bar")
    switch "\n\t" {
    case "foo bar":
        v6 = (197 == 251)
    }
    fmt.Println(v1, v29)
    if true {
        for fi70 := 0; fi70 < 11; fi70++ {
            var v71 bool = true
            v5 += 3.14159
            _ = fi70
            _ = v71
        }
        v29 = v6
    } else {
        var v72 int = 127
        if false {
            var v73 uint16 = (v3 & v3)
            v37 += v37
            _ = v73
        } else {
            for fi74 := 0; fi74 < 7; fi74++ {
                v72 = -1
                if len(v4) > 0 {
                    v4[0] = ("a" + ("\n\t" + ("foo bar" + ("hello" + "foo bar"))))
                }
                v75 := GenS2{F0: 42}
                _ = fi74
                _ = v75
            }
            v76 := GenS6{F0: 4294967295, F1: 2}
            var v77 int8 = -1
            _ = v76
            _ = v77
        }
        _ = v72
    }
    var v78 bool = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
