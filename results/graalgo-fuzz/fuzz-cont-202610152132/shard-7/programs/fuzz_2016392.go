package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16 }

func main() {
    var v1 float64 = 1.5
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    var v4 uint = 1
    v5 := []uint32{2147483647, 2147483648}
    var v6 uint32 = 222
    if true {
        v3[0] = true
        v1 -= 0.0
        var v7 uint32 = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 0}
        v9 := GenS2{F0: 2}
        _ = v8
        _ = v9
    }
    var v10 uint32 = (v6 << (((^(v4))) & 63))
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint64 = 1
        var v13 int64 = 42
        var v14 int64 = -296
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    switch -1 {
    case -1:
        v6 += (v6 & v10)
    case 2147483647:
        var v15 string = "a"
        _ = v15
    case 42:
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                if true {
                    v5 = append(v5, (^(v6)))
                    var v18 bool = true
                    _ = len(v3)
                    var v19 float64 = float64((0.1 + 1.5))
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
            _ = fi16
        }
    default:
        v20 := GenS1{F0: 9223372036854775807}
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS2{F0: 161}
        _ = fi21
        _ = v22
    }
    if false {
        var v23 int8 = -128
        switch "hello" {
        case "":
            fmt.Println(v2, v4, v2)
        case "foo bar":
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 bool = ((255 < 2) || (-1 > v23))
                _ = fi24
                _ = v25
            }
        default:
            if (0 > 0) {
                fmt.Println(v4, v10, v23)
                var v26 uint32 = v6
                v27 := GenS3{F0: false, F1: 0}
                _ = v26
                _ = v27
            } else {
                v28 := GenS1{F0: -135}
                _ = v28
            }
        }
        _ = v23
    }
    var v29 uint32 = (v6 + (v10 | 2147483647))
    var v30 float32 = 0.1
    if (true && true) {
        if len(v5) > 0 {
            v5[0] = (^(v29))
        }
        var v31 int64 = -38
        _ = v31
    }
    switch 32768 {
    case 376:
        if false {
            if true {
                v32 := GenS3{F0: false, F1: v2}
                _ = v32
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v6 = (v6 * 4294967295)
            var v33 uint64 = 1
            _ = v33
        }
    case 4294967295:
        switch "a" {
        case "foo bar":
            v34 := GenS3{F0: false, F1: (v2 + (v2 & (v2 ^ (v2 + (v2 & int32(v4))))))}
            _ = v34
        case "a":
            var v35 float32 = (float32(((1.5 + (v1 + 3.14159)) * ((v1 / v1) / v1))) - v30)
            _ = v35
        case "":
            delete(v3, 2147483647)
        }
    }
    v4 = (v4 | (v4 & v4))
    switch "a" {
    case "hello":
        for fi36 := 0; fi36 < 19; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                switch (fi37 * 1) {
                case -6:
                    switch "hello" {
                    case "\n\t":
                        switch fi36 {
                        case 957:
                            if len(v5) > 0 {
                                v5[0] = (^(v6))
                            }
                        case -808:
                            switch fi36 {
                            case 1:
                                if true {
                                    v38 := v3[232]
                                    v30 += v30
                                    switch fi37 {
                                    case 32768:
                                        v39 := GenS1{F0: int64(v2)}
                                        _ = v39
                                    case 4294967295:
                                        v40 := GenS3{F0: false, F1: -2147483648}
                                        _ = v40
                                    case 953:
                                        v41 := GenS3{F0: v38, F1: v2}
                                        _ = v41
                                    }
                                    _ = v38
                                }
                            case 127:
                                var v42 int8 = int8(v2)
                                _ = v42
                            case 114:
                                var v43 int64 = -894
                                _ = v43
                            default:
                                v44 := GenS1{F0: 1}
                                _ = v44
                            }
                        case -329:
                            if (!false) {
                                v45 := GenS3{F0: (32767 == -32768), F1: (v2 | v2)}
                                v30 -= v30
                                _ = v45
                            }
                        default:
                            fmt.Println(v29, v4, v10)
                        }
                    case "foo bar":
                        var v46 int64 = -965
                        _ = v46
                    case "hello":
                        if true {
                            var v47 float32 = ((1e10 / float32((1e10 + v1))) + (1e10 + v30))
                            switch (fi36 >> (((^(v4))) & 63)) {
                            case -2147483648:
                                var v48 string = ("foo bar" + "\n\t")
                                _ = v48
                            default:
                                if false {
                                    fmt.Println(fi36)
                                    if (127 <= 42) {
                                        var v49 float32 = float32(v1)
                                        _ = v49
                                    }
                                } else {
                                    switch -440 {
                                    case 128:
                                        if false {
                                            switch fi37 {
                                            case 32768:
                                                v4 = (^(v4))
                                            }
                                            fmt.Println(fi37)
                                        } else {
                                            var v50 int16 = -966
                                            fmt.Println(fi37, fi36, v1)
                                            v5 = append(v5, v29)
                                            _ = v50
                                        }
                                    default:
                                        if true {
                                            _ = len(v3)
                                            v1 -= float64(v30)
                                            fmt.Println(v30, v47, v30)
                                            switch fi37 {
                                            case -950:
                                                v51 := GenS2{F0: 1}
                                                _ = v51
                                            default:
                                                var v52 uint8 = 42
                                                v53 := GenS2{F0: 1}
                                                var v54 int64 = 367
                                                v54 = -1
                                                if true {
                                                    var v55 uint64 = 0
                                                    _ = v55
                                                } else {
                                                    v56 := GenS2{F0: 42}
                                                    v52 = v52
                                                    if (!false) {
                                                        v57 := GenS3{F0: (!(!((v1 == float64(((v47 / 1e10) * 1e10))) || false))), F1: int32(v52)}
                                                        if true {
                                                            if true {
                                                                v58 := v3[500]
                                                                if false {
                                                                    v59 := GenS4{F0: (v4 >> ((18446744073709551615) & 63))}
                                                                    v10 = 59
                                                                    _ = v59
                                                                }
                                                                _ = v58
                                                            } else {
                                                                v4 = v4
                                                                v1 -= v1
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v60 := GenS5{F0: (v4 >> (((v4 + uint(v10))) & 63))}
                                                                    _ = v60
                                                                case "hello":
                                                                    var v61 bool = ((true || (!true)) || true)
                                                                    _ = v61
                                                                }
                                                                var v62 int64 = 42
                                                                _ = v62
                                                            }
                                                            var v63 int = fi37
                                                            _ = v63
                                                        }
                                                        _ = v57
                                                    } else {
                                                        var v64 string = "\n\t"
                                                        _ = v64
                                                    }
                                                    _ = v56
                                                }
                                                v65 := GenS2{F0: 2}
                                                if (((v2 == 2147483647) && (!false)) || (!(v52 == (^(v52))))) {
                                                    v66 := GenS6{F0: 1}
                                                    fmt.Println(v30, v4)
                                                    if false {
                                                        fmt.Println(v10)
                                                        fmt.Println(v1, fi37, v30)
                                                        v52 = (v52 | (v52 / (((v52 % ((v52) | 1))) | 1)))
                                                    } else {
                                                        fmt.Println(v54)
                                                        switch "a" {
                                                        case "a":
                                                            v67 := GenS5{F0: (^(v4))}
                                                            _ = v67
                                                        case "hello":
                                                            var v68 uint32 = 42
                                                            _ = v68
                                                        default:
                                                            if true {
                                                                switch 32767 {
                                                                case -690:
                                                                    v54 = (^(v54))
                                                                case 32767:
                                                                    v69 := GenS7{F0: (v10 ^ (v29 << ((uint(v29)) & 63))), F1: -128}
                                                                    _ = v69
                                                                }
                                                                v70 := v3[1]
                                                                _ = len(v3)
                                                                _ = v70
                                                            }
                                                        }
                                                        if true {
                                                            var v71 uint8 = uint8(v29)
                                                            switch fi37 {
                                                            case 0:
                                                                v30 = (v47 + v30)
                                                            case -486:
                                                                v71 -= (v71 | (v52 % ((v52) | 1)))
                                                                var v72 bool = ((v6 + v29) >= 0)
                                                                _ = v72
                                                            }
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v66
                                                }
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                                _ = v65
                                            }
                                        }
                                    }
                                    v10 = (v10 + v10)
                                    var v73 uint8 = 2
                                    _ = v73
                                }
                                fmt.Println(v4)
                                v74 := GenS7{F0: v6, F1: 0}
                                if false {
                                    delete(v3, v6)
                                    var v75 int64 = -73
                                    _ = v75
                                } else {
                                    var v76 string = (("foo bar" + "foo bar") + "hello")
                                    v76 = ((v76 + v76) + "a")
                                    v77 := GenS3{F0: true, F1: (^(v2))}
                                    switch v76 {
                                    case "a":
                                        fmt.Println(v29)
                                    default:
                                        fmt.Println(v1, v76)
                                    }
                                    _ = v76
                                    _ = v77
                                }
                                _ = v74
                            }
                            switch fi37 {
                            case 994:
                                var v78 bool = (0 != 1)
                                _ = v78
                            default:
                                _ = len(v3)
                            }
                            _ = v47
                        } else {
                            v79 := GenS8{F0: 2}
                            if false {
                                var v80 int16 = -1
                                var v81 uint32 = (v10 << ((2) & 63))
                                _ = v80
                                _ = v81
                            }
                            if ((!(float32((0.1 / v1)) <= v30)) && true) {
                                v10 = (^(v10))
                            } else {
                                if ((fi36 - int(v29)) <= 32768) {
                                    v6 = (v29 + 0)
                                } else {
                                    fmt.Println(v10, v30, fi37)
                                    if false {
                                        var v82 uint32 = (v6 >> (((v4 + 805)) & 63))
                                        _ = v82
                                    }
                                    v5 = append(v5, v6)
                                }
                                if (false || true) {
                                    v2 = v2
                                    fmt.Println(fi37)
                                } else {
                                    v30 = 1e10
                                    switch "hello" {
                                    case "hello":
                                        v83 := GenS9{F0: 42}
                                        _ = v83
                                    case "foo bar":
                                        if false {
                                            v84 := v3[34]
                                            v85 := GenS1{F0: 0}
                                            _ = v84
                                            _ = v85
                                        } else {
                                            switch ("hello" + "foo bar") {
                                            case "\n\t":
                                                v2 = v2
                                            }
                                        }
                                    case "":
                                        var v86 float32 = (1e10 - -2.25)
                                        _ = v86
                                    }
                                }
                            }
                            _ = v79
                        }
                    }
                case -2147483648:
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v87 := GenS6{F0: v29}
                            _ = v87
                        }
                    } else {
                        var v88 string = ""
                        _ = len(v5)
                        if false {
                            switch fi37 {
                            case 0:
                                if false {
                                    v5 = append(v5, v6)
                                    v89 := GenS10{F0: (v4 << (((v4 % ((v4) | 1))) & 63))}
                                    var v90 int = (fi37 & fi36)
                                    _ = v89
                                    _ = v90
                                } else {
                                    switch fi37 {
                                    case 0:
                                        v91 := GenS10{F0: (v4 + v4)}
                                        _ = v91
                                    case 9223372036854775807:
                                        v4 = v4
                                    default:
                                        var v92 string = "\n\t"
                                        _ = v92
                                    }
                                    v6 -= (v29 ^ (v29 << ((4294967295) & 63)))
                                }
                            case 32767:
                                fmt.Println(v6, v10, v10)
                            }
                            var v93 int16 = -32768
                            fmt.Println(v29, v2)
                            var v94 uint64 = 0
                            _ = v93
                            _ = v94
                        }
                        _ = v88
                    }
                default:
                    v95 := GenS11{F0: 2}
                    _ = v95
                }
                fmt.Println(v4, v29)
                v3[(^(v6))] = false
                _ = fi37
            }
            _ = fi36
        }
    case "a":
        v1 -= (0.1 * v1)
    case "foo bar":
        for fi96 := 0; fi96 < 14; fi96++ {
            _ = len(v3)
            v30 -= float32(v1)
            _ = fi96
        }
    }
    switch 0 {
    case 477:
        var v97 string = ""
        _ = v97
    case 815:
        _ = len(v3)
    case 4294967295:
        for fi98 := 0; fi98 < 5; fi98++ {
            v99 := GenS12{F0: 32767}
            v30 += (v30 * 1.5)
            _ = fi98
            _ = v99
        }
    default:
        v100 := GenS1{F0: 1}
        _ = v100
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
