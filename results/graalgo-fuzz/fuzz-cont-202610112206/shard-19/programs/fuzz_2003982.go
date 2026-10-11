package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint64, p2 uint) bool {
    v4 := GenS1{F0: 152, F1: 1, F2: (9 > -128)}
    v5 := GenS1{F0: 128, F1: 1, F2: (!true)}
    var v6 int16 = -841
    v7 := GenS1{F0: 127, F1: -1, F2: true}
    v8 := []int{42, 9223372036854775807, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int8 = -26
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    v1 = v1
    var v9 int8 = (v1 * (-(v1)))
    fmt.Println(v9)
    delete(v2, "\n\t")
    v1 = v9
    v10 := []int{-2147483648, -1, 4294967295, 232}
    fmt.Println(v1)
    switch ("\n\t" + "\n\t") {
    case "hello":
        fmt.Println(v1)
    case "":
        switch "\n\t" {
        case "":
            switch -1 {
            case 1:
                v10 = append(v10, -1)
            default:
                switch int(v1) {
                case 1:
                    if (!false) {
                        v1 = (v9 - v1)
                        v9 = (v1 + (v1 & (v9 - (-(v1)))))
                        v2[("hello" + "\n\t")] = 4294967295
                    } else {
                        v11 := GenS1{F0: 0, F1: -9223372036854775808, F2: (true && (((3.14159 + (((((1e10 - (0.0 * 3.14159)) - (-2.25 + float32(3.14159))) + 1.5) - ((-2.25 - 0.1) / -2.25)) / -2.25)) == (0.0 * -2.25)) && false))}
                        var v12 string = ("a" + ("foo bar" + "foo bar"))
                        v13 := GenS2{F0: 482, F1: 65535}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
            }
        case "hello":
            switch -1 {
            case 1:
                var v14 int8 = -128
                _ = v14
            case -992:
                if (0 <= 1) {
                    v9 -= 127
                    fmt.Println(v9, v1)
                }
            }
        case "foo bar":
            if true {
                var v15 string = "foo bar"
                _ = v15
            } else {
                if true {
                    var v16 float32 = (3.14159 / float32((0.1 - (0.1 / (3.14159 - 0.0)))))
                    if len(v10) > 0 {
                        v10[0] = 716
                    }
                    var v17 bool = true
                    _ = v16
                    _ = v17
                }
                if (0.1 >= (1.5 * 1.5)) {
                    v3[-204] = "a"
                    switch "\n\t" {
                    case "foo bar":
                        switch (("a" + ("a" + "")) + "") {
                        case "\n\t":
                            var v18 uint = 42
                            _ = v18
                        case "foo bar":
                            v19 := GenS3{F0: ("" + "a"), F1: 424, F2: true}
                            _ = v19
                        case "a":
                            var v20 int8 = v1
                            _ = v20
                        default:
                            var v21 int64 = 42
                            _ = v21
                        }
                    case "\n\t":
                        fmt.Println(v9)
                    default:
                        switch "hello" {
                        case "hello":
                            fmt.Println(v1, v1)
                        case "":
                            if (!(0 == 658)) {
                                var v22 int32 = -255
                                v9 = 42
                                _ = v22
                            } else {
                                _ = len(v10)
                                v23 := GenS4{F0: v9}
                                _ = v23
                            }
                        case "foo bar":
                            fmt.Println(v9, v1)
                        default:
                            if (522 != -1) {
                                _ = len(v3)
                                if (4294967295 < 1) {
                                    v24 := GenS2{F0: 124, F1: 360}
                                    v25 := GenS1{F0: 2, F1: 9223372036854775807, F2: (-1 != 1)}
                                    var v26 int32 = -807
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                var v27 float32 = 0.1
                                _ = v27
                            } else {
                                var v28 uint64 = 622
                                _ = v28
                            }
                        }
                    }
                }
            }
        }
    default:
        var v29 float32 = 1.5
        _ = v29
    }
    switch "\n\t" {
    case "":
        fmt.Println(v1, v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
