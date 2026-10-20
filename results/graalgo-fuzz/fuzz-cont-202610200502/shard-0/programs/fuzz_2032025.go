package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 int8) int8 {
    v5 := []int{127}
    v6 := GenS1{F0: p0, F1: true, F2: 886}
    _ = p0
    _ = v5
    _ = v6
    return (p0 << ((65535) & 63))
}

func main() {
    v1 := []string{}
    v2 := []string{}
    var v3 uint32 = 65535
    v4 := make(map[uint32]bool)
    v7 := GenS1{F0: -98, F1: (-1 == 0), F2: 0}
    var v8 float64 = 0.1
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: -1, F1: true, F2: -488}
        for fi11 := 0; fi11 < 13; fi11++ {
            if false {
                v3 += v3
                switch fi11 {
                case -15:
                    v8 = v8
                case 329:
                    v1 = append(v1, "foo bar")
                }
                if ((255 > 30) && true) {
                    fmt.Println(v3, fi11)
                    switch (fi11 << ((42) & 63)) {
                    case 32768:
                        fmt.Println(fi11)
                    case 4294967295:
                        v8 = (v8 + v8)
                    }
                    v8 += (3.14159 / v8)
                } else {
                    v3 += v3
                    var v12 float64 = v8
                    var v13 uint16 = 42
                    _ = v12
                    _ = v13
                }
            }
            var v14 string = "\n\t"
            var v15 float32 = 0.0
            v16 := GenS1{F0: -1, F1: true, F2: 2147483647}
            v15 += ((3.14159 - (float32(((1.5 - v8) * v8)) * 0.0)) - ((v15 / v15) - 0.0))
            if true {
                fmt.Println(v3, v15, v15)
                if true {
                    if true {
                        if (1 > uint64(v3)) {
                            v3 -= (v3 | v3)
                            fmt.Println(v8, v3, fi9)
                        }
                        fmt.Println(v8)
                    }
                    if (-586 < (fi9 << ((4294967295) & 63))) {
                        switch v14 {
                        case "foo bar":
                            if (v3 >= v3) {
                                v3 += (^(v3))
                            }
                        default:
                            if true {
                                var v17 float32 = float32(1.5)
                                _ = v17
                            } else {
                                v18 := GenS2{F0: 42, F1: 2, F2: 383}
                                v19 := GenS3{F0: 338}
                                _ = len(v1)
                                _ = v18
                                _ = v19
                            }
                            var v20 bool = ((953 < 65535) && ((v8 / (0.0 + v8)) != v8))
                            _ = v20
                        }
                        var v21 int16 = int16(fi11)
                        v1 = append(v1, "\n\t")
                        _ = v21
                    }
                } else {
                    if true {
                        v15 = v15
                    }
                }
                var v22 int8 = 0
                _ = v22
            }
            v14 = (("" + v14) + v14)
            _ = fi11
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v3 += 4294967295
        _ = fi23
    }
    v3 -= v3
    v24 := v4[(^(v3))]
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS4{F0: -1, F1: -1}
        v8 -= ((0.0 - -2.25) + v8)
        _ = fi25
        _ = v26
    }
    v4[(v3 % ((2147483647) | 1))] = true
    v8 += float64(1.5)
    v27 := GenS3{F0: 32767}
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 float64 = float64((0.0 / 1.0))
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
