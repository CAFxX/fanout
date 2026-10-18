package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int16; F1 int64 }
func genFn0(p0 bool) bool {
    v6 := GenS1{F0: -1}
    v7 := []int{32767, 1}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v6
    _ = v7
    return (42 != 42)
}
func genFn1(p0 int32, p1 uint8, p2 string) int {
    p0 = (p0 | -833)
    v8 := []int{2147483647, -992, -602, 1}
    switch ("foo bar" + "a") {
    case "a":
        switch p2 {
        case "a":
            _ = len(v8)
        case "hello":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        default:
            v9 := GenS2{F0: p0, F1: p0, F2: p1}
            _ = v9
        }
    case "":
        v8 = append(v8, -454)
    case "\n\t":
        if (100 != 0) {
            var v10 uint32 = 518
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    v13 := GenS3{F0: (v10 * 1), F1: 18446744073709551615, F2: 359}
                    var v14 uint8 = 2
                    fmt.Println(p2, p1)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                v15 := GenS4{F0: p0}
                var v16 uint16 = 255
                _ = fi11
                _ = v15
                _ = v16
            }
            var v17 uint16 = 42
            _ = v10
            _ = v17
        }
    default:
        if true {
            if ((1.5 - (0.0 + -2.25)) >= 3.14159) {
                var v18 uint16 = 921
                v19 := GenS5{F0: 0, F1: p0, F2: ("a" + p2)}
                _ = v18
                _ = v19
            }
            fmt.Println(p2, p2)
        } else {
            if true {
                v20 := GenS6{F0: (p0 ^ (p0 & p0))}
                p0 -= (p0 & (p0 - 0))
                var v21 uint16 = 32767
                _ = v20
                _ = v21
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            switch 0 {
            case 127:
                p1 -= (p1 >> ((909) & 63))
            case -9223372036854775808:
                for fi22 := 0; fi22 < 5; fi22++ {
                    switch "" {
                    case "hello":
                        p2 = p2
                    case "a":
                        var v23 uint64 = 0
                        _ = v23
                    }
                    _ = fi22
                }
            case 4294967295:
                var v24 uint16 = 42
                _ = v24
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 964
    var v2 uint8 = 1
    v3 := make(map[string]int)
    var v4 bool = (!false)
    var v5 int32 = 42
    v25 := GenS7{F0: 127, F1: 1}
    var v26 int = 1
    fmt.Println(v5)
    v27 := GenS5{F0: uint(v2), F1: (v5 << ((1) & 63)), F2: ""}
    v5 -= -2147483648
    fmt.Println(v1)
    if (v4 && v4) {
        v3[""] = 32767
    } else {
        _ = len(v3)
    }
    if v4 {
        var v28 float64 = (((-2.25 - (((((0.0 / (0.0 - ((3.14159 / 3.14159) + (-2.25 / 1.5)))) - -2.25) / float64(0.1)) / 1e10) + float64((3.14159 / 1e10)))) / 1.0) + 1.5)
        _ = v28
    }
    switch v26 {
    case 1:
        _ = len(v3)
    case 32767:
        if (!v4) {
            if v4 {
                if true {
                    var v29 int32 = (v5 % ((v5) | 1))
                    switch genFn1(0, v2, "foo bar") {
                    case 2147483647:
                        if (!true) {
                            var v30 uint = 255
                            for fi31 := 0; fi31 < 4; fi31++ {
                                for fi32 := 0; fi32 < 4; fi32++ {
                                    switch int(v2) {
                                    case 996:
                                        v4 = true
                                    default:
                                        v33 := v3["\n\t"]
                                        _ = v33
                                    }
                                    _ = fi32
                                }
                                _ = fi31
                            }
                            v34 := GenS8{F0: 0, F1: -215}
                            _ = v30
                            _ = v34
                        }
                    default:
                        var v35 uint32 = 2
                        _ = v35
                    }
                    _ = v29
                } else {
                    v36 := GenS1{F0: 1}
                    var v37 bool = true
                    for fi38 := 0; fi38 < 7; fi38++ {
                        var v39 float64 = (0.1 / ((-2.25 + (-2.25 + float64(0.0))) + 0.0))
                        v26 += (v26 | fi38)
                        _ = fi38
                        _ = v39
                    }
                    _ = v36
                    _ = v37
                }
                switch 585 {
                case 42:
                    delete(v3, "a")
                case 2147483647:
                    v2 = 42
                case -61:
                    delete(v3, "")
                default:
                    for fi40 := 0; fi40 < 10; fi40++ {
                        fmt.Println(v2, v1)
                        v41 := []int{9223372036854775807, 128}
                        _ = fi40
                        _ = v41
                    }
                }
                v42 := []int{(v26 / ((2147483647) | 1))}
                _ = v42
            }
        }
    case 720:
        var v43 int64 = 0
        _ = v43
    }
    switch v26 {
    case 127:
        v44 := GenS8{F0: -1, F1: 933}
        _ = v44
    case 2147483647:
        v4 = true
    case 9223372036854775807:
        v5 += v5
    default:
        v1 = 274
    }
    switch ("a" + "a") {
    case "hello":
        v1 = v1
    }
    v1 = 4294967295
    v2 -= v2
    v26 = (v26 ^ 2147483647)
    if (0.1 >= -2.25) {
        v5 = (v5 ^ (^(v5)))
        switch (-(v26)) {
        case 42:
            var v45 bool = (0 < 646)
            _ = v45
        default:
            var v46 uint8 = v2
            _ = v46
        }
        v1 += uint64(v2)
    }
    var v47 uint16 = 520
    var v48 int = 9223372036854775807
    var v49 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
