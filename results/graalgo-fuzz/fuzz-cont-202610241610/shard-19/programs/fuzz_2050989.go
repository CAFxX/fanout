package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    if (997 > 1) {
        for fi4 := 0; fi4 < 2; fi4++ {
            v3 = (v3 - (v3 + 107))
            var v5 uint16 = 1
            _ = fi4
            _ = v5
        }
        fmt.Println(v3)
    }
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            if (!true) {
                for fi7 := 0; fi7 < 20; fi7++ {
                    switch (fi7 * fi6) {
                    case -1:
                        var v8 float64 = 0.0
                        _ = v8
                    default:
                        if (!false) {
                            v9 := GenS1{F0: 755, F1: (fi7 | (-(fi6))), F2: v3}
                            _ = v9
                        } else {
                            fmt.Println(fi7, v3, fi6)
                        }
                    }
                    if (-32768 < 42) {
                        v10 := GenS1{F0: 2147483647, F1: (fi6 + (fi7 / ((fi7) | 1))), F2: v3}
                        v3 -= v3
                        _ = v10
                    }
                    if ("a" <= "hello") {
                        if true {
                            v11 := GenS1{F0: 912, F1: fi6, F2: 2}
                            _ = v11
                        } else {
                            fmt.Println(fi6)
                            var v12 int64 = -1
                            _ = v12
                        }
                    }
                    _ = fi7
                }
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 12; fi15++ {
                    fmt.Println(fi6, v3, fi6)
                    _ = fi15
                }
            }
            _ = fi6
        }
        v3 = (v3 & (v3 + (v3 % ((149) | 1))))
        if ((2 != 124) || (!true)) {
            var v16 uint8 = (v3 >> ((0) & 63))
            var v17 float64 = 0.1
            v18 := []int{42}
            _ = v16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (490 != 0) {
            v3 = v3
            for fi20 := 0; fi20 < 4; fi20++ {
                if true {
                    var v21 uint32 = 0
                    var v22 uint64 = 345
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            var v23 int16 = -943
            _ = v23
        }
        v3 = (v3 / (((^(v3))) | 1))
        var v24 uint16 = 32767
        _ = fi19
        _ = v24
    }
    v3 = v3
    var v25 string = "foo bar"
    var v26 int = 180
    var v27 uint32 = 2147483648
    v1[v26] = (("foo bar" + (v25 + v25)) + (v25 + ("foo bar" + v25)))
    v28 := GenS1{F0: 323, F1: -9223372036854775808, F2: 20}
    if (!true) {
        var v29 float32 = (((float32((1e10 / 1e10)) / (((1.5 * 0.1) + 1.5) + float32(1e10))) + ((1.5 / -2.25) * -2.25)) * -2.25)
        for fi30 := 0; fi30 < 14; fi30++ {
            v31 := v1[-2147483648]
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
