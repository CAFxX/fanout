package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint16 = 635
    var v2 string = "\n\t"
    v3 := []string{v2, (v2 + v2), ("hello" + (v2 + "hello")), v2}
    var v4 uint64 = 34
    var v5 int = 0
    if (true || (!((v4 | (v4 & 18446744073709551615)) >= (v4 ^ 42)))) {
        fmt.Println(v2, v5, v1)
        var v6 int8 = 127
        var v7 int64 = 0
        _ = v6
        _ = v7
    }
    v2 = ""
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint32 = 766
            _ = fi8
            _ = v9
        }
    case 42:
        var v10 uint8 = 255
        _ = v10
    }
    switch v2 {
    case "\n\t":
        switch v5 {
        case -59:
            v4 -= (^(v4))
        case 71:
            fmt.Println(v4, v2, v1)
        case 32767:
            switch "foo bar" {
            case "hello":
                v5 = v5
            case "foo bar":
                if true {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        fmt.Println(v1, fi11)
                        _ = fi11
                    }
                    switch v2 {
                    case "foo bar":
                        if false {
                            v1 = v1
                            v5 = (v5 & v5)
                        } else {
                            _ = len(v3)
                            var v12 float32 = 1e10
                            _ = v12
                        }
                    case "hello":
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := GenS1{F0: false, F1: -2147483648, F2: 42}
                            _ = fi13
                            _ = v14
                        }
                    default:
                        v15 := GenS1{F0: false, F1: 0, F2: 42}
                        _ = v15
                    }
                    var v16 int = 32767
                    _ = v16
                }
            }
        default:
            var v17 int64 = 1
            _ = v17
        }
    case "foo bar":
        var v18 float32 = (float32(0.1) - (-2.25 / -2.25))
        _ = v18
    default:
        for fi19 := 0; fi19 < 7; fi19++ {
            v4 = v4
            _ = fi19
        }
    }
    var v20 uint32 = 42
    switch 140 {
    case 9223372036854775807:
        v3 = append(v3, (v2 + v2))
    }
    var v21 uint16 = v1
    if len(v3) > 0 {
        v3[0] = v2
    }
    if true {
        v22 := GenS2{F0: 791}
        _ = v22
    }
    v3 = append(v3, "foo bar")
    for fi23 := 0; fi23 < 5; fi23++ {
        v21 += (v21 << ((0) & 63))
        var v24 float32 = (((1e10 + 0.1) / 1.0) - 1.5)
        _ = fi23
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v25 := GenS1{F0: true, F1: -28, F2: 2}
    for fi26 := 0; fi26 < 1; fi26++ {
        if ((false && true) && (-2.25 == 3.14159)) {
            var v27 string = v2
            fmt.Println(v20)
            _ = v27
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
