package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{42}
    var v3 bool = (((1e10 + 0.0) >= 1.5) && (1 >= 65535))
    var v4 int = 199
    var v5 uint64 = 42
    v6 := []uint32{2147483647}
    if ((true || (v3 || false)) || v3) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v1 += v1
            _ = fi7
        }
        v8 := GenS1{F0: 42}
        _ = v8
    }
    var v9 string = ("hello" + ("foo bar" + "a"))
    _ = len(v2)
    for fi10 := 0; fi10 < 1; fi10++ {
        if (v5 >= v5) {
            var v11 bool = (!(v3 || v3))
            var v12 string = v9
            if ((2147483647 <= 376) && (!v3)) {
                fmt.Println(v1, v12)
                v5 = 932
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
            }
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: 10, F1: (v5 | 1), F2: 0}
            v1 -= v1
            _ = v13
        }
        if false {
            fmt.Println(v9)
            switch v9 {
            case "":
                var v14 int8 = -51
                _ = v14
            default:
                fmt.Println(fi10, v5)
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 2
            }
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 -= float64(1e10)
            switch fi10 {
            case 4294967295:
                if ((463 >= 0) || v3) {
                    v1 = float64(-2.25)
                    if (!((!(255 < 128)) || (false || false))) {
                        switch v9 {
                        case "\n\t":
                            var v16 int16 = 0
                            _ = v16
                        case "hello":
                            fmt.Println(fi10)
                        default:
                            v2 = append(v2, 1)
                            fmt.Println(v5, fi10, v5)
                            var v17 int8 = -1
                            _ = v17
                        }
                    }
                }
            }
            _ = fi15
        }
        _ = fi10
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint = 255
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint64 = v5
        var v22 int8 = 42
        v21 += 2
        _ = fi20
        _ = v21
        _ = v22
    }
    if v3 {
        fmt.Println(v9, v9)
        var v23 uint8 = 2
        _ = v23
    }
    if ((!v3) && true) {
        if (!((!true) && true)) {
            v24 := GenS1{F0: -128}
            var v25 int = v4
            _ = v24
            _ = v25
        } else {
            if false {
                v26 := GenS1{F0: -78}
                v6 = append(v6, 983)
                v27 := GenS3{F0: 0}
                _ = v26
                _ = v27
            } else {
                v28 := GenS2{F0: -1, F1: v5, F2: 0}
                if v3 {
                    var v29 int8 = -1
                    _ = v29
                }
                _ = v28
            }
        }
        var v30 uint = 492
        _ = v30
    } else {
        var v31 uint16 = 153
        _ = len(v6)
        switch (v4 * -371) {
        case 1:
            _ = len(v6)
        case 128:
            var v32 int16 = 260
            _ = v32
        case -9223372036854775808:
            if (127 != v4) {
                v9 = v9
            }
        default:
            fmt.Println(v4, v31)
        }
        _ = v31
    }
    _ = len(v2)
    v9 = v9
    var v33 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
