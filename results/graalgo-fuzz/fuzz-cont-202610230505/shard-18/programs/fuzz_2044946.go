package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint; F1 int64; F2 uint16 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32) uint16 {
    switch "a" {
    case "foo bar":
        var v5 int64 = 42
        _ = v5
    case "\n\t":
        if ((p0 << ((0) & 63)) == p0) {
            v6 := GenS1{F0: 18446744073709551615}
            v7 := []int{-309, -10, -9223372036854775808}
            _ = v6
            _ = v7
        } else {
            v8 := GenS2{F0: 640, F1: 1, F2: 219}
            _ = v8
        }
    case "a":
        v9 := []int{}
        _ = v9
    }
    var v10 int64 = 0
    switch "foo bar" {
    case "hello":
        v11 := []int{2147483647, -811, 1}
        _ = v11
    default:
        fmt.Println(v10, v10)
    }
    var v12 string = "hello"
    switch ("" + (v12 + "hello")) {
    case "hello":
        v13 := []int{-572, 128}
        _ = v13
    default:
        v14 := []int{15}
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v12
    return 2
}
func genFn1(p0 uint16) bool {
    if true {
        var v15 uint = 4294967295
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 4294967295, F1: 42, F2: (p0 / ((p0) | 1))}
        if true {
            if (255 != p0) {
                if (255 < 65535) {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        p0 -= 234
                        if (4294967295 != 915) {
                            p0 = (p0 / (((^(p0))) | 1))
                        }
                        _ = fi18
                    }
                }
            }
            var v19 float64 = (1.5 * 1.5)
            p0 = p0
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(p0, fi20, fi20)
                if false {
                    switch fi20 {
                    case 354:
                        v21 := []int{fi16, int(p0), 1, (fi20 * fi20)}
                        _ = v21
                    default:
                        if true {
                            var v22 int = 9223372036854775807
                            if (123 > 1) {
                                if true {
                                    fmt.Println(fi16, p0)
                                    switch v22 {
                                    case 128:
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch ("foo bar" + "hello") {
                                            case "foo bar":
                                                var v23 float64 = 1.5
                                                var v24 float32 = -2.25
                                                v24 -= -2.25
                                                _ = v23
                                                _ = v24
                                            case "hello":
                                                v25 := GenS3{F0: 2147483648, F1: -549}
                                                _ = v25
                                            }
                                        case "foo bar":
                                            fmt.Println(p0, v22)
                                        default:
                                            v26 := GenS3{F0: 1, F1: 948}
                                            _ = v26
                                        }
                                    }
                                    var v27 uint16 = (p0 - genFn0(42))
                                    _ = v27
                                } else {
                                    if (908 <= 2) {
                                        var v28 bool = false
                                        _ = v28
                                    } else {
                                        v29 := []int{(fi20 >> ((65535) & 63)), (fi20 >> ((255) & 63)), (v22 >> ((42) & 63)), 32767}
                                        _ = v29
                                    }
                                }
                            }
                            _ = v22
                        } else {
                            var v30 uint32 = uint32(p0)
                            var v31 int8 = -1
                            v31 -= (v31 >> ((849) & 63))
                            _ = v30
                            _ = v31
                        }
                    }
                }
                var v32 uint = 2
                _ = fi20
                _ = v32
            }
        }
        _ = fi16
        _ = v17
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        if (!(4294967295 != 2147483648)) {
            switch "\n\t" {
            case "hello":
                var v34 int64 = -1
                _ = v34
            case "foo bar":
                var v35 int32 = 0
                _ = v35
            case "\n\t":
                switch fi33 {
                case 2147483647:
                    p0 += 1
                case -2147483648:
                    p0 = (p0 & p0)
                case 1:
                    switch "\n\t" {
                    case "hello":
                        p0 = 65535
                    }
                default:
                    for fi36 := 0; fi36 < 12; fi36++ {
                        if true {
                            var v37 bool = (int16(p0) != 42)
                            var v38 uint = 18446744073709551615
                            fmt.Println(p0)
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS1{F0: 2}
                            fmt.Println(p0, p0, fi33)
                            _ = v39
                        }
                        _ = fi36
                    }
                }
            }
        }
        var v40 int16 = -1
        var v41 uint64 = 2
        _ = fi33
        _ = v40
        _ = v41
    }
    _ = p0
    return (!false)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    var v42 uint16 = 65535
    _ = len(v4)
    fmt.Println(v42, v1, v42)
    for fi43 := 0; fi43 < 2; fi43++ {
        v2 = float64(-2.25)
        var v44 int32 = 0
        _ = fi43
        _ = v44
    }
    v45 := GenS4{F0: 644, F1: 765, F2: (v42 * v42)}
    v46 := GenS3{F0: 1, F1: 0}
    v1 = 65535
    var v47 uint32 = v1
    delete(v3, 1)
    var v48 int16 = 0
    var v49 float32 = 0.0
    var v50 int32 = 42
    for fi51 := 0; fi51 < 7; fi51++ {
        for fi52 := 0; fi52 < 7; fi52++ {
            v50 -= 2147483647
            _ = fi52
        }
        v53 := GenS5{F0: ((44 >= 255) && (false || (false && false)))}
        v50 = 350
        _ = fi51
        _ = v53
    }
    var v54 string = "hello"
    if ("" > v54) {
        v55 := []int{9223372036854775807, 620, 635, 4294967295}
        v42 += (v42 >> ((319) & 63))
        var v56 float32 = float32(v2)
        _ = v55
        _ = v56
    }
    var v57 int = 4294967295
    var v58 uint16 = (v42 & 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v54
    _ = v57
    _ = v58
}
