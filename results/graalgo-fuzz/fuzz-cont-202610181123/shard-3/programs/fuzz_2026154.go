package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }
type GenS7 struct { F0 uint; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int8; F2 string }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 int16; F1 bool }
func genFn0(p0 uint64) uint16 {
    p0 -= (p0 | p0)
    var v6 string = "\n\t"
    _ = p0
    _ = v6
    return 255
}

func main() {
    v1 := []string{"foo bar", "foo bar", "a", ((("\n\t" + ("hello" + "a")) + "\n\t") + "\n\t")}
    var v2 float32 = 0.0
    var v3 bool = (966 < -32768)
    v4 := []uint32{982, 4294967295, 65535}
    var v5 int64 = 1
    v7 := GenS1{F0: 65535, F1: 942, F2: 127}
    if (987 >= 387) {
        v8 := GenS2{F0: 1}
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5, v3)
        case "a":
            if (!(0 != v5)) {
                v9 := GenS3{F0: 32767}
                switch -2147483648 {
                case 32767:
                    v2 = ((float32(1.5) + v2) * v2)
                case -9223372036854775808:
                    v2 += (v2 + 1.5)
                }
                _ = v9
            }
        default:
            if (!v3) {
                v5 += 9223372036854775807
            } else {
                var v10 float32 = float32(float64((float32((1e10 + 0.0)) - float32(0.1))))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v11 int32 = 465
                _ = v10
                _ = v11
            }
        }
        _ = v8
    }
    if v3 {
        switch ("foo bar" + "") {
        case "":
            _ = len(v4)
        case "a":
            if (42 == -1) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v12 int64 = v5
                switch -172 {
                case -130:
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v12 = (v5 ^ -1)
                        v3 = ((1.5 / 1.0) >= float64((1e10 - v2)))
                        v14 := GenS4{F0: 65535, F1: ("a" + "hello")}
                        _ = fi13
                        _ = v14
                    }
                case 32767:
                    v15 := GenS1{F0: 0, F1: 365, F2: 130}
                    _ = v15
                default:
                    var v16 int8 = 1
                    _ = v16
                }
                _ = v12
            }
        case "hello":
            var v17 uint16 = 65535
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        if (42 >= 4294967295) {
            if (4294967295 <= 18446744073709551615) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int64 = (v5 % ((v5) | 1))
                    if (!v3) {
                        var v21 uint64 = 1
                        _ = v21
                    } else {
                        var v22 int16 = -1
                        _ = v22
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v23 uint16 = 32767
                    _ = fi19
                    _ = v20
                    _ = v23
                }
                v24 := GenS5{F0: 174, F1: true}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 string = ("\n\t" + "a")
                    var v27 uint16 = 1
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v28 := GenS6{F0: v5, F1: 65535, F2: -32768}
                _ = v28
            }
            var v29 int8 = 100
            _ = v29
        } else {
            if (2 <= 118) {
                var v30 uint8 = 255
                v2 += 1.5
                _ = v30
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                if ((true || false) || true) {
                    v5 += v5
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v32 uint = 2
                        _ = v32
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "foo bar":
                        fmt.Println(fi18)
                    }
                }
                v5 += -9223372036854775808
                var v33 int16 = 1000
                if (37 < 128) {
                    switch (-(fi18)) {
                    case -220:
                        if (1 == 1) {
                            var v34 uint64 = 42
                            _ = v34
                        }
                    case 2147483647:
                        v5 -= (^(v5))
                    case -919:
                        v35 := GenS7{F0: 18446744073709551615, F1: true, F2: 42}
                        _ = v35
                    }
                    fmt.Println(fi18, fi31)
                    v36 := GenS8{F0: 2, F1: 42, F2: "hello"}
                    _ = v36
                } else {
                    var v37 uint16 = genFn0(2)
                    _ = len(v4)
                    v1 = append(v1, "\n\t")
                    if v3 {
                        v33 += (v33 - 1)
                        fmt.Println(v5)
                        v4 = append(v4, 2147483647)
                    } else {
                        fmt.Println(v33)
                    }
                    var v38 int16 = (-(v33))
                    _ = v37
                    _ = v38
                }
                v33 += (v33 ^ v33)
                _ = fi31
                _ = v33
            }
            v2 -= 1e10
        }
        var v39 bool = (-1 <= -648)
        for fi40 := 0; fi40 < 18; fi40++ {
            if false {
                switch "hello" {
                case "foo bar":
                    if (0 <= 369) {
                        v5 += 1
                        var v41 int32 = 953
                        _ = v41
                    } else {
                        v2 = 0.1
                        v2 -= v2
                    }
                case "":
                    switch 2147483647 {
                    case 0:
                        var v42 uint8 = 176
                        fmt.Println(v42, v3)
                        if v39 {
                            if v39 {
                                v39 = (42 == 244)
                                v43 := GenS9{F0: v42, F1: (-(fi40))}
                                _ = v43
                            } else {
                                switch (fi18 & fi40) {
                                case 127:
                                    if ((3.14159 > (-2.25 - ((-2.25 * -2.25) / (1.5 - float64(v2))))) && ((v2 <= (v2 - v2)) || true)) {
                                        var v44 int = fi18
                                        _ = v44
                                    } else {
                                        v4 = append(v4, 1)
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v45 := GenS9{F0: (v42 - 48), F1: fi18}
                                    v5 -= v5
                                    v3 = false
                                    _ = v45
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v46 uint32 = 224
                            _ = v46
                        }
                        _ = v42
                    case 1:
                        v47 := GenS2{F0: fi18}
                        _ = v47
                    default:
                        v2 -= (float32(float64(1.5)) / 1e10)
                    }
                case "a":
                    switch "" {
                    case "foo bar":
                        var v48 int = 387
                        _ = v48
                    case "a":
                        var v49 float32 = 1.5
                        _ = v49
                    case "\n\t":
                        fmt.Println(v2, v5)
                    }
                }
            } else {
                v50 := GenS10{F0: 127, F1: 1000}
                v39 = (65535 > 976)
                v1 = append(v1, "hello")
                _ = v50
            }
            _ = fi40
        }
        _ = fi18
        _ = v39
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        fmt.Println(fi51, v3)
        _ = fi51
    }
    v4 = append(v4, 969)
    for fi52 := 0; fi52 < 5; fi52++ {
        if true {
            var v53 int16 = -1
            _ = v53
        }
        _ = fi52
    }
    if true {
        v2 = float32(-2.25)
    }
    var v54 int16 = 42
    v2 = v2
    var v55 int8 = 100
    v3 = (float32(((float64((1.5 / -2.25)) * 0.1) + float64(v2))) < v2)
    for fi56 := 0; fi56 < 19; fi56++ {
        switch ("a" + (((("foo bar" + "a") + ("" + "\n\t")) + "foo bar") + "\n\t")) {
        case "foo bar":
            switch ("a" + "foo bar") {
            case "":
                var v57 uint64 = 453
                _ = v57
            case "a":
                v54 = -981
            default:
                var v58 uint16 = 2
                _ = v58
            }
        case "a":
            var v59 float64 = ((0.1 * 1.5) / 1.5)
            _ = v59
        case "":
            fmt.Println(v3)
        }
        v60 := GenS11{F0: -32768, F1: v3}
        _ = fi56
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v7
}
