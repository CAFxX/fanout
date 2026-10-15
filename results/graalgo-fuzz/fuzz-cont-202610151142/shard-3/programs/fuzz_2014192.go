package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.1
    var v3 int32 = 261
    v4 := []int{32768, -9223372036854775808}
    var v5 bool = true
    var v6 int16 = 1000
    _ = len(v1)
    v7 := GenS1{F0: v5, F1: 0, F2: -5}
    var v8 float64 = float64(3.14159)
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float64 = 3.14159
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "foo bar":
        if false {
            if true {
                v2 = float64(3.14159)
                for fi11 := 0; fi11 < 1; fi11++ {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        if len(v4) > 0 {
                            v4[0] = (fi12 | -2147483648)
                        }
                        var v13 uint16 = 1
                        v13 = 42
                        var v14 int16 = int16(fi12)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v6 -= (v6 << ((18446744073709551615) & 63))
                    _ = fi11
                }
                for fi15 := 0; fi15 < 2; fi15++ {
                    _ = len(v1)
                    if v5 {
                        v2 = v2
                        var v16 bool = true
                        var v17 int64 = 9223372036854775807
                        _ = v16
                        _ = v17
                    } else {
                        if true {
                            v3 -= -2147483648
                            v5 = (!(!((!v5) || v5)))
                            var v18 uint8 = 128
                            _ = v18
                        }
                        var v19 bool = v5
                        _ = v19
                    }
                    fmt.Println(v8)
                    _ = fi15
                }
            } else {
                if v5 {
                    v20 := GenS1{F0: false, F1: 1, F2: 42}
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v1[1] = (fi21 | fi21)
                        var v22 int32 = v3
                        fmt.Println(v6)
                        _ = fi21
                        _ = v22
                    }
                    if v5 {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            for fi24 := 0; fi24 < 15; fi24++ {
                                fmt.Println(v3, v5, v8)
                                v8 = v8
                                v8 -= v2
                                switch (fi23 >> ((973) & 63)) {
                                case 128:
                                    fmt.Println(v6)
                                }
                                _ = fi24
                            }
                            _ = fi23
                        }
                        var v25 int32 = v3
                        var v26 uint8 = 255
                        _ = v25
                        _ = v26
                    } else {
                        var v27 int64 = 9223372036854775807
                        _ = v27
                    }
                }
            }
            var v28 uint8 = 42
            _ = v28
        } else {
            v29 := GenS1{F0: v5, F1: 42, F2: 1}
            switch 219 {
            case 9223372036854775807:
                v30 := GenS1{F0: false, F1: 0, F2: int8(v6)}
                _ = v30
            case 42:
                if (false && v5) {
                    var v31 uint8 = 2
                    _ = v31
                }
            case 2147483647:
                var v32 int = 398
                _ = v32
            default:
                if (42 >= 42) {
                    fmt.Println(v8, v5, v5)
                }
            }
            _ = v29
        }
    case "a":
        v6 -= -392
    }
    delete(v1, 42)
    var v33 float64 = (((v8 / v8) / -2.25) / v2)
    if (v5 && v5) {
        if (v5 || true) {
            var v34 int16 = (v6 / (((^(v6))) | 1))
            if (v5 || (("hello" + ("foo bar" + ("" + (("\n\t" + "hello") + (("\n\t" + "a") + "hello"))))) >= "")) {
                v35 := GenS1{F0: (v5 && v5), F1: 1, F2: int8(v3)}
                if (!v5) {
                    var v36 float64 = -2.25
                    _ = v36
                } else {
                    var v37 int16 = v34
                    for fi38 := 0; fi38 < 20; fi38++ {
                        if len(v4) > 0 {
                            v4[0] = fi38
                        }
                        _ = fi38
                    }
                    _ = v37
                }
                _ = v35
            }
            _ = v34
        } else {
            switch "" {
            case "foo bar":
                v2 += v2
            case "\n\t":
                v39 := GenS1{F0: (!v5), F1: 740, F2: 100}
                _ = v39
            default:
                v6 = 1000
            }
            var v40 float32 = (0.0 / -2.25)
            _ = v40
        }
        v5 = (!((2 > 42) || v5))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
