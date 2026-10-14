package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int32 = -766
    var v2 int = 9223372036854775807
    v3 := []string{"", ("hello" + ""), ""}
    for fi4 := 0; fi4 < 1; fi4++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "hello")
        }
        _ = fi4
    }
    if (((!(246 == 0)) && false) || false) {
        if (!(!false)) {
            v3 = append(v3, "hello")
        }
        var v5 bool = (0.1 < -2.25)
        _ = v5
    }
    switch (v2 * v2) {
    case -2147483648:
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v2 = fi7
                _ = fi7
            }
            v1 = v1
            _ = fi6
        }
    }
    var v8 uint16 = 346
    var v9 uint16 = (v8 % ((34) | 1))
    v2 += -207
    v3 = append(v3, "\n\t")
    v10 := GenS1{F0: 0}
    _ = len(v3)
    v11 := GenS1{F0: -385}
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v9)
        _ = fi12
    }
    var v13 int8 = -49
    v9 = v8
    v8 += (v8 << ((2) & 63))
    v14 := GenS1{F0: (^(v1))}
    for fi15 := 0; fi15 < 11; fi15++ {
        if (3.14159 != -2.25) {
            if (2 < 161) {
                var v16 int32 = (v1 + -1)
                v1 -= (v16 ^ 42)
                for fi17 := 0; fi17 < 14; fi17++ {
                    v8 = v8
                    fmt.Println(fi15, fi17, v9)
                    v2 -= 4294967295
                    switch v2 {
                    case -2147483648:
                        var v18 int8 = v13
                        _ = v18
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "":
                            var v19 int = fi15
                            _ = v19
                        case "\n\t":
                            var v20 int64 = 771
                            _ = v20
                        case "foo bar":
                            fmt.Println(v1, v13)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    var v21 int16 = -855
                    _ = fi17
                    _ = v21
                }
                _ = v16
            } else {
                var v22 uint8 = 255
                var v23 int16 = 42
                _ = v22
                _ = v23
            }
            fmt.Println(v9, v2)
            v1 -= -1
        } else {
            var v24 uint32 = 660
            _ = v24
        }
        var v25 int64 = -9223372036854775808
        _ = fi15
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
