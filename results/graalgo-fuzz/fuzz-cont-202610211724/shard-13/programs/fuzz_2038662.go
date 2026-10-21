package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }
func genFn0() uint {
    if false {
        v5 := []int{-427}
        var v6 int = -192
        var v7 uint16 = 2
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 int16 = 33
        for fi9 := 0; fi9 < 15; fi9++ {
            switch (^(fi9)) {
            case 9223372036854775807:
                var v10 uint = 554
                _ = v10
            default:
                v11 := []int{fi9}
                _ = v11
            }
            if true {
                v12 := []int{}
                var v13 float64 = float64(float32(1.5))
                v13 += (v13 - v13)
                _ = v12
                _ = v13
            }
            var v14 uint16 = 2
            _ = fi9
            _ = v14
        }
        v15 := []int{272, 818, 1, 42}
        _ = v15
        _ = v8
    }
    v16 := []int{}
    var v17 uint8 = 127
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 float32 = 3.14159
        fmt.Println(v17, v19, v19)
        var v20 string = ""
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v16
    _ = v17
    return 255
}
func genFn1(p0 string) uint8 {
    v21 := GenS1{F0: 0, F1: "foo bar"}
    var v22 int64 = -339
    var v23 int64 = (v22 / ((1) | 1))
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (!false)
    v3 := make(map[uint32]int)
    var v4 int32 = -1
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int16 = 0
        switch "" {
        case "hello":
            var v26 uint8 = 1
            _ = v26
        default:
            v4 += (v4 | 361)
        }
        fmt.Println(v25, v2, v25)
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "hello":
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS1{F0: 0, F1: "a"}
            _ = fi27
            _ = v28
        }
    case "\n\t":
        v29 := GenS2{F0: -40, F1: 0}
        _ = v29
    }
    if (30 >= 293) {
        v30 := []int{2147483647, 128}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 1; fi31++ {
            if (!v2) {
                fmt.Println(v2, v4)
                v32 := GenS1{F0: genFn0(), F1: ("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "\n\t"))))}
                _ = v32
            }
            for fi33 := 0; fi33 < 15; fi33++ {
                if (1 >= 127) {
                    var v34 float32 = 0.1
                    if v2 {
                        var v35 int16 = 958
                        v1 -= (v1 ^ (v1 | 18446744073709551615))
                        v1 = v1
                        _ = v35
                    }
                    _ = v34
                }
                _ = fi33
            }
            _ = fi31
        }
    }
    var v36 string = "hello"
    var v37 int8 = 127
    var v38 string = ((v36 + v36) + v36)
    for fi39 := 0; fi39 < 13; fi39++ {
        v40 := GenS3{F0: 1, F1: 65535, F2: 4294967295}
        if (565 <= 9223372036854775807) {
            v41 := GenS4{F0: 65535, F1: 204, F2: 255}
            _ = v41
        } else {
            var v42 float64 = 0.1
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    var v43 string = ""
    switch 127 {
    case 465:
        fmt.Println(v37, v1)
    case 2147483647:
        var v44 uint8 = 127
        _ = v44
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        var v46 int16 = 465
        for fi47 := 0; fi47 < 5; fi47++ {
            if v2 {
                switch v38 {
                case "foo bar":
                    var v48 int = fi45
                    _ = v48
                case "hello":
                    if (!v2) {
                        switch "" {
                        case "a":
                            v49 := []int{(^(fi47)), -863, fi45, 874}
                            _ = v49
                        case "\n\t":
                            v1 = 901
                        }
                    }
                case "\n\t":
                    if false {
                        if (true && v2) {
                            if v2 {
                                if v2 {
                                    var v50 int16 = v46
                                    var v51 uint32 = 2
                                    switch v43 {
                                    case "foo bar":
                                        if ((v4 << ((genFn0()) & 63)) == (v4 | (v4 | v4))) {
                                            fmt.Println(v4)
                                        } else {
                                            var v52 uint16 = 962
                                            _ = v52
                                        }
                                    case "hello":
                                        v36 = "hello"
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        var v53 float32 = 1.5
                                        _ = v53
                                    case "foo bar":
                                        v54 := v3[4294967295]
                                        _ = v54
                                    }
                                    _ = v50
                                    _ = v51
                                }
                            } else {
                                if ((v4 <= v4) || false) {
                                    fmt.Println(v2)
                                    if v2 {
                                        v36 = v38
                                        if (!(false && (255 <= 83))) {
                                            var v55 int32 = v4
                                            _ = v55
                                        }
                                    }
                                    fmt.Println(fi45)
                                } else {
                                    var v56 float32 = 1e10
                                    _ = v56
                                }
                            }
                            switch v43 {
                            case "":
                                fmt.Println(fi45, fi45)
                            }
                            var v57 uint32 = 0
                            fmt.Println(v2, v1)
                            var v58 uint64 = v1
                            _ = v57
                            _ = v58
                        }
                        v4 += (v4 | v4)
                        v38 = "foo bar"
                    } else {
                        v3[2147483648] = (-(fi47))
                        v38 = ""
                    }
                default:
                    if (fi47 == fi45) {
                        v1 = v1
                    }
                }
            } else {
                var v59 int = fi47
                var v60 uint8 = 1
                if v2 {
                    fmt.Println(v2)
                    v61 := []int{(v59 - (fi45 + (-(fi47)))), (v59 >> ((18446744073709551615) & 63)), v59, (fi45 ^ fi47)}
                    v43 = v36
                    _ = v61
                } else {
                    if v2 {
                        v60 = v60
                    }
                }
                _ = v59
                _ = v60
            }
            v1 = (v1 / ((2) | 1))
            _ = fi47
        }
        for fi62 := 0; fi62 < 16; fi62++ {
            if true {
                switch (fi62 / (((fi62 / ((fi62) | 1))) | 1)) {
                case 127:
                    v36 = "a"
                case 32767:
                    var v63 uint32 = 878
                    _ = v63
                }
            } else {
                var v64 int8 = (v37 * (-(v37)))
                v65 := GenS4{F0: 255, F1: 494, F2: genFn0()}
                v3[2147483648] = 32768
                _ = v64
                _ = v65
            }
            var v66 uint64 = (v1 | (v1 + (v1 / ((v1) | 1))))
            v1 -= (v66 % (((v66 | 0)) | 1))
            _ = fi62
            _ = v66
        }
        _ = fi45
        _ = v46
    }
    if false {
        var v67 uint8 = 200
        _ = v67
    }
    if v2 {
        switch ((((((v43 + ("a" + v43)) + ("" + v43)) + (v43 + (v36 + v36))) + "\n\t") + v38) + (v38 + "hello")) {
        case "hello":
            switch (v36 + v36) {
            case "hello":
                var v68 float32 = 0.0
                _ = v68
            case "a":
                var v69 uint64 = v1
                _ = v69
            default:
                v43 = v36
            }
        case "a":
            var v70 int32 = v4
            _ = v70
        }
    }
    if v2 {
        v71 := []int{}
        v36 = ""
        switch 516 {
        case 32767:
            if (genFn1(v38) < uint8(v1)) {
                var v72 uint64 = v1
                for fi73 := 0; fi73 < 5; fi73++ {
                    if (((v72 - 42) <= 2) || v2) {
                        if false {
                            delete(v3, 2147483647)
                        }
                        for fi74 := 0; fi74 < 4; fi74++ {
                            if v2 {
                                v75 := GenS4{F0: 346, F1: (v4 + 801), F2: 573}
                                _ = v75
                            }
                            _ = fi74
                        }
                        if true {
                            fmt.Println(fi73, v1)
                        }
                    } else {
                        if v2 {
                            v76 := GenS2{F0: -1, F1: 942}
                            switch (((v38 + v38) + v36) + (v43 + ((("a" + ("\n\t" + ((((v38 + (v36 + "foo bar")) + ("hello" + ((v36 + ("hello" + v38)) + (("\n\t" + "hello") + "\n\t")))) + v43) + (((v38 + "foo bar") + v36) + v38)))) + (("a" + "foo bar") + ("" + "foo bar"))) + "hello"))) {
                            case "a":
                                v43 = ("\n\t" + "a")
                            case "foo bar":
                                v1 = (^(v1))
                            default:
                                var v77 int64 = -1
                                _ = v77
                            }
                            fmt.Println(v1, v36, fi73)
                            _ = v76
                        }
                    }
                    _ = fi73
                }
                switch -97 {
                case -523:
                    var v78 uint64 = 42
                    _ = v78
                }
                _ = v72
            } else {
                switch "foo bar" {
                case "":
                    var v79 int = -697
                    _ = v79
                }
                v3[0] = -2147483648
            }
        case 1:
            for fi80 := 0; fi80 < 3; fi80++ {
                if v2 {
                    var v81 int = 146
                    _ = v81
                }
                var v82 uint = 4294967295
                _ = fi80
                _ = v82
            }
        case -222:
            if v2 {
                v83 := GenS4{F0: 32767, F1: (-(v4)), F2: 824}
                switch v43 {
                case "foo bar":
                    var v84 uint = 60
                    _ = v84
                case "":
                    if (!(-392 >= 42)) {
                        fmt.Println(v2, v4)
                    }
                default:
                    v4 = 2147483647
                }
                var v85 uint64 = (^(v1))
                _ = v83
                _ = v85
            } else {
                switch 42 {
                case -1:
                    v71 = append(v71, 127)
                case -755:
                    v38 = "a"
                }
                var v86 uint = uint(v1)
                v37 = (v37 + v37)
                _ = v86
            }
        default:
            switch -657 {
            case 4294967295:
                fmt.Println(v37, v4, v1)
            case 819:
                for fi87 := 0; fi87 < 18; fi87++ {
                    var v88 string = (v43 + (v38 + v36))
                    _ = fi87
                    _ = v88
                }
            case 32767:
                v43 = v36
            default:
                var v89 uint16 = 535
                _ = v89
            }
        }
        _ = v71
    }
    v36 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
}
