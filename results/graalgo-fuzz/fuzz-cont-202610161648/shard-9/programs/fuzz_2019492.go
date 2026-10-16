package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int; F1 uint; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 459
    var v3 int16 = -32768
    v4 := make(map[string]string)
    if false {
        v5 := GenS1{F0: 128, F1: 255, F2: false}
        _ = v5
    }
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 13; fi7++ {
        v1 += 2147483648
        v8 := GenS1{F0: 42, F1: 32767, F2: (v6 >= (v6 / 3.14159))}
        _ = fi7
        _ = v8
    }
    switch (v2 / ((128) | 1)) {
    case 127:
        v2 = (v2 + (v2 / (((^(v2))) | 1)))
    }
    var v9 uint32 = 2
    fmt.Println(v1)
    v10 := v4["hello"]
    switch (("\n\t" + "") + "") {
    case "\n\t":
        var v11 float64 = ((-2.25 * float64(3.14159)) * v6)
        _ = v11
    default:
        var v12 int8 = 126
        _ = v12
    }
    var v13 float64 = v6
    if (3.14159 <= v6) {
        for fi14 := 0; fi14 < 5; fi14++ {
            switch ((v10 + (v10 + v10)) + v10) {
            case "hello":
                fmt.Println(v3, v9, v3)
            case "a":
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS1{F0: 241, F1: 816, F2: (!false)}
                    v6 += 0.0
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi14
        }
        if true {
            v17 := v4[v10]
            if true {
                var v18 uint8 = 255
                fmt.Println(v17, v3, v13)
                v3 = (v3 | -1)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 3; fi19++ {
                    if (true && false) {
                        if (false && true) {
                            var v20 int64 = 148
                            switch 1 {
                            case 128:
                                v13 -= v6
                            case 497:
                                v20 = (v20 & 0)
                            case 9223372036854775807:
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    v22 := []int{-9223372036854775808, -152, (fi21 << ((42) & 63))}
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            v23 := []int{-451, v2}
                            _ = v20
                            _ = v23
                        }
                    } else {
                        v24 := []int{fi19, fi19}
                        _ = v24
                    }
                    switch v2 {
                    case 499:
                        if (v3 == (v3 & 1)) {
                            var v25 int8 = 0
                            if true {
                                v26 := GenS2{F0: v3}
                                _ = v26
                            }
                            v1 = v1
                            _ = v25
                        }
                    case 518:
                        for fi27 := 0; fi27 < 20; fi27++ {
                            if (42 != -95) {
                                v3 = (v3 & (v3 << ((249) & 63)))
                            } else {
                                var v28 float64 = -2.25
                                _ = v28
                            }
                            _ = fi27
                        }
                    }
                    var v29 float64 = float64(0.1)
                    _ = fi19
                    _ = v29
                }
                v13 = (float64(1.5) + v13)
            }
            _ = v17
        }
        if (v1 >= 0) {
            switch ("hello" + v10) {
            case "a":
                v30 := []int{v2}
                _ = v30
            case "foo bar":
                if (!false) {
                    switch v2 {
                    case -1:
                        v3 = 1
                    case 128:
                        if (v1 < 4294967295) {
                            v31 := GenS1{F0: 2, F1: 255, F2: false}
                            _ = v31
                        } else {
                            fmt.Println(v10)
                            v13 = float64(1e10)
                        }
                    default:
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v6 -= ((0.1 / ((((float64(3.14159) + 1.5) * v13) - v13) + v13)) - ((0.0 / 0.1) + ((v13 + v6) - ((v6 * float64(0.1)) * 1e10))))
                            var v33 uint16 = 0
                            _ = fi32
                            _ = v33
                        }
                    }
                    v1 = (^(v1))
                    v34 := GenS3{F0: (v1 >> ((836) & 63))}
                    _ = v34
                }
            case "":
                v35 := v4[v10]
                _ = v35
            default:
                if false {
                    var v36 uint8 = 2
                    for fi37 := 0; fi37 < 15; fi37++ {
                        fmt.Println(v6, v36, v1)
                        v6 = float64(((-2.25 * (-2.25 / 3.14159)) + 3.14159))
                        v2 -= (fi37 << ((18446744073709551615) & 63))
                        _ = fi37
                    }
                    _ = v36
                }
            }
            v38 := GenS2{F0: 1}
            _ = v38
        } else {
            if (255 < uint8(v9)) {
                for fi39 := 0; fi39 < 6; fi39++ {
                    var v40 uint32 = (v9 ^ 590)
                    v1 = 1
                    _ = fi39
                    _ = v40
                }
                for fi41 := 0; fi41 < 8; fi41++ {
                    v9 = (v9 % ((v9) | 1))
                    v42 := []int{fi41, 42, v2}
                    var v43 uint64 = 1
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                v44 := GenS4{F0: 18446744073709551615, F1: 940}
                _ = v44
            } else {
                if ((!false) || ("" != ("hello" + (("\n\t" + v10) + ((("foo bar" + "") + "foo bar") + v10))))) {
                    v45 := GenS4{F0: 4294967295, F1: 805}
                    switch v10 {
                    case "a":
                        v46 := GenS5{F0: v2, F1: 0, F2: v3}
                        _ = v46
                    case "":
                        var v47 bool = (false || ((((0.1 + float32((v6 / float64(((3.14159 + 0.0) + 1.5))))) / float32((v13 * (-2.25 * float64(1.5))))) * 3.14159) == 0.0))
                        _ = v47
                    }
                    _ = v45
                } else {
                    v48 := GenS6{F0: 2}
                    for fi49 := 0; fi49 < 3; fi49++ {
                        for fi50 := 0; fi50 < 17; fi50++ {
                            switch (v10 + v10) {
                            case "\n\t":
                                var v51 bool = (!((!(false && (true && true))) && false))
                                _ = v51
                            case "foo bar":
                                switch (fi50 | fi49) {
                                case 4294967295:
                                    v9 += (v1 | v1)
                                case 0:
                                    var v52 uint = 868
                                    _ = v52
                                default:
                                    switch ((v10 + (v10 + (v10 + v10))) + v10) {
                                    case "hello":
                                        switch ((("a" + v10) + "hello") + v10) {
                                        case "hello":
                                            var v53 uint64 = 2
                                            _ = v53
                                        case "":
                                            var v54 uint32 = (v9 * v1)
                                            _ = v54
                                        default:
                                            if true {
                                                switch fi49 {
                                                case 2147483647:
                                                    fmt.Println(v2)
                                                case -2147483648:
                                                    v55 := GenS3{F0: (^(v1))}
                                                    _ = v55
                                                default:
                                                    v56 := GenS5{F0: 0, F1: uint(v9), F2: (v3 / ((v3) | 1))}
                                                    _ = v56
                                                }
                                                var v57 uint64 = 42
                                                _ = v57
                                            } else {
                                                var v58 int32 = 5
                                                switch fi50 {
                                                case 536:
                                                    if ((v3 & v3) != (^(v3))) {
                                                        fmt.Println(v9, v3)
                                                        fmt.Println(fi50)
                                                        switch (fi50 % (((fi49 + (fi49 ^ (v2 & (v2 % ((fi49) | 1)))))) | 1)) {
                                                        case -779:
                                                            var v59 uint16 = 157
                                                            v9 = (v9 * 65535)
                                                            v60 := GenS7{F0: 220, F1: 201, F2: 4294967295}
                                                            _ = v59
                                                            _ = v60
                                                        case -1:
                                                            fmt.Println(v9, v10, fi50)
                                                        }
                                                    }
                                                }
                                                var v61 float64 = (((0.0 * v13) + v13) - 0.1)
                                                switch -9223372036854775808 {
                                                case 0:
                                                    v62 := GenS1{F0: 127, F1: 42, F2: false}
                                                    _ = v62
                                                case 4294967295:
                                                    fmt.Println(v2, fi50, fi50)
                                                }
                                                fmt.Println(v58)
                                                _ = v58
                                                _ = v61
                                            }
                                        }
                                    case "\n\t":
                                        if true {
                                            if (685 != 0) {
                                                _ = len(v4)
                                                var v63 uint64 = 42
                                                _ = v63
                                            }
                                            var v64 float64 = 0.1
                                            v65 := v4["foo bar"]
                                            v66 := []int{1}
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 float32 = -2.25
                                            _ = v67
                                        }
                                    case "foo bar":
                                        v9 = (v1 * 42)
                                    }
                                }
                            case "hello":
                                var v68 int64 = 1
                                _ = v68
                            }
                            _ = fi50
                        }
                        for fi69 := 0; fi69 < 15; fi69++ {
                            var v70 int64 = 0
                            _ = fi69
                            _ = v70
                        }
                        _ = fi49
                    }
                    v71 := []int{32768, v2, v2}
                    _ = v48
                    _ = v71
                }
                v72 := GenS1{F0: 255, F1: 65535, F2: (2147483647 > 630)}
                v9 = v1
                _ = v72
            }
            if (v1 > (v1 >> ((0) & 63))) {
                var v73 float32 = 0.1
                fmt.Println(v6, v6, v2)
                fmt.Println(v2)
                _ = v73
            }
        }
    }
    var v74 uint32 = v9
    var v75 int8 = 78
    var v76 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v74
    _ = v75
    _ = v76
    _ = v9
}
