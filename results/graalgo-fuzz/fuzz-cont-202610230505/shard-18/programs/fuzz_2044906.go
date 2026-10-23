package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    v1 := []string{("\n\t" + "a")}
    var v2 int32 = 2147483647
    var v3 int8 = 127
    var v4 int64 = 1
    if false {
        var v5 uint16 = 111
        _ = v5
    } else {
        switch -977 {
        case 4294967295:
            v6 := GenS1{F0: ("hello" + ("" + "")), F1: (v2 ^ -1)}
            _ = v6
        case -444:
            for fi7 := 0; fi7 < 17; fi7++ {
                if (!(42 > 1)) {
                    fmt.Println(fi7, v3, v3)
                    v3 = v3
                    for fi8 := 0; fi8 < 4; fi8++ {
                        var v9 uint8 = 213
                        var v10 int32 = 1
                        var v11 float32 = 1.5
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                } else {
                    var v12 uint32 = 2147483647
                    _ = v12
                }
                _ = fi7
            }
        }
    }
    var v13 int = 632
    switch "" {
    case "":
        v14 := GenS2{F0: 1, F1: 0}
        _ = v14
    case "a":
        var v15 int32 = -1
        _ = v15
    case "foo bar":
        v16 := GenS1{F0: "foo bar", F1: (-(v2))}
        _ = v16
    default:
        var v17 int16 = -1
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: 255}
    if (1 >= 2) {
        for fi19 := 0; fi19 < 8; fi19++ {
            if false {
                v20 := GenS2{F0: 32767, F1: 0}
                _ = v20
            }
            _ = fi19
        }
    }
    _ = len(v1)
    v21 := GenS3{F0: 2147483648, F1: (!((0.0 * 1e10) <= 1.5))}
    v2 += (v2 - (v2 << ((0) & 63)))
    fmt.Println(v2, v13, v13)
    switch v13 {
    case 659:
        _ = len(v1)
    case 685:
        var v22 uint8 = 128
        _ = v22
    default:
        for fi23 := 0; fi23 < 12; fi23++ {
            if false {
                var v24 uint8 = 1
                v4 = 9223372036854775807
                _ = v24
            } else {
                v25 := GenS3{F0: 2147483647, F1: ((533 == 30) && false)}
                if (true || false) {
                    v3 -= v3
                    if (-232 >= -427) {
                        switch "a" {
                        case "":
                            v4 += v4
                        case "hello":
                            if false {
                                if (255 < 205) {
                                    v4 = (-(v4))
                                    var v26 int = (v13 & -9223372036854775808)
                                    for fi27 := 0; fi27 < 9; fi27++ {
                                        switch fi23 {
                                        case 32767:
                                            var v28 bool = true
                                            v29 := GenS2{F0: 1, F1: 0}
                                            _ = v28
                                            _ = v29
                                        case 32768:
                                            v3 += -116
                                        }
                                        _ = fi27
                                    }
                                    _ = v26
                                }
                            }
                        default:
                            v4 = v4
                        }
                    } else {
                        var v30 int16 = -20
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            var v31 float64 = ((-2.25 - ((float64(1.5) / ((3.14159 - float64(3.14159)) + 1.5)) - float64(float32((1.5 / 3.14159))))) + 3.14159)
                            _ = v31
                        case "a":
                            for fi32 := 0; fi32 < 18; fi32++ {
                                v2 += v2
                                _ = fi32
                            }
                        case "\n\t":
                            var v33 int16 = (v30 << ((4294967295) & 63))
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                if true {
                                    var v35 int32 = v2
                                    _ = len(v1)
                                    v2 += v35
                                    var v36 int64 = (v4 - -9223372036854775808)
                                    _ = v35
                                    _ = v36
                                } else {
                                    v3 = 0
                                    v30 = (v30 & v30)
                                }
                                _ = fi34
                            }
                        }
                        _ = v30
                    }
                    if true {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v1 = append(v1, "a")
                            v4 -= 0
                            _ = fi37
                        }
                        v4 += v4
                        var v38 int32 = v2
                        _ = v38
                    }
                } else {
                    for fi39 := 0; fi39 < 13; fi39++ {
                        switch 4294967295 {
                        case 4294967295:
                            _ = len(v1)
                        case 42:
                            v4 -= (v4 - -9223372036854775808)
                        }
                        v3 = int8(v4)
                        _ = len(v1)
                        _ = fi39
                    }
                    var v40 float64 = 1e10
                    _ = v40
                }
                var v41 string = ""
                _ = v25
                _ = v41
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 int64 = (v4 ^ (v4 ^ 9223372036854775807))
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                var v44 uint8 = 22
                _ = fi42
                _ = v43
                _ = v44
            }
            var v45 uint = 18446744073709551615
            _ = fi23
            _ = v45
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
