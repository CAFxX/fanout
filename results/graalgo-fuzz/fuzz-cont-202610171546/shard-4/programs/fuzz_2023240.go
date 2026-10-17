package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint32 = 579
    var v2 uint32 = v1
    v3 := []int{2147483647, 32767}
    var v4 uint16 = 65535
    var v5 int16 = -1
    var v6 int32 = -307
    switch "\n\t" {
    case "hello":
        var v7 int = 127
        _ = v7
    case "\n\t":
        fmt.Println(v1, v4, v4)
    case "a":
        var v8 int8 = 42
        _ = v8
    default:
        v6 += -97
    }
    v5 -= v5
    for fi9 := 0; fi9 < 1; fi9++ {
        if false {
            var v10 int16 = -494
            if true {
                switch "\n\t" {
                case "hello":
                    if ("" == ("\n\t" + "")) {
                        var v11 bool = (!((v10 / ((v5) | 1)) != 926))
                        v3 = append(v3, 1)
                        if false {
                            switch 1 {
                            case -735:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 127:
                                for fi12 := 0; fi12 < 7; fi12++ {
                                    switch "foo bar" {
                                    case "":
                                        v13 := GenS1{F0: (0.1 != -2.25), F1: 42}
                                        _ = v13
                                    default:
                                        v1 = 2
                                    }
                                    var v14 bool = v11
                                    _ = fi12
                                    _ = v14
                                }
                            default:
                                fmt.Println(v1, v5)
                            }
                        } else {
                            v15 := GenS1{F0: v11, F1: 275}
                            _ = v15
                        }
                        _ = v11
                    } else {
                        switch fi9 {
                        case 0:
                            var v16 int32 = v6
                            _ = v16
                        case 578:
                            var v17 bool = true
                            _ = v17
                        default:
                            switch -948 {
                            case 618:
                                fmt.Println(v5, v1)
                            }
                        }
                        switch "" {
                        case "":
                            v5 = v5
                        case "\n\t":
                            switch (fi9 / ((fi9) | 1)) {
                            case 805:
                                v18 := GenS1{F0: (true && (422 > 42)), F1: 42}
                                _ = v18
                            case 674:
                                v6 = (v6 % (((^(v6))) | 1))
                            default:
                                var v19 float32 = 3.14159
                                _ = v19
                            }
                        }
                        v6 = (^(v6))
                    }
                case "a":
                    var v20 uint16 = v4
                    _ = v20
                default:
                    _ = len(v3)
                }
                if (true || (false || false)) {
                    var v21 uint16 = v4
                    v6 = 31
                    var v22 float64 = 3.14159
                    _ = v21
                    _ = v22
                }
                var v23 int32 = (^(v6))
                _ = v23
            } else {
                var v24 float64 = 1e10
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 float64 = 1.5
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            _ = len(v3)
            _ = v10
        } else {
            var v27 int16 = v5
            _ = v27
        }
        if false {
            if len(v3) > 0 {
                v3[0] = (fi9 / ((fi9) | 1))
            }
            v28 := GenS1{F0: (65535 >= 18446744073709551615), F1: 18446744073709551615}
            _ = v28
        }
        var v29 int32 = (v6 * (v6 + (-(v6))))
        _ = fi9
        _ = v29
    }
    _ = len(v3)
    if (!true) {
        var v30 int64 = 190
        var v31 int32 = 42
        _ = v30
        _ = v31
    } else {
        v1 -= (v2 >> ((1) & 63))
    }
    v6 = (-(v6))
    v1 = v2
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint = 942
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS1{F0: false, F1: 2}
            switch fi34 {
            case 2147483647:
                var v36 uint8 = 163
                _ = v36
            case -777:
                var v37 int8 = int8(v33)
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    v2 += (v2 * (v2 >> ((65535) & 63)))
    var v38 float32 = ((0.1 * -2.25) * (-2.25 * ((3.14159 - 1.5) / 1.0)))
    v39 := GenS2{F0: 460}
    for fi40 := 0; fi40 < 20; fi40++ {
        _ = len(v3)
        v6 = v6
        _ = fi40
    }
    if false {
        if ((-2.25 + 3.14159) != 1e10) {
            v1 = (v1 - v2)
        }
    } else {
        v5 -= v5
        var v41 uint16 = v4
        var v42 float64 = float64((v38 - v38))
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
