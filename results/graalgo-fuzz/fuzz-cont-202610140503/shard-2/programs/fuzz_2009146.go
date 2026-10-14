package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []uint32{785}
    var v2 int8 = 110
    v3 := []uint32{0, 0, 42}
    var v4 bool = (1 <= 42)
    v5 := GenS1{F0: 255}
    v6 := GenS1{F0: 251}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 float64 = 1.5
    var v8 float32 = (1.5 - 1.5)
    var v9 uint = 388
    fmt.Println(v2)
    if v4 {
        var v10 int = 128
        _ = v10
    } else {
        switch 32768 {
        case 623:
            if v4 {
                v7 = float64(v8)
                switch 540 {
                case 32768:
                    var v11 float64 = (v7 - 0.0)
                    _ = v11
                case -9223372036854775808:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        fmt.Println(v9, fi12)
                        var v13 int64 = 1
                        _ = fi12
                        _ = v13
                    }
                default:
                    v4 = v4
                }
                var v14 int16 = 32767
                _ = v14
            } else {
                v8 -= v8
            }
        case 1:
            fmt.Println(v8, v8)
        }
        var v15 int64 = 1
        var v16 uint32 = 0
        _ = v15
        _ = v16
    }
    var v17 int32 = 1
    var v18 float64 = 1.5
    if (false || (v4 || (("" + ("" + ("hello" + "hello"))) == ("foo bar" + "")))) {
        var v19 uint = (v9 & v9)
        v20 := GenS1{F0: (v9 / (((v9 / ((v19) | 1))) | 1))}
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 int8 = v2
            _ = fi21
            _ = v22
        }
        switch 32767 {
        case 329:
            v17 = v17
        case -304:
            if (0 > 9223372036854775807) {
                v8 += (v8 + v8)
                if true {
                    var v23 string = ("a" + ("a" + "foo bar"))
                    if (-548 == -45) {
                        var v24 uint = v9
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            if (((1e10 * v8) / v8) == 3.14159) {
                                var v26 uint = v9
                                v27 := GenS2{F0: 32767, F1: 42}
                                if (65535 != 65535) {
                                    v28 := GenS1{F0: (v26 / ((42) | 1))}
                                    switch fi25 {
                                    case 4294967295:
                                        for fi29 := 0; fi29 < 17; fi29++ {
                                            if v4 {
                                                v1 = append(v1, 42)
                                            }
                                            var v30 int16 = -32768
                                            if ((v30 | (v30 + v30)) < 833) {
                                                fmt.Println(v30, fi29)
                                            } else {
                                                v3 = append(v3, 42)
                                                var v31 int = (-(fi29))
                                                _ = v31
                                            }
                                            if (-1 >= 1) {
                                                v4 = true
                                                var v32 int32 = (^(v17))
                                                fmt.Println(v32, v2)
                                                _ = v32
                                            }
                                            _ = fi29
                                            _ = v30
                                        }
                                    default:
                                        var v33 int16 = 1000
                                        _ = v33
                                    }
                                    _ = v28
                                }
                                _ = v26
                                _ = v27
                            }
                            _ = fi25
                        }
                    }
                    if v4 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        switch -84 {
                        case 70:
                            switch -651 {
                            case 127:
                                v34 := GenS2{F0: 1, F1: 0}
                                _ = v34
                            case -925:
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    var v36 int64 = 949
                                    if false {
                                        v4 = (!true)
                                        if ((fi35 == fi35) || false) {
                                            v2 -= (v2 | v2)
                                            var v37 int = (fi35 ^ int(v9))
                                            if len(v1) > 0 {
                                                v1[0] = uint32(v9)
                                            }
                                            _ = v37
                                        }
                                        if v4 {
                                            v17 -= (v17 << ((v9) & 63))
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            v2 += (v2 / ((int8(v17)) | 1))
                                        case "a":
                                            v8 -= (0.0 / ((v8 + v8) / (v8 - float32(v18))))
                                        case "\n\t":
                                            for fi38 := 0; fi38 < 8; fi38++ {
                                                v9 -= 948
                                                if v4 {
                                                    v23 = v23
                                                    v36 = v36
                                                    var v39 string = (((("hello" + v23) + ((("\n\t" + "a") + v23) + v23)) + v23) + v23)
                                                    var v40 int32 = v17
                                                    _ = v39
                                                    _ = v40
                                                }
                                                var v41 uint64 = 42
                                                _ = fi38
                                                _ = v41
                                            }
                                        default:
                                            fmt.Println(fi35, v9)
                                        }
                                        v8 += v8
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v42 uint32 = 65535
                                    _ = fi35
                                    _ = v36
                                    _ = v42
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 2147483648
                                }
                            }
                        case 128:
                            _ = len(v3)
                        case -2147483648:
                            if true {
                                switch 32768 {
                                case 4294967295:
                                    v3 = append(v3, 42)
                                }
                                v43 := GenS2{F0: 42, F1: 42}
                                var v44 int64 = 42
                                _ = v43
                                _ = v44
                            } else {
                                for fi45 := 0; fi45 < 16; fi45++ {
                                    fmt.Println(v9, v9, fi45)
                                    v46 := GenS1{F0: v9}
                                    var v47 int64 = -275
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                                for fi48 := 0; fi48 < 7; fi48++ {
                                    var v49 int = fi48
                                    _ = fi48
                                    _ = v49
                                }
                            }
                        default:
                            var v50 uint16 = 255
                            _ = v50
                        }
                    }
                    _ = v23
                }
                v51 := GenS2{F0: 32767, F1: 42}
                _ = v51
            }
        case -2147483648:
            v52 := GenS3{F0: 0, F1: 255}
            _ = v52
        }
    }
    if ((v2 / ((v2) | 1)) < v2) {
        if v4 {
            var v53 uint64 = 4294967295
            _ = v53
        }
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 float32 = 0.1
            _ = fi54
            _ = v55
        }
        if ((9223372036854775807 <= 42) || false) {
            fmt.Println(v18, v9)
            if len(v1) > 0 {
                v1[0] = 739
            }
            var v56 uint = v9
            _ = v56
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
