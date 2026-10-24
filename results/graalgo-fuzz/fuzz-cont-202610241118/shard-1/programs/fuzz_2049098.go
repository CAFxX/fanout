package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 float32 = ((1.5 * float32(((((-2.25 / 1.5) * 1e10) - 1e10) * float64(((-2.25 * 0.1) / (((1.5 - 1.5) + 0.1) / (-2.25 - ((1e10 + (-2.25 / 1.5)) * 3.14159)))))))) / 1.0)
    v2 := []uint32{0, 65535, 42}
    v3 := make(map[int]string)
    var v4 int = -1
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    var v7 int = (v4 >> ((439) & 63))
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float32 = ((((1.5 * (float32(1e10) + 1e10)) + v1) + float32((1.5 + 1e10))) - ((((float32((float64(((v1 + v1) * (0.0 * (1.5 * v1)))) * -2.25)) + v1) + v1) * 1.5) / v1))
        v6[""] = v5
        if (!true) {
            if (-149 > 0) {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                switch (v5 + ("" + ("hello" + v5))) {
                case "\n\t":
                    fmt.Println(v5)
                default:
                    var v10 int8 = 0
                    _ = v10
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v11 int32 = -2147483648
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v4 = v7
    v4 += (v7 % (((-(v7))) | 1))
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 -= v1
        fmt.Println(v1)
        _ = fi12
    }
    var v13 uint16 = 255
    v14 := GenS1{F0: 56, F1: 255}
    v15 := GenS1{F0: -1, F1: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
