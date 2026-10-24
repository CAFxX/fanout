package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 int8, p1 uint32, p2 int8) uint64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            if true {
                v6 := GenS1{F0: fi5}
                p2 += (p2 + (p2 + (^(p0))))
                _ = v6
            }
            var v7 int64 = -780
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: fi8}
        v10 := GenS2{F0: ("foo bar" + "hello")}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    switch ("a" + (("foo bar" + "\n\t") + "foo bar")) {
    case "":
        v12 := GenS1{F0: -270}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 float64 = (1.5 * (0.1 * 0.0))
            if (true && false) {
                fmt.Println(v11, p2)
                var v15 uint = 1
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            switch "hello" {
            case "\n\t":
                if (916 <= v11) {
                    v17 := GenS1{F0: 128}
                    var v18 int = (fi16 & -707)
                    v19 := []int{(fi16 & -684), 543, int(p1), -719}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        p2 = p0
                        switch (-(fi16)) {
                        case 1:
                            fmt.Println(p2, p0, fi16)
                        case 922:
                            fmt.Println(p0, fi20)
                        }
                        p1 += (p1 * p1)
                        var v21 int16 = -32768
                        _ = fi20
                        _ = v21
                    }
                    var v22 uint = 2
                    _ = v22
                }
            case "":
                v23 := GenS1{F0: 1}
                _ = v23
            default:
                switch (fi16 / (((fi16 % (((-(fi16))) | 1))) | 1)) {
                case 128:
                    switch (("hello" + "") + "hello") {
                    case "hello":
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v25 := GenS2{F0: ""}
                            var v26 int8 = -39
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        p2 = (p0 % (((p2 - (p0 / ((p0) | 1)))) | 1))
                    }
                case 4294967295:
                    if false {
                        var v27 int64 = 1
                        var v28 float64 = -2.25
                        _ = v27
                        _ = v28
                    }
                }
            }
            var v29 int32 = -31
            _ = fi16
            _ = v29
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (v11 + (v11 * (v11 * 0)))
}
func genFn1(p0 int, p1 string, p2 uint64) uint16 {
    var v30 uint16 = 42
    for fi31 := 0; fi31 < 19; fi31++ {
        if (255 < 1) {
            var v32 int8 = 127
            p0 -= fi31
            _ = v32
        }
        if true {
            for fi33 := 0; fi33 < 10; fi33++ {
                v34 := []int{fi31}
                fmt.Println(fi31)
                _ = fi33
                _ = v34
            }
            if true {
                v30 -= (^(v30))
                p0 -= (p0 + (p0 - fi31))
                var v35 uint8 = 1
                _ = v35
            } else {
                var v36 uint8 = 127
                var v37 uint16 = 32767
                _ = v36
                _ = v37
            }
            fmt.Println(p2, p1)
        }
        _ = fi31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 0
}

func main() {
    var v1 int64 = 0
    v2 := []string{}
    v3 := make(map[int]int)
    switch -453 {
    case -30:
        for fi38 := 0; fi38 < 16; fi38++ {
            v1 -= (v1 >> ((65535) & 63))
            v1 = v1
            var v39 float32 = 1e10
            _ = fi38
            _ = v39
        }
    case 850:
        fmt.Println(v1)
    default:
        for fi40 := 0; fi40 < 17; fi40++ {
            for fi41 := 0; fi41 < 18; fi41++ {
                v1 -= v1
                v1 += (v1 * -9223372036854775808)
                v42 := GenS3{F0: v1, F1: (("hello" + "") + ("foo bar" + "hello")), F2: 1}
                _ = fi41
                _ = v42
            }
            var v43 float32 = 0.0
            _ = fi40
            _ = v43
        }
    }
    var v44 uint32 = 551
    var v45 int16 = -493
    v46 := GenS4{F0: "a"}
    var v47 float32 = -2.25
    for fi48 := 0; fi48 < 8; fi48++ {
        v49 := GenS3{F0: (v1 | int64(v44)), F1: ("foo bar" + "a"), F2: 255}
        if false {
            delete(v3, fi48)
            v45 = v45
            var v50 int8 = 100
            _ = v50
        } else {
            v45 = (v45 + (v45 ^ (v45 % ((1000) | 1))))
        }
        if true {
            if false {
                v45 = 32767
                var v51 float64 = 0.0
                var v52 uint32 = 42
                _ = v51
                _ = v52
            }
            v53 := GenS1{F0: (fi48 % (((fi48 << ((630) & 63))) | 1))}
            _ = v53
        } else {
            v54 := GenS2{F0: "a"}
            v47 = (v47 * 1.5)
            v45 -= (v45 << ((409) & 63))
            _ = v54
        }
        _ = fi48
        _ = v49
    }
    if (0 != 95) {
        fmt.Println(v45, v47)
    }
    delete(v3, 32767)
    fmt.Println(v1, v44, v45)
    for fi55 := 0; fi55 < 2; fi55++ {
        fmt.Println(v1)
        v56 := GenS2{F0: ("hello" + "a")}
        v57 := GenS1{F0: fi55}
        _ = fi55
        _ = v56
        _ = v57
    }
    switch ((("a" + "\n\t") + (("a" + "\n\t") + (("hello" + "hello") + ""))) + ((("foo bar" + ("" + "hello")) + ((("foo bar" + "hello") + "") + "foo bar")) + "\n\t")) {
    case "foo bar":
        var v58 uint8 = 128
        _ = v58
    case "":
        if (v44 < 2147483647) {
            var v59 int32 = -1
            if true {
                v60 := GenS5{F0: 127, F1: 18446744073709551615}
                if false {
                    if true {
                        var v61 uint = 65535
                        if true {
                            v3[-2147483648] = -1
                            fmt.Println(v59, v45)
                            var v62 int16 = -1
                            _ = v62
                        }
                        _ = v61
                    }
                }
                _ = v60
            }
            v63 := GenS6{F0: ("a" + ""), F1: v45}
            _ = v59
            _ = v63
        }
    case "a":
        for fi64 := 0; fi64 < 4; fi64++ {
            if true {
                var v65 int8 = 1
                v66 := v3[(fi64 / (((fi64 & (fi64 / (((-(fi64))) | 1)))) | 1))]
                if ((!(true || false)) || ((0.0 / (1e10 + 0.1)) != 0.0)) {
                    v67 := GenS4{F0: (("hello" + "hello") + "\n\t")}
                    v3[v66] = (v66 >> ((2) & 63))
                    _ = v67
                } else {
                    v1 = -1
                }
                _ = v65
                _ = v66
            } else {
                var v68 uint32 = v44
                v1 -= v1
                switch fi64 {
                case 127:
                    var v69 string = "hello"
                    _ = v69
                case 0:
                    var v70 uint8 = 239
                    _ = v70
                case 910:
                    _ = len(v2)
                default:
                    v71 := GenS6{F0: ("foo bar" + "a"), F1: (v45 & (v45 + v45))}
                    _ = v71
                }
                _ = v68
            }
            var v72 string = "foo bar"
            switch "foo bar" {
            case "\n\t":
                if ((genFn1(fi64, ("a" + (((v72 + v72) + v72) + v72)), genFn0(-107, v44, -1)) > uint16(v44)) && (((v72 + v72) + "\n\t") < v72)) {
                    v45 += 1
                }
            default:
                v1 = 370
            }
            _ = fi64
            _ = v72
        }
    }
    v47 = (v47 - v47)
    var v73 int = 42
    if (32767 <= v45) {
        if (0 > 32767) {
            var v74 float64 = (3.14159 + (3.14159 + float64((v47 / 1.0))))
            _ = v74
        }
        fmt.Println(v44)
        v75 := v3[v73]
        _ = v75
    } else {
        switch ("foo bar" + "a") {
        case "\n\t":
            switch v73 {
            case 1:
                var v76 int32 = 2147483647
                _ = v76
            case -9223372036854775808:
                switch v73 {
                case 4294967295:
                    delete(v3, v73)
                }
            }
        case "foo bar":
            var v77 uint8 = 151
            _ = v77
        case "hello":
            var v78 string = "\n\t"
            _ = v78
        }
        if true {
            if true {
                var v79 int64 = (v1 >> ((4294967295) & 63))
                _ = v79
            } else {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
        } else {
            if (42 == 42) {
                v73 += v73
            } else {
                var v80 float64 = (3.14159 + (0.0 * 0.1))
                fmt.Println(v73, v1, v1)
                for fi81 := 0; fi81 < 19; fi81++ {
                    for fi82 := 0; fi82 < 5; fi82++ {
                        var v83 int = (fi82 >> ((65535) & 63))
                        if (!false) {
                            var v84 uint16 = 1
                            v45 -= (v45 + 620)
                            _ = v84
                        } else {
                            var v85 uint16 = 2
                            var v86 int8 = 42
                            var v87 float64 = float64((3.14159 * (v47 * v47)))
                            _ = v85
                            _ = v86
                            _ = v87
                        }
                        var v88 bool = (4294967295 >= genFn0(127, (v44 % ((65535) | 1)), 42))
                        v1 += v1
                        _ = fi82
                        _ = v83
                        _ = v88
                    }
                    _ = fi81
                }
                _ = v80
            }
        }
        v89 := GenS7{F0: v73}
        _ = v89
    }
    var v90 uint8 = 96
    v91 := GenS3{F0: v1, F1: ("a" + ""), F2: 42}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v73:", v73)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v73
    _ = v90
    _ = v91
}
