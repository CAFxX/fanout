package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 bool) uint32 {
    v7 := []int{}
    p0 = false
    _ = p0
    _ = v7
    return 2147483647
}
func genFn1(p0 string) int16 {
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64((1.5 + 0.1))
    var v3 int8 = 0
    v4 := []uint32{2, 65535}
    v5 := make(map[uint32]int)
    v6 := []int{420, 32767}
    var v9 int16 = 42
    var v10 int = 32768
    v11 := GenS1{F0: 255, F1: 1}
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 bool = (!(false || false))
        v12 = "a"
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int64 = -744
            v10 = fi15
            _ = len(v1)
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v12 = ("a" + v12)
    v17 := v5[2]
    fmt.Println(v10, v10)
    v3 -= v3
    if false {
        if ((^(v10)) > (v17 << ((42) & 63))) {
            _ = len(v5)
        }
        var v18 uint64 = 42
        _ = v18
    }
    if (!((0.1 * ((0.1 + -2.25) * (3.14159 + (1e10 * 0.1)))) > 1e10)) {
        v10 += 747
        v19 := GenS1{F0: 2, F1: 1}
        var v20 uint32 = genFn0((160 == 128))
        _ = v19
        _ = v20
    }
    var v21 int32 = 42
    v5[2147483647] = (^(v17))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = fi22
        _ = v23
    }
    _ = len(v4)
    for fi24 := 0; fi24 < 18; fi24++ {
        v9 = (v9 ^ v9)
        switch v12 {
        case "foo bar":
            fmt.Println(v21)
        case "hello":
            if false {
                v12 = v12
                v12 = (v12 + v12)
            }
        default:
            if (!false) {
                fmt.Println(v10, v3, v17)
                v9 = (v9 * (v9 ^ v9))
            }
        }
        _ = fi24
    }
    if true {
        for fi25 := 0; fi25 < 12; fi25++ {
            v3 -= (v3 & (v3 - -1))
            _ = fi25
        }
        v26 := GenS2{F0: v9, F1: 18446744073709551615, F2: 9223372036854775807}
        _ = v26
    } else {
        v27 := GenS1{F0: 276, F1: 2}
        v3 = (v3 & -10)
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                if false {
                    v2 = (3.14159 * float64((float32((v2 + v2)) + (0.1 + (0.0 + 0.1)))))
                    v30 := GenS3{F0: (v3 % ((v3) | 1)), F1: 32767, F2: (v9 + (v9 | 1000))}
                    v31 := v1[v12]
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint64 = 42
                    if (!true) {
                        fmt.Println(v10, v32, fi28)
                        v10 = 2147483647
                    }
                    _ = v32
                }
                var v33 float32 = -2.25
                v3 -= -76
                _ = fi29
                _ = v33
            }
            _ = fi28
        }
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
