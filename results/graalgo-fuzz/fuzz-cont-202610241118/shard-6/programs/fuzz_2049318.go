package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
func genFn0(p0 int64, p1 int) string {
    p1 = (^(p1))
    if true {
        v5 := []int{(-(p1)), 32768, p1}
        if false {
            if false {
                var v6 int64 = p0
                if true {
                    v6 -= 0
                } else {
                    var v7 string = "a"
                    p1 = p1
                    v8 := GenS1{F0: p1}
                    _ = v7
                    _ = v8
                }
                _ = v6
            } else {
                if (!false) {
                    if false {
                        p1 -= 4294967295
                    } else {
                        v9 := GenS1{F0: p1}
                        _ = v9
                    }
                }
                var v10 uint = 867
                _ = v10
            }
            p0 = -1
        }
        if true {
            fmt.Println(p1)
            var v11 int = p1
            for fi12 := 0; fi12 < 11; fi12++ {
                v11 -= -2147483648
                _ = fi12
            }
            _ = v11
        } else {
            var v13 uint64 = 2
            v13 = v13
            p1 += (p1 | p1)
            _ = v13
        }
        _ = v5
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 127
    var v3 uint8 = (v2 % ((v2) | 1))
    var v4 float64 = 0.0
    v14 := GenS1{F0: 32768}
    v15 := GenS2{F0: ("foo bar" + (("hello" + "a") + "a")), F1: -63}
    var v16 uint8 = v3
    var v17 string = "\n\t"
    if false {
        v16 += 42
        v18 := []int{32767, int(v16), -1}
        var v19 uint32 = 2
        _ = v18
        _ = v19
    }
    v20 := []int{-598, 42, 2147483647, -9223372036854775808}
    v2 = (^(v2))
    if (127 <= 4294967295) {
        var v21 bool = (true || false)
        if (false && false) {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 string = v17
                var v24 uint = 18446744073709551615
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(v21, v16, v21)
        }
        if ((0.1 - v4) < v4) {
            if v21 {
                if v21 {
                    var v25 uint = 255
                    _ = v25
                }
            } else {
                v3 = v3
                var v26 uint32 = 8
                v27 := GenS2{F0: "", F1: 1}
                _ = v26
                _ = v27
            }
        }
        _ = v21
    }
    switch -9223372036854775808 {
    case -796:
        v17 = ((("\n\t" + v17) + v17) + v17)
    case 0:
        for fi28 := 0; fi28 < 8; fi28++ {
            if true {
                if len(v20) > 0 {
                    v20[0] = 32767
                }
                _ = len(v20)
                for fi29 := 0; fi29 < 17; fi29++ {
                    switch -9223372036854775808 {
                    case 127:
                        if (float32(float64(1e10)) >= 0.1) {
                            v20 = append(v20, (fi28 >> ((255) & 63)))
                        }
                    case -212:
                        if false {
                            if ((fi28 | (fi28 / ((fi29) | 1))) >= 32767) {
                                var v30 uint16 = 255
                                if true {
                                    if ((float64((3.14159 + (float32(v4) * 0.1))) / v4) >= (((1e10 * v4) / ((0.0 + ((v4 - v4) + v4)) * (v4 - v4))) + 3.14159)) {
                                        if (!(0 < (v30 / ((v30) | 1)))) {
                                            switch (-(fi29)) {
                                            case 127:
                                                v31 := GenS1{F0: (-(fi28))}
                                                _ = v31
                                            case -2147483648:
                                                v32 := GenS3{F0: 42, F1: 1, F2: 2}
                                                _ = v32
                                            }
                                            v20 = append(v20, (-(fi28)))
                                        } else {
                                            var v33 string = v17
                                            v17 = (v33 + ("\n\t" + (v17 + "a")))
                                            if len(v20) > 0 {
                                                v20[0] = fi28
                                            }
                                            _ = v33
                                        }
                                    } else {
                                        v17 = v17
                                        if len(v20) > 0 {
                                            v20[0] = (fi29 + (fi28 | fi28))
                                        }
                                    }
                                } else {
                                    v34 := GenS4{F0: 255, F1: 0, F2: v30}
                                    switch (fi29 / ((fi28) | 1)) {
                                    case 128:
                                        _ = len(v20)
                                    case 2147483647:
                                        switch 4294967295 {
                                        case 42:
                                            switch v17 {
                                            case "\n\t":
                                                var v35 int16 = 32767
                                                _ = v35
                                            case "":
                                                var v36 int = (^(fi28))
                                                _ = v36
                                            case "hello":
                                                if (true || (v1 > v1)) {
                                                    if (510 == 0) {
                                                        if len(v20) > 0 {
                                                            v20[0] = 784
                                                        }
                                                        if true {
                                                            v37 := GenS3{F0: -87, F1: 534, F2: 42}
                                                            v20 = append(v20, fi29)
                                                            _ = v37
                                                        }
                                                        if false {
                                                            fmt.Println(fi28, v17, v3)
                                                            var v38 int64 = 147
                                                            _ = v38
                                                        } else {
                                                            _ = len(v20)
                                                            if false {
                                                                var v39 int64 = 1
                                                                v1 = (^(v1))
                                                                _ = v39
                                                            } else {
                                                                _ = len(v20)
                                                                v40 := GenS1{F0: 128}
                                                                _ = v40
                                                            }
                                                            if false {
                                                                v41 := GenS5{F0: 42, F1: int64(v30)}
                                                                _ = v41
                                                            } else {
                                                                if (fi28 == fi29) {
                                                                    switch fi28 {
                                                                    case 1:
                                                                        var v42 float64 = (((3.14159 + v4) * v4) - (v4 - 0.0))
                                                                        _ = v42
                                                                    default:
                                                                        var v43 int8 = 1
                                                                        _ = v43
                                                                    }
                                                                }
                                                                v44 := GenS3{F0: -104, F1: 1, F2: uint32(v30)}
                                                                var v45 int16 = -32768
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        }
                                                    }
                                                    var v46 uint16 = v30
                                                    _ = v46
                                                } else {
                                                    fmt.Println(fi29)
                                                }
                                            }
                                        default:
                                            if len(v20) > 0 {
                                                v20[0] = (fi29 * 0)
                                            }
                                        }
                                    }
                                    switch (-(fi28)) {
                                    case 127:
                                        fmt.Println(v16, v30, v3)
                                    default:
                                        if (v4 <= ((v4 - v4) + v4)) {
                                            v2 = (v16 ^ v2)
                                        } else {
                                            v30 -= v30
                                        }
                                    }
                                    _ = v34
                                }
                                _ = v30
                            } else {
                                if false {
                                    switch (-(fi28)) {
                                    case 883:
                                        var v47 int = fi29
                                        _ = v47
                                    default:
                                        v4 = (-2.25 * v4)
                                    }
                                    switch v17 {
                                    case "":
                                        if false {
                                            v17 = v17
                                        }
                                    }
                                    var v48 uint = 860
                                    _ = v48
                                }
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            }
                            switch ("" + ("\n\t" + ("a" + "\n\t"))) {
                            case "hello":
                                if len(v20) > 0 {
                                    v20[0] = fi28
                                }
                            default:
                                if len(v20) > 0 {
                                    v20[0] = (fi29 - fi29)
                                }
                            }
                            v17 = v17
                        }
                    case 0:
                        v4 += 0.0
                    default:
                        fmt.Println(v4, v2, v4)
                        var v49 uint32 = 65535
                        _ = v49
                    }
                    v4 += (v4 * v4)
                    _ = fi29
                }
            } else {
                if true {
                    v17 = v17
                    if true {
                        var v50 uint64 = 576
                        _ = v50
                    }
                } else {
                    v2 = v3
                }
                fmt.Println(v3)
                var v51 int = 881
                _ = v51
            }
            if false {
                var v52 uint16 = 32767
                if (-9223372036854775808 > 461) {
                    v1 += 1
                    if (1.5 > v4) {
                        switch ((v17 + "foo bar") + "\n\t") {
                        case "\n\t":
                            v1 = 0
                        case "foo bar":
                            v53 := GenS1{F0: 42}
                            _ = v53
                        default:
                            switch (fi28 - 32767) {
                            case 127:
                                var v54 float64 = 3.14159
                                _ = v54
                            default:
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    var v56 int32 = (-(v1))
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        }
                    } else {
                        if (!false) {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        }
                    }
                    for fi57 := 0; fi57 < 1; fi57++ {
                        var v58 float64 = v4
                        v16 -= (v2 << ((65535) & 63))
                        switch "foo bar" {
                        case "a":
                            if (18446744073709551615 == 42) {
                                fmt.Println(v17)
                                v20 = append(v20, (fi28 ^ fi57))
                                switch fi28 {
                                case 2147483647:
                                    var v59 int16 = 1
                                    _ = v59
                                }
                            } else {
                                _ = len(v20)
                                v17 = ((("foo bar" + v17) + (v17 + (v17 + ("a" + (v17 + v17))))) + ("\n\t" + v17))
                            }
                        }
                        if false {
                            var v60 uint16 = 32767
                            var v61 string = ("a" + v17)
                            fmt.Println(v3, v3, fi28)
                            switch "a" {
                            case "a":
                                if false {
                                    if false {
                                        switch fi28 {
                                        case 42:
                                            var v62 int16 = -263
                                            _ = v62
                                        case 127:
                                            v3 += (v2 / ((98) | 1))
                                        default:
                                            switch v61 {
                                            case "a":
                                                var v63 int16 = -32768
                                                _ = v63
                                            case "hello":
                                                var v64 float64 = ((v58 + 0.0) * float64(-2.25))
                                                _ = v64
                                            }
                                        }
                                    } else {
                                        fmt.Println(v3, v3)
                                        switch fi28 {
                                        case 32768:
                                            v65 := GenS6{F0: 2147483648}
                                            _ = v65
                                        case -124:
                                            var v66 int32 = v1
                                            _ = v66
                                        case 0:
                                            v60 -= (v60 << ((85) & 63))
                                            v17 = ""
                                            var v67 uint16 = v52
                                            _ = v67
                                        }
                                        v68 := GenS4{F0: 42, F1: 42, F2: (v52 >> ((1) & 63))}
                                        _ = v68
                                    }
                                    var v69 uint8 = 2
                                    _ = v69
                                } else {
                                    v16 += 200
                                    if false {
                                        switch v61 {
                                        case "foo bar":
                                            v4 = v58
                                        case "\n\t":
                                            var v70 uint32 = 2147483648
                                            _ = v70
                                        }
                                        v71 := GenS2{F0: v17, F1: 1}
                                        _ = v71
                                    } else {
                                        v58 = (float64(1e10) + v58)
                                        var v72 uint8 = v3
                                        _ = v72
                                    }
                                    var v73 uint32 = 944
                                    _ = v73
                                }
                            default:
                                fmt.Println(v1)
                            }
                            _ = v60
                            _ = v61
                        }
                        _ = fi57
                        _ = v58
                    }
                }
                var v74 uint32 = 2147483647
                _ = v52
                _ = v74
            }
            switch ("hello" + "a") {
            case "hello":
                v75 := GenS7{F0: 1}
                _ = v75
            }
            _ = fi28
        }
    default:
        v16 = 41
    }
    v76 := GenS8{F0: -159}
    v1 = (^(v1))
    fmt.Println(v16, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v76
}
