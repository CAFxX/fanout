package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    var v1 bool = (-128 == 100)
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int16 = 273
        _ = len(v3)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint16 = 1
            _ = fi7
            _ = v8
        }
        v3[757] = 1
        for fi9 := 0; fi9 < 16; fi9++ {
            if false {
                if v1 {
                    if v1 {
                        if false {
                            v1 = v1
                            v2 = "\n\t"
                            v3[2] = 1
                        }
                        var v10 uint = 1
                        switch (fi6 / ((fi9) | 1)) {
                        case 721:
                            v11 := GenS1{F0: v2, F1: 128}
                            _ = v11
                        case -207:
                            switch "\n\t" {
                            case "hello":
                                v10 -= v10
                            default:
                                v10 = v10
                            }
                        case 128:
                            var v12 int = (^(fi6))
                            _ = v12
                        }
                        v13 := GenS2{F0: 127, F1: 127}
                        _ = v10
                        _ = v13
                    } else {
                        var v14 string = ("a" + "\n\t")
                        _ = v14
                    }
                    fmt.Println(v1)
                }
            } else {
                delete(v3, 42)
                fmt.Println(fi9, v1)
            }
            _ = fi9
        }
        _ = fi6
    }
    fmt.Println(v2, v2)
    fmt.Println(v1)
    for fi15 := 0; fi15 < 19; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    if v1 {
        for fi16 := 0; fi16 < 7; fi16++ {
            v2 = v2
            _ = fi16
        }
    } else {
        v17 := GenS2{F0: -1, F1: -103}
        var v18 bool = (65535 != 2)
        _ = v17
        _ = v18
    }
    if (1 > 837) {
        for fi19 := 0; fi19 < 15; fi19++ {
            switch (v2 + "a") {
            case "foo bar":
                v20 := GenS2{F0: 0, F1: 122}
                _ = v20
            }
            v1 = v1
            _ = fi19
        }
    } else {
        fmt.Println(v1, v2)
        v21 := []int{1}
        v1 = v1
        _ = v21
    }
    v1 = false
    v3[146] = 566
    switch v2 {
    case "\n\t":
        if true {
            if false {
                if (1 <= 70) {
                    v22 := GenS2{F0: 34, F1: 68}
                    if (!v1) {
                        var v23 string = v2
                        var v24 float64 = (1e10 * 1e10)
                        var v25 uint32 = 1
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        v3[42] = 2147483648
                        var v26 float64 = 3.14159
                        _ = v26
                    }
                    _ = len(v3)
                    _ = v22
                }
                fmt.Println(v2)
                v2 = (v2 + (v2 + ""))
            }
            v1 = (105 < 42)
            v1 = ((0 == 18446744073709551615) && v1)
        }
    default:
        v1 = v1
    }
    v2 = ((v2 + v2) + v2)
    var v27 int16 = 32767
    for fi28 := 0; fi28 < 14; fi28++ {
        delete(v3, 702)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
