package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := []string{"\n\t"}
    v3 := []int{9223372036854775807, 42}
    var v4 int = -2147483648
    v5 := make(map[int]bool)
    fmt.Println(v1, v4)
    if (2 != 430) {
        v6 := GenS1{F0: true}
        switch v4 {
        case 128:
            var v7 float32 = (-2.25 / ((1.5 + 1e10) * 3.14159))
            _ = v7
        }
        v4 -= (-(v4))
        _ = v6
    }
    var v8 int32 = 42
    switch v4 {
    case 0:
        switch "" {
        case "":
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(v8)
                if (!true) {
                    var v10 uint64 = 0
                    _ = v10
                } else {
                    v11 := GenS2{F0: 0, F1: 18446744073709551615, F2: "hello"}
                    fmt.Println(v8)
                    _ = len(v5)
                    _ = v11
                }
                _ = fi9
            }
        case "hello":
            var v12 int32 = v8
            _ = v12
        case "a":
            if false {
                var v13 int32 = v8
                v1 -= 255
                for fi14 := 0; fi14 < 19; fi14++ {
                    switch "hello" {
                    case "":
                        v5[(v4 % ((fi14) | 1))] = true
                    default:
                        _ = len(v5)
                    }
                    v15 := GenS3{F0: false, F1: uint32(v1)}
                    var v16 uint16 = (v1 / (((v1 | v1)) | 1))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            } else {
                if (!(false || false)) {
                    switch "\n\t" {
                    case "hello":
                        if (v8 > (v8 ^ 1)) {
                            var v17 float32 = (0.1 - float32(0.0))
                            v5[-2147483648] = false
                            v18 := GenS4{F0: (-(v8)), F1: 128, F2: 521}
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS1{F0: ((42 < 18446744073709551615) && true)}
                            v1 = 32767
                            _ = v19
                        }
                    case "foo bar":
                        for fi20 := 0; fi20 < 9; fi20++ {
                            v4 += fi20
                            fmt.Println(v8, fi20)
                            switch "foo bar" {
                            case "foo bar":
                                v21 := GenS5{F0: 42, F1: 9223372036854775807, F2: (v1 / (((v1 * (v1 * v1))) | 1))}
                                _ = v21
                            case "":
                                v5[(^(fi20))] = (true || (!((!(true || ("hello" < "hello"))) || true)))
                            case "a":
                                if true {
                                    for fi22 := 0; fi22 < 12; fi22++ {
                                        v4 = fi22
                                        if (-2.25 <= 1.5) {
                                            var v23 float64 = -2.25
                                            _ = v23
                                        }
                                        _ = fi22
                                    }
                                    var v24 int = (v4 ^ 32768)
                                    switch (fi20 | fi20) {
                                    case -1:
                                        v4 = (v4 / (((fi20 | 32767)) | 1))
                                    case 178:
                                        switch v24 {
                                        case 910:
                                            var v25 int64 = 0
                                            _ = v25
                                        }
                                    case -993:
                                        _ = len(v5)
                                    }
                                    _ = v24
                                }
                            }
                            _ = fi20
                        }
                    case "\n\t":
                        var v26 uint32 = 0
                        _ = v26
                    }
                } else {
                    var v27 int32 = (^(v8))
                    _ = v27
                }
                var v28 uint64 = 42
                _ = v28
            }
        }
    case -9223372036854775808:
        if true {
            v8 = (v8 * 997)
        } else {
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 float32 = 3.14159
                var v31 int16 = 0
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    var v32 int = v4
    v5[-701] = false
    v33 := GenS2{F0: 498, F1: 255, F2: "foo bar"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 = -2147483648
    var v34 int32 = (v8 - (v8 / ((28) | 1)))
    var v35 bool = (14 >= 255)
    var v36 int16 = 1
    var v37 bool = ("hello" >= "")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v8
}
