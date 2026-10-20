package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int16; F1 uint32; F2 int64 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]int)
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    switch 32767 {
    case 0:
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := GenS1{F0: 4294967295, F1: -1, F2: uint16(v1)}
            var v7 float64 = 1.5
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    delete(v4, 4294967295)
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v3, v3)
        if true {
            if (v1 == 2) {
                switch ("foo bar" + "a") {
                case "foo bar":
                    v3 -= (v3 + (float32(float64(v3)) - v3))
                default:
                    var v9 uint = 4294967295
                    _ = v9
                }
            } else {
                v1 = v1
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 int16 = 1000
                    v1 -= (v1 / (((v1 | v1)) | 1))
                    v3 -= v3
                    _ = fi10
                    _ = v11
                }
                var v12 uint = 904
                _ = v12
            }
            var v13 bool = (!false)
            _ = v13
        }
        switch "a" {
        case "":
            var v14 int = fi8
            _ = v14
        case "a":
            var v15 int32 = 42
            _ = v15
        case "foo bar":
            var v16 float32 = v3
            _ = v16
        }
        _ = fi8
    }
    v17 := []int{-9223372036854775808, -2147483648, 42, 127}
    v17 = append(v17, 127)
    var v18 int16 = -392
    v19 := GenS2{F0: 42, F1: 127, F2: 1}
    var v20 uint64 = 620
    for fi21 := 0; fi21 < 13; fi21++ {
        if true {
            if (!false) {
                v18 += (^(v18))
            }
            v22 := GenS3{F0: (384 > 0), F1: fi21, F2: (v18 + v18)}
            var v23 int64 = 0
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 uint8 = (^(v1))
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v27 int = 1
    if (false || (577 >= 0)) {
        var v28 int32 = 0
        _ = len(v17)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 15; fi29++ {
            switch (^(v27)) {
            case 127:
                var v30 uint32 = 2
                _ = v30
            default:
                _ = len(v17)
            }
            if false {
                v31 := GenS2{F0: 2, F1: v27, F2: (v1 >> ((848) & 63))}
                _ = v31
            }
            switch "hello" {
            case "a":
                fmt.Println(v3, fi29)
            default:
                if (611 > 944) {
                    if (!((v20 - v20) < (v20 & v20))) {
                        v32 := GenS4{F0: 127, F1: 753, F2: v20}
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v27 = fi29
                            _ = fi33
                        }
                        _ = v32
                    } else {
                        v18 = (v18 - v18)
                        switch ("\n\t" + "a") {
                        case "":
                            if false {
                                for fi34 := 0; fi34 < 5; fi34++ {
                                    v35 := GenS2{F0: 255, F1: (fi34 * (fi29 % ((4294967295) | 1))), F2: (v1 & (v1 >> ((4294967295) & 63)))}
                                    v3 = v3
                                    switch fi29 {
                                    case -465:
                                        if (((v18 >> ((968) & 63)) > 139) || (3.14159 >= v3)) {
                                            if (!true) {
                                                v27 += fi29
                                                switch ("" + "hello") {
                                                case "hello":
                                                    if false {
                                                        _ = len(v4)
                                                    }
                                                    v17 = append(v17, v27)
                                                    v36 := GenS5{F0: ("\n\t" + "foo bar")}
                                                    if true {
                                                        v17 = append(v17, (^(v27)))
                                                        var v37 int32 = -1
                                                        _ = len(v2)
                                                        _ = v37
                                                    }
                                                    _ = v36
                                                default:
                                                    var v38 float32 = (0.0 + ((v3 + v3) / v3))
                                                    _ = v38
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v20 += (v20 / ((1) | 1))
                                                case "\n\t":
                                                    var v39 float32 = -2.25
                                                    _ = v39
                                                default:
                                                    var v40 int8 = 10
                                                    _ = v40
                                                }
                                            }
                                            if true {
                                                if (("hello" != "foo bar") && (!(!true))) {
                                                    fmt.Println(v3, v1, v3)
                                                    switch "" {
                                                    case "hello":
                                                        if true {
                                                            var v41 int8 = 100
                                                            v42 := GenS1{F0: 4294967295, F1: (v18 / (((v18 + (v18 << ((749) & 63)))) | 1)), F2: 1}
                                                            v20 += v20
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    default:
                                                        v43 := GenS4{F0: (v1 << ((183) & 63)), F1: 35, F2: v20}
                                                        _ = v43
                                                    }
                                                }
                                                if true {
                                                    v4[(^(fi29))] = 65535
                                                }
                                            }
                                        }
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                                for fi44 := 0; fi44 < 13; fi44++ {
                                    var v45 int = (v27 | (fi29 | fi29))
                                    if ((v18 >> ((0) & 63)) <= v18) {
                                        var v46 string = "foo bar"
                                        _ = v46
                                    }
                                    fmt.Println(fi29, v3)
                                    _ = fi44
                                    _ = v45
                                }
                            } else {
                                _ = len(v2)
                                if true {
                                    v2[(v27 + v27)] = v27
                                    for fi47 := 0; fi47 < 16; fi47++ {
                                        fmt.Println(v3, v3)
                                        _ = fi47
                                    }
                                } else {
                                    v48 := GenS6{F0: 65535}
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        var v50 int64 = 9223372036854775807
                                        v50 -= v50
                                        _ = fi49
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 3; fi51++ {
                                        v52 := GenS7{F0: -87}
                                        _ = fi51
                                        _ = v52
                                    }
                                    _ = v48
                                }
                                v1 -= 19
                            }
                        default:
                            v53 := v2[v27]
                            _ = v53
                        }
                        switch (^(fi29)) {
                        case 963:
                            v54 := GenS8{F0: false, F1: 127}
                            _ = v54
                        }
                    }
                    if true {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        delete(v4, (fi29 | -909))
                    } else {
                        if false {
                            v3 += 1.5
                            var v55 float32 = float32(-2.25)
                            for fi56 := 0; fi56 < 10; fi56++ {
                                v18 = v18
                                _ = fi56
                            }
                            _ = v55
                        } else {
                            if true {
                                v57 := GenS9{F0: v18, F1: 2147483648, F2: 42}
                                var v58 bool = ((v1 != v1) && (!(1 <= 0)))
                                _ = v57
                                _ = v58
                            } else {
                                v3 = v3
                            }
                            v59 := v2[v27]
                            _ = v59
                        }
                        v60 := GenS7{F0: 127}
                        _ = v60
                    }
                    var v61 uint16 = 42
                    _ = v61
                } else {
                    var v62 uint16 = 1
                    v17 = append(v17, v27)
                    _ = v62
                }
            }
            _ = fi29
        }
        var v63 string = (("\n\t" + ("foo bar" + "")) + "hello")
        _ = v63
    }
    if true {
        v27 = v27
        v27 -= v27
        fmt.Println(v18, v3, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
