package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{}
    var v3 int8 = 0
    var v4 string = "a"
    var v5 float64 = 1e10
    var v6 float64 = v5
    fmt.Println(v6, v3, v5)
    for fi7 := 0; fi7 < 12; fi7++ {
        switch 9223372036854775807 {
        case 442:
            v2 = append(v2, (fi7 ^ 2147483647))
        case 128:
            fmt.Println(v5)
        default:
            fmt.Println(v1)
        }
        if len(v2) > 0 {
            v2[0] = 769
        }
        if (!false) {
            fmt.Println(fi7, fi7)
            v4 = "\n\t"
        } else {
            v1 = 0
            v8 := GenS1{F0: -1, F1: -1}
            _ = len(v2)
            _ = v8
        }
        _ = fi7
    }
    var v9 float64 = (v5 + (float64(1e10) - v6))
    v3 = v3
    var v10 bool = false
    _ = len(v2)
    v2 = append(v2, 2147483647)
    switch v4 {
    case "foo bar":
        for fi11 := 0; fi11 < 11; fi11++ {
            v9 = v5
            v2 = append(v2, fi11)
            _ = fi11
        }
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch -2147483648 {
        case 9223372036854775807:
            v12 := GenS1{F0: -949, F1: -1}
            _ = v12
        case 289:
            var v13 int = 9223372036854775807
            _ = v13
        case -1:
            v2 = append(v2, 9223372036854775807)
        default:
            switch 0 {
            case -352:
                if false {
                    _ = len(v2)
                }
            case -2147483648:
                var v14 uint8 = 127
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v6, v5)
        if (18446744073709551615 < v1) {
            v16 := GenS1{F0: 278, F1: int16(v3)}
            var v17 int32 = 469
            _ = v16
            _ = v17
        }
        v10 = v10
        _ = fi15
    }
    v2 = append(v2, -811)
    v3 = v3
    if (-9223372036854775808 <= 364) {
        var v18 int16 = 905
        _ = v18
    }
    v4 = v4
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int64 = -9223372036854775808
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 577
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 int32 = -2147483648
    var v23 int8 = (v3 << ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
