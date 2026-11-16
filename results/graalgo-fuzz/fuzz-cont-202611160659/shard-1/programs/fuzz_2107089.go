package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    var v5 int16 = 42
    v6 := make(map[string]string)
    v3 = v3
    v7 := GenS1{F0: 100, F1: 700}
    if false {
        fmt.Println(v5, v3)
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS2{F0: 255, F1: 0, F2: 2147483647}
            v10 := []int{(^(fi8)), (^(fi8)), fi8, -2147483648}
            if (false && true) {
                v5 -= v5
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: 127, F1: 99}
                    v13 := GenS2{F0: 32767, F1: 255, F2: 583}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if true {
                    v3 = v3
                    if true {
                        if (!(0 == 32767)) {
                            if true {
                                for fi14 := 0; fi14 < 12; fi14++ {
                                    v3 = v3
                                    if true {
                                        var v15 int32 = 1
                                        v16 := GenS2{F0: 32767, F1: 2, F2: v15}
                                        v3 = v3
                                        _ = v15
                                        _ = v16
                                    } else {
                                        var v17 uint = 278
                                        fmt.Println(v17, v3, v5)
                                        switch fi14 {
                                        case 401:
                                            v17 -= v17
                                        case 9223372036854775807:
                                            v18 := v1[fi14]
                                            _ = v18
                                        case 2147483647:
                                            v17 -= 537
                                        }
                                        _ = v17
                                    }
                                    _ = fi14
                                }
                                delete(v1, -1)
                            } else {
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    v5 += (v5 - (-(v5)))
                                    if len(v10) > 0 {
                                        v10[0] = (fi8 + (fi19 >> ((156) & 63)))
                                    }
                                    v20 := GenS1{F0: 127, F1: 0}
                                    delete(v2, 65535)
                                    fmt.Println(v3)
                                    v21 := GenS1{F0: -1, F1: 32767}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                var v22 int16 = (v5 + (v5 - (^(v5))))
                                _ = v22
                            }
                            v4[2] = (!false)
                            v23 := GenS2{F0: 65535, F1: 169, F2: 708}
                            _ = v23
                        }
                        v24 := GenS3{F0: 73, F1: 32768}
                        v3 = "a"
                        _ = v24
                    }
                } else {
                    var v25 int = -958
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v3 = (v3 + "a")
                        _ = fi26
                    }
                    v27 := GenS4{F0: 661}
                    _ = v25
                    _ = v27
                }
            } else {
                v3 = "hello"
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v3 = v3
        delete(v6, (v3 + ""))
        switch 32768 {
        case 128:
            var v28 float64 = 1.5
            _ = v28
        case -460:
            switch ((v3 + "foo bar") + "a") {
            case "":
                for fi29 := 0; fi29 < 11; fi29++ {
                    var v30 uint64 = 4294967295
                    for fi31 := 0; fi31 < 1; fi31++ {
                        switch (v3 + v3) {
                        case "hello":
                            v32 := GenS5{F0: 2147483647}
                            _ = v32
                        case "foo bar":
                            fmt.Println(fi31, v30)
                            var v33 uint8 = 72
                            _ = v33
                        case "":
                            v30 = 4294967295
                        default:
                            fmt.Println(v5)
                        }
                        var v34 int = 280
                        _ = fi31
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 float64 = ((0.1 + 3.14159) * 0.1)
                        v5 -= v5
                        v3 = (((v3 + v3) + v3) + ("foo bar" + ((v3 + "foo bar") + ("hello" + "a"))))
                        if true {
                            switch "a" {
                            case "hello":
                                _ = len(v6)
                            }
                            var v37 int = (fi29 | fi35)
                            v37 -= (^(fi35))
                            _ = v37
                        } else {
                            v30 = (v30 % ((v30) | 1))
                            v30 = (v30 + (v30 % (((v30 << ((454) & 63))) | 1)))
                            v36 -= (1e10 * 0.1)
                        }
                        _ = fi35
                        _ = v36
                    }
                    _ = fi29
                    _ = v30
                }
            case "hello":
                for fi38 := 0; fi38 < 19; fi38++ {
                    v39 := GenS1{F0: int8(fi38), F1: 255}
                    _ = fi38
                    _ = v39
                }
            }
        default:
            v3 = "\n\t"
        }
    }
    v40 := GenS4{F0: 255}
    v5 = v5
    for fi41 := 0; fi41 < 15; fi41++ {
        if (!((!true) && (!false))) {
            var v42 bool = (255 >= 1)
            v5 = v5
            _ = v42
        }
        v3 = ("\n\t" + v3)
        _ = fi41
    }
    v43 := []int{-561}
    var v44 uint = 934
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
