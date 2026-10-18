package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string; F1 int64 }

func main() {
    var v1 int8 = 42
    var v2 uint64 = 4294967295
    var v3 bool = (!(1e10 <= 0.0))
    v4 := make(map[int]int)
    if (!((-(v1)) >= v1)) {
        switch "\n\t" {
        case "foo bar":
            var v5 int64 = -1
            _ = v5
        default:
            _ = len(v4)
        }
        switch 11 {
        case -2147483648:
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := GenS1{F0: 2}
                var v8 int32 = 662
                if v3 {
                    if (0 == 2147483647) {
                        if v3 {
                            if true {
                                v9 := GenS2{F0: "\n\t", F1: uint8(v2)}
                                _ = v9
                            }
                            v10 := GenS3{F0: (("foo bar" + "foo bar") + (("" + "hello") + "foo bar")), F1: -9223372036854775808}
                            v11 := []int{-142, (^(fi6)), (^(fi6))}
                            _ = v10
                            _ = v11
                        } else {
                            v12 := []int{(fi6 / ((130) | 1)), fi6, 0, (fi6 << ((473) & 63))}
                            _ = v12
                        }
                        var v13 float32 = (float32(0.1) * 1.5)
                        _ = len(v4)
                        _ = v13
                    }
                    v1 = 1
                    v14 := []int{fi6, (fi6 | (fi6 ^ fi6)), fi6}
                    _ = v14
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        default:
            var v15 int16 = 971
            _ = v15
        }
    }
    var v16 int64 = 457
    v16 = v16
    fmt.Println(v2, v16)
    for fi17 := 0; fi17 < 8; fi17++ {
        v4[fi17] = fi17
        fmt.Println(v2, v3)
        v2 = v2
        _ = fi17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v16 -= -720
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(fi19)
            _ = fi20
        }
        var v21 bool = v3
        _ = fi19
        _ = v21
    }
    v16 = 9223372036854775807
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int64 = v16
        _ = fi22
        _ = v23
    }
    var v24 int8 = 101
    fmt.Println(v16, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
