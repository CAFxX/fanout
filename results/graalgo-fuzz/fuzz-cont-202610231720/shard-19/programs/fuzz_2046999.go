package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 int; F2 uint16 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0(p0 bool, p1 string) uint64 {
    var v7 int8 = 0
    if (1 <= 32767) {
        if false {
            v7 = 0
            for fi8 := 0; fi8 < 9; fi8++ {
                v9 := []int{(fi8 | fi8), (fi8 + fi8), 370, 160}
                v10 := GenS1{F0: 0, F1: 465, F2: 1}
                var v11 bool = ((float64(1.5) * 1e10) < 0.0)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS2{F0: -699, F1: 1000, F2: -234}
                if true {
                    var v14 int = (-(fi12))
                    var v15 string = p1
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        var v17 uint64 = 42
                        switch fi12 {
                        case 581:
                            var v18 uint32 = 4294967295
                            _ = v18
                        case -9223372036854775808:
                            v7 -= v7
                        default:
                            v7 = (v7 & (v7 + 20))
                        }
                        v19 := GenS2{F0: 1, F1: 87, F2: 32767}
                        var v20 int = -19
                        v21 := GenS3{F0: 0, F1: -599, F2: p1}
                        _ = fi16
                        _ = v17
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v7 += (v7 * v7)
                }
                v22 := []int{(fi12 + fi12)}
                _ = fi12
                _ = v13
                _ = v22
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 172
}
func genFn1(p0 uint) uint8 {
    var v23 int32 = -1
    v24 := []int{95}
    if false {
        v25 := GenS1{F0: 0, F1: -9223372036854775808, F2: 2147483647}
        var v26 float64 = ((0.1 + (1.5 - 0.0)) / 1.0)
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 1, F1: 9223372036854775807, F2: 2}
    if false {
        v23 += v23
    }
    _ = p0
    _ = v23
    _ = v24
    _ = v27
    return 174
}

func main() {
    var v1 uint8 = 42
    v2 := []string{}
    var v3 uint8 = (v1 | (v1 * (v1 ^ 129)))
    var v4 string = ""
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 694
    v28 := GenS2{F0: int64(v6), F1: -1, F2: 1}
    fmt.Println(v3)
    if true {
        switch 32767 {
        case 4294967295:
            fmt.Println(v3)
        case -1:
            v1 = v3
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(v3)
            fmt.Println(fi29)
            _ = fi29
        }
    }
    v3 -= (v1 << ((255) & 63))
    var v30 int = -222
    v31 := GenS1{F0: 552, F1: 693, F2: 0}
    if true {
        fmt.Println(v1, v6)
        v32 := GenS3{F0: 0, F1: -32768, F2: v4}
        _ = v32
    }
    switch -571 {
    case 9223372036854775807:
        switch v30 {
        case -9223372036854775808:
            v30 = (v30 * v30)
        }
    default:
        v4 = v4
    }
    v4 = v4
    v33 := GenS4{F0: (v30 | v30), F1: (v30 ^ v30), F2: 157}
    if len(v2) > 0 {
        v2[0] = ((v4 + v4) + v4)
    }
    var v34 bool = true
    if v34 {
        if true {
            v35 := GenS4{F0: (v30 & v30), F1: 32768, F2: 95}
            switch ("" + (v4 + v4)) {
            case "":
                if (833 <= 42) {
                    switch "hello" {
                    case "foo bar":
                        var v36 uint16 = 1
                        _ = v36
                    case "":
                        var v37 int16 = 444
                        _ = v37
                    default:
                        v38 := GenS4{F0: v30, F1: -1, F2: 571}
                        _ = v38
                    }
                    var v39 int32 = 1
                    _ = v39
                }
            case "foo bar":
                v1 += (v1 / ((genFn1(0)) | 1))
            case "\n\t":
                var v40 uint64 = (v6 % ((v6) | 1))
                _ = v40
            }
            _ = v35
        } else {
            switch 4294967295 {
            case 32768:
                var v41 float64 = float64((3.14159 - float32((1.5 / ((3.14159 * 3.14159) / float64(-2.25))))))
                _ = v41
            }
            _ = len(v5)
            v42 := GenS4{F0: v30, F1: (^(v30)), F2: 1}
            _ = v42
        }
        if v34 {
            for fi43 := 0; fi43 < 2; fi43++ {
                v4 = ""
                fmt.Println(v4, v3)
                _ = fi43
            }
        }
    } else {
        for fi44 := 0; fi44 < 10; fi44++ {
            v1 = (v3 * (^(v1)))
            v6 -= (^(v6))
            if v34 {
                for fi45 := 0; fi45 < 17; fi45++ {
                    if (fi45 <= fi45) {
                        var v46 uint = 262
                        if true {
                            v34 = v34
                            switch ("foo bar" + ((v4 + (("a" + (v4 + (((v4 + v4) + ("a" + v4)) + v4))) + v4)) + v4)) {
                            case "\n\t":
                                var v47 int8 = 127
                                _ = v47
                            case "a":
                                v2 = append(v2, ("a" + v4))
                            default:
                                switch int(v46) {
                                case 32767:
                                    switch v4 {
                                    case "\n\t":
                                        var v48 bool = ((v46 < 1) && (!(65535 != 4294967295)))
                                        v4 = v4
                                        var v49 int32 = -1
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            }
                        }
                        v30 += (fi44 / ((fi45) | 1))
                        _ = v46
                    }
                    var v50 uint = 524
                    fmt.Println(v50, fi44)
                    _ = fi45
                    _ = v50
                }
                for fi51 := 0; fi51 < 18; fi51++ {
                    if (1000 < -32768) {
                        v4 = ""
                        if v34 {
                            _ = len(v5)
                            var v52 uint16 = 0
                            _ = v52
                        }
                        if v34 {
                            v1 = (v3 << ((65535) & 63))
                        } else {
                            v6 += (v6 << ((1) & 63))
                        }
                    } else {
                        v3 += (v3 ^ v1)
                        v53 := GenS2{F0: 9223372036854775807, F1: -744, F2: 1000}
                        var v54 uint64 = 1
                        var v55 bool = ((!(0 <= 42)) || ((v1 % ((genFn1(0)) | 1)) == 212))
                        fmt.Println(v30)
                        fmt.Println(v4, fi44)
                        if v55 {
                            _ = len(v2)
                        }
                        _ = len(v2)
                        var v56 string = v4
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    var v57 int32 = -2147483648
                    _ = fi51
                    _ = v57
                }
                v58 := GenS5{F0: v30, F1: 1}
                _ = v58
            } else {
                v59 := GenS5{F0: 310, F1: 42}
                v60 := v5[2147483648]
                v34 = true
                _ = v59
                _ = v60
            }
            _ = fi44
        }
    }
    v34 = (!v34)
    fmt.Println(v34, v3, v34)
    v61 := GenS6{F0: (v30 + v30), F1: -32768}
    var v62 float32 = (1e10 * 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
