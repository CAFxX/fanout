package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 65535
    var v4 int64 = 9223372036854775807
    var v5 string = "a"
    var v6 int32 = 1
    if true {
        v7 := v2[4294967295]
        fmt.Println(v5, v3)
        _ = v7
    }
    switch -102 {
    case 2147483647:
        switch -9223372036854775808 {
        case 4294967295:
            var v8 bool = false
            _ = v8
        case 0:
            var v9 string = v5
            _ = v9
        }
    case 128:
        if (1 != (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))) {
            v10 := []int{}
            for fi11 := 0; fi11 < 15; fi11++ {
                v3 -= v3
                v12 := GenS1{F0: -48, F1: (v4 * 1)}
                _ = fi11
                _ = v12
            }
            switch -977 {
            case -930:
                v10 = append(v10, 4294967295)
            case 1:
                var v13 uint8 = 42
                _ = v13
            case 583:
                fmt.Println(v5, v6, v3)
            default:
                if false {
                    v6 += v6
                    v14 := GenS2{F0: -183, F1: v4, F2: 42}
                    for fi15 := 0; fi15 < 11; fi15++ {
                        fmt.Println(v5)
                        _ = fi15
                    }
                    _ = v14
                } else {
                    if false {
                        var v16 uint64 = 0
                        v3 -= (v3 >> ((719) & 63))
                        _ = v16
                    } else {
                        if true {
                            var v17 float64 = 3.14159
                            for fi18 := 0; fi18 < 1; fi18++ {
                                var v19 float64 = v17
                                _ = len(v10)
                                fmt.Println(fi18, v5)
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        } else {
                            for fi20 := 0; fi20 < 14; fi20++ {
                                for fi21 := 0; fi21 < 20; fi21++ {
                                    if len(v10) > 0 {
                                        v10[0] = fi20
                                    }
                                    _ = fi21
                                }
                                v6 -= (v6 | v6)
                                _ = fi20
                            }
                            var v22 int8 = -1
                            delete(v1, "a")
                            _ = v22
                        }
                        var v23 string = "a"
                        _ = v23
                    }
                    v6 = v6
                }
            }
            _ = v10
        }
    default:
        v24 := []int{4294967295}
        _ = v24
    }
    if ((0.1 / 1.0) >= 3.14159) {
        var v25 int64 = v4
        fmt.Println(v25, v5, v4)
        var v26 int16 = 1000
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 int = fi28
                v29 -= -2147483648
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        fmt.Println(v3)
    }
    fmt.Println(v6, v5, v6)
    switch v5 {
    case "\n\t":
        v6 = (v6 % ((v6) | 1))
    case "foo bar":
        v30 := GenS2{F0: -2147483648, F1: v4, F2: (^(v4))}
        _ = v30
    }
    v31 := []int{300, 32768, 9223372036854775807}
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 8; fi33++ {
            v4 = (^(v4))
            v3 += (v3 | v3)
            fmt.Println(v5, v3, fi33)
            _ = fi33
        }
        _ = fi32
    }
    fmt.Println(v5)
    var v34 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
