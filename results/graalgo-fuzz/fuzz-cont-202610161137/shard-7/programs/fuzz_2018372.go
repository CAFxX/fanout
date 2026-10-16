package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 uint32, p2 int8) int8 {
    p0 = "a"
    for fi7 := 0; fi7 < 2; fi7++ {
        if (0 <= 819) {
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(p1, p2)
                var v9 int = fi8
                v10 := []int{v9, (v9 - (^(fi8))), fi7, (v9 - 128)}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            var v11 int8 = (p2 / ((-113) | 1))
            _ = v11
        }
        var v12 uint64 = 1
        _ = fi7
        _ = v12
    }
    var v13 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    v3 := []int{4294967295, 4294967295, -869, 32767}
    v4 := []uint32{93, 0, 42}
    v5 := make(map[uint32]bool)
    var v6 string = (("foo bar" + ("" + "hello")) + "foo bar")
    var v14 uint64 = 4294967295
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            v3 = append(v3, fi16)
            var v17 bool = true
            if (v17 || v17) {
                if (!(!(103 >= -102))) {
                    if (!v2) {
                        var v18 int64 = 42
                        if v2 {
                            if v2 {
                                v1[2147483647] = "\n\t"
                            } else {
                                if false {
                                    v2 = v2
                                    var v19 int8 = -1
                                    _ = v19
                                } else {
                                    var v20 uint32 = 42
                                    delete(v1, v20)
                                    _ = v20
                                }
                            }
                            var v21 uint64 = 18446744073709551615
                            v18 += v18
                            _ = v21
                        } else {
                            if v2 {
                                if false {
                                    var v22 int8 = -78
                                    fmt.Println(v2, v17)
                                    _ = v22
                                }
                            } else {
                                v1[924] = v6
                            }
                        }
                        v2 = (!v2)
                        _ = v18
                    }
                    v23 := v5[199]
                    _ = len(v3)
                    _ = v23
                } else {
                    if (65535 >= 4294967295) {
                        fmt.Println(v17)
                        fmt.Println(v17, v17, fi15)
                    }
                }
                if ((-2147483648 <= -2147483648) && (65535 > 2)) {
                    switch -1 {
                    case 127:
                        var v24 uint8 = 14
                        _ = v24
                    }
                }
            }
            _ = fi16
            _ = v17
        }
        if v2 {
            v14 -= (v14 + 42)
        }
        _ = fi15
    }
    fmt.Println(v2, v14, v6)
    var v25 float64 = 0.1
    fmt.Println(v14)
    var v26 string = "a"
    if v2 {
        var v27 float32 = 1e10
        var v28 uint16 = 32767
        _ = v27
        _ = v28
    } else {
        v14 = v14
        var v29 uint8 = 0
        _ = v29
    }
    v30 := GenS1{F0: 2147483648, F1: 255, F2: 511}
    v26 = "a"
    var v31 uint64 = (v14 + 4294967295)
    v32 := v1[uint32(v14)]
    v26 = ((((v6 + v26) + "a") + v6) + v26)
    v6 = v26
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
