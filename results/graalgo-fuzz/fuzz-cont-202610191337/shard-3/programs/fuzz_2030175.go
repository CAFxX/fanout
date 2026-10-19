package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 int32, p1 int, p2 int64) uint {
    var v6 string = ("\n\t" + ("" + ("a" + "hello")))
    p1 += p1
    var v7 float32 = 0.1
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 42
    v3 := make(map[uint32]bool)
    var v4 float64 = ((0.0 / -2.25) / 1.0)
    var v5 uint = 18446744073709551615
    v2 = v2
    v5 += (v5 << ((uint(v2)) & 63))
    v8 := GenS1{F0: (v5 & (^(v5))), F1: -121}
    v9 := GenS1{F0: (v5 + (v5 >> ((v5) & 63))), F1: 0}
    v10 := []int{-529, 1}
    if ((!(337 != 32768)) && ((((0 > 116) && (v1 != (-(v1)))) && true) || (((false && true) && (((42 != 295) && false) || (0 >= (v5 >> (((^(v5))) & 63))))) || true))) {
        _ = len(v3)
    }
    switch 2147483647 {
    case 32767:
        fmt.Println(v2)
    case 988:
        switch (("foo bar" + "\n\t") + ("" + "foo bar")) {
        case "\n\t":
            var v11 int = -2147483648
            _ = v11
        case "":
            _ = len(v10)
        default:
            _ = len(v10)
        }
    case 42:
        v2 = (v2 ^ 2)
    default:
        fmt.Println(v2, v2)
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
