package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = 42
    v2 := []uint32{1, 42, 4294967295}
    var v3 int64 = 1
    var v4 uint16 = 42
    var v5 string = "foo bar"
    var v6 bool = (!false)
    _ = len(v2)
    var v7 int8 = 100
    var v8 uint32 = 42
    v9 := GenS1{F0: v4, F1: (^(v7))}
    var v10 uint16 = v4
    v3 += int64(v10)
    switch (((v5 + v5) + v5) + v5) {
    case "a":
        switch -2147483648 {
        case 724:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint8 = 1
                _ = fi11
                _ = v12
            }
        case 9223372036854775807:
            switch "" {
            case "":
                v13 := GenS1{F0: (v10 << ((1) & 63)), F1: (v7 % (((v7 / ((v7) | 1))) | 1))}
                _ = v13
            case "hello":
                v2 = append(v2, v8)
            case "foo bar":
                if v6 {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v1 = v1
                        _ = fi14
                    }
                    fmt.Println(v8, v8)
                    var v15 int16 = (v1 * -32768)
                    _ = v15
                }
            default:
                switch 42 {
                case -1:
                    v1 = v1
                }
            }
        default:
            var v16 uint64 = 193
            _ = v16
        }
    }
    if false {
        v6 = v6
        switch v5 {
        case "":
            if (3.14159 != 1.5) {
                v5 = v5
            } else {
                switch 2147483647 {
                case 236:
                    if ((1.5 - 0.1) > ((-2.25 * 1e10) + 1.5)) {
                        v4 = (v4 & (^(v10)))
                    }
                default:
                    v8 = 260
                }
                var v17 int16 = (-(v1))
                fmt.Println(v5, v6, v17)
                _ = v17
            }
        case "\n\t":
            if v6 {
                if len(v2) > 0 {
                    v2[0] = v8
                }
                var v18 bool = (!(!false))
                _ = v18
            } else {
                _ = len(v2)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    if true {
                        v20 := GenS2{F0: (v8 / (((v8 / ((v8) | 1))) | 1))}
                        v7 += (v7 * v7)
                        _ = v20
                    } else {
                        v4 = (v4 % ((42) | 1))
                        v8 = 0
                        if (!((v4 - (v4 << ((65535) & 63))) >= (^(v4)))) {
                            fmt.Println(v7, v3)
                        } else {
                            v5 = v5
                            if v6 {
                                v21 := GenS1{F0: (v10 - v4), F1: v7}
                                _ = v21
                            }
                        }
                    }
                    _ = fi19
                }
            }
        case "foo bar":
            fmt.Println(v8, v4)
        default:
            if v6 {
                v6 = (!(v6 || (v1 != (^(v1)))))
            }
        }
        v4 += v10
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    var v22 float64 = 1.5
    v8 += v8
    if (!false) {
        if (4294967295 != 0) {
            if (0.0 >= 1.5) {
                switch (v5 + (v5 + v5)) {
                case "":
                    v2 = append(v2, 362)
                case "\n\t":
                    v23 := GenS2{F0: 0}
                    _ = v23
                case "foo bar":
                    v10 -= v10
                default:
                    fmt.Println(v5, v22)
                }
                fmt.Println(v22, v3, v5)
                switch (((v5 + ((v5 + (v5 + (v5 + v5))) + ("\n\t" + v5))) + "hello") + "foo bar") {
                case "":
                    var v24 float64 = 1e10
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS2{F0: v8}
                    _ = fi25
                    _ = v26
                }
                var v27 uint16 = (v10 / ((v10) | 1))
                var v28 uint64 = 4294967295
                _ = v27
                _ = v28
            }
        }
        v1 += -32768
    }
    var v29 int8 = -41
    v8 = (v8 ^ 58)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
