package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint = 1
    v2 := []string{"\n\t", ("\n\t" + "foo bar")}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 65535
    var v5 uint16 = (v4 % (((^(v4))) | 1))
    fmt.Println(v1, v3, v4)
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v5 = uint16(v1)
            if true {
                _ = len(v2)
                if len(v2) > 0 {
                    v2[0] = ("" + "")
                }
            }
            _ = fi7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            if (-1 == 32767) {
                if false {
                    v9 := GenS1{F0: "a", F1: 4294967295, F2: -2147483648}
                    _ = v9
                }
                if (((-9223372036854775808 == (^(fi8))) || true) || false) {
                    var v10 int16 = -32768
                    _ = v10
                } else {
                    fmt.Println(v1)
                }
                var v11 int8 = 127
                _ = v11
            }
            _ = fi8
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v5 -= 65535
        _ = fi12
    }
    var v13 int8 = -1
    var v14 int = 914
    switch 0 {
    case 0:
        if false {
            if false {
                switch -9223372036854775808 {
                case 583:
                    v15 := GenS1{F0: ("a" + ""), F1: (v3 + (v3 / ((v3) | 1))), F2: 0}
                    _ = v15
                case 4294967295:
                    if false {
                        var v16 float64 = (float64(-2.25) / -2.25)
                        _ = v16
                    }
                case -514:
                    v13 += (v13 / ((v13) | 1))
                default:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        switch ("" + (("" + "\n\t") + "foo bar")) {
                        case "\n\t":
                            fmt.Println(v3, v3)
                        case "a":
                            if (!((-233 <= 27) && false)) {
                                fmt.Println(v13, v14, v1)
                                var v18 string = "\n\t"
                                v5 += (v5 ^ 32767)
                                _ = v18
                            } else {
                                v1 = (v1 << (((v1 % (((^(v1))) | 1))) & 63))
                                v19 := GenS2{F0: "hello"}
                                if (!true) {
                                    fmt.Println(v14)
                                    v4 -= v5
                                }
                                _ = v19
                            }
                        case "foo bar":
                            v13 -= v13
                        default:
                            if false {
                                v1 = 2
                                var v20 uint32 = 2147483647
                                v2 = append(v2, "hello")
                                _ = v20
                            }
                        }
                        var v21 uint16 = (v5 + (v5 * (v5 ^ (v4 * 255))))
                        _ = fi17
                        _ = v21
                    }
                }
                v4 = v4
            }
            v2 = append(v2, "foo bar")
        } else {
            v5 -= (v4 * v5)
        }
    default:
        if (!false) {
            var v22 string = "hello"
            v4 = 67
            v3 = 1
            _ = v22
        } else {
            v23 := GenS1{F0: "hello", F1: v3, F2: -828}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v23
        }
    }
    v3 -= v3
    var v24 float32 = (1e10 + 1e10)
    if (true || false) {
        v5 -= (^(v5))
        var v25 int = (^(v14))
        fmt.Println(v5, v4, v3)
        _ = v25
    }
    if (false || (true && (v24 < v24))) {
        v24 = (v24 * (v24 + v24))
        if len(v2) > 0 {
            v2[0] = ("a" + "foo bar")
        }
        v5 -= v5
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 uint16 = (^(v5))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi27
                _ = v28
            }
            var v29 uint32 = 42
            var v30 float64 = 1.5
            _ = fi26
            _ = v29
            _ = v30
        }
        switch "" {
        case "":
            if false {
                var v31 float64 = 0.1
                switch (v14 | (v14 * 128)) {
                case 32767:
                    var v32 bool = (false || false)
                    _ = v32
                case 183:
                    v33 := GenS3{F0: 0}
                    _ = v33
                }
                _ = v31
            }
        default:
            v34 := GenS1{F0: "hello", F1: (^(v3)), F2: -761}
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
