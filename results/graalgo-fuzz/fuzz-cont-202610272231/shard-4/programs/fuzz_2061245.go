package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 bool = ((3.14159 + 0.0) != ((3.14159 * ((((-2.25 * 1e10) + 0.0) + (1e10 - ((0.1 + (0.1 - (-2.25 - float32(3.14159)))) / 1.0))) - (3.14159 - 1.5))) + 1.5))
    var v2 int = 32767
    var v3 uint64 = 18446744073709551615
    v4 := []string{("a" + (("\n\t" + ("hello" + "\n\t")) + "a")), ("a" + "a"), "", "foo bar"}
    var v5 int = -356
    if ((!v1) && v1) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    } else {
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 uint64 = (v3 % ((v3) | 1))
                var v9 uint8 = 2
                v10 := GenS1{F0: 18446744073709551615, F1: 485}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS2{F0: -2147483648}
                v2 = fi11
                v1 = false
                _ = fi11
                _ = v12
            }
            _ = fi6
        }
    }
    var v13 int8 = 100
    if v1 {
        var v14 uint16 = 588
        if v1 {
            var v15 uint64 = (v3 - (v3 * v3))
            v16 := GenS2{F0: int32(v14)}
            _ = v15
            _ = v16
        }
        _ = v14
    } else {
        switch (v2 - (v5 >> ((2) & 63))) {
        case 42:
            var v17 int32 = -367
            _ = v17
        default:
            var v18 float32 = 3.14159
            _ = v18
        }
        v19 := GenS1{F0: 4294967295, F1: 4294967295}
        _ = v19
    }
    var v20 int = (v5 / ((-1) | 1))
    v21 := GenS2{F0: 1}
    if ((!true) || true) {
        v3 -= 93
    }
    v20 = (v20 / (((v2 / (((^(v2))) | 1))) | 1))
    v5 = (^(v5))
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int32 = 2147483647
        var v24 uint = 108
        _ = fi22
        _ = v23
        _ = v24
    }
    if (1 < 2147483648) {
        var v25 float64 = 3.14159
        _ = v25
    } else {
        v4 = append(v4, "hello")
    }
    if (!(1 <= 42)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch v20 {
        case 9223372036854775807:
            switch "" {
            case "foo bar":
                v13 = (^(v13))
            case "\n\t":
                if (1.5 <= float32(((0.0 / ((((1.5 - 0.1) / 1.5) - 0.0) - 1e10)) / 1e10))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v26 := GenS2{F0: 42}
                    _ = v26
                }
            default:
                for fi27 := 0; fi27 < 3; fi27++ {
                    fmt.Println(v1, v20, v1)
                    fmt.Println(fi27, v20)
                    for fi28 := 0; fi28 < 5; fi28++ {
                        if v1 {
                            var v29 uint64 = v3
                            v30 := GenS3{F0: 88, F1: 0}
                            v2 = v2
                            _ = v29
                            _ = v30
                        } else {
                            var v31 int = 42
                            fmt.Println(fi28, v5, v31)
                            _ = v31
                        }
                        if v1 {
                            fmt.Println(v2, fi28)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi28
                    }
                    _ = fi27
                }
            }
        case -127:
            v4 = append(v4, "a")
        case 42:
            if (v13 >= (v13 >> ((549) & 63))) {
                if false {
                    var v32 int32 = -1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v3)
                    _ = v32
                } else {
                    if false {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            switch v2 {
                            case 9223372036854775807:
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                            case -1:
                                var v34 int32 = 42
                                _ = v34
                            }
                            var v35 string = "foo bar"
                            _ = fi33
                            _ = v35
                        }
                        v36 := GenS3{F0: 111, F1: 42}
                        _ = v36
                    }
                    fmt.Println(v1, v1)
                }
                v37 := GenS1{F0: 255, F1: 4294967295}
                v13 += v13
                _ = v37
            } else {
                if v1 {
                    var v38 uint8 = 176
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 6; fi39++ {
                        var v40 uint = 255
                        _ = fi39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 7; fi41++ {
                        if v1 {
                            v13 = v13
                            switch ("" + "hello") {
                            case "a":
                                switch "foo bar" {
                                case "foo bar":
                                    var v42 float64 = 0.0
                                    _ = v42
                                case "a":
                                    v20 = (fi41 & fi41)
                                case "":
                                    v5 = v2
                                default:
                                    for fi43 := 0; fi43 < 10; fi43++ {
                                        if (255 >= 1) {
                                            fmt.Println(fi41, v3)
                                        }
                                        fmt.Println(fi43)
                                        _ = fi43
                                    }
                                }
                            case "\n\t":
                                v1 = ((v1 && (!(-1 != -2147483648))) && v1)
                            case "foo bar":
                                var v44 int64 = -1
                                _ = v44
                            }
                        } else {
                            v5 -= -2147483648
                            v1 = (fi41 < 841)
                            v45 := GenS4{F0: 773}
                            _ = v45
                        }
                        _ = fi41
                    }
                    v5 += 4294967295
                }
                for fi46 := 0; fi46 < 19; fi46++ {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    _ = fi46
                }
            }
        }
        if true {
            v47 := GenS5{F0: 136}
            _ = v47
        }
    }
    if v1 {
        var v48 int8 = v13
        switch ((("" + ("\n\t" + ("" + "a"))) + ("hello" + (("foo bar" + "hello") + "a"))) + "hello") {
        case "\n\t":
            var v49 int = (v5 / ((v2) | 1))
            _ = v49
        }
        _ = v48
    }
    v5 -= (v5 + v20)
    var v50 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
