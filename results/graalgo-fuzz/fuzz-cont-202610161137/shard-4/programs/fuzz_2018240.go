package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0(p0 int, p1 uint16, p2 int64) int16 {
    var v4 int = p0
    v5 := []int{(v4 << ((358) & 63)), v4}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 string) bool {
    p0 = p0
    var v6 bool = false
    _ = p0
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 1
    var v3 int64 = 42
    var v7 int = 9223372036854775807
    v3 -= v3
    var v8 uint = 0
    switch "foo bar" {
    case "":
        fmt.Println(v7, v3)
    case "hello":
        for fi9 := 0; fi9 < 13; fi9++ {
            if (!true) {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 bool = false
                    var v12 bool = v11
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := v1[(v7 & -2147483648)]
                switch (v13 + v13) {
                case "":
                    v14 := GenS1{F0: v3, F1: 42, F2: 148}
                    _ = v14
                case "\n\t":
                    v3 += 0
                case "foo bar":
                    switch v13 {
                    case "":
                        var v15 int32 = -192
                        _ = v15
                    case "hello":
                        v3 -= v3
                    default:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v8 -= 65535
                            if (!(324 < v3)) {
                                v13 = v13
                                switch v13 {
                                case "foo bar":
                                    v2 = v2
                                case "hello":
                                    if ((float32(1e10) > 0.0) && (32767 >= 42)) {
                                        var v17 string = "a"
                                        var v18 uint64 = 18446744073709551615
                                        v1[947] = (v17 + (v17 + "\n\t"))
                                        _ = v17
                                        _ = v18
                                    }
                                }
                            } else {
                                v19 := GenS2{F0: 0, F1: 42, F2: (v3 * (-(v3)))}
                                _ = v19
                            }
                            fmt.Println(v7)
                            _ = fi16
                        }
                    }
                }
                _ = v13
            } else {
                v3 = (v3 * v3)
            }
            v2 = 0
            fmt.Println(v8, v3, v8)
            _ = fi9
        }
    }
    fmt.Println(v7)
    v3 = (v3 - (v3 % ((v3) | 1)))
    if false {
        var v20 int16 = -512
        var v21 uint8 = 128
        fmt.Println(v21, v20)
        _ = v20
        _ = v21
    } else {
        if ("" == ("" + "a")) {
            switch "hello" {
            case "foo bar":
                var v22 uint32 = 42
                _ = v22
            }
            v23 := GenS3{F0: v8, F1: 1000, F2: 0}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                fmt.Println(v8)
                fmt.Println(v3)
                _ = fi24
            }
            v1[1] = ((("hello" + "a") + "a") + ("\n\t" + "a"))
            v8 = (^(v2))
        }
    }
    var v25 uint32 = 289
    for fi26 := 0; fi26 < 15; fi26++ {
        if (42 > -1) {
            v1[fi26] = "hello"
        }
        _ = fi26
    }
    if (0 < 731) {
        v8 = v8
        if true {
            v25 = 90
            v27 := []int{-1}
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
