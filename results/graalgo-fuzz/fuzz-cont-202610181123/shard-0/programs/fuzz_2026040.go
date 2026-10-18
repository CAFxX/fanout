package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0() int {
    if (18446744073709551615 != 1) {
        v4 := GenS1{F0: 2147483647, F1: 42}
        _ = v4
    }
    var v5 float32 = 0.0
    var v6 int32 = -1
    var v7 float64 = (0.0 + float64((((v5 * (3.14159 / ((v5 * v5) * -2.25))) + v5) + (v5 + 1e10))))
    switch -514 {
    case 2147483647:
        v5 = v5
    }
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1() uint16 {
    var v8 int16 = 902
    fmt.Println(v8)
    for fi9 := 0; fi9 < 12; fi9++ {
        if ((255 < 65535) || false) {
            v10 := GenS2{F0: 2, F1: 1, F2: -1}
            _ = v10
        }
        _ = fi9
    }
    _ = v8
    return 65535
}

func main() {
    var v1 int = 453
    var v2 int8 = 1
    var v3 uint8 = 241
    var v11 int64 = 42
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            v3 = 127
            v13 := []int{genFn0(), (v1 & fi12)}
            var v14 float64 = 1e10
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    v11 -= (v11 + (v11 / ((-9223372036854775808) | 1)))
    fmt.Println(v2, v2)
    fmt.Println(v1, v3, v1)
    v1 = (^(v1))
    if ((^(v3)) == (^(v3))) {
        if (598 < 2) {
            v1 = v1
        } else {
            v1 += genFn0()
        }
        var v15 int32 = 1
        v2 += int8(v15)
        _ = v15
    }
    v16 := GenS3{F0: false, F1: (-32768 <= 1)}
    var v17 string = ("" + "\n\t")
    var v18 uint64 = 195
    v1 = genFn0()
    v19 := []int{genFn0()}
    for fi20 := 0; fi20 < 19; fi20++ {
        switch v17 {
        case "hello":
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        case "foo bar":
            v21 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v21
        case "\n\t":
            v3 = v3
        default:
            if true {
                v22 := GenS3{F0: true, F1: false}
                _ = v22
            } else {
                if true {
                    v2 += v2
                    v23 := GenS1{F0: 212, F1: 65535}
                    _ = v23
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int16 = -83
                    v26 := GenS3{F0: false, F1: true}
                    switch (fi20 / (((fi20 << ((42) & 63))) | 1)) {
                    case 126:
                        var v27 float64 = (((1e10 / (1e10 + 1.5)) / 1.0) / 1.5)
                        _ = v27
                    case -2147483648:
                        if true {
                            var v28 uint32 = 4294967295
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            if (570 <= (v18 - v18)) {
                                var v29 uint16 = 65535
                                var v30 int16 = v25
                                _ = v29
                                _ = v30
                            }
                            var v31 uint = 1
                            _ = v28
                            _ = v31
                        } else {
                            if true {
                                fmt.Println(v11, v3)
                                v19 = append(v19, (fi24 - 127))
                            }
                            var v32 int16 = 0
                            _ = v32
                        }
                    default:
                        v11 += (v11 % (((-(v11))) | 1))
                    }
                    switch "foo bar" {
                    case "hello":
                        v18 -= v18
                    default:
                        if true {
                            v2 = (v2 | v2)
                            var v33 int = (fi20 / ((fi20) | 1))
                            v2 += v2
                            fmt.Println(v2)
                            _ = v33
                        }
                    }
                    fmt.Println(v18)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        if false {
            switch (^(fi20)) {
            case 127:
                var v34 uint64 = 1
                _ = v34
            case -489:
                if (!(!((42 < 32767) && (1e10 < (float64(float32((1e10 / 1.0))) + (float64((1.5 - (3.14159 / 1e10))) * 0.0)))))) {
                    v35 := GenS4{F0: (!true), F1: 0}
                    var v36 uint32 = 42
                    var v37 uint64 = v18
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    var v38 int8 = (^(v2))
                    var v39 int16 = 42
                    _ = v38
                    _ = v39
                }
            case 2147483647:
                fmt.Println(v1, v18, v17)
            default:
                v11 = 9223372036854775807
            }
        } else {
            var v40 float64 = -2.25
            _ = v40
        }
        _ = fi20
    }
    if len(v19) > 0 {
        v19[0] = genFn0()
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
