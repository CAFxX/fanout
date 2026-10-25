package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int32; F1 int64; F2 int8 }
func genFn0() bool {
    v6 := GenS1{F0: -1, F1: 916, F2: 2147483647}
    if (1.5 != ((3.14159 / 1.0) + -2.25)) {
        var v7 int16 = 1000
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 1, F1: 462, F2: 42}
            _ = fi8
            _ = v9
        }
        var v10 bool = (true && true)
        _ = v10
    }
    v11 := []int{}
    if false {
        var v12 uint8 = 127
        _ = v12
    } else {
        var v13 int8 = 127
        _ = v13
    }
    _ = v11
    _ = v6
    return (-1 >= 1)
}
func genFn1(p0 int16) uint32 {
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 float32 = 1.5
        _ = fi14
        _ = v15
    }
    if false {
        switch (("" + "hello") + "\n\t") {
        case "hello":
            if false {
                for fi16 := 0; fi16 < 19; fi16++ {
                    v17 := GenS1{F0: -128, F1: 141, F2: 2147483647}
                    _ = fi16
                    _ = v17
                }
            }
        case "":
            v18 := GenS2{F0: -9223372036854775808}
            _ = v18
        default:
            var v19 uint = 42
            _ = v19
        }
        if (!true) {
            if ((-2.25 + 3.14159) >= (1e10 + 3.14159)) {
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 uint32 = 2147483647
                    v22 := GenS3{F0: (157 < 1), F1: 255}
                    v23 := []int{fi20, fi20}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 float32 = (float32(0.1) * 1e10)
                switch 177 {
                case 1:
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 uint = 550
                        v26 = v26
                        if (!true) {
                            v24 = (v24 - v24)
                            switch "" {
                            case "a":
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    if ((1.5 / 1e10) > (3.14159 * ((((1.5 + 3.14159) / (1e10 - 0.0)) / 3.14159) / (1.5 * 1.5)))) {
                                        p0 = p0
                                    }
                                    _ = fi27
                                }
                            }
                            fmt.Println(v26, v26)
                        } else {
                            v28 := []int{fi25, -429}
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                case 2147483647:
                    v24 = (1.5 + v24)
                default:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v31 := GenS4{F0: 675, F1: 255}
                            v24 += ((v24 - 0.1) - v24)
                            _ = fi30
                            _ = v31
                        }
                        switch fi29 {
                        case 0:
                            fmt.Println(fi29, p0)
                        case 32768:
                            fmt.Println(fi29)
                        }
                        var v32 string = (("a" + "\n\t") + ("\n\t" + (((("\n\t" + "foo bar") + "a") + "hello") + "\n\t")))
                        _ = fi29
                        _ = v32
                    }
                }
                _ = v24
            }
        }
        var v33 uint64 = 42
        _ = v33
    } else {
        v34 := []int{-2147483648}
        _ = v34
    }
    p0 = -32768
    if true {
        v35 := []int{-389}
        _ = v35
    } else {
        var v36 uint32 = 654
        for fi37 := 0; fi37 < 8; fi37++ {
            fmt.Println(fi37, p0, fi37)
            _ = fi37
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            var v39 string = "foo bar"
            if true {
                if (-700 < (p0 * 274)) {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        fmt.Println(fi38, v39)
                        if (!false) {
                            fmt.Println(fi40, v39)
                            v41 := GenS2{F0: -762}
                            _ = v41
                        } else {
                            fmt.Println(fi38, v36)
                            v36 -= v36
                        }
                        fmt.Println(fi40, p0, fi38)
                        _ = fi40
                    }
                    v42 := []int{32768}
                    fmt.Println(fi38)
                    _ = v42
                }
            } else {
                v43 := []int{(fi38 * (fi38 & (fi38 | fi38))), 32767}
                _ = v43
            }
            switch fi38 {
            case -912:
                v39 = ("" + "")
            case 2147483647:
                for fi44 := 0; fi44 < 8; fi44++ {
                    v45 := []int{}
                    v46 := GenS4{F0: p0, F1: 0}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            case 1:
                if true {
                    v47 := []int{(fi38 % (((fi38 << ((199) & 63))) | 1))}
                    v39 = (v39 + "")
                    _ = v47
                }
            }
            _ = fi38
            _ = v39
        }
        _ = v36
    }
    p0 += (p0 ^ p0)
    _ = p0
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 9223372036854775807
    var v3 int = 1
    var v4 float32 = (1.5 / 1.0)
    var v5 int64 = (v2 * 852)
    if false {
        v48 := GenS4{F0: -32768, F1: 140}
        _ = v48
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        fmt.Println(v4)
        for fi50 := 0; fi50 < 13; fi50++ {
            v5 = (v5 - (v5 << ((224) & 63)))
            v4 = 3.14159
            _ = fi50
        }
        _ = fi49
    }
    var v51 int64 = v2
    v1[65535] = 2147483647
    v52 := v1[42]
    v53 := GenS5{F0: -1, F1: (^(v51)), F2: int8(v52)}
    v52 = 0
    v54 := GenS4{F0: 42, F1: 2}
    var v55 uint8 = 0
    var v56 bool = (!(!false))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
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
}
