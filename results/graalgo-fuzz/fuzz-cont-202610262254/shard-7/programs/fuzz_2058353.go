package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint32; F1 string }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 bool = false
    var v3 uint16 = 0
    var v4 int16 = 991
    v5 := []uint32{2, 697, 42}
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: v4}
        _ = fi7
        _ = v8
    }
    if (!v2) {
        if true {
            fmt.Println(v3)
            v9 := GenS2{F0: 104, F1: v3, F2: 65535}
            fmt.Println(v3, v3)
            _ = v9
        } else {
            var v10 float32 = (1.5 + 0.0)
            v2 = v2
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 int16 = 597
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: 32767}
                _ = fi11
                _ = v14
            }
            _ = v10
        }
        v15 := GenS3{F0: 57, F1: v1, F2: 421}
        var v16 string = "foo bar"
        _ = v15
        _ = v16
    } else {
        v2 = (!v2)
    }
    switch 414 {
    case 1:
        for fi17 := 0; fi17 < 13; fi17++ {
            switch -1 {
            case 128:
                var v18 bool = v2
                _ = v18
            case -780:
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 uint32 = 2147483648
                    if len(v5) > 0 {
                        v5[0] = 274
                    }
                    if len(v5) > 0 {
                        v5[0] = v20
                    }
                    _ = fi19
                    _ = v20
                }
            case 4294967295:
                if true {
                    if (905 < -1) {
                        if false {
                            if len(v5) > 0 {
                                v5[0] = 4294967295
                            }
                            v4 = (v4 >> ((0) & 63))
                        } else {
                            for fi21 := 0; fi21 < 3; fi21++ {
                                var v22 int8 = 127
                                _ = fi21
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 5; fi23++ {
                                fmt.Println(v3, fi17, v3)
                                _ = fi23
                            }
                            fmt.Println(v2, v3)
                        }
                        v1 = ((((v1 + v1) + ((v1 + v1) + v1)) + ((("\n\t" + "\n\t") + v1) + (("a" + v1) + v1))) + "a")
                    }
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 string = v1
                        var v26 int16 = (v4 + 1)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    if true {
                        fmt.Println(fi17, fi17)
                    } else {
                        if ((1 <= 0) || v2) {
                            if (4294967295 == 1) {
                                v2 = false
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    v28 := GenS1{F0: v4}
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                switch (v1 + (v1 + (v1 + v1))) {
                                case "":
                                    for fi29 := 0; fi29 < 10; fi29++ {
                                        v30 := GenS4{F0: -1}
                                        fmt.Println(fi29, fi17, v3)
                                        var v31 float32 = ((3.14159 / -2.25) * ((0.0 * float32(0.1)) * 0.0))
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                default:
                                    switch ((v1 + (v1 + v1)) + (v1 + v1)) {
                                    case "\n\t":
                                        if v2 {
                                            fmt.Println(v1, v4, v1)
                                            fmt.Println(v2)
                                        } else {
                                            switch "" {
                                            case "hello":
                                                v1 = v1
                                            case "a":
                                                v6[0] = v1
                                            default:
                                                switch v1 {
                                                case "a":
                                                    for fi32 := 0; fi32 < 6; fi32++ {
                                                        var v33 int = 42
                                                        var v34 string = v1
                                                        _ = fi32
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                case "foo bar":
                                                    var v35 uint16 = (v3 % ((v3) | 1))
                                                    _ = v35
                                                default:
                                                    if v2 {
                                                        v36 := GenS3{F0: -128, F1: "foo bar", F2: 65535}
                                                        fmt.Println(v4, fi17, v2)
                                                        v5 = append(v5, 2)
                                                        _ = v36
                                                    } else {
                                                        switch v1 {
                                                        case "foo bar":
                                                            if v2 {
                                                                var v37 int16 = -816
                                                                var v38 string = ("a" + "foo bar")
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if len(v5) > 0 {
                                                            v5[0] = 2
                                                        }
                                                    }
                                                }
                                            }
                                            v39 := v6[2147483647]
                                            if ((^(v4)) >= v4) {
                                                var v40 int16 = int16(fi17)
                                                var v41 uint32 = 2
                                                _ = v40
                                                _ = v41
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                delete(v6, 2147483648)
                                            }
                                            _ = v39
                                        }
                                    case "foo bar":
                                        v3 -= (^(v3))
                                    }
                                }
                                for fi42 := 0; fi42 < 20; fi42++ {
                                    var v43 bool = (!v2)
                                    fmt.Println(v2, fi17, fi42)
                                    switch -2147483648 {
                                    case 1:
                                        if (-2147483648 <= -397) {
                                            v44 := GenS4{F0: 42}
                                            switch ((v1 + "") + v1) {
                                            case "\n\t":
                                                switch fi17 {
                                                case 641:
                                                    v6[920] = (v1 + v1)
                                                default:
                                                    v45 := GenS4{F0: 41}
                                                    _ = v45
                                                }
                                            case "hello":
                                                switch "hello" {
                                                case "\n\t":
                                                    if v2 {
                                                        var v46 int8 = -1
                                                        v47 := GenS5{F0: v46, F1: 4294967295, F2: v4}
                                                        v43 = (v1 != v1)
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        fmt.Println(v43)
                                                        var v48 int32 = 864
                                                        delete(v6, 1)
                                                        _ = v48
                                                    }
                                                }
                                            default:
                                                var v49 uint8 = 128
                                                _ = v49
                                            }
                                            v1 = v1
                                            _ = v44
                                        } else {
                                            v50 := GenS6{F0: (^(v3)), F1: 99, F2: 1}
                                            v51 := GenS7{F0: 2147483648, F1: v1}
                                            _ = v50
                                            _ = v51
                                        }
                                    case 975:
                                        var v52 int8 = 0
                                        _ = v52
                                    case 128:
                                        v53 := GenS6{F0: 26, F1: 0, F2: 42}
                                        _ = v53
                                    default:
                                        var v54 int16 = (v4 % ((v4) | 1))
                                        _ = v54
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                            }
                            _ = len(v6)
                        }
                        v4 = v4
                        _ = len(v5)
                    }
                }
            default:
                if v2 {
                    var v55 int8 = 42
                    _ = v55
                } else {
                    var v56 uint = 446
                    var v57 uint8 = 2
                    v5 = append(v5, 2)
                    _ = v56
                    _ = v57
                }
            }
            _ = fi17
        }
    case 0:
        var v58 uint8 = 127
        _ = v58
    case 9223372036854775807:
        var v59 int = 41
        _ = v59
    }
    if (127 < 1) {
        var v60 float32 = 3.14159
        var v61 uint8 = 255
        _ = v60
        _ = v61
    } else {
        var v62 float64 = float64((0.0 * 0.0))
        switch v1 {
        case "hello":
            if ((!false) && v2) {
                v4 = v4
            } else {
                v63 := GenS3{F0: 28, F1: v1, F2: 255}
                v3 += (v3 - (v3 ^ v3))
                _ = v63
            }
        default:
            fmt.Println(v1)
        }
        v1 = v1
        _ = v62
    }
    v4 -= v4
    if v2 {
        fmt.Println(v4)
    } else {
        if (396 >= 510) {
            fmt.Println(v1, v1, v3)
            if (true && v2) {
                delete(v6, 65535)
            }
        }
    }
    for fi64 := 0; fi64 < 2; fi64++ {
        if len(v5) > 0 {
            v5[0] = 212
        }
        _ = fi64
    }
    var v65 uint8 = 113
    v66 := GenS3{F0: -86, F1: v1, F2: 65535}
    if (592 < 255) {
        v67 := v6[4294967295]
        for fi68 := 0; fi68 < 15; fi68++ {
            v69 := GenS4{F0: 42}
            _ = fi68
            _ = v69
        }
        fmt.Println(v1, v2)
        _ = v67
    }
    for fi70 := 0; fi70 < 1; fi70++ {
        var v71 uint16 = 32767
        _ = fi70
        _ = v71
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    var v72 float64 = 0.1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v73 float32 = 0.0
    var v74 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v72
    _ = v73
    _ = v74
}
