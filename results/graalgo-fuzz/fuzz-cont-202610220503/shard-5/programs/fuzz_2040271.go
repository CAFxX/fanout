package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 uint }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{(v1 % (((v1 | (v1 << ((42) & 63)))) | 1))}
    v3 := []uint32{v1, v1, 2147483647}
    v4 := []string{"", "a"}
    var v5 int16 = 1000
    var v6 uint64 = 42
    switch ("" + "foo bar") {
    case "hello":
        if (255 <= 0) {
            var v7 int = 32768
            _ = v7
        } else {
            v8 := GenS1{F0: (^(v6)), F1: -2147483648}
            v6 = v6
            _ = v8
        }
    case "\n\t":
        v9 := GenS2{F0: 0, F1: -598, F2: -264}
        _ = v9
    case "foo bar":
        if (true && true) {
            if false {
                v6 += v6
                switch ("a" + "\n\t") {
                case "foo bar":
                    v10 := GenS3{F0: 0}
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                if false {
                    var v12 uint8 = 42
                    _ = v12
                } else {
                    fmt.Println(v6, v6, fi11)
                }
                _ = len(v3)
                _ = fi11
            }
            var v13 string = "hello"
            _ = v13
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS3{F0: v6}
                _ = fi14
                _ = v15
            }
            if (666 < 127) {
                v1 = 961
                if true {
                    var v16 uint64 = v6
                    v1 = 2147483648
                    _ = v16
                } else {
                    v17 := GenS1{F0: (v6 * (v6 << ((18446744073709551615) & 63))), F1: 42}
                    _ = v17
                }
            } else {
                v18 := GenS1{F0: (v6 | v6), F1: 212}
                _ = v18
            }
            v1 += (v1 + (^(v1)))
        }
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint32 = 732
            if (3.14159 >= -2.25) {
                var v21 float64 = (0.0 + -2.25)
                _ = v21
            } else {
                v3 = append(v3, v20)
                if (!false) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v22 := GenS3{F0: (v6 / ((42) | 1))}
                _ = v22
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v20
        }
    }
    if true {
        v6 -= (v6 * (v6 ^ v6))
    }
    v5 = 311
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if true {
            v5 += (v5 - v5)
            v5 = (^(v5))
            var v23 int16 = v5
            _ = v23
        } else {
            v24 := GenS4{F0: int16(v1), F1: "foo bar", F2: 18446744073709551615}
            _ = v24
        }
        switch ("" + "a") {
        case "":
            switch "" {
            case "hello":
                v6 -= v6
            }
        default:
            v25 := GenS2{F0: v6, F1: 272, F2: -354}
            _ = v25
        }
        var v26 int = 32768
        _ = v26
    }
    var v27 int32 = 1
    fmt.Println(v6, v5, v27)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
