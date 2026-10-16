package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = -730
    var v2 int8 = -106
    var v3 int8 = (v2 * v2)
    v4 := []string{("foo bar" + "hello")}
    var v5 bool = true
    if len(v4) > 0 {
        v4[0] = ("" + ("hello" + "\n\t"))
    }
    v5 = v5
    var v6 float64 = (1e10 + 0.1)
    var v7 uint32 = 828
    switch 0 {
    case 127:
        var v8 uint32 = (v7 + (^(v7)))
        _ = v8
    case -391:
        var v9 int8 = v3
        _ = v9
    case 4294967295:
        if len(v4) > 0 {
            v4[0] = "a"
        }
    default:
        fmt.Println(v7, v3, v2)
    }
    v5 = true
    for fi10 := 0; fi10 < 2; fi10++ {
        if (v5 && v5) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v2 += v2
                if v5 {
                    v12 := GenS1{F0: (v1 - 42)}
                    var v13 int16 = 0
                    _ = v12
                    _ = v13
                } else {
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v5 {
                            switch "\n\t" {
                            case "\n\t":
                                v3 = v2
                            case "":
                                v5 = (65535 == 0)
                            case "a":
                                if (100 == v2) {
                                    if true {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v1 -= int32(v7)
                                        v6 += 0.1
                                        v14 := GenS1{F0: (v1 << ((255) & 63))}
                                        fmt.Println(v1, v1, v2)
                                        _ = v14
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = len(v4)
                                    }
                                } else {
                                    v1 -= v1
                                }
                            default:
                                fmt.Println(v7, fi11)
                            }
                            fmt.Println(fi10, fi10, v3)
                        }
                        fmt.Println(v6, v6, fi11)
                    }
                    fmt.Println(v3)
                }
                var v15 bool = ((0 != 1000) && false)
                var v16 bool = (!true)
                _ = fi11
                _ = v15
                _ = v16
            }
        } else {
            fmt.Println(v2, v3)
            v2 += v3
        }
        var v17 int32 = 42
        _ = fi10
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint64 = 940
        fmt.Println(v2)
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 32767, F1: 42, F2: 2147483647}
    v6 -= v6
    v7 = v7
    v1 -= (v1 | v1)
    var v21 uint = 0
    v22 := GenS3{F0: int16(v21)}
    var v23 int64 = 1
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
