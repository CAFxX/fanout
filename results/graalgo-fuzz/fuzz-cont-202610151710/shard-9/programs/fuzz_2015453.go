package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    var v4 uint = 2
    var v5 uint64 = 18446744073709551615
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 bool, p1 uint) int16 {
    v6 := []int{}
    v6 = append(v6, 2147483647)
    _ = p0
    _ = p1
    _ = v6
    return 1000
}

func main() {
    var v1 uint8 = 42
    var v2 float32 = (0.0 - 0.0)
    var v3 bool = (100 <= 127)
    v2 = 1.5
    v7 := []int{32768, -9223372036854775808}
    v1 = v1
    if ((v1 ^ 183) > 255) {
        switch 9223372036854775807 {
        case 2147483647:
            switch "foo bar" {
            case "\n\t":
                v1 = 1
            }
        case 9223372036854775807:
            var v8 float64 = 3.14159
            _ = v8
        default:
            v1 = (v1 - (v1 % (((v1 - (v1 | v1))) | 1)))
        }
        v1 -= v1
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 string = "a"
            var v11 int = 32767
            if len(v7) > 0 {
                v7[0] = (fi9 | fi9)
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    switch "" {
    case "hello":
        var v12 float64 = ((0.1 / (3.14159 / (1.5 * 0.1))) * 0.1)
        _ = v12
    default:
        if false {
            switch -1 {
            case 32768:
                for fi13 := 0; fi13 < 1; fi13++ {
                    fmt.Println(v3, v2, v3)
                    v2 = 1e10
                    v14 := GenS1{F0: 0, F1: (v1 * (v1 * (^(v1)))), F2: 42}
                    _ = fi13
                    _ = v14
                }
            case -1:
                if len(v7) > 0 {
                    v7[0] = -62
                }
            case 383:
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        switch "hello" {
                        case "a":
                            switch ("\n\t" + "hello") {
                            case "foo bar":
                                v1 -= 1
                            case "":
                                v17 := GenS1{F0: 42, F1: v1, F2: -276}
                                _ = v17
                            default:
                                if (18446744073709551615 >= 0) {
                                    if (!v3) {
                                        v18 := GenS1{F0: -1, F1: (^(v1)), F2: 9223372036854775807}
                                        var v19 bool = v3
                                        var v20 uint8 = 74
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        if false {
                                            _ = len(v7)
                                            var v21 int = fi16
                                            v22 := GenS2{F0: 2}
                                            _ = v21
                                            _ = v22
                                        } else {
                                            fmt.Println(v3)
                                            if (42 == genFn1(v3, 4294967295)) {
                                                v3 = (!false)
                                            } else {
                                                _ = len(v7)
                                                v7 = append(v7, (fi15 - fi15))
                                            }
                                            v7 = append(v7, fi15)
                                            v1 += (v1 % ((170) | 1))
                                        }
                                        if v3 {
                                            fmt.Println(v3, v2, v3)
                                            v23 := GenS3{F0: -1}
                                            var v24 float64 = 0.1
                                            _ = v23
                                            _ = v24
                                        }
                                        v3 = false
                                        v7 = append(v7, fi15)
                                    }
                                }
                            }
                        case "foo bar":
                            fmt.Println(fi16, v3, v1)
                        default:
                            fmt.Println(v3, v3)
                        }
                        _ = fi16
                    }
                    _ = fi15
                }
            default:
                var v25 uint32 = 1
                _ = v25
            }
        } else {
            if v3 {
                for fi26 := 0; fi26 < 18; fi26++ {
                    v2 = v2
                    _ = fi26
                }
            }
            var v27 string = "hello"
            _ = v27
        }
    }
    if v3 {
        var v28 uint = 255
        _ = v28
    } else {
        var v29 uint = 990
        switch 42 {
        case 0:
            v2 = v2
        case -88:
            var v30 uint32 = 686
            _ = v30
        default:
            v3 = false
        }
        v1 += 0
        _ = v29
    }
    var v31 int16 = 42
    if v3 {
        fmt.Println(v2)
    }
    if ((float32(1.5) - v2) >= v2) {
        v7 = append(v7, 32767)
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    switch ("hello" + "hello") {
    case "foo bar":
        if (1 == 0) {
            v1 = v1
        }
    case "\n\t":
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 uint = 663
            for fi34 := 0; fi34 < 12; fi34++ {
                if len(v7) > 0 {
                    v7[0] = fi32
                }
                _ = fi34
            }
            var v35 bool = v3
            _ = fi32
            _ = v33
            _ = v35
        }
    default:
        _ = len(v7)
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v7
}
