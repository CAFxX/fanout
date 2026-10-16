package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 int; F1 string }
func genFn0(p0 uint32, p1 uint64) uint8 {
    var v6 uint = 630
    fmt.Println(v6, v6)
    _ = p0
    _ = p1
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1.5
    var v3 int = -960
    var v4 int8 = 0
    var v5 uint = 559
    v5 -= 4294967295
    fmt.Println(v3)
    var v7 int32 = 42
    _ = len(v1)
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            v9 := []int{(fi8 / (((fi8 % ((fi8) | 1))) | 1)), (fi8 & (v3 * (fi8 + (fi8 / ((v3) | 1))))), (fi8 << (((v5 & v5)) & 63))}
            switch "\n\t" {
            case "foo bar":
                v7 += v7
            case "\n\t":
                if false {
                    v10 := GenS1{F0: 32767, F1: 604}
                    _ = v10
                } else {
                    if len(v9) > 0 {
                        v9[0] = v3
                    }
                    v5 += (v5 + v5)
                }
            case "hello":
                switch v3 {
                case -2147483648:
                    var v11 int32 = int32(v4)
                    _ = v11
                case 2147483647:
                    v12 := GenS2{F0: 1, F1: true, F2: (627 == 0)}
                    _ = v12
                case 151:
                    _ = len(v1)
                default:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 uint32 = 2
                        if (1 < genFn0((v14 + (v14 / (((v14 * (v14 % ((v14) | 1)))) | 1))), 2)) {
                            switch ("a" + ("hello" + "")) {
                            case "foo bar":
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case "hello":
                                if true {
                                    var v15 uint = (v5 - 502)
                                    fmt.Println(v2, v2, fi8)
                                    var v16 uint8 = 255
                                    if true {
                                        v16 = v16
                                        if (299 > 2) {
                                            fmt.Println(v15, v15)
                                            var v17 uint8 = v16
                                            var v18 bool = (32767 < -1)
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v19 := GenS3{F0: 908, F1: -309, F2: (v16 + (v16 / ((v16) | 1)))}
                                            _ = v19
                                        }
                                        switch ("foo bar" + "hello") {
                                        case "":
                                            v16 = v16
                                        case "a":
                                            _ = len(v1)
                                        default:
                                            v14 = v14
                                        }
                                    }
                                    v20 := GenS2{F0: 42, F1: true, F2: (true && (true && true))}
                                    switch ((("" + "foo bar") + "foo bar") + "hello") {
                                    case "hello":
                                        if (false && true) {
                                            var v21 int = fi13
                                            var v22 int8 = 42
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 float32 = (v2 * (1.5 - v2))
                                            _ = v23
                                        }
                                    case "\n\t":
                                        v4 -= v4
                                    case "foo bar":
                                        v15 = v5
                                    default:
                                        switch "" {
                                        case "hello":
                                            if false {
                                                var v24 string = ("a" + "foo bar")
                                                v5 -= v15
                                                _ = v24
                                            } else {
                                                var v25 float32 = 0.0
                                                var v26 float64 = -2.25
                                                if false {
                                                    var v27 bool = ((((v26 - v26) * (0.0 / 1.0)) >= v26) || true)
                                                    _ = v27
                                                }
                                                var v28 uint16 = 984
                                                if (853 >= 1) {
                                                    var v29 int64 = 337
                                                    v30 := GenS4{F0: 4294967295, F1: uint8(v15), F2: 190}
                                                    v31 := GenS5{F0: true}
                                                    v3 = 1
                                                    v9 = append(v9, fi13)
                                                    switch fi8 {
                                                    case 4294967295:
                                                        v32 := GenS3{F0: int64(v7), F1: -654, F2: genFn0(v14, 14)}
                                                        _ = v32
                                                    case 1:
                                                        v29 += v29
                                                    }
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v33 float32 = -2.25
                                                        _ = v33
                                                    }
                                                }
                                                v4 = (^(v4))
                                                if (100 <= (v4 >> ((580) & 63))) {
                                                    if false {
                                                        var v34 int32 = 0
                                                        if len(v9) > 0 {
                                                            v9[0] = int(v28)
                                                        }
                                                        _ = v34
                                                    }
                                                    v35 := GenS6{F0: 0}
                                                    v26 = 1e10
                                                    if true {
                                                        fmt.Println(fi8, v26)
                                                        fmt.Println(fi13, v2, fi13)
                                                        fmt.Println(v7, fi8, v7)
                                                    }
                                                    fmt.Println(v5, v7, v26)
                                                    _ = v35
                                                }
                                                _ = v25
                                                _ = v26
                                                _ = v28
                                            }
                                        }
                                    }
                                    _ = v15
                                    _ = v16
                                    _ = v20
                                }
                            case "\n\t":
                                if ((!(!(1 < 755))) || (2147483647 != 886)) {
                                    var v36 bool = (v5 != (v5 * v5))
                                    var v37 int32 = v7
                                    if (1 >= 1) {
                                        v38 := v1["hello"]
                                        _ = v38
                                    } else {
                                        var v39 int8 = 104
                                        _ = v39
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                v2 += (1e10 / (float32((1.5 - float64(((v2 + 0.0) + 0.0)))) + (v2 / v2)))
                            default:
                                v40 := GenS2{F0: -710, F1: ((!false) && true), F2: (v7 == (^(v7)))}
                                var v41 bool = (v7 == v7)
                                v42 := GenS2{F0: -1, F1: false, F2: v41}
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        }
                        v2 = -2.25
                        fmt.Println(v5)
                        _ = fi13
                        _ = v14
                    }
                }
            default:
                v5 = 0
            }
            _ = v9
        } else {
            v43 := GenS1{F0: 406, F1: 42}
            v44 := GenS7{F0: 32767, F1: 2147483648, F2: 366}
            _ = v43
            _ = v44
        }
        switch "hello" {
        case "":
            v2 += v2
        case "hello":
            var v45 uint16 = 45
            _ = v45
        default:
            fmt.Println(v4)
        }
        _ = fi8
    }
    var v46 int = v3
    switch "\n\t" {
    case "foo bar":
        v7 -= v7
    case "":
        v4 -= v4
    case "a":
        var v47 uint64 = 1
        _ = v47
    default:
        if (!(!((v7 % ((v7) | 1)) >= v7))) {
            if (255 >= 117) {
                v48 := GenS3{F0: 632, F1: int64(v5), F2: 128}
                var v49 string = ""
                _ = v48
                _ = v49
            } else {
                for fi50 := 0; fi50 < 6; fi50++ {
                    v46 = (v46 | (fi50 ^ (fi50 * 431)))
                    fmt.Println(v46)
                    delete(v1, "")
                    _ = fi50
                }
                v2 = ((float32(0.0) / v2) - 3.14159)
            }
        }
    }
    switch v3 {
    case 0:
        v51 := GenS8{F0: -374, F1: (v5 % (((v5 + (v5 * 1))) | 1))}
        _ = v51
    case -187:
        v2 = v2
    case 860:
        if (v4 < (v4 >> (((v5 * v5)) & 63))) {
            if ((3.14159 - 1.5) <= (-2.25 + 1.5)) {
                if (!(2147483648 == 0)) {
                    if false {
                        if false {
                            v2 += (v2 / -2.25)
                            v3 += int(v4)
                        }
                        v5 = (^(v5))
                    } else {
                        switch v3 {
                        case 42:
                            v4 -= (v4 + v4)
                        }
                        var v52 int = 4294967295
                        v53 := GenS9{F0: 155, F1: (v4 ^ v4)}
                        _ = v52
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 4; fi54++ {
                        v55 := GenS8{F0: -32768, F1: (v5 & (v5 << ((v5) & 63)))}
                        _ = fi54
                        _ = v55
                    }
                    v56 := []int{(v3 >> (((v5 & (^(v5)))) & 63))}
                    _ = v56
                }
                for fi57 := 0; fi57 < 20; fi57++ {
                    v58 := GenS10{F0: v46, F1: ""}
                    v46 += -785
                    _ = fi57
                    _ = v58
                }
            }
        }
    default:
        fmt.Println(v4, v4, v3)
    }
    v59 := GenS10{F0: (v46 & v3), F1: "\n\t"}
    v60 := v1["\n\t"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v59
    _ = v60
    _ = v7
}
