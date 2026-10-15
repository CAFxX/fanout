package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 int64; F2 int }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS10 struct { F0 uint32; F1 uint16 }
func genFn0() uint64 {
    v5 := []int{-9223372036854775808, 42, -2147483648, -366}
    v6 := GenS1{F0: "foo bar", F1: (("\n\t" + "foo bar") + "\n\t")}
    if true {
        v5 = append(v5, 256)
    } else {
        if (1 != 1) {
            switch 128 {
            case 32767:
                var v7 float32 = ((float32(float64((0.1 * float32((1.5 * 0.0))))) / 1.5) + 0.1)
                _ = v7
            case -707:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        } else {
            if true {
                switch "foo bar" {
                case "foo bar":
                    var v8 uint16 = 715
                    _ = v8
                case "hello":
                    var v9 uint16 = 408
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v11 := GenS2{F0: 2147483648, F1: -1, F2: fi10}
                        var v12 int64 = 952
                        switch fi10 {
                        case -1:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case 0:
                            v5 = append(v5, fi10)
                        case 4294967295:
                            if (!(2 < 184)) {
                                var v13 uint32 = 947
                                v12 += v12
                                switch "a" {
                                case "\n\t":
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        fmt.Println(v12, v12, fi10)
                                    case 1:
                                        v14 := GenS3{F0: (v12 | v12), F1: 0, F2: v13}
                                        _ = v14
                                    }
                                case "a":
                                    var v15 string = ""
                                    _ = v15
                                }
                                _ = v13
                            }
                        default:
                            v5 = append(v5, (fi10 ^ (fi10 << ((255) & 63))))
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
            }
            if len(v5) > 0 {
                v5[0] = 1
            }
        }
    }
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int8 = -1
    v2 := []string{}
    var v3 int8 = v1
    var v4 int64 = 42
    var v16 uint64 = 1
    var v17 float32 = 3.14159
    if true {
        switch ("a" + ("\n\t" + "\n\t")) {
        case "a":
            var v18 uint = 18446744073709551615
            _ = v18
        case "\n\t":
            v19 := GenS2{F0: 130, F1: v4, F2: 0}
            _ = v19
        }
    } else {
        v20 := GenS4{F0: 1, F1: 874}
        var v21 string = "hello"
        v22 := GenS5{F0: (v4 ^ (v4 ^ (v4 - v4))), F1: v3}
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v2)
        _ = fi23
    }
    fmt.Println(v17, v4, v4)
    _ = len(v2)
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 int64 = (v4 * (v4 / ((v4) | 1)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!false) {
            var v26 int16 = 1
            for fi27 := 0; fi27 < 7; fi27++ {
                switch -534 {
                case -1:
                    var v28 float64 = 0.1
                    _ = v28
                case 42:
                    v16 = (v16 >> ((991) & 63))
                default:
                    v29 := GenS2{F0: 618, F1: -186, F2: fi27}
                    _ = v29
                }
                v26 = (v26 / (((-(v26))) | 1))
                v30 := GenS1{F0: "hello", F1: ""}
                var v31 int16 = v26
                _ = fi27
                _ = v30
                _ = v31
            }
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    switch "a" {
    case "\n\t":
        v32 := GenS3{F0: (^(v4)), F1: (v16 >> ((42) & 63)), F2: 46}
        _ = v32
    }
    if ((!false) && (0 != 42)) {
        if false {
            if true {
                fmt.Println(v3)
            } else {
                for fi33 := 0; fi33 < 1; fi33++ {
                    var v34 float32 = 0.0
                    if true {
                        v35 := GenS5{F0: v4, F1: (v3 + (^(v3)))}
                        _ = v35
                    } else {
                        var v36 uint64 = genFn0()
                        switch "foo bar" {
                        case "a":
                            var v37 uint = 18446744073709551615
                            _ = v37
                        case "\n\t":
                            var v38 int = 569
                            _ = v38
                        case "hello":
                            var v39 string = "hello"
                            _ = v39
                        }
                        _ = v36
                    }
                    switch ("\n\t" + "hello") {
                    case "hello":
                        if false {
                            v4 += v4
                            v17 -= float32(0.1)
                        }
                    case "":
                        var v40 uint64 = (v16 % (((v16 - v16)) | 1))
                        _ = v40
                    default:
                        if (!(((v16 >> ((42) & 63)) != v16) && true)) {
                            fmt.Println(fi33)
                            var v41 int64 = v4
                            _ = len(v2)
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                v43 := GenS6{F0: 128, F1: 65535, F2: 0}
                                var v44 uint64 = genFn0()
                                v16 -= v16
                                v44 -= 0
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                v45 := GenS7{F0: (v3 >> ((255) & 63)), F1: false, F2: -645}
                _ = v45
            }
            for fi46 := 0; fi46 < 18; fi46++ {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = fi46
            }
            if ((true || (false || (-590 > -9223372036854775808))) && true) {
                for fi47 := 0; fi47 < 3; fi47++ {
                    v16 -= v16
                    v1 = v1
                    switch ("a" + ("" + "hello")) {
                    case "foo bar":
                        if false {
                            var v48 int32 = 792
                            _ = v48
                        } else {
                            var v49 int64 = v4
                            _ = v49
                        }
                    case "hello":
                        var v50 bool = (80 > 217)
                        _ = v50
                    default:
                        _ = len(v2)
                    }
                    _ = fi47
                }
                if false {
                    _ = len(v2)
                    var v51 int = 32768
                    for fi52 := 0; fi52 < 9; fi52++ {
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            fmt.Println(v51, v17, v16)
                        case "foo bar":
                            v3 = 127
                        }
                        _ = fi52
                    }
                    _ = v51
                }
            }
        } else {
            v53 := GenS8{F0: 2}
            var v54 int64 = (v4 & 9223372036854775807)
            for fi55 := 0; fi55 < 14; fi55++ {
                _ = len(v2)
                _ = fi55
            }
            _ = v53
            _ = v54
        }
        if (-2147483648 >= 42) {
            if (!(127 == 0)) {
                v56 := GenS1{F0: "\n\t", F1: "a"}
                _ = v56
            }
            v1 += v3
        } else {
            switch 261 {
            case 1:
                for fi57 := 0; fi57 < 16; fi57++ {
                    for fi58 := 0; fi58 < 8; fi58++ {
                        v59 := GenS9{F0: -2147483648, F1: 627, F2: 32767}
                        fmt.Println(v4, v17)
                        var v60 bool = (32767 <= -32768)
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                    _ = fi57
                }
            case -113:
                v17 += 1.5
            case -829:
                for fi61 := 0; fi61 < 16; fi61++ {
                    v3 = v1
                    v17 += ((v17 * (v17 + (((v17 - ((v17 - 0.1) - float32(1.5))) / (v17 * 0.1)) - ((((v17 + v17) * float32(1e10)) / v17) + float32(0.0))))) * (1.5 - (v17 / (float32(1.5) / ((((v17 + v17) + v17) * 1e10) + (3.14159 * (3.14159 * ((-2.25 + float32(0.0)) + ((3.14159 / ((v17 + v17) - v17)) - v17)))))))))
                    _ = len(v2)
                    _ = fi61
                }
            }
            fmt.Println(v3, v1)
            switch "a" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = ("a" + ((((("hello" + (("hello" + "hello") + "\n\t")) + ((("" + ("foo bar" + "a")) + "foo bar") + "hello")) + "\n\t") + (("\n\t" + "\n\t") + "a")) + "foo bar"))
                }
            default:
                switch "a" {
                case "\n\t":
                    for fi62 := 0; fi62 < 20; fi62++ {
                        for fi63 := 0; fi63 < 17; fi63++ {
                            v1 = 100
                            _ = fi63
                        }
                        var v64 uint32 = 42
                        _ = fi62
                        _ = v64
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    var v65 float32 = 0.0
                    _ = v65
                default:
                    var v66 uint = 42
                    _ = v66
                }
            }
        }
        switch (("" + "") + "\n\t") {
        case "\n\t":
            v67 := GenS10{F0: 0, F1: 0}
            _ = v67
        case "hello":
            for fi68 := 0; fi68 < 9; fi68++ {
                if false {
                    v2 = append(v2, (("" + ((((("a" + ("" + "a")) + "a") + "foo bar") + "hello") + "a")) + "hello"))
                }
                _ = fi68
            }
        case "":
            v69 := GenS3{F0: 42, F1: 4294967295, F2: 2}
            _ = v69
        }
    } else {
        var v70 bool = (2 <= 439)
        _ = v70
    }
    var v71 uint16 = 347
    v72 := GenS4{F0: 1, F1: v71}
    fmt.Println(v17, v16, v4)
    fmt.Println(v17, v1, v71)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
}
