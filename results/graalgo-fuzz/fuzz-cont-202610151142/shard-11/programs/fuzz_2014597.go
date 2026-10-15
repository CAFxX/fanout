package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 bool = (!(!true))
    var v2 uint32 = 249
    var v3 uint32 = (v2 * v2)
    v4 := []uint32{2147483648}
    var v5 float32 = (((float32(0.0) * 3.14159) - 0.1) - (3.14159 / 1.0))
    var v6 uint = 65535
    var v7 int16 = -450
    v8 := GenS1{F0: "a", F1: -1, F2: 128}
    if (282 != 18446744073709551615) {
        var v9 uint16 = 2
        var v10 uint8 = 128
        v6 += 0
        _ = v10
        _ = v9
    }
    var v11 string = "hello"
    for fi12 := 0; fi12 < 11; fi12++ {
        switch ("a" + v11) {
        case "hello":
            if ((v7 - v7) >= 1) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 float32 = (v5 / v5)
                    if (v1 || v1) {
                        v11 = "a"
                        v5 = v14
                        var v15 int8 = int8(fi13)
                        _ = v15
                    } else {
                        var v16 uint = 65535
                        _ = v16
                    }
                    var v17 int32 = 1
                    v5 = (v5 / (v5 + (((((v5 + v5) * v5) / v14) * -2.25) - v5)))
                    _ = fi13
                    _ = v14
                    _ = v17
                }
                if v1 {
                    var v18 float32 = (1.5 * 0.1)
                    v2 = 65535
                    _ = v18
                }
            } else {
                v7 = v7
                if v1 {
                    v1 = v1
                    if ((v1 || v1) || v1) {
                        var v19 float32 = 1.5
                        _ = v19
                    }
                    if (v1 && true) {
                        v6 -= (v6 - v6)
                    } else {
                        var v20 uint = v6
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if v1 {
                            var v21 uint8 = 128
                            _ = v21
                        } else {
                            switch fi12 {
                            case 0:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 32767:
                                for fi22 := 0; fi22 < 15; fi22++ {
                                    v5 = v5
                                    v7 += (v7 + 42)
                                    switch "" {
                                    case "\n\t":
                                        v23 := GenS2{F0: (v3 - (v2 | (v3 ^ 42)))}
                                        _ = v23
                                    case "foo bar":
                                        _ = len(v4)
                                    default:
                                        if (!v1) {
                                            v24 := GenS1{F0: "foo bar", F1: -66, F2: 2}
                                            _ = v24
                                        }
                                    }
                                    _ = fi22
                                }
                            }
                            var v25 float64 = (1.5 - 0.0)
                            _ = v25
                        }
                        _ = v20
                    }
                }
                fmt.Println(v3)
            }
        }
        var v26 string = v11
        if len(v4) > 0 {
            v4[0] = v3
        }
        _ = fi12
        _ = v26
    }
    fmt.Println(v5, v6, v6)
    for fi27 := 0; fi27 < 20; fi27++ {
        if v1 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = v2
            }
            fmt.Println(v5, v7)
        } else {
            v7 = v7
            if len(v4) > 0 {
                v4[0] = v2
            }
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi27
    }
    v28 := GenS2{F0: v2}
    for fi29 := 0; fi29 < 1; fi29++ {
        switch v11 {
        case "a":
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int = (fi29 % ((fi29) | 1))
                v11 = (("" + (v11 + v11)) + v11)
                v3 = (^(v2))
                _ = fi30
                _ = v31
            }
        case "foo bar":
            v2 -= v2
        case "hello":
            for fi32 := 0; fi32 < 4; fi32++ {
                if ((0 >= 42) && v1) {
                    switch v11 {
                    case "":
                        fmt.Println(v2, v1, v6)
                    case "hello":
                        v7 -= 399
                    }
                } else {
                    if false {
                        fmt.Println(v6, v7)
                    }
                }
                v6 = v6
                v33 := GenS2{F0: 4294967295}
                _ = fi32
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 bool = v1
                var v36 float64 = (3.14159 + float64(-2.25))
                v37 := GenS1{F0: v11, F1: 42, F2: 128}
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        }
        var v38 uint64 = 4294967295
        var v39 int64 = 1
        _ = fi29
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 uint32 = 65535
        v42 := GenS2{F0: v41}
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 uint16 = 255
    if true {
        _ = len(v4)
        if len(v4) > 0 {
            v4[0] = (v2 - v2)
        }
        var v44 uint8 = 42
        _ = v44
    }
    v11 = v11
    switch v11 {
    case "foo bar":
        var v45 int32 = -2147483648
        _ = v45
    case "hello":
        fmt.Println(v11)
    case "\n\t":
        v6 = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
