package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 uint, p1 int8) int64 {
    p0 = (^(p0))
    p1 = -1
    v4 := []int{-610, 128, 378}
    p1 = -61
    _ = p0
    _ = p1
    _ = v4
    return 143
}

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 9223372036854775807
    var v3 int16 = 0
    switch -430 {
    case 1:
        v5 := GenS1{F0: false, F1: 128}
        _ = v5
    case 127:
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: (!(!false)), F1: 1}
            _ = fi6
            _ = v7
        }
    case -2147483648:
        v8 := GenS2{F0: 2}
        _ = v8
    default:
        fmt.Println(v1, v2, v3)
    }
    var v9 int16 = v3
    var v10 int16 = (v3 & (v3 + (v9 * 512)))
    if (true && ("" < "hello")) {
        switch 42 {
        case -111:
            if (false && false) {
                var v11 string = "hello"
                fmt.Println(v9, v10)
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 11; fi12++ {
                if true {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v14 := GenS3{F0: fi13}
                        v15 := GenS2{F0: 2147483647}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS4{F0: 65535, F1: ((v2 | 1) <= v2), F2: 255}
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if false {
                            v18 := []int{4294967295, fi17, (fi17 + fi12)}
                            _ = v18
                        }
                        _ = fi17
                    }
                    _ = v16
                }
                if true {
                    var v19 float64 = v1
                    fmt.Println(v10, v19)
                    _ = v19
                }
                _ = fi12
            }
        }
    } else {
        fmt.Println(v3)
    }
    v20 := GenS4{F0: 0, F1: (42 >= 127), F2: 467}
    v10 += (v9 | (-(v9)))
    v9 += 42
    v1 += -2.25
    v2 = v2
    v2 += v2
    v1 -= float64(1.5)
    var v21 int32 = -703
    v10 += (v9 >> ((42) & 63))
    var v22 int = -1
    v23 := []int{(v22 - v22), (v22 % ((128) | 1))}
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS4{F0: 1, F1: false, F2: 42}
        v26 := GenS4{F0: 1, F1: (false && (false && (!((!(("\n\t" + "a") > "a")) || (true || true))))), F2: 979}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int64 = v2
    if false {
        var v28 int32 = v21
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v9
}
