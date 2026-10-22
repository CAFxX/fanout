package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 string; F1 bool; F2 int64 }
func genFn0(p0 uint) string {
    v5 := []int{42, 127}
    v6 := GenS1{F0: 42, F1: -241}
    if (0 != 0) {
        var v7 uint16 = 67
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                v7 += v7
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                switch fi8 {
                case 32768:
                    var v9 int64 = 1
                    _ = v9
                case -573:
                    fmt.Println(fi8)
                default:
                    switch ("a" + ("foo bar" + "")) {
                    case "":
                        if false {
                            var v10 int16 = 0
                            _ = v10
                        } else {
                            var v11 int8 = -128
                            _ = v11
                        }
                    default:
                        var v12 uint64 = 1
                        _ = v12
                    }
                }
                var v13 uint8 = 25
                switch fi8 {
                case 1:
                    fmt.Println(v7)
                case 2147483647:
                    var v14 int32 = 2147483647
                    _ = v14
                case 336:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v16 := GenS1{F0: 42, F1: -9223372036854775808}
                        v17 := GenS2{F0: fi15, F1: -126}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                _ = v13
            }
            _ = fi8
        }
        _ = v7
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(p0, fi18)
        _ = fi18
    }
    _ = len(v5)
    _ = p0
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = 3.14159
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    var v19 float32 = 0.0
    v3[32768] = uint32(v1)
    if true {
        v20 := GenS3{F0: 100}
        v21 := GenS4{F0: 176, F1: 854}
        _ = v20
        _ = v21
    } else {
        fmt.Println(v1, v19, v2)
    }
    v19 = (v19 + (1e10 - (((((v19 - v19) + v4) / ((float32((v2 / 1.0)) + v4) + float32(0.0))) + v19) / v19)))
    v22 := GenS5{F0: 342, F1: -311}
    if false {
        var v23 bool = (42 < -9223372036854775808)
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS4{F0: 4294967295, F1: 255}
            if (!true) {
                var v26 uint32 = 2
                var v27 float64 = v2
                _ = len(v3)
                _ = v26
                _ = v27
            } else {
                v1 = (^(v1))
                v23 = (!((!v23) || true))
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    var v28 int32 = 1
    var v29 int16 = 32767
    if false {
        for fi30 := 0; fi30 < 11; fi30++ {
            if false {
                v28 = (v28 & (v28 * (v28 << ((0) & 63))))
                v31 := GenS6{F0: ("\n\t" + ""), F1: ((v29 << ((42) & 63)) != (v29 - v29)), F2: -986}
                _ = v31
            }
            _ = fi30
        }
        v32 := GenS5{F0: -799, F1: 0}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
