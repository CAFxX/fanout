package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }

func main() {
    var v1 bool = (42 == 2)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    var v5 uint64 = 42
    var v6 uint32 = 2
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int32 = -2147483648
        if (128 >= 78) {
            if (-2 > 123) {
                delete(v4, fi8)
                v7 -= (v7 >> ((uint(v5)) & 63))
                var v10 float32 = 0.0
                _ = v10
            } else {
                fmt.Println(fi8)
                if true {
                    if v1 {
                        if false {
                            for fi11 := 0; fi11 < 10; fi11++ {
                                switch "\n\t" {
                                case "a":
                                    if true {
                                        var v12 uint16 = (v7 << ((2) & 63))
                                        v13 := GenS1{F0: (v7 & (^(v7))), F1: v6, F2: (v7 << ((215) & 63))}
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v2[(^(v6))] = (4294967295 <= 929)
                                        switch -1 {
                                        case 9223372036854775807:
                                            var v14 float32 = (0.1 + 1e10)
                                            _ = v14
                                        case -229:
                                            v15 := GenS2{F0: v6}
                                            _ = v15
                                        default:
                                            fmt.Println(fi11, v5)
                                        }
                                        if (("hello" + "hello") >= ("hello" + "hello")) {
                                            v1 = false
                                        } else {
                                            var v16 int8 = 42
                                            v5 += v5
                                            v17 := GenS1{F0: (v7 ^ (v7 >> ((255) & 63))), F1: 2, F2: v7}
                                            if v1 {
                                                switch ("" + "\n\t") {
                                                case "\n\t":
                                                    v18 := GenS2{F0: (v6 >> ((976) & 63))}
                                                    _ = v18
                                                case "foo bar":
                                                    v19 := GenS2{F0: v6}
                                                    _ = v19
                                                }
                                            } else {
                                                var v20 int64 = 707
                                                _ = v20
                                            }
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                case "foo bar":
                                    _ = len(v4)
                                }
                                _ = fi11
                            }
                            v6 -= v6
                        }
                        v21 := GenS1{F0: v7, F1: (v6 - 1), F2: v7}
                        v22 := []int{int(v6), int(v5), (fi8 << ((65535) & 63)), -2147483648}
                        _ = v21
                        _ = v22
                    }
                    if (1 < fi8) {
                        v6 += v6
                        v23 := GenS3{F0: "a", F1: ""}
                        _ = v23
                    } else {
                        v24 := []int{-9223372036854775808}
                        switch fi8 {
                        case 370:
                            v1 = (!(-9223372036854775808 >= -275))
                        case -274:
                            fmt.Println(v9, v6)
                        case 1:
                            if ((true && (!(!(!(true && ((v1 && v1) || false)))))) && v1) {
                                v9 -= v9
                                switch 127 {
                                case 0:
                                    var v25 string = ("foo bar" + "hello")
                                    _ = v25
                                case -9223372036854775808:
                                    if (4294967295 >= 255) {
                                        var v26 bool = false
                                        var v27 uint8 = 2
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            } else {
                                v24 = append(v24, (fi8 / (((fi8 + (fi8 + (fi8 | fi8)))) | 1)))
                            }
                        default:
                            v9 += (v9 + v9)
                        }
                        _ = v24
                    }
                    _ = len(v2)
                } else {
                    var v28 uint32 = 2147483647
                    _ = v28
                }
                var v29 uint32 = v6
                _ = v29
            }
        }
        switch fi8 {
        case -1:
            v30 := []int{fi8, (fi8 % (((fi8 * int(v9))) | 1)), fi8}
            _ = v30
        }
        _ = fi8
        _ = v9
    }
    v31 := GenS2{F0: v6}
    var v32 uint16 = (v7 + v7)
    if ((451 < (^(v5))) || (((1.5 / (1.5 * (((1.5 - float64(float32(3.14159))) - ((1e10 * float64(0.0)) - 0.0)) + -2.25))) < -2.25) || v1)) {
        var v33 uint64 = (v5 << ((uint(v7)) & 63))
        if (v1 || (3.14159 != 0.1)) {
            var v34 uint64 = (v33 / (((v33 % ((18446744073709551615) | 1))) | 1))
            v7 = v32
            v34 = (v5 << ((42) & 63))
            _ = v34
        } else {
            if v1 {
                v6 = v6
            }
        }
        delete(v2, (v6 * (^(v6))))
        _ = v33
    }
    v6 += v6
    fmt.Println(v1, v1, v6)
    v35 := GenS4{F0: 127, F1: uint64(v7), F2: 0}
    v5 = (v5 << ((327) & 63))
    v36 := []int{}
    var v37 bool = ("hello" < "a")
    fmt.Println(v32, v6)
    v37 = v1
    if v37 {
        var v38 uint16 = v32
        fmt.Println(v7, v32)
        _ = v38
    } else {
        v37 = v1
        _ = len(v36)
    }
    var v39 uint8 = 55
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
