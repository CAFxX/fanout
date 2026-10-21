package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 uint64; F1 uint; F2 int64 }
type GenS9 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 uint16) int64 {
    var v4 bool = p0
    p0 = v4
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 bool = ((0.0 / 1.0) > (((-2.25 - (0.1 / (((3.14159 + (1e10 * (3.14159 * (0.1 + -2.25)))) * (0.1 / ((-2.25 + 1.5) - 0.0))) / float32(((3.14159 / (0.1 * 3.14159)) * 3.14159))))) + (float32((1e10 * (1.5 * 0.1))) - (1e10 * -2.25))) * 0.1))
    var v2 string = "hello"
    v3 := []int{1, 127, 4294967295}
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v1)
    case 0:
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := GenS1{F0: 1, F1: 18446744073709551615, F2: false}
            var v7 float32 = 3.14159
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v1 = v1
        case "\n\t":
            var v8 int16 = 42
            _ = v8
        case "a":
            var v9 uint32 = 838
            _ = v9
        default:
            fmt.Println(v2, v2)
        }
    }
    v3 = append(v3, 2147483647)
    v10 := GenS1{F0: 1, F1: 18446744073709551615, F2: v1}
    if len(v3) > 0 {
        v3[0] = -880
    }
    v2 = v2
    v2 = "a"
    var v11 uint64 = 4294967295
    if false {
        switch ("foo bar" + v2) {
        case "a":
            v1 = true
        case "foo bar":
            if true {
                var v12 int32 = -931
                for fi13 := 0; fi13 < 16; fi13++ {
                    if (18446744073709551615 != 2) {
                        for fi14 := 0; fi14 < 5; fi14++ {
                            _ = len(v3)
                            _ = fi14
                        }
                        v12 -= (v12 >> ((65535) & 63))
                    } else {
                        v15 := GenS1{F0: 42, F1: v11, F2: false}
                        for fi16 := 0; fi16 < 8; fi16++ {
                            if (!(true || (!true))) {
                                _ = len(v3)
                                v1 = ((880 < -641) || (3.14159 < 3.14159))
                                v1 = v1
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = fi16
                                }
                                var v17 float32 = 3.14159
                                var v18 string = v2
                                switch -9223372036854775808 {
                                case 299:
                                    v19 := GenS2{F0: 0}
                                    _ = v19
                                case 127:
                                    v20 := GenS3{F0: -1, F1: 90, F2: -2147483648}
                                    _ = v20
                                case -2147483648:
                                    if true {
                                        if (!(18446744073709551615 > 4294967295)) {
                                            v1 = v1
                                            v18 = (v2 + (v2 + v18))
                                        } else {
                                            if false {
                                                if (v1 && (128 > (^(fi16)))) {
                                                    if v1 {
                                                        if true {
                                                            v21 := GenS4{F0: genFn0(v1, 65535)}
                                                            _ = v21
                                                        }
                                                        v22 := GenS5{F0: 2, F1: 4294967295, F2: 255}
                                                        if v1 {
                                                            fmt.Println(fi13, v17, v17)
                                                        } else {
                                                            fmt.Println(fi16)
                                                            switch int(v11) {
                                                            case 42:
                                                                _ = len(v3)
                                                            case 1:
                                                                var v23 uint8 = 91
                                                                _ = v23
                                                            }
                                                            switch ("foo bar" + ("hello" + "hello")) {
                                                            case "":
                                                                var v24 uint32 = uint32(v11)
                                                                if v1 {
                                                                    var v25 int64 = 1
                                                                    v18 = v18
                                                                    _ = v25
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (fi16 & fi16)
                                                                    }
                                                                    v2 = "hello"
                                                                }
                                                                var v26 int16 = 1000
                                                                _ = v24
                                                                _ = v26
                                                            case "foo bar":
                                                                var v27 int32 = v12
                                                                _ = v27
                                                            default:
                                                                fmt.Println(v12, v1)
                                                            }
                                                        }
                                                        _ = v22
                                                    }
                                                    var v28 uint64 = (v11 / (((v11 & v11)) | 1))
                                                    _ = len(v3)
                                                    v12 += (^(v12))
                                                    _ = v28
                                                } else {
                                                    v29 := GenS6{F0: 0, F1: fi16}
                                                    v17 = (v17 * v17)
                                                    var v30 uint16 = 2
                                                    _ = v29
                                                    _ = v30
                                                }
                                            } else {
                                                v31 := GenS5{F0: 0, F1: v11, F2: 42}
                                                _ = v31
                                            }
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if true {
                                            v32 := GenS5{F0: 284, F1: v11, F2: uint(v11)}
                                            var v33 uint32 = 2147483647
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        fmt.Println(v11, v18)
                                        if (!(!(int64(fi16) <= 1))) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v34 := GenS7{F0: 65535, F1: 42}
                                                _ = v34
                                            case "\n\t":
                                                var v35 float32 = 3.14159
                                                _ = v35
                                            }
                                            _ = len(v3)
                                            switch v18 {
                                            case "":
                                                v36 := GenS8{F0: v11, F1: 0, F2: 1}
                                                _ = v36
                                            case "foo bar":
                                                var v37 uint32 = 2147483648
                                                _ = v37
                                            default:
                                                v38 := GenS9{F0: (v11 * v11), F1: 1}
                                                _ = v38
                                            }
                                            fmt.Println(fi13, v17, fi13)
                                        }
                                        var v39 uint16 = 208
                                        _ = v39
                                    }
                                default:
                                    v18 = ("hello" + v2)
                                }
                                var v40 uint8 = 255
                                _ = v17
                                _ = v18
                                _ = v40
                            }
                            if ((1e10 - 3.14159) != -2.25) {
                                if v1 {
                                    var v41 uint8 = 2
                                    fmt.Println(v2, v1, v41)
                                    if v1 {
                                        if v1 {
                                            var v42 int32 = (v12 - -972)
                                            var v43 uint64 = v11
                                            if (1 > v41) {
                                                fmt.Println(v2, fi13)
                                            }
                                            _ = v42
                                            _ = v43
                                        }
                                        v44 := GenS7{F0: 18446744073709551615, F1: 32767}
                                        _ = v44
                                    } else {
                                        var v45 int32 = v12
                                        if ((v1 && v1) && v1) {
                                            if len(v3) > 0 {
                                                v3[0] = fi13
                                            }
                                            if v1 {
                                                if false {
                                                    if (v41 >= v41) {
                                                        v12 = v45
                                                    }
                                                } else {
                                                    v12 -= v12
                                                }
                                                var v46 uint32 = 2147483647
                                                _ = v46
                                            }
                                        }
                                        v12 -= v45
                                        _ = v45
                                    }
                                    _ = len(v3)
                                    fmt.Println(fi16)
                                    _ = v41
                                }
                            } else {
                                fmt.Println(v11, fi16)
                            }
                            _ = fi16
                        }
                        _ = len(v3)
                        _ = v15
                    }
                    _ = fi13
                }
                _ = v12
            }
        case "":
            v2 = "\n\t"
        }
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(v11, fi47)
            v11 += v11
            _ = fi47
        }
        fmt.Println(v11, v1)
    } else {
        if v1 {
            _ = len(v3)
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi48 := 0; fi48 < 9; fi48++ {
        if false {
            var v49 int32 = -2147483648
            for fi50 := 0; fi50 < 8; fi50++ {
                fmt.Println(fi48)
                fmt.Println(fi50)
                _ = fi50
            }
            _ = v49
        } else {
            for fi51 := 0; fi51 < 19; fi51++ {
                fmt.Println(v11)
                _ = fi51
            }
            if v1 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = fi48
    }
    v11 = (^(v11))
    v1 = (18446744073709551615 < 42)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
