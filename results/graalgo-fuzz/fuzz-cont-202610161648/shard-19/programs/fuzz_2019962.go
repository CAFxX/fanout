package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 65535
    var v2 float32 = 0.1
    v3 := make(map[int]bool)
    var v4 int16 = 0
    v5 := []int{128, -447, 68}
    var v6 int8 = -128
    var v7 uint64 = 2
    if true {
        fmt.Println(v7, v4, v6)
        switch -885 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 14; fi8++ {
                delete(v3, fi8)
                var v9 uint32 = uint32(v7)
                _ = fi8
                _ = v9
            }
        case 1:
            var v10 bool = true
            _ = v10
        }
    }
    if (0 == 9223372036854775807) {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int32 = -847
            _ = fi11
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 65535, F1: v1, F2: 59}
        v1 += 4294967295
        _ = v13
    }
    v2 += 0.0
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            _ = len(v3)
            var v16 int = (fi15 & -2147483648)
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            switch ("foo bar" + "foo bar") {
            case "":
                switch 127 {
                case 2147483647:
                    fmt.Println(v7, v1)
                case 128:
                    fmt.Println(v4)
                default:
                    v18 := GenS1{F0: 162, F1: (v1 >> (((v1 & v1)) & 63)), F2: 32767}
                    _ = v18
                }
            case "\n\t":
                if ("foo bar" != "a") {
                    fmt.Println(v6, fi17)
                }
                fmt.Println(v6, v1)
            }
            v7 = 2
            _ = fi17
        }
        var v19 bool = (!((!true) || (!false)))
        _ = fi14
        _ = v19
    }
    switch "\n\t" {
    case "hello":
        if len(v5) > 0 {
            v5[0] = 727
        }
    case "":
        v20 := GenS1{F0: 2, F1: v1, F2: 0}
        _ = v20
    case "foo bar":
        v21 := GenS2{F0: 128}
        _ = v21
    }
    v4 = -1
    switch -2147483648 {
    case 32768:
        if true {
            fmt.Println(v6, v4, v2)
            switch 863 {
            case -194:
                fmt.Println(v1)
            default:
                var v22 bool = (!false)
                _ = v22
            }
            v7 = (v7 * (v7 - v7))
        } else {
            if ((true || false) || (!true)) {
                var v23 uint = v1
                if false {
                    var v24 int = 1
                    _ = v24
                }
                v4 = (v4 | 1000)
                _ = v23
            } else {
                if (!(("foo bar" > "hello") && (((true || (("foo bar" >= "") && false)) && false) || ((v6 * 113) != -30)))) {
                    switch "" {
                    case "hello":
                        v25 := GenS2{F0: 1}
                        _ = v25
                    }
                    var v26 bool = false
                    var v27 bool = v26
                    _ = v26
                    _ = v27
                } else {
                    var v28 int16 = v4
                    _ = v28
                }
                for fi29 := 0; fi29 < 2; fi29++ {
                    if false {
                        fmt.Println(v6, v4)
                    }
                    v6 -= v6
                    _ = fi29
                }
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                v31 := GenS1{F0: 2147483647, F1: (v1 >> ((v1) & 63)), F2: 77}
                switch ("" + "hello") {
                case "a":
                    var v32 string = "foo bar"
                    _ = v32
                case "hello":
                    v6 -= v6
                }
                _ = fi30
                _ = v31
            }
            for fi33 := 0; fi33 < 16; fi33++ {
                switch fi33 {
                case 9223372036854775807:
                    _ = len(v5)
                case 1:
                    v34 := GenS2{F0: 1}
                    _ = v34
                default:
                    switch "a" {
                    case "":
                        if (!false) {
                            delete(v3, 74)
                            var v35 float64 = -2.25
                            _ = v35
                        }
                    }
                }
                v5 = append(v5, (^(fi33)))
                v6 += v6
                _ = fi33
            }
        }
    case 456:
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 bool = false
            var v38 string = ((("\n\t" + ("\n\t" + "")) + ("foo bar" + "hello")) + "\n\t")
            v7 -= (v7 - 363)
            _ = fi36
            _ = v37
            _ = v38
        }
    case -2147483648:
        v39 := GenS2{F0: 224}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
