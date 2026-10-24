package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 int32) int {
    v5 := GenS1{F0: 91, F1: (true && true)}
    v6 := []int{}
    var v7 int32 = p0
    var v8 int32 = v7
    v9 := GenS1{F0: 531, F1: (!(true && true))}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint) uint8 {
    p0 = (^(p0))
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 -= 42
        v11 := []int{127}
        var v12 bool = (-9223372036854775808 <= 15)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint16 = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        p0 = p0
        if (128 != 2) {
            var v15 float32 = ((0.1 * float32(((1.5 * 1e10) * 1e10))) / 1.5)
            var v16 float64 = -2.25
            if (fi14 <= fi14) {
                v17 := []int{genFn0(42), genFn0(760)}
                _ = v17
            } else {
                var v18 uint8 = 97
                v19 := []int{160, (fi14 ^ fi14)}
                for fi20 := 0; fi20 < 10; fi20++ {
                    v21 := GenS2{F0: p0, F1: 42}
                    _ = fi20
                    _ = v21
                }
                _ = v18
                _ = v19
            }
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    var v22 float32 = -2.25
    _ = p0
    _ = v13
    _ = v22
    return 32
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 198
    var v3 uint = 221
    v4 := []int{32767}
    v3 -= 2
    v23 := GenS1{F0: 255, F1: (32768 <= int(v2))}
    v24 := GenS3{F0: -2147483648, F1: (^(v3))}
    var v25 int64 = 42
    v3 -= v3
    for fi26 := 0; fi26 < 6; fi26++ {
        if (false && false) {
            var v27 int8 = 127
            if false {
                if true {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        v4 = append(v4, genFn0(730))
                        v25 -= -542
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 19; fi29++ {
                        if true {
                            var v30 int64 = (v25 & -1)
                            v3 += 0
                            _ = v30
                        }
                        _ = fi29
                    }
                    if true {
                        v1[fi26] = true
                        if true {
                            v2 = v2
                            if (3.14159 != (((float32(1e10) / -2.25) * float32(1.5)) - (0.0 / float32(0.1)))) {
                                var v31 float64 = 3.14159
                                delete(v1, genFn0(42))
                                _ = v31
                            }
                        }
                        v25 -= (v25 & (^(v25)))
                    }
                } else {
                    v32 := GenS3{F0: 248, F1: (v3 >> (((v3 << ((v3) & 63))) & 63))}
                    _ = v32
                }
                fmt.Println(v27, v2, fi26)
                for fi33 := 0; fi33 < 8; fi33++ {
                    v1[127] = ((true || (true || (4294967295 != 731))) || (!(false || (-250 <= 2147483647))))
                    var v34 bool = (true || (71 != 0))
                    _ = fi33
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS4{F0: 1}
                var v37 int32 = 42
                _ = len(v4)
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v27
        } else {
            v38 := v1[genFn0(353)]
            switch "\n\t" {
            case "hello":
                if v38 {
                    var v39 uint64 = (v2 << (((^(v3))) & 63))
                    _ = v39
                } else {
                    switch "foo bar" {
                    case "a":
                        if (!(52 > 0)) {
                            fmt.Println(v3, fi26)
                            if (v3 < 42) {
                                for fi40 := 0; fi40 < 17; fi40++ {
                                    var v41 string = "\n\t"
                                    if (!v38) {
                                        var v42 int64 = v25
                                        var v43 uint8 = 241
                                        _ = v42
                                        _ = v43
                                    }
                                    v3 += (v3 - 777)
                                    _ = fi40
                                    _ = v41
                                }
                            }
                            for fi44 := 0; fi44 < 14; fi44++ {
                                v45 := v1[fi26]
                                var v46 string = "hello"
                                var v47 float64 = float64(0.1)
                                _ = fi44
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        } else {
                            switch (^(fi26)) {
                            case -298:
                                if v38 {
                                    var v48 uint32 = 42
                                    v49 := GenS2{F0: v3, F1: 65535}
                                    _ = v48
                                    _ = v49
                                }
                            case 507:
                                v38 = (false || v38)
                            }
                            fmt.Println(v25)
                            fmt.Println(v25, v25)
                        }
                    default:
                        v2 += 2
                    }
                }
            default:
                var v50 string = "\n\t"
                _ = v50
            }
            v51 := GenS1{F0: 1, F1: v38}
            _ = v38
            _ = v51
        }
        _ = fi26
    }
    if true {
        v52 := GenS2{F0: (^(v3)), F1: 548}
        v25 -= v25
        _ = v52
    } else {
        v53 := GenS5{F0: ((("hello" + ("" + "a")) + "hello") + ((("hello" + ("" + "\n\t")) + "hello") + "\n\t")), F1: (false && (!false)), F2: 18}
        var v54 int32 = 242
        _ = v53
        _ = v54
    }
    var v55 int16 = 1000
    var v56 float64 = 0.1
    fmt.Println(v3)
    fmt.Println(v25, v56, v56)
    v57 := GenS3{F0: 408, F1: (v3 & v3)}
    switch 815 {
    case 1:
        switch "foo bar" {
        case "\n\t":
            if (-2147483648 == 2147483647) {
                _ = len(v4)
            }
        }
    default:
        for fi58 := 0; fi58 < 20; fi58++ {
            _ = len(v4)
            for fi59 := 0; fi59 < 3; fi59++ {
                _ = len(v1)
                _ = fi59
            }
            _ = fi58
        }
    }
    var v60 uint32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v60
}
