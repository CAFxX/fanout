package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 uint64) uint64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int8 = 100
            fmt.Println(fi5)
            var v8 uint8 = 128
            _ = fi6
            _ = v7
            _ = v8
        }
        p0 += p0
        var v9 uint = 23
        _ = fi5
        _ = v9
    }
    p0 = (p0 / ((p0) | 1))
    _ = p0
    return (p0 % ((18446744073709551615) | 1))
}

func main() {
    var v1 int64 = 0
    var v2 int8 = -35
    var v3 uint8 = 128
    v4 := []uint32{4294967295}
    var v10 int16 = 1
    fmt.Println(v2, v1)
    v11 := GenS1{F0: "", F1: 18446744073709551615}
    v10 -= v10
    fmt.Println(v1, v2, v2)
    if (4294967295 > 2) {
        fmt.Println(v10)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= v1
    switch ("foo bar" + "\n\t") {
    case "a":
        v2 += v2
    case "foo bar":
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 = -128
            _ = fi12
        }
    case "hello":
        var v13 uint8 = v3
        _ = v13
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch "\n\t" {
    case "foo bar":
        v14 := GenS2{F0: "", F1: v3, F2: (4294967295 == 267)}
        _ = v14
    case "":
        for fi15 := 0; fi15 < 15; fi15++ {
            v2 = v2
            switch fi15 {
            case -9223372036854775808:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 241:
                v3 = v3
            case 1:
                var v16 uint = 42
                _ = v16
            default:
                if (-2.25 == 0.1) {
                    switch (fi15 % ((fi15) | 1)) {
                    case -752:
                        if false {
                            if (true && true) {
                                if true {
                                    v1 -= (-(v1))
                                }
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    fmt.Println(v3, v1, v10)
                                    _ = fi17
                                }
                            }
                            switch (-(fi15)) {
                            case 2147483647:
                                if (true && (475 >= 386)) {
                                    v1 = v1
                                    var v18 bool = (false || false)
                                    _ = v18
                                }
                            }
                        }
                    default:
                        v3 = v3
                    }
                    switch fi15 {
                    case 364:
                        if (2 == 957) {
                            v19 := GenS3{F0: 18446744073709551615, F1: (^(v3)), F2: 588}
                            _ = v19
                        }
                    case 4294967295:
                        v20 := GenS1{F0: "a", F1: genFn0(438)}
                        _ = v20
                    default:
                        v3 = (v3 * v3)
                    }
                }
            }
            v3 -= v3
            _ = fi15
        }
    default:
        fmt.Println(v2, v10, v3)
    }
    v21 := GenS2{F0: "hello", F1: 127, F2: (v1 != v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
