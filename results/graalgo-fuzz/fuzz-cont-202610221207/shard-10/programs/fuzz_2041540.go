package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []uint32{101}
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 uint64 = 0
    var v6 int = -402
    delete(v2, 689)
    fmt.Println(v6, v6)
    var v7 float64 = 0.1
    v5 = (v5 - 18446744073709551615)
    if (true && true) {
        v8 := GenS1{F0: 919, F1: 2}
        _ = v8
    }
    var v9 float32 = (3.14159 - (1e10 - 0.0))
    var v10 int64 = -95
    switch ("" + ("" + "")) {
    case "hello":
        if true {
            var v11 int8 = -128
            var v12 uint64 = (v5 - (v5 >> ((371) & 63)))
            switch (v6 - (v6 * (v6 % ((v6) | 1)))) {
            case 494:
                for fi13 := 0; fi13 < 20; fi13++ {
                    v7 = ((v7 * float64(v9)) * 0.1)
                    var v14 int = (^(fi13))
                    _ = fi13
                    _ = v14
                }
            default:
                switch "\n\t" {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = 633
                    }
                }
            }
            _ = v11
            _ = v12
        }
    case "":
        v9 += (-2.25 / v9)
    case "foo bar":
        v9 -= (0.1 - (v9 - ((v9 * v9) * (v9 - ((v9 / (v9 - -2.25)) + (v9 + v9))))))
    }
    var v15 uint8 = 255
    var v16 int16 = 835
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
