package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint64 }
func genFn0(p0 uint16, p1 int8, p2 uint64) bool {
    if (864 != 0) {
        var v5 uint64 = 18446744073709551615
        v6 := GenS1{F0: ("hello" + ("" + "\n\t"))}
        var v7 float64 = 1e10
        _ = v5
        _ = v6
        _ = v7
    } else {
        p0 += p0
        fmt.Println(p1)
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 bool = true
    v2 := []int{-781, 2147483647, -691}
    v3 := make(map[int]bool)
    v4 := []int{}
    if v1 {
        v8 := GenS1{F0: "hello"}
        _ = v8
    }
    if v1 {
        for fi9 := 0; fi9 < 15; fi9++ {
            if (128 < 2) {
                if false {
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v1 = (42 >= -2147483648)
                        if v1 {
                            v1 = v1
                            var v11 bool = v1
                            _ = v11
                        } else {
                            fmt.Println(v1, v1)
                            var v12 uint64 = 18446744073709551615
                            v12 += v12
                            _ = v12
                        }
                        v13 := GenS2{F0: 792, F1: 0, F2: v1}
                        var v14 int8 = -84
                        _ = fi10
                        _ = v13
                        _ = v14
                    }
                }
            } else {
                for fi15 := 0; fi15 < 11; fi15++ {
                    if v1 {
                        delete(v3, fi9)
                    } else {
                        var v16 int16 = -42
                        if (-1 == -2147483648) {
                            if v1 {
                                v17 := GenS3{F0: -387, F1: 1}
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                                if true {
                                    var v18 string = "hello"
                                    _ = v18
                                } else {
                                    if ((v16 != (^(v16))) || true) {
                                        v19 := v3[fi9]
                                        v4 = append(v4, int(v16))
                                        _ = v19
                                    } else {
                                        var v20 int = 32768
                                        _ = v20
                                    }
                                    var v21 int8 = 127
                                    var v22 float64 = float64(1.5)
                                    fmt.Println(v16, v1)
                                    delete(v3, fi9)
                                    _ = v21
                                    _ = v22
                                }
                                _ = v17
                            }
                            v23 := GenS4{F0: 2}
                            _ = v23
                        } else {
                            fmt.Println(fi15)
                            v1 = false
                        }
                        var v24 uint32 = 362
                        _ = v16
                        _ = v24
                    }
                    _ = fi15
                }
                v1 = (true && true)
                v25 := GenS5{F0: 255, F1: 1, F2: 18446744073709551615}
                _ = v25
            }
            var v26 uint32 = 65535
            _ = fi9
            _ = v26
        }
        var v27 uint64 = 449
        _ = v27
    } else {
        v28 := GenS3{F0: 4294967295, F1: 230}
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := GenS2{F0: 0, F1: 42, F2: ((1.5 / ((1e10 + 1.5) - (3.14159 * (float32(3.14159) + (0.0 * 3.14159))))) >= 1e10)}
        v31 := GenS1{F0: ("\n\t" + "\n\t")}
        _ = fi29
        _ = v30
        _ = v31
    }
    _ = len(v4)
    v32 := GenS4{F0: 798}
    v33 := v3[-9223372036854775808]
    var v34 int16 = 34
    fmt.Println(v34, v34)
    fmt.Println(v34)
    var v35 float64 = 3.14159
    switch 0 {
    case 128:
        var v36 int32 = 0
        _ = v36
    case -521:
        for fi37 := 0; fi37 < 3; fi37++ {
            v35 = -2.25
            delete(v3, (fi37 % (((fi37 << ((2) & 63))) | 1)))
            _ = fi37
        }
    }
    if v1 {
        if ((float64(1.5) + 0.0) <= v35) {
            v38 := v3[-647]
            _ = v38
        }
        var v39 float32 = -2.25
        var v40 int16 = v34
        _ = v39
        _ = v40
    } else {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
