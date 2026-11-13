package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 uint16 = 2
    v4 := make(map[int]uint32)
    if v1 {
        if (!("" == "a")) {
            if v1 {
                var v5 uint8 = 249
                v3 += (v3 + v3)
                v3 -= (v3 & v3)
                _ = v5
            } else {
                if (1 == 88) {
                    var v6 float64 = ((1e10 * float64((-2.25 - ((float32((((0.0 * -2.25) - 0.0) - ((-2.25 + 0.1) / 1.5))) * 1e10) / 0.1)))) - (3.14159 + 1e10))
                    _ = v6
                } else {
                    var v7 float32 = -2.25
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v3 = 42
                        fmt.Println(v3, v7, v1)
                        var v9 uint = 42
                        _ = fi8
                        _ = v9
                    }
                    _ = len(v2)
                    _ = v7
                }
                fmt.Println(v1, v3)
                if v1 {
                    v3 += v3
                } else {
                    v3 -= (v3 & v3)
                    var v10 bool = true
                    v3 = v3
                    _ = v10
                }
            }
            var v11 float32 = 1e10
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{fi12, 0, 128, fi12}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        v14 := GenS1{F0: "\n\t", F1: ""}
        _ = v14
    } else {
        var v15 int32 = -500
        v16 := GenS1{F0: "\n\t", F1: "hello"}
        var v17 int = -55
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 bool = v1
    var v19 int32 = -2147483648
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint8 = 128
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        if v1 {
            var v23 float64 = (3.14159 + float64(1.5))
            if v1 {
                v24 := GenS2{F0: -407, F1: (^(v3)), F2: (v3 - (v3 | (^(v3))))}
                v25 := GenS2{F0: 32767, F1: v3, F2: (v3 & v3)}
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                if (0 == 100) {
                    fmt.Println(v18, v3, v3)
                }
                _ = fi26
            }
            _ = v23
        } else {
            for fi27 := 0; fi27 < 12; fi27++ {
                if (235 >= 2) {
                    fmt.Println(v1)
                } else {
                    if true {
                        v18 = true
                        var v28 float32 = 0.0
                        var v29 float32 = -2.25
                        _ = v28
                        _ = v29
                    } else {
                        v3 = 32767
                        var v30 uint32 = 488
                        v19 = (v19 / ((v19) | 1))
                        _ = v30
                    }
                }
                _ = fi27
            }
            v31 := GenS1{F0: ("a" + (("hello" + ("a" + ("foo bar" + (("a" + "hello") + "hello")))) + ("hello" + "\n\t"))), F1: ("hello" + ("foo bar" + ("" + "\n\t")))}
            v19 = v19
            _ = v31
        }
        _ = fi22
    }
    v18 = v1
    v18 = false
    for fi32 := 0; fi32 < 12; fi32++ {
        v19 += -1
        v4[fi32] = 4294967295
        _ = fi32
    }
    fmt.Println(v3, v1, v3)
    fmt.Println(v1, v1)
    v3 = v3
    v33 := []int{-544}
    for fi34 := 0; fi34 < 9; fi34++ {
        v35 := GenS2{F0: -32768, F1: v3, F2: 689}
        _ = fi34
        _ = v35
    }
    _ = len(v4)
    var v36 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
}
