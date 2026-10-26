package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int32; F1 int }
func genFn0() string {
    var v4 int64 = 0
    v5 := GenS1{F0: (v4 ^ (v4 << ((1) & 63))), F1: 127, F2: 1}
    for fi6 := 0; fi6 < 15; fi6++ {
        if (1 != 42) {
            v4 += 9223372036854775807
        } else {
            for fi7 := 0; fi7 < 4; fi7++ {
                if true {
                    v8 := []int{-790, fi6, (fi6 - fi6)}
                    switch 128 {
                    case 1:
                        v4 = v4
                        switch fi7 {
                        case 446:
                            var v9 uint8 = 187
                            _ = v9
                        case -317:
                            if true {
                                var v10 uint64 = 42
                                _ = v10
                            }
                        }
                        v4 += (v4 << ((42) & 63))
                    }
                    _ = v8
                }
                fmt.Println(v4, v4, fi7)
                v11 := []int{(fi6 + fi6), (fi7 - (fi7 + (fi7 / ((139) | 1)))), (fi7 ^ fi7)}
                _ = fi7
                _ = v11
            }
        }
        if (114 < 242) {
            var v12 uint32 = 42
            _ = v12
        } else {
            var v13 int16 = 32767
            v14 := GenS2{F0: true, F1: "foo bar"}
            switch fi6 {
            case 217:
                var v15 uint = 18446744073709551615
                _ = v15
            case -450:
                v16 := []int{fi6, 32767, fi6}
                _ = v16
            }
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    var v17 int8 = -123
    v17 += v17
    _ = v17
    _ = v4
    _ = v5
    return ("hello" + "\n\t")
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = 1e10
    var v3 int16 = 32767
    var v18 uint = 65535
    switch -9223372036854775808 {
    case 127:
        v18 = v18
    case 551:
        var v19 int = 748
        _ = v19
    default:
        v20 := []int{}
        _ = v20
    }
    v2 += ((((v1 / float32((1e10 + 0.1))) / 1e10) / ((((-2.25 * v1) * (0.1 - ((((1e10 + v1) * 1e10) - (v2 + v1)) - 0.1))) * v2) + v2)) * float32(0.0))
    v2 += 1.5
    if true {
        v1 = (1e10 / ((v1 / v2) / v2))
        var v21 uint16 = 1
        _ = v21
    } else {
        if (true && false) {
            var v22 int16 = v3
            v18 = (v18 % (((v18 - (v18 / ((692) | 1)))) | 1))
            _ = v22
        } else {
            var v23 uint64 = 1
            v3 -= v3
            v24 := GenS2{F0: true, F1: "\n\t"}
            _ = v23
            _ = v24
        }
    }
    v25 := GenS3{F0: 42, F1: 32767}
    v18 -= (v18 << ((v18) & 63))
    var v26 uint = v18
    var v27 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
