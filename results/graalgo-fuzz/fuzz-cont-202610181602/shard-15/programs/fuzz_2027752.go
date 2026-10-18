package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 int8; F1 uint8; F2 int64 }
func genFn0(p0 int16, p1 uint8) uint {
    p1 = (p1 >> ((42) & 63))
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            switch "" {
            case "\n\t":
                p1 += p1
            default:
                v9 := GenS1{F0: false, F1: p0}
                _ = v9
            }
            p1 = (p1 % ((1) | 1))
            _ = fi8
        }
        v10 := []int{(fi7 - (fi7 >> ((65535) & 63))), fi7, (fi7 >> ((65535) & 63)), (fi7 + 4294967295)}
        p0 = (p0 % ((307) | 1))
        _ = fi7
        _ = v10
    }
    var v11 int64 = -1
    v12 := GenS2{F0: (^(v11)), F1: 65535}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 18446744073709551615
}
func genFn1(p0 int, p1 int) int16 {
    p0 = p1
    v13 := GenS3{F0: "a"}
    v14 := []int{p0, p0, p0, p1}
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]int)
    v3 := []int{2147483647, 160, -134, 435}
    var v4 int = 32768
    var v5 uint64 = 2
    v6 := make(map[uint32]uint32)
    var v15 uint8 = 26
    if (1 != 65535) {
        v16 := GenS4{F0: v4, F1: 989, F2: 65535}
        if (true || true) {
            var v17 bool = (!true)
            _ = v17
        }
        _ = v16
    }
    v15 = (^(v1))
    fmt.Println(v5, v4)
    v18 := GenS5{F0: int16(v4), F1: 2}
    v19 := GenS6{F0: 1}
    fmt.Println(v1, v15)
    v4 += v4
    if (false && ((0.1 - 1.5) == (-2.25 - (1.5 - 3.14159)))) {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 uint8 = v1
            switch (("\n\t" + "foo bar") + "hello") {
            case "":
                if true {
                    switch (v4 | 116) {
                    case -1:
                        if ((0.0 + 1.5) >= float64(1.5)) {
                            fmt.Println(v15, v15)
                            v5 = 0
                        } else {
                            var v22 uint32 = 42
                            if (!true) {
                                v22 = v22
                                v15 += (v1 / ((127) | 1))
                            } else {
                                v23 := v2[(^(v22))]
                                _ = v23
                            }
                            v3 = append(v3, fi20)
                            _ = v22
                        }
                    case 800:
                        if (!false) {
                            var v24 int16 = 861
                            v1 -= v1
                            _ = v24
                        } else {
                            var v25 uint64 = 2
                            v1 += (v1 >> ((4294967295) & 63))
                            _ = v25
                        }
                    case 2147483647:
                        var v26 int32 = -142
                        _ = v26
                    default:
                        switch "\n\t" {
                        case "a":
                            var v27 int64 = -65
                            _ = v27
                        case "foo bar":
                            switch (fi20 & 852) {
                            case 1:
                                var v28 int64 = -1
                                _ = v28
                            }
                        case "\n\t":
                            if true {
                                var v29 uint32 = 42
                                v6[v29] = v29
                                _ = len(v2)
                                _ = v29
                            } else {
                                switch -1 {
                                case 133:
                                    v2[222] = (v4 * 128)
                                }
                                switch "" {
                                case "":
                                    v21 = 42
                                }
                                var v30 int64 = 42
                                _ = v30
                            }
                        }
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        fmt.Println(v4, v4)
                        var v32 int32 = 0
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 17; fi33++ {
                        var v34 string = "a"
                        if (((^(v5)) >= (v5 + v5)) || true) {
                            fmt.Println(v1, v5, v5)
                            var v35 uint16 = 32767
                            fmt.Println(v5, v15, fi33)
                            _ = v35
                        }
                        var v36 uint64 = (^(v5))
                        v37 := GenS2{F0: 1, F1: 714}
                        v36 += (v36 ^ v5)
                        v38 := v6[42]
                        fmt.Println(fi33, v38)
                        var v39 string = v34
                        _ = fi33
                        _ = v34
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
            case "hello":
                for fi40 := 0; fi40 < 9; fi40++ {
                    switch "hello" {
                    case "foo bar":
                        v5 = (v5 - (v5 ^ uint64(v1)))
                    case "hello":
                        fmt.Println(v4)
                    case "a":
                        if false {
                            if (9223372036854775807 <= -9223372036854775808) {
                                v3 = append(v3, (v4 ^ (^(fi40))))
                                switch "foo bar" {
                                case "":
                                    var v41 int = v4
                                    _ = v41
                                case "hello":
                                    _ = len(v6)
                                default:
                                    switch (fi40 / (((fi20 << ((genFn0(int16(v4), v1)) & 63))) | 1)) {
                                    case 205:
                                        v42 := GenS6{F0: -128}
                                        _ = v42
                                    }
                                    if false {
                                        if (v5 > 0) {
                                            v15 -= 255
                                        } else {
                                            var v43 uint32 = 65535
                                            v43 -= (v43 % ((v43) | 1))
                                            _ = v43
                                        }
                                        var v44 float32 = (0.0 - 3.14159)
                                        _ = v44
                                    }
                                    v15 += v15
                                }
                                fmt.Println(v15)
                            } else {
                                v3 = append(v3, fi20)
                                v5 += (v5 & v5)
                            }
                        }
                    default:
                        v1 = (v1 >> ((255) & 63))
                    }
                    if true {
                        v21 = (v21 + (v21 >> ((uint(v5)) & 63)))
                    }
                    v45 := GenS7{F0: v21, F1: 2147483648}
                    _ = fi40
                    _ = v45
                }
            case "\n\t":
                var v46 uint64 = (^(v5))
                _ = v46
            default:
                var v47 uint8 = v21
                _ = v47
            }
            var v48 string = "hello"
            _ = fi20
            _ = v21
            _ = v48
        }
        for fi49 := 0; fi49 < 15; fi49++ {
            if true {
                var v50 uint64 = v5
                _ = v50
            }
            var v51 uint32 = 2147483648
            v52 := GenS2{F0: 215, F1: 255}
            _ = fi49
            _ = v51
            _ = v52
        }
    }
    v53 := GenS8{F0: -128, F1: 45, F2: 42}
    fmt.Println(v1, v4)
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
