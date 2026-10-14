package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int64, p1 uint64) bool {
    v5 := GenS1{F0: 43, F1: -391}
    p1 = p1
    v6 := GenS1{F0: 575, F1: -2147483648}
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]string)
    var v3 int8 = -21
    var v4 float32 = -2.25
    var v7 int16 = 464
    v8 := []int{334, -2147483648, 9223372036854775807, -683}
    var v9 int16 = -299
    switch 0 {
    case 128:
        if ((!true) || false) {
            fmt.Println(v3, v4, v3)
            for fi10 := 0; fi10 < 8; fi10++ {
                v7 -= (v7 / (((v9 << ((0) & 63))) | 1))
                _ = fi10
            }
        }
    case 869:
        if (745 == 4294967295) {
            fmt.Println(v9, v3, v3)
        }
    default:
        if false {
            v11 := GenS2{F0: 1}
            var v12 uint32 = 1
            switch 4294967295 {
            case 1:
                v1 -= 100
            case 32768:
                var v13 uint = 65535
                _ = v13
            case 396:
                switch ("hello" + "hello") {
                case "hello":
                    var v14 string = ""
                    _ = v14
                case "\n\t":
                    var v15 int8 = (v3 - v1)
                    _ = v15
                case "foo bar":
                    v4 = v4
                }
            default:
                switch "a" {
                case "foo bar":
                    if (42 <= 535) {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v9 -= v9
                            var v17 int16 = 770
                            var v18 uint16 = 2
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        fmt.Println(v4)
                    }
                case "hello":
                    for fi19 := 0; fi19 < 14; fi19++ {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            if (65535 <= 42) {
                                if (!((!true) && (!true))) {
                                    fmt.Println(v1, v12, fi20)
                                    v21 := GenS2{F0: 1}
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v21
                                }
                                switch -703 {
                                case 9223372036854775807:
                                    v22 := GenS2{F0: -2147483648}
                                    _ = v22
                                }
                            }
                            v23 := GenS1{F0: 65535, F1: -2147483648}
                            _ = fi20
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 uint8 = 1
                            v8 = append(v8, fi19)
                            var v26 int8 = (^(v1))
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v27 uint16 = 65535
                        _ = fi19
                        _ = v27
                    }
                }
            }
            _ = v11
            _ = v12
        } else {
            var v28 uint32 = 42
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        v1 = 42
        v1 = (^(v1))
        _ = fi29
    }
    if false {
        v1 += -1
    }
    switch -509 {
    case -645:
        switch "\n\t" {
        case "a":
            var v30 uint32 = 0
            _ = v30
        }
    default:
        v31 := GenS3{F0: 0, F1: 584}
        _ = v31
    }
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
