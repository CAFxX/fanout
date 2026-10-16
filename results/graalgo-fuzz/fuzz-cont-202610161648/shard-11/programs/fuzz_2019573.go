package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 uint; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t"}
    var v3 float32 = 3.14159
    v4 := []string{"", "", "foo bar", "a"}
    var v5 string = ((("hello" + ("foo bar" + ("hello" + "hello"))) + "") + ("" + ""))
    v6 := make(map[int]string)
    for fi7 := 0; fi7 < 5; fi7++ {
        v5 = ("a" + "foo bar")
        var v8 float32 = v3
        _ = fi7
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 13; fi9++ {
            v4 = append(v4, v5)
            v10 := GenS1{F0: 42}
            if true {
                for fi11 := 0; fi11 < 15; fi11++ {
                    v3 = float32((1e10 + 1.5))
                    _ = fi11
                }
            }
            _ = fi9
            _ = v10
        }
        var v12 uint16 = 2
        var v13 int8 = -1
        _ = v12
        _ = v13
    } else {
        var v14 uint8 = 128
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch -747 {
        case 2147483647:
            if (!false) {
                var v16 string = ""
                v2 = append(v2, v5)
                _ = v16
            } else {
                for fi17 := 0; fi17 < 5; fi17++ {
                    v3 = 3.14159
                    if (0 == -9223372036854775808) {
                        if (-1 >= -997) {
                            v18 := GenS2{F0: (fi15 ^ (fi15 / ((fi17) | 1))), F1: 42}
                            _ = v18
                        }
                        v3 += -2.25
                        v3 -= (v3 * (((v3 + v3) + 3.14159) / ((float32(3.14159) - v3) * v3)))
                    }
                    _ = fi17
                }
                v19 := v1[v5]
                for fi20 := 0; fi20 < 6; fi20++ {
                    if true {
                        v21 := GenS3{F0: -45}
                        if (false || (42 == (fi20 / (((-(fi20))) | 1)))) {
                            if true {
                                v22 := GenS1{F0: 128}
                                _ = v22
                            } else {
                                if (!true) {
                                    var v23 uint16 = 0
                                    _ = v23
                                }
                            }
                            var v24 int32 = 1
                            var v25 uint32 = 4294967295
                            _ = v24
                            _ = v25
                        }
                        switch "" {
                        case "\n\t":
                            if true {
                                fmt.Println(fi20, fi15)
                            }
                        }
                        fmt.Println(v3, v3, v3)
                        var v26 int64 = 9223372036854775807
                        _ = len(v1)
                        v3 += ((v3 * 1e10) + 1e10)
                        v27 := GenS4{F0: 206, F1: -9223372036854775808, F2: 511}
                        v28 := GenS3{F0: -128}
                        var v29 string = v5
                        _ = v21
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS3{F0: 1}
                        v19 = (v19 - v19)
                        v31 := GenS5{F0: 42}
                        _ = v30
                        _ = v31
                    }
                    v19 -= (v19 + (^(v19)))
                    _ = fi20
                }
                _ = v19
            }
        case 4294967295:
            if true {
                var v32 float32 = v3
                v3 -= v32
                _ = v32
            }
        default:
            fmt.Println(v5, v5, v5)
        }
        _ = fi15
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 string = (("hello" + ((("foo bar" + v5) + v5) + (v5 + v5))) + ((v5 + v5) + v5))
        if (!(128 == 65535)) {
            var v35 int16 = 32767
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    v36 := GenS1{F0: 110}
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v37 float32 = (0.0 + (v3 / (1e10 - (v3 / v3))))
    v38 := GenS3{F0: -1}
    var v39 int = 2147483647
    switch "foo bar" {
    case "":
        fmt.Println(v5)
    case "a":
        delete(v6, (v39 | v39))
    }
    v39 = v39
    switch v5 {
    case "foo bar":
        for fi40 := 0; fi40 < 1; fi40++ {
            v5 = v5
            _ = fi40
        }
    default:
        var v41 int = v39
        _ = v41
    }
    var v42 int8 = 42
    if true {
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 uint8 = 42
            if ((v5 + v5) > v5) {
                var v45 int16 = -583
                if false {
                    fmt.Println(v39, v42)
                    switch "" {
                    case "\n\t":
                        if true {
                            var v46 float64 = 1e10
                            _ = v46
                        } else {
                            fmt.Println(v44, v44, v45)
                            var v47 int64 = -954
                            _ = v47
                        }
                    default:
                        switch v5 {
                        case "\n\t":
                            v45 = 1000
                        default:
                            v48 := GenS6{F0: 661, F1: v42, F2: false}
                            _ = v48
                        }
                    }
                }
                v37 = -2.25
                _ = v45
            } else {
                for fi49 := 0; fi49 < 12; fi49++ {
                    v44 = v44
                    v42 = v42
                    _ = fi49
                }
                if (((1 == 2) || (false || (false || true))) && (!false)) {
                    var v50 float32 = 0.1
                    fmt.Println(v42)
                    _ = v50
                }
                var v51 int8 = v42
                _ = v51
            }
            _ = fi43
            _ = v44
        }
        v37 -= float32(float64(0.1))
        v1[(("" + ("foo bar" + (v5 + ((("foo bar" + "hello") + v5) + v5)))) + "\n\t")] = 42
    }
    if (1 > 112) {
        if (true || true) {
            var v52 uint32 = 2
            if false {
                v52 += 2147483647
            } else {
                fmt.Println(v42, v42)
                switch v39 {
                case 32767:
                    if ((!true) && (v39 != v39)) {
                        switch v39 {
                        case -69:
                            for fi53 := 0; fi53 < 19; fi53++ {
                                fmt.Println(v5, v52)
                                _ = fi53
                            }
                        case 4294967295:
                            v54 := GenS3{F0: -1}
                            _ = v54
                        default:
                            switch (("a" + (v5 + ("a" + v5))) + v5) {
                            case "hello":
                                v55 := GenS1{F0: 255}
                                _ = v55
                            case "":
                                v56 := GenS7{F0: 42}
                                _ = v56
                            default:
                                v57 := GenS7{F0: int32(v52)}
                                _ = v57
                            }
                        }
                    }
                case -740:
                    var v58 string = ((v5 + ("a" + v5)) + "a")
                    _ = v58
                case 42:
                    for fi59 := 0; fi59 < 9; fi59++ {
                        switch (v39 % (((fi59 - fi59)) | 1)) {
                        case 0:
                            var v60 uint16 = 370
                            _ = v60
                        default:
                            switch fi59 {
                            case -360:
                                v61 := GenS6{F0: 1000, F1: (v42 & (-(v42))), F2: (18446744073709551615 > 42)}
                                _ = v61
                            case -2147483648:
                                v39 -= (v39 * (-(v39)))
                            default:
                                v62 := GenS8{F0: v39, F1: 2}
                                _ = v62
                            }
                        }
                        var v63 uint16 = 0
                        _ = fi59
                        _ = v63
                    }
                }
                if ((120 <= 42) || false) {
                    var v64 float64 = 1.5
                    _ = v64
                } else {
                    if ((v3 + float32(-2.25)) < (float32(3.14159) / float32(1e10))) {
                        var v65 float64 = 0.1
                        fmt.Println(v52, v39)
                        v66 := GenS6{F0: 4294967295, F1: (v42 + v42), F2: false}
                        _ = v65
                        _ = v66
                    } else {
                        if (v52 < (^(v52))) {
                            var v67 int64 = 0
                            _ = v67
                        } else {
                            _ = len(v4)
                        }
                        v39 = (v39 ^ 128)
                    }
                    v68 := GenS5{F0: v42}
                    v69 := v6[v39]
                    _ = v68
                    _ = v69
                }
            }
            var v70 uint = 255
            _ = v52
            _ = v70
        }
        switch v39 {
        case 4294967295:
            var v71 int8 = -1
            _ = v71
        default:
            fmt.Println(v5)
        }
    } else {
        var v72 float32 = v37
        v73 := GenS9{F0: 18446744073709551615, F1: -1, F2: 1}
        var v74 uint64 = 0
        _ = v72
        _ = v73
        _ = v74
    }
    v75 := GenS6{F0: 825, F1: (^(v42)), F2: (!true)}
    if false {
        v4 = append(v4, (("" + v5) + v5))
    } else {
        var v76 int32 = -1
        var v77 int32 = (v76 - v76)
        _ = v76
        _ = v77
    }
    v78 := GenS4{F0: -1, F1: v39, F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v75
    _ = v78
}
