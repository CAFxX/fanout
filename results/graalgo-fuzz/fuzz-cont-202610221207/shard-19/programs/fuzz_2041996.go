package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 bool; F2 int16 }

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    v3 := make(map[string]uint32)
    var v4 string = ""
    v5 := []int{856, -784, -985, 4294967295}
    var v6 uint8 = 128
    switch 0 {
    case 42:
        fmt.Println(v1, v6, v6)
    }
    v7 := GenS1{F0: 2, F1: 1}
    v5 = append(v5, 916)
    fmt.Println(v2, v4, v1)
    if false {
        v8 := GenS1{F0: 346, F1: 687}
        if false {
            v9 := GenS1{F0: 255, F1: 65535}
            if (0 == (v1 ^ 42)) {
                v4 = "foo bar"
                v4 = ("hello" + ("" + (("foo bar" + "") + v4)))
            } else {
                fmt.Println(v2, v1)
                if false {
                    var v10 uint = 42
                    var v11 uint8 = 127
                    if true {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            v13 := GenS1{F0: 1, F1: 2}
                            switch (fi12 ^ fi12) {
                            case -2147483648:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 4294967295:
                                _ = len(v5)
                            case -9223372036854775808:
                                var v14 uint8 = v11
                                _ = v14
                            }
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v16 := GenS2{F0: 2147483648, F1: 555, F2: ((((((v4 + v4) + v4) + ("\n\t" + ("hello" + (v4 + (((v4 + ("hello" + v4)) + "a") + v4))))) + v4) + v4) + ("" + ""))}
                                v17 := GenS2{F0: 947, F1: 4294967295, F2: ((((("" + v4) + v4) + (v4 + (((v4 + "\n\t") + v4) + ""))) + v4) + "")}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            _ = fi12
                            _ = v13
                        }
                        var v18 uint = v10
                        _ = len(v5)
                        _ = v18
                    }
                    _ = v10
                    _ = v11
                } else {
                    v6 = (v6 | 0)
                    if true {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 uint8 = (^(v6))
                            _ = fi19
                            _ = v20
                        }
                        var v21 float64 = 0.0
                        if false {
                            v1 += (v1 ^ int16(v6))
                        } else {
                            v22 := GenS3{F0: v1, F1: 1}
                            v23 := GenS4{F0: -2147483648, F1: 42}
                            _ = v22
                            _ = v23
                        }
                        _ = v21
                    }
                    for fi24 := 0; fi24 < 9; fi24++ {
                        switch (v4 + "") {
                        case "foo bar":
                            for fi25 := 0; fi25 < 16; fi25++ {
                                var v26 uint32 = 463
                                if ((-2147483648 < 669) && true) {
                                    var v27 float64 = (float64(1e10) * (0.0 + (3.14159 + ((0.1 - 1e10) + 1e10))))
                                    var v28 int16 = -32768
                                    switch ((((v4 + v4) + v4) + v4) + v4) {
                                    case "foo bar":
                                        var v29 string = v4
                                        _ = v29
                                    default:
                                        if (((v2 <= 1) && false) || (("hello" != "hello") && (!true))) {
                                            fmt.Println(v2, v28)
                                            delete(v3, v4)
                                            v30 := GenS5{F0: 900, F1: v4, F2: 4294967295}
                                            _ = v30
                                        }
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                        v31 := GenS5{F0: 255, F1: v4, F2: uint64(v6)}
                        _ = fi24
                        _ = v31
                    }
                }
            }
            var v32 int = 2147483647
            _ = v32
            _ = v9
        } else {
            v6 += v6
            v33 := GenS1{F0: 490, F1: 65535}
            var v34 int8 = 0
            _ = v33
            _ = v34
        }
        _ = v8
    }
    if true {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int8 = -20
            _ = fi35
            _ = v36
        }
        switch ("\n\t" + (v4 + v4)) {
        case "":
            var v37 int32 = -1
            _ = v37
        case "\n\t":
            fmt.Println(v6, v6)
        case "hello":
            switch v4 {
            case "\n\t":
                var v38 bool = true
                _ = v38
            }
        default:
            v6 = v6
        }
    } else {
        if true {
            if (0.0 < float64(float32((0.0 - (1.5 + 1.5))))) {
                var v39 uint64 = 207
                var v40 uint8 = v6
                _ = v39
                _ = v40
            }
            fmt.Println(v6, v1)
        } else {
            v41 := v3[(v4 + v4)]
            v42 := v3[v4]
            _ = v41
            _ = v42
        }
    }
    if (!((^(v1)) < v2)) {
        v43 := GenS6{F0: -753, F1: 313}
        var v44 int32 = 42
        for fi45 := 0; fi45 < 8; fi45++ {
            if false {
                v46 := GenS4{F0: 0, F1: 387}
                v6 = v6
                _ = v46
            } else {
                var v47 int = (fi45 % ((fi45) | 1))
                switch v47 {
                case 632:
                    var v48 int8 = 42
                    _ = v48
                case 9223372036854775807:
                    if false {
                        if (true || (!(-727 == 0))) {
                            v1 -= v1
                            fmt.Println(v2, v2, v4)
                        } else {
                            for fi49 := 0; fi49 < 2; fi49++ {
                                var v50 float32 = 3.14159
                                switch ("\n\t" + v4) {
                                case "hello":
                                    if (true || true) {
                                        v1 += 1000
                                        var v51 int = 0
                                        _ = v51
                                    }
                                case "":
                                    if true {
                                        v1 -= (-(v2))
                                        if false {
                                            v1 += 42
                                        } else {
                                            fmt.Println(v1)
                                        }
                                    } else {
                                        v2 = (v1 * 1000)
                                        v2 = 1
                                    }
                                    v52 := GenS3{F0: (v1 % (((v2 - v1)) | 1)), F1: 1}
                                    switch fi45 {
                                    case -392:
                                        v6 += v6
                                    case 42:
                                        var v53 uint = 65535
                                        _ = v53
                                    case 1:
                                        v2 += (-(v1))
                                    default:
                                        v4 = v4
                                    }
                                    fmt.Println(v44, v4, v50)
                                    _ = v52
                                }
                                _ = fi49
                                _ = v50
                            }
                        }
                    } else {
                        fmt.Println(v1)
                        v6 = (^(v6))
                    }
                case 128:
                    v54 := GenS7{F0: 32767}
                    _ = v54
                }
                _ = v47
            }
            if false {
                v44 = v44
                var v55 bool = true
                if ((v55 && false) && (564 != 1)) {
                    var v56 int = 9223372036854775807
                    _ = v56
                } else {
                    switch (fi45 + -625) {
                    case 127:
                        var v57 int8 = 100
                        _ = v57
                    case 32767:
                        for fi58 := 0; fi58 < 13; fi58++ {
                            fmt.Println(v44)
                            v59 := GenS5{F0: 0, F1: v4, F2: 2}
                            var v60 int = fi45
                            var v61 uint = 65535
                            switch (fi45 + (-(fi58))) {
                            case -2147483648:
                                var v62 int64 = -242
                                _ = v62
                            case 745:
                                if (2 == 2147483647) {
                                    v63 := v3[""]
                                    var v64 bool = v55
                                    if v64 {
                                        fmt.Println(v64)
                                        fmt.Println(v44)
                                        v4 = ("\n\t" + v4)
                                    } else {
                                        fmt.Println(v2)
                                        fmt.Println(v44, v6, fi45)
                                        var v65 uint8 = v6
                                        switch fi45 {
                                        case 1:
                                            v66 := GenS8{F0: (("foo bar" + (v4 + v4)) + (("\n\t" + v4) + ("a" + v4)))}
                                            _ = v66
                                        case 127:
                                            var v67 int = (v60 + (fi58 >> (((^(v61))) & 63)))
                                            _ = v67
                                        default:
                                            switch fi45 {
                                            case -359:
                                                var v68 string = v4
                                                _ = v68
                                            case 548:
                                                var v69 uint = (v61 % ((v61) | 1))
                                                _ = v69
                                            case -102:
                                                var v70 int64 = 0
                                                _ = v70
                                            }
                                        }
                                        v71 := GenS4{F0: 363, F1: v61}
                                        v55 = v55
                                        v4 = v4
                                        var v72 uint8 = v6
                                        _ = v65
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v63
                                    _ = v64
                                }
                            case 128:
                                if (127 < 100) {
                                    var v73 int8 = 0
                                    v74 := GenS9{F0: 9223372036854775807, F1: true, F2: (v1 >> ((v61) & 63))}
                                    var v75 uint64 = 4294967295
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    var v76 int64 = -1
                                    delete(v3, v4)
                                    _ = v76
                                }
                            }
                            fmt.Println(v2, v4, v55)
                            var v77 bool = v55
                            _ = fi58
                            _ = v59
                            _ = v60
                            _ = v61
                            _ = v77
                        }
                    }
                    v55 = v55
                }
                _ = v55
            }
            _ = fi45
        }
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
