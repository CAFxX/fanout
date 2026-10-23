package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    var v1 float32 = (0.0 - (3.14159 - 3.14159))
    var v2 int64 = -1
    v3 := []string{("foo bar" + "a"), ""}
    v4 := make(map[uint32]uint32)
    switch -9223372036854775808 {
    case 333:
        fmt.Println(v1, v1)
    case 9223372036854775807:
        v5 := GenS1{F0: 255}
        _ = v5
    case -550:
        v1 = v1
    default:
        v6 := GenS1{F0: 180}
        _ = v6
    }
    if true {
        switch "hello" {
        case "":
            var v7 int64 = (v2 % (((^(v2))) | 1))
            _ = v7
        case "foo bar":
            fmt.Println(v1, v1)
        }
        var v8 int64 = (v2 & v2)
        v9 := GenS1{F0: 901}
        _ = v8
        _ = v9
    } else {
        v2 = v2
        if false {
            var v10 uint16 = 468
            var v11 int16 = 32767
            v12 := GenS1{F0: 42}
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: 1, F1: 2}
        if false {
            v1 = 0.1
            switch ("" + ((("foo bar" + ((("" + ("a" + ("\n\t" + (("hello" + "") + "foo bar")))) + (("" + ("\n\t" + (("hello" + "hello") + ""))) + ("hello" + ("a" + ("a" + "a"))))) + "\n\t")) + (("" + (("a" + "\n\t") + (((("a" + ("\n\t" + "hello")) + "hello") + "") + "hello"))) + "foo bar")) + ("\n\t" + ""))) {
            case "foo bar":
                var v15 uint64 = 0
                _ = v15
            case "a":
                v2 -= v2
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if false {
                    var v17 uint32 = 487
                    var v18 int64 = (v2 / (((-(v2))) | 1))
                    switch "hello" {
                    case "\n\t":
                        var v19 uint64 = 4294967295
                        if ((fi16 / ((2147483647) | 1)) >= (fi13 | (fi13 & (fi16 << ((uint(v17)) & 63))))) {
                            v1 -= v1
                            var v20 bool = false
                            _ = v20
                        } else {
                            _ = len(v3)
                            var v21 bool = (1.5 > 0.1)
                            _ = v21
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v19
                    }
                    _ = v17
                    _ = v18
                } else {
                    var v22 uint16 = 697
                    var v23 bool = (128 != 58)
                    v24 := GenS1{F0: 42}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = len(v3)
                if (!true) {
                    v1 += 0.1
                    if (!false) {
                        v2 = v2
                        v1 = (v1 / ((v1 + float32(((0.1 + 1e10) + 1.5))) - (0.1 + 1e10)))
                        if false {
                            fmt.Println(fi13, fi13, v2)
                        }
                    } else {
                        v1 = v1
                    }
                }
                _ = fi16
            }
        } else {
            switch fi13 {
            case 9223372036854775807:
                fmt.Println(v1)
            case -9223372036854775808:
                if true {
                    var v25 uint64 = 42
                    var v26 uint = 42
                    _ = v25
                    _ = v26
                } else {
                    v27 := GenS3{F0: 2}
                    _ = v27
                }
            case 42:
                _ = len(v4)
            default:
                var v28 float64 = float64(0.0)
                _ = v28
            }
        }
        delete(v4, 329)
        _ = fi13
        _ = v14
    }
    v29 := GenS4{F0: 42, F1: 68, F2: 2147483647}
    fmt.Println(v2)
    if true {
        var v30 float64 = float64((v1 - float32(0.1)))
        var v31 string = (("hello" + (((("" + "hello") + "a") + ("hello" + "foo bar")) + "foo bar")) + "")
        if (false || true) {
            v30 += (v30 / (float64(v1) * (v30 - (float64(v1) * 1.5))))
            switch -9223372036854775808 {
            case 32767:
                if (1 <= 89) {
                    if (!(157 != 2)) {
                        v32 := GenS5{F0: 741}
                        v30 -= ((v30 * (v30 / v30)) * v30)
                        v33 := GenS2{F0: 349, F1: 255}
                        _ = v32
                        _ = v33
                    }
                }
            case -866:
                if true {
                    if (!(255 >= 0)) {
                        if true {
                            fmt.Println(v31)
                            var v34 int = 1
                            for fi35 := 0; fi35 < 18; fi35++ {
                                for fi36 := 0; fi36 < 9; fi36++ {
                                    v30 = v30
                                    var v37 float32 = 0.1
                                    var v38 uint32 = 592
                                    v39 := GenS3{F0: 55}
                                    v31 = v31
                                    var v40 uint16 = 65535
                                    v41 := GenS1{F0: 588}
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                                var v42 uint16 = 65535
                                v43 := GenS4{F0: -32768, F1: v42, F2: 0}
                                _ = fi35
                                _ = v42
                                _ = v43
                            }
                            _ = v34
                        } else {
                            v44 := GenS6{F0: 1, F1: true, F2: -2147483648}
                            for fi45 := 0; fi45 < 12; fi45++ {
                                v2 = v2
                                v31 = "a"
                                _ = fi45
                            }
                            _ = v44
                        }
                        var v46 uint = 42
                        _ = v46
                    }
                    var v47 uint32 = 42
                    _ = v47
                }
            case 2147483647:
                fmt.Println(v31, v2, v30)
            default:
                if true {
                    fmt.Println(v31, v2)
                    for fi48 := 0; fi48 < 2; fi48++ {
                        if true {
                            v2 = v2
                            fmt.Println(v1, v2, v30)
                        }
                        switch ("a" + "a") {
                        case "":
                            for fi49 := 0; fi49 < 16; fi49++ {
                                v2 = 1
                                v31 = "hello"
                                v31 = "a"
                                _ = fi49
                            }
                        case "foo bar":
                            var v50 float32 = v1
                            _ = v50
                        case "\n\t":
                            v51 := GenS2{F0: 1, F1: 72}
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 7; fi52++ {
                            if false {
                                delete(v4, 2147483647)
                                _ = len(v4)
                            }
                            _ = fi52
                        }
                        _ = fi48
                    }
                    for fi53 := 0; fi53 < 15; fi53++ {
                        switch fi53 {
                        case 32767:
                            v1 = v1
                        case 1:
                            if true {
                                for fi54 := 0; fi54 < 12; fi54++ {
                                    v4[1] = 2
                                    v55 := GenS1{F0: 2}
                                    _ = fi54
                                    _ = v55
                                }
                            }
                        default:
                            if true {
                                var v56 int = fi53
                                v1 += (1.5 / ((v1 + (v1 + (0.1 / v1))) - 1.5))
                                _ = v56
                            }
                        }
                        _ = fi53
                    }
                }
            }
            v57 := GenS5{F0: 1}
            _ = v57
        }
        _ = v30
        _ = v31
    } else {
        for fi58 := 0; fi58 < 6; fi58++ {
            v59 := GenS5{F0: 586}
            _ = fi58
            _ = v59
        }
        v2 -= v2
    }
    var v60 int = 32767
    var v61 uint64 = 4294967295
    var v62 float32 = v1
    v63 := GenS3{F0: 42}
    for fi64 := 0; fi64 < 5; fi64++ {
        var v65 uint = 42
        _ = fi64
        _ = v65
    }
    v60 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
