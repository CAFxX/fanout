package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 bool; F2 int }
func genFn0() uint64 {
    v6 := GenS1{F0: (("hello" + "\n\t") + "a")}
    if ("a" >= ((("foo bar" + "\n\t") + ((("\n\t" + ("" + "hello")) + ("" + "a")) + "")) + ("\n\t" + ("\n\t" + (("" + ("a" + "")) + ("foo bar" + "a")))))) {
        switch 1 {
        case 127:
            for fi7 := 0; fi7 < 17; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    v9 := []int{fi7, 42}
                    fmt.Println(fi8)
                    _ = fi8
                    _ = v9
                }
                var v10 int8 = -123
                v11 := GenS2{F0: -1, F1: -1, F2: fi7}
                _ = fi7
                _ = v10
                _ = v11
            }
        case 112:
            var v12 bool = ((200 > 4294967295) || true)
            _ = v12
        case -384:
            v13 := GenS1{F0: ""}
            _ = v13
        }
        var v14 bool = (false && true)
        _ = v14
    } else {
        v15 := []int{806, 127, -240}
        if false {
            var v16 uint8 = 127
            if false {
                v15 = append(v15, 829)
            } else {
                var v17 uint32 = 45
                v18 := GenS1{F0: "hello"}
                var v19 uint16 = 2
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v16
        }
        _ = v15
    }
    var v20 int64 = 625
    v20 = (v20 - (v20 % ((v20) | 1)))
    fmt.Println(v20)
    _ = v20
    _ = v6
    return 0
}

