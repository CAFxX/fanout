package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 uint16; F1 bool; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 uint64 = 1
    v2 := []string{("hello" + "foo bar")}
    var v3 uint8 = 255
    var v4 float32 = (1e10 - 0.0)
    v3 = v3
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float64 = -2.25
        for fi7 := 0; fi7 < 20; fi7++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(fi7)
            v6 = (1e10 - (0.1 * (((v6 / v6) / v6) / float64(1e10))))
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v4, v4)
    v8 := GenS1{F0: -9223372036854775808}
    var v9 string = "\n\t"
    var v10 int32 = -1
    if (!false) {
        switch 0 {
        case -413:
            if len(v2) > 0 {
                v2[0] = v9
            }
        case -1:
            v11 := GenS2{F0: 657}
            _ = v11
        case -911:
            switch v9 {
            case "":
                switch 32767 {
                case -2147483648:
                    switch v9 {
                    case "a":
                        var v12 uint64 = (v1 >> ((496) & 63))
                        _ = v12
                    default:
                        switch "" {
                        case "\n\t":
                            v2 = append(v2, "\n\t")
                        case "":
                            v3 = 2
                        case "foo bar":
                            if false {
                                v2 = append(v2, (v9 + v9))
                                if true {
                                    for fi13 := 0; fi13 < 13; fi13++ {
                                        v2 = append(v2, "\n\t")
                                        _ = fi13
                                    }
                                }
                            } else {
                                v14 := GenS3{F0: 4294967295, F1: (false || ((false || (v4 != (float32((0.0 * 0.0)) / float32(0.1)))) && false)), F2: false}
                                for fi15 := 0; fi15 < 14; fi15++ {
                                    for fi16 := 0; fi16 < 1; fi16++ {
                                        if true {
                                            v17 := GenS4{F0: 514, F1: (-338 <= -1), F2: -300}
                                            _ = v17
                                        }
                                        var v18 int32 = v10
                                        _ = fi16
                                        _ = v18
                                    }
                                    fmt.Println(v4, v1)
                                    v19 := GenS5{F0: 18446744073709551615, F1: 0}
                                    _ = fi15
                                    _ = v19
                                }
                                if false {
                                    switch -462 {
                                    case -2147483648:
                                        fmt.Println(v1, v10, v9)
                                    case 32767:
                                        v20 := GenS4{F0: 1, F1: false, F2: 1}
                                        _ = v20
                                    case -938:
                                        var v21 string = v9
                                        _ = v21
                                    default:
                                        v9 = v9
                                    }
                                    v4 = float32(0.1)
                                    var v22 uint32 = 42
                                    _ = v22
                                }
                                _ = v14
                            }
                        }
                    }
                default:
                    var v23 int16 = 587
                    _ = v23
                }
            case "a":
                var v24 uint = 18446744073709551615
                _ = v24
            }
        }
        v25 := GenS2{F0: (v1 >> ((255) & 63))}
        _ = v25
    } else {
        var v26 int = -77
        if true {
            v9 = v9
        } else {
            switch v26 {
            case -659:
                v27 := GenS6{F0: 255}
                _ = v27
            default:
                var v28 uint32 = 112
                _ = v28
            }
            var v29 uint8 = v3
            _ = v29
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
