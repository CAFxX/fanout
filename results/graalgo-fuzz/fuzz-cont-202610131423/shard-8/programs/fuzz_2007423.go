package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS8 struct { F0 int }
func genFn0() uint64 {
    v4 := []int{146}
    var v5 int8 = 100
    if false {
        if (!false) {
            v5 = v5
            if (!false) {
                v6 := GenS1{F0: 562}
                v7 := GenS2{F0: 255, F1: 2, F2: ((!true) || (255 != 65535))}
                _ = v6
                _ = v7
            }
            v8 := GenS2{F0: 217, F1: 2, F2: (0 >= 0)}
            _ = v8
        }
        fmt.Println(v5)
    }
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint64 = 1
    var v2 bool = (!true)
    v3 := []uint32{0}
    if true {
        v9 := GenS3{F0: 715, F1: (!v2)}
        v10 := GenS4{F0: 32767}
        _ = v10
        _ = v9
    } else {
        v2 = (32767 == 4294967295)
        v1 = (v1 % ((genFn0()) | 1))
    }
    v11 := GenS5{F0: 65535, F1: 42}
    var v12 uint = 255
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint32 = 878
            v1 = (v1 % ((genFn0()) | 1))
            _ = fi14
            _ = v15
        }
        if (32767 == 255) {
            fmt.Println(v2, v1, v1)
            var v16 int16 = 1
            _ = v16
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                v1 = genFn0()
                switch "foo bar" {
                case "hello":
                    if v2 {
                        v1 -= (v1 * v1)
                    }
                case "foo bar":
                    if true {
                        var v18 uint64 = (v1 / (((v1 ^ (v1 * v1))) | 1))
                        var v19 uint16 = 372
                        var v20 int = -260
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "":
                    v12 = 2
                default:
                    v12 = v12
                }
                v21 := GenS6{F0: 42, F1: 328}
                _ = fi17
                _ = v21
            }
        }
        _ = fi13
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(fi22, v12)
            v24 := GenS7{F0: 1, F1: (v1 - (v1 >> (((v12 - v12)) & 63))), F2: 100}
            _ = fi23
            _ = v24
        }
        if (!false) {
            v3 = append(v3, 1000)
            var v25 int64 = -245
            _ = v25
        }
        v26 := GenS8{F0: fi22}
        _ = fi22
        _ = v26
    }
    switch (("a" + (("a" + (("" + "") + "hello")) + "\n\t")) + "hello") {
    case "a":
        var v27 int64 = 1
        _ = v27
    case "\n\t":
        v12 += (^(v12))
    }
    if (1000 > (v12 % (((v12 & v12)) | 1))) {
        fmt.Println(v2)
        if len(v3) > 0 {
            v3[0] = 2
        }
    } else {
        fmt.Println(v12)
    }
    switch 32767 {
    case -797:
        v2 = v2
    default:
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 15; fi29++ {
                switch ("a" + "hello") {
                case "\n\t":
                    v3 = append(v3, 42)
                }
                _ = fi29
            }
            v12 -= v12
            if v2 {
                var v30 uint16 = 0
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v1)
                    v2 = (!(v2 && ((-326 != -923) || ((255 != 255) && false))))
                    if ((^(v1)) != genFn0()) {
                        if ("foo bar" != (("hello" + "a") + "a")) {
                            v12 = v12
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 556
                            }
                            switch (fi28 - (fi31 / (((fi28 / (((fi28 & (-(fi28)))) | 1))) | 1))) {
                            case 0:
                                fmt.Println(fi31, v12, fi28)
                            case 1:
                                var v32 string = "a"
                                var v33 string = "foo bar"
                                _ = v32
                                _ = v33
                            default:
                                var v34 uint64 = v1
                                _ = v34
                            }
                            v2 = ((!(fi28 >= (fi31 << ((v12) & 63)))) || ((v1 * genFn0()) < (^(v1))))
                        }
                        v30 = (^(v30))
                        var v35 uint32 = 31
                        _ = v35
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                for fi36 := 0; fi36 < 19; fi36++ {
                    if v2 {
                        switch "" {
                        case "foo bar":
                            v1 = (v1 / ((v1) | 1))
                        }
                        var v37 uint = v12
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v37
                    }
                    _ = fi36
                }
                if (0.0 < -2.25) {
                    v2 = v2
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    switch 32768 {
                    case -892:
                        for fi38 := 0; fi38 < 8; fi38++ {
                            var v39 bool = false
                            v40 := GenS7{F0: 2, F1: (v1 * 610), F2: 1}
                            if (-918 >= 130) {
                                if false {
                                    v2 = ("\n\t" <= "a")
                                }
                            } else {
                                var v41 float64 = (3.14159 + 1e10)
                                v42 := GenS3{F0: 2, F1: (0.0 >= 1e10)}
                                _ = v41
                                _ = v42
                            }
                            v12 += uint(v1)
                            var v43 bool = v2
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v43
                        }
                    case 472:
                        var v44 int8 = 1
                        _ = v44
                    default:
                        var v45 int8 = -44
                        _ = v45
                    }
                    switch (fi28 + (-(fi28))) {
                    case 42:
                        var v46 float32 = 1e10
                        _ = v46
                    case 1:
                        v47 := GenS4{F0: 235}
                        _ = v47
                    default:
                        var v48 string = "\n\t"
                        _ = v48
                    }
                }
                v49 := GenS4{F0: 934}
                _ = v49
            }
            _ = fi28
        }
    }
    v50 := GenS1{F0: (v1 + (v1 / (((v1 | 4294967295)) | 1)))}
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println(v2, v2)
    v1 -= (v1 - (^(v1)))
    switch "\n\t" {
    case "":
        v12 += v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v50
}
