package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int32; F1 string; F2 string }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint64, p2 uint16) uint {
    p2 = p2
    var v7 int8 = 100
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 += v7
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint = 919
    var v2 uint64 = 2
    v3 := []int{430, -512}
    var v4 int16 = -821
    var v5 uint8 = 127
    v6 := []string{"", "a", "foo bar", "a"}
    switch "" {
    case "\n\t":
        if (!(float64(-2.25) <= ((0.1 / 1e10) / 0.1))) {
            var v10 uint8 = (^(v5))
            var v11 float64 = 0.0
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                v4 = (v4 * v4)
                v5 += 202
                _ = fi12
            }
            if true {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
            }
        }
    case "":
        if (v4 < (v4 << ((v1) & 63))) {
            var v13 uint64 = (v2 & v2)
            if true {
                v4 += 0
                var v14 uint16 = uint16(v1)
                _ = v14
            }
            _ = v13
        } else {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS1{F0: 1000}
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: 0, F1: 65535}
            if ((((1e10 + (1.5 / 1e10)) - 3.14159) / float32(3.14159)) == 1e10) {
                var v18 int64 = 42
                v4 += int16(v5)
                if (((0.0 + (3.14159 / 3.14159)) > -2.25) && (!true)) {
                    if true {
                        v3 = append(v3, 9223372036854775807)
                        v19 := GenS2{F0: 2147483647, F1: 42}
                        var v20 uint = (v1 * 408)
                        _ = v19
                        _ = v20
                    } else {
                        if (genFn0(-1, v2, uint16(v5)) <= uint(v2)) {
                            var v21 float32 = 0.0
                            if (!true) {
                                v22 := GenS2{F0: 0, F1: 42}
                                _ = v22
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v23 := GenS2{F0: 4294967295, F1: 0}
                                _ = v23
                            }
                            _ = v21
                        } else {
                            var v24 float32 = (0.0 + (float32(1e10) * (-2.25 + (-2.25 - (((((1e10 * (((float32(1.5) - ((float32(0.0) + ((float32(float64(1.5)) / 1.5) * 0.1)) + (3.14159 - (1e10 / 0.1)))) - 1.5) + 0.0)) + 0.0) * 3.14159) - 1e10) * 1e10)))))
                            v5 = v5
                            _ = v24
                        }
                    }
                } else {
                    var v25 uint8 = (v5 + v5)
                    _ = v25
                }
                _ = v18
            }
            _ = v17
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v26 := GenS1{F0: -279}
    v4 -= v4
    v2 -= v2
    switch 42 {
    case 32767:
        switch 2147483647 {
        case 127:
            v4 -= (v4 * v4)
        case -9223372036854775808:
            if (-903 == -224) {
                v5 -= 127
            } else {
                var v27 int8 = -124
                for fi28 := 0; fi28 < 20; fi28++ {
                    if ((v4 | v4) >= v4) {
                        v2 = 348
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v30 := GenS1{F0: (^(v4))}
                            if false {
                                var v31 uint = (v1 / (((v1 % (((v1 | (v1 + 255))) | 1))) | 1))
                                v32 := GenS1{F0: (v4 % (((v4 * v4)) | 1))}
                                if true {
                                    v2 += 18446744073709551615
                                }
                                _ = v31
                                _ = v32
                            }
                            switch "\n\t" {
                            case "a":
                                var v33 uint = genFn0(0, v2, 2)
                                _ = v33
                            default:
                                switch "a" {
                                case "hello":
                                    fmt.Println(v1, v1)
                                case "foo bar":
                                    var v34 int = -2147483648
                                    _ = v34
                                case "a":
                                    if false {
                                        var v35 string = "hello"
                                        _ = v35
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = -1
                                        }
                                    }
                                default:
                                    var v36 int8 = v27
                                    _ = v36
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                        _ = len(v6)
                    } else {
                        if ("\n\t" < "\n\t") {
                            v4 = v4
                        }
                    }
                    _ = fi28
                }
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
                _ = v27
            }
        case -1:
            v4 = 1
        default:
            if true {
                v4 -= (v4 << ((genFn0(2147483647, v2, 87)) & 63))
                v5 += 1
                v37 := GenS2{F0: 2, F1: 32767}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 int8 = 88
                    v2 = uint64(v5)
                    v40 := GenS1{F0: v4}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
        }
    case 127:
        for fi41 := 0; fi41 < 11; fi41++ {
            if (4294967295 < 479) {
                var v42 int32 = 0
                _ = v42
            } else {
                if false {
                    var v43 uint8 = v5
                    var v44 string = "\n\t"
                    _ = v43
                    _ = v44
                }
                if true {
                    v45 := GenS3{F0: "foo bar", F1: -9223372036854775808}
                    _ = v45
                } else {
                    if (int8(v2) <= -107) {
                        v4 = 42
                        if ((v5 & 255) <= (v5 + (v5 ^ 177))) {
                            switch (fi41 / ((fi41) | 1)) {
                            case 2147483647:
                                v46 := GenS2{F0: uint32(v5), F1: 42}
                                _ = v46
                            default:
                                if len(v6) > 0 {
                                    v6[0] = "\n\t"
                                }
                            }
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                for fi47 := 0; fi47 < 15; fi47++ {
                                    var v48 uint16 = 32767
                                    _ = len(v3)
                                    _ = fi47
                                    _ = v48
                                }
                            case "":
                                var v49 int = fi41
                                _ = v49
                            default:
                                v1 += v1
                            }
                            v4 = 32767
                        }
                        var v50 float32 = 1e10
                        _ = v50
                    }
                }
                v2 = 361
            }
            v2 -= 437
            for fi51 := 0; fi51 < 5; fi51++ {
                v52 := GenS4{F0: 420, F1: 42, F2: 42}
                _ = fi51
                _ = v52
            }
            _ = fi41
        }
    }
    switch "a" {
    case "\n\t":
        if false {
            var v53 float32 = -2.25
            _ = v53
        } else {
            fmt.Println(v2)
            if false {
                var v54 uint = v1
                v4 = -1
                _ = v54
            }
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
        }
    default:
        switch "a" {
        case "\n\t":
            v55 := GenS5{F0: 2, F1: 0, F2: (v2 << ((v1) & 63))}
            _ = v55
        case "hello":
            v4 = 42
        case "foo bar":
            v4 -= 646
        }
    }
    v56 := GenS3{F0: "hello", F1: 1}
    if ((3.14159 - 1e10) != (1e10 / 1e10)) {
        var v57 int = 318
        v3 = append(v3, v57)
        for fi58 := 0; fi58 < 12; fi58++ {
            v2 += v2
            v57 = v57
            _ = fi58
        }
        _ = v57
    }
    v59 := GenS6{F0: 2147483647, F1: ("hello" + ""), F2: "hello"}
    v60 := GenS7{F0: 537}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v61 uint8 = (v5 << ((genFn0(42, 2, 255)) & 63))
    switch 0 {
    case 665:
        for fi62 := 0; fi62 < 16; fi62++ {
            var v63 uint = (^(v1))
            var v64 int16 = v4
            switch (("hello" + ("a" + "")) + "foo bar") {
            case "hello":
                v65 := GenS3{F0: "\n\t", F1: int64(v4)}
                _ = v65
            case "\n\t":
                var v66 bool = (0 == (v1 ^ genFn0(-2147483648, v2, 866)))
                _ = v66
            case "foo bar":
                v4 += 0
            default:
                var v67 int64 = 856
                _ = v67
            }
            _ = fi62
            _ = v63
            _ = v64
        }
    case 127:
        v68 := GenS5{F0: 2, F1: 13, F2: v2}
        _ = v68
    }
    if true {
        for fi69 := 0; fi69 < 1; fi69++ {
            fmt.Println(v2, fi69, v2)
            _ = fi69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
