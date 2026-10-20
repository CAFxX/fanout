package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int = 127
    var v2 int = (v1 + v1)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 86, F1: 128, F2: 0}
        v2 += v1
        v1 -= fi4
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 65535, F1: 127, F2: 9223372036854775807}
    for fi7 := 0; fi7 < 10; fi7++ {
        switch (("foo bar" + "hello") + "\n\t") {
        case "":
            v1 -= (fi7 & fi7)
        case "hello":
            var v8 string = (("foo bar" + ("a" + "foo bar")) + ("\n\t" + "a"))
            _ = v8
        case "foo bar":
            var v9 uint32 = 568
            _ = v9
        default:
            fmt.Println(v1, v2)
        }
        var v10 int32 = -1
        for fi11 := 0; fi11 < 6; fi11++ {
            switch ("a" + "a") {
            case "\n\t":
                var v12 float32 = ((((float32(3.14159) - (float32(((1.5 / 1.0) - ((1e10 + ((1.5 / float64(1.5)) / 1.0)) - (((1.5 + 0.1) - 0.1) / 3.14159)))) * float32(1.5))) / (float32(3.14159) - float32(1e10))) + 0.0) / ((0.0 * 3.14159) - 3.14159))
                _ = v12
            }
            switch 2147483647 {
            case -2147483648:
                v10 -= v10
            case 42:
                v10 = -1
            case 293:
                v10 += v10
                var v13 float64 = -2.25
                _ = v13
            default:
                if false {
                    switch (("\n\t" + ("hello" + "")) + "a") {
                    case "\n\t":
                        if (3.14159 < 1.5) {
                            v14 := GenS1{F0: 0, F1: 128, F2: -1}
                            var v15 uint64 = 42
                            _ = v14
                            _ = v15
                        }
                    case "foo bar":
                        if (3.14159 <= ((1.5 + 3.14159) * 1.5)) {
                            switch "\n\t" {
                            case "\n\t":
                                v16 := []int{fi11, -324}
                                _ = v16
                            case "foo bar":
                                if false {
                                    _ = len(v3)
                                    var v17 uint32 = 580
                                    _ = v17
                                }
                            default:
                                v10 = v10
                                var v18 float64 = -2.25
                                _ = v18
                            }
                            if (!true) {
                                _ = len(v3)
                            } else {
                                v10 = (v10 % ((v10) | 1))
                                v2 = -2147483648
                            }
                            v19 := GenS2{F0: 42, F1: 77, F2: 28}
                            _ = v19
                        }
                    case "hello":
                        v20 := []int{-483, (v1 & v1)}
                        _ = v20
                    }
                    delete(v3, "a")
                    if false {
                        v21 := GenS2{F0: 2, F1: 0, F2: 255}
                        if (-87 < 32767) {
                            if ((668 > 1) || true) {
                                fmt.Println(v1, v2)
                                v3[("\n\t" + "\n\t")] = ("foo bar" + "a")
                                v22 := []int{(-(fi11)), fi7, 1}
                                _ = v22
                            }
                        } else {
                            var v23 string = ("a" + "")
                            v10 = 42
                            _ = v23
                        }
                        _ = v21
                    }
                } else {
                    var v24 uint32 = 2
                    _ = v24
                }
            }
            v3["hello"] = ("" + "\n\t")
            _ = fi11
        }
        _ = fi7
        _ = v10
    }
    v25 := GenS3{F0: 507, F1: 42, F2: ((!false) || (!true))}
    switch v1 {
    case 0:
        if true {
            v1 = (v2 / ((v2) | 1))
        }
    case -9223372036854775808:
        if true {
            if true {
                v1 = -9223372036854775808
                v2 += (-(v2))
                fmt.Println(v1)
            } else {
                v2 += (v1 + (v2 / ((892) | 1)))
                var v26 uint8 = 71
                _ = v26
            }
            var v27 int32 = -1
            _ = v27
        }
    }
    fmt.Println(v2)
    switch "" {
    case "foo bar":
        v28 := GenS2{F0: 18446744073709551615, F1: 127, F2: 0}
        _ = v28
    }
    fmt.Println(v1, v2)
    if (0 == -9223372036854775808) {
        v1 += v1
        fmt.Println(v2, v1)
        switch "hello" {
        case "a":
            var v29 uint8 = 2
            _ = v29
        case "\n\t":
            if false {
                var v30 string = (("\n\t" + "foo bar") + ("hello" + ("\n\t" + "")))
                _ = v30
            } else {
                v2 = (v1 + -74)
                var v31 int32 = 1
                fmt.Println(v31)
                _ = v31
            }
        case "hello":
            fmt.Println(v1, v2)
        }
    } else {
        v1 = (v2 % ((v2) | 1))
        v32 := GenS4{F0: 355}
        if true {
            fmt.Println(v2)
            v33 := []int{9223372036854775807}
            _ = v33
        }
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v6
}
