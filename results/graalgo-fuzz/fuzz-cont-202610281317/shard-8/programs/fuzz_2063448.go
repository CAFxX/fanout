package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 string = ("hello" + "")
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(v3)
        _ = fi5
    }
    switch (v4 + (v4 + (v4 + v4))) {
    case "\n\t":
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v3 -= 42
                var v8 uint8 = 42
                if ((2147483647 >= 42) || (v4 == ("hello" + ""))) {
                    v4 = "a"
                    var v9 uint = 483
                    switch 4294967295 {
                    case 127:
                        switch fi7 {
                        case -137:
                            _ = len(v2)
                        default:
                            switch 2147483647 {
                            case 9223372036854775807:
                                v10 := v2[fi7]
                                _ = v10
                            case -192:
                                fmt.Println(v4)
                            case 16:
                                v8 = (v8 / ((42) | 1))
                            }
                        }
                    }
                    _ = v9
                }
                v8 = v8
                if true {
                    _ = len(v2)
                    _ = len(v1)
                    v3 += (v3 - (v3 & (v3 & (v3 + v3))))
                } else {
                    switch 22 {
                    case 259:
                        v1[776] = fi6
                    case 530:
                        v11 := []int{(fi6 * fi7)}
                        _ = v11
                    case 676:
                        v12 := v1[1]
                        _ = v12
                    default:
                        v13 := []int{-428, fi6}
                        _ = v13
                    }
                    v4 = "hello"
                    v14 := GenS1{F0: 0, F1: int32(v8)}
                    v3 = v3
                    if (0 < 1) {
                        var v15 uint64 = 1
                        var v16 uint64 = v3
                        v17 := GenS2{F0: 42, F1: 791}
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v8, fi7, fi6)
                        v18 := []int{128}
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = v18
                    }
                    _ = v14
                }
                _ = fi7
                _ = v8
            }
            var v19 float64 = ((0.0 - (0.0 + float64((((1e10 + ((1.5 + (1e10 + ((0.1 * 0.1) + (0.0 / (0.0 - (0.1 - 3.14159)))))) - float32(float64(1.5)))) - 3.14159) - float32(1e10))))) - float64((-2.25 - 1e10)))
            var v20 int8 = 1
            _ = fi6
            _ = v19
            _ = v20
        }
    case "hello":
        fmt.Println(v3, v3)
    }
    fmt.Println(v4, v4)
    v3 = 970
    var v21 float32 = 3.14159
    if (18446744073709551615 == 65535) {
        v21 = 0.1
        v21 = (v21 - float32(0.1))
        v4 = v4
    }
    if true {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (!false) {
                v21 = v21
            }
            v3 = v3
            fmt.Println(v21)
            _ = fi22
        }
        var v23 int8 = 1
        _ = v23
    }
    v24 := []int{414, 548}
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 uint32 = 926
        _ = fi25
        _ = v26
    }
    switch v4 {
    case "a":
        switch 0 {
        case 9223372036854775807:
            for fi27 := 0; fi27 < 20; fi27++ {
                delete(v2, (fi27 >> ((0) & 63)))
                fmt.Println(v4, fi27)
                _ = fi27
            }
        case 743:
            v28 := GenS2{F0: 46, F1: 180}
            _ = v28
        case 2147483647:
            switch v4 {
            case "foo bar":
                var v29 uint = 1
                _ = v29
            default:
                var v30 int16 = 1
                _ = v30
            }
        default:
            v3 -= (v3 & 18446744073709551615)
        }
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint32 = 4294967295
        v4 = (v4 + v4)
        _ = fi31
        _ = v32
    }
    var v33 uint = 255
    v34 := v2[1]
    var v35 uint = (v33 * 2)
    switch 9223372036854775807 {
    case 42:
        v21 -= 0.1
    case 32768:
        var v36 int32 = 288
        _ = v36
    default:
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS1{F0: uint8(v3), F1: 2147483647}
            _ = fi37
            _ = v38
        }
    }
    fmt.Println(v33, v34)
    fmt.Println(v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
