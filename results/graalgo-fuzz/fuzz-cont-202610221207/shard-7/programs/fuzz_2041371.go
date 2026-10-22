package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint; F2 int64 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 uint64; F1 string }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 1
    v3 := make(map[int]string)
    var v4 int64 = 798
    var v5 float64 = (3.14159 * 1.5)
    v6 := GenS1{F0: 610}
    var v7 bool = false
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS2{F0: 690, F1: v2}
        v10 := GenS2{F0: 1, F1: v2}
        v11 := GenS2{F0: -1, F1: 1}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS2{F0: 2147483647, F1: (^(v2))}
    if v7 {
        fmt.Println(v4, v5, v7)
    } else {
        var v13 int32 = 2147483647
        if (v5 >= v5) {
            for fi14 := 0; fi14 < 1; fi14++ {
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := v3[(-(fi14))]
                    _ = len(v3)
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch fi17 {
                    case 32767:
                        v5 += ((0.1 / (-2.25 - v5)) + (float64((float32(v5) + ((1e10 * (float32(1.5) / 1e10)) * 0.1))) * (((v5 / v5) / 1e10) + v5)))
                    case -1:
                        _ = len(v3)
                    }
                    if ("a" > "\n\t") {
                        if (!(uint(v2) >= 42)) {
                            v4 = 1
                            if (("a" + "foo bar") == "") {
                                var v18 uint8 = v2
                                v2 = v18
                                if ((("" + "\n\t") != ("hello" + ("hello" + ("\n\t" + "foo bar")))) && (!v7)) {
                                    fmt.Println(v5, v18)
                                    if v7 {
                                        v4 -= (v4 + (v4 / ((v4) | 1)))
                                        fmt.Println(v13)
                                        v7 = (((fi14 / ((fi14) | 1)) <= (fi14 << ((2) & 63))) || (!true))
                                    }
                                } else {
                                    v2 = (^(v2))
                                    v19 := GenS3{F0: (((false || ("foo bar" >= ("foo bar" + "a"))) || v7) && (!((float64(1e10) - (v5 + (v5 + 1.5))) != v5))), F1: v18}
                                    v20 := []int{fi17}
                                    if v7 {
                                        fmt.Println(v18, v5, fi14)
                                    } else {
                                        fmt.Println(fi14)
                                        var v21 uint16 = (v1 & (v1 - 65535))
                                        var v22 bool = false
                                        _ = len(v20)
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                _ = v18
                            } else {
                                v23 := v3[-264]
                                v1 = (v1 * (v1 / (((v1 ^ (v1 - v1))) | 1)))
                                switch fi17 {
                                case -2147483648:
                                    v24 := GenS2{F0: v13, F1: v2}
                                    _ = v24
                                default:
                                    v25 := GenS1{F0: 42}
                                    v23 = v23
                                    fmt.Println(v13, v23)
                                    _ = v25
                                }
                                _ = v23
                            }
                        }
                    }
                    v26 := v3[(^(fi17))]
                    v27 := GenS3{F0: v7, F1: uint8(v1)}
                    _ = fi17
                    _ = v26
                    _ = v27
                }
                _ = fi14
            }
        } else {
            var v28 int = 32768
            _ = v28
        }
        _ = v13
    }
    v29 := []int{42}
    if (2147483647 == 1) {
        if (v5 == float64(1e10)) {
            if v7 {
                var v30 int8 = 47
                for fi31 := 0; fi31 < 9; fi31++ {
                    v2 = uint8(v1)
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 float32 = 0.1
                        _ = fi32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 5; fi34++ {
                        v1 += (v1 << ((42) & 63))
                        var v35 uint32 = 242
                        if true {
                            switch fi34 {
                            case 4294967295:
                                fmt.Println(v35)
                            case 32768:
                                var v36 float64 = v5
                                _ = v36
                            }
                        } else {
                            v7 = true
                            var v37 uint64 = 4294967295
                            var v38 int64 = v4
                            _ = v37
                            _ = v38
                        }
                        _ = fi34
                        _ = v35
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                if (v7 && v7) {
                    v4 -= (v4 >> ((0) & 63))
                    if true {
                        v5 = 1.5
                        v5 -= (3.14159 - (float64((0.1 - 0.0)) * (1e10 - 0.0)))
                    } else {
                        var v39 uint8 = 22
                        _ = v39
                    }
                }
                v4 = v4
            }
        } else {
            var v40 float64 = v5
            _ = v40
        }
        v41 := GenS2{F0: 0, F1: (v2 * 184)}
        _ = v41
    }
    v2 = (v2 << ((42) & 63))
    var v42 uint8 = 128
    v5 = v5
    for fi43 := 0; fi43 < 6; fi43++ {
        if (((fi43 * (-(fi43))) > (fi43 & 9223372036854775807)) || (!true)) {
            for fi44 := 0; fi44 < 18; fi44++ {
                v29 = append(v29, 32767)
                _ = fi44
            }
            fmt.Println(v42, v2)
        } else {
            if (1 != 1) {
                v4 = v4
            }
            if true {
                if v7 {
                    if (!v7) {
                        v1 += 0
                        var v45 int16 = -119
                        _ = len(v29)
                        _ = v45
                    }
                    fmt.Println(fi43, v2)
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                v46 := GenS4{F0: 748, F1: (-(v4)), F2: (^(v1))}
                _ = v46
            } else {
                fmt.Println(v4)
                v47 := GenS5{F0: 1, F1: 513, F2: (v4 * v4)}
                _ = v47
            }
        }
        if v7 {
            fmt.Println(v4, v2)
            delete(v3, -2147483648)
            for fi48 := 0; fi48 < 11; fi48++ {
                switch (("foo bar" + "") + (((((("\n\t" + "\n\t") + "foo bar") + "\n\t") + "foo bar") + "") + "foo bar")) {
                case "":
                    v4 -= (^(v4))
                default:
                    v49 := GenS6{F0: 314, F1: v2}
                    _ = v49
                }
                v50 := GenS7{F0: 0, F1: ("\n\t" + "\n\t")}
                _ = fi48
                _ = v50
            }
        } else {
            var v51 uint64 = 18446744073709551615
            switch fi43 {
            case -2147483648:
                fmt.Println(v2)
            case -1:
                var v52 uint8 = uint8(v51)
                _ = v52
            case 127:
                v53 := GenS6{F0: 4294967295, F1: v42}
                _ = v53
            }
            _ = v51
        }
        _ = fi43
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
