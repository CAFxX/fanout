package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 string; F1 bool; F2 uint }

func main() {
    var v1 float64 = 1.5
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    v4 := []string{"a", "\n\t", "hello"}
    for fi5 := 0; fi5 < 7; fi5++ {
        _ = len(v4)
        v6 := GenS1{F0: 65535, F1: 42}
        v7 := GenS1{F0: 65535, F1: -1}
        _ = fi5
        _ = v6
        _ = v7
    }
    if (false && (18446744073709551615 != 954)) {
        var v8 int32 = 1
        v1 = (-2.25 * v2)
        if (((v1 * float64(1e10)) >= ((v1 / 3.14159) + float64(0.1))) && (true && false)) {
            var v9 int = -164
            v10 := GenS2{F0: 1}
            _ = v10
            _ = v9
        }
        _ = v8
    } else {
        v11 := GenS3{F0: 4294967295}
        _ = v11
    }
    v2 -= (1e10 * v2)
    v12 := v3["\n\t"]
    var v13 int8 = 43
    switch "a" {
    case "foo bar":
        if v12 {
            fmt.Println(v13)
            v14 := GenS3{F0: 65535}
            v15 := GenS3{F0: 883}
            _ = v14
            _ = v15
        } else {
            v12 = (579 >= 4294967295)
            fmt.Println(v12, v13)
            v1 += v1
        }
    default:
        fmt.Println(v12, v12, v12)
    }
    v12 = (-155 <= 9223372036854775807)
    _ = len(v4)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint16 = 683
        _ = fi16
        _ = v17
    }
    v12 = v12
    switch 127 {
    case 4294967295:
        v18 := v3["foo bar"]
        _ = v18
    case -9223372036854775808:
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v13 -= v13
        case "hello":
            var v19 uint32 = 2147483648
            _ = v19
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 = (376 >= 42)
    if (!v12) {
        if (719 == 1) {
            for fi20 := 0; fi20 < 3; fi20++ {
                switch (-(fi20)) {
                case -9223372036854775808:
                    var v21 float64 = v1
                    _ = v21
                case 128:
                    var v22 int = 2147483647
                    _ = v22
                }
                if len(v4) > 0 {
                    v4[0] = ""
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi20
            }
        }
        if true {
            v23 := GenS4{F0: 656, F1: 0, F2: 0}
            switch 4294967295 {
            case 0:
                var v24 int16 = 42
                _ = v24
            case 32767:
                var v25 uint8 = 128
                _ = v25
            default:
                v1 = float64(-2.25)
            }
            _ = v23
        }
    }
    var v26 uint8 = 42
    if v12 {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS5{F0: v13, F1: 255}
            _ = fi27
            _ = v28
        }
        v29 := GenS6{F0: ("a" + ("\n\t" + ("hello" + "\n\t"))), F1: (float64((3.14159 - ((3.14159 + 1.5) / 0.1))) != v1), F2: 255}
        _ = v29
    } else {
        var v30 float64 = v1
        _ = v30
    }
    var v31 uint64 = 347
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
}
