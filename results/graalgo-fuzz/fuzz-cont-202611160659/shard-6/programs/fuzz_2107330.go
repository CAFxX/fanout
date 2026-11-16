package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
func genFn0(p0 uint64) int8 {
    switch -1 {
    case 32768:
        var v6 float32 = (-2.25 - 1.5)
        _ = v6
    case -132:
        p0 += p0
    case 1:
        var v7 int8 = 1
        _ = v7
    }
    var v8 float64 = 1.5
    if (!false) {
        v9 := GenS1{F0: 946, F1: -959, F2: -98}
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint16 = 0
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        v12 := []int{-9223372036854775808, 128, -2147483648}
        var v13 uint = 1
        if len(v12) > 0 {
            v12[0] = -9223372036854775808
        }
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v8
    return 100
}
func genFn1(p0 int16, p1 bool, p2 int64) int16 {
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: (fi15 % (((fi14 ^ fi14)) | 1)), F1: 134, F2: 9223372036854775807}
            p2 -= (p2 + p2)
            if p1 {
                if p1 {
                    v17 := []int{2147483647, fi15}
                    if (255 == 706) {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    } else {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        if (-2147483648 == 2147483647) {
                            var v18 uint16 = 42
                            _ = v18
                        }
                        var v19 uint32 = 496
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = v19
                    }
                    v20 := GenS1{F0: (-(fi14)), F1: 0, F2: p2}
                    _ = v17
                    _ = v20
                }
            } else {
                v21 := GenS1{F0: (fi15 << ((0) & 63)), F1: 42, F2: 42}
                var v22 int32 = 127
                p1 = (2 == 42)
                _ = v21
                _ = v22
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v23 int8 = -80
    v24 := []int{}
    if len(v24) > 0 {
        v24[0] = -875
    }
    var v25 bool = (-1 <= -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    _ = v25
    return -32768
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 string = ("a" + "a")
    var v4 string = ""
    var v5 uint16 = 307
    v3 = v4
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint64 = 556
        var v28 int16 = -32768
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := v2[19]
        _ = fi29
        _ = v30
    }
    v5 -= (^(v5))
    delete(v1, -9223372036854775808)
    var v31 uint = 688
    if (!true) {
        v32 := GenS1{F0: 0, F1: 1, F2: -9223372036854775808}
        v31 -= 0
        _ = v32
    }
    v33 := GenS2{F0: v5, F1: 209, F2: 0}
    var v34 string = v4
    for fi35 := 0; fi35 < 4; fi35++ {
        v3 = v3
        v4 = v3
        var v36 bool = false
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
