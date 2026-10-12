package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 int8; F1 int64 }
func genFn0(p0 int) bool {
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 += fi5
        _ = fi5
    }
    p0 = -2147483648
    p0 += p0
    _ = p0
    return true
}
func genFn1() uint64 {
    var v6 string = "foo bar"
    fmt.Println(v6)
    var v7 int64 = 23
    fmt.Println(v7)
    var v8 string = v6
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 42
    var v3 uint64 = 0
    v4 := []int{-9223372036854775808, 9223372036854775807, -1}
    if (true || (false || (false || (42 > -1)))) {
        v9 := GenS1{F0: -1, F1: 42}
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                if (!true) {
                    var v12 string = "hello"
                    _ = v12
                }
                var v13 bool = true
                v14 := GenS1{F0: 0, F1: 42}
                v1 -= v1
                var v15 bool = (!((1 == int8(fi11)) && (!(402 >= uint(v2)))))
                _ = fi11
                _ = v13
                _ = v14
                _ = v15
            }
            switch ("" + "foo bar") {
            case "\n\t":
                v16 := GenS1{F0: int16(fi10), F1: 18446744073709551615}
                _ = v16
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    v17 := GenS2{F0: 1, F1: true, F2: 0}
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 18; fi18++ {
                    if len(v4) > 0 {
                        v4[0] = (fi18 % ((fi18) | 1))
                    }
                    if false {
                        if (false && (!true)) {
                            var v19 int32 = -2147483648
                            if false {
                                var v20 float32 = float32(float64(float32(-2.25)))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch "" {
                                case "":
                                    v21 := GenS1{F0: 42, F1: 255}
                                    _ = v21
                                case "hello":
                                    var v22 int8 = -94
                                    _ = v22
                                case "foo bar":
                                    v1 -= (v1 | v1)
                                    v19 = v19
                                default:
                                    var v23 float32 = (v20 - v20)
                                    _ = v23
                                }
                                _ = v20
                            } else {
                                v2 = uint32(v3)
                                v2 -= (v2 % ((2) | 1))
                            }
                            _ = v19
                        } else {
                            var v24 uint8 = 92
                            var v25 int = fi18
                            _ = v24
                            _ = v25
                        }
                        v2 -= v2
                        var v26 float64 = 1.5
                        _ = v26
                    }
                    _ = fi18
                }
            }
            _ = fi10
        }
        var v27 int = 2147483647
        _ = v27
        _ = v9
    } else {
        switch int(v2) {
        case 9223372036854775807:
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
        case -2147483648:
            switch "foo bar" {
            case "foo bar":
                switch -652 {
                case -2147483648:
                    var v28 uint8 = (v1 + (^(v1)))
                    _ = v28
                case 4294967295:
                    v29 := GenS3{F0: v3, F1: (v1 << ((2) & 63))}
                    _ = v29
                case 128:
                    v3 += v3
                }
            case "hello":
                var v30 uint32 = (v2 / ((4294967295) | 1))
                _ = v30
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            switch -364 {
            case 925:
                v31 := GenS4{F0: 2}
                _ = v31
            case 32768:
                v32 := GenS4{F0: 18446744073709551615}
                _ = v32
            case 32767:
                switch ("" + "\n\t") {
                case "a":
                    fmt.Println(v1, v3)
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
        }
        switch 4294967295 {
        case -2147483648:
            v1 = v1
        case -1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v4 = append(v4, 127)
    var v33 uint64 = v3
    var v34 uint64 = genFn1()
    v2 = v2
    v35 := GenS5{F0: 416, F1: (v2 | (v2 * uint32(v33))), F2: v2}
    v36 := GenS1{F0: 1000, F1: 741}
    v37 := GenS6{F0: 697, F1: (^(v2))}
    if ((!false) && (-60 < 42)) {
        switch 0 {
        case -1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 42:
            if (("a" + "a") <= "") {
                switch 1 {
                case -2147483648:
                    v38 := GenS6{F0: 1000, F1: 967}
                    _ = v38
                case 133:
                    v3 -= v33
                default:
                    v39 := GenS7{F0: 2, F1: v1}
                    _ = v39
                }
            } else {
                v40 := GenS8{F0: 0, F1: -9223372036854775808}
                _ = v40
            }
        }
        fmt.Println(v1, v33, v33)
    }
    var v41 uint = 65535
    fmt.Println(v1)
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(fi42)
        var v43 float64 = (-2.25 / (float64(3.14159) * 0.1))
        _ = fi42
        _ = v43
    }
    switch ("hello" + "") {
    case "\n\t":
        var v44 float64 = (3.14159 * 0.1)
        _ = v44
    case "hello":
        if (false && (!(false || true))) {
            if (1 > (v33 % ((v33) | 1))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v45 int64 = 42
                _ = v45
            }
        } else {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            v34 = 1
        }
    case "":
        v4 = append(v4, -796)
    default:
        for fi46 := 0; fi46 < 6; fi46++ {
            v4 = append(v4, (fi46 | (fi46 >> (((v41 / (((v41 & (^(v41)))) | 1))) & 63))))
            v34 += (v34 * 848)
            switch fi46 {
            case 771:
                var v47 int16 = 42
                _ = v47
            case 114:
                fmt.Println(v1)
            case 32767:
                var v48 int16 = 1
                _ = v48
            }
            _ = fi46
        }
    }
    v3 -= (v33 << ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
