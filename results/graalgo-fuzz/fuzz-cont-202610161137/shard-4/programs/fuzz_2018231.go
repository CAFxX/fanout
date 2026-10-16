package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int8 = 0
    v2 := []uint32{587, 51, 654}
    var v3 int32 = -826
    var v4 float64 = (3.14159 - (3.14159 + ((((float64(1e10) - 0.1) / 3.14159) / 1.5) * 0.1)))
    var v5 uint = 42
    var v6 uint16 = 42
    var v7 int64 = 64
    _ = len(v2)
    var v8 uint64 = 728
    fmt.Println(v4, v3, v5)
    if true {
        v1 += -1
        v6 = (v6 & (v6 - v6))
        if false {
            fmt.Println(v4, v1)
        }
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "" {
        case "foo bar":
            v6 += v6
        case "hello":
            if true {
                if (true || ((!true) || false)) {
                    v7 = int64(v6)
                }
            } else {
                var v10 bool = (!((!(-1 != (-(v3)))) && false))
                v11 := GenS1{F0: v8, F1: "hello"}
                _ = v10
                _ = v11
            }
        default:
            if false {
                if true {
                    if true {
                        v5 = (v5 >> (((v5 - 42)) & 63))
                        var v12 uint8 = 1
                        _ = v12
                    }
                    v13 := GenS2{F0: 629}
                    _ = v13
                } else {
                    var v14 uint64 = uint64(v6)
                    if (("hello" >= "foo bar") && ("hello" >= "")) {
                        if ((^(v1)) <= (v1 & v1)) {
                            switch "\n\t" {
                            case "foo bar":
                                v4 += v4
                            case "":
                                v2 = append(v2, 2)
                            case "\n\t":
                                var v15 uint32 = 904
                                _ = v15
                            }
                            var v16 string = ""
                            for fi17 := 0; fi17 < 5; fi17++ {
                                v5 -= uint(v14)
                                _ = fi17
                            }
                            _ = v16
                        } else {
                            if false {
                                v2 = append(v2, 581)
                            }
                            v18 := GenS3{F0: v6, F1: v1}
                            var v19 int64 = (^(v7))
                            _ = v18
                            _ = v19
                        }
                        if false {
                            v8 = 42
                        } else {
                            var v20 int16 = -32768
                            _ = v20
                        }
                    } else {
                        fmt.Println(v4)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v2)
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        fmt.Println(v5, v4, v5)
                        var v22 uint64 = (v8 >> ((v5) & 63))
                        _ = fi21
                        _ = v22
                    }
                    _ = v14
                }
                v2 = append(v2, 917)
                v5 = (v5 % (((v5 * (v5 & 351))) | 1))
            }
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            switch "foo bar" {
            case "\n\t":
                var v24 int = (^(fi23))
                _ = v24
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v25 float64 = ((v4 / -2.25) + v4)
            _ = fi23
            _ = v25
        }
        _ = fi9
    }
    var v26 string = "a"
    for fi27 := 0; fi27 < 6; fi27++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v28 := GenS1{F0: v8, F1: v26}
        _ = fi27
        _ = v28
    }
    v29 := GenS1{F0: v8, F1: "a"}
    v1 -= -128
    if len(v2) > 0 {
        v2[0] = 875
    }
    var v30 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
