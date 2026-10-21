package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{0, 32768, 32768, 32767}
    var v2 uint = 255
    v3 := make(map[int]string)
    var v4 string = "foo bar"
    v5 := make(map[int]int)
    v3[-9223372036854775808] = "a"
    v2 -= 554
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        switch 369 {
        case 128:
            for fi6 := 0; fi6 < 15; fi6++ {
                if true {
                    var v7 bool = false
                    v8 := GenS1{F0: 0, F1: -857, F2: 2}
                    v9 := GenS2{F0: 0, F1: v7}
                    _ = v7
                    _ = v8
                    _ = v9
                } else {
                    if (!true) {
                        fmt.Println(v4, v2)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                var v10 int32 = 42
                _ = fi6
                _ = v10
            }
        }
        fmt.Println(v4)
    }
    switch -2147483648 {
    case -2147483648:
        v2 = (v2 * 0)
    case 127:
        for fi11 := 0; fi11 < 9; fi11++ {
            if (!false) {
                if (1 == 0) {
                    v12 := GenS1{F0: 167, F1: -715, F2: 4294967295}
                    v2 += (^(v2))
                    for fi13 := 0; fi13 < 3; fi13++ {
                        switch (-(fi13)) {
                        case 916:
                            delete(v5, -166)
                        case 32767:
                            var v14 uint32 = 32
                            _ = v14
                        default:
                            v15 := GenS1{F0: 255, F1: 2147483647, F2: 2147483647}
                            _ = v15
                        }
                        v2 = v2
                        if (255 >= 14) {
                            if true {
                                switch (fi13 | (fi13 << ((v2) & 63))) {
                                case 32768:
                                    var v16 int = 9223372036854775807
                                    _ = v16
                                case -9223372036854775808:
                                    v2 -= v2
                                    var v17 int8 = 100
                                    _ = v17
                                case 4294967295:
                                    fmt.Println(v4, fi13)
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v18 bool = (true && (false && false))
                                switch (fi11 * fi13) {
                                case 127:
                                    v1 = append(v1, fi13)
                                }
                                _ = len(v1)
                                _ = v18
                            }
                            var v19 int32 = 0
                            _ = v19
                        }
                        _ = fi13
                    }
                    _ = v12
                } else {
                    var v20 uint16 = 1
                    _ = v20
                }
            }
            _ = fi11
        }
    }
    delete(v5, 381)
    v21 := GenS2{F0: -9223372036854775808, F1: (!(42 == 0))}
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint16 = 914
            v4 = (("" + "foo bar") + v4)
            _ = fi23
            _ = v24
        }
        if (!(false && false)) {
            v1 = append(v1, fi22)
            var v25 uint32 = 693
            v2 = (v2 & v2)
            _ = v25
        } else {
            if (3.14159 == (0.1 / 1.0)) {
                v2 += 310
            }
            fmt.Println(fi22, v2)
        }
        v2 = (^(v2))
        _ = fi22
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int16 = 0
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 uint8 = 2
            _ = len(v5)
            _ = fi28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    _ = len(v1)
    fmt.Println(v2)
    fmt.Println(v2, v2)
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 int32 = 1
        _ = fi30
        _ = v31
    }
    var v32 uint32 = 770
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
