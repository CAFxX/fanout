package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 uint32 = 666
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 int8 = 0
    var v6 uint64 = 892
    v7 := []int{9223372036854775807, 1, 32768}
    switch 314 {
    case 4294967295:
        for fi8 := 0; fi8 < 19; fi8++ {
            switch (^(fi8)) {
            case 863:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 32768:
                var v9 uint8 = 217
                _ = v9
            case 2147483647:
                if (1 != -475) {
                    v1 += 1
                    var v10 float32 = 0.1
                    switch ("hello" + ((("hello" + "\n\t") + ((("hello" + "a") + ("" + "")) + ("a" + "hello"))) + "foo bar")) {
                    case "a":
                        switch (("\n\t" + "hello") + "") {
                        case "hello":
                            v5 -= v5
                        }
                    case "foo bar":
                        for fi11 := 0; fi11 < 13; fi11++ {
                            switch "a" {
                            case "\n\t":
                                v10 = v10
                            }
                            _ = fi11
                        }
                    case "":
                        switch "a" {
                        case "a":
                            for fi12 := 0; fi12 < 3; fi12++ {
                                var v13 uint8 = 240
                                _ = len(v7)
                                var v14 int64 = -1
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                var v15 uint16 = 32767
                                _ = v15
                            case "foo bar":
                                v16 := GenS1{F0: -1}
                                _ = v16
                            default:
                                v17 := GenS2{F0: (v1 / ((v1) | 1)), F1: 4294967295, F2: 1}
                                _ = v17
                            }
                        case "\n\t":
                            _ = len(v3)
                        default:
                            v18 := GenS2{F0: (^(v1)), F1: (fi8 >> ((2) & 63)), F2: int64(v5)}
                            _ = v18
                        }
                    }
                    _ = v10
                }
            default:
                var v19 uint32 = v1
                _ = v19
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 bool = (((((0.1 / float64(1.5)) + 0.0) * -2.25) + (((-2.25 + 3.14159) - (1e10 * 0.1)) - 1e10)) == -2.25)
                v6 = 228
                _ = fi20
                _ = v21
            }
            _ = fi8
        }
    case -1:
        v7 = append(v7, 427)
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(fi22)
        v23 := GenS2{F0: v1, F1: (^(fi22)), F2: -850}
        _ = fi22
        _ = v23
    }
    v1 += (v1 >> ((uint(v6)) & 63))
    v24 := GenS1{F0: -32768}
    v25 := GenS3{F0: "a", F1: ("a" + "")}
    fmt.Println(v5)
    if ((1 <= -1) && (false && (-2147483648 >= -100))) {
        var v26 bool = (-540 == 785)
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v5 -= (-(v5))
        _ = fi27
    }
    var v28 uint32 = (v1 | (^(v1)))
    fmt.Println(v1, v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
