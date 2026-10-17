package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int64; F1 int64; F2 int32 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
func genFn0() uint64 {
    switch ("\n\t" + ("\n\t" + ("foo bar" + (("foo bar" + "foo bar") + ("a" + "a"))))) {
    case "\n\t":
        var v7 string = "\n\t"
        _ = v7
    default:
        switch "a" {
        case "a":
            v8 := []int{128, -75, 42}
            _ = v8
        default:
            var v9 int16 = 1
            _ = v9
        }
    }
    var v10 int16 = -1
    v11 := GenS1{F0: 946}
    var v12 bool = (false || ((9223372036854775807 > 849) || (!false)))
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 string = ("a" + "\n\t")
    v2 := []uint32{504, 660}
    var v3 uint = 720
    var v4 uint64 = 828
    var v5 uint64 = v4
    v6 := []int{4294967295, 42, 128}
    var v13 uint64 = v4
    if len(v2) > 0 {
        v2[0] = 230
    }
    var v14 uint8 = 127
    var v15 uint8 = v14
    switch -828 {
    case -603:
        switch 9223372036854775807 {
        case -31:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if (!true) {
                var v16 uint16 = 1
                _ = v16
            } else {
                var v17 int32 = 1
                if (2 != 700) {
                    var v18 float64 = (-2.25 / 3.14159)
                    v15 = 128
                    _ = v18
                }
                var v19 int64 = 0
                _ = v17
                _ = v19
            }
        }
    case 42:
        if (true && false) {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint8 = 127
                if len(v2) > 0 {
                    v2[0] = 0
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v14 = 42
                    var v23 uint16 = 44
                    _ = fi22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            if ((3.14159 * 1e10) >= -2.25) {
                var v24 int8 = -32
                if ((v5 | v13) <= genFn0()) {
                    v1 = v1
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 uint64 = v13
                        if (-9223372036854775808 <= 147) {
                            _ = len(v2)
                            v4 = 4294967295
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v4, v26)
                            for fi27 := 0; fi27 < 2; fi27++ {
                                v4 = v26
                                v4 = (v13 >> ((42) & 63))
                                v28 := GenS2{F0: 687}
                                var v29 uint16 = 65535
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    var v30 int8 = 0
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (true || false) {
                        v6 = append(v6, -906)
                    }
                    _ = v30
                }
                _ = v24
            } else {
                v31 := GenS2{F0: 42}
                fmt.Println(v4)
                _ = v31
            }
        } else {
            switch ("\n\t" + v1) {
            case "hello":
                v6 = append(v6, -136)
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "foo bar":
                fmt.Println(v3)
            default:
                v1 = v1
            }
            v6 = append(v6, 4294967295)
        }
    case -242:
        var v32 int16 = 36
        _ = v32
    default:
        fmt.Println(v13, v5)
    }
    v5 = genFn0()
    v33 := GenS1{F0: 725}
    fmt.Println(v4, v14)
    var v34 int = 2147483647
    for fi35 := 0; fi35 < 3; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := GenS3{F0: 2, F1: 65535}
            if false {
                var v38 int16 = 1
                v39 := GenS4{F0: -166, F1: (^(v34))}
                switch v1 {
                case "a":
                    var v40 uint32 = 2
                    _ = v40
                case "hello":
                    switch (fi35 ^ int(v3)) {
                    case 1:
                        v5 -= (v5 & v13)
                        v13 -= v13
                    case 42:
                        if (true && ((234 > 128) || (0 > 2))) {
                            switch v1 {
                            case "hello":
                                var v41 int8 = 100
                                _ = v41
                            case "":
                                v42 := GenS2{F0: 42}
                                _ = v42
                            }
                        } else {
                            v43 := GenS5{F0: 9223372036854775807, F1: 1, F2: 2147483647}
                            var v44 bool = true
                            _ = v43
                            _ = v44
                        }
                    }
                default:
                    v45 := GenS6{F0: true, F1: -1}
                    _ = v45
                }
                _ = v38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    var v46 float32 = (1e10 * 0.1)
    for fi47 := 0; fi47 < 8; fi47++ {
        switch "" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v48 uint32 = 2147483648
            _ = v48
        }
        for fi49 := 0; fi49 < 3; fi49++ {
            switch (-(fi47)) {
            case 9223372036854775807:
                v50 := GenS5{F0: -9223372036854775808, F1: 42, F2: -1}
                _ = v50
            case -116:
                v51 := GenS7{F0: int(v4)}
                _ = v51
            }
            if (v14 != v14) {
                switch (v1 + "hello") {
                case "\n\t":
                    v52 := GenS8{F0: ((fi47 / ((32768) | 1)) >= (fi47 % (((fi47 | 32768)) | 1)))}
                    _ = v52
                case "foo bar":
                    v5 -= (^(v13))
                case "":
                    fmt.Println(v1, v5)
                }
                v46 -= (float32(1e10) + 0.1)
            }
            _ = fi49
        }
        var v53 int64 = 42
        _ = fi47
        _ = v53
    }
    var v54 float64 = 3.14159
    var v55 uint32 = 42
    var v56 bool = false
    var v57 uint32 = (v55 >> (((v3 * 65535)) & 63))
    if true {
        fmt.Println(v34, v13)
    }
    v58 := GenS5{F0: -9223372036854775808, F1: int64(v13), F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
