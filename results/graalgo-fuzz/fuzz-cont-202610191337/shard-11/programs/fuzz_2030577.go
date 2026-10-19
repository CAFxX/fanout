package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }

func main() {
    var v1 string = "hello"
    v2 := []string{"a"}
    var v3 float64 = ((1e10 + -2.25) - 1.5)
    var v4 uint = 203
    var v5 int16 = int16(v4)
    var v6 uint = (v4 << (((^(v4))) & 63))
    for fi7 := 0; fi7 < 9; fi7++ {
        if len(v2) > 0 {
            v2[0] = v1
        }
        if true {
            switch 32767 {
            case 808:
                if true {
                    v8 := GenS1{F0: 334}
                    fmt.Println(v3, v1)
                    var v9 float32 = 0.0
                    _ = v8
                    _ = v9
                } else {
                    v6 += v6
                    v10 := GenS2{F0: v6, F1: 65535, F2: 667}
                    _ = v10
                }
            case 1:
                for fi11 := 0; fi11 < 17; fi11++ {
                    v6 = (v6 / (((v6 / (((v6 - (v6 >> (((v6 >> ((639) & 63))) & 63)))) | 1))) | 1))
                    _ = fi11
                }
            default:
                v5 -= (v5 % (((v5 >> ((4294967295) & 63))) | 1))
            }
            switch fi7 {
            case 0:
                v3 -= v3
            case -9223372036854775808:
                for fi12 := 0; fi12 < 12; fi12++ {
                    if (2147483647 == 2147483647) {
                        var v13 float64 = (0.0 * (v3 - -2.25))
                        if false {
                            v4 = v4
                        } else {
                            var v14 int = 913
                            if false {
                                v13 -= float64(0.1)
                            }
                            v15 := GenS3{F0: false}
                            _ = v14
                            _ = v15
                        }
                        _ = v13
                    }
                    _ = fi12
                }
            }
        } else {
            var v16 uint16 = 2
            if len(v2) > 0 {
                v2[0] = v1
            }
            var v17 uint8 = 2
            _ = v16
            _ = v17
        }
        v18 := GenS2{F0: (v4 | v4), F1: 121, F2: v4}
        _ = fi7
        _ = v18
    }
    v6 = v6
    v19 := GenS3{F0: ((761 == 2) || false)}
    if (487 < 666) {
        v20 := GenS4{F0: 189}
        _ = v20
    }
    var v21 bool = ((!false) && true)
    v22 := GenS1{F0: 2}
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint16 = 2
        v4 = (v4 >> ((v6) & 63))
        _ = fi23
        _ = v24
    }
    v25 := GenS5{F0: (v1 + (v1 + "\n\t"))}
    v4 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
