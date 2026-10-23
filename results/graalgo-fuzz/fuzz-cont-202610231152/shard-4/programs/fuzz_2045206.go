package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint = 760
    var v5 int8 = 127
    if false {
        v5 = (v5 & v5)
        switch "" {
        case "\n\t":
            v1[65535] = ("hello" + "foo bar")
        }
        if false {
            v5 -= v5
        } else {
            fmt.Println(v5, v4)
            fmt.Println(v4, v5)
            fmt.Println(v4)
        }
    } else {
        var v6 uint64 = 4294967295
        fmt.Println(v4)
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "a" {
        case "hello":
            if true {
                v5 = 1
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := GenS1{F0: -9223372036854775808, F1: v4, F2: 4294967295}
                    if ((fi7 >> ((1) & 63)) >= 440) {
                        v5 = -128
                        var v10 uint8 = 128
                        var v11 uint8 = (v10 ^ 255)
                        _ = v10
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
            }
        }
        var v12 uint32 = 216
        var v13 uint32 = v12
        _ = fi7
        _ = v12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 13; fi14++ {
            v5 += (v5 - v5)
            delete(v1, 1)
            _ = fi14
        }
    }
    v15 := GenS2{F0: -2147483648, F1: (-9223372036854775808 >= -1), F2: 32767}
    v4 = (^(v4))
    fmt.Println(v4, v4)
    for fi16 := 0; fi16 < 11; fi16++ {
        if (fi16 <= fi16) {
            fmt.Println(v5, v5, v4)
            var v17 float32 = (1e10 - ((1.5 + 0.1) + 3.14159))
            _ = v17
        } else {
            v18 := GenS1{F0: 42, F1: (v4 >> ((v4) & 63)), F2: 4294967295}
            var v19 float32 = 1e10
            _ = v18
            _ = v19
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = fi16
    }
    if true {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 int8 = (v5 & v5)
            _ = fi20
            _ = v21
        }
        v5 += (v5 ^ 127)
        var v22 float64 = ((3.14159 * float64(0.0)) - (1.5 - (1e10 * 1.5)))
        _ = v22
    }
    var v23 bool = false
    if v23 {
        v23 = (653 == (v4 ^ v4))
    }
    var v24 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
