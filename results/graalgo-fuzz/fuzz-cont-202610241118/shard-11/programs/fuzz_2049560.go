package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int8 }
func genFn0() uint8 {
    v6 := []int{-688, 32768}
    if false {
        _ = len(v6)
        var v7 int8 = -1
        _ = v7
    }
    if (42 > 280) {
        var v8 int32 = 0
        switch 127 {
        case 819:
            if false {
                if true {
                    v9 := GenS1{F0: 0}
                    if (48 == 214) {
                        if true {
                            if (false && true) {
                                var v10 string = "\n\t"
                                if (1 < 42) {
                                    switch 2147483647 {
                                    case 128:
                                        v10 = (((v10 + v10) + ("a" + ("\n\t" + v10))) + ((v10 + ("" + "foo bar")) + v10))
                                    case -9223372036854775808:
                                        switch 42 {
                                        case 4294967295:
                                            if len(v6) > 0 {
                                                v6[0] = -2147483648
                                            }
                                        case -1:
                                            if false {
                                                v6 = append(v6, int(v8))
                                            }
                                        case 629:
                                            fmt.Println(v10, v10)
                                        default:
                                            fmt.Println(v8, v10)
                                        }
                                    default:
                                        v6 = append(v6, -2147483648)
                                    }
                                } else {
                                    var v11 int64 = -1
                                    for fi12 := 0; fi12 < 16; fi12++ {
                                        fmt.Println(v10, v11, v11)
                                        v6 = append(v6, (^(fi12)))
                                        _ = fi12
                                    }
                                    _ = v11
                                }
                                _ = v10
                            }
                            for fi13 := 0; fi13 < 18; fi13++ {
                                v14 := GenS1{F0: 430}
                                fmt.Println(v8, v8)
                                _ = fi13
                                _ = v14
                            }
                            _ = len(v6)
                        }
                        for fi15 := 0; fi15 < 13; fi15++ {
                            fmt.Println(fi15)
                            v16 := GenS2{F0: 4294967295}
                            v17 := GenS3{F0: "hello"}
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    } else {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v6 = append(v6, (fi18 | fi18))
                            switch "foo bar" {
                            case "a":
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    v8 -= (v8 >> ((493) & 63))
                                    _ = fi19
                                }
                            case "foo bar":
                                var v20 bool = ((0 != 1) || (!true))
                                _ = v20
                            default:
                                v6 = append(v6, -258)
                            }
                            _ = fi18
                        }
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v8 = v8
                            if ((true && (2 < 0)) || (true && (v8 > v8))) {
                                var v22 uint64 = 42
                                v6 = append(v6, fi21)
                                _ = v22
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v23 uint32 = 2147483647
                                v23 += v23
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v8 -= v8
                                switch ("" + "hello") {
                                case "hello":
                                    if (100 < 42) {
                                        switch (-(fi21)) {
                                        case 42:
                                            v8 += (v8 / ((v8) | 1))
                                        case -77:
                                            fmt.Println(v8, v8, fi24)
                                        case -518:
                                            v8 -= v8
                                        default:
                                            var v25 int16 = 1
                                            _ = v25
                                        }
                                    } else {
                                        v8 = (v8 / (((v8 - v8)) | 1))
                                        var v26 int8 = 0
                                        _ = v26
                                    }
                                default:
                                    var v27 int16 = 522
                                    _ = v27
                                }
                                v28 := GenS3{F0: ("" + ("\n\t" + "\n\t"))}
                                _ = fi24
                                _ = v28
                            }
                            _ = fi21
                        }
                        var v29 uint = 247
                        _ = v29
                    }
                    _ = v9
                } else {
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                }
            }
        case -1:
            fmt.Println(v8)
        case 32768:
            var v30 int8 = 0
            _ = v30
        default:
            var v31 float64 = (3.14159 / 1e10)
            _ = v31
        }
        v32 := GenS1{F0: 248}
        _ = v32
        _ = v8
    } else {
        if false {
            if true {
                v33 := GenS3{F0: ("foo bar" + ("" + "foo bar"))}
                _ = v33
            } else {
                _ = len(v6)
            }
        } else {
            var v34 uint16 = 1
            _ = v34
        }
        v35 := GenS1{F0: 255}
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        if true {
            var v37 float64 = (-2.25 / (-2.25 + 3.14159))
            switch "hello" {
            case "a":
                v37 = v37
            default:
                v38 := GenS1{F0: 65535}
                _ = v38
            }
            if true {
                var v39 string = "foo bar"
                _ = v39
            }
            _ = v37
        }
        var v40 uint64 = 1
        fmt.Println(fi36, v40)
        _ = fi36
        _ = v40
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 int = -2147483648
    var v3 float32 = 0.0
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    if true {
        for fi41 := 0; fi41 < 20; fi41++ {
            if ((v1 % ((v1) | 1)) != (v1 - (^(v1)))) {
                v1 += v1
                v3 = (1e10 / 1.5)
            }
            v3 += v3
            v42 := []int{(v2 & 667)}
            _ = fi41
            _ = v42
        }
    } else {
        v1 -= v1
        var v43 int8 = 1
        v2 += 389
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        if (0 <= 1) {
            fmt.Println(v3, fi44, v1)
            if (!(!(255 >= 76))) {
                fmt.Println(v3)
                var v45 float64 = 1e10
                fmt.Println(v2, fi44, fi44)
                _ = v45
            }
        } else {
            for fi46 := 0; fi46 < 6; fi46++ {
                v3 -= (v3 + 3.14159)
                _ = fi46
            }
            var v47 uint32 = 1
            switch "\n\t" {
            case "":
                v48 := GenS2{F0: (^(v47))}
                _ = v48
            }
            _ = v47
        }
        v3 += -2.25
        _ = fi44
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        switch 0 {
        case 9223372036854775807:
            switch "\n\t" {
            case "\n\t":
                var v50 float64 = 1e10
                _ = v50
            default:
                if false {
                    var v51 int32 = -1
                    _ = v51
                } else {
                    v1 -= (v1 % (((v1 / ((v1) | 1))) | 1))
                    var v52 int = v2
                    if (1e10 < (0.0 / 1.0)) {
                        if (128 > genFn0()) {
                            var v53 uint16 = 1
                            _ = v53
                        }
                        v54 := GenS4{F0: 2147483647, F1: 694}
                        var v55 bool = false
                        _ = v54
                        _ = v55
                    } else {
                        v56 := GenS5{F0: 100}
                        fmt.Println(fi49, v52)
                        switch v52 {
                        case -234:
                            delete(v5, (v2 | v2))
                        }
                        _ = v56
                    }
                    _ = v52
                }
            }
        case -2147483648:
            fmt.Println(v3, fi49, v3)
        }
        var v57 int8 = 118
        var v58 int32 = 1
        _ = fi49
        _ = v57
        _ = v58
    }
    if false {
        v59 := GenS1{F0: 420}
        v3 = (float32(-2.25) * (v3 / ((float32(-2.25) / v3) * ((1e10 / v3) * v3))))
        fmt.Println(v3)
        _ = v59
    }
    v1 += 471
    v1 = (^(v1))
    v1 -= 1
    if (1 < 255) {
        var v60 int16 = -32768
        _ = v60
    } else {
        var v61 float64 = -2.25
        switch -302 {
        case 127:
            v2 -= (v2 << ((0) & 63))
        case -2147483648:
            switch 846 {
            case 2147483647:
                var v62 uint16 = 65535
                _ = v62
            case 32768:
                if (!(true || true)) {
                    v63 := GenS3{F0: "foo bar"}
                    v1 = v1
                    _ = v63
                } else {
                    v64 := GenS3{F0: "hello"}
                    _ = v64
                }
            case -9223372036854775808:
                var v65 uint64 = (v1 + v1)
                _ = v65
            }
        case 42:
            v3 = 1.5
        default:
            var v66 int = v2
            _ = v66
        }
        _ = v61
    }
    switch v2 {
    case -9223372036854775808:
        if true {
            fmt.Println(v3, v2, v2)
        }
    }
    switch (v2 << ((42) & 63)) {
    case 568:
        var v67 int16 = 1
        _ = v67
    case -2147483648:
        var v68 int = 397
        _ = v68
    default:
        if ((v1 <= (v1 - (v1 >> ((4294967295) & 63)))) || false) {
            var v69 int16 = -390
            _ = v69
        }
    }
    if false {
        var v70 string = ("" + "\n\t")
        _ = v70
    }
    v2 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
