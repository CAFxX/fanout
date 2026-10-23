package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int32) bool {
    switch 592 {
    case 42:
        v7 := []int{-2147483648, -659, 127}
        _ = v7
    case 127:
        switch ("\n\t" + ("foo bar" + "hello")) {
        case "a":
            p0 = p0
        case "foo bar":
            var v8 uint8 = 94
            _ = v8
        default:
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 bool = (true && true)
                _ = fi9
                _ = v10
            }
        }
    }
    var v11 int = 4294967295
    v11 = 72
    _ = p0
    _ = v11
    return true
}
func genFn1(p0 uint, p1 uint32) int {
    var v12 uint32 = (p1 << ((255) & 63))
    v12 = p1
    var v13 int8 = 5
    p0 += p0
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 830
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 int32 = 0
    var v4 int16 = -228
    v5 := make(map[uint32]bool)
    v6 := []uint32{71, 4294967295, 556, 599}
    v3 = 1
    if false {
        if true {
            if false {
                var v14 string = "a"
                v2["a"] = 328
                if true {
                    fmt.Println(v3)
                }
                _ = v14
            } else {
                fmt.Println(v4, v3, v4)
                v1 += v1
            }
            v15 := v5[1]
            _ = v15
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS1{F0: 0, F1: 128, F2: "a"}
                _ = fi16
                _ = v17
            }
            var v18 uint32 = 115
            v19 := GenS2{F0: 2}
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: 1, F1: 4, F2: true}
        v21 := v5[281]
        _ = v20
        _ = v21
    } else {
        var v22 int8 = 0
        if (255 != 1) {
            var v23 int16 = -692
            var v24 int = 4294967295
            var v25 int32 = (-(v3))
            _ = v23
            _ = v24
            _ = v25
        } else {
            if len(v6) > 0 {
                v6[0] = 1
            }
            var v26 float32 = 1.5
            _ = v26
        }
        _ = v22
    }
    if (255 < 771) {
        switch "" {
        case "foo bar":
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint8 = 253
                _ = fi27
                _ = v28
            }
        default:
            fmt.Println(v1, v1)
        }
    } else {
        v4 = v4
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v29 := GenS1{F0: int8(v1), F1: 208, F2: "\n\t"}
    var v30 uint64 = 42
    v4 = -32768
    switch "hello" {
    case "\n\t":
        if (32767 > 759) {
            if (!false) {
                v31 := GenS1{F0: 0, F1: 42, F2: ("hello" + "")}
                fmt.Println(v1, v3)
                _ = v31
            } else {
                fmt.Println(v3, v3)
                if true {
                    var v32 uint8 = 127
                    if (0.1 >= 3.14159) {
                        v33 := GenS1{F0: 100, F1: 1, F2: ""}
                        v30 -= 851
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 11; fi34++ {
                        switch fi34 {
                        case -9223372036854775808:
                            v35 := GenS3{F0: fi34, F1: v30, F2: (-2.25 == (0.1 / 0.1))}
                            _ = v35
                        }
                        _ = fi34
                    }
                    _ = v32
                }
            }
            v36 := GenS4{F0: int8(v4), F1: 0}
            var v37 int32 = -1
            _ = v36
            _ = v37
        } else {
            if (!false) {
                delete(v2, "a")
                var v38 float32 = (1e10 * 3.14159)
                switch -980 {
                case -921:
                    var v39 uint32 = 505
                    _ = v39
                case 9223372036854775807:
                    v2["\n\t"] = 387
                default:
                    delete(v2, "\n\t")
                }
                _ = v38
            } else {
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
            }
            v40 := GenS4{F0: 127, F1: 140}
            var v41 float32 = (3.14159 / float32(-2.25))
            _ = v40
            _ = v41
        }
    case "hello":
        v42 := GenS3{F0: 9223372036854775807, F1: (v30 - v30), F2: (!true)}
        _ = v42
    default:
        fmt.Println(v1)
    }
    switch "" {
    case "a":
        v6 = append(v6, 4294967295)
    case "":
        if (0 >= 1) {
            v2[("hello" + ("\n\t" + ("a" + "foo bar")))] = 341
            if false {
                if (4294967295 < (^(v30))) {
                    v1 = v1
                    v6 = append(v6, 927)
                }
                var v43 int = -9223372036854775808
                v4 += (v4 % ((v4) | 1))
                _ = v43
            } else {
                v6 = append(v6, 780)
                if ((v4 > (-(v4))) || (2 == 241)) {
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v45 := GenS5{F0: v30}
                        _ = fi44
                        _ = v45
                    }
                }
            }
        }
    case "hello":
        v3 = v3
    }
    switch "hello" {
    case "\n\t":
        var v46 bool = true
        _ = v46
    default:
        v47 := v5[463]
        _ = v47
    }
    delete(v5, 400)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
