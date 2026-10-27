package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = -128
    var v2 bool = (2 > 417)
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 float64 = (-2.25 * (((0.1 * ((1e10 / (1e10 * -2.25)) - -2.25)) / 0.1) / -2.25))
    switch -9223372036854775808 {
    case 42:
        var v6 int16 = 0
        _ = v6
    default:
        if true {
            for fi7 := 0; fi7 < 7; fi7++ {
                v2 = false
                for fi8 := 0; fi8 < 18; fi8++ {
                    if (!(!(544 != 2))) {
                        if false {
                            switch fi7 {
                            case 127:
                                var v9 float32 = 3.14159
                                _ = v9
                            case 32767:
                                var v10 int = fi8
                                _ = v10
                            case 1:
                                var v11 int = 2147483647
                                _ = v11
                            default:
                                if (!v2) {
                                    v1 += v1
                                    if v2 {
                                        switch ("foo bar" + ("\n\t" + ("" + ("a" + "\n\t")))) {
                                        case "hello":
                                            v1 = v1
                                        case "\n\t":
                                            v12 := []int{(fi7 % ((0) | 1))}
                                            _ = v12
                                        }
                                        var v13 int32 = 0
                                        v2 = (224 < 23)
                                        _ = v13
                                    } else {
                                        _ = len(v4)
                                    }
                                }
                            }
                            v4["a"] = false
                            if true {
                                switch fi7 {
                                case -953:
                                    fmt.Println(v1, v5, v1)
                                default:
                                    var v14 int64 = -867
                                    _ = v14
                                }
                            }
                        }
                        var v15 int32 = 1
                        _ = v15
                    } else {
                        if v2 {
                            var v16 uint32 = 352
                            v17 := GenS1{F0: (v1 - (v1 | (^(v1)))), F1: (^(v16))}
                            v5 += v5
                            _ = v16
                            _ = v17
                        }
                    }
                    v18 := GenS1{F0: v1, F1: 1}
                    var v19 int = fi7
                    _ = fi8
                    _ = v18
                    _ = v19
                }
                _ = fi7
            }
            v20 := []int{-1, -87, 128}
            _ = v20
        }
    }
    v21 := GenS1{F0: v1, F1: 65535}
    fmt.Println(v2, v2)
    fmt.Println(v1)
    switch "" {
    case "a":
        v22 := GenS1{F0: v1, F1: 1}
        _ = v22
    case "foo bar":
        for fi23 := 0; fi23 < 10; fi23++ {
            if v2 {
                var v24 string = "hello"
                v5 += v5
                _ = v24
            }
            v25 := v3[72]
            _ = fi23
            _ = v25
        }
    case "":
        delete(v3, 395)
    default:
        switch 127 {
        case -1:
            var v26 uint32 = 2
            _ = v26
        default:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint32 = 4294967295
                if (2 <= 18446744073709551615) {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v30 := GenS1{F0: -1, F1: v28}
                        _ = fi29
                        _ = v30
                    }
                }
                var v31 uint8 = 255
                _ = fi27
                _ = v28
                _ = v31
            }
        }
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(v1, v2)
        var v33 uint32 = 4294967295
        var v34 uint8 = 63
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: v1, F1: 1}
    for fi36 := 0; fi36 < 6; fi36++ {
        switch "\n\t" {
        case "a":
            delete(v3, 42)
        default:
            v37 := GenS2{F0: v1}
            _ = v37
        }
        fmt.Println(v5, fi36, fi36)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
