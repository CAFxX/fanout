package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 uint }
func genFn0(p0 uint64) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 float32 = float32(-2.25)
        if false {
            var v10 bool = true
            _ = v10
        } else {
            var v11 bool = (18446744073709551615 < 65535)
            var v12 string = "\n\t"
            _ = v11
            _ = v12
        }
        var v13 uint64 = p0
        _ = fi8
        _ = v13
        _ = v9
    }
    _ = p0
    return 2
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -52
    v3 := make(map[uint32]string)
    var v4 int8 = v2
    v5 := []uint32{}
    v6 := []int{-1, 32768}
    v6 = append(v6, 42)
    var v14 int16 = 42
    v4 += v2
    if (true && (false || (v14 > v14))) {
        for fi15 := 0; fi15 < 1; fi15++ {
            if (0 != genFn0(18446744073709551615)) {
                fmt.Println(v4, v14, v4)
                for fi16 := 0; fi16 < 8; fi16++ {
                    if false {
                        switch (fi15 / ((fi15) | 1)) {
                        case 0:
                            _ = len(v3)
                        case -273:
                            switch "\n\t" {
                            case "hello":
                                if true {
                                    var v17 uint16 = 126
                                    fmt.Println(fi16, fi15, v17)
                                    _ = v17
                                }
                            case "":
                                if true {
                                    var v18 uint = 4294967295
                                    _ = v18
                                } else {
                                    var v19 int64 = (-(v1))
                                    fmt.Println(v2, fi15, v2)
                                    _ = v19
                                }
                            }
                        }
                        switch fi16 {
                        case -460:
                            var v20 uint8 = 2
                            _ = v20
                        default:
                            fmt.Println(v14)
                        }
                        v2 += (-(v4))
                        if (1 <= 858) {
                            switch "hello" {
                            case "":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v21 string = ("foo bar" + ("foo bar" + "a"))
                                switch -2147483648 {
                                case 32768:
                                    var v22 float32 = 1e10
                                    _ = v22
                                }
                                _ = v21
                            default:
                                switch fi16 {
                                case 1:
                                    v5 = append(v5, 711)
                                default:
                                    v23 := GenS1{F0: 42}
                                    _ = v23
                                }
                            }
                        }
                    }
                    _ = fi16
                }
            }
            for fi24 := 0; fi24 < 10; fi24++ {
                v2 = (^(v4))
                _ = fi24
            }
            _ = fi15
        }
    }
    v25 := GenS2{F0: (1 < 0), F1: -88, F2: 65535}
    for fi26 := 0; fi26 < 9; fi26++ {
        if false {
            if false {
                fmt.Println(v14)
            }
            var v27 uint16 = 2
            _ = v27
        } else {
            switch "" {
            case "hello":
                v2 = v4
            }
            var v28 float64 = 0.1
            _ = v28
        }
        var v29 float64 = 0.1
        _ = fi26
        _ = v29
    }
    fmt.Println(v4, v4)
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 bool = false
        v14 -= v14
        _ = fi30
        _ = v31
    }
    v1 = (v1 ^ int64(v2))
    delete(v3, 380)
    if true {
        switch -2147483648 {
        case 575:
            if true {
                if false {
                    var v32 int32 = -2147483648
                    v33 := GenS3{F0: (v14 >> ((65535) & 63)), F1: 127, F2: 65535}
                    switch -961 {
                    case 265:
                        v14 = 1000
                    }
                    _ = v32
                    _ = v33
                } else {
                    if len(v6) > 0 {
                        v6[0] = 940
                    }
                    var v34 float32 = (3.14159 + (1e10 + 0.1))
                    v14 += (v14 << ((255) & 63))
                    _ = v34
                }
            }
        case 128:
            var v35 float64 = (1e10 + (1.5 / (0.0 - -2.25)))
            _ = v35
        default:
            fmt.Println(v14)
        }
        var v36 uint8 = 145
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
