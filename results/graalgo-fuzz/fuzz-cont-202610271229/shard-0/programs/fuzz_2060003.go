package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int {
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 42
        fmt.Println(v7, v7, v9)
        v9 = (v9 ^ (^(v9)))
        _ = fi8
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 444
    v2 := make(map[string]int)
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    v5 := []int{32768, 9223372036854775807, -1, 9223372036854775807}
    var v6 uint64 = 42
    v10 := GenS1{F0: 1}
    v6 -= v6
    var v11 string = "\n\t"
    var v12 int8 = 100
    var v13 string = "foo bar"
    var v14 int16 = 927
    if (genFn0() != -1) {
        if (64 >= 593) {
            if (true || ((v6 << ((uint(v6)) & 63)) < v6)) {
                switch 42 {
                case 0:
                    v11 = v13
                case -880:
                    if false {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            fmt.Println(v14, fi15, v12)
                            if false {
                                v11 = v11
                                var v16 int16 = (v14 | (v14 ^ v14))
                                v13 = "hello"
                                _ = v16
                            } else {
                                if false {
                                    var v17 int8 = (^(v12))
                                    v1 -= 4294967295
                                    switch v13 {
                                    case "":
                                        _ = len(v5)
                                    }
                                    _ = v17
                                }
                                var v18 bool = false
                                _ = v18
                            }
                            v6 -= v6
                            _ = fi15
                        }
                        var v19 uint = 365
                        _ = v19
                    }
                }
                if (false || (!(-1 == 1))) {
                    v13 = v13
                    var v20 uint = 16
                    v14 += -1
                    _ = v20
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        delete(v4, (v1 >> ((18446744073709551615) & 63)))
        var v21 float32 = 0.0
        _ = v21
    }
    if true {
        v1 -= (v1 * 2)
        for fi22 := 0; fi22 < 1; fi22++ {
            switch v11 {
            case "a":
                if (128 == 127) {
                    v12 = 0
                    if ((uint32(v6) < 65535) && true) {
                        v23 := GenS1{F0: 128}
                        _ = v23
                    }
                    _ = len(v5)
                } else {
                    v13 = v13
                    v6 = (^(v6))
                }
            case "":
                if false {
                    fmt.Println(v13, v11, v13)
                    v6 += v6
                }
            default:
                fmt.Println(v6, v14)
            }
            fmt.Println(v14, v13, v13)
            fmt.Println(v14, fi22, v13)
            _ = fi22
        }
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            v4[v1] = fi24
            _ = fi24
        }
        v1 = v1
    }
    var v25 int16 = v14
    if false {
        _ = len(v2)
        var v26 float32 = float32(-2.25)
        _ = v26
    } else {
        for fi27 := 0; fi27 < 13; fi27++ {
            _ = len(v3)
            v28 := GenS1{F0: uint8(v6)}
            if (v12 != v12) {
                v29 := GenS1{F0: 42}
                for fi30 := 0; fi30 < 17; fi30++ {
                    if (!(true && true)) {
                        var v31 string = "a"
                        var v32 float32 = 1.5
                        _ = v31
                        _ = v32
                    }
                    _ = fi30
                }
                var v33 int = (fi27 | genFn0())
                _ = v29
                _ = v33
            }
            _ = fi27
            _ = v28
        }
    }
    v25 -= v25
    fmt.Println(v6, v1, v1)
    var v34 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
