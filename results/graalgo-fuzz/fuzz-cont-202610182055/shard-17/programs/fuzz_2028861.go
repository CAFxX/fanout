package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint8) int64 {
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 2; fi8++ {
        if false {
            var v9 int = (fi8 << ((2) & 63))
            var v10 uint64 = 2
            v10 = (^(v10))
            _ = v10
            _ = v9
        }
        fmt.Println(fi8, v7, v7)
        _ = fi8
    }
    if (true || (!true)) {
        switch -237 {
        case 9223372036854775807:
            v7 = v7
        case 32767:
            if false {
                v11 := []int{850, -240}
                _ = v11
            }
        default:
            v12 := []int{32767, -2147483648, 32767, 128}
            _ = v12
        }
        v13 := GenS1{F0: 42}
        _ = v13
    } else {
        var v14 uint16 = p0
        fmt.Println(v7, v7)
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: 2147483648, F2: 1}
    _ = p0
    _ = p1
    _ = v15
    _ = v7
    return 227
}

func main() {
    var v1 int = -446
    v2 := []uint32{943}
    var v3 uint16 = 65535
    var v4 bool = (v1 == -226)
    var v5 int32 = 590
    v6 := make(map[int]uint32)
    v5 -= 42
    if ((v5 - 909) != (^(v5))) {
        var v16 int8 = -128
        switch (-(v1)) {
        case 32768:
            if (v4 && true) {
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint16 = (v3 * (v3 * 0))
                    v16 -= -128
                    var v19 uint8 = 172
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int8 = (v16 * 1)
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v23 := GenS3{F0: 2147483648}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi21
                    }
                } else {
                    v1 = (^(v1))
                    if true {
                        if v4 {
                            v1 -= (v1 ^ v1)
                            fmt.Println(v1, v1, v5)
                        } else {
                            var v24 uint32 = 655
                            var v25 bool = (((v16 | (v16 << ((2) & 63))) == v16) && ((((("\n\t" + "") + "a") + "a") + (("\n\t" + "\n\t") + "\n\t")) == "hello"))
                            _ = v24
                            _ = v25
                        }
                    } else {
                        if v4 {
                            var v26 uint8 = 23
                            _ = v26
                        }
                        switch v1 {
                        case 127:
                            var v27 uint8 = 1
                            _ = v27
                        case -295:
                            v28 := GenS2{F0: 42, F1: 42, F2: 65535}
                            _ = v28
                        case -2147483648:
                            if (true || v4) {
                                v29 := GenS4{F0: v1, F1: v1}
                                v30 := GenS5{F0: 255, F1: 2}
                                _ = v29
                                _ = v30
                            } else {
                                v31 := v6[42]
                                v1 = int(v20)
                                _ = v31
                            }
                        default:
                            if (!(255 <= 520)) {
                                v2 = append(v2, 471)
                            }
                        }
                    }
                }
                _ = v20
            }
        default:
            v4 = v4
        }
        switch 1 {
        case -195:
            if v4 {
                v4 = v4
            }
        default:
            if ((!(v4 && v4)) && ((-(v1)) <= v1)) {
                v5 = (v5 % (((v5 % ((-2147483648) | 1))) | 1))
            }
        }
        _ = v16
    } else {
        v32 := GenS1{F0: 0}
        var v33 int = 32768
        _ = v32
        _ = v33
    }
    if true {
        switch (v1 * v1) {
        case 347:
            v3 -= 32767
        case 9223372036854775807:
            _ = len(v2)
        default:
            delete(v6, v1)
        }
    } else {
        var v34 string = "\n\t"
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v5 += 1
        _ = fi35
    }
    switch "" {
    case "foo bar":
        var v36 float32 = 3.14159
        _ = v36
    default:
        if len(v2) > 0 {
            v2[0] = 902
        }
    }
    v4 = v4
    var v37 uint = 533
    var v38 int16 = 32767
    for fi39 := 0; fi39 < 6; fi39++ {
        v4 = v4
        switch 42 {
        case 683:
            var v40 int = (-(v1))
            _ = v40
        case -1:
            var v41 int32 = v5
            _ = v41
        case 64:
            v42 := GenS2{F0: (v37 & uint(v3)), F1: 2, F2: 1}
            _ = v42
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            if ((true || (1.5 <= 0.1)) || v4) {
                v44 := GenS5{F0: 42, F1: 220}
                v45 := GenS5{F0: uint(v3), F1: 42}
                v38 = (v38 / ((v38) | 1))
                _ = v44
                _ = v45
            } else {
                var v46 float32 = 3.14159
                if (1 == genFn0(v3, 42)) {
                    v4 = true
                } else {
                    v38 -= v38
                    _ = len(v6)
                }
                v38 -= (v38 >> ((297) & 63))
                _ = v46
            }
            _ = fi43
        }
        _ = fi39
    }
    v4 = (!(1.5 == float64(float32(float64(0.1)))))
    switch "a" {
    case "hello":
        v4 = (v4 && v4)
    case "a":
        fmt.Println(v4, v3)
    default:
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    if v4 {
        v3 -= (^(v3))
        if (1.5 == ((0.0 + float64((0.0 + ((1e10 * 0.0) * 1e10)))) * (0.1 + 1e10))) {
            if ((v3 - (v3 ^ v3)) > 507) {
                var v47 string = "foo bar"
                switch -9223372036854775808 {
                case 127:
                    for fi48 := 0; fi48 < 19; fi48++ {
                        v49 := GenS6{F0: 35}
                        var v50 int8 = -1
                        var v51 uint32 = 2147483647
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                }
                _ = v47
            } else {
                v37 = (v37 / ((63) | 1))
                v4 = ((88 > 1) && v4)
            }
        } else {
            delete(v6, (v1 + 9223372036854775807))
            fmt.Println(v3, v38)
            v5 = 272
        }
        v3 = (v3 % ((2) | 1))
    }
    var v52 int = v1
    v1 = v52
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
