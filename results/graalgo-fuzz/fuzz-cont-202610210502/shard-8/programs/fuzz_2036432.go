package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int16 }

func main() {
    var v1 int32 = -247
    var v2 float64 = float64((1.5 * -2.25))
    v3 := make(map[string]int)
    v2 = 3.14159
    fmt.Println(v1)
    var v4 uint32 = 516
    v5 := GenS1{F0: -403, F1: 931, F2: 42}
    fmt.Println(v1, v4)
    var v6 uint16 = 32767
    var v7 int64 = 713
    if (-32768 != 708) {
        v8 := []int{}
        switch "a" {
        case "\n\t":
            v3["foo bar"] = 4294967295
        case "hello":
            var v9 string = ("a" + ("a" + (((("" + ("\n\t" + ("foo bar" + "a"))) + "foo bar") + "\n\t") + "hello")))
            _ = v9
        default:
            var v10 int8 = -1
            _ = v10
        }
        if ((v7 % ((v7) | 1)) >= 1) {
            _ = len(v8)
            v11 := GenS1{F0: 0, F1: -32768, F2: 1000}
            _ = v11
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v12 uint8 = 255
            _ = v12
        }
        _ = v8
    } else {
        var v13 uint16 = v6
        _ = v13
    }
    var v14 uint32 = v4
    if (!true) {
        if true {
            v3[("a" + ("" + "hello"))] = -1
            v15 := GenS1{F0: 0, F1: -1, F2: 42}
            switch -1 {
            case 9223372036854775807:
                switch -423 {
                case 42:
                    for fi16 := 0; fi16 < 7; fi16++ {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 int16 = 42
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(v7, v6)
                        v19 := []int{-901}
                        _ = fi16
                        _ = v19
                    }
                default:
                    v14 -= 2147483647
                }
            }
            _ = v15
        }
    } else {
        if false {
            v20 := v3[("" + "hello")]
            delete(v3, "\n\t")
            v14 = (v4 % (((v4 / ((v4) | 1))) | 1))
            _ = v20
        } else {
            var v21 uint64 = 42
            var v22 int64 = (v7 / ((0) | 1))
            _ = v21
            _ = v22
        }
        if (!true) {
            if (!false) {
                var v23 int32 = (-(v1))
                var v24 int8 = 127
                _ = v23
                _ = v24
            } else {
                v1 = v1
            }
            v25 := []int{}
            v4 = (^(v4))
            _ = v25
        } else {
            if (((42 != v7) || true) && true) {
                if (true || false) {
                    switch ("a" + "a") {
                    case "\n\t":
                        _ = len(v3)
                    default:
                        v4 -= 2147483648
                    }
                    v7 = v7
                    var v26 uint64 = 1
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        if true {
                            for fi28 := 0; fi28 < 5; fi28++ {
                                if false {
                                    v6 = (^(v6))
                                }
                                _ = fi28
                            }
                        } else {
                            var v29 float32 = 0.1
                            _ = v29
                        }
                        v30 := GenS1{F0: -32768, F1: 32767, F2: 450}
                        v31 := []int{(fi27 ^ 32768), int(v7), (fi27 & fi27)}
                        _ = fi27
                        _ = v30
                        _ = v31
                    }
                }
                v14 -= 776
                for fi32 := 0; fi32 < 5; fi32++ {
                    if false {
                        for fi33 := 0; fi33 < 3; fi33++ {
                            v2 = (float64(3.14159) - 1.5)
                            var v34 uint = 65535
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        v4 = (^(v14))
                        v35 := GenS1{F0: 1, F1: 32767, F2: -303}
                        var v36 int64 = (^(v7))
                        _ = v35
                        _ = v36
                    }
                    _ = fi32
                }
            }
            switch "\n\t" {
            case "hello":
                delete(v3, "hello")
            default:
                v6 = v6
            }
            v6 = v6
        }
        v2 += (0.1 / 1.5)
    }
    var v37 string = ("hello" + ("hello" + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
