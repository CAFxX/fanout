package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 string; F1 int8; F2 bool }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int64; F1 int32 }
func genFn0(p0 int) uint64 {
    v4 := []int{p0, 4294967295, 720, (p0 << ((42) & 63))}
    v5 := GenS1{F0: 1, F1: 58, F2: -2147483648}
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int16, p1 int32, p2 uint8) string {
    p0 = 1000
    var v6 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (("" + "\n\t") + (((("a" + ("hello" + "a")) + "") + ("hello" + "")) + ("" + "foo bar")))
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    if false {
        if (0 < 699) {
            v7 := GenS2{F0: 2, F1: 97}
            _ = v7
        }
        v1 -= v1
        if (586 >= 641) {
            if (-371 > -9223372036854775808) {
                v8 := GenS2{F0: 127, F1: 127}
                v9 := GenS1{F0: 65535, F1: v1, F2: -2147483648}
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 int8 = 1
                    _ = fi10
                    _ = v11
                }
                _ = v8
                _ = v9
            }
        } else {
            v1 = v1
        }
    } else {
        fmt.Println(v1)
        v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
    }
    v12 := GenS1{F0: 1, F1: (v1 ^ v1), F2: 0}
    for fi13 := 0; fi13 < 6; fi13++ {
        v1 += (v1 * v1)
        _ = fi13
    }
    v1 = v1
    if (false && false) {
        var v14 uint16 = v1
        v14 = 65535
        var v15 uint32 = 867
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: 181, F1: (v1 << ((255) & 63)), F2: 0}
        v17 := GenS3{F0: v1, F1: 255}
        _ = v16
        _ = v17
    }
    var v18 int32 = -2147483648
    v18 = (v18 ^ v18)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS2{F0: 169, F1: 84}
        if true {
            v21 := GenS1{F0: 1, F1: (v1 + v1), F2: v18}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    var v22 uint32 = 42
    switch "" {
    case "a":
        v23 := []int{}
        _ = v23
    case "hello":
        if true {
            var v24 float64 = 0.1
            for fi25 := 0; fi25 < 8; fi25++ {
                if true {
                    var v26 uint = 255
                    fmt.Println(v26)
                    var v27 int32 = v18
                    _ = v26
                    _ = v27
                }
                if true {
                    fmt.Println(fi25, v1)
                    v1 = 65535
                }
                _ = fi25
            }
            _ = v24
        }
    case "foo bar":
        if (!false) {
            var v28 int64 = -763
            v29 := []int{2147483647, 377, -9223372036854775808}
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        if true {
            fmt.Println(v22, fi30)
            var v31 int = int(v1)
            var v32 int32 = (v18 >> ((543) & 63))
            _ = v31
            _ = v32
        }
        _ = fi30
    }
    v33 := GenS4{F0: "foo bar", F1: 69, F2: (true || true)}
    fmt.Println(v22, v1, v1)
    switch 9223372036854775807 {
    case -9223372036854775808:
        v1 = v1
    case -1:
        v1 = v1
    case 32768:
        for fi34 := 0; fi34 < 7; fi34++ {
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := GenS3{F0: 65535, F1: 454}
                if (4294967295 > 669) {
                    delete(v3, (v22 - (^(v22))))
                }
                v22 = (v22 + 382)
                var v37 float64 = float64(0.0)
                v38 := GenS4{F0: "hello", F1: 0, F2: (!false)}
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
            _ = fi34
        }
    default:
        v1 += (v1 << ((1) & 63))
    }
    var v39 int = 127
    switch ("hello" + ("hello" + "hello")) {
    case "\n\t":
        v40 := GenS3{F0: (v1 >> ((366) & 63)), F1: 2}
        _ = v40
    case "hello":
        switch (^(v39)) {
        case 718:
            if true {
                if false {
                    if false {
                        v41 := GenS1{F0: 1, F1: 432, F2: -1}
                        switch "\n\t" {
                        case "\n\t":
                            v22 += (v22 + (v22 / (((v22 % ((4294967295) | 1))) | 1)))
                        case "":
                            v42 := GenS5{F0: (v18 % ((-1) | 1)), F1: 4294967295}
                            _ = v42
                        case "foo bar":
                            fmt.Println(v1)
                        }
                        _ = v41
                    } else {
                        v39 = 1
                        v18 = (v18 * (v18 & v18))
                    }
                    v18 -= v18
                } else {
                    v1 -= (v1 % (((v1 ^ v1)) | 1))
                }
            } else {
                v39 += (v39 | v39)
                var v43 int32 = (v18 & (^(v18)))
                switch "hello" {
                case "":
                    for fi44 := 0; fi44 < 15; fi44++ {
                        v45 := []int{(fi44 + (v39 | (fi44 / ((-2147483648) | 1)))), (v39 ^ (^(fi44))), fi44}
                        _ = fi44
                        _ = v45
                    }
                case "hello":
                    fmt.Println(v1, v18, v43)
                default:
                    for fi46 := 0; fi46 < 2; fi46++ {
                        for fi47 := 0; fi47 < 15; fi47++ {
                            switch fi47 {
                            case 32768:
                                v48 := GenS2{F0: 16, F1: 2}
                                _ = v48
                            case 0:
                                v22 -= (v22 >> ((1) & 63))
                            }
                            if (1 == 1000) {
                                v49 := GenS6{F0: 127, F1: (v22 & v22), F2: 42}
                                switch fi46 {
                                case 689:
                                    switch -2147483648 {
                                    case 2147483647:
                                        var v50 uint64 = 816
                                        _ = v50
                                    case -9223372036854775808:
                                        v51 := []int{}
                                        _ = v51
                                    case 128:
                                        switch (v39 * v39) {
                                        case 127:
                                            var v52 float32 = 0.0
                                            _ = v52
                                        case -687:
                                            if true {
                                                var v53 int16 = -32768
                                                if ((42 == -804) || false) {
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v39, v1)
                                                    case "a":
                                                        if true {
                                                            v39 -= (fi46 / (((v39 / ((0) | 1))) | 1))
                                                            var v54 uint8 = 127
                                                            _ = v54
                                                        } else {
                                                            switch fi47 {
                                                            case 2147483647:
                                                                v55 := GenS7{F0: 42, F1: 273}
                                                                _ = v55
                                                            case 276:
                                                                v43 = (v18 | v18)
                                                            default:
                                                                fmt.Println(fi46)
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v56 := GenS3{F0: v1, F1: 1}
                                                        _ = v56
                                                    default:
                                                        var v57 float64 = 0.1
                                                        _ = v57
                                                    }
                                                    var v58 uint64 = 42
                                                    switch fi46 {
                                                    case 2147483647:
                                                        fmt.Println(v39, v53, v53)
                                                    default:
                                                        v43 += 0
                                                    }
                                                    switch (v39 | (fi47 << ((837) & 63))) {
                                                    case -2147483648:
                                                        v59 := GenS7{F0: 1, F1: v18}
                                                        _ = v59
                                                    case -9223372036854775808:
                                                        if false {
                                                            v60 := GenS1{F0: 882, F1: (v1 * (v1 >> ((18446744073709551615) & 63))), F2: -441}
                                                            _ = v60
                                                        } else {
                                                            v61 := GenS7{F0: -797, F1: (v43 >> ((929) & 63))}
                                                            _ = v61
                                                        }
                                                    case 516:
                                                        v62 := []int{fi47}
                                                        _ = v62
                                                    }
                                                    _ = v58
                                                }
                                                if true {
                                                    var v63 int64 = 1
                                                    _ = v63
                                                } else {
                                                    fmt.Println(v39, fi47)
                                                    v1 += 0
                                                    v64 := v3[59]
                                                    _ = v64
                                                }
                                                _ = v53
                                            }
                                        case 846:
                                            v65 := GenS5{F0: 2147483647, F1: 0}
                                            _ = v65
                                        }
                                    }
                                default:
                                    v66 := GenS5{F0: v18, F1: genFn0((^(fi46)))}
                                    _ = v66
                                }
                                _ = v49
                            }
                            _ = fi47
                        }
                        _ = fi46
                    }
                }
                _ = v43
            }
        case -9223372036854775808:
            v67 := GenS5{F0: 465, F1: 278}
            _ = v67
        }
    }
    for fi68 := 0; fi68 < 18; fi68++ {
        var v69 float64 = 1e10
        _ = fi68
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v39
}
