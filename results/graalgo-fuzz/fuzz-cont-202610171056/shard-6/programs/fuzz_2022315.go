package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int16; F1 int16; F2 bool }
type GenS5 struct { F0 string; F1 uint64 }
func genFn0() int64 {
    var v4 uint64 = 1
    v4 -= (v4 >> ((18446744073709551615) & 63))
    _ = v4
    return 1
}
func genFn1() int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        if false {
            if true {
                var v6 int16 = -32768
                _ = v6
            } else {
                v7 := GenS1{F0: 0}
                var v8 float32 = 0.0
                _ = v7
                _ = v8
            }
            var v9 uint64 = 560
            var v10 bool = false
            _ = v10
            _ = v9
        } else {
            v11 := []int{}
            var v12 uint16 = 539
            v13 := GenS2{F0: fi5, F1: "\n\t", F2: ("\n\t" + "foo bar")}
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 string = ("foo bar" + "hello")
        _ = fi5
        _ = v14
    }
    var v15 int64 = 362
    _ = v15
    return -42
}

func main() {
    var v1 float32 = float32(((1.5 / 1.0) * (1.5 * 1.5)))
    v2 := make(map[string]string)
    var v3 uint16 = 0
    if (!false) {
        if false {
            for fi16 := 0; fi16 < 15; fi16++ {
                if (37 > 164) {
                    var v17 int8 = -28
                    if false {
                        var v18 int32 = -239
                        v3 -= v3
                        if ((!true) && false) {
                            v1 += 0.1
                            v19 := []int{fi16, (fi16 % (((fi16 ^ (-(fi16)))) | 1))}
                            _ = v19
                        }
                        _ = v18
                    }
                    v2["a"] = ""
                    _ = v17
                } else {
                    var v20 int8 = -128
                    _ = v20
                }
                var v21 int32 = 2147483647
                _ = fi16
                _ = v21
            }
            if true {
                switch "" {
                case "":
                    for fi22 := 0; fi22 < 5; fi22++ {
                        var v23 int = 0
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v3 = 873
                            switch (fi24 % (((^(v23))) | 1)) {
                            case -2147483648:
                                v1 += v1
                            case 42:
                                delete(v2, "a")
                                fmt.Println(v1, v1, v23)
                                if (!(!(0 < v3))) {
                                    v1 = v1
                                    var v25 uint32 = 808
                                    _ = v25
                                }
                                v23 -= 175
                            default:
                                var v26 float64 = 0.0
                                switch (fi24 | 9223372036854775807) {
                                case -9223372036854775808:
                                    delete(v2, "foo bar")
                                }
                                _ = v26
                            }
                            v3 = (v3 & 507)
                            _ = fi24
                        }
                        _ = fi22
                        _ = v23
                    }
                case "a":
                    v3 = (v3 * (v3 | v3))
                case "hello":
                    v1 += 0.0
                }
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 uint16 = v3
                    var v29 int16 = 32767
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch ("hello" + ("\n\t" + "foo bar")) {
                case "":
                    v1 += 1e10
                }
            } else {
                if false {
                    var v30 int16 = 1
                    v31 := GenS3{F0: (!(0 != -9223372036854775808)), F1: 1}
                    _ = v30
                    _ = v31
                }
            }
        }
        switch 1 {
        case 42:
            v32 := GenS3{F0: ((1.5 / (1e10 / (-2.25 - (-2.25 / -2.25)))) <= ((0.1 + 1.5) * (float64(v1) - 1.5))), F1: 128}
            _ = v32
        }
    }
    if true {
        if (false || true) {
            var v33 string = ("a" + "\n\t")
            var v34 uint64 = 4294967295
            v2[((v33 + "a") + v33)] = ("hello" + v33)
            _ = v33
            _ = v34
        }
        v35 := []int{}
        _ = v35
    } else {
        v36 := []int{-1, 9223372036854775807, 127, 2147483647}
        var v37 int8 = 100
        switch 32767 {
        case 0:
            var v38 int16 = int16(v3)
            _ = v38
        case 323:
            var v39 uint16 = 196
            _ = v39
        case -252:
            var v40 float64 = (-2.25 / 3.14159)
            _ = v40
        }
        _ = v36
        _ = v37
    }
    if (-41 < 42) {
        for fi41 := 0; fi41 < 9; fi41++ {
            _ = len(v2)
            _ = fi41
        }
        if true {
            v42 := GenS1{F0: 0}
            v43 := GenS4{F0: 1000, F1: 42, F2: (v1 != (float32((float64(((1.5 / 3.14159) - (v1 + float32(((0.1 - (0.0 * (1.5 * 0.0))) / -2.25))))) / (((((float64(float32((0.1 - (((-2.25 * float64((1e10 * 0.0))) * ((3.14159 / (-2.25 / 1.0)) / 1.0)) - (0.1 - 1e10))))) * 1e10) / float64(((v1 - float32(3.14159)) / ((v1 - (v1 / (v1 / 1.0))) - 3.14159)))) * float64(v1)) * -2.25) - (-2.25 / 0.1)))) / v1))}
            switch 494 {
            case -463:
                var v44 float64 = 0.1
                _ = v44
            case -843:
                var v45 int64 = genFn0()
                _ = v45
            default:
                var v46 string = ("a" + "\n\t")
                _ = v46
            }
            _ = v42
            _ = v43
        } else {
            for fi47 := 0; fi47 < 10; fi47++ {
                v48 := GenS4{F0: 201, F1: 0, F2: (!true)}
                if (-326 >= -1) {
                    switch "a" {
                    case "a":
                        v2["\n\t"] = "\n\t"
                    case "":
                        for fi49 := 0; fi49 < 17; fi49++ {
                            v1 = 1.5
                            var v50 uint = 255
                            _ = fi49
                            _ = v50
                        }
                    case "hello":
                        switch "" {
                        case "\n\t":
                            var v51 float32 = v1
                            _ = v51
                        case "":
                            v3 -= 705
                        default:
                            var v52 float32 = 1e10
                            _ = v52
                        }
                    default:
                        v3 = v3
                    }
                    if false {
                        v1 -= v1
                    }
                }
                v53 := GenS5{F0: "foo bar", F1: 4294967295}
                _ = fi47
                _ = v48
                _ = v53
            }
        }
        var v54 int8 = 127
        _ = v54
    } else {
        delete(v2, "\n\t")
        for fi55 := 0; fi55 < 5; fi55++ {
            fmt.Println(v3, v3)
            v3 = (^(v3))
            _ = fi55
        }
    }
    var v56 float32 = 3.14159
    fmt.Println(v1, v1)
    var v57 int8 = genFn1()
    switch ("a" + "a") {
    case "hello":
        v1 = (1e10 / ((float32(0.1) * 1e10) * 0.1))
    case "\n\t":
        for fi58 := 0; fi58 < 4; fi58++ {
            var v59 float32 = -2.25
            var v60 int8 = genFn1()
            _ = fi58
            _ = v59
            _ = v60
        }
    case "foo bar":
        v1 -= v1
    default:
        var v61 uint16 = (v3 - (^(v3)))
        _ = v61
    }
    var v62 uint8 = 1
    v63 := GenS5{F0: ("hello" + "a"), F1: 258}
    v1 += (float32(3.14159) / v56)
    v64 := v2["\n\t"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v62
    _ = v63
    _ = v64
}
