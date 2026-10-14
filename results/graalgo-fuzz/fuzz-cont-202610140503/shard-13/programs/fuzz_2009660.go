package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int64; F1 bool; F2 int32 }
type GenS5 struct { F0 int64; F1 bool; F2 int64 }
func genFn0(p0 bool, p1 int64) uint64 {
    if (!(9223372036854775807 != 127)) {
        fmt.Println(p1, p0)
    } else {
        for fi4 := 0; fi4 < 7; fi4++ {
            for fi5 := 0; fi5 < 12; fi5++ {
                p1 = (p1 & -9223372036854775808)
                _ = fi5
            }
            _ = fi4
        }
    }
    v6 := GenS1{F0: -1}
    v7 := []int{}
    if ((42 == p1) || (p1 == (p1 ^ 0))) {
        p0 = p0
        v7 = append(v7, 4294967295)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = -2147483648
    v3 := []int{-1, 0, -1}
    v2 -= v2
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "foo bar" {
        case "a":
            var v9 uint32 = 2
            _ = v9
        }
        v10 := GenS2{F0: -859}
        v1 += v1
        _ = fi8
        _ = v10
    }
    switch ("\n\t" + (("foo bar" + "\n\t") + "foo bar")) {
    case "":
        v2 = (v2 & -1)
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        var v11 float64 = -2.25
        _ = v11
    }
    v12 := GenS1{F0: 32767}
    var v13 bool = false
    switch -1 {
    case 127:
        if (0 < 929) {
            v2 = 492
            v1 += -2.25
            v3 = append(v3, 624)
        }
    case 4294967295:
        switch (("foo bar" + "\n\t") + "") {
        case "\n\t":
            v1 = 0.1
        case "":
            fmt.Println(v13)
        }
    }
    var v14 uint = 861
    var v15 int8 = 42
    if true {
        v1 = (0.0 / 3.14159)
        if (!v13) {
            if v13 {
                if (!((float32(v1) * 0.1) == 1e10)) {
                    fmt.Println(v13)
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        for fi16 := 0; fi16 < 14; fi16++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v2 = (v2 / (((v2 ^ (v2 ^ (v2 / ((v2) | 1))))) | 1))
                            _ = fi16
                        }
                        var v17 int8 = v15
                        _ = v17
                    }
                }
                v14 = v14
            } else {
                var v18 uint16 = uint16(v14)
                fmt.Println(v2, v13)
                _ = v18
            }
            if true {
                fmt.Println(v1, v13, v14)
                switch 279 {
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 379:
                    fmt.Println(v2)
                default:
                    v19 := GenS1{F0: 1}
                    _ = v19
                }
                _ = len(v3)
            } else {
                v1 = 0.1
                if v13 {
                    v20 := GenS1{F0: -1}
                    _ = v20
                } else {
                    switch (("" + "\n\t") + "hello") {
                    case "hello":
                        if (414 < 42) {
                            var v21 uint8 = 127
                            if (!(!v13)) {
                                fmt.Println(v13, v15)
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    var v23 int = (fi22 * (fi22 | fi22))
                                    v2 -= (v2 - (-(v2)))
                                    _ = fi22
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 7; fi24++ {
                                    if v13 {
                                        var v25 int64 = 42
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: 9223372036854775807}
                                    _ = fi24
                                    _ = v26
                                }
                            }
                            _ = v21
                        }
                    case "\n\t":
                        for fi27 := 0; fi27 < 16; fi27++ {
                            v28 := GenS3{F0: true, F1: 132}
                            v14 -= 65535
                            switch 9223372036854775807 {
                            case -480:
                                v29 := GenS3{F0: v13, F1: 4294967295}
                                _ = v29
                            case -789:
                                var v30 int64 = -9223372036854775808
                                _ = v30
                            case -220:
                                fmt.Println(v13)
                            }
                            _ = fi27
                            _ = v28
                        }
                    default:
                        if false {
                            v14 = v14
                        }
                    }
                }
            }
        } else {
            switch 983 {
            case 2147483647:
                var v31 uint16 = 42
                _ = v31
            default:
                switch ((("a" + "") + "hello") + ("\n\t" + (("" + "a") + "a"))) {
                case "foo bar":
                    var v32 uint = v14
                    _ = v32
                case "hello":
                    v33 := GenS4{F0: -862, F1: ("" == ("\n\t" + ("a" + ("a" + "")))), F2: 93}
                    _ = v33
                case "\n\t":
                    var v34 float32 = float32((((float64((-2.25 * ((1.5 - (((0.0 + ((1.5 + 1.5) * 3.14159)) + 1.5) + float32(v1))) / ((-2.25 / 1e10) + 1e10)))) + (3.14159 + float64(1.5))) - v1) - v1))
                    _ = v34
                default:
                    if (!(2 == 54)) {
                        _ = len(v3)
                        v14 += (v14 % ((v14) | 1))
                        v35 := GenS4{F0: 9223372036854775807, F1: (!(true || v13)), F2: v2}
                        _ = v35
                    }
                }
            }
            switch 4294967295 {
            case -2147483648:
                v36 := GenS1{F0: 42}
                _ = v36
            case 5:
                var v37 string = "hello"
                _ = v37
            case 42:
                v13 = ((v2 * (^(v2))) == (v2 / (((v2 & (v2 * v2))) | 1)))
            default:
                if (v14 != (v14 ^ v14)) {
                    v38 := GenS5{F0: -9223372036854775808, F1: v13, F2: 9223372036854775807}
                    fmt.Println(v2, v13)
                    for fi39 := 0; fi39 < 8; fi39++ {
                        var v40 float32 = 0.0
                        _ = fi39
                        _ = v40
                    }
                    _ = v38
                } else {
                    v15 += (v15 >> ((v14) & 63))
                    v14 -= (v14 + (v14 * 255))
                    var v41 float32 = 3.14159
                    _ = v41
                }
            }
            if len(v3) > 0 {
                v3[0] = -488
            }
        }
    } else {
        v3 = append(v3, 32768)
        var v42 int64 = -30
        _ = v42
    }
    v2 = (v2 << (((v14 % ((v14) | 1))) & 63))
    switch ("foo bar" + "hello") {
    case "a":
        v3 = append(v3, 32768)
    case "foo bar":
        for fi43 := 0; fi43 < 17; fi43++ {
            switch "foo bar" {
            case "\n\t":
                var v44 int64 = -9223372036854775808
                _ = v44
            default:
                var v45 bool = (("a" == "a") && (!(v13 && v13)))
                _ = v45
            }
            v1 = (float64((0.0 + 0.0)) + (v1 + 1.5))
            _ = fi43
        }
    case "hello":
        for fi46 := 0; fi46 < 5; fi46++ {
            v2 -= 1
            v3 = append(v3, 2147483647)
            v13 = ((v15 * v15) >= (v15 % (((v15 % ((-1) | 1))) | 1)))
            _ = fi46
        }
    }
    fmt.Println(v1, v14, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
