package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int16; F2 uint }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int32, p2 string) uint32 {
    switch 253 {
    case -770:
        for fi7 := 0; fi7 < 7; fi7++ {
            switch p2 {
            case "a":
                if (!((0.1 + 3.14159) != 1e10)) {
                    if true {
                        v8 := GenS1{F0: (p1 >> ((20) & 63)), F1: 937, F2: p2}
                        _ = v8
                    } else {
                        v9 := GenS2{F0: (p1 / ((p1) | 1))}
                        var v10 uint32 = 120
                        v11 := GenS1{F0: (p1 >> ((255) & 63)), F1: p1, F2: "a"}
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    var v12 int8 = 42
                    _ = v12
                } else {
                    p2 = p2
                }
            case "\n\t":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int32 = (^(p1))
                    fmt.Println(p1, fi7, p0)
                    _ = fi13
                    _ = v14
                }
            case "":
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 int = (fi7 * (fi15 | (fi15 ^ (fi15 | (fi15 * fi15)))))
                    _ = fi15
                    _ = v16
                }
            }
            var v17 int = (fi7 * (fi7 & (fi7 / ((fi7) | 1))))
            _ = fi7
            _ = v17
        }
    }
    switch "hello" {
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            p1 = (p1 >> ((65535) & 63))
            if false {
                p1 += p1
                var v19 int8 = 100
                fmt.Println(p1, p1, p1)
                _ = v19
            }
            _ = fi18
        }
    case "":
        switch p2 {
        case "":
            v20 := GenS2{F0: p1}
            _ = v20
        case "foo bar":
            switch 128 {
            case 128:
                var v21 uint16 = 2
                _ = v21
            default:
                if true {
                    fmt.Println(p1, p0, p0)
                    v22 := []int{-9223372036854775808, 32768}
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS2{F0: p1}
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                } else {
                    var v25 uint = 255
                    _ = v25
                }
            }
        case "\n\t":
            var v26 float32 = float32(3.14159)
            _ = v26
        default:
            var v27 bool = false
            _ = v27
        }
    default:
        p0 = (p0 % ((p0) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 348
    var v2 int16 = 0
    var v3 int32 = 0
    var v4 string = "a"
    var v5 float32 = 1e10
    var v6 float64 = 1e10
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 float64 = v6
        _ = fi28
        _ = v29
    }
    fmt.Println(v6, v5, v3)
    var v30 int16 = 1
    var v31 uint = (v1 << ((393) & 63))
    v32 := GenS3{F0: false, F1: 1, F2: (v30 > v30)}
    v33 := []int{32767, 32767}
    switch v4 {
    case "\n\t":
        fmt.Println(v5)
    case "foo bar":
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 uint = 42
            v33 = append(v33, (fi34 % ((9223372036854775807) | 1)))
            _ = fi34
            _ = v35
        }
    }
    fmt.Println(v30, v4)
    var v36 string = ((v4 + "a") + "\n\t")
    v6 = (v6 + (0.0 + 1e10))
    switch 32767 {
    case -9223372036854775808:
        if (127 >= 6) {
            for fi37 := 0; fi37 < 2; fi37++ {
                var v38 bool = (604 > 2)
                v33 = append(v33, (fi37 >> (((^(v31))) & 63)))
                _ = fi37
                _ = v38
            }
        } else {
            v30 = (v30 ^ v2)
            if false {
                var v39 bool = true
                switch v36 {
                case "foo bar":
                    v2 = v30
                case "\n\t":
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                }
                if true {
                    v36 = (("" + ("hello" + ("" + v4))) + (((v36 + "a") + (v36 + "hello")) + ("\n\t" + "foo bar")))
                    _ = len(v33)
                } else {
                    var v40 uint16 = 255
                    v40 = v40
                    _ = v40
                }
                _ = v39
            } else {
                if (!false) {
                    v2 += v2
                }
                v31 += (v1 - v1)
                switch v4 {
                case "":
                    v41 := GenS3{F0: (v5 != v5), F1: (v3 >> ((v31) & 63)), F2: (750 >= 1)}
                    _ = v41
                case "a":
                    v42 := GenS2{F0: v3}
                    _ = v42
                default:
                    var v43 int8 = -118
                    _ = v43
                }
            }
        }
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        v36 = v4
        _ = fi44
    }
    v4 = v4
    fmt.Println(v6, v4, v3)
    if false {
        v4 = v4
        if true {
            v31 -= (v1 | v31)
            var v45 string = v36
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v45
        } else {
            switch 128 {
            case 9223372036854775807:
                var v46 uint64 = 754
                _ = v46
            case 0:
                var v47 int16 = (v2 ^ v2)
                _ = v47
            default:
                if false {
                    v4 = v4
                    for fi48 := 0; fi48 < 3; fi48++ {
                        for fi49 := 0; fi49 < 4; fi49++ {
                            v50 := GenS4{F0: 361}
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            v36 = v36
                            _ = fi49
                            _ = v50
                        }
                        var v51 uint = (^(v31))
                        for fi52 := 0; fi52 < 15; fi52++ {
                            fmt.Println(fi48)
                            var v53 uint16 = 91
                            _ = fi52
                            _ = v53
                        }
                        _ = fi48
                        _ = v51
                    }
                    switch "" {
                    case "foo bar":
                        if false {
                            if true {
                                var v54 int = 2147483647
                                switch (v54 ^ (v54 >> (((v31 & 554)) & 63))) {
                                case 9223372036854775807:
                                    if true {
                                        v55 := GenS2{F0: -1}
                                        v56 := GenS4{F0: 4294967295}
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v33 = append(v33, 363)
                                        fmt.Println(v31)
                                    }
                                }
                                var v57 bool = false
                                _ = v54
                                _ = v57
                            }
                        }
                    }
                } else {
                    if (true && true) {
                        v4 = "foo bar"
                        v58 := GenS4{F0: 279}
                        if false {
                            var v59 string = v4
                            if false {
                                v4 = ""
                            }
                            v60 := GenS2{F0: 2147483647}
                            _ = v59
                            _ = v60
                        }
                        _ = v58
                    } else {
                        for fi61 := 0; fi61 < 17; fi61++ {
                            switch (fi61 << (((v1 >> ((436) & 63))) & 63)) {
                            case 32767:
                                fmt.Println(v3, v2, v6)
                            case -976:
                                v2 += (v30 & (v2 % ((v2) | 1)))
                            default:
                                if (!false) {
                                    v31 = 0
                                }
                            }
                            v62 := GenS5{F0: v3, F1: 42, F2: v1}
                            _ = fi61
                            _ = v62
                        }
                    }
                    v5 = -2.25
                }
            }
        }
        switch 0 {
        case 32768:
            for fi63 := 0; fi63 < 20; fi63++ {
                v6 = v6
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                _ = fi63
            }
        case -1:
            v64 := GenS6{F0: 65535}
            _ = v64
        }
    } else {
        var v65 bool = (false || false)
        for fi66 := 0; fi66 < 3; fi66++ {
            v6 += 1e10
            v2 = 1000
            var v67 uint = v31
            _ = fi66
            _ = v67
        }
        _ = v65
    }
    if true {
        switch 48 {
        case 42:
            var v68 float64 = 1.5
            _ = v68
        case -9223372036854775808:
            switch 9223372036854775807 {
            case 127:
                var v69 int16 = (^(v2))
                _ = v69
            case 32767:
                var v70 bool = true
                _ = v70
            default:
                for fi71 := 0; fi71 < 10; fi71++ {
                    var v72 uint16 = 596
                    v30 = v2
                    v73 := GenS3{F0: (v5 >= float32((v6 + v6))), F1: (v3 * v3), F2: ((v3 & v3) == -2147483648)}
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            }
        case 241:
            var v74 float32 = float32(v6)
            _ = v74
        default:
            if false {
                _ = len(v33)
                v1 -= v1
                var v75 uint64 = 4294967295
                _ = v75
            } else {
                v2 = v2
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
