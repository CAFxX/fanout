package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int; F1 bool; F2 uint }
type GenS5 struct { F0 uint64; F1 int64; F2 int8 }
type GenS6 struct { F0 int64; F1 string; F2 int64 }

func main() {
    var v1 bool = true
    var v2 int64 = 9223372036854775807
    var v3 uint = 4294967295
    var v4 int16 = 786
    var v5 uint32 = 800
    v3 -= (v3 ^ v3)
    var v6 float64 = (3.14159 + -2.25)
    v7 := []int{-1, -267}
    v6 = (1.5 - v6)
    _ = len(v7)
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 += v6
        v9 := GenS1{F0: true, F1: (fi8 * (fi8 / ((fi8) | 1))), F2: 18446744073709551615}
        switch (fi8 ^ fi8) {
        case 2147483647:
            if (0 < 2) {
                v4 -= (v4 % ((0) | 1))
                switch (-(fi8)) {
                case -438:
                    var v10 int32 = 492
                    _ = v10
                case 9223372036854775807:
                    if v1 {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v12 := GenS2{F0: (fi11 / (((fi8 - fi8)) | 1)), F1: false, F2: 1}
                            var v13 float64 = 0.0
                            var v14 string = "foo bar"
                            v4 = (v4 - v4)
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v4 = 42
                    }
                default:
                    v15 := GenS3{F0: 0, F1: 42}
                    _ = v15
                }
                var v16 uint64 = 2
                _ = v16
            }
        case -838:
            v5 = (v5 - v5)
        case -2147483648:
            var v17 int32 = 2147483647
            _ = v17
        default:
            v18 := GenS3{F0: 105, F1: -70}
            _ = v18
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v2, v3)
    for fi19 := 0; fi19 < 18; fi19++ {
        v5 -= 65535
        _ = fi19
    }
    if false {
        _ = len(v7)
        if (!true) {
            fmt.Println(v2)
            var v20 bool = (!((v5 ^ 42) == v5))
            v5 = v5
            _ = v20
        } else {
            for fi21 := 0; fi21 < 6; fi21++ {
                if true {
                    var v22 uint8 = 51
                    _ = v22
                }
                _ = fi21
            }
            if ("a" <= (("foo bar" + "foo bar") + ("hello" + (("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + ("\n\t" + "hello"))))) {
                var v23 float32 = (-2.25 / 1e10)
                _ = v23
            } else {
                v24 := GenS4{F0: -9223372036854775808, F1: v1, F2: v3}
                v4 += (v4 / ((42) | 1))
                switch "" {
                case "a":
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 uint8 = 127
                        var v27 bool = (238 < 580)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                default:
                    v5 = 0
                }
                _ = v24
            }
            v2 = 9223372036854775807
        }
        v28 := GenS5{F0: 4294967295, F1: v2, F2: 127}
        _ = v28
    } else {
        if v1 {
            v29 := GenS6{F0: -1, F1: "a", F2: -376}
            fmt.Println(v6)
            v4 += (v4 + 0)
            _ = v29
        } else {
            for fi30 := 0; fi30 < 2; fi30++ {
                v4 = (v4 & v4)
                var v31 uint8 = 128
                _ = fi30
                _ = v31
            }
            switch -522 {
            case 817:
                if (v6 > (v6 / float64(1.5))) {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        var v33 float64 = ((((v6 - v6) - float64((0.0 / 1e10))) * -2.25) - (v6 * v6))
                        _ = fi32
                        _ = v33
                    }
                } else {
                    switch -2147483648 {
                    case 297:
                        var v34 string = ""
                        _ = v34
                    default:
                        var v35 int32 = -2147483648
                        _ = v35
                    }
                }
            case 0:
                v7 = append(v7, 128)
            case 32768:
                var v36 int8 = -128
                _ = v36
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
