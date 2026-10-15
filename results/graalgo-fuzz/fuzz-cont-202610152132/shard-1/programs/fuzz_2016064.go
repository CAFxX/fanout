package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }

func main() {
    var v1 uint64 = 1
    var v2 uint64 = (v1 % (((^(v1))) | 1))
    var v3 int32 = -413
    var v4 uint64 = 1
    var v5 uint16 = 0
    if false {
        v3 -= v3
        var v6 uint64 = v2
        if (809 != v2) {
            v4 -= (v2 * 326)
        } else {
            var v7 uint = 2
            for fi8 := 0; fi8 < 8; fi8++ {
                for fi9 := 0; fi9 < 6; fi9++ {
                    if ((1.5 - (-2.25 * 3.14159)) >= 0.0) {
                        var v10 uint64 = 2
                        v2 = 42
                        if ((v3 * v3) != 0) {
                            switch "a" {
                            case "\n\t":
                                v11 := []int{fi9, fi9}
                                _ = v11
                            default:
                                var v12 int8 = 100
                                _ = v12
                            }
                        } else {
                            v13 := GenS1{F0: "", F1: (v3 / (((v3 >> (((v7 / (((^(v7))) | 1))) & 63))) | 1)), F2: (^(v1))}
                            _ = v13
                        }
                        _ = v10
                    }
                    _ = fi9
                }
                fmt.Println(v7, v6, v2)
                _ = fi8
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                v4 += (v1 % ((v1) | 1))
                if true {
                    var v15 int32 = (v3 / ((v3) | 1))
                    _ = v15
                }
                _ = fi14
            }
            _ = v7
        }
        _ = v6
    }
    v5 += v5
    v4 = 465
    v3 = v3
    v3 = v3
    var v16 int = 128
    for fi17 := 0; fi17 < 8; fi17++ {
        v1 = v4
        v4 = v2
        _ = fi17
    }
    var v18 uint8 = 0
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int64 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 float32 = (0.1 + (3.14159 - 3.14159))
            var v24 float32 = 0.0
            var v25 float32 = 1.5
            v3 += v3
            v26 := []int{32768, fi22, fi22, (fi22 & -9223372036854775808)}
            switch (fi22 % ((fi22) | 1)) {
            case 108:
                v27 := GenS1{F0: "foo bar", F1: v3, F2: 0}
                _ = v27
            case 42:
                v18 -= 236
            }
            _ = len(v26)
            v16 += (v16 % ((32767) | 1))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi21
    }
    var v28 int64 = 9223372036854775807
    var v29 uint64 = (v2 + (v1 - v4))
    v4 += (v29 / (((v29 ^ (v4 * v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
