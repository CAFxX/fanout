package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 uint8; F2 int32 }
func genFn0() int32 {
    var v4 int64 = 42
    switch "" {
    case "foo bar":
        for fi5 := 0; fi5 < 6; fi5++ {
            fmt.Println(v4, v4)
            _ = fi5
        }
    default:
        fmt.Println(v4)
    }
    _ = v4
    return 42
}

func main() {
    v1 := []uint32{667, 2147483647}
    var v2 string = (("\n\t" + ("a" + ("foo bar" + "a"))) + "hello")
    var v3 uint8 = 127
    v1 = append(v1, 1)
    v6 := GenS1{F0: 65535}
    v7 := GenS1{F0: 873}
    fmt.Println(v2, v3)
    v3 -= 42
    fmt.Println(v3, v3)
    for fi8 := 0; fi8 < 6; fi8++ {
        v2 = v2
        v3 = v3
        v9 := GenS2{F0: 65535, F1: (v3 + v3)}
        _ = fi8
        _ = v9
    }
    _ = len(v1)
    v3 -= 89
    if false {
        v10 := GenS2{F0: 0, F1: (v3 | v3)}
        _ = v10
    } else {
        fmt.Println(v2, v3)
    }
    switch "" {
    case "":
        v1 = append(v1, 2)
    case "\n\t":
        v11 := GenS2{F0: uint16(v3), F1: v3}
        _ = v11
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v12 int8 = 100
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int8 = 42
        if (!true) {
            v2 = v2
        } else {
            _ = len(v1)
            if true {
                switch (fi13 / (((fi13 ^ fi13)) | 1)) {
                case 32768:
                    if true {
                        v2 = "\n\t"
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v12 += v14
                            _ = fi15
                        }
                        var v16 string = v2
                        _ = v16
                    }
                case 0:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if false {
                            var v18 uint64 = 42
                            if (v14 <= (v12 - v12)) {
                                if false {
                                    var v19 int16 = 1000
                                    v1 = append(v1, 4294967295)
                                    _ = v19
                                } else {
                                    v20 := GenS2{F0: 494, F1: (^(v3))}
                                    v2 = v2
                                    _ = v20
                                }
                                v21 := GenS3{F0: (fi17 & 42), F1: (^(v14))}
                                _ = len(v1)
                                _ = v21
                            }
                            var v22 float64 = 0.0
                            if len(v1) > 0 {
                                v1[0] = 941
                            }
                            _ = v18
                            _ = v22
                        }
                        _ = fi17
                    }
                case 4294967295:
                    if (42 == 2) {
                        v23 := GenS4{F0: v3, F1: 2}
                        _ = v23
                    }
                default:
                    if true {
                        if ((0.1 / -2.25) >= 3.14159) {
                            for fi24 := 0; fi24 < 12; fi24++ {
                                if false {
                                    var v25 uint32 = 283
                                    var v26 int32 = -2147483648
                                    var v27 int16 = 815
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi24
                            }
                            switch v2 {
                            case "a":
                                v28 := GenS5{F0: 0, F1: 127}
                                _ = v28
                            }
                            v29 := GenS5{F0: (v3 + 42), F1: int(v14)}
                            _ = v29
                        }
                        var v30 int = -20
                        _ = v30
                    }
                }
                v1 = append(v1, 0)
                var v31 int32 = -394
                _ = v31
            } else {
                var v32 uint8 = v3
                v32 = 16
                for fi33 := 0; fi33 < 20; fi33++ {
                    v34 := GenS6{F0: 32767}
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
            v35 := GenS1{F0: 2}
            _ = v35
        }
        if len(v1) > 0 {
            v1[0] = uint32(v3)
        }
        _ = fi13
        _ = v14
    }
    if (0 != v3) {
        if true {
            var v36 int32 = 879
            _ = v36
        }
        v37 := GenS6{F0: 32767}
        var v38 float32 = 0.0
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 int32 = 2147483647
            _ = fi39
            _ = v40
        }
        if true {
            if true {
                var v41 int16 = -32768
                _ = v41
            } else {
                if true {
                    v1 = append(v1, 4294967295)
                    switch v2 {
                    case "foo bar":
                        for fi42 := 0; fi42 < 10; fi42++ {
                            var v43 uint32 = 55
                            _ = fi42
                            _ = v43
                        }
                    default:
                        for fi44 := 0; fi44 < 13; fi44++ {
                            v12 = v12
                            for fi45 := 0; fi45 < 16; fi45++ {
                                var v46 int8 = (v12 & (v12 | v12))
                                if (-1 != genFn0()) {
                                    v47 := GenS7{F0: 4294967295, F1: 65535}
                                    v48 := GenS8{F0: 42}
                                    _ = v47
                                    _ = v48
                                }
                                var v49 int = (^(fi45))
                                _ = fi45
                                _ = v46
                                _ = v49
                            }
                            _ = fi44
                        }
                    }
                    switch 753 {
                    case 9223372036854775807:
                        if (961 == 32767) {
                            fmt.Println(v2, v3)
                            v2 = (v2 + "foo bar")
                            var v50 uint32 = 502
                            _ = v50
                        } else {
                            fmt.Println(v2, v2, v12)
                            v51 := GenS5{F0: v3, F1: 0}
                            _ = len(v1)
                            _ = v51
                        }
                    case -9223372036854775808:
                        v12 -= (v12 % ((-70) | 1))
                    case -504:
                        var v52 bool = (!((!true) || true))
                        _ = v52
                    }
                } else {
                    v53 := GenS7{F0: 179, F1: 4294967295}
                    for fi54 := 0; fi54 < 15; fi54++ {
                        if (!true) {
                            var v55 float32 = 1.5
                            if false {
                                var v56 float32 = v55
                                _ = len(v1)
                                if false {
                                    var v57 int64 = -1
                                    _ = v57
                                } else {
                                    v55 += v56
                                    var v58 float32 = v55
                                    v3 -= v3
                                    _ = v58
                                }
                                _ = v56
                            } else {
                                switch (fi54 & 0) {
                                case 9223372036854775807:
                                    var v59 uint32 = 2
                                    _ = v59
                                case -669:
                                    v1 = append(v1, uint32(v3))
                                }
                            }
                            var v60 uint8 = v3
                            _ = v55
                            _ = v60
                        } else {
                            for fi61 := 0; fi61 < 11; fi61++ {
                                v62 := GenS9{F0: ("a" + ("a" + (v2 + (v2 + v2)))), F1: v3, F2: -771}
                                var v63 uint8 = (v3 + 238)
                                var v64 uint64 = 0
                                _ = fi61
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                        }
                        _ = fi54
                    }
                    _ = v53
                }
                var v65 int16 = -394
                v2 = "a"
                _ = v65
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v1 = append(v1, 425)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
