package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int; F2 uint32 }
func genFn0() int8 {
    var v4 uint32 = 1
    var v5 uint8 = 1
    v4 += (v4 | (v4 << ((704) & 63)))
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 int64 = 1
    var v2 int16 = -591
    var v3 int16 = (v2 << ((731) & 63))
    v1 += (v1 - 9223372036854775807)
    v6 := []int{}
    var v7 string = "foo bar"
    v3 = 137
    _ = len(v6)
    fmt.Println(v1)
    if (127 <= genFn0()) {
        switch 127 {
        case 2147483647:
            v8 := GenS1{F0: 1}
            _ = v8
        case 9223372036854775807:
            v7 = ""
        case 708:
            if (127 != 58) {
                if len(v6) > 0 {
                    v6[0] = 923
                }
                switch 32767 {
                case 1:
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := GenS2{F0: 2147483648}
                        var v11 uint64 = 42
                        var v12 string = v7
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                default:
                    if (!false) {
                        var v13 string = "hello"
                        if (803 >= 2147483647) {
                            if (128 <= 42) {
                                for fi14 := 0; fi14 < 1; fi14++ {
                                    if len(v6) > 0 {
                                        v6[0] = (fi14 % ((-2147483648) | 1))
                                    }
                                    v15 := GenS1{F0: 100}
                                    if (!(!true)) {
                                        v13 = v7
                                    } else {
                                        fmt.Println(v7)
                                        fmt.Println(v1)
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                            } else {
                                var v16 int32 = 42
                                v3 += v2
                                _ = v16
                            }
                            v17 := GenS1{F0: 0}
                            var v18 uint = 4294967295
                            _ = v17
                            _ = v18
                        }
                        switch "hello" {
                        case "hello":
                            var v19 uint8 = 106
                            _ = v19
                        case "foo bar":
                            var v20 string = v13
                            _ = v20
                        }
                        _ = v13
                    } else {
                        v3 = (v2 << ((607) & 63))
                    }
                }
                v3 -= v3
            }
        }
        v21 := GenS3{F0: 42, F1: 0}
        _ = v21
    }
    v2 = 95
    v1 -= v1
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 bool = (!true)
        v24 := GenS4{F0: 54, F1: fi22, F2: 2}
        _ = fi22
        _ = v23
        _ = v24
    }
    v3 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
