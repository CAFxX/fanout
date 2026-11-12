package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }

func main() {
    var v1 uint16 = 425
    var v2 int = 4294967295
    var v3 uint32 = 42
    var v4 string = ((("" + ("\n\t" + "\n\t")) + "foo bar") + "")
    v5 := make(map[uint32]int)
    v6 := []int{v2}
    if true {
        fmt.Println(v4, v3)
        var v7 int16 = 1
        var v8 uint64 = 4294967295
        _ = v7
        _ = v8
    }
    if true {
        if true {
            v3 = (v3 / ((42) | 1))
            v9 := GenS1{F0: 42, F1: "foo bar", F2: 42}
            _ = v9
        } else {
            v10 := GenS1{F0: 135, F1: v4, F2: -1}
            if (true && (("a" + v4) != (v4 + v4))) {
                for fi11 := 0; fi11 < 20; fi11++ {
                    fmt.Println(v2)
                    _ = fi11
                }
                var v12 int = 8
                if true {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        if false {
                            fmt.Println(v4, fi13)
                            var v14 uint16 = v1
                            _ = v14
                        } else {
                            var v15 int8 = -128
                            fmt.Println(v4, v4)
                            _ = v15
                        }
                        if true {
                            fmt.Println(fi13, v4)
                            v4 = v4
                            for fi16 := 0; fi16 < 1; fi16++ {
                                v17 := GenS1{F0: 0, F1: v4, F2: 107}
                                if (true || false) {
                                    v6 = append(v6, fi13)
                                } else {
                                    v18 := GenS1{F0: -1, F1: v4, F2: -7}
                                    _ = v18
                                }
                                var v19 float32 = ((1e10 - 1.5) + 0.1)
                                switch ("foo bar" + ((("hello" + "hello") + v4) + (v4 + v4))) {
                                case "a":
                                    fmt.Println(v1)
                                case "hello":
                                    v1 -= (v1 >> ((232) & 63))
                                default:
                                    fmt.Println(v3)
                                }
                                v5[v3] = -419
                                _ = fi16
                                _ = v17
                                _ = v19
                            }
                        }
                        for fi20 := 0; fi20 < 18; fi20++ {
                            var v21 uint32 = v3
                            if len(v6) > 0 {
                                v6[0] = (fi13 ^ 1)
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = fi13
                    }
                    v22 := GenS1{F0: 0, F1: v4, F2: 42}
                    var v23 float64 = 1e10
                    _ = v22
                    _ = v23
                }
                _ = v12
            } else {
                if ((v2 & v2) > v2) {
                    v4 = (v4 + v4)
                    fmt.Println(v2, v2, v1)
                }
                v4 = v4
            }
            v1 -= v1
            _ = v10
        }
        v4 = ((("hello" + v4) + v4) + v4)
        var v24 int = v2
        _ = v24
    }
    v4 = v4
    v2 = v2
    v25 := GenS1{F0: 124, F1: (v4 + (((v4 + (v4 + ("" + "hello"))) + v4) + "\n\t")), F2: 1}
    for fi26 := 0; fi26 < 16; fi26++ {
        switch (v4 + (v4 + "")) {
        case "":
            v4 = "foo bar"
        case "a":
            fmt.Println(v3, v4, fi26)
        case "foo bar":
            v3 += (v3 & 0)
        }
        v3 = (^(v3))
        v3 += v3
        _ = fi26
    }
    var v27 string = ((v4 + (v4 + ("hello" + v4))) + "a")
    if false {
        switch v2 {
        case -2147483648:
            if (65535 == 4294967295) {
                v5[(v3 / (((v3 / (((v3 + (^(v3)))) | 1))) | 1))] = 0
                var v28 string = v4
                _ = v28
            } else {
                v3 = (^(v3))
                v2 = v2
                v4 = "a"
            }
        default:
            v27 = v4
        }
        var v29 uint32 = 828
        _ = v29
    }
    v30 := GenS1{F0: 0, F1: v4, F2: -93}
    v31 := GenS1{F0: 331, F1: v27, F2: 127}
    v27 = (v27 + "a")
    v1 += (v1 * 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
