package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    var v2 float64 = (-2.25 * 0.1)
    var v3 int64 = -132
    v4 := []int{4294967295, 42}
    v2 -= 3.14159
    var v5 uint = 255
    v3 += v3
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int8 = 127
        var v9 float64 = 3.14159
        fmt.Println(fi7, v2, v3)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint = (v5 & v5)
    var v11 int32 = 1
    switch ("" + ("\n\t" + "foo bar")) {
    case "hello":
        if (v2 == v2) {
            v3 += (v3 | v3)
            switch 0 {
            case -342:
                switch 32767 {
                case -2147483648:
                    v11 = -1
                case 4294967295:
                    var v12 float64 = 0.1
                    _ = v12
                case 871:
                    if false {
                        fmt.Println(v5, v10)
                    } else {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            if len(v4) > 0 {
                                v4[0] = (fi13 ^ fi13)
                            }
                            v3 += 0
                            switch -9223372036854775808 {
                            case 32767:
                                switch (("a" + "a") + ("" + ("" + ("hello" + ("" + (("a" + "hello") + "hello")))))) {
                                case "hello":
                                    var v14 int64 = (v3 | v3)
                                    _ = v14
                                case "\n\t":
                                    for fi15 := 0; fi15 < 20; fi15++ {
                                        if (!true) {
                                            var v16 int = (-(fi15))
                                            var v17 uint32 = uint32(v1)
                                            _ = v16
                                            _ = v17
                                        } else {
                                            var v18 float32 = 0.0
                                            if false {
                                                var v19 float32 = (v6 / (v18 / v18))
                                                _ = v19
                                            }
                                            _ = v18
                                        }
                                        _ = fi15
                                    }
                                }
                            case 863:
                                v3 += (v3 % ((0) | 1))
                            case 1:
                                v20 := GenS1{F0: 128, F1: v3}
                                _ = v20
                            }
                            _ = fi13
                        }
                        v21 := GenS1{F0: 141, F1: (v3 - -1)}
                        _ = v21
                    }
                }
            }
            var v22 int = 9223372036854775807
            _ = v22
        }
    default:
        switch 32768 {
        case -23:
            var v23 int64 = (-(v3))
            _ = v23
        case -104:
            for fi24 := 0; fi24 < 18; fi24++ {
                v6 = v6
                var v25 string = (("a" + "hello") + "\n\t")
                v2 = 1.5
                _ = fi24
                _ = v25
            }
        case 737:
            v6 = -2.25
        default:
            var v26 int64 = v3
            _ = v26
        }
    }
    var v27 string = "\n\t"
    if (v3 < v3) {
        var v28 bool = false
        _ = v28
    }
    var v29 bool = (!true)
    if ((v10 << ((v5) & 63)) > v10) {
        var v30 bool = v29
        fmt.Println(v2, v11)
        _ = v30
    } else {
        v31 := GenS2{F0: v11}
        v6 = 1e10
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
