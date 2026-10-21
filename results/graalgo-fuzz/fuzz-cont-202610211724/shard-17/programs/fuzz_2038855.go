package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 uint; F2 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS10 struct { F0 uint64 }
func genFn0() int8 {
    var v4 string = "foo bar"
    switch v4 {
    case "\n\t":
        if true {
            v5 := []int{9223372036854775807, 128, -484, -682}
            var v6 int64 = -94
            _ = v5
            _ = v6
        }
    case "a":
        v7 := []int{9223372036854775807}
        _ = v7
    default:
        fmt.Println(v4)
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint32 = 2147483648
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{(^(fi10)), (fi10 << ((65535) & 63)), fi8, fi10}
            v12 := GenS1{F0: -2147483648, F1: 539}
            v11 = append(v11, 32768)
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: v9, F1: v4}
        _ = fi8
        _ = v13
        _ = v9
    }
    if true {
        v4 = v4
    } else {
        var v14 int16 = -861
        var v15 int8 = -61
        _ = v14
        _ = v15
    }
    _ = v4
    return 105
}

func main() {
    var v1 int8 = 1
    var v2 int64 = 9223372036854775807
    var v3 uint64 = 42
    v16 := GenS3{F0: 42, F1: -2147483648}
    for fi17 := 0; fi17 < 1; fi17++ {
        v1 = v1
        v2 = v2
        _ = fi17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        fmt.Println(v1, v1, v1)
        for fi19 := 0; fi19 < 8; fi19++ {
            if ((v3 - (v3 - v3)) > v3) {
                var v20 bool = (4294967295 <= 2147483647)
                var v21 int16 = -844
                var v22 float64 = -2.25
                _ = v20
                _ = v21
                _ = v22
            } else {
                if (!false) {
                    v23 := []int{(fi19 & -251), (fi18 + (fi19 ^ (fi18 << ((255) & 63))))}
                    v24 := GenS2{F0: 2147483647, F1: ("" + "\n\t")}
                    if true {
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                    } else {
                        var v25 int32 = 1
                        var v26 int16 = 548
                        _ = v25
                        _ = v26
                    }
                    var v27 uint8 = 0
                    _ = v23
                    _ = v24
                    _ = v27
                }
                switch (fi18 / ((42) | 1)) {
                case 2147483647:
                    fmt.Println(fi18)
                case 0:
                    v1 -= (v1 + genFn0())
                }
                fmt.Println(fi19, v1)
            }
            if (-1 <= -568) {
                v28 := []int{fi19}
                _ = v28
            } else {
                v29 := GenS3{F0: v3, F1: -2147483648}
                _ = v29
            }
            if true {
                switch 32768 {
                case 0:
                    v30 := GenS4{F0: 256, F1: 42, F2: 42}
                    _ = v30
                }
                v31 := []int{968, fi19, 42}
                if (1.5 == -2.25) {
                    if true {
                        fmt.Println(fi19, v2)
                        v31 = append(v31, (-(fi19)))
                    } else {
                        var v32 uint64 = (v3 ^ v3)
                        fmt.Println(fi18, fi19, v3)
                        var v33 uint8 = 75
                        if true {
                            if (73 != 864) {
                                v33 += (v33 - v33)
                                switch ("\n\t" + ("foo bar" + "\n\t")) {
                                case "a":
                                    switch ("" + "hello") {
                                    case "hello":
                                        if (false || (!(0.0 < ((float64(0.1) * 1e10) + 1.5)))) {
                                            if (1 > 0) {
                                                v31 = append(v31, fi19)
                                                if len(v31) >= 2 {
                                                    v31 = v31[1:len(v31)]
                                                }
                                            }
                                        } else {
                                            v1 = (-(v1))
                                            v34 := GenS5{F0: v1}
                                            _ = v34
                                        }
                                    case "\n\t":
                                        fmt.Println(v2, fi18, v1)
                                    case "foo bar":
                                        v1 = (v1 >> ((42) & 63))
                                        v33 += v33
                                    default:
                                        if (("hello" + "a") > ("hello" + "hello")) {
                                            v1 -= (v1 ^ (v1 + v1))
                                            v3 -= v32
                                        }
                                    }
                                case "foo bar":
                                    var v35 int = -213
                                    _ = v35
                                default:
                                    if (70 < (^(v33))) {
                                        if false {
                                            if (true || false) {
                                                if len(v31) > 0 {
                                                    v31[0] = fi18
                                                }
                                            }
                                            if (false && ((v33 - (v33 * (v33 ^ v33))) == (v33 >> ((42) & 63)))) {
                                                v33 += 85
                                                v33 = v33
                                                var v36 int8 = (-(v1))
                                                _ = v36
                                            }
                                            v1 -= v1
                                            var v37 int16 = 0
                                            _ = v37
                                        }
                                        v3 -= (^(v32))
                                    }
                                }
                                _ = len(v31)
                                var v38 float32 = (float32(1e10) - (3.14159 / 3.14159))
                                var v39 int64 = v2
                                _ = v38
                                _ = v39
                            } else {
                                v33 -= 255
                                v2 += v2
                            }
                            switch 0 {
                            case 692:
                                if false {
                                    fmt.Println(v33, v3, v32)
                                    v3 -= (v3 - (v32 >> ((18446744073709551615) & 63)))
                                }
                            case -233:
                                var v40 uint32 = uint32(v32)
                                _ = v40
                            case -181:
                                v3 += 1
                            }
                            var v41 int64 = v2
                            _ = v41
                        }
                        _ = v32
                        _ = v33
                    }
                    if true {
                        v42 := GenS6{F0: uint(v3), F1: 1, F2: 42}
                        var v43 uint8 = 31
                        _ = v42
                        _ = v43
                    }
                } else {
                    if len(v31) > 0 {
                        v31[0] = fi19
                    }
                    var v44 uint64 = (^(v3))
                    v45 := GenS7{F0: (("\n\t" + (("\n\t" + "") + "hello")) + "foo bar")}
                    _ = v44
                    _ = v45
                }
                _ = v31
            } else {
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi18)
                case "\n\t":
                    var v46 uint = 526
                    _ = v46
                default:
                    if false {
                        if true {
                            v47 := GenS8{F0: ((255 == 251) || ((false && (!(!(1000 < 42)))) || false))}
                            _ = v47
                        }
                        if (!(65535 == 2)) {
                            v48 := []int{-411}
                            v49 := GenS9{F0: (v3 * v3), F1: 9, F2: -1}
                            v50 := GenS3{F0: (v3 / ((v3) | 1)), F1: 174}
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    }
                }
                v2 += v2
                var v51 uint64 = (v3 ^ v3)
                switch fi18 {
                case 662:
                    v51 = (v51 - (v3 % ((v3) | 1)))
                }
                _ = v51
            }
            _ = fi19
        }
        _ = fi18
    }
    v52 := GenS10{F0: 2}
    fmt.Println(v3)
    if true {
        v53 := GenS5{F0: v1}
        if (!true) {
            if true {
                for fi54 := 0; fi54 < 3; fi54++ {
                    v2 = (v2 << ((33) & 63))
                    _ = fi54
                }
            }
            if false {
                fmt.Println(v3, v2)
                v55 := []int{-1, 128, 128}
                if len(v55) > 0 {
                    v55[0] = 644
                }
                _ = v55
            } else {
                v2 += 9223372036854775807
            }
        }
        v1 = (v1 / (((^(v1))) | 1))
        _ = v53
    } else {
        var v56 uint16 = 838
        var v57 uint8 = 128
        fmt.Println(v57)
        _ = v56
        _ = v57
    }
    v2 -= (v2 | (v2 << ((65535) & 63)))
    if (274 <= 2) {
        var v58 int32 = -776
        _ = v58
    }
    if ((v1 % ((v1) | 1)) <= 81) {
        v59 := []int{4294967295}
        _ = v59
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v52
}
