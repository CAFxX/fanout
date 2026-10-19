package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 string }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 0.0
    var v3 int8 = 42
    var v4 uint = 757
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 63}
        switch fi5 {
        case -1:
            var v7 int = fi5
            _ = v7
        case -9223372036854775808:
            var v8 uint16 = 801
            _ = v8
        case 1:
            if true {
                v9 := []int{fi5, 334, (fi5 - (fi5 % ((fi5) | 1))), (-(fi5))}
                v1 = v1
                _ = v9
            } else {
                v10 := []int{(fi5 | 9223372036854775807)}
                for fi11 := 0; fi11 < 18; fi11++ {
                    v4 = 18446744073709551615
                    _ = fi11
                }
                _ = v10
            }
        default:
            v12 := GenS2{F0: (v4 / ((65535) | 1)), F1: 128, F2: ("\n\t" + ((((("foo bar" + "foo bar") + ("" + "a")) + "hello") + "a") + (("" + "foo bar") + (("" + "foo bar") + "\n\t"))))}
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    if true {
        var v13 int = -9223372036854775808
        _ = v13
    } else {
        fmt.Println(v2, v3)
    }
    var v14 bool = true
    fmt.Println(v3)
    v3 = (v3 >> ((v4) & 63))
    v15 := []int{127, 2147483647}
    v15 = append(v15, -2147483648)
    v3 += (v3 - int8(v4))
    v16 := GenS1{F0: 118}
    if len(v15) > 0 {
        v15[0] = 9223372036854775807
    }
    var v17 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
