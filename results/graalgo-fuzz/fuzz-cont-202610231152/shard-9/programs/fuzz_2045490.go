package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint32 }
func genFn0(p0 string, p1 uint) int32 {
    if true {
        if false {
            var v6 float32 = ((-2.25 + -2.25) / 0.1)
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: 127, F1: (p1 - (p1 / ((p1) | 1))), F2: 0}
                _ = fi7
                _ = v8
            }
            p1 = 0
            _ = v6
        } else {
            var v9 int32 = 1
            v10 := []int{0, 334}
            v9 -= v9
            _ = v10
            _ = v9
        }
        p0 = ("" + ("" + (p0 + p0)))
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            if (fi11 == (fi11 + (fi11 % (((fi11 % (((^(fi11))) | 1))) | 1)))) {
                var v12 string = p0
                _ = v12
            } else {
                if (24 == 166) {
                    var v13 int64 = 0
                    var v14 float32 = (1.5 / (1e10 - (1e10 / 1e10)))
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v16 := []int{fi11, (-(fi15)), fi15, (fi11 * (^(fi11)))}
                        v17 := GenS2{F0: 0, F1: 255}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    var v18 int = -2147483648
                    var v19 int32 = int32(v18)
                    _ = v18
                    _ = v19
                }
                v20 := GenS2{F0: 1, F1: 127}
                v21 := GenS3{F0: -1, F1: 2, F2: 2147483647}
                _ = v20
                _ = v21
            }
            p1 = 255
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint64 = 4294967295
            var v24 int16 = -1
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi11
    }
    v25 := []int{2147483647, 0}
    _ = p0
    _ = p1
    _ = v25
    return 1
}
func genFn1(p0 int8) uint16 {
    p0 -= (p0 / ((p0) | 1))
    v26 := []int{9223372036854775807, 9223372036854775807, 123, 1}
    v27 := GenS3{F0: 67, F1: 133, F2: 1}
    v28 := GenS3{F0: -2147483648, F1: 42, F2: 0}
    _ = p0
    _ = v26
    _ = v27
    _ = v28
    return 65535
}

func main() {
    var v1 int = 567
    var v2 int = 42
    var v3 int16 = 1
    var v4 string = "foo bar"
    var v5 int8 = int8(v2)
    if false {
        v5 += v5
        var v29 uint8 = 1
        v30 := []int{0, 32767, v2, 2147483647}
        _ = v29
        _ = v30
    }
    fmt.Println(v1, v2, v4)
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 int8 = (v5 * 100)
        v33 := []int{}
        v33 = append(v33, (^(fi31)))
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v4, v5)
    var v34 float64 = (((1e10 * float64(1.5)) * (1.5 - -2.25)) - ((-2.25 - -2.25) / 1.5))
    v3 -= (v3 << ((86) & 63))
    var v35 int16 = v3
    if (4294967295 == 1) {
        v36 := []int{(^(v2)), v2, v2, -529}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
