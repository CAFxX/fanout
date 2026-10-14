package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
func genFn0() int {
    if (false || (697 != 2)) {
        if (false || ((0.0 >= 0.0) && ((65535 == 4294967295) && true))) {
            if true {
                switch -237 {
                case 2147483647:
                    var v5 bool = ("foo bar" > "")
                    _ = v5
                case -789:
                    switch ((("foo bar" + (("a" + "hello") + "foo bar")) + "hello") + "\n\t") {
                    case "\n\t":
                        if (-1 != 2147483647) {
                            for fi6 := 0; fi6 < 13; fi6++ {
                                if (1 != 1) {
                                    var v7 uint8 = 127
                                    _ = v7
                                } else {
                                    if (-128 != 58) {
                                        v8 := []int{}
                                        var v9 uint32 = 42
                                        _ = v8
                                        _ = v9
                                    } else {
                                        fmt.Println(fi6)
                                        v10 := GenS1{F0: 32767, F1: 32768, F2: 42}
                                        v11 := GenS1{F0: 65535, F1: fi6, F2: 127}
                                        _ = v10
                                        _ = v11
                                    }
                                }
                                var v12 float32 = -2.25
                                v13 := GenS1{F0: 42, F1: (-(fi6)), F2: 0}
                                _ = fi6
                                _ = v12
                                _ = v13
                            }
                            var v14 uint8 = 1
                            for fi15 := 0; fi15 < 12; fi15++ {
                                switch fi15 {
                                case 1:
                                    if true {
                                        v16 := GenS1{F0: 378, F1: (^(fi15)), F2: -31}
                                        _ = v16
                                    } else {
                                        fmt.Println(v14, v14)
                                    }
                                }
                                _ = fi15
                            }
                            _ = v14
                        } else {
                            var v17 bool = true
                            _ = v17
                        }
                    case "hello":
                        v18 := GenS1{F0: 255, F1: 32768, F2: 100}
                        _ = v18
                    case "a":
                        var v19 float64 = 1e10
                        _ = v19
                    }
                default:
                    for fi20 := 0; fi20 < 18; fi20++ {
                        v21 := []int{(fi20 & 594)}
                        if (2 >= 32767) {
                            if (!((-(fi20)) <= fi20)) {
                                fmt.Println(fi20)
                            }
                            _ = len(v21)
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = fi20
                        _ = v21
                    }
                }
            }
        }
        var v22 uint16 = 134
        _ = v22
    }
    var v23 int32 = -1
    var v24 uint64 = 2
    var v25 string = ""
    v26 := GenS1{F0: 32767, F1: 0, F2: -55}
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    return 127
}

func main() {
    var v1 uint16 = 475
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v27 int16 = -337
    v4[575] = 2147483648
    switch ((("\n\t" + "foo bar") + "") + (("a" + "a") + "a")) {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v28 int16 = -32768
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 float64 = float64((-2.25 / 0.1))
        var v31 float32 = float32((-2.25 / ((v30 + v30) - (1.5 * (v30 + (1e10 - ((float64(1e10) + (v30 - float64((3.14159 * (float32(v30) + 1.5))))) - 1e10)))))))
        v32 := GenS1{F0: (v1 % ((v1) | 1)), F1: 235, F2: -128}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 uint32 = 457
    fmt.Println(v27, v28)
    for fi34 := 0; fi34 < 20; fi34++ {
        delete(v4, fi34)
        var v35 bool = ((v33 - (v33 ^ v33)) <= v33)
        var v36 int8 = 0
        _ = fi34
        _ = v35
        _ = v36
    }
    v2 = append(v2, uint32(v1))
    v27 = (v27 % ((1) | 1))
    fmt.Println(v27)
    var v37 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v37
    _ = v4
}
