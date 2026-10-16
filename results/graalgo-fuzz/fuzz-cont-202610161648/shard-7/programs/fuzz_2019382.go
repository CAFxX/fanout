package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0() uint8 {
    var v5 int = 32768
    v5 -= (v5 * (v5 + (^(v5))))
    if false {
        var v6 int = (v5 / ((448) | 1))
        _ = v6
    }
    v7 := GenS1{F0: v5, F1: -9223372036854775808}
    v8 := GenS1{F0: (v5 / ((-780) | 1)), F1: 180}
    _ = v5
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 32767
    var v3 int = 4294967295
    v4 := []uint32{2147483648, 2, 2147483648, 693}
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v9 uint8 = 2
    v1 = (v1 + v1)
    var v10 uint = 2
    var v11 bool = (((v9 << ((864) & 63)) > v9) || ((1e10 != 1.5) || true))
    var v12 int32 = -2147483648
    switch ("\n\t" + ((("a" + ("\n\t" + (((("\n\t" + "a") + "foo bar") + "") + "foo bar"))) + "foo bar") + ("hello" + "\n\t"))) {
    case "":
        fmt.Println(v2)
    case "\n\t":
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int = (fi13 >> (((v10 % (((^(v10))) | 1))) & 63))
            if v11 {
                if ((127 == 0) && true) {
                    v3 -= -313
                    switch v3 {
                    case -538:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v15 int = (v14 ^ v14)
                        _ = v15
                    }
                    switch "\n\t" {
                    case "":
                        fmt.Println(v14, v12, v9)
                    }
                } else {
                    var v16 uint = uint(v2)
                    v1 = v1
                    var v17 string = ""
                    _ = v16
                    _ = v17
                }
                var v18 uint8 = v9
                _ = v18
            }
            v19 := GenS1{F0: fi13, F1: v1}
            _ = fi13
            _ = v14
            _ = v19
        }
    }
    switch "" {
    case "":
        fmt.Println(v3, v10)
    case "foo bar":
        var v20 float64 = (1.5 - (((1.5 * 1.5) * 3.14159) * (1.5 / 1.5)))
        _ = v20
    }
    v21 := GenS1{F0: v3, F1: 0}
    var v22 uint8 = (^(v9))
    v23 := GenS1{F0: v3, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
