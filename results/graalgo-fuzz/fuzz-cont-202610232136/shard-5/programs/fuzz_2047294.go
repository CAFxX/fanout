package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int8; F1 uint; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 bool }

func main() {
    var v1 uint16 = 1
    var v2 int8 = -128
    var v3 uint16 = v1
    var v4 uint = 18446744073709551615
    var v5 uint8 = 187
    var v6 int64 = 9223372036854775807
    var v7 uint16 = 42
    var v8 bool = true
    var v9 uint32 = 2
    v10 := []int{-1, 128, 42}
    v7 = v3
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch "a" {
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            switch fi11 {
            case -9223372036854775808:
                if len(v10) > 0 {
                    v10[0] = fi11
                }
            }
            var v12 int64 = (v6 >> ((v4) & 63))
            if (int16(fi11) > 884) {
                v13 := GenS1{F0: (v7 ^ (v1 * 0)), F1: ("\n\t" + ""), F2: 0}
                var v14 bool = true
                for fi15 := 0; fi15 < 11; fi15++ {
                    switch "hello" {
                    case "hello":
                        v7 = 0
                        var v16 int8 = v2
                        fmt.Println(v14)
                        var v17 int16 = -386
                        v18 := GenS2{F0: v8, F1: v8, F2: 127}
                        _ = v16
                        _ = v17
                        _ = v18
                    case "\n\t":
                        v2 = v2
                    }
                    _ = fi15
                }
                _ = v13
                _ = v14
            } else {
                fmt.Println(v9)
                if ((int(v9) < fi11) && false) {
                    var v19 float64 = 3.14159
                    _ = v19
                } else {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if (v8 && (872 != 42)) {
                        var v20 float64 = 0.1
                        _ = v20
                    } else {
                        var v21 float64 = 0.0
                        var v22 uint8 = v5
                        _ = v21
                        _ = v22
                    }
                    v6 = (-(v6))
                }
                fmt.Println(v9)
            }
            _ = fi11
            _ = v12
        }
    case "foo bar":
        fmt.Println(v7, v4)
    case "hello":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS1{F0: v7, F1: ("a" + ("foo bar" + ((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "a"))), F2: 42}
            _ = fi23
            _ = v24
        }
    }
    var v25 uint8 = v5
    v25 = v5
    v9 = v9
    v26 := GenS1{F0: (v7 * v1), F1: "\n\t", F2: -32768}
    fmt.Println(v9)
    var v27 int32 = 740
    var v28 string = "a"
    for fi29 := 0; fi29 < 6; fi29++ {
        v3 -= v3
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 bool = (!false)
            if true {
                v32 := GenS2{F0: (!v8), F1: v8, F2: (fi30 >> ((18446744073709551615) & 63))}
                var v33 bool = true
                if len(v10) > 0 {
                    v10[0] = (fi29 % ((fi30) | 1))
                }
                _ = v32
                _ = v33
            }
            switch v28 {
            case "a":
                if len(v10) > 0 {
                    v10[0] = fi29
                }
            case "\n\t":
                v3 = 567
            case "":
                v25 = 255
            default:
                var v34 float64 = 1e10
                _ = v34
            }
            var v35 uint8 = (v25 - v5)
            var v36 uint32 = v9
            v25 -= (v25 + (v35 + 128))
            if (!(4294967295 >= 18446744073709551615)) {
                v31 = (18446744073709551615 >= 877)
                v10 = append(v10, (fi30 & 467))
                fmt.Println(v5, v2)
            } else {
                fmt.Println(v8, v2)
                fmt.Println(v28, v3, v35)
                switch (v28 + v28) {
                case "\n\t":
                    v28 = v28
                case "foo bar":
                    var v37 uint8 = 0
                    _ = v37
                default:
                    switch "foo bar" {
                    case "a":
                        if (((1.5 * 1e10) == 1e10) || (!v31)) {
                            _ = len(v10)
                            var v38 uint16 = v3
                            v39 := GenS3{F0: (v2 - int8(v5)), F1: v28}
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v9)
                            var v40 float32 = 0.1
                            var v41 int64 = 0
                            v41 += v41
                            var v42 int64 = (v6 | v6)
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    case "foo bar":
                        var v43 int16 = -1
                        _ = v43
                    case "":
                        v44 := GenS4{F0: v2, F1: 4294967295, F2: 127}
                        if true {
                            var v45 float32 = (0.0 * 0.1)
                            v46 := GenS1{F0: (v1 | uint16(v9)), F1: v28, F2: -32768}
                            _ = v45
                            _ = v46
                        }
                        if len(v10) > 0 {
                            v10[0] = 127
                        }
                        _ = v44
                    default:
                        if ((true && v8) && v31) {
                            v47 := GenS2{F0: (((float32(1.5) * 0.1) >= 0.1) || v31), F1: (uint32(v7) < v9), F2: 4294967295}
                            fmt.Println(v2, v3)
                            _ = v47
                        }
                    }
                    v48 := GenS5{F0: (v36 & v36), F1: 4294967295, F2: ((fi30 << ((v4) & 63)) >= (fi29 * fi30))}
                    v8 = (v7 < uint16(v4))
                    _ = v48
                }
            }
            v8 = (!true)
            _ = fi30
            _ = v31
            _ = v35
            _ = v36
        }
        _ = fi29
    }
    var v49 float32 = 0.1
    v10 = append(v10, 4294967295)
    for fi50 := 0; fi50 < 3; fi50++ {
        if true {
            switch (fi50 + (fi50 ^ fi50)) {
            case 42:
                if ((v8 || false) || ((v25 << ((255) & 63)) != v5)) {
                    for fi51 := 0; fi51 < 4; fi51++ {
                        _ = len(v10)
                        v7 = (^(v7))
                        v52 := GenS4{F0: (v2 + -127), F1: (^(v4)), F2: (v2 * (v2 >> ((841) & 63)))}
                        _ = fi51
                        _ = v52
                    }
                } else {
                    _ = len(v10)
                }
            }
            v28 = ("hello" + v28)
        }
        var v53 int32 = (v27 / ((2147483647) | 1))
        _ = fi50
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
