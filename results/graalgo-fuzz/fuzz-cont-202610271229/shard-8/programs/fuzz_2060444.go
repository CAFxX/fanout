package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint16; F2 int }
type GenS6 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    v3 := []string{"a", "hello", ((v2 + "") + v2), (((v2 + (("a" + ("hello" + "hello")) + v2)) + v2) + v2)}
    var v4 int = 32767
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 2, F1: 127}
    v7 := GenS2{F0: false, F1: 42}
    if (604 > 102) {
        v2 = ((v2 + v2) + "hello")
        var v8 float64 = -2.25
        _ = v8
    } else {
        var v9 uint8 = 169
        _ = v9
    }
    fmt.Println(v2, v2)
    fmt.Println(v4)
    v10 := GenS1{F0: 0, F1: -34}
    v11 := v5[500]
    v4 -= v4
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 4294967295, F1: -37}
        var v14 uint = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int16 = 0
    var v16 bool = ((^(v4)) > v4)
    var v17 bool = (255 < 10)
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(fi18)
        var v19 int64 = -184
        _ = fi18
        _ = v19
    }
    v2 = "\n\t"
    var v20 float64 = float64(1e10)
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            if v17 {
                var v23 string = (v2 + ((v2 + v2) + ""))
                if (310 == 42) {
                    if v16 {
                        v23 = ("hello" + ("" + v23))
                        v15 += (^(v15))
                        if (977 == 0) {
                            v11 += v11
                            var v24 string = (v23 + (("" + v2) + v23))
                            _ = v24
                        } else {
                            switch fi22 {
                            case 32768:
                                fmt.Println(v11, v17, v23)
                            default:
                                v15 -= v15
                            }
                            v16 = false
                            v23 = (v23 + ("" + ((v23 + "foo bar") + "a")))
                        }
                        v4 -= (-(v4))
                    }
                    if ((v11 >= 2147483647) || v17) {
                        if len(v3) > 0 {
                            v3[0] = v23
                        }
                        v2 = v2
                        var v25 uint = 2
                        if (18446744073709551615 < 4294967295) {
                            if (110 == 43) {
                                v4 = (v4 ^ fi21)
                                fmt.Println(fi22)
                                if v16 {
                                    var v26 uint32 = v11
                                    _ = v26
                                }
                            }
                        } else {
                            v16 = (true || (uint16(v11) != uint16(v25)))
                            if (!v17) {
                                fmt.Println(v20, fi21, fi21)
                                if true {
                                    v16 = (!(!(v17 && (18446744073709551615 > v25))))
                                    v15 -= (v15 / ((v15) | 1))
                                    if v17 {
                                        if (v17 || false) {
                                            v1[v2] = ("a" + "\n\t")
                                            if v16 {
                                                if v16 {
                                                    var v27 string = ("\n\t" + "")
                                                    _ = v27
                                                }
                                            }
                                        }
                                        v17 = v17
                                    }
                                } else {
                                    var v28 bool = v17
                                    v16 = true
                                    _ = v28
                                }
                                _ = len(v1)
                            } else {
                                var v29 uint = 18446744073709551615
                                v11 = v11
                                var v30 float64 = ((v20 + 0.0) + v20)
                                fmt.Println(fi21)
                                var v31 bool = v17
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        delete(v1, "\n\t")
                        _ = v25
                    } else {
                        var v32 int64 = -379
                        v11 += (v11 | 2147483648)
                        v23 = v2
                        _ = v32
                    }
                    if true {
                        v17 = ((4294967295 == 42) && v16)
                        if (v16 && v17) {
                            v33 := GenS2{F0: (!v16), F1: 0}
                            if v16 {
                                fmt.Println(fi21, v15, v2)
                                v4 -= 2147483647
                                if (!v17) {
                                    var v34 string = (v2 + ("" + "foo bar"))
                                    _ = v34
                                }
                            }
                            v35 := GenS2{F0: v16, F1: 793}
                            var v36 uint = 65535
                            switch ("" + v2) {
                            case "hello":
                                v20 += v20
                            case "\n\t":
                                v37 := v1[""]
                                _ = v37
                            }
                            _ = v33
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v15, v11)
                            fmt.Println(v17, v2, v16)
                            v38 := GenS3{F0: (int64(fi21) >= -29)}
                            if v17 {
                                fmt.Println(v17, v23)
                                var v39 float32 = 0.0
                                var v40 bool = (356 >= 998)
                                _ = v39
                                _ = v40
                            } else {
                                var v41 uint8 = 42
                                _ = v41
                            }
                            if (0 <= -961) {
                                v42 := GenS3{F0: v17}
                                var v43 uint = 2
                                _ = v42
                                _ = v43
                            } else {
                                fmt.Println(fi21, v15)
                            }
                            _ = v38
                        }
                        var v44 int32 = 42
                        _ = v44
                    }
                    v2 = (v2 + v2)
                    if (false || (18446744073709551615 != 94)) {
                        switch (fi21 + fi22) {
                        case -954:
                            var v45 int8 = -76
                            _ = v45
                        }
                        if true {
                            switch "" {
                            case "a":
                                v11 -= 0
                            case "foo bar":
                                v46 := GenS4{F0: 255}
                                _ = v46
                            case "hello":
                                if ((float32(float64(float32(3.14159))) + (-2.25 / 1e10)) > 1.5) {
                                    v47 := GenS4{F0: 2}
                                    _ = v47
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = (("\n\t" + v23) + ("\n\t" + ""))
                                    }
                                    var v48 uint16 = 42
                                    _ = v48
                                }
                            }
                            var v49 int = fi22
                            _ = v49
                        }
                        if (2 == 32767) {
                            var v50 bool = v17
                            var v51 float64 = float64(((-2.25 - 0.0) + 0.0))
                            _ = v50
                            _ = v51
                        }
                        var v52 uint64 = 1
                        v15 += (v15 - v15)
                        switch -9223372036854775808 {
                        case -944:
                            _ = len(v3)
                        case 32767:
                            switch 2147483647 {
                            case -856:
                                v17 = true
                            case 2147483647:
                                v53 := GenS5{F0: "", F1: 781, F2: v4}
                                v17 = (v52 == (v52 / ((v52) | 1)))
                                if (127 <= 996) {
                                    var v54 int16 = (^(v15))
                                    _ = v54
                                } else {
                                    v16 = v16
                                }
                                _ = v53
                            case 1:
                                v5[2] = (v11 << ((42) & 63))
                            }
                        default:
                            v55 := GenS5{F0: "\n\t", F1: 836, F2: fi22}
                            _ = v55
                        }
                        _ = v52
                    } else {
                        v3 = append(v3, (v23 + v23))
                        var v56 uint = 174
                        if (!(false || (127 < 0))) {
                            var v57 uint16 = 65535
                            _ = v57
                        } else {
                            var v58 uint16 = 102
                            v58 = (v58 * (v58 ^ (v58 & (v58 ^ v58))))
                            _ = v58
                        }
                        _ = v56
                    }
                    if v16 {
                        v16 = v17
                        v20 = v20
                        v17 = (1 <= 127)
                    }
                    v2 = v23
                }
                var v59 uint8 = 1
                _ = v23
                _ = v59
            }
            switch fi22 {
            case 1:
                var v60 uint32 = (v11 - (v11 & (^(v11))))
                _ = v60
            case -9223372036854775808:
                v3 = append(v3, v2)
            case 32767:
                if false {
                    v15 = -1
                    v61 := v5[65535]
                    _ = v61
                } else {
                    v20 -= v20
                    var v62 string = (((("" + v2) + v2) + v2) + "")
                    _ = v62
                }
            default:
                var v63 string = v2
                _ = v63
            }
            _ = fi22
        }
        for fi64 := 0; fi64 < 4; fi64++ {
            fmt.Println(v16, v4, v4)
            var v65 int32 = 42
            var v66 float64 = v20
            _ = fi64
            _ = v65
            _ = v66
        }
        _ = fi21
    }
    switch 69 {
    case -1:
        var v67 uint = 2
        _ = v67
    case 1:
        var v68 uint32 = v11
        _ = v68
    default:
        if v16 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi69 := 0; fi69 < 20; fi69++ {
                switch v2 {
                case "hello":
                    var v70 int = int(v15)
                    _ = v70
                }
                for fi71 := 0; fi71 < 9; fi71++ {
                    _ = len(v3)
                    _ = fi71
                }
                _ = fi69
            }
        } else {
            var v72 uint8 = 128
            for fi73 := 0; fi73 < 17; fi73++ {
                v74 := GenS6{F0: 1, F1: ("\n\t" + v2)}
                _ = fi73
                _ = v74
            }
            _ = v72
        }
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
