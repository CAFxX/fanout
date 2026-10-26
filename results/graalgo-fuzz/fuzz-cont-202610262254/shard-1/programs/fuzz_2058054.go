package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 string = ("hello" + "foo bar")
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    v4 := make(map[uint32]uint32)
    var v5 int8 = 38
    var v6 bool = ((true && true) && true)
    v7 := GenS1{F0: 42, F1: (v5 - (v5 >> ((1) & 63))), F2: v1}
    v5 = v5
    v3 -= (v3 * v3)
    var v8 bool = ((!(!(v6 && true))) || (65535 >= 2147483648))
    var v9 float64 = 1e10
    var v10 uint32 = 2147483647
    v6 = (-694 == -62)
    v8 = (488 > 1000)
    for fi11 := 0; fi11 < 11; fi11++ {
        if (3.14159 == 0.0) {
            for fi12 := 0; fi12 < 8; fi12++ {
                fmt.Println(v10, v9, v6)
                v9 = ((float64(0.1) + (((v9 - float64(float32(float64((0.1 + 0.0))))) / 0.1) * (0.0 - -2.25))) - v9)
                var v13 bool = v8
                v9 = 0.1
                _ = fi12
                _ = v13
            }
            delete(v2, v10)
        }
        switch "" {
        case "foo bar":
            fmt.Println(v1, v5, v1)
        }
        _ = fi11
    }
    switch ((v1 + "foo bar") + v1) {
    case "hello":
        var v14 int64 = 0
        _ = v14
    case "":
        v15 := GenS2{F0: -893, F1: 42, F2: 227}
        _ = v15
    case "\n\t":
        v16 := []int{127, 2147483647}
        _ = v16
    default:
        v3 += v3
    }
    var v17 int32 = -842
    switch 2147483647 {
    case 9223372036854775807:
        for fi18 := 0; fi18 < 11; fi18++ {
            switch fi18 {
            case 542:
                v5 = (v5 & 100)
            case 383:
                v19 := []int{fi18, (fi18 << ((53) & 63)), -2147483648, int(v10)}
                _ = v19
            default:
                if true {
                    var v20 int8 = 0
                    switch ("foo bar" + (v1 + v1)) {
                    case "\n\t":
                        v6 = (((v3 <= v3) || (int(v17) > fi18)) || true)
                    case "foo bar":
                        switch fi18 {
                        case -668:
                            v6 = (uint64(v10) >= (v3 / ((v3) | 1)))
                        default:
                            v9 = (v9 / v9)
                        }
                    default:
                        if ((v1 <= (v1 + v1)) && (-1 >= v17)) {
                            if (v8 || (0 == 20)) {
                                if (v6 && v8) {
                                    for fi21 := 0; fi21 < 5; fi21++ {
                                        v10 = v10
                                        _ = fi21
                                    }
                                    v22 := GenS3{F0: 762, F1: v17, F2: v6}
                                    var v23 int16 = 911
                                    _ = v22
                                    _ = v23
                                } else {
                                    v8 = (473 <= 32767)
                                    var v24 uint = 65535
                                    _ = v24
                                }
                            } else {
                                var v25 uint64 = v3
                                _ = v25
                            }
                        } else {
                            v26 := GenS3{F0: 9223372036854775807, F1: v17, F2: v6}
                            var v27 uint8 = 127
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = v20
                }
            }
            v5 += (v5 << ((158) & 63))
            fmt.Println(v5, v3)
            _ = fi18
        }
    case -1:
        if (v6 && ((v3 | 451) > (v3 & 927))) {
            var v28 int = -249
            v29 := GenS4{F0: 0, F1: 128}
            _ = v28
            _ = v29
        }
    case 2147483647:
        v6 = (2 != 926)
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 int16 = 257
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        switch 172 {
        case 32767:
            v33 := v2[(^(v10))]
            _ = v33
        case -914:
            fmt.Println(v8, v9, v3)
        }
        var v34 bool = v6
        _ = fi32
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
