package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 uint = 435
    var v5 uint16 = 32767
    var v6 int64 = 42
    if false {
        if (925 >= 32768) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v1 += (v1 % ((-703) | 1))
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 int32 = 0
                    var v10 uint64 = 42
                    v5 = v5
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                switch ("a" + "a") {
                case "":
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if false {
                            fmt.Println(v5)
                            var v12 int32 = 1
                            v13 := v3[(^(fi7))]
                            _ = v12
                            _ = v13
                        } else {
                            v1 -= int16(v4)
                            var v14 int8 = -2
                            _ = v14
                        }
                        _ = fi11
                    }
                case "hello":
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS1{F0: -718, F1: 77}
                        _ = fi15
                        _ = v16
                    }
                case "\n\t":
                    var v17 float64 = 3.14159
                    _ = v17
                default:
                    if (42 <= 1) {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            v4 = (v4 - v4)
                            _ = fi18
                        }
                        v19 := []int{(-(fi7)), (^(fi7))}
                        _ = v19
                    }
                }
                _ = fi7
            }
            v20 := []int{}
            _ = v20
        } else {
            var v21 float64 = 1.5
            _ = v21
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            v6 = (v6 / ((-1) | 1))
            fmt.Println(v6, fi22)
            _ = fi22
        }
        var v23 int16 = -32768
        _ = v23
    } else {
        fmt.Println(v5, v4, v6)
        fmt.Println(v6)
    }
    if (390 != 42) {
        if ((0.0 - 0.1) > 1.5) {
            v1 = v1
        } else {
            fmt.Println(v6, v1, v1)
        }
        fmt.Println(v4)
    }
    var v24 float64 = (1.5 - float64(float32(0.0)))
    if true {
        var v25 uint8 = 42
        for fi26 := 0; fi26 < 18; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                _ = len(v3)
                if (true && true) {
                    var v28 int = fi26
                    _ = v28
                } else {
                    v29 := GenS2{F0: (fi27 << (((v4 & v4)) & 63))}
                    _ = v29
                }
                v24 = 1e10
                _ = fi27
            }
            _ = fi26
        }
        _ = v25
    } else {
        for fi30 := 0; fi30 < 17; fi30++ {
            v31 := v2[42]
            _ = fi30
            _ = v31
        }
    }
    if true {
        var v32 int32 = -1
        _ = v32
    } else {
        var v33 int = -1
        for fi34 := 0; fi34 < 7; fi34++ {
            fmt.Println(v33, v33, fi34)
            v6 = 1
            v35 := GenS1{F0: -2147483648, F1: 2}
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    var v36 float64 = float64(3.14159)
    var v37 int8 = int8(v4)
    var v38 uint = (v4 ^ 534)
    v24 -= (v36 * (0.1 * (((float64((3.14159 - ((1.5 * 3.14159) * (0.0 / 1.0)))) - (v36 / 1e10)) - 0.1) / 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
