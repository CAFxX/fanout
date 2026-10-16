package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int = -2147483648
    var v2 float32 = float32(float64(0.0))
    var v3 string = "\n\t"
    v4 := []uint32{2147483648, 2147483647, 4294967295}
    var v5 uint16 = 32767
    var v6 uint = 255
    var v7 uint = (v6 << ((v6) & 63))
    v3 = ("foo bar" + (((v3 + ((v3 + v3) + "\n\t")) + v3) + "a"))
    fmt.Println(v3)
    if ((false || false) && true) {
        v2 = v2
        if len(v4) > 0 {
            v4[0] = 752
        }
        v8 := GenS1{F0: 4294967295}
        _ = v8
    }
    if (-1 < -2147483648) {
        switch 0 {
        case 115:
            v2 += float32(3.14159)
        case -706:
            v4 = append(v4, 167)
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        switch v1 {
        case 4294967295:
            _ = len(v4)
        }
    } else {
        if (-966 != 496) {
            v3 = v3
        }
    }
    v3 = (v3 + (v3 + "hello"))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v5, v3, v6)
    v4 = append(v4, 2147483647)
    var v9 int8 = 127
    v2 = v2
    var v10 int16 = 0
    if (false && (((989 >= 369) && false) && true)) {
        fmt.Println(v6)
    }
    v3 = "a"
    switch "foo bar" {
    case "\n\t":
        var v11 int = v1
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    fmt.Println(v12, v12, v5)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
