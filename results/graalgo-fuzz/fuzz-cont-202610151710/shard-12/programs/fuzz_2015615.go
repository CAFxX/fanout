package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint8) bool {
    fmt.Println(p0, p0)
    p2 += 127
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            p1 = (^(p2))
            for fi8 := 0; fi8 < 14; fi8++ {
                fmt.Println(p0)
                var v9 float64 = 1e10
                _ = fi8
                _ = v9
            }
            p2 += 127
            _ = fi7
        }
        var v10 string = "\n\t"
        _ = v10
    }
    p1 = (^(p1))
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := []int{-1, -169, 0}
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 uint32 = 2147483648
    var v5 int = 4294967295
    var v6 string = ("hello" + "hello")
    var v11 uint32 = v4
    v6 = "a"
    if (2 >= 18446744073709551615) {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 0, F1: 1, F2: 1}
            _ = fi12
            _ = v13
        }
        var v14 bool = false
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int32 = -794
            _ = fi15
            _ = v16
        }
    }
    var v17 int64 = -9223372036854775808
    switch v6 {
    case "":
        if (false && true) {
            v5 = 639
        }
    case "foo bar":
        v18 := GenS1{F0: v17, F1: -125, F2: (v17 << ((4294967295) & 63))}
        _ = v18
    default:
        fmt.Println(v4, v11)
    }
    v6 = ((v6 + ("\n\t" + v6)) + "foo bar")
    v19 := GenS2{F0: (v6 + "hello"), F1: (!(382 != uint(v4)))}
    delete(v3, v5)
    v5 = (v5 % ((-281) | 1))
    fmt.Println(v4, v5)
    v1 = append(v1, -9223372036854775808)
    fmt.Println(v4, v17, v4)
    var v20 int32 = 42
    for fi21 := 0; fi21 < 5; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    switch 4294967295 {
    case 42:
        if ((-250 >= 32767) || (42 == 885)) {
            v4 = 4294967295
            v17 += (v17 | -9223372036854775808)
        } else {
            if false {
                v4 = v11
                var v22 string = "hello"
                _ = v22
            } else {
                switch v6 {
                case "a":
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS3{F0: 42, F1: 255}
                        switch v6 {
                        case "\n\t":
                            delete(v2, (((v6 + v6) + (("a" + (v6 + v6)) + (v6 + (v6 + v6)))) + ("" + v6)))
                        case "":
                            if len(v1) > 0 {
                                v1[0] = (fi23 % ((v5) | 1))
                            }
                        case "hello":
                            v11 -= (v4 / ((v4) | 1))
                        }
                        v6 = (v6 + "foo bar")
                        _ = fi23
                        _ = v24
                    }
                }
                var v25 string = v6
                var v26 uint = 18446744073709551615
                _ = v25
                _ = v26
            }
        }
    case 32767:
        v5 -= (v5 << ((4294967295) & 63))
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (379 < 255) {
        var v27 float64 = -2.25
        var v28 uint64 = 18446744073709551615
        _ = v27
        _ = v28
    } else {
        v29 := GenS3{F0: 1000, F1: 1}
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int32 = 0
            v32 := GenS2{F0: "\n\t", F1: false}
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
