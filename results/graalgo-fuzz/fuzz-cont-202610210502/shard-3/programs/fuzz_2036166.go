package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int {
    v6 := []int{-2147483648, 42, -42, -2147483648}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 127
        if ((-64 < 127) && (-9223372036854775808 != 1)) {
            v9 := GenS1{F0: 42}
            v8 = (v8 >> ((18446744073709551615) & 63))
            var v10 string = "hello"
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 660
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 2
    var v3 uint32 = 852
    var v4 float32 = 3.14159
    var v5 uint64 = 903
    v4 = (((float32((0.1 - float64(v4))) + 0.1) - float32((((1e10 / (0.1 / 0.1)) / (-2.25 + 1.5)) + 1.5))) * v4)
    v3 = v3
    v11 := GenS1{F0: v5}
    var v12 uint = 2
    v13 := []int{}
    switch -531 {
    case 32768:
        switch "\n\t" {
        case "hello":
            v1[""] = "a"
        case "\n\t":
            v5 -= (v5 / (((v5 | (v5 + v5))) | 1))
        default:
            fmt.Println(v2)
        }
    case 671:
        v14 := GenS1{F0: (^(v5))}
        _ = v14
    default:
        var v15 uint16 = v2
        _ = v15
    }
    var v16 float32 = v4
    v17 := GenS1{F0: (v5 % ((v5) | 1))}
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
