package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 int32 }
type GenS4 struct { F0 bool; F1 uint32; F2 int32 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint32) int {
    p0 = 428
    p0 += p0
    var v4 uint8 = 0
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint32 = (p1 << ((81) & 63))
        p1 = (^(v6))
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p0, v4)
    _ = p0
    _ = p1
    _ = v4
    return 32768
}
func genFn1(p0 int, p1 int64, p2 uint32) string {
    v9 := []int{}
    for fi10 := 0; fi10 < 15; fi10++ {
        switch "\n\t" {
        case "hello":
            v11 := GenS1{F0: "a"}
            _ = v11
        case "":
            if len(v9) > 0 {
                v9[0] = (^(p0))
            }
        default:
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
        }
        fmt.Println(p2, fi10)
        switch "foo bar" {
        case "":
            fmt.Println(p1, p1, p1)
        case "\n\t":
            var v12 string = "foo bar"
            _ = v12
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint64 = 42
                _ = fi13
                _ = v14
            }
        }
        _ = fi10
    }
    switch p0 {
    case -1:
        switch "foo bar" {
        case "\n\t":
            switch (p0 + p0) {
            case -9223372036854775808:
                switch int(p2) {
                case 127:
                    var v15 string = "a"
                    _ = v15
                case 1:
                    v16 := GenS2{F0: 1, F1: (p2 << ((196) & 63))}
                    _ = v16
                case -9223372036854775808:
                    for fi17 := 0; fi17 < 3; fi17++ {
                        for fi18 := 0; fi18 < 13; fi18++ {
                            fmt.Println(fi17, p0, fi18)
                            _ = fi18
                        }
                        var v19 int16 = -32768
                        var v20 uint16 = 944
                        _ = fi17
                        _ = v19
                        _ = v20
                    }
                default:
                    if true {
                        v21 := GenS3{F0: 1000, F1: 0, F2: 2147483647}
                        _ = v21
                    }
                }
            }
        default:
            var v22 int16 = 0
            _ = v22
        }
    case 32767:
        p2 = p2
    case 914:
        var v23 string = "foo bar"
        _ = v23
    }
    var v24 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int8 = 1
    if true {
        v25 := GenS1{F0: "foo bar"}
        if (32767 >= 1) {
            var v26 uint8 = 1
            v2["foo bar"] = true
            var v27 float64 = (-2.25 + (0.0 * (((((-2.25 * 1e10) * (((((float64(0.1) * 3.14159) - (1.5 + ((1.5 * float64(float32(1.5))) - float64((0.0 - 0.0))))) + (float64(float32(3.14159)) + float64((1.5 + 1e10)))) * (float64(3.14159) - float64(float32((1e10 / 1.0))))) + (float64(-2.25) - float64(-2.25)))) * 1.5) + -2.25) + ((0.0 * (((0.1 + 0.0) - 1.5) / 1.0)) - 0.0))))
            _ = v26
            _ = v27
        }
        var v28 uint16 = 65535
        _ = v25
        _ = v28
    } else {
        switch -1 {
        case -2147483648:
            v29 := GenS2{F0: 0, F1: 964}
            _ = v29
        case -1:
            if false {
                var v30 int16 = 0
                v3 = 100
                _ = v30
            }
        case 42:
            switch "hello" {
            case "\n\t":
                v3 += v3
            case "hello":
                v3 += (v3 >> ((1) & 63))
            }
        default:
            if true {
                for fi31 := 0; fi31 < 7; fi31++ {
                    v32 := []int{(fi31 ^ fi31)}
                    _ = fi31
                    _ = v32
                }
            } else {
                v1[32768] = ("\n\t" + "a")
                if false {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        var v34 int32 = -799
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS2{F0: 162, F1: 42}
                    var v36 float32 = (1.5 - (((1e10 * 0.1) / ((1.5 / ((0.0 * (0.0 - (float32((((((float64(0.0) * 3.14159) / (-2.25 + float64(0.0))) * 0.1) * 0.1) / 1.5)) + 0.1))) - float32(3.14159))) - 0.0)) + (-2.25 + 1.5)))
                    _ = v35
                    _ = v36
                }
            }
        }
    }
    v37 := v1[2147483647]
    var v38 uint32 = 0
    var v39 uint64 = 617
    if (-2.25 != 1.5) {
        var v40 int32 = 42
        if true {
            fmt.Println(v38, v39, v39)
        } else {
            for fi41 := 0; fi41 < 3; fi41++ {
                for fi42 := 0; fi42 < 6; fi42++ {
                    v43 := GenS2{F0: 1, F1: (v38 << ((255) & 63))}
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v38 += 2
                    }
                    v1[(fi42 | fi41)] = "foo bar"
                    fmt.Println(fi41, fi42, v40)
                    _ = fi42
                    _ = v43
                }
                var v44 uint = 384
                fmt.Println(v38, v40, v37)
                _ = fi41
                _ = v44
            }
            v1[1] = (v37 + (v37 + v37))
        }
        if true {
            for fi45 := 0; fi45 < 14; fi45++ {
                v46 := GenS4{F0: (false || (905 != 2)), F1: 1, F2: 370}
                for fi47 := 0; fi47 < 16; fi47++ {
                    if (!false) {
                        var v48 uint64 = (^(v39))
                        var v49 int = (^(fi47))
                        v49 = fi47
                        _ = v48
                        _ = v49
                    } else {
                        var v50 uint16 = 0
                        _ = v50
                    }
                    _ = fi47
                }
                for fi51 := 0; fi51 < 7; fi51++ {
                    v40 -= (v40 & v40)
                    if (!(!(865 <= 628))) {
                        switch v37 {
                        case "":
                            fmt.Println(fi51, fi51)
                        case "\n\t":
                            var v52 string = v37
                            _ = v52
                        case "a":
                            v3 -= v3
                        default:
                            v37 = ((((v37 + v37) + ((v37 + v37) + ("foo bar" + v37))) + (v37 + ((v37 + (v37 + "a")) + "\n\t"))) + "a")
                        }
                    }
                    fmt.Println(v3, fi51, v40)
                    _ = fi51
                }
                _ = fi45
                _ = v46
            }
            fmt.Println(v38, v37, v39)
        } else {
            switch "" {
            case "a":
                var v53 float64 = 0.0
                _ = v53
            case "foo bar":
                v38 -= (v38 - (v38 + v38))
            case "\n\t":
                fmt.Println(v38, v3, v37)
            }
        }
        _ = v40
    }
    v38 += (v38 >> ((294) & 63))
    v37 = (v37 + ("" + v37))
    if true {
        fmt.Println(v37, v37, v39)
    } else {
        fmt.Println(v39, v3, v3)
        v3 = int8(v38)
        v39 = v39
    }
    switch "\n\t" {
    case "":
        v37 = (((v37 + "a") + (v37 + ("\n\t" + ""))) + v37)
    case "a":
        v54 := []int{genFn0(1, uint32(v39)), -760, -2147483648}
        _ = v54
    }
    switch 4294967295 {
    case 42:
        if false {
            v38 += (v38 % ((v38) | 1))
            if true {
                v3 = (^(v3))
            } else {
                switch 42 {
                case 127:
                    fmt.Println(v39, v39)
                case -319:
                    v38 -= (v38 - (v38 % ((v38) | 1)))
                case -2147483648:
                    v3 += v3
                }
            }
            var v55 int64 = 9223372036854775807
            _ = v55
        } else {
            var v56 string = ("a" + v37)
            if (42 < -2147483648) {
                var v57 int64 = 0
                _ = v57
            } else {
                var v58 int16 = -32768
                switch -2147483648 {
                case 384:
                    fmt.Println(v3, v56, v37)
                case -545:
                    if false {
                        for fi59 := 0; fi59 < 14; fi59++ {
                            if ((0.1 + ((((float64(3.14159) / -2.25) - 1.5) - float64(0.0)) * 1.5)) > (0.1 - (-2.25 * 3.14159))) {
                                v60 := GenS1{F0: "\n\t"}
                                v61 := GenS3{F0: (v58 >> ((42) & 63)), F1: (v3 ^ (v3 % ((-17) | 1))), F2: 0}
                                _ = v60
                                _ = v61
                            }
                            v62 := GenS5{F0: false, F1: 225}
                            _ = fi59
                            _ = v62
                        }
                    }
                case 4294967295:
                    var v63 int64 = -358
                    _ = v63
                default:
                    v2[v56] = ((0.1 * (3.14159 - 0.0)) >= -2.25)
                }
                _ = v58
            }
            _ = v56
        }
    case -9223372036854775808:
        v64 := GenS6{F0: (v39 | 4294967295)}
        _ = v64
    default:
        for fi65 := 0; fi65 < 16; fi65++ {
            v38 += (v38 % (((v38 & v38)) | 1))
            var v66 uint16 = 32767
            _ = fi65
            _ = v66
        }
    }
    var v67 uint32 = v38
    v1[-9223372036854775808] = v37
    var v68 string = "hello"
    var v69 int64 = 0
    for fi70 := 0; fi70 < 12; fi70++ {
        switch fi70 {
        case 127:
            v3 = v3
        }
        _ = fi70
    }
    fmt.Println(v68)
    var v71 int8 = (v3 | -1)
    switch v37 {
    case "hello":
        v72 := GenS2{F0: 32767, F1: v67}
        _ = v72
    case "foo bar":
        v68 = ((v68 + (v68 + v68)) + v37)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v67
    _ = v68
    _ = v69
    _ = v71
}
