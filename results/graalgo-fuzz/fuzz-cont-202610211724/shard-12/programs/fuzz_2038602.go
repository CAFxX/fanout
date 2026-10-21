package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }

func main() {
    var v1 float64 = (1.5 * (1.5 * 1e10))
    var v2 int64 = 361
    var v3 float32 = 1.5
    var v4 uint8 = 1
    var v5 float32 = v3
    v6 := []int{1, 625, 32767}
    var v7 float64 = float64(v3)
    if (!true) {
        v5 = v3
    } else {
        if false {
            if true {
                var v8 float32 = v3
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(v8, v7, fi9)
                    for fi10 := 0; fi10 < 1; fi10++ {
                        switch ("hello" + ("" + "a")) {
                        case "foo bar":
                            var v11 int8 = -1
                            _ = v11
                        default:
                            var v12 int16 = 42
                            _ = v12
                        }
                        var v13 int64 = (v2 - -421)
                        _ = fi10
                        _ = v13
                    }
                    _ = fi9
                }
                switch (("foo bar" + ((("a" + "") + "a") + "\n\t")) + ("a" + (("foo bar" + (("a" + "hello") + "\n\t")) + "a"))) {
                case "\n\t":
                    v3 += v5
                case "hello":
                    v1 += (v1 * (v1 + ((v1 * v1) + (v1 + (v1 * 0.0)))))
                case "":
                    v8 = -2.25
                default:
                    fmt.Println(v7)
                }
                _ = v8
            }
            if true {
                for fi14 := 0; fi14 < 9; fi14++ {
                    for fi15 := 0; fi15 < 16; fi15++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if false {
                            var v16 int8 = -1
                            v17 := GenS1{F0: (-(fi15)), F1: 787, F2: uint32(v4)}
                            var v18 float32 = ((v3 / 1.0) + (-2.25 - ((float32(v7) - (v3 * v5)) / (v5 * 1.5))))
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            if len(v6) > 0 {
                                v6[0] = fi14
                            }
                        }
                        _ = fi15
                    }
                    for fi19 := 0; fi19 < 17; fi19++ {
                        fmt.Println(v7)
                        v5 -= v5
                        if false {
                            if len(v6) > 0 {
                                v6[0] = fi19
                            }
                            v7 = 1.5
                            v4 += (^(v4))
                        } else {
                            v7 -= 1.5
                            if (v3 == (v3 + v5)) {
                                var v20 string = ("a" + ((("foo bar" + "foo bar") + "") + "foo bar"))
                                var v21 string = "foo bar"
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS1{F0: (fi19 - (fi19 ^ -489)), F1: 42, F2: 996}
                                v23 := GenS1{F0: fi14, F1: -451, F2: 42}
                                _ = v22
                                _ = v23
                            }
                            v7 = v7
                        }
                        _ = fi19
                    }
                    _ = fi14
                }
                if (1 != 2147483648) {
                    v24 := GenS1{F0: -2147483648, F1: -1, F2: 4294967295}
                    _ = v24
                } else {
                    if (!(255 <= v4)) {
                        v3 += (v5 - ((((v3 * (float32((v7 + v7)) + float32((v7 * (((3.14159 + (((0.0 - ((v7 / 1.0) * v7)) / (v1 * 0.0)) / (1.5 + ((v7 - v1) * v7)))) + v7) + v7))))) / v3) + v3) / v5))
                    }
                }
                v25 := GenS1{F0: -703, F1: 44, F2: 426}
                _ = v25
            } else {
                _ = len(v6)
                var v26 int8 = 42
                _ = v26
            }
        } else {
            var v27 uint = 255
            _ = v27
        }
        if true {
            switch "foo bar" {
            case "a":
                v28 := GenS1{F0: -1, F1: 19, F2: 1}
                _ = v28
            case "":
                v2 -= -378
            case "hello":
                v29 := GenS1{F0: 2147483647, F1: -1, F2: 42}
                _ = v29
            }
        }
    }
    var v30 float32 = 1e10
    var v31 float32 = 1.5
    if true {
        var v32 uint8 = 42
        if true {
            var v33 int32 = -634
            _ = v33
        }
        if (1 > 4294967295) {
            var v34 int32 = 0
            var v35 uint64 = 2
            _ = v34
            _ = v35
        } else {
            v31 = (v3 + v30)
        }
        _ = v32
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v30 -= (v5 - (v3 * v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
