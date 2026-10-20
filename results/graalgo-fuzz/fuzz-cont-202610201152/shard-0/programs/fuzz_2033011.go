package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 uint16, p1 uint16) uint8 {
    p1 = 2
    var v7 uint16 = p1
    v8 := []int{127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []int{32767, -1, 0}
    var v2 uint16 = 1
    var v3 int = -285
    v4 := make(map[uint32]int)
    var v5 uint = 18446744073709551615
    v6 := []int{(v3 & int(v5))}
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v9 bool = true
        v2 = (v2 >> (((^(v5))) & 63))
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if true {
            v4[1] = fi10
            v11 := GenS1{F0: fi10, F1: (v3 * (fi10 ^ v3))}
            _ = v11
        }
        v2 -= 42
        _ = fi10
    }
    v4[65535] = v3
    switch "a" {
    case "\n\t":
        v1 = append(v1, v3)
    case "hello":
        _ = len(v1)
    default:
        v12 := GenS2{F0: false, F1: 1}
        _ = v12
    }
    if (-32768 > 1) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := v4[318]
            for fi15 := 0; fi15 < 13; fi15++ {
                v5 = v5
                _ = fi15
            }
            if (true || ((!(502 >= 281)) || ((2 == 383) || ((32767 <= (v2 >> (((v5 * v5)) & 63))) && (!true))))) {
                switch (fi13 >> (((v5 >> ((v5) & 63))) & 63)) {
                case 140:
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if (!false) {
                            fmt.Println(v14, v14, v2)
                            if false {
                                if true {
                                    v14 -= v14
                                    switch "a" {
                                    case "hello":
                                        switch "" {
                                        case "foo bar":
                                            if true {
                                                delete(v4, 640)
                                                if (!true) {
                                                    v5 -= 684
                                                    var v17 bool = false
                                                    _ = v17
                                                }
                                                var v18 uint16 = (v2 << (((v5 % (((v5 ^ (^(v5)))) | 1))) & 63))
                                                if ((-1 < -1) || (false && ((((-32768 < -1) && true) && (2147483647 != (fi13 / ((-747) | 1)))) || (!true)))) {
                                                    fmt.Println(v2, v18)
                                                    var v19 int8 = 1
                                                    _ = v19
                                                } else {
                                                    var v20 int32 = -1
                                                    v4[146] = fi16
                                                    var v21 int64 = 42
                                                    _ = v20
                                                    _ = v21
                                                }
                                                _ = v18
                                            }
                                        case "hello":
                                            v22 := GenS1{F0: 813, F1: fi13}
                                            _ = v22
                                        case "\n\t":
                                            var v23 int16 = -463
                                            _ = v23
                                        default:
                                            var v24 uint64 = 0
                                            _ = v24
                                        }
                                    case "\n\t":
                                        v5 = 255
                                    case "foo bar":
                                        var v25 string = "\n\t"
                                        if (!false) {
                                            if (false && ((0 <= (v2 - (v2 - 65535))) && (((1.5 == (0.1 - float64(0.0))) || true) || false))) {
                                                switch v25 {
                                                case "a":
                                                    fmt.Println(v5, fi16)
                                                case "hello":
                                                    fmt.Println(v14, v14)
                                                case "foo bar":
                                                    v2 -= v2
                                                }
                                            } else {
                                                switch (fi13 << (((v5 % ((v5) | 1))) & 63)) {
                                                case -766:
                                                    var v26 uint8 = 242
                                                    _ = v26
                                                }
                                            }
                                        } else {
                                            _ = len(v1)
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        if false {
                                            v27 := GenS3{F0: int32(v2), F1: 123}
                                            _ = v27
                                        } else {
                                            var v28 uint32 = 2147483648
                                            _ = len(v4)
                                            v3 = (fi16 - (fi13 - fi13))
                                            _ = v28
                                        }
                                        switch ((v25 + v25) + v25) {
                                        case "foo bar":
                                            switch fi13 {
                                            case 2:
                                                var v29 int32 = -3
                                                _ = v29
                                            default:
                                                switch "\n\t" {
                                                case "a":
                                                    fmt.Println(fi16)
                                                default:
                                                    switch (fi13 | (fi13 | v3)) {
                                                    case 143:
                                                        v3 = 426
                                                    case 127:
                                                        if (!(!true)) {
                                                            _ = len(v1)
                                                            switch v25 {
                                                            case "\n\t":
                                                                _ = len(v4)
                                                            default:
                                                                v2 += (v2 + v2)
                                                            }
                                                            var v30 int64 = 42
                                                            fmt.Println(fi16, fi13, v25)
                                                            var v31 float32 = (float32((((0.0 * -2.25) - 1e10) / float64(((0.0 * 0.0) - float32(((float64(float32(1e10)) + float64(1e10)) * (0.0 + (((3.14159 * 0.0) - ((3.14159 - float64(0.0)) - 3.14159)) - 0.1)))))))) - 0.0)
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                    default:
                                                        var v32 int16 = 0
                                                        _ = v32
                                                    }
                                                }
                                            }
                                        default:
                                            v25 = (("foo bar" + (v25 + ("foo bar" + (("foo bar" + v25) + v25)))) + "hello")
                                        }
                                        if false {
                                            _ = len(v1)
                                            if false {
                                                var v33 int16 = 0
                                                fmt.Println(v33, v14)
                                                var v34 float32 = 1.5
                                                _ = v33
                                                _ = v34
                                            } else {
                                                _ = len(v4)
                                                v35 := GenS1{F0: (v14 - fi16), F1: 0}
                                                v3 -= (fi13 / (((^(v3))) | 1))
                                                var v36 string = ((v25 + "") + (v25 + (v25 + (v25 + v25))))
                                                fmt.Println(v14, v25, v25)
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                        _ = v25
                                    default:
                                        fmt.Println(v14, fi13)
                                    }
                                } else {
                                    var v37 int16 = 32767
                                    _ = v37
                                }
                                v2 += v2
                            } else {
                                delete(v4, 63)
                            }
                        }
                        var v38 uint16 = v2
                        _ = fi16
                        _ = v38
                    }
                case -1:
                    if true {
                        if false {
                            v3 += fi13
                            var v39 int8 = -127
                            delete(v4, 0)
                            _ = v39
                        }
                        v2 = 0
                    }
                case 128:
                    v40 := GenS4{F0: 255}
                    _ = v40
                default:
                    if false {
                        v5 -= (v5 * 0)
                    }
                }
                var v41 int8 = -1
                v2 = (v2 ^ (v2 & 792))
                _ = v41
            }
            _ = fi13
            _ = v14
        }
        var v42 int8 = 100
        v5 = 30
        _ = v42
    }
    var v43 int32 = -673
    if (2 >= 857) {
        for fi44 := 0; fi44 < 13; fi44++ {
            if (!false) {
                var v45 uint16 = v2
                _ = v45
            }
            for fi46 := 0; fi46 < 20; fi46++ {
                _ = len(v6)
                if true {
                    _ = len(v1)
                    delete(v4, 0)
                } else {
                    v3 = v3
                }
                _ = fi46
            }
            _ = len(v4)
            _ = fi44
        }
        v2 += (v2 * v2)
        if (1e10 < ((-2.25 / 1.0) * (0.0 * 1e10))) {
            for fi47 := 0; fi47 < 3; fi47++ {
                for fi48 := 0; fi48 < 1; fi48++ {
                    switch "\n\t" {
                    case "":
                        if (-1 > (^(fi48))) {
                            v49 := GenS5{F0: v5}
                            switch "a" {
                            case "\n\t":
                                var v50 uint16 = v2
                                _ = v50
                            }
                            var v51 float32 = float32(1.5)
                            _ = v49
                            _ = v51
                        }
                    case "foo bar":
                        var v52 uint = v5
                        _ = v52
                    default:
                        switch "hello" {
                        case "hello":
                            var v53 int = fi48
                            _ = v53
                        case "a":
                            v54 := GenS6{F0: 1, F1: ("hello" + "foo bar")}
                            _ = v54
                        case "":
                            delete(v4, 559)
                        default:
                            fmt.Println(fi48)
                        }
                    }
                    v55 := GenS3{F0: (v43 + v43), F1: 128}
                    _ = fi48
                    _ = v55
                }
                var v56 int = (fi47 + v3)
                _ = fi47
                _ = v56
            }
        } else {
            v2 += (v2 * (v2 / (((^(v2))) | 1)))
            if ((!true) && true) {
                v3 -= v3
            } else {
                v57 := GenS4{F0: uint16(v5)}
                _ = v57
            }
            if (0.0 >= 0.0) {
                fmt.Println(v5)
                switch 502 {
                case 42:
                    v2 += v2
                default:
                    v2 -= (v2 / (((v2 & 324)) | 1))
                }
                var v58 int8 = 0
                _ = v58
            }
        }
    } else {
        var v59 uint16 = (v2 & 529)
        if false {
            for fi60 := 0; fi60 < 1; fi60++ {
                v59 = 65535
                var v61 bool = true
                var v62 uint = v5
                _ = fi60
                _ = v61
                _ = v62
            }
            switch int(v59) {
            case -943:
                v59 = (v2 - v59)
            case 0:
                if ("" < (("a" + "\n\t") + "a")) {
                    var v63 int32 = (v43 & 0)
                    _ = v63
                } else {
                    var v64 float32 = -2.25
                    fmt.Println(v64, v59, v43)
                    v65 := GenS7{F0: (v5 + (v5 % (((^(v5))) | 1))), F1: 0}
                    _ = v64
                    _ = v65
                }
            case -1:
                if false {
                    v66 := GenS8{F0: (^(v3))}
                    var v67 bool = (9223372036854775807 < -9223372036854775808)
                    _ = v66
                    _ = v67
                }
            }
        }
        var v68 int8 = -128
        _ = v59
        _ = v68
    }
    v69 := GenS1{F0: v3, F1: (v3 | (v3 | int(v2)))}
    v70 := GenS3{F0: -32, F1: 92}
    v2 -= v2
    if false {
        var v71 int16 = 1000
        _ = v71
    }
    v72 := v4[133]
    v73 := GenS9{F0: -128, F1: 42, F2: 450}
    if false {
        for fi74 := 0; fi74 < 3; fi74++ {
            for fi75 := 0; fi75 < 16; fi75++ {
                v3 = -810
                v72 -= v72
                var v76 int = (fi74 & 2147483647)
                _ = fi75
                _ = v76
            }
            _ = fi74
        }
    }
    for fi77 := 0; fi77 < 17; fi77++ {
        var v78 int16 = 849
        var v79 float64 = 1.5
        _ = fi77
        _ = v78
        _ = v79
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v72
    _ = v73
}
