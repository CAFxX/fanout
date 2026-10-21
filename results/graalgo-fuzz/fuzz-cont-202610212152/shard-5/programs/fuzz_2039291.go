package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int64, p1 uint32) string {
    v7 := GenS1{F0: 42, F1: 255, F2: -875}
    p0 = p0
    v8 := GenS2{F0: 1, F1: true, F2: 2}
    var v9 uint = 4294967295
    var v10 int32 = 70
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("hello" + "hello")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 int8 = 55
    var v4 float64 = (0.0 - 0.1)
    var v5 float32 = 3.14159
    v6 := []int{-168, 0, 42}
    v5 -= (-2.25 / v5)
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            switch ("foo bar" + (("" + ((("foo bar" + ((("a" + "a") + ("" + "")) + "hello")) + "\n\t") + "\n\t")) + "\n\t")) {
            case "\n\t":
                if true {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if (!(true && ((!false) || (42 >= 2147483647)))) {
                            if true {
                                v13 := GenS1{F0: 159, F1: 637, F2: fi12}
                                v1[65535] = fi11
                                _ = v13
                            }
                            v14 := v1[65535]
                            _ = v14
                        } else {
                            v15 := GenS3{F0: -32768, F1: -370, F2: 42}
                            var v16 uint64 = 1
                            _ = v15
                            _ = v16
                        }
                        _ = fi12
                    }
                    switch "a" {
                    case "":
                        if false {
                            if (int32(v2) != 2147483647) {
                                var v17 uint8 = (^(v2))
                                _ = v17
                            }
                            v2 = (v2 >> ((0) & 63))
                            var v18 int = (fi11 / (((fi11 << ((32) & 63))) | 1))
                            _ = v18
                        } else {
                            var v19 bool = false
                            _ = v19
                        }
                    case "a":
                        switch fi11 {
                        case 32767:
                            v5 = float32(v4)
                        case -2147483648:
                            if len(v6) > 0 {
                                v6[0] = fi11
                            }
                        case 42:
                            for fi20 := 0; fi20 < 14; fi20++ {
                                v21 := v1[2147483647]
                                _ = fi20
                                _ = v21
                            }
                        default:
                            v5 += (v5 + (v5 + (((3.14159 / v5) / v5) / (v5 * float32(v4)))))
                        }
                    case "\n\t":
                        v1[4294967295] = (fi11 % ((42) | 1))
                    }
                } else {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        var v23 bool = true
                        v3 -= -24
                        _ = fi22
                        _ = v23
                    }
                    var v24 int8 = (v3 >> ((1) & 63))
                    var v25 int = 1
                    _ = v24
                    _ = v25
                }
            case "foo bar":
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 uint8 = 0
                    v28 := GenS4{F0: 1}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                v3 = v3
            }
            if (-1 != (fi11 ^ fi11)) {
                switch "\n\t" {
                case "hello":
                    v29 := GenS1{F0: 4294967295, F1: 255, F2: (fi11 >> ((18446744073709551615) & 63))}
                    _ = v29
                case "foo bar":
                    switch -937 {
                    case 59:
                        var v30 int = -2147483648
                        _ = v30
                    case -1:
                        v4 = v4
                    case 487:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        var v31 float32 = (-2.25 / -2.25)
                        _ = v31
                    }
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 int = fi11
                    v34 := GenS3{F0: 1000, F1: 708, F2: 2}
                    var v35 int = int(v3)
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                v5 = 3.14159
            } else {
                v5 = v5
            }
        }
        if (1 > 295) {
            v36 := GenS5{F0: 248}
            _ = v36
        } else {
            var v37 uint16 = 48
            _ = v37
        }
        _ = fi11
    }
    v6 = append(v6, 128)
    if (795 >= -9223372036854775808) {
        fmt.Println(v4)
        var v38 float32 = 1.5
        if true {
            var v39 bool = false
            v40 := GenS1{F0: 1, F1: 32767, F2: 87}
            var v41 uint32 = 2147483647
            _ = v39
            _ = v40
            _ = v41
        } else {
            var v42 float32 = 0.0
            _ = v42
        }
        _ = v38
    } else {
        for fi43 := 0; fi43 < 11; fi43++ {
            for fi44 := 0; fi44 < 2; fi44++ {
                var v45 int = (-(fi44))
                _ = fi44
                _ = v45
            }
            if (368 == 2147483647) {
                switch "a" {
                case "a":
                    _ = len(v6)
                case "\n\t":
                    v4 += (1.5 + 1.5)
                case "hello":
                    var v46 string = "\n\t"
                    _ = v46
                default:
                    var v47 int16 = 1000
                    _ = v47
                }
                if false {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        var v49 uint64 = 0
                        _ = fi48
                        _ = v49
                    }
                    var v50 uint = 42
                    _ = v50
                }
            }
            var v51 int8 = 100
            _ = fi43
            _ = v51
        }
    }
    switch 42 {
    case 9223372036854775807:
        v4 = (1.5 - v4)
    default:
        var v52 bool = (!(1.5 == v4))
        _ = v52
    }
    _ = len(v6)
    var v53 uint32 = 468
    var v54 uint32 = (v53 << ((455) & 63))
    if true {
        switch "\n\t" {
        case "hello":
            v54 -= v53
        case "":
            if (-306 >= -210) {
                v5 -= v5
            }
        case "a":
            var v55 int16 = 42
            _ = v55
        }
        v1[v54] = -9223372036854775808
    }
    v3 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
