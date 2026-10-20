package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 int }
func genFn0(p0 uint32, p1 uint16, p2 uint) uint64 {
    v5 := []int{42, 0}
    var v6 int16 = 32767
    fmt.Println(p2)
    var v7 int16 = (v6 + (v6 << ((p2) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    var v4 string = "a"
    var v8 uint8 = 128
    v4 = ""
    v8 += v8
    switch 1 {
    case 2147483647:
        v9 := GenS1{F0: 1}
        _ = v9
    case 32768:
        if false {
            v10 := v1[547]
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: 2}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: 4294967295, F1: v4}
                switch 4294967295 {
                case -605:
                    v3 = 1
                case -228:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v16 := []int{fi15}
                        _ = fi15
                        _ = v16
                    }
                case 574:
                    if false {
                        v17 := v1[(fi13 * (-(fi13)))]
                        var v18 bool = ((0.1 < -2.25) && v10)
                        switch fi13 {
                        case 41:
                            v19 := []int{}
                            _ = v19
                        }
                        _ = v17
                        _ = v18
                    } else {
                        var v20 uint32 = 292
                        v3 -= (v3 / ((v3) | 1))
                        _ = v20
                    }
                }
                _ = fi13
                _ = v14
            }
            _ = v10
        } else {
            fmt.Println(v3, v4)
            v4 = v4
        }
    default:
        var v21 bool = (true && true)
        _ = v21
    }
    fmt.Println(v4)
    v8 = (v8 ^ (^(v8)))
    switch ("\n\t" + "a") {
    case "hello":
        if (!true) {
            if (!(true || true)) {
                if true {
                    var v22 int16 = 1
                    var v23 uint8 = 0
                    fmt.Println(v8)
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 14; fi24++ {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        if true {
                            switch (("foo bar" + "\n\t") + v4) {
                            case "\n\t":
                                var v26 int8 = -128
                                _ = v26
                            case "foo bar":
                                var v27 string = "foo bar"
                                _ = v27
                            case "hello":
                                if true {
                                    v28 := GenS1{F0: 83}
                                    _ = v28
                                } else {
                                    var v29 bool = (fi25 >= fi24)
                                    _ = v29
                                }
                            }
                        }
                        _ = fi25
                    }
                    _ = fi24
                }
                if (false || (((true || true) && false) && true)) {
                    v30 := GenS3{F0: 4294967295}
                    if (!(false && (0 <= -38))) {
                        if ((false && false) && true) {
                            v31 := GenS1{F0: 811}
                            if true {
                                fmt.Println(v4)
                                var v32 uint = 660
                                _ = v32
                            }
                            _ = v31
                        }
                        var v33 string = "\n\t"
                        _ = v33
                    }
                    delete(v1, 42)
                    _ = v30
                }
            }
        } else {
            switch v4 {
            case "a":
                v8 = (v8 << ((255) & 63))
            case "\n\t":
                var v34 uint32 = 2
                _ = v34
            case "":
                if true {
                    var v35 uint32 = 65535
                    fmt.Println(v8, v8, v4)
                    _ = v35
                } else {
                    if true {
                        v4 = "hello"
                        var v36 uint64 = 4294967295
                        for fi37 := 0; fi37 < 18; fi37++ {
                            if (1 == -81) {
                                switch fi37 {
                                case 446:
                                    v2[65535] = ((v8 | (v8 / ((128) | 1))) >= (v8 << ((4294967295) & 63)))
                                }
                                v3 -= v3
                                v36 += (v36 >> ((65535) & 63))
                            } else {
                                if (genFn0(2, 1, 11) >= genFn0(2147483647, uint16(v36), 41)) {
                                    for fi38 := 0; fi38 < 7; fi38++ {
                                        var v39 uint16 = 65535
                                        fmt.Println(v39)
                                        _ = fi38
                                        _ = v39
                                    }
                                    var v40 int8 = 16
                                    v36 = v36
                                    _ = v40
                                } else {
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        fmt.Println(v36)
                                        _ = len(v2)
                                        _ = fi41
                                    }
                                    fmt.Println(fi37, fi37)
                                    v4 = "hello"
                                }
                                v42 := GenS1{F0: 0}
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    if true {
                                        if (true && ((1e10 * 0.1) >= float64((0.0 / 1.0)))) {
                                            var v44 uint32 = 888
                                            var v45 int8 = -76
                                            v46 := []int{fi43, (fi37 + fi37)}
                                            switch fi37 {
                                            case 2147483647:
                                                if (false && false) {
                                                    switch (fi37 - -9223372036854775808) {
                                                    case 556:
                                                        switch v4 {
                                                        case "a":
                                                            v46 = append(v46, (fi37 >> ((2) & 63)))
                                                        case "foo bar":
                                                            v1[0] = (!(!false))
                                                        case "":
                                                            v4 = v4
                                                        }
                                                    default:
                                                        fmt.Println(v3, v3, v3)
                                                    }
                                                    if (!((float64((((1.5 + 3.14159) - 1e10) + 0.0)) == ((0.1 / 1.0) * 0.0)) || false)) {
                                                        var v47 uint8 = v8
                                                        v48 := GenS3{F0: 384}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    v49 := GenS4{F0: (v36 & (v36 % ((2) | 1))), F1: ("" + "\n\t")}
                                                    _ = v49
                                                }
                                            }
                                            v50 := GenS5{F0: true}
                                            if len(v46) >= 2 {
                                                v46 = v46[1:len(v46)]
                                            }
                                            var v51 float32 = 1.5
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                            _ = v50
                                            _ = v51
                                        } else {
                                            if false {
                                                if false {
                                                    v4 = ((v4 + v4) + v4)
                                                    var v52 int32 = -2147483648
                                                    _ = v52
                                                } else {
                                                    v53 := GenS2{F0: genFn0(2, 0, 860), F1: v4}
                                                    _ = v53
                                                }
                                                switch "hello" {
                                                case "":
                                                    var v54 uint16 = 255
                                                    _ = v54
                                                case "a":
                                                    var v55 uint64 = (v36 + v36)
                                                    _ = v55
                                                default:
                                                    var v56 int16 = 411
                                                    var v57 uint8 = 229
                                                    v56 -= (^(v56))
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                            var v58 float32 = (1e10 / 1e10)
                                            if (!false) {
                                                fmt.Println(fi37)
                                                v36 = (^(v36))
                                                if (v3 != (v3 * (v3 | int64(v36)))) {
                                                    v59 := []int{(^(fi37))}
                                                    v60 := GenS1{F0: 0}
                                                    v61 := GenS6{F0: v8}
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                            } else {
                                                var v62 int16 = -797
                                                v3 = -915
                                                v63 := []int{fi43, (fi37 & (^(fi37))), -376, (fi37 ^ fi37)}
                                                _ = v62
                                                _ = v63
                                            }
                                            _ = v58
                                        }
                                        switch (fi37 >> ((70) & 63)) {
                                        case 4294967295:
                                            v64 := GenS1{F0: (v36 * (v36 ^ 0))}
                                            _ = v64
                                        default:
                                            fmt.Println(fi43)
                                        }
                                    }
                                    fmt.Println(v36, v36)
                                    switch (fi37 % (((fi37 - fi43)) | 1)) {
                                    case 2147483647:
                                        fmt.Println(fi43, fi43, v3)
                                    case -9223372036854775808:
                                        v3 -= v3
                                    case 32767:
                                        var v65 int8 = -128
                                        v66 := []int{(fi37 - (^(fi43)))}
                                        _ = v65
                                        _ = v66
                                    default:
                                        fmt.Println(v36, v8, fi43)
                                    }
                                    _ = fi43
                                }
                                _ = v42
                            }
                            var v67 float32 = -2.25
                            fmt.Println(v4, v67)
                            _ = fi37
                            _ = v67
                        }
                        _ = v36
                    }
                    v68 := []int{128, -9223372036854775808}
                    _ = v68
                }
            default:
                for fi69 := 0; fi69 < 14; fi69++ {
                    v8 = v8
                    v70 := GenS7{F0: 128, F1: 4294967295}
                    _ = fi69
                    _ = v70
                }
            }
            v71 := []int{128, 127}
            _ = v71
        }
    case "foo bar":
        var v72 string = v4
        _ = v72
    }
    v73 := GenS5{F0: (42 < 0)}
    v74 := GenS8{F0: 922, F1: 272}
    switch v4 {
    case "\n\t":
        var v75 uint = 0
        _ = v75
    case "foo bar":
        for fi76 := 0; fi76 < 1; fi76++ {
            for fi77 := 0; fi77 < 2; fi77++ {
                v4 = ((v4 + (v4 + v4)) + (((v4 + ("\n\t" + v4)) + "foo bar") + (v4 + v4)))
                v78 := GenS9{F0: 313}
                _ = fi77
                _ = v78
            }
            fmt.Println(v8)
            _ = fi76
        }
    case "":
        v3 = v3
    }
    v3 = v3
    v3 -= (-(v3))
    if false {
        var v79 int16 = 42
        fmt.Println(v4)
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v74
    _ = v8
}
