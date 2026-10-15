package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []int{-1}
    var v2 float32 = float32(0.1)
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 1
    v3[1] = uint32(v4)
    v2 -= (v2 + (1.5 + 0.0))
    if len(v1) > 0 {
        v1[0] = 479
    }
    v5 := v3[2147483647]
    v6 := GenS1{F0: 42}
    var v7 int8 = 1
    var v8 int8 = -128
    switch 857 {
    case 321:
        v5 = (^(v5))
    default:
        v9 := GenS1{F0: -9223372036854775808}
        _ = v9
    }
    v4 = (v4 / (((v4 << ((255) & 63))) | 1))
    for fi10 := 0; fi10 < 10; fi10++ {
        v2 -= float32(0.1)
        fmt.Println(fi10, v8, v4)
        _ = fi10
    }
    v7 = (v8 * v7)
    fmt.Println(v4)
    switch 2147483647 {
    case -164:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 9223372036854775807:
        var v11 int32 = 71
        _ = v11
    case -136:
        for fi12 := 0; fi12 < 18; fi12++ {
            switch fi12 {
            case 9223372036854775807:
                for fi13 := 0; fi13 < 20; fi13++ {
                    delete(v3, v5)
                    v14 := GenS1{F0: 0}
                    v5 -= 2147483648
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi12
        }
    default:
        if (122 <= 200) {
            var v15 float64 = (3.14159 * float64(float32((-2.25 - 1e10))))
            v16 := GenS1{F0: 9223372036854775807}
            var v17 bool = false
            _ = v15
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(fi18)
                _ = fi18
            }
            v19 := GenS1{F0: -521}
            switch 944 {
            case 128:
                v20 := GenS1{F0: 1}
                _ = v20
            case 32768:
                if true {
                    var v21 uint16 = 42
                    for fi22 := 0; fi22 < 19; fi22++ {
                        v23 := GenS1{F0: -1}
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v5, v2, v2)
                    _ = v21
                } else {
                    var v24 int16 = -1
                    _ = v24
                }
            default:
                if (false && false) {
                    if (true || ((42 <= 65535) && ((v4 << ((2) & 63)) != 1))) {
                        v25 := GenS2{F0: -925, F1: 980}
                        switch 32767 {
                        case -2147483648:
                            v5 = 0
                        case -9223372036854775808:
                            v4 += 255
                        case -952:
                            switch "\n\t" {
                            case "hello":
                                var v26 bool = false
                                _ = v26
                            case "a":
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    var v28 int16 = 1000
                                    _ = fi27
                                    _ = v28
                                }
                            case "\n\t":
                                v4 += v4
                            default:
                                delete(v3, (v5 ^ (v5 % ((v5) | 1))))
                            }
                        }
                        switch -1 {
                        case -9223372036854775808:
                            switch 828 {
                            case -1:
                                fmt.Println(v4)
                            case 42:
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    if false {
                                        fmt.Println(v2, v4, v4)
                                        var v30 uint8 = 42
                                        fmt.Println(fi29, v4)
                                        _ = v30
                                    } else {
                                        fmt.Println(v5, v8, v7)
                                    }
                                    _ = fi29
                                }
                            case 2147483647:
                                var v31 int32 = 2147483647
                                _ = v31
                            }
                        }
                        _ = v25
                    }
                    var v32 string = "a"
                    _ = v32
                } else {
                    switch 32767 {
                    case 4:
                        var v33 float64 = (float64(((float32(float64(v2)) / 1.5) - 3.14159)) * 3.14159)
                        _ = v33
                    case 128:
                        delete(v3, (v5 % ((0) | 1)))
                    default:
                        v7 += v7
                    }
                    for fi34 := 0; fi34 < 7; fi34++ {
                        if true {
                            var v35 uint64 = 396
                            _ = v35
                        }
                        v8 = (v8 + int8(v4))
                        _ = fi34
                    }
                }
            }
            _ = v19
        }
    }
    v5 += (v5 + 0)
    v4 = v4
    var v36 int8 = (v7 ^ v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
