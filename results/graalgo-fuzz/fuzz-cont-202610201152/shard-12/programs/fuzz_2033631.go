package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint, p1 int) int64 {
    var v6 uint16 = 65535
    v7 := []int{p1, 127, (p1 >> ((p0) & 63)), p1}
    v8 := GenS1{F0: 65535}
    var v9 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 531
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (255 > 65535)
    v3 := make(map[int]string)
    var v4 bool = (((0 > -32768) || v2) || v2)
    v5 := make(map[uint32]int)
    v10 := GenS1{F0: 301}
    v4 = v2
    switch 42 {
    case -2147483648:
        v4 = (v2 || (!true))
    case -1:
        v11 := []int{666, 2147483647, -2147483648, 0}
        _ = v11
    case 298:
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 float64 = 3.14159
            _ = fi12
            _ = v13
        }
    default:
        v14 := GenS2{F0: 0}
        _ = v14
    }
    var v15 float64 = 0.0
    fmt.Println(v15)
    for fi16 := 0; fi16 < 14; fi16++ {
        if (0 != 255) {
            if v4 {
                v17 := GenS2{F0: 1}
                switch fi16 {
                case 2147483647:
                    var v18 uint32 = 65535
                    _ = v18
                case -1:
                    var v19 int8 = -128
                    _ = v19
                case 42:
                    if (65535 <= 968) {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            var v21 uint64 = 0
                            var v22 string = ""
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{128, fi16, fi16, 330}
                        _ = v23
                    }
                }
                _ = v17
            }
            var v24 uint16 = 0
            _ = v24
        }
        _ = fi16
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS3{F0: -819, F1: (fi25 * 4294967295)}
        var v27 uint = 18446744073709551615
        switch "hello" {
        case "\n\t":
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := []int{}
                _ = fi28
                _ = v29
            }
        case "":
            switch "\n\t" {
            case "foo bar":
                if v4 {
                    if ("a" >= (("hello" + "a") + ("a" + "a"))) {
                        if false {
                            v15 += 1.5
                            if false {
                                var v30 uint32 = 448
                                fmt.Println(v2, v30)
                                v27 = (v27 << (((v27 << ((1) & 63))) & 63))
                                _ = v30
                            } else {
                                v4 = v2
                                v31 := GenS2{F0: (fi25 / ((fi25) | 1))}
                                _ = v31
                            }
                            v27 = v27
                        }
                        var v32 int32 = 42
                        var v33 uint = (v27 / (((^(v27))) | 1))
                        _ = v32
                        _ = v33
                    } else {
                        if (!v4) {
                            v34 := GenS4{F0: v4}
                            v4 = (!(false || (42 >= 0)))
                            _ = v34
                        } else {
                            if (!v4) {
                                v15 = v15
                            } else {
                                v35 := GenS5{F0: 32768, F1: "hello"}
                                v36 := GenS5{F0: fi25, F1: ""}
                                switch (^(fi25)) {
                                case 32767:
                                    v37 := v1[""]
                                    _ = v37
                                }
                                _ = v35
                                _ = v36
                            }
                            v38 := []int{fi25, fi25}
                            _ = v38
                        }
                        var v39 bool = (!true)
                        if v2 {
                            v39 = v2
                            fmt.Println(v4, v4, fi25)
                        }
                        _ = v39
                    }
                    delete(v5, 2147483648)
                    var v40 int16 = -1
                    _ = v40
                } else {
                    if v2 {
                        if false {
                            v41 := GenS1{F0: 818}
                            if v2 {
                                v27 += (v27 * (v27 % ((v27) | 1)))
                                v42 := GenS3{F0: 32767, F1: fi25}
                                _ = v42
                            } else {
                                if v4 {
                                    for fi43 := 0; fi43 < 18; fi43++ {
                                        v4 = true
                                        delete(v1, ("\n\t" + "hello"))
                                        _ = fi43
                                    }
                                    v27 -= (v27 * (v27 & v27))
                                    v4 = v4
                                } else {
                                    fmt.Println(v4, v2)
                                    for fi44 := 0; fi44 < 7; fi44++ {
                                        v15 += 1e10
                                        if (true || (793 == 832)) {
                                            if v2 {
                                                fmt.Println(v2, v15, v2)
                                                v45 := []int{fi44, (-(fi44))}
                                                if v4 {
                                                    var v46 int = (fi44 | 32767)
                                                    _ = v46
                                                }
                                                _ = v45
                                            } else {
                                                if v4 {
                                                    if (v4 || (!v2)) {
                                                        v47 := GenS6{F0: (v4 || (100 <= -128)), F1: 2}
                                                        v15 += 0.1
                                                        _ = v47
                                                    } else {
                                                        v15 -= v15
                                                    }
                                                    v48 := GenS1{F0: uint32(v27)}
                                                    if v4 {
                                                        v15 = float64(1.5)
                                                    } else {
                                                        if (!((106 > -63) && (!(1 >= 128)))) {
                                                            v15 = float64(0.0)
                                                        }
                                                        v49 := GenS1{F0: 2147483647}
                                                        _ = v49
                                                    }
                                                    v50 := []int{}
                                                    var v51 int8 = -1
                                                    _ = v48
                                                    _ = v50
                                                    _ = v51
                                                } else {
                                                    v4 = false
                                                    if ((("\n\t" + "\n\t") + "") < "") {
                                                        var v52 int8 = 94
                                                        _ = v52
                                                    }
                                                }
                                                v53 := GenS7{F0: -1, F1: "foo bar"}
                                                _ = v53
                                            }
                                            v4 = v2
                                        } else {
                                            v4 = false
                                            var v54 uint = 255
                                            _ = v54
                                        }
                                        v5[4294967295] = (-(fi44))
                                        _ = fi44
                                    }
                                    var v55 uint64 = 419
                                    _ = v55
                                }
                                var v56 bool = v2
                                if (!(42 >= 127)) {
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        v15 -= v15
                                    case "hello":
                                        v3[fi25] = ""
                                    default:
                                        v57 := []int{(fi25 + (fi25 * int(v27))), fi25, 128}
                                        _ = v57
                                    }
                                    if (false && true) {
                                        var v58 int = fi25
                                        _ = v58
                                    } else {
                                        var v59 int16 = -32768
                                        var v60 uint = (v27 ^ 94)
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                _ = v56
                            }
                            for fi61 := 0; fi61 < 12; fi61++ {
                                v15 -= (v15 + (v15 / v15))
                                _ = fi61
                            }
                            _ = v41
                        }
                        fmt.Println(v27)
                        if (255 != 128) {
                            fmt.Println(v15, v2, fi25)
                        }
                    }
                }
            case "a":
                var v62 uint16 = 2
                _ = v62
            case "":
                v63 := GenS8{F0: -32768}
                _ = v63
            }
        default:
            v2 = (false && true)
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi64 := 0; fi64 < 10; fi64++ {
        if v4 {
            fmt.Println(v15, fi64, v15)
        }
        _ = fi64
    }
    var v65 int8 = 127
    if v2 {
        for fi66 := 0; fi66 < 7; fi66++ {
            fmt.Println(v15)
            _ = fi66
        }
    }
    for fi67 := 0; fi67 < 12; fi67++ {
        v68 := GenS9{F0: 0}
        _ = fi67
        _ = v68
    }
    v69 := []int{-119, -9223372036854775808, -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v69
}
