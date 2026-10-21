package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
func genFn0() uint8 {
    v5 := GenS1{F0: -2147483648, F1: "\n\t", F2: (("a" + "hello") + ("a" + "hello"))}
    if true {
        v6 := GenS1{F0: -690, F1: "", F2: (("a" + ("foo bar" + "a")) + ("" + "a"))}
        var v7 uint8 = 52
        v8 := GenS2{F0: 2}
        _ = v6
        _ = v7
        _ = v8
    }
    _ = v5
    return 255
}
func genFn1(p0 string, p1 uint16) uint16 {
    fmt.Println(p1)
    if (86 != 100) {
        v9 := GenS3{F0: 116, F1: 42}
        v10 := []int{-102, 672}
        _ = v10
        _ = v9
    } else {
        v11 := GenS3{F0: 0, F1: 255}
        var v12 uint = 4294967295
        _ = v11
        _ = v12
    }
    v13 := GenS4{F0: true}
    var v14 int64 = 42
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return p1
}

func main() {
    var v1 uint16 = 717
    v2 := make(map[uint32]bool)
    v3 := []uint32{2, 4294967295}
    var v4 int64 = -101
    v15 := GenS4{F0: false}
    v4 -= (v4 % (((v4 * (v4 / (((v4 & (v4 / ((808) | 1)))) | 1)))) | 1))
    fmt.Println(v1, v1)
    var v16 int32 = 42
    if false {
        fmt.Println(v4)
        v16 = v16
        fmt.Println(v1, v1, v1)
    } else {
        v17 := GenS1{F0: (v16 - v16), F1: "", F2: "a"}
        var v18 int32 = (v16 >> ((1) & 63))
        if (4294967295 > 4294967295) {
            v16 -= v16
            if (-2.25 <= -2.25) {
                fmt.Println(v18, v16)
                switch (((("\n\t" + ("foo bar" + (("\n\t" + ("hello" + (("\n\t" + ("a" + ("a" + "hello"))) + "a"))) + (("\n\t" + ("" + ("hello" + "foo bar"))) + ((("\n\t" + ("hello" + "")) + "hello") + ("a" + "hello")))))) + "\n\t") + (("a" + "\n\t") + "foo bar")) + "foo bar") {
                case "\n\t":
                    var v19 string = "a"
                    _ = v19
                case "":
                    for fi20 := 0; fi20 < 6; fi20++ {
                        _ = len(v2)
                        _ = fi20
                    }
                case "foo bar":
                    v21 := GenS3{F0: 1, F1: genFn0()}
                    _ = v21
                default:
                    var v22 float64 = 3.14159
                    _ = v22
                }
                var v23 bool = ((!false) || (190 != 18446744073709551615))
                _ = v23
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    v4 = v4
                    _ = fi24
                }
                if true {
                    switch 0 {
                    case -9223372036854775808:
                        if (388 == 2147483647) {
                            v4 = v4
                            var v25 uint = 663
                            v16 = (v18 | (v16 | (v18 * (v16 ^ v16))))
                            _ = v25
                        } else {
                            var v26 bool = true
                            _ = v26
                        }
                    case 32767:
                        if (((-2.25 / 1e10) * (((0.1 / float64(3.14159)) / -2.25) * 1e10)) == (1.5 - (((3.14159 - 1e10) * 1e10) + 0.0))) {
                            v27 := GenS5{F0: 945}
                            switch "" {
                            case "hello":
                                var v28 uint64 = 504
                                _ = v28
                            default:
                                var v29 float32 = 1.5
                                _ = v29
                            }
                            _ = v27
                        }
                    case 390:
                        v4 -= v4
                    }
                    var v30 string = (("" + ("hello" + (("hello" + "") + "a"))) + "a")
                    _ = v30
                } else {
                    _ = len(v3)
                    v31 := GenS2{F0: 42}
                    _ = v31
                }
            }
            v4 = (v4 ^ v4)
        } else {
            switch "a" {
            case "foo bar":
                var v32 int16 = 0
                _ = v32
            case "a":
                for fi33 := 0; fi33 < 3; fi33++ {
                    delete(v2, uint32(v1))
                    if false {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            switch "a" {
                            case "\n\t":
                                v35 := GenS1{F0: v16, F1: "", F2: ("foo bar" + "a")}
                                _ = v35
                            case "foo bar":
                                v4 = v4
                            case "":
                                if true {
                                    fmt.Println(v16, v4)
                                    var v36 uint64 = 2
                                    fmt.Println(v4, v1, v36)
                                    _ = v36
                                }
                            default:
                                v4 = v4
                            }
                            var v37 uint = 2
                            var v38 float32 = (float32((1e10 - 0.0)) + 0.0)
                            _ = fi34
                            _ = v37
                            _ = v38
                        }
                    }
                    _ = fi33
                }
            case "":
                var v39 int16 = -1
                _ = v39
            }
            v16 = (v16 - v16)
        }
        _ = v17
        _ = v18
    }
    var v40 float32 = (float32((((1.5 / 0.1) * ((0.0 / 1.0) - 1e10)) + 0.0)) * 3.14159)
    var v41 int8 = -37
    v40 = v40
    if true {
        fmt.Println(v16, v40)
        var v42 int = 1
        _ = v42
    } else {
        if true {
            if len(v3) > 0 {
                v3[0] = 541
            }
            delete(v2, 42)
            var v43 float32 = v40
            _ = v43
        }
    }
    v41 += (v41 ^ (v41 % (((v41 >> ((18446744073709551615) & 63))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
