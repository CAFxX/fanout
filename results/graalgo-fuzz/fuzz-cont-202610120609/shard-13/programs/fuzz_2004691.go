package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 uint32 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    v6 := []int{32768, -299, -999, 32767}
    fmt.Println(v3, v1)
    v1 = false
    delete(v2, 2147483647)
    var v7 int64 = -831
    if v1 {
        fmt.Println(v3, v7)
        v8 := GenS1{F0: -562, F1: ("" + "a")}
        v1 = (937 >= 65535)
        _ = v8
    }
    if v1 {
        v5[(^(v3))] = "a"
    }
    v7 = (v7 % (((^(v7))) | 1))
    fmt.Println(v3)
    v9 := GenS2{F0: v7}
    v10 := GenS2{F0: v7}
    switch -1 {
    case 343:
        v11 := GenS2{F0: (v7 / ((v7) | 1))}
        _ = v11
    case -2147483648:
        var v12 int16 = 32767
        _ = v12
    default:
        for fi13 := 0; fi13 < 16; fi13++ {
            if (0 >= 27) {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS1{F0: -323, F1: ("a" + "a")}
                    _ = fi14
                    _ = v15
                }
                if (2 != 255) {
                    v1 = true
                }
                var v16 int = fi13
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    if (715 < 18446744073709551615) {
                        v6 = append(v6, (fi13 & 703))
                    }
                    switch (^(fi17)) {
                    case 721:
                        var v18 string = ((("hello" + "a") + "hello") + "hello")
                        _ = v18
                    default:
                        var v19 int8 = -128
                        _ = v19
                    }
                    _ = fi17
                }
                v7 = (v7 >> ((18446744073709551615) & 63))
                var v20 int16 = 32767
                _ = v20
            }
            _ = fi13
        }
    }
    switch "foo bar" {
    case "a":
        var v21 uint64 = 4294967295
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 10; fi22++ {
            if v1 {
                v7 = v7
                fmt.Println(fi22, v7, v3)
            }
            _ = fi22
        }
    case "":
        v2[2] = ""
    }
    v23 := GenS2{F0: v7}
    switch ("\n\t" + ((("hello" + "hello") + "hello") + "a")) {
    case "":
        v3 = v3
    case "foo bar":
        v7 += (v7 / (((v7 | v7)) | 1))
    case "hello":
        if (64 == (-(v7))) {
            if v1 {
                if len(v6) > 0 {
                    v6[0] = 0
                }
            } else {
                if (197 < -9223372036854775808) {
                    switch 9223372036854775807 {
                    case -1:
                        if (true && v1) {
                            var v24 uint32 = (v3 ^ (v3 - (v3 << ((584) & 63))))
                            var v25 float64 = float64((0.0 + (((-2.25 / ((3.14159 * 3.14159) * 1e10)) + 1e10) / (-2.25 * (0.1 / (0.0 - -2.25))))))
                            _ = v24
                            _ = v25
                        }
                    case 94:
                        v7 = (^(v7))
                    default:
                        fmt.Println(v7, v7)
                    }
                    if true {
                        v26 := GenS3{F0: 127, F1: 255}
                        if v1 {
                            var v27 float32 = (3.14159 - (0.1 - 1e10))
                            v27 = v27
                            _ = v27
                        } else {
                            v28 := v4[(^(v3))]
                            fmt.Println(v28, v3)
                            _ = v28
                        }
                        v1 = (!true)
                        _ = v26
                    }
                } else {
                    v29 := GenS2{F0: v7}
                    for fi30 := 0; fi30 < 11; fi30++ {
                        for fi31 := 0; fi31 < 9; fi31++ {
                            if (v3 >= 42) {
                                v32 := GenS2{F0: 0}
                                var v33 string = ("\n\t" + "hello")
                                var v34 uint64 = 2
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                v35 := GenS3{F0: 141, F1: 658}
                                _ = v35
                            }
                            var v36 int64 = v7
                            var v37 bool = (746 < -2147483648)
                            _ = fi31
                            _ = v36
                            _ = v37
                        }
                        v6 = append(v6, fi30)
                        v38 := v4[940]
                        _ = fi30
                        _ = v38
                    }
                    _ = v29
                }
                var v39 int16 = 127
                _ = v39
            }
            switch -1 {
            case 0:
                var v40 int16 = -32768
                _ = v40
            }
            var v41 int16 = -32768
            _ = v41
        } else {
            v1 = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
