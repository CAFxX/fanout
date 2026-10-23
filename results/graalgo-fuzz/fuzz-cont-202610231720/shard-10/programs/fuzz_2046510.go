package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint = 1
    var v2 uint8 = uint8(v1)
    v3 := []string{"a", "", (("" + "\n\t") + "a")}
    var v4 string = "foo bar"
    v5 := []int{}
    v5 = append(v5, -2147483648)
    v6 := GenS1{F0: 42, F1: v4, F2: 730}
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
        fmt.Println(v4)
    } else {
        var v7 int = 4294967295
        _ = v7
    }
    v8 := GenS1{F0: 279, F1: v4, F2: 549}
    v4 = (v4 + "hello")
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v4, fi9, v1)
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS2{F0: v4}
            v4 = v4
            v1 -= v1
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9, v2, v4)
        _ = fi9
    }
    switch 29 {
    case 32767:
        var v12 uint32 = 2147483648
        _ = v12
    }
    v13 := GenS3{F0: (("hello" + v4) + v4)}
    switch v4 {
    case "\n\t":
        if (!(((1.5 - 0.0) + 1.5) == float32((3.14159 + ((3.14159 * float64(0.1)) - -2.25))))) {
            if false {
                v2 = (v2 ^ v2)
                v14 := GenS4{F0: 32768, F1: v1, F2: v2}
                _ = v14
            }
            v15 := GenS4{F0: 781, F1: 0, F2: 42}
            _ = v15
        }
    case "foo bar":
        if true {
            _ = len(v5)
        } else {
            v16 := GenS3{F0: ("" + v4)}
            v2 += 1
            _ = v16
        }
    case "hello":
        if true {
            v2 += v2
            for fi17 := 0; fi17 < 11; fi17++ {
                v1 = 0
                _ = fi17
            }
            v2 -= (v2 >> ((255) & 63))
        } else {
            fmt.Println(v1, v4)
            if true {
                fmt.Println(v1, v1, v2)
                var v18 uint64 = 0
                v19 := GenS5{F0: 1, F1: (v18 << ((v1) & 63))}
                _ = v18
                _ = v19
            } else {
                if (v4 > v4) {
                    var v20 uint8 = (v2 << ((v1) & 63))
                    var v21 int8 = 42
                    v21 = 51
                    _ = v20
                    _ = v21
                }
                v2 = (v2 + (^(v2)))
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch fi22 {
                    case 9223372036854775807:
                        var v23 uint = 65535
                        _ = v23
                    }
                    if (0 < 397) {
                        v24 := GenS4{F0: fi22, F1: (v1 / (((v1 >> (((v1 ^ v1)) & 63))) | 1)), F2: v2}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v4, v1, v1)
                        _ = v24
                    } else {
                        switch (v4 + "a") {
                        case "":
                            if (!((1e10 / 1e10) >= ((3.14159 - -2.25) - (((1.5 * 0.1) - 0.0) + 1e10)))) {
                                var v25 uint16 = 447
                                fmt.Println(v25)
                                var v26 uint8 = 0
                                _ = v25
                                _ = v26
                            }
                        case "a":
                            fmt.Println(v1)
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    }
                    _ = fi22
                }
            }
        }
    }
    v5 = append(v5, 127)
    if ("" == v4) {
        v2 = (v2 - v2)
    }
    v4 = v4
    var v27 float64 = 1e10
    v28 := GenS1{F0: -9223372036854775808, F1: v4, F2: 2147483647}
    v4 = "\n\t"
    v29 := GenS4{F0: -993, F1: v1, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
