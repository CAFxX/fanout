package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    var v1 string = "\n\t"
    var v2 string = "\n\t"
    var v3 int8 = -51
    for fi4 := 0; fi4 < 14; fi4++ {
        fmt.Println(v1, fi4)
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := []int{}
            var v7 uint8 = 128
            v8 := GenS1{F0: 1}
            if (662 == 82) {
                var v9 float32 = -2.25
                v3 = -1
                _ = v9
            }
            if (!false) {
                v1 = v2
                switch (fi4 >> ((4294967295) & 63)) {
                case 860:
                    v2 = ((("foo bar" + ("\n\t" + "\n\t")) + ((((v1 + ((v2 + v1) + v2)) + ("" + v1)) + ("hello" + (v1 + (v1 + v1)))) + "hello")) + v2)
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    v10 := GenS1{F0: 1}
    for fi11 := 0; fi11 < 11; fi11++ {
        if (0 >= 360) {
            var v12 uint64 = 42
            v12 = 662
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int64 = -481
                _ = fi13
                _ = v14
            }
            _ = v12
        } else {
            for fi15 := 0; fi15 < 12; fi15++ {
                if ((851 != 42) || ((77 < 42) || ((((3.14159 + 1.5) - float64((((0.1 / 1.0) + 0.1) + 3.14159))) * (0.0 * 1e10)) <= 0.0))) {
                    if ((!(!(!(!false)))) || true) {
                        if false {
                            if (0 == -259) {
                                if false {
                                    var v16 int = fi15
                                    switch fi15 {
                                    case -2147483648:
                                        var v17 int = (v16 - (fi15 * fi11))
                                        _ = v17
                                    default:
                                        var v18 uint16 = 694
                                        _ = v18
                                    }
                                    _ = v16
                                } else {
                                    if false {
                                        v3 -= (-(v3))
                                        var v19 uint32 = 0
                                        var v20 float64 = (0.0 + (0.0 + 0.1))
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v3 -= v3
                                        v21 := []int{-2147483648, fi11}
                                        var v22 int = fi15
                                        if (((fi15 | fi15) != v22) || true) {
                                            v2 = "hello"
                                        }
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 bool = false
                                    _ = v23
                                }
                                v24 := []int{-847, -514, (^(fi11))}
                                _ = v24
                            } else {
                                v25 := []int{fi11, (fi15 | (-(fi15)))}
                                _ = v25
                            }
                            v3 += (v3 ^ v3)
                        }
                        v26 := []int{}
                        _ = v26
                    }
                    var v27 int32 = 345
                    if false {
                        if ((fi15 & (fi11 - 128)) < 0) {
                            v1 = (v1 + ("" + v2))
                            var v28 int8 = v3
                            v29 := GenS1{F0: 2}
                            _ = v28
                            _ = v29
                        } else {
                            v30 := []int{-1, 127, 32768, fi15}
                            if len(v30) > 0 {
                                v30[0] = fi15
                            }
                            fmt.Println(v1)
                            _ = v30
                        }
                        v3 = (v3 - v3)
                        var v31 int8 = 85
                        _ = v31
                    } else {
                        v32 := GenS2{F0: 1, F1: 1}
                        if (true && (!true)) {
                            switch (fi11 + fi15) {
                            case -2147483648:
                                var v33 float64 = 1.5
                                _ = v33
                            case 9223372036854775807:
                                v1 = v2
                            }
                            v34 := GenS2{F0: 42, F1: 1}
                            _ = v34
                        }
                        v35 := []int{fi11, (fi15 << ((1) & 63)), fi15, -9223372036854775808}
                        v27 = (v27 * (v27 ^ v27))
                        _ = v32
                        _ = v35
                    }
                    v36 := []int{-40, (fi11 % ((fi11) | 1)), fi15}
                    v37 := GenS2{F0: 65535, F1: 1000}
                    _ = v27
                    _ = v36
                    _ = v37
                } else {
                    if false {
                        switch fi15 {
                        case 140:
                            var v38 int8 = (^(v3))
                            _ = v38
                        }
                        fmt.Println(v2)
                    }
                    fmt.Println(v2, v1)
                    var v39 uint64 = 392
                    _ = v39
                }
                var v40 int32 = -526
                _ = fi15
                _ = v40
            }
        }
        _ = fi11
    }
    fmt.Println(v3, v1)
    v1 = "hello"
    v41 := []int{}
    fmt.Println(v3, v3, v2)
    v41 = append(v41, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v41
}
