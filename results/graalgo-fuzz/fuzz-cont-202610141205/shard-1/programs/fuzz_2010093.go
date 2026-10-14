package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
func genFn0() int32 {
    if false {
        var v6 uint64 = 42
        var v7 uint32 = 921
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            var v9 int16 = -1
            if true {
                v10 := GenS1{F0: 42, F1: 19}
                fmt.Println(fi8)
                v9 = (v9 & v9)
                _ = v10
            } else {
                v11 := GenS1{F0: 255, F1: -9223372036854775808}
                v12 := []int{fi8, fi8, fi8}
                _ = v11
                _ = v12
            }
            switch "" {
            case "foo bar":
                v13 := []int{51, fi8, (fi8 - 0)}
                _ = v13
            case "hello":
                if ((false && false) && ((!(18446744073709551615 > 2)) && (0.0 <= 0.1))) {
                    v9 = v9
                    var v14 float64 = (1e10 * (-2.25 - (0.1 - (-2.25 * (1e10 * 3.14159)))))
                    if true {
                        v15 := []int{int(v9)}
                        _ = v15
                    } else {
                        if (709 < 180) {
                            v16 := []int{(fi8 >> ((1) & 63))}
                            _ = v16
                        } else {
                            var v17 int8 = -50
                            fmt.Println(fi8, v17, fi8)
                            var v18 uint8 = 2
                            _ = v17
                            _ = v18
                        }
                        switch (fi8 << ((255) & 63)) {
                        case 128:
                            var v19 uint16 = 65535
                            _ = v19
                        case 4294967295:
                            switch "" {
                            case "hello":
                                var v20 float32 = 3.14159
                                _ = v20
                            }
                        default:
                            var v21 uint = 255
                            _ = v21
                        }
                    }
                    _ = v14
                }
            }
            _ = v9
        } else {
            var v22 int16 = 1000
            if false {
                v22 = v22
                var v23 bool = true
                v22 = (-(v22))
                _ = v23
            } else {
                switch ("" + ("foo bar" + "")) {
                case "hello":
                    v22 += (v22 - v22)
                case "\n\t":
                    fmt.Println(v22, v22)
                }
            }
            _ = v22
        }
        _ = fi8
    }
    var v24 uint64 = 0
    _ = v24
    return 276
}

func main() {
    var v1 int = 32767
    var v2 int32 = 104
    var v3 int16 = 0
    var v4 bool = ((1.5 / 1.0) <= (0.1 - 1.5))
    v5 := []uint32{2147483648}
    v2 += (v2 | v2)
    if (1 != (v3 | 812)) {
        var v25 int8 = 1
        var v26 uint8 = 42
        fmt.Println(v4, v1)
        _ = v25
        _ = v26
    } else {
        var v27 int32 = v2
        if (!((v3 << ((42) & 63)) != 706)) {
            v3 = v3
            v4 = false
            v28 := GenS1{F0: 382, F1: -780}
            _ = v28
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            _ = len(v5)
            _ = fi29
        }
        _ = v27
    }
    switch v1 {
    case 392:
        fmt.Println(v3)
    case -2147483648:
        fmt.Println(v1)
    default:
        fmt.Println(v4, v2, v4)
    }
    v30 := GenS2{F0: (524 > 2), F1: 127, F2: (236 != 2)}
    var v31 uint64 = 42
    v3 = -348
    v32 := GenS2{F0: (0 == 9223372036854775807), F1: -126, F2: (!(false && v4))}
    var v33 uint = 748
    v33 -= v33
    var v34 int64 = -616
    var v35 int16 = (v3 / ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
