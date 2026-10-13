package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint16) int32 {
    fmt.Println(p0)
    var v4 uint16 = (p0 & (p0 % (((^(p0))) | 1)))
    fmt.Println(p0, v4)
    var v5 bool = true
    v4 = p0
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint32 = 949
    var v2 float64 = -2.25
    v3 := []string{"foo bar"}
    v6 := GenS1{F0: 608, F1: 127}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 float32 = (3.14159 / 0.1)
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 244, F1: 2}
    fmt.Println(v1)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 808, F1: 127}
            var v12 uint16 = 32767
            v3 = append(v3, ("foo bar" + ("\n\t" + "\n\t")))
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        if (!false) {
            v3 = append(v3, "a")
            v2 = (v2 * (float64(-2.25) / v2))
            v1 -= (v1 * v1)
        } else {
            var v13 int32 = 2147483647
            if (-1 != int64(v13)) {
                var v14 uint16 = 0
                _ = v14
            } else {
                v1 = v1
                if true {
                    switch 822 {
                    case 128:
                        var v15 int8 = -128
                        _ = v15
                    }
                }
            }
            _ = v13
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            v1 = 2147483647
            _ = fi16
        }
        switch ("hello" + "foo bar") {
        case "foo bar":
            fmt.Println(v1, v1)
        case "":
            if (-375 >= -9223372036854775808) {
                if (127 >= 79) {
                    v1 += (^(v1))
                }
                var v17 int8 = 42
                _ = v17
            } else {
                var v18 uint16 = 255
                v19 := GenS2{F0: -2147483648, F1: 0}
                v2 = (0.1 * 1e10)
                _ = v18
                _ = v19
            }
        default:
            var v20 int32 = genFn0(2)
            _ = v20
        }
    }
    switch "a" {
    case "":
        var v21 uint16 = 65535
        _ = v21
    case "foo bar":
        var v22 int = 4294967295
        _ = v22
    case "a":
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 float32 = 0.0
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := GenS1{F0: 1, F1: 127}
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                v27 := GenS3{F0: ""}
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    default:
        if (-128 <= -128) {
            _ = len(v3)
            v1 -= v1
        } else {
            var v28 int8 = 127
            if true {
                v28 += v28
            } else {
                var v29 int = 9223372036854775807
                if ((0.0 < -2.25) || (v2 == (1e10 * 3.14159))) {
                    switch (^(v29)) {
                    case 2147483647:
                        if (2 < 180) {
                            v30 := GenS1{F0: 1, F1: -90}
                            if false {
                                var v31 float64 = (v2 + v2)
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 2; fi32++ {
                                    var v33 string = ("foo bar" + "a")
                                    if false {
                                        var v34 string = v33
                                        _ = v34
                                    } else {
                                        switch 0 {
                                        case 128:
                                            var v35 float64 = v2
                                            _ = v35
                                        case 475:
                                            fmt.Println(fi32)
                                        }
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                var v36 uint16 = 734
                                for fi37 := 0; fi37 < 16; fi37++ {
                                    v38 := GenS1{F0: 4294967295, F1: 127}
                                    for fi39 := 0; fi39 < 15; fi39++ {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v40 float64 = v2
                                            _ = v40
                                        case "foo bar":
                                            if (fi39 != fi39) {
                                                v41 := GenS2{F0: -129, F1: 179}
                                                switch "a" {
                                                case "":
                                                    fmt.Println(fi37, v29)
                                                case "hello":
                                                    var v42 int8 = v28
                                                    _ = v42
                                                case "\n\t":
                                                    fmt.Println(fi37, v2)
                                                }
                                                _ = v41
                                            } else {
                                                v28 += (v28 >> ((4294967295) & 63))
                                                v28 += (v28 - v28)
                                                v36 += (v36 / ((v36) | 1))
                                                fmt.Println(v2, v2)
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            v3 = append(v3, "foo bar")
                                        }
                                        _ = fi39
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                _ = v36
                            }
                            _ = v30
                        } else {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v28, v29)
                            case "a":
                                _ = len(v3)
                            default:
                                for fi43 := 0; fi43 < 2; fi43++ {
                                    var v44 float32 = 1e10
                                    _ = fi43
                                    _ = v44
                                }
                            }
                            fmt.Println(v28)
                        }
                    }
                    var v45 string = "\n\t"
                    fmt.Println(v2, v45)
                    _ = v45
                } else {
                    v1 -= 4294967295
                    v2 = (float64(3.14159) + 1.5)
                    if ((702 != 42) || ((true || (true && true)) || (1 != 42))) {
                        v46 := GenS3{F0: ""}
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 12; fi47++ {
                            var v48 uint = 2
                            v1 += v1
                            v2 = (0.1 * v2)
                            _ = fi47
                            _ = v48
                        }
                        v1 = (v1 << ((65535) & 63))
                    }
                }
                var v49 int64 = 42
                _ = v29
                _ = v49
            }
            _ = v28
        }
    }
    _ = len(v3)
    var v50 float64 = v2
    var v51 int16 = 1
    _ = len(v3)
    v3 = append(v3, "foo bar")
    fmt.Println(v50)
    for fi52 := 0; fi52 < 17; fi52++ {
        v53 := GenS4{F0: fi52}
        for fi54 := 0; fi54 < 6; fi54++ {
            var v55 uint8 = 138
            v1 = (v1 & v1)
            v56 := GenS3{F0: "\n\t"}
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi52
        _ = v53
    }
    var v57 uint16 = 17
    fmt.Println(v1, v51)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v57
    _ = v6
    _ = v9
}
