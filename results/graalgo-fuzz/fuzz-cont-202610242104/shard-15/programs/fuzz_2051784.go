package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 bool = (true && true)
    v2 := []int{-1, 127, 658}
    var v3 int16 = 315
    var v4 uint64 = 4294967295
    var v5 string = ("\n\t" + "")
    v6 := GenS1{F0: 42, F1: -2147483648, F2: false}
    v2 = append(v2, 142)
    v7 := GenS2{F0: 297, F1: 18446744073709551615, F2: -37}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 string = ""
        var v10 int32 = -1
        var v11 int32 = v10
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v3 = -184
    v2 = append(v2, 2147483647)
    v2 = append(v2, 127)
    var v12 bool = (1.5 != -2.25)
    v13 := GenS2{F0: (v4 + (^(v4))), F1: 4294967295, F2: 0}
    var v14 uint16 = 390
    switch "" {
    case "\n\t":
        v15 := GenS2{F0: 1, F1: 27, F2: 0}
        _ = v15
    case "foo bar":
        var v16 float32 = 0.0
        _ = v16
    case "hello":
        if true {
            var v17 float64 = (3.14159 / -2.25)
            v12 = ((v3 | 1) <= (^(v3)))
            if (v1 && (!(2147483648 >= 4294967295))) {
                _ = len(v2)
                v18 := GenS1{F0: uint32(v14), F1: 128, F2: v1}
                _ = v18
            } else {
                switch int(v14) {
                case -9223372036854775808:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v20 := GenS2{F0: (v4 | (v4 << ((571) & 63))), F1: 946, F2: -1}
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 int8 = -1
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                switch (v5 + v5) {
                case "foo bar":
                    var v23 int16 = (v3 << ((255) & 63))
                    _ = v23
                case "\n\t":
                    switch 32767 {
                    case -1:
                        for fi24 := 0; fi24 < 18; fi24++ {
                            switch v5 {
                            case "foo bar":
                                switch fi24 {
                                case -2147483648:
                                    v5 = ""
                                default:
                                    var v25 int64 = 0
                                    _ = v25
                                }
                            case "a":
                                v14 = (v14 / ((v14) | 1))
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    var v26 float64 = float64(0.1)
                                    _ = v26
                                }
                            default:
                                fmt.Println(fi24)
                            }
                            if (-2.25 >= 3.14159) {
                                if v12 {
                                    var v27 int32 = 0
                                    if v1 {
                                        var v28 bool = (fi24 != fi24)
                                        v29 := GenS2{F0: v4, F1: 4294967295, F2: -593}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        for fi30 := 0; fi30 < 20; fi30++ {
                                            v31 := GenS3{F0: v5, F1: 127}
                                            v3 += (v3 + (v3 ^ v3))
                                            fmt.Println(fi30, v3, v3)
                                            _ = fi30
                                            _ = v31
                                        }
                                        var v32 uint = 255
                                        fmt.Println(v1)
                                        _ = v32
                                    }
                                    v33 := GenS1{F0: 65535, F1: fi24, F2: true}
                                    _ = v27
                                    _ = v33
                                }
                                var v34 uint64 = (^(v4))
                                _ = v34
                            }
                            var v35 int16 = -920
                            _ = fi24
                            _ = v35
                        }
                    case 0:
                        switch ((v5 + v5) + "foo bar") {
                        case "foo bar":
                            if v12 {
                                v17 = v17
                                v36 := GenS4{F0: 42, F1: (^(v3))}
                                fmt.Println(v17)
                                _ = v36
                            } else {
                                _ = len(v2)
                            }
                        }
                    case 42:
                        _ = len(v2)
                    }
                case "a":
                    for fi37 := 0; fi37 < 20; fi37++ {
                        for fi38 := 0; fi38 < 13; fi38++ {
                            v39 := GenS2{F0: (v4 >> ((113) & 63)), F1: v4, F2: -615}
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                }
            }
            _ = v17
        } else {
            switch 0 {
            case 535:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 349:
                v40 := GenS4{F0: 87, F1: 0}
                _ = v40
            }
        }
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 uint = 2
        _ = fi41
        _ = v42
    }
    fmt.Println(v1)
    v43 := GenS1{F0: 0, F1: -9223372036854775808, F2: (0 < 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
