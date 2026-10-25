package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int; F2 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{42}
    var v3 uint16 = 42
    v4 := []string{(("" + "") + "a"), ("foo bar" + "foo bar"), "hello"}
    var v5 string = ""
    switch (v5 + "") {
    case "":
        fmt.Println(v1)
    case "\n\t":
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint8 = 171
            _ = fi6
            _ = v7
        }
    case "a":
        fmt.Println(v5)
    }
    var v8 bool = ((v1 % ((18446744073709551615) | 1)) == 0)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v9 uint32 = 94
    var v10 int32 = -537
    var v11 uint64 = v1
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint8 = 38
        for fi14 := 0; fi14 < 19; fi14++ {
            switch v5 {
            case "\n\t":
                switch 32767 {
                case 9223372036854775807:
                    switch "a" {
                    case "":
                        v3 = (v3 ^ (v3 & v3))
                    case "hello":
                        v13 = 42
                    default:
                        fmt.Println(fi12, v13, v1)
                    }
                case 128:
                    v15 := GenS1{F0: (v10 / (((v10 % ((-2147483648) | 1))) | 1)), F1: (v5 + ""), F2: (v9 << ((2) & 63))}
                    _ = v15
                case 32768:
                    if v8 {
                        v16 := GenS2{F0: ((255 <= (v13 / ((v13) | 1))) && (((((v5 + v5) + "foo bar") + "hello") + v5) != v5))}
                        _ = v16
                    }
                default:
                    switch -528 {
                    case -2147483648:
                        if (-1 <= int32(fi12)) {
                            v2 = append(v2, uint32(v13))
                            fmt.Println(v8)
                            v9 += (v9 % (((v9 | (v9 + (v9 - (v9 % ((v9) | 1)))))) | 1))
                        } else {
                            fmt.Println(v8)
                            if v8 {
                                fmt.Println(fi14, v11, v13)
                                if ((v10 | v10) >= -2147483648) {
                                    var v17 uint32 = v9
                                    v9 = (^(v17))
                                    fmt.Println(fi12, v17)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v17
                                }
                                v4 = append(v4, "\n\t")
                                v3 -= (v3 ^ (^(v3)))
                                v1 = 4294967295
                            } else {
                                switch 32767 {
                                case 9223372036854775807:
                                    var v18 int32 = (v10 ^ v10)
                                    _ = v18
                                case -879:
                                    v10 -= (^(v10))
                                default:
                                    switch -353 {
                                    case 406:
                                        switch v5 {
                                        case "foo bar":
                                            if true {
                                                if (v13 < uint8(v1)) {
                                                    var v19 uint8 = v13
                                                    var v20 uint32 = v9
                                                    _ = v19
                                                    _ = v20
                                                }
                                            } else {
                                                var v21 int = -9223372036854775808
                                                v22 := GenS3{F0: uint32(v13), F1: (v21 << ((195) & 63)), F2: 108}
                                                _ = v21
                                                _ = v22
                                            }
                                        default:
                                            var v23 float64 = (0.0 / ((3.14159 / 1.0) + 1.5))
                                            _ = v23
                                        }
                                    }
                                }
                            }
                        }
                    case 32767:
                        if len(v4) > 0 {
                            v4[0] = ((((("foo bar" + v5) + ("\n\t" + v5)) + "hello") + v5) + (v5 + "a"))
                        }
                    default:
                        var v24 int32 = 2147483647
                        _ = v24
                    }
                }
            case "foo bar":
                v4 = append(v4, v5)
            }
            v3 = v3
            _ = fi14
        }
        v11 = v1
        _ = fi12
        _ = v13
    }
    v3 -= v3
    v11 += 1
    var v25 uint16 = (v3 * (v3 % ((v3) | 1)))
    var v26 bool = (!v8)
    fmt.Println(v8, v26)
    var v27 uint32 = 0
    v5 = (v5 + "\n\t")
    v26 = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
