package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 uint32 }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]uint32)
    var v3 int64 = -104
    v4 := GenS1{F0: v3, F1: 909, F2: 127}
    v1 += 2
    v1 += v1
    var v5 float32 = 0.1
    var v6 float32 = float32(1.5)
    v1 = v1
    var v7 uint8 = 42
    v8 := GenS1{F0: v3, F1: 712, F2: v7}
    v9 := GenS1{F0: v3, F1: 2, F2: v7}
    var v10 float32 = 0.0
    v3 -= 486
    if ((0.0 * 0.0) <= float64((((3.14159 - 1e10) - -2.25) / v5))) {
        v11 := []int{}
        if false {
            for fi12 := 0; fi12 < 7; fi12++ {
                switch fi12 {
                case 32768:
                    v7 += (v7 ^ (^(v7)))
                default:
                    v6 -= (v10 + ((0.0 - float32(1e10)) / 1.5))
                }
                _ = fi12
            }
            var v13 int16 = 42
            _ = v13
        } else {
            var v14 float32 = (v5 - float32(1e10))
            _ = v14
        }
        v3 -= v3
        _ = v11
    } else {
        v5 -= (v5 + v10)
        fmt.Println(v7, v10, v5)
    }
    if (v3 <= v3) {
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                v1 += v1
                if (v7 > v7) {
                    if (!true) {
                        v3 = (v3 + -309)
                        var v17 uint16 = (v1 << ((42) & 63))
                        _ = v17
                    }
                    v5 = v10
                    switch fi16 {
                    case -1:
                        switch ("a" + "hello") {
                        case "hello":
                            switch fi16 {
                            case 32768:
                                v3 = 42
                            default:
                                switch 127 {
                                case 1:
                                    switch (-(fi16)) {
                                    case 640:
                                        switch (("foo bar" + ("a" + "foo bar")) + "a") {
                                        case "":
                                            v18 := GenS1{F0: v3, F1: 847, F2: 255}
                                            _ = v18
                                        default:
                                            v6 -= v6
                                        }
                                    case 524:
                                        v19 := []int{fi16, (fi16 ^ (fi15 << ((489) & 63)))}
                                        _ = v19
                                    default:
                                        fmt.Println(v7, fi15)
                                    }
                                case 0:
                                    fmt.Println(v1, v5)
                                }
                            }
                        case "a":
                            var v20 float64 = (0.0 * (1e10 + 0.1))
                            _ = v20
                        default:
                            v1 = v1
                        }
                        var v21 int64 = v3
                        _ = v21
                    case 127:
                        var v22 int = fi15
                        _ = v22
                    case 87:
                        var v23 float32 = (1e10 / (v6 - v6))
                        _ = v23
                    }
                } else {
                    var v24 uint8 = v7
                    _ = v24
                }
                v25 := GenS2{F0: 0, F1: 42, F2: (!false)}
                _ = fi16
                _ = v25
            }
            var v26 uint64 = 18446744073709551615
            v6 -= v6
            _ = fi15
            _ = v26
        }
        var v27 int32 = -1
        for fi28 := 0; fi28 < 20; fi28++ {
            v29 := []int{9223372036854775807, (fi28 - -2147483648)}
            if (0 <= uint(v7)) {
                v30 := GenS3{F0: uint(v7)}
                fmt.Println(fi28, v10)
                var v31 uint16 = v1
                _ = v30
                _ = v31
            } else {
                _ = len(v29)
                v27 = (v27 * 1)
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 int16 = -753
                    switch fi28 {
                    case 9223372036854775807:
                        fmt.Println(v3)
                    case 4294967295:
                        v6 = ((v10 + (v5 - v10)) - (((-2.25 - float32(float64((float32(1e10) - float32(float64(float32((3.14159 / (1.5 / 0.1))))))))) - (v6 / 0.1)) + 1e10))
                    default:
                        v34 := GenS4{F0: (v7 + (^(v7))), F1: uint(v1), F2: 2147483648}
                        _ = v34
                    }
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        v35 := GenS1{F0: (v3 + (v3 << ((255) & 63))), F1: 749, F2: v7}
        _ = v35
    }
    switch "" {
    case "a":
        switch "hello" {
        case "":
            var v36 string = ("\n\t" + "hello")
            _ = v36
        }
    case "foo bar":
        var v37 uint32 = 1
        _ = v37
    default:
        var v38 int8 = 42
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
