package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    var v1 int = -2147483648
    var v2 uint8 = 0
    var v3 uint32 = 0
    v4 := []int{668}
    v5 := make(map[string]bool)
    fmt.Println(v3)
    if false {
        delete(v5, "foo bar")
    }
    fmt.Println(v2)
    v3 = v3
    fmt.Println(v1)
    v1 = (^(v1))
    v1 = (v1 | v1)
    v5["a"] = true
    fmt.Println(v3, v1, v3)
    v6 := GenS1{F0: 42, F1: ("" + "")}
    var v7 int8 = 0
    v2 = (v2 | 229)
    v8 := GenS1{F0: 65535, F1: ""}
    fmt.Println(v7, v3, v2)
    var v9 int64 = 42
    var v10 int32 = 1
    var v11 uint = 1
    if true {
        v4 = append(v4, int(v11))
        switch "foo bar" {
        case "a":
            switch "" {
            case "\n\t":
                v12 := v5[("hello" + "\n\t")]
                _ = v12
            case "hello":
                switch 478 {
                case 128:
                    v7 = v7
                case 32768:
                    v13 := GenS1{F0: 42, F1: ("\n\t" + ("\n\t" + "hello"))}
                    _ = v13
                case -9223372036854775808:
                    if false {
                        switch "hello" {
                        case "\n\t":
                            v11 = v11
                        default:
                            _ = len(v4)
                        }
                        v2 -= (v2 + v2)
                        v14 := GenS1{F0: (v11 >> (((v11 / (((v11 + (v11 + v11))) | 1))) & 63)), F1: ("hello" + ("" + "a"))}
                        _ = v14
                    } else {
                        switch "a" {
                        case "a":
                            v9 += (v9 % ((v9) | 1))
                        case "foo bar":
                            if (-2147483648 >= v10) {
                                if (false || (v3 != v3)) {
                                    v15 := GenS1{F0: v11, F1: "foo bar"}
                                    var v16 int = (-(v1))
                                    _ = v15
                                    _ = v16
                                } else {
                                    v17 := GenS1{F0: 1, F1: "hello"}
                                    switch "foo bar" {
                                    case "\n\t":
                                        v11 = 963
                                    case "a":
                                        v1 = (v1 % (((v1 & 1)) | 1))
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            for fi18 := 0; fi18 < 1; fi18++ {
                                                v5[("" + (("\n\t" + (("\n\t" + "\n\t") + ("a" + "hello"))) + "hello"))] = (!false)
                                                _ = fi18
                                            }
                                        case "\n\t":
                                            delete(v5, "a")
                                        case "foo bar":
                                            v3 = 2147483647
                                        default:
                                            if true {
                                                if true {
                                                    if (1 >= 18446744073709551615) {
                                                        var v19 int64 = (v9 % ((-1) | 1))
                                                        v20 := GenS1{F0: (v11 / ((135) | 1)), F1: ""}
                                                        if (true && true) {
                                                            v21 := GenS1{F0: 760, F1: "a"}
                                                            for fi22 := 0; fi22 < 6; fi22++ {
                                                                v2 += 178
                                                                _ = len(v4)
                                                                for fi23 := 0; fi23 < 10; fi23++ {
                                                                    switch ("hello" + (("" + ("foo bar" + "")) + "a")) {
                                                                    case "\n\t":
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi23
                                                                        }
                                                                    case "":
                                                                        v7 = (v7 << (((v11 & 4294967295)) & 63))
                                                                    }
                                                                    _ = fi23
                                                                }
                                                                _ = fi22
                                                            }
                                                            var v24 int = 32767
                                                            _ = v21
                                                            _ = v24
                                                        } else {
                                                            var v25 int64 = v19
                                                            _ = v25
                                                        }
                                                        _ = v19
                                                        _ = v20
                                                    }
                                                }
                                                fmt.Println(v10, v2)
                                                v9 = -9223372036854775808
                                            } else {
                                                _ = len(v5)
                                            }
                                        }
                                    }
                                    _ = v17
                                }
                                v10 = -1
                                switch "a" {
                                case "\n\t":
                                    var v26 uint8 = (v2 | v2)
                                    _ = v26
                                case "":
                                    var v27 uint = 255
                                    _ = v27
                                default:
                                    var v28 float32 = (0.0 - 1.5)
                                    _ = v28
                                }
                            }
                        case "\n\t":
                            switch (("" + (("" + "") + "foo bar")) + ("foo bar" + ("hello" + ""))) {
                            case "foo bar":
                                _ = len(v4)
                            default:
                                var v29 uint16 = 0
                                _ = v29
                            }
                        }
                        fmt.Println(v10)
                        var v30 uint32 = (v3 + 2147483648)
                        _ = v30
                    }
                default:
                    var v31 int32 = -2147483648
                    _ = v31
                }
            }
        }
    } else {
        var v32 int64 = 1
        v9 += (v9 % ((v9) | 1))
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
