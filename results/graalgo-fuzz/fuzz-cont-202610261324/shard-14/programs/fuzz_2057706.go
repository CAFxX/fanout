package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0() uint {
    if (!(780 < 2)) {
        var v7 int16 = -1
        if false {
            switch "a" {
            case "":
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int = fi8
                    v9 = fi8
                    fmt.Println(v7, v7)
                    _ = fi8
                    _ = v9
                }
            case "\n\t":
                var v10 int32 = 0
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 uint16 = 813
                    v13 := GenS1{F0: 42, F1: 226}
                    for fi14 := 0; fi14 < 11; fi14++ {
                        if true {
                            if (true || (true || true)) {
                                v12 += (^(v12))
                            } else {
                                var v15 int8 = 42
                                v15 = v15
                                _ = v15
                            }
                        } else {
                            v7 -= v7
                        }
                        _ = fi14
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            fmt.Println(v7)
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 uint32 = 2
                _ = fi16
                _ = v17
            }
        }
        _ = v7
    }
    var v18 int64 = 42
    if false {
        fmt.Println(v18)
    } else {
        v18 += -358
        var v19 int64 = v18
        _ = v19
    }
    _ = v18
    return 570
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = 784
    var v3 int64 = 550
    var v4 string = "a"
    v5 := []string{("\n\t" + ((("a" + "\n\t") + (v4 + v4)) + v4)), "", v4, (((v4 + ((("\n\t" + v4) + v4) + v4)) + v4) + (v4 + v4))}
    v6 := make(map[string]int)
    v3 = v3
    switch ("foo bar" + v4) {
    case "":
        var v20 uint8 = uint8(v1)
        _ = v20
    default:
        v21 := GenS2{F0: 1}
        _ = v21
    }
    v22 := GenS2{F0: 4294967295}
    v1 = (^(v1))
    var v23 int = 4294967295
    var v24 uint = genFn0()
    v24 = (v24 + (v24 ^ 4294967295))
    v2 = (^(v2))
    fmt.Println(v4, v1)
    v4 = ("hello" + "")
    for fi25 := 0; fi25 < 11; fi25++ {
        if (v4 < v4) {
            var v26 int = (fi25 / ((v23) | 1))
            _ = v26
        }
        var v27 uint32 = 422
        _ = fi25
        _ = v27
    }
    v28 := GenS3{F0: 75, F1: (^(v1)), F2: v3}
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS1{F0: (-(v2)), F1: 18446744073709551615}
        if len(v5) > 0 {
            v5[0] = v4
        }
        _ = fi29
        _ = v30
    }
    var v31 int64 = 905
    if true {
        if (false && true) {
            v1 = v1
            v5 = append(v5, (((("" + v4) + v4) + ((v4 + "hello") + ((v4 + v4) + ((v4 + "\n\t") + v4)))) + v4))
        } else {
            for fi32 := 0; fi32 < 12; fi32++ {
                v33 := GenS2{F0: 42}
                fmt.Println(v31)
                v34 := v6[v4]
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    } else {
        v2 = (^(v2))
        var v35 uint = genFn0()
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
