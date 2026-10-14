package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 string = "a"
    var v2 string = "a"
    v3 := []uint32{2147483648, 834}
    var v4 uint16 = 649
    var v5 int8 = 38
    var v6 uint16 = (v4 << ((978) & 63))
    fmt.Println(v4, v5, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if true {
            var v7 int = 560
            _ = v7
        }
        v4 = 55
    }
    v8 := GenS1{F0: -1}
    if (true && true) {
        var v9 uint = 2
        v3 = append(v3, uint32(v9))
        var v10 float32 = (3.14159 / 1e10)
        _ = v10
        _ = v9
    } else {
        v4 -= (v6 - v6)
        switch 2147483647 {
        case -1:
            v3 = append(v3, 2)
        case 0:
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        case 42:
            v5 += (v5 % ((1) | 1))
        }
    }
    _ = len(v3)
    var v11 int64 = 1
    v12 := GenS1{F0: 0}
    v1 = ((v2 + "\n\t") + (v2 + (v1 + v1)))
    switch v1 {
    case "":
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 string = (("hello" + (v2 + "")) + v2)
            _ = fi13
            _ = v14
        }
    case "foo bar":
        var v15 uint16 = (v4 | (v4 ^ v4))
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 11; fi16++ {
            v5 += v5
            _ = fi16
        }
    default:
        var v17 float64 = (((1.5 - float64(3.14159)) - 1.5) - 0.1)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
