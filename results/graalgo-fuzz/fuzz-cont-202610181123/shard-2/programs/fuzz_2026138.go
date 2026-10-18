package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint }
type GenS6 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 int64; F2 string }
func genFn0(p0 string) string {
    v4 := GenS1{F0: p0}
    if false {
        v5 := []int{127, -803, 773}
        if (!true) {
            if false {
                var v6 uint8 = 0
                fmt.Println(v6)
                v6 = (v6 % ((v6) | 1))
                _ = v6
            } else {
                var v7 uint64 = 42
                fmt.Println(p0, v7)
                _ = v7
            }
        }
        _ = v5
    }
    _ = p0
    _ = v4
    return (p0 + "foo bar")
}
func genFn1() uint {
    switch "\n\t" {
    case "":
        var v8 int16 = -1
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 float32 = -2.25
            _ = fi9
            _ = v10
        }
    default:
        var v11 uint = 612
        _ = v11
    }
    var v12 int8 = 42
    if true {
        v12 = (v12 % (((v12 ^ v12)) | 1))
        v13 := []int{}
        _ = v13
    } else {
        if (!((v12 >> ((536) & 63)) != (v12 * v12))) {
            for fi14 := 0; fi14 < 4; fi14++ {
                var v15 float64 = (-2.25 - ((((0.1 * -2.25) - 1.5) - float64(3.14159)) * 1e10))
                _ = fi14
                _ = v15
            }
        } else {
            if (((3.14159 * 3.14159) > 3.14159) || ((65535 != 32767) && (-975 >= -1))) {
                v12 += v12
                v16 := []int{}
                _ = v16
            } else {
                v12 = 127
                fmt.Println(v12)
                var v17 int16 = 32767
                _ = v17
            }
        }
    }
    _ = v12
    return 1
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{530, 2147483647}
    v3 := []int{32767, -484, -9223372036854775808, 32767}
    if ("" == "foo bar") {
        v18 := GenS1{F0: ("foo bar" + ("\n\t" + "hello"))}
        if false {
            v1 = v1
            var v19 float64 = -2.25
            _ = v19
        }
        _ = v18
    } else {
        v20 := GenS1{F0: "\n\t"}
        switch 4294967295 {
        case -1:
            v1 -= (^(v1))
        case -707:
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 float32 = 0.1
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := GenS1{F0: "\n\t"}
                    v25 := GenS2{F0: (false && false)}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        default:
            fmt.Println(v1)
        }
        _ = v20
    }
    v1 -= v1
    v1 = v1
    v26 := GenS3{F0: "hello", F1: -9223372036854775808, F2: 42}
    fmt.Println(v1)
    var v27 float64 = 0.1
    v27 = v27
    if false {
        if true {
            switch 127 {
            case -2147483648:
                var v28 int16 = 658
                _ = v28
            case 32768:
                if (!(-590 == 9223372036854775807)) {
                    v27 -= v27
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if true {
                            v1 = (v1 / ((v1) | 1))
                        }
                        _ = fi29
                    }
                    var v30 uint8 = 92
                    _ = v30
                } else {
                    if true {
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v32 := GenS4{F0: 255}
                            v33 := GenS3{F0: "a", F1: -9223372036854775808, F2: 2147483647}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    } else {
                        var v34 uint16 = 710
                        var v35 uint64 = 656
                        _ = v34
                        _ = v35
                    }
                    var v36 uint64 = 757
                    for fi37 := 0; fi37 < 7; fi37++ {
                        var v38 int32 = -566
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                }
            case 774:
                if (127 <= 127) {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        var v40 int64 = 0
                        v41 := GenS2{F0: (65535 <= 648)}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    if len(v3) > 0 {
                        v3[0] = -181
                    }
                    switch "\n\t" {
                    case "\n\t":
                        switch -9223372036854775808 {
                        case 741:
                            v27 = (v27 + ((0.1 - 0.0) / (float64(3.14159) - v27)))
                        case 851:
                            v1 = 255
                        case 128:
                            var v42 bool = false
                            _ = v42
                        default:
                            var v43 int = -1
                            _ = v43
                        }
                    case "hello":
                        var v44 uint = genFn1()
                        _ = v44
                    case "a":
                        v27 -= 0.0
                    }
                }
            default:
                var v45 float64 = 1e10
                _ = v45
            }
            var v46 uint8 = 0
            var v47 uint8 = v46
            _ = v46
            _ = v47
        } else {
            var v48 bool = (!((true || ((907 == 4294967295) || true)) && (42 == 42)))
            v49 := GenS2{F0: false}
            var v50 int8 = 127
            _ = v48
            _ = v49
            _ = v50
        }
        var v51 int16 = 42
        v52 := GenS5{F0: 128, F1: v51, F2: 2}
        _ = v51
        _ = v52
    }
    v53 := GenS6{F0: v1, F1: -19, F2: v1}
    if ((139 != 225) && false) {
        var v54 int64 = 42
        for fi55 := 0; fi55 < 2; fi55++ {
            v56 := GenS7{F0: 47}
            _ = fi55
            _ = v56
        }
        fmt.Println(v1, v54)
        _ = v54
    }
    v1 += (v1 << ((700) & 63))
    if (127 > 32) {
        var v57 uint32 = 2147483648
        var v58 int32 = -994
        v59 := GenS8{F0: 0, F1: -777, F2: (("a" + "hello") + ("hello" + "hello"))}
        _ = v57
        _ = v58
        _ = v59
    }
    var v60 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v53
    _ = v60
}
