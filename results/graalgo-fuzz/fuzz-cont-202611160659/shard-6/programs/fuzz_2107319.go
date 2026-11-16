package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 int = 708
    v3 := []string{"a", ((v1 + v1) + ""), "", v1}
    var v4 uint16 = 65535
    var v5 int = (v2 & v2)
    v6 := GenS1{F0: 2, F1: -2147483648}
    var v7 uint = 65535
    switch 4294967295 {
    case 32767:
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                if true {
                    _ = len(v3)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v7, v4)
            }
            v9 := GenS1{F0: 4294967295, F1: -2147483648}
            v10 := GenS1{F0: 2, F1: 42}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint32 = 1
        _ = fi11
        _ = v12
    }
    if true {
        var v13 uint16 = 1
        var v14 int16 = -32768
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v1, v4)
    switch v5 {
    case 1:
        v7 = (v7 % (((^(v7))) | 1))
    default:
        var v15 float64 = ((((1.5 - 0.1) * (-2.25 + 0.1)) * (1e10 / 3.14159)) * (1e10 - 0.0))
        _ = v15
    }
    switch 32767 {
    case -9223372036854775808:
        v16 := GenS1{F0: 0, F1: -2147483648}
        _ = v16
    case 32768:
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint8 = 127
            v3 = append(v3, ("hello" + "foo bar"))
            _ = fi17
            _ = v18
        }
    case -358:
        switch (-(v5)) {
        case -447:
            if true {
                if (-391 <= -32768) {
                    v5 = 128
                    v3 = append(v3, (v1 + v1))
                    if true {
                        var v19 int32 = 1
                        _ = v19
                    }
                } else {
                    switch "\n\t" {
                    case "hello":
                        var v20 int64 = -9223372036854775808
                        _ = v20
                    }
                }
                var v21 uint8 = 128
                _ = v21
            }
        case -577:
            fmt.Println(v7, v2, v5)
        case 32768:
            switch v1 {
            case "hello":
                var v22 float64 = (((-2.25 / float64(1.5)) * 1e10) + 0.1)
                _ = v22
            case "\n\t":
                var v23 int = (v2 & v2)
                _ = v23
            }
        }
    }
    var v24 bool = (42 == 187)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
