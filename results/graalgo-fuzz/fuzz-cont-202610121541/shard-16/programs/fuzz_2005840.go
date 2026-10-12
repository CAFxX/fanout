package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int; F2 int }

func main() {
    v1 := []uint32{4294967295, 4294967295}
    var v2 int32 = -1
    var v3 float32 = 0.1
    var v4 uint16 = 32767
    v3 = (1e10 - v3)
    var v5 float32 = float32(3.14159)
    if (1e10 > 0.0) {
        v1 = append(v1, 0)
        v6 := GenS1{F0: "hello", F1: 42}
        _ = v6
    }
    v3 += (v3 * (1e10 - 3.14159))
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int = fi7
        if (v8 < 32768) {
            if (287 >= -9223372036854775808) {
                switch ("" + "\n\t") {
                case "foo bar":
                    var v9 string = "foo bar"
                    _ = v9
                case "hello":
                    var v10 float64 = -2.25
                    _ = v10
                }
                switch v8 {
                case -9223372036854775808:
                    switch (("foo bar" + "a") + "foo bar") {
                    case "a":
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v3 -= (3.14159 - (v5 * v5))
                            v12 := GenS2{F0: -940, F1: (!(true && false)), F2: 65535}
                            fmt.Println(fi7, v2)
                            _ = fi11
                            _ = v12
                        }
                    case "hello":
                        if true {
                            v8 = fi7
                        } else {
                            var v13 float32 = (v3 - v3)
                            if (!(true || true)) {
                                fmt.Println(fi7, v2)
                                v4 = (v4 << ((65535) & 63))
                                v1 = append(v1, 4294967295)
                            }
                            v8 += 0
                            _ = v13
                        }
                    }
                default:
                    if ((float64(v3) > 0.1) || false) {
                        if (!((^(v8)) > (fi7 - (fi7 / (((v8 << ((72) & 63))) | 1))))) {
                            v8 += fi7
                        }
                        v4 = v4
                        v2 += -947
                    }
                }
            }
        }
        v14 := GenS2{F0: 42, F1: false, F2: 255}
        _ = fi7
        _ = v14
        _ = v8
    }
    fmt.Println(v5, v3)
    v2 -= (v2 / ((-450) | 1))
    if true {
        if (false && (!(!false))) {
            v2 = (^(v2))
        }
        var v15 int16 = 42
        _ = v15
    } else {
        _ = len(v1)
    }
    v3 = ((v5 / v3) - float32((-2.25 * (-2.25 * ((((-2.25 + 1.5) / (-2.25 - (float64(0.1) / 1.5))) - ((float64((((v5 - 0.0) / v3) * v3)) / -2.25) * -2.25)) / 1.0)))))
    var v16 int32 = (v2 - (v2 % ((v2) | 1)))
    v4 = v4
    v3 += v5
    fmt.Println(v5, v2, v4)
    if true {
        switch 807 {
        case 847:
            fmt.Println(v2, v2, v5)
        case 2147483647:
            var v17 uint32 = 2147483647
            _ = v17
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v18 := GenS3{F0: -1}
        _ = v18
    }
    if (-1 == -40) {
        v2 = 142
        var v19 uint = 255
        switch "hello" {
        case "hello":
            for fi20 := 0; fi20 < 14; fi20++ {
                if (1 >= 844) {
                    _ = len(v1)
                    var v21 uint8 = 224
                    _ = v21
                } else {
                    v3 -= v5
                    v22 := GenS4{F0: 42, F1: (fi20 | fi20), F2: fi20}
                    _ = v22
                }
                if false {
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    var v23 int = fi20
                    fmt.Println(v4)
                    _ = v23
                }
                v16 = v16
                _ = fi20
            }
        case "\n\t":
            var v24 int16 = 32767
            _ = v24
        case "foo bar":
            v25 := GenS3{F0: 833}
            _ = v25
        default:
            if false {
                var v26 int = 1
                v4 += v4
                _ = v26
            } else {
                if len(v1) > 0 {
                    v1[0] = 705
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int = -516
                    _ = fi27
                    _ = v28
                }
            }
        }
        _ = v19
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            fmt.Println(v16, v4, v4)
            if false {
                v3 += (v5 - (1e10 / 1.0))
                var v30 uint = 4294967295
                _ = v30
            } else {
                var v31 uint32 = 4294967295
                var v32 float32 = 1.5
                for fi33 := 0; fi33 < 15; fi33++ {
                    v34 := GenS3{F0: 1000}
                    v4 += (v4 & (v4 ^ v4))
                    switch (^(fi33)) {
                    case 32768:
                        if (!(!(false && true))) {
                            v35 := GenS1{F0: (("" + (("foo bar" + ("hello" + ("\n\t" + "\n\t"))) + ((("\n\t" + "\n\t") + "foo bar") + ("" + ("foo bar" + "foo bar"))))) + "hello"), F1: 42}
                            var v36 uint64 = 4294967295
                            _ = v35
                            _ = v36
                        } else {
                            var v37 uint8 = 2
                            if false {
                                fmt.Println(v37)
                                var v38 uint8 = v37
                                var v39 bool = (v38 > (^(v37)))
                                var v40 int16 = -1
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                v2 += (v16 - 2147483647)
                                switch "\n\t" {
                                case "a":
                                    var v41 uint64 = 738
                                    _ = v41
                                }
                                v4 += 0
                            }
                            v42 := GenS3{F0: 769}
                            var v43 uint8 = (v37 | v37)
                            _ = v37
                            _ = v42
                            _ = v43
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        fmt.Println(v3, v5, v4)
    }
    v5 += (((v5 * -2.25) / (v5 / v3)) / 3.14159)
    fmt.Println(v5, v3, v2)
    v3 = (v5 / v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
