package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 int32, p1 bool, p2 int16) int8 {
    switch "a" {
    case "hello":
        var v5 int16 = (p2 + (p2 * p2))
        _ = v5
    case "a":
        switch 2147483647 {
        case 42:
            switch ("\n\t" + "a") {
            case "foo bar":
                var v6 int8 = 31
                _ = v6
            case "hello":
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 int = fi7
                    fmt.Println(p1, p0, p1)
                    _ = fi7
                    _ = v8
                }
            }
        case 127:
            var v9 int = 32767
            _ = v9
        }
    case "":
        p0 = p0
    default:
        var v10 bool = (p1 && (128 < 255))
        _ = v10
    }
    var v11 float32 = -2.25
    fmt.Println(p0, p0)
    p1 = ("a" != (("hello" + ("a" + "hello")) + (("" + "\n\t") + "\n\t")))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 string = (("" + (("\n\t" + "a") + "")) + ("foo bar" + "a"))
    var v2 float64 = (3.14159 + 1.5)
    v3 := []int{}
    var v4 uint = 584
    switch -121 {
    case -465:
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v4, v4, v4)
            var v13 uint64 = 1
            _ = fi12
            _ = v13
        }
    case 42:
        if (0 > 18446744073709551615) {
            _ = len(v3)
            if true {
                var v14 uint64 = 42
                if false {
                    var v15 int = -2147483648
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 uint64 = (v14 ^ (v14 << (((v4 / ((279) | 1))) & 63)))
                        v2 = 3.14159
                        if (-2147483648 >= -2147483648) {
                            for fi18 := 0; fi18 < 4; fi18++ {
                                if false {
                                    v17 = v17
                                    v19 := GenS1{F0: (42 >= -1), F1: 854, F2: 697}
                                    v20 := GenS2{F0: 127, F1: 65535, F2: (625 == 32767)}
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint32 = 2147483647
                                    _ = v21
                                }
                                _ = fi18
                            }
                            if false {
                                fmt.Println(v1, v2)
                            } else {
                                switch (fi16 % ((128) | 1)) {
                                case 32767:
                                    var v22 int64 = 42
                                    _ = v22
                                case -9223372036854775808:
                                    if len(v3) > 0 {
                                        v3[0] = (fi16 - 32768)
                                    }
                                case 2147483647:
                                    v17 = v17
                                default:
                                    var v23 float64 = v2
                                    _ = v23
                                }
                                if true {
                                    switch v1 {
                                    case "foo bar":
                                        var v24 uint64 = (v14 | v14)
                                        _ = v24
                                    default:
                                        for fi25 := 0; fi25 < 8; fi25++ {
                                            v1 = ("foo bar" + "\n\t")
                                            _ = fi25
                                        }
                                    }
                                    v4 -= 933
                                }
                            }
                            for fi26 := 0; fi26 < 3; fi26++ {
                                if (1 >= (v4 * 175)) {
                                    v27 := GenS3{F0: v4}
                                    v3 = append(v3, fi26)
                                    _ = v27
                                }
                                v4 += 4294967295
                                _ = fi26
                            }
                        } else {
                            var v28 uint = (^(v4))
                            if false {
                                var v29 string = ("" + v1)
                                v30 := GenS4{F0: (v1 + "foo bar"), F1: 747}
                                _ = v29
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 6; fi31++ {
                                fmt.Println(v1, v2, v14)
                                _ = fi31
                            }
                            _ = v28
                        }
                        _ = fi16
                        _ = v17
                    }
                    _ = len(v3)
                }
                var v32 int = -663
                _ = v14
                _ = v32
            } else {
                for fi33 := 0; fi33 < 18; fi33++ {
                    fmt.Println(v2, v2, v2)
                    _ = fi33
                }
                v34 := GenS1{F0: true, F1: -1, F2: (v4 | v4)}
                v3 = append(v3, 9223372036854775807)
                _ = v34
            }
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 uint8 = uint8(v4)
                for fi37 := 0; fi37 < 14; fi37++ {
                    if len(v3) > 0 {
                        v3[0] = fi37
                    }
                    _ = fi37
                }
                v4 += 0
                _ = fi35
                _ = v36
            }
        }
    }
    var v38 int64 = 9223372036854775807
    fmt.Println(v1, v38)
    if ((779 > 0) && true) {
        v39 := GenS5{F0: 42, F1: v4, F2: -83}
        var v40 uint32 = 4294967295
        _ = v39
        _ = v40
    } else {
        fmt.Println(v2, v1, v4)
    }
    if (!((v38 % ((v38) | 1)) <= (v38 << ((541) & 63)))) {
        for fi41 := 0; fi41 < 11; fi41++ {
            fmt.Println(fi41)
            _ = len(v3)
            _ = len(v3)
            _ = fi41
        }
        fmt.Println(v38)
        v42 := GenS3{F0: (v4 + (v4 % (((v4 % ((v4) | 1))) | 1)))}
        _ = v42
    }
    v1 = v1
    fmt.Println(v1)
    switch 4294967295 {
    case 128:
        if true {
            v43 := GenS2{F0: genFn0(1, (!(-110 > 0)), -32768), F1: 41, F2: true}
            v2 -= v2
            _ = v43
        }
    case 1:
        v38 += -997
    case 2147483647:
        v1 = (v1 + "\n\t")
    }
    if (1 != 1) {
        for fi44 := 0; fi44 < 8; fi44++ {
            if false {
                v45 := GenS6{F0: 1}
                var v46 uint = 1
                _ = v45
                _ = v46
            }
            var v47 uint32 = 2
            if len(v3) > 0 {
                v3[0] = fi44
            }
            _ = fi44
            _ = v47
        }
        v2 = v2
        for fi48 := 0; fi48 < 14; fi48++ {
            var v49 int32 = 42
            var v50 uint8 = 220
            v51 := GenS7{F0: 0, F1: 1}
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
        }
    }
    var v52 bool = (int16(v38) <= 1)
    v53 := GenS8{F0: 1, F1: "", F2: 0}
    var v54 int32 = 2147483647
    for fi55 := 0; fi55 < 15; fi55++ {
        for fi56 := 0; fi56 < 16; fi56++ {
            if len(v3) > 0 {
                v3[0] = fi56
            }
            _ = fi56
        }
        _ = fi55
    }
    switch (v1 + (v1 + ("" + v1))) {
    case "hello":
        var v57 float64 = (v2 * float64((0.0 * -2.25)))
        _ = v57
    case "foo bar":
        v3 = append(v3, -952)
    case "a":
        v2 = (v2 * float64(-2.25))
    default:
        for fi58 := 0; fi58 < 1; fi58++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v59 float64 = ((v2 - (v2 + 1.5)) - (0.1 * (v2 + ((float64(3.14159) + v2) + (v2 * (-2.25 / 1.0))))))
            _ = fi58
            _ = v59
        }
    }
    v1 = ((v1 + v1) + ((v1 + v1) + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
