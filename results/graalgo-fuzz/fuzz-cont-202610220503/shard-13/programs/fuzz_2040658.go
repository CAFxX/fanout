package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 bool; F1 uint }

func main() {
    var v1 uint16 = 24
    var v2 int32 = -2147483648
    var v3 uint16 = (v1 + (v1 % ((v1) | 1)))
    v4 := []int{-2147483648}
    fmt.Println(v3)
    var v5 uint = 0
    v3 = (v3 >> (((^(v5))) & 63))
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v5, v5)
        _ = fi6
    }
    var v7 uint8 = 42
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi8 := 0; fi8 < 18; fi8++ {
            if false {
                if true {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 string = "\n\t"
                        _ = fi9
                        _ = v10
                    }
                    if ((((-2.25 - -2.25) - 0.1) < (1.5 - (0.0 - (1.5 * -2.25)))) || (((3.14159 - 1.5) - 3.14159) != (float32(0.1) / (1.5 - (0.0 - float32(-2.25)))))) {
                        switch 912 {
                        case 768:
                            if (false || false) {
                                v4 = append(v4, fi8)
                                v4 = append(v4, fi8)
                            } else {
                                for fi11 := 0; fi11 < 10; fi11++ {
                                    var v12 int64 = -1
                                    var v13 float64 = 1e10
                                    v13 += v13
                                    _ = fi11
                                    _ = v12
                                    _ = v13
                                }
                            }
                        default:
                            for fi14 := 0; fi14 < 2; fi14++ {
                                fmt.Println(v1, v1, v5)
                                _ = fi14
                            }
                        }
                        v15 := GenS1{F0: (v7 >> ((v5) & 63)), F1: (v3 >> ((1) & 63))}
                        var v16 uint = (^(v5))
                        _ = v15
                        _ = v16
                    } else {
                        v7 = v7
                        v5 -= v5
                    }
                    v17 := GenS2{F0: "foo bar"}
                    _ = v17
                }
            }
            _ = fi8
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v4 = append(v4, (fi18 & (fi18 | fi18)))
            v19 := GenS3{F0: (^(v2)), F1: 1, F2: -128}
            fmt.Println(fi18)
            _ = fi18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: (v7 << ((2) & 63)), F1: v1}
            var v22 uint16 = v1
            _ = fi20
            _ = v21
            _ = v22
        }
        if false {
            var v23 int64 = 1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v23
        } else {
            v1 = v1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if (true && true) {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint16 = (v1 << (((v5 & 255)) & 63))
                var v26 int = fi24
                _ = fi24
                _ = v25
                _ = v26
            }
            v1 = v1
            var v27 uint16 = v1
            _ = v27
        }
    }
    if (1 <= 648) {
        v4 = append(v4, 4294967295)
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 uint = uint(v7)
            v1 -= (v3 - v1)
            _ = fi28
            _ = v29
        }
    }
    if true {
        v30 := GenS3{F0: 525, F1: 109, F2: 101}
        v1 = v3
        var v31 bool = (906 <= v5)
        _ = v30
        _ = v31
    }
    if (!(2147483648 != 0)) {
        v32 := GenS3{F0: v2, F1: 100, F2: 42}
        _ = v32
    }
    v33 := GenS4{F0: true, F1: (v5 | (v5 ^ v5))}
    v5 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v7
}
