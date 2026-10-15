package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 0
    v3 := []uint32{}
    v2 += 591
    v4 := GenS1{F0: 9223372036854775807, F1: 837}
    v2 = v2
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := v1[1]
        for fi7 := 0; fi7 < 11; fi7++ {
            switch 457 {
            case 9223372036854775807:
                v8 := GenS1{F0: (fi5 - (fi7 >> (((^(v2))) & 63))), F1: 2147483647}
                _ = v8
            case 294:
                v6 = v6
            case 32767:
                fmt.Println(fi5)
            }
            v9 := GenS1{F0: 0, F1: 771}
            _ = fi7
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10)
        var v11 int64 = 42
        _ = fi10
        _ = v11
    }
    var v12 string = "hello"
    switch (v12 + v12) {
    case "a":
        var v13 string = (v12 + "\n\t")
        _ = v13
    case "foo bar":
        v2 -= 169
    case "hello":
        var v14 int64 = -104
        _ = v14
    default:
        fmt.Println(v2)
    }
    delete(v1, 1)
    if (32767 <= 32767) {
        v12 = "hello"
        if false {
            var v15 float64 = 0.0
            _ = v15
        }
    } else {
        fmt.Println(v12, v2)
        v2 = v2
    }
    switch -2147483648 {
    case -267:
        var v16 int8 = 1
        _ = v16
    }
    var v17 uint16 = 192
    var v18 int32 = 0
    var v19 int16 = 1000
    v20 := GenS2{F0: 100, F1: 1}
    if false {
        switch 907 {
        case 0:
            if (1 > 2147483648) {
                if true {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v18 += -1
                        for fi22 := 0; fi22 < 7; fi22++ {
                            switch -2147483648 {
                            case 198:
                                fmt.Println(v2, v2)
                            case 32768:
                                if false {
                                    switch 246 {
                                    case 42:
                                        fmt.Println(fi21, v17, v12)
                                    case 2147483647:
                                        v19 -= v19
                                    case -368:
                                        fmt.Println(v12, v2, fi22)
                                    default:
                                        var v23 int16 = (v19 % ((v19) | 1))
                                        _ = v23
                                    }
                                    if (uint8(v17) < 89) {
                                        var v24 bool = (v19 >= v19)
                                        v19 += int16(v17)
                                        _ = v24
                                    }
                                }
                            default:
                                var v25 uint8 = 0
                                _ = v25
                            }
                            var v26 uint = v2
                            var v27 int32 = 1
                            _ = fi22
                            _ = v26
                            _ = v27
                        }
                        switch "foo bar" {
                        case "a":
                            delete(v1, 4294967295)
                        case "":
                            fmt.Println(v19, v12, fi21)
                        case "\n\t":
                            for fi28 := 0; fi28 < 5; fi28++ {
                                v2 = (v2 | v2)
                                _ = fi28
                            }
                        default:
                            if (744 != 42) {
                                for fi29 := 0; fi29 < 10; fi29++ {
                                    switch "hello" {
                                    case "hello":
                                        var v30 bool = (18446744073709551615 != 0)
                                        _ = v30
                                    }
                                    _ = fi29
                                }
                                v2 = 2
                            }
                        }
                        _ = fi21
                    }
                }
                for fi31 := 0; fi31 < 10; fi31++ {
                    if true {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v17, v17, v19)
                        case "foo bar":
                            var v32 int64 = 0
                            _ = v32
                        }
                        switch -9223372036854775808 {
                        case -854:
                            v17 += v17
                        default:
                            if false {
                                var v33 int64 = 616
                                var v34 int64 = 1
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                    _ = fi31
                }
            } else {
                switch ("foo bar" + "") {
                case "hello":
                    var v35 int = 4294967295
                    _ = v35
                }
            }
        default:
            for fi36 := 0; fi36 < 19; fi36++ {
                _ = len(v3)
                _ = fi36
            }
        }
        _ = len(v1)
        v17 -= v17
    }
    var v37 float64 = -2.25
    v37 = (v37 - 1e10)
    for fi38 := 0; fi38 < 8; fi38++ {
        v19 += v19
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
}
