package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 uint16; F2 int }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int8; F1 int16; F2 string }
type GenS11 struct { F0 int16 }

func main() {
    var v1 int64 = -773
    var v2 uint64 = 826
    var v3 uint32 = 1
    var v4 int32 = 42
    var v5 string = ("hello" + ("a" + ("foo bar" + "\n\t")))
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if (-250 > 0) {
                fmt.Println(v3)
            } else {
                switch (v5 + "foo bar") {
                case "\n\t":
                    v4 += v4
                case "hello":
                    if false {
                        v2 += 204
                        fmt.Println(v5, v3)
                        var v9 float64 = 1.5
                        _ = v9
                    }
                }
                v2 = (v2 & 18446744073709551615)
                v1 = int64(v2)
            }
            v10 := GenS1{F0: v6}
            _ = fi8
            _ = v10
        }
        v1 += 197
        v11 := GenS2{F0: 797, F1: (v4 % ((v4) | 1)), F2: (v2 % ((v2) | 1))}
        _ = fi7
        _ = v11
    }
    var v12 uint = 718
    var v13 string = ("hello" + "")
    if true {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 int64 = (v1 + v1)
            v16 := []int{-2147483648, 32768, (fi14 ^ 4294967295)}
            if true {
                if true {
                    v17 := GenS3{F0: (^(v3)), F1: 1}
                    for fi18 := 0; fi18 < 16; fi18++ {
                        switch (v5 + (((v13 + ((v5 + "hello") + v5)) + ((v5 + v5) + "\n\t")) + (v5 + (v5 + "a")))) {
                        case "a":
                            v2 += v2
                        case "hello":
                            if true {
                                _ = len(v16)
                                fmt.Println(v15, v5)
                                var v19 int64 = v1
                                _ = v19
                            } else {
                                var v20 string = v13
                                _ = v20
                            }
                        case "":
                            switch v5 {
                            case "foo bar":
                                v2 = 42
                                if len(v16) > 0 {
                                    v16[0] = fi18
                                }
                            case "":
                                _ = len(v16)
                            }
                        default:
                            var v21 bool = ((42 > (v2 << ((v12) & 63))) || false)
                            _ = v21
                        }
                        switch -990 {
                        case 128:
                            _ = len(v16)
                        default:
                            var v22 int64 = v1
                            _ = v22
                        }
                        _ = fi18
                    }
                    var v23 bool = (-374 <= 0)
                    _ = v17
                    _ = v23
                } else {
                    var v24 int32 = (v4 | v4)
                    for fi25 := 0; fi25 < 20; fi25++ {
                        fmt.Println(v12, v3, fi25)
                        v12 -= v12
                        if (true && (!(128 <= 190))) {
                            v3 = (v3 << ((18446744073709551615) & 63))
                        } else {
                            v1 -= v15
                            fmt.Println(v3, fi14, v4)
                        }
                        v26 := GenS4{F0: v15, F1: 127}
                        var v27 uint64 = 0
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v4 += int32(v2)
                    _ = v24
                }
                if len(v16) > 0 {
                    v16[0] = fi14
                }
                var v28 uint32 = (v3 + 42)
                _ = v28
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v29 int = -300
        _ = v29
    }
    var v30 int = 120
    var v31 float32 = (float32(-2.25) * float32(((3.14159 - (((((3.14159 - (0.0 - 1e10)) / ((-2.25 + (1.5 - float64(1.5))) * 1e10)) / (0.1 / 3.14159)) - ((1e10 / -2.25) / 1.5)) + 0.1)) / float64(3.14159))))
    switch (v30 / ((v30) | 1)) {
    case 4294967295:
        switch v30 {
        case 2147483647:
            if true {
                if true {
                    v1 = v1
                    var v32 uint = 18446744073709551615
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v12 += (v32 * (v32 | 4294967295))
                        v34 := []int{-2147483648, 4294967295, fi33}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
                var v35 int = (v30 - int(v1))
                _ = v35
            }
        }
    case 1:
        v3 = v3
    case -446:
        for fi36 := 0; fi36 < 18; fi36++ {
            if false {
                switch "" {
                case "hello":
                    var v37 int64 = (v1 + 0)
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 15; fi38++ {
                        if true {
                            fmt.Println(v1, v6)
                        } else {
                            v39 := GenS5{F0: (^(v2))}
                            _ = v39
                        }
                        v40 := []int{-586, fi38}
                        v6 += v6
                        _ = fi38
                        _ = v40
                    }
                }
            }
            if true {
                var v41 bool = (-2147483648 == v30)
                _ = v41
            } else {
                var v42 int = 796
                for fi43 := 0; fi43 < 4; fi43++ {
                    v44 := []int{v30, (v30 ^ (fi43 + (fi43 % ((fi36) | 1))))}
                    _ = fi43
                    _ = v44
                }
                var v45 int32 = v4
                _ = v42
                _ = v45
            }
            _ = fi36
        }
    default:
        if (!(!(65535 <= (v12 % (((v12 + 2)) | 1))))) {
            var v46 uint8 = 127
            v30 = (v30 ^ -1)
            _ = v46
        } else {
            v13 = v5
        }
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        switch v30 {
        case -138:
            v5 = v5
        case 9223372036854775807:
            if (!true) {
                v48 := GenS1{F0: v6}
                _ = v48
            }
        case -625:
            var v49 int32 = 937
            _ = v49
        }
        if false {
            if false {
                v5 = (v5 + ((v5 + "") + ("" + v5)))
                v50 := GenS5{F0: (v2 % ((2) | 1))}
                v31 = v31
                _ = v50
            }
            switch "" {
            case "":
                if false {
                    var v51 int8 = -1
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 9; fi52++ {
                        var v53 int32 = v4
                        _ = fi52
                        _ = v53
                    }
                }
            case "\n\t":
                if true {
                    v54 := GenS4{F0: (v1 << ((v12) & 63)), F1: 214}
                    v55 := []int{(v30 * int(v1)), (v30 << ((v12) & 63)), 32768}
                    _ = v54
                    _ = v55
                } else {
                    v56 := GenS6{F0: (v5 + ("\n\t" + "a"))}
                    _ = v56
                }
            default:
                v57 := []int{int(v12), 127}
                _ = v57
            }
        }
        var v58 int = (v30 / ((v30) | 1))
        _ = fi47
        _ = v58
    }
    var v59 uint = v12
    var v60 uint64 = (^(v2))
    switch v30 {
    case -2147483648:
        v61 := []int{v30}
        _ = v61
    case 392:
        for fi62 := 0; fi62 < 3; fi62++ {
            var v63 int8 = -97
            for fi64 := 0; fi64 < 9; fi64++ {
                switch (v13 + v5) {
                case "a":
                    switch (fi64 << (((v59 | v12)) & 63)) {
                    case 9223372036854775807:
                        v65 := []int{}
                        _ = v65
                    case -587:
                        if false {
                            v66 := []int{(fi64 << ((uint(v6)) & 63)), (-(v30)), int(v4), v30}
                            v66 = append(v66, (^(fi62)))
                            v67 := GenS7{F0: true}
                            _ = v66
                            _ = v67
                        }
                    case -344:
                        v68 := []int{(fi64 + (v30 & -2147483648)), fi64, (^(fi64)), (fi64 * fi62)}
                        _ = v68
                    default:
                        v31 += 0.1
                    }
                case "\n\t":
                    v30 = (v30 * -837)
                case "foo bar":
                    v69 := GenS8{F0: (v60 >> (((^(v12))) & 63)), F1: 2, F2: (fi62 & fi62)}
                    _ = v69
                }
                _ = fi64
            }
            fmt.Println(v60)
            _ = fi62
            _ = v63
        }
    }
    for fi70 := 0; fi70 < 20; fi70++ {
        var v71 uint64 = v2
        _ = fi70
        _ = v71
    }
    if true {
        switch -500 {
        case 42:
            for fi72 := 0; fi72 < 16; fi72++ {
                v30 = 42
                _ = fi72
            }
        case 32767:
            v73 := GenS9{F0: -1, F1: v30}
            _ = v73
        case 2147483647:
            v74 := GenS5{F0: (v2 << ((v12) & 63))}
            _ = v74
        default:
            var v75 uint64 = (v60 >> (((v12 >> (((^(v12))) & 63))) & 63))
            _ = v75
        }
        switch 32768 {
        case -1:
            v76 := []int{}
            _ = v76
        case 450:
            var v77 uint16 = v6
            _ = v77
        default:
            v78 := GenS10{F0: 100, F1: int16(v1), F2: "a"}
            _ = v78
        }
    } else {
        for fi79 := 0; fi79 < 14; fi79++ {
            switch "foo bar" {
            case "foo bar":
                switch ("hello" + v5) {
                case "":
                    if ((v4 * v4) == v4) {
                        switch 128 {
                        case -9223372036854775808:
                            if ((uint16(v59) < (^(v6))) || (!(((true || ((!false) || false)) || (!true)) && (v4 != (v4 << (((v12 / ((755) | 1))) & 63)))))) {
                                v13 = (v5 + v13)
                                v80 := []int{}
                                v81 := GenS2{F0: (v6 | (v6 * (v6 << (((v12 ^ uint(v3))) & 63)))), F1: 748, F2: (v60 * uint64(v3))}
                                _ = v80
                                _ = v81
                            } else {
                                var v82 string = v5
                                if (v31 >= v31) {
                                    v2 -= uint64(v59)
                                    v4 = (v4 % ((v4) | 1))
                                    switch v82 {
                                    case "\n\t":
                                        var v83 int32 = v4
                                        _ = v83
                                    case "foo bar":
                                        v2 -= v60
                                    case "a":
                                        var v84 int8 = 100
                                        _ = v84
                                    default:
                                        v3 -= 183
                                    }
                                } else {
                                    v2 -= (v2 - (^(v60)))
                                    var v85 int16 = -32768
                                    _ = v85
                                }
                                _ = v82
                            }
                        case -1:
                            if (!false) {
                                if false {
                                    var v86 float32 = 3.14159
                                    _ = v86
                                }
                                switch 0 {
                                case 4294967295:
                                    v59 = (v12 & (^(v59)))
                                case -1:
                                    if true {
                                        v1 = v1
                                        v87 := []int{}
                                        v87 = append(v87, (v30 + (fi79 ^ 32768)))
                                        _ = v87
                                    }
                                }
                            } else {
                                var v88 uint64 = v60
                                fmt.Println(v1, v5, v2)
                                v89 := []int{(v30 | (v30 | (v30 % ((v30) | 1))))}
                                _ = v88
                                _ = v89
                            }
                        case 128:
                            if false {
                                v90 := []int{(fi79 - 9223372036854775807), 2147483647, v30, (fi79 >> (((^(v59))) & 63))}
                                _ = v90
                            }
                        default:
                            v1 += v1
                        }
                        for fi91 := 0; fi91 < 9; fi91++ {
                            v60 += (^(v60))
                            var v92 uint32 = 42
                            _ = fi91
                            _ = v92
                        }
                    }
                default:
                    v13 = v5
                }
            case "\n\t":
                for fi93 := 0; fi93 < 9; fi93++ {
                    v94 := GenS7{F0: false}
                    if (false || true) {
                        switch (-(v30)) {
                        case 127:
                            v95 := GenS11{F0: 32767}
                            _ = v95
                        case 32767:
                            v96 := []int{}
                            _ = v96
                        case 444:
                            v97 := []int{(fi93 - int(v2)), (v30 / (((fi79 >> ((v59) & 63))) | 1)), v30}
                            _ = v97
                        }
                        var v98 uint16 = 535
                        v31 = ((0.1 - (0.1 * v31)) * (((float32((0.0 - -2.25)) + v31) * v31) - v31))
                        _ = v98
                    } else {
                        v99 := GenS10{F0: -66, F1: 0, F2: (v5 + ((v5 + "hello") + v5))}
                        v100 := GenS5{F0: v2}
                        if (!(true && true)) {
                            var v101 float64 = -2.25
                            var v102 int = (v30 | -840)
                            var v103 int32 = v4
                            _ = v101
                            _ = v102
                            _ = v103
                        } else {
                            v2 -= (v2 | v60)
                            v104 := GenS7{F0: false}
                            v105 := []int{(fi93 | (^(fi93))), fi93, (^(v30))}
                            _ = v104
                            _ = v105
                        }
                        _ = v100
                        _ = v99
                    }
                    v30 -= (^(v30))
                    _ = fi93
                    _ = v94
                }
            }
            var v106 int = int(v3)
            var v107 uint16 = 376
            _ = fi79
            _ = v106
            _ = v107
        }
        v108 := []int{}
        fmt.Println(v31, v4)
        _ = v108
    }
    var v109 int = int(v3)
    var v110 float64 = (3.14159 - ((1e10 / 0.1) / 1e10))
    var v111 uint16 = (v6 << ((v59) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
