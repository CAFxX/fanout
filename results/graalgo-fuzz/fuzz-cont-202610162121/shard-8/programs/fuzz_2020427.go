package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []int{96}
    var v2 int64 = -1
    var v3 int16 = 42
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = v4
    var v6 float32 = 0.1
    fmt.Println(v5, v5)
    if (4294967295 <= 42) {
        if false {
            if (true && (128 <= 32768)) {
                var v7 int32 = 42
                switch 970 {
                case -1:
                    if (-1 > -128) {
                        v7 += v7
                        v5 = 0
                        var v8 float64 = 0.1
                        _ = v8
                    } else {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            switch (("foo bar" + "\n\t") + ("foo bar" + "")) {
                            case "\n\t":
                                var v10 int = (fi9 & 55)
                                _ = v10
                            }
                            fmt.Println(v6)
                            _ = fi9
                        }
                        v1 = append(v1, 2147483647)
                        fmt.Println(v6)
                    }
                default:
                    var v11 float32 = -2.25
                    _ = v11
                }
                fmt.Println(v3, v2, v4)
                _ = v7
            }
        }
        var v12 int64 = (v2 | (-(v2)))
        _ = v12
    }
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 3; fi14++ {
        if (false || (!((v5 + v5) >= (^(v4))))) {
            for fi15 := 0; fi15 < 12; fi15++ {
                v13 = "hello"
                v16 := GenS1{F0: 127, F1: (v3 ^ (v3 - (^(v3)))), F2: 870}
                if true {
                    var v17 float64 = 0.0
                    _ = v17
                }
                _ = len(v1)
                if false {
                    v18 := GenS2{F0: 1}
                    _ = v18
                }
                var v19 int16 = v3
                v20 := GenS3{F0: 42}
                var v21 int = (fi14 - fi14)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
                _ = v16
                _ = v19
                _ = v20
                _ = v21
            }
            if (false && true) {
                v5 += v4
            }
            var v22 int = -2147483648
            _ = v22
        }
        v4 = v4
        _ = fi14
    }
    switch (v13 + v13) {
    case "":
        fmt.Println(v13, v2)
    default:
        v13 = v13
    }
    fmt.Println(v5, v6, v3)
    v23 := GenS1{F0: 100, F1: (v3 << ((1) & 63)), F2: 0}
    v1 = append(v1, -2147483648)
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 uint16 = 255
        for fi26 := 0; fi26 < 10; fi26++ {
            v2 = 239
            v3 = v3
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v3)
    v13 = ("\n\t" + "a")
    if (true || false) {
        if (-977 != -2147483648) {
            for fi27 := 0; fi27 < 20; fi27++ {
                if true {
                    var v28 int8 = int8(fi27)
                    v2 = (^(v2))
                    if (!(((0.1 + 0.1) <= (1.5 - ((0.1 * 0.1) / -2.25))) || (!((!true) && true)))) {
                        if true {
                            var v29 string = v13
                            v4 -= (v4 - 18446744073709551615)
                            _ = v29
                        }
                    } else {
                        var v30 int32 = 42
                        _ = v30
                    }
                    _ = v28
                } else {
                    var v31 int32 = 0
                    v1 = append(v1, (^(fi27)))
                    _ = v31
                }
                if true {
                    v32 := GenS1{F0: -1, F1: (v3 / (((v3 | (v3 & (v3 * (v3 * 1000))))) | 1)), F2: 0}
                    _ = v32
                } else {
                    var v33 int8 = 42
                    fmt.Println(v13, v6)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v33
                }
                var v34 uint16 = 255
                _ = fi27
                _ = v34
            }
            _ = len(v1)
        } else {
            if len(v1) > 0 {
                v1[0] = 32768
            }
            if (v13 <= (v13 + "a")) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v35 uint = 255
                v4 += 499
                _ = v35
            }
            _ = len(v1)
        }
    }
    var v36 int = -360
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
