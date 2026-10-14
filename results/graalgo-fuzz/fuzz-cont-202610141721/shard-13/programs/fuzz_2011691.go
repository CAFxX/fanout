package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint32; F1 uint16; F2 bool }
func genFn0(p0 int, p1 uint8) int32 {
    fmt.Println(p1, p0)
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            var v8 int = (^(p0))
            if ((true && ((!false) && (102 > 42))) && ((float32((0.1 - v6)) / 1e10) <= 3.14159)) {
                var v9 uint16 = 861
                if (!false) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(fi10)
                        p0 += (p0 ^ fi10)
                        fmt.Println(v9)
                        var v11 int = (p0 << ((2) & 63))
                        var v12 uint32 = 415
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                } else {
                    v9 = (v9 ^ (v9 >> ((255) & 63)))
                    switch "hello" {
                    case "\n\t":
                        v13 := []int{}
                        _ = v13
                    case "":
                        if (int32(p0) <= 2147483647) {
                            var v14 int16 = 42
                            v15 := []int{}
                            _ = len(v15)
                            _ = v14
                            _ = v15
                        } else {
                            var v16 uint = 1
                            for fi17 := 0; fi17 < 12; fi17++ {
                                if (true && true) {
                                    switch (p0 | -223) {
                                    case 679:
                                        p1 += 91
                                    }
                                    v18 := GenS1{F0: (v8 | -2147483648), F1: 248, F2: (v9 % ((v9) | 1))}
                                    _ = v18
                                } else {
                                    v19 := []int{4294967295, (-(v8)), (v8 + -639)}
                                    var v20 float64 = v6
                                    var v21 uint32 = 693
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = fi17
                            }
                            _ = v16
                        }
                    }
                }
                var v22 float32 = (1e10 + 0.1)
                _ = v22
                _ = v9
            } else {
                v23 := []int{(p0 * p0), -2147483648}
                v6 += (((v6 * 3.14159) - v6) / v6)
                fmt.Println(p1, p1)
                _ = v23
            }
            _ = v8
        }
        _ = fi7
    }
    if (p0 < p0) {
        v24 := []int{p0}
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                if len(v24) > 0 {
                    v24[0] = (fi26 ^ (^(p0)))
                }
                _ = fi26
            }
            var v27 float32 = 1.5
            if ((^(fi25)) == (p0 << ((0) & 63))) {
                v27 = float32((v6 / v6))
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            } else {
                fmt.Println(p0)
                var v28 uint8 = p1
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                _ = v28
            }
            _ = fi25
            _ = v27
        }
        _ = v24
    } else {
        for fi29 := 0; fi29 < 10; fi29++ {
            p0 += 190
            _ = fi29
        }
        for fi30 := 0; fi30 < 11; fi30++ {
            if (1 <= -128) {
                if true {
                    p0 += (-(fi30))
                }
            } else {
                switch "hello" {
                case "a":
                    v6 -= float64((0.0 + 0.1))
                default:
                    fmt.Println(p1, p1)
                }
                switch "hello" {
                case "hello":
                    if true {
                        v31 := []int{}
                        var v32 uint16 = 42
                        _ = v31
                        _ = v32
                    } else {
                        p1 = p1
                        if true {
                            p0 += (fi30 ^ p0)
                        }
                    }
                default:
                    if ((p1 == (p1 / ((p1) | 1))) && true) {
                        fmt.Println(p1)
                        fmt.Println(fi30, p1)
                        fmt.Println(p1, p1, fi30)
                    } else {
                        v33 := GenS2{F0: 65535}
                        v34 := []int{fi30, fi30}
                        v35 := GenS3{F0: 2, F1: 101, F2: -1}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
                switch (p0 % (((fi30 >> ((4294967295) & 63))) | 1)) {
                case 866:
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := []int{fi36, p0, (fi30 | -1)}
                        _ = fi36
                        _ = v37
                    }
                case 128:
                    if false {
                        v38 := []int{32767}
                        if len(v38) > 0 {
                            v38[0] = (fi30 + fi30)
                        }
                        var v39 int16 = 0
                        _ = v38
                        _ = v39
                    } else {
                        p0 = -2147483648
                    }
                default:
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 float64 = 0.1
                        var v42 uint32 = 736
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
            }
            _ = fi30
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int32 = 282
    var v2 int = -9223372036854775808
    var v3 uint16 = 1
    v4 := []int{}
    v5 := make(map[uint32]int)
    v1 = genFn0((v2 / (((v2 * 128)) | 1)), 1)
    v2 += v2
    var v43 float64 = (0.0 + 0.1)
    v44 := GenS4{F0: 2147483648, F1: (v2 % ((v2) | 1))}
    if true {
        var v45 uint = 255
        v2 = (^(v2))
        _ = v45
    } else {
        var v46 uint32 = 2147483648
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v1 += (^(v1))
        v1 = (v1 >> ((0) & 63))
        v2 = (^(fi47))
        _ = fi47
    }
    v48 := GenS5{F0: 0, F1: 42, F2: true}
    var v49 int64 = 974
    var v50 string = "a"
    var v51 bool = (-2.25 >= (0.0 / 1.0))
    var v52 uint8 = 2
    switch v50 {
    case "foo bar":
        var v53 uint = 18446744073709551615
        _ = v53
    case "a":
        fmt.Println(v3)
    }
    var v54 int = 2147483647
    v51 = (65535 == 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v54
}
