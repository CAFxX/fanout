package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32 }
func genFn0() uint64 {
    var v4 uint = 42
    if (!true) {
        var v5 uint16 = 0
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6, v4)
        v7 := []int{(fi6 & 4294967295), (fi6 ^ 2147483647)}
        _ = fi6
        _ = v7
    }
    v4 = (^(v4))
    v8 := []int{408}
    _ = v4
    _ = v8
    return uint64(v4)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 int8 = -60
    var v9 uint64 = 1
    v9 += (v9 + 4294967295)
    if (!true) {
        v10 := []int{-2147483648, -1, 2147483647}
        for fi11 := 0; fi11 < 15; fi11++ {
            if false {
                v9 = (v9 - (^(v9)))
                v12 := v2[(fi11 & fi11)]
                _ = v12
            } else {
                fmt.Println(fi11, fi11, v9)
            }
            var v13 uint8 = 127
            _ = fi11
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v9 -= genFn0()
                if true {
                    if false {
                        fmt.Println(v3)
                    } else {
                        v16 := GenS1{F0: false, F1: -1}
                        v3 += (-(v3))
                        if (800 >= 0) {
                            v17 := GenS1{F0: false, F1: 42}
                            v9 = (v9 % ((v9) | 1))
                            var v18 int16 = 42
                            v3 = v3
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    }
                } else {
                    if true {
                        var v19 float64 = 0.0
                        if (1 >= -539) {
                            if (("" >= ("" + "")) && false) {
                                var v20 float32 = 0.1
                                _ = v20
                            }
                            v19 -= v19
                            v3 = v3
                        }
                        var v21 uint32 = 1
                        _ = v19
                        _ = v21
                    } else {
                        v3 = v3
                    }
                }
                switch "\n\t" {
                case "hello":
                    v3 += 1
                case "\n\t":
                    var v22 float32 = float32(-2.25)
                    _ = v22
                case "a":
                    var v23 int16 = -383
                    _ = v23
                default:
                    fmt.Println(fi15, v3, v3)
                }
                _ = fi15
            }
            var v24 int = fi14
            v25 := GenS1{F0: true, F1: -32768}
            _ = fi14
            _ = v24
            _ = v25
        }
        _ = v10
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        if true {
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 float32 = 1.5
                v28 += -2.25
                v29 := GenS2{F0: 1, F1: genFn0()}
                v30 := GenS2{F0: 267, F1: (v9 & 42)}
                fmt.Println(v9, fi27, v3)
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        _ = fi26
    }
    var v31 int8 = -1
    var v32 uint8 = uint8(v9)
    fmt.Println(v32)
    var v33 bool = true
    var v34 string = ""
    v35 := GenS2{F0: 2, F1: v9}
    v2[-1] = -387
    v36 := GenS1{F0: v33, F1: -32768}
    v37 := v2[523]
    var v38 int8 = (^(v3))
    v39 := GenS3{F0: 784}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v9
}
