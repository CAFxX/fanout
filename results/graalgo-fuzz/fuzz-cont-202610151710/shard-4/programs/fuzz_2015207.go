package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
type GenS5 struct { F0 string; F1 int16 }

func main() {
    var v1 int16 = -32768
    var v2 string = "a"
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    if true {
        var v6 float64 = (1.5 / float64(-2.25))
        _ = v6
    } else {
        v1 = 1
        if false {
            fmt.Println(v3, v3, v2)
            if (0.1 == -2.25) {
                v3 -= v3
            }
        }
        v7 := []int{}
        _ = v7
    }
    v1 = -221
    if ((-74 > 0) && true) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
    } else {
        v10 := GenS2{F0: 102}
        v11 := []int{2147483647, 42, 4294967295}
        switch "" {
        case "foo bar":
            v12 := GenS2{F0: uint8(v3)}
            _ = v12
        case "a":
            v13 := GenS3{F0: 65535, F1: 4294967295}
            _ = v13
        case "\n\t":
            delete(v5, "foo bar")
        }
        _ = v10
        _ = v11
    }
    var v14 float64 = -2.25
    var v15 float32 = 3.14159
    if false {
        v14 += (1e10 * (v14 + v14))
        if true {
            if (((((v1 % ((1) | 1)) <= v1) || (!(!(!(99 == 1))))) && false) || true) {
                v16 := GenS4{F0: 0, F1: (-(v1)), F2: true}
                _ = len(v4)
                _ = v16
            }
            if false {
                if ((!false) || true) {
                    if false {
                        switch v2 {
                        case "":
                            v1 = (v1 | 814)
                        case "hello":
                            switch (v2 + v2) {
                            case "a":
                                v17 := GenS5{F0: (v2 + (v2 + (v2 + v2))), F1: v1}
                                _ = v17
                            case "hello":
                                v3 += 94
                            default:
                                var v18 uint64 = (v3 >> ((2) & 63))
                                _ = v18
                            }
                        case "\n\t":
                            v19 := GenS3{F0: 279, F1: 260}
                            _ = v19
                        default:
                            v20 := GenS1{F0: 1}
                            _ = v20
                        }
                    } else {
                        _ = len(v4)
                        switch 4294967295 {
                        case -9223372036854775808:
                            v14 += 0.1
                        case -1:
                            var v21 uint8 = 255
                            _ = v21
                        }
                        v15 = 3.14159
                    }
                    v22 := []int{32767, 42}
                    if (false || false) {
                        v14 -= v14
                    }
                    _ = v22
                } else {
                    v2 = v2
                }
                var v23 uint64 = v3
                _ = v23
            }
            delete(v5, ("a" + "hello"))
        } else {
            v1 = (v1 | (v1 - 0))
        }
    } else {
        v15 = v15
        var v24 float64 = v14
        _ = v24
    }
    fmt.Println(v14)
    v25 := GenS3{F0: 4294967295, F1: 32768}
    v3 = (v3 * v3)
    for fi26 := 0; fi26 < 10; fi26++ {
        v2 = "a"
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
