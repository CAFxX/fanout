package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int64; F2 string }
func genFn0(p0 int16, p1 uint64, p2 uint64) int {
    if (255 != 702) {
        v6 := GenS1{F0: 0}
        fmt.Println(p1, p2)
        _ = v6
    } else {
        var v7 float64 = ((1e10 - (((0.1 * (1e10 * (1e10 / float64(3.14159)))) * 3.14159) * (1.5 * 3.14159))) * 1e10)
        if (p0 != (p0 >> ((782) & 63))) {
            switch 1 {
            case 32767:
                for fi8 := 0; fi8 < 18; fi8++ {
                    if (!false) {
                        if false {
                            v9 := GenS1{F0: 255}
                            if false {
                                fmt.Println(v7, v7, p1)
                            } else {
                                v10 := []int{(fi8 >> ((995) & 63)), fi8, fi8}
                                var v11 string = "a"
                                _ = v10
                                _ = v11
                            }
                            var v12 float64 = (v7 - float64(0.0))
                            _ = v12
                            _ = v9
                        }
                        switch (fi8 % ((-979) | 1)) {
                        case 32767:
                            if false {
                                var v13 bool = true
                                v14 := GenS1{F0: 1}
                                _ = v13
                                _ = v14
                            }
                        default:
                            if (false && false) {
                                v15 := GenS2{F0: p0, F1: p2, F2: 0}
                                fmt.Println(p0)
                                v16 := GenS1{F0: 28}
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint8 = 232
                                v18 := []int{(fi8 / ((fi8) | 1))}
                                _ = v17
                                _ = v18
                            }
                        }
                    } else {
                        p1 -= 2
                    }
                    switch ("a" + "a") {
                    case "":
                        v19 := GenS3{F0: 42, F1: 32767}
                        _ = v19
                    case "hello":
                        v20 := GenS1{F0: 281}
                        _ = v20
                    case "a":
                        fmt.Println(fi8, fi8)
                    }
                    fmt.Println(fi8, fi8)
                    _ = fi8
                }
            default:
                v21 := GenS4{F0: 2}
                _ = v21
            }
            if (float64(3.14159) != v7) {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 uint32 = 2147483648
                    _ = fi22
                    _ = v23
                }
            } else {
                if false {
                    var v24 bool = (2147483648 >= 2)
                    v25 := []int{0, -585, 1, 9223372036854775807}
                    var v26 int16 = (^(p0))
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            if ((v7 >= ((float64((1e10 - (((float32((((-2.25 + 0.0) / 1.5) * v7)) / 0.1) / 3.14159) / (-2.25 * (float32((v7 * (3.14159 * (float64(0.1) + v7)))) / 1.5))))) / ((0.1 * v7) / -2.25)) * v7)) || (1e10 == v7)) {
                for fi27 := 0; fi27 < 7; fi27++ {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        p2 += (p2 << ((255) & 63))
                        p0 -= p0
                        _ = fi28
                    }
                    _ = fi27
                }
                p1 -= 378
                v29 := []int{-20}
                _ = v29
            } else {
                v30 := []int{-573, 47, 1, 127}
                if len(v30) > 0 {
                    v30[0] = 32768
                }
                _ = v30
            }
        } else {
            v7 -= v7
        }
        fmt.Println(p2)
        _ = v7
    }
    v31 := GenS1{F0: 4294967295}
    switch "\n\t" {
    case "a":
        v32 := GenS5{F0: 817}
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 4294967295
}
func genFn1(p0 int8, p1 uint16, p2 uint64) int8 {
    switch "hello" {
    case "foo bar":
        var v33 uint16 = uint16(p2)
        _ = v33
    case "\n\t":
        var v34 bool = (0 < 1)
        _ = v34
    default:
        if false {
            p2 -= p2
            switch -2147483648 {
            case 4294967295:
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := []int{-691}
                    _ = fi35
                    _ = v36
                }
            case 128:
                var v37 uint8 = 2
                _ = v37
            case 507:
                for fi38 := 0; fi38 < 12; fi38++ {
                    p2 = 268
                    var v39 string = "\n\t"
                    var v40 float64 = float64(1.5)
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            default:
                fmt.Println(p0, p2, p2)
            }
            var v41 int16 = 0
            _ = v41
        } else {
            var v42 float64 = (3.14159 * 0.1)
            var v43 string = (((("hello" + "hello") + "\n\t") + "foo bar") + "a")
            _ = v42
            _ = v43
        }
    }
    v44 := GenS3{F0: 42, F1: -1}
    var v45 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    _ = v45
    return 1
}

func main() {
    var v1 uint16 = 76
    var v2 uint64 = 683
    v3 := make(map[uint32]int)
    var v4 string = "foo bar"
    var v5 float32 = 1.5
    switch ((v4 + (v4 + ("foo bar" + "foo bar"))) + v4) {
    case "foo bar":
        var v46 bool = false
        _ = v46
    }
    if (v2 != 0) {
        var v47 int64 = -9223372036854775808
        v5 = 1.5
        _ = v47
    } else {
        for fi48 := 0; fi48 < 2; fi48++ {
            var v49 int64 = 9223372036854775807
            var v50 bool = (!((fi48 % ((fi48) | 1)) == (fi48 / ((genFn0(-775, v2, v2)) | 1))))
            var v51 uint = 2
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
        }
        if false {
            for fi52 := 0; fi52 < 2; fi52++ {
                if (false || (1.5 == 1.5)) {
                    if (!false) {
                        for fi53 := 0; fi53 < 1; fi53++ {
                            var v54 bool = (true || (!(-115 >= 42)))
                            v5 = (0.0 * v5)
                            v5 = 0.1
                            _ = fi53
                            _ = v54
                        }
                        v3[97] = (fi52 % (((fi52 << ((uint(v1)) & 63))) | 1))
                    }
                    v55 := []int{fi52, (fi52 % (((fi52 / ((95) | 1))) | 1))}
                    _ = v55
                } else {
                    v4 = v4
                }
                _ = fi52
            }
        }
    }
    var v56 bool = true
    for fi57 := 0; fi57 < 7; fi57++ {
        fmt.Println(v1)
        v2 = v2
        _ = fi57
    }
    switch (v4 + (v4 + v4)) {
    case "foo bar":
        v58 := []int{}
        _ = v58
    }
    if v56 {
        v59 := GenS2{F0: -1, F1: v2, F2: (^(v1))}
        v60 := GenS6{F0: true, F1: 25}
        fmt.Println(v5, v1, v56)
        _ = v59
        _ = v60
    } else {
        var v61 string = v4
        _ = v61
    }
    var v62 uint16 = (v1 << ((4294967295) & 63))
    v63 := GenS7{F0: 65535, F1: 42, F2: v4}
    v4 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v62
    _ = v63
}
