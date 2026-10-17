package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    v1 := []string{"", "foo bar", "hello", ("\n\t" + "\n\t")}
    var v2 int = -1
    var v3 float64 = 3.14159
    fmt.Println(v2, v2)
    v4 := GenS1{F0: 2, F1: -32768, F2: 1}
    v2 += (v2 + (v2 >> ((4294967295) & 63)))
    v1 = append(v1, (("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "a"))
    var v5 bool = true
    v6 := GenS1{F0: 255, F1: -745, F2: 338}
    fmt.Println(v2, v3)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS2{F0: 42, F1: (729 != 993)}
        _ = fi7
        _ = v8
    }
    var v9 int8 = -1
    fmt.Println(v9, v3)
    v9 = (v9 >> ((18446744073709551615) & 63))
    var v10 uint = 42
    if v5 {
        v1 = append(v1, "")
        for fi11 := 0; fi11 < 19; fi11++ {
            switch 4294967295 {
            case 128:
                if v5 {
                    switch (v2 | -2147483648) {
                    case 32767:
                        var v12 uint8 = 127
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 8; fi13++ {
                        var v14 uint = (v10 + 18446744073709551615)
                        var v15 int64 = 9223372036854775807
                        var v16 uint = v10
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
            case 2147483647:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -1:
                fmt.Println(v5)
            default:
                fmt.Println(v10, v5, v2)
            }
            var v17 int8 = (^(v9))
            _ = len(v1)
            _ = fi11
            _ = v17
        }
        var v18 int8 = (v9 / ((v9) | 1))
        _ = v18
    } else {
        v19 := GenS1{F0: 441, F1: 1, F2: 0}
        var v20 uint64 = 531
        v10 = v10
        _ = v19
        _ = v20
    }
    var v21 float64 = (v3 - v3)
    var v22 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
