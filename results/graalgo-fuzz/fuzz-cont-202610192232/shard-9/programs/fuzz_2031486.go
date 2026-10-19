package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 float32 = float32(((1.5 + -2.25) + float64(float32(0.0))))
    v3 := []string{"hello", "hello", ""}
    switch 344 {
    case 127:
        if (-1 == -9223372036854775808) {
            if false {
                v2 += 3.14159
            }
            for fi4 := 0; fi4 < 4; fi4++ {
                if false {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            v1 += v1
                        } else {
                            switch (fi4 >> ((255) & 63)) {
                            case 4294967295:
                                for fi5 := 0; fi5 < 18; fi5++ {
                                    v2 = float32(1.5)
                                    var v6 int32 = 619
                                    _ = fi5
                                    _ = v6
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                            }
                            fmt.Println(v2, v2)
                        }
                    case "\n\t":
                        if (false && true) {
                            switch fi4 {
                            case -649:
                                for fi7 := 0; fi7 < 15; fi7++ {
                                    var v8 bool = ((float64(((1e10 - 0.1) - 1e10)) + ((-2.25 + ((1e10 - 0.0) - 0.0)) + 0.1)) >= 1e10)
                                    fmt.Println(fi7, fi7)
                                    _ = fi7
                                    _ = v8
                                }
                            default:
                                var v9 bool = true
                                _ = v9
                            }
                            for fi10 := 0; fi10 < 20; fi10++ {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    v1 -= v1
                                case "\n\t":
                                    v1 -= 0
                                default:
                                    var v11 int = (fi10 + fi4)
                                    _ = v11
                                }
                                v12 := GenS1{F0: (683 > 276)}
                                _ = fi10
                                _ = v12
                            }
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            var v13 int8 = 60
                            var v14 bool = true
                            _ = v13
                            _ = v14
                        }
                    case "":
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v2 = 1.5
                            _ = fi15
                        }
                    }
                    var v16 int64 = -186
                    switch "" {
                    case "\n\t":
                        if (false && false) {
                            var v17 float64 = (1.5 + 0.1)
                            v1 -= 2
                            var v18 bool = false
                            _ = v17
                            _ = v18
                        } else {
                            fmt.Println(v2, v16)
                            var v19 int = (fi4 & int(v1))
                            v16 += (v16 | v16)
                            _ = v19
                        }
                    case "foo bar":
                        v3 = append(v3, "hello")
                    case "":
                        v20 := GenS2{F0: 65535, F1: 0}
                        _ = v20
                    }
                    _ = v16
                }
                v2 = v2
                _ = fi4
            }
            v21 := GenS3{F0: 707, F1: 596}
            _ = v21
        }
    case 234:
        v22 := GenS1{F0: false}
        _ = v22
    case 32768:
        if (2147483647 != 42) {
            v2 = v2
            if true {
                var v23 uint16 = v1
                v24 := GenS4{F0: 255}
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: (241 >= 0)}
            _ = v25
        }
    default:
        v1 += v1
    }
    v1 += (^(v1))
    fmt.Println(v2, v2)
    _ = len(v3)
    var v26 int16 = -32768
    _ = len(v3)
    var v27 uint32 = 601
    fmt.Println(v1, v26, v1)
    v26 += 1
    var v28 float32 = -2.25
    v2 -= (v28 + (v2 + 1e10))
    var v29 int32 = 0
    if (18446744073709551615 == uint(v1)) {
        fmt.Println(v26, v26, v27)
    } else {
        if len(v3) > 0 {
            v3[0] = ""
        }
        v29 = (v29 - v29)
    }
    v1 -= 0
    v30 := GenS3{F0: 861, F1: v1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int = (fi31 / ((fi31) | 1))
        if ("foo bar" != "") {
            v3 = append(v3, ((("foo bar" + "hello") + "a") + "foo bar"))
        }
        _ = fi31
        _ = v32
    }
    switch -678 {
    case 2147483647:
        v3 = append(v3, ("hello" + ("a" + "a")))
    case -279:
        v2 += (float32(1e10) - -2.25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
