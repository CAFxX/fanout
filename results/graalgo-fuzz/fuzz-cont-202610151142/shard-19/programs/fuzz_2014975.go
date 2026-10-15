package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 int16, p1 int64, p2 bool) uint64 {
    var v7 int8 = -1
    if p2 {
        var v8 bool = (65535 > 18446744073709551615)
        p0 -= 32767
        v7 -= (v7 << ((65535) & 63))
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        var v9 string = "foo bar"
        _ = v9
    }
    switch 128 {
    case 32768:
        p1 -= p1
    case 9223372036854775807:
        fmt.Println(p2, v7)
    default:
        var v10 uint8 = 1
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 int32) uint64 {
    p0 += 1
    p0 = -2147483648
    v11 := []int{951}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = v11
    return 979
}

func main() {
    var v1 uint = 871
    v2 := make(map[string]bool)
    var v3 float64 = -2.25
    var v4 float32 = 1e10
    var v5 uint8 = 127
    var v6 int8 = 42
    var v12 int64 = 956
    if true {
        var v13 uint16 = 731
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint8 = v5
        switch (fi14 * (fi14 & (fi14 * (fi14 + fi14)))) {
        case 0:
            v5 = v5
        case 1:
            var v16 float64 = v3
            _ = v16
        case -1:
            var v17 uint32 = 514
            _ = v17
        default:
            v4 = ((-2.25 - (v4 / v4)) / (3.14159 - (-2.25 + float32(v3))))
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        if true {
            for fi19 := 0; fi19 < 8; fi19++ {
                if (true && ((v4 + 0.0) < v4)) {
                    v12 = (v12 % ((9223372036854775807) | 1))
                } else {
                    v20 := GenS1{F0: (v6 >> ((255) & 63)), F1: v12}
                    if false {
                        if true {
                            var v21 uint32 = 867
                            if true {
                                v5 = 127
                            } else {
                                v5 += v5
                                var v22 int = 295
                                v1 -= 18446744073709551615
                                switch (fi18 / ((32768) | 1)) {
                                case 42:
                                    switch 128 {
                                    case -9223372036854775808:
                                        fmt.Println(v5, v6, v21)
                                    case 702:
                                        var v23 int8 = v6
                                        _ = v23
                                    }
                                case 32768:
                                    v3 = v3
                                case 127:
                                    fmt.Println(v1)
                                }
                                _ = v22
                            }
                            var v24 bool = (-2147483648 >= fi19)
                            _ = v21
                            _ = v24
                        }
                        var v25 bool = (0 != -1)
                        _ = v25
                    } else {
                        fmt.Println(v3)
                        v12 -= v12
                        v26 := GenS2{F0: 680, F1: 18446744073709551615}
                        _ = v26
                    }
                    _ = v20
                }
                _ = fi19
            }
        }
        var v27 uint64 = 2
        _ = fi18
        _ = v27
    }
    v28 := []int{32767, -681, 496, 128}
    for fi29 := 0; fi29 < 13; fi29++ {
        switch 299 {
        case 2147483647:
            v30 := GenS1{F0: 57, F1: v12}
            _ = v30
        }
        var v31 int = (fi29 >> (((v1 >> ((v1) & 63))) & 63))
        _ = fi29
        _ = v31
    }
    v32 := GenS3{F0: 445, F1: (!false)}
    var v33 float64 = v3
    _ = len(v2)
    v34 := GenS3{F0: 495, F1: (int(v1) > 1)}
    for fi35 := 0; fi35 < 20; fi35++ {
        delete(v2, "a")
        _ = fi35
    }
    v12 -= (^(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
