package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int16; F2 int }
type GenS6 struct { F0 string; F1 uint8; F2 int16 }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 uint32 = 618
    v2 := []int{-2147483648}
    v3 := make(map[int]uint32)
    if false {
        v1 -= v1
        for fi4 := 0; fi4 < 15; fi4++ {
            for fi5 := 0; fi5 < 8; fi5++ {
                v6 := GenS1{F0: 18446744073709551615, F1: -1, F2: 1}
                var v7 uint8 = 0
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
    }
    if (1 >= -1) {
        v8 := GenS2{F0: -1, F1: ("a" + "\n\t")}
        var v9 float64 = (-2.25 - (-2.25 - float64(1.5)))
        _ = v8
        _ = v9
    }
    if len(v2) > 0 {
        v2[0] = -386
    }
    fmt.Println(v1)
    switch 1 {
    case 719:
        switch "hello" {
        case "":
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    _ = len(v3)
                    _ = fi11
                }
                if false {
                    fmt.Println(fi10, fi10)
                    v1 = v1
                    var v12 uint = 255
                    _ = v12
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1 -= (v1 << ((65535) & 63))
                    var v13 int32 = 621
                    _ = v13
                }
                v1 += (v1 - v1)
                _ = fi10
            }
        case "\n\t":
            v1 -= v1
        default:
            for fi14 := 0; fi14 < 6; fi14++ {
                if true {
                    fmt.Println(v1, fi14)
                    var v15 int = (fi14 - fi14)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if ((true && true) && false) {
                            v17 := GenS3{F0: 0}
                            v18 := GenS1{F0: 130, F1: 1, F2: 38}
                            _ = v17
                            _ = v18
                        } else {
                            var v19 float32 = 3.14159
                            switch v15 {
                            case 32768:
                                var v20 string = (("\n\t" + "") + ("foo bar" + ("a" + ("" + ""))))
                                v19 += -2.25
                                switch fi16 {
                                case 128:
                                    var v21 int16 = -746
                                    _ = v21
                                default:
                                    if false {
                                        if false {
                                            v22 := GenS4{F0: uint16(v1)}
                                            v23 := GenS3{F0: 127}
                                            v19 -= (((3.14159 + (v19 * (((v19 * float32(0.0)) * v19) / (3.14159 + -2.25)))) * float32(1.5)) / (v19 + v19))
                                            _ = v22
                                            _ = v23
                                        } else {
                                            v1 -= v1
                                            v1 = v1
                                            var v24 uint8 = 198
                                            if len(v2) > 0 {
                                                v2[0] = 32767
                                            }
                                            var v25 uint32 = v1
                                            _ = v24
                                            _ = v25
                                        }
                                    } else {
                                        if false {
                                            v26 := GenS5{F0: 142, F1: -565, F2: (fi16 % ((v15) | 1))}
                                            v27 := GenS6{F0: "foo bar", F1: 255, F2: -119}
                                            fmt.Println(v20, v19, fi14)
                                            _ = v26
                                            _ = v27
                                        } else {
                                            fmt.Println(fi14)
                                        }
                                        if false {
                                            if (!(true && (!true))) {
                                                var v28 int32 = -1
                                                _ = v28
                                            } else {
                                                var v29 bool = false
                                                _ = v29
                                            }
                                        }
                                    }
                                }
                                _ = v20
                            case 42:
                                var v30 bool = (!(2147483647 < -1))
                                _ = v30
                            }
                            _ = v19
                        }
                        v1 += v1
                        var v31 int = (^(v15))
                        _ = fi16
                        _ = v31
                    }
                    _ = v15
                } else {
                    fmt.Println(v1, v1)
                    v1 = v1
                    if true {
                        v1 = 3
                    }
                }
                if false {
                    if (!false) {
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v33 := v3[(-(fi32))]
                            _ = fi32
                            _ = v33
                        }
                        v1 += (v1 << ((4294967295) & 63))
                        var v34 int = (-(fi14))
                        _ = v34
                    } else {
                        var v35 bool = (false || ((!((fi14 >> ((2) & 63)) == 32768)) || (!(1 == 1))))
                        var v36 float32 = 0.1
                        _ = v35
                        _ = v36
                    }
                    if (!false) {
                        v3[(^(fi14))] = (v1 & v1)
                    } else {
                        fmt.Println(v1, v1)
                        var v37 string = "foo bar"
                        for fi38 := 0; fi38 < 13; fi38++ {
                            v1 = v1
                            if len(v2) > 0 {
                                v2[0] = fi14
                            }
                            switch -753 {
                            case -607:
                                fmt.Println(v37, v37, fi38)
                            }
                            _ = fi38
                        }
                        _ = v37
                    }
                } else {
                    v1 = (^(v1))
                    for fi39 := 0; fi39 < 7; fi39++ {
                        var v40 string = "\n\t"
                        _ = fi39
                        _ = v40
                    }
                }
                var v41 int = fi14
                _ = fi14
                _ = v41
            }
        }
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v43 := GenS7{F0: 495, F1: 17, F2: 0}
        v44 := GenS1{F0: 1, F1: 734, F2: 128}
        var v45 int8 = 0
        _ = fi42
        _ = v43
        _ = v44
        _ = v45
    }
    _ = len(v2)
    if (v1 <= 2147483648) {
        var v46 int64 = 9223372036854775807
        var v47 float32 = (float32(0.1) * 1e10)
        _ = v46
        _ = v47
    } else {
        v1 = 324
    }
    v1 -= (v1 - v1)
    switch "a" {
    case "hello":
        v1 = v1
    case "":
        var v48 int16 = 42
        _ = v48
    case "\n\t":
        for fi49 := 0; fi49 < 7; fi49++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v50 := GenS8{F0: -1}
            fmt.Println(fi49, fi49)
            _ = fi49
            _ = v50
        }
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        switch ("hello" + ("hello" + "\n\t")) {
        case "hello":
            var v52 uint32 = (v1 % ((v1) | 1))
            _ = v52
        case "":
            v2 = append(v2, fi51)
        default:
            if (18446744073709551615 <= 195) {
                v1 = (v1 - (v1 & v1))
                for fi53 := 0; fi53 < 2; fi53++ {
                    v54 := GenS9{F0: 128}
                    v1 = v1
                    _ = fi53
                    _ = v54
                }
            } else {
                v1 -= (v1 | (v1 - v1))
                var v55 float32 = ((1.5 / 1.5) + 1.5)
                _ = v55
            }
        }
        switch (fi51 - (fi51 % ((fi51) | 1))) {
        case -945:
            var v56 uint64 = 42
            _ = v56
        default:
            var v57 int16 = 1000
            _ = v57
        }
        _ = fi51
    }
    var v58 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
}
