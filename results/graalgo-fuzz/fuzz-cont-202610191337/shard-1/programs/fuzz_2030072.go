package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
func genFn0() int16 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 float64 = (-2.25 / (1e10 / -2.25))
        v5 -= ((3.14159 + (v5 - v5)) - 1.5)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{fi6, (-(fi6)), 127, fi6}
        var v8 uint8 = 105
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = 1.5
    v10 := GenS1{F0: 9223372036854775807}
    v11 := GenS1{F0: 283}
    _ = v10
    _ = v11
    _ = v9
    return -610
}

func main() {
    var v1 string = ""
    var v2 uint64 = 432
    var v3 int32 = 42
    if (953 > 65535) {
        if false {
            switch 9223372036854775807 {
            case 1:
                fmt.Println(v3)
            default:
                if (2 == 497) {
                    switch -444 {
                    case 32767:
                        var v12 int16 = -208
                        _ = v12
                    default:
                        v13 := GenS2{F0: 814, F1: genFn0()}
                        _ = v13
                    }
                    if false {
                        var v14 string = "foo bar"
                        _ = v14
                    }
                } else {
                    v2 = 18446744073709551615
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v3 -= 313
                        var v16 int32 = (v3 | v3)
                        _ = fi15
                        _ = v16
                    }
                }
            }
            v17 := []int{127, -2147483648, -1}
            _ = v17
        }
        switch 32768 {
        case -213:
            var v18 uint8 = 66
            _ = v18
        case 633:
            v3 = (^(v3))
        }
        v3 += v3
    }
    fmt.Println(v3)
    var v19 uint = 2
    v1 = "a"
    v20 := []int{42, 32768, 0, -1}
    var v21 uint8 = 128
    var v22 uint8 = (v21 ^ 80)
    if false {
        v23 := GenS3{F0: 1, F1: v1}
        v20 = append(v20, 1)
        _ = v23
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v24 := GenS4{F0: 0, F1: 2, F2: 827}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v1 = ("" + v1)
    if len(v20) > 0 {
        v20[0] = 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
}
