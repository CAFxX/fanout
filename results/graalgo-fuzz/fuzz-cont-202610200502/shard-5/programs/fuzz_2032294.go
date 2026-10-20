package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := []int{32767, 2147483647, -188, -856}
    var v4 int = 859
    v5 := []uint32{2147483647, 42, 2}
    var v6 uint32 = 103
    v6 += 2147483647
    v7 := GenS1{F0: (v4 / (((v4 ^ (v4 % ((v4) | 1)))) | 1)), F1: 300}
    if true {
        switch v4 {
        case 9223372036854775807:
            v3 = append(v3, 978)
        default:
            var v8 int = (v4 * v4)
            _ = v8
        }
        fmt.Println(v4, v6)
        if false {
            var v9 int = (v4 << ((0) & 63))
            _ = v9
        } else {
            var v10 float32 = 0.0
            var v11 uint64 = 4294967295
            _ = v10
            _ = v11
        }
    } else {
        v6 += 2
        v6 = (v6 ^ (v6 << ((255) & 63)))
        for fi12 := 0; fi12 < 7; fi12++ {
            _ = len(v3)
            v13 := GenS1{F0: int(v6), F1: -2147483648}
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: v4, F1: 248}
        var v16 uint64 = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 float64 = (0.0 + 0.1)
    if (("hello" + "a") < "hello") {
        v6 = 88
        if false {
            fmt.Println(v17)
            v4 -= (v4 << ((30) & 63))
            fmt.Println(v4, v17)
        } else {
            switch "" {
            case "foo bar":
                fmt.Println(v6, v6, v17)
            default:
                v6 -= v6
            }
            if (v4 <= int(v6)) {
                if true {
                    var v18 int32 = 42
                    var v19 string = "\n\t"
                    switch "hello" {
                    case "":
                        v18 += (^(v18))
                    case "hello":
                        v20 := GenS2{F0: true, F1: -314, F2: 257}
                        _ = v20
                    default:
                        _ = len(v3)
                    }
                    _ = v18
                    _ = v19
                }
                var v21 bool = false
                _ = v21
            } else {
                v17 -= v17
                v17 += 1e10
                for fi22 := 0; fi22 < 2; fi22++ {
                    if ((false && ("foo bar" != "foo bar")) && (80 <= 100)) {
                        var v23 int = 535
                        fmt.Println(v23, v17)
                        _ = len(v2)
                        _ = v23
                    } else {
                        v17 = v17
                        for fi24 := 0; fi24 < 13; fi24++ {
                            v25 := GenS1{F0: v4, F1: 818}
                            v4 = v4
                            v2["\n\t"] = (^(v6))
                            _ = fi24
                            _ = v25
                        }
                        if false {
                            if (!(-169 != 689)) {
                                v3 = append(v3, -381)
                                v26 := GenS1{F0: fi22, F1: 2147483647}
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    delete(v1, v6)
                                    _ = fi27
                                }
                                _ = v26
                            } else {
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    v29 := GenS3{F0: 0, F1: 127}
                                    _ = fi28
                                    _ = v29
                                }
                            }
                            for fi30 := 0; fi30 < 8; fi30++ {
                                v6 += v6
                                v31 := GenS2{F0: true, F1: 404, F2: 506}
                                _ = fi30
                                _ = v31
                            }
                            var v32 int8 = 100
                            _ = v32
                        }
                    }
                    v33 := GenS4{F0: -535}
                    _ = fi22
                    _ = v33
                }
            }
            v17 = (float64(3.14159) / float64((-2.25 - 1.5)))
        }
    }
    if false {
        if (14 < 127) {
            var v34 int32 = 22
            if (3.14159 < (float32(((1.5 / (float64(3.14159) * (v17 - (v17 + 0.1)))) / 1e10)) / (0.1 - (0.0 - (0.1 - 1.5))))) {
                v4 -= (v4 ^ 587)
                if false {
                    var v35 float64 = v17
                    _ = v35
                } else {
                    if len(v3) > 0 {
                        v3[0] = v4
                    }
                    fmt.Println(v34, v6)
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        if (0 <= 707) {
                            for fi36 := 0; fi36 < 5; fi36++ {
                                fmt.Println(v34, fi36)
                                if (-32768 < 0) {
                                    if (!(!false)) {
                                        v17 += v17
                                    } else {
                                        if ((int64(v6) < 42) && true) {
                                            var v37 uint = 593
                                            v38 := v2[""]
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    fmt.Println(v4, v17, fi36)
                                    fmt.Println(fi36, v4, fi36)
                                }
                                _ = fi36
                            }
                            for fi39 := 0; fi39 < 7; fi39++ {
                                for fi40 := 0; fi40 < 9; fi40++ {
                                    var v41 bool = (9223372036854775807 > 1)
                                    fmt.Println(v17)
                                    v4 -= fi40
                                    _ = fi40
                                    _ = v41
                                }
                                fmt.Println(fi39, fi39, v6)
                                _ = fi39
                            }
                            for fi42 := 0; fi42 < 3; fi42++ {
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    var v44 uint16 = 0
                                    v17 -= ((v17 - (((v17 * float64((((-2.25 - ((0.1 / -2.25) * 0.0)) / 0.1) / 0.1))) / 3.14159) * v17)) - (v17 / v17))
                                    _ = fi43
                                    _ = v44
                                }
                                v45 := GenS1{F0: 32767, F1: (v34 - -612)}
                                for fi46 := 0; fi46 < 13; fi46++ {
                                    v47 := GenS1{F0: fi46, F1: (v34 % ((v34) | 1))}
                                    var v48 uint32 = 2147483647
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                                _ = fi42
                                _ = v45
                            }
                        } else {
                            var v49 uint16 = 936
                            v50 := v1[v6]
                            fmt.Println(v4)
                            _ = v49
                            _ = v50
                        }
                    default:
                        for fi51 := 0; fi51 < 14; fi51++ {
                            var v52 bool = (!true)
                            switch "hello" {
                            case "\n\t":
                                var v53 bool = ((!v52) && false)
                                _ = v53
                            case "foo bar":
                                v54 := GenS5{F0: 132, F1: -9223372036854775808}
                                _ = v54
                            }
                            _ = fi51
                            _ = v52
                        }
                    }
                }
                for fi55 := 0; fi55 < 7; fi55++ {
                    var v56 uint = 599
                    _ = fi55
                    _ = v56
                }
            } else {
                if (-623 >= 172) {
                    v6 = 2147483648
                } else {
                    v6 -= 2147483647
                    for fi57 := 0; fi57 < 7; fi57++ {
                        for fi58 := 0; fi58 < 20; fi58++ {
                            var v59 uint = 0
                            var v60 int8 = -128
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                        if false {
                            var v61 string = "\n\t"
                            var v62 float32 = float32(v17)
                            _ = v61
                            _ = v62
                        }
                        if ((v6 - (^(v6))) != 4294967295) {
                            for fi63 := 0; fi63 < 16; fi63++ {
                                delete(v2, ("a" + "a"))
                                _ = fi63
                            }
                            v64 := GenS1{F0: (fi57 & 119), F1: (^(v34))}
                            _ = v64
                        }
                        _ = fi57
                    }
                    v65 := v2[("\n\t" + "hello")]
                    _ = v65
                }
            }
            v6 += (v6 >> ((65535) & 63))
            _ = v34
        } else {
            v17 -= float64(float32(1e10))
            switch "\n\t" {
            case "\n\t":
                var v66 uint16 = 574
                _ = v66
            default:
                v67 := GenS6{F0: 9223372036854775807}
                _ = v67
            }
        }
    } else {
        var v68 bool = (1 < 42)
        v4 = 543
        var v69 uint8 = 42
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 15; fi70++ {
        fmt.Println(v17, fi70, v17)
        _ = fi70
    }
    for fi71 := 0; fi71 < 15; fi71++ {
        fmt.Println(v4, fi71, fi71)
        var v72 uint32 = v6
        switch "a" {
        case "":
            var v73 uint = 18446744073709551615
            _ = v73
        }
        _ = fi71
        _ = v72
    }
    var v74 uint8 = 135
    var v75 int64 = 42
    var v76 uint64 = 258
    for fi77 := 0; fi77 < 13; fi77++ {
        for fi78 := 0; fi78 < 7; fi78++ {
            var v79 int = (fi78 * (fi78 ^ (^(fi77))))
            _ = fi78
            _ = v79
        }
        var v80 int8 = 127
        v75 += (v75 | 42)
        _ = fi77
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v75
    _ = v76
}
