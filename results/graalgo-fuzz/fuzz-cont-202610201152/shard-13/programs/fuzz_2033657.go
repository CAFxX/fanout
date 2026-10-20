package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint16; F2 bool }

func main() {
    var v1 int8 = -1
    var v2 float64 = 1e10
    var v3 uint16 = 42
    var v4 uint8 = 0
    var v5 uint = 65535
    var v6 int64 = 42
    if false {
        v7 := []int{206, -2147483648}
        v3 = v3
        _ = v7
    } else {
        if (127 < -1) {
            for fi8 := 0; fi8 < 8; fi8++ {
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint32 = 4294967295
                    fmt.Println(v10, v6, v6)
                    v2 -= (v2 * v2)
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            switch -1 {
            case 1:
                var v11 int8 = v1
                _ = v11
            case 389:
                if ((1 == -1) && (false || false)) {
                    fmt.Println(v6, v3)
                    if (!false) {
                        for fi12 := 0; fi12 < 2; fi12++ {
                            if false {
                                if true {
                                    v13 := []int{684, (fi12 << ((v5) & 63)), fi12}
                                    v6 += (v6 | 1)
                                    if len(v13) > 0 {
                                        v13[0] = (fi12 + (fi12 + (fi12 + (fi12 ^ fi12))))
                                    }
                                    _ = v13
                                }
                                v14 := GenS1{F0: 895, F1: 0}
                                var v15 string = "hello"
                                _ = v14
                                _ = v15
                            } else {
                                fmt.Println(v1)
                            }
                            _ = fi12
                        }
                        var v16 uint16 = v3
                        _ = v16
                    }
                }
            }
        }
        v17 := GenS2{F0: ((1.5 * 0.1) >= float32((v2 - 0.1))), F1: uint16(v4)}
        switch "" {
        case "":
            v6 -= v6
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                v5 = 42
                v19 := GenS1{F0: -278, F1: -1}
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    }
    fmt.Println(v2)
    if (17 != 0) {
        v20 := GenS1{F0: 0, F1: 1000}
        _ = v20
    }
    var v21 string = ""
    v22 := []int{414}
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS3{F0: (fi23 - (fi23 - (fi23 + (^(fi23)))))}
        _ = fi23
        _ = v24
    }
    var v25 uint16 = 32767
    v26 := GenS4{F0: v21, F1: (v25 ^ 216), F2: ((0 != 2147483647) || (v5 >= (v5 | 265)))}
    var v27 string = (("\n\t" + (v21 + ("\n\t" + v21))) + v21)
    var v28 int32 = 42
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint16 = (v3 / ((32767) | 1))
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 886
    v21 = "a"
    var v32 int = 42
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
