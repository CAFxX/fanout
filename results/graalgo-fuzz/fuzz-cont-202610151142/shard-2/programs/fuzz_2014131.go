package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := []int{-950}
    var v2 float64 = ((0.0 * 1e10) - -2.25)
    var v3 float64 = v2
    v4 := make(map[uint32]string)
    if true {
        var v5 int = -380
        _ = v5
    } else {
        var v6 int8 = -1
        v7 := GenS1{F0: 87}
        switch 42 {
        case 32768:
            var v8 uint16 = 0
            _ = v8
        default:
            v2 = v2
        }
        _ = v6
        _ = v7
    }
    var v9 int32 = -1
    if false {
        v2 -= 1e10
        if len(v1) > 0 {
            v1[0] = 323
        }
        var v10 uint8 = 42
        _ = v10
    } else {
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                for fi12 := 0; fi12 < 15; fi12++ {
                    switch "\n\t" {
                    case "":
                        if false {
                            switch fi11 {
                            case 1:
                                if true {
                                    switch (fi11 / ((125) | 1)) {
                                    case 128:
                                        v2 -= v3
                                    case 2147483647:
                                        if ((-2.25 - float64((((float32(((v2 - v2) * float64((float32((float64((1.5 + (3.14159 / -2.25))) / (v2 * v2))) + -2.25)))) - 0.1) * ((1.5 * 0.1) + (float32(v3) - (0.0 - 1.5)))) - 0.1))) <= float64((((-2.25 - 1.5) / 1e10) + float32((float64(0.1) + -2.25))))) {
                                            v2 = float64((-2.25 + (((0.0 + 3.14159) / (float32(3.14159) / 1e10)) - 1.5)))
                                            v13 := GenS1{F0: 255}
                                            var v14 float64 = v2
                                            _ = v13
                                            _ = v14
                                        } else {
                                            if true {
                                                v3 += float64((-2.25 / 0.1))
                                                v9 = 42
                                                v15 := v4[42]
                                                _ = v15
                                            }
                                            var v16 int64 = -1
                                            v16 += (v16 | v16)
                                            v2 -= v3
                                            switch (fi11 >> ((18446744073709551615) & 63)) {
                                            case 127:
                                                var v17 int64 = (v16 / ((v16) | 1))
                                                _ = v17
                                            case 436:
                                                _ = len(v4)
                                            case 553:
                                                var v18 string = ""
                                                _ = v18
                                            }
                                            _ = v16
                                        }
                                    default:
                                        fmt.Println(v9, fi11, fi12)
                                    }
                                    v3 = float64((float32(((0.1 + (3.14159 / (v3 - 0.0))) - 0.1)) + (0.1 / 0.1)))
                                    var v19 int16 = -1
                                    _ = v19
                                }
                            case -123:
                                delete(v4, 1)
                            default:
                                var v20 float64 = 1e10
                                _ = v20
                            }
                        } else {
                            switch ("foo bar" + ("\n\t" + (("" + ("hello" + ("" + (("foo bar" + ("\n\t" + "")) + "a")))) + "hello"))) {
                            case "foo bar":
                                v21 := GenS1{F0: 32767}
                                _ = v21
                            default:
                                fmt.Println(fi11, v2)
                            }
                        }
                    case "a":
                        v4[0] = "a"
                    }
                    var v22 int8 = -19
                    _ = fi12
                    _ = v22
                }
                v3 = v2
                var v23 int32 = v9
                _ = fi11
                _ = v23
            }
        }
    }
    v24 := GenS1{F0: 32767}
    v9 -= v9
    fmt.Println(v2)
    _ = len(v4)
    if (620 <= 32767) {
        v2 = (1.5 * ((v3 * 1.5) * v3))
        v2 -= v2
    }
    if true {
        if (true && (!(1e10 <= 1.5))) {
            v3 += v2
            v3 = (0.0 / v2)
            for fi25 := 0; fi25 < 8; fi25++ {
                delete(v4, 4294967295)
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch fi25 {
                    case 1:
                        var v27 int64 = 0
                        _ = v27
                    }
                    if (1e10 != 1.5) {
                        v2 -= v2
                        switch 2147483647 {
                        case 527:
                            fmt.Println(v2, v2)
                        case 42:
                            v2 = 1.5
                        default:
                            v3 -= float64(3.14159)
                        }
                        v28 := GenS1{F0: 2}
                        _ = v28
                    } else {
                        switch ("a" + "") {
                        case "\n\t":
                            v29 := GenS2{F0: "", F1: 9223372036854775807}
                            _ = v29
                        }
                        v2 = float64(0.1)
                    }
                    _ = fi26
                }
                fmt.Println(v3, v3)
                _ = fi25
            }
        }
        _ = len(v4)
        fmt.Println(v3, v3)
    } else {
        fmt.Println(v9, v2)
    }
    var v30 float32 = ((-2.25 * 0.0) / 0.1)
    v9 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
