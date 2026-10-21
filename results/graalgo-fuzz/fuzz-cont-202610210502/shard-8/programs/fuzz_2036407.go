package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    var v1 int = 127
    var v2 bool = ((true || false) || (57 >= 0))
    v3 := make(map[string]int)
    var v4 int8 = 66
    var v5 int64 = 105
    switch ("foo bar" + "foo bar") {
    case "a":
        v6 := GenS1{F0: (2147483648 < 4294967295)}
        _ = v6
    case "":
        if (v2 && v2) {
            delete(v3, ("\n\t" + ""))
            var v7 float64 = (float64((1.5 * (-2.25 + (1e10 / (3.14159 + (3.14159 + 0.1)))))) + 1e10)
            var v8 float64 = float64(((1e10 / 3.14159) + -2.25))
            _ = v7
            _ = v8
        }
    }
    if (623 < 4294967295) {
        v9 := GenS2{F0: (-(v4)), F1: (v4 & (v4 >> ((42) & 63)))}
        v4 += (v4 ^ -128)
        _ = v9
    }
    v10 := []int{}
    var v11 string = "foo bar"
    v12 := GenS3{F0: 1}
    switch v11 {
    case "foo bar":
        var v13 uint = 65535
        _ = v13
    default:
        v4 = (v4 * int8(v5))
    }
    v2 = (-328 < 0)
    var v14 uint8 = 209
    v15 := GenS4{F0: ("hello" + ("hello" + (((v11 + (("a" + ((v11 + "\n\t") + "foo bar")) + "")) + "foo bar") + v11))), F1: 32767, F2: 0}
    fmt.Println(v5, v1, v2)
    v16 := v3[(v11 + (v11 + (v11 + "")))]
    if len(v10) > 0 {
        v10[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
