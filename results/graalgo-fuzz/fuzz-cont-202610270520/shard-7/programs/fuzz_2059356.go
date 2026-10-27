package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 float32 = 1e10
    v4 := []int{-9223372036854775808}
    v5 := []uint32{691, 147, 141}
    switch ("" + ((("\n\t" + "") + ("a" + "hello")) + "")) {
    case "\n\t":
        v3 -= ((0.1 + ((v3 / v3) / ((1.5 * v3) * 0.1))) - ((v3 / v3) + ((v3 - (v3 - v3)) * v3)))
    case "hello":
        v6 := GenS1{F0: 4294967295, F1: 148, F2: 2}
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            var v8 float32 = v3
            _ = v8
        } else {
            if true {
                if false {
                    switch "a" {
                    case "\n\t":
                        var v9 int64 = 791
                        _ = v9
                    case "":
                        _ = len(v4)
                    case "hello":
                        var v10 uint16 = 255
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 int32 = 42
                            if (32767 < 1000) {
                                v13 := GenS1{F0: 4294967295, F1: 1, F2: 255}
                                v12 -= (^(v12))
                                _ = v13
                            }
                            v12 = v12
                            _ = fi11
                            _ = v12
                        }
                    }
                }
                var v14 uint64 = 479
                _ = v14
            }
            if true {
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int64 = 370
                    delete(v2, fi7)
                    var v17 int8 = -120
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if (false || (-32768 != 456)) {
                    if true {
                        var v18 uint32 = 2
                        var v19 uint = 887
                        _ = v18
                        _ = v19
                    }
                    v1[fi7] = 1
                }
                for fi20 := 0; fi20 < 17; fi20++ {
                    v3 = ((1e10 + float32(-2.25)) / (float32((1e10 / -2.25)) + ((v3 - 0.1) / (((-2.25 / v3) + 1.5) / (1.5 + (v3 + v3))))))
                    var v21 uint8 = 2
                    fmt.Println(v3, fi20)
                    _ = fi20
                    _ = v21
                }
            } else {
                v3 = (1e10 + v3)
                var v22 bool = (-22 > 2147483647)
                _ = v22
            }
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 int32 = 736
                _ = fi23
                _ = v24
            }
        }
        v2[(fi7 >> ((621) & 63))] = (("hello" + ("" + "foo bar")) + ("" + ""))
        _ = fi7
    }
    v25 := GenS1{F0: 554, F1: 128, F2: 1}
    v5 = append(v5, 42)
    if false {
        fmt.Println(v3)
        switch "\n\t" {
        case "":
            var v26 bool = true
            _ = v26
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            switch ((("hello" + "\n\t") + ("foo bar" + "\n\t")) + "") {
            case "a":
                if (65535 >= 209) {
                    var v28 uint = 0
                    _ = v28
                } else {
                    var v29 uint32 = 678
                    if (53 >= 255) {
                        var v30 uint64 = 1
                        var v31 int32 = 42
                        _ = v30
                        _ = v31
                    }
                    _ = v29
                }
            case "\n\t":
                v3 = 0.1
            }
            v3 += float32(((1e10 * (0.0 * -2.25)) + (0.0 - float64(v3))))
            _ = fi27
        }
    }
    v32 := GenS2{F0: 419, F1: 1, F2: -5}
    v33 := GenS2{F0: 127, F1: -1, F2: 100}
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
