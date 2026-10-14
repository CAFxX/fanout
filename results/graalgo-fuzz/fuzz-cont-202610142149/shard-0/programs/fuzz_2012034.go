package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{}
    var v3 int8 = 1
    v4 := []int{}
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int32 = 2147483647
        var v8 int = 840
        if len(v4) > 0 {
            v4[0] = fi6
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if (2 != 0) {
        if ((false || (float32((0.0 - 0.1)) < 0.0)) || (!(255 != 2))) {
            var v9 uint = (v1 + 2)
            _ = v9
        }
    } else {
        v1 -= 4294967295
        switch "foo bar" {
        case "":
            v3 -= v3
        default:
            if true {
                v1 += (^(v1))
            }
        }
    }
    var v10 int16 = 1000
    fmt.Println(v1)
    if false {
        var v11 int16 = v10
        if (2 < 860) {
            var v12 uint64 = 42
            if ((v1 / ((v1) | 1)) >= v1) {
                if false {
                    if false {
                        var v13 float64 = 0.1
                        if false {
                            fmt.Println(v11, v13, v3)
                            switch 853 {
                            case 4294967295:
                                v14 := GenS1{F0: 868, F1: 235, F2: -1}
                                _ = v14
                            }
                        } else {
                            var v15 int16 = v11
                            var v16 float32 = (float32(float64((0.0 - 0.0))) / 1.0)
                            if (!(4294967295 <= 2)) {
                                if (true || ((!true) && (!false))) {
                                    v3 += v3
                                    v11 += (v15 * v10)
                                    var v17 uint64 = (v12 - (v12 / ((v12) | 1)))
                                    _ = v17
                                }
                                if (v16 <= v16) {
                                    var v18 int64 = -1
                                    _ = v18
                                }
                            } else {
                                v19 := GenS1{F0: 0, F1: 1, F2: -2147483648}
                                _ = v19
                            }
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    }
                }
                fmt.Println(v12, v3)
                var v20 string = ("" + "foo bar")
                _ = v20
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v11 += 1000
                    var v22 uint = 42
                    _ = len(v5)
                    _ = fi21
                    _ = v22
                }
                if false {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if true {
                            v24 := GenS2{F0: 65535}
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 15; fi25++ {
                            delete(v5, 1)
                            _ = fi25
                        }
                        _ = fi23
                    }
                }
                var v26 uint8 = 1
                _ = v26
            }
            _ = v12
        } else {
            _ = len(v2)
        }
        fmt.Println(v1)
        _ = v11
    }
    switch 1 {
    case 418:
        v27 := GenS1{F0: 4294967295, F1: 406, F2: int32(v3)}
        _ = v27
    case -9223372036854775808:
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint16 = 65535
            fmt.Println(v1)
            _ = fi28
            _ = v29
        }
    case 986:
        v1 += v1
    }
    v30 := GenS1{F0: 2147483648, F1: 255, F2: -2147483648}
    fmt.Println(v10, v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
