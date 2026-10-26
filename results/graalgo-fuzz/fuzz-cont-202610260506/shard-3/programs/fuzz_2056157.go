package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint16) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    switch "\n\t" {
    case "hello":
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi8, p0)
            var v9 uint8 = uint8(p0)
            _ = fi8
            _ = v9
        }
    case "\n\t":
        p0 += (p0 | (p0 - 1))
    case "a":
        for fi10 := 0; fi10 < 14; fi10++ {
            switch 32767 {
            case 127:
                v11 := []int{}
                _ = v11
            case 9223372036854775807:
                var v12 bool = false
                _ = v12
            }
            _ = fi10
        }
    }
    if (9223372036854775807 < 1) {
        if (-9223372036854775808 != 0) {
            for fi13 := 0; fi13 < 1; fi13++ {
                p0 += (^(p0))
                _ = fi13
            }
        }
        var v14 int = -533
        p0 = (p0 * (^(p0)))
        _ = v14
    } else {
        fmt.Println(p0)
    }
    switch "a" {
    case "foo bar":
        var v15 uint8 = 255
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 string = ""
        switch 0 {
        case 808:
            switch "a" {
            case "foo bar":
                v17 = v17
            case "":
                v18 := []int{4294967295, (fi16 + fi16)}
                _ = v18
            case "a":
                var v19 uint64 = 18446744073709551615
                _ = v19
            default:
                fmt.Println(v17, p0)
            }
        case 1:
            v20 := GenS1{F0: (-(fi16)), F1: (p0 + 255), F2: fi16}
            _ = v20
        case 821:
            v21 := []int{-1, fi16, fi16, 42}
            _ = v21
        }
        var v22 int32 = 339
        _ = fi16
        _ = v17
        _ = v22
    }
    _ = p0
    return 480
}
func genFn1(p0 int64, p1 int) bool {
    var v23 uint = 759
    if (!true) {
        v24 := GenS1{F0: 42, F1: 518, F2: 4294967295}
        _ = v24
    }
    p0 += (p0 * 1)
    _ = p0
    _ = p1
    _ = v23
    return true
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{32768, -118, 127}
    var v3 bool = false
    v4 := []uint32{457, 480, 2, 42}
    var v5 int64 = 0
    v6 := make(map[int]int)
    switch v1 {
    case "\n\t":
        switch 42 {
        case -1:
            switch 4294967295 {
            case 1:
                v25 := v6[0]
                _ = v25
            case 127:
                v26 := v6[1]
                _ = v26
            case -2147483648:
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int32 = 42
                    v29 := GenS1{F0: 4294967295, F1: 81, F2: 2147483647}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            default:
                switch "" {
                case "foo bar":
                    v3 = (!v3)
                case "a":
                    _ = len(v6)
                default:
                    v1 = "hello"
                }
            }
        case 1:
            switch 32768 {
            case -863:
                v6[9223372036854775807] = 2147483647
            case 1:
                v30 := GenS1{F0: -867, F1: 474, F2: 32767}
                _ = v30
            case 466:
                _ = len(v4)
            }
        case 4294967295:
            if true {
                var v31 string = "hello"
                _ = v31
            }
        }
    case "foo bar":
        fmt.Println(v5)
    case "a":
        if v3 {
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 string = ("hello" + v1)
                var v34 int8 = 40
                delete(v6, fi32)
                _ = fi32
                _ = v33
                _ = v34
            }
            v1 = (("\n\t" + "\n\t") + v1)
        } else {
            v35 := GenS1{F0: int(v5), F1: 0, F2: -2147483648}
            var v36 uint = 255
            _ = v35
            _ = v36
        }
    default:
        delete(v6, 2147483647)
    }
    v37 := GenS2{F0: 65535, F1: 353}
    v2 = append(v2, 781)
    v3 = (42 < 255)
    switch 1 {
    case 0:
        v5 = v5
    case -55:
        var v38 int = 32767
        _ = v38
    case -16:
        v39 := GenS2{F0: 65535, F1: 99}
        _ = v39
    default:
        if (false || true) {
            delete(v6, -1)
        }
    }
    v5 -= v5
    fmt.Println(v3, v1, v1)
    switch -484 {
    case -2147483648:
        switch 127 {
        case 9223372036854775807:
            var v40 int8 = -101
            _ = v40
        case 42:
            v41 := GenS3{F0: -1, F1: 1}
            _ = v41
        case 488:
            var v42 int16 = 0
            _ = v42
        }
    case 122:
        var v43 float64 = 1e10
        _ = v43
    }
    if true {
        if v3 {
            fmt.Println(v1, v3)
            if true {
                if true {
                    var v44 string = v1
                    _ = v44
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v1 = (v1 + ("hello" + v1))
                }
                v45 := GenS2{F0: 32767, F1: 18446744073709551615}
                _ = v45
            }
        } else {
            fmt.Println(v5, v3)
            for fi46 := 0; fi46 < 9; fi46++ {
                var v47 float32 = ((float32(float64((-2.25 + 1.5))) / 1.0) - 0.0)
                switch 2147483647 {
                case -2147483648:
                    switch ((v1 + "hello") + "hello") {
                    case "\n\t":
                        v1 = ((v1 + "foo bar") + (v1 + "a"))
                    case "foo bar":
                        v48 := GenS3{F0: v5, F1: 128}
                        _ = v48
                    case "":
                        for fi49 := 0; fi49 < 4; fi49++ {
                            fmt.Println(v3)
                            var v50 float64 = -2.25
                            _ = fi49
                            _ = v50
                        }
                    default:
                        for fi51 := 0; fi51 < 6; fi51++ {
                            if (!(386 >= 4294967295)) {
                                v52 := GenS1{F0: (fi46 % (((fi51 / ((fi46) | 1))) | 1)), F1: 42, F2: fi46}
                                _ = v52
                            }
                            _ = fi51
                        }
                    }
                case 2147483647:
                    if false {
                        v47 = (0.0 * v47)
                        if (!(!((v47 * (0.1 / ((v47 / float32(3.14159)) / -2.25))) == 0.1))) {
                            for fi53 := 0; fi53 < 3; fi53++ {
                                v3 = (!(2 <= 886))
                                if len(v4) > 0 {
                                    v4[0] = 552
                                }
                                var v54 uint8 = 203
                                _ = len(v6)
                                _ = fi53
                                _ = v54
                            }
                            var v55 float32 = (-2.25 + 1.5)
                            if false {
                                var v56 int64 = v5
                                _ = v56
                            }
                            _ = v55
                        }
                        v5 = 542
                    }
                case 32767:
                    switch "a" {
                    case "hello":
                        v6[(fi46 * 2147483647)] = fi46
                    default:
                        if v3 {
                            if (2147483647 < 2) {
                                for fi57 := 0; fi57 < 19; fi57++ {
                                    v5 -= (-(v5))
                                    fmt.Println(v1)
                                    delete(v6, (fi46 * -826))
                                    _ = fi57
                                }
                                if ((false && v3) || false) {
                                    if v3 {
                                        for fi58 := 0; fi58 < 14; fi58++ {
                                            switch (fi58 % ((int(v5)) | 1)) {
                                            case 4294967295:
                                                v3 = v3
                                            case 42:
                                                switch v1 {
                                                case "":
                                                    fmt.Println(v1)
                                                case "\n\t":
                                                    v4 = append(v4, 1)
                                                case "a":
                                                    v1 = (v1 + v1)
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    if v3 {
                                                        _ = len(v6)
                                                        fmt.Println(v5)
                                                        var v59 string = v1
                                                        _ = v59
                                                    } else {
                                                        var v60 float64 = 3.14159
                                                        _ = v60
                                                    }
                                                }
                                            }
                                            var v61 int8 = 42
                                            var v62 uint8 = 0
                                            _ = fi58
                                            _ = v61
                                            _ = v62
                                        }
                                        for fi63 := 0; fi63 < 5; fi63++ {
                                            _ = len(v2)
                                            _ = fi63
                                        }
                                    } else {
                                        for fi64 := 0; fi64 < 6; fi64++ {
                                            var v65 uint64 = 4294967295
                                            _ = fi64
                                            _ = v65
                                        }
                                        v2 = append(v2, (fi46 + int(v5)))
                                    }
                                } else {
                                    v5 -= v5
                                }
                            } else {
                                var v66 uint = 255
                                _ = v66
                            }
                            var v67 uint = 18446744073709551615
                            fmt.Println(v3)
                            _ = v67
                        } else {
                            var v68 int64 = (-(v5))
                            _ = v68
                        }
                    }
                }
                _ = fi46
                _ = v47
            }
        }
    } else {
        v1 = v1
    }
    var v69 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
