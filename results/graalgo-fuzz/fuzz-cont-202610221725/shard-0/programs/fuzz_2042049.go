package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 bool = (2147483647 > 1)
    v2 := []string{("" + "\n\t")}
    var v3 int16 = 32767
    var v4 uint64 = 0
    var v5 uint32 = 831
    var v6 string = ""
    var v7 string = v6
    if len(v2) > 0 {
        v2[0] = (v7 + "")
    }
    if v1 {
        var v8 float64 = 0.0
        v8 += v8
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 string = "hello"
            var v11 bool = (2 > 655)
            v6 = ((v6 + (v7 + "foo bar")) + v6)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        if len(v2) > 0 {
            v2[0] = v6
        }
    }
    v4 = (v4 / ((18446744073709551615) | 1))
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            if false {
                v14 := GenS1{F0: 0, F1: 978}
                _ = v14
            }
            v15 := GenS1{F0: 1, F1: 42}
            switch v6 {
            case "hello":
                v16 := GenS1{F0: 1, F1: 1}
                _ = v16
            case "":
                var v17 string = v7
                if true {
                    if false {
                        var v18 uint8 = 128
                        _ = v18
                    }
                    v2 = append(v2, v6)
                    _ = len(v2)
                } else {
                    var v19 bool = true
                    v6 = v17
                    v4 += 0
                    _ = v19
                }
                _ = v17
            case "a":
                v3 += (v3 & 1000)
            default:
                switch int(v5) {
                case 128:
                    v20 := GenS1{F0: 726, F1: 0}
                    _ = v20
                case 436:
                    v1 = (v1 && ((v1 || true) || v1))
                default:
                    v21 := GenS1{F0: 1, F1: 65535}
                    _ = v21
                }
            }
            if v1 {
                fmt.Println(v6, v6, v3)
            } else {
                switch (fi13 | (fi12 >> ((1) & 63))) {
                case -603:
                    if len(v2) > 0 {
                        v2[0] = (v6 + "a")
                    }
                case -9223372036854775808:
                    v7 = (v6 + "hello")
                    v3 = v3
                case -659:
                    if (98 > 42) {
                        var v22 uint32 = (v5 % ((0) | 1))
                        if (!v1) {
                            v4 = (v4 | (v4 >> ((65535) & 63)))
                            v7 = v6
                        }
                        _ = v22
                    } else {
                        switch (fi12 << ((2) & 63)) {
                        case -427:
                            var v23 uint8 = 2
                            _ = v23
                        default:
                            if (false || (v5 != v5)) {
                                switch ((v6 + ("" + v7)) + ("" + "")) {
                                case "":
                                    var v24 uint8 = 107
                                    _ = v24
                                case "foo bar":
                                    var v25 float64 = (0.0 + -2.25)
                                    _ = v25
                                case "\n\t":
                                    var v26 int8 = 42
                                    _ = v26
                                default:
                                    v5 -= (v5 + v5)
                                }
                            }
                        }
                        v27 := GenS2{F0: v5, F1: 1}
                        var v28 uint8 = 0
                        _ = v27
                        _ = v28
                    }
                default:
                    if v1 {
                        switch "" {
                        case "hello":
                            v1 = (!v1)
                        case "\n\t":
                            v3 -= (v3 % ((42) | 1))
                        case "a":
                            var v29 uint64 = 42
                            _ = v29
                        }
                        fmt.Println(fi12, v1)
                        v6 = ((v6 + (v6 + "hello")) + v7)
                    }
                }
                if false {
                    v6 = v7
                    _ = len(v2)
                    v4 -= (v4 >> ((18446744073709551615) & 63))
                }
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                var v30 uint = 426
                _ = v30
            }
            if v1 {
                var v31 uint = 60
                var v32 int32 = 0
                _ = v31
                _ = v32
            }
            v1 = (fi13 == fi12)
            if (!(v4 == v4)) {
                switch 32768 {
                case -493:
                    var v33 uint64 = (v4 - 1)
                    _ = v33
                case -1:
                    var v34 uint = 42
                    _ = v34
                default:
                    var v35 float64 = ((0.1 * 1.5) + (3.14159 - 3.14159))
                    var v36 int32 = -805
                    v36 -= v36
                    var v37 int16 = v3
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v38 := GenS1{F0: 65535, F1: 1}
                _ = v38
            } else {
                fmt.Println(fi12, v5)
                v3 += v3
            }
            var v39 uint16 = 32767
            _ = fi13
            _ = v15
            _ = v39
        }
        if (0 > 41) {
            var v40 uint16 = 0
            _ = v40
        } else {
            fmt.Println(v7, v7)
            var v41 float64 = ((0.0 * float64(-2.25)) * -2.25)
            _ = v41
        }
        _ = fi12
    }
    for fi42 := 0; fi42 < 15; fi42++ {
        switch fi42 {
        case 4294967295:
            var v43 int32 = 42
            _ = v43
        }
        var v44 int = 32768
        _ = fi42
        _ = v44
    }
    if true {
        var v45 float64 = -2.25
        _ = v45
    }
    var v46 uint8 = 45
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v6, v7, v4)
    v4 = (v4 % (((v4 + (v4 >> ((199) & 63)))) | 1))
    _ = len(v2)
    var v47 int64 = 552
    switch (v6 + ((((v6 + "foo bar") + ("\n\t" + (v7 + v6))) + v6) + "a")) {
    case "a":
        fmt.Println(v46, v6)
    }
    v6 = (v6 + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
