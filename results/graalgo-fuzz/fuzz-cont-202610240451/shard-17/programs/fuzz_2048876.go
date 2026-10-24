package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 uint8, p1 uint8, p2 int16) int8 {
    fmt.Println(p2, p0, p0)
    p0 = 1
    var v6 int32 = -1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint16 = 1
    var v2 string = "\n\t"
    v3 := []uint32{942}
    v4 := make(map[string]bool)
    var v5 int16 = 0
    fmt.Println(v2)
    if (654 > 312) {
        switch 9223372036854775807 {
        case 4294967295:
            delete(v4, ((v2 + v2) + v2))
        case 42:
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS1{F0: 42}
                if true {
                    v4[("foo bar" + (("a" + (v2 + v2)) + v2))] = (true && true)
                } else {
                    v2 = ("hello" + "")
                    v9 := GenS1{F0: 126}
                    var v10 int32 = 843
                    _ = v10
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        default:
            switch v2 {
            case "hello":
                var v11 uint = 42
                _ = v11
            case "a":
                switch 0 {
                case -9223372036854775808:
                    v2 = v2
                default:
                    var v12 string = "a"
                    _ = v12
                }
            case "foo bar":
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := v4["hello"]
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v16 := v4[""]
                        if v14 {
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            v16 = v14
                            _ = len(v3)
                        }
                        var v17 uint = 42
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    if false {
                        var v18 bool = (v1 > (^(v1)))
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
            default:
                v5 -= v5
            }
        }
        v1 = (^(v1))
    } else {
        v5 = (v5 + (^(v5)))
        var v19 string = v2
        _ = v19
    }
    switch (v2 + v2) {
    case "hello":
        for fi20 := 0; fi20 < 12; fi20++ {
            if false {
                var v21 uint64 = 18446744073709551615
                fmt.Println(v1)
                var v22 float64 = (float64(3.14159) / 1.0)
                _ = v21
                _ = v22
            }
            v1 -= (^(v1))
            fmt.Println(v1)
            _ = fi20
        }
    case "\n\t":
        var v23 bool = (490 > 792)
        _ = v23
    case "a":
        var v24 float32 = 1.5
        _ = v24
    }
    switch 42 {
    case 162:
        if false {
            fmt.Println(v1)
        }
    default:
        v25 := GenS2{F0: -507, F1: 552}
        _ = v25
    }
    var v26 int = 0
    fmt.Println(v2, v1, v1)
    v27 := GenS1{F0: 992}
    v28 := GenS2{F0: -630, F1: 65535}
    v26 += (-(v26))
    fmt.Println(v1)
    v3 = append(v3, 0)
    for fi29 := 0; fi29 < 12; fi29++ {
        v2 = v2
        v30 := v4[v2]
        _ = fi29
        _ = v30
    }
    fmt.Println(v5)
    v3 = append(v3, 2)
    v31 := GenS3{F0: -32768, F1: v5, F2: (v26 | v26)}
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
