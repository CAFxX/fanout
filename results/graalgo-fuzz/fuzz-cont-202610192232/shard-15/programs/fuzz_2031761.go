package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 int32 = 2147483647
    var v4 int16 = -32768
    v5 := make(map[int]int)
    var v6 uint32 = 985
    delete(v5, 0)
    var v7 string = "foo bar"
    var v8 uint = 307
    v7 = v7
    if (v2 && false) {
        v8 = (v8 % ((42) | 1))
        var v9 uint8 = 128
        delete(v1, 1)
        _ = v9
    } else {
        v4 = v4
        v10 := GenS1{F0: v8, F1: "a"}
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                if (18446744073709551615 <= 4294967295) {
                    v12 := GenS1{F0: (^(v8)), F1: ((v7 + v7) + (((v7 + (("" + (v7 + v7)) + "foo bar")) + v7) + ("\n\t" + "")))}
                    var v13 float64 = (((0.1 + -2.25) - (((float64(3.14159) * (float64(3.14159) - -2.25)) + -2.25) * ((3.14159 + 1.5) * (3.14159 - -2.25)))) + float64((1.5 / ((-2.25 / -2.25) / float32(0.1)))))
                    for fi14 := 0; fi14 < 13; fi14++ {
                        var v15 uint16 = 1
                        v13 += v13
                        if (v3 != v3) {
                            fmt.Println(fi11, v6, v8)
                            var v16 int64 = -1
                            if ((int64(fi14) > (v16 * (v16 & (v16 - v16)))) && (0 != (^(v3)))) {
                                v15 = 664
                            }
                            var v17 int = fi11
                            if false {
                                var v18 uint = 42
                                v8 += (v18 % ((18446744073709551615) | 1))
                                _ = v18
                            } else {
                                var v19 float32 = 1.5
                                if (v13 > 0.1) {
                                    switch v7 {
                                    case "foo bar":
                                        v2 = (false || v2)
                                    case "a":
                                        fmt.Println(v6, v13, v13)
                                    default:
                                        var v20 bool = ("a" != "")
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: (v8 ^ (v8 + v8)), F1: "foo bar"}
                                    v15 = (v15 >> (((v8 & 512)) & 63))
                                    _ = v21
                                } else {
                                    v22 := v1[fi11]
                                    var v23 float64 = ((v13 - (v13 + float64(((((v19 + v19) / v19) / 0.1) / float32(v13))))) * 0.1)
                                    var v24 float64 = v23
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                _ = v19
                            }
                            var v25 uint = (v8 & v8)
                            _ = v16
                            _ = v17
                            _ = v25
                        } else {
                            var v26 uint16 = (v15 & (v15 | 42))
                            _ = v26
                        }
                        _ = fi14
                        _ = v15
                    }
                    _ = v12
                    _ = v13
                }
                if false {
                    if v2 {
                        var v27 float32 = -2.25
                        _ = v27
                    } else {
                        var v28 string = ((((((((v7 + (((v7 + ("\n\t" + "foo bar")) + (v7 + v7)) + v7)) + v7) + "hello") + ((((v7 + v7) + "\n\t") + "\n\t") + (v7 + v7))) + v7) + v7) + "hello") + v7)
                        for fi29 := 0; fi29 < 7; fi29++ {
                            v28 = "\n\t"
                            _ = fi29
                        }
                        _ = v28
                    }
                    v30 := GenS2{F0: v2}
                    var v31 float64 = 0.1
                    _ = v30
                    _ = v31
                }
                var v32 uint8 = 0
                _ = v32
            }
            _ = fi11
        }
        _ = v10
    }
    fmt.Println(v3, v2, v7)
    if (v7 != ((("a" + v7) + ("\n\t" + v7)) + "a")) {
        fmt.Println(v8)
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            fmt.Println(v6, v2, fi33)
            var v34 uint64 = 324
            v35 := GenS1{F0: (v8 | (v8 >> ((573) & 63))), F1: v7}
            _ = fi33
            _ = v34
            _ = v35
        }
        v36 := GenS2{F0: ("hello" < v7)}
        fmt.Println(v8)
        _ = v36
    }
    delete(v1, 128)
    v37 := []int{}
    v37 = append(v37, 9223372036854775807)
    if (v4 < v4) {
        var v38 uint32 = v6
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS2{F0: (v3 != v3)}
            _ = fi39
            _ = v40
        }
        _ = len(v5)
        _ = v38
    } else {
        for fi41 := 0; fi41 < 14; fi41++ {
            fmt.Println(v7)
            var v42 uint32 = 0
            var v43 int8 = 127
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 string = v7
        var v45 bool = (!(622 != 829))
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
