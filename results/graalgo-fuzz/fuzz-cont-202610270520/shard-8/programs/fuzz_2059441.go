package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 bool = true
    var v2 float64 = 1.5
    v3 := []uint32{42}
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 255
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    switch 128 {
    case 2147483647:
        if v1 {
            switch -9223372036854775808 {
            case -1:
                v5 += (v5 ^ 2)
            case 1:
                fmt.Println(v2, v2, v1)
            }
            var v6 bool = v1
            _ = v6
        } else {
            if (809 == -1) {
                switch "foo bar" {
                case "":
                    v2 += (v2 - v2)
                }
                v7 := GenS1{F0: 581, F1: -752, F2: (255 != 255)}
                _ = v7
            }
            for fi8 := 0; fi8 < 19; fi8++ {
                if v1 {
                    v2 = v2
                    var v9 int8 = 101
                    fmt.Println(v5)
                    _ = v9
                }
                _ = fi8
            }
        }
    case 0:
        var v10 int = 0
        _ = v10
    case -9223372036854775808:
        if (613 >= -110) {
            v2 += 3.14159
            v2 = 0.1
            v1 = ((-2.25 * 3.14159) < 0.0)
        }
    default:
        if ((true || (0 == -43)) && v1) {
            v11 := v4[2147483647]
            var v12 string = ""
            v13 := GenS2{F0: v5}
            _ = v11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v5)
    fmt.Println(v1, v2, v2)
    v14 := GenS1{F0: 2147483648, F1: -2147483648, F2: true}
    fmt.Println(v1)
    fmt.Println(v2, v5, v1)
    if v1 {
        v1 = v1
        if (831 <= 0) {
            v5 = v5
            var v15 uint32 = 351
            _ = v15
        } else {
            var v16 bool = (-32768 == 1)
            if v16 {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS1{F0: 4294967295, F1: (-(fi17)), F2: (!true)}
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v1, v5, v16)
                v16 = v16
            } else {
                v19 := GenS2{F0: v5}
                var v20 uint = 65535
                v20 = 255
                _ = v19
                _ = v20
            }
            _ = v16
        }
    } else {
        var v21 int8 = 1
        for fi22 := 0; fi22 < 11; fi22++ {
            _ = len(v3)
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 float64 = float64((0.0 * ((float32(v2) + (3.14159 * 0.0)) / 1.0)))
                if (127 == v21) {
                    switch "hello" {
                    case "foo bar":
                        if ((("a" + "") + ("a" + "a")) > "foo bar") {
                            fmt.Println(v1)
                            var v25 int = fi23
                            v5 -= (v5 + v5)
                            _ = v25
                        } else {
                            v3 = append(v3, 859)
                            if len(v3) > 0 {
                                v3[0] = 2147483648
                            }
                            var v26 float64 = 1e10
                            _ = v26
                        }
                    }
                }
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        fmt.Println(v21, v2)
        _ = v21
    }
    v27 := v4[237]
    switch ("" + "hello") {
    case "foo bar":
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 float32 = (3.14159 + 0.1)
            v30 := GenS2{F0: v5}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    var v31 float32 = 1.5
    var v32 int16 = 0
    var v33 uint64 = 18446744073709551615
    fmt.Println(v27, v32, v1)
    for fi34 := 0; fi34 < 20; fi34++ {
        v2 -= 1e10
        v35 := GenS2{F0: v5}
        fmt.Println(v27, v32)
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
