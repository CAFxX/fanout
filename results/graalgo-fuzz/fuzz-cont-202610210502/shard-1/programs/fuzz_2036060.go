package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int = 1
    var v2 string = "\n\t"
    var v3 int32 = 791
    var v4 uint16 = 448
    v5 := make(map[int]uint32)
    var v6 int = v1
    v7 := GenS1{F0: 2, F1: -32768}
    v3 -= 42
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 14; fi9++ {
        v8 = v8
        _ = fi9
    }
    v10 := GenS2{F0: -520}
    var v11 int8 = 1
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 3; fi13++ {
        v4 -= v4
        var v14 string = v2
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := []int{(v1 << ((0) & 63)), (v1 * 0), (v6 | int(v3))}
        var v17 int8 = 100
        if (false || true) {
            switch (fi15 / ((-2147483648) | 1)) {
            case 50:
                v18 := GenS1{F0: 4294967295, F1: 42}
                _ = v18
            default:
                for fi19 := 0; fi19 < 15; fi19++ {
                    if ((fi15 << ((1) & 63)) <= 1) {
                        v4 = 0
                        v17 += (v11 / (((v17 - (^(v11)))) | 1))
                    } else {
                        v5[42] = (v8 | v8)
                    }
                    var v20 uint16 = 2
                    _ = fi19
                    _ = v20
                }
            }
            if true {
                var v21 bool = true
                v16 = append(v16, (v6 >> ((302) & 63)))
                _ = v21
            } else {
                if true {
                    var v22 int64 = -62
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 uint8 = 2
                        _ = len(v16)
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS1{F0: 212, F1: 1}
                    _ = v22
                    _ = v25
                } else {
                    var v26 uint8 = 0
                    _ = v26
                }
                switch ("\n\t" + v2) {
                case "\n\t":
                    v8 = 324
                case "hello":
                    for fi27 := 0; fi27 < 15; fi27++ {
                        switch fi27 {
                        case 128:
                            if (!false) {
                                if len(v16) > 0 {
                                    v16[0] = (fi15 | fi27)
                                }
                                fmt.Println(v8)
                            }
                        default:
                            var v28 float32 = -2.25
                            v1 = (v6 ^ fi27)
                            _ = v28
                        }
                        fmt.Println(v1, fi15)
                        _ = fi27
                    }
                default:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if (-32768 != 1000) {
                            var v30 string = "a"
                            _ = v30
                        } else {
                            var v31 float32 = 0.1
                            var v32 string = "hello"
                            _ = v31
                            _ = v32
                        }
                        v17 += (-(v17))
                        v6 = (fi15 + fi15)
                        _ = fi29
                    }
                }
                if true {
                    _ = len(v16)
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 uint64 = 131
                        _ = fi33
                        _ = v34
                    }
                    switch v2 {
                    case "hello":
                        v35 := GenS1{F0: 255, F1: -910}
                        _ = v35
                    case "foo bar":
                        var v36 string = ("\n\t" + ((((("a" + ("\n\t" + "hello")) + "") + "hello") + "") + (v2 + v2)))
                        _ = v36
                    default:
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            fmt.Println(v6, v3)
                        case "":
                            v2 = v2
                        }
                    }
                } else {
                    var v37 uint = uint(v4)
                    var v38 float64 = 1.5
                    var v39 float32 = 1e10
                    _ = v37
                    _ = v38
                    _ = v39
                }
            }
        } else {
            v17 = (v17 << ((0) & 63))
            var v40 int8 = (v11 + (^(v17)))
            _ = v40
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v41 uint = 42
    v42 := []int{v1, v6}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
