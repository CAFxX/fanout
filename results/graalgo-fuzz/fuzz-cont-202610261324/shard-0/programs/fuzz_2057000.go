package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 string; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int16; F1 int; F2 uint }
func genFn0() bool {
    v7 := GenS1{F0: -1}
    var v8 float32 = 3.14159
    var v9 uint32 = 2147483647
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 int64 = 42
    var v2 int16 = 42
    v3 := []int{-1, 4294967295}
    var v4 uint32 = 0
    v5 := []string{"\n\t"}
    var v6 int32 = -2147483648
    for fi10 := 0; fi10 < 5; fi10++ {
        switch ("a" + "foo bar") {
        case "foo bar":
            v2 = v2
        case "\n\t":
            var v11 uint16 = uint16(v4)
            _ = v11
        default:
            var v12 uint16 = 32767
            _ = v12
        }
        _ = fi10
    }
    var v13 string = "hello"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (v4 >= (v4 - v4)) {
        var v14 int32 = 1
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (fi15 % ((fi15) | 1)) {
            case 42:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                var v16 bool = false
                _ = v16
            }
            var v17 uint = 0
            _ = fi15
            _ = v17
        }
        _ = v14
    } else {
        v6 -= v6
    }
    switch ((v13 + v13) + v13) {
    case "":
        switch v13 {
        case "\n\t":
            fmt.Println(v13, v2)
        case "":
            v18 := GenS2{F0: "foo bar", F1: 1}
            _ = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            switch ((v13 + ((v13 + v13) + ((v13 + v13) + (v13 + "\n\t")))) + (("a" + ((((v13 + ((v13 + "a") + v13)) + "") + v13) + (v13 + v13))) + v13)) {
            case "a":
                fmt.Println(v2, v13, v2)
            case "hello":
                switch 2147483647 {
                case 9223372036854775807:
                    v20 := GenS1{F0: 0}
                    _ = v20
                case -1:
                    if (true && false) {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v4 -= 0
                            if len(v5) > 0 {
                                v5[0] = v13
                            }
                            _ = fi21
                        }
                        v22 := GenS2{F0: v13, F1: -639}
                        var v23 uint32 = (v4 & v4)
                        _ = v22
                        _ = v23
                    } else {
                        var v24 uint8 = 0
                        fmt.Println(v2)
                        var v25 uint32 = 30
                        _ = v24
                        _ = v25
                    }
                }
            case "foo bar":
                v26 := GenS3{F0: 18446744073709551615, F1: 327, F2: (v6 + 0)}
                _ = v26
            }
            _ = fi19
        }
    case "\n\t":
        switch 4294967295 {
        case 32767:
            var v27 uint8 = 127
            _ = v27
        case 4294967295:
            switch -879 {
            case 0:
                var v28 string = ((v13 + v13) + v13)
                _ = v28
            default:
                if false {
                    switch v13 {
                    case "hello":
                        v29 := GenS4{F0: int16(v1), F1: 4294967295}
                        _ = v29
                    case "":
                        if false {
                            if (!((1e10 >= 0.0) || true)) {
                                fmt.Println(v2)
                            } else {
                                for fi30 := 0; fi30 < 16; fi30++ {
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        switch ((v13 + v13) + ("foo bar" + "\n\t")) {
                                        case "foo bar":
                                            v32 := GenS5{F0: (v13 + v13), F1: 4294967295, F2: 1}
                                            _ = v32
                                        }
                                        _ = fi31
                                    }
                                    if true {
                                        if (!false) {
                                            var v33 int8 = 127
                                            var v34 uint64 = 0
                                            _ = v33
                                            _ = v34
                                        } else {
                                            if true {
                                                switch int(v2) {
                                                case -2147483648:
                                                    if false {
                                                        var v35 uint8 = 1
                                                        switch fi30 {
                                                        case -9223372036854775808:
                                                            if (!true) {
                                                                for fi36 := 0; fi36 < 6; fi36++ {
                                                                    switch (^(fi36)) {
                                                                    case 9223372036854775807:
                                                                        v13 = ("foo bar" + "foo bar")
                                                                    case -9223372036854775808:
                                                                        var v37 int16 = (v2 >> ((249) & 63))
                                                                        _ = v37
                                                                    case -6:
                                                                        v35 += 1
                                                                    }
                                                                    fmt.Println(v35)
                                                                    v1 -= v1
                                                                    _ = fi36
                                                                }
                                                                v4 += 2
                                                            }
                                                        case 4294967295:
                                                            v35 = 0
                                                        case 32767:
                                                            for fi38 := 0; fi38 < 13; fi38++ {
                                                                switch fi38 {
                                                                case 42:
                                                                    v6 -= (v6 / ((-1) | 1))
                                                                case -669:
                                                                    if true {
                                                                        _ = len(v5)
                                                                        if len(v5) > 0 {
                                                                            v5[0] = "foo bar"
                                                                        }
                                                                    }
                                                                }
                                                                _ = fi38
                                                            }
                                                        }
                                                        v1 -= v1
                                                        _ = v35
                                                    }
                                                case -243:
                                                    for fi39 := 0; fi39 < 8; fi39++ {
                                                        var v40 uint16 = 1
                                                        v41 := GenS6{F0: fi39, F1: 724}
                                                        switch fi30 {
                                                        case 32767:
                                                            _ = len(v3)
                                                        }
                                                        _ = fi39
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                }
                                            } else {
                                                v42 := GenS7{F0: fi30}
                                                _ = v42
                                            }
                                            var v43 uint16 = 328
                                            _ = v43
                                        }
                                    } else {
                                        fmt.Println(v2, v4, v1)
                                        var v44 uint32 = (v4 + 2147483648)
                                        _ = v44
                                    }
                                    v45 := GenS7{F0: (fi30 % (((fi30 * fi30)) | 1))}
                                    _ = fi30
                                    _ = v45
                                }
                                var v46 int16 = 0
                                v46 = (v2 << ((838) & 63))
                                _ = v46
                            }
                        }
                    case "\n\t":
                        if (2 == 2) {
                            v6 = v6
                            fmt.Println(v1, v6, v13)
                        } else {
                            if false {
                                var v47 float32 = 0.1
                                _ = v47
                            } else {
                                switch -2147483648 {
                                case 693:
                                    v48 := GenS8{F0: 0, F1: -17}
                                    _ = v48
                                default:
                                    v6 = v6
                                }
                                for fi49 := 0; fi49 < 11; fi49++ {
                                    for fi50 := 0; fi50 < 3; fi50++ {
                                        switch fi50 {
                                        case 128:
                                            v2 = v2
                                        case -1:
                                            if (154 != 2) {
                                                if ((723 > 18446744073709551615) && false) {
                                                    v51 := GenS6{F0: 4294967295, F1: 42}
                                                    v52 := GenS4{F0: (-(v2)), F1: 65535}
                                                    _ = v51
                                                    _ = v52
                                                }
                                                switch v13 {
                                                case "":
                                                    v53 := GenS9{F0: v2, F1: fi49, F2: 502}
                                                    _ = v53
                                                case "hello":
                                                    var v54 int8 = -22
                                                    _ = v54
                                                case "foo bar":
                                                    _ = len(v5)
                                                }
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        fmt.Println(v6, v1)
                                        _ = fi50
                                    }
                                    fmt.Println(v6, fi49)
                                    _ = fi49
                                }
                                fmt.Println(v1)
                            }
                            var v55 int8 = 42
                            switch -214 {
                            case 9223372036854775807:
                                var v56 int8 = (v55 & (-(v55)))
                                _ = v56
                            case -901:
                                v57 := GenS4{F0: (v2 + v2), F1: uint(v4)}
                                _ = v57
                            default:
                                for fi58 := 0; fi58 < 7; fi58++ {
                                    v4 = v4
                                    _ = fi58
                                }
                            }
                            _ = v55
                        }
                    default:
                        if (1 < 84) {
                            var v59 uint = 1
                            var v60 int16 = v2
                            _ = v59
                            _ = v60
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                            v61 := GenS2{F0: (v13 + v13), F1: v1}
                            _ = v61
                        }
                    }
                }
            }
        }
    default:
        _ = len(v5)
    }
    fmt.Println(v2)
    fmt.Println(v13, v2)
    for fi62 := 0; fi62 < 2; fi62++ {
        v1 = v1
        v4 = v4
        _ = fi62
    }
    fmt.Println(v13, v13)
    v13 = ((v13 + ("a" + "foo bar")) + "")
    var v63 uint = 166
    var v64 int8 = -128
    var v65 string = v13
    v66 := GenS1{F0: v64}
    v65 = v65
    if true {
        fmt.Println(v13)
        for fi67 := 0; fi67 < 1; fi67++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi68 := 0; fi68 < 11; fi68++ {
                if len(v3) > 0 {
                    v3[0] = (-(fi68))
                }
                v69 := GenS3{F0: 788, F1: 696, F2: (v6 % (((v6 & v6)) | 1))}
                v1 -= (v1 + v1)
                _ = fi68
                _ = v69
            }
            _ = fi67
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
