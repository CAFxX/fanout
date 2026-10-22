package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 0
    var v4 bool = (0.0 <= 0.0)
    var v5 int32 = -2147483648
    v6 := make(map[string]bool)
    v1 = 200
    fmt.Println(v4, v5, v3)
    v7 := GenS1{F0: 127, F1: 42, F2: 42}
    var v8 int16 = 504
    var v9 bool = v4
    v2 = v2
    for fi10 := 0; fi10 < 9; fi10++ {
        v8 = (v8 | 1000)
        for fi11 := 0; fi11 < 1; fi11++ {
            v8 += 0
            _ = fi11
        }
        _ = fi10
    }
    _ = len(v6)
    if v9 {
        var v12 int32 = -661
        if (18446744073709551615 > (v3 & v3)) {
            fmt.Println(v2, v4, v9)
            fmt.Println(v4)
        }
        switch 32767 {
        case 0:
            for fi13 := 0; fi13 < 19; fi13++ {
                v9 = v4
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch (("\n\t" + "\n\t") + "") {
                    case "hello":
                        fmt.Println(v5, v5, fi13)
                    case "":
                        v5 = 0
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            var v15 int32 = int32(v3)
                            _ = v15
                        }
                    default:
                        var v16 uint8 = v1
                        _ = v16
                    }
                    _ = fi14
                }
                v6["\n\t"] = (!((!v9) || (!true)))
                _ = fi13
            }
        default:
            fmt.Println(v1)
        }
        _ = v12
    }
    v8 += (^(v8))
    switch 4294967295 {
    case -9223372036854775808:
        if (1.5 > 0.0) {
            var v17 uint = 2
            _ = v17
        }
    case 321:
        var v18 uint64 = v3
        _ = v18
    }
    var v19 int64 = v2
    var v20 float32 = 1.5
    v21 := GenS1{F0: 2147483647, F1: (v19 % (((v2 << ((1) & 63))) | 1)), F2: 65535}
    fmt.Println(v1, v3, v3)
    v22 := GenS2{F0: 2, F1: false}
    var v23 string = ""
    var v24 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
