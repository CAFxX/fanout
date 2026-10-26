package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0() uint32 {
    var v4 uint = 1
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v4, fi5)
        _ = fi5
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "a":
        var v6 uint32 = 4294967295
        _ = v6
    case "":
        if (!((v4 >= 67) || (!false))) {
            if (0 >= 4294967295) {
                switch 42 {
                case 1:
                    v4 = 429
                default:
                    v4 = 42
                }
                v7 := []int{1, -441, 2147483647}
                var v8 int16 = 297
                _ = v7
                _ = v8
            } else {
                if ((1 > 158) && ((v4 % ((v4) | 1)) > (^(v4)))) {
                    v4 = 4294967295
                }
                switch ("a" + ("a" + ((("hello" + "a") + ("foo bar" + ("foo bar" + "a"))) + ""))) {
                case "foo bar":
                    v4 += v4
                case "":
                    v9 := GenS1{F0: 12, F1: 42, F2: 128}
                    _ = v9
                }
                var v10 uint32 = 774
                _ = v10
            }
            v11 := GenS1{F0: 128, F1: 18446744073709551615, F2: 2}
            _ = v11
        } else {
            var v12 float64 = (1.5 * 1.5)
            v12 = v12
            v4 = (^(v4))
            _ = v12
        }
    case "foo bar":
        v4 = v4
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v4 += v4
            v4 = (v4 & v4)
            v15 := GenS1{F0: 2, F1: 2, F2: 127}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            switch (fi16 ^ fi16) {
            case 9223372036854775807:
                if (0.0 == 1.5) {
                    v4 = 0
                    var v17 float64 = float64(0.1)
                    _ = v17
                }
            }
            var v18 uint = v4
            v19 := []int{-395}
            _ = fi16
            _ = v18
            _ = v19
        }
        _ = fi13
    }
    _ = v4
    return 875
}

func main() {
    var v1 bool = ((!true) && true)
    var v2 int8 = -36
    v3 := make(map[uint32]uint32)
    var v20 uint32 = 345
    v1 = false
    v1 = (!false)
    _ = len(v3)
    var v21 uint8 = 63
    var v22 uint32 = v20
    if (!((32767 == 889) || (4294967295 == 42))) {
        if ((110 <= 0) || (1000 > 1)) {
            fmt.Println(v2, v22, v20)
        }
    } else {
        if ("" >= ("a" + ("hello" + ("" + ("" + ""))))) {
            v21 = v21
        } else {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                switch "foo bar" {
                case "":
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if (!(0 == -9223372036854775808)) {
                            v20 -= v22
                        } else {
                            if false {
                                v24 := GenS1{F0: 1, F1: 18446744073709551615, F2: v21}
                                var v25 uint64 = 665
                                v26 := []int{}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(v20)
                            }
                            var v27 int32 = 2147483647
                            _ = v27
                        }
                        _ = fi23
                    }
                case "hello":
                    var v28 float32 = float32(1e10)
                    _ = v28
                default:
                    v29 := GenS2{F0: 42}
                    _ = v29
                }
            }
            var v30 float64 = 1e10
            v31 := []int{}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS1{F0: (^(v21)), F1: 2, F2: (^(v21))}
            var v34 bool = (!((v21 * (v21 ^ (v21 >> ((1) & 63)))) < 0))
            var v35 int = fi32
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: (v21 + (^(v21))), F1: -2147483648}
        _ = v36
    }
    switch -1 {
    case 2147483647:
        switch ("a" + "\n\t") {
        case "foo bar":
            _ = len(v3)
        default:
            var v37 int32 = 1
            _ = v37
        }
    case 458:
        var v38 bool = v1
        _ = v38
    default:
        v21 = (v21 + (^(v21)))
    }
    v20 += (v20 & (v22 - genFn0()))
    fmt.Println(v21)
    var v39 float64 = 0.0
    for fi40 := 0; fi40 < 14; fi40++ {
        v21 -= v21
        fmt.Println(v1, v1)
        _ = fi40
    }
    switch ("a" + "foo bar") {
    case "\n\t":
        v1 = false
    default:
        fmt.Println(v39)
    }
    fmt.Println(v2, v21)
    v41 := GenS1{F0: 67, F1: 42, F2: 1}
    v42 := GenS2{F0: 85}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v41
    _ = v42
}
