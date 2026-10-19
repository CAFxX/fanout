package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 string }
func genFn0() bool {
    if false {
        v4 := GenS1{F0: 2, F1: 1}
        var v5 uint16 = 798
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint16 = v5
            for fi8 := 0; fi8 < 7; fi8++ {
                v7 -= v7
                v5 += (v5 >> ((1) & 63))
                v9 := GenS1{F0: 0, F1: 4294967295}
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    var v10 uint = 730
    v11 := GenS2{F0: 42, F1: 4294967295}
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int = fi12
        v13 -= fi12
        _ = fi12
        _ = v13
    }
    v10 = v10
    _ = v10
    _ = v11
    return ((0.0 * (0.0 / float32(3.14159))) <= (3.14159 / (1e10 / 1.0)))
}
func genFn1(p0 string, p1 int32, p2 int) int8 {
    var v14 int32 = (p1 & (p1 >> ((968) & 63)))
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"a", "\n\t", "a"}
    var v3 int16 = 32767
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int16 = 1
        switch fi15 {
        case 128:
            if (1 <= 508) {
                for fi17 := 0; fi17 < 15; fi17++ {
                    delete(v1, 713)
                    _ = fi17
                }
                v18 := GenS3{F0: 42, F1: 18446744073709551615}
                var v19 uint = 147
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v16 += -1
                    _ = fi20
                }
            }
        default:
            switch (-(fi15)) {
            case 2147483647:
                for fi21 := 0; fi21 < 13; fi21++ {
                    fmt.Println(v3, v3)
                    var v22 int16 = v3
                    _ = fi21
                    _ = v22
                }
            case -873:
                v16 = 0
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi15
        _ = v16
    }
    switch int(v3) {
    case 9223372036854775807:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3)
        default:
            v23 := GenS4{F0: 32767, F1: 2}
            _ = v23
        }
    case 499:
        v2 = append(v2, ("foo bar" + ("\n\t" + "a")))
    case -637:
        if true {
            v24 := GenS2{F0: v3, F1: 892}
            var v25 string = "foo bar"
            _ = v24
            _ = v25
        }
    }
    _ = len(v2)
    v3 = (v3 % ((v3) | 1))
    v26 := v1[4294967295]
    if ((v26 ^ (v26 % (((v26 - v26)) | 1))) <= v26) {
        var v27 bool = false
        _ = len(v2)
        var v28 uint = 571
        _ = v27
        _ = v28
    } else {
        v26 -= (^(v26))
        for fi29 := 0; fi29 < 12; fi29++ {
            if (!false) {
                var v30 int32 = 1
                if (!((fi29 - v26) <= v26)) {
                    var v31 int32 = 363
                    _ = v31
                }
                _ = v30
            } else {
                v3 -= v3
                var v32 uint8 = 7
                switch (v26 | fi29) {
                case 2147483647:
                    var v33 uint16 = 2
                    _ = v33
                }
                _ = v32
            }
            _ = fi29
        }
    }
    var v34 int64 = -9223372036854775808
    fmt.Println(v34, v34, v34)
    if (1 != 65535) {
        var v35 uint8 = 42
        _ = v35
    } else {
        fmt.Println(v3)
    }
    var v36 uint64 = 18446744073709551615
    v34 = (v34 / ((42) | 1))
    if len(v2) > 0 {
        v2[0] = ((("foo bar" + "") + "") + "")
    }
    if (1e10 > float64(3.14159)) {
        if false {
            fmt.Println(v3, v3, v26)
        } else {
            if false {
                v34 += (v34 * 1)
                v37 := GenS2{F0: -28, F1: v36}
                _ = v37
            } else {
                var v38 string = "foo bar"
                fmt.Println(v3, v26, v38)
                _ = v38
            }
            if false {
                v39 := GenS1{F0: v36, F1: v36}
                _ = v39
            } else {
                if (42 < 181) {
                    v34 = v34
                    v3 -= 0
                }
                var v40 int8 = genFn1(("\n\t" + "hello"), 0, 127)
                _ = v40
            }
        }
        if false {
            fmt.Println(v36, v3)
            if true {
                for fi41 := 0; fi41 < 8; fi41++ {
                    var v42 uint16 = 255
                    v43 := GenS5{F0: v36, F1: (("\n\t" + (("foo bar" + ("hello" + "")) + "hello")) + ("hello" + "a")), F2: v36}
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                if len(v2) > 0 {
                    v2[0] = ""
                }
            } else {
                for fi44 := 0; fi44 < 11; fi44++ {
                    v45 := GenS6{F0: 32767, F1: 4294967295}
                    _ = fi44
                    _ = v45
                }
                switch -456 {
                case 32767:
                    _ = len(v1)
                default:
                    for fi46 := 0; fi46 < 13; fi46++ {
                        var v47 float32 = ((1e10 - 0.1) + float32(((1e10 + 0.1) + -2.25)))
                        _ = fi46
                        _ = v47
                    }
                }
                v36 += (^(v36))
            }
            switch v26 {
            case 4294967295:
                var v48 uint32 = 1
                _ = v48
            case 9223372036854775807:
                for fi49 := 0; fi49 < 16; fi49++ {
                    var v50 int64 = (v34 & v34)
                    _ = fi49
                    _ = v50
                }
            case -752:
                v1[1] = v26
            }
        } else {
            v3 -= v3
            var v51 int16 = (v3 | (v3 & v3))
            switch (v26 - 215) {
            case -163:
                if (252 != 0) {
                    fmt.Println(v36, v26, v26)
                } else {
                    switch v26 {
                    case 128:
                        var v52 float32 = (((0.0 - 0.0) / 0.1) + (-2.25 + 0.1))
                        _ = v52
                    default:
                        v51 += v3
                    }
                    for fi53 := 0; fi53 < 5; fi53++ {
                        var v54 string = (("foo bar" + "a") + (("hello" + "") + "foo bar"))
                        _ = fi53
                        _ = v54
                    }
                }
            default:
                for fi55 := 0; fi55 < 19; fi55++ {
                    var v56 string = "foo bar"
                    v57 := GenS4{F0: (v26 >> ((475) & 63)), F1: 4294967295}
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            }
            _ = v51
        }
        for fi58 := 0; fi58 < 10; fi58++ {
            v59 := GenS5{F0: 1, F1: (("hello" + "a") + "a"), F2: v36}
            switch "" {
            case "a":
                if false {
                    v60 := GenS4{F0: 0, F1: 42}
                    v61 := GenS7{F0: 128, F1: 700}
                    _ = v60
                    _ = v61
                }
            }
            _ = fi58
            _ = v59
        }
    } else {
        fmt.Println(v34, v26, v34)
        var v62 int = (v26 % (((v26 << ((255) & 63))) | 1))
        if false {
            v63 := GenS8{F0: 749, F1: 1000}
            _ = v63
        } else {
            fmt.Println(v26, v34, v36)
        }
        _ = v62
    }
    v64 := GenS2{F0: (v3 + v3), F1: (v36 ^ (^(v36)))}
    v3 = (v3 | (v3 / ((950) | 1)))
    v26 = (v26 << ((255) & 63))
    v65 := GenS9{F0: "foo bar"}
    for fi66 := 0; fi66 < 20; fi66++ {
        var v67 string = "a"
        v68 := GenS7{F0: 237, F1: 32767}
        _ = fi66
        _ = v67
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v36
    _ = v64
    _ = v65
}
