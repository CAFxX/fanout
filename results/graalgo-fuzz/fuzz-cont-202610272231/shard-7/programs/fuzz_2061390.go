package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 bool = ((0 >= -27) && ((32767 <= 2) && true))
    v2 := []uint32{42, 65535, 170, 333}
    var v3 int32 = 1
    var v4 uint8 = 255
    v5 := []string{(("" + ("hello" + "a")) + "a"), "a"}
    var v6 string = (("a" + "\n\t") + (("\n\t" + "a") + ""))
    _ = len(v5)
    fmt.Println(v6, v4)
    switch 0 {
    case 128:
        for fi7 := 0; fi7 < 5; fi7++ {
            v4 += (v4 >> ((255) & 63))
            v6 = "hello"
            _ = fi7
        }
    default:
        v6 = (v6 + ("hello" + ""))
    }
    fmt.Println(v3, v1, v4)
    v5 = append(v5, (v6 + v6))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 uint16 = 0
    if v1 {
        if v1 {
            v9 := GenS1{F0: 0}
            if v1 {
                switch 1 {
                case -335:
                    var v10 float64 = (float64((0.0 - 0.0)) + 1.5)
                    _ = v10
                }
            }
            var v11 uint64 = 502
            _ = v11
            _ = v9
        } else {
            v2 = append(v2, 2147483648)
            var v12 uint = 18446744073709551615
            v3 = 42
            _ = v12
        }
    } else {
        var v13 bool = v1
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v2)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
