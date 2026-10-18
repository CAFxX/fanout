package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 int8 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 406
    v4 := make(map[int]bool)
    var v5 uint = 451
    for fi6 := 0; fi6 < 8; fi6++ {
        delete(v2, v3)
        v7 := GenS1{F0: (v3 / (((v3 << ((v5) & 63))) | 1)), F1: v1}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: v3, F1: v1}
    v3 = v3
    fmt.Println(v1, v3, v5)
    v9 := []int{}
    if (!(false || true)) {
        var v10 float64 = 3.14159
        _ = v10
    }
    if (v3 != 0) {
        v1 -= (v1 * 1)
        for fi11 := 0; fi11 < 5; fi11++ {
            _ = len(v9)
            v1 += int16(v3)
            fmt.Println(fi11, v5)
            _ = fi11
        }
        v12 := GenS2{F0: 128, F1: 32767, F2: 764}
        _ = v12
    }
    v1 = (v1 << (((^(v5))) & 63))
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v13 := GenS2{F0: 0, F1: 32767, F2: v1}
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS3{F0: (-(fi14)), F1: 42}
            if (false && true) {
                v17 := GenS4{F0: 2}
                if true {
                    v5 = (v5 & 0)
                } else {
                    if (1.5 < 0.0) {
                        fmt.Println(v3, v1, fi15)
                        v3 = (v3 * (^(v3)))
                        v5 += (v5 ^ (v5 << (((^(v5))) & 63)))
                    } else {
                        var v18 int8 = 42
                        _ = v18
                    }
                    if (float64((float32(1.5) + (-2.25 + (0.0 / (0.0 + (0.1 + 0.0)))))) >= float64(((3.14159 * -2.25) + (0.1 * (((3.14159 - 0.0) - -2.25) + 1.5))))) {
                        fmt.Println(fi14)
                        var v19 float64 = 1e10
                        _ = v19
                    } else {
                        var v20 int32 = -1
                        _ = v20
                    }
                    v3 -= 934
                }
                _ = v17
            }
            v1 = (v1 - (v1 - v1))
            v21 := GenS1{F0: 405, F1: v1}
            _ = fi15
            _ = v16
            _ = v21
        }
        if false {
            if true {
                for fi22 := 0; fi22 < 20; fi22++ {
                    if (1e10 < (3.14159 + -2.25)) {
                        var v23 uint8 = 168
                        v24 := GenS4{F0: 731}
                        _ = len(v2)
                        _ = v23
                        _ = v24
                    }
                    switch 127 {
                    case -544:
                        fmt.Println(v3, fi22)
                    case 387:
                        var v25 uint8 = 198
                        switch -533 {
                        case 225:
                            v5 = v5
                        }
                        v25 -= 2
                        _ = v25
                    default:
                        fmt.Println(v5, v1, fi22)
                    }
                    _ = fi22
                }
            } else {
                v26 := GenS5{F0: fi14, F1: -89}
                var v27 int = (fi14 ^ -2147483648)
                v27 = fi14
                _ = v26
                _ = v27
            }
            if (!(true || false)) {
                var v28 int32 = 0
                _ = v28
            } else {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 float32 = -2.25
                    var v31 int8 = 100
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        } else {
            v32 := GenS6{F0: v1, F1: -2147483648, F2: 0}
            _ = v32
        }
        var v33 uint64 = 328
        _ = fi14
        _ = v33
    }
    switch (("\n\t" + ("" + "")) + "hello") {
    case "":
        var v34 uint8 = 127
        _ = v34
    case "hello":
        var v35 bool = false
        _ = v35
    default:
        v1 -= (v1 % ((v1) | 1))
    }
    var v36 uint = (v5 % ((861) | 1))
    v5 = 236
    switch 42 {
    case 0:
        var v37 uint16 = 53
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
