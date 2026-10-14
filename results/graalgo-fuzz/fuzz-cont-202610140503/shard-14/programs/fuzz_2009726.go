package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 string = "a"
    var v2 int8 = 100
    var v3 uint8 = 1
    var v4 int64 = 157
    var v5 int16 = -488
    v6 := make(map[int]string)
    if ((v3 == v3) || false) {
        for fi7 := 0; fi7 < 18; fi7++ {
            if ("\n\t" != v1) {
                var v8 uint8 = (v3 - (v3 ^ (v3 << ((0) & 63))))
                _ = v8
            }
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: 391}
                _ = fi9
                _ = v10
            }
            _ = fi7
        }
    } else {
        if false {
            switch v1 {
            case "hello":
                v11 := []int{712}
                _ = v11
            case "a":
                if true {
                    if (!true) {
                        v5 -= v5
                        for fi12 := 0; fi12 < 7; fi12++ {
                            var v13 float64 = float64(0.0)
                            if true {
                                v14 := []int{(^(fi12)), fi12, (fi12 ^ 9223372036854775807), (fi12 + (-(fi12)))}
                                switch v1 {
                                case "\n\t":
                                    for fi15 := 0; fi15 < 15; fi15++ {
                                        if true {
                                            v2 = 127
                                            _ = len(v14)
                                        } else {
                                            var v16 int8 = v2
                                            _ = v16
                                        }
                                        v4 = v4
                                        _ = fi15
                                    }
                                default:
                                    var v17 uint32 = 1
                                    _ = v17
                                }
                                _ = v14
                            }
                            var v18 uint8 = v3
                            _ = fi12
                            _ = v13
                            _ = v18
                        }
                    }
                }
            }
            v4 = (v4 | (v4 * 523))
        }
        v2 += (v2 * int8(v5))
        var v19 uint64 = 457
        _ = v19
    }
    fmt.Println(v3)
    v20 := GenS2{F0: 4294967295}
    if false {
        v5 = v5
    } else {
        v3 -= v3
        v21 := GenS3{F0: (!(v4 <= -9223372036854775808)), F1: v5}
        _ = v21
    }
    var v22 int16 = (v5 & v5)
    v23 := GenS2{F0: 849}
    if true {
        v24 := GenS1{F0: 255}
        if ((-76 == -551) && false) {
            if (42 != -2147483648) {
                v1 = (("hello" + "\n\t") + (v1 + v1))
                switch -937 {
                case -9223372036854775808:
                    var v25 uint64 = 42
                    _ = v25
                case 32768:
                    if false {
                        fmt.Println(v2, v4, v1)
                        if true {
                            v5 += 1000
                        } else {
                            var v26 float64 = 3.14159
                            v26 = 1.5
                            _ = v26
                        }
                        v27 := GenS2{F0: 726}
                        _ = v27
                    }
                case -1:
                    var v28 int = 0
                    _ = v28
                }
            } else {
                var v29 uint64 = 1
                _ = len(v6)
                _ = v29
            }
        } else {
            v2 += v2
            var v30 bool = (!(0 == 761))
            v2 += v2
            _ = v30
        }
        _ = v24
    }
    v31 := []int{}
    v32 := GenS3{F0: (1e10 > -2.25), F1: (v22 ^ (v5 / (((v5 << ((0) & 63))) | 1)))}
    fmt.Println(v5)
    v6[-824] = v1
    var v33 string = v1
    if true {
        v3 += (v3 << ((4294967295) & 63))
        v3 = 2
        v5 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
