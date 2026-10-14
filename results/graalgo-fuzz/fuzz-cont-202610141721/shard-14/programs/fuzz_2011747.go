package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint16 }
func genFn0(p0 uint16, p1 int64) int16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint16 = (p0 >> ((42) & 63))
        _ = fi4
        _ = v5
    }
    var v6 int8 = 0
    var v7 int16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (v7 * v7)
}
func genFn1(p0 uint64, p1 bool) int32 {
    p0 = (p0 >> ((956) & 63))
    var v8 int16 = genFn0(280, 1)
    switch 127 {
    case -1:
        v9 := GenS1{F0: false, F1: 42}
        _ = v9
    case 0:
        v8 = (v8 * (v8 << ((2) & 63)))
    }
    p0 = p0
    v10 := []int{128, 32767}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return int32(p0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := []int{}
    v11 := GenS1{F0: ((((1 > 9223372036854775807) || (-128 != 64)) || (genFn1(396, false) <= -2147483648)) || true), F1: 4294967295}
    switch "\n\t" {
    case "hello":
        var v12 string = "foo bar"
        _ = v12
    case "\n\t":
        var v13 int64 = 9223372036854775807
        _ = v13
    case "":
        var v14 uint8 = 252
        _ = v14
    }
    switch "" {
    case "":
        var v15 uint = 2
        _ = v15
    default:
        v16 := GenS2{F0: -666, F1: 0}
        _ = v16
    }
    if true {
        var v17 string = (("\n\t" + "\n\t") + ("foo bar" + "hello"))
        fmt.Println(v17)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint = 1
            _ = fi18
            _ = v19
        }
        var v20 float32 = (float32((0.0 - ((3.14159 * -2.25) - 3.14159))) - 0.0)
        var v21 int64 = 160
        _ = v20
        _ = v21
    }
    if (2 >= 238) {
        var v22 int = 32767
        fmt.Println(v22)
        _ = v22
    } else {
        var v23 int64 = -425
        _ = v23
    }
    switch "hello" {
    case "hello":
        if true {
            switch -9223372036854775808 {
            case 127:
                switch 32767 {
                case 9223372036854775807:
                    var v24 uint32 = 2147483647
                    _ = v24
                case 498:
                    v25 := GenS1{F0: (-9223372036854775808 == 423), F1: 4294967295}
                    _ = v25
                case -581:
                    switch -441 {
                    case 372:
                        if false {
                            var v26 uint = 4294967295
                            for fi27 := 0; fi27 < 1; fi27++ {
                                v1[(fi27 % (((fi27 + fi27)) | 1))] = 708
                                var v28 int32 = -449
                                if false {
                                    v28 -= 0
                                } else {
                                    var v29 int8 = 117
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            if false {
                                                var v30 uint16 = 1
                                                v30 += (v30 | (v30 / (((^(v30))) | 1)))
                                                var v31 uint8 = 17
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                    case "foo bar":
                                        v32 := GenS1{F0: true, F1: 2147483648}
                                        _ = v32
                                    default:
                                        v28 -= v28
                                    }
                                    fmt.Println(v29)
                                    _ = v29
                                }
                                _ = fi27
                                _ = v28
                            }
                            var v33 int = 127
                            _ = v26
                            _ = v33
                        }
                    default:
                        v34 := GenS2{F0: 0, F1: 4294967295}
                        _ = v34
                    }
                default:
                    var v35 uint32 = 2
                    _ = v35
                }
            }
            var v36 uint8 = 42
            switch 42 {
            case 42:
                var v37 int8 = 97
                _ = v37
            case 4294967295:
                var v38 uint32 = 65535
                _ = v38
            case 166:
                v36 = (v36 % ((v36) | 1))
            default:
                v39 := GenS1{F0: true, F1: 2}
                _ = v39
            }
            _ = v36
        } else {
            switch "hello" {
            case "a":
                if (!(1 < 127)) {
                    switch "hello" {
                    case "foo bar":
                        for fi40 := 0; fi40 < 8; fi40++ {
                            v41 := GenS2{F0: 1000, F1: 2}
                            v42 := GenS3{F0: 2, F1: 1}
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    case "":
                        for fi43 := 0; fi43 < 13; fi43++ {
                            if (!true) {
                                fmt.Println(fi43)
                                switch (fi43 << ((4294967295) & 63)) {
                                case -659:
                                    var v44 bool = ((!((fi43 << ((1) & 63)) <= 32768)) && (42 <= (fi43 << ((255) & 63))))
                                    _ = v44
                                case 165:
                                    v45 := GenS4{F0: int16(fi43)}
                                    _ = v45
                                case 4294967295:
                                    _ = len(v3)
                                default:
                                    var v46 int64 = 42
                                    _ = v46
                                }
                                var v47 int = fi43
                                _ = v47
                            }
                            _ = fi43
                        }
                    default:
                        var v48 uint8 = 2
                        _ = v48
                    }
                }
            case "":
                _ = len(v3)
            case "\n\t":
                if (!(0 < 0)) {
                    v49 := GenS4{F0: 32767}
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 float64 = (1.5 / (0.1 / 1e10))
                        v51 -= 1.5
                        switch "foo bar" {
                        case "\n\t":
                            switch ("a" + "") {
                            case "hello":
                                var v52 int32 = -2147483648
                                _ = v52
                            case "a":
                                v51 = float64(-2.25)
                            case "":
                                switch ("\n\t" + ("hello" + ("foo bar" + ("foo bar" + "\n\t")))) {
                                case "\n\t":
                                    v51 -= v51
                                }
                            default:
                                var v53 float32 = 1.5
                                _ = v53
                            }
                        default:
                            var v54 int16 = 0
                            _ = v54
                        }
                        _ = fi50
                        _ = v51
                    }
                }
            default:
                var v55 uint = 2
                _ = v55
            }
        }
    }
    switch -2147483648 {
    case -9223372036854775808:
        if (!(-1 <= genFn0(32767, 0))) {
            switch -78 {
            case 1:
                if (0 >= -9223372036854775808) {
                    var v56 float32 = (1e10 * -2.25)
                    v56 = v56
                    v56 = (v56 + v56)
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 5; fi57++ {
                        var v58 int32 = -2147483648
                        for fi59 := 0; fi59 < 20; fi59++ {
                            v60 := GenS4{F0: 32767}
                            v58 = (v58 / ((genFn1(2, true)) | 1))
                            v58 += genFn1(765, (4294967295 >= 1))
                            _ = fi59
                            _ = v60
                        }
                        v58 -= (v58 | -214)
                        _ = fi57
                        _ = v58
                    }
                }
            case 42:
                if (true && false) {
                    var v61 float64 = (-2.25 - 3.14159)
                    _ = v61
                }
            }
            var v62 uint8 = 255
            switch 470 {
            case 32768:
                if false {
                    var v63 int32 = 0
                    v64 := GenS4{F0: genFn0(65535, -1)}
                    var v65 int16 = 0
                    _ = v63
                    _ = v64
                    _ = v65
                } else {
                    fmt.Println(v62)
                    switch "\n\t" {
                    case "foo bar":
                        v66 := GenS1{F0: true, F1: 65535}
                        _ = v66
                    }
                    fmt.Println(v62)
                }
            }
            _ = v62
        } else {
            var v67 int32 = 2147483647
            _ = len(v1)
            var v68 uint32 = 2
            _ = v67
            _ = v68
        }
    case 674:
        for fi69 := 0; fi69 < 9; fi69++ {
            var v70 float64 = 1.5
            switch fi69 {
            case -2147483648:
                v70 += 3.14159
            case 32767:
                v71 := GenS5{F0: "", F1: 2}
                _ = v71
            case -82:
                var v72 uint = 18446744073709551615
                _ = v72
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi69
            _ = v70
        }
    }
    for fi73 := 0; fi73 < 5; fi73++ {
        v74 := v1[fi73]
        _ = fi73
        _ = v74
    }
    if (801 < 4294967295) {
        var v75 bool = false
        _ = v75
    }
    for fi76 := 0; fi76 < 6; fi76++ {
        var v77 int32 = 42
        switch "a" {
        case "\n\t":
            var v78 bool = true
            _ = v78
        case "a":
            var v79 int16 = -626
            _ = v79
        }
        _ = fi76
        _ = v77
    }
    var v80 uint32 = 854
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v80
}
