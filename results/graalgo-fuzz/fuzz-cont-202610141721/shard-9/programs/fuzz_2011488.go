package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint) bool {
    var v6 float32 = ((((float32((-2.25 * ((0.1 - -2.25) / 1e10))) / 1.5) - 0.0) - 1.5) * (1e10 * 1e10))
    v7 := GenS1{F0: 42, F1: 0, F2: (^(p0))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 0
    v3 := []uint32{}
    var v4 uint8 = 42
    v5 := make(map[uint32]uint32)
    v8 := GenS2{F0: 80, F1: int64(v4)}
    v9 := GenS3{F0: 639, F1: 4294967295, F2: 0}
    v2 = (v2 | v2)
    v3 = append(v3, 975)
    var v10 bool = false
    v11 := GenS4{F0: 760, F1: ""}
    v2 += 883
    fmt.Println(v4, v2)
    if v10 {
        for fi12 := 0; fi12 < 15; fi12++ {
            _ = len(v5)
            for fi13 := 0; fi13 < 17; fi13++ {
                if v10 {
                    v14 := GenS4{F0: 1000, F1: "hello"}
                    if (32767 != -1) {
                        v10 = (!((v4 < v4) || true))
                    }
                    v5[1] = 2147483647
                    _ = v14
                } else {
                    v4 = v4
                    v2 = v2
                    var v15 uint16 = 1
                    v16 := v5[4294967295]
                    _ = v15
                    _ = v16
                }
                switch "\n\t" {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        v2 = -1
                        v17 := GenS5{F0: fi13, F1: -884}
                        var v18 string = "\n\t"
                        _ = v17
                        _ = v18
                    default:
                        v4 += (v4 ^ (v4 & (^(v4))))
                    }
                case "hello":
                    switch fi12 {
                    case 0:
                        v19 := GenS2{F0: 4, F1: 42}
                        _ = v19
                    default:
                        var v20 uint32 = 65535
                        switch (fi12 >> ((2) & 63)) {
                        case 42:
                            v20 += 727
                        }
                        var v21 bool = true
                        _ = v20
                        _ = v21
                    }
                case "a":
                    v2 -= -2147483648
                }
                _ = fi13
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                v10 = true
                v3 = append(v3, 42)
                _ = fi22
            }
            _ = fi12
        }
        switch ("" + "foo bar") {
        case "\n\t":
            var v23 int8 = 50
            _ = v23
        case "foo bar":
            var v24 uint8 = (^(v4))
            _ = v24
        default:
            v25 := GenS6{F0: 42, F1: (v2 + (v2 - (v2 >> ((0) & 63))))}
            _ = v25
        }
        if v10 {
            for fi26 := 0; fi26 < 16; fi26++ {
                if (v10 && (!(18446744073709551615 <= 255))) {
                    v10 = true
                } else {
                    _ = len(v5)
                    v10 = v10
                }
                _ = fi26
            }
            v2 -= (v2 ^ 0)
            if (32768 > -2147483648) {
                v27 := GenS7{F0: v2}
                _ = v27
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v10, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
