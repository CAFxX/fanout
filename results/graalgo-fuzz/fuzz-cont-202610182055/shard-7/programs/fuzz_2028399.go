package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 bool; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint, p2 int64) string {
    fmt.Println(p1, p1, p2)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -639
    v3 := make(map[int]string)
    var v4 int64 = -1
    var v5 int8 = 1
    v6 := []uint32{2, 2, 765}
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: (^(fi8))}
            _ = fi8
            _ = v9
        }
    case "":
        var v10 int8 = v5
        _ = v10
    case "foo bar":
        fmt.Println(v2)
    }
    v2 += v2
    fmt.Println(v5, v4)
    v4 += v4
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            v5 = v5
            var v12 bool = (254 < 255)
            _ = v12
        }
        _ = fi11
    }
    v2 = (v2 / (((-(v2))) | 1))
    var v13 int8 = v5
    if false {
        v5 += v13
    } else {
        v14 := GenS2{F0: -564}
        switch 9223372036854775807 {
        case -2147483648:
            fmt.Println(v4, v4, v4)
        case 32768:
            for fi15 := 0; fi15 < 10; fi15++ {
                v4 = -1
                _ = fi15
            }
        case 0:
            if false {
                if (!(false && true)) {
                    var v16 float32 = -2.25
                    v17 := GenS3{F0: 475, F1: (!(!false)), F2: 120}
                    v18 := GenS2{F0: 32767}
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v19 := GenS4{F0: 65535}
                    fmt.Println(v4)
                    switch ("foo bar" + "a") {
                    case "a":
                        var v20 uint16 = 255
                        _ = v20
                    case "foo bar":
                        if true {
                            v21 := v1[32767]
                            _ = v21
                        }
                    case "hello":
                        for fi22 := 0; fi22 < 15; fi22++ {
                            v23 := v3[(-(fi22))]
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = v19
                }
                v2 -= v2
                var v24 uint = 18446744073709551615
                _ = v24
            } else {
                v2 = v2
            }
        default:
            v25 := GenS3{F0: 2147483648, F1: (87 >= 42), F2: 32768}
            _ = v25
        }
        _ = v14
    }
    v13 = (v13 & v13)
    fmt.Println(v5, v2)
    var v26 float64 = ((-2.25 * 1.5) / 1.0)
    if false {
        v5 -= (v5 + (v13 << ((65535) & 63)))
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 float64 = v26
        var v29 uint64 = 4294967295
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
