package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int16, p1 uint64, p2 uint64) uint16 {
    p1 = (p1 >> ((1) & 63))
    p2 = (p1 | 0)
    var v4 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint16(v4)
}

func main() {
    v1 := []int{128, 32768, 128}
    v2 := make(map[int]string)
    var v3 uint = 2
    fmt.Println(v3)
    v3 -= (v3 % (((v3 & 791)) | 1))
    switch 2147483647 {
    case 9223372036854775807:
        var v5 int8 = 127
        _ = v5
    case -2147483648:
        v3 = (v3 - v3)
    case -978:
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int16 = 32767
                v8 = v8
                var v9 uint16 = 32767
                v8 -= -32768
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := v2[fi6]
            _ = fi6
            _ = v10
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int8 = 95
    switch "hello" {
    case "hello":
        v11 = (v11 ^ 42)
    case "a":
        v12 := GenS1{F0: 150}
        _ = v12
    case "":
        fmt.Println(v11, v11)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            fmt.Println(fi13, v3)
        }
        _ = fi13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if true {
            v3 += 2
            v11 -= v11
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v2[fi14] = (("foo bar" + "\n\t") + "\n\t")
                var v15 float64 = (3.14159 - float64((1.5 - float32(1.5))))
                for fi16 := 0; fi16 < 2; fi16++ {
                    v3 = 255
                    v17 := v2[191]
                    fmt.Println(v3, v11, v17)
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        } else {
            if false {
                if (128 > fi14) {
                    switch (("foo bar" + "") + "hello") {
                    case "a":
                        v11 = (v11 / ((v11) | 1))
                    }
                    var v18 uint32 = 616
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        fmt.Println(fi14, fi14)
                        _ = fi19
                    }
                    if (65535 != 42) {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v11 = (v11 * 100)
                            if len(v1) > 0 {
                                v1[0] = 618
                            }
                            var v21 int64 = 9223372036854775807
                            var v22 int32 = -2147483648
                            v11 = v11
                            var v23 uint = v3
                            v22 -= -1
                            if (!(true || false)) {
                                if false {
                                    var v24 int = -9223372036854775808
                                    v24 += fi14
                                    _ = v24
                                } else {
                                    var v25 float32 = 1e10
                                    var v26 uint32 = 1
                                    _ = v25
                                    _ = v26
                                }
                                v23 -= v3
                                if false {
                                    v27 := GenS2{F0: 360}
                                    _ = v27
                                } else {
                                    var v28 int16 = 1000
                                    if (("\n\t" + "hello") >= ("hello" + "a")) {
                                        v11 = (v11 ^ (v11 & 127))
                                        switch fi14 {
                                        case -399:
                                            if (147 != (v3 >> ((1) & 63))) {
                                                var v29 int8 = (^(v11))
                                                _ = v29
                                            } else {
                                                v30 := GenS3{F0: 9223372036854775807}
                                                v11 = 85
                                                var v31 float32 = float32(0.1)
                                                v32 := GenS2{F0: 2147483648}
                                                v33 := GenS4{F0: (v21 & (v21 >> (((^(v23))) & 63)))}
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                        case 9223372036854775807:
                                            switch (("" + "") + ("" + "")) {
                                            case "\n\t":
                                                if (2 < (v23 << (((v23 ^ 710)) & 63))) {
                                                    v21 = (v21 - (^(v21)))
                                                    if ((v11 << (((^(v23))) & 63)) >= v11) {
                                                        var v34 uint64 = 0
                                                        _ = v34
                                                    }
                                                    var v35 float64 = 1.5
                                                    _ = v35
                                                }
                                            case "a":
                                                _ = len(v2)
                                            default:
                                                v21 += v21
                                                v28 = 1
                                                _ = len(v2)
                                                fmt.Println(v22, v22)
                                                var v36 uint8 = 127
                                                fmt.Println(fi14)
                                                fmt.Println(v21)
                                                _ = v36
                                            }
                                            v28 -= (v28 & 42)
                                            var v37 int64 = 0
                                            var v38 string = "a"
                                            v11 = v11
                                            _ = v37
                                            _ = v38
                                        case 4294967295:
                                            fmt.Println(v23, v3, v3)
                                        default:
                                            var v39 string = "foo bar"
                                            _ = v39
                                        }
                                    } else {
                                        var v40 float64 = (float64((0.0 * (1e10 * 3.14159))) + 1.5)
                                        fmt.Println(fi20)
                                        if (!(!(255 >= 862))) {
                                            fmt.Println(v11, v28, fi14)
                                            if (((fi20 ^ (fi14 / (((^(fi14))) | 1))) < fi14) && true) {
                                                fmt.Println(v3, v40)
                                                var v41 uint64 = 42
                                                v23 += (^(v23))
                                                _ = v41
                                            } else {
                                                v42 := GenS2{F0: 2}
                                                _ = v42
                                            }
                                            v43 := v2[(fi20 * -1)]
                                            _ = v43
                                        } else {
                                            fmt.Println(v21, v28)
                                            fmt.Println(v3)
                                        }
                                        v2[fi20] = ""
                                        v3 = v3
                                        _ = v40
                                    }
                                    v23 -= v23
                                    _ = v28
                                }
                            }
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    } else {
                        fmt.Println(v11, v11, v11)
                        var v44 uint = v3
                        if ((42 >= -1) && (!((true && false) && false))) {
                            v3 += 50
                            var v45 uint32 = 2
                            _ = v45
                        }
                        _ = v44
                    }
                    for fi46 := 0; fi46 < 7; fi46++ {
                        var v47 uint64 = 331
                        v47 = v47
                        _ = fi46
                        _ = v47
                    }
                }
            } else {
                v3 = (v3 | v3)
                v48 := GenS2{F0: 416}
                _ = v48
            }
            if false {
                for fi49 := 0; fi49 < 3; fi49++ {
                    v3 = v3
                    _ = fi49
                }
            }
            v3 = v3
        }
        for fi50 := 0; fi50 < 12; fi50++ {
            switch "" {
            case "":
                fmt.Println(fi50, v11, v3)
            case "foo bar":
                var v51 int64 = 9223372036854775807
                _ = v51
            case "\n\t":
                fmt.Println(v3, v3)
            default:
                v52 := GenS3{F0: 0}
                _ = v52
            }
            v11 = (v11 ^ v11)
            _ = fi50
        }
        switch (^(fi14)) {
        case -620:
            if false {
                for fi53 := 0; fi53 < 12; fi53++ {
                    v54 := v2[fi53]
                    v55 := v2[fi14]
                    _ = fi53
                    _ = v54
                    _ = v55
                }
            } else {
                var v56 int32 = -2147483648
                v56 += (v56 << (((v3 * v3)) & 63))
                _ = v56
            }
        }
        _ = fi14
    }
    v11 = (-(v11))
    if true {
        for fi57 := 0; fi57 < 19; fi57++ {
            v58 := GenS2{F0: 181}
            var v59 bool = ((!(fi57 < fi57)) || ((0 > (v3 % ((v3) | 1))) && true))
            _ = fi57
            _ = v58
            _ = v59
        }
        var v60 uint16 = 809
        fmt.Println(v3)
        _ = v60
    }
    var v61 uint8 = 126
    var v62 uint8 = v61
    var v63 uint32 = 2
    v3 = (v3 ^ v3)
    var v64 uint = (v3 + uint(v63))
    if true {
        if false {
            var v65 bool = (218 != 127)
            var v66 uint32 = (v63 % (((v63 / (((v63 | v63)) | 1))) | 1))
            if (-2147483648 < -1) {
                v65 = v65
                _ = len(v1)
                var v67 string = "hello"
                _ = v67
            } else {
                v62 = 255
            }
            _ = v65
            _ = v66
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
