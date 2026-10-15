package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    var v1 uint = 498
    var v2 uint8 = 81
    var v3 int16 = 1
    v4 := make(map[string]string)
    v5 := []int{-558, -9223372036854775808, -632}
    var v6 uint16 = 843
    var v7 bool = true
    var v8 uint32 = 807
    fmt.Println(v1, v6, v7)
    if (42 > 571) {
        var v9 uint64 = 0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v9
    }
    var v10 int8 = 100
    _ = len(v4)
    switch "a" {
    case "":
        var v11 float64 = 0.0
        _ = v11
    case "a":
        var v12 int8 = v10
        _ = v12
    case "\n\t":
        var v13 uint64 = 1
        _ = v13
    }
    if ((uint64(v2) >= 18446744073709551615) || v7) {
        if ("a" < "\n\t") {
            v7 = false
            if ((v2 + v2) > (v2 & (^(v2)))) {
                v4["hello"] = "hello"
                var v14 int32 = 2147483647
                _ = v14
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    v4[("a" + "\n\t")] = "\n\t"
                    fmt.Println(v1, v10)
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v17 := GenS1{F0: -9223372036854775808, F1: ("hello" + ""), F2: (("\n\t" + "\n\t") + "foo bar")}
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
            }
            var v18 int32 = -2147483648
            _ = v18
        }
        fmt.Println(v10, v10, v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if ((^(v10)) != v10) {
            v6 = (v6 & 42)
            var v19 int8 = v10
            if (((76 != 2147483647) && (!true)) && (v7 && v7)) {
                if (false || ((v6 | 42) <= 368)) {
                    var v20 int8 = v10
                    _ = v20
                }
                var v21 uint64 = 1
                _ = v21
            }
            _ = v19
        } else {
            v22 := GenS2{F0: 0, F1: v10}
            v23 := GenS1{F0: -563, F1: ("\n\t" + ("foo bar" + "a")), F2: "\n\t"}
            _ = v22
            _ = v23
        }
        v8 += 4294967295
        var v24 uint8 = (v2 ^ 2)
        _ = v24
    }
    v25 := GenS3{F0: (-(v3)), F1: 72, F2: uint32(v6)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
