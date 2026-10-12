package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint16 = 255
    var v2 int32 = -335
    var v3 string = ""
    v4 := []int{9223372036854775807, -720, 127}
    var v5 float32 = 3.14159
    var v6 uint16 = v1
    for fi7 := 0; fi7 < 5; fi7++ {
        if (true && (true && ((v1 == (v1 ^ 0)) || false))) {
            fmt.Println(v5, fi7)
        } else {
            if true {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := GenS1{F0: "\n\t", F1: 118}
                    var v10 string = (("foo bar" + v3) + ("\n\t" + ((((v3 + v3) + (v3 + v3)) + v3) + v3)))
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            } else {
                switch v3 {
                case "hello":
                    fmt.Println(v6)
                case "":
                    v3 = "a"
                default:
                    if len(v4) > 0 {
                        v4[0] = fi7
                    }
                }
            }
            var v11 float64 = 3.14159
            v12 := GenS2{F0: -644}
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v4 = append(v4, 782)
    switch v3 {
    case "foo bar":
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint16 = 1
            v15 := GenS1{F0: (v3 + ((v3 + v3) + "\n\t")), F1: -105}
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        var v16 uint16 = v6
        _ = v16
    }
    if (v6 <= (v1 + 255)) {
        var v17 uint16 = 2
        var v18 int64 = 9223372036854775807
        _ = v17
        _ = v18
    }
    if (true && true) {
        switch "a" {
        case "a":
            v19 := GenS3{F0: 1000}
            _ = v19
        }
        switch 128 {
        case 127:
            for fi20 := 0; fi20 < 4; fi20++ {
                v21 := GenS1{F0: "hello", F1: 100}
                _ = fi20
                _ = v21
            }
        }
    }
    v6 -= 65535
    v5 = (v5 * ((1.5 + (-2.25 + v5)) - v5))
    v22 := GenS3{F0: -355}
    v3 = v3
    var v23 string = "\n\t"
    var v24 float32 = ((0.0 / (v5 + v5)) - v5)
    var v25 int8 = 100
    var v26 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
