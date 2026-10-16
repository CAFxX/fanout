package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    var v1 int16 = 756
    var v2 int16 = (v1 ^ -1)
    var v3 uint16 = 32767
    var v4 uint8 = 255
    v5 := []int{0, 2147483647, 127}
    v3 = v3
    v4 = v4
    var v6 int8 = 100
    v7 := GenS1{F0: -619, F1: true}
    if false {
        if true {
            v8 := GenS1{F0: -277, F1: (0 >= 2)}
            v9 := GenS2{F0: 1, F1: -597, F2: -2147483648}
            _ = v8
            _ = v9
        } else {
            v1 -= (v2 % ((v2) | 1))
            v5 = append(v5, 32767)
            var v10 float64 = 0.0
            _ = v10
        }
    } else {
        v4 -= 3
        v11 := GenS3{F0: v3, F1: 2, F2: 2}
        var v12 int = 4294967295
        _ = v11
        _ = v12
    }
    v6 = v6
    if (!true) {
        for fi13 := 0; fi13 < 20; fi13++ {
            if (416 >= 18446744073709551615) {
                v4 -= v4
                for fi14 := 0; fi14 < 8; fi14++ {
                    v2 += 1000
                    if false {
                        var v15 float64 = (1.5 / float64((1.5 / ((((-2.25 + -2.25) * -2.25) - 0.0) - 1.5))))
                        _ = v15
                    } else {
                        if true {
                            if len(v5) > 0 {
                                v5[0] = fi13
                            }
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v16 := GenS1{F0: -2147483648, F1: true}
                            _ = v16
                        }
                        var v17 bool = (0 <= -258)
                        _ = v17
                    }
                    var v18 float32 = (3.14159 * 0.1)
                    _ = fi14
                    _ = v18
                }
            } else {
                switch "foo bar" {
                case "a":
                    v19 := GenS3{F0: (v3 ^ v3), F1: 2, F2: uint64(v3)}
                    _ = v19
                case "\n\t":
                    switch "" {
                    case "":
                        if ("\n\t" < (("" + "foo bar") + ((("\n\t" + "") + "hello") + ("" + (("" + ("" + "a")) + (("" + "a") + "hello")))))) {
                            var v20 uint16 = v3
                            if false {
                                if false {
                                    v6 += 127
                                }
                                v3 -= v3
                            }
                            _ = v20
                        }
                    default:
                        v5 = append(v5, fi13)
                    }
                }
                v3 += v3
            }
            switch fi13 {
            case -650:
                for fi21 := 0; fi21 < 1; fi21++ {
                    v2 = v2
                    _ = fi21
                }
            default:
                v1 -= (v1 | v1)
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 int16 = v2
                _ = fi22
                _ = v23
            }
            _ = fi13
        }
        v24 := GenS1{F0: 781, F1: (!true)}
        v1 = v1
        _ = v24
    }
    var v25 int64 = 1
    v6 = 1
    var v26 int64 = (-(v25))
    v26 += 1
    switch "hello" {
    case "":
        v1 += (v2 << ((42) & 63))
    case "a":
        var v27 uint16 = uint16(v4)
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v5 = append(v5, fi28)
        v5 = append(v5, -1)
        _ = fi28
    }
    v25 -= (-(v26))
    fmt.Println(v26, v3, v1)
    if true {
        v2 = v1
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 float32 = 0.1
            switch ("" + "hello") {
            case "\n\t":
                switch -543 {
                case 2147483647:
                    fmt.Println(v30, v25)
                }
            default:
                switch "" {
                case "\n\t":
                    v31 := GenS1{F0: (fi29 * fi29), F1: ((!false) || true)}
                    _ = v31
                case "":
                    var v32 int64 = -670
                    _ = v32
                }
            }
            _ = fi29
            _ = v30
        }
        v1 = v1
    }
    v26 += (v26 - 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
