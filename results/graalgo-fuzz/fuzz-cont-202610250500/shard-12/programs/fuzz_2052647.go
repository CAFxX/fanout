package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int16; F1 int8; F2 uint8 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 bool; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    v3 := []string{(("foo bar" + "foo bar") + "foo bar"), "hello", "\n\t"}
    var v4 int = 9223372036854775807
    var v5 int8 = 0
    v4 = (^(v4))
    v4 += 1
    v3 = append(v3, "foo bar")
    v5 = (v5 / (((^(v5))) | 1))
    if true {
        switch "hello" {
        case "foo bar":
            for fi6 := 0; fi6 < 8; fi6++ {
                if (127 == 2) {
                    if false {
                        v1 = 392
                    } else {
                        if (v5 != int8(v1)) {
                            v3 = append(v3, (("" + "hello") + ""))
                            fmt.Println(v4)
                            if true {
                                for fi7 := 0; fi7 < 14; fi7++ {
                                    if (!(false && (v5 != 127))) {
                                        v8 := GenS1{F0: 0, F1: 74, F2: 65535}
                                        if false {
                                            _ = len(v2)
                                            var v9 int = (fi7 >> ((409) & 63))
                                            _ = v9
                                        }
                                        var v10 uint8 = 178
                                        var v11 int64 = -1
                                        var v12 bool = (true && (false && true))
                                        _ = v10
                                        _ = v11
                                        _ = v12
                                        _ = v8
                                    } else {
                                        v2[(fi6 * fi6)] = 1
                                        delete(v2, v4)
                                        v13 := GenS2{F0: v1, F1: (!false), F2: 0}
                                        if (true || false) {
                                            v5 = v5
                                        }
                                        _ = v13
                                    }
                                    v5 += 1
                                    _ = fi7
                                }
                            } else {
                                if true {
                                    v14 := GenS3{F0: true, F1: 255, F2: 2}
                                    _ = v14
                                } else {
                                    if (false || true) {
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        v5 = v5
                                        var v15 string = ("foo bar" + ("a" + ("foo bar" + ("a" + ((("\n\t" + "a") + "hello") + ("\n\t" + "foo bar"))))))
                                        _ = v15
                                    }
                                }
                                delete(v2, 2147483647)
                                v4 = (fi6 & fi6)
                            }
                        } else {
                            v16 := GenS4{F0: 255, F1: (!false), F2: 972}
                            _ = v16
                        }
                    }
                    v17 := GenS1{F0: (fi6 | (fi6 | fi6)), F1: 255, F2: 132}
                    _ = v17
                } else {
                    v18 := GenS5{F0: 42, F1: 255, F2: 18446744073709551615}
                    for fi19 := 0; fi19 < 13; fi19++ {
                        switch 1 {
                        case 120:
                            switch "\n\t" {
                            case "a":
                                switch ("" + "foo bar") {
                                case "a":
                                    var v20 int64 = 9223372036854775807
                                    _ = v20
                                case "":
                                    _ = len(v3)
                                case "foo bar":
                                    v21 := v2[fi6]
                                    _ = v21
                                }
                            }
                            v5 += 100
                        default:
                            switch (v4 & (fi6 + (fi19 / ((-9223372036854775808) | 1)))) {
                            case 128:
                                var v22 float64 = (1e10 + 3.14159)
                                _ = v22
                            case 0:
                                var v23 uint64 = 1
                                _ = v23
                            case 32768:
                                v24 := GenS6{F0: 127, F1: 42}
                                _ = v24
                            }
                        }
                        v5 -= -128
                        _ = fi19
                    }
                    v25 := GenS7{F0: v1, F1: (v5 % ((90) | 1))}
                    _ = v18
                    _ = v25
                }
                _ = fi6
            }
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v26 uint = 0
        if (((-(v4)) > -2147483648) || false) {
            v4 += int(v5)
        } else {
            v5 -= (v5 + -91)
        }
        _ = v26
    }
    fmt.Println(v5)
    if ((4294967295 != 1) || false) {
        if false {
            switch "\n\t" {
            case "":
                for fi27 := 0; fi27 < 3; fi27++ {
                    _ = len(v2)
                    v4 = fi27
                    var v28 int64 = 0
                    _ = fi27
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 uint32 = 440
                    v32 := GenS8{F0: -825, F1: int8(v1), F2: 0}
                    fmt.Println(v4, fi29)
                    v33 := GenS2{F0: v1, F1: ((!true) && false), F2: 65535}
                    if true {
                        v34 := GenS9{F0: (-(v1)), F1: 42}
                        _ = v34
                    } else {
                        v5 = 52
                        switch (^(fi29)) {
                        case 4294967295:
                            switch "a" {
                            case "foo bar":
                                v1 = v1
                            case "a":
                                v4 = -1
                            }
                        default:
                            v5 = -128
                        }
                        v35 := GenS5{F0: 214, F1: 48, F2: 0}
                        _ = v35
                    }
                    v1 -= v1
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                var v36 float64 = 0.1
                var v37 int8 = (v5 >> ((57) & 63))
                _ = fi29
                _ = v36
                _ = v37
            }
            v38 := GenS10{F0: false, F1: 434, F2: 1}
            _ = v38
        }
        var v39 uint8 = 1
        fmt.Println(v5, v39)
        _ = v39
    } else {
        if (2 > 4294967295) {
            _ = len(v2)
            v4 -= (^(v4))
            v40 := GenS6{F0: 127, F1: 0}
            _ = v40
        } else {
            var v41 float64 = 0.0
            var v42 float32 = -2.25
            fmt.Println(v41, v41)
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 8; fi43++ {
            v44 := GenS11{F0: 687}
            var v45 string = ("" + "a")
            _ = fi43
            _ = v44
            _ = v45
        }
        var v46 uint32 = 2147483647
        _ = v46
    }
    var v47 string = "a"
    v5 += (-(v5))
    switch int(v5) {
    case -2147483648:
        v48 := GenS8{F0: 32767, F1: v5, F2: 247}
        _ = v48
    case -50:
        v5 -= (v5 % (((v5 - (v5 % ((42) | 1)))) | 1))
    }
    _ = len(v3)
    var v49 float32 = 0.1
    switch "a" {
    case "":
        if (255 == 1) {
            for fi50 := 0; fi50 < 19; fi50++ {
                var v51 float32 = 0.1
                fmt.Println(fi50, v51, v51)
                fmt.Println(fi50, v5, v4)
                _ = fi50
                _ = v51
            }
        } else {
            for fi52 := 0; fi52 < 11; fi52++ {
                var v53 uint = 0
                for fi54 := 0; fi54 < 3; fi54++ {
                    var v55 string = ("" + "foo bar")
                    var v56 uint8 = 128
                    var v57 uint = v53
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v57
                }
                _ = fi52
                _ = v53
            }
            var v58 float32 = v49
            var v59 int16 = 1000
            _ = v58
            _ = v59
        }
    case "foo bar":
        var v60 int = (v4 / (((v4 >> ((1) & 63))) | 1))
        _ = v60
    case "\n\t":
        var v61 uint8 = 116
        _ = v61
    }
    fmt.Println(v47)
    for fi62 := 0; fi62 < 7; fi62++ {
        if false {
            var v63 uint16 = 65535
            _ = v63
        } else {
            v64 := GenS10{F0: (!true), F1: 4294967295, F2: -1}
            var v65 int8 = v5
            _ = v64
            _ = v65
        }
        _ = fi62
    }
    delete(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
}
