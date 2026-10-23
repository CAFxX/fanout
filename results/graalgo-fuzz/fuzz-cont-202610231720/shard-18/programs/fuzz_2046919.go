package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    var v1 int = -969
    var v2 bool = (422 <= (v1 | (v1 % (((v1 | -260)) | 1))))
    v3 := []uint32{42}
    var v4 uint64 = 0
    var v5 int8 = -1
    v6 := make(map[uint32]int)
    if false {
        v4 -= v4
        v3 = append(v3, 38)
        v2 = true
    } else {
        var v7 int16 = -32768
        v1 = (v1 & (v1 % (((v1 % ((-1) | 1))) | 1)))
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                v4 += (v4 << ((130) & 63))
                if (!((-2.25 / (float64((3.14159 * (0.1 + 0.1))) / (float64(0.0) + -2.25))) >= 0.0)) {
                    v10 := GenS1{F0: 962, F1: 941, F2: "a"}
                    if false {
                        v11 := GenS1{F0: 794, F1: 18446744073709551615, F2: "a"}
                        if v2 {
                            v2 = v2
                        }
                        var v12 float64 = float64(-2.25)
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS1{F0: 65535, F1: 18446744073709551615, F2: ("hello" + "a")}
                    _ = v10
                    _ = v13
                }
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    }
    if v2 {
        var v14 float32 = 0.1
        var v15 float64 = 0.1
        _ = v14
        _ = v15
    } else {
        v3 = append(v3, 703)
    }
    var v16 uint = 255
    v17 := v6[2147483647]
    v1 -= (v1 / (((v17 - 0)) | 1))
    v5 = (v5 / ((v5) | 1))
    switch ("" + "hello") {
    case "":
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 string = ""
                switch ("a" + v20) {
                case "a":
                    var v21 uint64 = (v4 << ((570) & 63))
                    if v2 {
                        var v22 string = v20
                        var v23 int64 = 9223372036854775807
                        fmt.Println(v5, v4, v20)
                        _ = v22
                        _ = v23
                    }
                    if (v20 > v20) {
                        fmt.Println(v1)
                        v24 := GenS1{F0: uint16(v16), F1: 2, F2: v20}
                        v25 := GenS1{F0: 1, F1: v16, F2: ("a" + v20)}
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                default:
                    v17 = (fi19 * v1)
                }
                _ = fi19
                _ = v20
            }
            v1 = v17
            _ = fi18
        }
    case "\n\t":
        for fi26 := 0; fi26 < 14; fi26++ {
            v2 = (v2 || v2)
            for fi27 := 0; fi27 < 3; fi27++ {
                v16 = v16
                v28 := GenS2{F0: 100, F1: ("hello" + ("foo bar" + (("a" + "") + "a")))}
                _ = fi27
                _ = v28
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi26
        }
    default:
        for fi29 := 0; fi29 < 14; fi29++ {
            switch ("" + "") {
            case "":
                v30 := GenS1{F0: 32767, F1: (v16 * v16), F2: ("a" + "\n\t")}
                _ = v30
            case "\n\t":
                for fi31 := 0; fi31 < 14; fi31++ {
                    if v2 {
                        if true {
                            var v32 uint32 = 2147483647
                            _ = v32
                        }
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        v2 = false
                    }
                    _ = fi31
                }
            default:
                if ((!(!(!v2))) || (!v2)) {
                    delete(v6, 2)
                    var v33 float32 = 0.0
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 4; fi34++ {
                v5 -= (v5 >> (((^(v16))) & 63))
                fmt.Println(v5)
                _ = fi34
            }
            v35 := GenS3{F0: fi29, F1: 1, F2: (-(v5))}
            _ = fi29
            _ = v35
        }
    }
    switch "" {
    case "foo bar":
        for fi36 := 0; fi36 < 15; fi36++ {
            v5 -= v5
            _ = fi36
        }
    }
    if false {
        var v37 int64 = 0
        switch (v1 * 127) {
        case -170:
            var v38 float64 = 1e10
            _ = v38
        default:
            var v39 float64 = ((-2.25 * (((1e10 / 1e10) / 1.5) - (0.1 + 0.1))) * (0.1 * -2.25))
            _ = v39
        }
        _ = v37
    } else {
        fmt.Println(v5)
        v40 := GenS4{F0: 239, F1: (4294967295 != 1)}
        v17 -= -336
        _ = v40
    }
    _ = len(v3)
    v41 := GenS3{F0: (v17 >> (((v16 >> (((v16 + 0)) & 63))) & 63)), F1: 697, F2: (v5 - (v5 & v5))}
    var v42 float64 = float64((0.0 * 3.14159))
    if (65535 <= 65535) {
        v6[2147483648] = -468
        var v43 int = int(v16)
        fmt.Println(v42)
        _ = v43
    } else {
        var v44 int8 = (v5 + -67)
        _ = v44
    }
    v45 := GenS3{F0: v17, F1: v4, F2: int8(v16)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
