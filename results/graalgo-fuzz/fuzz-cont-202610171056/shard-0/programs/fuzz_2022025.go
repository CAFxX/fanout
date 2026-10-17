package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int64; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 bool }
func genFn0(p0 uint, p1 bool, p2 uint16) uint32 {
    p0 += uint(p2)
    v5 := GenS1{F0: -32768}
    var v6 uint8 = 215
    var v7 uint32 = 945
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (v7 + (v7 << ((2) & 63)))
}

func main() {
    var v1 int32 = 0
    var v2 bool = false
    var v3 float64 = 1.5
    v4 := make(map[uint32]bool)
    switch 0 {
    case 4294967295:
        if (-2.25 < ((0.0 + -2.25) + 0.0)) {
            if (-9223372036854775808 == 9223372036854775807) {
                var v8 uint32 = 2147483648
                _ = v8
            } else {
                v4[393] = (127 <= 78)
                var v9 uint8 = 42
                var v10 float32 = 1.5
                _ = v10
                _ = v9
            }
        } else {
            fmt.Println(v2, v2, v1)
            if (!(!v2)) {
                v3 = -2.25
                v2 = (v2 && true)
            } else {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v2 = (!v2)
                    _ = fi11
                }
                if true {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v13 := []int{2147483647}
                        _ = fi12
                        _ = v13
                    }
                } else {
                    v3 = 1.5
                    if true {
                        v2 = ((((v3 + (float64(1.5) * v3)) / 3.14159) <= (float64(((3.14159 / float32(1.5)) - (0.1 * -2.25))) + 0.0)) && true)
                        v14 := GenS2{F0: 9223372036854775807}
                        _ = v14
                    } else {
                        v15 := []int{9223372036854775807, -830, 2147483647}
                        v3 += v3
                        if (("hello" > "hello") || (false || v2)) {
                            if (668 < 2147483647) {
                                if false {
                                    v1 = 2147483647
                                    if (341 >= 65535) {
                                        if (0 < -973) {
                                            for fi16 := 0; fi16 < 6; fi16++ {
                                                for fi17 := 0; fi17 < 20; fi17++ {
                                                    v1 = -680
                                                    v18 := GenS3{F0: (292 <= 495), F1: 114, F2: -1}
                                                    _ = fi17
                                                    _ = v18
                                                }
                                                var v19 float32 = -2.25
                                                v3 = (0.1 / 3.14159)
                                                _ = fi16
                                                _ = v19
                                            }
                                            var v20 int64 = -1
                                            var v21 int8 = -1
                                            _ = v20
                                            _ = v21
                                        }
                                        v2 = v2
                                        if true {
                                            var v22 int8 = 127
                                            fmt.Println(v3)
                                            _ = v22
                                        } else {
                                            v2 = false
                                            v23 := GenS1{F0: -1}
                                            _ = len(v15)
                                            _ = v23
                                        }
                                    }
                                    switch ("" + (("foo bar" + (("\n\t" + "a") + "a")) + "foo bar")) {
                                    case "hello":
                                        for fi24 := 0; fi24 < 16; fi24++ {
                                            v1 = (v1 | 561)
                                            _ = fi24
                                        }
                                    }
                                } else {
                                    if ((v3 != (1.5 / float64(1.5))) && true) {
                                        v3 = v3
                                    }
                                    var v25 uint64 = 1
                                    if v2 {
                                        var v26 uint8 = 254
                                        var v27 int = 42
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v25
                                }
                                if (898 > 323) {
                                    for fi28 := 0; fi28 < 7; fi28++ {
                                        var v29 uint32 = 2147483647
                                        _ = fi28
                                        _ = v29
                                    }
                                    var v30 uint64 = 819
                                    if len(v15) > 0 {
                                        v15[0] = -9223372036854775808
                                    }
                                    _ = v30
                                }
                            }
                            v3 -= 1.5
                        }
                        _ = v15
                    }
                    switch "foo bar" {
                    case "":
                        v31 := []int{2147483647, 2147483647, 729, 0}
                        _ = v31
                    case "foo bar":
                        var v32 float64 = v3
                        _ = v32
                    case "a":
                        for fi33 := 0; fi33 < 4; fi33++ {
                            switch "\n\t" {
                            case "":
                                switch fi33 {
                                case 145:
                                    var v34 float64 = ((v3 - (v3 + v3)) - float64(((1e10 / -2.25) * ((float32(v3) + 1.5) - 1e10))))
                                    _ = v34
                                case 127:
                                    if (("hello" <= "hello") && v2) {
                                        if (4294967295 <= 2) {
                                            v3 += v3
                                            var v35 uint64 = 507
                                            v3 -= (v3 - (v3 - (v3 - 1.5)))
                                            _ = v35
                                        }
                                    }
                                }
                            case "hello":
                                var v36 uint16 = 65535
                                _ = v36
                            }
                            if (276 == 42) {
                                v37 := GenS4{F0: (fi33 >= fi33)}
                                v38 := GenS4{F0: false}
                                v39 := GenS2{F0: fi33}
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                            switch (("a" + ("" + ("hello" + "\n\t"))) + "a") {
                            case "foo bar":
                                v40 := GenS2{F0: (-(fi33))}
                                _ = v40
                            default:
                                v41 := []int{32767, fi33, (^(fi33))}
                                _ = v41
                            }
                            _ = fi33
                        }
                    default:
                        var v42 uint8 = 2
                        _ = v42
                    }
                }
                fmt.Println(v1, v1, v2)
            }
            v1 += v1
        }
    case -511:
        if v2 {
            var v43 uint8 = 42
            _ = v43
        } else {
            var v44 uint32 = 65535
            _ = v44
        }
    case -386:
        for fi45 := 0; fi45 < 4; fi45++ {
            var v46 uint64 = 42
            _ = fi45
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        v2 = (59 > 0)
        _ = fi47
    }
    v3 -= float64(float32((v3 * 3.14159)))
    if v2 {
        var v48 bool = false
        var v49 int32 = v1
        fmt.Println(v2)
        _ = v48
        _ = v49
    } else {
        v50 := GenS5{F0: v2}
        v3 = float64((1e10 + 3.14159))
        _ = v50
    }
    var v51 uint64 = 148
    v1 = v1
    v51 = v51
    switch ("" + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v52 int32 = v1
        _ = v52
    case "\n\t":
        fmt.Println(v2)
    case "":
        v53 := v4[genFn0(65535, (v2 || v2), 392)]
        _ = v53
    default:
        var v54 string = "a"
        _ = v54
    }
    switch ((("foo bar" + (("" + ("" + "hello")) + "")) + (("a" + (("" + "hello") + ((("" + "a") + "hello") + ("hello" + "hello")))) + "\n\t")) + "a") {
    case "hello":
        for fi55 := 0; fi55 < 5; fi55++ {
            v56 := GenS4{F0: (v2 && (v2 && (v2 || v2)))}
            for fi57 := 0; fi57 < 2; fi57++ {
                fmt.Println(v3, v3, fi57)
                fmt.Println(v2, fi57, fi55)
                _ = fi57
            }
            _ = fi55
            _ = v56
        }
    case "foo bar":
        fmt.Println(v2, v1)
    case "\n\t":
        fmt.Println(v3)
    default:
        for fi58 := 0; fi58 < 13; fi58++ {
            for fi59 := 0; fi59 < 20; fi59++ {
                var v60 uint32 = 4294967295
                v1 = int32(v60)
                var v61 int16 = -432
                var v62 float64 = (v3 / v3)
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
            v2 = v2
            if v2 {
                fmt.Println(v51, v2)
                v63 := []int{}
                _ = v63
            } else {
                v64 := GenS6{F0: (v51 >> ((590) & 63))}
                v51 = 0
                _ = len(v4)
                _ = v64
            }
            _ = fi58
        }
    }
    fmt.Println(v2)
    var v65 uint16 = 42
    v66 := GenS7{F0: v51, F1: (!(100 <= 40))}
    var v67 float64 = v3
    var v68 uint64 = (v51 << ((1) & 63))
    if (1 != 374) {
        v3 = (0.1 - (((-2.25 / v3) / ((((v3 / 1e10) + ((float64(-2.25) + v67) * v67)) / v67) - ((v67 - float64(float32((v67 - 1.5)))) - ((v67 + 0.1) - 3.14159)))) / -2.25))
        var v69 string = "a"
        _ = v69
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        if v2 {
            var v71 uint32 = 2
            var v72 string = ""
            _ = v71
            _ = v72
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
