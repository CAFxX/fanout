package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint = 776
    var v3 int = 9223372036854775807
    var v4 int16 = -539
    if (15 <= 42) {
        v4 += v4
    }
    v5 := GenS1{F0: ("a" < ("foo bar" + (("" + ("" + "foo bar")) + "\n\t"))), F1: (v4 & (v4 >> ((v2) & 63)))}
    v2 -= 18446744073709551615
    if false {
        var v6 uint = (^(v2))
        _ = v6
    } else {
        var v7 float64 = 0.1
        v8 := []int{v3, v3, (v3 | (v3 % ((-9223372036854775808) | 1)))}
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 string = "\n\t"
            var v11 string = v10
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    v12 := GenS1{F0: false, F1: (-(v4))}
    var v13 uint = 255
    var v14 int32 = v1
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v4, v1, v13)
        _ = fi15
    }
    fmt.Println(v14, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
