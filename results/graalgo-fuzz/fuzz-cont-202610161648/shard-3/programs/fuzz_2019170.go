package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64) uint16 {
    var v5 float64 = ((3.14159 - 0.0) * ((0.0 * (3.14159 + (0.0 * 0.1))) / 3.14159))
    v6 := []int{32767, 329}
    fmt.Println(v5)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 65535
    var v3 string = ""
    v4 := make(map[int]uint32)
    for fi7 := 0; fi7 < 16; fi7++ {
        v3 = v3
        _ = fi7
    }
    v8 := GenS1{F0: 9223372036854775807}
    if false {
        var v9 uint32 = 42
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(v1, v2)
            var v11 uint16 = 324
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            switch v3 {
            case "":
                var v13 float32 = 1.5
                _ = v13
            case "a":
                v2 = 4294967295
            default:
                delete(v4, -58)
            }
            _ = fi12
        }
        _ = v9
    }
    fmt.Println(v3)
    var v14 bool = (4294967295 == -386)
    v3 = "a"
    switch v3 {
    case "\n\t":
        for fi15 := 0; fi15 < 6; fi15++ {
            v14 = v14
            v3 = ""
            _ = fi15
        }
    }
    v1 += 264
    var v16 uint = v1
    var v17 uint32 = v2
    v18 := GenS1{F0: -199}
    fmt.Println(v3, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
