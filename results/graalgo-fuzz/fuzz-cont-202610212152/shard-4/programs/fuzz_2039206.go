package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int8; F1 bool; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 int; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 int64 = 1
    var v4 int64 = v3
    var v5 uint16 = 935
    if false {
        v5 -= v5
    } else {
        if true {
            fmt.Println(v5)
        } else {
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := []int{32768, (fi6 - 1)}
                switch fi6 {
                case -8:
                    v4 = v4
                case -1:
                    var v8 float64 = ((1.5 + 1.5) + 0.1)
                    _ = v8
                case 2147483647:
                    var v9 int8 = 42
                    _ = v9
                default:
                    v5 -= v5
                }
                _ = fi6
                _ = v7
            }
            if true {
                v10 := []int{284, -173}
                _ = v10
            } else {
                if (false || false) {
                    delete(v1, 42)
                    v4 = 0
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int8 = int8(fi11)
                    _ = fi11
                    _ = v12
                }
                var v13 bool = true
                _ = v13
            }
            switch -2147483648 {
            case 9223372036854775807:
                if false {
                    switch 1 {
                    case 128:
                        v14 := GenS1{F0: -117, F1: "hello", F2: 61}
                        _ = v14
                    case 2147483647:
                        v1[0] = 2
                    case 0:
                        v15 := GenS2{F0: 42, F1: -1}
                        _ = v15
                    default:
                        v16 := GenS2{F0: -28, F1: 25}
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        var v18 int32 = 618
                        if (true && (!false)) {
                            fmt.Println(v3, fi17, fi17)
                        } else {
                            v1[4294967295] = 1
                            if (42 < 0) {
                                fmt.Println(v3, v18, v18)
                            } else {
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    switch (fi19 | (^(fi19))) {
                                    case 125:
                                        v5 -= 847
                                    case -2147483648:
                                        delete(v1, 2)
                                    case -707:
                                        v5 -= (v5 | (v5 * v5))
                                    default:
                                        fmt.Println(fi17, v18)
                                    }
                                    var v20 int8 = 42
                                    v18 = (v18 + v18)
                                    _ = fi19
                                    _ = v20
                                }
                                v3 -= v3
                            }
                        }
                        _ = fi17
                        _ = v18
                    }
                    fmt.Println(v3)
                }
            case 860:
                switch "a" {
                case "\n\t":
                    v21 := GenS3{F0: 2, F1: "", F2: 1000}
                    _ = v21
                case "foo bar":
                    v4 = (v4 + v4)
                case "hello":
                    var v22 uint8 = 78
                    _ = v22
                default:
                    var v23 uint16 = v5
                    _ = v23
                }
            case 838:
                var v24 int32 = -154
                _ = v24
            }
        }
    }
    fmt.Println(v3)
    v25 := GenS3{F0: 500, F1: "a", F2: 42}
    if (0 != 0) {
        var v26 int64 = (v3 << ((2) & 63))
        v2[-105] = "hello"
        fmt.Println(v5, v4, v3)
        _ = v26
    } else {
        switch ("foo bar" + ("hello" + ("" + "\n\t"))) {
        case "foo bar":
            v4 -= (v3 & (^(v3)))
        }
    }
    v27 := v1[2]
    var v28 bool = (!true)
    var v29 uint8 = 128
    v1[(v27 - v27)] = v27
    for fi30 := 0; fi30 < 10; fi30++ {
        if ("\n\t" > "") {
            v31 := v1[65535]
            _ = v31
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            v3 -= (v4 | (v3 * -1))
            _ = fi32
        }
        _ = fi30
    }
    v27 += uint32(v29)
    var v33 int64 = (v4 % ((int64(v29)) | 1))
    switch -2147483648 {
    case -2147483648:
        if v28 {
            v34 := GenS4{F0: v28, F1: 956}
            var v35 uint8 = v29
            _ = v34
            _ = v35
        } else {
            v36 := []int{-280, 127}
            var v37 int64 = v4
            var v38 uint16 = (v5 - 2)
            _ = v36
            _ = v37
            _ = v38
        }
    case 32767:
        v28 = (v28 || ("hello" >= ""))
    case -308:
        for fi39 := 0; fi39 < 3; fi39++ {
            delete(v1, v27)
            var v40 bool = true
            _ = fi39
            _ = v40
        }
    }
    switch 111 {
    case 0:
        v41 := GenS3{F0: (v27 / ((0) | 1)), F1: "\n\t", F2: -32768}
        _ = v41
    case 32768:
        if v28 {
            v28 = v28
            if false {
                if ((^(v27)) > (v27 << ((255) & 63))) {
                    if true {
                        fmt.Println(v27, v27)
                    }
                    if v28 {
                        switch 1 {
                        case -2147483648:
                            var v42 uint64 = 0
                            _ = v42
                        }
                        var v43 int16 = 1
                        v2[4294967295] = "foo bar"
                        _ = v43
                    } else {
                        if (!false) {
                            v44 := v2[-2147483648]
                            _ = v44
                        }
                    }
                    v33 = v3
                } else {
                    v45 := []int{int(v4), -863}
                    v29 = (v29 - v29)
                    if v28 {
                        v46 := GenS5{F0: 0, F1: true, F2: 127}
                        switch "a" {
                        case "":
                            for fi47 := 0; fi47 < 5; fi47++ {
                                switch "" {
                                case "foo bar":
                                    v48 := GenS1{F0: 127, F1: ("" + "hello"), F2: (-(fi47))}
                                    _ = v48
                                }
                                v4 = (v4 + v33)
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    if (!(float64(1.5) != (0.1 - 0.0))) {
                                        v50 := GenS6{F0: 52, F1: 2}
                                        v51 := GenS7{F0: (v29 | v29)}
                                        v1[(v27 << ((652) & 63))] = (v27 * 2)
                                        v52 := GenS8{F0: (fi49 & (-(fi49))), F1: (fi47 ^ 0), F2: 254}
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = fi49
                                }
                                _ = fi47
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            if len(v45) > 0 {
                                v45[0] = 2147483647
                            }
                        case "a":
                            var v53 int8 = 42
                            _ = v53
                        case "":
                            if (true || true) {
                                var v54 float32 = -2.25
                                for fi55 := 0; fi55 < 4; fi55++ {
                                    var v56 string = (("\n\t" + "\n\t") + ("\n\t" + (("\n\t" + "a") + ("foo bar" + ("foo bar" + "hello")))))
                                    _ = fi55
                                    _ = v56
                                }
                                _ = v54
                            }
                        default:
                            var v57 string = "a"
                            _ = v57
                        }
                        _ = v46
                    } else {
                        for fi58 := 0; fi58 < 1; fi58++ {
                            var v59 int32 = -2147483648
                            v28 = true
                            var v60 int8 = -1
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                        var v61 int8 = 127
                        v33 -= v4
                        _ = v61
                    }
                    _ = v45
                }
                fmt.Println(v27, v28)
            } else {
                for fi62 := 0; fi62 < 8; fi62++ {
                    if (!(1.5 == 0.0)) {
                        fmt.Println(v33, v29, v5)
                        _ = len(v1)
                        if false {
                            v63 := GenS6{F0: v5, F1: 286}
                            _ = v63
                        } else {
                            v27 -= v27
                        }
                    } else {
                        _ = len(v1)
                        v5 = v5
                    }
                    switch "hello" {
                    case "":
                        v4 -= v4
                    }
                    _ = fi62
                }
            }
            if v28 {
                if v28 {
                    if v28 {
                        v64 := []int{-417}
                        var v65 int = 246
                        _ = v64
                        _ = v65
                    }
                    if false {
                        var v66 int8 = 0
                        for fi67 := 0; fi67 < 12; fi67++ {
                            var v68 uint = 239
                            v69 := v1[(v27 * 65535)]
                            switch "hello" {
                            case "foo bar":
                                _ = len(v2)
                            }
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                        _ = v66
                    } else {
                        for fi70 := 0; fi70 < 11; fi70++ {
                            v71 := GenS2{F0: 0, F1: -102}
                            _ = fi70
                            _ = v71
                        }
                        v29 = uint8(v5)
                    }
                    v4 -= (v3 % ((v3) | 1))
                }
            }
        }
    case -877:
        delete(v1, 81)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
