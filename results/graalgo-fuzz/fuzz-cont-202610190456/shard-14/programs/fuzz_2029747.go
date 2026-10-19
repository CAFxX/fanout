package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{fi5, 128, -2147483648}
        fmt.Println(v4, v3)
        _ = fi5
        _ = v6
    }
    switch "" {
    case "a":
        if (!true) {
            v7 := []int{-2147483648, -9223372036854775808, 32768}
            var v8 bool = (18446744073709551615 > 156)
            _ = v7
            _ = v8
        } else {
            var v9 float32 = 1e10
            _ = v9
        }
    default:
        var v10 int8 = -128
        _ = v10
    }
    var v11 int32 = 42
    v11 -= v11
    if false {
        var v12 int64 = 1
        v3 -= (v3 << ((uint(v4)) & 63))
        _ = v12
    } else {
        v13 := []int{}
        switch -2147483648 {
        case 127:
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int8 = 100
                if true {
                    v3 = v3
                    for fi16 := 0; fi16 < 13; fi16++ {
                        switch 32767 {
                        case 128:
                            _ = len(v2)
                        case 32768:
                            v17 := GenS1{F0: -1, F1: v3}
                            _ = v17
                        case 42:
                            v18 := GenS1{F0: 9223372036854775807, F1: (v3 + 681)}
                            _ = v18
                        }
                        _ = fi16
                    }
                } else {
                    v19 := GenS1{F0: -172, F1: v3}
                    _ = v19
                }
                v15 -= v15
                _ = fi14
                _ = v15
            }
        default:
            var v20 int16 = 688
            _ = v20
        }
        if (!true) {
            switch -970 {
            case 2147483647:
                switch -413 {
                case -145:
                    var v21 int32 = (v11 + (v11 % (((v11 << ((0) & 63))) | 1)))
                    _ = v21
                }
            case 895:
                switch 579 {
                case 113:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS2{F0: -128, F1: 255}
                        switch fi22 {
                        case -2147483648:
                            v3 -= v3
                        case 2147483647:
                            v4 += 36
                        case 32767:
                            v24 := GenS3{F0: 1, F1: 0, F2: 0}
                            _ = v24
                        }
                        _ = fi22
                        _ = v23
                    }
                default:
                    v25 := GenS4{F0: 42, F1: 961}
                    _ = v25
                }
            }
            if true {
                v4 = (v4 ^ uint8(v3))
                v3 = (v3 - v3)
                v3 = (v3 + v3)
            }
        } else {
            switch 77 {
            case -980:
                fmt.Println(v11, v3)
            case 32768:
                v26 := GenS5{F0: 110}
                _ = v26
            case 127:
                var v27 uint64 = 2
                _ = v27
            }
            var v28 int32 = -1
            _ = v28
        }
        _ = v13
    }
    v29 := GenS6{F0: (^(v3))}
    for fi30 := 0; fi30 < 14; fi30++ {
        if (false || (-2.25 < 1e10)) {
            switch fi30 {
            case -1:
                fmt.Println(v3, v4, fi30)
            default:
                if ((32767 > 820) || ((-2.25 <= 3.14159) || (42 <= 1))) {
                    v2["\n\t"] = (fi30 - (fi30 ^ (fi30 | (fi30 % ((int(v3)) | 1)))))
                    v31 := GenS4{F0: -1, F1: 2}
                    _ = v31
                }
            }
            if false {
                v32 := GenS3{F0: 930, F1: 0, F2: 0}
                if true {
                    fmt.Println(v3, v3, v3)
                    for fi33 := 0; fi33 < 13; fi33++ {
                        if (127 <= 1) {
                            switch ("foo bar" + ((("a" + (("a" + ("foo bar" + "foo bar")) + "a")) + "a") + "hello")) {
                            case "\n\t":
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        switch fi30 {
                                        case -524:
                                            v34 := v1[4294967295]
                                            _ = v34
                                        case 1:
                                            v3 -= 32767
                                        }
                                        if true {
                                            v35 := GenS7{F0: fi33}
                                            v11 -= 42
                                            _ = v35
                                        }
                                        var v36 uint16 = v3
                                        fmt.Println(v4)
                                        _ = v36
                                    case "a":
                                        v37 := GenS8{F0: fi30}
                                        _ = v37
                                    case "foo bar":
                                        fmt.Println(v11, fi33, v3)
                                    }
                                    var v38 float64 = 3.14159
                                    _ = v38
                                } else {
                                    v4 += 127
                                    switch -98 {
                                    case 127:
                                        v4 -= (v4 | 79)
                                    default:
                                        v3 = (^(v3))
                                    }
                                }
                            case "a":
                                var v39 bool = (-128 != 100)
                                _ = v39
                            case "hello":
                                var v40 uint = 2
                                _ = v40
                            }
                            if ((fi30 - (fi33 - (fi30 >> ((65535) & 63)))) >= fi33) {
                                var v41 uint = 1
                                _ = v41
                            } else {
                                var v42 int16 = -1
                                _ = v42
                            }
                        } else {
                            v43 := GenS9{F0: 42}
                            _ = v43
                        }
                        switch fi30 {
                        case -905:
                            var v44 uint16 = v3
                            _ = v44
                        }
                        var v45 bool = true
                        _ = fi33
                        _ = v45
                    }
                }
                for fi46 := 0; fi46 < 1; fi46++ {
                    v2["a"] = (fi46 | (^(fi46)))
                    _ = len(v1)
                    v47 := []int{(fi30 >> ((255) & 63)), 290}
                    _ = fi46
                    _ = v47
                }
                _ = v32
            } else {
                var v48 int16 = -32768
                v1[65535] = (true && false)
                _ = v48
            }
        }
        v3 -= v3
        var v49 uint64 = 42
        _ = fi30
        _ = v49
    }
    if true {
        for fi50 := 0; fi50 < 18; fi50++ {
            for fi51 := 0; fi51 < 14; fi51++ {
                switch ("\n\t" + (("foo bar" + "a") + "a")) {
                case "":
                    v4 = (v4 + (v4 - v4))
                case "\n\t":
                    fmt.Println(v3)
                }
                v52 := []int{fi50}
                _ = fi51
                _ = v52
            }
            switch ((("\n\t" + ("a" + (("foo bar" + "a") + "a"))) + ("\n\t" + "hello")) + "") {
            case "\n\t":
                var v53 int = fi50
                _ = v53
            case "foo bar":
                if true {
                    var v54 uint = 18446744073709551615
                    switch 32768 {
                    case 1:
                        v4 -= v4
                    case 4294967295:
                        v54 = (^(v54))
                    case -247:
                        fmt.Println(v3, v3)
                    default:
                        switch -1 {
                        case 32767:
                            if true {
                                var v55 uint16 = v3
                                _ = v55
                            } else {
                                v3 = (v3 ^ v3)
                                for fi56 := 0; fi56 < 3; fi56++ {
                                    var v57 float64 = ((-2.25 * 1.5) * 0.1)
                                    if (1 <= 1) {
                                        if (v4 != 0) {
                                            var v58 uint8 = v4
                                            var v59 uint32 = 223
                                            var v60 float32 = 3.14159
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v11 = (v11 * v11)
                                        }
                                        var v61 uint16 = v3
                                        v2[("" + "\n\t")] = (fi50 | 662)
                                        var v62 uint64 = uint64(v4)
                                        _ = v61
                                        _ = v62
                                    } else {
                                        fmt.Println(v54, fi56, v57)
                                    }
                                    v63 := GenS10{F0: 540}
                                    if (!false) {
                                        var v64 uint64 = 4294967295
                                        var v65 int16 = 223
                                        v11 -= (v11 >> ((v54) & 63))
                                        _ = v64
                                        _ = v65
                                    }
                                    if (v11 > (v11 << ((4294967295) & 63))) {
                                        if true {
                                            v66 := GenS4{F0: 42, F1: 2}
                                            switch fi50 {
                                            case 32767:
                                                var v67 float64 = -2.25
                                                _ = v67
                                            case -2147483648:
                                                v4 = (v4 ^ (v4 | v4))
                                            case 0:
                                                var v68 float64 = 1.5
                                                _ = v68
                                            default:
                                                v54 = v54
                                                v3 = v3
                                                v11 = (^(v11))
                                            }
                                            v69 := GenS9{F0: -32768}
                                            var v70 uint16 = v3
                                            _ = v66
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                    v71 := []int{(fi50 + fi56), fi50, fi56, fi56}
                                    _ = fi56
                                    _ = v57
                                    _ = v63
                                    _ = v71
                                }
                            }
                        default:
                            if ((!true) && true) {
                                var v72 uint = v54
                                v72 += (v72 & (v72 % ((v72) | 1)))
                                v73 := GenS3{F0: 1, F1: (v72 - v72), F2: -1}
                                _ = v72
                                _ = v73
                            }
                        }
                    }
                    _ = v54
                } else {
                    v74 := GenS11{F0: 89, F1: (fi50 | 32767)}
                    v11 = v11
                    fmt.Println(v4)
                    _ = v74
                }
            }
            var v75 int32 = -1
            _ = fi50
            _ = v75
        }
        v3 -= (^(v3))
    }
    fmt.Println(v3, v11, v3)
    var v76 int64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v76
}
