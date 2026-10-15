package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 bool }
func genFn0() uint8 {
    var v5 uint64 = 18446744073709551615
    var v6 uint8 = 0
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 float32 = ((1.5 / float32(1e10)) / 1.0)
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    var v7 uint32 = 719
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: (!(-2.25 != 0.1)), F1: "\n\t", F2: 2}
        _ = fi8
        _ = v9
    }
    v7 -= (^(v7))
    if false {
        var v10 uint = 42
        if (v7 == v7) {
            v11 := GenS1{F0: (100 < -126), F1: "foo bar", F2: 517}
            var v12 uint16 = 65535
            _ = v11
            _ = v12
        }
        fmt.Println(v10)
        _ = v10
    }
    var v13 int64 = -9223372036854775808
    v7 = v7
    switch (v4 + "\n\t") {
    case "foo bar":
        if (0 <= -1) {
            delete(v2, -473)
            v7 += (v7 * v7)
        } else {
            var v14 uint = 4294967295
            v1 = 1.5
            switch v4 {
            case "a":
                fmt.Println(v14, v1, v4)
            case "\n\t":
                for fi15 := 0; fi15 < 12; fi15++ {
                    if true {
                        v1 += v1
                    } else {
                        v16 := []int{fi15, (fi15 >> (((v14 >> ((v14) & 63))) & 63)), 128, 32768}
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v18 := GenS1{F0: true, F1: v4, F2: 1}
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    _ = fi15
                }
            default:
                v19 := GenS2{F0: -1, F1: -2147483648, F2: false}
                _ = v19
            }
            _ = v14
        }
    case "":
        for fi20 := 0; fi20 < 6; fi20++ {
            if true {
                _ = len(v3)
            } else {
                var v21 float64 = 3.14159
                v3[fi20] = v4
                switch v4 {
                case "\n\t":
                    if false {
                        v21 += ((1e10 - v21) * v21)
                        fmt.Println(v4, v7)
                    } else {
                        if true {
                            v22 := v2[fi20]
                            var v23 float32 = v1
                            delete(v2, (fi20 & fi20))
                            _ = v22
                            _ = v23
                        }
                        fmt.Println(v21, v21)
                        var v24 bool = (552 != 18446744073709551615)
                        _ = v24
                    }
                case "":
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v1 -= v1
                        v26 := []int{(-(fi25)), 128, fi20, 639}
                        var v27 uint16 = 308
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                _ = v21
            }
            if false {
                fmt.Println(fi20, v1, fi20)
                v1 = 0.1
            } else {
                var v28 int16 = 0
                v29 := GenS1{F0: (v13 == (v13 * (-(v13)))), F1: "", F2: 1}
                _ = v28
                _ = v29
            }
            fmt.Println(v13, v13)
            _ = fi20
        }
    case "a":
        if ((float32(((((0.1 + 3.14159) - ((float64(((3.14159 - ((v1 / v1) * 0.1)) * float32(-2.25))) * 0.0) + ((3.14159 + ((3.14159 * float64((1.5 + 1e10))) + 1.5)) / -2.25))) / 1.5) + 1e10)) - ((float32(0.1) / (v1 - v1)) - float32(float64(v1)))) <= v1) {
            v7 -= (v7 ^ v7)
            switch v4 {
            case "a":
                v7 -= v7
            }
        }
    default:
        v13 = 42
    }
    switch v4 {
    case "a":
        var v30 string = v4
        _ = v30
    default:
        var v31 uint64 = 2
        _ = v31
    }
    var v32 bool = false
    v33 := GenS1{F0: (!((v7 - (v7 << ((67) & 63))) == (^(v7)))), F1: (("" + "\n\t") + (("hello" + (v4 + "")) + v4)), F2: 1}
    var v34 int8 = -1
    v7 += (v7 * 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v7
}
