package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = 42
    var v2 uint64 = 4294967295
    var v3 int16 = v1
    v4 := make(map[uint32]int)
    var v5 uint8 = 255
    v1 -= (v3 * (v3 ^ (v1 - v1)))
    v4[2147483648] = 4294967295
    delete(v4, 2147483648)
    var v6 int64 = 42
    v7 := GenS1{F0: 0}
    var v8 int64 = 711
    fmt.Println(v3)
    var v9 float64 = 1.5
    var v10 float32 = (0.1 - -2.25)
    v10 -= float32((v9 - v9))
    switch ("" + ("\n\t" + "hello")) {
    case "":
        var v11 uint16 = 42
        _ = v11
    case "foo bar":
        v12 := []int{}
        _ = v12
    case "\n\t":
        var v13 int = -763
        _ = v13
    }
    var v14 uint = 18446744073709551615
    switch "\n\t" {
    case "":
        if ((-2147483648 != 903) && (1 < 1)) {
            fmt.Println(v8, v3, v10)
            switch 32767 {
            case 4294967295:
                if false {
                    var v15 uint = v14
                    var v16 bool = true
                    _ = v15
                    _ = v16
                }
            case 42:
                var v17 uint64 = v2
                _ = v17
            default:
                v5 = (v5 << ((v14) & 63))
            }
            fmt.Println(v8, v9)
        }
    case "\n\t":
        var v18 int8 = 1
        _ = v18
    case "hello":
        var v19 int = 32768
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
