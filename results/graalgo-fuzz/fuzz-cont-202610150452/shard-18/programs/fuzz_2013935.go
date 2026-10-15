package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 string) string {
    p0 = p0
    v7 := GenS1{F0: -596}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            if (2 <= 255) {
                fmt.Println(fi8)
                var v10 int8 = 81
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return (p0 + p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 677
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    var v6 uint32 = 395
    v6 += v6
    v6 = (v6 / ((v6) | 1))
    v2 = (v2 * (v2 & 2147483647))
    switch 2147483647 {
    case 0:
        for fi11 := 0; fi11 < 8; fi11++ {
            if (-172 < -903) {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v2 -= v2
                    _ = fi12
                }
                v6 -= v6
                switch 1 {
                case -9223372036854775808:
                    var v13 int32 = 2147483647
                    _ = v13
                case 745:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        fmt.Println(v6, fi14)
                        v6 += (v6 * v6)
                        switch "a" {
                        case "":
                            fmt.Println(v6, v2)
                        case "foo bar":
                            fmt.Println(v2, fi11)
                        case "hello":
                            if (!(0.1 != ((((1e10 - (float32((0.0 * 1e10)) * 3.14159)) - 1.5) * 0.0) * 3.14159))) {
                                var v15 uint8 = 128
                                _ = v15
                            }
                        default:
                            fmt.Println(v6, v6)
                        }
                        fmt.Println(fi14, fi14, v6)
                        v16 := GenS1{F0: 32767}
                        _ = fi14
                        _ = v16
                    }
                case 42:
                    if (!true) {
                        var v17 uint8 = 128
                        if (!((v17 | 127) != v17)) {
                            for fi18 := 0; fi18 < 20; fi18++ {
                                var v19 int16 = 42
                                _ = fi18
                                _ = v19
                            }
                            var v20 uint16 = 752
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v17, v17, fi11)
                            case "a":
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    v2 -= (^(v2))
                                    v22 := []int{(fi11 - fi11)}
                                    v2 -= (v2 - (v2 % ((v2) | 1)))
                                    v23 := GenS2{F0: "", F1: 65535}
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            default:
                                var v24 uint = uint(v6)
                                _ = v24
                            }
                            _ = v20
                        } else {
                            if true {
                                v6 = (v6 ^ 2147483647)
                            }
                        }
                        v17 += (v17 ^ 1)
                        _ = v17
                    }
                }
            } else {
                fmt.Println(v2, v2)
                var v25 int8 = 1
                _ = v25
            }
            _ = fi11
        }
    }
    if (2147483647 > (v2 - v2)) {
        var v26 uint16 = 2
        var v27 float32 = (0.0 + -2.25)
        v26 = 2
        _ = v26
        _ = v27
    }
    if true {
        v6 = (v6 << ((0) & 63))
    }
    var v28 int16 = 32767
    for fi29 := 0; fi29 < 3; fi29++ {
        _ = len(v4)
        v30 := []int{}
        switch ("hello" + "") {
        case "\n\t":
            for fi31 := 0; fi31 < 3; fi31++ {
                _ = len(v30)
                _ = fi31
            }
        default:
            if true {
                var v32 uint64 = 1
                _ = v32
            } else {
                v33 := GenS3{F0: 2, F1: 0, F2: ("foo bar" + "")}
                _ = v33
            }
        }
        _ = fi29
        _ = v30
    }
    v34 := GenS4{F0: v28, F1: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
