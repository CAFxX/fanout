package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int64) int {
    var v6 float64 = (3.14159 * 3.14159)
    fmt.Println(v6, v6)
    fmt.Println(p0, p0)
    v7 := GenS1{F0: 42}
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int64 = 605
    v2 := []uint32{}
    var v3 int8 = -1
    var v4 int64 = -1
    var v5 float32 = ((3.14159 + 1e10) / 1.5)
    fmt.Println(v5)
    v2 = append(v2, 65535)
    if (-9223372036854775808 >= 1) {
        v8 := GenS1{F0: -1}
        _ = v8
    } else {
        v9 := GenS1{F0: 42}
        _ = v9
    }
    v10 := GenS1{F0: 0}
    switch 127 {
    case 127:
        v11 := GenS2{F0: false}
        _ = v11
    case 32768:
        if true {
            var v12 float64 = 1.5
            switch ("" + ((("hello" + "foo bar") + ("foo bar" + ("a" + "\n\t"))) + "a")) {
            case "hello":
                v13 := GenS3{F0: "\n\t", F1: 42}
                _ = v13
            case "":
                if false {
                    var v14 uint32 = 2147483648
                    _ = v14
                }
            default:
                v2 = append(v2, 1)
            }
            _ = v12
        }
    }
    var v15 uint8 = 127
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int8 = 0
        _ = len(v2)
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 float32 = float32(0.1)
            if ((!true) || (!(1.5 == ((((((v19 / v5) / (float32(0.0) + (v19 / float32(((3.14159 + float64(1e10)) * 1.5))))) + v5) / v5) + 0.0) - v19)))) {
                if len(v2) > 0 {
                    v2[0] = 726
                }
            }
            v20 := GenS4{F0: 32767, F1: 42}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    if true {
        v21 := GenS4{F0: -662, F1: 42}
        v22 := GenS1{F0: -977}
        _ = v21
        _ = v22
    } else {
        if false {
            for fi23 := 0; fi23 < 18; fi23++ {
                if (4294967295 != 2) {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    switch ("\n\t" + ("foo bar" + ((("a" + "foo bar") + "") + "foo bar"))) {
                    case "a":
                        v3 = v3
                    case "\n\t":
                        v1 += (v4 << ((2) & 63))
                    case "hello":
                        var v24 uint8 = v15
                        _ = v24
                    }
                } else {
                    if false {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v5 -= v5
                            var v26 uint16 = 42
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        switch ("a" + "foo bar") {
                        case "a":
                            v2 = append(v2, 42)
                        }
                        v27 := GenS5{F0: 415}
                        _ = v27
                    }
                    if false {
                        switch 127 {
                        case 127:
                            switch 1 {
                            case 755:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -67:
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            default:
                                fmt.Println(v5, fi23)
                            }
                        case 32767:
                            v15 = (v15 % ((v15) | 1))
                        }
                    }
                    if (float32(-2.25) == ((v5 * 0.1) * (-2.25 + v5))) {
                        if false {
                            v28 := GenS3{F0: "\n\t", F1: 2}
                            _ = v28
                        }
                        v5 = float32(-2.25)
                    } else {
                        fmt.Println(v3, v5, v4)
                        for fi29 := 0; fi29 < 13; fi29++ {
                            var v30 uint32 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi29
                            _ = v30
                        }
                        if (0 == 32767) {
                            var v31 string = "hello"
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((v31 + "foo bar") < (v31 + (v31 + ((v31 + ((v31 + "") + ((v31 + v31) + (v31 + v31)))) + (v31 + v31))))) {
                                var v32 int8 = -5
                                _ = v32
                            }
                            _ = v31
                        }
                    }
                }
                if (2 != 4294967295) {
                    v33 := GenS6{F0: 18446744073709551615, F1: 42, F2: 42}
                    fmt.Println(v15)
                    _ = v33
                }
                _ = fi23
            }
            var v34 int32 = 851
            switch ("foo bar" + ("" + "")) {
            case "foo bar":
                if true {
                    v3 += 22
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v34 = (v34 & (-(v34)))
                        var v36 bool = true
                        _ = fi35
                        _ = v36
                    }
                }
            case "a":
                v37 := GenS4{F0: -548, F1: 4294967295}
                _ = v37
            }
            _ = v34
        }
        var v38 uint8 = (v15 * v15)
        _ = v38
    }
    switch "foo bar" {
    case "":
        var v39 float64 = -2.25
        _ = v39
    }
    if false {
        switch 0 {
        case 2147483647:
            v2 = append(v2, 1)
        case -377:
            v40 := GenS7{F0: ((true || false) && false)}
            _ = v40
        default:
            v41 := GenS8{F0: 351, F1: 2147483648}
            _ = v41
        }
    }
    fmt.Println(v15, v15)
    fmt.Println(v5)
    fmt.Println(v5)
    switch -2147483648 {
    case 2147483647:
        switch int(v3) {
        case -370:
            if true {
                var v42 int32 = -696
                if true {
                    _ = len(v2)
                }
                _ = v42
            }
        default:
            switch 32767 {
            case 421:
                if (!(1e10 > -2.25)) {
                    if len(v2) > 0 {
                        v2[0] = 433
                    }
                } else {
                    v43 := GenS5{F0: 2147483648}
                    if (false || (127 == (v3 << ((1) & 63)))) {
                        var v44 float64 = 1e10
                        _ = v44
                    } else {
                        v45 := GenS2{F0: (1 != 255)}
                        _ = v45
                    }
                    _ = len(v2)
                    _ = v43
                }
            }
        }
    case -2147483648:
        v5 -= ((1e10 / v5) + (v5 - 3.14159))
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v47 int64 = v1
        switch "" {
        case "":
            v1 = (-(v1))
        case "a":
            var v48 int32 = 0
            _ = v48
        case "\n\t":
            v4 = (v4 << ((862) & 63))
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
