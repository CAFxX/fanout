package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 int = -2147483648
    v3 := make(map[int]uint32)
    var v4 uint = 905
    v5 := make(map[string]uint32)
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := v3[(v2 << ((v4) & 63))]
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v2, v1)
    v9 := GenS1{F0: (-(v2)), F1: (v2 ^ v2)}
    var v10 string = ("a" + ("a" + v6))
    var v11 uint = 18446744073709551615
    v12 := GenS2{F0: 127, F1: 42, F2: 1}
    if false {
        v4 = 627
        _ = len(v3)
        v10 = ("foo bar" + v10)
    } else {
        var v13 uint = 881
        if false {
            if false {
                v11 = 515
            } else {
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(v6, v4, v6)
                    switch (v6 + "hello") {
                    case "foo bar":
                        var v15 float64 = (float64(((3.14159 + 0.1) * 1.5)) - 0.1)
                        _ = v15
                    default:
                        v10 = "hello"
                    }
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 bool = ((42 <= 999) && false)
                        _ = fi16
                        _ = v17
                    }
                    _ = fi14
                }
            }
        }
        _ = v13
    }
    var v18 int8 = -70
    fmt.Println(v18, v11)
    fmt.Println(v11, v18, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
