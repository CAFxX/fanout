package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int32; F2 bool }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 uint8 }

func main() {
    var v1 string = "hello"
    var v2 int32 = -944
    var v3 uint32 = 96
    var v4 int32 = 0
    var v5 string = v1
    v6 := make(map[int]uint32)
    var v7 int = 128
    fmt.Println(v7, v5)
    if (3.14159 >= float64((0.1 + 1.5))) {
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                v1 = "foo bar"
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: -606, F1: 2, F2: 65535}
                    v1 = v1
                    if (!(-2.25 > (float64(1.5) + (3.14159 + float64(3.14159))))) {
                        fmt.Println(v7)
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        } else {
            if true {
                if ((!true) || true) {
                    fmt.Println(v1, v5)
                } else {
                    var v11 int = (v7 % (((v7 % ((int(v2)) | 1))) | 1))
                    if (32767 == 0) {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 uint64 = 1
                            if false {
                                for fi14 := 0; fi14 < 8; fi14++ {
                                    var v15 string = v5
                                    if (v3 <= v3) {
                                        fmt.Println(v2, v7)
                                        v16 := GenS2{F0: v13}
                                        _ = v16
                                    }
                                    var v17 float64 = (((0.1 / -2.25) + 0.1) * -2.25)
                                    _ = fi14
                                    _ = v15
                                    _ = v17
                                }
                            } else {
                                if true {
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        if true {
                                            v19 := v6[fi12]
                                            v20 := GenS3{F0: ("hello" + ("" + v5)), F1: 1}
                                            _ = v19
                                            _ = v20
                                        }
                                        if false {
                                            var v21 uint8 = 125
                                            _ = v21
                                        }
                                        _ = fi18
                                    }
                                }
                                if (-443 <= (v2 << ((1) & 63))) {
                                    var v22 uint32 = v3
                                    if ((false && (false && (true || true))) || false) {
                                        switch "hello" {
                                        case "":
                                            var v23 int = int(v2)
                                            _ = v23
                                        case "a":
                                            switch (v11 << ((2) & 63)) {
                                            case 32767:
                                                v11 += (v7 + (v7 ^ (fi12 + fi12)))
                                            case 9223372036854775807:
                                                var v24 uint16 = 45
                                                _ = v24
                                            case 870:
                                                var v25 uint32 = (^(v3))
                                                _ = v25
                                            default:
                                                _ = len(v6)
                                            }
                                        }
                                    } else {
                                        switch (^(v7)) {
                                        case 9223372036854775807:
                                            var v26 uint = 255
                                            _ = v26
                                        }
                                    }
                                    _ = v22
                                } else {
                                    fmt.Println(v4, v1)
                                    fmt.Println(v1, v7, v7)
                                    _ = len(v6)
                                }
                            }
                            fmt.Println(v4, v1)
                            _ = fi12
                            _ = v13
                        }
                        for fi27 := 0; fi27 < 12; fi27++ {
                            for fi28 := 0; fi28 < 5; fi28++ {
                                if false {
                                    v29 := []int{(fi27 >> ((967) & 63)), -600, (v11 % ((-9223372036854775808) | 1))}
                                    v2 += (^(v2))
                                    if false {
                                        if false {
                                            if false {
                                                v30 := GenS4{F0: fi28, F1: 255}
                                                if len(v29) >= 2 {
                                                    v29 = v29[1:len(v29)]
                                                }
                                                v31 := v6[v7]
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 uint8 = 127
                                                v7 = fi28
                                                var v33 uint = 4294967295
                                                _ = v32
                                                _ = v33
                                            }
                                            if true {
                                                _ = len(v6)
                                                var v34 int = (v7 ^ (fi27 + v7))
                                                if false {
                                                    v6[v7] = (v3 >> ((18446744073709551615) & 63))
                                                    var v35 int = v11
                                                    var v36 uint32 = 2147483647
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    v4 += v4
                                                    v37 := GenS1{F0: v34, F1: 255, F2: 1}
                                                    _ = v37
                                                }
                                                _ = v34
                                            }
                                            var v38 uint16 = 1
                                            _ = v38
                                        } else {
                                            switch (v11 ^ (v11 | v11)) {
                                            case 32768:
                                                var v39 uint32 = v3
                                                _ = v39
                                            case 32767:
                                                v40 := GenS4{F0: (fi28 + -1), F1: 255}
                                                v41 := GenS5{F0: fi27}
                                                _ = v40
                                                _ = v41
                                            default:
                                                var v42 uint16 = 32767
                                                _ = v42
                                            }
                                            var v43 string = v5
                                            v44 := GenS6{F0: (v2 + -2147483648), F1: v2, F2: false}
                                            fmt.Println(v2, v7, v5)
                                            fmt.Println(v2, v4)
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = len(v29)
                                        v3 += v3
                                        var v45 int64 = -9223372036854775808
                                        fmt.Println(v45)
                                        var v46 float32 = -2.25
                                        v29 = append(v29, 2147483647)
                                        v3 = v3
                                        switch v11 {
                                        case -2147483648:
                                            var v47 uint8 = 222
                                            _ = v47
                                        case -1:
                                            v45 -= v45
                                        case 32768:
                                            switch "" {
                                            case "":
                                                v48 := GenS4{F0: fi28, F1: 127}
                                                _ = v48
                                            case "\n\t":
                                                if len(v29) >= 2 {
                                                    v29 = v29[1:len(v29)]
                                                }
                                            case "hello":
                                                v11 = (^(fi28))
                                            }
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v29
                                } else {
                                    v3 -= v3
                                }
                                _ = fi28
                            }
                            v3 -= v3
                            _ = fi27
                        }
                        v11 = v11
                    } else {
                        var v49 uint8 = 162
                        v50 := GenS2{F0: 515}
                        v11 -= v7
                        _ = v49
                        _ = v50
                    }
                    _ = v11
                }
            }
        }
    } else {
        switch v7 {
        case 4294967295:
            v1 = (v5 + v5)
        case -2147483648:
            v51 := []int{}
            _ = v51
        case -205:
            for fi52 := 0; fi52 < 7; fi52++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(v3, v3)
                default:
                    _ = len(v6)
                }
                _ = fi52
            }
        default:
            var v53 int64 = 1
            _ = v53
        }
        v54 := []int{v7, -1, v7}
        _ = v54
    }
    var v55 uint = 18446744073709551615
    var v56 uint = v55
    if (-246 <= -9223372036854775808) {
        var v57 int = 990
        for fi58 := 0; fi58 < 7; fi58++ {
            v59 := []int{fi58, 2147483647, 32767, v7}
            var v60 bool = ((false || (true || (170 == 128))) || false)
            fmt.Println(fi58)
            _ = fi58
            _ = v59
            _ = v60
        }
        _ = v57
    }
    v61 := v6[(v7 << ((42) & 63))]
    if ((65535 <= (v55 | 966)) || false) {
        fmt.Println(v56)
        var v62 float32 = 1.5
        _ = v62
    } else {
        if (!(4294967295 >= v7)) {
            v63 := []int{}
            _ = v63
        } else {
            v64 := GenS4{F0: (v7 ^ v7), F1: 0}
            fmt.Println(v5, v2)
            v65 := []int{-826, (v7 ^ (v7 / (((v7 & v7)) | 1)))}
            _ = v64
            _ = v65
        }
    }
    v66 := v6[32767]
    fmt.Println(v4, v7)
    if true {
        v67 := GenS7{F0: 1, F1: (v56 | v56)}
        for fi68 := 0; fi68 < 10; fi68++ {
            var v69 string = ""
            if (255 <= 255) {
                v69 = "hello"
                var v70 float32 = (1e10 - 1.5)
                _ = v70
            }
            var v71 bool = ((0.1 / ((3.14159 / 1.0) / ((0.1 * (1.5 * -2.25)) - 0.0))) < 1e10)
            _ = fi68
            _ = v69
            _ = v71
        }
        v56 = v56
        _ = v67
    } else {
        switch "hello" {
        case "hello":
            for fi72 := 0; fi72 < 10; fi72++ {
                var v73 uint16 = 174
                v61 -= (^(v61))
                for fi74 := 0; fi74 < 2; fi74++ {
                    v56 = 137
                    var v75 uint32 = (v66 / ((2) | 1))
                    _ = fi74
                    _ = v75
                }
                _ = fi72
                _ = v73
            }
        default:
            v6[(v7 | v7)] = v61
        }
        switch "" {
        case "hello":
            v61 += v66
        }
    }
    switch 42 {
    case 563:
        if false {
            var v76 int64 = 512
            fmt.Println(v5, v4, v2)
            _ = v76
        } else {
            v77 := GenS8{F0: 42, F1: 144, F2: 127}
            _ = v77
        }
    case 937:
        var v78 float64 = -2.25
        _ = v78
    default:
        for fi79 := 0; fi79 < 5; fi79++ {
            switch (v5 + ("a" + v1)) {
            case "foo bar":
                v80 := []int{}
                _ = v80
            case "a":
                switch (^(fi79)) {
                case -312:
                    v2 -= 2147483647
                case -1:
                    v81 := v6[(v7 - -95)]
                    _ = v81
                default:
                    v7 += (v7 | v7)
                }
            case "\n\t":
                fmt.Println(v1, v4, v7)
            }
            _ = fi79
        }
    }
    v82 := []int{(^(v7))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v61
    _ = v66
    _ = v7
    _ = v82
}
