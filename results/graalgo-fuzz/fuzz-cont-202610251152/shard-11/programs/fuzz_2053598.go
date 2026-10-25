package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int64; F2 int }
type GenS5 struct { F0 int8; F1 int64; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int64 = 9223372036854775807
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 856, F1: (32768 < 128), F2: 255}
    _ = v9
    return 2
}

func main() {
    var v1 int64 = -28
    v2 := []int{4294967295, 9223372036854775807, -856, 9223372036854775807}
    var v3 uint16 = 375
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    var v6 uint16 = v3
    _ = len(v5)
    fmt.Println(v3, v1, v1)
    var v10 int64 = 0
    v11 := v5[1]
    v2 = append(v2, 32767)
    for fi12 := 0; fi12 < 18; fi12++ {
        v6 += 32767
        if (128 > 255) {
            fmt.Println(v6, v11, v3)
            if false {
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := v4[1]
                    _ = fi13
                    _ = v14
                }
                v15 := GenS2{F0: "", F1: (^(v11)), F2: (v3 & (v3 - 255))}
                fmt.Println(v11, v6, v6)
                _ = v15
            } else {
                _ = len(v5)
                fmt.Println(v11, v3)
                v16 := v4[(^(v11))]
                _ = v16
            }
        } else {
            if (!true) {
                switch fi12 {
                case -573:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v11 -= (v11 >> ((1) & 63))
                        var v18 string = ((((("" + "hello") + "a") + ("hello" + "")) + "foo bar") + (("a" + "\n\t") + "a"))
                        v19 := GenS3{F0: 283}
                        fmt.Println(v6, fi17, v3)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case -2147483648:
                    v6 = (v3 & v6)
                case -658:
                    var v20 uint32 = 680
                    _ = v20
                }
                v21 := GenS4{F0: "foo bar", F1: v10, F2: fi12}
                _ = v21
            } else {
                var v22 uint8 = 128
                switch ("\n\t" + "foo bar") {
                case "":
                    fmt.Println(v11, v11)
                case "\n\t":
                    var v23 uint16 = v6
                    _ = v23
                case "hello":
                    var v24 float64 = 0.1
                    _ = v24
                default:
                    v22 += (^(v22))
                }
                switch fi12 {
                case 653:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        switch (fi25 * (fi12 >> ((genFn0()) & 63))) {
                        case 9223372036854775807:
                            var v26 int16 = 975
                            _ = v26
                        case -961:
                            var v27 string = "foo bar"
                            _ = v27
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v6 = v6
                            v28 := GenS5{F0: -1, F1: v1, F2: uint(v11)}
                            v29 := GenS3{F0: 18446744073709551615}
                            _ = v28
                            _ = v29
                        }
                        v22 += (v22 / ((v22) | 1))
                        _ = fi25
                    }
                case 9223372036854775807:
                    delete(v5, (v11 ^ (v11 << ((uint(v11)) & 63))))
                case 957:
                    var v30 uint8 = (v22 & (v22 << ((979) & 63)))
                    _ = v30
                }
                _ = v22
            }
        }
        _ = fi12
    }
    v1 = -775
    v1 = v1
    v10 = -884
    v31 := GenS3{F0: 42}
    if (-709 == 0) {
        if false {
            v32 := GenS5{F0: 100, F1: v1, F2: 0}
            switch 2147483647 {
            case 42:
                v1 += v1
            case -181:
                var v33 float32 = (-2.25 * float32(float64(0.0)))
                _ = v33
            case 0:
                fmt.Println(v10, v3, v3)
            default:
                for fi34 := 0; fi34 < 12; fi34++ {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 int32 = 2147483647
                        var v37 int16 = 878
                        var v38 uint16 = (^(v6))
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi34
                }
            }
            var v39 int64 = v1
            _ = v32
            _ = v39
        }
    }
    var v40 uint8 = 1
    v41 := GenS6{F0: (v6 % ((v3) | 1))}
    if false {
        v2 = append(v2, -1)
        v40 = (v40 * (^(v40)))
    }
    v42 := v4[v11]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
