package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []uint32{2, 30}
    var v2 string = "foo bar"
    var v3 int = 42
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 100, F1: false, F2: -1}
        if ((v2 + (v2 + (v2 + ("hello" + "a")))) <= "") {
            var v6 bool = true
            var v7 int8 = 1
            for fi8 := 0; fi8 < 12; fi8++ {
                fmt.Println(fi8)
                if ((!(true && v6)) || v6) {
                    switch v2 {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = 654
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v9 := GenS1{F0: v7, F1: (("\n\t" > v2) && (v6 && ((!(2 > 18446744073709551615)) || (v7 > v7)))), F2: 1000}
                    v2 = v2
                    _ = v9
                } else {
                    v1 = append(v1, 65535)
                }
                v1 = append(v1, 2147483648)
                v10 := GenS1{F0: v7, F1: (!v6), F2: 1000}
                var v11 uint32 = 2147483648
                v1 = append(v1, (v11 >> ((346) & 63)))
                _ = len(v1)
                if (0.0 < 1.5) {
                    var v12 int64 = 42
                    switch (^(v3)) {
                    case 2147483647:
                        v6 = false
                    }
                    _ = v12
                }
                fmt.Println(v7, fi4)
                v1 = append(v1, 1)
                v13 := GenS1{F0: -128, F1: v6, F2: 32767}
                _ = fi8
                _ = v10
                _ = v11
                _ = v13
            }
            _ = v6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v14 uint64 = 0
    v1 = append(v1, 0)
    v14 = (^(v14))
    switch (v2 + ("" + v2)) {
    case "":
        if len(v1) > 0 {
            v1[0] = 537
        }
    case "\n\t":
        var v15 uint8 = 127
        _ = v15
    default:
        fmt.Println(v14, v3)
    }
    v2 = ("" + v2)
    fmt.Println(v2, v14, v2)
    fmt.Println(v14, v3, v14)
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(v14, v14)
            _ = fi16
        }
        v17 := GenS2{F0: 189, F1: 708}
        _ = v17
    }
    switch v2 {
    case "\n\t":
        v3 = (v3 | 2147483647)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
