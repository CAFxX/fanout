package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (3.14159 + 0.0)
    v3 := make(map[uint32]uint32)
    v4 := []int{128, 1, 127}
    fmt.Println(v2)
    if (!((1.5 > -2.25) && false)) {
        v2 += (v2 + v2)
    }
    if (false || (false || false)) {
        var v5 int64 = 42
        _ = v5
    } else {
        if (!false) {
            switch (("a" + "\n\t") + "\n\t") {
            case "hello":
                if true {
                    _ = len(v3)
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v6 := GenS1{F0: -1, F1: -2147483648, F2: 978}
                    _ = v6
                }
            case "\n\t":
                switch ("a" + "") {
                case "\n\t":
                    for fi7 := 0; fi7 < 19; fi7++ {
                        for fi8 := 0; fi8 < 7; fi8++ {
                            var v9 uint32 = 4294967295
                            if false {
                                var v10 int64 = 1
                                var v11 uint8 = 148
                                _ = v10
                                _ = v11
                            }
                            _ = fi8
                            _ = v9
                        }
                        _ = fi7
                    }
                case "foo bar":
                    v12 := GenS1{F0: -832, F1: 1, F2: 42}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        switch "" {
                        case "a":
                            if false {
                                var v14 string = ""
                                fmt.Println(fi13, v2, v2)
                                fmt.Println(fi13, v14, fi13)
                                _ = v14
                            }
                        case "hello":
                            var v15 int8 = 0
                            _ = v15
                        case "":
                            if (!true) {
                                fmt.Println(v2)
                                v4 = append(v4, (fi13 << ((221) & 63)))
                            }
                        default:
                            if true {
                                if true {
                                    fmt.Println(fi13, fi13)
                                    v2 -= (float32(1e10) - (v2 + v2))
                                }
                            }
                        }
                        _ = fi13
                    }
                }
            case "":
                fmt.Println(v2)
            default:
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint8 = 245
                    _ = fi16
                    _ = v17
                }
            }
            v18 := GenS1{F0: 42, F1: -1, F2: 1}
            var v19 uint64 = 1
            _ = v18
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS1{F0: (fi20 >> ((569) & 63)), F1: 831, F2: 32767}
                var v22 string = "\n\t"
                if true {
                    var v23 uint16 = 465
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v2 = v2
                    _ = v23
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            _ = len(v1)
        }
    }
    var v24 int8 = 100
    var v25 bool = (-1 >= 1)
    fmt.Println(v2)
    _ = len(v1)
    v25 = true
    var v26 uint64 = 0
    fmt.Println(v25, v24)
    var v27 uint32 = 0
    v24 -= (v24 | 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
