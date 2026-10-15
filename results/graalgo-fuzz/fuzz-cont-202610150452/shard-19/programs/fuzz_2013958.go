package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int16) uint {
    if true {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p0, p0)
        fmt.Println(p0, p1)
        var v6 string = ""
        _ = v6
    }
    var v7 uint64 = 42
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 int32, p2 uint32) uint8 {
    switch -758 {
    case 1:
        var v8 int32 = (-(p1))
        _ = v8
    case 2147483647:
        fmt.Println(p2)
    default:
        p2 -= uint32(p0)
    }
    fmt.Println(p1, p1, p1)
    p1 -= (p1 << ((300) & 63))
    var v9 float32 = float32(3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    v1 := []uint32{2, 0}
    v2 := []int{}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint = 2
    switch 9223372036854775807 {
    case -1:
        if true {
            v5 += v5
        } else {
            v10 := v4[2147483647]
            v11 := GenS1{F0: 151, F1: 65535, F2: 9223372036854775807}
            _ = v10
            _ = v11
        }
    case 32768:
        v12 := GenS2{F0: ("foo bar" + "a")}
        _ = v12
    default:
        fmt.Println(v5)
    }
    v5 = v5
    _ = len(v4)
    var v13 float64 = float64(-2.25)
    var v14 int32 = -1
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint64 = 18446744073709551615
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 int8 = 1
        fmt.Println(v18)
        _ = fi17
        _ = v18
    }
    switch 128 {
    case 128:
        if len(v1) > 0 {
            v1[0] = 248
        }
    case 32768:
        switch -9223372036854775808 {
        case 1:
            v19 := GenS3{F0: 286, F1: -1}
            _ = v19
        case 0:
            if true {
                var v20 uint8 = 128
                if true {
                    switch 32767 {
                    case 465:
                        if (false && true) {
                            if true {
                                v1 = append(v1, 65535)
                                var v21 bool = false
                                fmt.Println(v5)
                                _ = v21
                            }
                        } else {
                            v20 += 42
                            var v22 int64 = -1
                            for fi23 := 0; fi23 < 18; fi23++ {
                                _ = len(v3)
                                _ = fi23
                            }
                            _ = v22
                        }
                    case 42:
                        fmt.Println(v14, v20, v20)
                    default:
                        switch "a" {
                        case "":
                            switch "\n\t" {
                            case "foo bar":
                                v5 = (v5 % ((42) | 1))
                            case "\n\t":
                                delete(v4, 4294967295)
                            case "a":
                                var v24 string = ("a" + ((("hello" + "foo bar") + "\n\t") + "foo bar"))
                                _ = v24
                            }
                        }
                    }
                    v5 += (v5 * v5)
                }
                v20 = v20
                _ = v20
            } else {
                v5 -= 2
            }
        case 4294967295:
            var v25 string = "a"
            _ = v25
        default:
            fmt.Println(v14, v13)
        }
    case 4294967295:
        v13 = float64((1.5 / ((float32(v13) / 1.0) * ((3.14159 + 1.5) - ((1.5 + ((((0.0 / 3.14159) + 0.1) - 0.0) + 0.1)) - 3.14159)))))
    default:
        v26 := GenS4{F0: 174}
        _ = v26
    }
    switch "a" {
    case "a":
        var v27 int64 = 9223372036854775807
        _ = v27
    case "foo bar":
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 int16 = -918
            if ("\n\t" == ("\n\t" + "a")) {
                if false {
                    var v30 uint64 = 18446744073709551615
                    var v31 uint32 = 2
                    _ = v30
                    _ = v31
                }
                fmt.Println(v14)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi28
            _ = v29
        }
    default:
        v5 = (v5 / ((v5) | 1))
    }
    v14 = v14
    fmt.Println(v5)
    fmt.Println(v13, v13, v13)
    var v32 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
