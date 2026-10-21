package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int }
func genFn0(p0 int16, p1 uint8, p2 uint64) int {
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 1 {
    case 2147483647:
        v7 := GenS1{F0: 127, F1: -1}
        _ = v7
    case -763:
        v6 = append(v6, -2147483648)
    case 622:
        var v8 int = 118
        _ = v8
    default:
        v9 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v9
    }
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 880
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 uint32 = 2
    var v4 bool = (42 != -32768)
    v5 := []int{127, 4294967295, 9223372036854775807}
    switch -9223372036854775808 {
    case 127:
        v3 = v3
    }
    v10 := GenS2{F0: 2, F1: 143, F2: 42}
    var v11 uint8 = 0
    v11 += v11
    var v12 int = 32768
    if false {
        var v13 uint16 = 0
        if (1 > -32768) {
            v14 := GenS2{F0: v11, F1: (^(v11)), F2: v13}
            v15 := GenS1{F0: (v12 % ((32767) | 1)), F1: 629}
            _ = v14
            _ = v15
        }
        _ = v13
    } else {
        if false {
            v11 = (v11 + 127)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v16 int8 = 100
    v17 := GenS3{F0: 1, F1: "a", F2: 42}
    var v18 int64 = 0
    fmt.Println(v18)
    v11 = v11
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            if false {
                if v4 {
                    v4 = (0 >= 2)
                    if len(v2) > 0 {
                        v2[0] = (fi19 | int(v16))
                    }
                    switch 32768 {
                    case 127:
                        if (-1 < -32768) {
                            var v21 bool = ((2 > 42) || v4)
                            _ = v21
                        } else {
                            fmt.Println(fi19, v11)
                            var v22 int32 = 929
                            fmt.Println(v4, v12)
                            _ = v22
                        }
                        if v4 {
                            switch (("\n\t" + "hello") + "\n\t") {
                            case "":
                                if ((v12 >> ((856) & 63)) > (^(v12))) {
                                    v23 := GenS4{F0: 337, F1: (v4 && v4), F2: -9223372036854775808}
                                    var v24 uint8 = 255
                                    v25 := GenS2{F0: 177, F1: (v11 >> ((484) & 63)), F2: 42}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    if false {
                                        var v26 uint32 = (^(v3))
                                        fmt.Println(v26, fi19)
                                        switch (fi19 ^ (fi20 << ((uint(v3)) & 63))) {
                                        case -661:
                                            v11 = (^(v11))
                                        }
                                        _ = v26
                                    }
                                    if (243 <= 0) {
                                        switch genFn0(-32768, 176, 0) {
                                        case 9223372036854775807:
                                            var v27 int32 = 689
                                            _ = v27
                                        case 1:
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                delete(v1, v12)
                                            case "a":
                                                v18 = (v18 | 42)
                                            case "":
                                                v3 = (v3 + (^(v3)))
                                            default:
                                                var v28 float32 = -2.25
                                                _ = v28
                                            }
                                        case -62:
                                            var v29 uint8 = v11
                                            _ = v29
                                        }
                                        _ = len(v2)
                                        v16 = (v16 - (-(v16)))
                                        v30 := GenS4{F0: 2, F1: (-2.25 <= (float32((-2.25 / 1.0)) / 1e10)), F2: -9223372036854775808}
                                        v31 := GenS2{F0: 127, F1: v11, F2: 32767}
                                        _ = v30
                                        _ = v31
                                    }
                                    switch "" {
                                    case "hello":
                                        var v32 uint64 = uint64(v3)
                                        _ = v32
                                    case "\n\t":
                                        if v4 {
                                            v2 = append(v2, (fi19 * fi19))
                                            v16 -= (v16 & (^(v16)))
                                            v4 = v4
                                        }
                                    }
                                }
                            case "a":
                                var v33 float64 = 3.14159
                                _ = v33
                            case "\n\t":
                                v34 := GenS5{F0: v11}
                                _ = v34
                            default:
                                var v35 string = (("" + "") + ("foo bar" + "hello"))
                                _ = v35
                            }
                        }
                    case 2147483647:
                        var v36 int8 = (v16 * (^(v16)))
                        _ = v36
                    case 0:
                        switch "\n\t" {
                        case "a":
                            var v37 uint64 = 348
                            _ = v37
                        default:
                            v16 = (v16 & (v16 & 59))
                        }
                    }
                }
                v12 += (-(fi20))
                v38 := GenS6{F0: 141, F1: -766}
                _ = v38
            }
            switch fi19 {
            case 4294967295:
                v4 = v4
            case 928:
                if (!(1 < v18)) {
                    fmt.Println(v3, fi20)
                }
            default:
                var v39 uint = 564
                _ = v39
            }
            switch "hello" {
            case "":
                fmt.Println(v12, fi19, v12)
            }
            _ = fi20
        }
        switch "a" {
        case "hello":
            fmt.Println(v16)
        default:
            v40 := v1[(-(fi19))]
            _ = v40
        }
        _ = fi19
    }
    var v41 string = "foo bar"
    v11 -= v11
    if (-32768 != 42) {
        for fi42 := 0; fi42 < 6; fi42++ {
            if len(v2) > 0 {
                v2[0] = (v12 % (((v12 << ((0) & 63))) | 1))
            }
            v3 = (v3 >> ((725) & 63))
            v16 = (v16 - v16)
            _ = fi42
        }
    }
    v16 -= -66
    fmt.Println(v41)
    if len(v2) > 0 {
        v2[0] = (v12 ^ v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
