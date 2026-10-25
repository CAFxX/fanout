package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 float64 = 0.1
    var v2 int64 = 1
    v3 := []string{((("\n\t" + (("hello" + "") + "")) + ("foo bar" + "foo bar")) + "foo bar"), "a"}
    var v4 int8 = -48
    fmt.Println(v1)
    v1 += float64(1e10)
    v3 = append(v3, "foo bar")
    v3 = append(v3, "hello")
    switch ("a" + ("a" + "hello")) {
    case "":
        for fi5 := 0; fi5 < 12; fi5++ {
            v4 -= v4
            if true {
                switch (-(fi5)) {
                case -621:
                    v1 += ((v1 * (3.14159 - v1)) - 1.5)
                case 468:
                    if false {
                        v4 -= (v4 / ((-128) | 1))
                        v6 := GenS1{F0: 2, F1: 127}
                        _ = v6
                    } else {
                        if true {
                            switch "" {
                            case "hello":
                                if false {
                                    if ((v2 <= (v2 & 1)) || true) {
                                        var v7 uint8 = 0
                                        v1 += (v1 - (v1 * v1))
                                        var v8 uint16 = 65535
                                        _ = v7
                                        _ = v8
                                    } else {
                                        v2 += 0
                                    }
                                }
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "a":
                                fmt.Println(fi5, v1, v2)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = v1
                        } else {
                            for fi9 := 0; fi9 < 6; fi9++ {
                                fmt.Println(v2)
                                _ = fi9
                            }
                            if false {
                                var v10 bool = true
                                v2 = v2
                                v4 += v4
                                _ = v10
                            } else {
                                var v11 float64 = v1
                                _ = v11
                            }
                        }
                        var v12 bool = true
                        _ = v12
                    }
                case 885:
                    var v13 int16 = 1
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 16; fi14++ {
                    switch ("" + "") {
                    case "hello":
                        v4 -= 1
                    case "\n\t":
                        switch ("" + "hello") {
                        case "":
                            _ = len(v3)
                        default:
                            var v15 uint8 = 207
                            _ = v15
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if ((0.0 * ((1.5 * 0.1) / 1.0)) == 0.0) {
                        fmt.Println(fi5, fi5, v1)
                        v16 := GenS1{F0: 4294967295, F1: 128}
                        _ = v16
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v17 uint8 = 128
                        _ = v17
                    }
                    _ = fi14
                }
                if (-850 <= 0) {
                    v2 -= (v2 & (v2 >> ((908) & 63)))
                    switch fi5 {
                    case 0:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -1:
                        if false {
                            if len(v3) > 0 {
                                v3[0] = (("a" + "foo bar") + "\n\t")
                            }
                        }
                    case 4294967295:
                        v3 = append(v3, "\n\t")
                    }
                    var v18 uint8 = 245
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                v2 -= v2
                var v20 uint = 750
                v20 += v20
                _ = fi19
                _ = v20
            }
            _ = fi5
        }
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v2 += 0
    }
    var v21 bool = true
    v22 := GenS1{F0: 65535, F1: 127}
    switch -627 {
    case 1:
        v23 := GenS2{F0: 1000, F1: 4294967295}
        _ = v23
    case -728:
        var v24 uint16 = 255
        _ = v24
    }
    switch "foo bar" {
    case "hello":
        var v25 uint16 = 355
        _ = v25
    case "foo bar":
        var v26 string = "a"
        _ = v26
    }
    if (v21 || (0 < 582)) {
        v3 = append(v3, "foo bar")
        var v27 uint16 = 65535
        _ = v27
    } else {
        var v28 uint8 = 1
        for fi29 := 0; fi29 < 17; fi29++ {
            fmt.Println(v4)
            var v30 float32 = ((-2.25 + 0.0) - (-2.25 / (1.5 / 0.1)))
            var v31 bool = v21
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    v2 += (v2 | v2)
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
