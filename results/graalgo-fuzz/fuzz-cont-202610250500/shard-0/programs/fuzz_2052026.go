package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := []string{"\n\t", "", "a"}
    v2 := make(map[int]uint32)
    v3 := []int{}
    var v4 float64 = -2.25
    var v5 bool = (!true)
    fmt.Println(v4, v5)
    v6 := GenS1{F0: 42, F1: -2147483648}
    fmt.Println(v4)
    var v7 int64 = 42
    var v8 int16 = -32768
    v5 = (672 == 2147483647)
    switch ("\n\t" + "") {
    case "\n\t":
        v9 := GenS1{F0: -41, F1: -443}
        _ = v9
    case "foo bar":
        fmt.Println(v8, v7, v7)
    }
    v7 = (v7 / (((v7 - v7)) | 1))
    if (65535 != 975) {
        v10 := v2[-1]
        delete(v2, -2147483648)
        v11 := GenS1{F0: 0, F1: 246}
        _ = v10
        _ = v11
    }
    fmt.Println(v8, v7)
    if (v4 < (v4 - (1.5 - 1e10))) {
        delete(v2, 127)
        if true {
            if true {
                var v12 float32 = 3.14159
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint16 = 68
                var v15 uint32 = 405
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint16 = 42
                for fi18 := 0; fi18 < 9; fi18++ {
                    fmt.Println(fi18)
                    v5 = (!true)
                    v19 := GenS1{F0: -1, F1: 2147483647}
                    _ = fi18
                    _ = v19
                }
                v7 -= 42
                _ = fi16
                _ = v17
            }
        } else {
            var v20 uint = 42
            if ((false || v5) || v5) {
                for fi21 := 0; fi21 < 5; fi21++ {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        if v5 {
                            if (false || ((fi22 + (fi22 ^ fi22)) <= fi21)) {
                                v20 = (^(v20))
                                v23 := GenS1{F0: 0, F1: 322}
                                _ = len(v2)
                                _ = v23
                            } else {
                                _ = len(v1)
                                v4 -= (v4 - v4)
                                v24 := GenS1{F0: 0, F1: 0}
                                _ = v24
                            }
                            var v25 uint32 = 42
                            _ = v25
                        }
                        v20 = 195
                        _ = fi22
                    }
                    switch (fi21 / ((fi21) | 1)) {
                    case 1:
                        v7 = 430
                    case 816:
                        v26 := GenS2{F0: fi21, F1: (v7 | v7)}
                        _ = v26
                    case 2147483647:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v28 := GenS2{F0: fi21, F1: 9223372036854775807}
                            v29 := GenS2{F0: fi21, F1: 9223372036854775807}
                            v8 -= v8
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 float32 = ((1e10 * (1e10 - -2.25)) / float32(float64((float32(v4) + ((3.14159 + float32((0.0 / v4))) / 1.0)))))
                        _ = v30
                    }
                    _ = len(v1)
                    _ = fi21
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v3 = append(v3, 32768)
            } else {
                v5 = v5
            }
            _ = v20
        }
    } else {
        var v31 uint16 = 661
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
