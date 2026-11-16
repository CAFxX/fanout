package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 int8 }

func main() {
    var v1 int64 = -740
    var v2 int64 = v1
    v3 := []int{}
    var v4 uint32 = 4294967295
    switch 1 {
    case -2147483648:
        v3 = append(v3, 128)
    }
    v5 := GenS1{F0: (^(v4)), F1: 18446744073709551615, F2: (18446744073709551615 < 803)}
    var v6 string = "hello"
    _ = len(v3)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    v2 = -915
    if ((1e10 + -2.25) > (0.0 - 3.14159)) {
        v9 := GenS2{F0: 32767, F1: 4294967295, F2: (^(v2))}
        fmt.Println(v2)
        switch v6 {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            v2 = 1
        }
        _ = v9
    } else {
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                switch (fi10 >> ((65535) & 63)) {
                case -2147483648:
                    var v11 float64 = 0.1
                    _ = v11
                default:
                    var v12 int64 = v2
                    _ = v12
                }
            } else {
                var v13 uint = 18446744073709551615
                _ = v13
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi10
        }
        switch 2147483647 {
        case 1:
            for fi14 := 0; fi14 < 13; fi14++ {
                v6 = "a"
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 int64 = 0
                    _ = fi15
                    _ = v16
                }
                v17 := GenS2{F0: 65535, F1: 18446744073709551615, F2: (-(v1))}
                _ = fi14
                _ = v17
            }
        default:
            v1 = v1
        }
        v3 = append(v3, 658)
    }
    fmt.Println(v4, v6, v1)
    v18 := GenS3{F0: 0, F1: 42, F2: -77}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
