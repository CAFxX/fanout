package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 string }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 int; F2 uint32 }

func main() {
    var v1 int32 = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 int16 = -1
    if false {
        fmt.Println(v4, v1)
    } else {
        if ((128 >= 42) || (false || (true || true))) {
            v3[-2147483648] = 0
        } else {
            delete(v3, 2147483647)
            var v5 uint8 = 85
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := GenS1{F0: -128, F1: -1, F2: 573}
                var v8 uint8 = (v5 + (v5 % (((v5 - (v5 % ((v5) | 1)))) | 1)))
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint16 = 65535
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 113, F1: -698, F2: 32767}
        _ = v11
    }
    v12 := GenS1{F0: -60, F1: 1, F2: 271}
    v1 = (-(v1))
    var v13 string = "\n\t"
    v1 = -344
    if false {
        if false {
            switch -834 {
            case -9223372036854775808:
                var v14 int16 = (v4 ^ v4)
                _ = v14
            }
            var v15 uint8 = 1
            _ = v15
        } else {
            if false {
                switch v13 {
                case "hello":
                    var v16 string = v13
                    _ = v16
                case "a":
                    switch 32768 {
                    case 1:
                        var v17 bool = (16 <= -128)
                        _ = v17
                    default:
                        switch 1 {
                        case 4294967295:
                            fmt.Println(v1, v1, v4)
                        }
                    }
                case "\n\t":
                    switch -2147483648 {
                    case -528:
                        _ = len(v3)
                    case 898:
                        if ((127 <= 127) || true) {
                            for fi18 := 0; fi18 < 8; fi18++ {
                                switch fi18 {
                                case 1:
                                    for fi19 := 0; fi19 < 9; fi19++ {
                                        var v20 int = -725
                                        v20 = (-(fi19))
                                        v21 := v2[v13]
                                        _ = fi19
                                        _ = v20
                                        _ = v21
                                    }
                                case 4294967295:
                                    v22 := GenS1{F0: 42, F1: -1, F2: 636}
                                    _ = v22
                                }
                                v23 := []int{(fi18 / ((fi18) | 1)), 32768, (fi18 * fi18)}
                                var v24 uint = 65535
                                _ = fi18
                                _ = v23
                                _ = v24
                            }
                            v1 = 0
                        } else {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 uint8 = 127
                                for fi27 := 0; fi27 < 13; fi27++ {
                                    var v28 int32 = -384
                                    v29 := GenS2{F0: -128, F1: (fi27 + (^(fi25)))}
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                    default:
                        for fi30 := 0; fi30 < 16; fi30++ {
                            v31 := GenS3{F0: 451, F1: v1}
                            _ = fi30
                            _ = v31
                        }
                    }
                }
                v13 = ((v13 + v13) + v13)
            }
            v32 := []int{42}
            var v33 uint = 1
            _ = v32
            _ = v33
        }
        if (-1 < -9223372036854775808) {
            var v34 uint32 = 1
            _ = v34
        }
    }
    var v35 uint = 2
    var v36 int = 1
    var v37 float64 = (-2.25 + (float64(((3.14159 + 0.1) - 0.0)) / ((1e10 / 1.0) * float64((1e10 - -2.25)))))
    for fi38 := 0; fi38 < 18; fi38++ {
        fmt.Println(v35, fi38, v13)
        switch v13 {
        case "":
            v39 := GenS4{F0: (!(v13 <= "a")), F1: "a", F2: (v4 % ((v4) | 1))}
            _ = v39
        case "hello":
            v4 += -32768
        case "a":
            if false {
                var v40 int8 = 127
                var v41 bool = true
                v42 := GenS3{F0: 848, F1: (v1 / ((v1) | 1))}
                _ = v40
                _ = v41
                _ = v42
            } else {
                switch fi38 {
                case -1:
                    v1 -= v1
                case 274:
                    v43 := GenS5{F0: 0}
                    _ = v43
                case 0:
                    switch -9223372036854775808 {
                    case 4294967295:
                        v44 := GenS1{F0: -31, F1: 1, F2: 1}
                        _ = v44
                    default:
                        delete(v2, v13)
                    }
                default:
                    if true {
                        var v45 int = (v36 >> ((v35) & 63))
                        if (1 == 65535) {
                            var v46 uint32 = 0
                            switch v13 {
                            case "":
                                v1 += 2147483647
                            case "a":
                                v36 = (-(v45))
                            case "foo bar":
                                v35 -= v35
                            }
                            v47 := GenS3{F0: 612, F1: (^(v1))}
                            _ = v46
                            _ = v47
                        }
                        _ = v45
                    }
                }
                v48 := GenS6{F0: 32767, F1: 2, F2: v13}
                _ = v48
            }
        }
        _ = fi38
    }
    v13 = ""
    for fi49 := 0; fi49 < 8; fi49++ {
        v13 = v13
        var v50 bool = (!(v13 <= "a"))
        _ = fi49
        _ = v50
    }
    v1 = (v1 << (((v35 / (((v35 & v35)) | 1))) & 63))
    if false {
        v3[v36] = 42
        for fi51 := 0; fi51 < 17; fi51++ {
            var v52 float32 = (-2.25 + 1e10)
            for fi53 := 0; fi53 < 8; fi53++ {
                v54 := []int{}
                _ = len(v54)
                v55 := GenS7{F0: 4294967295, F1: (79 == 42)}
                v56 := GenS8{F0: (fi51 & fi53), F1: (-21 > 46)}
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            _ = fi51
            _ = v52
        }
        if ((v36 & (^(v36))) != (v36 / ((-2147483648) | 1))) {
            for fi57 := 0; fi57 < 10; fi57++ {
                switch (v13 + (v13 + v13)) {
                case "hello":
                    switch v13 {
                    case "hello":
                        v58 := GenS5{F0: (v4 ^ (v4 % ((v4) | 1)))}
                        _ = v58
                    case "foo bar":
                        v37 += ((v37 * float64(0.1)) - v37)
                    case "\n\t":
                        if false {
                            v59 := GenS9{F0: int64(fi57)}
                            v35 = 922
                            v60 := []int{4294967295, v36}
                            _ = v59
                            _ = v60
                        }
                    }
                default:
                    if (false && (-1 < -749)) {
                        var v61 uint8 = 2
                        var v62 int = 339
                        _ = v61
                        _ = v62
                    } else {
                        v63 := GenS3{F0: 943, F1: v1}
                        _ = v63
                    }
                }
                _ = fi57
            }
            for fi64 := 0; fi64 < 5; fi64++ {
                if (974 <= 3) {
                    switch fi64 {
                    case 127:
                        var v65 uint = v35
                        _ = v65
                    case -279:
                        v66 := []int{}
                        _ = v66
                    }
                    v67 := GenS7{F0: 18446744073709551615, F1: (3.14159 >= 0.1)}
                    if false {
                        for fi68 := 0; fi68 < 15; fi68++ {
                            fmt.Println(v1, fi68, v36)
                            var v69 int16 = v4
                            v70 := GenS10{F0: -45, F1: fi64, F2: 2147483648}
                            fmt.Println(v4, v13, v4)
                            if true {
                                v35 = (v35 >> (((v35 ^ (v35 - v35))) & 63))
                                v71 := []int{v36, fi68}
                                v13 = v13
                                _ = v71
                            } else {
                                _ = len(v3)
                            }
                            if true {
                                v1 = v1
                                fmt.Println(v13, fi68)
                            } else {
                                var v72 int16 = (v4 << ((370) & 63))
                                _ = v72
                            }
                            switch int(v35) {
                            case 9223372036854775807:
                                v1 = -551
                            }
                            _ = fi68
                            _ = v69
                            _ = v70
                        }
                        for fi73 := 0; fi73 < 4; fi73++ {
                            if false {
                                v13 = "a"
                                var v74 string = "foo bar"
                                delete(v2, "foo bar")
                                _ = v74
                            } else {
                                var v75 uint8 = 255
                                v76 := []int{(fi73 >> ((v35) & 63)), -2147483648}
                                v37 = v37
                                _ = v75
                                _ = v76
                            }
                            _ = fi73
                        }
                    } else {
                        var v77 string = (((v13 + "hello") + ((v13 + v13) + (v13 + (v13 + ((("a" + v13) + v13) + v13))))) + "")
                        for fi78 := 0; fi78 < 5; fi78++ {
                            var v79 int = (-(fi64))
                            var v80 uint = v35
                            v3[(fi78 % (((v36 / (((fi78 * -2147483648)) | 1))) | 1))] = (fi64 % (((v36 & (fi78 ^ fi64))) | 1))
                            fmt.Println(v4, v13)
                            v81 := []int{(fi78 & (-(v36)))}
                            var v82 float64 = v37
                            _ = fi78
                            _ = v79
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                        _ = v77
                    }
                    _ = v67
                }
                _ = fi64
            }
        }
    }
    var v83 bool = (uint32(v35) >= 0)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v83
}
