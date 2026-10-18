package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]bool)
    v3 := []int{2147483647, 128, 447}
    v4 := []uint32{}
    v5 := make(map[int]int)
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            _ = len(v2)
        } else {
            fmt.Println(v6)
            var v8 float64 = ((v1 - v1) - 3.14159)
            if (true && ((0.0 != 0.1) || true)) {
                fmt.Println(fi7)
                if true {
                    fmt.Println(v6)
                    if (!(true || (-9223372036854775808 == 11))) {
                        v6 = (^(v6))
                        if (true || true) {
                            var v9 float64 = float64(0.0)
                            v1 += (3.14159 / ((float64(1e10) - ((((v9 - ((float64(1.5) + v9) / 1.0)) - (v9 * -2.25)) - (v9 + (float64(1.5) + v1))) * (1.5 + 1.5))) - 0.0))
                            v10 := v2[fi7]
                            _ = v10
                            _ = v9
                        } else {
                            if len(v3) > 0 {
                                v3[0] = -9223372036854775808
                            }
                        }
                        v11 := GenS1{F0: 32767}
                        _ = v11
                    } else {
                        v1 = (((v8 + ((1e10 / v1) / (float64(0.0) * v8))) * (v1 + v8)) / -2.25)
                        v12 := v5[(^(fi7))]
                        v8 = v1
                        _ = v12
                    }
                    v6 -= (v6 - v6)
                }
                v1 -= 0.0
            }
            _ = v8
        }
        var v13 int16 = 1
        if (!(("" >= "") || (!false))) {
            if false {
                v13 = (v13 - (v13 << ((42) & 63)))
                var v14 uint32 = (v6 ^ 4294967295)
                _ = v14
            }
            _ = len(v3)
            _ = len(v5)
        } else {
            var v15 uint64 = 460
            fmt.Println(fi7)
            v1 = (v1 + (v1 + 0.1))
            _ = v15
        }
        _ = fi7
        _ = v13
    }
    fmt.Println(v6, v6)
    v16 := GenS1{F0: 255}
    v17 := GenS2{F0: 42, F1: 32767}
    var v18 int32 = 42
    fmt.Println(v18)
    _ = len(v5)
    v1 += (float64(0.1) - ((v1 + (float64(1e10) + v1)) * (v1 + (v1 / float64(1e10)))))
    v6 -= (^(v6))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ("\n\t" < "foo bar") {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v6 = (v6 << ((65535) & 63))
        if ((v18 * -2147483648) < v18) {
            var v19 int16 = 42
            fmt.Println(v6, v1, v19)
            _ = v19
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            if false {
                switch -246 {
                case 32768:
                    switch "hello" {
                    case "hello":
                        var v20 bool = (false && (false && true))
                        _ = v20
                    case "a":
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            if false {
                                fmt.Println(v18, v18)
                                v5[-1] = 32767
                                v21 := GenS2{F0: 1, F1: 565}
                                _ = v21
                            } else {
                                v22 := GenS2{F0: 127, F1: 42}
                                _ = v22
                            }
                        }
                    default:
                        v23 := v5[4294967295]
                        _ = v23
                    }
                case -2147483648:
                    _ = len(v2)
                case -57:
                    fmt.Println(v1, v18, v6)
                }
            } else {
                if false {
                    var v24 uint8 = 0
                    var v25 float64 = v1
                    var v26 int32 = 42
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        case "\n\t":
            fmt.Println(v18)
        case "a":
            v18 = 42
        default:
            v6 -= (v6 & 451)
        }
        v6 = 2147483647
        v1 = float64(1.5)
    }
    switch "foo bar" {
    case "foo bar":
        v18 -= v18
    case "\n\t":
        if true {
            if len(v3) > 0 {
                v3[0] = 32768
            }
            _ = len(v3)
        }
    case "":
        delete(v2, 51)
    }
    switch ("" + "\n\t") {
    case "\n\t":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
