package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0() string {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(fi5)
            var v6 uint32 = 65535
            _ = fi5
            _ = v6
        }
        var v7 int8 = 127
        fmt.Println(v7, fi4)
        _ = fi4
        _ = v7
    }
    var v8 uint8 = 128
    var v9 uint = 652
    fmt.Println(v9)
    v9 = (^(v9))
    _ = v8
    _ = v9
    return ("" + "\n\t")
}
func genFn1(p0 uint8, p1 int32) int32 {
    v10 := []int{-2147483648, 4294967295}
    p1 = p1
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 2; fi12++ {
        if true {
            switch (fi12 << ((405) & 63)) {
            case -1:
                if true {
                    if (18446744073709551615 <= 1) {
                        if (true || false) {
                            switch "" {
                            case "\n\t":
                                var v13 string = (("\n\t" + ("" + "\n\t")) + "hello")
                                _ = v13
                            default:
                                var v14 uint8 = 211
                                _ = v14
                            }
                            v10 = append(v10, (fi12 << ((797) & 63)))
                            v15 := GenS1{F0: ("a" + ("\n\t" + "foo bar")), F1: 1}
                            _ = v15
                        } else {
                            fmt.Println(fi12, p1)
                            v11 = 1
                        }
                    } else {
                        if ((((-2.25 - -2.25) / 3.14159) - 0.1) < (((1e10 / (-2.25 - 0.0)) - ((1e10 * (1e10 - ((3.14159 - (3.14159 - ((1.5 * (1.5 * ((((((0.1 / float32(1e10)) / (1.5 - (0.1 + (0.1 * 1e10)))) - ((0.0 - float32((float64(0.0) - (1e10 - (1e10 - (1e10 + (float64(-2.25) - 3.14159))))))) / 1.0)) + -2.25) * -2.25) + -2.25))) - 0.0))) + 1e10))) + (1.5 * (-2.25 - (0.1 * ((0.1 + 1e10) + (3.14159 - 0.1))))))) / (((0.0 + -2.25) / 3.14159) + 1.5))) {
                            var v16 int8 = 1
                            for fi17 := 0; fi17 < 10; fi17++ {
                                p0 = p0
                                _ = fi17
                            }
                            if false {
                                v16 = v16
                                var v18 float64 = 1e10
                                var v19 float32 = -2.25
                                _ = v18
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 19; fi20++ {
                                    if len(v10) > 0 {
                                        v10[0] = (fi12 << ((65535) & 63))
                                    }
                                    switch ("" + "foo bar") {
                                    case "":
                                        v11 += (v11 - 488)
                                    case "a":
                                        v11 = 2147483648
                                    }
                                    var v21 uint16 = 255
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            _ = v16
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v11 -= (v11 / ((v11) | 1))
                        p0 += (^(p0))
                        _ = fi22
                    }
                    var v23 string = "a"
                    _ = v23
                }
            case -466:
                v11 += 2147483648
            }
        } else {
            fmt.Println(v11, p1, fi12)
            v24 := GenS2{F0: (684 <= 282), F1: -9223372036854775808}
            v25 := GenS3{F0: 619, F1: (!(true && (2 <= 87))), F2: 151}
            _ = v24
            _ = v25
        }
        _ = fi12
    }
    v10 = append(v10, 42)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p1
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 32767
    var v3 bool = false
    v3 = (128 < 255)
    switch -192 {
    case 389:
        v3 = (2 == 255)
    case 127:
        fmt.Println(v1)
    }
    if true {
        fmt.Println(v3)
    } else {
        switch 67 {
        case -9223372036854775808:
            fmt.Println(v2, v1, v2)
        default:
            v26 := []int{-818}
            _ = v26
        }
        var v27 int8 = 0
        _ = v27
    }
    var v28 int32 = int32(v2)
    v29 := GenS4{F0: 2147483647, F1: -431}
    v30 := []int{9223372036854775807, -9223372036854775808}
    var v31 uint = 18446744073709551615
    if v3 {
        if len(v30) > 0 {
            v30[0] = 126
        }
        v28 += v28
        var v32 int8 = 127
        _ = v32
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 uint32 = 1
            _ = fi33
            _ = v34
        }
        if (((3.14159 >= 3.14159) || v3) && v3) {
            v35 := GenS3{F0: 4294967295, F1: (!v3), F2: 255}
            _ = v35
        } else {
            v36 := GenS5{F0: v2}
            for fi37 := 0; fi37 < 4; fi37++ {
                v30 = append(v30, 449)
                v38 := GenS3{F0: 42, F1: v3, F2: v31}
                _ = fi37
                _ = v38
            }
            switch ("hello" + "hello") {
            case "":
                var v39 int16 = -1
                _ = v39
            case "\n\t":
                var v40 int = 1
                _ = v40
            default:
                _ = len(v30)
            }
            _ = v36
        }
    }
    switch 4294967295 {
    case 42:
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS3{F0: 18446744073709551615, F1: (true && v3), F2: v31}
            if true {
                v28 += v28
            }
            fmt.Println(v28, v3)
            _ = fi41
            _ = v42
        }
    case -2147483648:
        fmt.Println(v1, v1, v2)
    case 32768:
        if (-128 <= 127) {
            if len(v30) > 0 {
                v30[0] = 2147483647
            }
        }
    default:
        var v43 uint32 = 2147483647
        _ = v43
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        fmt.Println(v28, v31, v1)
        v45 := GenS6{F0: 4294967295, F1: 65535}
        _ = fi44
        _ = v45
    }
    var v46 int32 = (v28 << (((^(v31))) & 63))
    fmt.Println(v31)
    if v3 {
        var v47 string = "\n\t"
        switch v47 {
        case "hello":
            v48 := GenS6{F0: 571, F1: 24}
            _ = v48
        case "foo bar":
            for fi49 := 0; fi49 < 19; fi49++ {
                for fi50 := 0; fi50 < 15; fi50++ {
                    if v3 {
                        v46 += v28
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        v1 += 42
                        var v51 int32 = v28
                        _ = v51
                    }
                    v31 = v31
                    fmt.Println(v2)
                    _ = fi50
                }
                v3 = false
                v30 = append(v30, (^(fi49)))
                _ = fi49
            }
        case "":
            for fi52 := 0; fi52 < 11; fi52++ {
                switch (fi52 | -225) {
                case 128:
                    for fi53 := 0; fi53 < 14; fi53++ {
                        if (2 < 289) {
                            v31 += (v31 & 2)
                            v54 := GenS4{F0: v28, F1: -1}
                            fmt.Println(v2)
                            v55 := GenS2{F0: (18446744073709551615 == 936), F1: (^(v1))}
                            v31 = v31
                            _ = v54
                            _ = v55
                        } else {
                            var v56 int64 = (v1 - (v1 | v1))
                            _ = v56
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v1)
                        case "hello":
                            v46 = v28
                        }
                        _ = fi53
                    }
                default:
                    v57 := GenS5{F0: (v2 * (^(v2)))}
                    _ = v57
                }
                v47 = (v47 + v47)
                _ = fi52
            }
        default:
            v58 := GenS3{F0: 287, F1: ((1 <= 42) || false), F2: (v31 << ((586) & 63))}
            _ = v58
        }
        v1 = v1
        _ = v47
    }
    v28 += v46
    v59 := GenS7{F0: 868, F1: 42, F2: 92}
    if len(v30) > 0 {
        v30[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v46
    _ = v59
}
