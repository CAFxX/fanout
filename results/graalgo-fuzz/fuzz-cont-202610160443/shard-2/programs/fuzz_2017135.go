package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 string; F1 string; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int16 }

func main() {
    var v1 uint = 290
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    if ((-1 < v3) || (0 != 2)) {
        var v4 int = 42
        _ = v4
    } else {
        v5 := GenS1{F0: 42}
        v6 := GenS2{F0: (v1 + (v1 ^ (v1 << ((0) & 63))))}
        fmt.Println(v3, v3)
        _ = v5
        _ = v6
    }
    var v7 int16 = -988
    switch 2147483647 {
    case 240:
        var v8 float64 = (float64(0.0) * (1.5 + (0.1 + (0.0 + 1.5))))
        _ = v8
    case 128:
        var v9 int16 = 32767
        _ = v9
    case 0:
        var v10 float64 = (0.0 * float64(3.14159))
        _ = v10
    }
    v1 -= v1
    if (!false) {
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{128}
            if false {
                v13 := GenS2{F0: 1}
                fmt.Println(v3)
                _ = v13
            } else {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS1{F0: 1}
                    _ = fi14
                    _ = v15
                }
                var v16 bool = ("" != ((("foo bar" + ("" + "foo bar")) + "foo bar") + "\n\t"))
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        v17 := []int{976, -1, 881, -967}
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := v2[fi18]
        v20 := GenS2{F0: v1}
        v21 := []int{fi18, (-(fi18)), (^(fi18))}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 int32 = (v3 / ((v3) | 1))
    switch ("\n\t" + "\n\t") {
    case "hello":
        switch 2147483647 {
        case 2147483647:
            if true {
                if (!(((float32(0.1) == 1e10) && true) && ((3.14159 + (1e10 * (1.5 * 0.1))) != (0.0 - float64(((3.14159 / 1e10) + (-2.25 * 1e10))))))) {
                    v22 = (v22 / ((v22) | 1))
                }
                var v23 int16 = (v7 & v7)
                _ = v23
            } else {
                var v24 string = ("" + "\n\t")
                fmt.Println(v22, v1)
                if (true && false) {
                    v25 := GenS1{F0: 32767}
                    if false {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            v27 := GenS3{F0: int64(v22), F1: "foo bar"}
                            _ = fi26
                            _ = v27
                        }
                    }
                    _ = v25
                }
                _ = v24
            }
        }
    }
    if true {
        var v28 float64 = -2.25
        v2[2147483647] = 2147483647
        _ = v28
    } else {
        var v29 bool = (100 > -128)
        for fi30 := 0; fi30 < 9; fi30++ {
            _ = len(v2)
            if v29 {
                switch "foo bar" {
                case "":
                    fmt.Println(v3)
                default:
                    v31 := []int{fi30}
                    _ = v31
                }
            } else {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v3 += (v22 << (((v1 >> (((^(v1))) & 63))) & 63))
                    v3 = v22
                    if false {
                        delete(v2, 111)
                        switch (fi32 + fi32) {
                        case 2147483647:
                            var v33 uint32 = 2
                            var v34 bool = v29
                            if true {
                                var v35 uint = (^(v1))
                                _ = v35
                            } else {
                                if v34 {
                                    var v36 uint8 = 0
                                    if (v34 && (262 < 9223372036854775807)) {
                                        var v37 float64 = 1.5
                                        v38 := GenS1{F0: 2}
                                        v39 := []int{(-(fi32)), (fi30 << ((1) & 63)), fi32}
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                var v40 float32 = 0.1
                                var v41 int = 128
                                _ = v40
                                _ = v41
                            }
                            _ = v33
                            _ = v34
                        default:
                            var v42 int16 = v7
                            _ = v42
                        }
                    } else {
                        fmt.Println(fi30, fi30, v22)
                        var v43 bool = false
                        _ = v43
                    }
                    _ = fi32
                }
                var v44 uint = 255
                _ = v44
            }
            fmt.Println(v3, v1)
            _ = fi30
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            fmt.Println(v29)
            v46 := GenS4{F0: "\n\t", F1: (("hello" + "a") + ("" + "foo bar")), F2: (v3 / (((v3 + 1)) | 1))}
            _ = fi45
            _ = v46
        }
        _ = v29
    }
    v7 = v7
    var v47 uint32 = 811
    if false {
        v1 = (v1 | uint(v47))
        switch -964 {
        case 0:
            switch -1 {
            case 401:
                var v48 uint = v1
                _ = v48
            case 565:
                var v49 int8 = -52
                _ = v49
            case 1:
                v50 := GenS5{F0: v47}
                _ = v50
            default:
                var v51 uint32 = 65535
                _ = v51
            }
        case 127:
            v2[-9223372036854775808] = 4294967295
        case -9223372036854775808:
            v3 -= v22
        }
    } else {
        if true {
            fmt.Println(v1)
        }
    }
    var v52 uint16 = 390
    var v53 uint16 = 42
    for fi54 := 0; fi54 < 1; fi54++ {
        fmt.Println(v52)
        v55 := GenS1{F0: v53}
        for fi56 := 0; fi56 < 2; fi56++ {
            if true {
                if false {
                    var v57 uint16 = 358
                    v53 = (v57 >> ((v1) & 63))
                    _ = v57
                } else {
                    v22 = (v3 | v22)
                }
                var v58 string = ""
                _ = v58
            } else {
                v59 := GenS3{F0: -1, F1: "foo bar"}
                v22 += (v22 - -1)
                var v60 uint16 = (v53 % ((2) | 1))
                if (!(0.0 > ((1.5 - 0.0) * 3.14159))) {
                    v2[128] = v47
                }
                v61 := GenS6{F0: (fi56 >> ((uint(v53)) & 63)), F1: true}
                _ = v59
                _ = v60
                _ = v61
            }
            v52 -= 523
            if (621 != 42) {
                v62 := []int{int(v1)}
                v53 += (^(v53))
                v7 += (v7 | v7)
                _ = v62
            } else {
                delete(v2, (fi56 - fi54))
                v63 := []int{}
                fmt.Println(fi54, fi56, v7)
                var v64 float32 = 0.1
                _ = v63
                _ = v64
            }
            _ = fi56
        }
        _ = fi54
        _ = v55
    }
    if (true && true) {
        v65 := []int{4294967295, -2147483648}
        if (true && true) {
            v53 = v53
            switch ((("a" + "") + "") + ("foo bar" + ("" + ""))) {
            case "foo bar":
                fmt.Println(v3, v1)
            case "a":
                v65 = append(v65, 42)
            case "hello":
                if len(v65) >= 2 {
                    v65 = v65[1:len(v65)]
                }
            default:
                for fi66 := 0; fi66 < 11; fi66++ {
                    switch 9223372036854775807 {
                    case 1:
                        for fi67 := 0; fi67 < 1; fi67++ {
                            fmt.Println(fi67, v3, v22)
                            var v68 uint16 = v52
                            var v69 string = ""
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                    case 4294967295:
                        v52 -= 675
                    case -2147483648:
                        switch fi66 {
                        case 2147483647:
                            for fi70 := 0; fi70 < 19; fi70++ {
                                v2[-81] = v47
                                v47 = 65535
                                switch int(v3) {
                                case 567:
                                    var v71 int = (fi70 >> ((v1) & 63))
                                    _ = v71
                                case 2147483647:
                                    var v72 bool = true
                                    _ = v72
                                }
                                var v73 int16 = v7
                                _ = fi70
                                _ = v73
                            }
                        case -107:
                            v3 = (-(v22))
                        case -1:
                            var v74 int = (fi66 & int(v1))
                            _ = v74
                        }
                    default:
                        if false {
                            fmt.Println(v52)
                            switch "hello" {
                            case "foo bar":
                                v75 := GenS7{F0: 1000}
                                _ = v75
                            case "":
                                var v76 uint32 = 611
                                _ = v76
                            default:
                                var v77 float64 = ((((1e10 - (1.5 / 1.5)) / 1.5) / 1.5) / 1.0)
                                _ = v77
                            }
                        }
                    }
                    v52 -= 0
                    v52 = 1
                    _ = fi66
                }
            }
        }
        _ = v65
    }
    for fi78 := 0; fi78 < 2; fi78++ {
        if (false && false) {
            var v79 string = ("\n\t" + "a")
            var v80 int8 = 100
            _ = v79
            _ = v80
        } else {
            var v81 int32 = v22
            var v82 string = ""
            _ = v81
            _ = v82
        }
        if false {
            fmt.Println(v47, v22)
        }
        _ = fi78
    }
    v83 := []int{120, -9223372036854775808, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v83:", len(v83), v83)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v47
    _ = v52
    _ = v53
    _ = v7
    _ = v83
}
