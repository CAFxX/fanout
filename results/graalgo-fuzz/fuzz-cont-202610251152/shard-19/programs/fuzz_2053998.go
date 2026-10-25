package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint; F1 int16 }

func main() {
    var v1 string = ""
    var v2 string = v1
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    v4 += 9223372036854775807
    fmt.Println(v1, v1, v1)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int64 = v4
        if false {
            var v7 uint64 = 2
            _ = v7
        }
        var v8 uint = 0
        _ = fi5
        _ = v6
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: true}
    switch (v2 + "a") {
    case "foo bar":
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 103, F1: ((!false) && (127 == 255))}
            _ = fi10
            _ = v11
        }
    case "\n\t":
        v4 = v4
    }
    fmt.Println(v1, v1)
    switch 128 {
    case 32767:
        switch 4294967295 {
        case 0:
            if (118 > 2) {
                fmt.Println(v4, v1, v4)
                var v12 int32 = 709
                fmt.Println(v4, v12)
                _ = v12
            } else {
                if (((740 < 2147483647) || true) || (922 == 0)) {
                    var v13 int32 = 42
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        fmt.Println(v4, v2, v4)
                        var v15 float32 = (((0.0 / 1.5) * (3.14159 * 0.1)) - (((0.0 - 0.0) * (0.1 - 3.14159)) * -2.25))
                        _ = fi14
                        _ = v15
                    }
                    _ = len(v3)
                    v16 := []int{127}
                    _ = v16
                }
            }
        }
    case 0:
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi17)
            _ = fi17
        }
    }
    if (-528 == -32768) {
        fmt.Println(v2, v4, v2)
        var v18 float32 = (float32((float64(0.1) * float64(-2.25))) / (0.1 * 1.5))
        var v19 float32 = v18
        _ = v18
        _ = v19
    }
    switch 1 {
    case 9223372036854775807:
        if (!false) {
            v20 := []int{-659, 224, 42, 945}
            v1 = v1
            if (767 > 32768) {
                v2 = (("" + v1) + ("a" + v1))
            } else {
                v1 = "hello"
                v21 := GenS1{F0: 255, F1: (!false)}
                _ = v21
            }
            _ = v20
        } else {
            v22 := []int{-786, 9223372036854775807}
            _ = v22
        }
    case 468:
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int = fi23
            _ = fi23
            _ = v24
        }
    case 2147483647:
        v1 = (v1 + v2)
    }
    fmt.Println(v1, v2, v4)
    for fi25 := 0; fi25 < 7; fi25++ {
        switch (v1 + (v2 + ((v2 + ("foo bar" + "foo bar")) + v1))) {
        case "":
            switch fi25 {
            case 127:
                var v26 string = ""
                _ = v26
            case 2147483647:
                v27 := GenS1{F0: 1, F1: (true || (true && false))}
                _ = v27
            case -625:
                fmt.Println(v4)
            default:
                fmt.Println(v4, v2, v2)
            }
        case "\n\t":
            fmt.Println(v2)
        }
        if false {
            switch fi25 {
            case 2147483647:
                v3[4294967295] = (v1 + "a")
            case 32768:
                fmt.Println(v4)
            case 32767:
                v4 -= (v4 - (v4 + (v4 | v4)))
            default:
                var v28 int32 = 78
                _ = v28
            }
        } else {
            if true {
                if (false && true) {
                    if true {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 int = (^(fi29))
                            v31 := GenS2{F0: 42, F1: -1}
                            if (true && false) {
                                var v32 float64 = 0.0
                                if ((1e10 < v32) && (65535 < 2)) {
                                    var v33 uint64 = 365
                                    if (!(v32 == 1.5)) {
                                        v2 = "\n\t"
                                        var v34 int32 = -462
                                        if true {
                                            v3[42] = ("foo bar" + ("hello" + "a"))
                                            v30 += fi25
                                        }
                                        v35 := v3[4294967295]
                                        _ = v34
                                        _ = v35
                                    }
                                    v33 = 18446744073709551615
                                    _ = v33
                                } else {
                                    switch v2 {
                                    case "a":
                                        fmt.Println(v1)
                                    case "\n\t":
                                        switch (("\n\t" + "\n\t") + v1) {
                                        case "hello":
                                            var v36 uint64 = 2
                                            _ = v36
                                        case "":
                                            fmt.Println(v30, v30)
                                        default:
                                            v37 := GenS2{F0: int8(v30), F1: 1}
                                            _ = v37
                                        }
                                    case "foo bar":
                                        v4 = (v4 / (((-(v4))) | 1))
                                    default:
                                        v38 := []int{v30, (v30 | 128)}
                                        _ = v38
                                    }
                                    var v39 uint = 1
                                    _ = v39
                                }
                                if false {
                                    switch v30 {
                                    case 9223372036854775807:
                                        var v40 int32 = 2147483647
                                        _ = v40
                                    case 32768:
                                        switch (fi25 / ((fi25) | 1)) {
                                        case -38:
                                            var v41 int = fi25
                                            _ = v41
                                        }
                                    case 423:
                                        fmt.Println(v32, fi29, v32)
                                    default:
                                        v42 := GenS3{F0: (v30 % ((fi29) | 1))}
                                        _ = v42
                                    }
                                    v30 += fi29
                                    v43 := v3[2147483648]
                                    _ = v43
                                }
                                _ = v32
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        if (1000 != -1) {
                            v44 := GenS4{F0: v4, F1: 32767}
                            switch 128 {
                            case 32767:
                                fmt.Println(v1, v1)
                            case -2147483648:
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    if (!true) {
                                        v1 = ("hello" + (v1 + ""))
                                        v46 := []int{}
                                        _ = v46
                                    } else {
                                        v47 := []int{fi45, (fi25 / (((fi25 ^ 4294967295)) | 1))}
                                        if true {
                                            var v48 int32 = int32(fi45)
                                            _ = v48
                                        } else {
                                            fmt.Println(fi45, fi25)
                                            var v49 float32 = float32(0.1)
                                            var v50 uint = 1
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v47
                                    }
                                    v1 = (v2 + "\n\t")
                                    _ = fi45
                                }
                            case 761:
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    var v52 int = fi51
                                    fmt.Println(fi25, v2)
                                    var v53 uint16 = 0
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                            default:
                                var v54 float64 = 3.14159
                                _ = v54
                            }
                            _ = v44
                        }
                    }
                    for fi55 := 0; fi55 < 14; fi55++ {
                        switch (^(fi55)) {
                        case 127:
                            v56 := GenS5{F0: -128}
                            _ = v56
                        default:
                            var v57 float64 = 1e10
                            _ = v57
                        }
                        _ = fi55
                    }
                }
            } else {
                if (v4 > -1) {
                    if ((!false) && true) {
                        var v58 float64 = (0.0 * -2.25)
                        switch 2147483647 {
                        case 1:
                            var v59 uint = 4294967295
                            _ = v59
                        case -1:
                            for fi60 := 0; fi60 < 2; fi60++ {
                                v58 += ((float64((0.0 - -2.25)) - ((v58 - v58) / v58)) / (float64((0.0 - (1.5 / (0.1 / 0.1)))) - (float64((-2.25 - 0.0)) - (v58 + v58))))
                                fmt.Println(v1)
                                v4 = v4
                                _ = fi60
                            }
                        default:
                            v61 := GenS1{F0: 255, F1: (!false)}
                            _ = v61
                        }
                        _ = v58
                    }
                    switch -1 {
                    case 9223372036854775807:
                        v4 += (v4 + (v4 * (v4 - int64(fi25))))
                    case -1:
                        v62 := GenS5{F0: -128}
                        _ = v62
                    }
                } else {
                    v63 := GenS1{F0: 42, F1: (2 == 18446744073709551615)}
                    _ = v63
                }
            }
            fmt.Println(fi25, v1, fi25)
        }
        fmt.Println(fi25)
        _ = fi25
    }
    if false {
        fmt.Println(v2, v1, v1)
        v3[0] = ((v2 + ((v1 + v2) + (v1 + (v1 + v1)))) + "\n\t")
        if true {
            v64 := GenS6{F0: 32767, F1: 269, F2: 2147483647}
            v65 := GenS5{F0: 1}
            _ = v64
            _ = v65
        } else {
            var v66 uint64 = 29
            _ = v66
        }
    }
    var v67 uint32 = 351
    v2 = v1
    switch v2 {
    case "hello":
        v68 := []int{127}
        _ = v68
    case "":
        var v69 int16 = 316
        _ = v69
    }
    if true {
        fmt.Println(v2, v4, v1)
        switch (v2 + (("a" + (v2 + v1)) + ((((v2 + v1) + v1) + v1) + v2))) {
        case "a":
            v70 := GenS7{F0: 42, F1: 1000}
            _ = v70
        case "":
            var v71 float32 = (0.1 - 1.5)
            _ = v71
        default:
            v72 := GenS7{F0: 4294967295, F1: 0}
            _ = v72
        }
        if (false || (!(-32768 == 461))) {
            delete(v3, v67)
            v67 = v67
        }
    } else {
        var v73 bool = true
        _ = v73
    }
    v1 = (v1 + "a")
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v9
}
