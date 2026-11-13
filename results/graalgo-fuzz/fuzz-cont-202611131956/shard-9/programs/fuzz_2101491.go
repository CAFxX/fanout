package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 bool; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint; F1 int64; F2 uint }
type GenS8 struct { F0 uint; F1 uint; F2 bool }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint16, p2 int64) uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        switch 4294967295 {
        case 2147483647:
            p2 = -9223372036854775808
        default:
            v6 := []int{fi5, (fi5 + (fi5 ^ (fi5 | (fi5 | (fi5 << ((p0) & 63)))))), (fi5 + 661)}
            _ = v6
        }
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int64 = 336
            switch fi5 {
            case -9223372036854775808:
                p2 = (p2 >> ((p0) & 63))
            case 127:
                fmt.Println(p2, v8, fi7)
            default:
                var v9 float32 = 3.14159
                _ = v9
            }
            switch "" {
            case "":
                fmt.Println(fi7, v8)
            default:
                fmt.Println(p1, p1, fi7)
            }
            _ = fi7
            _ = v8
        }
        v10 := GenS1{F0: 65535, F1: 4, F2: "foo bar"}
        _ = fi5
        _ = v10
    }
    switch -1 {
    case 127:
        switch -545 {
        case 1:
            switch -859 {
            case 2147483647:
                var v11 string = "hello"
                _ = v11
            case 9223372036854775807:
                p0 = (p0 << ((913) & 63))
            case 553:
                var v12 bool = true
                _ = v12
            default:
                if false {
                    p2 += p2
                }
            }
        case 42:
            var v13 float64 = -2.25
            _ = v13
        case 32768:
            var v14 uint64 = 938
            _ = v14
        }
    case 885:
        for fi15 := 0; fi15 < 3; fi15++ {
            if (int32(p0) > 42) {
                v16 := GenS1{F0: (p1 & p1), F1: 42, F2: ""}
                _ = v16
            } else {
                fmt.Println(fi15, fi15, fi15)
                p0 -= (p0 ^ 255)
            }
            _ = fi15
        }
    default:
        v17 := GenS2{F0: -506, F1: "\n\t"}
        _ = v17
    }
    v18 := GenS3{F0: 112}
    switch ("" + "\n\t") {
    case "\n\t":
        fmt.Println(p1, p1)
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (p1 >> (((^(p0))) & 63))
}
func genFn1(p0 int, p1 string) int64 {
    p0 = 0
    if (42 < 65535) {
        for fi19 := 0; fi19 < 14; fi19++ {
            p0 += p0
            _ = fi19
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 42
    var v3 uint = v1
    v4 := []int{-9223372036854775808, 4294967295, 0, -828}
    if true {
        v1 = 65535
        v20 := GenS2{F0: 9223372036854775807, F1: "a"}
        v4 = append(v4, 0)
        _ = v20
    }
    v3 += (v1 ^ 42)
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 = (v3 << (((v3 ^ v1)) & 63))
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int8 = 100
            _ = fi22
            _ = v23
        }
        if false {
            v24 := GenS2{F0: 0, F1: ((("" + ("a" + "foo bar")) + "a") + (("a" + "hello") + "foo bar"))}
            _ = v24
        } else {
            switch 9223372036854775807 {
            case -2147483648:
                v2 += v2
            case 42:
                if len(v4) > 0 {
                    v4[0] = fi21
                }
            case 133:
                switch "" {
                case "a":
                    for fi25 := 0; fi25 < 20; fi25++ {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            v3 -= 65535
                        case "":
                            if false {
                                v26 := GenS1{F0: 65535, F1: -128, F2: ""}
                                _ = v26
                            } else {
                                if (!true) {
                                    v27 := GenS4{F0: false, F1: (true && (!false)), F2: (^(fi25))}
                                    v28 := GenS2{F0: genFn1((fi25 & fi21), "a"), F1: ""}
                                    v3 += (v1 % (((v3 | (v1 % ((v3) | 1)))) | 1))
                                    v3 += (v3 | 42)
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 int64 = 0
                                    if (false || false) {
                                        v29 = genFn1(fi21, "hello")
                                        fmt.Println(v1, v2, v2)
                                        if false {
                                            var v30 uint = v1
                                            v31 := GenS5{F0: (1 <= 100)}
                                            fmt.Println(v29, fi25)
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    switch (^(fi25)) {
                                    case 0:
                                        v32 := GenS5{F0: ((^(v2)) > uint64(v3))}
                                        fmt.Println(v2, fi21, v2)
                                        _ = v32
                                    default:
                                        if (-670 <= 1000) {
                                            if len(v4) > 0 {
                                                v4[0] = (fi21 - fi21)
                                            }
                                            v2 = 18446744073709551615
                                            v2 = 865
                                        } else {
                                            if true {
                                                v3 += 108
                                                v33 := GenS2{F0: genFn1((fi21 << ((65535) & 63)), "a"), F1: ("\n\t" + "hello")}
                                                _ = v33
                                            }
                                            switch (fi21 << (((v3 + 4294967295)) & 63)) {
                                            case -9223372036854775808:
                                                if (!false) {
                                                    v29 = (v29 - -563)
                                                } else {
                                                    fmt.Println(fi21, fi25, v29)
                                                    var v34 uint32 = 2147483648
                                                    _ = v34
                                                }
                                            case 127:
                                                var v35 float32 = 0.0
                                                _ = v35
                                            case 32767:
                                                var v36 bool = (964 <= 42)
                                                _ = v36
                                            }
                                            var v37 int16 = 1000
                                            switch "" {
                                            case "\n\t":
                                                v38 := GenS2{F0: v29, F1: (("foo bar" + ("" + "hello")) + "foo bar")}
                                                _ = v38
                                            case "foo bar":
                                                _ = len(v4)
                                            case "":
                                                if false {
                                                    v4 = append(v4, fi25)
                                                } else {
                                                    _ = len(v4)
                                                    var v39 bool = (v2 < v2)
                                                    _ = v39
                                                }
                                            }
                                            _ = v37
                                        }
                                        v4 = append(v4, fi25)
                                    }
                                    v40 := GenS1{F0: 255, F1: -23, F2: ("\n\t" + "hello")}
                                    var v41 uint16 = 1
                                    var v42 int16 = 1
                                    _ = v29
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                        v3 = 18446744073709551615
                        v3 = (v1 & (v1 / ((v3) | 1)))
                        _ = fi25
                    }
                case "":
                    _ = len(v4)
                }
            default:
                if (true && (2147483648 != 0)) {
                    _ = len(v4)
                    var v43 uint32 = 0
                    _ = v43
                }
            }
        }
        _ = fi21
    }
    switch ((("foo bar" + ("hello" + "a")) + "\n\t") + "") {
    case "foo bar":
        if true {
            v44 := GenS5{F0: ((("hello" + "a") + "\n\t") < ("hello" + ((("foo bar" + "hello") + "hello") + "\n\t")))}
            v1 = v3
            v45 := GenS5{F0: false}
            _ = v44
            _ = v45
        } else {
            var v46 uint64 = 18446744073709551615
            for fi47 := 0; fi47 < 10; fi47++ {
                for fi48 := 0; fi48 < 1; fi48++ {
                    if (((1.5 + (0.1 + 0.1)) != (0.1 / 1.5)) && (1 != -1)) {
                        var v49 float32 = (0.1 / 1.0)
                        _ = v49
                    } else {
                        v50 := GenS6{F0: false, F1: 1}
                        _ = v50
                    }
                    var v51 uint = v3
                    v2 = v2
                    _ = fi48
                    _ = v51
                }
                var v52 float32 = 0.1
                _ = fi47
                _ = v52
            }
            if false {
                v53 := GenS1{F0: 65535, F1: 127, F2: ""}
                var v54 int32 = -744
                _ = v53
                _ = v54
            }
            _ = v46
        }
    default:
        for fi55 := 0; fi55 < 13; fi55++ {
            v1 = (v1 / (((v3 + v3)) | 1))
            _ = fi55
        }
    }
    v1 = v3
    v1 += v3
    fmt.Println(v1, v3)
    if true {
        switch 700 {
        case -135:
            fmt.Println(v1, v1, v1)
        case 658:
            for fi56 := 0; fi56 < 5; fi56++ {
                switch fi56 {
                case 128:
                    if (518 >= 0) {
                        var v57 int16 = int16(fi56)
                        _ = v57
                    }
                case -1:
                    var v58 int16 = 1
                    _ = v58
                case -9223372036854775808:
                    var v59 string = "\n\t"
                    _ = v59
                default:
                    v60 := GenS7{F0: (^(v1)), F1: 0, F2: 65535}
                    _ = v60
                }
                v61 := GenS8{F0: 4294967295, F1: v1, F2: false}
                _ = fi56
                _ = v61
            }
        case -911:
            fmt.Println(v2)
        }
        var v62 uint64 = (^(v2))
        _ = v62
    } else {
        fmt.Println(v3, v3, v3)
        v3 -= 537
    }
    if (float64((float32(0.1) * (float32((3.14159 * 1.5)) / (1e10 / 3.14159)))) == 0.0) {
        for fi63 := 0; fi63 < 11; fi63++ {
            v4 = append(v4, fi63)
            for fi64 := 0; fi64 < 8; fi64++ {
                v3 += (v1 & (v1 + v1))
                _ = len(v4)
                _ = fi64
            }
            _ = fi63
        }
        v1 = v1
        if (32767 > -2147483648) {
            v2 -= (v2 ^ v2)
            var v65 string = ("hello" + ("foo bar" + "foo bar"))
            v2 = 1
            _ = v65
        } else {
            if ((genFn0(v3, genFn0(2, genFn0((v3 + 255), 1, -917), -649), -1) != 432) || (!false)) {
                v66 := GenS4{F0: ((v3 & (v1 << (((^(v1))) & 63))) > 42), F1: false, F2: -348}
                _ = v66
            } else {
                fmt.Println(v3, v1)
                v2 = v2
                var v67 uint8 = 128
                _ = v67
            }
            v68 := GenS2{F0: 0, F1: "a"}
            v4 = append(v4, 42)
            _ = v68
        }
    } else {
        v2 = v2
        v69 := GenS9{F0: uint16(v2)}
        v70 := GenS4{F0: true, F1: ("hello" >= (("hello" + ("" + "\n\t")) + "")), F2: -2147483648}
        _ = v69
        _ = v70
    }
    fmt.Println(v2, v2)
    v71 := GenS3{F0: 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
}
