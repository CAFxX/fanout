package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 uint16 = 830
    var v2 uint16 = v1
    var v3 uint64 = 42
    v4 := []int{4294967295, 42}
    var v5 int = -9223372036854775808
    fmt.Println(v3, v5)
    v2 -= uint16(v3)
    if (0 != 2) {
        if false {
            v5 += (v5 << ((65535) & 63))
            var v6 uint = 42
            _ = v6
        } else {
            if (-280 > 14) {
                var v7 int = 2147483647
                if (-867 == -1) {
                    v8 := GenS1{F0: v5, F1: (v3 + v3), F2: "foo bar"}
                    if true {
                        if (!(float64((3.14159 / 0.1)) < (0.1 / 0.1))) {
                            if len(v4) > 0 {
                                v4[0] = 32768
                            }
                            v3 = v3
                        } else {
                            var v9 uint64 = (^(v3))
                            v10 := GenS1{F0: 255, F1: (v3 >> ((65535) & 63)), F2: "\n\t"}
                            v1 -= v1
                            _ = v10
                            _ = v9
                        }
                        var v11 uint32 = 65535
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v11
                    }
                    _ = v8
                }
                _ = v7
            } else {
                v3 += v3
            }
            var v12 uint8 = 2
            _ = v12
        }
        var v13 bool = (!false)
        _ = v13
    }
    var v14 int32 = 0
    v2 = v1
    v2 = v1
    var v15 uint64 = (^(v3))
    v16 := GenS1{F0: (v5 + (v5 | 128)), F1: 715, F2: ""}
    v2 = (^(v1))
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v15, v3)
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: (-(fi18)), F1: (^(v15)), F2: "foo bar"}
            fmt.Println(v3, v15)
            switch ("hello" + ("\n\t" + "a")) {
            case "a":
                fmt.Println(v3)
            case "foo bar":
                v3 -= v3
            case "hello":
                v15 = 4294967295
            default:
                v20 := GenS2{F0: 403, F1: 42}
                _ = v20
            }
            v14 += -1
            switch "foo bar" {
            case "foo bar":
                var v21 int16 = -32768
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
