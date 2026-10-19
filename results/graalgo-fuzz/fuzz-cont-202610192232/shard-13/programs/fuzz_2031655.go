package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32; F1 int8; F2 int }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint8; F1 int16 }
func genFn0(p0 int16) int8 {
    if (!(((!(42 >= -388)) || false) || false)) {
        if true {
            switch 55 {
            case 32768:
                v6 := GenS1{F0: 42, F1: 255}
                _ = v6
            default:
                for fi7 := 0; fi7 < 2; fi7++ {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        var v9 uint16 = 255
                        if (763 == (v9 >> ((0) & 63))) {
                            v10 := []int{fi8, (fi7 & 1), -1}
                            p0 += (p0 << ((1) & 63))
                            p0 += (p0 % (((p0 | p0)) | 1))
                            if len(v10) > 0 {
                                v10[0] = (fi7 * (fi7 / ((fi8) | 1)))
                            }
                            _ = v10
                        }
                        _ = fi8
                        _ = v9
                    }
                    v11 := []int{fi7, 32767, fi7, fi7}
                    var v12 uint8 = 255
                    _ = fi7
                    _ = v11
                    _ = v12
                }
            }
            switch 128 {
            case 704:
                p0 = p0
            case 32768:
                var v13 float64 = float64(3.14159)
                _ = v13
            case -797:
                var v14 string = ""
                _ = v14
            default:
                fmt.Println(p0)
            }
            p0 += (p0 << ((255) & 63))
        } else {
            if false {
                v15 := GenS2{F0: 801, F1: 601}
                _ = v15
            } else {
                v16 := GenS2{F0: 65535, F1: 2}
                _ = v16
            }
            var v17 int32 = 2147483647
            _ = v17
        }
    }
    var v18 int8 = 42
    var v19 uint = 255
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            fmt.Println(v19)
        } else {
            v18 = 0
            v21 := GenS2{F0: 389, F1: 65535}
            _ = v21
        }
        v22 := []int{fi20, fi20, fi20, fi20}
        switch ("foo bar" + "a") {
        case "a":
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(fi20, v19, fi23)
                _ = fi23
            }
        case "":
            v19 = v19
        case "hello":
            v22 = append(v22, (fi20 - (fi20 << ((v19) & 63))))
        default:
            v18 += (v18 | 0)
        }
        _ = fi20
        _ = v22
    }
    _ = p0
    _ = v18
    _ = v19
    return v18
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]string)
    v3 := []int{4294967295}
    v4 := []int{-308}
    v5 := []string{("a" + ("foo bar" + "\n\t"))}
    v24 := GenS3{F0: -116, F1: -84, F2: 32767}
    v2[-9223372036854775808] = ("foo bar" + ("\n\t" + ("hello" + ("foo bar" + "a"))))
    fmt.Println(v1)
    v25 := GenS3{F0: 2147483647, F1: genFn0(1000), F2: 32768}
    if false {
        switch "foo bar" {
        case "foo bar":
            switch -2147483648 {
            case -605:
                if (!(-563 >= -279)) {
                    v26 := GenS4{F0: 0, F1: 559}
                    _ = v26
                }
            case 4294967295:
                for fi27 := 0; fi27 < 16; fi27++ {
                    if (42 >= 0) {
                        var v28 int32 = 0
                        var v29 uint = 2
                        _ = v28
                        _ = v29
                    } else {
                        if true {
                            v1 += (v1 / ((v1) | 1))
                        } else {
                            if true {
                                switch 2147483647 {
                                case -108:
                                    v1 = v1
                                case -414:
                                    var v30 uint16 = 42
                                    _ = v30
                                }
                            }
                            switch "hello" {
                            case "hello":
                                switch "hello" {
                                case "hello":
                                    if (!(384 == -49)) {
                                        v31 := v2[fi27]
                                        var v32 int16 = (v1 << ((529) & 63))
                                        v1 += v1
                                        _ = v31
                                        _ = v32
                                    } else {
                                        v33 := GenS3{F0: int32(v1), F1: 127, F2: (fi27 << ((0) & 63))}
                                        _ = v33
                                    }
                                case "foo bar":
                                    var v34 float32 = 3.14159
                                    _ = v34
                                }
                            case "":
                                var v35 bool = true
                                _ = v35
                            case "\n\t":
                                var v36 int64 = -498
                                _ = v36
                            }
                            var v37 int = -70
                            _ = v37
                        }
                        v38 := GenS3{F0: 42, F1: -1, F2: (fi27 / ((fi27) | 1))}
                        _ = v38
                    }
                    var v39 int32 = 1
                    var v40 float32 = (1e10 / (3.14159 - 1e10))
                    _ = fi27
                    _ = v39
                    _ = v40
                }
            }
        case "\n\t":
            v41 := GenS3{F0: 0, F1: 100, F2: 0}
            _ = v41
        }
    } else {
        v42 := GenS3{F0: -2147483648, F1: -96, F2: -60}
        v1 = -128
        var v43 uint16 = 0
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        v1 = v1
        if false {
            v1 = (^(v1))
            v45 := GenS5{F0: 774}
            _ = v45
        }
        v46 := v2[fi44]
        _ = fi44
        _ = v46
    }
    v1 = v1
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 uint32 = 268
        _ = fi47
        _ = v48
    }
    fmt.Println(v1)
    v1 = v1
    for fi49 := 0; fi49 < 7; fi49++ {
        _ = len(v2)
        v1 += 32767
        if true {
            if (fi49 == 2147483647) {
                fmt.Println(fi49, v1)
            }
            v50 := GenS2{F0: 2, F1: 65535}
            var v51 bool = (-2147483648 != -50)
            _ = v50
            _ = v51
        }
        _ = fi49
    }
    v52 := GenS6{F0: "\n\t", F1: ""}
    switch 127 {
    case 32768:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -2147483648:
        fmt.Println(v1)
    default:
        v53 := v2[-840]
        _ = v53
    }
    if (((3.14159 * 0.1) <= (float32(1.5) - float32(float64((3.14159 * 0.0))))) && false) {
        if false {
            v1 -= 32767
        }
        switch 1 {
        case 32767:
            v54 := GenS7{F0: 83, F1: v1}
            _ = v54
        default:
            fmt.Println(v1)
        }
    } else {
        if ((2 < 255) && (2 < 794)) {
            fmt.Println(v1)
            fmt.Println(v1)
            var v55 float64 = 3.14159
            _ = v55
        }
        var v56 string = ("a" + "foo bar")
        _ = v56
    }
    var v57 string = "a"
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
}
