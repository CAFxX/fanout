package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }

func main() {
    var v1 uint32 = 513
    var v2 int16 = 42
    v3 := make(map[uint32]int)
    var v4 int8 = 42
    switch 32767 {
    case 42:
        switch ("a" + "") {
        case "a":
            for fi5 := 0; fi5 < 6; fi5++ {
                if false {
                    fmt.Println(v2, v4)
                    if (1 <= 531) {
                        v1 -= 65535
                    } else {
                        delete(v3, (v1 - v1))
                        fmt.Println(v2, v4, fi5)
                        v2 -= (v2 * v2)
                    }
                } else {
                    v4 = -128
                    var v6 string = "hello"
                    v4 -= v4
                    _ = v6
                }
                _ = fi5
            }
        case "hello":
            fmt.Println(v2, v4)
        }
    case 471:
        v1 += v1
    case 32768:
        var v7 int16 = v2
        _ = v7
    }
    v2 = (^(v2))
    var v8 uint64 = 42
    var v9 int8 = v4
    var v10 int = 42
    if false {
        if false {
            var v11 int32 = -2147483648
            var v12 int64 = -9223372036854775808
            var v13 int8 = 0
            _ = v11
            _ = v12
            _ = v13
        } else {
            switch "" {
            case "hello":
                fmt.Println(v4, v8, v10)
            default:
                switch v10 {
                case 32767:
                    var v14 bool = (("\n\t" + "foo bar") >= "\n\t")
                    _ = v14
                case 1:
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            v15 := GenS1{F0: 169, F1: 720}
                            _ = v15
                        default:
                            var v16 int16 = 32767
                            _ = v16
                        }
                        var v17 uint64 = 2
                        for fi18 := 0; fi18 < 19; fi18++ {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v2 -= 32767
                                if true {
                                    v9 -= (v4 & (^(v9)))
                                } else {
                                    v20 := GenS1{F0: 659, F1: 185}
                                    _ = v20
                                }
                                var v21 int64 = 0
                                var v22 float32 = 0.0
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    var v23 int16 = (v2 & (v2 / ((v2) | 1)))
                                    _ = v23
                                case "\n\t":
                                    fmt.Println(v17)
                                default:
                                    fmt.Println(v9, v9, v1)
                                }
                                _ = fi19
                                _ = v21
                                _ = v22
                            }
                            _ = fi18
                        }
                        _ = v17
                    }
                default:
                    if true {
                        switch (v10 >> ((255) & 63)) {
                        case -641:
                            var v24 float64 = 1.5
                            _ = v24
                        }
                    }
                }
            }
            v25 := v3[(v1 % (((v1 | (v1 ^ v1))) | 1))]
            _ = v25
        }
        v4 += -128
        for fi26 := 0; fi26 < 5; fi26++ {
            switch (v10 & 128) {
            case 518:
                v27 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = v27
            case -222:
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 int32 = 2147483647
                    var v30 float32 = 3.14159
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            if (("\n\t" > "a") || (!(true && ((fi26 & (fi26 & v10)) >= fi26)))) {
                v31 := GenS2{F0: (!(-1 != 1)), F1: v1, F2: ("foo bar" + "\n\t")}
                _ = v31
            }
            _ = fi26
        }
    }
    var v32 int16 = 32767
    fmt.Println(v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}
