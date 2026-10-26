package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int8 = 42
    var v2 int32 = -39
    v3 := []uint32{2147483647, 962, 4294967295, 2}
    var v4 uint32 = 65535
    var v5 float32 = (0.1 / ((0.0 * 0.0) - -2.25))
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: true}
        _ = fi6
        _ = v7
    }
    v1 = (v1 | (v1 - v1))
    _ = len(v3)
    if (-2.25 > 0.0) {
        if false {
            v1 = (v1 >> ((65535) & 63))
            var v8 string = "\n\t"
            v5 -= -2.25
            _ = v8
        }
        v2 += (^(v2))
    } else {
        if true {
            fmt.Println(v2, v4, v2)
        }
        fmt.Println(v4, v2)
        var v9 uint64 = 496
        _ = v9
    }
    var v10 int32 = 1
    for fi11 := 0; fi11 < 10; fi11++ {
        switch (fi11 << ((42) & 63)) {
        case 801:
            var v12 int8 = v1
            _ = v12
        case 32768:
            switch fi11 {
            case -2147483648:
                v2 = v10
            case -972:
                if (!true) {
                    v13 := GenS2{F0: 1, F1: int16(v4)}
                    _ = v13
                } else {
                    var v14 int64 = 42
                    switch -553 {
                    case 122:
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v16 := GenS1{F0: (fi15 != -1)}
                            _ = fi15
                            _ = v16
                        }
                    case 32768:
                        v17 := GenS1{F0: (true && true)}
                        _ = v17
                    case 2147483647:
                        if (v1 >= (v1 / ((v1) | 1))) {
                            var v18 float32 = 1e10
                            v19 := GenS2{F0: 42, F1: -106}
                            fmt.Println(v18)
                            _ = v18
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                if true {
                                    fmt.Println(v10, v2, v1)
                                    v5 -= v5
                                    var v21 int = fi20
                                    _ = v21
                                }
                                var v22 int8 = v1
                                switch "\n\t" {
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "foo bar":
                                    v23 := GenS1{F0: false}
                                    _ = v23
                                case "":
                                    v4 = 644
                                default:
                                    var v24 float32 = (v5 + 3.14159)
                                    _ = v24
                                }
                                fmt.Println(v14, v10)
                                v25 := GenS2{F0: 32767, F1: 0}
                                fmt.Println(v14)
                                _ = fi20
                                _ = v22
                                _ = v25
                            }
                        }
                    default:
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 int32 = -55
                            v3 = append(v3, v4)
                            _ = fi26
                            _ = v27
                        }
                    }
                    v5 = (1e10 - float32(1e10))
                    _ = v14
                }
            default:
                if ((!(0 != 635)) || true) {
                    switch "hello" {
                    case "\n\t":
                        v10 = (v2 & -477)
                    }
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v2 -= 2147483647
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 4; fi29++ {
                        _ = len(v3)
                        var v30 int64 = -9223372036854775808
                        _ = fi29
                        _ = v30
                    }
                } else {
                    v2 += 1
                    if false {
                        if (true && (true && (v5 >= 3.14159))) {
                            var v31 uint64 = 18446744073709551615
                            _ = len(v3)
                            _ = len(v3)
                            _ = v31
                        } else {
                            switch fi11 {
                            case 128:
                                switch 127 {
                                case 32768:
                                    v32 := GenS2{F0: 0, F1: -789}
                                    _ = v32
                                default:
                                    fmt.Println(fi11)
                                }
                            case 53:
                                var v33 uint64 = 298
                                _ = v33
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            var v34 float64 = 0.1
                            _ = v34
                        }
                    } else {
                        var v35 uint64 = 2
                        var v36 int64 = -9223372036854775808
                        _ = v35
                        _ = v36
                    }
                }
            }
        case 42:
            var v37 uint64 = 866
            _ = v37
        default:
            v4 += 232
        }
        _ = fi11
    }
    var v38 int8 = v1
    for fi39 := 0; fi39 < 19; fi39++ {
        fmt.Println(v10)
        _ = fi39
    }
    v4 = (^(v4))
    v40 := GenS2{F0: 65535, F1: 1}
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
