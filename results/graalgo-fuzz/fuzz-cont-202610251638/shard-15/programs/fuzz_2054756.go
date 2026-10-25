package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int16; F1 bool; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 uint) uint16 {
    p0 -= (^(p0))
    var v7 int16 = 1
    p0 += 761
    fmt.Println(v7)
    v7 = (v7 ^ 1)
    _ = p0
    _ = v7
    return 2
}
func genFn1(p0 int16) string {
    v8 := GenS1{F0: p0, F1: 1, F2: p0}
    v9 := []int{9223372036854775807}
    if (!false) {
        for fi10 := 0; fi10 < 6; fi10++ {
            p0 = (^(p0))
            var v11 int64 = -1
            _ = fi10
            _ = v11
        }
        switch "foo bar" {
        case "foo bar":
            v12 := GenS2{F0: "\n\t"}
            _ = v12
        case "hello":
            v13 := GenS1{F0: 0, F1: 42, F2: (p0 >> ((819) & 63))}
            _ = v13
        case "":
            p0 = p0
        default:
            fmt.Println(p0)
        }
    }
    if ((p0 * p0) > p0) {
        for fi14 := 0; fi14 < 14; fi14++ {
            p0 -= (^(p0))
            _ = fi14
        }
        if ((!((42 > 253) && false)) && false) {
            p0 -= p0
            if len(v9) > 0 {
                v9[0] = 664
            }
            var v15 bool = (-2.25 < 1.5)
            _ = v15
        } else {
            var v16 uint16 = 153
            var v17 int = int(v16)
            _ = v16
            _ = v17
        }
        var v18 uint = 18446744073709551615
        _ = v18
    }
    _ = p0
    _ = v8
    _ = v9
    return ""
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 uint16 = 2
    v4 := make(map[string]string)
    var v5 float32 = ((0.0 + (1.5 * 0.1)) + -2.25)
    var v6 bool = true
    var v19 uint32 = 4294967295
    v20 := GenS3{F0: ""}
    var v21 float32 = v5
    if (!v6) {
        for fi22 := 0; fi22 < 1; fi22++ {
            if (0 < 42) {
                for fi23 := 0; fi23 < 19; fi23++ {
                    switch (fi22 << ((42) & 63)) {
                    case 32767:
                        if ((^(fi22)) > (fi23 << ((985) & 63))) {
                            var v24 uint32 = 0
                            _ = v24
                        }
                    default:
                        var v25 float64 = 0.1
                        _ = v25
                    }
                    _ = fi23
                }
                v19 -= 691
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                v21 = v5
                _ = fi26
            }
            if (((42 == 292) || (127 != 42)) || v6) {
                if (fi22 < 0) {
                    var v27 string = "a"
                    fmt.Println(v6, fi22, fi22)
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        switch ("a" + "") {
                        case "foo bar":
                            v29 := GenS1{F0: -1, F1: 9223372036854775807, F2: -40}
                            _ = v29
                        case "hello":
                            fmt.Println(v5, fi28, fi28)
                        case "":
                            v21 -= (v5 / (0.1 * v21))
                        default:
                            if ((!v6) || (1.5 <= float64(v21))) {
                                var v30 int32 = int32(v3)
                                _ = v30
                            } else {
                                v5 += (v5 / float32(-2.25))
                                var v31 uint16 = (v3 << ((18446744073709551615) & 63))
                                var v32 uint = 102
                                switch ("hello" + "a") {
                                case "foo bar":
                                    var v33 float32 = v21
                                    v3 = (v3 << ((255) & 63))
                                    if v6 {
                                        v31 = 65535
                                        v34 := []int{fi28}
                                        _ = v34
                                    } else {
                                        v35 := []int{-643}
                                        var v36 int16 = 4
                                        switch (fi28 >> ((4294967295) & 63)) {
                                        case -1:
                                            v37 := GenS4{F0: (-(v36)), F1: 18446744073709551615}
                                            _ = v37
                                        case 463:
                                            var v38 uint = v32
                                            _ = v38
                                        }
                                        v33 += (float32(-2.25) + ((float32((3.14159 / -2.25)) / v21) + (1.5 + v33)))
                                        var v39 float64 = 1.5
                                        _ = v35
                                        _ = v36
                                        _ = v39
                                    }
                                    _ = v33
                                case "":
                                    if v6 {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v19 -= (v19 & v19)
                                        case "hello":
                                            var v40 uint = v32
                                            v41 := []int{(fi28 - 32768), -529, 0, int(v19)}
                                            _ = v40
                                            _ = v41
                                        case "a":
                                            if (((uint64(v32) <= 369) && v6) && (!(true || false))) {
                                                var v42 uint = (v32 - (v32 / (((v32 + (v32 & 960))) | 1)))
                                                if v6 {
                                                    switch (fi22 / ((fi28) | 1)) {
                                                    case 1:
                                                        v3 = (v31 << ((255) & 63))
                                                    default:
                                                        v32 -= (v42 >> (((v32 >> (((^(v42))) & 63))) & 63))
                                                    }
                                                }
                                                _ = v42
                                            } else {
                                                var v43 uint8 = 255
                                                if v6 {
                                                    v2["a"] = fi22
                                                } else {
                                                    var v44 float64 = 0.1
                                                    v21 = v5
                                                    v44 = 1.5
                                                    v45 := GenS5{F0: 2, F1: fi22}
                                                    v46 := []int{(fi28 % (((-(fi22))) | 1))}
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 float32 = v5
                                                _ = v43
                                                _ = v47
                                            }
                                        default:
                                            if v6 {
                                                v31 -= v31
                                                v19 = 2
                                            } else {
                                                fmt.Println(v21)
                                                if false {
                                                    v4["foo bar"] = ("a" + "a")
                                                    v3 += v31
                                                }
                                            }
                                        }
                                    } else {
                                        var v48 int = -570
                                        var v49 uint = 241
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    delete(v2, ("hello" + "\n\t"))
                                }
                                v50 := []int{(fi28 / (((fi22 % ((848) | 1))) | 1))}
                                _ = v31
                                _ = v32
                                _ = v50
                            }
                        }
                        _ = fi28
                    }
                    fmt.Println(v6)
                    fmt.Println(fi22)
                }
                if (3.14159 != 1.5) {
                    if v6 {
                        v3 += genFn0(4294967295)
                        var v51 uint64 = 18446744073709551615
                        v3 -= (v3 & genFn0(611))
                        _ = v51
                    } else {
                        v52 := GenS2{F0: "a"}
                        var v53 int8 = -128
                        _ = v52
                        _ = v53
                    }
                    v54 := GenS2{F0: "foo bar"}
                    switch "foo bar" {
                    case "":
                        for fi55 := 0; fi55 < 16; fi55++ {
                            v56 := []int{}
                            _ = fi55
                            _ = v56
                        }
                    case "\n\t":
                        switch (fi22 % ((int(v19)) | 1)) {
                        case 391:
                            for fi57 := 0; fi57 < 11; fi57++ {
                                switch "hello" {
                                case "hello":
                                    v21 = v5
                                case "a":
                                    fmt.Println(v21, v3, v5)
                                default:
                                    v58 := []int{}
                                    _ = v58
                                }
                                var v59 float32 = (v21 * float32(float64((v5 / float32(0.1)))))
                                _ = fi57
                                _ = v59
                            }
                        case 42:
                            fmt.Println(v21)
                        case 224:
                            v60 := GenS3{F0: "hello"}
                            _ = v60
                        default:
                            var v61 int8 = -1
                            _ = v61
                        }
                    case "hello":
                        fmt.Println(v5, v5, v21)
                    default:
                        switch fi22 {
                        case 353:
                            var v62 uint32 = v19
                            _ = v62
                        case 433:
                            v6 = ((v3 >> ((357) & 63)) != (^(v3)))
                        case 0:
                            v3 = v3
                        default:
                            if false {
                                if (!(((v19 + v19) > (v19 << ((265) & 63))) || v6)) {
                                    for fi63 := 0; fi63 < 14; fi63++ {
                                        var v64 int16 = -300
                                        _ = fi63
                                        _ = v64
                                    }
                                    if true {
                                        switch ("a" + "a") {
                                        case "foo bar":
                                            var v65 int64 = -9223372036854775808
                                            _ = v65
                                        }
                                    } else {
                                        if (fi22 <= fi22) {
                                            for fi66 := 0; fi66 < 7; fi66++ {
                                                switch int(v3) {
                                                case 2147483647:
                                                    v67 := GenS4{F0: 864, F1: 961}
                                                    v21 -= (v21 + (v5 * float32(0.1)))
                                                    _ = v67
                                                case -244:
                                                    v68 := v1["hello"]
                                                    _ = v68
                                                default:
                                                    if (!v6) {
                                                        if (v6 && v6) {
                                                            v69 := []int{(fi66 - fi22), fi66, 32767}
                                                            _ = v69
                                                        }
                                                    } else {
                                                        var v70 int16 = 1
                                                        var v71 int64 = -776
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                                _ = fi66
                                            }
                                        }
                                    }
                                    v3 = genFn0(708)
                                }
                                var v72 float64 = (0.1 - 0.0)
                                var v73 bool = (!true)
                                _ = v72
                                _ = v73
                            }
                        }
                    }
                    _ = v54
                } else {
                    v74 := []int{fi22, (fi22 ^ fi22)}
                    _ = v74
                }
            }
            _ = fi22
        }
        var v75 int64 = 3
        if ((v6 || (0 == 1)) || false) {
            if ((591 == (v75 ^ v75)) || (!v6)) {
                for fi76 := 0; fi76 < 18; fi76++ {
                    for fi77 := 0; fi77 < 1; fi77++ {
                        fmt.Println(v19)
                        fmt.Println(fi77)
                        _ = fi77
                    }
                    for fi78 := 0; fi78 < 12; fi78++ {
                        if true {
                            v19 = v19
                            var v79 int8 = 42
                            v80 := GenS1{F0: 42, F1: (v75 & (v75 ^ v75)), F2: -580}
                            _ = v79
                            _ = v80
                        }
                        v19 = 0
                        if (0 != 609) {
                            v21 = (0.0 + ((0.0 / (v5 + v21)) * (0.0 - float32(1.5))))
                            v81 := GenS6{F0: 1, F1: ("hello" <= "hello"), F2: 42}
                            _ = v81
                        } else {
                            fmt.Println(v5)
                            var v82 int8 = -1
                            v82 += 42
                            _ = v82
                        }
                        _ = fi78
                    }
                    for fi83 := 0; fi83 < 17; fi83++ {
                        var v84 int16 = -1
                        v19 = v19
                        _ = fi83
                        _ = v84
                    }
                    _ = fi76
                }
                if v6 {
                    var v85 float32 = v5
                    _ = v85
                } else {
                    if v6 {
                        v21 -= (v5 + v21)
                    }
                    v86 := GenS4{F0: 32767, F1: 1}
                    if (-598 != -521) {
                        v5 -= (float32((0.0 / 1.0)) - v5)
                        v3 = 1
                        v87 := GenS7{F0: ("hello" + "hello")}
                        _ = v87
                    } else {
                        switch "foo bar" {
                        case "":
                            var v88 float32 = (v21 + float32(0.1))
                            _ = v88
                        }
                        var v89 bool = true
                        _ = v89
                    }
                    _ = v86
                }
                switch 1 {
                case 1:
                    delete(v4, ("hello" + "foo bar"))
                case 9223372036854775807:
                    for fi90 := 0; fi90 < 12; fi90++ {
                        v75 += -440
                        _ = fi90
                    }
                case 4294967295:
                    var v91 uint64 = 0
                    _ = v91
                default:
                    v19 = 578
                }
            }
            v92 := GenS5{F0: 42, F1: 127}
            _ = v92
        }
        _ = v75
    } else {
        if false {
            fmt.Println(v5, v3, v6)
        }
    }
    v93 := GenS7{F0: "foo bar"}
    var v94 int64 = 42
    v95 := v4["foo bar"]
    v6 = v6
    if (v6 && true) {
        v96 := GenS8{F0: -1, F1: 3, F2: (v19 + 2147483648)}
        if v6 {
            var v97 bool = v6
            fmt.Println(v6)
            _ = len(v1)
            _ = v97
        } else {
            for fi98 := 0; fi98 < 13; fi98++ {
                for fi99 := 0; fi99 < 17; fi99++ {
                    v100 := []int{(-(fi98))}
                    var v101 int64 = -323
                    _ = fi99
                    _ = v100
                    _ = v101
                }
                fmt.Println(fi98, v19)
                var v102 uint8 = 233
                _ = fi98
                _ = v102
            }
            delete(v2, ("\n\t" + "hello"))
            v95 = "hello"
        }
        v19 -= v19
        _ = v96
    }
    var v103 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v103
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v94
    _ = v95
}
