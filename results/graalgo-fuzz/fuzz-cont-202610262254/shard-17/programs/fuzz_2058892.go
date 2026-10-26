package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    var v1 int8 = -128
    v2 := []string{"", "foo bar"}
    var v3 bool = (!("" > "hello"))
    var v4 int32 = -72
    var v5 int8 = (v1 << ((0) & 63))
    fmt.Println(v1, v1)
    if (!(!v3)) {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 bool = v3
            for fi8 := 0; fi8 < 17; fi8++ {
                switch "hello" {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                case "":
                    v5 -= 42
                case "foo bar":
                    switch "hello" {
                    case "hello":
                        v9 := GenS1{F0: false, F1: -1, F2: 255}
                        _ = v9
                    case "a":
                        var v10 string = "hello"
                        _ = v10
                    default:
                        v1 = (v1 / ((127) | 1))
                    }
                default:
                    var v11 float64 = (0.1 / 0.1)
                    _ = v11
                }
                var v12 string = ("hello" + (("foo bar" + ("a" + "foo bar")) + "\n\t"))
                _ = fi8
                _ = v12
            }
            _ = fi6
            _ = v7
        }
    }
    if v3 {
        var v13 uint16 = 32767
        _ = v13
    }
    var v14 float64 = 3.14159
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 int = -9223372036854775808
        var v17 int8 = v5
        v16 = (fi15 % ((fi15) | 1))
        _ = fi15
        _ = v16
        _ = v17
    }
    switch "foo bar" {
    case "a":
        var v18 int32 = v4
        _ = v18
    case "foo bar":
        switch "" {
        case "a":
            var v19 uint32 = 825
            _ = v19
        case "":
            if true {
                var v20 uint8 = 78
                _ = v20
            } else {
                var v21 uint8 = 2
                v21 -= v21
                _ = v21
            }
        }
    }
    var v22 int64 = -9223372036854775808
    for fi23 := 0; fi23 < 18; fi23++ {
        if v3 {
            v2 = append(v2, "a")
        } else {
            var v24 bool = true
            _ = v24
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            if (!(v3 && v3)) {
                v22 += -576
                if (v3 && (!v3)) {
                    v26 := GenS1{F0: (!false), F1: 1, F2: 498}
                    _ = v26
                }
            } else {
                v4 = (v4 & (v4 << ((65535) & 63)))
                v1 -= (v5 << ((864) & 63))
            }
            fmt.Println(v5, v14)
            var v27 float32 = (0.0 - (1.5 - (-2.25 / (1e10 - (0.1 + 1.5)))))
            var v28 int64 = (v22 / (((v22 | (v22 ^ (v22 >> ((4294967295) & 63))))) | 1))
            fmt.Println(v28)
            v29 := GenS1{F0: ("hello" <= "hello"), F1: 32767, F2: 1}
            _ = fi25
            _ = v27
            _ = v28
            _ = v29
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi23
    }
    fmt.Println(v14, v5)
    for fi30 := 0; fi30 < 8; fi30++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v31 uint32 = 737
        if v3 {
            if v3 {
                _ = len(v2)
            }
            var v32 int16 = -562
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v5, v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
