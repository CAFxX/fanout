package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool, p1 int64, p2 uint8) int {
    v5 := []int{128, 2147483647}
    var v6 uint8 = (p2 ^ (^(p2)))
    var v7 int32 = -2147483648
    v8 := GenS1{F0: v7, F1: 42}
    if ((!p0) || (60 != 2)) {
        p1 -= p1
        var v9 uint32 = uint32(v6)
        switch "hello" {
        case "a":
            switch ("" + (("hello" + ("\n\t" + "")) + "\n\t")) {
            case "a":
                v5 = append(v5, -2147483648)
            case "":
                fmt.Println(v9)
            case "foo bar":
                if false {
                    v5 = append(v5, 32768)
                    v10 := GenS2{F0: 4294967295, F1: 1}
                    _ = v10
                }
            default:
                v6 = p2
            }
        case "foo bar":
            for fi11 := 0; fi11 < 11; fi11++ {
                _ = len(v5)
                fmt.Println(p0, v6, v7)
                for fi12 := 0; fi12 < 13; fi12++ {
                    _ = len(v5)
                    _ = fi12
                }
                _ = fi11
            }
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1() int32 {
    var v13 float32 = 3.14159
    switch "foo bar" {
    case "hello":
        v14 := []int{32768}
        _ = v14
    default:
        var v15 uint32 = 4294967295
        _ = v15
    }
    switch "foo bar" {
    case "\n\t":
        if (!(1 != 776)) {
            v16 := []int{}
            _ = v16
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                switch 4294967295 {
                case -58:
                    v18 := GenS1{F0: 0, F1: 1}
                    _ = v18
                default:
                    var v19 int64 = 9223372036854775807
                    _ = v19
                }
                switch fi17 {
                case 9223372036854775807:
                    switch -147 {
                    case -9223372036854775808:
                        switch "hello" {
                        case "\n\t":
                            if (1e10 != 1.5) {
                                switch (fi17 >> ((0) & 63)) {
                                case 42:
                                    var v20 uint = 255
                                    _ = v20
                                default:
                                    v13 = (v13 * v13)
                                }
                                var v21 uint16 = 42
                                _ = v21
                            } else {
                                v13 = v13
                                if false {
                                    var v22 int16 = 0
                                    _ = v22
                                } else {
                                    fmt.Println(fi17, v13)
                                    v13 = v13
                                    for fi23 := 0; fi23 < 4; fi23++ {
                                        var v24 int = 32767
                                        v24 = fi17
                                        v13 -= v13
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                            }
                        case "foo bar":
                            v25 := []int{-143, -2147483648, genFn0(((!true) || (939 >= 0)), 42, 255)}
                            _ = v25
                        default:
                            fmt.Println(fi17, v13)
                        }
                    case 310:
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v13 = v13
                            _ = fi26
                        }
                    default:
                        v27 := GenS3{F0: 42, F1: 216}
                        _ = v27
                    }
                }
                _ = fi17
            }
        }
    case "a":
        v28 := []int{-9223372036854775808, 42, 1}
        _ = v28
    default:
        v13 = 0.1
    }
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (false || true)
    var v3 uint8 = 254
    var v4 float64 = 0.1
    v29 := GenS1{F0: -525, F1: 16}
    v2 = (-8 >= 8)
    if (!false) {
        for fi30 := 0; fi30 < 14; fi30++ {
            for fi31 := 0; fi31 < 6; fi31++ {
                if (280 != 269) {
                    v3 = (v3 * (v3 >> ((4294967295) & 63)))
                    v3 = v3
                    v32 := GenS2{F0: 1, F1: 255}
                    _ = v32
                }
                v33 := GenS2{F0: 1, F1: 32767}
                v3 += v3
                _ = fi31
                _ = v33
            }
            v4 = v4
            if (genFn1() < 42) {
                for fi34 := 0; fi34 < 15; fi34++ {
                    v2 = false
                    _ = fi34
                }
                var v35 uint32 = 978
                _ = v35
            } else {
                v3 = v3
                v36 := []int{genFn0(true, 681, (v3 << ((1) & 63)))}
                fmt.Println(v4, fi30)
                _ = v36
            }
            _ = fi30
        }
    } else {
        switch ("foo bar" + "") {
        case "a":
            for fi37 := 0; fi37 < 15; fi37++ {
                if false {
                    fmt.Println(v4, v4, v4)
                    v38 := GenS3{F0: 4294967295, F1: 0}
                    fmt.Println(v3, v4, v2)
                    _ = v38
                }
                _ = fi37
            }
        }
        var v39 int64 = 42
        switch "hello" {
        case "":
            if true {
                fmt.Println(v3)
                if true {
                    v40 := v1[898]
                    _ = v40
                }
                v4 = v4
            }
        case "foo bar":
            fmt.Println(v4)
        default:
            switch (("" + "a") + "foo bar") {
            case "\n\t":
                var v41 float64 = v4
                _ = v41
            case "a":
                var v42 int32 = genFn1()
                _ = v42
            case "":
                for fi43 := 0; fi43 < 20; fi43++ {
                    v3 = v3
                    switch fi43 {
                    case 4294967295:
                        switch ("hello" + "hello") {
                        case "a":
                            v44 := []int{fi43}
                            _ = v44
                        case "foo bar":
                            v3 = v3
                        case "hello":
                            v45 := []int{}
                            _ = v45
                        }
                    default:
                        fmt.Println(v39, v3, v3)
                    }
                    _ = fi43
                }
            default:
                v2 = (!(float32(v4) <= 1e10))
            }
        }
        _ = v39
    }
    switch (("" + ("\n\t" + "\n\t")) + "foo bar") {
    case "hello":
        var v46 int64 = 204
        _ = v46
    default:
        v3 = (v3 / (((^(v3))) | 1))
    }
    delete(v1, 276)
    for fi47 := 0; fi47 < 18; fi47++ {
        v48 := GenS3{F0: 4294967295, F1: 0}
        if (1000 == -296) {
            v49 := GenS3{F0: 0, F1: 1}
            _ = v49
        } else {
            v50 := GenS3{F0: 42, F1: 506}
            v3 += (v3 >> ((255) & 63))
            if v2 {
                if (("foo bar" + "a") != "hello") {
                    v51 := []int{(fi47 % ((468) | 1))}
                    v51 = append(v51, (fi47 & fi47))
                    _ = v51
                } else {
                    if (!((-2.25 - v4) >= v4)) {
                        v4 += v4
                        v52 := v1[0]
                        _ = v52
                    } else {
                        if (!false) {
                            v53 := GenS1{F0: 0, F1: 0}
                            v2 = ("a" >= "\n\t")
                            v3 = 128
                            _ = v53
                        }
                        if v2 {
                            var v54 int = fi47
                            _ = v54
                        } else {
                            if false {
                                v3 -= (v3 << ((2) & 63))
                            }
                        }
                    }
                }
            }
            _ = v50
        }
        v55 := []int{}
        _ = fi47
        _ = v48
        _ = v55
    }
    if (((v4 - (float64(0.0) + 0.1)) > v4) || (!v2)) {
        switch 9223372036854775807 {
        case 42:
            var v56 bool = (v2 || ("" > "hello"))
            _ = v56
        case 127:
            v3 -= (v3 << ((1) & 63))
        default:
            v57 := []int{4294967295}
            _ = v57
        }
        var v58 uint16 = 42
        _ = v58
    } else {
        v59 := GenS4{F0: 682, F1: -5, F2: -324}
        v3 = v3
        v4 -= v4
        _ = v59
    }
    v4 -= v4
    var v60 bool = ((v2 && ((false && (genFn1() > genFn1())) || (v2 && (65535 > 268)))) && v2)
    var v61 uint16 = 971
    if (354 > 0) {
        var v62 int8 = 0
        _ = v62
    } else {
        var v63 uint32 = 817
        for fi64 := 0; fi64 < 3; fi64++ {
            v61 = v61
            v61 = v61
            for fi65 := 0; fi65 < 9; fi65++ {
                var v66 uint = 4294967295
                var v67 float32 = 1e10
                _ = fi65
                _ = v66
                _ = v67
            }
            _ = fi64
        }
        _ = v63
    }
    v68 := GenS5{F0: (v61 / ((v61) | 1))}
    var v69 float32 = 3.14159
    v69 = v69
    _ = len(v1)
    fmt.Println(v4, v60, v60)
    for fi70 := 0; fi70 < 11; fi70++ {
        v2 = v2
        _ = fi70
    }
    if true {
        v4 = 0.0
        v3 -= v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v68
    _ = v69
}
