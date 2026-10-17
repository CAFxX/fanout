package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }

func main() {
    v1 := []uint32{2}
    var v2 uint8 = 1
    var v3 int8 = 1
    var v4 float64 = 1e10
    v5 := []string{"a", "\n\t"}
    var v6 int8 = v3
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint32 = 2
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 := GenS1{F0: 32767, F1: 65535, F2: 0}
    if false {
        var v11 int8 = v6
        var v12 uint16 = 387
        if (!((false || false) && (!(165 != 65535)))) {
            var v13 string = ""
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS1{F0: -341, F1: v12, F2: 18446744073709551615}
                for fi16 := 0; fi16 < 5; fi16++ {
                    v13 = "a"
                    if (!false) {
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                    }
                    v17 := GenS1{F0: 1000, F1: (v12 ^ v12), F2: 42}
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            var v18 string = "a"
            _ = v18
        }
        _ = v11
        _ = v12
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(fi19, v3)
        _ = fi19
    }
    v20 := GenS1{F0: 1, F1: 0, F2: 58}
    v4 += v4
    v2 += v2
    var v21 string = ("a" + ("hello" + "hello"))
    switch "hello" {
    case "":
        fmt.Println(v2, v3, v6)
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        fmt.Println(v6, v3)
    default:
        if (false && true) {
            var v22 string = v21
            for fi23 := 0; fi23 < 2; fi23++ {
                if (52 < 1) {
                    v5 = append(v5, v22)
                    if (42 != 2147483648) {
                        var v24 int64 = -118
                        if true {
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            var v25 int64 = (^(v24))
                            _ = v25
                        }
                        _ = v24
                    }
                } else {
                    if (!false) {
                        var v26 float32 = 1.5
                        fmt.Println(v3)
                        v21 = (v21 + ((v21 + v22) + (v21 + ("a" + v22))))
                        _ = v26
                    } else {
                        var v27 float64 = v4
                        if false {
                            v3 = v3
                        }
                        _ = v27
                    }
                }
                fmt.Println(v22, v3, v4)
                _ = fi23
            }
            _ = v22
        }
    }
    switch v21 {
    case "\n\t":
        if true {
            if len(v1) > 0 {
                v1[0] = 331
            }
        } else {
            fmt.Println(v4)
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int16 = -382
                _ = fi28
                _ = v29
            }
        }
    case "hello":
        v21 = ("\n\t" + v21)
    default:
        fmt.Println(v3)
    }
    v4 += ((v4 - -2.25) + ((v4 - 1.5) / v4))
    if true {
        v6 = (v6 - 115)
    }
    v30 := GenS1{F0: 1, F1: 469, F2: 418}
    var v31 bool = (-2147483648 == -62)
    v2 = 125
    switch v21 {
    case "\n\t":
        v3 = (v3 >> ((42) & 63))
    case "foo bar":
        if (42 <= 42) {
            v4 += v4
        } else {
            v21 = v21
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    var v32 string = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
