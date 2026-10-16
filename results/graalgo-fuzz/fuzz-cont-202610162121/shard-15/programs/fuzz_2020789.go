package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 string }

func main() {
    var v1 uint8 = 0
    var v2 float32 = float32((1e10 + (1e10 / (1.5 - 0.0))))
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 2147483647, F1: 1, F2: "hello"}
    var v6 uint64 = 124
    if true {
        v7 := GenS2{F0: 9223372036854775807, F1: 42}
        _ = v7
    } else {
        if (-1 != -1) {
            v8 := GenS3{F0: (v6 % ((v6) | 1)), F1: 0, F2: ""}
            v6 += (v6 | v6)
            for fi9 := 0; fi9 < 12; fi9++ {
                if false {
                    var v10 uint32 = 823
                    var v11 uint16 = 32767
                    v2 -= 1e10
                    _ = v10
                    _ = v11
                } else {
                    v12 := []int{fi9}
                    _ = v12
                }
                var v13 int8 = 0
                v14 := GenS1{F0: fi9, F1: v6, F2: "a"}
                _ = fi9
                _ = v13
                _ = v14
            }
            _ = v8
        } else {
            var v15 uint8 = 255
            v2 -= v2
            v1 = v1
            _ = v15
        }
        fmt.Println(v6, v1)
        v2 = -2.25
    }
    v6 = (v6 * (v6 << ((998) & 63)))
    v2 -= v2
    if true {
        switch -2147483648 {
        case 9223372036854775807:
            if (17 > -1) {
                fmt.Println(v1)
                v16 := []int{2147483647, 9223372036854775807}
                _ = v16
            }
        case 128:
            v6 = 856
        case -1:
            switch 997 {
            case 221:
                fmt.Println(v2, v1, v6)
            default:
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch "" {
                    case "\n\t":
                        _ = len(v4)
                    case "foo bar":
                        if false {
                            var v18 uint16 = 65535
                            for fi19 := 0; fi19 < 10; fi19++ {
                                var v20 float64 = (0.0 / ((1e10 / 1.0) - (-2.25 / float64((v2 * float32(-2.25))))))
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                    case "":
                        if false {
                            v1 -= v1
                        } else {
                            v21 := []int{int(v6), fi17}
                            if true {
                                var v22 int64 = -314
                                var v23 int64 = 1
                                _ = v22
                                _ = v23
                            } else {
                                var v24 uint16 = 1
                                _ = v24
                            }
                            var v25 uint32 = 1
                            _ = v21
                            _ = v25
                        }
                    default:
                        v1 = (v1 | v1)
                    }
                    _ = fi17
                }
            }
        default:
            v3[4294967295] = (0.0 >= 0.1)
        }
        var v26 float64 = (3.14159 - -2.25)
        _ = v26
    } else {
        v4[713] = 2
        v27 := []int{}
        _ = v27
    }
    v28 := GenS2{F0: 1, F1: -1}
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
