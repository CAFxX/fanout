package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 string }
func genFn0() uint16 {
    switch "foo bar" {
    case "foo bar":
        var v7 uint8 = 0
        _ = v7
    case "\n\t":
        var v8 string = ""
        _ = v8
    default:
        v9 := GenS1{F0: 1, F1: (((("hello" + "foo bar") + "foo bar") > "a") && true), F2: 2}
        _ = v9
    }
    switch 0 {
    case -2147483648:
        var v10 int32 = 7
        _ = v10
    case 4294967295:
        var v11 float64 = (1.5 - ((1.5 + 3.14159) * (0.0 + (3.14159 + float64(1e10)))))
        _ = v11
    }
    return 32767
}

func main() {
    var v1 float32 = float32((float64(-2.25) - 0.1))
    var v2 uint32 = 65535
    var v3 uint16 = 464
    v4 := make(map[int]string)
    var v5 bool = (float64((float32(1e10) * (v1 * float32(3.14159)))) <= -2.25)
    var v6 uint16 = (v3 + (v3 % ((962) | 1)))
    var v12 uint64 = 0
    var v13 int32 = 2147483647
    if ((v5 && ((!v5) || true)) || false) {
        v2 += v2
        var v14 uint = 110
        var v15 string = ((("a" + "foo bar") + "a") + ("hello" + (("hello" + "hello") + "hello")))
        _ = v14
        _ = v15
    } else {
        if v5 {
            v16 := []int{}
            v17 := GenS1{F0: 2, F1: v5, F2: (v12 - v12)}
            fmt.Println(v13, v3, v13)
            _ = v16
            _ = v17
        } else {
            var v18 int = -2147483648
            switch "hello" {
            case "\n\t":
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS1{F0: 128, F1: true, F2: (v12 & 42)}
                    var v21 int64 = -9223372036854775808
                    if (v5 && v5) {
                        _ = len(v4)
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "":
                v22 := GenS1{F0: 2, F1: v5, F2: (^(v12))}
                _ = v22
            default:
                if (!(0 < 127)) {
                    v5 = (-32768 != 952)
                }
            }
            v23 := GenS1{F0: 72, F1: ((0 == 1) && (-1 <= -7)), F2: (^(v12))}
            _ = v18
            _ = v23
        }
    }
    v13 = (v13 ^ v13)
    switch "\n\t" {
    case "\n\t":
        v12 -= 203
    case "foo bar":
        switch "a" {
        case "hello":
            v1 = 0.0
        default:
            switch "foo bar" {
            case "":
                switch 128 {
                case 32768:
                    var v24 uint64 = (v12 ^ v12)
                    _ = v24
                default:
                    var v25 string = "hello"
                    _ = v25
                }
            default:
                switch 275 {
                case 32768:
                    if (9223372036854775807 == 0) {
                        if true {
                            for fi26 := 0; fi26 < 16; fi26++ {
                                var v27 uint = 358
                                v28 := []int{-1, (fi26 - fi26)}
                                v27 += 65535
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        } else {
                            if ((!v5) || false) {
                                v29 := []int{127, 42}
                                _ = v29
                            }
                            if v5 {
                                switch "foo bar" {
                                case "":
                                    v6 = v6
                                case "\n\t":
                                    var v30 int = 281
                                    _ = v30
                                }
                                var v31 uint8 = 127
                                _ = v31
                            } else {
                                v32 := GenS2{F0: false, F1: ""}
                                _ = v32
                            }
                        }
                    }
                case -450:
                    if v5 {
                        var v33 int = -519
                        _ = v33
                    }
                default:
                    v34 := v4[32768]
                    _ = v34
                }
            }
        }
    default:
        v2 = (v2 | v2)
    }
    delete(v4, 32768)
    v35 := GenS1{F0: 127, F1: (true || v5), F2: v12}
    v4[2147483647] = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
