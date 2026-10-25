package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int; F1 int64; F2 uint }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 uint = 885
    var v5 float32 = 3.14159
    var v6 int64 = 9223372036854775807
    var v7 int8 = 100
    if true {
        v5 = ((v5 + v5) - v5)
        _ = len(v3)
        if (false || true) {
            v8 := GenS1{F0: (-84 <= 1), F1: 790}
            switch 109 {
            case 42:
                v9 := GenS1{F0: false, F1: -32768}
                _ = v9
            case 9223372036854775807:
                switch 42 {
                case 812:
                    fmt.Println(v5, v5, v5)
                case -787:
                    v5 = ((v5 + 1.5) * 0.1)
                default:
                    v1 = append(v1, 42)
                }
            }
            for fi10 := 0; fi10 < 10; fi10++ {
                if (!(v5 >= float32(-2.25))) {
                    v6 = 878
                    fmt.Println(fi10)
                    v6 = v6
                }
                _ = fi10
            }
            _ = v8
        } else {
            switch 2147483647 {
            case 42:
                if true {
                    fmt.Println(v4, v7, v4)
                    var v11 int8 = 127
                    _ = v11
                } else {
                    fmt.Println(v5, v4)
                }
            case 9:
                var v12 string = ((("a" + "\n\t") + ("foo bar" + "a")) + "")
                _ = v12
            }
            delete(v3, 42)
            if (-1 == 601) {
                for fi13 := 0; fi13 < 8; fi13++ {
                    if len(v1) > 0 {
                        v1[0] = 188
                    }
                    v1 = append(v1, 494)
                    _ = fi13
                }
                _ = len(v1)
                if false {
                    if false {
                        var v14 uint32 = 0
                        _ = v14
                    }
                    switch 658 {
                    case 32768:
                        v15 := GenS2{F0: false, F1: (true && false), F2: (v7 << ((v4) & 63))}
                        _ = v15
                    case -231:
                        v16 := GenS1{F0: true, F1: 32767}
                        _ = v16
                    case -946:
                        switch "hello" {
                        case "\n\t":
                            v7 -= (v7 * 70)
                        case "":
                            fmt.Println(v5)
                        default:
                            v17 := GenS1{F0: true, F1: 42}
                            _ = v17
                        }
                    }
                    v1 = append(v1, 4294967295)
                }
            }
        }
    } else {
        var v18 uint32 = 2147483647
        v5 += 1e10
        v7 = v7
        _ = v18
    }
    v7 = -1
    v7 += (v7 - (v7 | v7))
    v1 = append(v1, 2147483647)
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    v19 := v2["hello"]
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := v2["foo bar"]
        var v22 int8 = v7
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        for fi23 := 0; fi23 < 19; fi23++ {
            v2[("" + ("hello" + ("hello" + "")))] = v19
            v7 += (^(v7))
            v19 = true
            _ = fi23
        }
    }
    if (0 <= -32768) {
        v5 = 1e10
        v5 -= v5
    } else {
        fmt.Println(v4, v6, v4)
        if v19 {
            switch "\n\t" {
            case "hello":
                v6 = v6
            case "foo bar":
                if (419 > 2147483648) {
                    var v24 uint32 = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 float32 = 0.0
                        _ = fi25
                        _ = v26
                    }
                }
            case "a":
                if v19 {
                    if (v19 && v19) {
                        v27 := GenS2{F0: false, F1: (1e10 < v5), F2: (v7 * (v7 % ((v7) | 1)))}
                        v5 = v5
                        v5 -= ((float32(((-2.25 / 1e10) / 0.1)) - (v5 - ((v5 / (1.5 + (float32(float64((1.5 * v5))) - 3.14159))) - v5))) / 3.14159)
                        _ = v27
                    } else {
                        var v28 int32 = -1
                        _ = v28
                    }
                    v29 := GenS2{F0: true, F1: (!((!v19) || (!(!v19)))), F2: 1}
                    if ((v7 + v7) == v7) {
                        v30 := GenS3{F0: v6, F1: 0}
                        _ = v30
                    }
                    _ = v29
                } else {
                    var v31 uint16 = 42
                    if (float64(v5) > 1e10) {
                        if false {
                            fmt.Println(v31, v4, v31)
                            var v32 int32 = -2147483648
                            _ = v32
                        }
                    } else {
                        v33 := GenS2{F0: true, F1: true, F2: 32}
                        _ = v33
                    }
                    _ = v31
                }
            }
            v34 := GenS2{F0: v19, F1: v19, F2: v7}
            _ = v34
        }
        v35 := GenS4{F0: 128, F1: (v6 & (v6 >> ((876) & 63))), F2: 255}
        _ = v35
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 int64 = v6
        _ = fi36
        _ = v37
    }
    v38 := GenS5{F0: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
