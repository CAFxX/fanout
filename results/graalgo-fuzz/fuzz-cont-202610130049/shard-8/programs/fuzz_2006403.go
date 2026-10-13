package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint; F2 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 bool; F2 int32 }

func main() {
    var v1 int = 127
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 uint16 = 648
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 8; fi6++ {
        switch (fi6 + v1) {
        case -400:
            var v7 int = (v1 & v1)
            _ = v7
        default:
            v8 := GenS1{F0: (fi6 % (((^(fi6))) | 1)), F1: 0}
            _ = v8
        }
        _ = fi6
    }
    fmt.Println(v1, v1)
    switch ("" + "foo bar") {
    case "foo bar":
        if true {
            if (1e10 > (((1.5 * 0.0) / 3.14159) + 3.14159)) {
                v9 := GenS2{F0: 35, F1: "foo bar"}
                _ = v9
            } else {
                fmt.Println(v4, v4)
            }
            v10 := []int{v1, 32768}
            _ = v10
        } else {
            switch "a" {
            case "foo bar":
                v11 := GenS3{F0: 65535, F1: -337, F2: 123}
                _ = v11
            case "\n\t":
                for fi12 := 0; fi12 < 20; fi12++ {
                    v1 -= 0
                    _ = fi12
                }
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                v4 += v4
                _ = fi13
            }
            v14 := GenS3{F0: 539, F1: -2147483648, F2: 1}
            _ = v14
        }
    case "":
        fmt.Println(v1)
    }
    var v15 uint64 = 0
    v4 = 0
    v16 := GenS1{F0: v1, F1: 42}
    v15 -= v15
    var v17 float32 = ((1e10 + 0.1) - 3.14159)
    if false {
        fmt.Println(v1)
        v18 := []int{(v1 - -1), v1, 308, (v1 << ((255) & 63))}
        _ = v18
    } else {
        var v19 float64 = 1.5
        var v20 bool = true
        var v21 uint32 = 1
        _ = v19
        _ = v20
        _ = v21
    }
    if (((v4 <= 65535) && (42 <= 1)) || ((v15 % ((v15) | 1)) >= 42)) {
        v22 := GenS1{F0: v1, F1: 65535}
        _ = v22
    } else {
        if (1 == 1) {
            if true {
                for fi23 := 0; fi23 < 12; fi23++ {
                    v5[2] = "hello"
                    _ = fi23
                }
                v24 := GenS4{F0: ((((v17 + (v17 + (v17 - float32(-2.25)))) + (v17 * (3.14159 + 1e10))) - v17) != (0.0 + v17))}
                if (21 < 0) {
                    v25 := GenS3{F0: 1, F1: -1, F2: 255}
                    fmt.Println(v4)
                    _ = v25
                } else {
                    v4 += (v4 / ((v4) | 1))
                    if true {
                        v26 := []int{(v1 / (((v1 & v1)) | 1)), v1, v1, 32767}
                        fmt.Println(v4, v17, v4)
                        _ = v26
                    } else {
                        switch (v1 & -9223372036854775808) {
                        case -500:
                            v27 := GenS3{F0: 65535, F1: -2147483648, F2: 128}
                            _ = v27
                        case -2147483648:
                            if true {
                                if (true && true) {
                                    v28 := GenS4{F0: true}
                                    if (true && ((("hello" + ("" + "foo bar")) + "foo bar") < "\n\t")) {
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            v29 := []int{2147483647, v1, (v1 ^ (v1 ^ v1))}
                                            _ = v29
                                        case "\n\t":
                                            var v30 float32 = v17
                                            _ = v30
                                        }
                                    }
                                    _ = v28
                                } else {
                                    v1 = (v1 * -808)
                                }
                                switch (v1 % ((1) | 1)) {
                                case 1:
                                    switch 4294967295 {
                                    case 32767:
                                        for fi31 := 0; fi31 < 6; fi31++ {
                                            v32 := GenS5{F0: -907}
                                            var v33 uint64 = 959
                                            _ = fi31
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                case 204:
                                    switch "foo bar" {
                                    case "":
                                        v17 -= v17
                                    case "hello":
                                        v4 -= v4
                                    default:
                                        v1 += v1
                                    }
                                case 207:
                                    for fi34 := 0; fi34 < 5; fi34++ {
                                        v15 = v15
                                        for fi35 := 0; fi35 < 20; fi35++ {
                                            var v36 uint16 = 2
                                            v37 := []int{}
                                            switch (fi34 % ((-1) | 1)) {
                                            case 128:
                                                v1 -= fi35
                                            case 1:
                                                fmt.Println(v36, v1, v4)
                                            }
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = fi34
                                    }
                                }
                            } else {
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    v39 := []int{(v1 / ((-9223372036854775808) | 1)), (fi38 >> ((30) & 63)), (fi38 >> ((255) & 63))}
                                    if false {
                                        fmt.Println(v1, v1)
                                        if true {
                                            var v40 string = "foo bar"
                                            _ = v40
                                        } else {
                                            v1 += 127
                                            fmt.Println(fi38, v1, v1)
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                }
                                v4 -= 1
                                v41 := GenS5{F0: 42}
                                _ = v41
                            }
                        }
                        for fi42 := 0; fi42 < 18; fi42++ {
                            for fi43 := 0; fi43 < 12; fi43++ {
                                switch fi42 {
                                case 4294967295:
                                    v44 := GenS5{F0: -453}
                                    var v45 bool = false
                                    v46 := GenS6{F0: false, F1: uint(v15), F2: -550}
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                default:
                                    v17 -= (float32(1.5) / v17)
                                }
                                _ = fi43
                            }
                            var v47 uint8 = 127
                            _ = fi42
                            _ = v47
                        }
                    }
                    v48 := []int{}
                    _ = v48
                }
                _ = v24
            } else {
                v49 := []int{}
                fmt.Println(v15)
                v50 := GenS1{F0: -569, F1: 0}
                _ = v49
                _ = v50
            }
        }
        if (v1 <= (v1 + v1)) {
            for fi51 := 0; fi51 < 18; fi51++ {
                fmt.Println(v1)
                var v52 int16 = 1
                var v53 bool = (-1 == v52)
                _ = fi51
                _ = v52
                _ = v53
            }
            v15 += 443
            fmt.Println(v17, v17, v4)
        } else {
            var v54 float32 = float32(((float64(v17) * 0.0) * 0.1))
            _ = v54
        }
        if false {
            var v55 uint8 = 127
            for fi56 := 0; fi56 < 3; fi56++ {
                for fi57 := 0; fi57 < 7; fi57++ {
                    v58 := []int{fi57, (fi57 / ((0) | 1)), v1, (fi57 % ((fi57) | 1))}
                    if false {
                        v59 := GenS2{F0: 4294967295, F1: "\n\t"}
                        if (!false) {
                            var v60 int32 = int32(fi56)
                            if true {
                                v60 = v60
                                var v61 uint16 = (v4 & (v4 >> ((1) & 63)))
                                v17 += 1e10
                                if (4294967295 > 148) {
                                    var v62 int16 = 1
                                    _ = v62
                                }
                                _ = v61
                            } else {
                                v55 -= (v55 % ((v55) | 1))
                                v1 -= 9223372036854775807
                            }
                            _ = v60
                        }
                        v63 := GenS7{F0: 1}
                        _ = v59
                        _ = v63
                    }
                    if ((v55 + 202) != (v55 >> ((1) & 63))) {
                        var v64 uint8 = v55
                        _ = v64
                    }
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
            fmt.Println(v17, v1)
            _ = v55
        }
    }
    var v65 uint8 = 128
    if true {
        var v66 uint = 0
        _ = v66
    } else {
        v67 := []int{v1, (^(v1)), v1, v1}
        v15 -= uint64(v4)
        _ = v67
    }
    for fi68 := 0; fi68 < 6; fi68++ {
        v69 := GenS8{F0: -1, F1: (true || false), F2: 243}
        if false {
            switch 2147483647 {
            case 32767:
                v65 = 128
            case 128:
                v70 := GenS2{F0: 842, F1: ""}
                _ = v70
            case 0:
                v71 := GenS8{F0: 44, F1: false, F2: -335}
                _ = v71
            }
            v72 := GenS7{F0: 2}
            v73 := GenS7{F0: 0}
            _ = v72
            _ = v73
        } else {
            v74 := []int{v1, 32768, 42}
            _ = v74
        }
        var v75 uint16 = 32767
        _ = fi68
        _ = v69
        _ = v75
    }
    if (65535 > 859) {
        if ((v17 * v17) > 0.0) {
            fmt.Println(v4, v15)
        }
        v76 := v2[v1]
        _ = v76
    }
    var v77 string = ""
    for fi78 := 0; fi78 < 1; fi78++ {
        if true {
            v79 := GenS7{F0: 2147483648}
            if (v4 < v4) {
                for fi80 := 0; fi80 < 2; fi80++ {
                    v15 += (v15 - (v15 >> ((65535) & 63)))
                    v4 = (v4 - v4)
                    _ = fi80
                }
                delete(v2, -1)
                switch (v1 & fi78) {
                case 0:
                    fmt.Println(v1, fi78)
                }
            } else {
                v15 = (v15 ^ v15)
                var v81 uint64 = v15
                var v82 uint8 = (v65 >> ((1) & 63))
                _ = v81
                _ = v82
            }
            var v83 bool = false
            _ = v79
            _ = v83
        }
        v77 = "\n\t"
        _ = fi78
    }
    switch (v1 & -1) {
    case 4294967295:
        switch v1 {
        case -9223372036854775808:
            for fi84 := 0; fi84 < 4; fi84++ {
                v77 = v77
                v17 -= v17
                _ = fi84
            }
        case -306:
            v85 := []int{}
            _ = v85
        }
    case -9223372036854775808:
        var v86 uint16 = v4
        _ = v86
    case 721:
        var v87 float32 = (v17 - v17)
        _ = v87
    }
    var v88 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v77:", v77)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v77
    _ = v88
}
