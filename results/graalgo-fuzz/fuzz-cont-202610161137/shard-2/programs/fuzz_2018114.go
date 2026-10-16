package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8; F1 uint16; F2 int }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint; F1 bool; F2 uint64 }
func genFn0() string {
    switch ("" + (("" + "") + "\n\t")) {
    case "foo bar":
        if (true && (-9223372036854775808 >= 1)) {
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint8 = 0
                    fmt.Println(fi6)
                    var v9 bool = true
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                if false {
                    var v10 int32 = 42
                    var v11 bool = false
                    v12 := GenS1{F0: "foo bar", F1: 255, F2: 1}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        var v14 float32 = 0.1
                        _ = fi13
                        _ = v14
                    }
                    var v15 int32 = 1
                    fmt.Println(v15, v15)
                    _ = v15
                }
                _ = fi6
            }
        } else {
            var v16 float64 = (1.5 * 1e10)
            _ = v16
        }
    default:
        if true {
            v17 := []int{-9223372036854775808, 0, 42, 32768}
            var v18 uint8 = 42
            _ = v17
            _ = v18
        }
    }
    var v19 uint16 = 645
    var v20 uint16 = (v19 % ((v19) | 1))
    _ = v19
    _ = v20
    return "\n\t"
}
func genFn1(p0 int8, p1 uint8) int32 {
    fmt.Println(p0)
    p1 = p1
    p0 += 42
    for fi21 := 0; fi21 < 15; fi21++ {
        switch fi21 {
        case -224:
            v22 := GenS1{F0: "\n\t", F1: 4294967295, F2: uint32(p1)}
            _ = v22
        case 201:
            var v23 uint32 = 2147483648
            _ = v23
        case 1:
            var v24 uint32 = 2147483647
            _ = v24
        default:
            if (true || false) {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 bool = (32767 > 0)
                    _ = fi25
                    _ = v26
                }
            } else {
                v27 := []int{(fi21 >> ((1) & 63))}
                p1 += p1
                _ = v27
            }
        }
        var v28 uint8 = p1
        v29 := []int{fi21, (fi21 + (fi21 << ((0) & 63))), (fi21 | (fi21 - -9223372036854775808))}
        _ = fi21
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 203
    var v3 uint = 59
    v4 := make(map[int]uint32)
    var v5 string = "foo bar"
    fmt.Println(v1)
    if (1e10 == (1.5 / (-2.25 + float32(float64(0.0))))) {
        switch 0 {
        case -2147483648:
            var v30 uint8 = (^(v2))
            _ = v30
        case 0:
            v2 += uint8(v3)
        case 683:
            v31 := GenS2{F0: 2, F1: 846, F2: ("" + ("" + (v5 + v5)))}
            _ = v31
        default:
            fmt.Println(v5)
        }
        v2 -= v2
        var v32 uint32 = 2147483647
        _ = v32
    }
    v3 = 0
    v33 := GenS3{F0: 1, F1: (v2 / (((v2 | (^(v2)))) | 1)), F2: ((true && (v1 >= v1)) || false)}
    switch -855 {
    case 9223372036854775807:
        v3 += 42
    case 42:
        v1 += (^(v1))
    case -912:
        switch 864 {
        case -1:
            v1 += (^(v1))
        default:
            if false {
                if ((v3 * v3) < v3) {
                    switch v5 {
                    case "":
                        fmt.Println(v3, v1)
                    case "foo bar":
                        v1 += v1
                    }
                    v1 = 42
                    v2 += v2
                } else {
                    v3 += 427
                    for fi34 := 0; fi34 < 3; fi34++ {
                        v35 := GenS4{F0: v2, F1: (v1 ^ v1), F2: fi34}
                        _ = fi34
                        _ = v35
                    }
                }
                var v36 int64 = -990
                if true {
                    switch v5 {
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            var v37 int16 = -989
                            _ = v37
                        case "hello":
                            var v38 bool = false
                            _ = v38
                        case "foo bar":
                            var v39 uint32 = 2147483648
                            _ = v39
                        }
                    case "":
                        var v40 int32 = 42
                        _ = v40
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        v36 = (v36 >> (((v3 % ((1) | 1))) & 63))
                    }
                }
                _ = v36
            } else {
                fmt.Println(v5, v5)
                v41 := GenS5{F0: 0, F1: -111}
                _ = v41
            }
        }
    default:
        var v42 int = 32768
        _ = v42
    }
    var v43 uint8 = (^(v2))
    v44 := []int{4294967295}
    for fi45 := 0; fi45 < 10; fi45++ {
        v4[fi45] = 2147483647
        v1 += (^(v1))
        _ = fi45
    }
    v46 := GenS1{F0: ("foo bar" + ("\n\t" + v5)), F1: v3, F2: 42}
    v47 := GenS5{F0: 65535, F1: 0}
    v43 = v43
    for fi48 := 0; fi48 < 10; fi48++ {
        v49 := v4[fi48]
        var v50 int = fi48
        _ = fi48
        _ = v49
        _ = v50
    }
    if false {
        for fi51 := 0; fi51 < 2; fi51++ {
            var v52 uint8 = 127
            if (!(0 == 8)) {
                delete(v4, fi51)
                if false {
                    for fi53 := 0; fi53 < 11; fi53++ {
                        var v54 uint8 = v2
                        _ = fi53
                        _ = v54
                    }
                }
                var v55 bool = true
                _ = v55
            }
            v5 = "foo bar"
            _ = fi51
            _ = v52
        }
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        v43 -= 0
        _ = fi56
    }
    fmt.Println(v43)
    v57 := GenS6{F0: 2, F1: false, F2: 18446744073709551615}
    v5 = (v5 + ("a" + ((((v5 + (v5 + v5)) + (v5 + v5)) + "a") + v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v57
}
