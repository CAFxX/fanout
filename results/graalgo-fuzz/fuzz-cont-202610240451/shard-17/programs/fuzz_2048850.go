package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((-2.25 - -2.25) / 1.5)
    var v3 float64 = v2
    var v4 int = 127
    v5 := make(map[int]string)
    var v6 bool = (0 > 62)
    _ = len(v5)
    if (true || (177 > 2)) {
        switch (("" + "foo bar") + "") {
        case "\n\t":
            v2 -= v2
        case "hello":
            var v7 float64 = 3.14159
            _ = v7
        default:
            v8 := []int{v4, (v4 * -469)}
            _ = v8
        }
        v9 := []int{v4, (v4 + (v4 % (((v4 - (v4 & v4))) | 1))), (v4 % ((v4) | 1)), v4}
        _ = v9
    }
    v6 = v6
    var v10 uint = 4294967295
    fmt.Println(v4)
    var v11 uint = 4294967295
    v3 = ((v3 * 1.5) * v2)
    fmt.Println(v2, v2, v3)
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := []int{(fi13 | fi13)}
            var v15 int16 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
        v6 = v6
        var v16 float32 = -2.25
        _ = fi12
        _ = v16
    }
    var v17 uint16 = 32767
    v11 = (^(v11))
    var v18 int64 = 42
    v3 += (3.14159 / 1.0)
    v2 += ((0.1 - float64(float32(((float64(0.0) + float64((0.1 / 3.14159))) - ((v2 * v3) + ((v3 - v3) + v3)))))) + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
