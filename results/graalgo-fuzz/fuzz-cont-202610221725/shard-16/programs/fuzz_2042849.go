package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 uint32) int {
    var v7 float32 = 3.14159
    var v8 int16 = 224
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return -1
}
func genFn1() int8 {
    var v9 string = ("\n\t" + "foo bar")
    v9 = "a"
    _ = v9
    return 1
}

func main() {
    var v1 uint = 500
    var v2 int64 = -632
    v3 := make(map[uint32]bool)
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]int)
    v6 := []string{"", "hello"}
    v10 := GenS1{F0: 2}
    v11 := GenS1{F0: 246}
    var v12 int32 = 421
    if ((v2 - v2) == v2) {
        v13 := GenS1{F0: 0}
        _ = v13
    }
    v12 = -252
    v4 = genFn0(402)
    fmt.Println(v4, v4)
    if true {
        delete(v5, 194)
    } else {
        var v14 bool = ("" >= (("" + "foo bar") + "hello"))
        _ = v14
    }
    fmt.Println(v1)
    v12 = (v12 + v12)
    if (!(-883 == 1000)) {
        var v15 float64 = 0.0
        var v16 int = v4
        v16 += (v16 ^ -461)
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 180}
        for fi18 := 0; fi18 < 1; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                if (125 == 128) {
                    var v20 int64 = -9223372036854775808
                    v6 = append(v6, ("hello" + ""))
                    _ = v20
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if ((-(v2)) == (v2 << (((v1 - (v1 << ((18446744073709551615) & 63)))) & 63))) {
                    if len(v6) > 0 {
                        v6[0] = "a"
                    }
                    switch "" {
                    case "hello":
                        if (1 >= 0) {
                            v21 := GenS1{F0: 255}
                            v3[2] = false
                            _ = v21
                        } else {
                            if false {
                                v22 := GenS2{F0: 42, F1: 171, F2: 42}
                                switch "" {
                                case "foo bar":
                                    if false {
                                        v23 := GenS3{F0: 864, F1: 0}
                                        _ = v23
                                    }
                                case "":
                                    fmt.Println(v4, v12)
                                default:
                                    v2 -= (v2 / ((1) | 1))
                                }
                                _ = len(v5)
                                _ = v22
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            v24 := GenS4{F0: 85, F1: 128, F2: 0}
                            _ = v24
                        }
                    default:
                        var v25 float64 = -2.25
                        _ = v25
                    }
                }
                _ = fi19
            }
            _ = fi18
        }
        _ = v17
    }
    v2 = (v2 & 0)
    _ = len(v6)
    switch "foo bar" {
    case "\n\t":
        if len(v6) > 0 {
            v6[0] = "foo bar"
        }
    case "":
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
