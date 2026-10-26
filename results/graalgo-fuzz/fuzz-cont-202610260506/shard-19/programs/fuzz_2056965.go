package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
type GenS6 struct { F0 int }
func genFn0() string {
    var v5 bool = (!true)
    var v6 uint = 447
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(fi7, fi7)
        v5 = v5
        _ = fi7
    }
    v8 := GenS1{F0: uint8(v6), F1: 1}
    _ = v5
    _ = v6
    _ = v8
    return "a"
}
func genFn1() int {
    if true {
        if (false || true) {
            var v9 float32 = 1.5
            _ = v9
        } else {
            var v10 uint = 571
            _ = v10
        }
        var v11 int64 = 42
        _ = v11
    }
    if (2147483647 == 65535) {
        var v12 uint16 = 65535
        var v13 int8 = -58
        v14 := []int{0, 128}
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 4294967295
        _ = v15
    }
    v16 := GenS1{F0: 203, F1: 42}
    v17 := []int{}
    _ = v16
    _ = v17
    return -346
}

func main() {
    var v1 string = "a"
    var v2 float64 = -2.25
    v3 := make(map[uint32]int)
    var v4 uint64 = 147
    switch 42 {
    case 128:
        var v18 int32 = 1
        _ = v18
    case -459:
        for fi19 := 0; fi19 < 14; fi19++ {
            switch (fi19 & (fi19 / (((fi19 % ((-1) | 1))) | 1))) {
            case 1:
                for fi20 := 0; fi20 < 17; fi20++ {
                    v21 := []int{-816, (^(fi20)), 98}
                    _ = fi20
                    _ = v21
                }
            }
            delete(v3, 1)
            _ = fi19
        }
    default:
        fmt.Println(v2, v1, v2)
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(fi22, v4, fi22)
            var v24 uint16 = 223
            fmt.Println(v24)
            _ = fi23
            _ = v24
        }
        var v25 uint64 = (^(v4))
        for fi26 := 0; fi26 < 12; fi26++ {
            if false {
                v27 := []int{}
                _ = v27
            }
            _ = fi26
        }
        _ = fi22
        _ = v25
    }
    v1 = v1
    if (false || true) {
        switch v1 {
        case "\n\t":
            fmt.Println(v4, v4)
        }
        var v28 float32 = (0.0 / (1.5 * 3.14159))
        v29 := []int{genFn1(), 0, 32767, 32768}
        _ = v28
        _ = v29
    } else {
        var v30 uint32 = 65535
        _ = v30
    }
    v31 := GenS1{F0: 128, F1: 9223372036854775807}
    if (v1 == v1) {
        v32 := []int{genFn1(), 768, 32768}
        _ = v32
    } else {
        v33 := GenS1{F0: 255, F1: -921}
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS2{F0: (fi34 << ((2) & 63)), F1: v4}
        v3[42] = fi34
        v1 = (v1 + (("\n\t" + "hello") + ""))
        _ = fi34
        _ = v35
    }
    v1 = (v1 + (("a" + (v1 + (((v1 + "\n\t") + "foo bar") + v1))) + "foo bar"))
    var v36 int64 = 1
    var v37 uint = 65535
    v38 := GenS3{F0: (!(v4 >= (v4 ^ (v4 >> ((v37) & 63))))), F1: 2}
    if (!false) {
        v37 += (v37 / (((^(v37))) | 1))
        if true {
            var v39 uint = v37
            if true {
                if true {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v41 := GenS4{F0: "\n\t"}
                        var v42 int64 = -1
                        var v43 uint64 = 18446744073709551615
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    if false {
                        if (2 > 93) {
                            v44 := GenS1{F0: 2, F1: v36}
                            v1 = ((((v1 + v1) + v1) + v1) + ((v1 + ("" + ((v1 + (v1 + ("" + v1))) + v1))) + v1))
                            v45 := GenS1{F0: 255, F1: v36}
                            _ = v44
                            _ = v45
                        }
                    }
                }
                for fi46 := 0; fi46 < 17; fi46++ {
                    if ((v4 ^ 1) == (v4 & (^(v4)))) {
                        v47 := GenS5{F0: (fi46 >> ((v37) & 63)), F1: 102, F2: 1}
                        for fi48 := 0; fi48 < 5; fi48++ {
                            var v49 int8 = 56
                            v36 += 281
                            v3[2147483647] = (fi46 >> (((v37 ^ 4294967295)) & 63))
                            switch (fi46 - fi46) {
                            case -493:
                                v50 := []int{(^(fi46)), 4294967295, genFn1()}
                                v39 -= (v39 / (((^(v37))) | 1))
                                if false {
                                    v50 = append(v50, genFn1())
                                    v51 := GenS5{F0: -2147483648, F1: 4294967295, F2: -461}
                                    _ = v51
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        var v52 uint16 = uint16(v37)
                                        _ = v52
                                    default:
                                        v53 := GenS2{F0: 32767, F1: 0}
                                        _ = v53
                                    }
                                    fmt.Println(v1, fi46, v37)
                                }
                                _ = v50
                            case -462:
                                v54 := []int{352, fi46, (fi46 / ((fi48) | 1)), (-(fi46))}
                                _ = v54
                            default:
                                v55 := GenS5{F0: fi48, F1: 2, F2: -2147483648}
                                _ = v55
                            }
                            var v56 int32 = -1
                            var v57 int32 = (v56 << ((v39) & 63))
                            _ = fi48
                            _ = v49
                            _ = v56
                            _ = v57
                        }
                        _ = v47
                    } else {
                        for fi58 := 0; fi58 < 16; fi58++ {
                            switch v1 {
                            case "a":
                                v37 += 2
                            case "foo bar":
                                if false {
                                    v59 := GenS1{F0: 255, F1: (v36 / ((1) | 1))}
                                    _ = v59
                                }
                            }
                            if true {
                                v37 = 0
                                var v60 uint = 255
                                var v61 int32 = 2147483647
                                var v62 float64 = (float64((1.5 * 0.0)) * v2)
                                v37 = 2
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            v63 := GenS6{F0: fi46}
                            _ = fi58
                            _ = v63
                        }
                    }
                    var v64 int16 = -32768
                    var v65 int8 = 111
                    _ = fi46
                    _ = v64
                    _ = v65
                }
            }
            var v66 uint8 = 1
            _ = v39
            _ = v66
        }
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
