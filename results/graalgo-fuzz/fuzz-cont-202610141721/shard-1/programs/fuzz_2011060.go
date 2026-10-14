package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 42
    var v3 uint = 285
    v4 := []string{"hello", "\n\t"}
    var v5 float32 = ((((1e10 - (3.14159 * 1.5)) / 0.1) - (1.5 - -2.25)) / (((float32(3.14159) - (3.14159 - 0.0)) - 0.1) - 0.0))
    fmt.Println(v3, v5, v5)
    var v6 bool = (-192 > 1)
    var v7 string = ""
    v7 = v7
    v7 = ((v7 + ((v7 + v7) + v7)) + "a")
    for fi8 := 0; fi8 < 10; fi8++ {
        v2 -= (v2 | (v2 - (-(v2))))
        _ = fi8
    }
    var v9 int32 = 0
    v5 += (1e10 / (v5 * 0.0))
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v4) > 0 {
            v4[0] = v7
        }
        switch fi10 {
        case 32767:
            v2 += int64(v3)
        case 32768:
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int = fi11
                v12 -= -1
                var v13 int = v12
                v9 += int32(v2)
                if false {
                    switch "\n\t" {
                    case "":
                        v1[(v13 / (((v13 % (((^(v13))) | 1))) | 1))] = "a"
                        _ = len(v1)
                    case "foo bar":
                        v6 = false
                    }
                    switch (v7 + "hello") {
                    case "a":
                        v12 -= (-(v12))
                        switch v7 {
                        case "hello":
                            v14 := GenS1{F0: 180}
                            _ = v14
                        default:
                            var v15 bool = (v9 < (-(v9)))
                            _ = v15
                        }
                        v4 = append(v4, v7)
                    case "hello":
                        v16 := GenS2{F0: (v2 << ((42) & 63)), F1: v3, F2: 441}
                        _ = v16
                    }
                    v17 := GenS3{F0: (v2 << ((v3) & 63)), F1: 2, F2: (v2 + v2)}
                    v18 := GenS2{F0: (v2 + (v2 ^ 0)), F1: (v3 * v3), F2: 42}
                    v19 := GenS4{F0: v2, F1: 816}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    if ((v9 << ((573) & 63)) > 0) {
                        delete(v1, v12)
                        switch 128 {
                        case 592:
                            if (-969 > 32767) {
                                _ = len(v1)
                            }
                        case 847:
                            switch fi11 {
                            case -9223372036854775808:
                                v20 := GenS5{F0: v3, F1: 42}
                                _ = v20
                            case 0:
                                v13 -= 128
                                v4 = append(v4, v7)
                                switch 4294967295 {
                                case 338:
                                    v21 := GenS2{F0: (v2 * v2), F1: 2, F2: -444}
                                    _ = v21
                                case -955:
                                    var v22 float64 = (1.5 + 3.14159)
                                    _ = v22
                                default:
                                    v6 = v6
                                }
                            case 211:
                                var v23 float64 = (1e10 * 1.5)
                                _ = v23
                            }
                        case 1:
                            var v24 uint64 = 0
                            _ = v24
                        }
                        var v25 uint64 = 1
                        _ = v25
                    }
                    v7 = v7
                }
                fmt.Println(v9, v2, v3)
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            switch "foo bar" {
            case "foo bar":
                for fi26 := 0; fi26 < 13; fi26++ {
                    v4 = append(v4, (v7 + ""))
                    _ = fi26
                }
            case "a":
                v27 := GenS4{F0: 42, F1: 2}
                _ = v27
            case "\n\t":
                v28 := GenS3{F0: v2, F1: 2, F2: (v2 - (v2 * int64(v9)))}
                _ = v28
            }
        }
        fmt.Println(v5, v7, v7)
        _ = fi10
    }
    var v29 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
