package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint64; F1 uint }
func genFn0() uint8 {
    if (42 == 255) {
        v6 := []int{128, 0}
        _ = v6
    } else {
        switch "hello" {
        case "foo bar":
            v7 := []int{2147483647, -728, 32767, 9223372036854775807}
            _ = v7
        case "\n\t":
            var v8 string = ("foo bar" + ((("a" + "") + ("" + "foo bar")) + (("hello" + ("foo bar" + (("a" + ("" + "\n\t")) + "hello"))) + "a")))
            _ = v8
        case "hello":
            v9 := GenS1{F0: 235}
            _ = v9
        default:
            v10 := GenS1{F0: 179}
            _ = v10
        }
    }
    if (255 == 42) {
        switch ((("" + ("" + "\n\t")) + "a") + "") {
        case "":
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint32 = 4294967295
                var v13 int64 = 42
                v12 += v12
                _ = fi11
                _ = v12
                _ = v13
            }
        case "\n\t":
            v14 := []int{681, 0}
            _ = v14
        default:
            var v15 int = 128
            _ = v15
        }
        v16 := GenS1{F0: 2}
        _ = v16
    } else {
        v17 := GenS2{F0: -2147483648, F1: false, F2: 4294967295}
        _ = v17
    }
    v18 := []int{32768, 1, 42}
    var v19 int64 = -90
    _ = v18
    _ = v19
    return 93
}
func genFn1() uint8 {
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := []int{(fi20 ^ (fi21 % (((fi20 ^ (fi20 % ((fi21) | 1)))) | 1)))}
            var v23 float64 = -2.25
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    var v24 bool = (!(42 <= 128))
    fmt.Println(v24)
    var v25 string = "foo bar"
    fmt.Println(v24, v24)
    _ = v24
    _ = v25
    return 114
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{(("\n\t" + "a") + ("a" + ("a" + ""))), ("hello" + (("a" + "") + "hello"))}
    var v3 int64 = -1
    v4 := make(map[uint32]int)
    var v5 int8 = 127
    v3 -= v3
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int16 = int16(v3)
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 float64 = (1e10 - ((((3.14159 * 1.5) / 1.0) - (3.14159 * (1e10 - 0.0))) - (0.0 - (3.14159 + 0.0))))
            v29 -= v29
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    v30 := GenS1{F0: 148}
    switch 1 {
    case -858:
        v3 += 1
    case 726:
        var v31 bool = true
        _ = v31
    case -435:
        v2 = append(v2, "a")
    }
    v5 += v5
    v32 := GenS3{F0: 0}
    v3 = v3
    var v33 uint64 = 4294967295
    _ = len(v1)
    for fi34 := 0; fi34 < 14; fi34++ {
        v5 += v5
        _ = fi34
    }
    v5 = 127
    if ((true && (true && false)) && (0 == -497)) {
        for fi35 := 0; fi35 < 14; fi35++ {
            v3 = (v3 + (^(v3)))
            switch ("" + ("hello" + "foo bar")) {
            case "hello":
                var v36 uint = 65535
                _ = v36
            case "foo bar":
                if (42 == 478) {
                    for fi37 := 0; fi37 < 13; fi37++ {
                        var v38 int = fi35
                        var v39 uint16 = 32767
                        var v40 int8 = -33
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 4; fi41++ {
                        if (float64(1e10) == ((((((1e10 + (3.14159 / -2.25)) * 0.0) / 1.0) * 0.1) - (float64((float32(float64(float32(((1e10 - (1e10 + float64(0.1))) * 0.0)))) + 3.14159)) - 0.1)) + (0.0 / 1.0))) {
                            if (true || false) {
                                v2 = append(v2, "foo bar")
                            } else {
                                var v42 int8 = 27
                                _ = v42
                            }
                            v43 := GenS4{F0: 1, F1: 355}
                            v3 = (v3 * (v3 + v3))
                            _ = v43
                        }
                        v44 := GenS1{F0: 2}
                        _ = fi41
                        _ = v44
                    }
                    switch "\n\t" {
                    case "":
                        v45 := GenS5{F0: true, F1: 0}
                        _ = v45
                    default:
                        var v46 int64 = (v3 + (-(v3)))
                        _ = v46
                    }
                }
            case "a":
                switch (-(fi35)) {
                case -968:
                    v3 -= (^(v3))
                }
            }
            v47 := GenS6{F0: 32767, F1: -1}
            _ = fi35
            _ = v47
        }
        var v48 int64 = int64(v33)
        _ = v48
    } else {
        var v49 uint = 2
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        for fi51 := 0; fi51 < 10; fi51++ {
            v4[1] = (fi50 % ((901) | 1))
            switch (fi51 & fi51) {
            case 841:
                var v52 int8 = (v5 << ((4294967295) & 63))
                _ = v52
            case -715:
                v53 := GenS6{F0: 961, F1: 321}
                v5 = v5
                v5 -= 100
                _ = v53
            }
            switch "\n\t" {
            case "a":
                if false {
                    v54 := GenS3{F0: 286}
                    v55 := GenS7{F0: v33, F1: 65535}
                    _ = v54
                    _ = v55
                }
            case "hello":
                var v56 int8 = v5
                var v57 int = 9223372036854775807
                switch v57 {
                case 150:
                    if true {
                        v3 -= v3
                        v33 = 42
                    } else {
                        if true {
                            var v58 uint64 = 797
                            fmt.Println(fi51, v56, v3)
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            _ = v58
                        }
                        var v59 int64 = 120
                        _ = v59
                    }
                case 2147483647:
                    fmt.Println(v3, v5)
                }
                fmt.Println(v57, v56, v3)
                v5 += 1
                _ = v56
                _ = v57
            }
            _ = fi51
        }
        _ = fi50
    }
    v2 = append(v2, "hello")
    for fi60 := 0; fi60 < 10; fi60++ {
        for fi61 := 0; fi61 < 7; fi61++ {
            v62 := v4[2147483647]
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        var v64 uint64 = (^(v33))
        var v65 int32 = 0
        if (-1 >= -766) {
            fmt.Println(v3, v3)
            var v66 uint32 = 65535
            _ = v66
        }
        _ = fi63
        _ = v64
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
