package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0() uint16 {
    var v5 int64 = 42
    fmt.Println(v5)
    if true {
        var v6 float64 = float64(-2.25)
        switch 2147483647 {
        case -855:
            fmt.Println(v6)
        case 42:
            v5 += v5
        case 1:
            var v7 uint32 = 147
            _ = v7
        default:
            if false {
                switch ("a" + "\n\t") {
                case "hello":
                    v5 += v5
                case "foo bar":
                    v5 -= v5
                case "a":
                    v8 := []int{32767, 32768, -9223372036854775808}
                    _ = v8
                default:
                    fmt.Println(v5)
                }
                if (42 != 0) {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v5 = v5
                        v10 := GenS1{F0: 2, F1: 779}
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: 0, F1: 1}
                            v13 := GenS1{F0: 65535, F1: 42}
                            switch "hello" {
                            case "\n\t":
                                var v14 uint16 = 2
                                _ = v14
                            default:
                                var v15 uint = 1
                                _ = v15
                            }
                            fmt.Println(fi11)
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        _ = fi9
                        _ = v10
                    }
                    switch "" {
                    case "":
                        v5 += 410
                    case "hello":
                        v16 := GenS1{F0: 511, F1: 714}
                        _ = v16
                    case "foo bar":
                        if false {
                            if false {
                                v6 = float64(3.14159)
                            } else {
                                v17 := GenS2{F0: 255, F1: -128}
                                _ = v17
                            }
                        } else {
                            v18 := []int{-9223372036854775808, 127, 808}
                            if true {
                                switch ("" + "foo bar") {
                                case "hello":
                                    var v19 uint64 = 4294967295
                                    _ = v19
                                case "":
                                    v5 -= v5
                                default:
                                    if true {
                                        if false {
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            v20 := GenS1{F0: 593, F1: 58}
                                            _ = v20
                                        }
                                        v5 = 1
                                    }
                                }
                                fmt.Println(v6, v6)
                            } else {
                                fmt.Println(v5, v5)
                                var v21 uint8 = 128
                                fmt.Println(v21)
                                _ = v21
                            }
                            _ = v18
                        }
                    }
                    if ((true || ("a" <= ("foo bar" + "hello"))) && false) {
                        var v22 bool = true
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: 18446744073709551615, F1: 0}
                    _ = v23
                }
            } else {
                var v24 int8 = -1
                var v25 int64 = v5
                _ = v24
                _ = v25
            }
        }
        _ = v6
    }
    v5 = (^(v5))
    _ = v5
    return 1
}

func main() {
    var v1 int8 = -128
    v2 := []uint32{65535}
    var v3 uint64 = 4294967295
    var v4 string = (("\n\t" + ("" + "a")) + ("\n\t" + ("" + "")))
    v4 = "\n\t"
    v2 = append(v2, 4294967295)
    fmt.Println(v4, v4)
    if false {
        var v26 int = 32768
        _ = v26
    }
    v1 = -128
    if (((!(2 > 219)) && ((-2147483648 > 32768) && (1 < 440))) && false) {
        var v27 int32 = 0
        _ = v27
    } else {
        fmt.Println(v1, v4, v3)
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 uint = 255
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 float64 = -2.25
        _ = fi30
        _ = v31
    }
    v32 := GenS1{F0: 4294967295, F1: 551}
    var v33 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
