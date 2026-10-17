package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := []int{697, 1, 397}
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    v4[9223372036854775807] = "\n\t"
    if (!true) {
        if (128 >= 42) {
            switch -285 {
            case 839:
                if false {
                    v5 := GenS1{F0: false}
                    var v6 float64 = (1.5 + -2.25)
                    if true {
                        for fi7 := 0; fi7 < 2; fi7++ {
                            v6 -= v6
                            var v8 int64 = 0
                            _ = fi7
                            _ = v8
                        }
                    } else {
                        switch 0 {
                        case -806:
                            var v9 uint8 = 201
                            _ = v9
                        case 0:
                            v6 -= (0.1 - (1e10 + v6))
                        case 2147483647:
                            for fi10 := 0; fi10 < 10; fi10++ {
                                v11 := GenS2{F0: -9223372036854775808, F1: -628, F2: 65535}
                                fmt.Println(v6, v6)
                                v6 += v6
                                _ = fi10
                                _ = v11
                            }
                        default:
                            v6 = 0.1
                        }
                        v6 = v6
                    }
                    _ = v5
                    _ = v6
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v12 bool = (!true)
                        _ = v12
                    }
                    var v13 int16 = 0
                    _ = v13
                }
            }
            v14 := GenS3{F0: 1, F1: 32767}
            var v15 int8 = 100
            _ = v14
            _ = v15
        } else {
            var v16 int64 = 47
            for fi17 := 0; fi17 < 17; fi17++ {
                v16 = (-(v16))
                v18 := GenS4{F0: 1}
                fmt.Println(fi17, v16)
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    }
    var v19 int32 = 2147483647
    fmt.Println(v19)
    if false {
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS3{F0: -1, F1: 504}
            if true {
                switch (fi20 << ((255) & 63)) {
                case 1:
                    if len(v2) > 0 {
                        v2[0] = fi20
                    }
                default:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v23 := GenS2{F0: -9223372036854775808, F1: -511, F2: 132}
                        if len(v2) > 0 {
                            v2[0] = (-(fi20))
                        }
                        v24 := v1[("foo bar" + "hello")]
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = len(v4)
                var v25 uint64 = 0
                _ = v25
            } else {
                var v26 float32 = -2.25
                _ = v26
            }
            v19 -= (v19 * -2147483648)
            _ = fi20
            _ = v21
        }
    }
    v19 += v19
    var v27 uint64 = 2
    v19 = (-(v19))
    v28 := GenS4{F0: 778}
    fmt.Println(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
