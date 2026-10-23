package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint32; F1 int16; F2 int32 }
type GenS7 struct { F0 int; F1 uint16; F2 int }
type GenS8 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 int16) uint16 {
    var v6 int64 = 0
    var v7 float64 = 1.5
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{2147483648}
    var v3 float32 = (1e10 - float32(-2.25))
    var v4 int = -835
    var v5 uint8 = 128
    v8 := GenS1{F0: 18446744073709551615}
    v9 := GenS1{F0: 18446744073709551615}
    v5 -= 127
    v10 := GenS2{F0: "\n\t", F1: 0}
    v4 -= (v4 | v4)
    for fi11 := 0; fi11 < 15; fi11++ {
        if ((11 > 20) && true) {
            fmt.Println(v5, v3)
            if true {
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                if ((!(0.1 > 0.1)) && false) {
                    v12 := GenS3{F0: v5}
                    var v13 int64 = -1
                    _ = v12
                    _ = v13
                } else {
                    v4 -= fi11
                    for fi14 := 0; fi14 < 5; fi14++ {
                        if true {
                            if false {
                                var v15 int64 = -9223372036854775808
                                _ = v15
                            }
                        } else {
                            var v16 float64 = 3.14159
                            _ = v16
                        }
                        v17 := GenS4{F0: -1, F1: -102, F2: (v5 >> ((218) & 63))}
                        _ = fi14
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if false {
                            var v19 string = "foo bar"
                            _ = v19
                        }
                        fmt.Println(fi11, v4)
                        var v20 float64 = -2.25
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi18
                        _ = v20
                    }
                }
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                v3 -= v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "hello" {
                case "\n\t":
                    if true {
                        var v22 int = (fi11 / ((v4) | 1))
                        _ = len(v2)
                        _ = v22
                    }
                case "foo bar":
                    switch "a" {
                    case "a":
                        switch "" {
                        case "\n\t":
                            fmt.Println(v4, v3)
                            v23 := GenS5{F0: 703, F1: -14}
                            _ = v23
                        case "hello":
                            fmt.Println(fi21, fi11, fi21)
                        default:
                            var v24 uint8 = (v5 << ((42) & 63))
                            _ = v24
                        }
                    case "hello":
                        var v25 int32 = -646
                        _ = v25
                    }
                case "hello":
                    var v26 int16 = 42
                    _ = v26
                default:
                    if (-112 <= -123) {
                        v2 = append(v2, 65535)
                        v27 := GenS6{F0: 2147483647, F1: 1000, F2: -751}
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((((3.14159 * float32(((-2.25 + 0.0) - 0.0))) / v1) * v3) <= ((-2.25 / ((1e10 + (((float32(3.14159) / v3) / 1.5) * v1)) / 1e10)) / 3.14159)) {
                                fmt.Println(v4)
                                var v28 int64 = -1
                                _ = v28
                            }
                            var v29 float32 = (0.0 / 3.14159)
                            _ = v29
                        } else {
                            switch (fi21 % ((fi11) | 1)) {
                            case 0:
                                v30 := GenS5{F0: 319, F1: 1}
                                _ = v30
                            default:
                                fmt.Println(fi11, v3, v4)
                            }
                            v31 := GenS7{F0: fi11, F1: 2, F2: 42}
                            _ = v31
                        }
                        _ = v27
                    }
                }
                _ = fi21
            }
        }
        _ = fi11
    }
    v4 -= 9223372036854775807
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 string = ("\n\t" + "")
        var v34 uint32 = 2147483648
        _ = fi32
        _ = v33
        _ = v34
    }
    if ((!true) && ((42 > 18446744073709551615) || (!(!(-670 > -808))))) {
        if (!(!(!(18446744073709551615 < 4294967295)))) {
            var v35 uint64 = 4294967295
            for fi36 := 0; fi36 < 17; fi36++ {
                if true {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    var v37 uint64 = (v35 << ((0) & 63))
                    if (true && true) {
                        var v38 bool = false
                        _ = v38
                    } else {
                        v39 := GenS8{F0: (609 > 2147483648), F1: v5, F2: 40}
                        _ = v39
                    }
                    _ = v37
                }
                _ = fi36
            }
            _ = v35
        }
        fmt.Println(v4)
    } else {
        v40 := GenS5{F0: 0, F1: 1}
        for fi41 := 0; fi41 < 3; fi41++ {
            var v42 int32 = -2147483648
            var v43 float64 = 0.1
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = v40
    }
    var v44 int16 = 42
    _ = len(v2)
    var v45 float32 = v1
    v45 += v1
    var v46 string = ((("hello" + "a") + "a") + ("hello" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v8
    _ = v9
}
