package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    switch -1 {
    case 670:
        for fi7 := 0; fi7 < 11; fi7++ {
            switch "\n\t" {
            case "a":
                fmt.Println(fi7)
            }
            fmt.Println(fi7)
            fmt.Println(fi7)
            _ = fi7
        }
    }
    var v8 uint64 = 601
    var v9 int = 4294967295
    v9 -= v9
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint32 = 42
    v2 := []int{9223372036854775807, 127}
    var v3 float32 = 1.5
    v4 := make(map[string]string)
    var v5 uint16 = 960
    v6 := make(map[int]bool)
    v10 := GenS1{F0: (v1 % ((v1) | 1))}
    v1 = v1
    fmt.Println(v3, v5, v5)
    v5 += v5
    var v11 int8 = 0
    var v12 string = "a"
    v11 += (v11 & (v11 % ((42) | 1)))
    v13 := GenS2{F0: 65535}
    v3 = v3
    fmt.Println(v12, v5)
    var v14 uint16 = 0
    switch v12 {
    case "a":
        delete(v6, -943)
    default:
        if (true && true) {
            var v15 int32 = -1
            _ = v15
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                v3 += v3
                switch "foo bar" {
                case "foo bar":
                    if (42 > v11) {
                        var v17 int32 = 620
                        switch v12 {
                        case "a":
                            _ = len(v2)
                        case "hello":
                            v14 += v14
                        case "":
                            v12 = v12
                        }
                        _ = v17
                    }
                case "\n\t":
                    if (2 != 18446744073709551615) {
                        switch v12 {
                        case "a":
                            for fi18 := 0; fi18 < 7; fi18++ {
                                v14 = (v5 << ((113) & 63))
                                v2 = append(v2, fi18)
                                _ = fi18
                            }
                        case "":
                            v3 -= (float32((((float64(1.5) * 3.14159) + 0.0) + (((-2.25 + 3.14159) * 1.5) - 3.14159))) - v3)
                        default:
                            if (true && true) {
                                v12 = (("hello" + (v12 + "foo bar")) + v12)
                                var v19 uint32 = (^(v1))
                                var v20 int8 = v11
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                case "hello":
                    _ = len(v4)
                default:
                    var v21 uint16 = (v14 >> ((65535) & 63))
                    _ = v21
                }
                v5 = (v14 - v5)
                _ = fi16
            }
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint16 = v14
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 bool = (0 != 1)
                var v26 uint64 = 18446744073709551615
                switch v12 {
                case "":
                    v27 := GenS2{F0: v1}
                    _ = v27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            var v28 uint8 = 128
            var v29 uint8 = 128
            fmt.Println(v3, v14)
            _ = v28
            _ = v29
        }
        v30 := v4[v12]
        _ = fi22
        _ = v23
        _ = v30
    }
    v31 := v6[9223372036854775807]
    var v32 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
