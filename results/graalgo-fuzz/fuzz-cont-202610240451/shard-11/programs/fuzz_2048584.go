package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint16; F1 uint; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16 }
func genFn0(p0 int16) uint16 {
    if (1 <= 1) {
        for fi6 := 0; fi6 < 20; fi6++ {
            if false {
                if false {
                    if (!((true || false) || false)) {
                        v7 := GenS1{F0: ((p0 - (-(p0))) == (-(p0))), F1: 61, F2: 0}
                        fmt.Println(fi6, p0)
                        p0 += p0
                        _ = v7
                    }
                    var v8 string = "a"
                    _ = v8
                }
                if true {
                    fmt.Println(fi6)
                    for fi9 := 0; fi9 < 9; fi9++ {
                        var v10 float64 = 3.14159
                        switch "a" {
                        case "a":
                            var v11 float32 = (1.5 + 1.5)
                            _ = v11
                        case "hello":
                            var v12 float32 = float32(v10)
                            _ = v12
                        }
                        _ = fi9
                        _ = v10
                    }
                    switch (fi6 * (fi6 % ((1) | 1))) {
                    case 4294967295:
                        p0 = -747
                    default:
                        if (((!false) || true) && (!true)) {
                            if false {
                                if false {
                                    v13 := []int{(fi6 >> ((16) & 63)), fi6}
                                    _ = len(v13)
                                    _ = v13
                                }
                                switch "hello" {
                                case "\n\t":
                                    for fi14 := 0; fi14 < 20; fi14++ {
                                        fmt.Println(fi6, fi14)
                                        p0 = 0
                                        v15 := GenS1{F0: false, F1: -1, F2: 816}
                                        _ = fi14
                                        _ = v15
                                    }
                                case "a":
                                    var v16 bool = false
                                    _ = v16
                                case "foo bar":
                                    p0 = -322
                                default:
                                    switch (-(fi6)) {
                                    case -723:
                                        switch (fi6 << ((4294967295) & 63)) {
                                        case 841:
                                            for fi17 := 0; fi17 < 15; fi17++ {
                                                switch "hello" {
                                                case "":
                                                    switch ("" + "") {
                                                    case "a":
                                                        fmt.Println(fi6, p0, fi17)
                                                    case "\n\t":
                                                        p0 = (p0 + (^(p0)))
                                                    case "":
                                                        fmt.Println(p0, fi17, fi17)
                                                    default:
                                                        v18 := []int{fi6}
                                                        _ = v18
                                                    }
                                                default:
                                                    var v19 string = ""
                                                    _ = v19
                                                }
                                                if false {
                                                    v20 := []int{}
                                                    _ = v20
                                                } else {
                                                    v21 := GenS2{F0: 27, F1: 111, F2: ""}
                                                    _ = v21
                                                }
                                                if (!(0 <= 1)) {
                                                    p0 -= (p0 + -32768)
                                                }
                                                _ = fi17
                                            }
                                        case -671:
                                            var v22 uint32 = 2
                                            _ = v22
                                        }
                                    case 4294967295:
                                        switch ("a" + ("a" + "")) {
                                        case "":
                                            for fi23 := 0; fi23 < 4; fi23++ {
                                                fmt.Println(fi23, p0, p0)
                                                if ("\n\t" == "foo bar") {
                                                    var v24 uint8 = 255
                                                    _ = v24
                                                }
                                                v25 := []int{fi23, (fi6 << ((779) & 63)), (fi6 ^ fi6), fi6}
                                                _ = fi23
                                                _ = v25
                                            }
                                        case "foo bar":
                                            if true {
                                                for fi26 := 0; fi26 < 12; fi26++ {
                                                    p0 -= p0
                                                    _ = fi26
                                                }
                                            }
                                        case "\n\t":
                                            var v27 int32 = -1
                                            _ = v27
                                        default:
                                            v28 := []int{fi6, fi6, fi6}
                                            _ = v28
                                        }
                                    }
                                }
                                fmt.Println(fi6, fi6)
                            } else {
                                if false {
                                    fmt.Println(p0, p0)
                                    p0 += (-(p0))
                                }
                                v29 := GenS3{F0: ("foo bar" + "")}
                                switch fi6 {
                                case -900:
                                    p0 = p0
                                default:
                                    var v30 uint16 = 42
                                    _ = v30
                                }
                                _ = v29
                            }
                            v31 := GenS3{F0: ("a" + "foo bar")}
                            _ = v31
                        } else {
                            var v32 float64 = 3.14159
                            _ = v32
                        }
                    }
                } else {
                    if true {
                        p0 = 1000
                        v33 := GenS2{F0: 2147483647, F1: 42, F2: ("hello" + "hello")}
                        _ = v33
                    }
                    p0 = (p0 + (p0 + p0))
                    v34 := []int{fi6, fi6}
                    _ = v34
                }
            } else {
                v35 := GenS4{F0: 26}
                v36 := GenS5{F0: (^(fi6)), F1: 1}
                _ = v35
                _ = v36
            }
            _ = fi6
        }
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        p0 -= (p0 | 42)
        _ = fi37
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 1
    v3 := []int{}
    var v4 int8 = 127
    v5 := []string{("\n\t" + "hello")}
    fmt.Println(v2, v2)
    fmt.Println(v4)
    switch 32767 {
    case -9223372036854775808:
        v38 := GenS6{F0: true}
        _ = v38
    case 32768:
        var v39 float32 = 1e10
        _ = v39
    case 0:
        v1[""] = 127
    default:
        if (("a" + "") <= "foo bar") {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 bool = (!true)
                var v42 float32 = (1e10 - (0.1 + -2.25))
                v4 = (v4 | (v4 & v4))
                _ = fi40
                _ = v41
                _ = v42
            }
            fmt.Println(v2, v4)
        }
    }
    v4 += 100
    if true {
        if false {
            var v43 string = ""
            _ = v43
        } else {
            fmt.Println(v2, v2)
        }
        if (502 >= 57) {
            v4 = 42
            v44 := GenS6{F0: false}
            switch "foo bar" {
            case "a":
                v45 := GenS2{F0: 4294967295, F1: 1, F2: "foo bar"}
                _ = v45
            case "foo bar":
                for fi46 := 0; fi46 < 20; fi46++ {
                    for fi47 := 0; fi47 < 3; fi47++ {
                        v4 += v4
                        if false {
                            v2 = 128
                            v4 = -18
                            v1["a"] = fi46
                            var v48 uint32 = 2147483647
                            _ = v48
                        }
                        v2 = (^(v2))
                        _ = fi47
                    }
                    var v49 int16 = 1000
                    _ = fi46
                    _ = v49
                }
            }
            _ = v44
        }
    } else {
        var v50 float64 = ((0.0 + (1.5 * float64((1.5 * (1.5 - 1.5))))) / 1.0)
        _ = v50
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        if (!(!false)) {
            var v52 uint16 = 255
            v53 := GenS4{F0: (v2 % (((v2 - v2)) | 1))}
            v4 = (v4 >> ((0) & 63))
            _ = v52
            _ = v53
        } else {
            fmt.Println(v4)
            fmt.Println(fi51)
            v4 -= 42
        }
        fmt.Println(fi51, v2)
        var v54 bool = ((620 > 465) || true)
        _ = fi51
        _ = v54
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        if false {
            switch "a" {
            case "":
                v56 := GenS7{F0: v4, F1: (v2 << ((811) & 63)), F2: 1}
                _ = v56
            case "foo bar":
                var v57 uint = 65535
                _ = v57
            case "a":
                fmt.Println(v4, fi55)
            default:
                var v58 uint8 = (^(v2))
                _ = v58
            }
            v59 := GenS8{F0: 533, F1: 4294967295, F2: -1}
            for fi60 := 0; fi60 < 12; fi60++ {
                v4 = v4
                v4 -= v4
                if (((!(614 > genFn0(1000))) || true) || false) {
                    if false {
                        v61 := GenS8{F0: 255, F1: 689, F2: -1}
                        _ = v61
                    } else {
                        v62 := GenS9{F0: 0}
                        v5 = append(v5, (("" + ("\n\t" + ("hello" + ("hello" + (("foo bar" + ("hello" + "a")) + ("a" + ((("\n\t" + "\n\t") + "\n\t") + ""))))))) + "\n\t"))
                        _ = v62
                    }
                    switch fi60 {
                    case 0:
                        v63 := GenS10{F0: int16(fi60)}
                        _ = v63
                    case 4294967295:
                        if false {
                            var v64 int64 = 30
                            _ = v64
                        }
                    }
                    v65 := GenS5{F0: fi60, F1: 1}
                    _ = v65
                }
                _ = fi60
            }
            _ = v59
        }
        _ = len(v1)
        v2 = 127
        _ = fi55
    }
    var v66 uint = 390
    var v67 uint = (v66 ^ v66)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
}
