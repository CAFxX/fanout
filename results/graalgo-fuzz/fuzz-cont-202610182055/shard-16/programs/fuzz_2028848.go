package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 bool = (true && true)
    var v3 uint64 = 698
    var v4 uint64 = 661
    v5 := []uint32{}
    v1 = "a"
    v1 = v1
    v5 = append(v5, 2147483648)
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 12; fi7++ {
        if ((v6 << ((2) & 63)) != 32767) {
            if ((1e10 - 0.0) < (1.5 * float64(3.14159))) {
                v2 = v2
                fmt.Println(v1, v2, v6)
                var v8 float64 = 1e10
                _ = v8
            } else {
                switch v1 {
                case "a":
                    var v9 uint32 = 2147483648
                    _ = v9
                case "":
                    v10 := GenS1{F0: 42}
                    _ = v10
                default:
                    fmt.Println(v2, v2, v3)
                }
                v6 -= 65535
                var v11 uint8 = 1
                _ = v11
            }
            fmt.Println(v6, v1, v1)
        } else {
            v1 = v1
        }
        _ = fi7
    }
    _ = len(v5)
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 12; fi13++ {
        if (!true) {
            fmt.Println(fi13)
        }
        _ = fi13
    }
    switch 1 {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 11; fi14++ {
            switch ((v1 + v1) + v1) {
            case "":
                var v15 float32 = (float32(0.0) / -2.25)
                _ = v15
            case "\n\t":
                var v16 string = v1
                _ = v16
            case "hello":
                v4 = (v3 / ((4294967295) | 1))
            default:
                for fi17 := 0; fi17 < 13; fi17++ {
                    if true {
                        v4 = v4
                    } else {
                        var v18 int8 = 100
                        _ = v18
                    }
                    var v19 float32 = (((0.1 - ((1.5 * (0.1 - 3.14159)) / float32(1.5))) + float32(3.14159)) * -2.25)
                    _ = fi17
                    _ = v19
                }
            }
            v3 -= 1
            for fi20 := 0; fi20 < 8; fi20++ {
                fmt.Println(v1)
                v21 := GenS1{F0: 2147483648}
                if (!v2) {
                    if len(v5) > 0 {
                        v5[0] = 2
                    }
                } else {
                    switch (fi14 - (fi20 & -357)) {
                    case -508:
                        fmt.Println(fi14, v4, fi14)
                    case 128:
                        if false {
                            v22 := GenS1{F0: 0}
                            fmt.Println(v3, v3)
                            _ = v22
                        }
                    default:
                        v23 := GenS2{F0: (^(v3)), F1: 49, F2: (!v2)}
                        _ = v23
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = fi14
        }
    }
    switch "foo bar" {
    case "hello":
        if len(v5) > 0 {
            v5[0] = 2
        }
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint64 = 42
        var v26 string = "\n\t"
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(v1)
        if v2 {
            fmt.Println(v1, v3, v6)
            switch ("foo bar" + v1) {
            case "":
                var v28 string = "\n\t"
                _ = v28
            }
        }
        _ = fi27
    }
    var v29 float32 = 1e10
    fmt.Println(v2, v29)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
