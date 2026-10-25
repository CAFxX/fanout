package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int }

func main() {
    var v1 int = -2147483648
    var v2 uint32 = 1
    v3 := []string{"foo bar", "foo bar"}
    switch -204 {
    case 42:
        var v4 uint16 = 2
        _ = v4
    default:
        v5 := GenS1{F0: (v1 * -312), F1: 32768, F2: 0}
        _ = v5
    }
    if true {
        if (v2 == (v2 - v2)) {
            switch "hello" {
            case "\n\t":
                if false {
                    if (42 > 65535) {
                        v1 = v1
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        var v6 uint32 = 511
                        _ = v6
                    }
                    v1 = v1
                    var v7 int64 = 42
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 15; fi8++ {
                        if (uint16(v2) < 44) {
                            _ = len(v3)
                        } else {
                            var v9 uint16 = 0
                            _ = v9
                        }
                        fmt.Println(v2)
                        _ = fi8
                    }
                    if (1 < 0) {
                        var v10 uint16 = 1
                        _ = v10
                    } else {
                        v1 -= 127
                    }
                    var v11 uint16 = 65535
                    _ = v11
                }
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            var v12 uint16 = 620
            fmt.Println(v12, v12)
            v13 := GenS2{F0: 721, F1: 389}
            _ = v12
            _ = v13
        }
        fmt.Println(v1, v2)
        var v14 uint32 = (v2 & (v2 + 42))
        _ = v14
    }
    var v15 string = "\n\t"
    var v16 int32 = -2147483648
    fmt.Println(v16)
    switch (v15 + ((v15 + v15) + ("hello" + "a"))) {
    case "foo bar":
        switch v15 {
        case "foo bar":
            v17 := GenS1{F0: (v1 % (((v1 << ((1) & 63))) | 1)), F1: (v1 >> ((1) & 63)), F2: 18446744073709551615}
            _ = v17
        default:
            fmt.Println(v15, v15, v2)
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v16 -= 42
            v2 = 4294967295
            var v20 uint8 = 255
            _ = fi19
            _ = v20
        }
        fmt.Println(v16)
        _ = fi18
    }
    var v21 uint = 1
    var v22 uint = 225
    v22 -= uint(v2)
    if true {
        var v23 int16 = -1
        _ = v23
    }
    v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
    v2 -= 681
    if true {
        switch (-(v1)) {
        case 9223372036854775807:
            if (!(((float32(3.14159) + (1e10 / (1e10 * -2.25))) + 0.0) >= 0.0)) {
                switch (v15 + "") {
                case "a":
                    var v24 int32 = (v16 - 2147483647)
                    _ = v24
                }
                fmt.Println(v22)
            }
        }
        var v25 uint = v22
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS3{F0: (fi26 & v1), F1: 417}
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float32 = ((float32(3.14159) * (0.0 + -2.25)) - -2.25)
            v30 := GenS4{F0: (fi28 + 0)}
            var v31 int8 = int8(v1)
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
        _ = fi26
        _ = v27
    }
    v1 = 4294967295
    var v32 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
}
