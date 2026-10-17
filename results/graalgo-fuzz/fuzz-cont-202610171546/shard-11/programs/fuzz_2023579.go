package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint16; F2 int32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 uint) int8 {
    var v7 int8 = 0
    fmt.Println(v7)
    v8 := GenS1{F0: 1000, F1: (^(p0)), F2: 65535}
    p0 += 2
    switch "" {
    case "foo bar":
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 0, F1: (p0 ^ p0), F2: 2147483648}
            v11 := []int{(^(fi9)), fi9, (fi9 ^ 574), (^(fi9))}
            var v12 bool = (!false)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 42
    v3 := []uint32{2147483647}
    var v4 int32 = 0
    var v5 uint = 42
    var v6 uint32 = 193
    if (((v5 - (^(v5))) == 1) || (-1 >= 1)) {
        if false {
            v13 := GenS2{F0: 42, F1: 0, F2: 125}
            switch "hello" {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = (v6 % ((v6) | 1))
                }
            case "foo bar":
                if true {
                    v2 += (-(v2))
                    if (false && (true && ((v4 ^ v4) != (v4 << (((v5 ^ 18446744073709551615)) & 63))))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v14 bool = (255 > 32767)
                        if (!v14) {
                            for fi15 := 0; fi15 < 2; fi15++ {
                                var v16 uint64 = 0
                                switch ("a" + "hello") {
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = v6
                                    }
                                case "hello":
                                    v17 := GenS3{F0: (float64(float32(0.0)) >= -2.25), F1: 676, F2: 2147483647}
                                    _ = v17
                                default:
                                    var v18 uint8 = 1
                                    _ = v18
                                }
                                var v19 int = 128
                                _ = fi15
                                _ = v16
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 4; fi20++ {
                                if v14 {
                                    fmt.Println(fi20)
                                    switch 42 {
                                    case 4294967295:
                                        v4 = (v4 << (((v5 >> ((142) & 63))) & 63))
                                    case -2147483648:
                                        v4 = (v4 % ((v4) | 1))
                                    case 714:
                                        v5 -= (v5 + v5)
                                    }
                                }
                                v6 = (v6 | v6)
                                _ = fi20
                            }
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 uint = v5
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    var v24 int32 = (v4 + (v4 ^ (v4 - (v4 << ((v5) & 63)))))
                                    _ = fi23
                                    _ = v24
                                }
                                if (991 <= 55) {
                                    v25 := GenS2{F0: 32767, F1: 32767, F2: 2}
                                    var v26 uint = v22
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi21
                                _ = v22
                            }
                        } else {
                            delete(v1, 32768)
                        }
                        _ = v14
                    } else {
                        for fi27 := 0; fi27 < 9; fi27++ {
                            switch "\n\t" {
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = v6
                                }
                            case "\n\t":
                                v28 := GenS3{F0: (!true), F1: 0, F2: 0}
                                _ = v28
                            case "foo bar":
                                var v29 int = fi27
                                _ = v29
                            }
                            if (v2 < (v2 * (-(v2)))) {
                                v5 = v5
                            } else {
                                v30 := GenS4{F0: v4, F1: (fi27 * fi27)}
                                for fi31 := 0; fi31 < 16; fi31++ {
                                    var v32 int64 = v2
                                    switch ((("" + ("" + "hello")) + "a") + ((("hello" + "\n\t") + ("foo bar" + ("" + "a"))) + ("a" + "a"))) {
                                    case "a":
                                        v1[fi31] = int(v6)
                                    case "\n\t":
                                        switch (fi27 - 42) {
                                        case -300:
                                            var v33 uint16 = 42
                                            _ = v33
                                        case -309:
                                            var v34 uint8 = 255
                                            _ = v34
                                        default:
                                            var v35 bool = (!(false || true))
                                            _ = v35
                                        }
                                    case "":
                                        if false {
                                            v3 = append(v3, 2147483647)
                                        }
                                    default:
                                        v6 = v6
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                                _ = v30
                            }
                            _ = fi27
                        }
                    }
                }
            case "":
                fmt.Println(v2, v6, v5)
            }
            v6 += (^(v6))
            _ = v13
        } else {
            for fi36 := 0; fi36 < 5; fi36++ {
                if true {
                    if len(v3) > 0 {
                        v3[0] = (v6 / ((42) | 1))
                    }
                    v37 := GenS5{F0: 1, F1: 18446744073709551615}
                    _ = v37
                }
                _ = fi36
            }
            if len(v3) > 0 {
                v3[0] = (v6 / ((v6) | 1))
            }
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        v2 = v2
        _ = fi38
    }
    switch 32767 {
    case 32767:
        var v39 uint32 = 2147483648
        _ = v39
    }
    if true {
        if true {
            v4 = (v4 | (v4 % ((-1) | 1)))
        } else {
            v3 = append(v3, uint32(v5))
            var v40 uint8 = 255
            _ = v40
        }
        switch -917 {
        case -1:
            if false {
                if true {
                    _ = len(v1)
                    var v41 string = ""
                    _ = v41
                }
                var v42 float64 = 3.14159
                fmt.Println(v2, v5, v6)
                _ = v42
            } else {
                v43 := GenS6{F0: 0}
                switch -1 {
                case -9223372036854775808:
                    v44 := v1[780]
                    _ = v44
                case 971:
                    v2 -= (v2 * v2)
                default:
                    var v45 string = (("\n\t" + ("foo bar" + "\n\t")) + "hello")
                    _ = v45
                }
                _ = v43
            }
        case 42:
            if true {
                v2 = (-(v2))
            } else {
                if (128 >= int(v5)) {
                    for fi46 := 0; fi46 < 10; fi46++ {
                        v4 += (^(v4))
                        v47 := GenS7{F0: v2}
                        fmt.Println(v2, v4)
                        _ = fi46
                        _ = v47
                    }
                } else {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        var v49 int8 = -69
                        v6 = 2147483648
                        v5 = 0
                        _ = fi48
                        _ = v49
                    }
                    fmt.Println(v6)
                }
            }
        case -294:
            _ = len(v3)
        }
    } else {
        var v50 uint16 = 526
        _ = v50
    }
    var v51 int32 = (v4 - 846)
    var v52 bool = ((42 != 2) || true)
    v2 = 1
    var v53 float32 = float32(1.5)
    var v54 uint16 = 516
    var v55 int32 = (-(v4))
    if v52 {
        fmt.Println(v6, v6)
        fmt.Println(v5, v6)
    }
    var v56 float64 = 3.14159
    var v57 uint = (^(v5))
    for fi58 := 0; fi58 < 14; fi58++ {
        fmt.Println(v2)
        switch "\n\t" {
        case "":
            for fi59 := 0; fi59 < 5; fi59++ {
                v57 = (v5 | v5)
                _ = fi59
            }
        default:
            if v52 {
                var v60 int16 = 42
                _ = v60
            }
        }
        v61 := GenS1{F0: -32768, F1: (v5 - 1), F2: 4294967295}
        _ = fi58
        _ = v61
    }
    v62 := GenS7{F0: 17}
    v53 = float32(v56)
    for fi63 := 0; fi63 < 9; fi63++ {
        if (!false) {
            v4 = (v51 / ((v4) | 1))
        } else {
            for fi64 := 0; fi64 < 4; fi64++ {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v65 := GenS7{F0: (^(v2))}
                    _ = v65
                case "":
                    if (v52 && v52) {
                        v57 -= (v57 << (((^(v57))) & 63))
                    }
                    var v66 uint32 = (v6 & v6)
                    v51 = 1
                    _ = v66
                case "foo bar":
                    v67 := GenS6{F0: 206}
                    _ = v67
                }
                var v68 uint = (^(v57))
                var v69 int = (fi63 & fi64)
                _ = fi64
                _ = v68
                _ = v69
            }
        }
        var v70 float32 = -2.25
        v71 := GenS6{F0: 255}
        _ = fi63
        _ = v70
        _ = v71
    }
    var v72 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v62
    _ = v72
}
