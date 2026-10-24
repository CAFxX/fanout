package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 uint, p2 int) uint8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        p2 = p2
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(p2)
            v9 := []int{-92, (p2 + 130), (p2 - 32767)}
            _ = fi8
            _ = v9
        }
        p0 = (p0 + (p1 / (((p1 - (^(p0)))) | 1)))
        _ = fi7
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{p2}
        if true {
            var v12 float64 = (0.0 / 1e10)
            if len(v11) > 0 {
                v11[0] = 2147483647
            }
            v13 := GenS1{F0: false, F1: 102, F2: 42}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    p1 = 42
    switch (p2 % ((p2) | 1)) {
    case 127:
        for fi14 := 0; fi14 < 11; fi14++ {
            switch -710 {
            case -9223372036854775808:
                switch (p2 - 42) {
                case 9223372036854775807:
                    p0 = (p1 % (((p0 * 18446744073709551615)) | 1))
                case 0:
                    var v15 int8 = 120
                    _ = v15
                case 42:
                    fmt.Println(fi14, fi14, fi14)
                default:
                    var v16 int8 = -1
                    _ = v16
                }
            default:
                p2 = (fi14 ^ p2)
            }
            v17 := GenS1{F0: (246 > 127), F1: 42, F2: 165}
            _ = fi14
            _ = v17
        }
    default:
        p2 = (p2 << ((p1) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    return 139
}
func genFn1(p0 int, p1 uint32, p2 int8) int {
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            fmt.Println(p0, p0)
            var v20 uint16 = 65535
            _ = fi19
            _ = v20
        }
        var v21 uint = 208
        fmt.Println(fi18, p2, fi18)
        _ = fi18
        _ = v21
    }
    var v22 uint32 = 2
    var v23 uint64 = uint64(v22)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return (p0 * (p0 ^ (p0 << ((2) & 63))))
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 2
    var v3 uint16 = v2
    var v4 float64 = (1e10 / 1.0)
    v5 := make(map[string]string)
    var v6 uint16 = v3
    v24 := GenS2{F0: 196, F1: -628}
    v4 += v4
    v25 := GenS2{F0: genFn0(4294967295, 65535, genFn1(32767, 2147483647, (v1 - (v1 - (v1 / ((v1) | 1)))))), F1: -1}
    var v26 int8 = 52
    var v27 uint8 = genFn0(255, 0, 296)
    var v28 uint8 = (v27 + (^(v27)))
    switch -1 {
    case 32767:
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "":
            var v29 float32 = (float32(v4) / float32((v4 / (v4 * ((((3.14159 - v4) + v4) * 1e10) + v4)))))
            _ = v29
        case "\n\t":
            v30 := v5["\n\t"]
            _ = v30
        case "hello":
            v28 = (v28 | (v27 - v27))
        }
    case -9223372036854775808:
        for fi31 := 0; fi31 < 14; fi31++ {
            fmt.Println(v3)
            v6 -= v3
            if false {
                fmt.Println(v1)
                var v32 int = fi31
                _ = v32
            } else {
                if true {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        if false {
                            delete(v5, "")
                        } else {
                            switch (^(fi33)) {
                            case 32768:
                                v34 := GenS1{F0: (true && false), F1: 18446744073709551615, F2: (v6 % ((v3) | 1))}
                                _ = v34
                            default:
                                if (0 != 697) {
                                    v2 -= (v2 | (v6 << ((255) & 63)))
                                }
                            }
                            v35 := GenS2{F0: (v27 / (((v27 * v28)) | 1)), F1: 823}
                            _ = v35
                        }
                        switch (fi33 * fi33) {
                        case 9223372036854775807:
                            var v36 int32 = -2147483648
                            _ = v36
                        }
                        var v37 uint32 = 65535
                        _ = fi33
                        _ = v37
                    }
                    v3 += (^(v2))
                } else {
                    v38 := v5[("hello" + "hello")]
                    v3 = 1
                    _ = v38
                }
                var v39 uint16 = v6
                switch (fi31 + (fi31 | (fi31 ^ (fi31 - genFn1(genFn1(fi31, 823, 127), 418, int8(v39)))))) {
                case 105:
                    v40 := GenS3{F0: (!(18446744073709551615 < 967)), F1: 9223372036854775807}
                    _ = v40
                case 32767:
                    v27 += genFn0(660, 2, genFn1(fi31, 65535, v26))
                }
                _ = v39
            }
            _ = fi31
        }
    case -2147483648:
        switch int(v3) {
        case 127:
            v41 := []int{}
            _ = v41
        case 128:
            for fi42 := 0; fi42 < 18; fi42++ {
                for fi43 := 0; fi43 < 12; fi43++ {
                    if true {
                        v27 = genFn0(407, 4294967295, 127)
                        fmt.Println(v28, v27)
                        v1 = -83
                    }
                    _ = fi43
                }
                _ = fi42
            }
        case 366:
            for fi44 := 0; fi44 < 2; fi44++ {
                _ = len(v5)
                v45 := []int{4294967295, (fi44 & (-(fi44))), fi44}
                _ = fi44
                _ = v45
            }
        default:
            if ((v2 ^ v2) >= 2) {
                switch -2147483648 {
                case 2147483647:
                    var v46 uint32 = 1
                    _ = v46
                }
                v26 = v1
            }
        }
    }
    v47 := GenS1{F0: (v6 == (^(v3))), F1: 335, F2: v6}
    var v48 uint32 = 42
    v49 := v5["a"]
    v50 := GenS2{F0: genFn0(1, 841, 32767), F1: -1}
    if (false && false) {
        for fi51 := 0; fi51 < 6; fi51++ {
            if true {
                v48 = v48
                v52 := []int{fi51, fi51}
                _ = v52
            }
            for fi53 := 0; fi53 < 20; fi53++ {
                if (false || (42 != fi51)) {
                    fmt.Println(v2)
                    var v54 int64 = 0
                    _ = v54
                } else {
                    switch 127 {
                    case 42:
                        v55 := []int{(fi51 % (((fi53 & (-(fi53)))) | 1)), -2147483648, (fi53 % (((^(fi51))) | 1))}
                        _ = v55
                    default:
                        switch v49 {
                        case "":
                            switch "hello" {
                            case "\n\t":
                                v56 := GenS4{F0: fi51}
                                _ = v56
                            case "hello":
                                v4 = v4
                                switch v49 {
                                case "foo bar":
                                    if false {
                                        switch ("hello" + "foo bar") {
                                        case "\n\t":
                                            v57 := []int{(^(fi53)), genFn1((fi51 + fi51), (v48 - v48), 42), fi53, fi53}
                                            _ = v57
                                        default:
                                            v58 := []int{int(v48), (-(fi53))}
                                            _ = v58
                                        }
                                        if false {
                                            fmt.Println(v26, v4)
                                            v6 -= v6
                                            v28 -= 127
                                            var v59 float32 = float32(v4)
                                            v60 := GenS1{F0: false, F1: 0, F2: (v3 + v6)}
                                            _ = v59
                                            _ = v60
                                        }
                                    }
                                }
                                var v61 int32 = 0
                                _ = v61
                            case "foo bar":
                                fmt.Println(v48, fi53, v28)
                            default:
                                var v62 bool = false
                                _ = v62
                            }
                        case "\n\t":
                            var v63 uint8 = (v27 << ((1) & 63))
                            _ = v63
                        default:
                            v64 := GenS2{F0: (v27 / ((v27) | 1)), F1: -1}
                            _ = v64
                        }
                    }
                    v2 -= (v3 << ((4294967295) & 63))
                    v65 := GenS3{F0: true, F1: 765}
                    var v66 string = v49
                    _ = v65
                    _ = v66
                }
                _ = fi53
            }
            _ = fi51
        }
        v67 := []int{32768}
        var v68 int32 = 42
        _ = v67
        _ = v68
    } else {
        var v69 bool = true
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
