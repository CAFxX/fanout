package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 int = 680
    var v2 float32 = (1.5 - 1.5)
    v3 := make(map[int]uint32)
    var v4 int16 = 38
    var v5 float32 = 1.5
    var v6 float64 = (1.5 * (0.1 * 0.0))
    if true {
        v1 = v1
    } else {
        fmt.Println(v4, v1, v5)
        v3[v1] = 2147483648
    }
    v5 = v5
    if false {
        var v7 bool = false
        switch ("\n\t" + ("foo bar" + "hello")) {
        case "hello":
            if (42 <= 220) {
                fmt.Println(v1)
            } else {
                var v8 uint64 = 0
                switch (v1 | 4294967295) {
                case 2147483647:
                    if v7 {
                        for fi9 := 0; fi9 < 10; fi9++ {
                            var v10 int64 = 0
                            v3[(v1 / (((fi9 / ((v1) | 1))) | 1))] = 65535
                            v5 = v2
                            _ = fi9
                            _ = v10
                        }
                        var v11 int64 = 42
                        _ = v11
                    } else {
                        var v12 int32 = -1
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v7 = true
                            fmt.Println(v5, v5)
                            _ = fi13
                        }
                        _ = v12
                    }
                default:
                    v2 -= v2
                }
                _ = v8
            }
        case "a":
            switch 127 {
            case 32768:
                v7 = (v7 || (-44 == 0))
            case -2147483648:
                fmt.Println(v5)
            case 199:
                v14 := GenS1{F0: 359, F1: v4}
                _ = v14
            }
        case "foo bar":
            if (10 <= -666) {
                var v15 int32 = 548
                _ = len(v3)
                for fi16 := 0; fi16 < 11; fi16++ {
                    switch "" {
                    case "\n\t":
                        var v17 int16 = v4
                        _ = v17
                    case "":
                        fmt.Println(v6)
                    }
                    _ = fi16
                }
                _ = v15
            } else {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint32 = 4294967295
                    _ = fi18
                    _ = v19
                }
            }
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            v5 = (0.1 / (v5 * v5))
            _ = fi20
        }
        _ = v7
    }
    switch "\n\t" {
    case "hello":
        v21 := GenS1{F0: 841, F1: 1000}
        _ = v21
    default:
        switch ("hello" + "a") {
        case "":
            v22 := GenS2{F0: v1}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if (2147483647 > -2147483648) {
            v24 := []int{32768}
            _ = v24
        }
        _ = fi23
    }
    var v25 int8 = 110
    var v26 uint8 = 42
    v27 := GenS1{F0: 705, F1: (v4 * (v4 % (((^(v4))) | 1)))}
    var v28 int64 = 42
    var v29 uint32 = 2147483647
    v29 = 2147483647
    v2 -= v2
    v2 -= 3.14159
    if false {
        v26 = (v26 - v26)
    } else {
        var v30 float32 = v2
        _ = v30
    }
    delete(v3, v1)
    v31 := GenS1{F0: 836, F1: (v4 % (((v4 % ((32767) | 1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
