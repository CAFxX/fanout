package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]int)
    v3 := []string{("\n\t" + ("a" + ("" + ((((("a" + "hello") + "a") + "") + "") + (("a" + ("\n\t" + (("hello" + "hello") + ((("a" + "\n\t") + "foo bar") + "\n\t")))) + "hello")))))}
    var v4 int32 = -1
    var v5 int = -958
    var v6 int8 = 1
    if true {
        switch "hello" {
        case "hello":
            v5 = 0
        case "":
            v7 := GenS1{F0: -793, F1: 1}
            _ = v7
        }
        v8 := GenS1{F0: -9223372036854775808, F1: 679}
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v2)
    var v9 int32 = 572
    v10 := GenS2{F0: 128, F1: 99}
    switch 350 {
    case 0:
        if (!(1 >= 0)) {
            switch v5 {
            case 1:
                for fi11 := 0; fi11 < 6; fi11++ {
                    if ((!true) || (false || (true && false))) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v12 := GenS1{F0: 42, F1: 18446744073709551615}
                        _ = v12
                    } else {
                        if true {
                            if (-99 <= -1) {
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    var v13 uint16 = 553
                                    _ = v13
                                case "hello":
                                    for fi14 := 0; fi14 < 1; fi14++ {
                                        v4 += -2147483648
                                        var v15 uint = 42
                                        _ = fi14
                                        _ = v15
                                    }
                                case "":
                                    v16 := GenS1{F0: 861, F1: 1}
                                    _ = v16
                                default:
                                    if (0 <= 377) {
                                        if (!((1e10 > (((v1 - v1) / (v1 + v1)) * -2.25)) || true)) {
                                            fmt.Println(v9, v1, v9)
                                            delete(v2, "")
                                        }
                                        v5 = int(v4)
                                    } else {
                                        var v17 bool = true
                                        _ = v17
                                    }
                                }
                            }
                            v6 = (v6 / ((42) | 1))
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                        } else {
                            for fi18 := 0; fi18 < 16; fi18++ {
                                fmt.Println(v9)
                                var v19 uint32 = 65535
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                                _ = fi18
                                _ = v19
                            }
                        }
                    }
                    _ = fi11
                }
            case 32767:
                v20 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                _ = v20
            }
            var v21 int32 = int32(v5)
            _ = v21
        }
    default:
        switch v5 {
        case -359:
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v4, v1)
                    v1 = ((v1 / float64(1.5)) / ((v1 + v1) + 1.5))
                    _ = fi23
                }
                v5 -= -2147483648
                _ = fi22
            }
        default:
            fmt.Println(v9, v6)
        }
    }
    var v24 bool = (true && (true || (v4 > (v9 ^ v9))))
    delete(v2, "")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
