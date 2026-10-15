package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 uint16; F2 int8 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 int; F1 string; F2 int64 }
func genFn0() int8 {
    var v7 uint32 = 2
    if false {
        v8 := []int{42, 127, 42, 127}
        _ = v8
    } else {
        switch 40 {
        case 4294967295:
            switch "foo bar" {
            case "hello":
                for fi9 := 0; fi9 < 1; fi9++ {
                    v7 = 2147483648
                    fmt.Println(v7, fi9)
                    _ = fi9
                }
            case "\n\t":
                switch "\n\t" {
                case "":
                    switch 128 {
                    case 42:
                        v10 := GenS1{F0: v7}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 12; fi11++ {
                            var v12 uint16 = 65535
                            if (2 <= 42) {
                                var v13 int16 = -32768
                                var v14 uint16 = (^(v12))
                                _ = v13
                                _ = v14
                            } else {
                                if (607 <= -361) {
                                    fmt.Println(fi11)
                                    for fi15 := 0; fi15 < 17; fi15++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v16 float32 = 0.0
                                            _ = v16
                                        default:
                                            v17 := GenS2{F0: -872, F1: 962, F2: -9223372036854775808}
                                            _ = v17
                                        }
                                        _ = fi15
                                    }
                                    if true {
                                        v18 := GenS2{F0: -32768, F1: v7, F2: 138}
                                        var v19 float64 = 3.14159
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v20 := []int{fi11, 0, fi11, fi11}
                                        _ = v20
                                    }
                                } else {
                                    if (1 >= -1) {
                                        switch "hello" {
                                        case "":
                                            if ((float64((0.0 - -2.25)) * -2.25) == 3.14159) {
                                                v7 = v7
                                                v21 := GenS3{F0: v7, F1: -1}
                                                v22 := GenS4{F0: -729}
                                                _ = v21
                                                _ = v22
                                            } else {
                                                v12 += (v12 >> ((65535) & 63))
                                                var v23 int16 = 42
                                                var v24 float32 = 0.1
                                                _ = v23
                                                _ = v24
                                            }
                                        case "\n\t":
                                            var v25 int16 = -413
                                            _ = v25
                                        case "hello":
                                            v26 := []int{fi11, fi11}
                                            _ = v26
                                        default:
                                            v27 := []int{32768, 9223372036854775807, 1}
                                            _ = v27
                                        }
                                    }
                                    for fi28 := 0; fi28 < 4; fi28++ {
                                        switch (^(fi11)) {
                                        case -1:
                                            v29 := []int{fi11, (fi28 ^ fi28), (fi28 << ((465) & 63))}
                                            _ = v29
                                        case -2147483648:
                                            v30 := GenS5{F0: -9223372036854775808, F1: 1, F2: 73}
                                            _ = v30
                                        case 2147483647:
                                            v31 := []int{128, (fi11 << ((255) & 63))}
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                            _ = v31
                                        }
                                        _ = fi28
                                    }
                                }
                                fmt.Println(v7, v12)
                                v32 := GenS3{F0: (v7 % ((v7) | 1)), F1: 42}
                                _ = v32
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                case "a":
                    v7 = v7
                default:
                    v33 := GenS4{F0: 0}
                    _ = v33
                }
            case "":
                v7 = v7
            default:
                v7 = (v7 >> ((42) & 63))
            }
        case 128:
            v34 := GenS5{F0: -103, F1: 2, F2: -1}
            _ = v34
        case 32768:
            switch 128 {
            case -8:
                var v35 float64 = 0.0
                _ = v35
            default:
                switch -987 {
                case 1:
                    v36 := []int{-1, 9223372036854775807, -41, 290}
                    _ = v36
                case 42:
                    v37 := GenS3{F0: (^(v7)), F1: 42}
                    _ = v37
                }
            }
        }
    }
    v38 := []int{9223372036854775807, -889, 9223372036854775807}
    _ = len(v38)
    _ = v38
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 int, p2 int32) int {
    if false {
        p1 = (p1 + p1)
        v39 := GenS4{F0: -1}
        _ = v39
    } else {
        v40 := []int{-2147483648}
        if true {
            var v41 int32 = 1
            if len(v40) > 0 {
                v40[0] = 727
            }
            var v42 uint = 65535
            _ = v41
            _ = v42
        }
        var v43 int8 = 0
        _ = v40
        _ = v43
    }
    var v44 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    return (p1 & p1)
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 170
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    var v5 uint16 = 65535
    var v6 uint32 = v1
    if true {
        fmt.Println(v1, v6)
    }
    v3 = 3.14159
    v45 := []int{1}
    if true {
        var v46 float32 = ((0.0 - (0.0 / 1.0)) - v3)
        var v47 bool = true
        var v48 float32 = (1.5 / (v3 * v46))
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println(v2)
    fmt.Println(v5, v2)
    fmt.Println(v1)
    v49 := GenS6{F0: 32767, F1: 0}
    var v50 int32 = (v2 + v2)
    if (true || true) {
        v1 = v1
        v51 := GenS7{F0: 4294967295, F1: "\n\t", F2: -266}
        switch 970 {
        case -807:
            v52 := GenS4{F0: 42}
            _ = v52
        case 32767:
            var v53 int32 = v2
            _ = v53
        case -2147483648:
            var v54 int32 = (v2 + -895)
            _ = v54
        }
        _ = v51
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 uint64 = 390
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
        v6 = (v1 / ((v6) | 1))
        _ = fi55
        _ = v56
    }
    v6 = 2
    var v57 uint = 255
    v1 += (v6 << (((^(v57))) & 63))
    v5 -= (v5 << ((2) & 63))
    var v58 uint = (v57 - v57)
    v1 = (v1 * (v1 | v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v57
    _ = v58
    _ = v6
}
