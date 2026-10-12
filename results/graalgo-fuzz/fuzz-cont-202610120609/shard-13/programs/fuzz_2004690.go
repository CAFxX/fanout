package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    var v3 int32 = 646
    var v4 uint64 = 1
    v5 := []string{"hello", "hello"}
    var v6 uint64 = v4
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 255
        v9 := GenS1{F0: -9223372036854775808, F1: 255}
        switch (v1 + 2147483647) {
        case 32768:
            var v10 string = "foo bar"
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = v4
    var v11 uint8 = 176
    var v12 int16 = -525
    fmt.Println(v3, v6)
    v13 := v2["a"]
    _ = len(v5)
    var v14 uint = 2
    if (1 < 32767) {
        switch "" {
        case "\n\t":
            var v15 int64 = -9223372036854775808
            _ = v15
        }
        var v16 float32 = -2.25
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v6, v14, v6)
            var v18 string = ("foo bar" + (v13 + v13))
            var v19 int8 = 1
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        if true {
            if (false || (v3 == int32(v14))) {
                fmt.Println(v6)
            } else {
                var v20 int32 = v3
                if (541 == 10) {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        var v22 string = (v13 + "hello")
                        v23 := GenS1{F0: -9223372036854775808, F1: 46}
                        if false {
                            var v24 int64 = 0
                            v12 -= v12
                            v14 -= (v14 | 255)
                            _ = v24
                        } else {
                            switch "a" {
                            case "hello":
                                for fi25 := 0; fi25 < 13; fi25++ {
                                    v22 = v22
                                    if false {
                                        var v26 float64 = 1.5
                                        _ = v26
                                    }
                                    v4 = v6
                                    _ = fi25
                                }
                            }
                            fmt.Println(v3, v22, v6)
                        }
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    delete(v2, "a")
                } else {
                    v27 := GenS1{F0: 9223372036854775807, F1: v14}
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 float32 = 1e10
                        fmt.Println(v4)
                        for fi30 := 0; fi30 < 15; fi30++ {
                            if ((v29 + 0.1) != (v29 * v29)) {
                                v13 = v13
                                _ = len(v5)
                            } else {
                                var v31 int8 = 127
                                _ = v31
                            }
                            if true {
                                v14 = (v14 ^ 255)
                                v32 := GenS1{F0: -474, F1: (v14 * (v14 - 322))}
                                switch v13 {
                                case "foo bar":
                                    if false {
                                        if len(v5) > 0 {
                                            v5[0] = v13
                                        }
                                    }
                                }
                                _ = v32
                            }
                            _ = fi30
                        }
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
                for fi33 := 0; fi33 < 9; fi33++ {
                    switch "foo bar" {
                    case "hello":
                        var v34 uint16 = uint16(v14)
                        _ = v34
                    }
                    var v35 float32 = 0.0
                    _ = fi33
                    _ = v35
                }
                _ = v20
            }
            v4 = (v6 ^ v6)
        }
        var v36 float64 = ((((1e10 + 0.0) - (((0.1 + 1.5) / (1.5 / 1.5)) * 1.5)) - (3.14159 + ((1e10 * 1e10) / float64(float32((1e10 - 0.0)))))) * (((0.0 / ((1.5 + (-2.25 * -2.25)) * (-2.25 / 0.1))) - (1.5 - ((1e10 + 0.1) - ((0.0 + 1.5) + (1.5 / -2.25))))) + 3.14159))
        _ = v36
    }
    var v37 uint64 = (v6 % ((v6) | 1))
    fmt.Println(v6, v13)
    var v38 uint64 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
