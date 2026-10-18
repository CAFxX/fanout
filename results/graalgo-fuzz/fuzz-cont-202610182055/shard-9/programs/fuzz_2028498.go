package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint16 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = (fi7 >> ((2) & 63))
        if false {
            var v9 uint16 = 255
            var v10 uint8 = 2
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            switch (("a" + ("\n\t" + "")) + "\n\t") {
            case "":
                var v12 bool = false
                _ = v12
            case "foo bar":
                v13 := GenS1{F0: 1}
                _ = v13
            }
            _ = fi11
        }
        _ = fi7
        _ = v8
    }
    var v14 int64 = -135
    if true {
        v15 := []int{int(v14), 686, 9223372036854775807}
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint32 = 42
            var v18 uint8 = 127
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 42}
        _ = v15
        _ = v19
    }
    fmt.Println(v14)
    v20 := []int{128, -91}
    _ = v14
    _ = v20
    return (v14 * v14)
}
func genFn1() int16 {
    var v21 string = "\n\t"
    for fi22 := 0; fi22 < 7; fi22++ {
        v23 := []int{(fi22 / ((fi22) | 1)), fi22}
        var v24 string = ("hello" + ("" + (v21 + ("" + v21))))
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 4294967295}
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            if true {
                v28 := []int{fi26, (fi27 + -2147483648), (fi26 / ((-387) | 1))}
                v21 = (((v21 + v21) + "hello") + ((v21 + v21) + ""))
                var v29 int64 = 544
                _ = v28
                _ = v29
            }
            switch fi26 {
            case 128:
                var v30 int8 = 42
                _ = v30
            default:
                var v31 int16 = 132
                _ = v31
            }
            _ = fi27
        }
        switch fi26 {
        case -732:
            var v32 float64 = 0.1
            _ = v32
        default:
            for fi33 := 0; fi33 < 20; fi33++ {
                v34 := GenS2{F0: 2147483647, F1: 465, F2: ""}
                v21 = v21
                _ = fi33
                _ = v34
            }
        }
        _ = fi26
    }
    if true {
        var v35 float32 = 1e10
        v36 := GenS1{F0: 2}
        _ = v35
        _ = v36
    }
    _ = v21
    _ = v25
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("\n\t" + "\n\t")}
    var v3 uint = 255
    var v4 float32 = (1.5 / 1e10)
    var v5 int = 32767
    v6 := []uint32{uint32(v3), 2147483648}
    v4 = (v4 / v4)
    v4 -= float32(0.0)
    if false {
        var v37 int32 = 506
        var v38 int = (^(v5))
        _ = v37
        _ = v38
    } else {
        fmt.Println(v4, v3)
    }
    v3 = v3
    switch "\n\t" {
    case "":
        v39 := GenS3{F0: 76}
        _ = v39
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "hello":
        v40 := GenS3{F0: 188}
        _ = v40
    default:
        if ((v5 + (v5 ^ v5)) != (v5 * v5)) {
            if false {
                v5 += (v5 / (((v5 ^ (v5 % ((0) | 1)))) | 1))
                var v41 bool = ((false && true) && false)
                var v42 int16 = -521
                _ = v41
                _ = v42
            }
            var v43 uint64 = 2
            v6 = append(v6, 133)
            _ = v43
        } else {
            fmt.Println(v5, v4)
            switch 128 {
            case 798:
                var v44 string = ""
                _ = v44
            case 193:
                fmt.Println(v3)
            case 42:
                var v45 float64 = 0.0
                _ = v45
            default:
                if false {
                    if (!(!(65535 == 65535))) {
                        v46 := GenS3{F0: 32767}
                        v47 := GenS4{F0: 4294967295, F1: 42, F2: 0}
                        _ = v46
                        _ = v47
                    }
                    var v48 string = "foo bar"
                    if (42 != 127) {
                        switch v48 {
                        case "a":
                            fmt.Println(v4, v4, v4)
                        }
                        for fi49 := 0; fi49 < 8; fi49++ {
                            for fi50 := 0; fi50 < 11; fi50++ {
                                v5 = fi50
                                _ = fi50
                            }
                            v1[v48] = ("\n\t" <= ("\n\t" + ""))
                            _ = fi49
                        }
                        for fi51 := 0; fi51 < 16; fi51++ {
                            for fi52 := 0; fi52 < 5; fi52++ {
                                switch (^(fi52)) {
                                case -1:
                                    switch v48 {
                                    case "foo bar":
                                        if false {
                                            v48 = v48
                                        }
                                    default:
                                        fmt.Println(v3, v3)
                                    }
                                case 840:
                                    if (!true) {
                                        var v53 int16 = -417
                                        v54 := v1[v48]
                                        _ = v53
                                        _ = v54
                                    }
                                case 32767:
                                    if false {
                                        v3 -= (v3 >> ((v3) & 63))
                                        v4 += ((-2.25 / (1.5 / v4)) / v4)
                                    }
                                }
                                fmt.Println(v3)
                                v55 := GenS3{F0: 0}
                                _ = fi52
                                _ = v55
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi51
                        }
                    } else {
                        var v56 string = "\n\t"
                        v57 := GenS5{F0: 4294967295}
                        _ = v56
                        _ = v57
                    }
                    _ = v48
                }
            }
        }
    }
    if ("\n\t" <= "hello") {
        for fi58 := 0; fi58 < 6; fi58++ {
            if (!((0.0 - 3.14159) >= 1.5)) {
                v4 = 0.0
                if (-1 <= -2147483648) {
                    var v59 int64 = 1
                    _ = v59
                } else {
                    switch "hello" {
                    case "hello":
                        for fi60 := 0; fi60 < 13; fi60++ {
                            switch "a" {
                            case "foo bar":
                                v4 = v4
                                var v61 int64 = genFn0()
                                _ = v61
                            default:
                                switch "hello" {
                                case "":
                                    switch (fi58 ^ 0) {
                                    case 127:
                                        var v62 string = "foo bar"
                                        _ = v62
                                    default:
                                        var v63 uint32 = 2147483647
                                        _ = v63
                                    }
                                case "foo bar":
                                    fmt.Println(v5, fi58)
                                case "\n\t":
                                    v1[""] = false
                                default:
                                    if true {
                                        v3 = (v3 << ((v3) & 63))
                                        if false {
                                            v4 -= v4
                                            if (false && true) {
                                                var v64 int8 = 100
                                                _ = v64
                                            }
                                            if false {
                                                fmt.Println(fi58)
                                            }
                                        } else {
                                            v4 = 0.1
                                            v65 := v1["a"]
                                            _ = v65
                                        }
                                        var v66 int64 = 1
                                        v67 := v1[(("\n\t" + "a") + "a")]
                                        _ = v66
                                        _ = v67
                                    } else {
                                        var v68 float32 = 3.14159
                                        _ = v68
                                    }
                                    fmt.Println(fi58, fi60, v4)
                                    v69 := GenS6{F0: (128 <= 0), F1: 324}
                                    _ = v69
                                }
                            }
                            if (false && (!((!(1000 >= 846)) && (!(!false))))) {
                                var v70 int32 = 721
                                v71 := GenS3{F0: 65535}
                                _ = len(v1)
                                _ = v70
                                _ = v71
                            } else {
                                v72 := GenS4{F0: 4294967295, F1: 953, F2: genFn1()}
                                _ = v72
                            }
                            v5 -= (fi60 % ((fi58) | 1))
                            v1[("hello" + "hello")] = ((("\n\t" + "a") + "a") > "hello")
                            _ = len(v1)
                            _ = fi60
                        }
                    default:
                        switch "" {
                        case "\n\t":
                            v73 := GenS7{F0: 65535}
                            _ = v73
                        case "":
                            v74 := GenS7{F0: 1}
                            _ = v74
                        case "hello":
                            v4 = v4
                        default:
                            v75 := GenS4{F0: 871, F1: 835, F2: 1000}
                            _ = v75
                        }
                    }
                }
                var v76 uint = 42
                _ = v76
            } else {
                var v77 float64 = 0.1
                _ = v77
            }
            if true {
                var v78 int32 = -1
                switch ("" + "a") {
                case "":
                    if (false || true) {
                        v5 -= 4294967295
                    } else {
                        v3 -= (^(v3))
                        delete(v1, ("a" + "foo bar"))
                        if true {
                            var v79 int16 = -768
                            var v80 string = ""
                            _ = v79
                            _ = v80
                        }
                    }
                case "foo bar":
                    if (2 <= 4294967295) {
                        delete(v1, "foo bar")
                        v3 += (v3 - 2)
                        if false {
                            v5 = (fi58 | fi58)
                            for fi81 := 0; fi81 < 6; fi81++ {
                                var v82 uint64 = uint64(v3)
                                delete(v1, "a")
                                var v83 bool = true
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                            v3 += (v3 - v3)
                        }
                    } else {
                        v4 = 3.14159
                        if false {
                            var v84 bool = (!true)
                            _ = v84
                        } else {
                            var v85 int = (v5 / ((v5) | 1))
                            v85 = (v5 | (^(fi58)))
                            _ = v85
                        }
                        var v86 uint8 = 242
                        _ = v86
                    }
                case "a":
                    if true {
                        v87 := GenS3{F0: 159}
                        var v88 int8 = -20
                        for fi89 := 0; fi89 < 20; fi89++ {
                            fmt.Println(v5, fi89, fi89)
                            var v90 uint = v3
                            var v91 uint16 = 255
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                        _ = v87
                        _ = v88
                    }
                default:
                    delete(v1, ("hello" + "a"))
                }
                fmt.Println(v3, v78, v5)
                _ = v78
            } else {
                var v92 string = "a"
                _ = v92
            }
            var v93 int64 = 296
            _ = fi58
            _ = v93
        }
    }
    var v94 uint64 = 4294967295
    v5 += v5
    if false {
        v5 -= (v5 - v5)
        var v95 uint = 210
        var v96 string = "a"
        _ = v95
        _ = v96
    } else {
        v1["foo bar"] = false
        v3 = 0
    }
    v94 += 4294967295
    var v97 int = (v5 / (((v5 >> (((v3 | 461)) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v97
}
