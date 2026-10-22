package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 127
    var v2 int16 = 1
    var v3 int32 = 2147483647
    var v4 int64 = 42
    if true {
        if (((!true) && false) && (4294967295 > 946)) {
            switch "" {
            case "hello":
                switch ((("hello" + ("\n\t" + "foo bar")) + "a") + "hello") {
                case "\n\t":
                    fmt.Println(v2, v2)
                }
            case "foo bar":
                var v5 uint32 = 4294967295
                _ = v5
            default:
                fmt.Println(v2)
            }
            v4 = int64(v3)
        }
        v6 := []int{-2147483648, -789, -2147483648, 0}
        _ = v6
    }
    switch -4 {
    case 985:
        v2 -= 32767
    case -248:
        fmt.Println(v3)
    case 4294967295:
        switch -1 {
        case -2147483648:
            v1 = (v1 ^ v1)
        case 129:
            v7 := GenS1{F0: 356, F1: 0, F2: v4}
            _ = v7
        case 32767:
            var v8 uint16 = 2
            _ = v8
        default:
            if (42 != 1) {
                for fi9 := 0; fi9 < 4; fi9++ {
                    switch fi9 {
                    case -1:
                        v2 = -32768
                    case 2147483647:
                        var v10 float64 = -2.25
                        _ = v10
                    case 128:
                        v1 -= (v1 - (v1 - (v1 | (^(v1)))))
                    default:
                        if ("\n\t" != "hello") {
                            if (0 <= 42) {
                                if false {
                                    if false {
                                        v11 := GenS1{F0: (fi9 - 490), F1: 18446744073709551615, F2: (v4 << ((4294967295) & 63))}
                                        if false {
                                            v1 += 128
                                        } else {
                                            fmt.Println(v4, v2)
                                        }
                                        if false {
                                            var v12 string = (("\n\t" + "") + "hello")
                                            _ = v12
                                        }
                                        _ = v11
                                    }
                                }
                                fmt.Println(fi9, v4, v3)
                                if false {
                                    v13 := GenS1{F0: 9223372036854775807, F1: 1, F2: v4}
                                    v14 := []int{}
                                    _ = v13
                                    _ = v14
                                }
                            }
                        }
                    }
                    var v15 uint8 = 2
                    var v16 uint16 = 910
                    _ = fi9
                    _ = v15
                    _ = v16
                }
            }
        }
    default:
        switch 365 {
        case 42:
            fmt.Println(v1)
        case 32767:
            v3 += v3
        default:
            var v17 uint64 = 660
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v3 -= (v3 & (^(v3)))
        v3 += v3
        fmt.Println(v1)
        _ = fi18
    }
    var v19 float32 = (0.0 / float32(1.5))
    if (((true && true) || (!(!false))) || true) {
        v4 = (v4 % (((v4 % ((-567) | 1))) | 1))
    }
    v20 := GenS1{F0: -807, F1: 4294967295, F2: (v4 << ((42) & 63))}
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 = 228
        _ = fi21
    }
    if false {
        for fi22 := 0; fi22 < 11; fi22++ {
            v23 := GenS2{F0: 960}
            v1 -= v1
            _ = fi22
            _ = v23
        }
        var v24 uint16 = 2
        _ = v24
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v19, v1)
            if (false && (!(((2147483648 < 2147483648) && (2 <= 617)) || false))) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v3 = 622
                    _ = fi26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 float64 = 1.5
                    v28 -= 0.0
                    v29 := GenS1{F0: fi27, F1: 18446744073709551615, F2: 9223372036854775807}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
            switch (fi25 + (^(fi25))) {
            case -194:
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint16 = 0
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi25
        }
        v32 := []int{74, 1, 42, -9223372036854775808}
        fmt.Println(v1, v4)
        _ = v32
    }
    fmt.Println(v19)
    var v33 int8 = 127
    v34 := []int{2147483647, -1}
    if true {
        var v35 int16 = v2
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
