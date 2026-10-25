package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int64 = -1
    var v2 int32 = 1
    var v3 int = -9223372036854775808
    var v4 uint16 = 435
    var v5 uint8 = 42
    v6 := []int{(v3 << ((65535) & 63)), (-(v3)), 17}
    v7 := GenS1{F0: (v1 * -9223372036854775808), F1: -81, F2: -208}
    if true {
        switch 9223372036854775807 {
        case 975:
            if false {
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 uint16 = 65535
                    _ = fi8
                    _ = v9
                }
                v10 := GenS1{F0: 1, F1: 100, F2: (v3 >> ((65535) & 63))}
                if false {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := GenS2{F0: 4294967295}
                        _ = fi11
                        _ = v12
                    }
                    if ((3.14159 < float64(0.0)) && true) {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            for fi14 := 0; fi14 < 17; fi14++ {
                                var v15 float32 = -2.25
                                _ = fi14
                                _ = v15
                            }
                            _ = len(v6)
                            v6 = append(v6, (-(v3)))
                            _ = fi13
                        }
                        var v16 uint16 = 42
                        _ = v16
                    } else {
                        v2 = (^(v2))
                    }
                    for fi17 := 0; fi17 < 9; fi17++ {
                        fmt.Println(v3, v1)
                        if true {
                            if true {
                                var v18 uint = 388
                                _ = v18
                            }
                        } else {
                            if false {
                                fmt.Println(fi17, v3)
                            } else {
                                v3 = (v3 % ((9223372036854775807) | 1))
                                switch v3 {
                                case -1:
                                    if false {
                                        var v19 uint8 = (v5 + (v5 % ((v5) | 1)))
                                        if false {
                                            v4 = v4
                                            var v20 bool = (false || ((!true) && false))
                                            if len(v6) > 0 {
                                                v6[0] = (-(v3))
                                            }
                                            _ = v20
                                        } else {
                                            v21 := GenS2{F0: 2}
                                            _ = len(v6)
                                            _ = len(v6)
                                            _ = v21
                                        }
                                        _ = v19
                                    }
                                default:
                                    if (42 == 4294967295) {
                                        v4 = v4
                                        fmt.Println(v2, v5)
                                        var v22 uint64 = 844
                                        _ = v22
                                    }
                                }
                                for fi23 := 0; fi23 < 16; fi23++ {
                                    _ = len(v6)
                                    v24 := GenS2{F0: 42}
                                    v5 += 25
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        }
                        fmt.Println(fi17)
                        _ = fi17
                    }
                } else {
                    v6 = append(v6, (^(v3)))
                    if false {
                        if false {
                            if false {
                                v6 = append(v6, (^(v3)))
                            }
                            fmt.Println(v5, v4, v5)
                            switch (^(v3)) {
                            case 0:
                                var v25 uint16 = (v4 & v4)
                                _ = v25
                            case -2147483648:
                                switch ((("" + "\n\t") + ("foo bar" + "hello")) + ("\n\t" + ((("hello" + "foo bar") + ((("foo bar" + "") + "foo bar") + "")) + "\n\t"))) {
                                case "":
                                    if (v1 > (v1 / ((v1) | 1))) {
                                        var v26 uint16 = v4
                                        var v27 string = (("foo bar" + "a") + (("hello" + ("a" + "hello")) + "a"))
                                        _ = v26
                                        _ = v27
                                    }
                                case "foo bar":
                                    v6 = append(v6, 429)
                                }
                            }
                        }
                    } else {
                        v3 += (v3 | -942)
                    }
                    v2 -= 2147483647
                }
                _ = v10
            }
        case -511:
            v4 = uint16(v5)
        case 9223372036854775807:
            if false {
                if (false || false) {
                    switch (v3 & -9223372036854775808) {
                    case 2147483647:
                        v2 = (^(v2))
                    case 32767:
                        v28 := GenS1{F0: v1, F1: -117, F2: 2147483647}
                        _ = v28
                    case 42:
                        for fi29 := 0; fi29 < 5; fi29++ {
                            switch "hello" {
                            case "":
                                fmt.Println(v5, v4, v1)
                            }
                            var v30 uint64 = 1
                            fmt.Println(v2, v4)
                            _ = fi29
                            _ = v30
                        }
                    }
                } else {
                    fmt.Println(v1, v4)
                }
                var v31 uint8 = (v5 & uint8(v4))
                var v32 int16 = 42
                _ = v31
                _ = v32
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v33 := GenS2{F0: 2}
                _ = v33
            }
        }
        if true {
            if len(v6) > 0 {
                v6[0] = v3
            }
            var v34 int = (v3 + int(v1))
            _ = v34
        }
    }
    switch (v3 & (v3 / ((v3) | 1))) {
    case 32767:
        switch "" {
        case "a":
            if false {
                v5 = v5
                v35 := GenS1{F0: v1, F1: 100, F2: int(v5)}
                _ = v35
            }
        case "hello":
            for fi36 := 0; fi36 < 8; fi36++ {
                fmt.Println(v2)
                _ = fi36
            }
        case "\n\t":
            v3 -= 42
        default:
            v6 = append(v6, 32768)
        }
    case 0:
        v37 := GenS1{F0: v1, F1: -128, F2: (v3 - (v3 & 9223372036854775807))}
        _ = v37
    default:
        var v38 int16 = -1
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int8 = 1
        var v41 int64 = 0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        if true {
            var v43 int = fi42
            v44 := GenS3{F0: 65535, F1: 222}
            var v45 int16 = 0
            _ = v43
            _ = v44
            _ = v45
        } else {
            v46 := GenS3{F0: 2147483647, F1: 32767}
            fmt.Println(v4, fi42, v4)
            _ = len(v6)
            _ = v46
        }
        _ = fi42
    }
    v47 := GenS4{F0: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
