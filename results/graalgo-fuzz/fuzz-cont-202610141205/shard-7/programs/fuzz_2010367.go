package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 42
    var v3 uint8 = uint8(v2)
    v4 := make(map[int]int)
    var v5 int8 = 50
    v6 := make(map[string]int)
    var v7 float64 = 1.5
    if (176 > -260) {
        var v8 float32 = (float32(v7) + (-2.25 / 1e10))
        switch 32768 {
        case -2147483648:
            if true {
                fmt.Println(v7)
                var v9 int = 42
                _ = v9
            } else {
                if (669 >= -32768) {
                    if true {
                        var v10 uint16 = v2
                        _ = v10
                    } else {
                        var v11 bool = false
                        if true {
                            fmt.Println(v11, v7, v8)
                            var v12 uint16 = (v2 | (v2 % (((v2 << ((18446744073709551615) & 63))) | 1)))
                            v13 := v6["hello"]
                            _ = v12
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 1; fi14++ {
                                v4[(^(fi14))] = (fi14 & 1)
                                var v15 string = "foo bar"
                                for fi16 := 0; fi16 < 11; fi16++ {
                                    var v17 int32 = -1
                                    _ = fi16
                                    _ = v17
                                }
                                _ = fi14
                                _ = v15
                            }
                        }
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v19 := GenS1{F0: 32767, F1: -1}
                            _ = fi18
                            _ = v19
                        }
                        _ = v11
                    }
                    v20 := v6["foo bar"]
                    var v21 string = ""
                    _ = v20
                    _ = v21
                }
                v22 := GenS1{F0: -521, F1: -2147483648}
                _ = v22
            }
        default:
            v23 := GenS1{F0: int16(v2), F1: -505}
            _ = v23
        }
        v6[("hello" + "\n\t")] = 127
        _ = v8
    }
    v7 = v7
    v24 := []int{127, 792}
    v25 := GenS2{F0: (!true), F1: 2147483647, F2: v2}
    var v26 int32 = -1
    var v27 uint16 = 32767
    switch ("\n\t" + "") {
    case "\n\t":
        v5 += 127
    default:
        if (1 > 612) {
            var v28 float32 = 1.5
            switch -9223372036854775808 {
            case 128:
                fmt.Println(v28, v1)
            default:
                v7 = (3.14159 / 1.5)
            }
            _ = v28
        }
    }
    _ = len(v6)
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int64 = 9223372036854775807
        var v31 bool = ((uint(v3) >= 65535) || true)
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 float32 = 0.0
    v24 = append(v24, int(v27))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
