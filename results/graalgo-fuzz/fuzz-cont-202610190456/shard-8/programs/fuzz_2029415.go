package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint64; F1 int }
func genFn0() int {
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int8 = 1
        v9 := []int{fi7, fi7}
        for fi10 := 0; fi10 < 18; fi10++ {
            if len(v9) > 0 {
                v9[0] = 4294967295
            }
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v11 := []int{42, -345, -1, 32768}
    v6 = 65535
    v12 := GenS1{F0: 0, F1: 0}
    _ = v11
    _ = v12
    _ = v6
    return 490
}

func main() {
    var v1 float64 = 0.0
    var v2 int8 = 1
    var v3 string = "\n\t"
    var v4 float32 = ((float32((v1 + v1)) / 1.0) * 3.14159)
    v5 := []uint32{0, 65535, 911}
    var v13 string = v3
    fmt.Println(v3, v2, v1)
    v4 = float32(float64((v4 * 1e10)))
    var v14 uint16 = 2
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: 1, F1: (("" + "\n\t") + "foo bar")}
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS3{F0: 952, F1: 32767}
            var v19 uint64 = 4294967295
            if (0 > 2) {
                switch fi15 {
                case 128:
                    if ((v19 + v19) <= 943) {
                        if true {
                            var v20 int16 = 283
                            switch genFn0() {
                            case 0:
                                var v21 bool = true
                                _ = v21
                            case 269:
                                if (0.1 > (v4 + v4)) {
                                    v5 = append(v5, 30)
                                } else {
                                    v14 = 42
                                    if ((!false) && true) {
                                        v22 := GenS3{F0: (v19 ^ v19), F1: -906}
                                        var v23 uint64 = (v19 | (v19 % (((v19 / ((v19) | 1))) | 1)))
                                        var v24 uint = 42
                                        v19 = v23
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    }
                                }
                            default:
                                fmt.Println(v1, v4, v2)
                                var v25 uint8 = 2
                                _ = v25
                            }
                            fmt.Println(v13)
                            _ = v20
                        }
                        if len(v5) > 0 {
                            v5[0] = 2
                        }
                        v14 += v14
                    } else {
                        var v26 uint32 = 744
                        var v27 bool = false
                        v4 -= 1.5
                        _ = v26
                        _ = v27
                    }
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        var v28 uint32 = 42
        _ = fi15
        _ = v16
        _ = v28
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint = 2
        _ = fi29
        _ = v30
    }
    var v31 uint32 = 4294967295
    v32 := GenS3{F0: 42, F1: 32767}
    v31 -= v31
    switch genFn0() {
    case 42:
        v3 = ("hello" + v3)
    case 4294967295:
        if len(v5) > 0 {
            v5[0] = 0
        }
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
