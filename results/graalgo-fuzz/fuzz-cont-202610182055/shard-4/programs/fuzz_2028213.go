package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint8 }

func main() {
    v1 := []int{127, -719, 128, 9223372036854775807}
    var v2 bool = (((-2.25 / -2.25) + 0.1) <= (1.5 / 1e10))
    v3 := []int{127, -281, 32767}
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 128
    _ = len(v1)
    v4 = -373
    v6 := GenS1{F0: ((!(((!true) && v2) && true)) || (!v2)), F1: 32767, F2: (v5 - (v5 + 0))}
    var v7 int = -2147483648
    v5 = v5
    var v8 uint16 = 683
    fmt.Println(v2)
    var v9 int8 = 100
    v7 = v7
    var v10 bool = v2
    for fi11 := 0; fi11 < 14; fi11++ {
        _ = len(v1)
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    switch ("foo bar" + ("" + "\n\t")) {
    case "hello":
        var v13 bool = (((v2 || true) || (!false)) || (1e10 != float32(0.0)))
        _ = v13
    case "":
        var v14 int = (v7 % ((v7) | 1))
        _ = v14
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v15 float64 = 0.1
        _ = v15
    case 769:
        var v16 string = ("" + "")
        _ = v16
    case 32768:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
