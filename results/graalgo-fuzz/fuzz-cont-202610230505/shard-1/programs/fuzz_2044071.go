package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint }
func genFn0() int16 {
    if (221 < 42) {
        var v7 uint = 255
        fmt.Println(v7)
        v8 := GenS1{F0: 100}
        _ = v7
        _ = v8
    }
    var v9 int8 = 2
    _ = v9
    return 0
}
func genFn1(p0 uint, p1 int8, p2 int8) int64 {
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10, p0)
        _ = fi10
    }
    v11 := []int{2147483647, -658}
    p2 -= 87
    switch 316 {
    case -804:
        if ((p0 / (((^(p0))) | 1)) != 255) {
            var v12 uint32 = 750
            v13 := GenS1{F0: p1}
            if true {
                fmt.Println(p1, p1)
                if (false || (false && false)) {
                    var v14 int8 = p2
                    switch 752 {
                    case -9223372036854775808:
                        fmt.Println(p1)
                    case 0:
                        v11 = append(v11, 1)
                    case 9223372036854775807:
                        var v15 uint8 = 255
                        _ = v15
                    default:
                        if (4294967295 >= 0) {
                            v16 := GenS1{F0: p1}
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    v14 -= (p1 | p1)
                                    var v19 int = (fi17 - (fi18 - fi17))
                                    if (true && true) {
                                        v20 := GenS1{F0: (p1 / ((p1) | 1))}
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        v14 += p2
                                        _ = v20
                                    } else {
                                        var v21 int8 = (p2 ^ v14)
                                        _ = v21
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                _ = fi17
                            }
                            if (128 >= 32768) {
                                switch "" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "a":
                                        v12 -= 482
                                    case "\n\t":
                                        switch -631 {
                                        case 558:
                                            v22 := GenS1{F0: (v14 >> ((p0) & 63))}
                                            _ = v22
                                        default:
                                            for fi23 := 0; fi23 < 3; fi23++ {
                                                v14 -= 100
                                                v24 := GenS2{F0: uint64(v12)}
                                                _ = fi23
                                                _ = v24
                                            }
                                        }
                                    case "foo bar":
                                        p0 -= p0
                                    }
                                case "hello":
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                            } else {
                                v25 := GenS3{F0: -1, F1: (^(p0))}
                                _ = v25
                            }
                            _ = v16
                        }
                    }
                    _ = v14
                }
                p0 -= (^(p0))
            } else {
                var v26 bool = (p0 > uint(v12))
                p1 += (p1 + int8(v12))
                _ = v26
            }
            _ = v12
            _ = v13
        } else {
            fmt.Println(p2)
        }
    case 32768:
        var v27 uint64 = 455
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 int8 = 110
    v2 := make(map[uint32]int)
    var v3 int8 = v1
    var v4 uint8 = 2
    v5 := []string{(("a" + "") + ("\n\t" + "hello"))}
    var v6 bool = (1 <= 65535)
    for fi28 := 0; fi28 < 5; fi28++ {
        if (-32768 <= -987) {
            v4 += v4
        }
        _ = fi28
    }
    var v29 float32 = -2.25
    v29 += v29
    var v30 uint = 995
    fmt.Println(v4, v4, v1)
    if ((v1 << (((v30 * (v30 << ((2) & 63)))) & 63)) == v1) {
        fmt.Println(v29, v4)
    } else {
        var v31 bool = v6
        _ = v31
    }
    var v32 string = "foo bar"
    v29 += 1e10
    fmt.Println(v30, v29)
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 int16 = 42
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
