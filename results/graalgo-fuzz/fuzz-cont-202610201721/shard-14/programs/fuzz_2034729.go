package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := []uint32{}
    v2 := []uint32{42, 4294967295}
    v3 := []uint32{4294967295, 2}
    var v4 uint64 = 1
    var v5 string = "a"
    var v6 int = 32768
    fmt.Println(v5, v5, v4)
    if (3.14159 < ((3.14159 + 0.0) * 1.5)) {
        fmt.Println(v5)
        var v7 uint32 = 4294967295
        if ((v6 >> ((343) & 63)) != (-(v6))) {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 bool = true
                var v10 uint = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
            v5 = (((((("" + (v5 + v5)) + v5) + "a") + v5) + "\n\t") + ("hello" + v5))
        } else {
            switch v5 {
            case "a":
                var v11 int8 = -50
                _ = v11
            case "\n\t":
                v7 = v7
            }
        }
        _ = v7
    } else {
        v12 := GenS1{F0: 1, F1: 0, F2: (v5 + (v5 + (v5 + v5)))}
        v2 = append(v2, 845)
        if (!false) {
            _ = len(v2)
            v13 := GenS1{F0: 42, F1: 376, F2: ("" + ("foo bar" + ("\n\t" + v5)))}
            _ = v13
        }
        _ = v12
    }
    v6 -= v6
    v4 = (v4 & v4)
    if (false || true) {
        var v14 uint = 791
        var v15 uint8 = 255
        var v16 float64 = 3.14159
        _ = v14
        _ = v15
        _ = v16
    } else {
        if false {
            fmt.Println(v4, v6, v4)
        }
    }
    var v17 uint16 = 255
    if (-128 > 53) {
        for fi18 := 0; fi18 < 6; fi18++ {
            v6 = (v6 % (((-(v6))) | 1))
            _ = len(v2)
            v19 := GenS1{F0: 0, F1: 42, F2: ""}
            _ = fi18
            _ = v19
        }
        if (3.14159 <= -2.25) {
            v6 = -176
            for fi20 := 0; fi20 < 5; fi20++ {
                if false {
                    v6 -= 128
                } else {
                    v2 = append(v2, 1)
                    var v21 int = (v6 * v6)
                    _ = v21
                }
                switch fi20 {
                case 4294967295:
                    if false {
                        var v22 string = v5
                        var v23 bool = true
                        v3 = append(v3, 0)
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            _ = len(v1)
                            _ = fi24
                        }
                        v25 := GenS2{F0: 0, F1: -33}
                        _ = len(v3)
                        _ = v25
                    }
                case 42:
                    v4 = v4
                case -9223372036854775808:
                    var v26 float64 = (3.14159 * 1e10)
                    _ = v26
                }
                _ = fi20
            }
        } else {
            var v27 int8 = 0
            _ = v27
        }
        var v28 int8 = 11
        _ = v28
    }
    fmt.Println(v5, v5)
    var v29 uint32 = 56
    if false {
        var v30 string = v5
        var v31 int = (v6 - (v6 & v6))
        _ = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
