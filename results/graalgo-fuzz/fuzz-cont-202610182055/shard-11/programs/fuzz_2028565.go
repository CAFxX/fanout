package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int; F1 uint32; F2 int32 }
type GenS10 struct { F0 int64; F1 int8 }
func genFn0(p0 string) int16 {
    v5 := []int{645}
    if true {
        switch 32767 {
        case 0:
            var v6 uint = 0
            _ = v6
        case 1:
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 12; fi8++ {
                    switch fi7 {
                    case 32768:
                        p0 = "foo bar"
                    case -1:
                        switch p0 {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = (fi7 - fi8)
                            }
                        case "foo bar":
                            if (("\n\t" + "\n\t") != "foo bar") {
                                var v9 uint = 255
                                p0 = p0
                                var v10 uint64 = 2
                                _ = v10
                                _ = v9
                            } else {
                                p0 = ("foo bar" + p0)
                                switch (fi8 * 127) {
                                case -2147483648:
                                    v11 := GenS1{F0: 127, F1: -128}
                                    _ = v11
                                case 32767:
                                    var v12 int = fi8
                                    _ = v12
                                }
                                var v13 uint8 = 42
                                _ = v13
                            }
                        case "\n\t":
                            switch fi7 {
                            case -2147483648:
                                var v14 int32 = -737
                                _ = v14
                            }
                        }
                    default:
                        var v15 string = p0
                        if true {
                            v16 := GenS2{F0: (fi7 / (((fi7 + fi7)) | 1)), F1: -405, F2: 32767}
                            v5 = append(v5, fi7)
                            var v17 int64 = -996
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = fi8
                }
                fmt.Println(p0, p0)
                _ = fi7
            }
        case -1:
            if (!(65535 < 18446744073709551615)) {
                p0 = "\n\t"
                v5 = append(v5, 9223372036854775807)
                _ = len(v5)
            } else {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: 128, F1: 127}
                    fmt.Println(fi18)
                    fmt.Println(p0, fi18)
                    _ = fi18
                    _ = v19
                }
            }
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 uint64 = 339
                var v23 float32 = float32(1e10)
                v24 := GenS3{F0: -569, F1: 272, F2: int32(fi20)}
                if false {
                    v22 = v22
                    if (!true) {
                        var v25 uint8 = 42
                        var v26 uint32 = 65535
                        _ = v25
                        _ = v26
                    }
                }
                if true {
                    v27 := GenS3{F0: -1, F1: 4294967295, F2: 42}
                    v28 := GenS4{F0: 42, F1: 95}
                    _ = v27
                    _ = v28
                }
                if true {
                    if true {
                        var v29 uint64 = 734
                        _ = v29
                    }
                }
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(p0, p0)
            p0 = p0
            _ = fi20
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 int32 = -1
                switch ("" + "\n\t") {
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = (fi30 ^ (fi31 * (fi30 + (-(fi30)))))
                    }
                }
                v33 := GenS5{F0: -1, F1: 289}
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi30
        }
    } else {
        for fi34 := 0; fi34 < 4; fi34++ {
            switch fi34 {
            case 127:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case -385:
                p0 = (p0 + "")
            case -195:
                p0 = p0
            default:
                var v35 int16 = -1
                _ = v35
            }
            _ = fi34
        }
    }
    _ = p0
    _ = v5
    return -889
}
func genFn1(p0 uint8, p1 int64) uint {
    var v36 int32 = 1
    fmt.Println(v36)
    p1 -= (p1 - -981)
    p0 -= 1
    _ = p0
    _ = p1
    _ = v36
    return 255
}

func main() {
    var v1 int16 = 847
    var v2 float64 = 0.1
    var v3 int = 1
    var v4 uint8 = 0
    fmt.Println(v4, v3, v4)
    fmt.Println(v1)
    for fi37 := 0; fi37 < 15; fi37++ {
        v1 = (v1 % ((int16(fi37)) | 1))
        if (132 <= 65535) {
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := GenS6{F0: -613, F1: 42, F2: 2147483647}
                v40 := []int{(v3 ^ (^(v3))), (fi37 & (fi38 * fi37)), fi38}
                _ = fi38
                _ = v39
                _ = v40
            }
            var v41 int = (fi37 - (fi37 ^ fi37))
            _ = v41
        }
        _ = fi37
    }
    var v42 string = "foo bar"
    v43 := []int{}
    fmt.Println(v1)
    if true {
        switch v42 {
        case "hello":
            if len(v43) >= 2 {
                v43 = v43[1:len(v43)]
            }
        case "a":
            fmt.Println(v1, v3)
        }
        var v44 int64 = 0
        switch ((v42 + ("" + (v42 + "a"))) + "hello") {
        case "":
            var v45 int16 = v1
            _ = v45
        default:
            if (0 >= 943) {
                v42 = (("a" + v42) + "foo bar")
            } else {
                for fi46 := 0; fi46 < 10; fi46++ {
                    fmt.Println(v2, fi46, v44)
                    v3 -= fi46
                    for fi47 := 0; fi47 < 7; fi47++ {
                        var v48 int8 = 0
                        _ = fi47
                        _ = v48
                    }
                    _ = fi46
                }
                for fi49 := 0; fi49 < 3; fi49++ {
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v51 := GenS5{F0: v44, F1: 495}
                        _ = fi50
                        _ = v51
                    }
                    _ = fi49
                }
                v1 -= (^(v1))
            }
        }
        _ = v44
    }
    v52 := GenS7{F0: 42}
    var v53 float32 = 1.5
    v54 := GenS2{F0: 749, F1: 2147483647, F2: 0}
    switch v3 {
    case 4294967295:
        switch 619 {
        case 32767:
            v55 := GenS5{F0: -1, F1: 468}
            _ = v55
        case -976:
            v56 := GenS8{F0: 27, F1: 18446744073709551615}
            _ = v56
        case 1:
            for fi57 := 0; fi57 < 15; fi57++ {
                var v58 int = (v3 & v3)
                v53 += (v53 / ((1e10 / (float32(v2) * 1.5)) / 3.14159))
                _ = fi57
                _ = v58
            }
        }
    case 128:
        switch v42 {
        case "a":
            var v59 float32 = -2.25
            _ = v59
        default:
            var v60 float64 = (1e10 + float64(v53))
            _ = v60
        }
    case -2147483648:
        v3 += v3
    }
    v3 = v3
    v61 := GenS9{F0: v3, F1: 0, F2: 42}
    v62 := GenS10{F0: 1, F1: 100}
    for fi63 := 0; fi63 < 1; fi63++ {
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
        _ = fi63
    }
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    v53 -= v53
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v52
    _ = v53
    _ = v54
    _ = v61
    _ = v62
}
