package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int8; F1 uint32; F2 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int64; F1 int64; F2 uint64 }
type GenS12 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{(("a" + "\n\t") + ("\n\t" + "hello")), ""}
    var v3 uint = 577
    v3 = (v3 * (v3 - v3))
    if false {
        var v4 int = 626
        var v5 uint64 = 18446744073709551615
        _ = v4
        _ = v5
    } else {
        var v6 int = 1
        v3 = v3
        if true {
            v3 += (v3 | (v3 | (v3 ^ v3)))
            if (-127 == 9223372036854775807) {
                if false {
                    fmt.Println(v6)
                    var v7 int64 = -1
                    for fi8 := 0; fi8 < 20; fi8++ {
                        v9 := GenS1{F0: 1}
                        v3 = 2
                        if true {
                            var v10 int8 = 0
                            for fi11 := 0; fi11 < 8; fi11++ {
                                var v12 float32 = 1e10
                                var v13 float32 = 0.0
                                v14 := GenS1{F0: -686}
                                _ = fi11
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                            _ = v10
                        }
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                }
            } else {
                if true {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        var v16 uint = v3
                        _ = len(v1)
                        _ = fi15
                        _ = v16
                    }
                    v2 = append(v2, "a")
                    for fi17 := 0; fi17 < 9; fi17++ {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            v6 = 379
                            if (true || true) {
                                var v19 uint8 = 0
                                v3 = (v3 - (v3 - (v3 + 2)))
                                if (!(!false)) {
                                    var v20 uint8 = v19
                                    v21 := GenS2{F0: (!false), F1: 2147483647}
                                    _ = v20
                                    _ = v21
                                }
                                _ = v19
                            } else {
                                if (42 > 42) {
                                    v6 += 42
                                } else {
                                    _ = len(v1)
                                }
                                if false {
                                    var v22 int16 = 42
                                    v22 += 1000
                                    _ = v22
                                } else {
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        fmt.Println(fi18)
                                        v23 := GenS1{F0: 850}
                                        _ = v23
                                    default:
                                        var v24 float64 = 1e10
                                        _ = v24
                                    }
                                }
                                v25 := v1[0]
                                _ = v25
                            }
                            switch "a" {
                            case "foo bar":
                                var v26 string = "foo bar"
                                _ = v26
                            case "hello":
                                fmt.Println(v3, fi17)
                                switch "a" {
                                case "hello":
                                    v1[1] = "foo bar"
                                }
                                var v27 bool = false
                                v28 := GenS3{F0: 577}
                                _ = v27
                                _ = v28
                            case "\n\t":
                                fmt.Println(fi17)
                            default:
                                if true {
                                    v29 := GenS3{F0: 0}
                                    var v30 string = ("foo bar" + ((("foo bar" + ("a" + "\n\t")) + (("foo bar" + "hello") + "")) + ""))
                                    switch v6 {
                                    case -9223372036854775808:
                                        if (1e10 >= 3.14159) {
                                            if false {
                                                var v31 string = v30
                                                delete(v1, 2147483647)
                                                var v32 uint32 = 106
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v31, v3)
                                                case "":
                                                    if (0 < 32767) {
                                                        var v33 int = 127
                                                        _ = v33
                                                    }
                                                case "a":
                                                    switch 1 {
                                                    case 1:
                                                        v2 = append(v2, v30)
                                                    default:
                                                        v30 = (v30 + v31)
                                                    }
                                                }
                                                if ((v32 <= (v32 | v32)) && (!(1e10 >= (0.0 - (float32((3.14159 / 1.0)) / 1e10))))) {
                                                    v31 = ("foo bar" + v31)
                                                    v34 := GenS4{F0: (v32 % (((v32 | 2147483647)) | 1)), F1: v31}
                                                    var v35 float32 = float32((1.5 - ((3.14159 / ((((1e10 + 0.0) + (1.5 - ((3.14159 * 0.0) - 1.5))) * -2.25) - (3.14159 + 1.5))) / -2.25)))
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    if true {
                                                        fmt.Println(fi17, fi17, v6)
                                                        v36 := GenS3{F0: 203}
                                                        _ = v36
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = ((v30 + "hello") + "a")
                                                        }
                                                    }
                                                }
                                                _ = v31
                                                _ = v32
                                            } else {
                                                v3 = v3
                                                var v37 int = -28
                                                _ = v37
                                            }
                                            var v38 uint8 = 255
                                            switch (^(fi17)) {
                                            case 127:
                                                v39 := GenS5{F0: fi17, F1: 2, F2: 1}
                                                _ = v39
                                            case 1:
                                                var v40 float32 = (1.5 + 0.0)
                                                _ = v40
                                            default:
                                                if (0.0 >= 1e10) {
                                                    if ((-1 < 42) && false) {
                                                        var v41 uint32 = 2147483648
                                                        _ = v41
                                                    } else {
                                                        fmt.Println(fi17, fi18)
                                                        if (-128 == 0) {
                                                            v42 := v1[0]
                                                            var v43 float64 = 3.14159
                                                            v38 = 128
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            v6 += (v6 ^ 2147483647)
                                                            v44 := GenS6{F0: (v38 % ((42) | 1)), F1: 707, F2: 2147483648}
                                                            v1[42] = ((("foo bar" + v30) + (v30 + v30)) + v30)
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v38
                                        }
                                    }
                                    v30 = (((("" + (v30 + ("" + v30))) + v30) + v30) + v30)
                                    var v45 bool = (2 != 396)
                                    _ = v29
                                    _ = v30
                                    _ = v45
                                }
                            }
                            _ = fi18
                        }
                        _ = fi17
                    }
                } else {
                    _ = len(v2)
                }
            }
        }
        _ = v6
    }
    var v46 bool = true
    fmt.Println(v3, v46)
    var v47 uint16 = 2
    v48 := GenS7{F0: 1, F1: -1}
    switch 42 {
    case -593:
        switch -1 {
        case 32767:
            fmt.Println(v3, v47, v46)
        case 1:
            fmt.Println(v46)
        case 32768:
            if v46 {
                v46 = (-2.25 == 0.1)
                if (-954 > 9223372036854775807) {
                    if (true && ((!(203 == uint8(v3))) && (252 > 2147483647))) {
                        if true {
                            var v49 bool = v46
                            for fi50 := 0; fi50 < 13; fi50++ {
                                if false {
                                    fmt.Println(v47)
                                } else {
                                    switch (fi50 >> (((^(v3))) & 63)) {
                                    case -806:
                                        v3 = (v3 & 797)
                                    case -2147483648:
                                        var v51 int32 = -395
                                        _ = v51
                                    }
                                    for fi52 := 0; fi52 < 6; fi52++ {
                                        v3 -= (v3 ^ (v3 ^ (v3 - v3)))
                                        var v53 float32 = (0.0 * 1.5)
                                        _ = fi52
                                        _ = v53
                                    }
                                }
                                _ = fi50
                            }
                            fmt.Println(v46, v47, v49)
                            _ = v49
                        } else {
                            var v54 uint8 = 120
                            for fi55 := 0; fi55 < 6; fi55++ {
                                for fi56 := 0; fi56 < 6; fi56++ {
                                    v57 := GenS8{F0: -1, F1: 4294967295, F2: 255}
                                    v47 = (^(v47))
                                    v54 -= (v54 / ((v54) | 1))
                                    _ = fi56
                                    _ = v57
                                }
                                fmt.Println(v3, v3)
                                _ = fi55
                            }
                            _ = v54
                        }
                    } else {
                        var v58 string = "hello"
                        _ = v58
                    }
                }
                v3 = v3
            } else {
                for fi59 := 0; fi59 < 3; fi59++ {
                    v60 := GenS2{F0: v46, F1: 1}
                    v3 += v3
                    switch fi59 {
                    case -549:
                        if (v3 < v3) {
                            v61 := GenS9{F0: (v3 % (((v3 / ((v3) | 1))) | 1))}
                            switch 1 {
                            case -409:
                                var v62 uint16 = (v47 * 894)
                                _ = v62
                            }
                            fmt.Println(fi59)
                            _ = v61
                        }
                    case 4294967295:
                        v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
                    default:
                        v63 := GenS10{F0: 0, F1: (v3 + (v3 >> ((v3) & 63)))}
                        _ = v63
                    }
                    _ = fi59
                    _ = v60
                }
                var v64 float32 = 3.14159
                _ = v64
            }
        }
    case 647:
        if true {
            var v65 int8 = -1
            _ = v65
        }
    case 9223372036854775807:
        v3 += (^(v3))
    default:
        if v46 {
            v66 := GenS11{F0: 0, F1: 913, F2: 2}
            _ = v66
        } else {
            if (1 != 1) {
                v67 := GenS8{F0: -128, F1: 42, F2: 651}
                v68 := GenS4{F0: 1, F1: ""}
                fmt.Println(v47, v3)
                _ = v67
                _ = v68
            } else {
                var v69 uint32 = 2147483648
                v70 := GenS12{F0: 42, F1: ("" == "hello")}
                v3 -= v3
                _ = v69
                _ = v70
            }
            fmt.Println(v3, v3, v3)
        }
    }
    var v71 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v71
}
