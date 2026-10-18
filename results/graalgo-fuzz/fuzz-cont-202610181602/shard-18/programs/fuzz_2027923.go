package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 uint16) uint {
    fmt.Println(p0)
    if true {
        v5 := GenS1{F0: -126}
        _ = v5
    } else {
        switch 2147483647 {
        case -2147483648:
            if (!((false || true) || (65535 >= 246))) {
                if ((0 <= 1000) && (-1 != 0)) {
                    v6 := GenS1{F0: 100}
                    p0 = (p0 & (^(p0)))
                    _ = v6
                }
            }
        case 127:
            switch "" {
            case "foo bar":
                switch ("" + ("" + "a")) {
                case "hello":
                    switch -744 {
                    case 73:
                        var v7 int8 = -1
                        _ = v7
                    case -920:
                        for fi8 := 0; fi8 < 18; fi8++ {
                            switch "a" {
                            case "":
                                switch ("hello" + ("foo bar" + ("hello" + "\n\t"))) {
                                case "":
                                    p0 = p0
                                case "foo bar":
                                    v9 := []int{0, fi8}
                                    _ = v9
                                default:
                                    p0 = 1
                                }
                            case "foo bar":
                                var v10 int8 = 1
                                _ = v10
                            }
                            switch 128 {
                            case 283:
                                var v11 uint8 = 2
                                _ = v11
                            default:
                                var v12 uint32 = 1
                                _ = v12
                            }
                            _ = fi8
                        }
                    }
                case "foo bar":
                    p0 = 755
                }
            }
        case -184:
            switch -608 {
            case 128:
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS2{F0: -1, F1: (2147483647 > 1)}
                    var v15 bool = false
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            default:
                v16 := GenS3{F0: ""}
                _ = v16
            }
        }
        switch "a" {
        case "a":
            for fi17 := 0; fi17 < 9; fi17++ {
                p0 = p0
                if true {
                    v18 := []int{(^(fi17)), fi17}
                    if false {
                        var v19 bool = (2 > 612)
                        if false {
                            if v19 {
                                var v20 int = fi17
                                v21 := GenS4{F0: 65535, F1: (-1 <= 1)}
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v23 := GenS4{F0: 952, F1: (fi17 >= fi17)}
                                    _ = fi22
                                    _ = v23
                                }
                                _ = v20
                                _ = v21
                            }
                            var v24 uint8 = 238
                            _ = v24
                        }
                        v25 := GenS2{F0: 104, F1: true}
                        _ = v19
                        _ = v25
                    }
                    _ = v18
                }
                p0 -= 32767
                _ = fi17
            }
        case "":
            if false {
                if true {
                    v26 := GenS5{F0: 632, F1: ("\n\t" + "hello")}
                    _ = v26
                } else {
                    var v27 float64 = 1e10
                    _ = v27
                }
            }
        default:
            for fi28 := 0; fi28 < 9; fi28++ {
                switch (fi28 ^ (fi28 | (fi28 | fi28))) {
                case 0:
                    var v29 uint16 = p0
                    _ = v29
                case 9223372036854775807:
                    p0 = (p0 << ((0) & 63))
                }
                fmt.Println(fi28, fi28)
                for fi30 := 0; fi30 < 9; fi30++ {
                    v31 := GenS6{F0: -2147483648}
                    _ = fi30
                    _ = v31
                }
                _ = fi28
            }
        }
    }
    _ = p0
    return 233
}
func genFn1(p0 uint16) bool {
    v32 := []int{}
    var v33 int16 = -885
    for fi34 := 0; fi34 < 4; fi34++ {
        p0 = 302
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
        _ = fi34
    }
    var v35 string = "foo bar"
    _ = p0
    _ = v32
    _ = v33
    _ = v35
    return (!false)
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{937, 42, 2147483647, 4294967295}
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    v1 += v1
    var v36 uint16 = 0
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v36)
    v37 := GenS7{F0: 127, F1: (^(v36)), F2: 2147483647}
    var v38 float32 = (3.14159 + -2.25)
    if true {
        v39 := GenS5{F0: 220, F1: "a"}
        if true {
            v36 = 42
            v4 = v4
        } else {
            _ = len(v3)
        }
        if true {
            var v40 uint32 = 65535
            _ = v40
        } else {
            for fi41 := 0; fi41 < 18; fi41++ {
                for fi42 := 0; fi42 < 15; fi42++ {
                    if (1.5 == (float64(0.1) / 1.0)) {
                        fmt.Println(v4, fi41, v36)
                        if false {
                            switch "hello" {
                            case "a":
                                v43 := GenS6{F0: -2147483648}
                                _ = v43
                            default:
                                switch ("\n\t" + "\n\t") {
                                case "foo bar":
                                    if ((0 < 0) && false) {
                                        v36 = (^(v36))
                                        var v44 float32 = 0.0
                                        if ((fi41 % (((fi42 ^ (fi42 % ((fi41) | 1)))) | 1)) == (fi42 * 1)) {
                                            if ((fi42 != (-(fi41))) && (v4 == (-(v4)))) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v2 = append(v2, 65535)
                                                }
                                                v44 += (float32(3.14159) + v44)
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                                v45 := GenS6{F0: 42}
                                                v4 = v4
                                                if false {
                                                    var v46 int64 = (v4 >> ((65535) & 63))
                                                    var v47 bool = (!true)
                                                    switch (("" + "a") + "") {
                                                    case "hello":
                                                        v48 := v3["\n\t"]
                                                        _ = v48
                                                    }
                                                    v4 -= v4
                                                    fmt.Println(v38, v4, v4)
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    var v49 uint = 1
                                                    fmt.Println(v44, fi42, fi41)
                                                    _ = v49
                                                }
                                                _ = v45
                                            } else {
                                                _ = len(v2)
                                                var v50 uint8 = 191
                                                if false {
                                                    v51 := GenS8{F0: "foo bar", F1: (fi41 << ((563) & 63))}
                                                    if false {
                                                        var v52 uint32 = 2147483648
                                                        fmt.Println(fi42, v52, v52)
                                                        var v53 float64 = float64((float32(1e10) * v44))
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    var v54 uint = 417
                                                    _ = v51
                                                    _ = v54
                                                }
                                                _ = v50
                                            }
                                            v55 := GenS9{F0: (1 < 9223372036854775807), F1: 24, F2: 1}
                                            _ = v55
                                        }
                                        var v56 int8 = -128
                                        _ = v44
                                        _ = v56
                                    } else {
                                        v1 -= v1
                                        v36 -= v36
                                    }
                                case "":
                                    delete(v3, "a")
                                }
                            }
                            v57 := GenS4{F0: 906, F1: false}
                            _ = v57
                        }
                    }
                    v58 := GenS8{F0: "hello", F1: 127}
                    var v59 int8 = -112
                    _ = fi42
                    _ = v58
                    _ = v59
                }
                _ = fi41
            }
        }
        _ = v39
    }
    var v60 uint = 4294967295
    for fi61 := 0; fi61 < 13; fi61++ {
        var v62 uint = (^(v60))
        _ = fi61
        _ = v62
    }
    v60 += v60
    var v63 string = "a"
    for fi64 := 0; fi64 < 14; fi64++ {
        v60 -= v60
        _ = fi64
    }
    delete(v3, "")
    _ = len(v3)
    switch v63 {
    case "\n\t":
        if true {
            v1 = (v1 * v1)
            v60 = (v60 & (v60 >> (((v60 % ((v60) | 1))) & 63)))
            var v65 uint8 = 2
            _ = v65
        }
    case "":
        _ = len(v2)
    case "a":
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    if true {
        for fi66 := 0; fi66 < 1; fi66++ {
            v67 := GenS5{F0: -9223372036854775808, F1: (v63 + v63)}
            v68 := GenS4{F0: 0, F1: ("a" < v63)}
            var v69 uint8 = 233
            _ = fi66
            _ = v67
            _ = v68
            _ = v69
        }
        if len(v2) > 0 {
            v2[0] = 542
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v60
    _ = v63
}
