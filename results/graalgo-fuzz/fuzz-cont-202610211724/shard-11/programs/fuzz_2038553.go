package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0() uint {
    var v6 int16 = -1
    for fi7 := 0; fi7 < 5; fi7++ {
        v6 += (-(v6))
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float32 = 1.5
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: false, F1: 2, F2: 18446744073709551615}
    v11 := GenS1{F0: false, F1: 161, F2: 42}
    v6 -= (v6 - (^(v6)))
    _ = v10
    _ = v11
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 65535
    v3 := make(map[int]string)
    var v4 uint = 255
    var v5 uint64 = 0
    fmt.Println(v4)
    if ((9223372036854775807 == 1) || (false || (3.14159 < -2.25))) {
        var v12 int8 = -128
        _ = v12
    }
    v4 = (v2 << ((genFn0()) & 63))
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            if false {
                if (!false) {
                    v15 := []int{(fi14 << ((v2) & 63)), (fi14 - fi14), (-(fi14))}
                    _ = v15
                } else {
                    var v16 uint64 = v5
                    if true {
                        var v17 uint8 = 2
                        _ = v17
                    }
                    _ = v16
                }
                if false {
                    var v18 bool = ((("a" + "") >= "") && ((!false) && (v4 >= 0)))
                    if (0 >= 1) {
                        v19 := []int{}
                        var v20 float64 = float64(float32(1.5))
                        _ = v19
                        _ = v20
                    } else {
                        switch "" {
                        case "foo bar":
                            if ((1.5 > (((((-2.25 / 1.5) + 1e10) * 0.0) / 1.0) - 1.5)) || (-1 < -32768)) {
                                var v21 uint32 = 439
                                switch "hello" {
                                case "hello":
                                    v22 := GenS1{F0: (!(868 > 255)), F1: 1, F2: v5}
                                    _ = v22
                                }
                                _ = v21
                            } else {
                                switch "" {
                                case "a":
                                    var v23 uint8 = 128
                                    _ = v23
                                case "":
                                    switch 475 {
                                    case 32768:
                                        v18 = (-2.25 != 1.5)
                                        fmt.Println(v18, v4)
                                    }
                                case "hello":
                                    var v24 int16 = 948
                                    _ = v24
                                }
                                v25 := GenS1{F0: ((32767 >= -791) || v18), F1: 199, F2: 0}
                                _ = v25
                            }
                        case "":
                            v26 := GenS1{F0: (94 > 255), F1: 0, F2: (v5 / (((^(v5))) | 1))}
                            _ = v26
                        }
                    }
                    _ = v18
                } else {
                    var v27 float32 = 3.14159
                    fmt.Println(v4)
                    _ = v27
                }
                v28 := GenS1{F0: false, F1: 172, F2: v5}
                v29 := GenS1{F0: (!true), F1: 125, F2: v5}
                fmt.Println(fi14, v5)
                v5 = 517
                _ = v28
                _ = v29
            }
            _ = fi14
        }
        v2 = v2
        _ = fi13
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 string = "hello"
        _ = fi30
        _ = v31
    }
    var v32 uint16 = 1
    _ = len(v1)
    v33 := GenS2{F0: (^(v2)), F1: 42}
    var v34 uint = (v2 * v4)
    v32 = v32
    v35 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
