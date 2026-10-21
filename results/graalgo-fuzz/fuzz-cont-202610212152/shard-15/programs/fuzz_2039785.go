package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0() uint8 {
    v6 := GenS1{F0: -606}
    v7 := GenS2{F0: 0, F1: 1, F2: 1}
    var v8 uint = 18446744073709551615
    _ = v6
    _ = v7
    _ = v8
    return 155
}

func main() {
    v1 := []string{}
    var v2 uint16 = 1
    var v3 int32 = 2147483647
    v4 := []uint32{622, 2}
    var v5 uint8 = 128
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if true {
        v2 = v2
        switch "\n\t" {
        case "a":
            var v9 uint64 = 18446744073709551615
            _ = v9
        case "\n\t":
            v5 = (v5 % ((v5) | 1))
        }
        var v10 float32 = 0.1
        _ = v10
    }
    var v11 float64 = -2.25
    v5 -= 255
    var v12 int = -401
    var v13 uint64 = 2
    if false {
        var v14 int = v12
        var v15 uint64 = 953
        _ = v14
        _ = v15
    } else {
        v2 += 255
    }
    var v16 uint = 42
    v4 = append(v4, 896)
    v2 -= (v2 | (v2 * (v2 | v2)))
    v17 := GenS2{F0: v3, F1: (^(v16)), F2: genFn0()}
    for fi18 := 0; fi18 < 4; fi18++ {
        v5 -= 1
        switch "" {
        case "":
            v19 := GenS2{F0: 734, F1: v16, F2: v5}
            _ = v19
        default:
            v5 += (v5 / (((v5 - (v5 - v5))) | 1))
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi18
    }
    if true {
        fmt.Println(v2, v3, v16)
        for fi20 := 0; fi20 < 7; fi20++ {
            if (1 != (-(v3))) {
                v16 = 210
                var v21 float32 = 1e10
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if (!(true || false)) {
                        switch ("\n\t" + "a") {
                        case "a":
                            switch ("" + "\n\t") {
                            case "foo bar":
                                var v23 float32 = 1e10
                                _ = v23
                            case "hello":
                                var v24 uint32 = 360
                                _ = v24
                            }
                        case "\n\t":
                            var v25 string = "hello"
                            _ = v25
                        default:
                            v11 += ((float64((0.1 / 1e10)) / 3.14159) + 1e10)
                        }
                        var v26 string = "\n\t"
                        _ = v26
                    } else {
                        switch (fi20 | (v12 ^ 853)) {
                        case -344:
                            v16 -= (v16 * v16)
                        default:
                            var v27 uint = v16
                            var v28 int64 = 42
                            v28 = int64(v13)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v27
                            _ = v28
                        }
                        v16 -= v16
                    }
                    var v29 int = fi20
                    _ = fi22
                    _ = v29
                }
            }
            var v30 uint16 = (^(v2))
            _ = fi20
            _ = v30
        }
        v31 := GenS3{F0: 9223372036854775807, F1: -127}
        _ = v31
    } else {
        v1 = append(v1, "")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
