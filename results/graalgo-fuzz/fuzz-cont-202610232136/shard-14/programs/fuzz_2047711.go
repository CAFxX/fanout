package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0() int16 {
    var v7 uint = 154
    v7 -= 18446744073709551615
    _ = v7
    return 42
}
func genFn1(p0 string) uint16 {
    switch 405 {
    case -2147483648:
        fmt.Println(p0)
    case 4294967295:
        fmt.Println(p0)
    default:
        if (0.0 == 0.1) {
            p0 = (p0 + "\n\t")
        } else {
            if (1000 < -242) {
                var v8 float64 = 1e10
                if false {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        fmt.Println(p0)
                        for fi10 := 0; fi10 < 4; fi10++ {
                            var v11 int8 = 0
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                    fmt.Println(p0, p0)
                } else {
                    if false {
                        switch "a" {
                        case "hello":
                            if true {
                                fmt.Println(v8)
                            }
                        case "\n\t":
                            v12 := []int{-770, 1, -9223372036854775808, -58}
                            _ = v12
                        case "a":
                            if false {
                                v13 := []int{}
                                fmt.Println(p0, v8)
                                if (-361 > 966) {
                                    var v14 string = (p0 + p0)
                                    var v15 int = 2147483647
                                    _ = v14
                                    _ = v15
                                }
                                _ = v13
                            }
                        }
                        var v16 int64 = 683
                        _ = v16
                    }
                    var v17 int64 = 1
                    _ = v17
                }
                _ = v8
            } else {
                v18 := GenS1{F0: 569, F1: "foo bar"}
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 int16 = 32767
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
            fmt.Println(p0)
        }
    }
    var v21 uint32 = 65535
    _ = p0
    _ = v21
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{981}
    var v3 uint16 = 478
    var v4 float32 = 1e10
    v5 := []int{2147483647, -36, (^(v1))}
    v6 := make(map[uint32]uint32)
    for fi22 := 0; fi22 < 11; fi22++ {
        v2 = append(v2, (fi22 % (((^(v1))) | 1)))
        _ = fi22
    }
    v3 = 430
    v23 := GenS1{F0: 1, F1: ("a" + "\n\t")}
    var v24 uint64 = 4294967295
    v3 += 42
    v4 += (v4 - 3.14159)
    var v25 bool = true
    fmt.Println(v25)
    v1 = v1
    fmt.Println(v3, v25)
    v3 += v3
    var v26 uint8 = 42
    var v27 uint64 = v24
    var v28 float32 = (v4 * v4)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
