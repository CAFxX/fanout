package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ""
    v3 := []int{-735, -1, 32768}
    v3 = append(v3, 2147483647)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint = 109
        v6 := GenS1{F0: 657}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint8 = 127
    var v8 uint64 = 349
    v2 = ("a" + (v2 + ((v2 + ("\n\t" + v2)) + (v2 + v2))))
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: v8}
        v2 = v2
        _ = fi9
        _ = v10
    }
    var v11 string = ""
    v2 = ((v2 + ((v2 + "") + v2)) + v11)
    var v12 uint32 = 42
    if (((true && false) && false) && ((false || true) || true)) {
        v13 := GenS1{F0: (v8 % ((v8) | 1))}
        for fi14 := 0; fi14 < 12; fi14++ {
            if (true || true) {
                var v15 int8 = -1
                _ = v15
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v16 uint8 = 128
                v8 += v8
                _ = v16
            }
            switch 32767 {
            case 32767:
                var v17 int64 = -1
                _ = v17
            case 32768:
                var v18 uint64 = 449
                _ = v18
            default:
                var v19 int8 = 1
                _ = v19
            }
            _ = fi14
        }
        v2 = v11
        _ = v13
    } else {
        v3 = append(v3, -1)
        var v20 bool = ((!(((false || (false && ((3.14159 <= 1.5) || ((1 < 42) || (191 <= 0))))) || (true && true)) || (false || false))) || (-2147483648 >= -1))
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := GenS1{F0: v8}
            if (v20 && (1 < 65535)) {
                var v23 int64 = -933
                v24 := v1[(v12 * v12)]
                _ = v23
                _ = v24
            } else {
                if ((-947 == 9223372036854775807) || v20) {
                    var v25 uint8 = v7
                    fmt.Println(v8, v20)
                    for fi26 := 0; fi26 < 13; fi26++ {
                        if v20 {
                            var v27 int64 = 1
                            var v28 float64 = 0.0
                            _ = v27
                            _ = v28
                        }
                        if true {
                            var v29 int8 = -28
                            if ((0 == uint(v8)) && (!(!(1 <= 42)))) {
                                var v30 float64 = (0.0 * (1e10 / 1.5))
                                _ = v30
                            }
                            _ = v29
                        } else {
                            var v31 uint64 = (v8 & uint64(v12))
                            v12 -= 0
                            _ = v31
                        }
                        v7 += (^(v25))
                        _ = fi26
                    }
                    _ = v25
                } else {
                    if v20 {
                        if (820 < 32767) {
                            v32 := v1[v12]
                            v33 := v1[v12]
                            _ = v32
                            _ = v33
                        }
                        fmt.Println(fi21, v2, v11)
                        fmt.Println(v7, v7)
                    }
                }
            }
            v34 := GenS2{F0: 2, F1: (v7 + uint8(v8))}
            _ = fi21
            _ = v22
            _ = v34
        }
        _ = v20
    }
    switch 1 {
    case -513:
        var v35 int8 = 127
        _ = v35
    case 0:
        switch 9223372036854775807 {
        case 42:
            v11 = ((("foo bar" + ((v11 + v2) + ((v11 + v2) + ""))) + v11) + v11)
        default:
            if (!false) {
                v36 := GenS2{F0: 944, F1: (v7 | uint8(v12))}
                var v37 int = -86
                v38 := GenS3{F0: 32767}
                _ = v36
                _ = v37
                _ = v38
            } else {
                var v39 int16 = 1
                v8 -= v8
                _ = v39
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
