package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool) uint64 {
    switch 409 {
    case 42:
        fmt.Println(p0)
    default:
        p0 = ((!p0) || (-915 <= 837))
    }
    switch "foo bar" {
    case "hello":
        var v6 int8 = 42
        _ = v6
    case "foo bar":
        fmt.Println(p0)
    case "a":
        switch -789 {
        case 603:
            p0 = p0
        case 675:
            switch 42 {
            case 9223372036854775807:
                switch "a" {
                case "hello":
                    fmt.Println(p0)
                default:
                    for fi7 := 0; fi7 < 16; fi7++ {
                        var v8 float32 = (-2.25 - (((3.14159 / 3.14159) - float32(3.14159)) + -2.25))
                        p0 = (0 >= 646)
                        var v9 string = "foo bar"
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                }
            case 900:
                for fi10 := 0; fi10 < 3; fi10++ {
                    if (0 <= 42) {
                        v11 := GenS1{F0: fi10, F1: 879, F2: "hello"}
                        if p0 {
                            for fi12 := 0; fi12 < 19; fi12++ {
                                v13 := []int{fi10, (fi10 - fi12)}
                                if p0 {
                                    var v14 uint = 2
                                    p0 = (753 > 192)
                                    if true {
                                        if p0 {
                                            _ = len(v13)
                                        }
                                        fmt.Println(v14)
                                        v15 := GenS2{F0: -1}
                                        _ = v15
                                    } else {
                                        if (false && (0 >= 42)) {
                                            fmt.Println(fi10, fi12, fi10)
                                        } else {
                                            var v16 string = "foo bar"
                                            if len(v13) > 0 {
                                                v13[0] = 912
                                            }
                                            _ = v16
                                        }
                                    }
                                    _ = v14
                                }
                                v17 := GenS1{F0: fi12, F1: 878, F2: "a"}
                                _ = fi12
                                _ = v13
                                _ = v17
                            }
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        } else {
                            p0 = p0
                            switch "hello" {
                            case "\n\t":
                                var v19 float32 = float32(0.1)
                                _ = v19
                            case "a":
                                fmt.Println(p0)
                            }
                        }
                        switch fi10 {
                        case 1:
                            switch "" {
                            case "\n\t":
                                var v20 uint8 = 255
                                _ = v20
                            case "hello":
                                var v21 uint64 = 42
                                _ = v21
                            case "":
                                v22 := []int{(fi10 * (fi10 ^ -1))}
                                _ = v22
                            default:
                                v23 := []int{fi10}
                                _ = v23
                            }
                        case -651:
                            if false {
                                switch fi10 {
                                case 2147483647:
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        p0 = true
                                    case "hello":
                                        fmt.Println(p0)
                                    }
                                case 9223372036854775807:
                                    v24 := []int{(^(fi10)), fi10, fi10, -2147483648}
                                    _ = v24
                                }
                            }
                        case 819:
                            fmt.Println(p0, fi10)
                        }
                        _ = v11
                    } else {
                        v25 := GenS3{F0: (!p0), F1: fi10, F2: 968}
                        for fi26 := 0; fi26 < 9; fi26++ {
                            p0 = true
                            if (!p0) {
                                var v27 int64 = -9223372036854775808
                                v27 = v27
                                _ = v27
                            } else {
                                var v28 uint32 = 2147483648
                                _ = v28
                            }
                            v29 := []int{fi10, 4294967295, (fi26 - fi10)}
                            _ = fi26
                            _ = v29
                        }
                        _ = v25
                    }
                    if p0 {
                        fmt.Println(p0)
                        if true {
                            v30 := []int{}
                            _ = v30
                        } else {
                            var v31 int32 = -2147483648
                            for fi32 := 0; fi32 < 9; fi32++ {
                                v31 -= (v31 * (v31 - 1))
                                if p0 {
                                    var v33 uint64 = 1
                                    _ = v33
                                } else {
                                    p0 = (p0 && true)
                                }
                                var v34 int32 = 12
                                _ = fi32
                                _ = v34
                            }
                            var v35 int = fi10
                            _ = v31
                            _ = v35
                        }
                    } else {
                        var v36 int64 = 1
                        _ = v36
                    }
                    v37 := []int{(fi10 + (fi10 / ((fi10) | 1)))}
                    _ = fi10
                    _ = v37
                }
            case 243:
                var v38 bool = (67 < 152)
                _ = v38
            default:
                switch 32767 {
                case 32768:
                    v39 := []int{4294967295}
                    _ = v39
                case -793:
                    var v40 uint8 = 1
                    _ = v40
                default:
                    var v41 float64 = 1.5
                    _ = v41
                }
            }
        case 127:
            switch "\n\t" {
            case "foo bar":
                var v42 int32 = -1
                _ = v42
            case "\n\t":
                fmt.Println(p0)
            default:
                var v43 float32 = 1e10
                _ = v43
            }
        }
    }
    _ = p0
    return 4294967295
}
func genFn1(p0 int16, p1 bool) uint8 {
    var v44 int = 4294967295
    if (135 > 2) {
        if ((255 <= 159) && true) {
            p1 = p1
            v45 := []int{(v44 / ((1) | 1))}
            v46 := GenS4{F0: (v44 + v44), F1: 1}
            _ = v45
            _ = v46
        } else {
            v44 = (^(v44))
            for fi47 := 0; fi47 < 13; fi47++ {
                var v48 int8 = 42
                if p1 {
                    v49 := GenS5{F0: int64(fi47), F1: 149, F2: (v44 * v44)}
                    _ = v49
                }
                _ = fi47
                _ = v48
            }
            var v50 float64 = 1.5
            _ = v50
        }
        var v51 uint = 0
        var v52 int64 = 1
        _ = v51
        _ = v52
    } else {
        var v53 uint64 = genFn0((-576 == 42))
        _ = v53
    }
    for fi54 := 0; fi54 < 4; fi54++ {
        if p1 {
            var v55 float64 = (0.0 / 1.0)
            _ = v55
        }
        _ = fi54
    }
    _ = p0
    _ = p1
    _ = v44
    return 127
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    var v3 int32 = 442
    v4 := make(map[uint32]string)
    var v5 uint = 0
    switch "hello" {
    case "\n\t":
        var v56 int16 = -192
        _ = v56
    case "a":
        for fi57 := 0; fi57 < 6; fi57++ {
            if true {
                var v58 bool = false
                switch (fi57 * fi57) {
                case 4294967295:
                    var v59 int8 = 42
                    _ = v59
                case 42:
                    fmt.Println(v3)
                }
                _ = v58
            }
            v3 -= 0
            _ = fi57
        }
    case "foo bar":
        _ = len(v4)
    default:
        var v60 string = "hello"
        _ = v60
    }
    var v61 int32 = v3
    for fi62 := 0; fi62 < 8; fi62++ {
        for fi63 := 0; fi63 < 4; fi63++ {
            v64 := GenS6{F0: 127, F1: "foo bar"}
            v65 := GenS7{F0: -1, F1: 2}
            var v66 uint32 = 889
            _ = fi63
            _ = v64
            _ = v65
            _ = v66
        }
        _ = fi62
    }
    v61 = v3
    if true {
        switch "" {
        case "\n\t":
            v61 = v61
        case "a":
            if false {
                v1 -= float64(1e10)
                for fi67 := 0; fi67 < 14; fi67++ {
                    v61 = 42
                    for fi68 := 0; fi68 < 9; fi68++ {
                        v3 -= (v3 & (v3 * (v3 + (v61 << (((v5 ^ v5)) & 63)))))
                        var v69 float64 = (v1 / (float64(((-2.25 + 1e10) / (-2.25 / (-2.25 + 0.1)))) / v1))
                        _ = fi68
                        _ = v69
                    }
                    _ = fi67
                }
                for fi70 := 0; fi70 < 11; fi70++ {
                    for fi71 := 0; fi71 < 19; fi71++ {
                        var v72 int = fi71
                        var v73 int32 = int32(v5)
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    _ = fi70
                }
            } else {
                for fi74 := 0; fi74 < 12; fi74++ {
                    v75 := []int{}
                    _ = fi74
                    _ = v75
                }
                switch 372 {
                case 902:
                    v3 = (v61 & (v3 / (((v61 ^ v3)) | 1)))
                case -9223372036854775808:
                    var v76 int8 = 109
                    _ = v76
                default:
                    v77 := []int{4294967295, -398, 0, -1}
                    _ = v77
                }
            }
        }
    }
    fmt.Println(v5, v5, v61)
    v4[4294967295] = "foo bar"
    v5 = (v5 >> (((^(v5))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
