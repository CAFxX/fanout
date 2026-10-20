package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 int16 }

func main() {
    v1 := []string{"\n\t"}
    var v2 uint8 = 0
    v3 := make(map[string]int)
    v4 := []uint32{42, 65535}
    if false {
        if (2147483648 > 2147483647) {
            var v5 string = ("a" + "a")
            v5 = (((v5 + v5) + "a") + v5)
            v5 = "hello"
            _ = v5
        } else {
            delete(v3, ((("a" + (("\n\t" + ("foo bar" + "foo bar")) + "a")) + "") + (((("\n\t" + "foo bar") + ((("a" + "foo bar") + "") + ("\n\t" + ""))) + "a") + "a")))
        }
    }
    var v6 uint32 = 2
    var v7 int8 = 100
    if false {
        var v8 int8 = v7
        fmt.Println(v7, v6)
        _ = v8
    } else {
        if (false || (390 < 65535)) {
            v9 := GenS1{F0: false, F1: 4294967295}
            var v10 uint = 639
            var v11 float64 = float64(0.1)
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 string = "a"
        var v13 uint8 = 1
        _ = v12
        _ = v13
    }
    switch ("hello" + "") {
    case "hello":
        if false {
            v14 := GenS1{F0: ((^(v2)) > (^(v2))), F1: 1}
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "a":
                var v15 int8 = v7
                _ = v15
            }
            _ = len(v3)
            _ = v14
        }
    default:
        v6 += v6
    }
    v16 := v3["hello"]
    _ = len(v3)
    v16 = 9223372036854775807
    if (-618 <= 42) {
        switch v16 {
        case -745:
            var v17 uint32 = 608
            _ = v17
        case -956:
            var v18 int16 = -883
            _ = v18
        case -1:
            if (!(((v7 <= v7) && (-644 >= -1)) && false)) {
                var v19 float64 = 1e10
                var v20 int = v16
                _ = v19
                _ = v20
            } else {
                var v21 int32 = 0
                v7 += (^(v7))
                v7 = 42
                _ = v21
            }
        }
        var v22 int64 = 1
        _ = v22
    } else {
        _ = len(v1)
        var v23 int64 = 1
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((0.0 * 0.0) != 3.14159) {
                var v25 int64 = v23
                _ = v25
            } else {
                var v26 int16 = -960
                _ = v26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS1{F0: (false && (true || (v6 >= (v6 + (v6 ^ v6))))), F1: 4294967295}
                if true {
                    var v29 int = (^(fi24))
                    switch "" {
                    case "foo bar":
                        fmt.Println(v16)
                    case "":
                        v29 = (^(v16))
                    case "a":
                        switch v16 {
                        case 123:
                            fmt.Println(v23, v23, v16)
                        default:
                            switch 9223372036854775807 {
                            case -575:
                                var v30 int8 = (v7 + (v7 / ((v7) | 1)))
                                _ = v30
                            case -962:
                                if (v2 >= v2) {
                                    if false {
                                        var v31 int16 = 1
                                        _ = v31
                                    } else {
                                        var v32 uint64 = 2
                                        _ = v32
                                    }
                                    v23 = (v23 & 9223372036854775807)
                                    var v33 int = (v29 - v16)
                                    v34 := GenS2{F0: 1, F1: "foo bar"}
                                    _ = v33
                                    _ = v34
                                } else {
                                    switch ("a" + ("a" + "\n\t")) {
                                    case "hello":
                                        var v35 float32 = float32(0.0)
                                        _ = v35
                                    case "":
                                        if (v23 >= v23) {
                                            if false {
                                                var v36 float64 = (1.5 / 1.5)
                                                v37 := GenS3{F0: (v6 % (((^(v6))) | 1)), F1: (float32((float64((((1e10 / ((-2.25 / -2.25) / 1.5)) + -2.25) - (0.0 + (0.1 - 0.1)))) + 0.0)) >= -2.25), F2: 0}
                                                fmt.Println(fi24, v2)
                                                _ = v36
                                                _ = v37
                                            }
                                            switch v29 {
                                            case 42:
                                                _ = len(v3)
                                            case 726:
                                                v4 = append(v4, 2)
                                            case -653:
                                                switch "" {
                                                case "\n\t":
                                                    fmt.Println(v7, v16)
                                                }
                                            default:
                                                v38 := GenS1{F0: true, F1: 18446744073709551615}
                                                _ = v38
                                            }
                                        }
                                    default:
                                        v29 -= (-(v16))
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v16 = (fi24 << ((226) & 63))
                                    }
                                    var v39 float64 = (float64(3.14159) - float64(0.0))
                                    v40 := GenS2{F0: -955, F1: "a"}
                                    _ = v39
                                    _ = v40
                                }
                            }
                        }
                    }
                    _ = v29
                }
                _ = fi27
                _ = v28
            }
            _ = fi24
        }
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
