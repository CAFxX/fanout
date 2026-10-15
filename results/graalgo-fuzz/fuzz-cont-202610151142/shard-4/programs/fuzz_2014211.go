package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint32; F1 string }
func genFn0() uint32 {
    v7 := []int{128, 2147483647}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint32 = 695
        for fi10 := 0; fi10 < 20; fi10++ {
            v9 = v9
            _ = fi10
        }
        var v11 uint = 42
        _ = fi8
        _ = v11
        _ = v9
    }
    switch -186 {
    case 1:
        var v12 uint8 = 1
        _ = v12
    case 2147483647:
        var v13 float32 = 1.5
        _ = v13
    case 42:
        switch "hello" {
        case "a":
            var v14 int8 = -1
            _ = v14
        }
    default:
        switch (("foo bar" + "") + "hello") {
        case "":
            v15 := GenS1{F0: -1, F1: 790, F2: 32768}
            _ = v15
        default:
            var v16 uint8 = 0
            _ = v16
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    return 65535
}

func main() {
    var v1 float64 = (((0.0 - 0.0) * (0.0 * float64(((1.5 + -2.25) * (-2.25 + -2.25))))) * float64(0.0))
    var v2 int32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 int32 = (v2 + 145)
    v5 := []int{1, 0, 32768, 32767}
    var v6 float32 = float32(v1)
    v6 -= 0.0
    v6 = 1.5
    fmt.Println(v6, v6)
    switch "a" {
    case "hello":
        var v17 int8 = -66
        _ = v17
    }
    fmt.Println(v6, v6, v2)
    fmt.Println(v1, v2, v4)
    v18 := GenS1{F0: 127, F1: 0, F2: 42}
    for fi19 := 0; fi19 < 18; fi19++ {
        if (float32((v1 - v1)) >= (v6 * v6)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v20 int32 = (v4 % ((-2147483648) | 1))
            _ = v20
        }
        _ = fi19
    }
    var v21 int = 42
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 float32 = v6
        fmt.Println(fi22)
        _ = fi22
        _ = v23
    }
    switch "hello" {
    case "hello":
        v24 := GenS2{F0: ((true && false) && (!true))}
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS3{F0: 65535}
            _ = len(v3)
            _ = fi25
            _ = v26
        }
    case "a":
        v27 := v3[564]
        _ = v27
    default:
        v6 += 0.0
    }
    if true {
        var v28 int64 = -1
        for fi29 := 0; fi29 < 18; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                switch ("hello" + "") {
                case "hello":
                    v3[0] = genFn0()
                case "foo bar":
                    v1 = v1
                case "a":
                    var v31 int64 = (v28 | (v28 - (v28 & (v28 + (v28 / (((v28 >> ((65535) & 63))) | 1))))))
                    _ = v31
                default:
                    var v32 uint32 = 42
                    _ = v32
                }
                _ = fi30
            }
            var v33 int8 = 42
            switch (^(v21)) {
            case 127:
                switch "hello" {
                case "foo bar":
                    if ((2147483647 >= 0) && (65535 >= 1)) {
                        var v34 float32 = 1.5
                        _ = v34
                    }
                case "hello":
                    if true {
                        if (!("\n\t" > ("\n\t" + "hello"))) {
                            for fi35 := 0; fi35 < 13; fi35++ {
                                v6 = (v6 / v6)
                                fmt.Println(v2)
                                v36 := GenS4{F0: 255, F1: 2}
                                _ = fi35
                                _ = v36
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            for fi37 := 0; fi37 < 7; fi37++ {
                                v2 += 0
                                _ = fi37
                            }
                        }
                        v3[genFn0()] = 2147483648
                    }
                default:
                    fmt.Println(fi29)
                }
            case -2147483648:
                fmt.Println(v28, v4, v1)
            default:
                var v38 uint16 = 1
                _ = v38
            }
            _ = fi29
            _ = v33
        }
        _ = v28
    } else {
        v21 = (v21 - -661)
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "":
        var v39 int32 = (-(v2))
        _ = v39
    case "a":
        v40 := GenS5{F0: 0, F1: (("a" + ("" + "\n\t")) + ("a" + "a"))}
        _ = v40
    }
    v41 := v3[0]
    v41 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
