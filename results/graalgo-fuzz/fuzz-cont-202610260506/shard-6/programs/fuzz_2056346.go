package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int; F1 int32; F2 int32 }
type GenS11 struct { F0 uint64; F1 uint }

func main() {
    var v1 int = 32767
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 int64 = 865
    switch ("\n\t" + "a") {
    case "":
        var v5 int = (-(v1))
        _ = v5
    case "\n\t":
        fmt.Println(v4, v4)
    case "hello":
        var v6 float32 = -2.25
        _ = v6
    }
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            if (0 != int16(v4)) {
                switch (v1 - (-(fi8))) {
                case -663:
                    v9 := GenS2{F0: 0, F1: 18446744073709551615, F2: 1}
                    _ = v9
                case -732:
                    if (127 == 2) {
                        v1 = (fi8 << ((439) & 63))
                    }
                }
                switch "" {
                case "a":
                    var v10 int64 = (v4 << ((1) & 63))
                    _ = v10
                default:
                    var v11 int32 = (v3 | (v3 >> ((461) & 63)))
                    _ = v11
                }
                v3 = 2147483647
            } else {
                switch v1 {
                case 4294967295:
                    if false {
                        if true {
                            if false {
                                v4 -= -83
                            } else {
                                if (false && true) {
                                    v1 = (fi8 / (((-(v1))) | 1))
                                    v12 := GenS3{F0: 0, F1: 122}
                                    _ = v12
                                } else {
                                    fmt.Println(v1, v1, v3)
                                }
                            }
                        }
                    } else {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v14 := GenS1{F0: -1}
                            _ = fi13
                            _ = v14
                        }
                        switch 128 {
                        case 42:
                            var v15 uint = 471
                            _ = v15
                        }
                    }
                default:
                    switch (v1 ^ 4294967295) {
                    case 344:
                        var v16 int8 = 127
                        _ = v16
                    case 128:
                        v17 := GenS1{F0: 40}
                        _ = v17
                    case 42:
                        v18 := []int{fi8, fi8, (v1 % ((fi8) | 1)), (^(v1))}
                        _ = v18
                    default:
                        switch "a" {
                        case "foo bar":
                            var v19 float32 = float32(0.0)
                            _ = v19
                        case "\n\t":
                            var v20 string = "hello"
                            _ = v20
                        case "hello":
                            for fi21 := 0; fi21 < 9; fi21++ {
                                if (42 != 845) {
                                    var v22 int16 = 0
                                    if (255 == 2) {
                                        var v23 int32 = v3
                                        if true {
                                            switch (v1 / (((fi8 % (((v1 ^ v1)) | 1))) | 1)) {
                                            case -271:
                                                var v24 uint32 = 1
                                                _ = v24
                                            case -2147483648:
                                                fmt.Println(v3, v3)
                                            case 9223372036854775807:
                                                v1 -= (v1 % ((v1) | 1))
                                            }
                                            if true {
                                                var v25 uint = 18446744073709551615
                                                var v26 uint8 = 127
                                                _ = v25
                                                _ = v26
                                            } else {
                                                fmt.Println(fi8)
                                                var v27 bool = (0.1 >= 0.0)
                                                _ = v27
                                            }
                                            var v28 string = ("a" + "")
                                            _ = v28
                                        }
                                        _ = v23
                                    }
                                    _ = v22
                                }
                                v29 := []int{}
                                _ = fi21
                                _ = v29
                            }
                        default:
                            for fi30 := 0; fi30 < 1; fi30++ {
                                v3 += (v3 | v3)
                                var v31 int16 = 42
                                if true {
                                    v32 := GenS4{F0: 18446744073709551615}
                                    _ = v32
                                } else {
                                    v33 := GenS5{F0: (v31 | v31)}
                                    _ = v33
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                    }
                }
            }
            v34 := []int{(fi8 >> ((65535) & 63))}
            _ = v34
        } else {
            v1 = fi8
            delete(v2, "\n\t")
        }
        _ = fi8
    }
    fmt.Println(v4, v4)
    if ((true || true) || true) {
        var v35 int32 = (-(v3))
        v3 -= 42
        _ = v35
    }
    var v36 int8 = 100
    if false {
        v36 -= -128
        v37 := GenS2{F0: v3, F1: 526, F2: 18446744073709551615}
        _ = v37
    }
    v38 := GenS6{F0: 120}
    if ((!true) || false) {
        fmt.Println(v1, v4)
        var v39 int64 = (v4 << ((0) & 63))
        _ = v39
    }
    v1 = v1
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 float32 = 3.14159
        v4 = (v4 ^ (^(v4)))
        var v42 float32 = ((float32((0.1 * 1.5)) * (float32((3.14159 / 1e10)) * (v41 / 1.0))) + (v41 - v41))
        _ = fi40
        _ = v41
        _ = v42
    }
    switch int(v4) {
    case 541:
        for fi43 := 0; fi43 < 17; fi43++ {
            fmt.Println(v4, v4, v4)
            v1 -= (fi43 << ((42) & 63))
            if false {
                v1 = (v1 + (v1 + (fi43 / ((fi43) | 1))))
                var v44 float32 = (1e10 / ((((3.14159 + 1e10) - 0.0) + 1.5) + float32(1.5)))
                if (true || (v1 == -2147483648)) {
                    fmt.Println(v44, v36, v36)
                } else {
                    for fi45 := 0; fi45 < 17; fi45++ {
                        switch (^(v1)) {
                        case 0:
                            v46 := []int{}
                            _ = v46
                        case 9223372036854775807:
                            v47 := GenS7{F0: (v36 | (v36 * -29))}
                            _ = v47
                        }
                        v48 := GenS6{F0: 173}
                        var v49 int16 = -32768
                        _ = fi45
                        _ = v48
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 2; fi50++ {
                        switch (^(fi50)) {
                        case 496:
                            v51 := []int{}
                            _ = v51
                        case -9223372036854775808:
                            if false {
                                if false {
                                    v4 -= -1
                                } else {
                                    fmt.Println(fi50, v44, fi50)
                                    v4 = v4
                                    v44 = (-2.25 * 0.1)
                                }
                                var v52 string = "a"
                                _ = v52
                            } else {
                                v53 := GenS8{F0: (v3 * v3), F1: 445}
                                _ = v53
                            }
                        case 91:
                            if false {
                                if ("" != ("\n\t" + "")) {
                                    v4 -= 0
                                    switch "\n\t" {
                                    case "hello":
                                        v36 += v36
                                    case "a":
                                        if false {
                                            switch "" {
                                            case "foo bar":
                                                var v54 int16 = 32767
                                                _ = v54
                                            }
                                            var v55 int64 = v4
                                            v56 := GenS1{F0: v36}
                                            _ = v55
                                            _ = v56
                                        } else {
                                            if (128 >= 20) {
                                                switch 775 {
                                                case 1:
                                                    v57 := []int{993, 128}
                                                    _ = v57
                                                case -766:
                                                    _ = len(v2)
                                                default:
                                                    if false {
                                                        var v58 uint32 = 65535
                                                        fmt.Println(v4, v3)
                                                        v58 -= (v58 ^ (v58 * v58))
                                                        _ = v58
                                                    } else {
                                                        if false {
                                                            if false {
                                                                delete(v2, "a")
                                                            }
                                                            v59 := []int{fi50}
                                                            v60 := GenS7{F0: (v36 & v36)}
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v44 += (float32(float64(1e10)) * -2.25)
                                                            v61 := GenS9{F0: 32767, F1: 0}
                                                            var v62 int16 = -398
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        v63 := GenS10{F0: fi43, F1: v3, F2: (v3 & v3)}
                                                        _ = v63
                                                    }
                                                }
                                                v3 = (v3 & v3)
                                            }
                                            v3 -= -2147483648
                                        }
                                    default:
                                        v64 := GenS11{F0: 1, F1: 1}
                                        _ = v64
                                    }
                                } else {
                                    v65 := GenS11{F0: 2, F1: 490}
                                    delete(v2, "")
                                    v44 -= 0.0
                                    _ = v65
                                }
                                var v66 bool = (true || (v44 <= v44))
                                var v67 int8 = (v36 % (((v36 | -128)) | 1))
                                _ = v66
                                _ = v67
                            }
                        default:
                            if false {
                                var v68 string = "a"
                                v69 := []int{}
                                _ = v68
                                _ = v69
                            } else {
                                var v70 int = fi50
                                _ = v70
                            }
                            if (((-372 == (^(v4))) && ((v36 + v36) >= 27)) && true) {
                                var v71 uint64 = 0
                                var v72 uint32 = 1
                                v73 := GenS4{F0: (v71 + (^(v71)))}
                                _ = v71
                                _ = v72
                                _ = v73
                            } else {
                                _ = len(v2)
                                var v74 int16 = 973
                                _ = v74
                            }
                        }
                        v75 := GenS4{F0: 4294967295}
                        _ = fi50
                        _ = v75
                    }
                    switch (v1 ^ v1) {
                    case -729:
                        switch fi43 {
                        case 1:
                            var v76 bool = true
                            _ = v76
                        }
                    case 32768:
                        for fi77 := 0; fi77 < 16; fi77++ {
                            var v78 float64 = 1e10
                            _ = fi77
                            _ = v78
                        }
                    case 127:
                        v4 = v4
                    default:
                        v79 := []int{fi43}
                        _ = v79
                    }
                }
                _ = v44
            } else {
                v4 += int64(v36)
                var v80 int32 = (^(v3))
                switch (fi43 >> ((18446744073709551615) & 63)) {
                case -556:
                    var v81 uint = 0
                    _ = v81
                }
                _ = v80
            }
            _ = fi43
        }
    case 9223372036854775807:
        for fi82 := 0; fi82 < 4; fi82++ {
            v3 += 0
            switch "foo bar" {
            case "\n\t":
                if (209 <= 42) {
                    var v83 uint16 = 0
                    switch v1 {
                    case 0:
                        var v84 uint32 = 4294967295
                        _ = v84
                    case 127:
                        var v85 uint64 = 18446744073709551615
                        _ = v85
                    case 9223372036854775807:
                        delete(v2, "")
                    }
                    _ = v83
                } else {
                    for fi86 := 0; fi86 < 6; fi86++ {
                        delete(v2, "foo bar")
                        fmt.Println(v1)
                        _ = fi86
                    }
                }
            default:
                v1 -= fi82
            }
            v4 = (v4 >> ((2) & 63))
            _ = fi82
        }
    default:
        v2[("hello" + "")] = false
    }
    v87 := GenS10{F0: 1, F1: -877, F2: -46}
    v3 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v7
    _ = v87
}
