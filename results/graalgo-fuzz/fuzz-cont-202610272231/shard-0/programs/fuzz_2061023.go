package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int = 1
    var v2 float64 = 0.1
    var v3 float64 = ((v2 * 1e10) / (((v2 / float64(0.1)) - (v2 + float64(1.5))) + (1e10 - 1.5)))
    var v4 int8 = -128
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(v1)
        _ = fi5
    }
    switch ("\n\t" + ("a" + "foo bar")) {
    case "foo bar":
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 255, F1: 534}
            if (((!(140 == 116)) && (true && true)) && false) {
                v4 = 127
            }
            switch (("\n\t" + "") + "a") {
            case "hello":
                if (v4 > v4) {
                    v8 := []int{(fi6 >> ((42) & 63)), -2147483648}
                    fmt.Println(v2)
                    var v9 int64 = 451
                    _ = v8
                    _ = v9
                } else {
                    var v10 int64 = -616
                    _ = v10
                }
            default:
                var v11 int32 = 71
                _ = v11
            }
            _ = fi6
            _ = v7
        }
    case "\n\t":
        fmt.Println(v4, v1, v4)
    }
    v12 := []int{(v1 ^ -33), -839}
    v13 := GenS1{F0: 2, F1: 32767}
    v4 = (v4 << ((4294967295) & 63))
    var v14 float64 = (((3.14159 * v3) + (0.1 / v2)) - (3.14159 + (v3 + v3)))
    switch v1 {
    case 9223372036854775807:
        v3 = v3
    default:
        var v15 bool = true
        _ = v15
    }
    fmt.Println(v14, v14, v3)
    fmt.Println(v4)
    var v16 int16 = 32767
    if true {
        if len(v12) > 0 {
            v12[0] = 127
        }
        var v17 bool = (!(!true))
        var v18 int32 = 2147483647
        _ = v17
        _ = v18
    } else {
        v3 = (0.0 / (v14 + 0.1))
        var v19 int32 = 2147483647
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
