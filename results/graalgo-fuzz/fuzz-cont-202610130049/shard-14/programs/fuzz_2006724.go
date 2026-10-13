package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{}
    var v3 bool = (270 >= 301)
    var v4 float64 = 3.14159
    v5 := make(map[uint32]int)
    v4 = 1.5
    var v6 int16 = 32767
    if (379 <= 42) {
        fmt.Println(v3, v3)
    } else {
        v1 -= 3.14159
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(v6)
        v6 -= 32767
        fmt.Println(v3, v3, v3)
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    if v3 {
        for fi8 := 0; fi8 < 3; fi8++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v4, v1)
            _ = fi8
        }
    } else {
        fmt.Println(v3, v1)
        for fi9 := 0; fi9 < 10; fi9++ {
            fmt.Println(v3, v1)
            _ = fi9
        }
        var v10 uint32 = 384
        _ = v10
    }
    fmt.Println(v1, v6)
    delete(v5, 288)
    v3 = ((127 != -22) && ((!true) || ((!(false && v3)) && (((883 <= (v6 * v6)) || (42 == 9223372036854775807)) || v3))))
    if (42 <= 1) {
        var v11 float32 = v1
        v6 = (v6 | -100)
        switch -893 {
        case -698:
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 string = ""
                v4 += v4
                v14 := GenS1{F0: 0}
                _ = fi12
                _ = v13
                _ = v14
            }
        case -79:
            v6 += 1
        default:
            var v15 uint64 = 2
            _ = v15
        }
        _ = v11
    } else {
        v4 = (float64(1.5) - v4)
        if (v3 || v3) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v1, v1)
            switch "" {
            case "hello":
                var v16 int = -2147483648
                _ = v16
            default:
                var v17 float64 = v4
                _ = v17
            }
        }
        var v18 int64 = -970
        _ = v18
    }
    v3 = (!v3)
    fmt.Println(v3)
    v1 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
