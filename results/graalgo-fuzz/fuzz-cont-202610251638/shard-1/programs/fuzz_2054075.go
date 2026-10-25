package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 int = 963
    var v2 string = "foo bar"
    v3 := make(map[string]int)
    var v4 int32 = 0
    var v5 bool = false
    fmt.Println(v4, v1, v4)
    if true {
        switch v2 {
        case "hello":
            v6 := GenS1{F0: 2, F1: 2}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := []int{(fi7 * (v1 >> ((811) & 63))), 127}
                _ = fi7
                _ = v8
            }
        case "foo bar":
            var v9 float64 = 0.1
            _ = v9
        }
        v4 -= -517
        switch 42 {
        case -177:
            if (v4 != (v4 & (v4 + (v4 | (-(v4)))))) {
                v4 += 0
                v4 += (v4 * v4)
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v5 = (!(42 >= 112))
                    v11 := GenS2{F0: 42, F1: 255}
                    var v12 int64 = int64(v4)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        default:
            for fi13 := 0; fi13 < 13; fi13++ {
                switch v2 {
                case "foo bar":
                    var v14 string = ((v2 + "\n\t") + ("\n\t" + "foo bar"))
                    _ = v14
                case "hello":
                    v5 = ((((((float64(-2.25) * (float64(0.0) * float64(float32(3.14159)))) * float64(float32((3.14159 / -2.25)))) - float64(3.14159)) - -2.25) * 1e10) < 1.5)
                default:
                    v15 := []int{}
                    _ = v15
                }
                switch v2 {
                case "foo bar":
                    fmt.Println(v4, v1)
                case "hello":
                    v16 := []int{fi13, v1}
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v5 = v5
                        switch v2 {
                        case "foo bar":
                            var v18 uint32 = 607
                            _ = v18
                        case "\n\t":
                            if v5 {
                                v19 := GenS1{F0: 0, F1: 206}
                                if v5 {
                                    v2 = (v2 + v2)
                                    v4 -= v4
                                    fmt.Println(v5, fi13, v4)
                                    v20 := GenS2{F0: -79, F1: 2}
                                    _ = v20
                                }
                                fmt.Println(v1, v4)
                                _ = v19
                            }
                            v21 := []int{(v1 / (((fi13 | 791)) | 1)), fi17}
                            switch 1 {
                            case 4294967295:
                                fmt.Println(v5, fi17, v1)
                            }
                            _ = v21
                        case "a":
                            delete(v3, ((((v2 + ("a" + "\n\t")) + "foo bar") + (v2 + ((v2 + ("" + v2)) + v2))) + ((v2 + v2) + "")))
                        }
                        var v22 int8 = 6
                        _ = fi17
                        _ = v22
                    }
                }
                _ = fi13
            }
        }
    }
    switch v1 {
    case -9223372036854775808:
        switch (v1 >> ((223) & 63)) {
        case 4294967295:
            v4 = v4
        case 32768:
            if (0 == 323) {
                v5 = (v5 && (107 < 100))
                fmt.Println(v4, v4, v1)
                if (!v5) {
                    v23 := []int{v1, v1}
                    _ = v23
                }
            } else {
                v1 = (v1 << ((18446744073709551615) & 63))
                if v5 {
                    fmt.Println(v2)
                    delete(v3, "\n\t")
                    var v24 int8 = 1
                    _ = v24
                }
            }
        case 1:
            switch -9223372036854775808 {
            case 42:
                var v25 uint64 = 4294967295
                _ = v25
            case 32767:
                var v26 uint32 = 42
                _ = v26
            case 2147483647:
                v1 -= v1
            default:
                if ((!(-60 < 0)) && (2 > 754)) {
                    switch 535 {
                    case 4294967295:
                        v2 = ("a" + "hello")
                    default:
                        var v27 float64 = 0.1
                        _ = v27
                    }
                } else {
                    if v5 {
                        v2 = "\n\t"
                        var v28 bool = (v5 || v5)
                        for fi29 := 0; fi29 < 12; fi29++ {
                            var v30 int = 0
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                    v2 = v2
                    v4 = v4
                }
            }
        default:
            if false {
                v31 := v3[("" + "\n\t")]
                fmt.Println(v31, v4, v4)
                _ = v31
            }
        }
    case 4294967295:
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 uint8 = 1
            for fi34 := 0; fi34 < 3; fi34++ {
                delete(v3, v2)
                _ = fi34
            }
            v35 := []int{int(v33), -281, v1}
            _ = fi32
            _ = v33
            _ = v35
        }
    default:
        var v36 float64 = 1e10
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        for fi38 := 0; fi38 < 5; fi38++ {
            switch (-(v1)) {
            case 4294967295:
                fmt.Println(fi38, v2, v1)
            case 128:
                var v39 float32 = (0.1 / 3.14159)
                _ = v39
            case -2147483648:
                v40 := GenS3{F0: 2, F1: 819, F2: (^(fi38))}
                _ = v40
            }
            var v41 string = (("a" + v2) + v2)
            v42 := GenS4{F0: (-(v4)), F1: -32768}
            v43 := []int{(-(fi37)), fi38, (-(v1)), v1}
            _ = fi38
            _ = v41
            _ = v42
            _ = v43
        }
        v1 += int(v4)
        v44 := GenS5{F0: 0}
        _ = fi37
        _ = v44
    }
    v45 := []int{(v1 >> ((793) & 63)), -9223372036854775808, (v1 - 4294967295)}
    v1 += -9223372036854775808
    for fi46 := 0; fi46 < 10; fi46++ {
        if (3.14159 >= 0.1) {
            v5 = (!v5)
            if false {
                var v47 uint32 = 4294967295
                v48 := GenS2{F0: 100, F1: 2}
                v2 = ("a" + v2)
                _ = v47
                _ = v48
            }
            for fi49 := 0; fi49 < 19; fi49++ {
                var v50 uint8 = 120
                _ = fi49
                _ = v50
            }
        } else {
            switch ("a" + v2) {
            case "hello":
                fmt.Println(v5)
            case "foo bar":
                v51 := GenS5{F0: 42}
                _ = v51
            }
            for fi52 := 0; fi52 < 2; fi52++ {
                v53 := GenS6{F0: (v4 + (v4 & (v4 + (-(v4)))))}
                if (548 >= 2) {
                    var v54 string = v2
                    _ = v54
                }
                _ = fi52
                _ = v53
            }
        }
        _ = fi46
    }
    v55 := v3[v2]
    v56 := GenS5{F0: 11}
    if true {
        var v57 float64 = 1.5
        _ = v57
    } else {
        switch "\n\t" {
        case "a":
            for fi58 := 0; fi58 < 11; fi58++ {
                if false {
                    for fi59 := 0; fi59 < 9; fi59++ {
                        fmt.Println(v55, v5)
                        _ = fi59
                    }
                    var v60 bool = true
                    v2 = v2
                    _ = v60
                } else {
                    if v5 {
                        v61 := GenS7{F0: 2147483648, F1: 65535}
                        fmt.Println(v55, fi58, v55)
                        v62 := GenS8{F0: 2147483648, F1: 9223372036854775807}
                        _ = v61
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 7; fi63++ {
                            switch (^(v55)) {
                            case -800:
                                if (517 >= 255) {
                                    fmt.Println(v2, v55)
                                    v45 = append(v45, 32767)
                                    fmt.Println(fi63)
                                } else {
                                    fmt.Println(fi63, v55)
                                    var v64 uint8 = 2
                                    v4 -= v4
                                    _ = v64
                                }
                            case 2147483647:
                                var v65 uint64 = 18446744073709551615
                                _ = v65
                            case 32768:
                                v66 := v3["foo bar"]
                                _ = v66
                            default:
                                v67 := GenS9{F0: v2}
                                _ = v67
                            }
                            fmt.Println(v1)
                            delete(v3, "a")
                            delete(v3, "foo bar")
                            _ = fi63
                        }
                    }
                }
                _ = fi58
            }
        case "hello":
            var v68 uint32 = 0
            _ = v68
        default:
            v69 := GenS6{F0: 2147483647}
            _ = v69
        }
        var v70 int16 = -32768
        _ = len(v45)
        _ = v70
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        for fi72 := 0; fi72 < 8; fi72++ {
            v4 = (v4 >> ((18446744073709551615) & 63))
            if (42 == -285) {
                v2 = (v2 + v2)
                if true {
                    var v73 float32 = 0.0
                    _ = v73
                } else {
                    v74 := GenS10{F0: 366, F1: 86, F2: (v4 * (v4 ^ (v4 | v4)))}
                    v2 = v2
                    switch v2 {
                    case "":
                        v75 := GenS5{F0: 1}
                        _ = v75
                    case "\n\t":
                        switch (-(v55)) {
                        case 127:
                            v55 = 128
                        case -2147483648:
                            v76 := GenS11{F0: 42}
                            _ = v76
                        case 2147483647:
                            var v77 float64 = 1.5
                            _ = v77
                        }
                    case "a":
                        fmt.Println(fi72)
                    default:
                        v78 := GenS1{F0: 0, F1: 2}
                        _ = v78
                    }
                    fmt.Println(v1, v2)
                    _ = v74
                }
                v45 = append(v45, fi72)
            }
            _ = fi72
        }
        var v79 int = v55
        _ = fi71
        _ = v79
    }
    if v5 {
        switch "hello" {
        case "a":
            var v80 float64 = -2.25
            _ = v80
        case "foo bar":
            if (10 == 2) {
                if len(v45) > 0 {
                    v45[0] = v55
                }
                v1 = v55
            } else {
                var v81 int64 = 95
                v4 += v4
                _ = v81
            }
        case "\n\t":
            v82 := GenS10{F0: 545, F1: 2, F2: (v4 | v4)}
            _ = v82
        default:
            switch -9223372036854775808 {
            case 32768:
                v83 := GenS2{F0: -128, F1: 136}
                _ = v83
            default:
                v3[v2] = (^(v1))
            }
        }
    } else {
        if (0 == 42) {
            v55 = 1
            switch v2 {
            case "":
                fmt.Println(v1, v4)
            case "foo bar":
                switch (v1 | v55) {
                case 32767:
                    fmt.Println(v2, v5, v2)
                }
            }
        }
        v2 = "foo bar"
    }
    _ = len(v3)
    fmt.Println(v55)
    v5 = (1 != 74)
    var v84 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v55
    _ = v56
    _ = v84
}
