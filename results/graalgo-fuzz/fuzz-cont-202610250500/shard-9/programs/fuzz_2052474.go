package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 int16 = 285
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    var v6 uint8 = 86
    var v7 uint32 = 2147483647
    v8 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 2}
    v9 := GenS2{F0: 845, F1: -32768}
    if (42 == (v6 + v6)) {
        if true {
            if (!false) {
                if (-1 >= -1) {
                    fmt.Println(v4, v6)
                }
            } else {
                v10 := []int{-1, 517, 42}
                fmt.Println(v3)
                for fi11 := 0; fi11 < 10; fi11++ {
                    if (!(!(v3 <= (v3 ^ v3)))) {
                        _ = len(v5)
                        if ((v7 / ((v7) | 1)) <= v7) {
                            if false {
                                v10 = append(v10, fi11)
                                v4 = (3.14159 / float32(1.5))
                                switch "a" {
                                case "hello":
                                    for fi12 := 0; fi12 < 6; fi12++ {
                                        v4 = v4
                                        v4 = (v4 / v4)
                                        v5[("a" + "a")] = "foo bar"
                                        _ = fi12
                                    }
                                default:
                                    v7 -= (v7 & v7)
                                }
                            }
                            var v13 float32 = 1.5
                            fmt.Println(v6, fi11, v3)
                            _ = v13
                        } else {
                            var v14 uint8 = v6
                            var v15 int32 = 124
                            if len(v10) > 0 {
                                v10[0] = fi11
                            }
                            _ = v14
                            _ = v15
                        }
                        var v16 uint64 = 589
                        _ = v16
                    }
                    if len(v10) > 0 {
                        v10[0] = (fi11 ^ -1)
                    }
                    _ = fi11
                }
                _ = v10
            }
            v7 = v7
            var v17 int16 = v3
            _ = v17
        } else {
            v7 = v7
        }
    } else {
        fmt.Println(v6, v3, v4)
        if (1 <= -766) {
            v7 = 518
        }
    }
    if (!true) {
        v6 = (v6 ^ 42)
    } else {
        v18 := GenS3{F0: "a", F1: 42}
        if false {
            v6 -= v6
            switch ("" + "foo bar") {
            case "a":
                var v19 int64 = -1
                _ = v19
            default:
                v20 := []int{-960, 32767, -9223372036854775808}
                _ = v20
            }
            v21 := GenS2{F0: 42, F1: (v3 << ((853) & 63))}
            _ = v21
        } else {
            v22 := []int{2147483647, 32767, -9223372036854775808}
            v4 = (v4 / (1e10 + (-2.25 * (((v4 - 0.0) + (v4 / v4)) - ((0.0 - v4) / ((float32(-2.25) + (v4 + v4)) - 3.14159))))))
            v7 += (v7 >> ((0) & 63))
            _ = v22
        }
        v23 := v2[(v7 - (v7 % ((v7) | 1)))]
        _ = v18
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 16; fi24++ {
            switch (^(fi24)) {
            case 1:
                v6 = 127
            case 388:
                var v25 uint64 = 18446744073709551615
                _ = v25
            case 2147483647:
                v26 := GenS1{F0: 18446744073709551615, F1: 255, F2: v6}
                _ = v26
            }
            _ = fi24
        }
        v3 -= (v3 | (^(v3)))
    } else {
        if (244 >= 18446744073709551615) {
            v6 = 2
        }
        if false {
            var v27 int16 = v3
            _ = v27
        }
        var v28 int16 = v3
        _ = v28
    }
    v3 -= v3
    var v29 uint32 = 0
    if true {
        var v30 uint8 = (v6 << ((uint(v29)) & 63))
        fmt.Println(v4)
        _ = v30
    }
    v29 += (v7 % ((2147483647) | 1))
    v3 = (v3 & (v3 % ((659) | 1)))
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v6, v3, v6)
        _ = fi31
    }
    var v32 uint = 1
    var v33 uint8 = (^(v6))
    fmt.Println(v29)
    if true {
        fmt.Println(v7)
        switch "" {
        case "foo bar":
            v34 := GenS3{F0: "", F1: 255}
            _ = v34
        case "\n\t":
            fmt.Println(v4, v32)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
