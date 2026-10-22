package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint8; F2 int8 }
func genFn0(p0 int64, p1 int8) uint8 {
    p1 -= -1
    if ((p0 * p0) == (p0 >> ((42) & 63))) {
        p1 += 37
    }
    switch ("" + "foo bar") {
    case "a":
        var v6 int64 = -9223372036854775808
        _ = v6
    case "":
        var v7 string = "\n\t"
        _ = v7
    case "hello":
        var v8 float32 = (0.1 / -2.25)
        _ = v8
    default:
        var v9 int8 = p1
        _ = v9
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 uint8 = 196
    v4 := make(map[uint32]int)
    var v5 uint8 = 127
    v10 := []int{9223372036854775807}
    v3 = v5
    var v11 uint32 = 2147483648
    fmt.Println(v3, v5)
    v12 := GenS1{F0: -128, F1: 0}
    v10 = append(v10, 128)
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: 0, F1: -128}
            _ = fi14
            _ = v15
        }
        var v16 int8 = int8(fi13)
        if (4294967295 > 4294967295) {
            v11 = v11
        } else {
            var v17 int64 = -1
            v18 := v4[(v11 << ((0) & 63))]
            for fi19 := 0; fi19 < 14; fi19++ {
                if len(v10) > 0 {
                    v10[0] = (^(fi19))
                }
                _ = fi19
            }
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v16
    }
    v10 = append(v10, -9223372036854775808)
    if true {
        v20 := GenS2{F0: 255, F1: 7, F2: (^(v5))}
        v3 = 31
        _ = v20
    }
    var v21 uint16 = 141
    var v22 int8 = -1
    v23 := GenS2{F0: 0, F1: 2, F2: uint8(v21)}
    for fi24 := 0; fi24 < 15; fi24++ {
        if (false && true) {
            v11 -= 4294967295
            for fi25 := 0; fi25 < 5; fi25++ {
                switch (fi25 & (fi24 + 0)) {
                case -9223372036854775808:
                    var v26 float32 = 3.14159
                    _ = v26
                case -2147483648:
                    fmt.Println(v21, fi25, v21)
                    v10 = append(v10, (^(fi24)))
                default:
                    var v27 uint8 = v3
                    _ = v27
                }
                v28 := GenS3{F0: v22, F1: v3, F2: v22}
                v21 -= (v21 & (v21 % (((v21 % ((255) | 1))) | 1)))
                _ = fi25
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                v2[("foo bar" + ((((("foo bar" + ("\n\t" + "\n\t")) + "foo bar") + "a") + "\n\t") + ""))] = v11
                _ = fi29
            }
        } else {
            v10 = append(v10, (fi24 % ((-429) | 1)))
        }
        v22 += 127
        if false {
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 float32 = 3.14159
                if true {
                    var v32 uint8 = v5
                    v21 -= v21
                    v33 := GenS1{F0: v22, F1: v22}
                    _ = v32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
        } else {
            v34 := v4[v11]
            v4[663] = (fi24 >> ((18446744073709551615) & 63))
            var v35 int8 = 100
            _ = v34
            _ = v35
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
