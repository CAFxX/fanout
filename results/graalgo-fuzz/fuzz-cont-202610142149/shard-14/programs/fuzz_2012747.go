package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []int{32768, -1, 0}
    var v2 float32 = (float32((0.1 - 0.1)) / 3.14159)
    var v3 int32 = -1
    var v4 uint64 = 888
    v5 := make(map[int]int)
    var v6 int = 4294967295
    for fi7 := 0; fi7 < 19; fi7++ {
        if (90 <= -503) {
            v6 += (fi7 ^ 4294967295)
            _ = len(v1)
            var v8 int = (fi7 ^ 982)
            _ = v8
        } else {
            v6 = -605
            var v9 int8 = 0
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        switch fi10 {
        case 127:
            v11 := v5[(fi10 * (fi10 % ((int(v4)) | 1)))]
            _ = v11
        case 32768:
            if true {
                v12 := GenS1{F0: 2}
                switch 9223372036854775807 {
                case 9223372036854775807:
                    for fi13 := 0; fi13 < 6; fi13++ {
                        var v14 string = ""
                        var v15 uint32 = 65535
                        if (1 != -128) {
                            if (100 <= 1) {
                                if (v3 <= 1) {
                                    var v16 float32 = 1.5
                                    if (v2 >= v16) {
                                        var v17 string = v14
                                        v18 := GenS1{F0: 781}
                                        _ = v17
                                        _ = v18
                                    } else {
                                        fmt.Println(v6, v4)
                                        v2 = (v16 / 1.5)
                                        v16 = 0.0
                                    }
                                    _ = v16
                                } else {
                                    var v19 uint = 1
                                    v4 = (v4 / ((uint64(v15)) | 1))
                                    if (v3 != (v3 >> (((^(v19))) & 63))) {
                                        v2 += v2
                                    }
                                    _ = v19
                                }
                                fmt.Println(fi13)
                            }
                            var v20 int8 = -124
                            _ = v20
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    for fi21 := 0; fi21 < 6; fi21++ {
                        fmt.Println(v6, fi10, fi21)
                        var v22 uint16 = 32767
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v12
            }
        }
        v4 = v4
        var v23 int32 = v3
        _ = fi10
        _ = v23
    }
    fmt.Println(v4)
    _ = len(v5)
    fmt.Println(v3)
    var v24 uint = 0
    fmt.Println(v3, v6, v6)
    var v25 uint = (v24 | v24)
    switch (v6 >> ((v24) & 63)) {
    case 1:
        if false {
            var v26 int64 = -9223372036854775808
            v24 = (^(v25))
            var v27 float64 = 1.5
            _ = v26
            _ = v27
        } else {
            fmt.Println(v24, v2, v2)
            var v28 float32 = 0.1
            _ = v28
        }
    case 9223372036854775807:
        fmt.Println(v25)
    case -433:
        if false {
            _ = len(v1)
        }
    default:
        delete(v5, v6)
    }
    v4 += 927
    var v29 bool = false
    fmt.Println(v3, v2)
    if ("a" < "") {
        var v30 float64 = -2.25
        var v31 int32 = 42
        v32 := GenS1{F0: 0}
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
