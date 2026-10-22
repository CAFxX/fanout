package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 string; F1 bool; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 uint32 }
func genFn0(p0 int16) uint32 {
    v7 := GenS1{F0: -942, F1: 2147483647, F2: 42}
    if false {
        fmt.Println(p0)
    } else {
        p0 = (p0 >> ((255) & 63))
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint = 0
            v10 := []int{(-(fi8))}
            var v11 uint64 = 2
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 int8 = 42
    v13 := []int{898, 4294967295}
    _ = p0
    _ = v12
    _ = v13
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 uint8) int {
    var v14 int64 = 42
    fmt.Println(v14, p1)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v14
    return -1
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 42
    v3 := []string{"hello"}
    v4 := make(map[string]int)
    var v5 int = 32768
    var v6 int16 = 560
    var v15 int32 = 2147483647
    v16 := GenS2{F0: (v6 % ((30) | 1)), F1: -1}
    switch v5 {
    case 813:
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS3{F0: 371, F1: 4294967295, F2: -1}
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint32 = 460
                var v21 bool = (0 <= int16(v20))
                _ = fi19
                _ = v20
                _ = v21
            }
            v2 -= -128
            _ = fi17
            _ = v18
        }
    case -2147483648:
        v1 = (v1 ^ (v1 | v1))
    default:
        if (("hello" + "\n\t") == "") {
            var v22 int8 = v2
            var v23 int64 = 9223372036854775807
            _ = v22
            _ = v23
        }
    }
    v1 = v1
    var v24 string = "a"
    var v25 int8 = (^(v2))
    if false {
        if false {
            var v26 int = (v5 | (v5 ^ genFn1(986, 128)))
            for fi27 := 0; fi27 < 2; fi27++ {
                v24 = v24
                for fi28 := 0; fi28 < 5; fi28++ {
                    v5 -= 562
                    v1 -= 993
                    switch fi28 {
                    case -266:
                        var v29 float32 = ((((1.5 * 1e10) - (float32(((0.1 * 0.0) + 0.0)) * float32((1e10 / (1e10 / (-2.25 * 3.14159)))))) - 0.0) * 1e10)
                        var v30 int64 = 42
                        _ = v29
                        _ = v30
                    case 32767:
                        v31 := GenS4{F0: 42, F1: 42}
                        _ = v31
                    }
                    if false {
                        v32 := GenS3{F0: 17, F1: 730, F2: fi27}
                        _ = v32
                    } else {
                        if true {
                            var v33 string = ("" + v24)
                            _ = v33
                        }
                        v5 -= genFn1(4294967295, 2)
                        var v34 int8 = (v2 | int8(fi27))
                        _ = v34
                    }
                    var v35 int8 = (v2 & (-(v25)))
                    _ = fi28
                    _ = v35
                }
                var v36 int16 = -244
                _ = fi27
                _ = v36
            }
            _ = len(v3)
            _ = v26
        } else {
            for fi37 := 0; fi37 < 14; fi37++ {
                v38 := GenS3{F0: 2, F1: 665, F2: v5}
                if false {
                    var v39 int = (v5 - genFn1(0, 42))
                    var v40 int8 = (v2 + (v25 << ((255) & 63)))
                    v41 := GenS2{F0: -393, F1: (^(v6))}
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    switch genFn1(1, 0) {
                    case -332:
                        v42 := GenS5{F0: (v24 + v24), F1: (!true), F2: 225}
                        _ = v42
                    case 621:
                        switch 32768 {
                        case -9223372036854775808:
                            if (!true) {
                                var v43 float32 = 1.5
                                v3 = append(v3, "")
                                _ = v43
                            }
                        case 0:
                            if false {
                                if (!true) {
                                    for fi44 := 0; fi44 < 1; fi44++ {
                                        fmt.Println(v6)
                                        var v45 uint64 = 4294967295
                                        fmt.Println(v2)
                                        _ = fi44
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 1; fi46++ {
                                        v47 := GenS2{F0: (v6 << ((2) & 63)), F1: (-(v6))}
                                        switch (v24 + "hello") {
                                        case "hello":
                                            v48 := GenS3{F0: 0, F1: 1, F2: fi46}
                                            _ = v48
                                        case "":
                                            v49 := v4[""]
                                            _ = v49
                                        }
                                        v15 += (v15 & v15)
                                        _ = fi46
                                        _ = v47
                                    }
                                } else {
                                    v50 := v4[v24]
                                    fmt.Println(fi37)
                                    _ = v50
                                }
                                if ((1.5 / ((-2.25 + (1e10 - ((1.5 * 1e10) + -2.25))) + (0.1 + 0.0))) > float64(float32((float64((0.0 - float32(1.5))) * 1.5)))) {
                                    for fi51 := 0; fi51 < 7; fi51++ {
                                        var v52 int8 = v25
                                        _ = fi51
                                        _ = v52
                                    }
                                }
                                var v53 int32 = 0
                                _ = v53
                            }
                        }
                    case 0:
                        var v54 uint = v1
                        _ = v54
                    default:
                        var v55 int32 = 1
                        _ = v55
                    }
                }
                v56 := v4["a"]
                _ = fi37
                _ = v38
                _ = v56
            }
        }
        if false {
            var v57 uint = (v1 ^ v1)
            var v58 int32 = v15
            if false {
                v59 := v4["\n\t"]
                var v60 uint16 = 255
                v57 += (v1 / (((v1 % ((v1) | 1))) | 1))
                _ = v59
                _ = v60
            } else {
                if false {
                    var v61 float32 = 1.5
                    switch (((v24 + v24) + "foo bar") + "\n\t") {
                    case "foo bar":
                        switch v24 {
                        case "a":
                            var v62 uint = 65535
                            _ = v62
                        case "":
                            if true {
                                if (-32768 >= v6) {
                                    for fi63 := 0; fi63 < 13; fi63++ {
                                        _ = len(v3)
                                        v64 := GenS1{F0: 1, F1: (v5 ^ 42), F2: 42}
                                        _ = fi63
                                        _ = v64
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v65 := GenS5{F0: "foo bar", F1: (1 > 821), F2: 222}
                                    _ = v65
                                }
                            }
                        case "foo bar":
                            switch int(v58) {
                            case 42:
                                fmt.Println(v25, v24, v2)
                            case -2147483648:
                                fmt.Println(v58, v5)
                            case 1:
                                switch v5 {
                                case 9223372036854775807:
                                    v15 = (-(v58))
                                default:
                                    v15 = v58
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                    case "a":
                        var v66 uint16 = 32767
                        _ = v66
                    default:
                        v5 += genFn1(1, 128)
                    }
                    var v67 uint16 = 32767
                    _ = v61
                    _ = v67
                }
                v15 -= v58
                switch v24 {
                case "hello":
                    _ = len(v4)
                case "a":
                    v68 := GenS6{F0: (v2 | (v25 & v2))}
                    _ = v68
                case "":
                    var v69 bool = true
                    _ = v69
                }
            }
            _ = v57
            _ = v58
        }
    } else {
        var v70 int = genFn1(4294967295, 68)
        _ = v70
    }
    if true {
        v5 = -2147483648
    } else {
        switch v5 {
        case -9223372036854775808:
            if true {
                var v71 uint = 255
                fmt.Println(v25)
                _ = len(v3)
                _ = v71
            } else {
                v6 = v6
                _ = len(v3)
            }
        }
        var v72 string = ""
        _ = v72
    }
    var v73 uint8 = 127
    v74 := GenS7{F0: 963, F1: 529}
    var v75 uint16 = 32767
    v76 := GenS4{F0: (v73 ^ (v73 % (((^(v73))) | 1))), F1: (^(v1))}
    v15 -= (v15 + (v15 << ((50) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
