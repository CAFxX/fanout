package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 uint8 }
func genFn0(p0 int64) int8 {
    p0 = 350
    var v6 int = 521
    v7 := []int{9223372036854775807, 497, (v6 * v6)}
    p0 -= p0
    var v8 uint32 = 736
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 int64 = -255
    var v2 int = 32767
    v3 := make(map[uint32]uint32)
    var v4 int16 = 1
    v5 := make(map[uint32]uint32)
    var v9 uint32 = 1
    for fi10 := 0; fi10 < 14; fi10++ {
        if (false && true) {
            v9 = v9
            if false {
                if true {
                    if false {
                        if false {
                            v9 = 282
                            for fi11 := 0; fi11 < 2; fi11++ {
                                v4 = v4
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    var v12 float32 = -2.25
                                    _ = v12
                                default:
                                    v1 = (v1 % ((v1) | 1))
                                }
                                _ = fi11
                            }
                            fmt.Println(v2, v4, v9)
                        } else {
                            var v13 uint16 = 42
                            v2 = v2
                            _ = v13
                        }
                        var v14 uint64 = 1
                        var v15 int8 = 42
                        _ = v14
                        _ = v15
                    }
                }
            } else {
                var v16 float64 = float64((1e10 + (3.14159 + (-2.25 - (0.0 * 3.14159)))))
                _ = v16
            }
        }
        v17 := GenS1{F0: 0}
        v4 = -709
        _ = fi10
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 string = "a"
        var v20 string = "a"
        switch 42 {
        case 644:
            for fi21 := 0; fi21 < 10; fi21++ {
                v9 = v9
                v20 = ("a" + v19)
                _ = fi21
            }
        default:
            for fi22 := 0; fi22 < 8; fi22++ {
                var v23 int = fi18
                _ = fi22
                _ = v23
            }
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v24 int16 = v4
    for fi25 := 0; fi25 < 16; fi25++ {
        v1 = (v1 + v1)
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 int32 = 2147483647
            switch "" {
            case "":
                v1 += v1
            case "hello":
                v28 := GenS1{F0: (v27 ^ 645)}
                _ = v28
            default:
                v5[v9] = v9
            }
            v29 := []int{32768, 0}
            _ = fi26
            _ = v27
            _ = v29
        }
        _ = fi25
    }
    if true {
        switch "hello" {
        case "":
            if ((0.1 / 1.0) <= (((0.1 + (1e10 + (((((0.1 * float64(1.5)) * (((1.5 - (float64((1e10 / (0.0 + -2.25))) - 1.5)) - (3.14159 - 3.14159)) + ((-2.25 / 3.14159) / 1.0))) * 0.1) / 1e10) / 0.1))) - ((1.5 / (1e10 + float64((-2.25 + 0.0)))) + ((0.0 - (0.0 - float64(0.0))) - (((0.1 - ((3.14159 + (3.14159 / float64((((1.5 + 3.14159) + ((float32(1e10) - 1.5) - 3.14159)) - 1.5)))) - 1e10)) * 0.0) + (-2.25 / 1.0))))) * (3.14159 * (-2.25 / 1e10)))) {
                v30 := GenS2{F0: 9223372036854775807, F1: 4294967295, F2: 42}
                v1 = (v1 & v1)
                v31 := GenS2{F0: (-(v2)), F1: 1, F2: 128}
                _ = v30
                _ = v31
            }
        case "a":
            switch ("a" + ("\n\t" + "hello")) {
            case "foo bar":
                var v32 uint32 = 2
                _ = v32
            case "\n\t":
                v33 := []int{(v2 % (((v2 >> ((65535) & 63))) | 1))}
                _ = v33
            }
        default:
            fmt.Println(v2)
        }
    }
    v2 -= 127
    v34 := GenS2{F0: v2, F1: 65535, F2: 105}
    _ = len(v3)
    fmt.Println(v9)
    fmt.Println(v1, v4, v24)
    v35 := GenS2{F0: v2, F1: 65535, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
