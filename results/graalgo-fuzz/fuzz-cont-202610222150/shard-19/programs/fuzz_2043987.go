package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int64; F2 uint }
type GenS7 struct { F0 bool; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS9 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 uint8) int {
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            switch fi5 {
            case 931:
                switch ("" + ("a" + "")) {
                case "a":
                    var v7 bool = false
                    _ = v7
                case "foo bar":
                    fmt.Println(fi5)
                default:
                    var v8 float32 = 0.0
                    _ = v8
                }
            }
            _ = fi6
        }
        switch fi5 {
        case 369:
            p0 -= p0
        case 338:
            switch (fi5 ^ (fi5 + 2147483647)) {
            case 1:
                v9 := []int{fi5, fi5, fi5}
                _ = v9
            case 57:
                p0 = 42
            }
        }
        v10 := GenS1{F0: 127, F1: 0, F2: "a"}
        _ = fi5
        _ = v10
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 2147483647
}
func genFn1(p0 string, p1 int64, p2 uint16) string {
    var v11 int32 = 952
    if true {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 float64 = 1.5
            _ = fi12
            _ = v13
        }
        v14 := []int{}
        if true {
            v15 := GenS2{F0: 127, F1: false, F2: 548}
            _ = v15
        }
        _ = v14
    }
    v16 := []int{-136, 872}
    v17 := GenS2{F0: 1, F1: true, F2: 2147483648}
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v16
    _ = v17
    return ""
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -1
    var v3 int64 = -1
    var v4 string = ("" + "a")
    var v18 int32 = -1
    var v19 float64 = (-2.25 / 1.0)
    v3 += (v3 ^ v3)
    v20 := GenS3{F0: 2}
    v21 := []int{803, genFn0(254), 32767, 9223372036854775807}
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            fmt.Println(fi22, v18)
            fmt.Println(fi23)
            switch "hello" {
            case "a":
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
            _ = fi23
        }
        var v24 uint64 = 747
        v25 := GenS4{F0: int16(v1), F1: v1}
        _ = fi22
        _ = v24
        _ = v25
    }
    v26 := GenS5{F0: true}
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS5{F0: true}
        fmt.Println(v18, v4)
        _ = fi27
        _ = v28
    }
    var v29 int16 = 0
    _ = len(v21)
    v30 := GenS3{F0: 74}
    v31 := GenS2{F0: -2147483648, F1: (!((((v1 | (v1 % ((v1) | 1))) == v1) && false) || false)), F2: 65535}
    fmt.Println(v29, v29)
    for fi32 := 0; fi32 < 2; fi32++ {
        switch (-(fi32)) {
        case -2147483648:
            switch v4 {
            case "foo bar":
                if (980 > 42) {
                    v29 = int16(v3)
                } else {
                    v33 := GenS6{F0: (42 > -1), F1: v3, F2: 42}
                    v34 := GenS7{F0: false, F1: 128, F2: 1}
                    if true {
                        for fi35 := 0; fi35 < 15; fi35++ {
                            v36 := GenS8{F0: 128, F1: v1, F2: (v29 * v29)}
                            if false {
                                v4 = v4
                                if true {
                                    v37 := GenS9{F0: v18, F1: -1, F2: (false || (42 != 68))}
                                    _ = v37
                                } else {
                                    v18 -= (v18 >> ((0) & 63))
                                    v2 = (-(v2))
                                    switch (fi35 >> ((0) & 63)) {
                                    case 42:
                                        v38 := GenS5{F0: false}
                                        _ = v38
                                    case 821:
                                        var v39 int = genFn0(uint8(v1))
                                        _ = v39
                                    default:
                                        if len(v21) > 0 {
                                            v21[0] = 2147483647
                                        }
                                    }
                                    if false {
                                        var v40 int8 = 127
                                        _ = v40
                                    } else {
                                        v21 = append(v21, genFn0(231))
                                    }
                                }
                                if (!false) {
                                    var v41 uint64 = v1
                                    var v42 uint = 2
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 uint = 0
                                    if false {
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    }
                                    v44 := GenS7{F0: (!false), F1: 2, F2: (v18 & 1)}
                                    _ = v43
                                    _ = v44
                                }
                            } else {
                                var v45 int64 = 42
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                v45 -= (v3 ^ v3)
                                _ = v45
                            }
                            var v46 uint32 = 571
                            _ = fi35
                            _ = v36
                            _ = v46
                        }
                    } else {
                        v47 := GenS7{F0: true, F1: 255, F2: int32(v1)}
                        if (635 <= 255) {
                            v21 = append(v21, -1)
                        }
                        _ = v47
                    }
                    _ = v33
                    _ = v34
                }
            case "\n\t":
                if true {
                    switch (fi32 - fi32) {
                    case -9223372036854775808:
                        v48 := GenS6{F0: true, F1: v3, F2: 42}
                        _ = v48
                    }
                } else {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    v49 := GenS8{F0: 255, F1: v1, F2: v29}
                    var v50 uint8 = 255
                    _ = v49
                    _ = v50
                }
            case "":
                v51 := GenS6{F0: false, F1: (v3 - (v3 ^ 863)), F2: 0}
                _ = v51
            }
        default:
            for fi52 := 0; fi52 < 18; fi52++ {
                v53 := GenS3{F0: 252}
                if false {
                    fmt.Println(v18, v3, v2)
                } else {
                    var v54 uint32 = 65535
                    v19 += (((0.1 * (((v19 + 0.0) / v19) / (3.14159 + v19))) + v19) * ((-2.25 - 1e10) * v19))
                    v19 += v19
                    _ = v54
                }
                _ = len(v21)
                _ = fi52
                _ = v53
            }
        }
        var v55 uint = 65535
        _ = fi32
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
