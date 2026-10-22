package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    var v1 float64 = (-2.25 * (0.0 - float64(float32(0.1))))
    var v2 int64 = 9223372036854775807
    var v3 bool = ((716 >= 2) || true)
    var v4 uint8 = 135
    if (v3 || (18446744073709551615 < 820)) {
        if false {
            var v5 uint = 600
            var v6 int64 = v2
            fmt.Println(v3, v6)
            _ = v5
            _ = v6
        }
    } else {
        v7 := []int{827}
        _ = v7
    }
    var v8 uint64 = 703
    fmt.Println(v1, v8, v2)
    fmt.Println(v1)
    v9 := []int{260, 32768}
    if (((-2.25 <= v1) || (127 > 2147483647)) && (true && true)) {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v4)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v2 = (v2 - int64(v8))
            _ = fi10
        }
    }
    var v11 uint16 = 255
    var v12 uint = 1
    var v13 int = -418
    v11 = 0
    fmt.Println(v4, v11)
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            if (false || (3.14159 <= (float32((v1 + (v1 - v1))) - 0.0))) {
                fmt.Println(v11, v1, v3)
            }
            _ = fi14
        }
    } else {
        v12 += 255
        for fi15 := 0; fi15 < 6; fi15++ {
            v3 = v3
            _ = fi15
        }
        if ((0 >= 42) || (v3 && false)) {
            if ((85 >= v4) && (v4 != v4)) {
                v2 = v2
                v3 = false
                var v16 string = "a"
                _ = v16
            } else {
                if (false || ((v2 | -1) > v2)) {
                    switch "" {
                    case "a":
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v18 := GenS1{F0: 391, F1: 18446744073709551615, F2: v11}
                            var v19 float32 = (float32(float64((3.14159 + 1e10))) + ((0.0 / 1.0) + 3.14159))
                            switch (("" + ("hello" + "a")) + "foo bar") {
                            case "\n\t":
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            default:
                                var v20 uint64 = 2
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        switch v13 {
                        case -1:
                            v12 -= 255
                        case 2147483647:
                            var v21 uint16 = (v11 - 0)
                            _ = v21
                        }
                    default:
                        v13 += 128
                    }
                    var v22 string = ""
                    var v23 float64 = (1.5 + 1.5)
                    _ = v22
                    _ = v23
                } else {
                    if true {
                        v12 = (^(v12))
                        if v3 {
                            var v24 int64 = v2
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 string = "hello"
                                _ = fi25
                                _ = v26
                            }
                        }
                    }
                }
                var v27 int64 = (v2 | (^(v2)))
                _ = v27
            }
            var v28 int = 2147483647
            _ = v28
        } else {
            for fi29 := 0; fi29 < 10; fi29++ {
                v9 = append(v9, 173)
                v30 := GenS2{F0: (^(v11)), F1: fi29}
                _ = fi29
                _ = v30
            }
            if len(v9) > 0 {
                v9[0] = v13
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
