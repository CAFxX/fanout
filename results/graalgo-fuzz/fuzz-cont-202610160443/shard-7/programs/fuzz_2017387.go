package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (float64(0.0) * (3.14159 - 0.1))
    v3 := []string{("foo bar" + ""), "", "foo bar"}
    var v4 int = 9223372036854775807
    var v5 float32 = 0.0
    fmt.Println(v5)
    if false {
        v2 += -2.25
        v4 += v4
        fmt.Println(v2)
    }
    var v6 uint64 = 1
    if (-1 < -324) {
        _ = len(v1)
        fmt.Println(v2, v5, v5)
        switch v4 {
        case -2147483648:
            for fi7 := 0; fi7 < 5; fi7++ {
                _ = len(v1)
                _ = fi7
            }
        case 32767:
            for fi8 := 0; fi8 < 3; fi8++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(fi8, v2)
                default:
                    if true {
                        var v9 float64 = v2
                        if (255 != 42) {
                            if (false || (true || true)) {
                                v5 += ((0.1 + (((((float32(float64(float32((((v2 - v9) / v9) * 0.0)))) - (float32(-2.25) - 0.1)) * v5) / v5) * (v5 / v5)) - v5)) / (v5 + v5))
                                if true {
                                    var v10 int64 = -885
                                    var v11 bool = true
                                    var v12 bool = true
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                } else {
                                    var v13 uint32 = 0
                                    v14 := GenS1{F0: 32767}
                                    _ = v13
                                    _ = v14
                                }
                            } else {
                                v15 := GenS1{F0: -1}
                                _ = v15
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v5 += v5
                            case "foo bar":
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 float32 = (v5 - v5)
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            var v18 uint32 = 65535
                            _ = v18
                        }
                        _ = v9
                    } else {
                        v19 := v1[(("hello" + "hello") + "hello")]
                        _ = v19
                    }
                }
                _ = fi8
            }
        }
    }
    switch int(v6) {
    case 1:
        v4 -= 675
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v4 = (v4 % ((v4) | 1))
        _ = fi20
    }
    fmt.Println(v2)
    fmt.Println(v2, v6, v2)
    var v21 float64 = -2.25
    v22 := GenS2{F0: 42, F1: 522}
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint = 1
        v5 = (1.5 + float32(v2))
        v25 := v1[(("foo bar" + "hello") + ((("a" + ("a" + "\n\t")) + "a") + "\n\t"))]
        _ = fi23
        _ = v24
        _ = v25
    }
    v21 -= -2.25
    if (!false) {
        for fi26 := 0; fi26 < 17; fi26++ {
            switch v4 {
            case 9223372036854775807:
                fmt.Println(v2)
            case 4294967295:
                var v27 float64 = 0.1
                _ = v27
            case 42:
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 bool = true
                    _ = fi28
                    _ = v29
                }
            }
            if ((!(!false)) || false) {
                fmt.Println(v6, v21)
                if false {
                    v3 = append(v3, "\n\t")
                    var v30 int16 = -45
                    _ = v30
                }
            } else {
                v1["hello"] = 42
                switch v4 {
                case 2147483647:
                    if (!(-9223372036854775808 < fi26)) {
                        v2 += (-2.25 / (0.1 + 0.1))
                        var v31 uint64 = v6
                        for fi32 := 0; fi32 < 6; fi32++ {
                            fmt.Println(v21)
                            v2 = v2
                            _ = fi32
                        }
                        _ = v31
                    } else {
                        v33 := GenS1{F0: 32767}
                        _ = v33
                    }
                case -461:
                    var v34 uint8 = 1
                    _ = v34
                case -1:
                    v3 = append(v3, ("\n\t" + ""))
                }
            }
            _ = fi26
        }
    }
    var v35 int32 = -2147483648
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
