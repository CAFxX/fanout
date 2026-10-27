package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 int16 = -333
    v2 := make(map[int]int)
    var v3 int64 = 0
    v4 := []string{}
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 15; fi6++ {
        v3 += v3
        _ = fi6
    }
    var v7 bool = true
    if (((1.5 + 1e10) / 1e10) != 0.0) {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch fi8 {
            case 0:
                var v9 float64 = float64(1.5)
                _ = v9
            case 32767:
                if (!v7) {
                    var v10 int64 = (-(v3))
                    _ = v10
                } else {
                    v1 += v1
                }
            }
            var v11 string = "a"
            _ = fi8
            _ = v11
        }
        var v12 float32 = 0.1
        if (int32(v3) <= 42) {
            v7 = v7
            fmt.Println(v12)
        } else {
            v13 := GenS1{F0: 4294967295, F1: 2, F2: false}
            v14 := GenS1{F0: -323, F1: 128, F2: (804 != 705)}
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS2{F0: "a", F1: -274, F2: 1}
                    fmt.Println(v1, fi15, fi16)
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            _ = v13
            _ = v14
        }
        _ = v12
    } else {
        v18 := GenS2{F0: (("foo bar" + ((("" + ("\n\t" + "")) + "\n\t") + ((("" + ("hello" + "\n\t")) + "a") + ("" + "foo bar")))) + "a"), F1: -686, F2: -1}
        _ = v18
    }
    var v19 float64 = 1.5
    v2[-463] = 1
    switch "" {
    case "\n\t":
        if (v3 < v3) {
            if (!(128 == 0)) {
                fmt.Println(v7)
                var v20 uint32 = 1
                _ = v20
            }
        }
    case "a":
        delete(v5, -469)
    }
    fmt.Println(v7, v19)
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS3{F0: ("hello" + "foo bar"), F1: 42, F2: 444}
        switch (fi21 + fi21) {
        case 0:
            v23 := GenS3{F0: "a", F1: 0, F2: 880}
            _ = v23
        }
        v7 = (!(-77 >= -1))
        _ = fi21
        _ = v22
    }
    _ = len(v5)
    var v24 bool = (false || (!(v19 == 1e10)))
    for fi25 := 0; fi25 < 17; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 string = "foo bar"
            _ = len(v4)
            v24 = v24
            var v28 string = v27
            switch ("hello" + "\n\t") {
            case "hello":
                var v29 uint8 = 1
                v30 := GenS1{F0: (-(fi25)), F1: 1, F2: (0 >= 461)}
                var v31 float64 = v19
                _ = v29
                _ = v30
                _ = v31
            case "":
                if (!(66 <= 83)) {
                    if v7 {
                        _ = len(v2)
                        var v32 uint = 65535
                        _ = v32
                    }
                    fmt.Println(fi25, v28, v24)
                    var v33 int16 = v1
                    _ = v33
                } else {
                    var v34 uint = 255
                    v34 = (^(v34))
                    _ = v34
                }
            default:
                v28 = v27
            }
            switch v28 {
            case "a":
                switch v28 {
                case "foo bar":
                    if ((v3 ^ v3) < v3) {
                        var v35 int8 = 42
                        switch fi25 {
                        case -2147483648:
                            switch (fi25 & int(v3)) {
                            case 75:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -870:
                                if v24 {
                                    switch v28 {
                                    case "foo bar":
                                        v7 = v7
                                    case "":
                                        switch "hello" {
                                        case "hello":
                                            if (v7 && (v24 || (false && false))) {
                                                switch (fi25 - fi26) {
                                                case 32767:
                                                    var v36 int64 = (v3 % ((v3) | 1))
                                                    _ = v36
                                                case 128:
                                                    var v37 uint16 = 1
                                                    _ = v37
                                                }
                                            }
                                        case "foo bar":
                                            v38 := GenS4{F0: 9223372036854775807, F1: 2147483647}
                                            _ = v38
                                        }
                                    default:
                                        fmt.Println(v3)
                                    }
                                } else {
                                    var v39 string = (v28 + v28)
                                    _ = v39
                                }
                            case -9223372036854775808:
                                v3 = 1
                            default:
                                var v40 int8 = v35
                                _ = v40
                            }
                        case -674:
                            if false {
                                var v41 uint8 = 0
                                v42 := GenS5{F0: (v41 * (v41 & 128)), F1: 2, F2: 2}
                                var v43 uint16 = 32767
                                _ = v41
                                _ = v42
                                _ = v43
                            } else {
                                switch (fi25 * fi26) {
                                case -1:
                                    v1 = (^(v1))
                                case 987:
                                    v28 = (v28 + ("" + v27))
                                default:
                                    v1 = v1
                                }
                            }
                        case 80:
                            if v7 {
                                if (!(!(-55 == 1))) {
                                    var v44 int = 32767
                                    v45 := v2[722]
                                    _ = v44
                                    _ = v45
                                }
                                switch -87 {
                                case 127:
                                    v35 -= v35
                                }
                                v46 := GenS2{F0: "\n\t", F1: int32(v1), F2: 0}
                                _ = v46
                            } else {
                                v7 = ((-(v3)) == v3)
                            }
                        default:
                            var v47 uint64 = 1
                            _ = v47
                        }
                        var v48 int = fi26
                        _ = v35
                        _ = v48
                    }
                case "\n\t":
                    fmt.Println(fi25)
                    var v49 float64 = ((v19 / 3.14159) * ((((v19 - v19) - ((((0.0 - (v19 + v19)) / v19) / (v19 / -2.25)) / 1.0)) - v19) / (v19 * (float64(3.14159) + 0.0))))
                    var v50 string = v28
                    var v51 int8 = 127
                    var v52 uint8 = 255
                    _ = v49
                    _ = v50
                    _ = v51
                    _ = v52
                case "hello":
                    switch ("" + v28) {
                    case "\n\t":
                        v1 = v1
                    }
                default:
                    switch v27 {
                    case "hello":
                        var v53 bool = (true || (156 != 143))
                        _ = v53
                    case "\n\t":
                        var v54 float32 = 0.1
                        _ = v54
                    default:
                        var v55 float32 = -2.25
                        _ = v55
                    }
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        var v56 uint64 = 1
        v1 = v1
        _ = fi25
        _ = v56
    }
    fmt.Println(v1, v19)
    if true {
        v57 := GenS6{F0: 2147483648}
        _ = v57
    }
    switch ("" + "a") {
    case "":
        var v58 float64 = float64(0.1)
        _ = v58
    case "foo bar":
        v1 -= (v1 & (^(v1)))
    case "hello":
        v59 := GenS2{F0: "hello", F1: -872, F2: 0}
        _ = v59
    default:
        if v24 {
            v60 := GenS7{F0: 243}
            switch 127 {
            case 4294967295:
                switch 127 {
                case -2147483648:
                    v19 -= (v19 - 1e10)
                case 2147483647:
                    var v61 uint32 = 2
                    _ = v61
                case 400:
                    if v24 {
                        v7 = (32768 == int(v1))
                        if v24 {
                            v3 = 9223372036854775807
                            v62 := GenS1{F0: -9223372036854775808, F1: 255, F2: (!((("a" + ("" + ("foo bar" + ("foo bar" + "foo bar")))) <= (("a" + "hello") + "\n\t")) && (228 > 1)))}
                            fmt.Println(v3, v7, v19)
                            _ = v62
                        }
                        if v7 {
                            fmt.Println(v24, v1)
                            switch -2147483648 {
                            case -2147483648:
                                fmt.Println(v19)
                            }
                        }
                    } else {
                        var v63 uint16 = 255
                        _ = v63
                    }
                default:
                    v1 = (^(v1))
                }
            default:
                v1 = 42
            }
            var v64 int = 9223372036854775807
            _ = v60
            _ = v64
        } else {
            if false {
                var v65 uint = 761
                fmt.Println(v1)
                _ = v65
            }
            var v66 uint64 = 334
            for fi67 := 0; fi67 < 10; fi67++ {
                var v68 float64 = -2.25
                _ = fi67
                _ = v68
            }
            _ = v66
        }
    }
    var v69 uint32 = 99
    v24 = v7
    var v70 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v7
    _ = v70
}
