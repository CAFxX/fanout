package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int8 = 127
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    v4 := []string{v2, "hello"}
    var v5 uint = 2
    var v6 uint64 = 137
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 2
        _ = fi7
        _ = v8
    }
    if (331 == 1000) {
        var v9 float32 = 3.14159
        v5 = v5
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: (v5 >> ((v5) & 63)), F1: 49}
                _ = fi11
                _ = v12
            }
            var v13 int64 = -9223372036854775808
            switch (fi10 % (((fi10 ^ int(v1))) | 1)) {
            case 2147483647:
                _ = len(v4)
            case 32767:
                fmt.Println(v6, v5, v5)
            default:
                v14 := GenS2{F0: v5, F1: 147}
                _ = v14
            }
            _ = v13
        } else {
            v3[(fi10 << (((v5 ^ (v5 >> ((4294967295) & 63)))) & 63))] = v2
        }
        v5 += uint(v6)
        _ = fi10
    }
    if false {
        var v15 uint16 = 42
        fmt.Println(v6, v5, v6)
        _ = v15
    } else {
        v16 := GenS2{F0: (v5 / (((v5 >> (((v5 | 254)) & 63))) | 1)), F1: 42}
        if (!(!false)) {
            for fi17 := 0; fi17 < 12; fi17++ {
                switch v2 {
                case "a":
                    v6 = (v6 / (((^(v6))) | 1))
                case "foo bar":
                    fmt.Println(v2, v5)
                case "\n\t":
                    if (!(-2.25 <= 0.1)) {
                        v18 := GenS2{F0: v5, F1: 42}
                        v1 += (^(v1))
                        switch fi17 {
                        case 1:
                            switch (v2 + v2) {
                            case "":
                                v2 = v2
                            default:
                                if (!false) {
                                    if true {
                                        v6 = v6
                                    }
                                }
                            }
                        case 32767:
                            var v19 float32 = 0.1
                            _ = v19
                        case -9223372036854775808:
                            v6 -= (v6 - (v6 | 18446744073709551615))
                        }
                        _ = v18
                    }
                default:
                    v20 := GenS3{F0: (v6 << ((v5) & 63))}
                    _ = v20
                }
                _ = fi17
            }
        } else {
            v2 = (v2 + v2)
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(fi21, fi21)
                if false {
                    fmt.Println(fi21, v2)
                }
                var v22 uint = v5
                _ = fi21
                _ = v22
            }
            switch -293 {
            case 1:
                var v23 bool = ((true || (!true)) || (!false))
                _ = v23
            }
        }
        if len(v4) > 0 {
            v4[0] = ("hello" + (v2 + ((((v2 + v2) + v2) + ("hello" + v2)) + "\n\t")))
        }
        _ = v16
    }
    var v24 uint32 = 675
    if false {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 bool = (-1 != -1)
            var v27 int8 = (v1 | (v1 - v1))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 int32 = 1
    var v29 uint64 = (v6 * (v6 | 689))
    v24 = v24
    var v30 float64 = 3.14159
    if (v30 >= ((v30 * ((v30 - (v30 - 0.0)) * float64(float32(v30)))) * 1.5)) {
        fmt.Println(v5)
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
