package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }

func main() {
    var v1 uint16 = 563
    var v2 int = 127
    var v3 uint = 4294967295
    var v4 bool = false
    v5 := []string{"", "", ("" + ""), ""}
    fmt.Println(v1, v3, v3)
    var v6 uint16 = v1
    for fi7 := 0; fi7 < 2; fi7++ {
        v4 = v4
        if (0.0 > 1e10) {
            if (!(((-2.25 / 1.0) + (1e10 / 3.14159)) != 1e10)) {
                v2 += 380
            }
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float32 = -2.25
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v4 = v4
    for fi10 := 0; fi10 < 12; fi10++ {
        switch 2147483647 {
        case 1:
            var v11 float32 = float32((1e10 + (0.1 * (1e10 - (1.5 / 1e10)))))
            _ = v11
        case 0:
            var v12 string = ""
            _ = v12
        case 408:
            var v13 int16 = 42
            _ = v13
        default:
            v1 = 1
        }
        if v4 {
            switch v2 {
            case 2147483647:
                v6 = (v1 << ((v3) & 63))
            case -259:
                var v14 int16 = 32767
                _ = v14
            }
            var v15 string = ""
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            switch v2 {
            case 122:
                var v17 int32 = -2147483648
                _ = v17
            case 127:
                switch fi16 {
                case 32767:
                    v1 = v6
                case 0:
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        var v18 int32 = 1
                        _ = v18
                    default:
                        fmt.Println(v1)
                    }
                case 108:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    v19 := GenS1{F0: 1, F1: (!true)}
                    _ = v19
                }
            default:
                var v20 float32 = ((1e10 / (float32(1.5) - (0.0 + 3.14159))) - 0.1)
                _ = v20
            }
            switch (fi10 + -9223372036854775808) {
            case 128:
                v1 += (v1 >> (((^(v3))) & 63))
            case 4294967295:
                var v21 uint16 = (v1 >> ((198) & 63))
                _ = v21
            case -9223372036854775808:
                v5 = append(v5, "\n\t")
            }
            if ((!v4) && true) {
                var v22 bool = (65535 < 631)
                fmt.Println(v22)
                _ = v22
            }
            _ = fi16
        }
        _ = fi10
    }
    var v23 uint = v3
    var v24 uint = 242
    var v25 uint16 = (v6 >> (((v3 + 4294967295)) & 63))
    var v26 string = ("" + ("\n\t" + ("hello" + "a")))
    _ = len(v5)
    var v27 int32 = -2147483648
    var v28 float32 = -2.25
    if (24 == 10) {
        v29 := GenS1{F0: v2, F1: (0 < 128)}
        var v30 float64 = -2.25
        _ = v29
        _ = v30
    } else {
        var v31 string = (((v26 + ("foo bar" + ("" + ((v26 + v26) + v26)))) + v26) + (v26 + (v26 + "hello")))
        switch 740 {
        case 400:
            v23 += (v24 + (v23 + 65535))
        case 42:
            var v32 int32 = -2147483648
            _ = v32
        case -534:
            v33 := GenS1{F0: (-(v2)), F1: (0 >= 144)}
            _ = v33
        default:
            switch v26 {
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v34 uint16 = v6
                    v27 += 0
                    _ = v34
                }
            }
        }
        _ = v31
    }
    v23 -= v24
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
