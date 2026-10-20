package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint8; F1 bool; F2 uint8 }

func main() {
    v1 := []string{((("" + "\n\t") + "foo bar") + ("a" + ("foo bar" + ""))), "a", "", (("\n\t" + ("a" + ("" + ("\n\t" + "foo bar")))) + (("\n\t" + "") + "a"))}
    v2 := make(map[int]bool)
    var v3 bool = true
    var v4 uint32 = 2147483647
    v5 := []int{9223372036854775807, -2147483648, 438, 42}
    var v6 uint8 = 121
    if false {
        v6 += (v6 << ((173) & 63))
        var v7 int = 128
        _ = v7
    }
    fmt.Println(v6, v4, v4)
    fmt.Println(v6, v4)
    for fi8 := 0; fi8 < 5; fi8++ {
        if (0 > 1) {
            var v9 int16 = 32767
            v6 += (^(v6))
            _ = v9
        } else {
            if v3 {
                v10 := GenS1{F0: 2, F1: 4294967295, F2: 2}
                var v11 uint = 42
                _ = v10
                _ = v11
            }
            var v12 bool = v3
            _ = v12
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 float64 = (1.5 * (((3.14159 * -2.25) / 1e10) / ((0.1 + (0.0 / (1.5 + (1.5 * 0.1)))) + float64(3.14159))))
        _ = fi8
        _ = v13
    }
    v6 -= 189
    switch 486 {
    case -809:
        fmt.Println(v6, v4, v3)
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            v6 += (v6 - v6)
            delete(v2, fi14)
            v15 := GenS2{F0: (v4 | 65535), F1: v3}
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        _ = len(v2)
        var v17 uint8 = (^(v6))
        var v18 float32 = -2.25
        _ = fi16
        _ = v17
        _ = v18
    }
    if v3 {
        if ("" < ("" + "")) {
            var v19 float64 = float64(1.5)
            _ = v19
        }
        switch -1 {
        case -2147483648:
            v1 = append(v1, "hello")
        case -9223372036854775808:
            var v20 int32 = 42
            _ = v20
        }
    }
    if v3 {
        v21 := GenS3{F0: 32767}
        _ = v21
    } else {
        if (0 <= 65535) {
            if v3 {
                fmt.Println(v6, v4, v6)
                var v22 int16 = -1
                _ = v22
            } else {
                v6 = (v6 >> ((481) & 63))
                if len(v5) > 0 {
                    v5[0] = 128
                }
                var v23 bool = v3
                _ = v23
            }
        } else {
            var v24 uint16 = 0
            _ = v24
        }
        if (778 <= -32768) {
            for fi25 := 0; fi25 < 5; fi25++ {
                if false {
                    v26 := GenS1{F0: 255, F1: 18, F2: 0}
                    _ = v26
                }
                _ = fi25
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                for fi28 := 0; fi28 < 3; fi28++ {
                    fmt.Println(fi27)
                    var v29 uint32 = v4
                    var v30 int16 = 1000
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        } else {
            if true {
                for fi31 := 0; fi31 < 16; fi31++ {
                    _ = len(v2)
                    v6 = v6
                    v4 -= uint32(v6)
                    _ = fi31
                }
                var v32 int = -1
                _ = v32
            } else {
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                v6 -= 1
                v3 = v3
            }
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            switch "foo bar" {
            case "":
                switch (-(fi33)) {
                case 127:
                    if ((0.1 + float64((3.14159 / 3.14159))) <= 0.0) {
                        _ = len(v5)
                        v6 = 158
                    } else {
                        var v34 uint8 = 157
                        _ = v34
                    }
                case 2147483647:
                    v6 = v6
                case -2147483648:
                    var v35 int64 = 0
                    _ = v35
                }
            case "foo bar":
                var v36 int8 = 56
                _ = v36
            case "a":
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 int16 = 65
                    switch ("" + "") {
                    case "hello":
                        v2[-2147483648] = (v3 || false)
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            if (588 < 65535) {
                                if v3 {
                                    var v39 string = ("" + ("a" + "foo bar"))
                                    v40 := v2[(fi33 & (fi37 & fi33))]
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 string = ("hello" + "\n\t")
                                    _ = v41
                                }
                            } else {
                                fmt.Println(fi37)
                                fmt.Println(v4, v6)
                                v42 := GenS1{F0: 0, F1: 262, F2: 843}
                                _ = v42
                            }
                        default:
                            v6 -= 105
                        }
                    }
                    _ = fi37
                    _ = v38
                }
            default:
                if (!(!(!(1 < -2147483648)))) {
                    var v43 int = fi33
                    for fi44 := 0; fi44 < 5; fi44++ {
                        fmt.Println(v43, fi44, fi44)
                        v6 = 255
                        _ = fi44
                    }
                    for fi45 := 0; fi45 < 14; fi45++ {
                        v4 = v4
                        if false {
                            v1 = append(v1, "foo bar")
                        }
                        _ = fi45
                    }
                    _ = v43
                }
            }
            _ = fi33
        }
    }
    if (3.14159 >= (0.0 * float32(((0.0 * (1.5 * (1e10 + -2.25))) * 1e10)))) {
        var v46 int64 = -577
        if true {
            v46 -= (v46 - int64(v6))
        } else {
            if (-647 != 1) {
                if (813 > 2147483648) {
                    if (v46 != (v46 * -226)) {
                        var v47 uint8 = 2
                        _ = len(v2)
                        fmt.Println(v47, v47, v4)
                        _ = v47
                    } else {
                        v48 := GenS4{F0: v6, F1: -1}
                        fmt.Println(v46, v46)
                        v49 := GenS5{F0: 32767, F1: 65535, F2: 255}
                        _ = v48
                        _ = v49
                    }
                } else {
                    v50 := GenS6{F0: (v6 - v6), F1: (794 < 832), F2: (v6 + v6)}
                    _ = v50
                }
            }
        }
        v51 := GenS4{F0: v6, F1: -1}
        _ = v46
        _ = v51
    } else {
        for fi52 := 0; fi52 < 8; fi52++ {
            for fi53 := 0; fi53 < 3; fi53++ {
                var v54 uint64 = 18446744073709551615
                var v55 float32 = (0.0 * ((((3.14159 - 0.0) / 1.0) + (3.14159 * (float32(0.1) / 1.5))) / float32((1e10 * -2.25))))
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = fi52
        }
        var v56 string = "foo bar"
        for fi57 := 0; fi57 < 14; fi57++ {
            v2[int(v6)] = true
            delete(v2, fi57)
            var v58 float64 = 1e10
            _ = fi57
            _ = v58
        }
        _ = v56
    }
    v59 := GenS5{F0: 0, F1: 0, F2: (v6 ^ 128)}
    v2[-537] = (!v3)
    if len(v5) > 0 {
        v5[0] = -364
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
