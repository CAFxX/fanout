package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int16; F1 uint16; F2 int16 }
type GenS7 struct { F0 int64; F1 string }
func genFn0(p0 int, p1 bool) uint16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 bool = (!(!p1))
        v7 := []int{p0, -9223372036854775808, p0, p0}
        if p1 {
            if len(v7) > 0 {
                v7[0] = fi5
            }
            if (!(42 <= 0)) {
                v8 := GenS1{F0: 42}
                v6 = (!(v6 && true))
                _ = v8
            } else {
                var v9 uint8 = 255
                _ = v9
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            for fi11 := 0; fi11 < 12; fi11++ {
                switch "\n\t" {
                case "a":
                    p1 = false
                case "":
                    switch (^(p0)) {
                    case 42:
                        v12 := GenS2{F0: 1}
                        _ = v12
                    case 9223372036854775807:
                        v13 := GenS3{F0: 0, F1: 4294967295}
                        _ = v13
                    default:
                        v14 := []int{1, (p0 ^ -364), fi11, (fi11 + (fi11 + 9223372036854775807))}
                        _ = v14
                    }
                    var v15 uint = 255
                    p1 = (!(!p1))
                    if p1 {
                        p1 = p1
                    } else {
                        v15 -= v15
                        p0 -= fi11
                    }
                    v16 := []int{(-(p0)), -2147483648}
                    _ = v15
                    _ = v16
                case "foo bar":
                    p0 -= fi11
                    p0 += fi11
                    v17 := GenS1{F0: 72}
                    switch 128 {
                    case 4294967295:
                        var v18 uint32 = 2147483647
                        _ = v18
                    case 796:
                        var v19 uint32 = 2
                        _ = v19
                    case -463:
                        if true {
                            p1 = p1
                            if p1 {
                                if true {
                                    fmt.Println(p1, p0)
                                    if false {
                                        var v20 uint = 255
                                        _ = v20
                                    }
                                } else {
                                    if false {
                                        var v21 int = (fi11 << ((1) & 63))
                                        p1 = p1
                                        _ = v21
                                    } else {
                                        var v22 float64 = 1e10
                                        _ = v22
                                    }
                                    fmt.Println(p0, fi11, p1)
                                    p1 = p1
                                    v23 := []int{(p0 & (p0 << ((4294967295) & 63))), p0}
                                    _ = v23
                                }
                            }
                            v24 := GenS1{F0: -132}
                            _ = v24
                        }
                    default:
                        fmt.Println(p0, fi11, p0)
                    }
                    _ = v17
                default:
                    var v25 string = (("\n\t" + (("" + (("hello" + (("hello" + ((("\n\t" + "\n\t") + ((("a" + "foo bar") + "hello") + ("" + ""))) + ("a" + "a"))) + "hello")) + ((("foo bar" + "a") + ("hello" + "a")) + ("\n\t" + "a")))) + "\n\t")) + (("foo bar" + "foo bar") + ""))
                    _ = v25
                }
                p0 = (p0 & (p0 + p0))
                p0 = fi10
                _ = fi11
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 uint32 = 544
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS2{F0: 42}
                _ = fi28
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS1{F0: 1}
                _ = fi30
                _ = v31
            }
            v32 := []int{fi10, p0, p0}
            _ = v32
        }
        _ = fi10
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        for fi33 := 0; fi33 < 13; fi33++ {
            if false {
                v34 := GenS4{F0: 128, F1: 105, F2: 2147483647}
                v35 := []int{}
                for fi36 := 0; fi36 < 17; fi36++ {
                    if len(v35) >= 2 {
                        v35 = v35[1:len(v35)]
                    }
                    p0 = fi33
                    p1 = (1 != -883)
                    var v37 int = (-(fi36))
                    _ = fi36
                    _ = v37
                }
                _ = v34
                _ = v35
            } else {
                if (-1 != -26) {
                    v38 := []int{p0}
                    if true {
                        switch p0 {
                        case 128:
                            var v39 uint16 = 42
                            _ = v39
                        }
                    } else {
                        p0 = (-(p0))
                        switch p0 {
                        case 9223372036854775807:
                            var v40 bool = true
                            _ = v40
                        case 32768:
                            var v41 uint64 = 66
                            _ = v41
                        default:
                            p0 -= -9223372036854775808
                        }
                        fmt.Println(p0, fi33, fi33)
                    }
                    var v42 uint16 = 930
                    _ = v38
                    _ = v42
                } else {
                    var v43 uint8 = 42
                    _ = v43
                }
                v44 := GenS5{F0: "hello", F1: (1 > 302)}
                _ = v44
            }
            if p1 {
                p0 += 32768
                if p1 {
                    if (p1 || ((3.14159 / ((1e10 - -2.25) - 3.14159)) > 1.5)) {
                        var v45 uint8 = 144
                        if p1 {
                            switch (p0 & 733) {
                            case 127:
                                var v46 uint = 65535
                                _ = v46
                            case 42:
                                switch "\n\t" {
                                case "":
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        switch -2147483648 {
                                        case -878:
                                            v48 := []int{}
                                            _ = v48
                                        case -9223372036854775808:
                                            var v49 string = "foo bar"
                                            _ = v49
                                        case 127:
                                            v50 := []int{p0}
                                            _ = v50
                                        }
                                        p0 += fi47
                                        fmt.Println(p1)
                                        _ = fi47
                                    }
                                case "foo bar":
                                    v45 += v45
                                }
                            }
                        }
                        v51 := GenS2{F0: 1}
                        _ = v45
                        _ = v51
                    }
                    v52 := GenS5{F0: "hello", F1: true}
                    p0 = (fi33 << ((468) & 63))
                    _ = v52
                }
            }
            if p1 {
                var v53 float64 = 3.14159
                var v54 int = (p0 >> ((65535) & 63))
                _ = v53
                _ = v54
            } else {
                for fi55 := 0; fi55 < 19; fi55++ {
                    var v56 float32 = 3.14159
                    _ = fi55
                    _ = v56
                }
            }
            _ = fi33
        }
    case "a":
        p0 -= p0
    default:
        var v57 uint64 = 1
        _ = v57
    }
    p0 -= 4294967295
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 int = 127
    v2 := []uint32{}
    v3 := []uint32{0, 63, 641}
    var v4 uint8 = 191
    v2 = append(v2, 999)
    for fi58 := 0; fi58 < 18; fi58++ {
        if len(v2) > 0 {
            v2[0] = 1
        }
        v59 := GenS6{F0: 0, F1: 125, F2: 901}
        _ = fi58
        _ = v59
    }
    v2 = append(v2, 0)
    v4 = v4
    fmt.Println(v4, v4)
    v1 += v1
    var v60 uint64 = 1
    var v61 uint = 333
    if false {
        v1 -= 458
        var v62 int = v1
        var v63 int16 = 1
        _ = v62
        _ = v63
    } else {
        switch "" {
        case "hello":
            v61 -= 4294967295
        case "a":
            fmt.Println(v4)
        }
    }
    if false {
        for fi64 := 0; fi64 < 4; fi64++ {
            _ = len(v2)
            _ = fi64
        }
        v65 := GenS3{F0: 601, F1: (v1 >> ((v61) & 63))}
        v66 := GenS2{F0: v4}
        _ = v65
        _ = v66
    } else {
        switch (v1 * (v1 / ((int(v61)) | 1))) {
        case -2147483648:
            switch ("hello" + "a") {
            case "a":
                var v67 bool = (v4 > (v4 << ((v61) & 63)))
                _ = v67
            case "":
                v1 -= v1
            case "hello":
                for fi68 := 0; fi68 < 12; fi68++ {
                    for fi69 := 0; fi69 < 18; fi69++ {
                        if (104 >= 933) {
                            switch fi69 {
                            case 32767:
                                if true {
                                    switch "a" {
                                    case "a":
                                        v60 = v60
                                    }
                                }
                            }
                            v70 := GenS5{F0: "hello", F1: false}
                            _ = v70
                        } else {
                            var v71 uint64 = (^(v60))
                            _ = v71
                        }
                        if (-2.25 <= 1.5) {
                            var v72 uint16 = 0
                            v73 := GenS7{F0: 1, F1: "foo bar"}
                            _ = v72
                            _ = v73
                        } else {
                            var v74 uint = (^(v61))
                            _ = v74
                        }
                        v4 += v4
                        _ = fi69
                    }
                    _ = fi68
                }
            }
        case 668:
            v1 -= v1
        case 448:
            for fi75 := 0; fi75 < 14; fi75++ {
                v60 = 0
                switch fi75 {
                case 97:
                    if true {
                        switch fi75 {
                        case 128:
                            v3 = append(v3, 328)
                        default:
                            var v76 float64 = 3.14159
                            _ = v76
                        }
                        fmt.Println(v61)
                        var v77 uint16 = 979
                        _ = v77
                    }
                default:
                    v1 -= -2147483648
                }
                _ = fi75
            }
        default:
            v78 := GenS1{F0: int32(v60)}
            _ = v78
        }
        v1 += (v1 * v1)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v79 uint8 = 2
    v61 = v61
    v61 += v61
    fmt.Println(v60, v61)
    var v80 int32 = -700
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v79
    _ = v80
}
