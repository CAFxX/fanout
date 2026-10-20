package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint32 = 2
    var v2 float64 = (0.0 - (0.1 / 1.0))
    v3 := []int{}
    v4 := []string{"\n\t", (("a" + ("\n\t" + "\n\t")) + ("foo bar" + "a")), "hello", "foo bar"}
    v5 := make(map[uint32]int)
    var v6 int8 = 42
    v1 += v1
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    var v9 string = "foo bar"
    fmt.Println(v6)
    var v10 float32 = -2.25
    v11 := GenS1{F0: 0, F1: 18446744073709551615}
    var v12 bool = true
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint = 65535
        v15 := GenS2{F0: (-(v6)), F1: ((v1 ^ (v1 / ((65535) | 1))) == 42)}
        _ = fi13
        _ = v14
        _ = v15
    }
    if (255 >= 32767) {
        var v16 uint32 = 216
        if ((2 <= 0) || (!(v9 == v9))) {
            switch (((v9 + ((v9 + (v9 + (v9 + (v9 + "hello")))) + ("\n\t" + v9))) + ((v9 + (v9 + (v9 + "foo bar"))) + v9)) + ("\n\t" + "hello")) {
            case "a":
                v16 -= (^(v1))
            case "foo bar":
                fmt.Println(v16)
            case "hello":
                switch 9223372036854775807 {
                case 32768:
                    v17 := GenS1{F0: 255, F1: 65535}
                    _ = v17
                case 128:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 float32 = v10
                        if v12 {
                            fmt.Println(v16, v10)
                            v20 := GenS2{F0: (v6 << ((924) & 63)), F1: v12}
                            _ = v20
                        } else {
                            var v21 int16 = 134
                            _ = v21
                        }
                        v19 = (3.14159 + v10)
                        _ = fi18
                        _ = v19
                    }
                case 42:
                    fmt.Println(v9, v1, v1)
                default:
                    if v12 {
                        v16 = v1
                    } else {
                        var v22 int16 = -32768
                        _ = v22
                    }
                }
            }
            v23 := GenS2{F0: v6, F1: v12}
            _ = v23
        }
        _ = v16
    }
    v9 = v9
    if (0 != -1) {
        var v24 int64 = 543
        v12 = true
        _ = v24
    }
    switch "a" {
    case "foo bar":
        var v25 uint = 18446744073709551615
        _ = v25
    case "a":
        v26 := GenS2{F0: v6, F1: (!((-2.25 + (v10 + ((((float32(((v2 * v2) / 1.0)) / 0.1) + v10) * v10) * v10))) > v10))}
        _ = v26
    case "\n\t":
        var v27 int8 = (v6 ^ 127)
        _ = v27
    }
    v28 := GenS3{F0: 1, F1: -249}
    var v29 int8 = v6
    v6 += (v6 / (((v29 << ((1) & 63))) | 1))
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS4{F0: 4294967295, F1: 9223372036854775807}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        v6 = (v6 << ((0) & 63))
        switch 2147483647 {
        case 374:
            v33 := GenS3{F0: 42, F1: 9223372036854775807}
            _ = v33
        case 1:
            var v34 uint32 = (v1 - v1)
            _ = v34
        default:
            if true {
                var v35 uint16 = 1
                var v36 float32 = v10
                v1 = (v1 & (v1 << ((uint(v35)) & 63)))
                _ = v35
                _ = v36
            } else {
                v1 = (v1 ^ 2147483648)
                if false {
                    if (!(-1 < 9223372036854775807)) {
                        v37 := GenS1{F0: 580, F1: 42}
                        var v38 int = (^(fi32))
                        _ = v37
                        _ = v38
                    }
                    switch fi32 {
                    case -913:
                        if (559 != (v1 ^ v1)) {
                            var v39 uint64 = 0
                            var v40 int8 = 61
                            _ = v39
                            _ = v40
                        }
                    case 4294967295:
                        v29 -= (v29 ^ (v29 + v29))
                    default:
                        fmt.Println(fi32, fi32)
                    }
                } else {
                    for fi41 := 0; fi41 < 10; fi41++ {
                        var v42 uint64 = 42
                        switch -560 {
                        case -390:
                            fmt.Println(fi32)
                            v43 := GenS5{F0: (v9 + "a")}
                            _ = v43
                        case 32768:
                            fmt.Println(v6)
                        case 32767:
                            v44 := GenS1{F0: 65535, F1: 2}
                            _ = v44
                        }
                        v45 := GenS6{F0: 58, F1: 1}
                        _ = fi41
                        _ = v42
                        _ = v45
                    }
                    v46 := GenS1{F0: 32767, F1: 0}
                    v47 := v5[v1]
                    _ = v46
                    _ = v47
                }
                for fi48 := 0; fi48 < 18; fi48++ {
                    v49 := GenS7{F0: 2147483647, F1: 1}
                    _ = len(v5)
                    v1 = 209
                    v10 += 1e10
                    v12 = (!(9223372036854775807 <= 1))
                    _ = fi48
                    _ = v49
                }
            }
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
