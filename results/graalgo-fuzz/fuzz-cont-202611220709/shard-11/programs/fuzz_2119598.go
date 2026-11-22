package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 bool) int8 {
    if (-682 >= -9223372036854775808) {
        fmt.Println(p0)
        v5 := []int{42, -2147483648}
        _ = v5
    } else {
        var v6 int16 = 890
        var v7 int64 = 920
        _ = v6
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: 192, F1: 238, F2: 1}
    v10 := GenS1{F0: 214, F1: 128, F2: 554}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 100
}

func main() {
    v1 := []string{"foo bar", "a", ("a" + "hello"), ""}
    var v2 int8 = 100
    var v3 int = 127
    v4 := make(map[uint32]uint32)
    for fi11 := 0; fi11 < 2; fi11++ {
        delete(v4, 0)
        switch ("" + ("\n\t" + "")) {
        case "a":
            v12 := GenS2{F0: 255, F1: v3}
            _ = v12
        case "\n\t":
            v4[2147483648] = 42
        }
        v3 -= fi11
        _ = fi11
    }
    var v13 int16 = -522
    v2 = v2
    v2 = genFn0((!false))
    var v14 int16 = 1000
    v15 := GenS2{F0: 1, F1: 442}
    switch "foo bar" {
    case "hello":
        v13 = 0
    case "foo bar":
        v16 := GenS2{F0: 128, F1: v3}
        _ = v16
    case "a":
        switch v3 {
        case 882:
            v17 := GenS1{F0: 127, F1: 255, F2: 2}
            _ = v17
        case -197:
            if (379 >= 0) {
                fmt.Println(v14, v13, v13)
                if (((float64(0.1) * float64(1e10)) - ((1e10 - float64((1e10 - float32(((0.0 + (3.14159 - -2.25)) + -2.25))))) + 1.5)) != 0.1) {
                    v14 = 206
                } else {
                    if true {
                        switch v3 {
                        case -702:
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                        case 0:
                            fmt.Println(v3, v3)
                        case 32768:
                            switch v3 {
                            case 631:
                                fmt.Println(v14, v14, v13)
                            }
                        default:
                            if (125 > 42) {
                                var v18 uint = 2
                                fmt.Println(v3, v14, v3)
                                _ = v18
                            }
                        }
                        v19 := v4[0]
                        _ = v19
                    }
                    v2 += v2
                }
            } else {
                switch 4294967295 {
                case 0:
                    v20 := GenS2{F0: 102, F1: v3}
                    _ = v20
                }
            }
        case 32768:
            fmt.Println(v2)
        default:
            switch "foo bar" {
            case "hello":
                v3 += v3
            case "":
                var v21 bool = true
                _ = v21
            case "\n\t":
                v14 = (^(v14))
            default:
                v22 := v4[4294967295]
                _ = v22
            }
        }
    default:
        fmt.Println(v14, v2, v13)
    }
    v13 += -870
    v2 = v2
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS2{F0: 1, F1: v3}
        v14 = v13
        _ = fi23
        _ = v24
    }
    var v25 int8 = (v2 >> ((18446744073709551615) & 63))
    fmt.Println(v13, v3)
    var v26 float32 = ((-2.25 + -2.25) / 1e10)
    v1 = append(v1, ("hello" + "foo bar"))
    if false {
        fmt.Println(v14)
        var v27 int = v3
        _ = v27
    } else {
        var v28 int64 = -430
        if ("foo bar" == "hello") {
            v26 -= (1.5 * ((v26 - ((v26 * (1e10 * 1e10)) - v26)) * -2.25))
        } else {
            if ((((float64((float32((0.1 + (0.0 + (0.1 - 1e10)))) * v26)) + 0.1) / 1e10) == float64(v26)) && false) {
                if false {
                    fmt.Println(v2)
                    var v29 float32 = float32(1e10)
                    fmt.Println(v13, v13)
                    _ = v29
                }
            }
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            if ((false || (!((!false) && (42 != 2)))) || true) {
                fmt.Println(v25, v25)
            } else {
                for fi31 := 0; fi31 < 18; fi31++ {
                    v32 := GenS1{F0: 127, F1: 0, F2: 0}
                    fmt.Println(v28, v25)
                    _ = fi31
                    _ = v32
                }
            }
            var v33 string = "\n\t"
            _ = fi30
            _ = v33
        }
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
