package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0(p0 uint16) int {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: 42, F1: 42}
        var v6 uint64 = 2
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 = 0
    var v7 int16 = -98
    var v8 uint16 = (p0 % ((p0) | 1))
    _ = p0
    _ = v7
    _ = v8
    return int(v7)
}
func genFn1(p0 int16, p1 int32, p2 bool) int64 {
    var v9 uint32 = 158
    var v10 float64 = 0.1
    v10 = v10
    v9 += (v9 << ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 bool = false
    var v2 float64 = (((3.14159 * (0.1 + 1e10)) + 0.0) + 0.1)
    v3 := []uint32{2147483647, 4294967295, 0, 0}
    switch ("" + "foo bar") {
    case "hello":
        var v11 uint8 = 42
        _ = v11
    default:
        v2 -= v2
    }
    _ = len(v3)
    var v12 int32 = 2147483647
    switch "foo bar" {
    case "\n\t":
        switch -769 {
        case -2147483648:
            switch int(v12) {
            case 775:
                for fi13 := 0; fi13 < 14; fi13++ {
                    switch fi13 {
                    case 32767:
                        for fi14 := 0; fi14 < 10; fi14++ {
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            var v15 uint = 65535
                            var v16 int32 = int32(fi14)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    default:
                        _ = len(v3)
                    }
                    switch ("hello" + "a") {
                    case "\n\t":
                        fmt.Println(v1, v12, v1)
                    default:
                        if len(v3) > 0 {
                            v3[0] = 660
                        }
                    }
                    v12 += v12
                    _ = fi13
                }
            case -175:
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 uint = 96
                    var v19 int32 = v12
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint16 = 1
                    if (!(v1 && false)) {
                        switch (fi20 - -1) {
                        case 1:
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v23 := GenS2{F0: 1}
                                fmt.Println(fi22, v2)
                                _ = fi22
                                _ = v23
                            }
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            var v25 float32 = float32(-2.25)
                            var v26 int16 = 1
                            v26 += 42
                            var v27 int = fi20
                            var v28 int16 = (-(v26))
                            v29 := GenS2{F0: 102}
                            switch ("foo bar" + ((("hello" + "a") + "") + "\n\t")) {
                            case "hello":
                                _ = len(v3)
                            default:
                                v30 := GenS3{F0: -128, F1: true}
                                _ = v30
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        var v31 bool = (2 <= 77)
                        _ = v31
                    }
                    _ = fi20
                    _ = v21
                }
            }
        }
    case "foo bar":
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 int64 = -9223372036854775808
            _ = fi32
            _ = v33
        }
    case "hello":
        if true {
            switch "foo bar" {
            case "hello":
                v2 = v2
            case "a":
                if (-9223372036854775808 < 4294967295) {
                    var v34 uint = 18446744073709551615
                    var v35 int8 = 1
                    v3 = append(v3, 291)
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(v12, v2)
                    switch "a" {
                    case "foo bar":
                        if (310 > 42) {
                            switch genFn0(702) {
                            case 32768:
                                var v36 uint = 4294967295
                                _ = v36
                            default:
                                var v37 int32 = 1
                                _ = v37
                            }
                        } else {
                            v38 := GenS2{F0: 1}
                            v1 = v1
                            v12 += -2147483648
                            _ = v38
                        }
                    default:
                        if v1 {
                            v39 := GenS4{F0: 42}
                            _ = v39
                        } else {
                            v3 = append(v3, 96)
                        }
                    }
                }
            }
        }
    }
    fmt.Println(v12)
    v1 = false
    v2 += -2.25
    fmt.Println(v1, v12, v1)
    for fi40 := 0; fi40 < 13; fi40++ {
        var v41 int16 = 32767
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        v1 = (279 >= 4294967295)
        for fi43 := 0; fi43 < 20; fi43++ {
            var v44 string = "foo bar"
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    var v45 float32 = 0.1
    if ((-(v12)) > v12) {
        for fi46 := 0; fi46 < 6; fi46++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (4294967295 > 42) {
                for fi47 := 0; fi47 < 12; fi47++ {
                    _ = len(v3)
                    _ = fi47
                }
                if v1 {
                    v48 := GenS3{F0: 1, F1: (228 != 580)}
                    _ = v48
                }
                var v49 int16 = 1
                _ = v49
            }
            _ = fi46
        }
        v50 := GenS1{F0: 42, F1: 2}
        _ = v50
    }
    for fi51 := 0; fi51 < 2; fi51++ {
        switch (("hello" + "a") + ("a" + ("foo bar" + ("foo bar" + "a")))) {
        case "":
            switch "foo bar" {
            case "":
                v52 := GenS5{F0: 0, F1: 2, F2: 255}
                _ = v52
            case "foo bar":
                v2 = v2
            default:
                if true {
                    switch -740 {
                    case 0:
                        var v53 int8 = 100
                        _ = v53
                    case 2147483647:
                        var v54 bool = (2 >= 65535)
                        _ = v54
                    default:
                        v1 = (32767 <= genFn0(1))
                    }
                    for fi55 := 0; fi55 < 11; fi55++ {
                        if v1 {
                            v45 += ((1e10 / 0.1) / (0.1 * v45))
                            if ((42 == 0) && (1e10 < (v2 - 1e10))) {
                                var v56 uint64 = 1
                                if (-1 <= (v12 - (v12 ^ 1))) {
                                    var v57 int = genFn0(0)
                                    if ((963 == 4294967295) && ((!(!v1)) || v1)) {
                                        fmt.Println(v57)
                                    } else {
                                        var v58 uint = 42
                                        var v59 string = ("hello" + "\n\t")
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v57
                                }
                                _ = v56
                            } else {
                                if (v1 && (726 == 673)) {
                                    v1 = v1
                                    v12 -= v12
                                } else {
                                    fmt.Println(v45, v45, v1)
                                    var v60 uint64 = 2
                                    if (v1 || (v1 && v1)) {
                                        v61 := GenS1{F0: 32767, F1: 65535}
                                        if (-986 > (v12 + v12)) {
                                            if (60 <= 0) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            fmt.Println(v60, v1)
                                            if v1 {
                                                var v62 float64 = float64(-2.25)
                                                v63 := GenS6{F0: 4294967295, F1: 128, F2: 106}
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                        _ = v61
                                    }
                                    _ = v60
                                }
                            }
                        }
                        fmt.Println(v2, fi55)
                        v1 = (100 < 42)
                        _ = fi55
                    }
                    fmt.Println(v1)
                } else {
                    if v1 {
                        v64 := GenS1{F0: 32767, F1: 42}
                        _ = v64
                    }
                }
            }
        case "a":
            v65 := GenS3{F0: -120, F1: v1}
            _ = v65
        case "foo bar":
            var v66 uint16 = 0
            _ = v66
        default:
            v67 := GenS5{F0: 127, F1: 255, F2: 42}
            _ = v67
        }
        _ = fi51
    }
    v3 = append(v3, 0)
    var v68 float64 = (float64(v45) - 1.5)
    var v69 int = 32767
    var v70 int16 = 42
    v45 -= v45
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v45
    _ = v68
    _ = v69
    _ = v70
}
