package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 int32; F1 uint; F2 uint }
func genFn0(p0 uint64, p1 int64) uint32 {
    v6 := []int{4294967295, -185}
    var v7 uint32 = 2
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (v7 ^ v7)
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := []uint32{932, 409}
    var v5 uint32 = 65535
    var v8 bool = false
    if ((3.14159 - 1e10) < 0.1) {
        fmt.Println(v8)
    } else {
        v1 = v1
    }
    v9 := GenS1{F0: 255, F1: 0, F2: v1}
    var v10 int = 373
    v11 := GenS2{F0: 42, F1: -127}
    v8 = (v8 || v8)
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v10, v1)
            _ = fi13
        }
        _ = fi12
    }
    var v14 string = "a"
    switch v10 {
    case 32767:
        fmt.Println(v5, v10, v8)
    case -676:
        v15 := GenS1{F0: 65535, F1: 1, F2: v1}
        _ = v15
    default:
        if v8 {
            for fi16 := 0; fi16 < 12; fi16++ {
                if v8 {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v10 = (fi16 / (((fi16 + fi17)) | 1))
                        var v18 float64 = float64(1.5)
                        v5 -= v5
                        _ = fi17
                        _ = v18
                    }
                    switch "foo bar" {
                    case "\n\t":
                        for fi19 := 0; fi19 < 18; fi19++ {
                            switch (v10 << ((4294967295) & 63)) {
                            case -1:
                                var v20 uint64 = 4294967295
                                _ = v20
                            }
                            _ = fi19
                        }
                    }
                } else {
                    var v21 uint64 = 4294967295
                    v21 -= v21
                    _ = v21
                }
                var v22 uint32 = genFn0(4294967295, -9223372036854775808)
                if true {
                    v23 := GenS3{F0: true, F1: (v10 / ((886) | 1)), F2: (42 != 236)}
                    var v24 int16 = v1
                    var v25 int8 = 109
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    if (!(42 < 18446744073709551615)) {
                        var v26 bool = v8
                        _ = v26
                    } else {
                        v4 = append(v4, (v5 - v5))
                        switch "a" {
                        case "foo bar":
                            var v27 bool = v8
                            _ = v27
                        }
                    }
                    for fi28 := 0; fi28 < 12; fi28++ {
                        switch v10 {
                        case -58:
                            if (((369 == 23) || (false && v8)) || ((v1 / ((-248) | 1)) < (^(v1)))) {
                                var v29 float32 = -2.25
                                switch fi16 {
                                case 4294967295:
                                    var v30 int32 = int32(v22)
                                    _ = v30
                                }
                                _ = v29
                            }
                        default:
                            var v31 uint8 = 24
                            _ = v31
                        }
                        v2[((v14 + "foo bar") + (((v14 + "\n\t") + "") + v14))] = int(v1)
                        _ = fi28
                    }
                }
                _ = fi16
                _ = v22
            }
        }
    }
    v1 = v1
    v32 := GenS2{F0: -2147483648, F1: 1}
    v33 := GenS4{F0: 42, F1: 42, F2: 18446744073709551615}
    v5 = (^(v5))
    var v34 int64 = 0
    v1 = 32767
    switch v14 {
    case "hello":
        v1 = v1
    case "foo bar":
        v35 := GenS2{F0: 201, F1: v34}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