func main() {
    var v1 bool = true
    var v2 int8 = -67
    var v3 string = ("foo bar" + "a")
    v4 := make(map[string]bool)
    var v5 uint8 = 0
    _ = len(v4)
    v21 := []int{-9223372036854775808, 127, 855}
    v22 := GenS1{F0: (("\n\t" + ((v3 + "a") + ("\n\t" + v3))) + "hello")}
    switch "a" {
    case "\n\t":
        switch ("foo bar" + ("foo bar" + v3)) {
        case "":
            if false {
                v23 := GenS2{F0: 1, F1: 42, F2: 106}
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 int64 = 172
                    var v26 string = v3
                    var v27 int16 = 1
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v23
            }
        }
    case "":
        fmt.Println(v5, v2, v5)
    default:
        v28 := GenS3{F0: uint32(v5), F1: 151}
        _ = v28
    }
    var v29 uint16 = 402
    v3 = v3
    if v1 {
        v3 = "foo bar"
        var v30 uint16 = v29
        _ = v30
    } else {
        v31 := GenS4{F0: (v2 / ((127) | 1))}
        switch -44 {
        case -770:
            if (225 == (v5 | (v5 / (((v5 % (((v5 & v5)) | 1))) | 1)))) {
                fmt.Println(v3, v3, v5)
            }
        case 128:
            v32 := GenS2{F0: -774, F1: 42, F2: 42}
            _ = v32
        case 296:
            for fi33 := 0; fi33 < 18; fi33++ {
                v2 = (v2 ^ v2)
                var v34 int64 = 998
                v35 := GenS5{F0: 42, F1: v29}
                _ = fi33
                _ = v34
                _ = v35
            }
        default:
            for fi36 := 0; fi36 < 8; fi36++ {
                if ((4294967295 != 65535) || v1) {
                    v37 := GenS6{F0: 729, F1: 411}
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 1; fi38++ {
                        if v1 {
                            fmt.Println(fi36)
                            v3 = (v3 + ((("" + v3) + v3) + ("a" + v3)))
                            var v39 uint64 = 488
                            v5 -= v5
                            _ = v39
                        } else {
                            if (!false) {
                                var v40 int = (^(fi38))
                                fmt.Println(v5, v29, v2)
                                v5 = v5
                                _ = v40
                            }
                            var v41 int16 = 1
                            switch v3 {
                            case "a":
                                if len(v21) > 0 {
                                    v21[0] = (fi38 % ((834) | 1))
                                }
                            case "\n\t":
                                v3 = "hello"
                            }
                            v41 = (v41 >> ((uint(v5)) & 63))
                            _ = v41
                        }
                        v42 := GenS7{F0: fi36}
                        _ = fi38
                        _ = v42
                    }
                    if (v2 > (^(v2))) {
                        var v43 int8 = v2
                        if (!v1) {
                            var v44 bool = (2 < 156)
                            var v45 int8 = (v2 >> ((511) & 63))
                            for fi46 := 0; fi46 < 4; fi46++ {
                                var v47 uint8 = (^(v5))
                                if (!v44) {
                                    if (3.14159 < 1e10) {
                                        v48 := GenS6{F0: 0, F1: 1}
                                        _ = v48
                                    } else {
                                        v49 := GenS3{F0: 857, F1: 32767}
                                        delete(v4, v3)
                                        var v50 int = fi36
                                        v51 := GenS1{F0: ("hello" + (("hello" + "a") + (v3 + v3)))}
                                        var v52 uint = 0
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 uint8 = v47
                                    if (4294967295 <= 18446744073709551615) {
                                        v54 := v4[v3]
                                        var v55 int32 = 1
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v53
                                } else {
                                    if (70 == 2) {
                                        v56 := GenS8{F0: ((v44 || v1) || (1 > 0)), F1: v1, F2: (fi46 | fi36)}
                                        v1 = (2 <= 2147483648)
                                        _ = v56
                                    } else {
                                        var v57 int = fi36
                                        var v58 uint16 = 1
                                        v44 = (1 > 2147483647)
                                        _ = v57
                                        _ = v58
                                    }
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    var v59 int = -1
                                    switch 389 {
                                    case 796:
                                        if v1 {
                                            v60 := GenS2{F0: -32768, F1: 1, F2: v59}
                                            _ = v60
                                        } else {
                                            if v44 {
                                                v3 = v3
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                            } else {
                                                if v1 {
                                                    var v61 int32 = 2147483647
                                                    var v62 int32 = 0
                                                    v61 = (v62 - (v62 * 1))
                                                    _ = v61
                                                    _ = v62
                                                }
                                                v45 += int8(v29)
                                            }
                                            fmt.Println(v59)
                                            v63 := GenS6{F0: 1, F1: 1}
                                            _ = v63
                                        }
                                    case -835:
                                        if (v3 < "\n\t") {
                                            v64 := GenS3{F0: 2147483648, F1: -843}
                                            _ = v64
                                        }
                                    default:
                                        fmt.Println(v3, v47, v45)
                                        var v65 float64 = (3.14159 + (0.1 * float64(0.1)))
                                        v66 := GenS4{F0: (^(v45))}
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 uint = 233
                                    _ = v59
                                    _ = v67
                                }
                                fmt.Println(v43, v45)
                                _ = fi46
                                _ = v47
                            }
                            _ = v44
                            _ = v45
                        } else {
                            _ = len(v4)
                            for fi68 := 0; fi68 < 19; fi68++ {
                                fmt.Println(fi68, v2, fi68)
                                _ = fi68
                            }
                        }
                        v5 = (^(v5))
                        _ = v43
                    }
                    for fi69 := 0; fi69 < 20; fi69++ {
                        var v70 uint32 = 837
                        var v71 bool = (-9223372036854775808 <= -432)
                        v71 = ((!false) || (!(-475 != -9223372036854775808)))
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                }
                fmt.Println(v2, v2)
                _ = fi36
            }
        }
        _ = v31
    }
    if (0.0 >= (1.5 - 3.14159)) {
        for fi72 := 0; fi72 < 2; fi72++ {
            if (((v29 / ((v29) | 1)) < v29) || v1) {
                v73 := GenS8{F0: false, F1: (!v1), F2: fi72}
                switch v3 {
                case "":
                    fmt.Println(fi72, v1)
                case "\n\t":
                    v29 = 1
                }
                _ = v73
            }
            _ = fi72
        }
        var v74 int16 = -1
        _ = v74
    } else {
        v75 := GenS7{F0: -334}
        v2 += (-(v2))
        _ = v75
    }
    var v76 uint32 = 140
    var v77 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
}
