package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 bool; F1 int16 }
func genFn0(p0 int16, p1 bool, p2 string) int8 {
    fmt.Println(p2, p2, p0)
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float64 = float64(0.0)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 65535
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    var v5 float64 = ((-2.25 / 0.1) - 3.14159)
    var v8 uint16 = 839
    v9 := []int{2147483647, 42, 9223372036854775807}
    v10 := GenS1{F0: 0}
    v11 := GenS2{F0: -128, F1: 1}
    for fi12 := 0; fi12 < 8; fi12++ {
        switch "" {
        case "foo bar":
            var v13 float64 = v5
            _ = v13
        case "\n\t":
            fmt.Println(v5)
        default:
            if (0 < 127) {
                if (!((-2147483648 != -461) && (uint8(v2) < 249))) {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 float64 = 1.5
                        _ = fi14
                        _ = v15
                    }
                    v5 = ((3.14159 + v5) - v5)
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v8 = v8
                        var v17 float32 = 0.1
                        v17 = (v17 / 0.1)
                        _ = fi16
                        _ = v17
                    }
                } else {
                    if false {
                        var v18 uint16 = 2
                        if len(v9) > 0 {
                            v9[0] = int(v2)
                        }
                        switch "a" {
                        case "hello":
                            switch (fi12 - 0) {
                            case -723:
                                var v19 uint16 = (^(v18))
                                _ = v19
                            default:
                                v20 := GenS3{F0: fi12, F1: -95, F2: 42}
                                _ = v20
                            }
                        case "foo bar":
                            var v21 float64 = v5
                            _ = v21
                        case "\n\t":
                            var v22 float32 = -2.25
                            _ = v22
                        default:
                            for fi23 := 0; fi23 < 6; fi23++ {
                                var v24 float64 = v5
                                v18 += (v8 & 1)
                                v25 := GenS4{F0: -88, F1: (fi12 << ((522) & 63))}
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = v18
                    } else {
                        for fi26 := 0; fi26 < 1; fi26++ {
                            if (393 < 789) {
                                v8 += 65535
                                v5 = v5
                                var v27 uint64 = 18446744073709551615
                                _ = v27
                            } else {
                                var v28 float64 = 1.5
                                var v29 int8 = 100
                                fmt.Println(v8, v28)
                                _ = v28
                                _ = v29
                            }
                            v5 = 0.1
                            v2 -= 42
                            _ = fi26
                        }
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v30 := GenS4{F0: int8(fi12), F1: (-(fi12))}
                        _ = v30
                    }
                    fmt.Println(fi12)
                    var v31 float64 = float64(0.0)
                    _ = v31
                }
                v32 := v4[(^(v2))]
                v33 := GenS5{F0: true, F1: 1000}
                _ = v32
                _ = v33
            } else {
                var v34 uint8 = 42
                var v35 int16 = int16(v34)
                var v36 string = "\n\t"
                _ = v34
                _ = v35
                _ = v36
            }
        }
        var v37 int32 = 2147483647
        v38 := GenS5{F0: (!(42 != 627)), F1: 1000}
        _ = fi12
        _ = v37
        _ = v38
    }
    var v39 int32 = 0
    if (((^(v8)) <= (v8 ^ v8)) && false) {
        v2 = 65535
    } else {
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 bool = ((v8 << ((1) & 63)) >= v8)
            fmt.Println(v2, v41, v2)
            _ = fi40
            _ = v41
        }
        var v42 uint32 = (v2 << ((566) & 63))
        switch ("\n\t" + ("a" + "foo bar")) {
        case "\n\t":
            if false {
                if (32767 < -1) {
                    v8 = v8
                    for fi43 := 0; fi43 < 18; fi43++ {
                        v2 = (v2 >> ((0) & 63))
                        v42 += (^(v42))
                        var v44 uint = 0
                        _ = fi43
                        _ = v44
                    }
                } else {
                    switch -419 {
                    case 311:
                        fmt.Println(v5, v39, v42)
                    case 42:
                        v42 -= (v2 & v2)
                    }
                    var v45 float32 = (1e10 * (0.1 * 1.5))
                    _ = len(v4)
                    _ = v45
                }
            }
        case "":
            v46 := GenS1{F0: v39}
            _ = v46
        case "hello":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = v42
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
