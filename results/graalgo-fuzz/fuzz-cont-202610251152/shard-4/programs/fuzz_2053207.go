package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 1
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int32 = 0
        _ = fi5
        _ = v6
    }
    delete(v3, -2147483648)
    v2 = v2
    var v7 int16 = 0
    v8 := v4[652]
    var v9 uint16 = 60
    v2 -= v2
    v2 -= (v2 & v2)
    v10 := []int{42, -826}
    switch -1 {
    case 646:
        v11 := GenS1{F0: v7, F1: 152, F2: -891}
        _ = v11
    case 4294967295:
        if true {
            fmt.Println(v8)
        } else {
            fmt.Println(v9, v2, v9)
            if true {
                v2 += v2
            } else {
                if (2 == 0) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 int8 = -128
                            if true {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v14, fi12)
                                case "\n\t":
                                    v2 -= v2
                                    var v15 int64 = -673
                                    v14 = (v14 ^ (^(v14)))
                                    _ = v15
                                case "hello":
                                    fmt.Println(fi13)
                                default:
                                    var v16 string = v8
                                    _ = v16
                                }
                            } else {
                                var v17 uint64 = 42
                                fmt.Println(v7, v2)
                                if false {
                                    v18 := GenS1{F0: (v7 | -32768), F1: 128, F2: fi13}
                                    if (67 <= 255) {
                                        var v19 uint8 = 1
                                        _ = v19
                                    }
                                    _ = v18
                                }
                                _ = v17
                            }
                            v7 += (v7 ^ v7)
                            _ = fi13
                            _ = v14
                        }
                        delete(v1, 2147483648)
                        if true {
                            if (961 != 2147483647) {
                                var v20 int8 = 1
                                _ = v20
                            }
                            v21 := GenS1{F0: v7, F1: 128, F2: int(v7)}
                            for fi22 := 0; fi22 < 18; fi22++ {
                                v23 := GenS2{F0: 2}
                                v9 += 255
                                fmt.Println(v2, v7, v2)
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        } else {
                            var v24 float32 = (-2.25 / 1e10)
                            _ = v24
                        }
                        _ = fi12
                    }
                    v4[660] = "a"
                    v2 = (v2 + 441)
                } else {
                    v25 := GenS2{F0: 1}
                    v8 = v8
                    _ = v25
                }
                for fi26 := 0; fi26 < 12; fi26++ {
                    v8 = v8
                    _ = fi26
                }
                v7 += v7
            }
            delete(v1, 0)
        }
    case 2147483647:
        delete(v4, 437)
    default:
        var v27 uint8 = 127
        _ = v27
    }
    v8 = (("" + v8) + "\n\t")
    _ = len(v10)
    var v28 int64 = (v2 << ((1) & 63))
    v29 := GenS1{F0: (v7 >> ((886) & 63)), F1: 0, F2: 4294967295}
    var v30 int64 = v2
    var v31 int16 = 1000
    v8 = ((v8 + v8) + v8)
    var v32 string = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
