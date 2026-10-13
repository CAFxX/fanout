package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint16; F1 int; F2 string }
type GenS6 struct { F0 int32; F1 string }
func genFn0() uint16 {
    v5 := GenS1{F0: 127, F1: 128, F2: -1}
    var v6 int = 276
    v7 := GenS1{F0: v6, F1: 142, F2: (-(v6))}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(v6 * fi8), (fi8 >> ((609) & 63)), (v6 << ((18446744073709551615) & 63)), (-(v6))}
        v6 = v6
        switch "a" {
        case "\n\t":
            var v10 int16 = -1
            _ = v10
        default:
            var v11 uint32 = 42
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 17
}

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{65535}
    v3 := []string{("" + "a")}
    var v4 float32 = 1e10
    var v12 int64 = -736
    var v13 uint = 461
    fmt.Println(v1)
    v14 := GenS2{F0: 1}
    for fi15 := 0; fi15 < 12; fi15++ {
        v3 = append(v3, "")
        _ = fi15
    }
    switch "a" {
    case "\n\t":
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v1, v12, v12)
            v16 := GenS3{F0: 2147483648, F1: 494}
            _ = v16
        } else {
            var v17 string = "a"
            _ = v17
        }
    case "foo bar":
        var v18 int = 32768
        _ = v18
    default:
        if false {
            if (((v12 * -535) != (v12 % ((v12) | 1))) || true) {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if (-2.25 > v4) {
                        var v20 uint = v13
                        if true {
                            if (!false) {
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    var v22 uint16 = genFn0()
                                    _ = fi21
                                    _ = v22
                                }
                                var v23 uint8 = 42
                                _ = v23
                            } else {
                                if true {
                                    _ = len(v2)
                                    fmt.Println(v13, v1, v13)
                                    if false {
                                        switch "a" {
                                        case "a":
                                            var v24 bool = (-936 != 0)
                                            _ = v24
                                        }
                                        if false {
                                            if (0 == 2147483647) {
                                                var v25 int16 = 1000
                                                _ = v25
                                            } else {
                                                var v26 int64 = (v12 & (v12 | (v12 - v12)))
                                                if true {
                                                    v26 = v12
                                                }
                                                fmt.Println(v1)
                                                _ = v26
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v4, v4)
                                }
                                for fi27 := 0; fi27 < 16; fi27++ {
                                    if (true || false) {
                                        var v28 int32 = 335
                                        v29 := GenS1{F0: fi27, F1: 1, F2: (fi27 | fi27)}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        if false {
                                            var v30 uint16 = 32767
                                            var v31 int = (fi19 * (fi27 & (fi19 | fi19)))
                                            switch (^(fi27)) {
                                            case 128:
                                                if false {
                                                    var v32 uint8 = 99
                                                    var v33 float32 = 3.14159
                                                    _ = v32
                                                    _ = v33
                                                }
                                            case 9223372036854775807:
                                                if true {
                                                    v31 = fi27
                                                }
                                            }
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = len(v2)
                                        v12 -= v12
                                    }
                                    v34 := GenS2{F0: uint16(v13)}
                                    _ = fi27
                                    _ = v34
                                }
                                if true {
                                    v20 = (v20 * (v20 ^ 18446744073709551615))
                                    v20 = 65535
                                } else {
                                    var v35 uint8 = 230
                                    v13 += v13
                                    v20 += 208
                                    _ = v35
                                }
                            }
                            v36 := GenS4{F0: -168, F1: 32767}
                            if false {
                                fmt.Println(v20, v4)
                                if true {
                                    _ = len(v2)
                                    var v37 int16 = 42
                                    v38 := GenS2{F0: 32767}
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 uint = 258
                                    _ = v39
                                }
                                v40 := GenS5{F0: 42, F1: fi19, F2: ("a" + (("\n\t" + "a") + ("hello" + "hello")))}
                                _ = v40
                            }
                            _ = v36
                        } else {
                            for fi41 := 0; fi41 < 12; fi41++ {
                                var v42 uint8 = 78
                                var v43 int32 = 2147483647
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                        }
                        _ = v20
                    }
                    if (v4 != (3.14159 + (v4 - v4))) {
                        v1 = (1.5 / -2.25)
                        v44 := GenS5{F0: 42, F1: fi19, F2: "\n\t"}
                        v1 = v1
                        _ = v44
                    } else {
                        if false {
                            fmt.Println(v13, v12)
                            v3 = append(v3, "foo bar")
                        }
                        _ = len(v2)
                        var v45 bool = true
                        _ = v45
                    }
                    _ = fi19
                }
                if (4294967295 > 4294967295) {
                    v4 = v4
                } else {
                    if (false && false) {
                        v13 += 502
                    } else {
                        switch 32768 {
                        case 9223372036854775807:
                            v13 = v13
                        case 32768:
                            var v46 string = "hello"
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 3; fi47++ {
                            for fi48 := 0; fi48 < 9; fi48++ {
                                var v49 uint16 = 2
                                var v50 int16 = 1000
                                var v51 int64 = v12
                                _ = fi48
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            v52 := GenS1{F0: (fi47 * (fi47 / (((fi47 >> ((v13) & 63))) | 1))), F1: 152, F2: (fi47 / ((fi47) | 1))}
                            _ = fi47
                            _ = v52
                        }
                        if len(v3) > 0 {
                            v3[0] = (("hello" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar")
                        }
                    }
                }
                switch ("hello" + "a") {
                case "a":
                    fmt.Println(v4, v4, v13)
                default:
                    v2 = append(v2, 4294967295)
                }
            }
            v53 := GenS6{F0: -2147483648, F1: ("a" + "a")}
            _ = v53
        }
    }
    var v54 int16 = 1000
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v13 += (v13 | 2)
    switch "\n\t" {
    case "\n\t":
        var v55 int8 = 123
        _ = v55
    }
    if false {
        switch ("hello" + ("" + "hello")) {
        case "foo bar":
            _ = len(v3)
        }
    }
    var v56 bool = ((((("\n\t" + "\n\t") > (((("" + "a") + "\n\t") + "foo bar") + "hello")) || ((^(v13)) == (v13 / ((1) | 1)))) || (-2.25 != v1)) && false)
    if (906 > 2147483648) {
        for fi57 := 0; fi57 < 19; fi57++ {
            var v58 float32 = (1e10 * (((v4 + v4) - float32(v1)) / v4))
            fmt.Println(v12, v58)
            _ = fi57
            _ = v58
        }
        v59 := GenS3{F0: 4294967295, F1: 1}
        var v60 float32 = (float32(-2.25) * (v4 * (v4 - 0.1)))
        _ = v59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v56
}
