package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int; F1 uint; F2 string }
type GenS4 struct { F0 string; F1 int32; F2 uint32 }
func genFn0(p0 int, p1 int64, p2 bool) int {
    p1 -= p1
    if ("\n\t" <= ("foo bar" + "")) {
        for fi5 := 0; fi5 < 7; fi5++ {
            switch p0 {
            case -827:
                var v6 int8 = 0
                _ = v6
            default:
                switch "\n\t" {
                case "foo bar":
                    var v7 uint32 = 65535
                    _ = v7
                case "":
                    fmt.Println(fi5)
                default:
                    p1 = (p1 << ((0) & 63))
                }
            }
            _ = fi5
        }
        v8 := GenS1{F0: 2147483648, F1: -210, F2: 127}
        var v9 int8 = 0
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 | p0)
}
func genFn1(p0 uint8) uint8 {
    if true {
        switch 127 {
        case -786:
            p0 = 0
        default:
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{genFn0((fi10 + 127), -9223372036854775808, ((((fi10 ^ genFn0(fi10, -9223372036854775808, false)) <= (-(fi10))) || false) && (127 < 79))), (fi10 + fi10), fi10}
                _ = fi10
                _ = v11
            }
        }
        var v12 bool = false
        v13 := GenS2{F0: 32767, F1: v12}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: 32768, F1: (true || (float64((3.14159 - 1e10)) == (0.1 * (0.1 * -2.25))))}
        _ = fi14
        _ = v15
    }
    v16 := []int{genFn0(898, 0, true), 4294967295, 0}
    if false {
        var v17 uint = uint(p0)
        switch 32768 {
        case 128:
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(p0, v17, v17)
                if (true || (275 == 1)) {
                    v19 := GenS3{F0: 128, F1: 400, F2: "foo bar"}
                    v16 = append(v16, genFn0((^(fi18)), 637, (3.14159 < (0.0 - -2.25))))
                    _ = len(v16)
                    _ = v19
                }
                switch (fi18 / (((fi18 % ((593) | 1))) | 1)) {
                case 9223372036854775807:
                    var v20 int16 = 0
                    _ = v20
                case 516:
                    if (true && true) {
                        if true {
                            var v21 uint32 = 42
                            v17 += v17
                            v22 := GenS2{F0: (fi18 >> (((v17 % ((v17) | 1))) & 63)), F1: (-1 > 1)}
                            _ = v21
                            _ = v22
                        }
                        var v23 int64 = 0
                        _ = v23
                    }
                default:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
                _ = fi18
            }
        case 32768:
            var v24 string = (("\n\t" + ("hello" + "")) + "foo bar")
            _ = v24
        }
        _ = v17
    } else {
        if false {
            var v25 int32 = -688
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 bool = (42 <= 628)
                fmt.Println(v27)
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    }
    _ = p0
    _ = v16
    return 127
}

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    var v3 int16 = 1000
    v4 := []uint32{514}
    v3 = 1000
    var v28 int8 = 42
    switch v2 {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        var v29 string = v2
        _ = v29
    }
    if false {
        v3 += -32768
        var v30 int32 = 2147483647
        v30 = v30
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint16 = 65535
            if true {
                fmt.Println(v3, fi32)
            } else {
                var v34 uint = 255
                fmt.Println(v33, v28, v34)
                switch v2 {
                case "\n\t":
                    v35 := GenS3{F0: fi31, F1: (v34 | (^(v34))), F2: ((v2 + "\n\t") + "hello")}
                    _ = v35
                case "hello":
                    v36 := GenS3{F0: genFn0(9223372036854775807, 284, (0.1 <= 1.5)), F1: v34, F2: (v2 + "foo bar")}
                    _ = v36
                case "a":
                    fmt.Println(fi32, v28, v34)
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v34
            }
            fmt.Println(fi31, v3, v3)
            var v37 string = v2
            _ = fi32
            _ = v33
            _ = v37
        }
        _ = fi31
    }
    var v38 int8 = v28
    var v39 int8 = v28
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v39 = v28
    var v40 string = ("a" + "")
    for fi41 := 0; fi41 < 13; fi41++ {
        fmt.Println(v2, v38, v28)
        _ = fi41
    }
    for fi42 := 0; fi42 < 4; fi42++ {
        var v43 float64 = 0.1
        for fi44 := 0; fi44 < 9; fi44++ {
            var v45 uint = 427
            var v46 uint8 = genFn1(2)
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if false {
        v40 = "foo bar"
        switch 0 {
        case -46:
            v47 := GenS4{F0: v2, F1: -1, F2: 19}
            _ = v47
        case 646:
            for fi48 := 0; fi48 < 3; fi48++ {
                var v49 uint16 = 180
                _ = fi48
                _ = v49
            }
        case -726:
            var v50 int64 = -1
            _ = v50
        default:
            var v51 uint16 = 65535
            _ = v51
        }
    }
    switch "foo bar" {
    case "hello":
        v28 -= v39
    case "\n\t":
        var v52 int = 860
        _ = v52
    case "a":
        for fi53 := 0; fi53 < 16; fi53++ {
            var v54 int16 = int16(fi53)
            v4 = append(v4, 106)
            for fi55 := 0; fi55 < 3; fi55++ {
                if (v3 <= (v3 >> ((18446744073709551615) & 63))) {
                    switch "" {
                    case "\n\t":
                        var v56 uint64 = 42
                        _ = v56
                    case "":
                        fmt.Println(fi53)
                    case "foo bar":
                        var v57 uint8 = 91
                        _ = v57
                    default:
                        v40 = v2
                    }
                    if (!(20 != -2147483648)) {
                        v40 = "foo bar"
                        v38 = (v39 / (((^(v38))) | 1))
                    }
                    var v58 uint64 = 0
                    _ = v58
                }
                _ = fi55
            }
            _ = fi53
            _ = v54
        }
    default:
        if len(v4) > 0 {
            v4[0] = 2
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
