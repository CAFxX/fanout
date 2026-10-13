package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16 }
func genFn0() int8 {
    var v5 int32 = -2147483648
    var v6 string = (("hello" + "\n\t") + "a")
    var v7 int16 = 1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int = (fi8 << ((2) & 63))
        v10 := []int{-53, v9}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 int32 = 75
    var v2 uint8 = 199
    v3 := []uint32{42}
    var v4 int8 = -128
    if true {
        v2 -= (v2 >> ((4294967295) & 63))
        if false {
            v4 += genFn0()
        }
        v4 -= -1
    } else {
        if (24 > genFn0()) {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch "foo bar" {
                case "a":
                    var v12 float64 = 0.0
                    _ = v12
                }
                _ = fi11
            }
        }
    }
    fmt.Println(v2, v1)
    if (127 >= (^(v4))) {
        if false {
            var v13 uint = 2
            if len(v3) > 0 {
                v3[0] = 146
            }
            _ = v13
        } else {
            v14 := GenS1{F0: 1, F1: true}
            _ = v14
        }
        v1 += (v1 & (v1 + (-(v1))))
    }
    v15 := GenS1{F0: (v2 - (v2 >> ((2) & 63))), F1: (65535 >= 1)}
    if false {
        v4 = v4
        fmt.Println(v4)
    } else {
        if (96 != 2) {
            fmt.Println(v1, v1, v2)
            if false {
                var v16 uint16 = 333
                fmt.Println(v16, v1, v2)
                v3 = append(v3, 297)
                _ = v16
            }
        }
        _ = len(v3)
    }
    switch "" {
    case "\n\t":
        switch 1 {
        case 163:
            v4 = v4
        default:
            v2 -= 2
        }
    case "hello":
        fmt.Println(v2, v2, v2)
    case "":
        v3 = append(v3, 307)
    }
    v17 := GenS1{F0: v2, F1: true}
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            if (175 < (^(v2))) {
                var v20 int8 = (v4 % ((v4) | 1))
                if (!true) {
                    if false {
                        v21 := GenS2{F0: genFn0(), F1: 413, F2: 9223372036854775807}
                        _ = v21
                    }
                    fmt.Println(v20)
                } else {
                    v22 := GenS3{F0: 133}
                    v20 = v4
                    _ = v22
                }
                v3 = append(v3, 42)
                switch ((("hello" + ("foo bar" + "a")) + ("foo bar" + (("\n\t" + "hello") + "\n\t"))) + "") {
                case "foo bar":
                    var v23 uint16 = 65535
                    _ = v23
                default:
                    var v24 float64 = 0.1
                    _ = v24
                }
                v25 := GenS1{F0: v2, F1: ((65535 > 4294967295) || (0.0 < ((float32(((0.0 - (1e10 * float64(0.1))) - 1e10)) + 1.5) / 3.14159)))}
                _ = v20
                _ = v25
            }
            v3 = append(v3, 42)
            _ = fi19
        }
        var v26 float32 = 3.14159
        v27 := GenS4{F0: 0}
        _ = fi18
        _ = v26
        _ = v27
    }
    v28 := GenS5{F0: 0, F1: v4}
    v4 = 127
    v29 := GenS6{F0: 32767}
    if (((false && false) || (!(2147483648 != 4294967295))) && false) {
        var v30 int64 = 858
        v1 = v1
        _ = v30
    }
    v4 += v4
    v1 = (v1 + v1)
    switch 837 {
    case 127:
        v31 := GenS3{F0: 479}
        _ = v31
    case 32768:
        for fi32 := 0; fi32 < 14; fi32++ {
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 string = "\n\t"
                if len(v3) > 0 {
                    v3[0] = 1
                }
                _ = fi33
                _ = v34
            }
            var v35 float64 = float64(1e10)
            _ = fi32
            _ = v35
        }
    default:
        if true {
            v4 -= v4
            var v36 int64 = 9223372036854775807
            _ = v36
        } else {
            var v37 bool = false
            _ = v37
        }
    }
    v4 = (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
