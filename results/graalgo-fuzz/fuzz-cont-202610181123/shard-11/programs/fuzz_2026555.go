package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 string; F1 uint8; F2 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 float64 = (1e10 + float64((0.1 / 1.0)))
    var v2 float32 = 0.1
    var v3 int8 = -65
    var v4 uint16 = 569
    v2 = v2
    switch 32768 {
    case 128:
        v5 := []int{42, 0}
        _ = v5
    case -9223372036854775808:
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: v3}
            _ = fi6
            _ = v7
        }
    case 960:
        v4 = (^(v4))
    }
    if true {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: v3}
            for fi10 := 0; fi10 < 8; fi10++ {
                switch (fi8 * fi10) {
                case 32767:
                    var v11 int = fi8
                    _ = v11
                case -2147483648:
                    v12 := GenS2{F0: 127, F1: (v4 << ((112) & 63)), F2: 675}
                    _ = v12
                }
                switch fi8 {
                case -1:
                    v1 = (v1 - 0.0)
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    } else {
        v13 := GenS3{F0: -9223372036854775808, F1: 19}
        var v14 bool = (!false)
        var v15 uint32 = 0
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 bool = (("\n\t" + ("\n\t" + ("hello" + ("\n\t" + "foo bar")))) >= "a")
    if (false && (1 >= -657)) {
        switch "hello" {
        case "":
            v17 := []int{-9223372036854775808}
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 12; fi18++ {
                if v16 {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        if (v3 != (-(v3))) {
                            if true {
                                v1 += v1
                                v20 := []int{fi18, -1, (fi19 + (fi19 / (((fi19 | fi18)) | 1))), fi19}
                                _ = v20
                            } else {
                                fmt.Println(v4)
                                var v21 uint32 = 775
                                v2 += float32(1e10)
                                _ = v21
                            }
                            if v16 {
                                v22 := []int{(fi19 % (((fi19 >> ((382) & 63))) | 1)), fi18, fi18, fi18}
                                switch fi19 {
                                case -302:
                                    var v23 bool = false
                                    _ = v23
                                case 128:
                                    v24 := GenS3{F0: fi19, F1: 157}
                                    _ = v24
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v25 float64 = (float64(0.1) * (v1 + v1))
                                    _ = v25
                                case "a":
                                    fmt.Println(fi19, fi19, v1)
                                case "foo bar":
                                    if (0 != 397) {
                                        v4 += v4
                                        v26 := GenS2{F0: 128, F1: 42, F2: 1}
                                        _ = v26
                                    } else {
                                        switch fi19 {
                                        case 128:
                                            fmt.Println(v2)
                                        case -2147483648:
                                            v1 = float64(3.14159)
                                            if (v16 || (v16 && (1 >= 18446744073709551615))) {
                                                v27 := GenS4{F0: fi18, F1: v16}
                                                v28 := GenS2{F0: 42, F1: v4, F2: 1}
                                                _ = v27
                                                _ = v28
                                            }
                                            fmt.Println(fi18, v2)
                                            var v29 int16 = 1000
                                            _ = v29
                                        default:
                                            if (128 <= 42) {
                                                if v16 {
                                                    v16 = v16
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                    _ = len(v22)
                                                }
                                                v3 = v3
                                            }
                                        }
                                    }
                                default:
                                    if (133 != 359) {
                                        switch "" {
                                        case "hello":
                                            v16 = (0 > 1)
                                        case "":
                                            if (925 != 2147483647) {
                                                v30 := GenS3{F0: int(v4), F1: 182}
                                                var v31 float32 = (v2 - v2)
                                                var v32 uint64 = 0
                                                var v33 uint8 = 127
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        var v34 uint64 = 807
                                        _ = v34
                                    }
                                }
                                _ = v22
                            } else {
                                if true {
                                    v35 := GenS1{F0: (v3 + 127)}
                                    var v36 uint32 = 2147483648
                                    _ = v35
                                    _ = v36
                                }
                                var v37 int32 = -2147483648
                                _ = v37
                            }
                        } else {
                            v38 := GenS4{F0: (fi19 | (fi19 ^ (fi18 - 42))), F1: (!true)}
                            v4 -= 830
                            _ = v38
                        }
                        _ = fi19
                    }
                }
                v2 = (3.14159 / -2.25)
                _ = fi18
            }
        default:
            v39 := GenS5{F0: 1, F1: -2147483648}
            _ = v39
        }
        v2 += 1.5
    }
    if (255 > 2) {
        for fi40 := 0; fi40 < 10; fi40++ {
            if ((32767 != -32768) && v16) {
                v2 = (v2 * (((((v2 - ((v2 * v2) / v2)) / v2) * 0.1) * v2) - v2))
            }
            fmt.Println(v4, v1)
            var v41 float32 = (v2 / v2)
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v3)
    if v16 {
        switch -9223372036854775808 {
        case 42:
            if (0 <= 0) {
                v42 := GenS2{F0: 2, F1: (v4 * v4), F2: 1}
                _ = v42
            }
        case 9223372036854775807:
            for fi43 := 0; fi43 < 1; fi43++ {
                for fi44 := 0; fi44 < 2; fi44++ {
                    v3 = 42
                    _ = fi44
                }
                _ = fi43
            }
        default:
            for fi45 := 0; fi45 < 17; fi45++ {
                v46 := GenS6{F0: "", F1: 0, F2: 65535}
                var v47 int32 = -1
                _ = fi45
                _ = v46
                _ = v47
            }
        }
        v2 += (v2 / v2)
    } else {
        var v48 uint8 = 127
        if v16 {
            if ("a" >= ("\n\t" + "foo bar")) {
                for fi49 := 0; fi49 < 20; fi49++ {
                    v50 := []int{}
                    if false {
                        v51 := GenS7{F0: -1}
                        v48 -= (v48 % ((255) | 1))
                        _ = v51
                    }
                    v3 -= 1
                    _ = fi49
                    _ = v50
                }
                if ((float64(1.5) / v1) >= (v1 - (v1 / v1))) {
                    var v52 int = 1
                    _ = v52
                } else {
                    v2 -= v2
                }
            } else {
                var v53 uint16 = 0
                if false {
                    fmt.Println(v2, v4, v53)
                    fmt.Println(v3)
                }
                _ = v53
            }
            v54 := []int{-733, 32767}
            for fi55 := 0; fi55 < 14; fi55++ {
                if ((124 > 42) || (v3 != (^(v3)))) {
                    v1 += float64(float32((float64((3.14159 * v2)) - (v1 * 1e10))))
                    for fi56 := 0; fi56 < 6; fi56++ {
                        v48 = v48
                        var v57 int16 = 1000
                        var v58 int16 = v57
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                } else {
                    if (true || v16) {
                        for fi59 := 0; fi59 < 2; fi59++ {
                            fmt.Println(v4, v16)
                            v60 := GenS8{F0: v16, F1: 2, F2: v3}
                            if v16 {
                                v61 := GenS9{F0: v4, F1: "a"}
                                var v62 uint = 2
                                v63 := GenS10{F0: (v48 ^ v48), F1: 877}
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                            _ = fi59
                            _ = v60
                        }
                        var v64 string = (("a" + "foo bar") + "a")
                        if len(v54) >= 2 {
                            v54 = v54[1:len(v54)]
                        }
                        _ = v64
                    }
                    if (!v16) {
                        v3 += (v3 - -57)
                    }
                    var v65 float64 = v1
                    _ = v65
                }
                _ = fi55
            }
            _ = v54
        }
        _ = v48
    }
    fmt.Println(v3, v16, v1)
    var v66 uint32 = 968
    v1 += v1
    var v67 int16 = -875
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
}
