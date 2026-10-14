package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint8, p2 uint8) uint16 {
    v7 := GenS1{F0: ("\n\t" == "\n\t")}
    p1 = p2
    for fi8 := 0; fi8 < 13; fi8++ {
        p1 = p1
        fmt.Println(fi8, p1, p1)
        var v9 string = ("\n\t" + "\n\t")
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 255
}

func main() {
    var v1 string = ("" + ("\n\t" + "foo bar"))
    v2 := make(map[string]bool)
    var v3 uint32 = 1
    var v4 uint16 = 42
    var v5 float64 = 3.14159
    v6 := []uint32{v3, (v3 / (((v3 << ((42) & 63))) | 1)), (v3 | (v3 * v3))}
    v1 = ("foo bar" + (v1 + "\n\t"))
    var v11 int32 = 42
    var v12 int64 = 330
    if false {
        _ = len(v6)
        if false {
            var v13 int8 = -128
            v6 = append(v6, v3)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                if false {
                    switch v1 {
                    case "a":
                        var v15 float64 = float64(((float32(((v5 * v5) * v5)) / (1e10 / 1.0)) - (1e10 / 3.14159)))
                        _ = v15
                    case "foo bar":
                        fmt.Println(v12, v1)
                    default:
                        v4 += v4
                    }
                    switch "foo bar" {
                    case "":
                        switch ("foo bar" + "foo bar") {
                        case "":
                            v16 := v2[("\n\t" + v1)]
                            _ = v16
                        }
                    default:
                        switch -1 {
                        case 32767:
                            fmt.Println(v5, v12)
                        case 0:
                            if (v11 <= v11) {
                                if ((fi14 ^ fi14) == fi14) {
                                    _ = len(v6)
                                } else {
                                    v17 := GenS1{F0: false}
                                    v5 = (0.1 / v5)
                                    v3 += 4294967295
                                    _ = v17
                                }
                            } else {
                                v18 := GenS1{F0: ((0.0 >= v5) && false)}
                                if false {
                                    v4 -= genFn0(false, 116, 60)
                                }
                                _ = v18
                            }
                        }
                    }
                }
                v19 := GenS1{F0: (v5 != v5)}
                _ = fi14
                _ = v19
            }
            v5 += v5
        }
    } else {
        switch 42 {
        case 32767:
            var v20 uint64 = 18446744073709551615
            _ = v20
        case 42:
            var v21 uint8 = 127
            _ = v21
        case 127:
            var v22 int8 = -128
            _ = v22
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v5 += ((((((v5 / v5) * 1.5) / 0.1) + (v5 - -2.25)) * v5) + v5)
    }
    var v23 float64 = 3.14159
    if (!false) {
        var v24 int8 = 118
        v25 := GenS1{F0: (4294967295 < 42)}
        if ((!(32768 < int(v11))) || (((^(v3)) < v3) && (false || (v11 == v11)))) {
            if len(v6) > 0 {
                v6[0] = 626
            }
        }
        _ = v24
        _ = v25
    } else {
        var v26 uint64 = 311
        var v27 int32 = v11
        if (!(((v11 % (((v27 * 0)) | 1)) != v11) || true)) {
            var v28 uint8 = 1
            fmt.Println(v28, v28)
            _ = v28
        } else {
            v1 = ("\n\t" + "foo bar")
        }
        _ = v26
        _ = v27
    }
    if true {
        v29 := GenS2{F0: (v1 + "")}
        _ = v29
    }
    if (1 >= 0) {
        switch (v1 + v1) {
        case "\n\t":
            for fi30 := 0; fi30 < 1; fi30++ {
                v31 := GenS1{F0: (false && false)}
                v32 := GenS2{F0: (v1 + "")}
                switch 32768 {
                case 0:
                    var v33 int8 = 0
                    _ = v33
                default:
                    var v34 uint32 = (v3 * 2147483647)
                    _ = v34
                }
                _ = fi30
                _ = v31
                _ = v32
            }
        case "":
            v12 += 10
        }
        for fi35 := 0; fi35 < 4; fi35++ {
            v6 = append(v6, v3)
            var v36 uint = 0
            _ = fi35
            _ = v36
        }
        var v37 uint32 = (v3 << ((164) & 63))
        _ = v37
    }
    delete(v2, ("a" + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
