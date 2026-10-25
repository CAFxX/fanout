package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int32; F2 uint }
func genFn0() int {
    v6 := []int{-1, 32768, 42}
    switch "hello" {
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float32 = 0.0
            v8 = ((v8 - v8) / ((v8 + v8) - float32((1e10 / (float64((1e10 * (1.5 - v8))) - 1.5)))))
            _ = fi7
            _ = v8
        }
    case "\n\t":
        var v9 uint16 = 391
        _ = v9
    }
    var v10 string = "\n\t"
    fmt.Println(v10)
    _ = v10
    _ = v6
    return 2147483647
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]string)
    var v3 int8 = -73
    var v4 string = ("\n\t" + "\n\t")
    var v5 int = 857
    var v11 string = (v4 + (("\n\t" + "\n\t") + v4))
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 = "\n\t"
        _ = fi12
    }
    var v13 uint = 255
    v14 := GenS1{F0: v13, F1: -51, F2: 1}
    v2[v11] = v4
    switch v4 {
    case "foo bar":
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS2{F0: 255}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case "":
        v18 := GenS3{F0: 988, F1: v1, F2: v13}
        _ = v18
    case "hello":
        if ((42 > v13) && false) {
            switch 634 {
            case -1:
                v5 += v5
            case 809:
                var v19 float64 = 0.1
                _ = v19
            }
            v20 := GenS1{F0: (v13 - (v13 | (^(v13)))), F1: (v1 & (v1 * (v1 + (^(v1))))), F2: 1}
            _ = v20
        } else {
            if true {
                for fi21 := 0; fi21 < 1; fi21++ {
                    delete(v2, v11)
                    fmt.Println(v5)
                    var v22 uint8 = 42
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v11)
                    if (-488 >= -1) {
                        if (!true) {
                            v4 = v4
                            var v24 int8 = (v3 + 63)
                            switch "" {
                            case "a":
                                var v25 uint8 = 127
                                _ = v25
                            case "":
                                var v26 uint16 = 792
                                _ = v26
                            }
                            _ = v24
                        } else {
                            v1 += -2147483648
                        }
                        v27 := []int{(fi23 ^ 42), -292, (fi23 % ((0) | 1))}
                        _ = v27
                    } else {
                        switch (v4 + ("hello" + v11)) {
                        case "a":
                            v4 = ("" + "foo bar")
                        case "foo bar":
                            var v28 bool = ((0 > 42) || (954 <= 0))
                            _ = v28
                        case "":
                            v3 += v3
                        }
                    }
                    _ = fi23
                }
                v3 = v3
            }
            v29 := []int{v5, (-(v5)), 453}
            _ = v29
        }
    default:
        var v30 int64 = -9223372036854775808
        _ = v30
    }
    var v31 string = v11
    var v32 float32 = (0.0 * (-2.25 / 3.14159))
    delete(v2, ((v11 + v31) + (("foo bar" + "\n\t") + (v4 + v11))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
