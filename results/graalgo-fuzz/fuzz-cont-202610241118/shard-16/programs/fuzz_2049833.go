package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int; F1 int8; F2 int32 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint8; F1 string; F2 int }
func genFn0(p0 int) string {
    var v7 uint16 = 2
    switch ("\n\t" + "foo bar") {
    case "a":
        var v8 string = "foo bar"
        _ = v8
    case "hello":
        var v9 float64 = -2.25
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 14; fi10++ {
            p0 += 716
            fmt.Println(p0)
            v7 = (^(v7))
            _ = fi10
        }
    }
    p0 = (p0 + 32767)
    v7 = (v7 / (((v7 / ((904) | 1))) | 1))
    var v11 int16 = -32768
    _ = p0
    _ = v11
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 32767
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    var v5 uint8 = 128
    var v6 int8 = 1
    var v12 uint8 = v5
    v13 := []int{161, 32767, 956}
    var v14 float64 = 1e10
    var v15 int8 = v6
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v3 += -2.25
    switch "a" {
    case "a":
        v6 = v6
    case "hello":
        if ("foo bar" == (("hello" + ("hello" + ((("\n\t" + "foo bar") + "\n\t") + "\n\t"))) + ("" + ("hello" + ("a" + (("" + "hello") + (("foo bar" + "") + "foo bar"))))))) {
            v16 := GenS1{F0: false, F1: -1}
            if len(v13) > 0 {
                v13[0] = 2147483647
            }
            _ = v16
        }
    default:
        if (int16(v2) != 42) {
            var v17 uint64 = 2
            _ = v17
        } else {
            v14 = float64((-2.25 * 1.5))
            v18 := GenS2{F0: (v2 << ((890) & 63))}
            delete(v4, 32768)
            _ = v18
        }
    }
    switch -426 {
    case -395:
        var v19 int64 = 0
        _ = v19
    }
    if true {
        fmt.Println(v12, v3)
        v20 := GenS3{F0: 532, F1: 964, F2: -1}
        var v21 float32 = (((((v3 + (1e10 + float32(v14))) + v3) * v3) / 3.14159) / (1e10 * -2.25))
        _ = v20
        _ = v21
    } else {
        var v22 string = (("hello" + "foo bar") + "hello")
        v14 = v14
        _ = v22
    }
    fmt.Println(v12)
    if false {
        var v23 int16 = 898
        switch "\n\t" {
        case "hello":
            var v24 string = "foo bar"
            _ = v24
        case "foo bar":
            v25 := GenS4{F0: v6, F1: int(v2)}
            _ = v25
        default:
            for fi26 := 0; fi26 < 12; fi26++ {
                v14 = (1.5 / ((v14 + float64(v3)) - (v14 - 0.1)))
                switch ("foo bar" + "") {
                case "\n\t":
                    if (int64(fi26) >= -9223372036854775808) {
                        if true {
                            v6 = v6
                            if len(v13) > 0 {
                                v13[0] = -534
                            }
                            v6 += (v15 ^ (v15 & 1))
                        }
                    }
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 float64 = 1.5
                    v15 = v15
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        }
        if (9223372036854775807 >= 9223372036854775807) {
            if false {
                fmt.Println(v2, v23, v14)
            } else {
                v29 := GenS5{F0: 2147483647, F1: (v6 ^ v15), F2: 42}
                var v30 int64 = 0
                v5 -= (v5 % ((v12) | 1))
                _ = v29
                _ = v30
            }
            if ((v14 / v14) != 3.14159) {
                var v31 int32 = -1
                v23 = (v23 + v23)
                var v32 float32 = v3
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 7; fi33++ {
                    v12 += v12
                    v12 = (^(v5))
                    if false {
                        switch (fi33 + fi33) {
                        case -1:
                            fmt.Println(v6, fi33)
                        }
                        v23 += 952
                    } else {
                        switch ("foo bar" + "a") {
                        case "a":
                            switch "" {
                            case "foo bar":
                                if (v2 < 42) {
                                    var v34 int32 = int32(v12)
                                    var v35 int32 = v34
                                    v12 = 1
                                    _ = v34
                                    _ = v35
                                } else {
                                    v15 = v6
                                    var v36 float32 = ((1e10 + (1e10 * (0.0 - 0.1))) - 0.0)
                                    switch (^(fi33)) {
                                    case 128:
                                        fmt.Println(v23, v3)
                                    case 9223372036854775807:
                                        if (527 <= 9223372036854775807) {
                                            v14 -= v14
                                            var v37 int = 9223372036854775807
                                            _ = v37
                                        }
                                    case 127:
                                        if false {
                                            for fi38 := 0; fi38 < 10; fi38++ {
                                                v39 := GenS4{F0: v15, F1: (-(fi38))}
                                                v4[int(v6)] = ("" + "hello")
                                                var v40 float32 = v3
                                                var v41 uint = 65535
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            }
                                            fmt.Println(fi33, v2, v14)
                                        } else {
                                            fmt.Println(fi33, v36, v6)
                                            var v42 uint16 = v2
                                            for fi43 := 0; fi43 < 14; fi43++ {
                                                if false {
                                                    var v44 uint32 = 435
                                                    v45 := GenS6{F0: (v12 / ((v12) | 1)), F1: (v23 << ((255) & 63)), F2: (^(v42))}
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    v6 = 67
                                                    var v46 uint8 = 162
                                                    v47 := GenS7{F0: v12, F1: "a", F2: (fi43 - (fi33 / (((^(fi43))) | 1)))}
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 uint64 = 112
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                                _ = fi43
                                                _ = v48
                                            }
                                            _ = v42
                                        }
                                    }
                                    _ = v36
                                }
                            }
                        case "foo bar":
                            for fi49 := 0; fi49 < 18; fi49++ {
                                v14 = v14
                                _ = fi49
                            }
                        case "":
                            for fi50 := 0; fi50 < 2; fi50++ {
                                fmt.Println(v2, v3, v23)
                                v5 += (v5 + v12)
                                fmt.Println(v6)
                                _ = fi50
                            }
                        default:
                            var v51 uint = 18446744073709551615
                            _ = v51
                        }
                    }
                    _ = fi33
                }
            }
            var v52 int32 = 182
            _ = v52
        }
        _ = v23
    } else {
        var v53 string = "hello"
        v54 := GenS3{F0: 1, F1: 254, F2: 127}
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 5; fi55++ {
        v12 = 127
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 string = ("hello" + "hello")
            if false {
                if (-9223372036854775808 != -536) {
                    v12 = (v5 * (v5 >> ((18446744073709551615) & 63)))
                    v5 += (v5 >> ((0) & 63))
                } else {
                    var v58 int16 = 938
                    if (!false) {
                        if (!(v14 < (-2.25 - ((v14 - 1e10) + (0.1 * 1.5))))) {
                            if true {
                                v5 -= (v5 - (v5 + (v5 * v12)))
                                fmt.Println(v5)
                            }
                            if (-2147483648 < 0) {
                                var v59 int32 = 1
                                v2 = (v2 + (v2 << ((4294967295) & 63)))
                                _ = v59
                            } else {
                                if (!(!true)) {
                                    v4[(fi55 >> ((626) & 63))] = ("a" + (((v57 + v57) + ("a" + "")) + ("a" + v57)))
                                    v58 += v58
                                    _ = len(v13)
                                    v2 -= 65535
                                } else {
                                    v60 := GenS6{F0: v12, F1: v58, F2: (v2 ^ (v2 / (((v2 & v2)) | 1)))}
                                    v2 = v2
                                    var v61 int32 = -32
                                    _ = v60
                                    _ = v61
                                }
                            }
                        } else {
                            fmt.Println(v6)
                            fmt.Println(v6, v58)
                            var v62 string = v57
                            _ = v62
                        }
                    } else {
                        var v63 float32 = v3
                        var v64 uint = 42
                        v64 = 0
                        _ = v63
                        _ = v64
                    }
                    _ = v58
                }
            }
            if (v14 == v14) {
                v5 -= (v12 ^ v12)
                var v65 float64 = (1e10 - (v14 / v14))
                v66 := GenS6{F0: (^(v12)), F1: -327, F2: (v2 - 4)}
                _ = len(v13)
                var v67 int64 = int64(v2)
                _ = v65
                _ = v66
                _ = v67
            } else {
                v57 = v57
            }
            fmt.Println(v2)
            _ = fi56
            _ = v57
        }
        fmt.Println(v14)
        _ = fi55
    }
    var v68 int8 = (^(v6))
    for fi69 := 0; fi69 < 13; fi69++ {
        var v70 int32 = 168
        _ = fi69
        _ = v70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
