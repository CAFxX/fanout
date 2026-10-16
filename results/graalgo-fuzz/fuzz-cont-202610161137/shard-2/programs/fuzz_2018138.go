package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    var v1 float64 = 1e10
    var v2 uint32 = 2147483648
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 18446744073709551615}
    v6 := GenS1{F0: 29}
    v2 -= v2
    v7 := []int{9223372036854775807, 2147483647, 32767, -1}
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int32 = 1
            _ = fi8
            _ = v9
        }
    }
    switch ("" + ("\n\t" + "")) {
    case "foo bar":
        var v10 int = -9223372036854775808
        _ = v10
    default:
        var v11 string = ""
        _ = v11
    }
    v2 -= (v2 | (^(v2)))
    var v12 int16 = 1
    for fi13 := 0; fi13 < 8; fi13++ {
        v4[(v2 & (v2 | (^(v2))))] = -1
        v14 := GenS1{F0: 412}
        _ = fi13
        _ = v14
    }
    switch 831 {
    case 42:
        v12 -= v12
    case -228:
        fmt.Println(v2)
    default:
        if true {
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 int32 = 42
                _ = fi15
                _ = v16
            }
        } else {
            if true {
                v17 := GenS2{F0: 2, F1: 0, F2: 1}
                switch "" {
                case "\n\t":
                    for fi18 := 0; fi18 < 10; fi18++ {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            switch "foo bar" {
                            case "hello":
                                switch (fi18 - 381) {
                                case -923:
                                    if (false || (!(!(!false)))) {
                                        if false {
                                            v20 := GenS1{F0: 0}
                                            _ = v20
                                        }
                                        var v21 uint16 = 2
                                        var v22 int32 = -942
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            case "\n\t":
                                var v23 int8 = -1
                                _ = v23
                            case "":
                                if (v12 >= (v12 & (^(v12)))) {
                                    if true {
                                        fmt.Println(fi18, fi18, v1)
                                        var v24 int = (fi19 ^ (fi19 % (((fi18 + (fi18 * -316))) | 1)))
                                        var v25 int32 = 586
                                        _ = v24
                                        _ = v25
                                    }
                                    v12 = v12
                                    v2 -= (v2 ^ (v2 % ((1) | 1)))
                                    delete(v3, fi18)
                                } else {
                                    fmt.Println(v1, fi18)
                                }
                            }
                            _ = fi19
                        }
                        var v26 int32 = -1
                        v26 += 42
                        _ = fi18
                        _ = v26
                    }
                default:
                    var v27 uint64 = 1
                    _ = v27
                }
                _ = v17
            }
            fmt.Println(v12, v2)
        }
    }
    var v28 string = ("foo bar" + "a")
    var v29 uint8 = 0
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 int32 = -2147483648
        _ = fi30
        _ = v31
    }
    var v32 int64 = 1
    var v33 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
