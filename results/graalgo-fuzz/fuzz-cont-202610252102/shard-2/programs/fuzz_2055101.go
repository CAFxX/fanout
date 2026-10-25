package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := []int{-2147483648, 32767}
    var v4 uint16 = 42
    switch "" {
    case "":
        if (3.14159 < 0.0) {
            delete(v1, 42)
        }
    case "a":
        fmt.Println(v4)
    case "foo bar":
        var v5 int64 = 9223372036854775807
        _ = v5
    default:
        fmt.Println(v4)
    }
    if (!((-2147483648 >= 1) || (!true))) {
        v6 := GenS1{F0: 4294967295, F1: (true || (853 <= -207))}
        for fi7 := 0; fi7 < 3; fi7++ {
            v4 += v4
            _ = fi7
        }
        var v8 int64 = 42
        _ = v6
        _ = v8
    } else {
        v9 := GenS2{F0: "foo bar", F1: 22, F2: -9223372036854775808}
        _ = v9
    }
    switch 4294967295 {
    case 4294967295:
        if (!(771 <= 42)) {
            if (false && ((v4 + (v4 | (v4 << ((4294967295) & 63)))) >= v4)) {
                v4 += (v4 / ((v4) | 1))
                var v10 int32 = 1
                _ = v10
            } else {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v2[(fi11 + 0)] = (fi11 / ((fi11) | 1))
                    v12 := GenS3{F0: 4294967295, F1: 214}
                    v4 = (v4 & (v4 / (((v4 & (v4 << ((1) & 63)))) | 1)))
                    _ = fi11
                    _ = v12
                }
            }
            var v13 uint32 = 2147483648
            switch "\n\t" {
            case "foo bar":
                var v14 int16 = -761
                _ = v14
            case "a":
                var v15 uint16 = 32767
                _ = v15
            case "hello":
                v16 := GenS1{F0: 127, F1: false}
                _ = v16
            }
            _ = v13
        } else {
            v17 := GenS1{F0: -565, F1: (128 <= 0)}
            v4 += (v4 << ((255) & 63))
            _ = v17
        }
    case 127:
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(fi18, v4)
            v19 := v2[(fi18 & (^(fi18)))]
            for fi20 := 0; fi20 < 17; fi20++ {
                switch "\n\t" {
                case "\n\t":
                    v21 := v1[(fi18 % ((237) | 1))]
                    _ = v21
                case "foo bar":
                    var v22 int8 = -24
                    _ = v22
                case "":
                    var v23 int = (-(v19))
                    _ = v23
                default:
                    switch (fi20 % ((fi20) | 1)) {
                    case 128:
                        if len(v3) > 0 {
                            v3[0] = 128
                        }
                    case 4294967295:
                        if true {
                            var v24 uint16 = v4
                            v3 = append(v3, fi18)
                            var v25 float64 = (((1.5 * 1.5) + (1e10 / -2.25)) - 1.5)
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS4{F0: -9223372036854775808, F1: 932}
                            if (false && true) {
                                var v27 int = (fi20 * (v19 | fi20))
                                if (32768 != v27) {
                                    v1[-1] = 32768
                                    switch (^(v19)) {
                                    case 0:
                                        if len(v3) > 0 {
                                            v3[0] = fi20
                                        }
                                    default:
                                        v2[-1] = -252
                                    }
                                    var v28 float64 = -2.25
                                    if len(v3) > 0 {
                                        v3[0] = v27
                                    }
                                    v28 = ((((v28 + ((((v28 * ((((float64(1e10) - v28) / ((v28 / v28) * 0.1)) / v28) + float64((1.5 + ((((float32(((float64(1.5) - 1e10) * (v28 + v28))) - 1.5) - 0.0) - -2.25) * 3.14159))))) - float64((float32((-2.25 * v28)) * 3.14159))) + (1e10 / v28)) * (v28 - 1.5))) / -2.25) + ((3.14159 / v28) - -2.25)) * 0.1)
                                    v4 -= 72
                                    v29 := GenS4{F0: -2147483648, F1: 0}
                                    v30 := GenS1{F0: v27, F1: false}
                                    v28 = v28
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS5{F0: 0, F1: ("\n\t" + "a")}
                                    var v32 uint16 = v4
                                    _ = v31
                                    _ = v32
                                }
                                if len(v3) > 0 {
                                    v3[0] = -2147483648
                                }
                                _ = v27
                            } else {
                                _ = len(v2)
                                fmt.Println(fi18, fi20)
                                v33 := GenS1{F0: -105, F1: (32767 != 42)}
                                _ = v33
                            }
                            var v34 uint64 = 18446744073709551615
                            _ = v26
                            _ = v34
                        }
                    case 32768:
                        var v35 uint32 = 2147483647
                        _ = v35
                    default:
                        if ((-40 > 116) && (0 == (v4 / (((v4 + (^(v4)))) | 1)))) {
                            if len(v3) > 0 {
                                v3[0] = fi18
                            }
                        }
                    }
                }
                v36 := GenS1{F0: (fi18 | 42), F1: true}
                _ = len(v1)
                switch "hello" {
                case "foo bar":
                    var v37 string = "hello"
                    _ = v37
                case "":
                    delete(v2, fi18)
                case "\n\t":
                    if true {
                        fmt.Println(fi18)
                        var v38 float64 = 0.0
                        _ = v38
                    }
                default:
                    v4 -= v4
                }
                var v39 int64 = 210
                _ = fi20
                _ = v36
                _ = v39
            }
            _ = fi18
            _ = v19
        }
    case -307:
        switch int(v4) {
        case 1:
            var v40 uint16 = (v4 | (v4 ^ (v4 | (v4 * v4))))
            _ = v40
        case 32768:
            v4 = 11
        default:
            if false {
                v4 = (v4 | (v4 * v4))
                var v41 uint16 = v4
                _ = v41
            } else {
                if false {
                    switch "a" {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "a":
                        switch 32767 {
                        case 2147483647:
                            var v42 float32 = 0.0
                            _ = v42
                        case 9223372036854775807:
                            var v43 bool = true
                            _ = v43
                        case -408:
                            _ = len(v2)
                        default:
                            v44 := GenS3{F0: 42, F1: 255}
                            _ = v44
                        }
                    case "hello":
                        for fi45 := 0; fi45 < 11; fi45++ {
                            v4 = v4
                            _ = fi45
                        }
                    }
                    if (!(32767 != 255)) {
                        var v46 int64 = 1
                        _ = len(v3)
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 8; fi47++ {
                        var v48 uint16 = 887
                        _ = fi47
                        _ = v48
                    }
                }
                for fi49 := 0; fi49 < 20; fi49++ {
                    switch fi49 {
                    case -281:
                        v50 := GenS1{F0: fi49, F1: false}
                        _ = v50
                    case 127:
                        if ((fi49 >> ((65535) & 63)) < 42) {
                            v3 = append(v3, fi49)
                            switch ("\n\t" + ("hello" + "a")) {
                            case "":
                                v51 := GenS5{F0: 406, F1: "hello"}
                                _ = v51
                            }
                            v4 = (v4 / (((^(v4))) | 1))
                        }
                    default:
                        if (false || false) {
                            v52 := GenS5{F0: 1, F1: "hello"}
                            var v53 uint = 255
                            _ = v52
                            _ = v53
                        } else {
                            v4 += (v4 + (v4 / ((v4) | 1)))
                            v4 = (v4 - v4)
                            var v54 float32 = (0.0 * ((0.0 * 3.14159) * 0.0))
                            _ = v54
                        }
                    }
                    _ = fi49
                }
                switch (("" + ("foo bar" + "\n\t")) + "a") {
                case "foo bar":
                    v4 = 694
                case "hello":
                    for fi55 := 0; fi55 < 7; fi55++ {
                        fmt.Println(v4, fi55)
                        var v56 int = fi55
                        v57 := GenS4{F0: v56, F1: (v4 + 1)}
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                case "":
                    if len(v3) > 0 {
                        v3[0] = 725
                    }
                default:
                    var v58 string = "foo bar"
                    _ = v58
                }
            }
        }
    }
    switch "" {
    case "a":
        var v59 float64 = 0.0
        _ = v59
    case "":
        var v60 float64 = float64(0.0)
        _ = v60
    case "\n\t":
        var v61 int16 = 135
        _ = v61
    default:
        var v62 uint64 = 2
        _ = v62
    }
    var v63 string = ""
    _ = len(v2)
    v63 = (v63 + v63)
    fmt.Println(v63, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
}
