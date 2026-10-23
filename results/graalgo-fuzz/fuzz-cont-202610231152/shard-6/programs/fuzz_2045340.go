package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (-2.25 + 3.14159)
    var v3 string = "hello"
    var v4 int64 = -262
    var v5 int64 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, v3, v5)
    v8 := GenS1{F0: 0}
    fmt.Println(v3, v5, v3)
    switch v3 {
    case "hello":
        v1["a"] = 2147483648
    case "a":
        v4 = -9223372036854775808
    case "":
        fmt.Println(v5, v4)
    default:
        if ((18446744073709551615 != 1) && false) {
            v9 := []int{}
            v10 := GenS1{F0: 1}
            _ = v10
            _ = v9
        }
    }
    if (-209 > 1) {
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(fi11)
            v5 -= (v4 >> ((42) & 63))
            v3 = v3
            _ = fi11
        }
    }
    v12 := []int{943}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v13 int = 4294967295
    if (v13 <= (v13 * 128)) {
        v3 = "a"
    }
    if (2147483648 == 1) {
        if true {
            if (-2.25 <= 0.1) {
                v3 = ("a" + v3)
            }
        } else {
            switch "" {
            case "\n\t":
                var v14 uint64 = 56
                _ = v14
            case "":
                v4 -= -1
            case "foo bar":
                var v15 int16 = -32768
                _ = v15
            }
            var v16 float64 = ((v2 - ((1.5 + v2) + float64(1.5))) / (((0.1 + -2.25) - (float64(1e10) - v2)) - (v2 - (v2 + v2))))
            fmt.Println(v5, v3, v3)
            _ = v16
        }
    } else {
        switch v3 {
        case "foo bar":
            v17 := GenS2{F0: v5, F1: 9223372036854775807}
            _ = v17
        default:
            var v18 bool = (!false)
            _ = v18
        }
        v4 = v4
    }
    fmt.Println(v2, v2, v4)
    v5 -= (v4 / ((42) | 1))
    switch ("" + v3) {
    case "foo bar":
        if len(v12) > 0 {
            v12[0] = (v13 << ((1) & 63))
        }
    case "\n\t":
        if true {
            if (false || false) {
                var v19 float32 = 1.5
                _ = v19
            }
            v20 := GenS1{F0: 1}
            var v21 uint = 18446744073709551615
            _ = v20
            _ = v21
        } else {
            switch v3 {
            case "a":
                switch ((((("a" + "") + v3) + v3) + v3) + (((v3 + "\n\t") + v3) + v3)) {
                case "foo bar":
                    var v22 string = ("\n\t" + (v3 + (("\n\t" + v3) + ((v3 + "hello") + v3))))
                    _ = v22
                case "\n\t":
                    if false {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            if len(v12) > 0 {
                                v12[0] = (fi23 / ((4294967295) | 1))
                            }
                            v2 = v2
                            _ = fi23
                        }
                        switch (v13 * 128) {
                        case -9223372036854775808:
                            v4 = (-(v4))
                        }
                        switch (v13 * (v13 / ((v13) | 1))) {
                        case -2147483648:
                            v2 += v2
                        case 865:
                            if true {
                                switch v13 {
                                case -2147483648:
                                    v4 = (v5 * -9223372036854775808)
                                case 128:
                                    v2 = (0.0 + v2)
                                case 4294967295:
                                    v3 = v3
                                default:
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        if true {
                                            if false {
                                                if false {
                                                    v13 = fi24
                                                    if true {
                                                        v25 := GenS3{F0: 1, F1: v3}
                                                        v2 = (-2.25 / v2)
                                                        _ = v25
                                                    }
                                                    var v26 int64 = -969
                                                    _ = v26
                                                } else {
                                                    fmt.Println(v2, v13)
                                                }
                                            }
                                            v27 := GenS4{F0: v3, F1: 1}
                                            _ = v27
                                        } else {
                                            v3 = v3
                                            fmt.Println(v3, v4)
                                        }
                                        _ = fi24
                                    }
                                }
                                var v28 uint64 = 242
                                var v29 float32 = ((1.5 - 0.1) / ((-2.25 + float32((v2 / float64((1e10 * (float32(v2) * 1e10)))))) - ((0.0 - (1.5 * ((3.14159 / ((-2.25 / float32(v2)) / 1.0)) - float32((((v2 * (v2 / float64(float32(v2)))) * v2) / (v2 / 0.1)))))) * 3.14159)))
                                _ = v28
                                _ = v29
                            }
                        }
                    } else {
                        if true {
                            for fi30 := 0; fi30 < 10; fi30++ {
                                var v31 uint = 620
                                _ = fi30
                                _ = v31
                            }
                            var v32 bool = true
                            _ = v32
                        }
                        v5 -= (v4 | (v5 << ((18446744073709551615) & 63)))
                    }
                }
            case "":
                v33 := GenS2{F0: (v5 * (v4 << ((293) & 63))), F1: (v5 >> ((305) & 63))}
                _ = v33
            case "foo bar":
                var v34 int = v13
                _ = v34
            }
            var v35 float64 = v2
            delete(v1, v3)
            _ = v35
        }
    case "hello":
        v2 = (v2 + (-2.25 + (float64((-2.25 / 0.1)) + v2)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
