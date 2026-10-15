package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{633}
    var v3 uint32 = 665
    var v4 int = 1
    v5 := GenS1{F0: 536, F1: (!((true && (-128 != -28)) || (120 == 100))), F2: -897}
    v6 := GenS1{F0: 915, F1: true, F2: -9223372036854775808}
    v7 := GenS1{F0: 255, F1: (2 > v3), F2: v1}
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v1 += (v1 % (((v1 % ((-703) | 1))) | 1))
        _ = fi8
    }
    var v9 uint8 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        if (fi10 > v4) {
            for fi11 := 0; fi11 < 6; fi11++ {
                if true {
                    if len(v2) > 0 {
                        v2[0] = (^(v4))
                    }
                    if false {
                        switch -609 {
                        case 751:
                            if len(v2) > 0 {
                                v2[0] = (v4 >> ((918) & 63))
                            }
                        case 556:
                            var v12 string = "hello"
                            _ = v12
                        case 9223372036854775807:
                            v13 := GenS1{F0: 0, F1: (-128 <= 35), F2: (^(v1))}
                            _ = v13
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = -2147483648
                        }
                        switch -1 {
                        case -888:
                            var v14 uint8 = (v9 & v9)
                            _ = v14
                        case 127:
                            fmt.Println(v4)
                        }
                    }
                } else {
                    var v15 float64 = (-2.25 * (1.5 + 1e10))
                    v3 = 2
                    _ = v15
                }
                v3 -= v3
                var v16 uint64 = 115
                _ = fi11
                _ = v16
            }
            fmt.Println(fi10, v4)
        }
        _ = fi10
    }
    switch (v4 * (v4 << ((2) & 63))) {
    case -280:
        if (2147483647 <= 42) {
            v17 := GenS2{F0: "\n\t"}
            var v18 float32 = 0.1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v17
            _ = v18
        }
    case -991:
        v4 -= v4
    case 127:
        v1 += (v1 | v1)
    default:
        fmt.Println(v4)
    }
    v19 := GenS1{F0: 42, F1: false, F2: v1}
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 string = ((((("a" + "foo bar") + "hello") + "\n\t") + ("foo bar" + "hello")) + ("\n\t" + "foo bar"))
        var v22 int = v4
        v23 := GenS3{F0: 42}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if ((0.0 * -2.25) <= 3.14159) {
        for fi24 := 0; fi24 < 6; fi24++ {
            if (!false) {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 float32 = 1e10
                    _ = fi25
                    _ = v26
                }
                if (true && false) {
                    v27 := GenS4{F0: (v3 ^ (v3 | 65535))}
                    for fi28 := 0; fi28 < 9; fi28++ {
                        if ((((614 >= 18446744073709551615) || true) || true) || true) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v29 := GenS5{F0: v3, F1: 255, F2: 871}
                            _ = v29
                        } else {
                            v9 = 255
                        }
                        _ = fi28
                    }
                    _ = v27
                } else {
                    v1 += 0
                    var v30 bool = false
                    fmt.Println(v9, v1)
                    _ = v30
                }
            }
            v31 := GenS5{F0: (v3 | (v3 / ((v3) | 1))), F1: 1, F2: 1}
            _ = fi24
            _ = v31
        }
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(fi32)
            if len(v2) > 0 {
                v2[0] = v4
            }
            v9 = v9
            _ = fi32
        }
        if len(v2) > 0 {
            v2[0] = v4
        }
    } else {
        v3 += v3
        fmt.Println(v1, v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
