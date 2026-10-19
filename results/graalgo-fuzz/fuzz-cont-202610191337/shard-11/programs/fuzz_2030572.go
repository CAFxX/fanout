package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 string = ""
    v4 := make(map[uint32]string)
    v3 = v3
    v3 = v3
    v5 := v4[65535]
    if (0 == 65535) {
        delete(v2, 32767)
        v1[v3] = -166
    } else {
        var v6 uint16 = 0
        var v7 int16 = int16(v6)
        _ = v6
        _ = v7
    }
    v3 = v5
    var v8 float64 = (1e10 - (float64((((0.1 - 0.0) / (((0.0 * 1e10) + -2.25) - (1.5 + (0.1 - 0.1)))) * 0.1)) + float64((-2.25 - (0.1 + (((0.1 * (1e10 * 3.14159)) / 1e10) - 3.14159))))))
    v5 = v3
    v8 = ((v8 * ((v8 / 1.5) + (v8 / 0.1))) * 0.1)
    var v9 float64 = float64(3.14159)
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint64 = 42
            v12 -= 2
            _ = fi11
            _ = v12
        }
        fmt.Println(v9)
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
