package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    var v3 int16 = -497
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    var v6 bool = false
    var v7 bool = true
    v7 = (!(v6 || (float64((float32(((1.5 * 0.1) / 1.0)) * (1.5 - 0.1))) > (float64((((3.14159 / (0.1 - float32(1e10))) * 1e10) - 3.14159)) - (((1e10 * 3.14159) / 3.14159) * 0.0)))))
    v8 := GenS1{F0: 42, F1: 810, F2: 42}
    v7 = (4 < 127)
    if (!(!((v3 >> ((1) & 63)) != v3))) {
        if true {
            var v9 float32 = 1.5
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "" {
                case "foo bar":
                    if (!v6) {
                        switch 2147483647 {
                        case 627:
                            var v11 uint = 255
                            _ = v11
                        case 1:
                            delete(v1, 9223372036854775807)
                        default:
                            if len(v2) > 0 {
                                v2[0] = -2147483648
                            }
                        }
                    }
                case "hello":
                    var v12 int64 = -1
                    _ = v12
                default:
                    var v13 float32 = (v9 / -2.25)
                    _ = v13
                }
                var v14 string = "a"
                _ = fi10
                _ = v14
            }
            _ = v9
        }
    }
    if (v6 || (v6 && v7)) {
        var v15 int = 32768
        var v16 int = -706
        _ = v15
        _ = v16
    }
    fmt.Println(v7, v3, v6)
    v6 = (float64((1.5 * -2.25)) != 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
