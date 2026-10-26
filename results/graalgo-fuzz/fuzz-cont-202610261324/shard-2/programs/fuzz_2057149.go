package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]int)
    v3 := []uint32{(v1 / (((v1 - v1)) | 1)), (v1 / (((v1 - v1)) | 1)), v1, 65535}
    var v4 int16 = 901
    v5 := []string{("\n\t" + "foo bar"), "", "a", "hello"}
    for fi6 := 0; fi6 < 12; fi6++ {
        switch 4294967295 {
        case 42:
            for fi7 := 0; fi7 < 17; fi7++ {
                v4 += (v4 / ((v4) | 1))
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi7)
                case "foo bar":
                    fmt.Println(v4)
                case "\n\t":
                    var v8 uint64 = 18446744073709551615
                    _ = v8
                default:
                    if len(v3) > 0 {
                        v3[0] = (v1 * 0)
                    }
                }
                _ = fi7
            }
        }
        v9 := GenS1{F0: ((v4 | 42) != (^(v4)))}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    fmt.Println(v4, v4)
    fmt.Println(v1, v1)
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 = v4
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi12
    }
    v1 += v1
    var v13 uint32 = 0
    v2[127] = -1
    v1 += v13
    fmt.Println(v13, v13, v4)
    v4 -= v4
    v13 -= v1
    fmt.Println(v1, v1)
    var v14 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
