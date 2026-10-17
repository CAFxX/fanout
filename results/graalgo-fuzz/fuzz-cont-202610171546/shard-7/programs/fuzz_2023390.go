package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 uint }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 int8 = 0
    var v3 bool = (65535 <= 4294967295)
    v4 := make(map[uint32]uint32)
    var v5 int64 = -1
    var v6 uint64 = 2
    switch v1 {
    case "foo bar":
        switch -2147483648 {
        case 128:
            v7 := GenS1{F0: 0, F1: 42}
            _ = v7
        case -2147483648:
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 uint64 = 143
                _ = fi8
                _ = v9
            }
        }
    case "":
        var v10 float32 = (float32(3.14159) / 3.14159)
        _ = v10
    }
    var v11 string = "\n\t"
    switch -751 {
    case 0:
        var v12 int64 = 406
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch "hello" {
        case "hello":
            for fi14 := 0; fi14 < 9; fi14++ {
                if v3 {
                    v15 := GenS1{F0: -32768, F1: 718}
                    v16 := []int{(fi13 | (fi14 / ((-2147483648) | 1)))}
                    if (true || false) {
                        v5 -= v5
                        var v17 string = ((("hello" + ("foo bar" + v11)) + v11) + ((v11 + v11) + v1))
                        _ = v17
                    } else {
                        if true {
                            if false {
                                v5 += -1
                                delete(v4, 867)
                                var v18 int32 = -1
                                _ = v18
                            } else {
                                fmt.Println(v6, fi13)
                                _ = len(v16)
                            }
                            fmt.Println(v3, v1, fi13)
                        } else {
                            var v19 uint32 = 6
                            v20 := GenS1{F0: -1, F1: 0}
                            var v21 uint8 = 92
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        delete(v4, 2147483648)
                        if (v3 && false) {
                            if v3 {
                                if false {
                                    if v3 {
                                        v1 = "a"
                                    } else {
                                        v2 = 100
                                    }
                                } else {
                                    var v22 uint16 = 32767
                                    _ = v22
                                }
                            } else {
                                v5 += int64(v6)
                                v1 = "\n\t"
                                v3 = v3
                            }
                            v11 = (v1 + v11)
                        }
                    }
                    var v23 float32 = 1.5
                    v24 := GenS1{F0: 42, F1: -454}
                    _ = v15
                    _ = v16
                    _ = v23
                    _ = v24
                }
                v5 = 0
                _ = fi14
            }
        case "":
            var v25 int = (fi13 % ((-416) | 1))
            _ = v25
        case "foo bar":
            if (42 == 128) {
                var v26 uint16 = 682
                var v27 int16 = 0
                v28 := GenS1{F0: (v27 >> ((uint(v6)) & 63)), F1: (v27 + v27)}
                _ = v26
                _ = v27
                _ = v28
            }
        default:
            v29 := []int{}
            _ = v29
        }
        v30 := GenS1{F0: 42, F1: -1}
        _ = fi13
        _ = v30
    }
    if (!true) {
        v5 -= (v5 >> ((227) & 63))
        fmt.Println(v2, v5)
        var v31 int8 = (v2 - v2)
        _ = v31
    }
    v32 := v4[65535]
    var v33 int = 32768
    v34 := GenS2{F0: (-(v2)), F1: ((v5 + (v5 + (-(v5)))) >= v5), F2: 65535}
    if v3 {
        if false {
            delete(v4, v32)
        } else {
            if (1 <= 854) {
                switch v33 {
                case 477:
                    if (float32((-2.25 - (1.5 * float64(0.0)))) == 1.5) {
                        var v35 int16 = -58
                        v36 := GenS1{F0: 32767, F1: v35}
                        for fi37 := 0; fi37 < 20; fi37++ {
                            if (false || true) {
                                if true {
                                    var v38 uint64 = (v6 % ((v6) | 1))
                                    fmt.Println(v32, v6, v35)
                                    _ = v38
                                } else {
                                    for fi39 := 0; fi39 < 11; fi39++ {
                                        var v40 uint16 = 831
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                if false {
                                    fmt.Println(v32)
                                    v41 := GenS1{F0: v35, F1: (v35 >> ((42) & 63))}
                                    v42 := []int{fi37, v33}
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 string = v1
                                    for fi44 := 0; fi44 < 18; fi44++ {
                                        if v3 {
                                            var v45 string = ((v43 + (v11 + v1)) + ((((("" + (v1 + ((v43 + v11) + (("hello" + v43) + ("\n\t" + v43))))) + v1) + "\n\t") + (v11 + "")) + v11))
                                            fmt.Println(v35, v45)
                                            fmt.Println(v5, v5, fi37)
                                            _ = v45
                                        }
                                        if v3 {
                                            v46 := GenS1{F0: (v35 % (((v35 % ((42) | 1))) | 1)), F1: 1000}
                                            v35 += 791
                                            var v47 uint = 979
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v11 = v1
                                        }
                                        switch -1 {
                                        case 2147483647:
                                            v32 = 1
                                        }
                                        _ = fi44
                                    }
                                    switch (v1 + ((v1 + v1) + "a")) {
                                    case "":
                                        if false {
                                            v35 = v35
                                            switch ((v43 + v1) + v43) {
                                            case "\n\t":
                                                switch fi37 {
                                                case 0:
                                                    if ((v6 - (v6 & v6)) == v6) {
                                                        var v48 uint32 = v32
                                                        _ = v48
                                                    }
                                                case 9223372036854775807:
                                                    fmt.Println(v11, v35)
                                                }
                                            case "foo bar":
                                                v49 := GenS2{F0: v2, F1: false, F2: 650}
                                                _ = v49
                                            case "":
                                                var v50 bool = (32767 > 42)
                                                _ = v50
                                            }
                                        } else {
                                            fmt.Println(v2)
                                            var v51 uint16 = 2
                                            v52 := v4[v32]
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    _ = v43
                                }
                                var v53 uint64 = v6
                                _ = v53
                            }
                            _ = fi37
                        }
                        _ = v35
                        _ = v36
                    } else {
                        var v54 uint64 = v6
                        var v55 uint = 40
                        _ = v54
                        _ = v55
                    }
                case 0:
                    v33 = 658
                case -9223372036854775808:
                    fmt.Println(v33, v32)
                default:
                    v5 += 0
                }
                var v56 int64 = (v5 << ((1) & 63))
                _ = v56
            }
            var v57 int = v33
            _ = v57
        }
        for fi58 := 0; fi58 < 3; fi58++ {
            var v59 float32 = 3.14159
            fmt.Println(v3)
            v60 := []int{int(v5), 9223372036854775807}
            _ = fi58
            _ = v59
            _ = v60
        }
        v33 += (v33 & v33)
    } else {
        if ((!(((((1.5 * ((0.0 + -2.25) / 3.14159)) * 0.1) == 0.1) || ((646 < 65535) && (v33 > 128))) && true)) && (!(1 == 428))) {
            for fi61 := 0; fi61 < 20; fi61++ {
                var v62 uint16 = 2
                v63 := GenS2{F0: v2, F1: v3, F2: 1}
                for fi64 := 0; fi64 < 13; fi64++ {
                    var v65 int32 = -1
                    switch (fi64 | fi61) {
                    case -623:
                        switch 128 {
                        case 4294967295:
                            delete(v4, v32)
                        default:
                            v62 += (^(v62))
                        }
                    case 127:
                        var v66 int16 = -1
                        _ = v66
                    case 2147483647:
                        var v67 uint16 = (v62 / ((0) | 1))
                        _ = v67
                    }
                    v6 -= v6
                    fmt.Println(v65)
                    _ = fi64
                    _ = v65
                }
                _ = fi61
                _ = v62
                _ = v63
            }
            v68 := v4[v32]
            _ = v68
        }
        var v69 bool = v3
        _ = v69
    }
    v70 := GenS2{F0: (v2 >> ((795) & 63)), F1: (1 <= 1), F2: 2}
    var v71 bool = (!(v33 >= v33))
    v72 := GenS2{F0: v2, F1: v71, F2: uint(v6)}
    if (32767 > v33) {
        switch 128 {
        case 192:
            v73 := []int{(v33 & (v33 ^ 42))}
            _ = v73
        case 32767:
            fmt.Println(v11, v3)
        case 2147483647:
            v74 := []int{v33, (v33 / ((-2147483648) | 1))}
            _ = v74
        }
        for fi75 := 0; fi75 < 1; fi75++ {
            switch fi75 {
            case -2147483648:
                var v76 int16 = 0
                _ = v76
            }
            _ = fi75
        }
        v5 = 42
    } else {
        var v77 uint16 = 0
        v6 += v6
        if v3 {
            v78 := v4[(v32 - (v32 | v32))]
            if v71 {
                var v79 bool = ((-(v2)) > (^(v2)))
                _ = v79
            }
            _ = v78
        } else {
            fmt.Println(v2, v11)
            switch (v33 % ((v33) | 1)) {
            case 1:
                fmt.Println(v3, v11)
            case 0:
                var v80 float64 = 3.14159
                _ = v80
            default:
                v81 := GenS3{F0: 255, F1: -592}
                _ = v81
            }
        }
        _ = v77
    }
    switch "foo bar" {
    case "hello":
        v82 := GenS1{F0: 0, F1: -32768}
        _ = v82
    }
    if v71 {
        var v83 int16 = 1000
        if (v32 >= (v32 / ((uint32(v6)) | 1))) {
            v4[(v32 | (v32 / ((v32) | 1)))] = (v32 ^ (^(v32)))
            var v84 uint64 = 775
            _ = v84
        }
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
}
