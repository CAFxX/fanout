package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32, p1 int, p2 int8) string {
    var v6 int32 = p0
    var v7 float64 = 0.1
    v8 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []int{-1, 9223372036854775807, 9223372036854775807}
    var v3 int64 = -364
    v4 := make(map[uint32]string)
    var v5 uint16 = 42
    if len(v1) > 0 {
        v1[0] = 0
    }
    v5 = 32767
    var v9 uint16 = (^(v5))
    v10 := GenS1{F0: 116}
    fmt.Println(v5, v9)
    var v11 uint64 = 18446744073709551615
    v9 += 1
    var v12 uint = 827
    fmt.Println(v5)
    fmt.Println(v5, v9)
    fmt.Println(v3)
    var v13 bool = (v5 >= v5)
    _ = len(v2)
    var v14 uint32 = 2147483648
    if true {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint8 = 127
            var v17 uint16 = (v9 & v5)
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    if v13 {
        switch 1 {
        case 452:
            switch "\n\t" {
            case "":
                v13 = (v13 || (true || (2147483647 != -1)))
            default:
                v1 = append(v1, (^(v14)))
            }
        case -450:
            v18 := GenS1{F0: v11}
            _ = v18
        case 0:
            v13 = v13
        }
    } else {
        v5 = 65535
        var v19 uint16 = (v9 / ((65535) | 1))
        v14 -= v14
        _ = v19
    }
    v3 -= v3
    v13 = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
