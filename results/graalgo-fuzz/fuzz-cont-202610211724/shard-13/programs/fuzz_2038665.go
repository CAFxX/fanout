package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 uint64 = 788
    var v2 uint16 = 42
    v3 := make(map[string]int)
    var v4 string = (((("\n\t" + "a") + ("\n\t" + "")) + "\n\t") + "\n\t")
    var v5 float32 = (((0.1 / 3.14159) + (-2.25 * (0.1 / -2.25))) - 1e10)
    v2 = v2
    if false {
        var v6 float32 = float32(float64(3.14159))
        fmt.Println(v5, v1)
        v7 := []int{0, 32767}
        _ = v6
        _ = v7
    } else {
        if (!(!false)) {
            var v8 int16 = -1
            _ = v8
        } else {
            v2 = v2
            var v9 uint32 = 560
            _ = v9
        }
    }
    v10 := GenS1{F0: 2, F1: v2, F2: 22}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 = v5
            _ = fi12
        }
        _ = fi11
    }
    var v13 uint = 327
    v14 := []int{-2147483648, -9223372036854775808, -874, -1}
    if true {
        if (2 != 2147483647) {
            if true {
                v2 = 940
                v4 = ((v4 + (v4 + "\n\t")) + ("foo bar" + ("hello" + "foo bar")))
                if false {
                    if true {
                        v15 := GenS1{F0: 2, F1: v2, F2: 2}
                        v13 += (v13 - 4294967295)
                        _ = v15
                    } else {
                        var v16 float64 = ((1.5 - (1e10 / 1e10)) + (((0.1 * (-2.25 + 1e10)) - 3.14159) + 1e10))
                        v4 = ""
                        fmt.Println(v1)
                        _ = v16
                    }
                    v2 = (v2 % (((^(v2))) | 1))
                }
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    if false {
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v2 += 255
                            switch (fi18 & 695) {
                            case -160:
                                v19 := GenS1{F0: 105, F1: (^(v2)), F2: 42}
                                _ = v19
                            case -9223372036854775808:
                                var v20 string = v4
                                _ = v20
                            }
                            var v21 int8 = 0
                            switch v4 {
                            case "":
                                v4 = ("hello" + v4)
                            case "\n\t":
                                var v22 int = (-(fi17))
                                _ = v22
                            default:
                                if ((42 < 42) || false) {
                                    var v23 int = 4294967295
                                    var v24 float32 = (float32(float64(((0.0 + (1e10 - (v5 - (v5 * v5)))) - ((0.1 * 0.1) - (float32(-2.25) - (0.1 - v5)))))) * v5)
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = fi18
                            _ = v21
                        }
                        v5 += float32(0.0)
                    }
                    fmt.Println(v4, v4)
                    _ = fi17
                }
            }
        }
    }
    var v25 int16 = -273
    var v26 int64 = int64(v25)
    var v27 int16 = 510
    var v28 int64 = (v26 & 9223372036854775807)
    v3[v4] = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
