package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 uint8 = 20
    var v2 float64 = 1e10
    var v3 uint = 388
    v4 := make(map[uint32]uint32)
    switch 431 {
    case 164:
        v2 = (0.1 - ((v2 + 1.5) - (v2 / 1.0)))
    }
    v5 := v4[uint32(v1)]
    v4[(v5 + (^(v5)))] = v5
    if (32767 == -9223372036854775808) {
        v1 -= (^(v1))
        for fi6 := 0; fi6 < 1; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                if (!(!true)) {
                    v5 = v5
                    var v8 int = 4294967295
                    v5 = v5
                    _ = v8
                }
                if (42 == -1) {
                    v1 = (v1 & 1)
                    if ((!(fi7 >= 127)) && true) {
                        v3 -= v3
                        if ((^(v1)) == v1) {
                            v1 = v1
                            if false {
                                if (42 == 1) {
                                    v9 := []int{2147483647, 42, (-(fi6)), (fi7 - fi7)}
                                    var v10 string = "\n\t"
                                    if false {
                                        var v11 int8 = 1
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = v11
                                    } else {
                                        v2 = v2
                                        v3 += 0
                                        var v12 uint64 = 660
                                        _ = v12
                                    }
                                    var v13 float64 = (-2.25 / v2)
                                    _ = v10
                                    _ = v13
                                    _ = v9
                                }
                            } else {
                                if false {
                                    v1 = uint8(v5)
                                }
                                v14 := GenS1{F0: -32768, F1: 156, F2: -128}
                                fmt.Println(v3)
                                _ = v14
                            }
                        } else {
                            v15 := GenS1{F0: -462, F1: 1, F2: 1}
                            _ = v15
                        }
                    }
                }
                _ = fi7
            }
            v5 -= (v5 * (v5 % ((v5) | 1)))
            switch -196 {
            case 42:
                var v16 uint8 = (^(v1))
                _ = v16
            case 1000:
                v17 := GenS1{F0: 1, F1: 1, F2: 42}
                _ = v17
            case 9223372036854775807:
                if (!false) {
                    var v18 float64 = v2
                    fmt.Println(v1, v2, v3)
                    _ = v18
                }
            default:
                if false {
                    fmt.Println(fi6, fi6)
                } else {
                    v19 := GenS1{F0: -32768, F1: 1, F2: 100}
                    if ("" > "hello") {
                        switch (("foo bar" + "\n\t") + "a") {
                        case "a":
                            if false {
                                var v20 int16 = 533
                                fmt.Println(v1, fi6, v20)
                                _ = v20
                            }
                        }
                    } else {
                        v1 -= v1
                        var v21 float32 = ((-2.25 + 3.14159) + -2.25)
                        v22 := GenS1{F0: 233, F1: 42, F2: -1}
                        _ = v21
                        _ = v22
                    }
                    _ = v19
                }
            }
            _ = fi6
        }
        var v23 bool = ((!(!(v1 < v1))) || true)
        _ = v23
    }
    var v24 int32 = -1
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := []int{}
        var v27 uint16 = 2
        v28 := GenS1{F0: 1000, F1: v27, F2: 0}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    if (!false) {
        v5 = (^(v5))
    } else {
        if false {
            v29 := GenS1{F0: -80, F1: 42, F2: 127}
            _ = v29
        } else {
            v1 = (v1 & (^(v1)))
        }
    }
    fmt.Println(v24, v5, v24)
    switch (((("foo bar" + ("" + ("\n\t" + ""))) + "a") + "\n\t") + ((("\n\t" + "") + "hello") + ("\n\t" + ""))) {
    case "hello":
        var v30 uint8 = (v1 ^ v1)
        _ = v30
    default:
        var v31 int32 = (v24 & (^(v24)))
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
