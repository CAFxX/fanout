package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint; F2 string }
type GenS4 struct { F0 int8; F1 string; F2 int32 }
type GenS5 struct { F0 int; F1 int32 }

func main() {
    v1 := []uint32{65535, 2147483648}
    v2 := make(map[int]uint32)
    v3 := []int{-1, 1}
    var v4 uint64 = 297
    var v5 int64 = 318
    var v6 int64 = v5
    v7 := GenS1{F0: 581, F1: 176, F2: (true || ((0.1 < ((float64((3.14159 / (-2.25 - 1.5))) / -2.25) * (-2.25 + 1e10))) || (-1 == 0)))}
    v8 := GenS1{F0: 590, F1: 42, F2: false}
    if true {
        v4 = 4294967295
        var v9 uint64 = 945
        _ = len(v2)
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if len(v3) > 0 {
            v3[0] = (fi10 / (((fi10 - (^(fi10)))) | 1))
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if true {
            v12 := v2[-351]
            _ = len(v2)
            _ = v12
        } else {
            fmt.Println(v6, v6, v5)
        }
        var v13 float64 = float64(1.5)
        v6 += (v6 - (v5 >> ((2) & 63)))
        _ = fi11
        _ = v13
    }
    v5 -= v5
    switch -2147483648 {
    case 128:
        switch 9223372036854775807 {
        case 2147483647:
            switch (("" + "\n\t") + ("\n\t" + ("a" + "hello"))) {
            case "a":
                var v14 uint64 = 18446744073709551615
                _ = v14
            case "":
                _ = len(v1)
            case "hello":
                v15 := GenS2{F0: 0}
                _ = v15
            }
        default:
            if true {
                if (!false) {
                    v5 += (-(v5))
                    var v16 float64 = (0.1 * -2.25)
                    var v17 string = "a"
                    _ = v16
                    _ = v17
                } else {
                    if (!((true || (0 == 32768)) || true)) {
                        v18 := GenS2{F0: 0}
                        _ = v18
                    }
                    var v19 string = ("\n\t" + ("\n\t" + "\n\t"))
                    _ = v19
                }
            } else {
                v4 = 2
                switch "foo bar" {
                case "":
                    var v20 int16 = -32768
                    _ = v20
                default:
                    if false {
                        if false {
                            switch 32767 {
                            case -122:
                                if (0.0 == (((0.0 * -2.25) * 1.5) + 1.5)) {
                                    v5 = (v6 | (v6 - (v6 & v5)))
                                    if (!(!(!(-22 <= 0)))) {
                                        var v21 uint = 257
                                        _ = v21
                                    } else {
                                        v22 := GenS3{F0: 289, F1: 65535, F2: "hello"}
                                        v23 := GenS4{F0: -128, F1: ("\n\t" + ""), F2: int32(v4)}
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                            var v24 uint64 = v4
                            v25 := GenS1{F0: 15, F1: 2, F2: ((((float64(float32(0.0)) / 1.0) * (float64(3.14159) + ((-2.25 / 3.14159) - 1.5))) != 0.0) && false)}
                            _ = v24
                            _ = v25
                        } else {
                            if ((0 == -7) || true) {
                                switch 5 {
                                case 42:
                                    if (-1 > 110) {
                                        v26 := GenS5{F0: 727, F1: 694}
                                        var v27 bool = true
                                        _ = v26
                                        _ = v27
                                    } else {
                                        if ((uint(v4) <= 134) && (false && true)) {
                                            v6 = 9223372036854775807
                                        }
                                    }
                                case -695:
                                    for fi28 := 0; fi28 < 18; fi28++ {
                                        _ = len(v2)
                                        for fi29 := 0; fi29 < 10; fi29++ {
                                            fmt.Println(v4)
                                            v6 -= -9223372036854775808
                                            _ = fi29
                                        }
                                        v5 = (v5 * v6)
                                        _ = fi28
                                    }
                                }
                                if (true || false) {
                                    v4 -= v4
                                    var v30 int = -2147483648
                                    _ = v30
                                } else {
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        delete(v2, fi31)
                                        v6 = (v5 + (v6 + (v5 >> ((498) & 63))))
                                        _ = fi31
                                    }
                                    var v32 uint = 0
                                    _ = v32
                                }
                            }
                        }
                        v4 -= v4
                    }
                }
            }
        }
    case 333:
        v33 := GenS2{F0: 385}
        _ = v33
    }
    v5 = (v6 * v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
