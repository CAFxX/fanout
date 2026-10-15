package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"foo bar", ("" + "foo bar"), "a", "a"}
    var v3 int = -250
    var v4 uint8 = 42
    v5 := []string{"hello"}
    var v6 int = (v3 + -2147483648)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v6, fi7)
        if (false || (-32768 > 32767)) {
            if true {
                fmt.Println(v1, v4, v6)
            }
            var v8 float64 = -2.25
            _ = v8
        } else {
            var v9 int64 = 9223372036854775807
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 444, F1: (v1 >> ((905) & 63)), F2: 2}
    for fi11 := 0; fi11 < 19; fi11++ {
        v3 = -2147483648
        _ = fi11
    }
    v3 += (-(v3))
    var v12 float64 = -2.25
    v1 += (v1 << ((2) & 63))
    v12 = ((((v12 + ((v12 - (v12 + 1e10)) / float64((((-2.25 * 1e10) * 1.5) + -2.25)))) - v12) / (v12 - ((-2.25 * 0.1) + 3.14159))) / v12)
    switch (-(v3)) {
    case 0:
        var v13 int16 = -574
        _ = v13
    case 4294967295:
        var v14 bool = true
        _ = v14
    }
    fmt.Println(v6, v3)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
