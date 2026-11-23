package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 uint16; F1 int64 }
type GenS13 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 552
    var v3 string = ""
    v4 := make(map[int]bool)
    v5 := []int{-244, 32767, 127, 4294967295}
    fmt.Println(v2, v3)
    if (-19 < -1) {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int64 = (v2 / ((0) | 1))
            if (false && false) {
                if ((v3 + ("\n\t" + "")) == ("hello" + "")) {
                    v8 := GenS1{F0: 606, F1: -94}
                    v9 := GenS2{F0: 251, F1: v7, F2: 2}
                    _ = v8
                    _ = v9
                } else {
                    if false {
                        var v10 string = "a"
                        v7 = v7
                        _ = v10
                    } else {
                        v11 := GenS2{F0: int16(fi6), F1: 9223372036854775807, F2: 364}
                        v2 -= (v2 | v7)
                        if true {
                            if true {
                                if true {
                                    v12 := v4[fi6]
                                    v3 = "foo bar"
                                    delete(v4, -789)
                                    _ = v12
                                } else {
                                    v3 = "hello"
                                    v13 := GenS3{F0: 591}
                                    delete(v4, fi6)
                                    _ = v13
                                }
                            }
                            var v14 bool = ((true && false) && false)
                            v15 := GenS4{F0: 65535}
                            _ = v14
                            _ = v15
                        } else {
                            fmt.Println(v3)
                        }
                        _ = v11
                    }
                    switch ((("\n\t" + v3) + v3) + ("hello" + v3)) {
                    case "\n\t":
                        v3 = (v3 + v3)
                    }
                }
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v7 = (-(v7))
                _ = fi16
            }
            _ = fi6
            _ = v7
        }
    } else {
        v3 = v3
        v2 = (v2 & v2)
        for fi17 := 0; fi17 < 10; fi17++ {
            if (!(2 < 1)) {
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 uint = 0
                    v19 -= 255
                    if ((fi17 >> ((v19) & 63)) <= fi18) {
                        switch ((v3 + "foo bar") + v3) {
                        case "":
                            fmt.Println(fi18, v19)
                        case "hello":
                            if (1 == -2147483648) {
                                v20 := GenS5{F0: 43, F1: 42, F2: 2147483647}
                                switch v3 {
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "":
                                    var v21 uint = v19
                                    _ = v21
                                case "\n\t":
                                    fmt.Println(fi17, v2, v19)
                                default:
                                    v3 = ""
                                }
                                if ("" >= "foo bar") {
                                    var v22 int16 = -1
                                    _ = v22
                                }
                                switch ("hello" + ((v3 + ("a" + "\n\t")) + (v3 + v3))) {
                                case "foo bar":
                                    v4[(fi18 >> ((v19) & 63))] = ((1.5 < 1e10) && (true && (984 <= 0)))
                                }
                                var v23 bool = (1e10 != float64(((1e10 * -2.25) + (1.5 - ((((1e10 + 1e10) * 1.5) - 0.1) * 1.5)))))
                                v1 = append(v1, 235)
                                switch (fi18 & -2147483648) {
                                case 0:
                                    v23 = (42 == int8(v19))
                                }
                                delete(v4, 0)
                                _ = v20
                                _ = v23
                            }
                        }
                        v24 := GenS6{F0: 2}
                        fmt.Println(v3, fi17)
                        fmt.Println(v2, v19)
                        _ = v24
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi17
        }
    }
    v2 = v2
    _ = len(v5)
    if true {
        if true {
            if ((4294967295 < 246) && (true || true)) {
                if true {
                    v3 = ("hello" + v3)
                    v2 = -1
                } else {
                    v25 := GenS7{F0: -294, F1: -32768}
                    v26 := GenS8{F0: 599}
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    v3 = v3
                    _ = fi27
                }
                switch "a" {
                case "a":
                    if false {
                        v28 := v4[127]
                        if (65535 >= 255) {
                            if v28 {
                                v2 = v2
                                v29 := GenS8{F0: -106}
                                v3 = v3
                                _ = v29
                            }
                        } else {
                            v30 := GenS7{F0: 2147483647, F1: -448}
                            _ = v30
                        }
                        fmt.Println(v28, v2)
                        _ = v28
                    } else {
                        v31 := GenS9{F0: 72}
                        v3 = v3
                        _ = v31
                    }
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = 128
                    }
                }
            } else {
                for fi32 := 0; fi32 < 18; fi32++ {
                    fmt.Println(v3, fi32)
                    v33 := GenS4{F0: 2}
                    _ = fi32
                    _ = v33
                }
                v34 := v4[-867]
                v2 = (^(v2))
                _ = v34
            }
            var v35 bool = (91 == 42)
            _ = v35
        } else {
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 uint8 = 255
                _ = fi36
                _ = v37
            }
        }
        var v38 uint = 961
        v39 := GenS10{F0: v3, F1: "hello"}
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            switch v3 {
            case "":
                fmt.Println(fi40)
            default:
                var v41 int8 = 86
                _ = v41
            }
            var v42 string = v3
            if (true || false) {
                v43 := GenS6{F0: 145}
                var v44 float64 = 1.5
                v3 = v3
                _ = v43
                _ = v44
            } else {
                var v45 uint64 = 4294967295
                var v46 int64 = v2
                for fi47 := 0; fi47 < 19; fi47++ {
                    _ = len(v4)
                    _ = fi47
                }
                _ = v45
                _ = v46
            }
            _ = fi40
            _ = v42
        }
    }
    if ((1 < 628) || (2 != 2147483647)) {
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 int = 2147483647
            _ = fi48
            _ = v49
        }
        var v50 string = v3
        for fi51 := 0; fi51 < 11; fi51++ {
            switch ("foo bar" + (v3 + v50)) {
            case "":
                switch fi51 {
                case -2147483648:
                    v50 = ("hello" + v50)
                }
            default:
                v52 := GenS11{F0: 2, F1: 1}
                _ = v52
            }
            for fi53 := 0; fi53 < 5; fi53++ {
                fmt.Println(fi53, v50)
                v3 = (("\n\t" + "\n\t") + "a")
                v3 = v3
                _ = fi53
            }
            _ = fi51
        }
        _ = v50
    }
    fmt.Println(v3, v3)
    v54 := GenS12{F0: 65535, F1: -9223372036854775808}
    var v55 uint32 = 65535
    var v56 float64 = 1e10
    var v57 uint16 = 1
    if true {
        var v58 string = v3
        v59 := GenS1{F0: 981, F1: int8(v2)}
        fmt.Println(v57)
        _ = v58
        _ = v59
    } else {
        for fi60 := 0; fi60 < 20; fi60++ {
            v3 = "foo bar"
            v2 = 42
            v61 := GenS9{F0: 2}
            _ = fi60
            _ = v61
        }
        var v62 uint32 = v55
        var v63 float64 = (0.0 - v56)
        _ = v62
        _ = v63
    }
    v64 := GenS13{F0: 42, F1: (v57 >> ((211) & 63))}
    var v65 float32 = 1e10
    var v66 bool = (-2.25 > 0.1)
    var v67 uint8 = 1
    delete(v4, 42)
    var v68 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
