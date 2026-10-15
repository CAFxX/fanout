package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    var v3 int32 = -252
    var v4 int = 42
    delete(v2, 32767)
    fmt.Println(v4, v1)
    fmt.Println(v3, v1)
    for fi5 := 0; fi5 < 13; fi5++ {
        if (556 != 65535) {
            fmt.Println(v3)
            var v6 int32 = v3
            for fi7 := 0; fi7 < 12; fi7++ {
                v1 += (^(v1))
                v8 := []int{}
                v3 = int32(v1)
                _ = fi7
                _ = v8
            }
            _ = v6
        }
        switch v4 {
        case 32768:
            fmt.Println(v3, fi5, fi5)
        case -768:
            v3 -= (v3 >> ((2) & 63))
        default:
            v2[v4] = "foo bar"
        }
        _ = fi5
    }
    var v9 float64 = 0.1
    v10 := []int{(v4 << ((297) & 63)), (^(v4))}
    var v11 float64 = v9
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
