package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int8) uint16 {
    var v4 int8 = (p0 / ((p0) | 1))
    v5 := GenS1{F0: 2147483647, F1: (true && false)}
    v6 := []int{1, 0}
    v4 = 0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]string)
    var v3 uint = 255
    v7 := []int{-1}
    for fi8 := 0; fi8 < 20; fi8++ {
        if ("hello" > "foo bar") {
            v9 := GenS1{F0: 66, F1: true}
            if true {
                var v10 int = fi8
                fmt.Println(v10, v1, fi8)
                _ = v10
            } else {
                switch fi8 {
                case 244:
                    var v11 float64 = 0.0
                    _ = v11
                case 42:
                    switch "hello" {
                    case "hello":
                        v1 = v1
                    case "foo bar":
                        var v12 uint64 = 476
                        _ = v12
                    case "\n\t":
                        v13 := GenS2{F0: "\n\t", F1: 9223372036854775807}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 7; fi14++ {
                            v1 += (v1 / ((v1) | 1))
                            var v15 int8 = 100
                            var v16 int = fi8
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                }
            }
            v3 -= v3
            _ = v9
        }
        _ = fi8
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v2[65535] = "hello"
        v18 := GenS3{F0: 0}
        _ = fi17
        _ = v18
    }
    if false {
        fmt.Println(v3, v1)
        var v19 uint64 = 0
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS4{F0: false, F1: 0}
        var v22 uint32 = 1
        _ = fi20
        _ = v21
        _ = v22
    }
    v1 = v1
    v3 = v3
    v1 = v1
    v3 += (v3 << (((v3 ^ (v3 | v3))) & 63))
    for fi23 := 0; fi23 < 7; fi23++ {
        v3 -= (v3 % (((v3 | v3)) | 1))
        var v24 string = (("\n\t" + "hello") + ((("\n\t" + "\n\t") + ("foo bar" + "hello")) + ""))
        for fi25 := 0; fi25 < 7; fi25++ {
            if false {
                switch v24 {
                case "foo bar":
                    v1 = v1
                case "\n\t":
                    var v26 uint16 = 522
                    _ = v26
                case "a":
                    switch (fi25 ^ (fi25 >> ((v3) & 63))) {
                    case 127:
                        v24 = (("foo bar" + v24) + ("hello" + ("\n\t" + "\n\t")))
                    default:
                        if false {
                            v27 := GenS5{F0: 18446744073709551615, F1: 0}
                            _ = v27
                        } else {
                            delete(v2, 2147483648)
                        }
                    }
                default:
                    _ = len(v2)
                }
            } else {
                var v28 int8 = 1
                _ = v28
            }
            if (true || false) {
                if ((v1 - (v1 | (v1 << ((v3) & 63)))) >= v1) {
                    if false {
                        if false {
                            fmt.Println(v1, v24)
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v3 -= (v3 / (((v3 & (v3 ^ (v3 - v3)))) | 1))
                        v24 = v24
                    } else {
                        var v29 uint32 = 251
                        if true {
                            fmt.Println(fi23, fi23)
                        }
                        _ = v29
                    }
                    var v30 int8 = -128
                    v1 = v1
                    _ = v30
                }
                fmt.Println(fi25, v1)
                v24 = "foo bar"
            } else {
                fmt.Println(fi25, v3)
                if len(v7) > 0 {
                    v7[0] = -235
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    var v31 int32 = (v1 ^ (v1 * (v1 & (v1 % ((v1) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v7
}
