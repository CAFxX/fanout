package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -2147483648
    var v3 bool = ((v2 | v2) != v2)
    v4 := make(map[int]bool)
    var v5 float32 = ((((float32((0.0 + (((float64((float32((((3.14159 / 1e10) / (((0.0 * (0.0 / 1.0)) - 1e10) / (1e10 + -2.25))) * -2.25)) * 1.5)) - (1.5 - (-2.25 / ((0.1 / -2.25) / 1.5)))) / 1.5) - float64(3.14159)))) + 1e10) + 0.1) + (float32(float64(-2.25)) * ((0.1 - 1e10) + 3.14159))) / 1.0)
    var v6 uint8 = 56
    v2 = 1
    if (1 > 0) {
        _ = len(v1)
        var v7 int = 317
        fmt.Println(v5, v5, v5)
        _ = v7
    } else {
        var v8 uint64 = 42
        _ = v8
    }
    v9 := GenS1{F0: 651}
    v3 = v3
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 float32 = (((v5 - v5) / v5) + v5)
        _ = fi10
        _ = v11
    }
    var v12 int64 = 505
    for fi13 := 0; fi13 < 3; fi13++ {
        if true {
            var v14 uint32 = 2147483647
            var v15 int = -2147483648
            var v16 float64 = (-2.25 + ((-2.25 + (1.5 / (float64(v5) + float64((v5 / v5))))) / (0.1 * (-2.25 + (1.5 * (3.14159 + 0.1))))))
            _ = v14
            _ = v15
            _ = v16
        } else {
            v6 += (^(v6))
            if ((v3 || ((!(65535 >= 42)) && ((^(fi13)) <= fi13))) || v3) {
                var v17 uint16 = 949
                switch "a" {
                case "hello":
                    var v18 int = 128
                    _ = v18
                case "a":
                    if ((!false) && false) {
                        if true {
                            var v19 uint32 = 65535
                            v2 = 42
                            _ = v19
                        }
                        var v20 uint16 = v17
                        var v21 int8 = 83
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int16 = 0
                        if false {
                            var v23 uint16 = v17
                            _ = v23
                        }
                        v24 := GenS1{F0: -1}
                        _ = v22
                        _ = v24
                    }
                case "":
                    if (!("hello" > "foo bar")) {
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v26 := GenS1{F0: 42}
                            v6 -= v6
                            switch "hello" {
                            case "hello":
                                v27 := []int{784, (fi25 ^ (fi25 % (((fi25 | (fi13 ^ (^(fi13))))) | 1))), fi13}
                                _ = v27
                            case "a":
                                var v28 float32 = (0.1 - (v5 * v5))
                                _ = v28
                            default:
                                var v29 uint = 4294967295
                                _ = v29
                            }
                            if (!(true && v3)) {
                                v30 := GenS2{F0: v17, F1: -9223372036854775808, F2: 2}
                                v6 -= 60
                                v31 := []int{fi13, -578, 128, (fi25 | (fi25 << ((0) & 63)))}
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS1{F0: -480}
                            v12 = (v12 << ((953) & 63))
                            _ = fi25
                            _ = v26
                            _ = v32
                        }
                        v1[2147483647] = v3
                        v33 := GenS1{F0: 0}
                        _ = v33
                    }
                default:
                    v3 = false
                }
                _ = v17
            } else {
                v3 = (v3 && ((uint32(v6) == 42) && (!(false || (-127 == 0)))))
            }
            fmt.Println(v6)
        }
        var v34 int64 = v12
        fmt.Println(v6)
        _ = fi13
        _ = v34
    }
    fmt.Println(v2)
    var v35 uint8 = v6
    v12 += v12
    v36 := GenS3{F0: v3, F1: v3, F2: 0}
    fmt.Println(v6, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
