package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }

func main() {
    var v1 int16 = 870
    v2 := []int{-2147483648}
    var v3 uint = 255
    v4 := []int{}
    v5 := make(map[uint32]int)
    var v6 float32 = 0.1
    v7 := GenS1{F0: 860}
    var v8 uint32 = 2
    v9 := v5[v8]
    var v10 float64 = (float64((v6 * (v6 * v6))) - ((((3.14159 / ((((0.1 * 0.1) * -2.25) / -2.25) * -2.25)) * ((-2.25 - float64(v6)) * (((1e10 / 1.0) + 3.14159) / 1e10))) + ((3.14159 + 0.0) / ((1.5 + -2.25) + 1e10))) / float64(v6)))
    var v11 int = v9
    if true {
        v11 -= v11
        var v12 int32 = 1
        _ = v12
    } else {
        fmt.Println(v1)
    }
    v5[v8] = (^(v9))
    fmt.Println(v9)
    if (false && true) {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS1{F0: 1}
                if true {
                    if false {
                        v1 += v1
                        fmt.Println(v8)
                    }
                }
                fmt.Println(v10)
                _ = fi14
                _ = v15
            }
            v16 := GenS2{F0: 1}
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    v1 -= 947
                }
            case "":
                v17 := GenS3{F0: -1, F1: "a", F2: 2}
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    if false {
                        v1 += v1
                        fmt.Println(v3, v1, v8)
                        var v19 bool = (int16(fi13) <= v1)
                        _ = v19
                    }
                    if true {
                        if len(v2) > 0 {
                            v2[0] = fi18
                        }
                    }
                    var v20 uint = (v3 >> (((v3 * (v3 - (v3 | (^(v3)))))) & 63))
                    fmt.Println(fi18, v8)
                    v10 += -2.25
                    _ = fi18
                    _ = v20
                }
            }
            _ = fi13
            _ = v16
        }
        fmt.Println(v3, v6)
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            _ = len(v5)
            var v22 bool = true
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            if (!false) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    v10 = float64(((((1e10 / v6) / v6) + (v6 * v6)) * float32(float64((v6 / v6)))))
                    var v25 uint8 = 42
                    var v26 float32 = (0.0 + v6)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 int8 = 42
                switch fi23 {
                case 127:
                    fmt.Println(v8, v9, v10)
                case 1:
                    v28 := GenS1{F0: 32767}
                    _ = v28
                case -9223372036854775808:
                    if true {
                        for fi29 := 0; fi29 < 16; fi29++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi29
                        }
                        v27 -= v27
                    }
                default:
                    if false {
                        v30 := GenS4{F0: 65535, F1: false, F2: v3}
                        _ = v30
                    } else {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v3, v1)
                        }
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v27
            }
            v31 := GenS1{F0: uint16(v3)}
            _ = fi23
            _ = v31
        }
        _ = len(v5)
    }
    v1 = v1
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint32 = v8
        v1 -= (^(v1))
        v33 -= v33
        _ = fi32
        _ = v33
    }
    fmt.Println(v8)
    var v34 uint = v3
    var v35 uint = 2
    v4 = append(v4, (v9 / (((v9 ^ int(v34))) | 1)))
    var v36 int16 = v1
    for fi37 := 0; fi37 < 20; fi37++ {
        v36 = (v1 * 32767)
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
