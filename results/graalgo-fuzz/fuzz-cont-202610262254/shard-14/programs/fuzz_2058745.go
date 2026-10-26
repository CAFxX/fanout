package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
func genFn0(p0 int64, p1 uint16) int16 {
    v6 := GenS1{F0: (("hello" + "") + "\n\t"), F1: false, F2: ("foo bar" + "a")}
    v7 := GenS2{F0: 0, F1: p1}
    var v8 float64 = -2.25
    var v9 int64 = (p0 / (((p0 / ((-365) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 72
    v2 := []uint32{65535, 2147483648, 65535, 42}
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    var v5 int8 = 109
    delete(v4, -903)
    var v10 bool = false
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            v12 := GenS3{F0: v5, F1: 2, F2: 626}
            fmt.Println(v3, v10)
            _ = v12
        } else {
            var v13 float64 = 1e10
            v2 = append(v2, 42)
            var v14 uint64 = 1
            _ = v13
            _ = v14
        }
        var v15 uint = 42
        _ = fi11
        _ = v15
    }
    if v10 {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint8 = v1
            v18 := GenS3{F0: (v5 & -127), F1: 18446744073709551615, F2: 948}
            fmt.Println(fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        switch "" {
        case "\n\t":
            switch 32767 {
            case 543:
                if (721 <= uint64(v1)) {
                    if true {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            switch ("hello" + "hello") {
                            case "":
                                var v19 int32 = -2147483648
                                _ = v19
                            }
                        }
                    } else {
                        var v20 uint = 255
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v22 := v4[fi21]
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                }
            default:
                var v23 uint64 = 0
                _ = v23
            }
        case "hello":
            var v24 string = ""
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v2 = append(v2, 65535)
            _ = fi25
        }
    }
    switch 128 {
    case 42:
        for fi26 := 0; fi26 < 6; fi26++ {
            fmt.Println(fi26, v1)
            _ = fi26
        }
    }
    v3 += 0.0
    v27 := GenS2{F0: 9223372036854775807, F1: 0}
    v28 := GenS3{F0: v5, F1: 42, F2: 0}
    v10 = (-9223372036854775808 != -1)
    v1 = v1
    v1 = 42
    var v29 uint8 = 42
    v10 = (400 != 463)
    v5 = (v5 * (v5 - v5))
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
