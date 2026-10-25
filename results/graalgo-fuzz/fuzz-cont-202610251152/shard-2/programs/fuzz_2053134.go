package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 string; F2 int32 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0(p0 uint16, p1 int8) int8 {
    v4 := []int{}
    var v5 int8 = (p1 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}
func genFn1(p0 bool, p1 uint32) uint {
    v6 := GenS1{F0: 165, F1: (p1 + (p1 + (p1 & 65535))), F2: p1}
    p1 = (p1 % ((p1) | 1))
    fmt.Println(p1, p0)
    p1 -= (p1 << ((319) & 63))
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    v1 = v1
    v1 = v1
    v1 = ("a" + (v1 + v1))
    v1 = v1
    for fi7 := 0; fi7 < 2; fi7++ {
        if (!false) {
            v1 = ((("hello" + v1) + ("foo bar" + (v1 + ("\n\t" + v1)))) + (v1 + ""))
        }
        v8 := GenS1{F0: 1, F1: 772, F2: 0}
        _ = fi7
        _ = v8
    }
    var v9 int64 = -682
    if ("a" < ("a" + (v1 + v1))) {
        if (false || false) {
            if true {
                v10 := GenS2{F0: 7, F1: 288, F2: 255}
                delete(v3, v1)
                _ = v10
            }
            var v11 uint8 = 42
            if ((!true) || false) {
                var v12 uint32 = 1
                if true {
                    v13 := GenS2{F0: 42, F1: 109, F2: (v11 * (v11 * v11))}
                    var v14 float32 = 1.5
                    _ = v13
                    _ = v14
                }
                _ = v12
            } else {
                var v15 float32 = 3.14159
                _ = v15
            }
            _ = v11
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                if false {
                    fmt.Println(v9)
                } else {
                    fmt.Println(v9, fi16)
                    v17 := GenS2{F0: 1, F1: 255, F2: 146}
                    fmt.Println(v1, v1)
                    _ = v17
                }
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 int = fi16
                    v19 += fi18
                    var v20 int32 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(fi16, v1, v9)
                _ = fi16
            }
            v21 := GenS1{F0: 18446744073709551615, F1: 42, F2: 4294967295}
            _ = v21
        }
        var v22 float32 = 1.5
        if true {
            v23 := []int{4294967295, 0, -4}
            _ = v23
        } else {
            v24 := []int{632, -378, 32768}
            v25 := GenS3{F0: 42, F1: 42, F2: 269}
            _ = v24
            _ = v25
        }
        _ = v22
    } else {
        fmt.Println(v1, v9)
        v9 = (^(v9))
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 int32 = 951
            var v28 uint8 = 127
            v29 := v2[742]
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    }
    v30 := []int{553}
    v31 := GenS4{F0: 584, F1: "a", F2: -1}
    if false {
        for fi32 := 0; fi32 < 20; fi32++ {
            if (!(265 >= genFn1((255 == 65535), 2147483647))) {
                var v33 int32 = 1
                _ = v33
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 bool = true
                    var v36 int = fi34
                    v37 := GenS3{F0: -1, F1: 202, F2: 4294967295}
                    fmt.Println(fi34, v9)
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                var v38 float64 = 3.14159
                v39 := GenS5{F0: -1, F1: 4294967295}
                _ = v38
                _ = v39
            }
            var v40 int32 = 0
            switch (v1 + "a") {
            case "\n\t":
                var v41 uint32 = 654
                _ = v41
            case "foo bar":
                var v42 int16 = 1
                _ = v42
            case "":
                v40 = (v40 % ((v40) | 1))
            default:
                var v43 float64 = 3.14159
                _ = v43
            }
            _ = fi32
            _ = v40
        }
        if (490 != 2147483647) {
            var v44 int8 = -128
            _ = len(v3)
            _ = v44
        } else {
            var v45 bool = false
            for fi46 := 0; fi46 < 12; fi46++ {
                v1 = ""
                for fi47 := 0; fi47 < 17; fi47++ {
                    if ((v45 && ((((float64((3.14159 * (-2.25 - -2.25))) / (0.1 + (float64(0.1) * -2.25))) - 1.5) <= 3.14159) && (true && v45))) && false) {
                        var v48 string = ("foo bar" + ("a" + ("hello" + (v1 + (v1 + ("hello" + ((((v1 + ("foo bar" + v1)) + v1) + "foo bar") + "")))))))
                        v9 = 0
                        if ((0 > 181) || v45) {
                            switch v1 {
                            case "\n\t":
                                v48 = v48
                            }
                            if (!v45) {
                                var v49 uint64 = 42
                                _ = v49
                            } else {
                                delete(v2, 973)
                            }
                        } else {
                            v45 = (42 < 2147483647)
                            var v50 float32 = 0.1
                            _ = v50
                        }
                        _ = v48
                    } else {
                        var v51 bool = ((fi47 << ((255) & 63)) <= (fi47 * (fi46 + 908)))
                        v52 := GenS6{F0: -895, F1: v1}
                        v53 := GenS1{F0: 2, F1: 683, F2: 611}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi47
                }
                _ = fi46
            }
            v45 = (1 >= 641)
            _ = v45
        }
        v54 := GenS7{F0: v1}
        _ = v54
    } else {
        _ = len(v30)
        v55 := GenS1{F0: 681, F1: 2147483648, F2: 693}
        _ = v55
    }
    switch 1 {
    case 42:
        var v56 int16 = -686
        _ = v56
    case 127:
        var v57 bool = false
        _ = v57
    case 32767:
        for fi58 := 0; fi58 < 5; fi58++ {
            for fi59 := 0; fi59 < 11; fi59++ {
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                var v60 uint16 = 42
                _ = fi59
                _ = v60
            }
            _ = fi58
        }
    }
    switch v1 {
    case "hello":
        if true {
            switch 32767 {
            case -1:
                v61 := GenS5{F0: -1, F1: 18446744073709551615}
                _ = v61
            case 128:
                v3[""] = (24 > 1)
            case 2147483647:
                _ = len(v30)
            }
        }
    default:
        switch -9223372036854775808 {
        case 4294967295:
            for fi62 := 0; fi62 < 13; fi62++ {
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                for fi63 := 0; fi63 < 17; fi63++ {
                    v64 := GenS8{F0: 2, F1: (fi63 - fi62), F2: 2}
                    var v65 string = ((v1 + v1) + "")
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = fi62
            }
        case -9223372036854775808:
            v9 = 435
        default:
            var v66 bool = true
            _ = v66
        }
    }
    for fi67 := 0; fi67 < 8; fi67++ {
        var v68 int8 = -93
        if (255 <= 127) {
            v9 = -1
        } else {
            for fi69 := 0; fi69 < 19; fi69++ {
                fmt.Println(fi67, v1, v1)
                var v70 string = v1
                v2[2] = fi67
                _ = fi69
                _ = v70
            }
            switch v1 {
            case "a":
                if (-2147483648 != 2147483647) {
                    delete(v3, ("a" + "a"))
                }
            case "hello":
                fmt.Println(v68)
            case "":
                switch fi67 {
                case 9223372036854775807:
                    switch (fi67 - (fi67 - fi67)) {
                    case -657:
                        v9 = v9
                    default:
                        v71 := v2[0]
                        _ = v71
                    }
                }
            default:
                if ((((3.14159 / 3.14159) - 3.14159) >= 1e10) || false) {
                    v9 -= (-(v9))
                } else {
                    fmt.Println(v9, v9)
                    if (false && true) {
                        v68 = -119
                    }
                    var v72 uint64 = 496
                    _ = v72
                }
            }
            v9 = v9
        }
        _ = fi67
        _ = v68
    }
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v9
}
